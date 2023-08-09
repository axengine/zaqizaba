package alioss

import (
	"encoding/json"
	"testing"
)

const (
	endpoint        = "oss-cn-hangzhou.aliyuncs.com"
	accessKeyId     = "-------"
	accessKeySecret = "------------"
	my_bucket       = "------------"
)

func TestExample(t *testing.T) {
	cli := New(endpoint, accessKeyId, accessKeySecret)

	var data = map[string]interface{}{
		"A": "a",
	}
	bz, _ := json.Marshal(&data)

	objKey := "uripath/0"
	err := cli.UploadFile(my_bucket, objKey, bz, "application/json")
	if err != nil {
		panic(err)
	}
	// https://domain/uripath/0
}
