# west2-online-golang-2025-test

> 2025.07 - future

## task1

task1是很久之前写的，写的非常乱，而且对于 go routine 相关的理解不足

## task2

task2也已经是蛮久之前写的了，也写得很乱，而且代码用了很多AI，b站爬虫不写了， python 模拟浏览器搞定， go 爬虫主要核心在于并行， py 即便是 httpx 也只是并发。

fzu_go文件我主要的失误在于锁的问题，向表内写入的时候没有进行锁的约束，从而导致报错

## task3

task3目前已经实现所有基本功能，完成三个bonus，详见[task3](task3/)，注意了数据库交互保存的是哈希加密的密码，考虑了service层 ，并且采用了自动生成接口文档

主要缺陷在于未使用redis，接口返回值很不规范。并且还有注入重复编写delete接口等问题

## task4

[task4](task4/)在完成丐版最基本的17个接口基础上，额外实现了MFA认证，以图搜图（文字版），视频流，like redis，并且优化了一部分代码结构（上传文件的逻辑），接下来的优化放在task5实现

## task5

[task5](task5/)的技术路线使用hertz+kitex+protobuf+etcd3，我认为这样的技术路线是比较合理的

大概思路就是先看示例demo，然后参考福uu实现吧

感觉大概就是kitex承包了service层，因为正常的hertz得额外创建service层，而这部分内容现在背kitex代替了

## algorithm

[go algorithm](https://github.com/ShaddockNH3/algorithm-study/tree/main/0_go(%E4%BB%A3%E7%A0%81%E9%9A%8F%E6%83%B3%E5%BD%95))
