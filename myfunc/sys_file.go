package myfunc

import (
	"chianoob_server_linux/models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
)

//根据路径找到该路径下的所有文件，并返回
func GetFilesFromDir(path string) []models.UpFiles {
	files := []models.UpFiles{}
	mfile := models.UpFiles{}
	//从目录path下找到所有的目录文件
	allDir, err := ioutil.ReadDir(path)
	//如果有错误发生就返回nil，不再进行查找
	if err != nil {
		return files
	}
	//遍历获取到的每一个目录信息
	for _, dir := range allDir {
		if dir.IsDir() {
			//如果还是目录，就继续向下进行递归查找,并追加到返回切片中
			retfiles := GetFilesFromDir(path + "/" + dir.Name())
			for _, temp := range retfiles {
				files = append(files, temp)
			}
			continue
		}
		//如果不是目录,就读取该文件并加入到files中
		fileName := path + "/" + dir.Name()
		file, fileErr := ioutil.ReadFile(fileName)
		if fileErr != nil {
			return files
		}
		mfile.PathName = dir.Name()
		mfile.Md5 = ByteToMd5(file)
		//打印文件名和MD5
		fmt.Println(mfile.PathName + " 的MD5为 " + mfile.Md5)
		files = append(files, mfile)
	}
	return files
}

//根据路径找到该文件，并返回
func GetFilesMd5(path string) string {
	mfile := models.UpFiles{}

	file, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		return ""
	}
	mfile.PathName = path
	mfile.Md5 = ByteToMd5(file)
	//打印文件名和MD5
	fmt.Println(mfile.PathName + " 的MD5为 " + mfile.Md5)
	return mfile.Md5
}

//[]byte转md5
func ByteToMd5(fileByte []byte) string {
	has := md5.Sum(fileByte)
	//用新的切片存放
	has2 := has[:]
	md5Str := hex.EncodeToString(has2)
	return md5Str
}
