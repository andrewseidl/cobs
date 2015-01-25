package main

import (
	"encoding/json"
	"fmt"
	"github.com/iron-io/iron_go/mq"
	"syscall"
	"time"
)

type Image struct {
	Id     string `json:"id"`
	Tag    string `json:"tag"`
	Source string `json:"source"`
}

type Message struct {
	Body string
	Id   string
}

func UtsnameToByte(name [65]int8) []byte {
	b := make([]byte, len(name))
	i := 0
	for ; i < len(name); i++ {
		if name[i] == 0 {
			break
		}
		b[i] = byte(name[i])
	}
	return b
}

func GetMachineName() string {
	uts := &syscall.Utsname{}
	err := syscall.Uname(uts)
	if err != nil {
		fmt.Println(err)
	}

	return string(UtsnameToByte(uts.Machine))
}

func GetBuildRequests(queue *mq.Queue, wait time.Duration) <-chan Message {
	c := make(chan Message)

	go func() {
		for {
			msg, err := queue.Get()
			if err != nil {
				time.Sleep(wait)
			} else {
				fmt.Println("Request: " + msg.Body)
				c <- Message{msg.Body, msg.Id}
			}
		}
	}()
	return c
}

func main() {
	//machine := GetMachineName()

	queue := mq.New("builder-x86_64")

	request := GetBuildRequests(queue, 10*time.Second)

	for {
		select {
		case msg := <-request:
			var mymsg Image
			json.Unmarshal([]byte(msg.Body), &mymsg)
			fmt.Println(mymsg)
			queue.DeleteMessage(msg.Id)
		}
	}
}
