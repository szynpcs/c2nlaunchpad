package customerror

type BrainError struct {
	ErrorName string
}

func (brainError *BrainError) Error() string {
	return brainError.ErrorName
}
