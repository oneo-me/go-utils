package fastpath

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Create 创建目录
func Create(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

// IsDir 是目录
func IsDir(p string) bool {
	info, err := os.Stat(p)
	return err == nil && info.IsDir()
}

// GetExeDir 获取程序所在目录
func GetExeDir() (string, error) {
	exeFile, err := GetExeFile()
	if err == nil {
		return filepath.Dir(exeFile), nil
	}
	return "", err
}

// GetDirs 获取子目录
func GetDirs(dir string) []string {
	dirs := []string{}
	dir, err := filepath.Abs(dir)
	if err == nil && IsDir(dir) {
		infos, err := ioutil.ReadDir(dir)
		if err == nil {
			for _, info := range infos {
				if info.IsDir() {
					dirs = append(dirs, Join(dir, info.Name()))
				}
			}
		}
	}
	return dirs
}

// GetFiles 获取子文件
func GetFiles(dir string) []string {
	files := []string{}
	dir, err := filepath.Abs(dir)
	if err == nil && IsDir(dir) {
		infos, err := ioutil.ReadDir(dir)
		if err == nil {
			for _, info := range infos {
				if !info.IsDir() {
					files = append(files, Join(dir, info.Name()))
				}
			}
		}
	}
	return files
}

// GetList 获取子内容
func GetList(dir string, includeChid bool) []string {
	list := []string{}
	for _, fdir := range GetDirs(dir) {
		list = append(list, fdir)
		if includeChid {
			for _, l := range GetList(fdir, includeChid) {
				list = append(list, l)
			}
		}
	}
	for _, file := range GetFiles(dir) {
		list = append(list, file)
	}
	return list
}

// GetDirSize 获取目录大小
func GetDirSize(dir string) int64 {
	var size int64
	list := GetList(dir, true)
	for _, file := range list {
		info, err := os.Stat(file)
		if err == nil && !info.IsDir() {
			size += info.Size()
		}
	}
	return size
}

// ForEach 遍历目录，返回 true 跳出遍历
func ForEach(dir string, includeChid bool, action func(string, bool) bool) {
	list := GetList(dir, includeChid)
	for _, file := range list {
		info, err := os.Stat(file)
		if err == nil {
			if action(file, !info.IsDir()) {
				break
			}
		}
	}
}

// Move 移动
func Move(src, dst string) error {
	err := Copy(src, dst)
	if err == nil {
		return Delete(src)
	}
	return err
}

// Copy 复制
func Copy(src, dst string) error {
	src, err := Abs(src)
	if err != nil {
		return err
	}
	dst, err = Abs(dst)
	if err != nil {
		return err
	}
	if IsFile(src) {
		CopyFile(src, dst)
	} else {
		if err := Create(dst); err != nil {
			return err
		}
		for _, file := range GetList(src, false) {
			if err := Copy(file, strings.Replace(file, src, dst, -1)); err != nil {
				return err
			}
		}
	}
	return nil
}
