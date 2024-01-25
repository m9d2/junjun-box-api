package api

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-contrib/sse"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	"io"
	"junjun-box-api/config"
	"log/slog"
	"net/http"
	"net/url"
)

type ChatHandler struct {
}

func (h ChatHandler) InitRouter(g *gin.RouterGroup) {
	g.GET("completion", h.completion)
}

func (h ChatHandler) completion(c *gin.Context) {
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
		MaxTokens: 20,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: c.Query("q"),
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

	c.Header("Content-Type", "text/event-stream; charset=utf-8")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("Access-Control-Allow-Origin", "*")
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			slog.Error(err.Error())
			return
		}
		err = sse.Encode(c.Writer, sse.Event{
			Event: "message",
			Data: map[string]interface{}{
				"content": response.Choices[0].Delta.Content,
			},
		})
		if err != nil {
			slog.Error(err.Error())
			return
		}
		fmt.Print(response.Choices[0].Delta.Content)
	}

}
