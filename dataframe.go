package fuse

import (
	"context"

	gen "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
)

type DataFrame struct {
	df_uid uuid.UUID
	sess   *Session
}

func newDataFrame(uid uuid.UUID, sess *Session) *DataFrame {
	return &DataFrame{
		df_uid: uid,
		sess:   sess,
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
		DataframeUid: d.df_uid.String(),
	})
	if err != nil {
		return nil, err
	}
	return mapSlice(contents.Rows, func(r *gen.Row) *Row {
		return &Row{r: r}
	}), nil

}

func (d *DataFrame) createDescendant(dfUID *gen.DataFrameUID) (*DataFrame, error) {
	return dfUIDToDF(dfUID, d.sess)
}

func dfUIDToDF(dfUID *gen.DataFrameUID, sess *Session) (*DataFrame, error) {
	parsed, err := uuid.Parse(dfUID.DataframeUid)
	if err != nil {
		return nil, err
	}
	return &DataFrame{df_uid: parsed, sess: sess}, nil
}
