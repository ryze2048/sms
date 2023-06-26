package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111" // 引入sms
)

type TencentSMS struct {
	SecretId      string   `json:"secret_id"`
	SecretKey     string   `json:"secret_key"`
	Endpoint      string   `json:"endpoint"`
	AppId         string   `json:"app_id"`
	SignName      string   `json:"sign_name"`
	TemplateId    string   `json:"template_id"`
	TemplateParam []string `json:"template_param"`
	PhoneNumber   []string `json:"phone_number"`
	Region        string   `json:"region"`
}

func (t *TencentSMS) NewClient() (*sms.Client, error) {
	credential := common.NewCredential(t.SecretId, t.SecretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.Endpoint = t.Endpoint
	cpf.SignMethod = "HmacSHA1"
	return sms.NewClient(credential, t.Region, cpf)
}
