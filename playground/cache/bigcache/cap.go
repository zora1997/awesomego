package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache/v3"
	"runtime"
	"time"
)

const (
	// appid bigCache
	key                = "appidbc"
	defaultShards      = 1024
	defaultLifeWindow  = 30 * 24 * time.Hour
	defaultCleanWindow = 0
	loopTimes          = 100000000
	uidStart           = 2100000001
	wesingAppID        = 1000366
)

type UID2AppID struct {
	AppIDBC map[uint64]uint32
}

func main() {
	// bigcache
	bigCache, err := bigcache.New(context.Background(), bigcache.Config{
		Shards:             defaultShards,
		LifeWindow:         defaultLifeWindow,
		CleanWindow:        defaultCleanWindow,
		MaxEntriesInWindow: 1000 * 10 * 60,
		MaxEntrySize:       500,
		// StatsEnabled:         false,
		Verbose: true,
		// Hasher:               nil,
		HardMaxCacheSize: 0,
		OnRemove:         nil,
		// OnRemoveWithMetadata: nil,
		OnRemoveWithReason: nil,
		// Logger:               nil,
	})
	if err != nil {
		fmt.Printf("bigcache new err: %s", err)
	}

	setUID2AppID := UID2AppID{AppIDBC: map[uint64]uint32{}}
	for i := 0; i < loopTimes; i++ {
		setUID2AppID.AppIDBC[uint64(i)+uidStart] = wesingAppID
	}

	err = setAppIDBigCache(bigCache, setUID2AppID)
	if err != nil {
		fmt.Printf("setAppIDBigCache err: %s\n", err)
		return
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("total: %+v MB\n", m.TotalAlloc/1024/1024)
	fmt.Printf("os %d\n", m.Sys)
}

/*
func getAppIDBigCache(bigCache *bigcache.BigCache) (*UID2AppID, error) {
	uid2AppID := &UID2AppID{}

	// 1. 拿全量的数据
	bytes, err := bigCache.Get(key)
	if err == bigcache.ErrEntryNotFound {
		// 说明对应 key 在 bigcache 里没有数据
		fmt.Printf("%s no value\n", key)
		return uid2AppID, bigcache.ErrEntryNotFound
	}
	// check err
	if err != nil {
		fmt.Printf("bigcache get err: %s, key: %s\n", err, key)
		return uid2AppID, err
	}

	// 2. 解析
	err = json.Unmarshal(bytes, uid2AppID)
	if err != nil {
		fmt.Printf("Unmarshal err: %s\n", err)
		return uid2AppID, err
	}
	// TODO: delete
	fmt.Printf("bigcache get success!\n")

	return uid2AppID, nil
}
*/
func setAppIDBigCache(bigCache *bigcache.BigCache, uid2AppID UID2AppID) error {
	// 1. Marshal
	bytesSet, err := json.Marshal(uid2AppID)
	if err != nil {
		fmt.Printf("Marshal err: %s\n", err)
	}

	// 2. bigcache set
	err = bigCache.Set(key, bytesSet)
	if err != nil {
		fmt.Printf("bigcache set err: %s, key: %s", err, key)
		return err
	}

	// TODO: delete
	fmt.Printf("bigcache set success!\n")

	return nil
}
