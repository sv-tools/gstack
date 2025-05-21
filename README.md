# gstack

Simple Generic Stack implementation in Go using linked list.

The repository is archived because the feature is fully implemented.

## Usage

```shell
go get github.com/sv-tools/gstack
```

```go
s := gstack.New(1, 2, 3, 4)
fmt.Println(s.Len())
fmt.Println(s.Pop())
fmt.Println(s.Pop())
fmt.Println(s.Pop())
fmt.Println(s.Pop())
s.Push(5)
fmt.Println(s.Peek())
// Output: 4
// 4
// 3
// 2
// 1
// 5
```

## Changelog
### v0.1.0
- minimal supported version of Go is 1.20
- initial implementation

### v1.0.0
- minimal supported version of Go is 1.23
- added `Iter` method to iterate over the stack and consume all items
```go
s := gstack.New(1, 2, 3, 4)
for v := range s.Iter() {
    fmt.Println(v)
}
fmt.Println("len:", s.Len())
// Output: 4
// 3
// 2
// 1
// len: 0
```

### v1.1.0
- added `IsEmpty` method to check if the stack is empty
