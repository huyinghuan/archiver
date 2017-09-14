## Zip File or Direcotry

### Get

```
go get github.com/huyinghuan/archiver
```

### Doc

```
use 'godoc cmd/archiver/' for documentation on the archiver/ command 

PACKAGE DOCUMENTATION

package archiver
    import "archiver"


FUNCTIONS

func Unzip(source string, target string) error
    Unzip unzip file

func UnzipFromBytes(body []byte, target string) error

func Zip(source string, target string) error
    Zip zip file or diretory to target file path

func ZipToBytes(source string) ([]byte, error)
    ZipToBytes zip file or directory to []byte @params source<string> need
    zip file path or directory path @return []btye, error


```