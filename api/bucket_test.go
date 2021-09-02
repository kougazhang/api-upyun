package api

import (
	"fmt"
	"os"
	"purge/assert"
	"testing"
)

func TestUpyun_Buckets(t *testing.T) {
	upyun := NewUpyun(UpyunConfig{
		Authorization: os.Getenv("api-upyun-test-auth"),
	})
	resp, err := upyun.Buckets(BucketParams{
		Limit: 10000,
	})
	assert.Equal(t, nil, err)
	fmt.Println(resp)
}
