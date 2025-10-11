# west2-online-golang-2025-test

## task1

task1在其他仓库，很久之前写的，写的非常乱，未来有机会再移过来~~（搬史）~~

## task2

task2也已经是蛮久之前写的了，也写得很乱，而且代码用了很多AI，b站爬虫不写了，python模拟浏览器搞定，go爬虫主要核心在于并行，py即便是httpx也只是并发。

fzu_go文件我主要的失误在于锁的问题，向表内写入的时候没有进行锁的约束，从而导致报错

## task3

task3目前已经实现所有基本功能，完成三个bonus，详见[task3](https://github.com/ShaddockNH3/west2-online-golang-2025-test/tree/main/task3)，注意了数据库交互保存的是哈希加密的密码，考虑了service层 ，并且采用了自动生成接口文档

## task4

这是一条测试commit签名的测试