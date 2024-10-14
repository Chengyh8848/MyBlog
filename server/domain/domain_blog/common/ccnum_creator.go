package common

import (
	"sync"
	"time"
)

var ccnumCreator *ccnumStore

type ccnumHolder struct {
	ccnum     int
	createdAt int64

	lock sync.Mutex
}

type ccnumStore struct {
	store *sync.Map

	lifeTime int //uuid 存在时长(s),超过该时长未更新的将被清理
	done     chan int
}

// 清理超过生存期的uuid
func (c *ccnumStore) clean() (cleanNum int) {
	c.store.Range(func(key, value interface{}) bool {
		if (time.Now().Unix() - value.(*ccnumHolder).createdAt) > int64(c.lifeTime) {
			cleanNum += 1
			c.store.Delete(key)
		}
		return true
	})
	return cleanNum
}

// 获取该uuid的调用序号
// 在生存期的uuid返回上一次调用的序号加1，否则返回1
func (c *ccnumStore) GenCcnum(uuid string) int {
	h, ok := c.store.LoadOrStore(uuid, &ccnumHolder{ccnum: 1, createdAt: time.Now().Unix()})
	if ok {
		h.(*ccnumHolder).lock.Lock()
		h.(*ccnumHolder).createdAt = time.Now().Unix()
		h.(*ccnumHolder).ccnum += 1
		h.(*ccnumHolder).lock.Unlock()
		c.store.Store(uuid, h)

		return h.(*ccnumHolder).ccnum
	} else {
		return 1
	}
}

func (c *ccnumStore) Close() {
	close(c.done)
}

func NewCCNumStore(lifeTime int) *ccnumStore {
	if lifeTime <= 0 {
		panic("lifeTime必须大于0")
	}
	store := &ccnumStore{
		store:    &sync.Map{},
		lifeTime: lifeTime,
		done:     make(chan int),
	}
	go func() {
		cleanInterval := (store.lifeTime - 5)
		if cleanInterval <= 0 {
			cleanInterval = 1
		}
		ticker := time.NewTicker((time.Duration(cleanInterval)) * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-store.done:
				return
			case <-ticker.C:
				store.clean()
			}
		}
	}()

	return store
}
func init() {
	ccnumCreator = NewCCNumStore(30)
}
