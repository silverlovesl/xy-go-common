package utils

import (
	"archive/zip"
	"io/ioutil"
	"os"
)

// CreateZipFileContent Zipファイルの中身を設定する
func CreateZipFileContent(writer *zip.Writer, zipPath string, targetFilePath string) error {
	// zip化対象のファイルを設定
	f, err := os.Open(targetFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := os.Stat(targetFilePath)
	if err != nil {
		return err
	}

	// zipファイル内のディレクトリ構造を設定
	header, _ := zip.FileInfoHeader(info)
	header.Name = zipPath
	if !info.IsDir() {
		header.Method = zip.Deflate
	}

	zf, err := writer.CreateHeader(header)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	body, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	if _, err := zf.Write(body); err != nil {
		return err
	}
	return nil
}
