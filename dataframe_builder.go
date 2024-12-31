package fuse

import (
	"context"
	"fmt"
)

type dataFrameTransformFunc func(
	ctx context.Context,
	df *DataFrame,
) (*DataFrame, error)

// DataFrameBuilder is an error-safe construct to connect a series of
// DataFrame transformations together. Create one of these from an existing
// DataFrame by calling its Builder() method.
type DataFrameBuilder struct {
	origDF *DataFrame
	funcs  []dataFrameTransformFunc
}

// Select appends a call to df.Select(ctx, cols) to the end of the set of
// transforms.
func (d DataFrameBuilder) Select(cols ...Column) DataFrameBuilder {
	d.funcs = append(
		d.funcs,
		func(ctx context.Context, df *DataFrame) (*DataFrame, error) {
			return df.Select(ctx, cols...)
		},
	)

	return d
}

// Execute executes all transforms in d and returns the resulting DataFrame.
// If any transform fails, execution is immediately halted and the error
// is returned.
func (d DataFrameBuilder) Execute(ctx context.Context) (*DataFrame, error) {
	res := d.origDF
	for i, fn := range d.funcs {
		nextDF, err := fn(ctx, res)
		if err != nil {
			return nil, fmt.Errorf(
				"executing step %d in DataFrameBuilder: %w",
				i,
				err,
			)
		}

		res = nextDF
	}

	return res, nil
}

// Collect is a convenience function for roughly the following:
//
//	df, err := d.Execute(ctx)
//	if err != nil {
//		return nil, fmt.Errorf("collecting: %w", err)
//	}
//	return df.Collect(ctx)
func (d DataFrameBuilder) Collect(ctx context.Context) ([]*Row, error) {
	df, err := d.Execute(ctx)
	if err != nil {
		return nil, fmt.Errorf("collecting: %w", err)
	}

	return df.Collect(ctx)
}
