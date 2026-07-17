package file

import (
	"fmt"
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

	srcInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}
	if dstInfo, statErr := os.Stat(dst); statErr == nil {
		if os.SameFile(srcInfo, dstInfo) {
			return fmt.Errorf("file: source and destination are the same file")
		}
	} else if !os.IsNotExist(statErr) {
		return statErr
	}

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	dstFile, err := os.CreateTemp(filepath.Dir(dst), ".copy-*")
	if err != nil {
		return err
	}
	tmpPath := dstFile.Name()
	keepTemp := false
	defer func() {
		if !keepTemp {
			_ = os.Remove(tmpPath)
		}
	}()

	if err = dstFile.Chmod(srcInfo.Mode().Perm()); err == nil {
		_, err = io.Copy(dstFile, srcFile)
	}
	if err == nil {
		err = dstFile.Sync()
	}
	if closeErr := dstFile.Close(); closeErr != nil && err == nil {
		err = closeErr
	}
	if err != nil {
		return err
	}
	if err = os.Rename(tmpPath, dst); err != nil {
		return err
	}
	keepTemp = true
	return nil
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
