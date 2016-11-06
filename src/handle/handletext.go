package handle

import (
	"beans"
	"bytes"
	"constants"
	"dao"
	"fmt"
	"strconv"
	"strings"
)

// 处理文本信息
func HandleText(body *beans.TextRequestBody) ([]byte, error) {

	msg := body.Content
	fmt.Println("Wechat Service : ", msg)

	b := bytes.Buffer{}
	b.WriteString("你 : " + msg + "\r\n")

	b.WriteString(analyseText(msg))

	b.WriteString("-------------------\r\n")
	b.WriteString(constants.MESSAGE_TAIL)

	responseText := b.String()
	responseTextBody, err := MakeTextResponseBody(body.ToUserName,
		body.FromUserName,
		responseText)

	return responseTextBody, err
}

// 分析内容,返回结果
func analyseText(word string) string {

	// output := search.Keyword(word)
	// fmt.Println("分词:", output)

	word = strings.Replace(word, " ", "%", -1)
	b := bytes.Buffer{}

	db, err := dao.GetConnction()
	if err != nil {
		return ""
	}
	sql := "select question_id , question_content , answer_count from aws_question where question_content like '%" + word + "%' order by update_time desc limit 0,10"
	fmt.Println("sql = ", sql)
	rows, err := dao.Select(db, sql)
	var qid int
	var content string
	var answerCount int
	count := 0
	for rows.Next() {
		count++
		err = rows.Scan(&qid, &content, &answerCount)
		b.WriteString("<a href='http://www.linesno.com/?/question/" + strconv.Itoa(qid) + "'>" + content + "(" + strconv.Itoa(answerCount) + "个回复)</a>\r\n")
	}

	if count == 0 {
		b.WriteString("尝试分词搜索，如'广西民族大学'修改为'广西 民大',或直接把你的疑问提交给象伢问答平台:-)\r\n")
		b.WriteString("象伢，一个认真分享经验的高校平台!\r\n")
	}
	dao.CloseDb(db)

	return b.String()
}
