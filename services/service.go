package services

import (
	"context"

	"github.com/kiriyms/server_devops_practice_go/common"
)

type Service interface {
	Greet(context.Context) (string, error)
}

type Greeter struct{}

func NewGreeter() *Greeter {
	return &Greeter{}
}

func (g *Greeter) Greet(ctx context.Context) (string, error) {
	uid := common.GetUserId()
	return "Hello, user " + uid + "!", nil
}
