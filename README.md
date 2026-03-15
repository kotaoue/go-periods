# go-periods
Package for date period

## Usage
```go
import "github.com/kotaoue/go-periods"

func main() {
	period, _ := periods.Split("20220101")
	fmt.Println(period)
	// [20220101]

	days, _ := periods.Split("20220101-20220105")
	fmt.Println(days)
	// [20220101 20220102 20220103 20220104 20220105]

	months, _ := periods.Split("202201-202203")
	fmt.Println(months)
	// [202201 202202 202203]
}
```