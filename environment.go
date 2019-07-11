package environment

import (
	"crypto/sha256"
	"encoding/hex"
)

type EnvironmentalContext struct {
	nickname2DbMap map[string]string
	databases map[string]*databaseContext
	objects 	map[string]interface{}
}

func InitEnvironmentalContext() *EnvironmentalContext {
	return &EnvironmentalContext{nickname2DbMap:make(map[string]string),
							databases:make(map[string]*databaseContext),
							objects:make(map[string]interface{})}
}

func(e *EnvironmentalContext) ConnectDatabase(nickname, driverName, connectionInfo string) error {

	bytes := []byte(driverName+connectionInfo)
	b := sha256.Sum256(bytes)
	connInfoHash := hex.EncodeToString(b[:])

	if _, ok := e.databases[connInfoHash]; !ok {
		dc, err := initDatabaseContext(driverName,connectionInfo)
		if err != nil {
			return err
		}
		e.databases[connInfoHash] = dc
	}

	e.nickname2DbMap[nickname] = connInfoHash

	return nil
}

func(e *EnvironmentalContext) GetConnection(nickname string) *databaseContext {
	if val, ok := e.nickname2DbMap[nickname]; ok {
		return e.databases[val]
	}
	return nil
}

func (e *EnvironmentalContext) Put(key string, val interface{}) error {
	if len(key) <= 0 {
		return ErrorKeyIsZeroLength
	}
	e.objects[key] = val
	return nil
}

func (e *EnvironmentalContext) Get(key string) interface{} {
	return e.objects[key]
}
