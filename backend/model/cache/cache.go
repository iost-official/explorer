package cache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

// Create a cache with a default expiration time of 1 minutes, and which
// purges expired items every 5 minutes
var GlobalCache = cache.New(time.Minute, time.Second * 20)