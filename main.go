package main

import (
	"flag"
	"fmt"
	"github.com/mzky/zip"
	"io"
	"os"
)

func main() {
	// 文件名
	var filename string
	// 密码
	var password string
	// 输出文件名
	var outputFileName string

	flag.StringVar(&filename, "f", "", "文件名,默认为空")
	flag.StringVar(&password, "p", "", "密码,默认为空")
	flag.StringVar(&outputFileName, "o", "", "输出文件名,默认为空")
	flag.Parse()

	fmt.Printf("filename=%v,password=%v，outputFileName=%v",filename,password,outputFileName)
	var array []string
	array = append(array, filename)
	err := Zip(outputFileName, password, array)
	if err != nil {
		fmt.Printf("err=%v",err.Error())
	}
	fmt.Printf("result=success")
}

// Zip password值可以为空""
func Zip(zipPath, password string, fileList []string) error {
	if len(fileList) < 1 {
		return fmt.Errorf("将要压缩的文件列表不能为空")
	}
	fz, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(fz)
	defer zw.Close()

	for _, fileName := range fileList {
		fr, err := os.Open(fileName)
		if err != nil {
			return err
		}

		// 写入文件的头信息
		var w io.Writer
		if password != "" {
			w, err = zw.Encrypt(fileName, password, zip.AES256Encryption)
		} else {
			w, err = zw.Create(fileName)
		}

		if err != nil {
			return err
		}

		// 写入文件内容
		_, err = io.Copy(w, fr)
		if err != nil {
			return err
		}
	}
	return zw.Flush()
}
