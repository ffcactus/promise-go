package main

import (
	"os"
	"os/signal"
	"flag"
	"log"
	"net/url"
	"time"
	"github.com/gorilla/websocket"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	var host = flag.String("host", "10.61.18.217:8081", "remote address")
	flag.Parse()
	log.Println("Args count = %d", len(flag.Args()))
	if len(flag.Args()) < 1 {
		log.Println("Usage: connector remote_URI username password.")
	}

	u := url.URL{Scheme: "ws", Host: *host, Path: "/ws"}
	log.Println("Dialing to %v", u)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("Dial remote failed, err = %v", err)
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read message failed, err = %v", err)
				return
			}
			log.Println("%s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
			if err != nil {
				log.Println("write message failed, err = %v", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")
			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return			
		}
	}
}