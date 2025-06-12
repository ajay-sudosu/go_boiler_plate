package lib_adapter

import (
	"context"
)

type NetworkAdapter interface {
	CreateNetwork(ctx context.Context) error
	DeleteNetwork(ctx context.Context) error
}
