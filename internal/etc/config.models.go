package etc

type Configuration struct {
	Db    db
	Web   web
	Redis redis
	Mqtt  mqtt
}

type web struct {
	Listen string
}

type redis struct {
	Enable   bool
	Addr     string
	Password string
	Db       int
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}

type mqtt struct {
	Broker   string
	Username string
	Password string
	Qos      int
}
