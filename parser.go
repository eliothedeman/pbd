package pbd

import (
	"bufio"
	"errors"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	lineParsers = map[string]lineParser{
		"ATOM":   parseAtom,
		"AUTHOR": parseAuthor,
		"COMPND": parseCompound,
	}
	whiteSpace         = regexp.MustCompile(`\s+`)
	INSUFFICENT_LENGTH = errors.New("pdb: line is is of incorrect length")
)

// ParsePDB read each line of a pdb file and return the PDB object it creates
func ParsePDB(r io.Reader) (*PDB, error) {
	scanner := bufio.NewScanner(r)
	pdb := &PDB{
		Mutex: &sync.Mutex{},
	}
	var err error
	for scanner.Scan() {
		line := scanner.Text()
		f := selectParser(line)

		if f != nil {
			err = f(line, pdb)
			if err != nil {
				log.Println(err)
			}
		} else {
			// log.Println("could not find parser for ", line)
		}
	}
	return pdb, nil
}

func selectParser(line string) lineParser {
	line = whiteSpace.Split(line, 2)[0]
	line = strings.ToUpper(line)

	f, ok := lineParsers[line]
	if !ok {
		return nil
	}
	return f
}

func parseFloat(s string) float64 {
	s = strings.TrimSpace(s)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println(err)
		panic(err)
		return float64(0)
	}
	return f
}

// lineParser parsers a line of text, and assigns it to the correct
type lineParser func(line string, pdb *PDB) error

func parseAtom(line string, p *PDB) error {
	if len(line) < 80 {
		return INSUFFICENT_LENGTH
	}
	a := &Atom{}
	a.Serial, _ = strconv.Atoi(strings.TrimSpace(line[0:7]))
	a.AtomName = strings.TrimSpace(line[6:12])
	a.AltLocation = strings.TrimSpace(line[11:14])
	a.ResidueName = strings.TrimSpace(line[13:16])
	a.ChainIdentifier = strings.TrimSpace(string(line[17]))
	a.ResidueSequenceNumber, _ = strconv.Atoi(strings.TrimSpace(line[18:20]))
	a.InsertionCode, _ = strconv.Atoi(strings.TrimSpace(string(line[22])))
	a.X = parseFloat(line[31:38])
	a.Y = parseFloat(line[39:46])
	a.Z = parseFloat(line[47:54])
	a.Occupancy = parseFloat(line[55:60])
	a.TemperatureFactor = parseFloat(line[61:66])
	a.Segment = strings.TrimSpace(line[73:76])
	a.Element = strings.TrimSpace(line[77:78])
	a.Charge = strings.TrimSpace(line[79:80])
	p.Lock()
	p.Atoms = append(p.Atoms, a)
	p.Unlock()
	return nil
}

func parseAuthor(line string, p *PDB) error {
	p.Lock()
	line = strings.TrimSpace(line[6:])
	s := strings.Split(line, ",")
	for i := range s {
		a := Author(s[i])
		p.Authors = append(p.Authors, &a)
	}
	p.Unlock()
	return nil
}

func parseCompound(line string, p *PDB) error {
	p.Lock()
	// strop prefix
	line = strings.TrimSpace(line[10:])
	s := whiteSpace.Split(line, 2)
	var err error
	switch s[0] {
	case "MOL_ID:":
		p.Compound.ID, err = strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(s[1], ";")))
	case "MOLECULE:":
		p.Compound.Name = strings.TrimSpace(strings.TrimSuffix(s[1], ";"))
	case "CHAIN:":
		chain := strings.Split(s[1], ",")
		for i := range chain {
			p.Compound.Chain = append(p.Compound.Chain, strings.TrimSpace(chain[i]))
		}
	}
	p.Unlock()

	return err
}
