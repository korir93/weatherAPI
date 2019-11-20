package config

type Config struct {
    Apikey    string
    
    
}

var _configs Config

func SetConfig(conf Config) {
    _configs = conf
}

func GetConfig() Config {
    return _configs
}