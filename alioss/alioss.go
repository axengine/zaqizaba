package alioss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	cli *oss.Client
}

func New(endpoint, accessKeyId, accessKeySecret string) *Client {
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		panic(err)
	}
	return &Client{
		cli: client,
	}
}

func (c *Client) UploadFile(bucket string, objKey string, bz []byte, contentType string) error {
	bkt, err := c.cli.Bucket(bucket)
	if err != nil {
		return err
	}
	// URL: filepath.Join(bucket,objKey)
	return bkt.PutObject(objKey, bytes.NewReader(bz), oss.ContentType(contentType))
}
