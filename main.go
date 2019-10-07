package main

import (
	"go-project/crawler/danke/parser"
	"go-project/crawler/engine"
	"go-project/crawler/persist"
	"go-project/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{}, //调度器
		WorkerCount: 10,                           // 线程数
		ItemChan:    persist.ItemSaverSql(),       //管道
	}

	//解析数据
	e.Run(engine.Request{
		Url:        "https://www.danke.com/room/sh/d.html",
		ParserFunc: parser.ParseAreaList,
	})
}
