package executor

import (
	"os"
	"path"
)

type LocalExecutor struct {
}

func NewLocalExecutor() LocalExecutor {
	e := LocalExecutor{}

	return e
}

// Install installs the rke binary with the desired config
func (e *LocalExecutor) Install(opts InitOpts) error {
	path := path.Join(opts.ServiceFilePath, serviceFileName)
	f, err := os.Create(path)
	if err != nil {
		return err
	}

	b := createInitConfig(opts)

	_, err = f.Write(b.Bytes())
	if err != nil {
		f.Close()
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}
