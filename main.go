// Load Balancer backend discovery over IRC

package main

import (
	"fmt"
	"github.com/lrstanley/girc"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	client := girc.New(girc.Config{
		Server: "localhost",
		Port:   6667,
		Nick:   "test",
		User:   "user",
		Name:   "Example bot",
		Debug:  os.Stdout,
	})

	client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
		c.Cmd.Join("#dev")
	})

	client.Handlers.Add(girc.PRIVMSG, func(c *girc.Client, e girc.Event) {
		if strings.Contains(e.Last(), "add") {
			c.Cmd.ReplyTo(e, fmt.Sprintf("%s", e.Last()))
			return
		}

		if strings.Contains(e.Last(), "quit") {
			c.Close()
		}
	})

	// An example of how you would add reconnect logic.
	for {
		if err := client.Connect(); err != nil {
			log.Printf("error: %s", err)

			log.Println("reconnecting in 30 seconds...")
			time.Sleep(30 * time.Second)
		} else {
			return
		}
	}
}
