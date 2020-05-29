package factoryPattern

import (
	"context"
	"encoding/json"
	"log"
)

type ComponentRepository interface {
	AddComponent(context.Context, Component) (*Component, error)
	GetComponent(context.Context, string) (*Component, error)
}

type componentServer struct {
	repo ComponentRepository
}

func NewComponentServer(repo ComponentRepository) *componentServer {
	return &componentServer{repo: repo}
}

func (c *componentServer) AddComponent(ctx context.Context, comp Component) error {
	added, err := c.repo.AddComponent(ctx, comp)
	if err != nil {
		return err
	}
	log.Println("[svc] added component : ", added)
	return nil
}

func (c *componentServer) GetComponent(ctx context.Context, code string) (string, error) {
	found, err := c.repo.GetComponent(ctx, code)
	if err != nil {
		return "", err
	}
	bytes, err := json.Marshal(found)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
