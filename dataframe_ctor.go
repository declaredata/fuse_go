package fuse

import (
	"context"
	"fmt"

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
	dfID, err := sess.client.LoadCSV(ctx, &gen.LoadFileRequest{
		SessionId: sess.uid.String(),
		Source:    fileName,
	})
	if err != nil {
		return nil, fmt.Errorf(
			"loading CSV %s for session %s: %w",
			fileName,
			sess.uid.String(),
			err,
		)
	}

	return dfUIDToDF(dfID, sess)
}

func dfUIDToDF(dfUID *gen.DataFrameUID, sess *Session) (*DataFrame, error) {
	parsed, err := uuid.Parse(dfUID.GetDataframeUid())
	if err != nil {
		return nil, fmt.Errorf(
			"parsing raw DataFrameUID %s: %w",
			dfUID.GetDataframeUid(),
			err,
		)
	}

	return newDataFrame(parsed, sess), nil
}
