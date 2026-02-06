package yellow

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn      *websocket.Conn
	mu        sync.Mutex
	done      chan struct{}
	closeOnce sync.Once
}

func NewClient(endpoint string) (*Client, error) {
	dialer := websocket.DefaultDialer
	conn, _, err := dialer.Dial(endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &Client{
		conn: conn,
		done: make(chan struct{}),
	}, nil
}

func (c *Client) Listen(ctx context.Context, handler func(RPCResponse)) {
	defer close(c.done)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			c.mu.Lock()
			if c.conn == nil {
				c.mu.Unlock()
				return
			}
			conn := c.conn
			c.mu.Unlock()

			_, message, err := conn.ReadMessage()
			if err != nil {
				return
			}

			var resp RPCResponse
			if err := json.Unmarshal(message, &resp); err == nil {
				handler(resp)
			}
		}
	}
}

// Close should clean up
func (c *Client) Close() error {
	var err error
	c.closeOnce.Do(func() {
		c.mu.Lock()
		defer c.mu.Unlock()
		if c.conn != nil {
			if dlErr := c.conn.SetWriteDeadline(time.Now().Add(time.Second)); dlErr == nil {
				_ = c.conn.WriteMessage(websocket.CloseMessage,
					websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			}
			err = c.conn.Close()
			c.conn = nil
		}
	})
	return err
}
