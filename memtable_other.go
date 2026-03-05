//go:build !windows
// +build !windows

package badger

func isUserMappedFileError(_ error) bool {
	return false
}
