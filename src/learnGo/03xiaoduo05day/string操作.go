package main

import "fmt"
import "strings"

func main() {
	//1、Contains,字符串s中是否包含substr，返回bool值
	//func Contains(s, substr string) bool
	fmt.Println(strings.Contains("seafood","foo"))	//ture
	fmt.Println(strings.Contains("seafood","foor"))	//false
	fmt.Println(strings.Contains("seafood",""))		//ture
	fmt.Println(strings.Contains("",""))				//ture

	//2、Join,字符串连接,通过sep连接起来
	s := []string{"abc","de","fg"}
	fmt.Println(strings.Join(s, ","))	//abc,de,fg
	fmt.Println(strings.Join(s, ""))	//abcdefg

	//3、Index,在字符串s中查找sep所在的位置，返回位置值，找不到返回-1
	//func Index(s, sep string) int
	fmt.Println(strings.Index("chicken", "ke"))	//4
	fmt.Println(strings.Index("chicken", "dm"))	//-1

	//4、Repeat,重复字符串
	//func Repeat(s string,count int) string
	fmt.Println("ba"+ strings.Repeat("na", 2))		//banana

	//5、Replace，替换
	//func Replace(s, old, new string,n int ) string ,n为替换次数，小于0表示全部
	fmt.Println(strings.Replace("oink oink oink oink", "k", "y", 2))		//oiny oiny oink oink
	fmt.Println(strings.Replace("oink oink oink oink", "k", "y", -1))	//oiny oiny oiny oiny

	//4、Split分离
	//func Split(s, sep string) []string ,把s按照sep分离存入切片中
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))					//["a" "b" "c"]
	fmt.Printf("%q\n", strings.Split("a man a pen a apple", "a"))	//["" " m" "n " " pen " " " "pple"]
	fmt.Printf("%q\n", strings.Split(" x yz ", " "))					//["" "x" "yz" ""]
	fmt.Printf("%q\n", strings.Split("", "asdaszxczxc"))				//[""]

	//5、Trim,在s头部尾部去除cutest指定字符串
	//func Trim(s string, cutest string) string
	fmt.Printf("[%q]\n", strings.Trim("!!!acfunca!!", "!"))			//["acfunca"]

	//6、Fields, 去除s的所有空格，并返回一个切片
	//func Fields(s string) []string
	fmt.Printf("[%q]\n", strings.Fields(" foo bar baz "))					//[["foo" "bar" "baz"]]

}
