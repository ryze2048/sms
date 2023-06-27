package unisms

import (
	"github.com/apistd/uni-go-sdk"
	unisms "github.com/apistd/uni-go-sdk/sms"
)

func (u *UniSMS) Send() (response *uni.UniResponse, err error) {
	client := u.NewClient()
	message := unisms.BuildMessage()
	message.SetTo(u.Phone)
	message.SetSignature(u.Signature)
	message.SetTemplateId(u.TemplateId)
	message.SetTemplateData(u.Data)
	return client.Send(message)
}
