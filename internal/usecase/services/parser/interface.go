package parser

import "context"

type Interface interface {
	Parse(ctx context.Context, urls []string) ([][]byte, error)
	pullId(url string) string
}
