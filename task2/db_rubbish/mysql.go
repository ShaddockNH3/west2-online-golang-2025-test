package db_rubbish

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

// connectDB 负责连接到数据库
func connectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("数据库连接配置无效: %w", err)
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("无法连接到数据库: %w", err)
	}
	return db, nil
}

// setupDepartments 负责创建并填充科室表
func setupDepartments(db *sql.DB, departments []string) error {
	createSQL := "CREATE TABLE IF NOT EXISTS `departments` (`department_id` INT PRIMARY KEY, `department_name` VARCHAR(50) UNIQUE) CHARACTER SET utf8mb4;"
	if _, err := db.Exec(createSQL); err != nil {
		return fmt.Errorf("创建departments表失败: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开启事务失败: %w", err)
	}
	stmt, err := tx.Prepare("INSERT IGNORE INTO `departments` (`department_id`, `department_name`) VALUES (?, ?)")
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("准备科室插入语句失败: %w", err)
	}
	defer stmt.Close()

	for id, name := range departments {
		if _, err := stmt.Exec(id, name); err != nil {
			tx.Rollback()
			return fmt.Errorf("插入科室 '%s' 失败: %w", name, err)
		}
	}
	return tx.Commit()
}

// loadDepartmentsToMap 从数据库加载科室信息到map
func loadDepartmentsToMap(db *sql.DB) (map[string]int, error) {
	departmentMap := make(map[string]int)
	rows, err := db.Query("SELECT department_id, department_name FROM `departments`")
	if err != nil {
		return nil, fmt.Errorf("查询科室信息失败: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			return nil, fmt.Errorf("扫描科室信息失败: %w", err)
		}
		departmentMap[name] = id
	}
	return departmentMap, rows.Err()
}

// processExcelData 负责创建主表并插入处理后的Excel数据
func processExcelData(db *sql.DB, excelRows [][]string, departmentMap map[string]int) error {
	if len(excelRows) < 2 {
		return fmt.Errorf("Excel数据不足")
	}

	headers := excelRows[0]
	authorColumnIndex := -1
	var mainTableColumns []string
	for i, header := range headers {
		if strings.EqualFold(header, "author") {
			authorColumnIndex = i
			mainTableColumns = append(mainTableColumns, "`department_id` INT")
		} else {
			mainTableColumns = append(mainTableColumns, fmt.Sprintf("`%s` VARCHAR(255)", header))
		}
	}
	if authorColumnIndex == -1 {
		return fmt.Errorf("在Excel表头中未找到 'Author' 列")
	}

	mainTableName := "main_data"
	createSQL := fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` (%s) CHARACTER SET utf8mb4;", mainTableName, strings.Join(mainTableColumns, ", "))
	if _, err := db.Exec(createSQL); err != nil {
		return fmt.Errorf("创建主数据表失败: %w", err)
	}
	fmt.Printf("主数据表 '%s' 已准备就绪\n", mainTableName)

	var insertHeaders []string
	for _, h := range headers {
		if !strings.EqualFold(h, "author") {
			insertHeaders = append(insertHeaders, fmt.Sprintf("`%s`", h))
		}
	}
	insertHeaders = append(insertHeaders, "`department_id`")
	insertSQL := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (?%s)", mainTableName, strings.Join(insertHeaders, ","), strings.Repeat(",?", len(insertHeaders)-1))

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("无法开启事务: %w", err)
	}
	stmt, err := tx.Prepare(insertSQL)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("无法准备主数据插入语句: %w", err)
	}
	defer stmt.Close()

	fmt.Println("开始转换并插入主数据...")
	for i, row := range excelRows[1:] {
		authorName := row[authorColumnIndex]
		departmentID, ok := departmentMap[authorName]
		if !ok {
			log.Printf("警告: 在第 %d 行找到未知科室 '%s'，将跳过此行", i+2, authorName)
			continue
		}

		var valuesToInsert []interface{}
		for j, cell := range row {
			if j != authorColumnIndex {
				valuesToInsert = append(valuesToInsert, cell)
			}
		}
		valuesToInsert = append(valuesToInsert, departmentID)
		if _, err := stmt.Exec(valuesToInsert...); err != nil {
			tx.Rollback()
			return fmt.Errorf("插入第 %d 行数据失败: %w", i+2, err)
		}
	}
	return tx.Commit()
}

// MySQL 是整个数据迁移任务的总指挥官
func MySQL(mysqlDSN, excelFileName, sheetName string, departments []string) error {
	fmt.Println("开始数据处理任务...")

	db, err := connectDB(mysqlDSN)
	if err != nil {
		return err
	}
	defer db.Close()
	fmt.Println("数据库连接成功")

	if err := setupDepartments(db, departments); err != nil {
		return err
	}
	fmt.Println("科室数据准备就绪")

	departmentMap, err := loadDepartmentsToMap(db)
	if err != nil {
		return err
	}
	fmt.Println("科室信息已加载到内存")

	excelFile, err := excelize.OpenFile(excelFileName)
	if err != nil {
		return fmt.Errorf("打开Excel文件失败: %w", err)
	}
	defer excelFile.Close()
	excelRows, err := excelFile.GetRows(sheetName)
	if err != nil {
		return fmt.Errorf("读取工作表失败: %w", err)
	}
	fmt.Printf("Excel文件读取成功，共 %d 行\n", len(excelRows))

	if err := processExcelData(db, excelRows, departmentMap); err != nil {
		return err
	}

	fmt.Println("所有任务已成功完成")
	return nil
}

// main 是程序的唯一入口
func main() {
	// --- 在这里配置你的信息 ---
	mysqlDSN := "rooting:mmmmmmmmmmmmmmmmmmmmmm@tcp(127.0.0.1:3306)/fzutest?charset=utf8mb4"
	excelFileName := "working1.xlsx"
	sheetName := "Sheet1"
	departments := []string{
		"综合科", "教学运行", "教研教改", "计划科", "实践科",
		"质量办", "电教中心", "教材中心", "铜盘校区管理科",
	}
	// --------------------------

	// 只需调用这一个函数，即可启动整个流程
	if err := MySQL(mysqlDSN, excelFileName, sheetName, departments); err != nil {
		log.Fatalf("任务失败: %v", err)
	}
}
