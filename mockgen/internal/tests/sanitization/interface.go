package sanitization

import (
	"github.com/nigoroll/mock/mockgen/internal/tests/sanitization/any"
)

//go:generate mockgen -destination mockout/mock.go -package mockout . AnyMock

type AnyMock interface {
	Do(a *any.Any, b int)
}
