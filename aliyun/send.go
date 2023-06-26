package aliyun

import (
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

func (a *AliyunSMS) Send() (response *dysmsapi20170525.SendSmsResponse, err error) {
	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		TemplateParam: tea.String(a.TemplateParam),
		TemplateCode:  tea.String(a.TemplateCode),
		SignName:      tea.String(a.SignName),
		PhoneNumbers:  tea.String(a.PhoneNumber),
	}
	var client *dysmsapi20170525.Client
	if client, err = a.NewClient(); err != nil {
		return nil, err
	}
	runtime := &util.RuntimeOptions{}
	return client.SendSmsWithOptions(sendSmsRequest, runtime)
}
