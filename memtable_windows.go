//go:build windows
// +build windows

package badger

import (
	"errors"
	"syscall"
)

const errnoUserMappedFile = syscall.Errno(1224)

func isUserMappedFileError(err error) bool {
	return errors.Is(err, errnoUserMappedFile)
}
