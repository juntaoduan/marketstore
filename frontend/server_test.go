package frontend_test

import (
	"testing"

	"github.com/alpacahq/marketstore/v4/frontend"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	tearDown, rootDir, metadata := setup(t, "TestNewServer")
	defer tearDown()

	serv, _ := frontend.NewServer(rootDir, metadata.CatalogDir)
	assert.True(t, serv.HasMethod("DataService.Query"))
}
