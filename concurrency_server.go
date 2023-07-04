package main

import (
	"fmt"
	"time"
)

type Message struct {
	From    string
	Payload string
}

type Server struct {
	msgch  chan Message
	quitch chan struct{}
}

func (s *Server) startAndListen() {
running:
	for {
		select {
		// block here until someone is sending a message to the channel
		case msg := <-s.msgch:
			fmt.Printf("received message from %s payload %s\n", msg.From, msg.Payload)
		case <-s.quitch:
			fmt.Printf("The server is doing a gracefull shutdown")
			break running
		}

	}

	fmt.Println("The server is shut down.")
}

func sendMessageToServer(msgch chan Message, payload string) {
	msg := Message{
		From:    "YoeBuyDem",
		Payload: payload,
	}
	msgch <- msg
}

func graceFullQitServer(quitch chan struct{}) {
	close(quitch)
}

func run() {

	s := &Server{
		msgch:  make(chan Message),
		quitch: make(chan struct{}),
	}
	go s.startAndListen()

	go func() {
		time.Sleep(2 * time.Second)
		sendMessageToServer(s.msgch, "Hello Sailor!")
	}()

	go func() {
		time.Sleep(4 * time.Second)
		graceFullQitServer(s.quitch)
	}()

	time.Sleep(5 * time.Second)

}
