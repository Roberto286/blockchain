// Package server exposes a chain over raw TCP.
package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/davecgh/go-spew/spew"

	"blockchain/internal/chain"
)

const broadcastInterval = 30 * time.Second

type Server struct {
	chain *chain.Chain
}

func New(c *chain.Chain) *Server {
	return &Server{chain: c}
}

// Serve accepts connections on l until it fails.
func (s *Server) Serve(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()

	done := make(chan struct{})
	defer close(done)
	go s.broadcast(conn, done)

	io.WriteString(conn, "Enter a new BPM:")

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		bpm, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("%q not a number: %v", scanner.Text(), err)
			fmt.Fprint(conn, "\nEnter a new BPM:")
			continue
		}

		if _, err := s.chain.Add(bpm); err != nil {
			log.Println(err)
		} else {
			spew.Dump(s.chain.Blocks())
		}
		fmt.Fprint(conn, "\nEnter a new BPM:")
	}
}

// broadcast periodically writes the chain as JSON to conn until done is closed
// or a write fails.
func (s *Server) broadcast(conn net.Conn, done <-chan struct{}) {
	t := time.NewTicker(broadcastInterval)
	defer t.Stop()

	for {
		select {
		case <-done:
			return
		case <-t.C:
			out, err := json.Marshal(s.chain.Blocks())
			if err != nil {
				log.Println(err)
				return
			}
			if _, err := conn.Write(out); err != nil {
				return
			}
		}
	}
}
