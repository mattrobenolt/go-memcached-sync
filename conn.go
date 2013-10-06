package sync

import "github.com/bradfitz/gomemcache/memcache"

func NewConnection(addrs string...) *memcache.Client {
	return memcache.New(addrs...)
}
