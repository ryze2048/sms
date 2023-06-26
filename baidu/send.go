package baidu

import (
	"github.com/baidubce/bce-sdk-go/services/sms"
	"github.com/baidubce/bce-sdk-go/services/sms/api"
)

func (b *BaiduSMS) Send() (*api.SendSmsResult, error) {
	var err error
	var client *sms.Client
	if client, err = b.NewClient(); err != nil {
		panic(err)
	}
	sendSmsArgs := &api.SendSmsArgs{
		Mobile:      b.Phone,
		Template:    b.Template,
		SignatureId: b.SignatureId,
		ContentVar:  b.ContentMap,
	}
	return client.SendSms(sendSmsArgs)
}
