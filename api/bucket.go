// https://api.upyun.com/doc#/api/operation/domain/GET%20%2Fbuckets%2Fdomains
package api

import (
	"encoding/json"
	"fmt"
)

type BucketParams struct {
	Limit int `json:"limit"`
}

type BucketResponse struct {
	Buckets []Bucket `json:"buckets"`
	Pager   Pager    `json:"pager"`
}

// todo add other fields
type Bucket struct {
	Domains []Domain `json:"domains"`
}

type Domain struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
}

func (u Upyun) url(params BucketParams) string {
	if params.Limit == 0 {
		return fmt.Sprintf("%s/buckets?limit=%d", u.host, 100)
	}
	return fmt.Sprintf("%s/buckets?limit=%d", u.host, params.Limit)
}

// document: https://api.upyun.com/doc#/api/operation/domain/GET%20%2Fbuckets%2Fdomains
func (u Upyun) Buckets(params BucketParams) (*BucketResponse, error) {
	bytes, err := u.Get(u.url(params), nil)
	if err != nil {
		return nil, err
	}
	var response BucketResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}
	if len(response.Buckets) == 0 {
		return nil, fmt.Errorf("%s", string(bytes))
	}
	return &response, nil
}
