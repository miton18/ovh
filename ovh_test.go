package ovh

import (
	"github.com/miton18/ovh/config"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := config.Auto()
	if err != nil {
		t.Errorf("Auto() error = %v", err)
		return
	}

	_, err = New(c)
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
}
