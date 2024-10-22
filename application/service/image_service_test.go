package service_test

import (
	"context"
	"testing"

	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/stretchr/testify/mock"
)

type MockImageRepository struct {
	mock.Mock
}

func (m *MockImageRepository) Save(ctx context.Context, image entity.Image) error {
	args := m.Called(ctx, image)
	return args.Error(0)
}

func TestCreateImage(t *testing.T) {

}
