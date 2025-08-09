package urlprocessor

import "context"

//go:generate mockery --name=UseCase --output=./mocks
type UseCase interface {
	CleanURL(ctx context.Context, operation, url string) (string, error)
}
