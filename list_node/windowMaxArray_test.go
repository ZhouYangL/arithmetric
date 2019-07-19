package list_node

import (
	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
	"fmt"
	"testing"
)

type IdxData struct {
	idx int
	vae int
}

// 生成窗口的最大数组
func getMaxWindow(data []int,  w int) []int {
	if len(data) == 0 || w == 0{
		return make([]int, 0)
	}
	queue := deque.New()
	var res []int
	for index, value := range data {
		//在添加元素前 先将队列中小于该值的数据删除掉，确保队列的值是有序的
		for !queue.Empty() && queue.Right().(IdxData).vae <= value {
			queue.PopRight()
		}
		queue.PushRight(IdxData{index, value})
		//处理过期的数据
		if queue.Right().(IdxData).idx == index - w {
			queue.PopLeft()
		}
		//记录最大的值
		if !queue.Empty(){
			res = append(res, queue.Left().(IdxData).vae)
		}
	}
	return res
}


func Test_getMaxWindow(t *testing.T)  {
	data1 := []int{1, 2, 4, 10, 3, 6, 100, 0}
	res := getMaxWindow(data1, 4)
	fmt.Println(res)
}

