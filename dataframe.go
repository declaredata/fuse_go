package fuse

import (
	"context"

	gen "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
)

type DataFrame struct {
	dfUID uuid.UUID
	sess  *Session
}

func newDataFrame(uid uuid.UUID, sess *Session) *DataFrame {
	return &DataFrame{
		dfUID: uid,
		sess:  sess,
	}
}

func DataFrameFromCSV(ctx context.Context, sess *Session, fileName string) (*DataFrame, error) {
	df_id, err := sess.client.LoadCSV(ctx, &gen.LoadFileRequest{
		SessionId: sess.uid.String(),
		Source:    fileName,
	})
	if err != nil {
		return nil, err
	}
	return dfUIDToDF(df_id, sess)
}

func (d *DataFrame) Collect(ctx context.Context) ([]*Row, error) {
	contents, err := d.sess.client.Collect(ctx, &gen.DataFrameUID{
		DataframeUid: d.dfUID.String(),
	})
	if err != nil {
		return nil, err
	}
	return mapSlice(contents.Rows, func(r *gen.Row) *Row {
		return &Row{r: r}
	}), nil
}

func (d *DataFrame) Select(ctx context.Context, cols ...Column) (*DataFrame, error) {
	resp, err := d.sess.client.Select(ctx, &gen.SelectRequest{
		DfUid: d.dfUID.String(),
		Columns: mapSlice(cols, func(c Column) *gen.Column {
			return c.toColumn()
		}),
	})
	if err != nil {
		return nil, err
	}
	return d.createDescendant(resp)
}

func (d *DataFrame) createDescendant(dfUID *gen.DataFrameUID) (*DataFrame, error) {
	return dfUIDToDF(dfUID, d.sess)
}

func dfUIDToDF(dfUID *gen.DataFrameUID, sess *Session) (*DataFrame, error) {
	parsed, err := uuid.Parse(dfUID.DataframeUid)
	if err != nil {
		return nil, err
	}
	return &DataFrame{dfUID: parsed, sess: sess}, nil
}
