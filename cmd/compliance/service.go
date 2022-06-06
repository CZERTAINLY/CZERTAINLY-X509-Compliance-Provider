package compliance

import (
	"context"
)

type Service interface {
	ComplianceCheck(ctx context.Context, kind string, request Request) (Response, error)
}
