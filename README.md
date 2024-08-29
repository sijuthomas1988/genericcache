# genericcache
A simple Thread Safe LRU Generic Cache

```bash
go get github.com/sijuthomas1988/genericcache
```
## Examples
Basic usage.
```go
func main() {
    cache := genericcache.New[string, string](lru.WithCapacity(int))
    cache.Set("key", "value")
    value, _ := cache.Get("key")
    fmt.Println(value)
}
```

