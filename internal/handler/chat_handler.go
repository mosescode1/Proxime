package handler

import "github.com/labstack/echo/v4"

type ChatHandler struct{}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (h *ChatHandler) Health(c echo.Context) error {
	return nil
}

func (h *ChatHandler) GetChat(c echo.Context) error {
	return nil
}
