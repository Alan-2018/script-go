package baidubce

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/sms"
	"github.com/baidubce/bce-sdk-go/services/sms/api"
)

func TestBaiduBceSms() {
	var (
		accessId  string = "****"
		accessKey string = "****"
		sign      string = "****"
		template  string = "****"

		endpoint string = "http://smsv3.bj.baidubce.com"
	)

	client, err := sms.NewClient(accessId, accessKey, endpoint)
	if err != nil {
		fmt.Printf("sms new client error, %s", err)
		return
	}

	client.Config.ConnectionTimeoutInMillis = 30 * 1000

	contentMap := make(map[string]interface{})
	contentMap["code"] = "12345"
	contentMap["minute"] = "10"
	sendSmsArgs := &api.SendSmsArgs{
		Mobile:      "xxxxxxxxxxx,xxxxxxxxxxx",
		SignatureId: sign,
		Template:    template,
		ContentVar:  contentMap,
	}

	sendSmsResult, err := client.SendSms(sendSmsArgs)
	if err != nil {
		fmt.Printf("send sms error, %s", err)
		return
	}

	fmt.Printf("send sms success. %s", sendSmsResult)
}
