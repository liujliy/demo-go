package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func main() {
	/*http.HandleFunc("/", httpTest)           //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
	// CPU性能数据
	/*pprof.StartCPUProfile(os.Stdout)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}
	fmt.Println(n)*/

	// array/slice
	/*words := []string{"GO", "Java", "C++"}
	for i, s := range words {
		words = append(words, "test")
		fmt.Println(i, s)
	}
	fmt.Println(words)

	// map
	testMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	for k, v := range testMap {
		fmt.Println(k, v)
	}

	// chan
	ch := make(chan string)
	go func() {
		ch <- "Go"
		ch <- "Java"
		ch <- "C++"
		close(ch)
	}()
	for n := range ch {
		fmt.Println(n)
	}*/
	os.Setenv("CONFIG_NAME", "global_server")
	os.Setenv("CONFIG_IP", "10.0.0.1")
	os.Setenv("CONFIG_URL", "geektutu.com")
	c := readConfig()
	fmt.Printf("%+v", c)
	fmt.Println(unsafe.Sizeof(c))
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
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

type MakeTemplate struct {
	ptr *[]int
	len int
	cap int
}

type Config struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	URL  string `json:"url"`
}

func readConfig() *Config {
	config := Config{}
	// 反射
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}
