package gopdb

// PDB
type PDB struct {
	Atoms              []*Atom
	Authors            []*Author
	Compounts          []*Compound
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
}

// TODO add full structures for each of these
type Atom string
type Author string
type Compound string
type Connection string
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
