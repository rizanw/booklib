package urlprocessor

import "context"

type UseCase interface {
	CleanURL(ctx context.Context, operation, url string) (string, error)
}
