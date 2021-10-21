package main

//import (
//	"encoding/hex"
//	"fmt"
//	"net"
//	"time"
//)
//
//func main1() {
//	var buf [512]byte
//	addr, err := net.ResolveTCPAddr("tcp", "192.168.132.129:3306")
//	result := checkError(err)
//	if !result {
//		return
//	}
//	conn, err := net.DialTCP("tcp", nil, addr)
//	result = checkError(err)
//	if !result {
//		return
//	}
//	remoteAddr := conn.RemoteAddr()
//	readByte, err := conn.Read(buf[0:])
//	result = checkError(err)
//	if !result {
//		return
//	}
//	fmt.Println("Reply from Server", remoteAddr.String(), string(buf[0:readByte]))
//	fmt.Println(buf)
//	dstEncode := make([]byte, hex.EncodedLen(len(buf)))
//	hex.Encode(dstEncode, buf[0:readByte])
//	fmt.Println("16进制")
//	fmt.Printf("%x\n", dstEncode)
//
//	fmt.Println("dump")
//	fmt.Println(hex.Dump(buf[0:readByte]))
//	time.Sleep(time.Second * 2)
//
//	conn.Close()
//}
//
//func checkError(err error) bool {
//	if err != nil {
//		fmt.Errorf("fatal error:%s", err)
//		return false
//	}
//	return true
//}
