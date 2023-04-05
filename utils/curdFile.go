package utils

import "os"

// CurdFile 对文件的增删改查操作
type CurdFile struct {
}

func (receiver *CurdFile) Create(path string, content string, errorOpt ErrorLogOpt) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		ErrorLog(errorOpt)
	}
	_, _ = file.WriteString(content)
	return file
}

func (receiver *CurdFile) Update() *CurdFile {
	return receiver
}

func (receiver *CurdFile) Read(path string, errorOpt ErrorLogOpt) string {
	file, err := os.ReadFile(path)
	if err != nil {
		ErrorLog(errorOpt)
	}
	return string(file)
}

func (receiver *CurdFile) Delete(path string) *CurdFile {
	_ = os.RemoveAll(path)
	return receiver
}

// OpenFile open file
func (receiver *CurdFile) OpenFile(path string, errorOpt ErrorLogOpt) *os.File {
	file, readErr := os.Open(path)
	if readErr != nil {
		ErrorLog(errorOpt)
	}
	return file
}
