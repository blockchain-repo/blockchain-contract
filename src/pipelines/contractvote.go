package pipelines

import (
	"io"
	"os"
	"log"
	"bufio"
	"bytes"
	"encoding/json"
	r "unicontract/src/core/db/rethinkdb"
)

//TODO
func changefeed(in io.Reader, out io.Writer) {
	var value interface{}
	res := r.Changefeed("bigchain", "backlog")
	for res.Next(&value) {
		v, err := json.Marshal(value)
		if err != nil {
			log.Fatalf(err.Error())
		}
		out.Write(v)
	}
}

func validateContract(in io.Reader, out io.Writer) {
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

func vote(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte,10)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			break
		}
		t := bytes.ToLower(p[:n])
		out.Write(t)
	}
}

func writeVote(in io.Reader, out io.Writer) {
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

func startContractVote() {
	p := Pipe(
	changefeed,
	validateContract,
	vote,
	writeVote)

	p(os.Stdin, os.Stdout)
}
