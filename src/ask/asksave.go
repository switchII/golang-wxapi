package ask

import (
	"beans"
	"fmt"
	"net/http"
	"strings"
)

// 保存问题
func AskSave(r *http.Request) beans.BackMsg {

	openId := strings.Join(r.Form["openId"], "")
	content := strings.Join(r.Form["content"], "")
	urgent := strings.Join(r.Form["urgent"], "")

	sql := "insert into wx_question(openid , content , urgent ,addtime) value(" + openId + "," + content + "," + urgent + ",now())"
	fmt.Println("open id = ", openId, " , content = ", content, " , urgent = ", urgent, " , sql = ", sql)

	back := beans.BackMsg{Tip: true, Msg: "保存成功"}

	return back
}
