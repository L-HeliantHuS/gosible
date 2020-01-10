package main

import (
	"flag"
	"fmt"
	"os"
	"gosible/Utils"
)

var (
	Help     string
	Host     string
	Username string
	Password string
	Command  string
	Port     int
)

func usage() {
	fmt.Fprintf(os.Stderr, `gosbile version: gosbile/1.0.0
Usage: gosbile [-h hostname] [-u username] [-P password] [-p port] [-c command]

Options:
`)
	flag.PrintDefaults()
}

func init() {
	flag.StringVar(&Help, "help", "", "显示帮助")
	flag.StringVar(&Host, "h", "", "填写主机地址")
	flag.StringVar(&Username, "u", "", "填写用户名")
	flag.StringVar(&Password, "P", "", "填写密码, 填写空会自动找当前主机的私钥进行连接")
	flag.StringVar(&Command, "c", "", "填写要执行的命令, 多条命令使用','分开")
	flag.IntVar(&Port, "p", 22, "填写端口, 默认22")
	flag.Usage = usage
}

func main() {
	flag.Parse()

	if Username != "" && Host != "" && Command != "" {
		Utils.SSHConnect(Host, Username, Password, Port, Command)
	} else {
		flag.Usage()
	}

}
