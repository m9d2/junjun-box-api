package model

type Client struct {
	Id   string
	Chan chan []byte
}
