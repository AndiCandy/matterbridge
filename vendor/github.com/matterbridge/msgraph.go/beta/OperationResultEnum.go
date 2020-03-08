// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// OperationResult undocumented
type OperationResult int

const (
	// OperationResultVSuccess undocumented
	OperationResultVSuccess OperationResult = 0
	// OperationResultVFailure undocumented
	OperationResultVFailure OperationResult = 1
	// OperationResultVTimeout undocumented
	OperationResultVTimeout OperationResult = 2
	// OperationResultVUnknownFutureValue undocumented
	OperationResultVUnknownFutureValue OperationResult = 3
)

// OperationResultPSuccess returns a pointer to OperationResultVSuccess
func OperationResultPSuccess() *OperationResult {
	v := OperationResultVSuccess
	return &v
}

// OperationResultPFailure returns a pointer to OperationResultVFailure
func OperationResultPFailure() *OperationResult {
	v := OperationResultVFailure
	return &v
}

// OperationResultPTimeout returns a pointer to OperationResultVTimeout
func OperationResultPTimeout() *OperationResult {
	v := OperationResultVTimeout
	return &v
}

// OperationResultPUnknownFutureValue returns a pointer to OperationResultVUnknownFutureValue
func OperationResultPUnknownFutureValue() *OperationResult {
	v := OperationResultVUnknownFutureValue
	return &v
}