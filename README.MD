# gomapper

A tiny ObjectMapper

```go
// Create your converter
type MyConverter struct {
    gomapper.BaseConverter
}
func (c *MyConverter) Supports(src interface{}, dest interface{}, options gomapper.Options) bool {
    return c.IsType(src, FromObject{}) && c.IsType(dest, ToObject{})
}
func (c *MyConverter) Convert(obj interface{}, options gomapper.Options) interface{} {
    in := obj.(FromObject)
    return ToObject{
        ....
    }
}

// Create mapper
mapper := gomapper.New()

// Register your converter
mapper.Register(&MyConverter{})

// ....
// Somewhere in your code
// ....
in := FromObjec{}
dest := mapper.MustMap(in, ToObject{}, nil}

assert.IsType(t, ToObject{}, dest)
```

```go
// Add context to mapping process
dest := mapper.MustMap(in, ToObject{}, gomapper.Options{"yourContextData": 1234})
```

````go
// Map nested struct
func (c *MyConverter) Convert(obj interface{}, options gomapper.Options) interface{} {
    in := obj.(FromObject)

    return ToObject{
        ....,
        SomeCoolObject: c.GetMapper().MustMap(in.NestedObject, ToSuperModel{}, nil),
    }
}
````
