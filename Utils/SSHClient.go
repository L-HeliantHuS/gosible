package Utils

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"runtime"
	"time"
)

// SSHConnect 连接SSH
func SSHConnect(hostname string, username string, password string, port int, cmd string) {
	/*
		hostname: 主机ip
		username: 用户名
		password: 密码 (用私钥的话可以空着 "")
		port: 端口
		cmd: 要执行的命令
	*/
	sshHost := hostname
	sshUser := username
	sshPass := password
	sshPort := port
	sshCommand := cmd
	sshKeyPath := ""
	sshWindowsKeyPath := "C:/Users/Administrator/.ssh/id_rsa"
	sshLinuxKeyPath := "~/.ssh/id_rsa.pub"

	if runtime.GOOS == "windows" {
		sshKeyPath = sshWindowsKeyPath
	} else if runtime.GOOS == "linux" {
		sshKeyPath = sshLinuxKeyPath
	}

	// 创建config
	config := &ssh.ClientConfig{
		User:            sshUser,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Second,
	}
	if sshPass != "" {
		config.Auth = []ssh.AuthMethod{ssh.Password(sshPass)}
	} else {
		config.Auth = []ssh.AuthMethod{publicKeyAuthFunc(sshKeyPath)}
	}

	// 构造ssh主机:端口
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		panic(fmt.Sprintf("连接主机: %s 失败, 错误描述: %s", addr, err))
	}

	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		panic(fmt.Sprintf("%s, 创建Session失败: %s", addr, err))
	}
	defer session.Close()

	output, err := session.CombinedOutput(sshCommand)
	if err != nil {
		panic(fmt.Sprintf("%s 执行命令: %s 失败, 错误描述: %s", addr, sshCommand, err))
	}
	fmt.Println(fmt.Sprintf("%s执行命令%s成功:\n %s", addr, sshCommand, string(output)))

}
