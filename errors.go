package environment

import "github.com/pkg/errors"

var (
	ErrorKeyIsZeroLength = errors.New("key value is empty")
	ErrorDatabaseIsNil = errors.New("Database is nil")
	ErrorExistingTrans = errors.New("A transaction already exists")
	ErrorNoTransaction = errors.New("No current database transaction")
)
