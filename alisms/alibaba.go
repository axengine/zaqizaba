package alisms

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
)

type AliSMS struct {
	client                                *dysmsapi.Client
	signName, templateCode, templateParam string
}

func NewAliSMS(accessKeyId, accessKeySecret, signName, templateCode, templateParam string) *AliSMS {
	config := &openapi.Config{
		AccessKeyId:     &accessKeyId,
		AccessKeySecret: &accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	client, _ := dysmsapi.NewClient(config)

	return &AliSMS{
		client:        client,
		signName:      signName,
		templateCode:  templateCode,
		templateParam: templateParam,
	}
}

func (p *AliSMS) Send(receiver, code string) (string, error) {
	sendSmsRequest := &dysmsapi.SendSmsRequest{
		PhoneNumbers:  tea.String(receiver),
		SignName:      tea.String(p.signName),
		TemplateCode:  tea.String(p.templateCode),
		TemplateParam: tea.String(fmt.Sprintf(p.templateParam, code)),
	}
	result, err := p.client.SendSmsWithOptions(sendSmsRequest, &util.RuntimeOptions{})
	if err != nil {
		return "", err
	}
	if result != nil && result.StatusCode != nil && tea.Int32Value(result.StatusCode) != 200 {
		return "", errors.Errorf("Err:%v", result)
	}
	if result != nil && result.Body != nil && tea.StringValue(result.Body.Code) != "OK" {
		return "", errors.Errorf("Err:%v", result)
	}
	if result != nil && result.Body != nil && result.Body.RequestId != nil {
		return tea.StringValue(result.Body.RequestId), nil
	}

	return "", nil
}
