package beans

import (
	"encoding/xml"
	"time"
)

// 返回格式
type CDATAText struct {
	Text string `xml:",innerxml"`
}

// 文本消息
type TextRequestBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   time.Duration
	MsgType      string
	Event        string
	Content      string
	MsgId        int
	Recognition  string
	Format       string
}

// 文本返回信息
type TextResponseBody struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATAText
	FromUserName CDATAText
	CreateTime   time.Duration
	MsgType      CDATAText
	Content      CDATAText
}

// 提示返回信息
type BackMsg struct {
	Tip bool   `json:"tip"`
	Msg string `json:"msg"`
}
