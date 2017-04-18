package pipelines

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"

	r "unicontract/src/core/db/rethinkdb"
)

func txeChangefeed(in io.Reader, out io.Writer) {
	var value interface{}
	//TODO table name
	res := r.Changefeed("Unicontract", "Contract")
	for res.Next(&value) {
		m := value.(map[string]interface{})
		v, err := json.Marshal(m["new_val"])
		if err != nil {
			log.Fatalf(err.Error())
			continue
		}
		if bytes.Equal(v, []byte("null")) {
			continue
		}
		out.Write(v)
	}
}

func txeHeadFilter(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		//TODO head filter
		out.Write(t)
	}
}

func txeQueryEists(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		//TODO query
		out.Write(t)
	}
}

func txeSend(in io.Reader, out io.Writer) {
	rd := bufio.NewReader(in)
	p := make([]byte, MaxSizeTX)
	for {
		n, _ := rd.Read(p)
		if n == 0 {
			continue
		}
		t := p[:n]
		//TODO send
		out.Write(t)
	}
}

func starttxElection() {

	p := Pipe(
		txeChangefeed,
		txeHeadFilter,
		txeQueryEists,
		txeSend)

	f, err := os.OpenFile("/dev/null", os.O_RDWR, 0)
	if err != nil {
		log.Fatalf(err.Error())
	}
	w := bufio.NewWriter(f)
	p(nil, w)
}
