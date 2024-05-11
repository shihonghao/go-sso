package conf

type DbConfig struct {
	DriverName   string
	Dsn          string
	ShowSql      bool
	ShowExecTime bool
	MaxIdle      int
	MaxOpen      int
}

var Db = map[string]DbConfig{
	"db1": {
		DriverName:   "mysql",
		Dsn:          "root:096514@tcp([2409:8a00:c9b:87a1::75]:3306)/sso?charset=utf8mb4&parseTime=true&loc=Local",
		ShowSql:      true,
		ShowExecTime: false,
		MaxIdle:      10,
		MaxOpen:      200,
	},
}
