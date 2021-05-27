package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", httpTest)           //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func httpTest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("Path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key: ", k)
		fmt.Println("Value: ", strings.Join(v, ","))
	}
	fmt.Fprintf(w, "Hello astaxie!")
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
