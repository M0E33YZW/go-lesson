package main

/*
このプログラムは、IPv4 アドレスを表す型 Ipv4 を定義します。
Ipv4 型は他の IPv4 アドレスと同じサブネットにあるかどうかを判定するメソッドを持ちます。

参考：
https://ja.wikipedia.org/wiki/IP%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9#/media/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB:Ipv4_address_ja.svg
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unsafe"
)

type Ipv4 uint32

func NewIpv4(s string) (*Ipv4, error) {
	var bin uint32

	a := strings.Split(s, ".")
	// fmt.Println("a  :", a)
	if len(a) != 4 {
		return nil, fmt.Errorf("IPアドレスではありません")
	}

	offset := 24
	for _, s := range a {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		if n < 0 {
			return nil, errors.New("0以上の値でなければいけません")
		}
		if n > 256 {
			return nil, errors.New("255以下の値でなければいけません")
		}

		bin |= uint32(n) << offset
		offset -= 8
	}
	fmt.Print("\n&Bin:", &bin)
	fmt.Print("\nBin:", bin)
	return (*Ipv4)(&bin), nil
}

func (i *Ipv4) String() string {
	ip := *(*Ipv4)(unsafe.Pointer(i))
	fmt.Printf("\n%032b", ip)

	b1 := (ip >> 24) & 0xFF
	b2 := (ip >> 16) & 0xFF
	b3 := (ip >> 8) & 0xFF
	b4 := ip & 0xFF

	return fmt.Sprintf("%d.%d.%d.%d", b1, b2, b3, b4)
}

// サブネットのネットワークアドレスを&計算
// ネットワークアドレスが同じであれば、同じサブネットに属する
func (i *Ipv4) IsInSameSubnet(ip *Ipv4, mask *Ipv4) bool {
	a := *(*Ipv4)(unsafe.Pointer(i))
	b := *(*Ipv4)(unsafe.Pointer(ip))
	msk := *(*Ipv4)(unsafe.Pointer(mask))

	return (a&msk == b&msk)
}

func main() {
	NewIpv4("172.16.254.1")  // test data

	// IP アドレス 1
	ip1, err := NewIpv4("192.168.0.1")
	if err != nil {
		panic(err)
	}

	// IP アドレス 2
	ip2, err := NewIpv4("192.168.0.100")
	if err != nil {
		panic(err)
	}

	// IP アドレス 3
	ip3, err := NewIpv4("192.168.1.1")
	if err != nil {
		panic(err)
	}

	mask, err := NewIpv4("255.255.255.0")
	if err != nil {
		panic(err)
	}

	fmt.Println(ip1.IsInSameSubnet(ip2, mask)) // true
	fmt.Println(ip1.IsInSameSubnet(ip3, mask)) // false
}
