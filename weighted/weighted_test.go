package weighted

import (
	"fmt"
	"testing"
)

func TestWeightBalance(t *testing.T) {
	r := &WeightBalance{}
	r.Add("127.0.0.1:2003", "4")
	r.Add("127.0.0.1:2004", "3")
	r.Add("127.0.0.1:2005", "2")

	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
	fmt.Println(r.Next())
}
