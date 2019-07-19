package serach

import "testing"

func binary(data []int, value int) int {
	hight := len(data) - 1
	low := 0
	for{
		if low > hight{
			return -1
		}
		middle := (hight + low)/2
		if value == data[middle]{
			return middle
		} else if value > data[middle] {
			low = middle + 1
		} else {
			hight = middle - 1
		}
	}
	return  -1
}




func Test_Main(t *testing.T){
	data1 := []int{1,4,6,9,10}
	data2 := []int{}
	data3 := []int{0,1,1,1,1,1,1,1,1}
	data4 := []int{0,1,2,3,4,5,10,11}
	
	if binary(data1,4) != 1{
		t.Error("第1个测试错误")
	}
	
	if binary(data2,4) != -1{
		t.Error("第2个测试错误")
	}
	
	if binary(data3,1) != 4{
		t.Error("第3个测试错误")
	}
	
	if binary(data4,10) != 6{
		t.Error("第4个测试错误")
	}
}