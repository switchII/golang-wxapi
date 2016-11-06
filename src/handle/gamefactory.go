package handle

import (
	"beans"
	"bytes"
	"fmt"
	// "game"
)

//凡科领取信息
func FkGameFactory(body *beans.TextRequestBody) ([]byte, error) {

	// gift := game.PlayGiftgame(body.FromUserName)
	msg := body.Content
	fmt.Println("Wechat FkGame Service : ", msg)
	b := bytes.Buffer{}

	b.WriteString("你 : " + msg + "\r\n")
	b.WriteString("你的兑换信息已经录入,请于9月17日(百团大战)至广西民族大学二坡篮球场(万人坑)计算机俱乐部摊点领取奖品.\r\n")
	b.WriteString("感谢参与'奔跑吧!外卖小哥'活动，对我们平台有任何建议直接在公众号内留言，感谢对平台支持!\r\n")
	b.WriteString("以上为暂定兑奖地点!\r\n")

	b.WriteString("-------------------\r\n")
	b.WriteString("<a href='https://hd.faisco.cn/10634913/9/load.html?style=16'>@奔跑吧!外卖小哥</a>")

	responseText := b.String()
	responseTextBody, err := MakeTextResponseBody(body.ToUserName,
		body.FromUserName,
		responseText)

	return responseTextBody, err

}

//游戏处理工厂
func GameFactory(body *beans.TextRequestBody) ([]byte, error) {

	// gift := game.PlayGiftgame(body.FromUserName)
	msg := body.Content
	fmt.Println("Wechat Game Service : ", msg)
	b := bytes.Buffer{}

	/*
		if gift.GiftNum == 2 {
			b.WriteString("你 : 奖品已抽完:)\r\n")
			b.WriteString("感谢你的参与本次活动!\r\n")
		} else if gift.GiftNum == 1 {
			b.WriteString("你 : 恭喜你:) \r\n")
			b.WriteString("<a href='http://www.linesno.com'>" + gift.GiftName + "</a>\r\n请到我们代理小店领取，感谢你的参与!\r\n")
			b.WriteString("领取码" + gift.GetCode + "\r\n")
		} else {
			b.WriteString("你 : 没抽中:( \r\n")
			b.WriteString("感谢你的参考，我们中奖率接近0.75 ,还有很多奖品哦，更有机会赢取我们的iPhone6s哦，邀请你同学来参与吧.\r\n")
		}
	*/

	b.WriteString("游戏开始啦:)\r\n")
	b.WriteString("感谢你的参与本次活动!,点击直接进入\r\n")
	b.WriteString("<a href='https://hd.faisco.cn/10634913/9/load.html?style=16'>奔跑吧！外卖小哥</a>\r\n")

	b.WriteString("-------------------\r\n")
	b.WriteString("<a href='http://www.linesno.com'>@象伢问答</a>")

	// b.WriteString("抽奖调度阶段，未正式开始抽奖\r\n<a href='http://www.linesno.com'>@象伢问答</a>")

	responseText := b.String()
	responseTextBody, err := MakeTextResponseBody(body.ToUserName,
		body.FromUserName,
		responseText)

	return responseTextBody, err

}
