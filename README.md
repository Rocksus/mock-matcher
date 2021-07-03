# Mock Matcher

Golang gomock matcher examples.

## Usage

You can use the matchers listed here by running

```bash
go get github.com/Rocksus/mock-matcher
```

## Examples

```go
// compare two slices with same values
matcher := NewInt64SliceMatcher([]int64{1,2,3,4,5}])
fmt.Println(matcher.Matches([]int64{5,4,2,1,3})) // true
```
