package file

import "os"

// Exists returns whether the given file or directory exists or not
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

func Size(path string) int64 {
	f, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return f.Size()
}
