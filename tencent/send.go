package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
)

func (t *TencentSMS) Send() (response *sms.SendSmsResponse, err error) {
	var client *sms.Client
	if client, err = t.NewClient(); err != nil {
		return nil, err
	}
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(t.AppId)
	request.SignName = common.StringPtr(t.SignName)
	request.TemplateId = common.StringPtr(t.TemplateId)
	request.TemplateParamSet = common.StringPtrs(t.TemplateParam)
	request.PhoneNumberSet = common.StringPtrs(t.PhoneNumber)
	return client.SendSms(request)
}
