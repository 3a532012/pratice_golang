// 场景：在一个高并发的web服务器中，要限制IP的频繁访问。现模拟100个IP同时并发访问服务器，每个IP要重复访问1000次。

// 每个IP三分钟之内只能访问一次。修改以下代码完成该过程，要求能成功输出 success:100

package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type VisitIP struct {
	ip   map[string]*time.Time
	lock sync.Mutex
}

func main() {
	ipMap := &VisitIP{ip: make(map[string]*time.Time)}
	success := int64(0)
	fail := int64(0)
	timer := time.NewTimer(3 * time.Minute)
	wait := &sync.WaitGroup{}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(c context.Context) {
		for {
			select {
			case <-timer.C:
				for ip, ipTime := range ipMap.ip {
					different := time.Now().Sub(*ipTime)
					if different >= time.Minute*3 {
						delete(ipMap.ip, ip)
					}
				}
			case <-c.Done():
				return
			}
		}
	}(ctx)
	wait.Add(1000 * 100)
	go func() {
		for i := 0; i < 1000; i++ {
			for j := 0; j < 100; j++ {

				go func(index int) {
					defer wait.Done()
					ip := fmt.Sprintf("192.168.1.%d", index)
					if _, ok := ipMap.get(ip); !ok {
						now := time.Now()
						ipMap.set(ip, &now)
						atomic.AddInt64(&success, 1)
					} else {
						atomic.AddInt64(&fail, 1)
					}
				}(j)
			}
		}
	}()
	wait.Wait()
	fmt.Printf("success: %d \n", success)
	fmt.Printf("fail: %d \n", fail)
}
func (v *VisitIP) set(key string, value *time.Time) {
	v.lock.Lock()
	defer v.lock.Unlock()
	v.ip[key] = value
}

func (v *VisitIP) get(key string) (*time.Time, bool) {
	v.lock.Lock()
	defer v.lock.Unlock()
	time, ok := v.ip[key]
	if !ok {
		return nil, false
	}
	return time, true
}
