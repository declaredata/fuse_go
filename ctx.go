package fuse

import (
	"context"
	"fmt"

	fuse "github.com/declaredata/fuse_go/gen"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	fuse.SdsClient
}

func NewClient(ctx context.Context, host string, port int) (*Client, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	cl := fuse.NewSdsClient(conn)
	return &Client{SdsClient: cl}, nil
}

func (c *Client) NewSession(ctx context.Context) (*Session, error) {
	sess, err := c.CreateSession(ctx, &fuse.Empty{})
	if err != nil {
		return nil, err
	}
	uid, err := uuid.Parse(sess.SessionUid)
	if err != nil {
		return nil, err
	}

	return &Session{client: c, uid: uid}, nil
}
