package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"SettleGuard/internal/yellow"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	client, err := yellow.NewClient(yellow.SandboxEndpoint)
	if err != nil {
		log.Fatalf("Failed to connect to Yellow Network: %v", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			fmt.Printf("Error during graceful shutdown: %v\n", err)
		} else {
			fmt.Println("Connection closed gracefully.")
		}
	}()

	go client.Listen(ctx, func(resp yellow.RPCResponse) {
		if resp.Error != nil {
			fmt.Printf("RPC Error [%d]: %s\n", resp.Error.Code, resp.Error.Message)
		} else {
			if len(resp.Result) == 0 || string(resp.Result) == "null" {
				fmt.Println("Success: (Empty Response)")
			} else {
				fmt.Printf("Success: %s\n", string(resp.Result))
			}
		}
	})

	fmt.Println("Yellow Client is running. Press Ctrl+C to exit.")

	<-ctx.Done()

	time.Sleep(100 * time.Millisecond)
}
