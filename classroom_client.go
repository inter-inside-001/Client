package main

import (
	"github.com/astaxie/beego"
	"net"
	"fmt"
)

// 处理错误
func CheckError(err error){
	if err != nil{
		beego.Error(err)
	}
}

func main(){
	fmt.Println("客户端已启动...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	CheckError(err)

	conn.Write([]byte("hello world"))
}

