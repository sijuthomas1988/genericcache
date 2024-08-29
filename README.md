# genericcache
A simple Thread Safe Generic Cache

## Examples
Basic usage.
```go
func main() {
    cache := lru.New[string, string]()
    cache.Set("key", "value")
    value, _ := cache.Get("key")
    fmt.Println(value)
}
```
Set the capacity using the `lru.WithCapacity` option. The default capacity is set to 10000.
```go
func main() {
    cache := lru.New[string, string](lru.WithCapacity(100))
    ...
}
```

