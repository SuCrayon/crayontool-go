package shell

const (
	// Success 0
	Success = iota
	// OperationNotPermitted 1
	OperationNotPermitted
	// NoSuchFileOrDict 2
	NoSuchFileOrDict

	CommandNotFound = 127
)