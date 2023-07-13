package api

import (
	"fmt"
	"os"
	"testing"
)

func TestUpyun_Buckets(t *testing.T) {
	upyun := NewUpyun(UpyunConfig{
		Authorization: os.Getenv("api_upyun_bucket"),
	})
	domains, err := upyun.BucketDomains()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(domains)
	fmt.Println(len(domains))
}
