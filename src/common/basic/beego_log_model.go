package basic

import (
	"github.com/astaxie/beego/logs"
	"os"
	"path"
)

// logs.AdapterFile
type MyBeegoLogAdapterFile struct {
	// filename 保存的文件名
	FileName string `json:"filename"`

	// maxlines 每个文件保存的最大行数，默认值 1000000
	MaxLines int64 `json:"maxlines"`

	// maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB
	MaxSize int64 `json:"maxsize"`

	// daily 是否按照每天 logrotate，默认是 true  // bool must set
	Daily bool `json:"daily"`

	// 文件最多保存多少天，默认保存 7 天
	MaxDays int16 `json:"maxdays"`

	// 是否开启 logrotate，默认是 true  // bool must set
	Rotate bool `json:"rotate"`

	// 日志保存的时候的级别，默认是 Trace 级别
	Level int `json:"level"`

	// 日志文件权限
	Perm string `json:"perm"`
}

type MyBeegoLogAdapterMultiFile struct {
	MyBeegoLogAdapterFile
	Separate []string `json:"separate"` //需要单独写入文件的日志级别,设置后命名类似 test.error.log
}

// newMyBeegoLogAdapterFile create a FileLogWriter returning as LoggerInterface.
func NewMyBeegoLogAdapterFile(myBeego *MyBeegoLogAdapterFile) *MyBeegoLogAdapterFile {
	filename := "project.log"
	maxlines := int64(1000000)
	maxsize := int64(1 << 28) //256 MB
	//daily := true // bool must set
	maxdays := int16(7)
	//rotate := true  // bool must set
	level := logs.LevelTrace
	perm := "0660"

	if myBeego.FileName != "" {
		filename = myBeego.FileName
	}

	if myBeego.MaxLines >= 0 {
		maxlines = myBeego.MaxLines
	}

	if myBeego.MaxSize >= 0 {
		maxsize = myBeego.MaxSize
	}

	if myBeego.MaxDays >= 0 {
		maxdays = myBeego.MaxDays
	}

	level = myBeego.Level

	if myBeego.Perm != "" {
		perm = myBeego.Perm
	}

	mm := &MyBeegoLogAdapterFile{
		FileName: filename,
		MaxLines: maxlines,
		MaxSize:  maxsize,
		Daily:    myBeego.Daily,
		MaxDays:  maxdays,
		Rotate:   myBeego.Rotate,
		Level:    level,
		Perm:     perm,
	}
	return mm
}

func NewMyBeegoLogAdapterMultiFile(myBeego *MyBeegoLogAdapterMultiFile) *MyBeegoLogAdapterMultiFile {
	filename := "project.log"
	maxlines := int64(1000000)
	maxsize := int64(1 << 28) //256 MB
	//daily := true // bool must set
	maxdays := int16(7)
	//rotate := true  // bool must set
	level := logs.LevelTrace
	perm := "0660"
	separate := []string{"emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"}

	if myBeego.FileName != "" {
		filename = myBeego.FileName
	}
	// create filepath if not exist!
	var sepator string
	if os.IsPathSeparator('\\') {  //前边的判断是否是系统的分隔符
		sepator = "\\"
	} else {
		sepator = "/"
	}
	log_dir := path.Dir(filename)
	_, err := os.Stat(log_dir)
	if err != nil{
		err := os.Mkdir(log_dir + sepator , os.ModePerm)
		if err != nil {
			logs.Error("create log dir error!", err)
		}else{
			logs.Info("API log dir", log_dir, " create success!")
		}

	}
	logs.Info("API log will store in dir", log_dir)

	if myBeego.MaxLines >= 0 {
		maxlines = myBeego.MaxLines
	}

	if myBeego.MaxSize >= 0 {
		maxsize = myBeego.MaxSize
	}

	if myBeego.MaxDays >= 0 {
		maxdays = myBeego.MaxDays
	}

	level = myBeego.Level

	if myBeego.Perm != "" {
		perm = myBeego.Perm
	}
	if myBeego.Separate != nil {
		separate = myBeego.Separate
	}
	mm := &MyBeegoLogAdapterMultiFile{
		MyBeegoLogAdapterFile: MyBeegoLogAdapterFile{
			FileName: filename,
			MaxLines: maxlines,
			MaxSize:  maxsize,
			Daily:    myBeego.Daily,
			MaxDays:  maxdays,
			Rotate:   myBeego.Rotate,
			Level:    level,
			Perm:     perm,
		},
		Separate: separate,
	}
	return mm
}
