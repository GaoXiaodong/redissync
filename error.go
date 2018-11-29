package redissync

import "errors"

var ErrFailed = errors.New("redissync: failed to acquire lock")
