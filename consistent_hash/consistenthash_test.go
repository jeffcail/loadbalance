package consistent_hash

import (
	"fmt"
	"testing"
)

func TestNewConsistentHashBalance(t *testing.T) {
	r := NewConsistentHashBalance(10, nil)
	r.Add("127.0.0.1:2003")
	r.Add("127.0.0.1:2004")
	r.Add("127.0.0.1:2005")
	r.Add("127.0.0.1:2006")
	r.Add("127.0.0.1:2007")

	fmt.Println(r.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(r.Get("http://127.0.0.1:2002/base/errinfo"))
	fmt.Println(r.Get("http://127.0.0.1:2002/base/getinfo"))
	fmt.Println(r.Get("http://127.0.0.1:2002/base/pwd"))

	fmt.Println(r.Get("127.0.0.1"))
	fmt.Println(r.Get("192.168.0.1"))
	fmt.Println(r.Get("127.0.0.1"))
}
