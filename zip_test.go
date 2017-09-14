package archiver

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestZipToBytesFromFile(t *testing.T) {
	buf, err := ZipToBytes("/Users/hyh/Downloads/test/a")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	ioutil.WriteFile("/Users/hyh/Downloads/test/a.zip", buf, os.FileMode(0644))
}

func TestZip(t *testing.T) {
	err := Zip("/Users/hyh/Downloads/test/a", "/Users/hyh/test/a.zip")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestZip2(t *testing.T) {
	inFilePath := "/Users/hyh/Downloads/ShadowsocksX-NG.zip"
	outFilePath := "/Users/hyh/Downloads"
	err := Unzip(inFilePath, outFilePath)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestZipFromBytes(t *testing.T) {

	body, _ := ioutil.ReadFile("/Users/hyh/Downloads/ShadowsocksX-NG.zip")
	err := UnzipFromBytes(body, "/Users/hyh/Downloads")
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
