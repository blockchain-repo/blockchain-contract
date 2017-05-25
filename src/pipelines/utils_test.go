package pipelines

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

//读取args slice的每个字符串，将其作为文件名，读取文件,并在文件的每一行首部加上行号，写入到out中
//此处in没有使用到，主要是为了保证管道定义的一致性
func app1(in io.Reader, out io.Writer, args []string) {
	for _, v := range args {
		//fmt.Println(v)
		file, err := os.Open(v)
		if err != nil {
			continue
		}
		defer file.Close()
		buf := bufio.NewReader(file)
		for i := 1; ; i++ {
			line, err := buf.ReadBytes('\n')
			if err != nil {
				break
			}
			linenum := strconv.Itoa(i)
			nline := []byte(linenum + " ")
			nline = append(nline, line...)
			out.Write(nline)
		}
	}
}

//app2 主要是将字节流转化为大写,中文可能会有点问题，不过主要是演示用，重在理解思想
//read from in, convert byte to Upper ,write the result to out
func app2(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, 10)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			break
		}
		t := bytes.ToUpper(p[:n])
		out.Write(t)
	}
}

func app3(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, 10)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			break
		}
		t := bytes.ToLower(p[:n])
		out.Write(t)
	}
}

func Test_pipe(t *testing.T) {
	args := os.Args[1:]
	for _, v := range args {
		fmt.Println(v)
	}

	p := Pipe(
		Bind(app1, args),
		app2,
		app3,
		app2)

	p(os.Stdin, os.Stdout)
}

func TestInit(t *testing.T) {
	Init()
}
