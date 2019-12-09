package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

func main() {

	// 设置最大开n个原生线程
	runtime.GOMAXPROCS(8)

	go PostTestDemo()
	go PostTestDemo2()
	go PostTestDemo3()
	go PostTestDemo4()
	go PostTestDemo5()

	var input string
	fmt.Scan(&input)
}

func PostTestDemo() {
	for i := 5000; i < 7000; i++ {
		PostTest(fmt.Sprintf("%d@qq.com", i))
	}
}
func PostTestDemo2() {
	for i := 7000; i < 9000; i++ {
		PostFormTest(fmt.Sprintf("%d@qq.com", i))
	}
}
func PostTestDemo3() {
	for i := 9000; i < 11000; i++ {
		PostFormTest(fmt.Sprintf("%d@qq.com", i))
	}
}
func PostTestDemo4() {
	for i := 11000; i < 13000; i++ {
		PostFormTest(fmt.Sprintf("%d@qq.com", i))
	}
}
func PostTestDemo5() {
	for i := 13000; i < 15000; i++ {
		PostFormTest(fmt.Sprintf("%d@qq.com", i))
	}
}

func PostTest(email string) {
	fmt.Println("PostTest " + email)
	u := "http://localhost:8080/api/register"

	////序列化对象
	//body, err := json.Marshal(user)
	//if err != nil {
	//	fmt.Println("marshal fail")
	//	return
	//}
	data := fmt.Sprintf("email=%s&password=123456", email) //设置提交数据
	resp, err1 := http.Post(u, "application/x-www-form-urlencoded", strings.NewReader(data))
	if err1 != nil {
		fmt.Println("post fail")
		return
	}
	//解析响应的数据
	defer resp.Body.Close()
	b, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("read body fail")
		return
	}
	fmt.Println(string(b))
}
func PostFormTest(email string) {
	fmt.Println("PostFormTest " + email)
	//构造参数
	u, err := url.Parse("http://localhost:8080/api/register")
	if err != nil {
		fmt.Println("parse is fail")
		return
	}
	q := u.Query()
	q.Set("name", "大大王")
	q.Set("sex", "男")

	mapdata := url.Values{"email": {email}, "password": {"123456"}}
	//发起post表单请求
	resp, err1 := http.PostForm(u.String(), mapdata)
	if err1 != nil {
		fmt.Println("post fail")
		return
	}
	//解析响应
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("body fail")
		return
	}
	fmt.Println(string(body))
}
