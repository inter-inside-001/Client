package main

import (
	"github.com/astaxie/beego"
	"net"
	"fmt"
	"bufio"
	"os"
	"strings"
)

// 处理错误
func CheckError(err error){
	if err != nil{
		beego.Error(err)
	}
}

// 消息发送协程
func MessageSend(conn net.Conn){
	for{
		// 从控制台读取信息
		reader := bufio.NewReader(os.Stdin)
		data,_,_ := reader.ReadLine()
		input := string(data)

		// 如果信息为exit，则退出
		if strings.ToUpper(strings.Trim(input, " ")) == "EXIT"{
			conn.Close()
			break
		}

		// 如果信息不为exit,则把信息发往服务器端
		_,err := conn.Write([]byte(input))
		if err != nil{
			conn.Close()
			fmt.Println("client connect failure:" + err.Error())
			break
		}
	}
}

func main(){
	fmt.Println("客户端已启动...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	defer conn.Close()
	CheckError(err)
	// 得到本地地址和服务器地址
	localAddr := conn.LocalAddr()
	localAddrStr := fmt.Sprintf("%s",localAddr)
	remoteAddr := conn.RemoteAddr()
	remoteAddrStr := fmt.Sprintf("%s",remoteAddr)
	fmt.Printf("客户端(%s)已经连接上服务器端(%s)\n", localAddrStr, remoteAddrStr)

	// 开启消息发送协程
	go MessageSend(conn)

	// 主协程负责接收消息
	buf := make([]byte, 1024)
	for{
		num,_ := conn.Read(buf)
		if num > 0 {
			fmt.Println("服务器端(%s): %s", remoteAddr, buf[0:num])
		}
	}
}

