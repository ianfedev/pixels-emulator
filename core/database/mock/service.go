package database_test

import (
	"context"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// ModelServiceMock is a mock of the IDataService interface.
type ModelServiceMock[T any] struct {
	mock.Mock
}

// CreateSync mocks the CreateSync method.
func (m *ModelServiceMock[T]) CreateSync(ctx context.Context, entity *T) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

// Create mocks the Create method.
func (m *ModelServiceMock[T]) Create(ctx context.Context, entity *T) <-chan error {
	args := m.Called(ctx, entity)
	return args.Get(0).(<-chan error)
}

// GetSync mocks the GetSync method.
func (m *ModelServiceMock[T]) GetSync(ctx context.Context, id uint) (*T, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*T), args.Error(1)
}

// Get mocks the Get method.
func (m *ModelServiceMock[T]) Get(ctx context.Context, id uint) <-chan struct {
	Data  *T
	Error error
} {
	args := m.Called(ctx, id)
	return args.Get(0).(<-chan struct {
		Data  *T
		Error error
	})
}

// UpdateSync mocks the UpdateSync method.
func (m *ModelServiceMock[T]) UpdateSync(ctx context.Context, entity *T) error {
	args := m.Called(ctx, entity)
	return args.Error(0)
}

// Update mocks the Update method.
func (m *ModelServiceMock[T]) Update(ctx context.Context, entity *T) <-chan error {
	args := m.Called(ctx, entity)
	return args.Get(0).(<-chan error)
}

// DeleteSync mocks the DeleteSync method.
func (m *ModelServiceMock[T]) DeleteSync(ctx context.Context, id uint) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// Delete mocks the Delete method.
func (m *ModelServiceMock[T]) Delete(ctx context.Context, id uint) <-chan error {
	args := m.Called(ctx, id)
	return args.Get(0).(<-chan error)
}

// FindByQuerySync mocks the FindByQuerySync method.
func (m *ModelServiceMock[T]) FindByQuerySync(ctx context.Context, query map[string]interface{}) ([]T, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]T), args.Error(1)
}

// FindByQuery mocks the FindByQuery method.
func (m *ModelServiceMock[T]) FindByQuery(ctx context.Context, query map[string]interface{}) <-chan struct {
	Data  []T
	Error error
} {
	args := m.Called(ctx, query)
	return args.Get(0).(<-chan struct {
		Data  []T
		Error error
	})
}

// BeginTransactionSync mocks the BeginTransactionSync method.
func (m *ModelServiceMock[T]) BeginTransactionSync(ctx context.Context) (*gorm.DB, error) {
	args := m.Called(ctx)
	return args.Get(0).(*gorm.DB), args.Error(1)
}

// BeginTransaction mocks the BeginTransaction method.
func (m *ModelServiceMock[T]) BeginTransaction(ctx context.Context) <-chan struct {
	Tx    *gorm.DB
	Error error
} {
	args := m.Called(ctx)
	return args.Get(0).(<-chan struct {
		Tx    *gorm.DB
		Error error
	})
}

// CommitTransactionSync mocks the CommitTransactionSync method.
func (m *ModelServiceMock[T]) CommitTransactionSync(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

// CommitTransaction mocks the CommitTransaction method.
func (m *ModelServiceMock[T]) CommitTransaction(tx *gorm.DB) <-chan error {
	args := m.Called(tx)
	return args.Get(0).(<-chan error)
}

// RollbackTransactionSync mocks the RollbackTransactionSync method.
func (m *ModelServiceMock[T]) RollbackTransactionSync(tx *gorm.DB) error {
	args := m.Called(tx)
	return args.Error(0)
}

// RollbackTransaction mocks the RollbackTransaction method.
func (m *ModelServiceMock[T]) RollbackTransaction(tx *gorm.DB) <-chan error {
	args := m.Called(tx)
	return args.Get(0).(<-chan error)
}
