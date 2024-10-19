/*package main

import (
	"fmt"
)*/

/*Напишите программу, которая запрашивает у пользователя число и выводит на экран сообщение: Число отрицательное и четное,
Число отрицательное и нечетное, Число положительное и четное и Число положительное и нечетное по результатам сравнения.
Да, ещё нужно подготовить сообщение: Число равно нулю при вводе 0 (Гоша, конечно, не видел таких книг, но решил предусмотреть
даже крайний случай).

func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if a > b {
		fmt.Println("Первое число больше второго")
	} else if a < b {
		fmt.Println("Второе число больше первого")
	} else {
		fmt.Println("Числа равны")
	}
}
*/

/*Напишите программу, которая запрашивает у пользователя число и выводит на экран сообщение:
Число отрицательное и четное, Число отрицательное и нечетное, Число положительное и четное и Число положительное
и нечетное по результатам сравнения. Да, ещё нужно подготовить сообщение:
 Число равно нулю при вводе 0 (Гоша, конечно, не видел таких книг, но решил предусмотреть даже крайний случай).*/

/*func main() {
	var a int
	fmt.Scan(&a)
	if a > 0 && a%2 == 0 {
		fmt.Println("Число положительное и четное")
	} else if a < 0 && a%2 == 0 {
		fmt.Println("Число отрицательное и четное")
	} else if a > 0 && a%2 != 0 {
		fmt.Println("Число положительное и нечетное")
	} else if a < 0 && a%2 != 0 {
		fmt.Println("Число отрицательное и нечетное")
	} else {
		fmt.Println("Число равно нулю")
}

}
*/

/*Напишите программу, которая запрашивает у пользователя два числа и выводит на экран сообщение:
 Оба числа положительные, Оба числа отрицательные, Одно число
положительное, а другое отрицательное по результатам их проверки. Выведите на экран сообщение: Одно из чисел равно нулю, если это так.*/
/*func main() {
	var a, b int

	fmt.Scan(&a, &b)

	if a > 0 && b > 0 {
		fmt.Println("Оба числа положительные")
	} else if a < 0 && b < 0 {
		fmt.Println("Оба числа отрицательные")
	} else if a < 0 && b > 0 || a > 0 && b < 0 {
		fmt.Println("Одно число положительное, а другое отрицательное")
	} else {
		fmt.Println("Одно из чисел равно нулю")
	}
}*/

/*Напишите программу, которая запрашивает у пользователя три числа и выводит на экран сообщение:
 Все числа равны, Два числа равны и Все числа разные по результату их сравнения.
Если ни одно из условий не выполняется, программа должна выводить сообщение: Некорректный ввод.*/

/*
package main

import (

	"fmt"

)

	func main() {
		var a, b, c int

		fmt.Scan(&a, &b, &c)
		_, d := fmt.Scan(&a, &b, &c)

		if d != 3 {
			fmt.Println("Некорректный ввод")
			return
		} else if a == b && c == a {
			fmt.Println("Все числа равны")
		} else if a == b && c != a || a == c && b != a || b == c && a != b {
			fmt.Println("Два числа равны")
		} else if a != b && c != a {
			fmt.Println("Все числа разные")
		}
	}
*/
/*package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Считываем значения и проверяем их количество
	d, err := fmt.Scan(&a, &b, &c)
	if err != nil || d < 3 || d > 3 {
		fmt.Println("Некорректный ввод")
		return
	}

	// Проверяем равенство чисел
	if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if (a == b && c != a) || (a == c && b != a) || (b == c && a != b) {
		fmt.Println("Два числа равны")
	} else {
		fmt.Println("Все числа разные")
	}
}
*/

/*
func main() {
	var a, b, c int

	// Считываем три целых числа
	n, err := fmt.Scanf("%d %d %d", &a, &b, &c)
	if err != nil || n != 3 {
		fmt.Println("Некорректный ввод")
		return
	}

	// Проверяем равенство чисел
	if a == b && b == c {
		fmt.Println("Все числа равны")
	} else if (a == b && c != a) || (a == c && b != a) || (b == c && a != b) {
		fmt.Println("Два числа равны")
	} else if a != b && a != c && b != c {
		fmt.Println("Все числа разные")
	} else {
		fmt.Println("Некорректный ввод!")
	}
}*/

/*
Напишите программу, которая запрашивает у пользователя число и выводит на экран сообщение: Число меньше 10,

	Число меньше 100 (если оно больше или равно 10, но меньше 100), Число меньше 1000 (если оно больше или равно 100, но меньше 1000) и
	Число больше или равно 1000 по результатам оценки числа. Если число отрицательное, программа должна выводить сообщение: Некорректный ввод.
*/

/*
	func main() {
		var a int

/		fmt.Scan(&a)

		if a > -1 && a < 10 {
			fmt.Println("Число меньше 10 ")

		} else if a >= 10 && a < 100 {
			fmt.Println("Число меньше 100 ")

		} else if a >= 100 && a < 1000 {
			fmt.Println("Число меньше 1000 ")

		} else if a >= 1000 {
			fmt.Println("Число больше или равно 1000")
		} else {
			fmt.Println("Некорректный ввод")

		}
	}
*/

/*
	func main() {
		var input int
		fmt.Scanln(&input)
		i := 1
		c := 0
		for i < input+1 {
			if (i%1 == 0 || i%2 == 0 || i%4 == 0 || i%6 == 0 || i%7 == 0 || i%8 == 0 || i%9 == 0) && (i%3 != 0 && i%5 != 0) {
				c = c + i

			}

			i++
		}
		fmt.Println(c)
	}
*/
/*
package main

import (
	"fmt"
	"math"
)

func SqRoots() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)

	discriminant := b*b - 4*a*c
	// Два различных вещественных корня
	if discriminant > 0 {
		nvalue1 := ((-b + math.Sqrt(discriminant)) / (2.0 * a))
		nvalue2 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
		fmt.Printf("%f %f ", nvalue2, nvalue1)
	} else if discriminant == 0 {
		nvalue3 := (-b - math.Sqrt(discriminant)) / (2.0 * a)
		fmt.Printf("%f", nvalue3)
	} else {
		fmt.Println("0 0")
	}
}


func main() {
	SqRoots()

}
*/

/*
func Add(a, b float64) float64 {
	f := a + b
	return f
}
func Multiply(j, u float64) float64 {
	g := j * u
	return g
}
func PrintNumbersAscending(n int) int {
	for i := 1; i <= n; i++ {
		if i == n {
			fmt.Print(i) // не добавляем пробел после последнего числа
		} else {
			fmt.Print(i, " ")
		}
	}
	return n
}*/

/*
func main() {
	PrintNumbersAscending(25)
}
*/
/*func SumDigitsRecursive(n int) int {
	var v, c int
	v = 0
	for n > 0 {
		c = n % 10
		v += c
		n = n / 10
	}

	return v
}


func main() {
	fmt.Println(SumDigitsRecursive(123))
}*/
/*
func IsPowerOfTwoRecursive(N int) {
	var n int
	n = N
	for N > 2 {
		N /= 2
	}
	if N == 2 || n == 1 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}


func main() {
	IsPowerOfTwoRecursive(8)
}
*/
/*
func CalculateSeriesSum(n float64) float64 {
	var h, i float64
	i = 1
	h = 0
	for i != n+1 {
		h += 1 / i
		i++
	}
	return h
}


func main() {
	fmt.Println(CalculateSeriesSum(6))
}
*/
/*
func FindIntersection(k1, b1, k2, b2 float64) (float64, float64) {
	var a, b float64
	if k1 == k2 {
		return math.NaN(), math.NaN()
	}
	a = (b1 - b2) / (k2 - k1)
	b = k1*a + b1
	return a, b

}


func main() {
	fmt.Println(FindIntersection(5, 20, 5, 10))
}
*/
/*
func CalculateDigitalRoot(n int) int {
	var v, c int
	v = 0
	for n > 0 {
		c = n % 10
		v += c
		n = n / 10
	}
	var k, f int
	k = 0
	for v > 0 {
		f = v % 10
		k += f
		v = v / 10
	}
	return k

}


func main() {
	fmt.Println(CalculateDigitalRoot(12345))
}
*/
/*
func FiveSteps(array [5]int) [5]int {
	var reversedArray [5]int
	for i := 0; i < len(array); i = i + 1 {
		reversedArray[i] = array[len(array)-1-i]

	}
	return reversedArray
}


func main() {
	input := [5]int{1, 2, 3, 4, 5}
	output := FiveSteps(input)
	for i := 0; i < len(output); i++ {
		fmt.Printf("%d ", output[i])
	}
}
*/
/*
func ThirdElementInArray(array [6]int) int {
	return array[2]

}


func main() {
	input := [6]int{1, 1, 1, 1, 1, 1}
	output := ThirdElementInArray(input)
	fmt.Print(output)
}
*/
/*
func FindMinMaxInArray(array [10]int) (int, int) {
	min := array[0]
	max := array[0]
	for _, value := range array {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}

	return max, min
}


func main() {
	input := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	max, min := FindMinMaxInArray(input)

	fmt.Print(max, min)
}
*/
/*
func PrettyArrayOutput(array [9]string) {
	c := 1
	for i := 0; i < 9; i++ {
		if i < 7 {
			fmt.Println(c, "я уже сделал:", array[i])
			c += 1
		} else if i >= 7 {
			fmt.Println(c, "не успел сделать:", array[i])
			c += 1
		}
	}
}


func main() {
	input := [9]string{
		"проснуться",
		"позавтракать",
		"сходить в школу",
		"пообедать",
		"погулять с друзьями",
		"сделать домашнюю работу",
		"попрограммировать на Go",
		"поужинать",
		"лечь спать",
	}
	PrettyArrayOutput(input)
}
*/

/*Напишите программу, которая запрашивает у пользователя число и выводит
на экран 10 чисел Фибоначчи, начиная с введённого пользователем, если оно является таким числом, или с ближайшего большего.
func main() {
	var input int
	n := input
	fmt.Scan(&n)
	var list = []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711}
	var return_i []int

	for i := 0; i < len(list); i++ {
		if list[i] >= n {
			return_i = append(return_i, list[i])
		}
		if len(return_i) >= 10 {
			break
		}
	}
	printSlice(return_i)

}
func printSlice(return_i []int) {
	for _, value := range return_i {
		fmt.Printf("%v\n", value)
	}
}
*/
/*
func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	ret := make([]int, 0, n)
	if n < 0 {
		return ret, nil
	}
	for i := 0; i < len(nums); i++ {
		if len(ret) >= n {
			break
		}
		if nums[i] < limit {
			ret = append(ret, nums[i])
	}

	return ret, nil
}
*/

/*
import (
	"errors"
	"fmt"
)

func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if n < 0 {
		return nil, errors.New("true")
	}

	ret := make([]int, 0, n)
	for _, num := range nums {
		if len(ret) >= n {
			break
		}
		if num < limit {
			ret = append(ret, num)
		}
	}
	return ret, nil
}

func main() {
	nums := []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}
	var limit = 3
	var n = 5
	fmt.Println(UnderLimit(nums, limit, n))
}
*/
/*
func Clean(nums []int, x int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == x {
			nums = append(nums[:i], nums[i+1:]...)
			if len(nums) == 1 && nums[0] == x {
				return nums[:len(nums)-1]
			}
		} else if len(nums) == 0 {
			return nums

		}
	}
	return nums

}


func main() {
	nums := []int{3, -7}
	var x = -7
	fmt.Println(Clean(nums, x))
}
*/
/*
func SliceCopy(nums []int) []int {
	var reversedArray []int
	reversedArray = append(reversedArray, nums...)
	return reversedArray
}*/
/*
func Join(nums1, nums2 []int) []int {
	var A []int

	for i := 0; i < len(nums1); i++ {
		A = append(A, nums1[i])
	}
	for i := 0; i < len(nums2); i++ {
		A = append(A, nums2[i])
	}
	fmt.Println(len(cap(A)))
	return A
}

func main() {
	nums1 := []int{}
	nums2 := []int{4, 5, 6, 7}
	fmt.Println(Join(nums1, nums2))
}
*/
/*
func Mix(nums []int) []int {
	z := make([]int, 0, len(nums)/2)
	u := make([]int, 0, len(nums)/2)
	g := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		if i < len(nums)/2 {
			z = append(z, nums[i])
		} else {
			u = append(u, nums[i])
		}
	}

	for i := 0; i < len(z); i++ {
		g = append(g, z[i])
		g = append(g, u[i])
	}
	return g
}
*/
/*
func main() {
	nums := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(Mix(nums))
}
*/
/*
func ConcatenateStrings(str1, str2 string) string {
	fullName := str1 + " " + str2
	return fullName
}
func main() {
	fmt.Println(ConcatenateStrings("Irina", "Khomyakova"))
}*/
/*
func isLatin(input string) bool {
	var a bool
	for _, r := range input {
		if unicode.In(r, unicode.Latin) {
			a = true
		} else {
			a = false
		}
	}
	return a
}
func main() {
	word := "こんにちは"
	fmt.Println(isLatin(word))
}
*/
/*
func IsPalindrome(input string) bool {
	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPalindrome(" a b c       cba  "))
}
*/
/*
func CountingSort(contacts []string) map[string]int {
	swappedMap := make(map[string]int)
	c := 0
	c1 := 0
	c2 := 0
	c3 := 0
	c4 := 0
	c5 := 0
	c6 := 0
	c7 := 0

	for i := 0; i < len(contacts); i++ {

		if contacts[i] == "1" {
			c++
		}
		if contacts[i] == "2" {
			c1++
		}
		if contacts[i] == "3" {
			c2++
		}
		if contacts[i] == "4" {
			c3++
		}
		if contacts[i] == "5" {
			c4++
		}
		if contacts[i] == "6" {
			c5++
		}
		if contacts[i] == "7" {
			c6++
		}
		if contacts[i] == "8" {
			c7++
		}

	}
	if c > 0 {
		swappedMap["1"] = c
	}
	if c1 > 0 {
		swappedMap["2"] = c1
	}
	if c2 > 0 {
		swappedMap["3"] = c2
	}
	if c3 > 0 {
		swappedMap["4"] = c3
	}
	if c4 > 0 {
		swappedMap["5"] = c4
	}
	if c5 > 0 {
		swappedMap["6"] = c5
	}
	if c6 > 0 {
		swappedMap["7"] = c6
	}
	if c7 > 0 {
		swappedMap["8"] = c7
	}
	return swappedMap

}
func main() {
	s := []string{"1", "2", "3", "2", "3", "3"}
	fmt.Println(CountingSort(s))
}
*/
/*
func DeleteLongKeys(m map[string]int) map[string]int {
	newMap := make(map[string]int)
	for key, value := range m {
		if len(key) >= 6 {
			newMap[key] = value
		}
	}
	return newMap
}

func main() {
	m := map[string]int{"aaaaaaa": 1, "bb": 2, "cccccc": 3}
	fmt.Println(DeleteLongKeys(m))
}
*/
/*
import (
	"fmt"
)

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) Print() {
	fmt.Println("Name:", p.name)

	fmt.Println("Age:", p.age)

	fmt.Println("Address:", p.address)
}
func main() {
	man := Person{name: "Anton", age: 31, address: "Krasnogorsk"}
	man.Print()
}
*/
/*
package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (p Employee) CalculateTotalSalary() {
	fmt.Println("Employee:", p.name)

	fmt.Println("Position:", p.position)

	fmt.Printf("Total Salary: %.2f", p.salary+p.bonus)
}
func main() {
	man := Employee{name: "Anton", position: "product manager", salary: 100.0, bonus: 100.0}
	man.CalculateTotalSalary()
}
*/
/*
package main

import "fmt"

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (p Student) IsExcellentStudent() bool {
	if float64(p.solvedProblems)*p.scoreForOneTask >= p.passingScore {
		return true
	}
	return false

}
func main() {
	student := Student{name: "Gosha", solvedProblems: 30, scoreForOneTask: 10.0, passingScore: 290.0}
	fmt.Print(student.IsExcellentStudent())
}
*/
/*
package main

import (
	"fmt"
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (p Task) IsOverdue() bool {
	return time.Now().After(p.deadline)
}

func (p Task) IsTopPriority() bool {
	return p.priority == 4 || p.priority == 5
}
func main() {
	task := Task{summary: "Make Yandex Lyceum", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 4}
	fmt.Println(task.IsOverdue())
	fmt.Print(task.IsTopPriority())
}
*/
/*
package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Circle struct {
	radius float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (s Circle) Area() float64 {
	return math.Pi * math.Pow(s.radius, 2)
}
func (f Rectangle) Area() float64 {
	return f.width * f.height
}
func main() {
	figure_1 := Circle{radius: 1.0}
	fmt.Println(figure_1.Area())
	figure_2 := Rectangle{width: 57.2, height: 10.2}
	fmt.Println(figure_2.Area())
}
*/
/*
package main

import "fmt"

type Animal interface {
	MakeSound() string
}
type Dog struct {
}
type Cat struct {
}

func (d Dog) MakeSound() {
	fmt.Println("Гав!")

}
func (f Cat) MakeSound() {
	fmt.Println("Мяу!")
}
func main() {
	cat1 := Cat{}
	cat1.MakeSound()
	dog1 := Dog{}
	dog1.MakeSound()
}
*/
/*
package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	LogLevelError LogLevel = "ERROR"
	LogLevelInfo  LogLevel = "INFO"
)

type Log struct {
	Level LogLevel
}

func (l *Log) Log(message string) {
	switch l.Level {
	case LogLevelError:
		fmt.Printf("ERROR: %s\n", message)
	case LogLevelInfo:
		fmt.Printf("INFO: %s\n", message)
	default:
		fmt.Printf("UNKNOWN LEVEL: %s\n", message)
	}
}

func main() {
	errorLog := &Log{Level: Error}
	errorLog.Log("This is an error message")
}
*/
/*
package main

import "strconv"

func ConcatStringsAndInt(str1, str2 string, num int) string {
	c := strconv.Itoa(num)
	return str1 + str2 + c
}
*/
/*Напишите функцию DivideIntegers(a, b int) (float64, error), которая возвращает результат деления первого числа на второе. .
Если второе число равно нулю, функция должна возвращать в качестве
ответа 0 и сообщение об ошибке (division by zero is not allowed). Если второе число не равно нулю – верните в качестве ошибки nil.*/
/*
package main

import "fmt"

func DivideIntegers(a, b int) (float64, error) {
	var floatNum float64 = float64(a)
	var floatNum1 float64 = float64(b)
	s := floatNum / floatNum1
	if a == 0 && b == 0 {
		return 0, fmt.Errorf("division by zero is not allowed")
	} else if b == 0 {
		return 0, fmt.Errorf("division by zerodivision by zero is not allowed")
	}
	return s, nil
}
func main() {
	fmt.Println(DivideIntegers(10, 2))
}
*/
/*
package main

import (
	"errors"
	"fmt"
)

func GetCharacterAtPosition(str string, position int) (rune, error) {
	g := []rune(str)
	if position < 0 || position >= len(g) {
		return '\u0000', errors.New("position out of range")
	}

	return g[position], nil
}

func main() {
	fmt.Println(GetCharacterAtPosition("Оле ола вперёд Спартак Москва", 0))
}
*/
/*
Напишите функцию Factorial(n int) (int, error) для нахождения факториала без ошибок. Функция получает на вход целое число и выводит его факториал.
 Если пользователь ввёл отрицательное число, программа должна вернуть сообщение об ошибке (factorial is not defined for negative numbers).*/
/*package main

import "errors"

func Factorial(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("factorial is not defined for negative numbers")
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result, nil
}
*/
/*
Гоша хочет научиться переводить в двоичную систему положительные целые числа. Напишите функцию IntToBinary(num int) (string, error),
 которая принимает на вход целое число и возвращает его двоичное
представление. Если пользователь ввёл отрицательное число, программа должна возвращать сообщение об ошибке (negative numbers are not allowed).
*/
/*
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func IntToBinary(num int) (string, error) {
	g := strconv.FormatInt(int64(num), 2)
	if num < 0 {
		return "", errors.New("negative numbers are not allowed")
	}
	return g, nil
}
func main() {
	fmt.Println(IntToBinary(-20))
}
*/

/*Напишите функцию SumTwoIntegers(a, b string) (int, error), которая получает две строки и, если это целые числа, возвращает их сумму.
Если пользователь ввёл данные не целого типа, функция должна вернуть сообщение об ошибке (invalid input, please provide two integers).*/
/*
package main

import (
	"errors"
	"fmt"
	"strconv"
)

func SumTwoIntegers(a, b string) (int, error) {
	i, err1 := strconv.Atoi(a)
	f, err2 := strconv.Atoi(b)
	cum := i + f
	if err1 != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	if err2 != nil {
		return 0, errors.New("invalid input, please provide two integers")
	}
	return cum, nil
}
func main() {
	fmt.Println(SumTwoIntegers("0,2", "12"))
}
*/
/*
package main

import (
	"fmt"
	"time"
)

func TimeDifference(start, end time.Time) time.Duration {
	diff := end.Sub(start)
	return diff

}
func main() {
	start := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	end := time.Date(2023, 10, 23, 4, 41, 49, 0, time.UTC)
	diff := TimeDifference(start, end)
	fmt.Println(diff) // Output: 2h0m0s
}
*/
/*Создайте функцию FormatTimeToString(timestamp time.Time, format string) string, которая форматирует заданное значение time.Time как строку
в определенном формате. Функция должна принимать значение time.Time и строку формата и возвращать форматированную строку.*/
/*
package main

import "time"

func FormatTimeToString(timestamp time.Time, format string) string {
	return timestamp.Format(format)
}
*/
/*
Напишите функцию ParseStringToTime(dateString, format string) (time.Time, error), которая анализирует строку в определенном формате
и преобразует ее в значение time.Time. Функция должна принимать строку и строку формата, возвращая проанализированное значение time.Time.
*/
/*
package main

import (
	"fmt"
	"time"
)

func ParseStringToTime(dateString, format string) (time.Time, error) {
	u, err := time.Parse(format, dateString)
	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	}
	return u, nil
}
func main() {
	dateString := "2023-10-23 02:41:49"
	format := "2006-01-02 15:04:05"
	result, err := ParseStringToTime(dateString, format)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
*/
/*
Напишите функцию TimeAgo(pastTime time.Time) string,который вычисляет время, прошедшее с момента заданного значения time.Time,
 и возвращает удобочитаемую строку, указывающую, как давно это было.
*/
/*
Напишите функцию NextWorkday(start time.Time) time.Time, которая вычисляет дату следующего
рабочего дня (исключая выходные). Учитывая дату начала, функция должна возвращать дату следующего рабочего дня.*/
/*
package main

import (
	"fmt"
	"time"
)

func NextWorkday(startTime time.Time) time.Time {
	day := startTime

	day = day.AddDate(0, 0, 1)
	if day.Weekday() != time.Saturday || day.Weekday() != time.Sunday {
		if day.Weekday() == time.Saturday {
			day = day.AddDate(0, 0, 2)

		}
		if day.Weekday() == time.Sunday {
			day = day.AddDate(0, 0, 1)

		}
	}
	return day
}
func main() {
	start := time.Date(2023, time.October, 6, 0, 0, 0, 0, time.UTC) // Friday
	nextWorkday := NextWorkday(start)
	fmt.Println(nextWorkday)
}
*/
/*
package main

func Length(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "short"
	case a < 100:
		return "long"
	}
	return "very long"
}
*/
/*
package main

func Multiply(a, b int) int {
	return a * b
}
*/
/*
package main

import "unicode"

func DeleteVowels(s string) string {
	var result string
	for i := 0; i < len(s); i++ {
		switch unicode.ToLower(rune(s[i])) {
		case 'a':
			continue
		case 'e':
			continue
		case 'i':
			continue
		case 'o':
			continue
		case 'u':
			continue
		}
		result += string(s[i])
	}
	return result
}
*/

package main

import (
	"errors"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("invalid utf8")

func GetUTFLength(input []byte) (int, error) {

	if !utf8.Valid(input) {
		return 0, ErrInvalidUTF8
	}

	return utf8.RuneCount(input), nil
}
