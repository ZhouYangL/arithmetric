package main

import (
	"fmt"
	"time"
)

/*
	编写一个程序，通过已填充的空格来解决数独问题。
	一个数独的解法需遵循如下规则：
	数字 1-9 在每一行只能出现一次。
	数字 1-9 在每一列只能出现一次。
	数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
	空白格用 '.' 表示。
 */

func PrintBoard( board [][]byte)  {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			fmt.Printf("%s ", string(board[i][j]))
		}
		fmt.Print("\n")
	}
	fmt.Println("================")
}

func checkVild(board [][]byte, x int, y int) bool  {
	for j := 0; j < len(board); j++ {
		if j != y &&  board[x][j] == board[x][y] {
			return false
		}
	}
	
	for j := 0; j < len(board); j++ {
		if j != x &&  board[j][y] == board[x][y] {
			return false
		}
	}
	
	for i := x/3*3; i < (x/3+1)*3; i++ {
		for j := y/3*3; j < (y/3+1)*3; j++ {
			if x == i {
				break
			}
			
			if i != x && j != y && board[x][y] == board[i][j] {
				return false
			}
		}
	}
	return true
}


func OrderSudoku(board [][]byte, value byte, index int, maxIndex int, mMap [][2]int) bool {
	if index >= maxIndex {
		return true
	}
	i, j := mMap[index][0], mMap[index][1]
	for ; value <= '9'; value++{
		board[i][j] = value
		if checkVild(board, i, j) {
			if OrderSudoku(board,'1', index+1, maxIndex, mMap) {
				return true
			}
		}
		continue
	}
	board[i][j] = '.'
	return false
}


func solveSudoku(board [][]byte)  {
	mMap := make([][2]int, 81)
	index := 0
	for i := 0; i < len(board); i++ {
		for j:= 0; j < len(board); j++ {
			if board[i][j] == '.' {
				mMap[index] = [2]int{i, j}
				index++
			}
		}
	}
	OrderSudoku(board, '1', 0, index, mMap)
}


func main() {
	a := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	start := time.Now()
	solveSudoku(a)
	fmt.Println(time.Now().Sub(start))
	PrintBoard(a)
}