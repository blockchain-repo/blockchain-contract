package pipelines

import (
//	"time"
//	"fmt"
	"os"
	"io"
	"log"
	"bufio"
	"bytes"
	"encoding/json"
	r "unicontract/src/core/db/rethinkdb"
)

//TODO table UPDATE? mutil
func ceChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("Unicontract", "Contract")
	for res.Next(&value) {
		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			log.Fatalf(err.Error())
		}
		out.Write(v)
	}
}

func ceHeadFilter(in io.Reader, out io.Writer) {
    rd := bufio.NewReader(in)
    p := make([]byte,MaxSizeTX)
    for {
        n, _ := rd.Read(p)
        if n == 0 {
            break
        }
        t := bytes.ToLower(p[:n])
        out.Write(t)
    }
}

func ceQueryEists(in io.Reader, out io.Writer) {
    rd := bufio.NewReader(in)
    p := make([]byte,MaxSizeTX)
    for {
        n, _ := rd.Read(p)
        if n == 0 {
            break
        }
        t := bytes.ToLower(p[:n])
        out.Write(t)
    }
}

func ceSend(in io.Reader, out io.Writer) {
    rd := bufio.NewReader(in)
    p := make([]byte,MaxSizeTX)
    for {
        n, _ := rd.Read(p)
        if n == 0 {
            break
        }
        t := bytes.ToLower(p[:n])
        out.Write(t)
    }


func startContractElection() {

	p:=Pipe(
	ceChangefeed,
	ceHeadFilter,
	ceQueryEists,
	ceSend)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf(err.Error())
	w := bufio.NewWriter(f)
	writeVote)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil,w)
}
