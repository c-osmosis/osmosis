package types

import (
	"fmt"
	"time"
)

const (
	ModuleName = "poolincentives"

	StoreKey = ModuleName

	RouterKey = ModuleName

	QuerierRoute = ModuleName
)

var (
	LockableDurationsKey = []byte("lockable_durations")
	DistrInfoKey         = []byte("distr_info")
)

func GetPoolPodIdStoreKey(poolId uint64, duration time.Duration) []byte {
	return []byte(fmt.Sprintf("pool-incentives/%d/%s", poolId, duration.String()))
}
