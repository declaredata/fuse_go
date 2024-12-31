package fuse

import (
	"context"
	"fmt"

	gen "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
)

type Session struct {
	client *Client
	uid    uuid.UUID
}

func (s *Session) Close(ctx context.Context) error {
	_, err := s.client.CloseSession(
		ctx,
		&gen.SessionUID{SessionUid: s.uid.String()},
	)

	return fmt.Errorf("closing session: %w", err)
}
