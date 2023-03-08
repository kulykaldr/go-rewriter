package main

import (
	"context"
	"fmt"
	"github.com/kulykaldr/go-rewriter/chatgpt"
	"log"
	"os"
)

func main() {
	cfg := &chatgpt.Config{
		Login:    "vadnik759@gmail.com",
		Password: "Qwerty2023",
		IsLogin:  true,
		Headless: false,
		Debug:    false,
		Timeout:  15,
	}

	fmt.Println(cfg)

	cl := chatgpt.NewClient(cfg)
	ctx, cancel := cl.CreateContext(context.Background())
	defer cancel()

	if cfg.IsLogin {
		err := cl.SignIn(ctx)
		if err != nil {
			log.Fatalf("signin error: %v\n", err)
		}
	}

	text, _ := os.ReadFile("text.txt")

	rText, err := cl.Rewrite(ctx, string(text), "REWRITE")
	if err != nil {
		log.Fatalf("rewrite error: %v\n", err)
	}

	if err = os.WriteFile("answer.txt", []byte(rText), 0644); err != nil {
		log.Fatalf("write text file error: %v\n", err)
	}

	err = cl.Close(ctx)
	if err != nil {
		log.Fatalf("close: %v\n", err)
	}
}
