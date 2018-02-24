# Base62 algorithm

```
go get github.com/denismitr/base62.git
```

#### From integer To base62
```go
ToBase62(1) // "1"
ToBase62(2) // "2"
ToBase62(10) // "a"
ToBase62(209) // "3n"
ToBase62(209098) // "Soy"
ToBase62(20909809875) // "mP5qJZ"
```

#### From base62 to integer
```go
ToBase10("mP5qJZ") // 20909809875
ToBase10("Soy") // 209098
ToBase10("3n") // 209
ToBase10("a") // 10
ToBase10("2") // 2
ToBase10("1") // 1
```

### License
MIT