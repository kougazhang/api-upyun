package api

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestUpyun_Statistic(t *testing.T) {
	upyun := NewUpyun(UpyunConfig{
		Authorization: os.Getenv("api_upyun_statistic"),
	})
	end := time.Now()
	start := end.Add(-time.Hour)
	res, err := upyun.Statistic(StatisticParam{
		Domain:      "sns-video-hw.xhscdn.com",
		Start:       start,
		End:         end,
		AccountName: "huaweicloud",
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, item := range res {
		fmt.Printf("%+v\n", item)
	}
}

// 1689838200, 2023-07-20 15:30:00
