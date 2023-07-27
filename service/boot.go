package service

import (
	_ "embed"
	"github.com/redis/go-redis/v9"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/store"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/store/memory"
	redis2 "gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/store/redis"

	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/tests"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/core/utils/maps"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/config"
	"gitlab.com/tmds-io/core-model/hyperion/kore.git/v2/service/kernel"
)

const (
	Version = "latest"
)

func Boot() error {
	err := kernel.New(
		kernel.Name("MASTERDATA"),
		kernel.Version(Version),
		kernel.ConfigMap(maps.H{}),
		//kernel.RegisterHandler(func(svc service.ServiceInterface) error {
		//	return proto.RegisterHandlers(svc.Server(), container.GetAll()...)
		//}),
	).Init()
	if err != nil {
		return err
	}

	// Redis store
	initializeWhoAmIStore()

	return nil
}

func initializeWhoAmIStore() {
	if tests.IsRunning() {
		store.DefaultStore = memory.NewStore()
	} else {
		createRedisStore()
	}
}

func createRedisStore() {
	credis := redis.NewUniversalClient(&redis.UniversalOptions{
		MasterName: config.Get("redis.master").String(),
		Addrs:      []string{config.Get("redis.address").String()},
		Username:   config.Get("redis.username").String(),
		Password:   config.Get("redis.password").String(),
	})

	redisStore := redis2.NewStore(
		redis2.WithClient(credis),
		store.Table("MASTERDATA-"),
	)
	store.DefaultStore = redisStore
}
