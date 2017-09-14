package archiver

import (
	"archive/zip"
	"bytes"
	"encoding/binary"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//ZipToBytes zip file or directory to []byte
//@params source<string> need zip file path or directory path
//@return []btye, error
func ZipToBytes(source string) ([]byte, error) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)
	// Create a new zip archive.
	w := zip.NewWriter(buf)
	basePath := filepath.Dir(source)
	err := filepath.Walk(source, func(path string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return nil
		}
		relativeFilePath, err := filepath.Rel(basePath, path)
		if err != nil {
			return err
		}
		f, err := w.Create(relativeFilePath)
		if err != nil {
			return err
		}
		fileContent := []byte(path)
		fileContent, err = ioutil.ReadFile(path)
		_, err = f.Write([]byte(fileContent))
		return err
	})
	if err != nil {
		return nil, err
	}
	// Make sure to check the error on Close.
	// Cannot use `defer w.close()` , Because filepath.Walk's closure maybe have not finish after `w.close()` run
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//Zip zip file or diretory to target file path
func Zip(source string, target string) error {
	buf, err := ZipToBytes(source)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(target, buf, 0644)
}

//Unzip unzip file
func Unzip(source string, target string) error {
	body, err := ioutil.ReadFile(source)
	if err != nil {
		return err
	}
	err = UnzipFromBytes(body, target)
	return err
}

//UnzipFromBytes unzip a file bytes
func UnzipFromBytes(body []byte, target string) error {
	r, err := zip.NewReader(bytes.NewReader(body), int64(binary.Size(body)))
	if err != nil {
		return nil
	}
	for _, zf := range r.File {
		if zf.FileInfo().IsDir() {
			continue
		}
		targetFile := filepath.Join(target, filepath.Join(strings.Split(zf.Name, "/")...))
		err := os.MkdirAll(filepath.Dir(targetFile), os.FileMode(0755))
		if err != nil {
			return err
		}
		dst, err := os.Create(targetFile)
		if err != nil {
			return err
		}
		src, err := zf.Open()
		if err != nil {
			return err
		}
		io.Copy(dst, src)
		dst.Close()
		src.Close()
	}
	return nil

}
