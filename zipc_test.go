package zipc_test

import (
	"testing"

	"github.com/mcandre/zipc"
)

func TestVersion(t *testing.T) {
	if zipc.Version == "" {
		t.Errorf("Expected zipc version to be non-blank")
	}
}
