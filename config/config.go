package config

type Config struct {
	Host string `env:"SHORTENER_HOST,required"`
	Port int    `env:"SHORTENER_PORT" envDefault:"7777"`
}
