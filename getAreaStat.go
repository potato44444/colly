package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type GetAreaStat struct {
	ProvinceName          string   // 省名
	ProvinceShortName     string   //	省的简称
	CurrentConfirmedCount int      //	现存确诊
	ConfirmedCount        int      // 累计确诊
	SuspectedCount        int      //	可疑人数
	CuredCount            int      //治愈人数
	DeadCount             int      //死亡人数
	Comment               int      //评论
	LocationId            int      // 位置编号
	StatisticsData        string   // 统计数据
	Cities                []Cities //城市
}
type Cities struct {
	CityName              string // 城市
	CurrentConfirmedCount int    // 现存确诊
	ConfirmedCount        int    // 累计确诊
	SuspectedCount        int    // 可疑人数
	CuredCount            int    // 治愈人数
	DeadCount             int    // 死亡人数
	LocationId            int    // 位置编号
}

// 地区数据解析
func areaParse(csvFile *os.File, data []GetAreaStat) {
	// 写入数据
	writer := csv.NewWriter(csvFile)
	// 写入标题
	title := []string{
		"名称",
		"现存确诊",
		"累计确诊",
		"可疑人数",
		"治愈人数",
		"死亡人数",
		"编号",
	}

	writer.Write(title)
	for _, post := range data {
		line := []string{
			post.ProvinceName,
			strconv.Itoa(post.CurrentConfirmedCount),
			strconv.Itoa(post.ConfirmedCount),
			strconv.Itoa(post.SuspectedCount),
			strconv.Itoa(post.CuredCount),
			strconv.Itoa(post.DeadCount),
			strconv.Itoa(post.LocationId),
		}
		err := writer.Write(line)
		for _, city := range post.Cities {
			cityLine := []string{
				city.CityName,
				strconv.Itoa(city.CurrentConfirmedCount),
				strconv.Itoa(city.ConfirmedCount),
				strconv.Itoa(city.SuspectedCount),
				strconv.Itoa(city.CuredCount),
				strconv.Itoa(city.DeadCount),
				strconv.Itoa(city.LocationId),
			}
			writer.Write(cityLine)
		}
		writer.Write([]string{})

		if err != nil {
			panic(err)
		}
		if err != nil {
			panic(err)
		}
	}
}
