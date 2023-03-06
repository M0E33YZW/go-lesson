package main

/*
このプログラムは、IPv4 アドレスを表す型 Ipv4 を定義します。
Ipv4 型は他の IPv4 アドレスと同じサブネットにあるかどうかを判定するメソッドを持ちます。

1. サブネットマスクとIPアドレスを2進数に変換する
2. IPアドレスとサブネットマスクとの論理積（AND）を求める
3. 求めた結果を10進数に戻す

参考：
https://ja.wikipedia.org/wiki/IP%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9#/media/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB:Ipv4_address_ja.svg
*/

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Ipv4 uint32

func NewIpv4(s string) (*Ipv4, error) {
	var bin uint32

	a := strings.Split(s, ".")
	// fmt.Println("a:", a)
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
		if n > 255 {
			return nil, errors.New("255以下の値でなければいけません")
		}

		bin |= uint32(n) << offset
		offset -= 8
	}
	fmt.Printf("%032b\n", bin)

	return (*Ipv4)(&bin), nil
}

func (i *Ipv4) String() string {

	return "192.168.0.1"
}

func (i *Ipv4) IsInSameSubnet(ip *Ipv4, mask *Ipv4) bool {

	return false
}

func main() {
	NewIpv4("172.16.254.1")
	_, err := NewIpv4("192.168.22.11")

	// IP アドレス 1
	// ip1, err := NewIpv4("192.168.0.1")
	if err != nil {
		panic(err)
	}

	/* // IP アドレス 2
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
	fmt.Println(ip1.IsInSameSubnet(ip3, mask)) // false */
}
