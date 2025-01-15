package main

import (
	"fmt"
	"log"
	"net/mail"

	"github.com/google/go-cmp/cmp"
	sendgrid_mail "github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Customer struct {
	ID      string
	Balance float64
}

func main() {
	s1 := []int{1, 2, 3}
	// バグる
	// s2 := s1[:2]
	// 完全スライス式で容量も指定することで、s2を変更したときの影響範囲を先頭２つの要素のみに絞ることができる
	s2 := s1[:2:2]
	s3 := append(s2, 10)
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(s2, len(s2), cap(s2))
	fmt.Println(s3, len(s3), cap(s3))

	ints := []int{1, 2, 3}
	for i, v := range ints {
		fmt.Println(i, &v, &ints[i])
	}
	fmt.Println(ints)

	bytes := []byte{'a', 'b'}
	strings := []string{"A", "B"}
	fmt.Println(bytes, string(bytes))
	fmt.Println(strings)

	hello := "Hello, 田中, World"
	fmt.Println(hello[:5])
	fmt.Println(hello[:9])

	got := &Customer{
		ID:      "1a",
		Balance: 0.12,
	}
	want := &Customer{
		ID:      "1b",
		Balance: 0.12,
	}
	fmt.Println(cmp.Diff(got, want))

	// address := "Alice , <alice@example.com>"
	// address := "[Alice] <alice@example.com>"
	// address := "Alice： <alice@example.com>"
	// address := "Alice : <alice@example.com>"
	// address := "<Alice><alice@example.com>"
	// address := "alice@example.com <alice@example.com>"
	address := "\"Alice : [ ] , @ \" <alice@example.com>"
	// sendgrid-go
	sendgridMail, err := sendgrid_mail.ParseEmail(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(sendgridMail.Name, sendgridMail.Address)

	// net/mail
	e, err := mail.ParseAddress(address)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.Name, e.Address)
}
