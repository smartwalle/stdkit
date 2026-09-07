package stdkit

import (
	"context"
	"time"
)

type Context struct {
	parent context.Context
	cancel context.CancelFunc
}

func NewContext(parent context.Context) *Context {
	parent, cancel := context.WithCancel(parent)
	return &Context{
		parent: parent,
		cancel: cancel,
	}
}

func ContextWithTimeout(parent context.Context, timeout time.Duration, cause ...error) *Context {
	var err error
	if len(cause) > 0 {
		err = cause[0]
	}
	parent, cancel := context.WithTimeoutCause(parent, timeout, err)
	return &Context{
		parent: parent,
		cancel: cancel,
	}
}

func ContextWithDeadline(parent context.Context, deadline time.Time, cause ...error) *Context {
	var err error
	if len(cause) > 0 {
		err = cause[0]
	}
	parent, cancel := context.WithDeadlineCause(parent, deadline, err)
	return &Context{
		parent: parent,
		cancel: cancel,
	}
}

func (c *Context) Deadline() (deadline time.Time, ok bool) {
	return c.parent.Deadline()
}

func (c *Context) Done() <-chan struct{} {
	return c.parent.Done()
}

func (c *Context) Err() error {
	return c.parent.Err()
}

func (c *Context) Value(key any) any {
	return c.parent.Value(key)
}

func (c *Context) Wait() {
	<-c.parent.Done()
}

func (c *Context) Cancel() {
	c.cancel()
}

func (c *Context) Cancelled() bool {
	select {
	case <-c.parent.Done():
		return true
	default:
		return false
	}
}

func (c *Context) Cause() error {
	return context.Cause(c.parent)
}
