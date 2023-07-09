package main

import "fmt"

func main() {
	//////////////// ตัวแปร //////////////////////////
	// // การสร้างตัวแปรแบบที่ 1
	// var name string = "passapol"
	// var age int = 19
	// var score float32 = 89.92
	// // var result bool = true
	// // // การสร้างตัวแปรแบบที่ 2
	// // name1 := "passapol"
	// // age1 := 19
	// // score1 := 89.92
	// // result1 := true
	// // const number0 int = 99

	// // fmt.Println("Hi", name)
	// // fmt.Println("Name:", name)
	// // fmt.Println("Age:", age)
	// // fmt.Println("score:", score)
	// // fmt.Println("pass Exam:", result)
	// // name1 = "saf"
	// // fmt.Println("Name:", name1)
	// // fmt.Println("Age:", age1)
	// // fmt.Println("score:", score1)
	// // fmt.Println("pass Exam:", result1)
	// ///////////////// การตรวจสอบค่าข้อมูล ในตัวแปร /////////////////////////////
	// fmt.Printf("my name is %v Type : %T\n", name, name)
	// fmt.Printf("Age = %v Type : %T\n", age, age)
	// fmt.Printf("Age = %v Type : %T\n", score, score)
	// ///////////////// รับค่าข้อมูล /////////////////////////////
	// var number1 int
	// // %s:str %d:int %f:float
	// fmt.Print("input number1 :")
	// fmt.Scanf("%d", &number1)
	// fmt.Println("your number1 = ", number1)
	///////////////// if else /////////////////////////////
	var num int
	fmt.Print("input number (0,1):")
	fmt.Scanf("%d", &num)
	if num == 0 {
		fmt.Println("No!")
	} else if num == 1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("error")
	}
}
