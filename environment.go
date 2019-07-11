package environment

type EnvironmentalContext struct {
	databases map[string]*databaseContext
	objects 	map[string]interface{}
}

func InitEnvironmentalContext() *EnvironmentalContext {
	return &EnvironmentalContext{make(map[string]*databaseContext), make(map[string]interface{})}
}

func(e *EnvironmentalContext) ConnectDatabase(nickname, driverName, connectionInfo string) error {
	dc, err := initDatabaseContext(driverName,connectionInfo)
	if err != nil {
		return err
	}
	e.databases[nickname] = dc
	return nil
}

func(e *EnvironmentalContext) GetConnection(nickname string) *databaseContext {
	return e.databases[nickname]
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
