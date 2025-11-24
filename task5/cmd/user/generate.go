//go:build ignore

package main

import (
	"github.com/ShaddockNH3/west2-online-golang-2025-test/task5/app/user/infrastructure/mysql"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "app/user/infrastructure/mysql/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.ApplyBasic(
		mysql.User{},
		mysql.Image{},
	)

	g.Execute()
}
