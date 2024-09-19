package rotation

import (
	"fmt"
	"testing"
)

func TestRoundRotationBalance(t *testing.T) {
	r := &RoundRotationBalance{}
	r.Add("127.0.0.1:2003")
	r.Add("127.0.0.1:2004")
	r.Add("127.0.0.1:2005")
	r.Add("127.0.0.1:2006")
	r.Add("127.0.0.1:2007")

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
