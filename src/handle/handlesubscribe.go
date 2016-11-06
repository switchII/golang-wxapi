package handle

import (
	"beans"
	"bytes"
	"constants"
	"fmt"
)

// 处理关注信息
func HandleSubscribe(body *beans.TextRequestBody) ([]byte, error) {

	msg := body.Content
	fmt.Println("Wechat Service : ", msg)

	b := bytes.Buffer{}
	b.WriteString("你 : 关键词\r\n")

	b.WriteString("欢迎关注象伢\r\n一个认真分享经验的高校平台\r\n输入关键词可以查询出相关问题及回答\r\n如没有查询到，可在象伢平台提交你的疑问\r\n我们会邀请有经验的学姐学长回答你的问题:-)\r\n")

	b.WriteString("-------------------\r\n")
	b.WriteString(constants.MESSAGE_TAIL)

	responseText := b.String()
	responseTextBody, err := MakeTextResponseBody(body.ToUserName,
		body.FromUserName,
		responseText)

	return responseTextBody, err
}
