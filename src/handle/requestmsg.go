package handle

import (
	"beans"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// 接收文本消息
func ParserTextRequestBody(r *http.Request) *beans.TextRequestBody {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	fmt.Println(string(body))
	requestBody := &beans.TextRequestBody{}
	xml.Unmarshal(body, requestBody)
	return requestBody
}

// 数据结果转换
func value2CDATA(v string) beans.CDATAText {
	return beans.CDATAText{"<![CDATA[" + v + "]]>"}
}

// 接收文本返回信息
func MakeTextResponseBody(fromUserName, toUserName, content string) ([]byte, error) {

	body := beans.TextResponseBody{}

	body.ToUserName = value2CDATA(toUserName)
	body.FromUserName = value2CDATA(fromUserName)
	body.MsgType = value2CDATA("text")
	body.Content = value2CDATA(content)
	body.CreateTime = time.Duration(time.Now().Unix())

	return xml.MarshalIndent(body, "", "")

}
