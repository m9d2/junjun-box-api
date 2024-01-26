package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"junjun-box-api/config"
	"junjun-box-api/model"
	"log/slog"
	"net/http"
	"net/url"
	"sync"
	"time"
)

var clients = make(map[chan []byte]struct{})
var mu sync.Mutex

type ChatHandler struct {
}

func (h ChatHandler) InitRouter(g *gin.RouterGroup) {
	g.GET("message", h.handleMessage)
	g.GET("events", h.events)
}

func (h ChatHandler) events(c *gin.Context) {
	w := c.Writer
	r := c.Request
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan []byte)
	clients[messageChan] = struct{}{}

	defer func() {
		delete(clients, messageChan)
		close(messageChan)
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher.Flush()
	for {
		select {
		case <-r.Context().Done():
			return
		case msg := <-messageChan:
			_, err := fmt.Fprintf(w, "data: %s\n\n", msg)
			if err != nil {
				mu.Unlock()
				return
			}
			flusher.Flush()
		}
	}
}

func (h ChatHandler) handleMessage(c *gin.Context) {
	q := c.Query("q")
	h.completion(q)
	JSON(c, nil)
}

func (h ChatHandler) broadcastMessages() {
	for {
		// Simulate periodic broadcasts
		time.Sleep(time.Second * 5)

		message := model.Event{
			Event: "message",
			Data:  "",
		}
		bytes, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for client := range clients {
			client <- bytes
		}
	}
}

func (h ChatHandler) completion(question string) {
	cfg := openai.DefaultConfig(config.Conf.Openai.Token)

	if config.Conf.Openai.ProxyUrl != "" {
		proxyUrl, err := url.Parse(config.Conf.Openai.ProxyUrl)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		cfg.HTTPClient = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}

	client := openai.NewClientWithConfig(cfg)
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		MaxTokens: 2048,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		},
		Stream: true,
	}
	stream, err := client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	defer stream.Close()
	currentTime := time.Now()
	nanoTimestamp := currentTime.UnixNano()
	millisecondTimestamp := nanoTimestamp / int64(time.Millisecond)
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			slog.Error(err.Error())
			return
		}

		if err != nil {
			slog.Error(err.Error())
			return
		}
		fmt.Print(response.Choices[0].Delta.Content)
		message := model.Event{
			Event: "message",
			Id:    millisecondTimestamp,
			Data:  response.Choices[0].Delta.Content,
			Time:  time.Now(),
		}

		bytes, err := json.Marshal(message)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for client := range clients {
			client <- bytes
		}
	}
}
