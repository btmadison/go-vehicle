package inmem_test

import (
	"testing"

	"github.com/btmadison/go-vehicle/pkg/data/inmem"
)

func TestNewRepository_CreatesSeededRepo(t *testing.T) {
	r := inmem.NewRepository()
	if r == nil {
		t.Error("shit")
	}
}
