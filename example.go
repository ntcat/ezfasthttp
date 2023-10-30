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
	cxt := client.NewFastClient(url)
	cxt.SetHeadReq("Accept-Language", "zh-CN,zh;q=0.9")
	cxt.SetHeadReq("Upgrade-Insecure-Requests", "1")
	cxt.SetHeadReq("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	cxt.SetHeadResp("Cache-Control", "no-cache")
	cxt.SetHeadResp("Connection", "keep-alive")
	cxt.SetHeadResp("Pragma", "no-cache")
	cxt.SetHeadReq("Cookie", "JSESSIONID=6D8C8BC8BE99F662DDA3CB07C67F0BAC; insert_cookie=37836164; routeId=.uc2; _sp_ses.2141=*; _sp_id.2141=545d02af-1b78-4e19-948b-c469ae31ae52.1696392078.5.1698682342.1698083751.b2ee5ad7-17eb-4641-9d21-251234ccea8f")
	cxt.SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")

	var test = "test"

	data, err := cxt.GetJsonDo(func(bodyIn client.BodyType) (client.BodyType, error) {
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
		dMap := dataMap.(client.DataMapType)
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
			lh.TotalBuy = data[5].(float64)
			lh.TotalAsk = data[6].(float64)
			lh.Turnover = data[7].(float64)
			result = append(result, lh)
		}
		return
	}
	return nil, fmt.Errorf("dataMap is nill")
}
