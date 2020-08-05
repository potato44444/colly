package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type IncrVo struct {
	CurrentConfirmedIncr int
	ConfirmedIncr        int
	CuredIncr            int
	DeadIncr             int
	ShowRank             bool
}

type GetListByCountryType struct {
	Id               int
	ProvinceName     string // 国家
	CountryShortCode string // 国家短代码
	CountryFullName  string // 国家全名
	LocationID       int    // 位置编码
	CountryType      int    // 国家类型
	Continents       string // 所属大洲

	CurrentConfirmedCount int // 现存确诊
	ConfirmedCount        int // 累计确诊
	DeadCount             int // 死亡人数
	CuredCount            int // 治愈人数

	CreateTime int // 创建时间
	ModifyTime int // 修改时间

	DeadRate           string // 死亡率
	DeadCountRank      int    // 死亡人数排名
	DeadRateRank       int    // 病死率排序
	ConfirmedCountRank int    // 确认计数等级
	SuspectedCount     int    // 可疑人数
	StatisticsData     string // 统计数据json

	IncrVo IncrVo

	Operator          string
	Comment           string
	Sort              int
	Tags              string
	ProvinceId        int
	CityName          string
	ProvinceShortName string
}

// world数据解析
func wolrdParse(csvFile *os.File, data []GetListByCountryType) {
	// 写入数据
	writer := csv.NewWriter(csvFile)

	// 写入标题
	title := []string{
		"ID",
		"国家名称",
		"所属洲",
		"现存确诊",
		"可疑人数",
		"累计确诊",
		"死亡人数",
		"治愈人数",
		"确认计数等级",
		"病死率排序",
		"死亡率",
		"历史记录",
	}
	writer.Write(title)
	// 写入相应的数据
	for _, post := range data {
		line := []string{
			strconv.Itoa(post.Id),
			post.ProvinceName,
			post.Continents,
			strconv.Itoa(post.CurrentConfirmedCount),
			strconv.Itoa(post.SuspectedCount),
			strconv.Itoa(post.ConfirmedCount),
			strconv.Itoa(post.DeadCount),
			strconv.Itoa(post.CuredCount),
			strconv.Itoa(post.ConfirmedCountRank),
			strconv.Itoa(post.DeadRateRank),
			post.DeadRate,
			post.StatisticsData,
		}
		writer.Write(line)
	}
	writer.Flush()
}
