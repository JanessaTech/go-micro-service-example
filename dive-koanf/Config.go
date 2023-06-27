package divekoanf

import "time"

var defaultConfig = map[string]interface{}{
	"server.port":             8080,
	"server.readTimeout":      "5s",
	"server.writeTimeout":     "10s",
	"server.gracefulShutdown": "30s",

	"logging.level":       -1,
	"logging.encoding":    "console",
	"logging.development": true,

	"jwt.secret":      "secret-key",
	"jwt.sessionTime": "864000s",

	"db.dataSourceName":   "root:password@tcp(127.0.0.1:3306)/local_db?charset=utf8&parseTime=True&multiStatements=true",
	"db.logLevel":         1,
	"db.migrate.enable":   false,
	"db.migrate.dir":      "",
	"db.pool.maxOpen":     10,
	"db.pool.maxIdle":     5,
	"db.pool.maxLifetime": "5m",

	"cache.enabled":            false,
	"cache.prefix":             "article-",
	"cache.type":               "redis",
	"cache.ttl":                60 * time.Second,
	"cache.redis.cluster":      false,
	"cache.redis.endpoints":    []string{"localhost:6379"},
	"cache.redis.readTimeout":  "3s",
	"cache.redis.writeTimeout": "3s",
	"cache.redis.dialTimeout":  "5s",
	"cache.redis.poolSize":     10,
	"cache.redis.poolTimeout":  "1m",
	"cache.redis.maxConnAge":   "0",
	"cache.redis.idleTimeout":  "5m",

	"metrics.namespace": "article_server",
	"metrics.subsystem": "",
}

type Config struct {
	ServerConfig  ServerConfig  `json:"server"`
	LoggingConfig LoggingConfig `json:"logging" yaml:"logging"`
	JwtConfig     JWTConfig     `json:"jwt"`
	DBConfig      DBConfig      `json:"db"`
	CacheConfig   CacheConfig   `json:"cache"`
	MetricsConfig MetricsConfig `json:"metrics"`
	Accounts      []struct {
		Name    string `json:"name"`
		Key     string `json:"key"`
		Address string `json:"address"`
	} `json:"accounts"`
}

type ServerConfig struct {
	Port             int           `json:"port"`
	ReadTimeout      time.Duration `json:"readTimeout"`
	WriteTimeout     time.Duration `json:"writeTimeout"`
	GracefulShutdown time.Duration `json:"gracefulShutdown"`
}

type LoggingConfig struct {
	Level       int    `json:"level"`
	Encoding    string `json:"encoding"`
	Development bool   `json:"development"`
}

type JWTConfig struct {
	Secret      string        `json:"secret"`
	SessionTime time.Duration `json:"sessionTime"`
}

type DBConfig struct {
	DataSourceName string `json:"dataSourceName"`
	LogLevel       int    `json:"logLevel"`
	Migrate        struct {
		Enable bool   `json:"enable"`
		Dir    string `json:"dir"`
	} `json:"migrate"`
	Pool struct {
		MaxOpen     int           `json:"maxOpen"`
		MaxIdle     int           `json:"maxIdle"`
		MaxLifetime time.Duration `json:"maxLifetime"`
	} `json:"pool"`
}

type CacheConfig struct {
	Enabled     bool          `json:"enabled"`
	Prefix      string        `json:"prefix"`
	Type        string        `json:"type"`
	TTL         time.Duration `json:"ttl"`
	RedisConfig RedisConfig   `json:"redis"`
}

type RedisConfig struct {
	Cluster      bool          `json:"cluster"`
	Endpoints    []string      `json:"endpoints"`
	ReadTimeout  time.Duration `json:"readTimeout"`
	WriteTimeout time.Duration `json:"writeTimeout"`
	DialTimeout  time.Duration `json:"dialTimeout"`
	PoolSize     int           `json:"poolSize"`
	PoolTimeout  time.Duration `json:"poolTimeout"`
	MaxConnAge   time.Duration `json:"maxConnAge"`
	IdleTimeout  time.Duration `json:"idleTimeout"`
}

type MetricsConfig struct {
	Namespace string `json:"namespace"`
	Subsystem string `json:"subsystem"`
}
