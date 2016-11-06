package handle

import (
	"beans"
	"constants"
	"fmt"
	"strings"
	"util"
)

// 处理消息的工厂类
func MsgFactory(body *beans.TextRequestBody) ([]byte, error) {

	fmt.Println("msg type = ", body.MsgType)
	// 文本
	if strings.EqualFold(body.MsgType, "text") {

		fmt.Println("Wechat Service : recv Msg is text!")

		//抽奖处理工厂
		content := strings.Replace(body.Content, " ", "", -1)

		if strings.EqualFold(content, constants.GAME_CODE) {
			return GameFactory(body)
		} else {

			subContent := util.Substr(body.Content, 0, 1)

			if strings.EqualFold(subContent, "a") {
				return FkGameFactory(body)
			} else {
				return HandleText(body)
			}
		}

	} else if strings.EqualFold(body.MsgType, "voice") { //声音

		fmt.Println("wechat service : recv msg is voice!")
		return HandleVoice(body)

	} else if strings.EqualFold(body.MsgType, "image") {

		fmt.Println("wechat service : recv msg is image!")

	} else if strings.EqualFold(body.MsgType, "event") {

		fmt.Println("wechat service : recv msg is event!")
		// 订阅
		if strings.EqualFold(body.Event, "subscribe") {

			fmt.Println("wechat service : recv msg is subscribe!")
			return HandleSubscribe(body)

		} else if strings.EqualFold(body.Event, "unsubscribe") { //取消订阅

			fmt.Println("wechat service : recv msg is unsubscribe!")

		}
	}

	return MakeTextResponseBody(body.ToUserName, body.FromUserName, "微小象感谢你的支持!")
}
