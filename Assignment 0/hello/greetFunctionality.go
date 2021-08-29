package hello

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayGreetings(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	if in.Body == "bye" {
		return &Message{Body: "Bye! - From the Server"}, nil
	} else {
		return &Message{Body: "Hello, There! - From the Server"}, nil
	}
}
