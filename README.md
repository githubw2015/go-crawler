# crawler
golang并发爬虫，以蛋壳网为例<br/>
大体思路讲解:<br/>
1.danke：目标网站。<br/>
fetcher：处理url，得到返回结果。<br/>
parser： 解析fetcher返回的结果<br/>
2.scheduler：任务调度器。<br/>
3.engine：实现Scheduler接口，建立工作引擎，中间件。<br/>
4.model：定义存储数据类型<br/>
5.persist：对数据做持久化处理（本例实现的是存储本地mysql）<br/>
6.main.go：项目入口<br/>

<br/>
整体：
将request传递给engine <br/>
engine将request交给scheduler调度<br/>
scheduler将request放入channel，然后由worker进行解析，将解析结果再放入channel,scheduler从channel中获取<br/>
parseResult继续调度任务<br/>

##engine：
根据用户配置WorkerCount，开启多个goroutine实现并发 <br/>
创建了供所有worker使用的channel<br/>
循环消费channel中的request，有就交给scheduler进行调度<br/>

##scheduler：
持续将request放入channel，安排worker进行工作

<br/>
##如何使用:
前提这两个包需要下载下来：<br/>
git clone https://github.com/golang/text.git   <br/>
git clone https://github.com/golang/net.git 

##如何下载：可以参考一下文章
https://blog.csdn.net/weixin_37677804/article/details/82877984

##数据库结构 参见rent.sql 文件
<br/>

##抓取数据的效果图:
![实际效果图](./82951.png)