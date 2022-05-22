package config

// config of this application
// see config.yaml

type Config struct {
	App AppConfig `mapstructure:"application"`
	DB  DBConfig  `mapstructure:"db"`
	Log LogConfig `mapstructure:"log"`
	Cos CosConfig `mapstruct:"cos"`
}

type AppConfig struct {
	Name     string `mapstruct:"name"`
	Port     int    `mapstructure:"port"`
	CertFile string `mapstructure:"certFile"`
	KeyFile  string `mapstructure:"keyFile"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

type LogConfig struct {
	Mode       string `mapstructure:"mode"`
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxSize"`
	MaxAge     int    `mapstructure:"maxAge"`
	MaxBackups int    `mapstructure:"maxBackups"`
}

type CosConfig struct {
	BucketId  string `mapstruct:"bucketId"`
	Region    string `mapstruct:"region"`
	SecretId  string `mapstruct:"secretId"`
	SecretKey string `mapstruct:"secretKey"`
}
