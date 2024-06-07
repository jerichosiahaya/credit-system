package usecase

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"kredit-plus/src/domain"
	mock_domain "kredit-plus/src/mock"

	"github.com/stretchr/testify/assert"
)

func TestCreateCustomer(t *testing.T) {
	ctrl := gomock.NewController(t)
	ctx := context.Background()

	mockDatabase := mock_domain.NewMockDatabase(ctrl)

	u := New(mockDatabase)

	t.Run("success", func(t *testing.T) {
		mockDatabase.EXPECT().InsertCustomer(ctx, domain.Customer{}).Return(domain.Customer{}, nil)
		_, err := u.CreateCustomer(ctx, domain.Customer{})
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockDatabase.EXPECT().InsertCustomer(ctx, domain.Customer{}).Return(nil, errors.New("error"))
		_, err := u.CreateCustomer(ctx, domain.Customer{})
		assert.Error(t, err)
	})

}