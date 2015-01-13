package pbd

import "fmt"

// PBD
type PBD struct {
	Atoms              []*Atom
	Authors            []*Author
	Compound           Compound
	Connections        []*Connection
	Crystal            Crystal
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
}

func NewPBD() *PBD {
	return &PBD{}
}

// TODO add full structures for each of these
type Atom struct {
	Serial                int     `json:"serial"`
	AtomName              string  `json:"atom_name"`
	AltLocation           string  `json:"alternat_location"`
	ResidueName           string  `json:"residue_name"`
	ChainIdentifier       string  `json:"chain_identifier"`
	ResidueSequenceNumber int     `json:"residue_sequence_number"`
	InsertionCode         int     `json:"insertion_code"`
	X                     float64 `json:"x"`
	Y                     float64 `json:"y"`
	Z                     float64 `json:"z"`
	Occupancy             float64 `json:"occupancy"`
	TemperatureFactor     float64 `json:"temperature_factor"`
	Segment               string  `json:"segment"`
	Element               string  `json:"element"`
	Charge                string  `json:"charge"`
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

type Crystal struct {
	A, B, C           float64
	Alpha, Beta, Gama float64
	SGroup            string
	ZVal              int
}

type DatabaseReference struct {
	ID                          string
	ChainID                     string
	SequenceBegin               int
	InsertBegin                 string
	SequenceEnd                 int
	InsertEnd                   string
	Database                    string
	DBAccession                 string
	DBIDCode                    string
	DBSeqBegin                  int
	InitialResidueInsertionCode string
	DBSeqEnd                    int
	DBInsertionEnd              string
}
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
