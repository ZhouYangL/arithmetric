package main

import (
	"fmt"
)

/*
	按下述要求实现 StreamChecker 类：
	StreamChecker(words)：构造函数，用给定的字词初始化数据结构。
	query(letter)：如果存在某些 k >= 1，可以用查询的最后 k个字符（按从旧到新顺序，包括刚刚查询的字母）
	拼写出给定字词表中的某一字词时，返回 true。否则，返回 false。
	示例：
		StreamChecker streamChecker = new StreamChecker(["cd","f","kl"]); // 初始化字典
		streamChecker.query('a');          // 返回 false
		streamChecker.query('b');          // 返回 false
		streamChecker.query('c');          // 返回 false
		streamChecker.query('d');          // 返回 true，因为 'cd' 在字词表中
		streamChecker.query('e');          // 返回 false
		streamChecker.query('f');          // 返回 true，因为 'f' 在字词表中
		streamChecker.query('g');          // 返回 false
		streamChecker.query('h');          // 返回 false
		streamChecker.query('i');          // 返回 false
		streamChecker.query('j');          // 返回 false
		streamChecker.query('k');          // 返回 false
		streamChecker.query('l');          // 返回 true，因为 'kl' 在字词表中。
 */

type Tree struct {
	IsEnd bool
	childs [26]*Tree
}

type StreamChecker struct {
	WordDict Tree
	QueryRecode []byte
}


func Constructor(words []string) StreamChecker {
	stream := StreamChecker{
		QueryRecode: make([]byte, 0),
	}
	for _, ele := range words {
		InsertWord(&stream.WordDict, ele)
	}
	return stream
}

func InsertWord(wordDict *Tree, insert string)  {
	if insert == "" {
		return
	}
	t := wordDict
	for i := len(insert) -1 ; i >= 0; i-- {
		ch := insert[i] - 'a'
		if t.childs[ch] == nil {
			t.childs[ch] = &Tree{}
		}
		t = t.childs[ch]
	}
	t.IsEnd = true
}


func (this *StreamChecker) Query(letter byte) bool {
	this.QueryRecode = append(this.QueryRecode, letter)
	t := &this.WordDict
	for i := len(this.QueryRecode) - 1; i >= 0; i-- {
		ch := this.QueryRecode[i] - 'a'
		if t.childs[ch] != nil && t.childs[ch].IsEnd {
			return true
		}else if t.childs[ch] == nil {
			return false
		}
		t = t.childs[ch]
	}
	if t.IsEnd == true {
		return true
	}
	return false
}


//func bytes2string(b []byte) string {
//	head := (*reflect.SliceHeader)(unsafe.Pointer(&b))
//	s := reflect.StringHeader{
//		Data: head.Data,
//		Len:  head.Len,
//	}
//	return *(*string)(unsafe.Pointer(&s))
//}
//
//
//type stringStruct struct {
//	array unsafe.Pointer
//	len int
//}

func main() {
	//a := []int64{1, 2}
	//x := (*stringStruct)(unsafe.Pointer(&a))
	//
	//s := []byte{'a', 'b', 'c'}
	//
	//t := string(s)
	//fmt.Println(t)
	//var ss string
	//z := (*stringStruct)(unsafe.Pointer(&ss))
	//z.array = unsafe.Pointer(&s[0])
	//z.len = len(s)
	//fmt.Println(ss)
	//e := unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + 1*unsafe.Sizeof(a[0]))
	//fmt.Println(*(*int64)(e))
	//fmt.Printf("%p, %p, %p, %p, %d", &a[0], &a[1], &a, x.array, x.len)
	obj := Constructor([]string{"cd","f","kl"})
	fmt.Println(obj)
	fmt.Println(obj.Query('a'))         // 返回 false
	fmt.Println(obj.Query('b'))
	fmt.Println(obj.Query('c'))
	fmt.Println(obj.Query('d'))
	fmt.Println(obj.Query('e'))
	fmt.Println(obj.Query('f'))
	fmt.Println(obj.Query('g'))
	fmt.Println(obj.Query('h'))
	fmt.Println(obj.Query('i'))
	fmt.Println(obj.Query('j'))
	fmt.Println(obj.Query('k'))
	fmt.Println(obj.Query('l'))
}
