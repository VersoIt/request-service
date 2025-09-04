package cfg

type Config struct {
	Server Server
}

type Server struct {
	Port int32
}

type Postgres struct {
}
