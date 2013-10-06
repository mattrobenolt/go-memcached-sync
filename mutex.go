package sync

import (
	"time"
	"github.com/bradfitz/gomemcache/memcache"
)

var (
	defaultPollFrequency = 100 * time.Millisecond
	lockValue = []byte("1")
)

type Mutex struct {
	Key string
	Client *memcache.Client
	PollFrequency time.Duration
}

func (m *Mutex) createLock(ttl int32) error {
	return l.client.Add(&memcache.Item{Key: key, Value: lockValue, Expiration: ttl})
}

func (m *Mutex) MaybeLock() bool {
	return m.MaybeLockTtl(0)
}

func (m *Mutex) MaybeLockTtl(ttl int32) bool {
	err := m.createLock(ttl)
	return err == nil
}

func (m *Mutex) Lock() {
	m.LockTtl(0)
}

func (m *Mutex) LockTtl(ttl int32) {
	for {
		if ok := m.MaybeLockTtl(ttl); ok {
			return
		}
		time.Sleep(m.PollFrequency)
	}
}

func (m *Mutex) Unlock() {
	if err := m.client.Delete(key); err != nil {
		panic("sync: unlock of unlocked mutex")
	}
}

func NewMutex(key string, client *memcache.Client) *Mutex {
	return &Mutex{key, client, defaultPollFrequency}
}
