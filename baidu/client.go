package baidu

import "github.com/baidubce/bce-sdk-go/services/sms"

type BaiduSMS struct {
	AccessKey   string                 `json:"access-key"`
	SecretKey   string                 `json:"secret_key"`
	Endpoint    string                 `json:"endpoint"`
	ContentMap  map[string]interface{} `json:"content_map"`
	Phone       string                 `json:"phone"`
	Template    string                 `json:"template"`
	SignatureId string                 `json:"signature_id"`
}

func (b *BaiduSMS) NewClient() (*sms.Client, error) {
	return sms.NewClient(b.AccessKey, b.SecretKey, b.Endpoint)
}
