package models

type VarEnviroment struct {
	Host        string
	Port        int32
	User        string
	Pass        string
	DB          string
	Service     string
	ServicePort string
	RedisHost   string
	RedisPort   string
	RedisUser   string
	RedisPass   string
}

type VarSchema struct {
	Core string
	Fin  string
	DBO  string
	Ent  string
}
