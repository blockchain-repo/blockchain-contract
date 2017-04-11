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

const MaxSizeTX = 16*1024

func changefeed(in io.Reader, out io.Writer) {
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

//TODO: core validate func
func validateContract(in io.Reader, out io.Writer) {
    rd := bufio.NewReader(in)
    p := make([]byte,MaxSizeTX)
    for {
        n, _ := rd.Read(p)
        if n == 0 {
            break
        }
        t := bytes.ToUpper(p[:n])
        out.Write(t)
    }
}

//TODO: core make vote
func vote(in io.Reader, out io.Writer) {
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


//TODO:core write vote
func writeVote(in io.Reader, out io.Writer) {
    rd := bufio.NewReader(in)
    p := make([]byte,MaxSizeTX)
    for {
        n, _ := rd.Read(p)
        if n == 0 {
            break
        }
        t := p[:n]
	r.Insert("Unicontract","Votes",string(t))
	out.Write(t)
    }
}

func startContractVote() {
	p:=Pipe(
	changefeed,
	validateContract,
	vote,
	writeVote)

	p(os.Stdin, os.Stdout)
}
