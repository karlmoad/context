package environment

type environment struct {
	databases map[string]*databaseContext
	objects 	map[string]interface{}
}



func InitEnvironmentalContext() *environment {
	return &environment{make(map[string]*databaseContext), make(map[string]interface{})}
}

func(e *environment) ConnectDatabase(nickname, driverName, connectionInfo string) error {
	dc, err := initDatabaseContext(driverName,connectionInfo)
	if err != nil {
		return err
	}
	e.databases[nickname] = dc
	return nil
}

func(e *environment) GetConnection(nickname string) *databaseContext {
	return e.databases[nickname]
}

func (e *environment) Put(key string, val interface{}) error {
	if len(key) <= 0 {
		return ErrorKeyIsZeroLength
	}
	e.objects[key] = val
	return nil
}

func (e *environment) Get(key string) interface{} {
	return e.objects[key]
}
