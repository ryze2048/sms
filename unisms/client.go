package unisms

import unisms "github.com/apistd/uni-go-sdk/sms"

type UniSMS struct {
	Signature       string            `json:"signature"`
	AccessKeyId     string            `json:"access-key-id"`
	AccessKeySecret string            `json:"access-key-secret"`
	Phone           string            `json:"phone"`
	TemplateId      string            `json:"template_id"`
	Data            map[string]string `json:"data"`
}

func (u *UniSMS) NewClient() *unisms.UniSMSClient {
	return unisms.NewClient(u.AccessKeyId, u.AccessKeySecret)
}
