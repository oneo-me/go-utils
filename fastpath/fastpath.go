package fastpath

import (
	"os"
	"path"
	"path/filepath"
)

// Abs 获取绝对路径
func Abs(p string) (string, error) {
	return filepath.Abs(p)
}

// Dir 获取所在目录
func Dir(p string) string {
	return filepath.Dir(p)
}

// Name 获取名称
func Name(p string) string {
	return filepath.Base(p)
}

// Ext 获取扩展名
func Ext(p string) string {
	return filepath.Ext(p)
}

// Join 连接路径
func Join(elem ...string) string {
	return path.Join(elem...)
}

// Delete 删除文件或目录
func Delete(p string) error {
	return os.RemoveAll(p)
}
