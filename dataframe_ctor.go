package fuse

import (
	"context"

	gen "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
)

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

func dfUIDToDF(dfUID *gen.DataFrameUID, sess *Session) (*DataFrame, error) {
	parsed, err := uuid.Parse(dfUID.DataframeUid)
	if err != nil {
		return nil, err
	}
	return newDataFrame(parsed, sess), nil
}
