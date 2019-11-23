package main

import (
	"fmt"

	"github.com/NDLPhat/demo_golang/type_system/student"
)

// "strconv"

func main() {
	// 	// Int
	// 	// aInt := 1
	// 	// var bInt int

	// 	// // String
	// 	// aString := "Hello world"

	// 	// // Array/Slice
	// 	// aSlice := []string{"a", "b", "c"}

	// 	// fmt.Printf("bInt = %d \r\n", bInt)
	// 	// fmt.Printf("aInt = %d \r\n", aInt)
	// 	// fmt.Printf("aString = %s \r\n", aString)
	// 	// fmt.Printf("aSlice = %s \r\n", aSlice)

	// 	// nSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 	// for slice := range nSlice {
	// 	// 	if slice >= 5 {
	// 	// 		fmt.Printf("%d \r\n", slice)
	// 	// 	}
	// 	// }
	// 	// // Map
	// 	// aMap := map[string]int{"age": 100, "level": 123}
	// 	// fmt.Println(aMap)

	// 	// // Struct

	// 	// // s := Student{
	// 	// // fmt.Println(s)
	// 	// // fmt.Printf("%+v", s)
	// 	// s := Student{firstName: "Phat", lastName: "Nguyen", Age: 23, Email: "paht@gmail.com"}
	// 	// // Interface
	// 	// // var i interface{}
	// 	// // // i = s
	// 	// // fmt.Println("interface", i)

	// 	// fmt.Println(s)
	// 	// // Chanel
	// 	// ch := make(chan int, 2)
	// 	// ch <- 1  // write to channel
	// 	// ch <- 10 // write to channel

	// 	// // fmt.Println(<-ch) // read from channel
	// 	// // fmt.Println(<-ch) // read from channel

	email1 := student.Student{"1", "2", 3, "old@mail"}
	email1.SetEmail("new@mail")
	fmt.Println(email1)

	// 	// cInt := 100
	// 	// cStr := strconv.Itoa(cInt)
	// 	// cString := fmt.Sprintf("%d hello", cInt)
	// 	// fmt.Println(reflect.TypeOf(cStr), reflect.TypeOf(cString), cString)

	// 	// cStrr := "23"
	// 	// fmt.Println(strconv.ParseInt(cStrr, 10, 64))

	// 	inputJson := `[
	// 		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
	// 		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
	// 	]`
	// 	var out interface{}
	// 	out.parseJson(inputJson)
	// 	// in := []byte(inputJson);
	// 	// err := json.Unmarshal(in, &out);

	// 	// fmt.Printf("%+v", out)
	// 	// fmt.Println(err, "err")
	// }

	// func (d interface{}) parseJson(iJson string) error {
	// 	in := []byte(jJson)
	// 	err := json.Unmarshal(in, &d)
	// 	return error.Error("Sdsa")
}
