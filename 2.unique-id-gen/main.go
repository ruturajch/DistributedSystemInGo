package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	rand.Seed(time.Now().UnixNano())
	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}
		body["type"] = "generate_ok"
		body["id"] = msg.Src + "-" + strconv.FormatInt(time.Now().UnixNano(), 10) + "-" + strconv.Itoa(rand.Intn(1000000))

		return n.Reply(msg, body)
	})
	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
