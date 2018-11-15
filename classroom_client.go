package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)
func CheckError(err error){
	if err != nil{
		//fmt.Println("Error :%s", err.Error())
		//os.Exit(1)
		panic(err)
	}
}

func main(){
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	CheckError(err)
	defer conn.Close()
	go MessageSend(conn)

	buf := make([]byte, 1024)
	for{
		_,err := conn.Read(buf)
		if err != nil{
			fmt.Println("您已经退出 欢迎使用")
			os.Exit(0)
		}
		//CheckError(err)
		fmt.Println("Receive server message content:" + string(buf))
	}
	//conn.Write([]byte("hello beifengwang3!"))
	fmt.Println("Client program end")
}

func MessageSend(conn net.Conn){
	var input string
	for {
		reader := bufio.NewReader(os.Stdin)
		data,_,_ := reader.ReadLine()
		input = string(data)

		if strings.ToUpper(input) == "EXIT"{
			conn.Close()
			break
		}

		_,err := conn.Write(data)
		if err != nil{
			conn.Close()
			fmt.Println("Client connect failure: " + err.Error())
		}
	}
}

