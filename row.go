package fuse

import gen "github.com/declaredata/fuse_go/gen"

type Row struct {
	r *gen.Row
}

func (r *Row) Data() map[string]*gen.Value {
	return r.r.Data
}
