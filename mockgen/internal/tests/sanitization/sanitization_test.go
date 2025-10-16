package sanitization

import (
	"testing"

	"github.com/nigoroll/mock/gomock"
	any0 "github.com/nigoroll/mock/mockgen/internal/tests/sanitization/any"
	"github.com/nigoroll/mock/mockgen/internal/tests/sanitization/mockout"
)

func TestSanitization(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mockout.NewMockAnyMock(ctrl)
	m.EXPECT().Do(gomock.Any(), 1)
	m.Do(&any0.Any{}, 1)
}
