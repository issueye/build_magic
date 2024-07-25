package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 获取程序运行目录
func GetWorkDir() string {
	pwd, _ := os.Getwd()
	return pwd
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

func FileGetName(path string) string {
	base := filepath.Base(path)
	fileSuffix := filepath.Ext(base)
	return strings.TrimSuffix(base, fileSuffix)
}

func FileExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

// 创建 tar 文件
func CreateTarFile(sourceDir, targetFile string) error {
	// 创建输出文件
	file, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建 gzip 写入器
	gzipWriter := gzip.NewWriter(file)
	defer gzipWriter.Close()

	// 创建 tar 写入器
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// 遍历源目录并添加到 tar 文件
	err = filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 创建 tar 头信息
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = filepath.Join(filepath.Base(sourceDir), strings.TrimPrefix(path, sourceDir))

		// 写入头信息
		if err := tarWriter.WriteHeader(header); err != nil {
			return err
		}

		// 如果是目录,则无需写入内容
		if info.IsDir() {
			return nil
		}

		// 写入文件内容
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(tarWriter, file)
		return err
	})
	return err
}

// 解压 tar 文件
func ExtractTarFile(tarFile, targetDir string) error {
	// 打开 tar 文件
	file, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建 gzip 读取器
	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	// 创建 tar 读取器
	tarReader := tar.NewReader(gzipReader)

	// 遍历 tar 文件内容并提取
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		// 创建文件或目录
		path := filepath.Join(targetDir, header.Name)
		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			_, err = io.Copy(file, tarReader)
			if err != nil {
				file.Close()
				return err
			}
			file.Close()
		default:
			fmt.Printf("Unsupported file type %c in %s\n", header.Typeflag, path)
		}
	}

	return nil
}
