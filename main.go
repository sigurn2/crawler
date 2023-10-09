package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "https://www.xx.cn/"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)
	exist := strings.Contains(string(body), "gold")
	fmt.Printf("是否存在黄金：%v\n", exist)

}
