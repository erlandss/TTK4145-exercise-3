package main

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

const (
	ADDRESS = "localhost:20000"
)

func main() {
	exec.Command("cmd", "/C", "start", "powershell", "go", "run", "phoenix.go").Run()
	time.Sleep(time.Second * 3)

	addr, _ := net.ResolveUDPAddr("udp", ADDRESS)
	con, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		msg := fmt.Sprintf("%d", i)
		con.Write([]byte(msg))
		time.Sleep(time.Second)
	}
}
