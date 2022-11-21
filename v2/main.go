/*
 * @Author: hanshan-macbookair 2625406970@qq.com
 * @Date: 2022-11-20 20:13:32
 * @LastEditors: hanshan-macbookair 2625406970@qq.com
 * @LastEditTime: 2022-11-20 20:13:55
 * @FilePath: /K8s-Learning/v2/main.go
 * @Description:
 *
 * Copyright (c) 2022 by hanshan-macbookair 2625406970@qq.com, All Rights Reserved.
 */
package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v2] Hello, Kubernetes!")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
