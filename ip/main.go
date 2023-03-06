package main

/*
このプログラムは、IPv4 アドレスを表す型 Ipv4 を定義します。
Ipv4 型は他の IPv4 アドレスと同じサブネットにあるかどうかを判定するメソッドを持ちます。

参考：
https://ja.wikipedia.org/wiki/IP%E3%82%A2%E3%83%89%E3%83%AC%E3%82%B9#/media/%E3%83%95%E3%82%A1%E3%82%A4%E3%83%AB:Ipv4_address_ja.svg
*/

import (
	"fmt"
	"strconv"
	"strings"
)

type Ipv4 uint32

func NewIpv4(s string) (*Ipv4, error) {
	arr1 := strings.Split(s, ".")
	// fmt.Println(arr1)
	if len(arr1) != 4 {
		return nil, fmt.Errorf("IPアドレスではありません")
	}

	for _, s := range arr1 {
		fmt.Printf("%s\n", s)
		i, _ := strconv.Atoi(s)
		fmt.Printf("%08b\n", i)
	}

	return nil, nil
}

func (i *Ipv4) String() string {

	return "192.168.0.1"
}

func (i *Ipv4) IsInSameSubnet(ip *Ipv4, mask *Ipv4) bool {

	return false
}

func main() {
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
