package factoryPattern

import (
	"context"
	"encoding/json"
	"fmt"
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

	val, err := found.Compliance.Calculate()
	if err != nil {
		return "", err
	}
	fmt.Println("calculated compliance[SG/IN/...] : ", val)

	switch found.CountryCode {
	case "SG":
		singaporeCompliance := found.Compliance.(*SingaporeCompliance)
		singaporeCompliance.SingaporeSpecificFunction()
	case "IN":
		indiaCompliance := found.Compliance.(*IndiaCompliance)
		indiaCompliance.IndiaSpecificFunction()
	}

	bytes, err := json.Marshal(found)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
