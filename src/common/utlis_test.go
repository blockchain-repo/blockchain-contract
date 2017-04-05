package common

import(
	"testing"
	"fmt"
)

func Test_GenDate(t *testing.T) {
	dat := GenDate()
	fmt.Println("----------------------dat:", dat)
}

func Test_GenTimestamp(t *testing.T) {
        tim := GenTimestamp()
        fmt.Println("----------------------tim:", tim)
}

func Test_Deserialize(t *testing.T) {
	jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
        data := Deserialize(jsonStr)
        fmt.Println("----------------------des:", data)
}

func Test_Serialize(t *testing.T) {
        jsonStr := `{"host": "http://localhost:9090","port": 9090,"analytics_file": "","static_file_version": 1,"static_dir": "E:/Project/goTest/src/","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits": ["apple", "peach"]}`
        data := Deserialize(jsonStr)
	data2 := Serialize(data) 
        fmt.Println("----------------------ser:", data2)
}
