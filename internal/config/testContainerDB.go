package config

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type TestDatabase struct {
	Instance testcontainers.Container
	Host     string
	Port     int
}

func InitPostgresContainer() (*TestDatabase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	req := testcontainers.ContainerRequest{
		Image:        "postgres",
		ExposedPorts: []string{"5432/tcp"},
		AutoRemove:   true,
		Env: map[string]string{
			"POSTGRES_USER":     "postgres",
			"POSTGRES_PASSWORD": "postgres",
			"POSTGRES_DB":       "postgres",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}
	postgres, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	mappedPort, err := postgres.MappedPort(ctx, "5432")
	if err != nil {
		return nil, err
	}

	hostIP, err := postgres.Host(ctx)
	if err != nil {
		return nil, err
	}

	return &TestDatabase{
		Instance: postgres,
		Host:     hostIP,
		Port:     mappedPort.Int(),
	}, nil
}
