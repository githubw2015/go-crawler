package persist

import (
	"go-project/crawler/model"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		zoneCount := 0
		houseCount := 0
		for {
			item := <-out
			switch t := item.(type) {
			case int:
				zoneCount += 1
				switch t {
				case 1:
					log.Printf("获取徐汇区子区域url：%d个", zoneCount)
				case 2:
					log.Printf("获取闵行区子区域url：%d个", zoneCount)
				case 3:
					log.Printf("获取静安区子区域url：%d个", zoneCount)
				case 4:
					log.Printf("获取松江区子区域url：%d个", zoneCount)
				case 5:
					log.Printf("获取卢湾区子区域url：%d个", zoneCount)
				case 6:
					log.Printf("获取黄埔区子区域url：%d个", zoneCount)
				case 7:
					log.Printf("获取闸北区子区域url：%d个", zoneCount)
				case 8:
					log.Printf("获取宝山区子区域url：%d个", zoneCount)
				case 9:
					log.Printf("获取浦东新区子区域url：%d个", zoneCount)
				case 10:
					log.Printf("获取长宁区子区域url：%d个", zoneCount)
				case 11:
					log.Printf("获取青浦区子区域url：%d个", zoneCount)
				case 12:
					log.Printf("获取虹口区子区域url：%d个", zoneCount)
				case 13:
					log.Printf("获取杨浦区子区域url：%d个", zoneCount)
				case 14:
					log.Printf("获取普陀区子区域url：%d个", zoneCount)
				case 15:
					log.Printf("获取嘉定区子区域url：%d个", zoneCount)
				}
			case string:
				houseCount += 1
				log.Printf("获取子区域%v房源url", t)
				log.Printf("获取房屋url：%d个", houseCount)
			case model.Attribute:
				itemCount++
				log.Printf("解析到第%d个房源信息: #%v", itemCount, item)
			}
		}
	}()
	return out
}
