# go-periods
Package for date period

## Usage
```go
import "github.com/kotaoue/go-periods"

func main() {
   	period, _ := Split("20220101")
	fmt.Println(period) 

    periods, _ := Split("20220101-20220105")
	fmt.Println(periods) 
}
```