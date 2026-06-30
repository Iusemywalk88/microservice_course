package env

import (
	"net"
	"os"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

const (
	redisHostEnvName                 = "REDIS_HOST"
	redisPortEnvName                 = "REDIS_PORT"
	redisConnectionTimeoutSecEnvName = "REDIS_CONNECTION_TIMEOUT_SEC"
	redisMaxIdleEnvName              = "REDIS_MAX_IDLE"
	redisIdleTimeoutSecEnvName       = "REDIS_IDLE_TIMEOUT_SEC"
	expirationSecEnvName             = "REDIS_EXPIRATION_SEC"
)

type redisConfig struct {
	host string
	port string

	connectionTimeout time.Duration

	maxIdle     int
	idleTimeout time.Duration
	expiration  time.Duration
}

func NewRedisConfig() (*redisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("redis host not found")
	}

	port := os.Getenv(redisPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("redis port not found")
	}

	connectionTimeoutStr := os.Getenv(redisConnectionTimeoutSecEnvName)
	if len(connectionTimeoutStr) == 0 {
		return nil, errors.New("redis connection timeout not found")
	}

	connectionTimeout, err := strconv.ParseInt(connectionTimeoutStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse connection timeout")
	}

	expirationTime := os.Getenv(expirationSecEnvName)
	if len(expirationTime) == 0 {
		return nil, errors.New("redis expiration not found")
	}

	expirationDuration, err := strconv.ParseInt(expirationTime, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse expiration")
	}

	maxIdleStr := os.Getenv(redisMaxIdleEnvName)
	if len(maxIdleStr) == 0 {
		return nil, errors.New("redis max idle not found")
	}

	maxIdle, err := strconv.Atoi(maxIdleStr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse max idle")
	}

	idleTimeoutStr := os.Getenv(redisIdleTimeoutSecEnvName)
	if len(idleTimeoutStr) == 0 {
		return nil, errors.New("redis idle timeout not found")
	}

	idleTimeout, err := strconv.ParseInt(idleTimeoutStr, 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse idle timeout")
	}

	return &redisConfig{
		host:              host,
		port:              port,
		connectionTimeout: time.Duration(connectionTimeout) * time.Second,
		maxIdle:           maxIdle,
		idleTimeout:       time.Duration(idleTimeout) * time.Second,
		expiration:        time.Duration(expirationDuration) * time.Second,
	}, nil
}

func (cfg *redisConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}

func (cfg *redisConfig) ConnectionTimeout() time.Duration {
	return cfg.connectionTimeout
}

func (cfg *redisConfig) MaxIdle() int {
	return cfg.maxIdle
}

func (cfg *redisConfig) IdleTimeout() time.Duration {
	return cfg.idleTimeout
}

func (cfg *redisConfig) Expiration() time.Duration { return cfg.expiration }
