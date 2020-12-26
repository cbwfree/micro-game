package tool

import "os"

// ExistDir 检查目录是否存在
func ExistFolder(folder string) bool {
	if _, err := os.Stat(folder); !os.IsNotExist(err) {
		return true
	}
	return false
}

// MkFolder 创建目录
func MkFolder(folder string, perm ...os.FileMode) error {
	var p os.FileMode
	if len(perm) > 0 {
		p = perm[0]
	} else {
		p = 0755
	}
	return os.MkdirAll(folder, p)
}

func InitFolder(folder string, perm ...os.FileMode) error {
	if !ExistFolder(folder) {
		if err := MkFolder(folder, perm...); err != nil {
			return err
		}
	}
	return nil
}
