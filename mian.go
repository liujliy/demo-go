package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello Go!")

	// struct
	/*test := ColorPoint{
		Point: Point{3, 4},
		Color: color.RGBA{G: 225, B: 225, A: 1},
	}
	fmt.Println(test)*/

	fmt.Println("-----------File Tree--------------")
	fmt.Println("E:\\config")
	FileTree("E:\\config", 0)
	fmt.Println("--------------End-----------------")
}

func FileTree(filePath string, level int) {
	fileList, _ := ioutil.ReadDir(filePath)
	for i := 0; i < len(fileList); i++ {
		file := fileList[i]
		var tag, pre string

		if i+1 == len(fileList) {
			tag = "`--"
		} else {
			tag = "|--"
		}
		for j := 0; j < level; j++ {
			pre += "|  "
		}
		pre = pre + tag
		fmt.Println(pre + file.Name())
		if file.IsDir() {
			FileTree(filePath+"\\"+file.Name(), level+1)
		}
	}
	/*for _, file := range fileList {
		fmt.Println(tag + file.Name())
		if file.IsDir() {
			FileTree(filePath + "\\" + file.Name(), "|  " + tag)
		}
	}*/
}

type Point struct {
	x, y float64
}

type ColorPoint struct {
	Point
	Color color.RGBA
}
