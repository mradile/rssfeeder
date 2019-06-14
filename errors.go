package rssfeeder

import "errors"

var (
	ErrEmptyLogin   = errors.New("empty login")
	ErrEmptyURI     = errors.New("uri is empty")
	ErrNotAllowed   = errors.New("not allowed")
	ErrEntryMissing = errors.New("entry does not exist")
)
