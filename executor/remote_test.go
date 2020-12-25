package executor

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_NewSSHExecutor(t *testing.T) {
	ReadFile = func(filename string) ([]byte, error) {
		return []byte("foo"), nil
	}

	e := NewSSHExecutor("0.0.0.0", "ubuntu", "", InitOpts{})

	assert.Equal(t, "0.0.0.0", e.Address)
	assert.Equal(t, "ubuntu", e.ClientConfig.User)
}
