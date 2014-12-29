package pbd

import (
	"fmt"
	"sync"
)

// PBD
type PBD struct {
	Atoms              []*Atom
	Authors            []*Author
	Compound           Compound
	Connections        []*Connection
	Crystals           []*Crystal
	DatabaseReferences []*DatabaseReference
	ExperimentalData   []*ExperimentalData
	Formulas           []*Formula
	Headers            []*Header
	Helixes            []*Helix
	HeteroAtoms        []*HeteroAtom
	HeteroAtomNames    []*HeteroAtomName
	JournalEntries     []*Journal
	Keywords           []*Keywords
	Links              []*Link
	Master             []*Master
	ModifiedResidues   []*ModifiedResidue
	Origins            []*Orign
	Remarks            []*Remark
	ResidueSequences   []*ResidueSequence
	Revisions          []*Revision
	Scales             []*Scale
	Sheets             []*Sheet
	Sites              []*Site
	Source             []*Source
	SSBonds            []*SSBond
	Title              Title
	*sync.Mutex
}

func NewPBD() *PBD {
	return &PBD{}
}

// TODO add full structures for each of these
type Atom struct {
	Serial                int
	AtomName              string
	AltLocation           string
	ResidueName           string
	ChainIdentifier       string
	ResidueSequenceNumber int
	InsertionCode         int
	X, Y, Z               float64
	Occupancy             float64
	TemperatureFactor     float64
	Segment               string
	Element               string
	Charge                string
}

func (a *Atom) String() string {
	return fmt.Sprintf("%+v", *a)

}

type Author string

func (a *Author) String() string {
	return fmt.Sprint(*a)
}

type Compound struct {
	ID    int
	Name  string
	Chain []string
}

func (c *Compound) String() string {
	return fmt.Sprint(*c)
}

type Connection struct {
	BaseAtom       int
	ConnectedAtoms []int
}

func (c *Connection) String() string {
	return fmt.Sprint(*c)
}

type Crystal string
type DatabaseReference string
type ExperimentalData string
type Formula string
type Header string
type Helix string
type HeteroAtom string
type HeteroAtomName string
type Journal string
type Keywords string
type Link string
type Master string
type ModifiedResidue string
type Orign string
type Remark string
type ResidueSequence string
type Revision string
type Scale string
type Sheet string
type Site string
type Source string
type SSBond string
type Title string
