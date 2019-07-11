package environment

import "database/sql"

type databaseContext struct {
	db 	*sql.DB
	tx 	*sql.Tx
}


func initDatabaseContext(driverName, connectionInfo string) (*databaseContext, error) {
	db, err := sql.Open(driverName, connectionInfo)
	if err != nil{
		return nil, err
	}
	return &databaseContext{db,nil}, nil
}

func(dc *databaseContext) HasTransaction() bool {
	return dc.tx != nil
}

func(dc *databaseContext) CreateTransaction() error {
	if dc.db == nil {
		return ErrorDatabaseIsNil
	}
	if dc.tx != nil {
		return ErrorExistingTrans
	}

	tx, err := dc.db.Begin()
	if err != nil {
		return err
	}
	dc.tx = tx
	return nil
}

func (dc *databaseContext) CommitTransaction() error {
	if dc.db == nil {
		return ErrorDatabaseIsNil
	}
	if dc.tx == nil {
		return ErrorNoTransaction
	}

	return dc.tx.Commit()
}


func (dc *databaseContext) RollbackTransaction() error {
	if dc.db == nil {
		return ErrorDatabaseIsNil
	}
	if dc.tx == nil {
		return ErrorNoTransaction
	}

	return dc.tx.Rollback()
}