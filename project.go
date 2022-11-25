package main

import "fmt"

func main() {
	//fmt.Println("Hello World ----------------------------------------------------------")

	//var age = 12
	//var age2 int = 16
	//fmt.Println(age)
	//var num float64 = 2.3453
	//fmt.Println(num, age)

	//--//var res float64 = age + num -- invalid operation: age + num (mismatched types int and float64)
	//var res int = age + age2
	//fmt.Println(res)

	//const pi = 3.14 ---- constant

	//var web string = "itProper"
	//fmt.Println("web = " + web + "\n Babe")
	//fmt.Println(len(web))
	//fmt.Printf("%f \n", num)
	//fmt.Printf("%.2f \n", num)
	//fmt.Printf("%f.2 \n", num)

	//var isDone bool = true
	//fmt.Printf("this is %t \n", isDone)

	//var age = 8
	//if age < 5 {
	//	fmt.Println("Вам пора в детский сад")
	//} else if age == 5 {
	//	fmt.Println("Вам пора идти в прескул")
	//
	//} else if (age > 5) && (age <= 18) {
	//	var grade = age - 5
	//	fmt.Println("Пора идти в", grade, "класс")
	//} else {
	//	fmt.Println("Вам пора в универ")
	//}

	//var age = 10
	//switch age {
	//case 5:
	//	fmt.Println("Вам 5 лет")
	//case 15:
	//	fmt.Println("Вам 15 лет")
	//case 51:
	//	fmt.Println("Вам 51 лет")
	//default:
	//	fmt.Println("Вам неизвестно сколько лет")
	//}

	//var i = 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i++
	//}

	//for i := 0; i <= 5; i++ {
	//	fmt.Println(i)
	//}

	//var arr [3]int
	//arr[0] = 45
	//arr[1] = 97
	//arr[2] = 76
	//fmt.Println(arr[0])

	//nums := [3]float64{3.23, 5.23, 67.42}
	//for i, value := range nums {
	//	fmt.Println(value, i)
	//}

	//webSites := make(map[string]float64)
	//webSites["itProper"] = 0.8
	//webSites["yandex"] = 0.99
	//fmt.Println(webSites)
	//fmt.Println(webSites["itProper"])

	//fmt.Println(summ(1, 2))
	//
	//fmt.Println(summ2(1, 2))

	//var num = 3
	//multiple := func() int {
	//	num *= 2
	//	return num
	//}
	//fmt.Println(multiple())

	//defer two() //откладывает функцию на потом
	//one()

	//var x = 0
	//var y = 0
	//pointer(x)
	//pointery(&y)
	//fmt.Println(x)
	//fmt.Println(y)

	bob := Cats{"Bob", 7, 0.87}
	fmt.Println("Bob age is", bob.age)
	fmt.Println("Bob function is", bob.test())
}

func summ(n1 int, n2 int) int {
	var res int
	res = n1 + n2
	return res
}

func summ2(n1 int, n2 int) (int, int) {
	var res int
	res = n1 + n2
	return res, n1 * n2
}

func one() {
	fmt.Println("1")
}
func two() {
	fmt.Println("2")
}

func pointer(x int) {
	x = 2
}

func pointery(x *int) {
	*x = 2
}

type Cats struct {
	name      string
	age       int
	happiness float64
}

func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness
}
