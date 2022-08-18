package handlers

import (
	"context"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/thiiluh/kombibeer/internal/config"
)

func TestHandlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers Suite")
}

var tc *config.TestDatabase
var ctx context.Context

var _ = BeforeSuite(func() {
	tc, _ = config.InitPostgresContainer()
	ctx = context.Background()
	config.Connect(tc.Host, tc.Port)
	go InitRoutes()
	time.Sleep(2 * time.Second)
})

var _ = AfterSuite(func() {
	tc.Instance.Terminate(ctx)
})
