package entitiy

type ServerConfig struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	ReadTimeout  string `yaml:"read_timeout"`
	WriteTimeout string `yaml:"write_timeout"`
}

type Config struct {
	Server    ServerConfig `yaml:"server"`
	EndPoints []EndPoint   `yaml:"end_points"`
}
type EndPoint struct {
	Name     string            `yaml:"name"`
	Versions []EndPointVersion `yaml:"versions"`
}

type EndPointVersion struct {
	Version  string   `yaml:"version"`
	Method   string   `yaml:"method"`
	Path     string   `yaml:"path"`
	Response Response `yaml:"response"`
}

type Response struct {
	Status int               `yaml:"status"`
	Header map[string]string `yaml:"header"`
	Body   string            `yaml:"body"`
}
