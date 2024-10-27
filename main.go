package main

import (
	"log"

	"github.com/gliderlabs/ssh"
	"golang.org/x/term"
)

func main() {
	ssh.Handle(func(s ssh.Session) {
		term := term.NewTerminal(s, "> ")
		line := ""
		for {
			line, _ = term.ReadLine()
			if line == "quit" {
				break
			}
		}
	})

	log.Println("starting ssh server on port 2222...")
	log.Fatal(ssh.ListenAndServe(":2222", nil))
}