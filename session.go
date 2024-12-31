package fuse_client

import (
	"context"

	fuse "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
)

type Session struct {
	client *Client
	uid    uuid.UUID
}

func (s *Session) Close(ctx context.Context) error {
	_, err := s.client.CloseSession(
		ctx,
		&fuse.SessionUID{SessionUid: s.uid.String()},
	)
	return err
}
