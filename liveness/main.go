/*
 * @Author: hanshan-macbookair 2625406970@qq.com
 * @Date: 2022-11-20 20:33:38
 * @LastEditors: hanshan-macbookair 2625406970@qq.com
 * @LastEditTime: 2022-11-20 20:33:44
 * @FilePath: /K8s-Learning/liveness/main.go
 * @Description:
 *
 * Copyright (c) 2022 by hanshan-macbookair 2625406970@qq.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "[v2] Hello, Kubernetes!")
}

func main() {
	started := time.Now()
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() > 15 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("ok"))
		}
	})

	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
