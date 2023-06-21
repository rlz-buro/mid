package mid

import (
	"sync"
)

type Publisher struct {
	ch   chan []byte
	done chan struct{}
	wg   sync.WaitGroup
	once sync.Once
}

func NewPublisher() *Publisher {
	return &Publisher{
		ch:   make(chan []byte),
		done: make(chan struct{}),
		wg:   sync.WaitGroup{},
		once: sync.Once{},
	}
}

func (p *Publisher) Read() <-chan []byte {
	return p.ch
}

func (p *Publisher) Write(data []byte) {
	p.wg.Add(1)
	defer p.wg.Done()
	select {
	case <-p.done:
	case p.ch <- data:
	}
}

func (p *Publisher) Close() {
	p.once.Do(func() {
		close(p.done)
		go func() {
			for range p.ch {
			}
		}()
		p.wg.Wait()
		close(p.ch)
	})
}
