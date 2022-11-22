/*
 * @Author: hanshan-macbookair 2625406970@qq.com
 * @Date: 2022-11-22 23:54:40
 * @LastEditors: hanshan-macbookair 2625406970@qq.com
 * @LastEditTime: 2022-11-22 23:54:45
 * @FilePath: /K8s-Learn/v3/main.go
 * @Description:
 *
 * Copyright (c) 2022 by hanshan-macbookair 2625406970@qq.com, All Rights Reserved.
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	io.WriteString(w, fmt.Sprintf("[v3] Hello, Kubernetes!, From host: %s", host))
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":3000", nil)
}
