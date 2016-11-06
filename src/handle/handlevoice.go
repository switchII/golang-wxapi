package handle

import (
	"beans"
	"bytes"
	"constants"
	"fmt"
)

// 处理文本信息
func HandleVoice(body *beans.TextRequestBody) ([]byte, error) {

	msg := body.Recognition
	fmt.Println("Wechat Service : ", msg)

	b := bytes.Buffer{}
	b.WriteString("你 : " + msg + "\r\n")

	b.WriteString(analyseVoice(msg))

	b.WriteString("-------------------\r\n")
	b.WriteString(constants.MESSAGE_TAIL)

	responseText := b.String()
	responseTextBody, err := MakeTextResponseBody(body.ToUserName,
		body.FromUserName,
		responseText)

	return responseTextBody, err
}

// 分析内容,返回结果
func analyseVoice(word string) string {

	// output := search.Keyword(word)
	// fmt.Println("分词:", output)

	b := bytes.Buffer{}
	b.WriteString(PostWord(word))

	return b.String()
}
