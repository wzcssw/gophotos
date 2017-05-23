package models

import (
	"io/ioutil"
	"photos/config"
	"strconv"
)

// Photo 图片
type Photo struct {
	FileName string
	FileSize int64
}

// GetNiceSize 格式化输出FileSize字段
func (photo *Photo) GetNiceSize() string {
	var num int64
	size := "Byte"
	if photo.FileSize > 1024 {
		num = photo.FileSize / 1024
		size = "KB"
	}
	return strconv.FormatInt(num, 10) + size
}

// Add 模板中得到序号+1
func (photo *Photo) Add(x, y int) int {
	return x + y
}

// GetFiles 获取文件
func GetFiles() []Photo {
	var photos []Photo
	sourceFolder := config.Conf.FilePath
	// 获取所有文件
	files, _ := ioutil.ReadDir(sourceFolder)
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			photos = append(photos, Photo{FileName: file.Name(), FileSize: file.Size()})
		}
	}
	return photos
}
