package syncLog

type Node struct {
	isSelf      bool
	address     []string
	connections []connection
	ownerLog    []*Log
}

func
