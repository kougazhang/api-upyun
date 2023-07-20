// Package api https://api.upyun.com/doc#/api/operation/bucket/GET%20%2Fbuckets
package api

import (
	"encoding/json"
	"fmt"
	"strings"
)

// BucketDomains https://api.upyun.com/doc#/api/operation/bucket/GET%20%2Fbuckets
func (u Upyun) BucketDomains() (res []DomainInfo, err error) {
	url := fmt.Sprintf("%s/buckets?limit=100", u.host)
	for {
		resp, _err := u.bucketDomains(url)
		if _err != nil {
			return nil, _err
		}
		for _, bucket := range resp.Buckets {
			for _, dm := range bucket.Domains {
				if strings.HasSuffix(dm.Domain, "test.upcdn.net") {
					continue
				}
				res = append(res, dm)
			}
		}
		if resp.Pager.Max == 0 {
			return
		}
		url = fmt.Sprintf("%s/buckets?limit=100&max=%d", u.host, resp.Pager.Max)
	}
}

type respBucketDomains struct {
	Pager   Pager        `json:"pager"`
	Buckets []BucketInfo `json:"buckets"`
}

type BucketInfo struct {
	Domains []DomainInfo `json:"domains"`
}

type DomainInfo struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
}

func (u Upyun) bucketDomains(url string) (*respBucketDomains, error) {
	bytes, err := u.Get(url, nil)
	if err != nil {
		return nil, err
	}
	var response respBucketDomains
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
