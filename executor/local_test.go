package executor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLocalExecuto(t *testing.T) {
	e := NewLocalExecutor()
	assert.IsType(t, LocalExecutor{}, e)
}

func Test_Install(t *testing.T) {
	e := NewLocalExecutor()
	i := InitOpts{ServiceFilePath: "/tmp"}
	err := e.Install(i)
	assert.Nil(t, err)
}

func Test_InstallErr(t *testing.T) {
	e := NewLocalExecutor()
	i := InitOpts{ServiceFilePath: "/not/valid/dir"}
	err := e.Install(i)
	assert.Equal(t, "open /not/valid/dir/j.unit: no such file or directory", err.Error())
}
