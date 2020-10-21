package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func WriteFile(path string){
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//使用后需要关闭，用defer延迟关闭
	defer f.Close()

	var buf string
	for i := 1; i < 10; i++ {
		//将i = %d\n 这个字符串存入buf中
		buf = fmt.Sprintf("i = %d\n", i)
		_, err2 := f.WriteString(buf)
		if err2 != nil {
			fmt.Println("err = ", err)
			return
		}
		//fmt.Println("n = ", n) 每次写入的长度，上面用_没有接受n
	}
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 1024*2)//2k
	//n表示从文件中读取的长度
	n, err2 := f.Read(buf)
	if err2 != nil && err2 != io.EOF {//文件出差同时没到结尾，EOF结束符
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))//内容是在切片中，转string类型
}

//一行一行读取
func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	//新建一个缓冲区，把内容放在缓冲区
	r := bufio.NewReader(f)
	for {
		buf, err2 := r.ReadBytes('\n') //读取到\n就停止读取,但是\n也被读取进去了
		if err2 != nil {
			if err2 == io.EOF {//表示文件读取完
				break
			}
			fmt.Println("err = ", err2)
			return
		}

		fmt.Printf("buf = #%s#", string(buf))
	}
}

func main() {
	path := "./demo1.txt"
	//WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}
