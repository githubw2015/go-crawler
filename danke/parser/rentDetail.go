package parser

import (
	"github.com/antchfx/htmlquery"
	"go-project/crawler/engine"
	"go-project/crawler/model"
	"strconv"
	"strings"
)

func ParseRentDetail(contents []byte, url string) engine.ParseResult {

	doc, _ := htmlquery.Parse(strings.NewReader(string(contents)))
	attribute := model.Attribute{}

	//标题
	name := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[1]/h1")
	attribute.Name = name[0].FirstChild.Data

	//价格
	price := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[3]/div[2]/div[2]/div")
	if len(price) > 0 {
		p := price[0].FirstChild.Data
		p = strings.ReplaceAll(p, " ", "")
		p = strings.ReplaceAll(p, "\n", "")
		Price, _ := strconv.Atoi(p)
		attribute.Price = Price
	} else {
		attribute.Price = 0
	}

	//建筑面积
	area := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[1]/label")
	attribute.Area = area[0].FirstChild.Data

	//编号
	number := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[2]/label")
	attribute.Number = number[0].FirstChild.Data

	//户型
	structure := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[3]/label")
	Structure := structure[0].FirstChild.Data
	Structure = strings.ReplaceAll(Structure, " ", "")
	Structure = strings.ReplaceAll(Structure, "\n", "")
	attribute.Structure = Structure

	//付款方式
	pay := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[1]/div[4]/label/a")
	attribute.Pay = pay[0].FirstChild.Data

	//朝向
	orientation := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[1]/label")
	attribute.Orientation = orientation[0].FirstChild.Data

	//楼层
	floor := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[3]/label")
	attribute.Floor = floor[0].FirstChild.Data

	//区域
	region := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[4]/label/div")
	if len(region) > 0 {
		attribute.Region = htmlquery.SelectAttr(region[0], "title")
	} else {
		attribute.Region = ""
	}

	//距离地铁
	metro := htmlquery.Find(doc, "/html/body/div[3]/div[1]/div[2]/div[2]/div[4]/div[2]/div[5]/label")
	attribute.Metro = metro[0].FirstChild.Data

	//链接URL
	attribute.Url = url

	result := engine.ParseResult{
		Requests: nil,
		Items:    []interface{}{attribute},
	}
	return result
}
