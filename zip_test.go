package archiver

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestZipToWriter(t *testing.T) {
	outFile, err := os.Create("/Users/hyh/Downloads/test/a.zip")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	defer func() {
		_ = outFile.Close()
	}()

	ZipToWriter("/Users/hyh/Downloads/test/a", outFile)
}

func TestZipToBytesFromFile(t *testing.T) {
	buf, err := ZipToBytes("/Users/hyh/Downloads/httpd/node.msi")
	if err != nil {
		fmt.Println(err)
		t.Fail()
		return
	}
	ioutil.WriteFile("/Users/hyh/Downloads/node.msi.zip", buf, 0644)
}
