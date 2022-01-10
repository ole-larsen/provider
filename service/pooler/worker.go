package pooler

import (
	"context"
	"sync"
	"time"

	"github.com/olelarssen/provider/service/client"
	"github.com/olelarssen/provider/service/logger"
)

// Pooler is a worker group that runs a number of operations at concurrency mode depended on operations count.
type Pooler struct {
	logger logger.Logger
	ctx    context.Context
	cancel     context.CancelFunc
	operations map[string][]Operator
	wg         sync.WaitGroup
	done       chan struct{}
	client     *client.Client
	sync.Mutex
	sync.Once
}

// NewPooler initializes a new pool with the given operations at the given concurrency.
func NewPooler(client *client.Client, log logger.Logger, ctx context.Context, cancel context.CancelFunc) *Pooler {
	p := &Pooler{
		client:     client,
		logger:     log,
		ctx:        ctx,
		cancel:     cancel,
		operations: make(map[string][]Operator),
		done:       make(chan struct{}),
	}
	return p
}

// Scheduler for Pooler worker
func (p *Pooler) Scheduler(ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			p.logger.Infoln("ticker")
		case <-p.done:
			p.cancel()
			return
		case <-p.ctx.Done():
			p.done <- struct{}{}
		}
	}
}

// Operator interface
type Operator interface {
	Exec() error
}

type operatorFunc func() error

// operator encapsulates a work item that should go in a work pool.
type operator struct {
	// Err holds an error that occurred during a operation. Its
	// result is only meaningful after run has been called
	// for the pool that holds it.
	Err  error
	f    operatorFunc
	uuid string
}

// NewOperator initializes a new operation based on a given work function.
func NewOperator(f operatorFunc) Operator {
	return &operator{f: f}
}

// Exec runs an operator
func (t *operator) Exec() error {
	t.Err = t.f()
	return t.Err
}

// run all work within the pool and block until it's finished.
func (p *Pooler) Run(source string) {
	p.wg.Add(len(p.operations[source]))
	for _, o := range p.operations[source] {
		go func(o Operator) {
			defer p.wg.Done()
			if err := (o).Exec(); err != nil {
				p.logger.Errorln(err)
			}
		}(o)
	}
	p.wg.Wait()
}
