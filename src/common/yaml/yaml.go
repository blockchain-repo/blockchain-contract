package yaml

import (
	"errors"
	"io/ioutil"
	"os"
)

import (
	"gopkg.in/yaml.v2"
)

/*
 * yaml文件读取函数
 * \param [in] filenpath : yaml文件的全路径
 * \param [in, out] out : 读取出的yaml文件内容，可以是struct或者map
 * \return nil 正常
 * \       not nil 有错误
 */
func Unmarshal(filenpath string, out interface{}) error {
	detail, err := readFile(filenpath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(detail, out)
}

/*
 * yaml文件写入函数
 * \param [in] filenpath : 需要写入的yaml文件的全路径，如果存在同名文件，将覆盖原文件
 * \param [in] in : 想要写入到yaml文件的内容，可以是struct或者map
 * \return nil 正常
 * \       not nil 有错误
 */
func Marshal(filenpath string, in interface{}) error {
	detail, err := yaml.Marshal(in)
	if err != nil {
		return err
	}

	if n, err := writeFile(filenpath, string(detail)); err != nil {
		return err
	} else {
		if n != len(detail) {
			return errors.New("write file failed")
		} else {
			return err
		}
	}
}

func readFile(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return []byte(""), err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func writeFile(fileName, content string) (int, error) {
	file, err := os.Create(fileName)
	//file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return 0, err
	}
	defer file.Close()
	return file.WriteString(content)
}
