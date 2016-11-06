package main

import (
	"ask"
	"encoding/json"
	"fmt"
	"handle"
	"log"
	"net/http"
)

// 处理接收的提问信息
func ProcAsk(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("Wechat Service : Ask Request Method : [%s]", r.Method)
	if r.Method == "POST" {
		back := ask.AskSave(r)
		if b, err := json.Marshal(back); err == nil {
			fmt.Fprintf(w, string(b))
		}
	}
}

// 处理接收的微信消息
func ProcRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("Wechat Service : Request Method : [%s]", r.Method)
	if r.Method == "POST" {
		textRequestBody := handle.ParserTextRequestBody(r)
		if textRequestBody != nil {

			responseTextBody, err := handle.MsgFactory(textRequestBody)
			if err != nil {
				log.Println("Wechat Service : MakeTextResponseBody error", err)
				return
			}
			fmt.Fprintf(w, string(responseTextBody))
		}
	} else if r.Method == "GET" {
		fmt.Fprintf(w, string("Wechat Request Method!"))
	}

}

// 微信服务器
func server() {

	log.Println("Wechat Service start!")
	http.HandleFunc("/", ProcRequest)
	http.HandleFunc("/ask", ProcAsk)
	// http.HandleFunc("/", ValidateRequest)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("Wechat Service : ListenAndServe failed , ", err)
	}
	log.Println("Wechat Service : stop!")
}

func main() {
	fmt.Println("wxapi runing")
	server()
}
