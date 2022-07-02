package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ntcat/ezfasthttp/client"
)

func main() {
	var result Longhu
	theDate := "2022-03-03"
	url := fmt.Sprintf("https://proxy.finance.qq.com/cgi/cgi-bin/longhubang/lhbDetail?&date=%s", theDate)
	cli := client.NewFastClient(url)
	var test = "test"
	data, err := cli.GetJsonDo(func(bodyIn client.BodyType) (client.BodyType, error) {
		sBody := string(bodyIn)
		sBody = strings.ReplaceAll(sBody, test, test)
		return []byte(sBody), nil
	}, dataMapPrase)

	if err != nil {
		fmt.Print(err)
		return
	}

	for i, v := range data {
		result = v.(Longhu)
		fmt.Println(i, result.StockName)
	}

}

type Longhu struct {
	ID        uint64  `gorm:"primaryKey;column:id" json:"-"`
	Date      string  `gorm:"column:date" json:"date"`             // 日期
	StockID   string  `gorm:"column:stock_id" json:"stock_id"`     // 代码
	StockName string  `gorm:"column:stock_name" json:"stock_name"` // 名称
	BanCount  int8    `gorm:"column:ban_count" json:"ban_count"`   // 第几板
	NetBuy    float64 `gorm:"column:net_buy" json:"net_buy"`       // 净买
	Chg       float64 `gorm:"column:chg" json:"chg"`               // 涨跌幅
	TotalBuy  float64 `gorm:"column:total_buy" json:"total_buy"`   // 总买入
	TotalAsk  float64 `gorm:"column:total_ask" json:"total_ask"`   // 总卖出
	Turnover  float64 `gorm:"column:turnover" json:"turnover"`     // 换手率
}

func dataMapPrase(dataMap any) (result []any, err error) {
	var lh Longhu
	if dataMap != nil {
		dMap := dataMap.(map[string]any)
		date := dMap["date"].(string)
		lh.Date = date
		if dMap["all"] == nil {
			return nil, fmt.Errorf("[%s]无龙虎榜数据", date)
		}
		all := dMap["all"].([]interface{})
		for i := 0; i < len(all); i++ {
			data := (all[i]).([]interface{})
			lh.StockID = data[0].(string)
			lh.StockName = data[1].(string)
			tmp, _ := strconv.Atoi(data[2].(string))
			lh.BanCount = int8(tmp)
			lh.NetBuy = data[3].(float64)
			lh.Chg = data[4].(float64)
			lh.TotalBuy = data[5].(float64)
			lh.TotalAsk = data[6].(float64)
			lh.Turnover = data[7].(float64)
			result = append(result, lh)
		}
		return
	}
	return nil, fmt.Errorf("dataMap is nill")
}
