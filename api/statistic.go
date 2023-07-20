package api

import (
	"encoding/json"
	"net/url"
	"time"
)

type StatisticParam struct {
	Domain      string
	AccountName string
	Start, End  time.Time
}

func (u Upyun) Statistic(param StatisticParam) ([]StatisticItem, error) {
	req, err := newStatisticReq(u, param.AccountName)
	if err != nil {
		return nil, err
	}
	// request
	_url, err := req.generateUrl(param.Domain, param.Start, param.End)
	if err != nil {
		return nil, err
	}
	data, err := u.Get(_url, nil)
	if err != nil {
		return nil, err
	}
	// response
	var resp statisticResp
	err = json.Unmarshal(data, &resp)
	if err != nil {
		return nil, err
	}
	return convertStatisticRespToResult(resp)
}

func convertStatisticRespToResult(resp statisticResp) ([]StatisticItem, error) {
	// improve the resp.Result to be more readable
	result := make([]StatisticItem, 0, len(resp.Result))
	for _, item := range resp.Result {
		ins := StatisticItem{
			// convert the end time as the start time
			// fix the strange setting using the end time as the time flag
			Time:          time.Unix(item.Time, 0).Add(-5 * time.Minute),
			BillBytes:     item.BillBytes,
			BillBandwidth: item.BillBandwidth,
			LogBytes:      item.LogBytes,
			LogBandwidth:  item.LogBandwidth,
			RealBytes:     item.RealBytes,
			RealBandwidth: item.RealBandwidth,
		}
		result = append(result, ins)
	}
	return result, nil
}

type statisticResp struct {
	Result []StatisticRespItem `json:"result"`
}

type StatisticRespItem struct {
	Time          int64   `json:"time"`
	BillBytes     int64   `json:"bill_bytes"`
	BillBandwidth float64 `json:"bill_bandwidth"`
	LogBytes      int64   `json:"log_bytes"`
	LogBandwidth  float64 `json:"log_bandwidth"`
	RealBytes     int64   `json:"real_bytes"`
	RealBandwidth float64 `json:"real_bandwidth"`
}

type StatisticItem struct {
	Time          time.Time
	BillBytes     int64
	BillBandwidth float64
	LogBytes      int64
	LogBandwidth  float64
	RealBytes     int64
	RealBandwidth float64
}

type statisticReq struct {
	baseURL     string
	accountName string
	firm        string
}

func newStatisticReq(ins Upyun, accountName string) (sta statisticReq, err error) {
	if len(sta.firm) == 0 {
		sta.firm = "all"
	}
	sta.accountName = accountName
	sta.baseURL = ins.host
	return
}

func (s statisticReq) generateUrl(domain string, start, end time.Time) (string, error) {
	ins, err := url.Parse(s.baseURL)
	if err != nil {
		return "", err
	}
	ins.Path = "gifshow/log/statistic"
	query := ins.Query()
	query.Add("start_time", start.Add(-8*time.Hour).Format("2006-01-02T15:04:05Z"))
	query.Add("end_time", end.Add(-8*time.Hour).Format("2006-01-02T15:04:05Z"))
	query.Add("account_name", s.accountName)
	query.Add("domain", domain)
	query.Add("firm", s.firm)
	ins.RawQuery = query.Encode()
	return ins.String(), nil
}
