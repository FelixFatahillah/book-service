package exception

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
)

type ErrWithCode struct {
	Code int
	Err  error
}

type RPCError struct {
	Code    codes.Code
	Message string
}

func (e *RPCError) Error() string {
	return e.Message
}

func (e *ErrWithCode) Error() string {
	return fmt.Sprintf("error[%d]: %v", e.Code, e.Err)
}

type ErrValidation struct {
	Message string
}

func (e *ErrValidation) Error() string {
	return e.Message
}

var ErrOptimisticLock = errors.New("error trx lock")
