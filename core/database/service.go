package database

import "gorm.io/gorm"

// DataService defines a generic interface for CRUD operations and transactions.
type DataService[T any] interface {

	// CreateSync adds a new record to the database synchronously.
	CreateSync(entity *T) error

	// Create adds a new record to the database asynchronously.
	Create(entity *T) <-chan error

	// GetSync retrieves a record by its primary key synchronously.
	GetSync(id uint) (*T, error)

	// Get retrieves a record by its primary key asynchronously.
	Get(id uint) <-chan struct {
		Data  *T
		Error error
	}

	// UpdateSync updates an existing record synchronously.
	UpdateSync(entity *T) error

	// Update updates an existing record asynchronously.
	Update(entity *T) <-chan error

	// DeleteSync performs a soft delete on a record synchronously.
	DeleteSync(id uint) error

	// Delete performs a soft delete on a record asynchronously.
	Delete(id uint) <-chan error

	// FindByQuerySync retrieves records based on a query synchronously.
	FindByQuerySync(query map[string]interface{}) ([]T, error)

	// FindByQuery retrieves records based on a query asynchronously.
	FindByQuery(query map[string]interface{}) <-chan struct {
		Data  []T
		Error error
	}

	// BeginTransactionSync starts a new database transaction synchronously.
	BeginTransactionSync() (*gorm.DB, error)

	// BeginTransaction starts a new database transaction asynchronously.
	BeginTransaction() <-chan struct {
		Tx    *gorm.DB
		Error error
	}

	// CommitTransactionSync commits the transaction synchronously.
	CommitTransactionSync(tx *gorm.DB) error

	// CommitTransaction commits the transaction asynchronously.
	CommitTransaction(tx *gorm.DB) <-chan error

	// RollbackTransactionSync rolls back the transaction synchronously.
	RollbackTransactionSync(tx *gorm.DB) error

	// RollbackTransaction rolls back the transaction asynchronously.
	RollbackTransaction(tx *gorm.DB) <-chan error
}

// ModelService provides CRUD operations for any GORM model.
type ModelService[T any] struct {
	DB *gorm.DB
}

// CreateSync adds a new record to the database (synchronous).
func (s *ModelService[T]) CreateSync(entity *T) error {
	return s.DB.Create(entity).Error
}

// Create adds a new record to the database (asynchronous).
func (s *ModelService[T]) Create(entity *T) <-chan error {
	result := make(chan error, 1)
	go func() {
		defer close(result)
		result <- s.DB.Create(entity).Error
	}()
	return result
}

// GetSync retrieves a record by its primary key (synchronous).
func (s *ModelService[T]) GetSync(id uint) (*T, error) {
	entity := new(T)
	if err := s.DB.First(entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// Get retrieves a record by its primary key (asynchronous).
func (s *ModelService[T]) Get(id uint) <-chan struct {
	Data  *T
	Error error
} {
	result := make(chan struct {
		Data  *T
		Error error
	}, 1)
	go func() {
		defer close(result)
		entity := new(T)
		if err := s.DB.First(entity, id).Error; err != nil {
			result <- struct {
				Data  *T
				Error error
			}{nil, err}
			return
		}
		result <- struct {
			Data  *T
			Error error
		}{entity, nil}
	}()
	return result
}

// UpdateSync modifies an existing record in the database (synchronous).
func (s *ModelService[T]) UpdateSync(entity *T) error {
	return s.DB.Save(entity).Error
}

// Update modifies an existing record in the database (asynchronous).
func (s *ModelService[T]) Update(entity *T) <-chan error {
	result := make(chan error, 1)
	go func() {
		defer close(result)
		result <- s.DB.Save(entity).Error
	}()
	return result
}

// DeleteSync performs a soft delete on a record (synchronous).
func (s *ModelService[T]) DeleteSync(id uint) error {
	return s.DB.Delete(new(T), id).Error
}

// Delete performs a soft delete on a record (asynchronous).
func (s *ModelService[T]) Delete(id uint) <-chan error {
	result := make(chan error, 1)
	go func() {
		defer close(result)
		result <- s.DB.Delete(new(T), id).Error
	}()
	return result
}

// FindByQuerySync retrieves records based on a query (synchronous).
func (s *ModelService[T]) FindByQuerySync(query map[string]interface{}) ([]T, error) {
	var entities []T
	if err := s.DB.Where(query).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// FindByQuery retrieves records based on a query (asynchronous).
func (s *ModelService[T]) FindByQuery(query map[string]interface{}) <-chan struct {
	Data  []T
	Error error
} {
	result := make(chan struct {
		Data  []T
		Error error
	}, 1)
	go func() {
		defer close(result)
		var entities []T
		if err := s.DB.Where(query).Find(&entities).Error; err != nil {
			result <- struct {
				Data  []T
				Error error
			}{nil, err}
			return
		}
		result <- struct {
			Data  []T
			Error error
		}{entities, nil}
	}()
	return result
}

// Transaction methods

// BeginTransactionSync starts a new database transaction (synchronous).
func (s *ModelService[T]) BeginTransactionSync() (*gorm.DB, error) {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

// BeginTransaction starts a new database transaction (asynchronous).
func (s *ModelService[T]) BeginTransaction() <-chan struct {
	Tx    *gorm.DB
	Error error
} {
	result := make(chan struct {
		Tx    *gorm.DB
		Error error
	}, 1)
	go func() {
		defer close(result)
		tx := s.DB.Begin()
		result <- struct {
			Tx    *gorm.DB
			Error error
		}{tx, tx.Error}
	}()
	return result
}

// CommitTransactionSync commits the transaction (synchronous).
func (s *ModelService[T]) CommitTransactionSync(tx *gorm.DB) error {
	return tx.Commit().Error
}

// CommitTransaction commits the transaction (asynchronous).
func (s *ModelService[T]) CommitTransaction(tx *gorm.DB) <-chan error {
	result := make(chan error, 1)
	go func() {
		defer close(result)
		result <- tx.Commit().Error
	}()
	return result
}

// RollbackTransactionSync rolls back the transaction in case of an error (synchronous).
func (s *ModelService[T]) RollbackTransactionSync(tx *gorm.DB) error {
	return tx.Rollback().Error
}

// RollbackTransaction rolls back the transaction in case of an error (asynchronous).
func (s *ModelService[T]) RollbackTransaction(tx *gorm.DB) <-chan error {
	result := make(chan error, 1)
	go func() {
		defer close(result)
		result <- tx.Rollback().Error
	}()
	return result
}
