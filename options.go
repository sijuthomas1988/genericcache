package genericcache

const (
	// DefaultCapacity is the default cache capacity, can be extended for making it configurable.
	DefaultCapacity = 10000
)

// CacheOption configures a lru cache.
type CacheOption interface {
	apply(*options)
}

// funcCacheOption wraps a function to implement the CacheOption interface.
type funcCacheOption func(o *options)

func (f funcCacheOption) apply(o *options) {
	f(o)
}

// WithCapacity configures how many items can be stored before old items begin
// to be deleted.
func WithCapacity(capacity int) CacheOption {
	return funcCacheOption(func(o *options) {
		o.capacity = capacity
	})
}

// options for a cache instance.
type options struct {
	capacity int
}

// defaultOptions returns options with default values set.
func defaultOptions() *options {
	return &options{
		capacity: DefaultCapacity,
	}
}
