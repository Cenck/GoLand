package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://fundgz.1234567.com.cn/js/161725.js?v=20210208111814"
	res, _ := http.Get(url)
	get, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(get))
}
