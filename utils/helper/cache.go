package helper

import (
	"fmt"
	"os"
	"time"

	cache "github.com/SporkHubr/echo-http-cache"
	"github.com/SporkHubr/echo-http-cache/adapter/memory"
)

func Cache() *cache.Client {
	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(1*time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return cacheClient
}
