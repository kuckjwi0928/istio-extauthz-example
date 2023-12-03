package configs

import "flag"

var Cfg = new(Config)

type Config struct {
	Port int
}

func init() {
	flag.IntVar(&Cfg.Port, "port", 8080, "Http server port")
	flag.Parse()
}
