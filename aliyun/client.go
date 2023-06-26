package aliyun

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
)

type AliyunSMS struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyId     string `json:"access-key-id"`
	AccessKeySecret string `json:"access-key-secret"`
	PhoneNumber     string `json:"phone_number"`   // 手机号码
	SignName        string `json:"sign_name"`      // 签名
	TemplateParam   string `json:"template_param"` // 模版 {"code":"xxx"}
	TemplateCode    string `json:"template_code"`  // 短信模版code SMS_XXX
}

func (a *AliyunSMS) NewClient() (*dysmsapi20170525.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(a.AccessKeyId),
		AccessKeySecret: tea.String(a.AccessKeySecret),
	}
	config.Endpoint = tea.String(a.Endpoint)
	return dysmsapi20170525.NewClient(config)
}
