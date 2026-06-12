package file

import (
	"os"
	"testing"

	. "github.com/Ccmuyu/gopkg/test"
)

func TestExists(t *testing.T) {
	AssertTrue(t, Exists("file.go"))
	AssertTrue(t, !Exists("nonexistent_file_xyz"))
}

func TestIsDir(t *testing.T) {
	AssertTrue(t, IsDir("."))
	AssertTrue(t, !IsDir("file.go"))
}

func TestIsFile(t *testing.T) {
	AssertTrue(t, IsFile("file.go"))
	AssertTrue(t, !IsFile("."))
}

func TestWriteAndRead(t *testing.T) {
	path := "test_write_read.tmp"
	defer os.Remove(path)

	err := Write(path, []byte("hello world"), 0644)
	AssertTrue(t, err == nil)

	data, err := Read(path)
	AssertTrue(t, err == nil)
	AssertEqual(t, string(data), "hello world")
}

func TestWriteStringAndReadString(t *testing.T) {
	path := "test_write_str.tmp"
	defer os.Remove(path)

	err := WriteString(path, "hello world", 0644)
	AssertTrue(t, err == nil)

	data, err := ReadString(path)
	AssertTrue(t, err == nil)
	AssertEqual(t, data, "hello world")
}

func TestCopy(t *testing.T) {
	src := "test_copy_src.tmp"
	dst := "test_copy_dst.tmp"
	defer os.Remove(src)
	defer os.Remove(dst)

	WriteString(src, "copy content", 0644)
	err := Copy(src, dst)
	AssertTrue(t, err == nil)

	data, _ := ReadString(dst)
	AssertEqual(t, data, "copy content")
}

func TestMove(t *testing.T) {
	src := "test_move_src.tmp"
	dst := "test_move_dst.tmp"
	defer os.Remove(dst)

	WriteString(src, "move content", 0644)
	err := Move(src, dst)
	AssertTrue(t, err == nil)
	AssertTrue(t, !Exists(src))
	AssertTrue(t, Exists(dst))

	data, _ := ReadString(dst)
	AssertEqual(t, data, "move content")
}

func TestRemove(t *testing.T) {
	path := "test_remove.tmp"
	WriteString(path, "to remove", 0644)
	AssertTrue(t, Exists(path))
	err := Remove(path)
	AssertTrue(t, err == nil)
	AssertTrue(t, !Exists(path))
}

func TestListDir(t *testing.T) {
	names, err := ListDir(".")
	AssertTrue(t, err == nil)
	AssertTrue(t, len(names) > 0)
}

func TestTempDir(t *testing.T) {
	dir, err := TempDir("test-")
	AssertTrue(t, err == nil)
	AssertTrue(t, Exists(dir))
	os.RemoveAll(dir)
}

func TestTempFile(t *testing.T) {
	f, err := TempFile("test-")
	AssertTrue(t, err == nil)
	path := f.Name()
	AssertTrue(t, Exists(path))
	f.Close()
	os.Remove(path)
}

func TestExt(t *testing.T) {
	AssertEqual(t, Ext("file.go"), ".go")
	AssertEqual(t, Ext("file"), "")
}

func TestBasename(t *testing.T) {
	AssertEqual(t, Basename("/path/to/file.go"), "file.go")
}

func TestDir(t *testing.T) {
	AssertEqual(t, Dir("/path/to/file.go"), "/path/to")
}
