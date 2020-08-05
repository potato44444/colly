package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

func dataIntercept(text string) string {
	// 数据截取
	index := strings.Index(text, "[")
	lastindex := strings.LastIndex(text, "}catch(e){}")
	return strings.TrimSpace(text[index:lastindex])
}
func writeFile(text string, name string) {
	// 创建目录
	now := time.Now().Format("2006-01-02")
	dir := "COVID-19_" + now
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	// 创建csv文件

	fileName := name + ".csv"
	csvFile, err := os.Create(dir + "/" + fileName)
	if err != nil {
		panic(err)
	}
	// 在结束时关闭该文件
	defer csvFile.Close()
	// 写入UTF-8 BOM，防止中文乱码
	csvFile.WriteString("\xEF\xBB\xBF")
	str := dataIntercept(text)

	// 将字符串进行转换
	jsonBlob := []byte(str)

	if name == "World" {
		var data []GetListByCountryType
		json.Unmarshal(jsonBlob, &data)
		wolrdParse(csvFile, data)
	} else if name == "Area" {
		var data []GetAreaStat
		json.Unmarshal(jsonBlob, &data)
		areaParse(csvFile, data)
	}
}
func collector() *colly.Collector {
	// Instantiate default collector
	c := colly.NewCollector()
	c.OnHTML("script[id]", func(e *colly.HTMLElement) {
		id := e.Attr("id")
		// 全球数据
		if id == "getListByCountryTypeService2true" {
			writeFile(e.Text, "World")
			// 地区数据
		} else if id == "getAreaStat" {
			writeFile(e.Text, "Area")
		}
	})
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	return c
}
func main() {
	c := collector()
	// 访问地址
	c.Visit("https://ncov.dxy.cn/ncovh5/view/pneumonia")
}
