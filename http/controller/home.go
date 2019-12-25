package controller

import (
	"fmt"
	"net/http"
)
// w表示response对象，返回给客户端的内容都在对象里处理
// r表示客户端请求对象，包含了请求头，请求参数等等
func Index(w http.ResponseWriter, r *http.Request){
	// 往w里写入内容，就会在浏览器里输出
	fmt.Fprintf(w, "你好，欢迎访问阿苏若软体!")
}
