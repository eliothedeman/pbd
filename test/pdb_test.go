package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"testing"

	"github.com/eliothedeman/pbd"
)

func TestParse(t *testing.T) {
	p, err := pbd.ParsePBD(buff)
	if err != nil {
		t.Error(err)
	}

	b, jErr := json.Marshal(p.Atoms)
	if jErr != nil {
		t.Fatal(err)
	}
	fmt.Print(string(b))

}

func BenchmarkParse(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		p, err := pbd.ParsePBD(buff)
		if err != nil {
			b.Error(err)
		}
		if p == nil {
			b.Fail()
		}
	}
}

var (
	buff = func() *bytes.Buffer {
		b, err := ioutil.ReadFile("test.pbd")
		if err != nil {
			log.Fatal(err)
		}
		return bytes.NewBuffer(b)
	}()
)
