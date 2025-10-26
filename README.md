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

[task4](task4/)目前正在完成所有东西，包括bonus，优化等等。

遇到的几个问题。

第一，处理数据库的dal层，逻辑越写越史，需要大量重构

第二，对于redis运行其实不是很熟悉

第三，对于sql的命令，基本是靠AI写的，自己没有很深入的去了解sql命令

第四，测试问题。目前的测试都是写了一个或者好几个模块，然后一起测试的，也没有使用go语言自带的单元测试，就是直接使用的api fox测试

第五，架构问题。架构的话也是觉得越写越奇怪，需要重构

第六，性能问题，这个就不用多说了，现在是can run

第七，规范问题。这个后续引入.golangci.yml再看看

大概主要就是这些问题了，目前正在继续优化

## task5

[task5](task5/)的技术路线使用hertz+kitex+protobuf+etcd3，我认为这样的技术路线是比较合理的

大概思路就是先看示例demo，然后参考福uu实现吧

感觉大概就是kitex承包了service层，因为正常的hertz得额外创建service层，而这部分内容现在背kitex代替了

## algorithm

[go algorithm](https://github.com/ShaddockNH3/algorithm-study/tree/main/0_go(%E4%BB%A3%E7%A0%81%E9%9A%8F%E6%83%B3%E5%BD%95))
