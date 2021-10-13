package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//import log "github.com/sirupsen/logrus"
//
//func main() {
//	fmt.Println("hello")
//	log.Info("hello,logrus framework")
//	log.Info("读取文件内容")
//	open, err := os.Open("C:\\Users\\raytine\\Desktop\\test.txt")
//	if err != nil {
//		log.Errorf("打开文件异常")
//		return
//	}
//	//buf:=make([]byte,126)
//	//read, err := open.Read(buf)
//	//if err!=nil {
//	//	log.Error("文件读取异常")
//	//	return
//	//}
//	//fmt.Printf("%d=%q",read,buf)
//	//log.Infof("%d=%q",read,buf)
//
//	defer open.Close()
//	reader := bufio.NewReader(open)
//
//	for {
//		readString, err := reader.ReadString('\n')
//		if err!=nil {
//			//log.Info("读取失败")
//			return
//		}
//		log.Infof("读取内容=%s",readString)
//	}
//}
