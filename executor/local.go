package executor

type LocalExecutor struct {
	InitOpts InitOpts
}

func NewLocalExecutor(opts InitOpts) LocalExecutor {
	e := LocalExecutor{
		InitOpts: opts,
	}

	return e
}
