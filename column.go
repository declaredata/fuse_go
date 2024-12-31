package fuse

import gen "github.com/declaredata/fuse_go/gen"

type Column interface {
	// Name returns the current name of this column
	Name() string
	// toColumn converts this Column to a gRPC-compatible Column. it is
	// purposefully not exported. clients of the fuse_go library must not
	// implement Column
	toColumn() *gen.Column
}

type existingCol struct {
	name string
}

func (e *existingCol) Name() string {
	return e.name
}

func (e existingCol) toColumn() *gen.Column {
	return &gen.Column{
		ColSpec: &gen.Column_ColName{ColName: e.name},
		Window:  nil,
	}
}
