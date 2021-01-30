package utils

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

type FileUtils struct {
	_filename   string
	_flag       int
	_flagAppend int
	_fileMode   os.FileMode
}

func NewFileUtils(filename string) *FileUtils {
	// windows 路径 反斜杠 替换成 斜杠
	filename = strings.ReplaceAll(filename, "\\", "/")
	// 操作方式
	flag := os.O_RDWR | os.O_CREATE
	flagAppend := os.O_RDWR | os.O_CREATE | os.O_APPEND
	// 文件权限
	filemode := os.ModePerm

	return &FileUtils{
		_filename:   filename,
		_flag:       flag,
		_fileMode:   filemode,
		_flagAppend: flagAppend,
	}
}

// 创建目录
func (f FileUtils) CreateDir(path string) error {

	// 创建目录所有的父目录
	err := os.MkdirAll(f._filename, os.ModePerm)
	if err != nil {
		return err
	}
	// 改变权限 777
	err = os.Chmod(path, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// 读取文件
func (f FileUtils) ReadData(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		println("ReadData 异常: ", err.Error())
		return ""
	}

	return string(data)
}

// 写入数据
func (f FileUtils) WriteData(content string) {
	data := []byte(content)
	err := ioutil.WriteFile(f._filename, data, os.ModePerm)
	if err != nil {
		println("WriteData 异常: ", err.Error())
	}
}

// 追加内容
func (f FileUtils) AppendData(content string) {
	// 打开文件
	file, err := os.OpenFile(f._filename, f._flagAppend, f._fileMode)

	if err != nil {
		println("AppendData  OpenFile 异常: ", err.Error())
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)

	if err != nil {
		println("AppendData WriteString 异常: ", err.Error())
		return
	}
}

// 获取文件扩展名
// (例子: a/name.txt -> .txt)
func (f FileUtils) GetExtension() string {
	return path.Ext(f._filename)
}

// 获取最后一级 目录/文件  名字
// (例子: a/b/c  -> c)
func (f FileUtils) GetLastLevelPath() string {
	return path.Base(f._filename)
}

// 获取 除了最后一级以外的地址
// (例子: a/b/c -> a/b)
func (f FileUtils) GetExceptLastLevelPath() string {
	return path.Dir(f._filename)
}

// 获取等效的最短路径名
func (f FileUtils) GetEqualShortPath() string {
	return path.Clean(f._filename)
}

// 是否是 绝对路径
// 相对路径 ->  ../a/b/c.txt
// 绝对路径 ->  /a/b/c.txt
func (f FileUtils) IsAbsolutePath() bool {
	return path.IsAbs(f._filename)
}

// 获取 目录名  与  文件名
func (f FileUtils) GetDirAndFileName() (dir, fileName string) {
	return path.Split(f._filename)
}

// 是否是 文件
func (f FileUtils) IsFile() bool {
	fileInfo, err := os.Stat(f._filename)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func (f FileUtils) IsDir() bool {
	s, err := os.Stat(f._filename)
	if err != nil {
		return false
	}

	return s.IsDir()
}

// 判断 文件夹/文件 是否存在
func (f FileUtils) IsExist() (bool, error) {
	_, err := os.Stat(f._filename)
	if err != nil {
		return true, nil // 文件存在
	} else {
		if os.IsNotExist(err) {
			return false, nil // 文件不存在
		} else {
			return false, err // 文件不存在, 未知错误
		}
	}
}

func GetDirItemList(dirPath string) *[]os.FileInfo {
	fileInfoList, _ := ioutil.ReadDir(dirPath)
	return &fileInfoList
}
