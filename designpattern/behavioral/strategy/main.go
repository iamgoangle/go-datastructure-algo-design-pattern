package main

import (
	"fmt"
	"log"
)

// CachingStrategy interface declares operations common to all
// supported versions of some algorithm. The context uses this
// interface to call the algorithm defined by the concrete
// strategies.
type CachingStrategy interface {
	Execute() error
}

type redis struct{}

// RedisStrategy represent concrete strategy for Redis caching
func RedisStrategy() CachingStrategy {
	return &redis{}
}

func (r *redis) Execute() error {
	fmt.Println("Implement Redis strategy")
	return nil
}

type memcache struct{}

// MemcacheStrategy represent concrete strategy for Memcache caching
func MemcacheStrategy() CachingStrategy {
	return &memcache{}
}

func (r *memcache) Execute() error {
	fmt.Println("Implement Memcache strategy")
	return nil
}

// Context type
type Context struct {
	Strategy CachingStrategy
}

// NewContext defines the interface of interest to clients
func NewContext() *Context {
	return &Context{}
}

// SetStrategy usually the context accepts a strategy through the
// constructor, and also provides a setter so that the
// strategy can be switched at runtime.
func (c *Context) SetStrategy(s CachingStrategy) error {
	c.Strategy = s

	return nil
}

// ExecuteStrategy delegates some work to the strategy object
// instead of implementing multiple versions of the
// algorithm on its own.
func (c *Context) ExecuteStrategy() error {
	return c.Strategy.Execute()
}

func main() {
	redis := RedisStrategy()
	// memcache := MemcacheStrategy()

	context := NewContext()
	err := context.SetStrategy(redis)
	if err != nil {
		log.Panic(err)
	}

	err = context.ExecuteStrategy()
	if err != nil {
		log.Panic(err)
	}
}
