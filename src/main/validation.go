package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strings"
)

// Token常量
const (
	token = "switchjs"
)

// 生成加密签名
func makeSignature(timestmp string, noice string) string {
	sl := []string{token, timestmp, noice}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// 检验的链接
func validateUrl(w http.ResponseWriter, r *http.Request) bool {

	timestmp := strings.Join(r.Form["timestamp"], "")
	noice := strings.Join(r.Form["nonce"], "")
	signatureGen := makeSignature(timestmp, noice)

	signatureIn := strings.Join(r.Form["signature"], "")

	if signatureGen != signatureIn {
		return false

	}

	echostr := strings.Join(r.Form["echostr"], "")
	fmt.Fprintf(w, echostr)

	return true
}

// 获取Url
func ValidateRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if !validateUrl(w, r) {
		log.Println("Wechat Service : this http request is not from Wechat platform")
		fmt.Fprintf(w, string("Wechat Service : this http request is not from Wechat platform!"))
		return
	}
	log.Println("Wechat Service : validateUrl OK!")
}
