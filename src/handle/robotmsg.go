package handle

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 接入图灵机器人
const (
	APIKEY  = "94df8b1f5952f4dbcd2d4c8e2be4780a"
	POSTURL = "http://www.tuling123.com/openapi/api"
)

// 消息返回解析
type RobotResponseBody struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// post请求
func PostWord(word string) string {
	resp, err := http.PostForm(POSTURL,
		url.Values{"key": {APIKEY}, "userid": {"123"}, "info": {word}})
	if err != nil {
		return "象小伢在抢修 :-("
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "象小伢在抢修 :-("
	}
	var robot RobotResponseBody
	responseStr := ""
	err = json.Unmarshal([]byte(string(body)), &robot)
	if err == nil {
		responseStr = robot.Text
	} else {
		responseStr = "象小伢在抢修中!"
	}
	b := bytes.Buffer{}
	// b.WriteString("象小伢 : \n")
	b.WriteString(responseStr + "\r\n")
	return b.String()
}
