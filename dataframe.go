package fuse

import (
	"context"
	"fmt"

	gen "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
)

type DataFrame struct {
	dfUID uuid.UUID
	sess  *Session
}

func (d *DataFrame) Collect(ctx context.Context) ([]*Row, error) {
	contents, err := d.sess.client.Collect(ctx, &gen.DataFrameUID{
		DataframeUid: d.dfUID.String(),
	})
	if err != nil {
		return nil, fmt.Errorf("collecting dataframe: %w", err)
	}

	return mapSlice(contents.GetRows(), func(r *gen.Row) *Row {
		return &Row{r: r}
	}), nil
}

func (d *DataFrame) Col(colName string) *ExistingCol {
	return Col(colName)
}

func (d *DataFrame) Select(ctx context.Context, cols ...Column) (*DataFrame, error) {
	resp, err := d.sess.client.Select(ctx, &gen.SelectRequest{
		DfUid: d.dfUID.String(),
		Columns: mapSlice(cols, func(c Column) *gen.Column {
			return c.toColumn()
		}),
	})
	if err != nil {
		return nil, fmt.Errorf(
			"calling select on %d columns: %w",
			len(cols),
			err,
		)
	}

	return d.createDescendant(resp)
}

func (d *DataFrame) Builder() DataFrameBuilder {
	return DataFrameBuilder{origDF: d, funcs: nil}
}

func (d *DataFrame) createDescendant(dfUID *gen.DataFrameUID) (*DataFrame, error) {
	return dfUIDToDF(dfUID, d.sess)
}
