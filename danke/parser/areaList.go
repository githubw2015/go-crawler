package parser

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"go-project/crawler/engine"
	"strconv"
	"strings"
)

/*
获得某个大的行政区的子区域列表信息
比如普陀区
就获得普陀区下的李子园。。。等子区域的url列表信息
*/
var areaList = [...]string{
	"徐汇区",
	"闵行区",
	"静安区",
	"松江区",
	"卢湾区",
	"黄浦区",
	"闸北区",
	"宝山区",
	"浦东新区",
	"长宁区",
	"青浦区",
	"虹口区",
	"杨浦区",
	"普陀区",
	"嘉定区",
}

func ParseAreaList(contents []byte) engine.ParseResult {
	doc, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	result := engine.ParseResult{}

	for index, _ := range areaList {

		zone := index + 1
		///徐汇区对应于html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[1]/div/a
		//普陀区对应于/html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[2]/div/a
		//嘉定区对应于/html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[3]/div/a
		// 依次递进  普陀区div[4]，徐汇区div[5]，嘉定区div[15]
		patern := "/html/body/div[3]/div/div[4]/div[2]/dl[2]/dd/div/div[" + strconv.Itoa(zone) + "]/div/a"
		for _, v := range htmlquery.Find(doc, patern) {
			//获取所有的地区下列表
			result.Items = append(result.Items, zone)
			result.Requests = append(result.Requests, engine.Request{
				Url:        htmlquery.SelectAttr(v, "href"),
				ParserFunc: ParseArea,
			})
		}
	}
	fmt.Println("抓取的数据URL总数:", len(result.Items))
	return result
}
