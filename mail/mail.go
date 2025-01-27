package mail

import (
	"fmt"
	"net/mail"
	"strings"
)

var addresses = []string{
	// "Who? <one@y.test>",
	// "\"Joe Q. Public\" <john.q.public@example.com>",
	// "Mary Smith <mary@x.test>",
	// "<Alice><alice@example.com>",
	// "<boss@nil.test>",
	"boss@nil.test <boss@nil.test>",
	// "\"Alice : [ ] , @ \" <alice@example.com>",
	// "\"Giant; \"Big\" Box\" <sysservices@example.net>",
	// "\"Giant; Big Box\" <sysservices@example.net>",
	// "Giant; Box <sysservices@example.net>",
	// "Sales Team: alice@example.com, bob@example.com ; (Contact for sales inquiries)",
	"Alice <alice@example.com>, Bob <bob@example.com>, Eve <eve@example.com>",
	// "A Group:Ed Jones <c@a.test>,joe@where.test,John <jdoe@one.test>;",
}

func customParse(address string) (*mail.Address, error) {
	fmt.Println("customParse実行")
	i := strings.LastIndex(address, "<")
	l := strings.LastIndex(address, ">")

	// "<"や">"が無い場合はエラーを返す
	if i == -1 || l == -1 {
		return nil, fmt.Errorf("< > は含まれている前提 : %s", address)
	}

	// アドレス
	email := address[i+1 : l]
	// 名前
	name := address[:i]
	name = strings.TrimSuffix(name, " ")
	// 名前に"をつける
	name = "\"" + name + "\""

	tmp := name + " " + "<" + email + ">"
	fmt.Println("デバッグ1", tmp)

	return mail.ParseAddress(tmp)
}

func Parse(address string) (*mail.Address, error) {
	// まずParseAddress実行
	a, err := mail.ParseAddress(address)
	if err == nil {
		return a, nil
	}

	// 複数を渡していないかチェック
	as, err := mail.ParseAddressList(address)
	if err == nil {
		return nil, fmt.Errorf("複数の宛先は期待してない : %v", as)
	}

	// ParseAddressでもParseAddressListでもエラー起きたら、""で囲う
	a, err = customParse(address)
	if err == nil {
		return a, nil
	}

	return nil, err
}

func ParseTest() {
	for _, s := range addresses {
		fmt.Println("_____________________________________")
		fmt.Println("Before:", "---"+s+"---")

		a, err := Parse(s)
		if err != nil {
			// TODO: ちゃんとエラー処理しないといけない
			fmt.Println(s, err)
			continue
		}

		fmt.Println("After:")
		fmt.Println("---" + a.Name + "---")
		fmt.Println("---" + a.Address + "---")
	}
}
