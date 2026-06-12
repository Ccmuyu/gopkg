# file

文件系统操作工具。

## 安装

```go
import "github.com/Ccmuyu/gopkg/file"
```

## 快速开始

```go
file.Exists("/tmp/foo")          // true/false
file.IsDir("/tmp")               // true
file.IsFile("/tmp/foo.txt")      // true

data, _ := file.Read("/tmp/foo.txt")
_ = file.Write("/tmp/bar.txt", []byte("hello"), 0644)

file.Copy("/tmp/a.txt", "/tmp/b.txt")
file.Move("/tmp/a.txt", "/tmp/c.txt")
```

## API 参考

```go
func Exists(path string) bool
func IsDir(path string) bool
func IsFile(path string) bool
func Read(path string) ([]byte, error)
func ReadString(path string) (string, error)
func Write(path string, data []byte, perm os.FileMode) error
func WriteString(path string, data string, perm os.FileMode) error
func Copy(src, dst string) error
func Move(src, dst string) error
func Remove(path string) error
func RemoveAll(path string) error
func ListDir(path string) ([]string, error)
func TempDir(prefix string) (string, error)
func TempFile(prefix string) (*os.File, error)
func Ext(path string) string
func Basename(path string) string
func Dir(path string) string
```
