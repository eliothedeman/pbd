package pbd

import (
	"bufio"
	"errors"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	lineParsers = map[string]lineParser{
		"ATOM":   parseAtom,
		"AUTHOR": parseAuthor,
		"COMPND": parseCompound,
		"CONECT": parseConnection,
		"CRYST1": parseCrystal,
	}
	whiteSpace         = regexp.MustCompile(`\s+`)
	INSUFFICENT_LENGTH = errors.New("pbd: line is is of incorrect length")
)

// ParsePBD read each line of a pbd file and return the PBD object it creates
func ParsePBD(r io.Reader) (*PBD, error) {
	scanner := bufio.NewScanner(r)
	pbd := &PBD{}
	var err error
	for scanner.Scan() {
		line := scanner.Text()
		f := selectParser(line)

		if f != nil {
			err = f(line, pbd)
			if err != nil {
				log.Println(err)
			}
		} else {
			// log.Println("could not find parser for ", line)
		}
	}
	return pbd, nil
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

func parseFloat(s string) (float64, error) {
	s = strings.TrimSpace(s)
	f, err := strconv.ParseFloat(s, 64)
	return f, err
}

func parseInt(s string) (int, error) {
	s = strings.TrimSpace(s)
	i, err := strconv.Atoi(s)
	return i, err
}

// lineParser parsers a line of text, and assigns it to the correct
type lineParser func(line string, pbd *PBD) error

func parseAtom(line string, p *PBD) error {
	if len(line) < 80 {
		return INSUFFICENT_LENGTH
	}
	var err error
	a := &Atom{}
	a.Serial, _ = strconv.Atoi(strings.TrimSpace(line[7:11]))
	a.AtomName = strings.TrimSpace(line[13:16])
	a.AltLocation = strings.TrimSpace(line[17:18])
	a.ResidueName = strings.TrimSpace(line[18:20])
	a.ChainIdentifier = line[22:23]
	a.ResidueSequenceNumber, err = parseInt(line[23:26])
	a.InsertionCode, err = parseInt(line[27:28])
	a.X, err = parseFloat(line[31:38])
	a.Y, err = parseFloat(line[39:46])
	a.Z, err = parseFloat(line[47:54])
	a.Occupancy, err = parseFloat(line[55:60])
	a.TemperatureFactor, err = parseFloat(line[61:66])
	a.Segment = strings.TrimSpace(line[73:76])
	a.Element = strings.TrimSpace(line[77:78])
	a.Charge = strings.TrimSpace(line[79:80])
	p.Atoms = append(p.Atoms, a)
	return err
}

func parseAuthor(line string, p *PBD) error {
	line = strings.TrimSpace(line[6:])
	s := strings.Split(line, ",")
	for i := range s {
		a := Author(s[i])
		p.Authors = append(p.Authors, &a)
	}
	return nil
}

func parseCompound(line string, p *PBD) error {
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

	return err
}

// TODO at the end of the parsing, link all of the actual atom structures to eachother with pointers
func parseConnection(line string, p *PBD) error {
	c := &Connection{}
	var err error
	c.BaseAtom, err = strconv.Atoi(strings.TrimSpace(line[7:11]))
	s := whiteSpace.Split(strings.TrimSpace(line[11:80]), 8)
	if err != nil {
		return err
	}
	var a int
	for i := range s {
		a, err = strconv.Atoi(strings.TrimSpace(s[i]))
		if err != nil {
			return err
		}
		c.ConnectedAtoms = append(c.ConnectedAtoms, a)
	}
	p.Connections = append(p.Connections, c)
	return nil
}

func parseCrystal(line string, p *PBD) error {
	c := Crystal{}
	var err error

	c.A, err = parseFloat(line[7:15])
	c.B, err = parseFloat(line[16:24])
	c.C, err = parseFloat(line[25:33])
	c.Alpha, err = parseFloat(line[34:40])
	c.Beta, err = parseFloat(line[41:47])
	c.Gama, err = parseFloat(line[48:54])
	c.SGroup = strings.TrimSpace(line[56:66])
	c.ZVal, err = parseInt(line[67:70])

	p.Crystal = c
	return err
}
