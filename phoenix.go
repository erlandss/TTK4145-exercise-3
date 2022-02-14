package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"time"
)

const (
	ADDRESS = "localhost:20000"
)

func main() {
	//ime.Sleep(10 * time.Second)
	go backupListener()
	for {

	}

}

func backupListener() {
	count := 0
	addr, _ := net.ResolveUDPAddr("udp", ADDRESS)
	pc, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	isCounting := false

	for !isCounting {
		buf := make([]byte, 1024)
		pc.SetReadDeadline(time.Now().Add(time.Second * 2))

		n, _, err := pc.ReadFromUDP(buf)
		if err != nil {
			if e, ok := err.(net.Error); !ok || !e.Timeout() {
				fmt.Println("Not timeout error; panic??")
			}

			if e, _ := err.(net.Error); e.Timeout() {
				//This shit is timed out yo
				fmt.Println("Timeout reached")
				time.Sleep(time.Second)
				//panic(err)
				exec.Command("cmd", "/C", "start", "powershell", "go", "run", "phoenix.go").Run()
				time.Sleep(time.Second * 1)
				fmt.Println("Starting new backup")
				//start ny instans
				go counter(count)
				//slutt å listen
				pc.Close()
				isCounting = true
			}
		}
		//parse msg
		count, _ = strconv.Atoi(string(buf[:n]))

	}
}

func counter(count int) {
	fmt.Println("Resuming counting")
	addr, _ := net.ResolveUDPAddr("udp", ADDRESS)
	con, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5; i++ {
		count++
		fmt.Println(count)
		msg := fmt.Sprintf("%d", count)
		con.Write([]byte(msg))
		time.Sleep(time.Second)

	}
	os.Exit(0)
}
