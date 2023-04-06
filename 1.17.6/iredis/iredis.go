package iredis

import (
	"github.com/go-redis/redis/v8"
)

const (
	RdsAddrs            = "127.0.0.1:26379"
	RdsPassword         = "****"
	RdsSentinelAddrs    = "127.0.0.1:26379"
	RdsSentinelPassword = "****"
	RdsMasterName       = "mymaster"
)

var (
	RdsCli *redis.Client
)

func init() {

	RdsCli = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:       RdsMasterName,
		SentinelAddrs:    []string{RdsSentinelAddrs},
		SentinelPassword: RdsSentinelPassword,
		Password:         RdsPassword,
	})

}

func TestIRedisConcurrent() {

}
