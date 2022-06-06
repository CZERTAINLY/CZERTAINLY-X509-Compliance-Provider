package rules

import (
	"context"
)

type Service interface {
	GetRules(ctx context.Context, kind string, certificateType []string) ([]Response, error)
	GetGroups(ctx context.Context, kind string) ([]GroupResponse, error)
	GetGroupDetails(ctx context.Context, uuid string, kind string) ([]Response, error)
}
