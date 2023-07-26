# gstack
Simple Generic Stack implementation in Go using linked list

## Usage

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

