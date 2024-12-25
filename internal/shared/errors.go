package shared

import "errors"

var (
	ErrorDuplicateTrx = errors.New("duplicate trx, please create new one")
)
