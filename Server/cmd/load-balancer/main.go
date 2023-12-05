package main

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
	"time"
)

// https://kasvith.me/posts/lets-create-a-simple-lb-go/

type Backend struct {
	URL          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backend []*Backend
	current uint64
}

// Add index with synchronization
func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backend)))
}

// Mutex lock set alive to backend
func (b *Backend) SetAlive(alive bool) {
	b.mux.Lock()
	b.Alive = alive
	b.mux.Unlock()
}

func (b *Backend) CheckAlive() bool {
	alive := false
	b.mux.RLock()
	alive = b.Alive
	b.mux.RUnlock()
	return alive
}

// Return next active peer
func (s *ServerPool) GetNextPeer() *Backend {
	// Loop and find alive one
	nextIdx := s.NextIndex()
	nextCycle := len(s.backend) + nextIdx
	// Looping through idx -> idx+len
	for idx := nextIdx; idx < nextCycle; idx++ {
		pos := idx % len(s.backend)
		if s.backend[pos].CheckAlive() {
			atomic.StoreUint64(&s.current, uint64(pos))
			return s.backend[pos]
		}
	}
	return nil
}

var serverPool ServerPool

func load_balancer(w http.ResponseWriter, r *http.Request) {
	peer := serverPool.GetNextPeer()
	if peer != nil {
		peer.ReverseProxy.ServeHTTP(w, r)
		return
	}
	http.Error(w, "Service not available!", http.StatusServiceUnavailable)
}

func GetRetries(r *http.Request) int {
	if retry, ok := r.Context().Value("Retry").(int); ok {
		return retry
	}
	return 0
}

func check_healthy_handler(serverUrl url.URL, reverseProxy httputil.ReverseProxy) func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, r *http.Request, e error) {
		log.Printf("[%s] %s\n", serverUrl.Host, e.Error())
		retries := GetRetries(r)
		if retries < 5 {
			select {
			// Delay 100ms each retries
			case <-time.After(100 * time.Millisecond):
				// Accumulate retries in request through context
				ctx := context.WithValue(r.Context(), "Retry", retries+1)
				reverseProxy.ServeHTTP(w, r.WithContext(ctx))

			}
		}

	}
}

func main() {
}
