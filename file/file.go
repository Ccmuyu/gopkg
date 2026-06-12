package file

import (
	"io"
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func Read(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func ReadString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Write(path string, data []byte, perm os.FileMode) error {
	if perm == 0 {
		perm = 0644
	}
	return os.WriteFile(path, data, perm)
}

func WriteString(path string, data string, perm os.FileMode) error {
	return Write(path, []byte(data), perm)
}

func Copy(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func Move(src, dst string) error {
	if err := os.Rename(src, dst); err == nil {
		return nil
	}
	if err := Copy(src, dst); err != nil {
		return err
	}
	return os.Remove(src)
}

func Remove(path string) error {
	return os.Remove(path)
}

func RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func ListDir(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	names := make([]string, len(entries))
	for i, entry := range entries {
		names[i] = entry.Name()
	}
	return names, nil
}

func TempDir(prefix string) (string, error) {
	return os.MkdirTemp("", prefix)
}

func TempFile(prefix string) (*os.File, error) {
	return os.CreateTemp("", prefix)
}

func Ext(path string) string {
	return filepath.Ext(path)
}

func Basename(path string) string {
	return filepath.Base(path)
}

func Dir(path string) string {
	return filepath.Dir(path)
}
