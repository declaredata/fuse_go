package fuse

import fusegen "github.com/declaredata/fuse_go/gen"

type Column interface {
	// Name returns the current name of this column
	Name() string
	// toColumn converts this Column to a gRPC-compatible Column. it is
	// purposefully not exported. clients of the fuse_go library must not
	// implement Column
	toColumn() *fusegen.Column
}

// ExistingCol is a Column implementation that simply references a column
// that already exists in a DataFrame.
//
// Call either the Col or DataFrame.Col functions to create one of these.
type ExistingCol struct {
	name string
}

var _ Column = &ExistingCol{name: ""}

func Col(name string) *ExistingCol {
	return &ExistingCol{name: name}
}

func (e *ExistingCol) Name() string {
	return e.name
}

func (e *ExistingCol) toColumn() *fusegen.Column {
	return &fusegen.Column{
		ColSpec: &fusegen.Column_ColName{ColName: e.name},
		Window:  nil,
	}
}
