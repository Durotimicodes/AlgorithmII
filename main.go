package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// n := "IV"
	// log.Println(romanToInt(n))

	sl := []string{"joha:180"}
	an := []string{"leopard:180"}

	fmt.Println(SplitString(sl, an))

}

func SplitString(s, a []string) []string {
	var result []string
	var si, amn []string

	for i := 0; i < len(s); i++ {
		si = strings.Split(s[i], ":")
		amn = strings.Split(a[i], ":")
	}

	sVal := si[:1]
	sLen := si[1:]
	aVal := amn[:1]
	aLen := amn[1:]

	convLen, _ := strconv.Atoi(strings.Join(sLen, ""))
	convVal := strings.Join(sVal, "")
	convALen, _ := strconv.Atoi(strings.Join(aLen, ""))
	convAVal := strings.Join(aVal, "")

	log.Println(convVal, convLen)
	log.Println(convAVal, convALen)

	result = append(result, convVal, strconv.Itoa(3))
	// }

	return result

}

func RomanToInt(s string) int {

	//have a record of roman values
	myMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := myMap[s[len(s)-1]]
	log.Println("initial result", result)

	for i := len(s) - 2; i >= 0; i-- {
		if myMap[s[i]] < myMap[s[i+1]] {
			result -= myMap[s[i]]
			log.Println("2nd result", result)
		} else {
			result += myMap[s[i]]
			log.Println("final result", result)
		}
	}

	return result

}

func StrStr(haystack string, needle string) int {
	if needle == "" || haystack == needle {
		return 0
	}
	for i := 0; i <= (len(haystack) - len(needle)); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func ToWeirdCase(str string) string {
	newStr := ""
	count := 0
	for _, j := range str {
		if string(j) != " " {
			if count%2 == 0 {
				newStr += strings.ToUpper(string(j))
			} else {
				newStr += strings.ToLower(string(j))
			}
			count++
		} else {
			count = 0
			newStr += " "
		}
	}
	return newStr
}

type Stack struct {
	items []int32
}

// adds an item to the end of a slice
func (s *Stack) Push(n int32) {
	s.items = append(s.items, n)
}

// remove the value at the end of a
// slice and returns removed value
func (s *Stack) Pop() int32 {
	position := len(s.items) - 1
	valueAtPosition := s.items[position]
	s.items = s.items[:position]
	return valueAtPosition
}

// returns true if stack is empty, else false
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// checks the last item of a stack and returns it
func (s *Stack) Peek() int32 {
	value := s.items[len(s.items)-1]
	return value
}

// Main Function
func EqualStacks(h1 []int32, h2 []int32, h3 []int32) int32 {
	// Write your code here
	stack1 := Stack{}
	stack2 := Stack{}
	stack3 := Stack{}
	cumulativeHeight1 := int32(0)
	cumulativeHeight2 := int32(0)
	cumulativeHeight3 := int32(0)
	//fills up stack1 stack2 and stack3
	for i := len(h1) - 1; i >= 0; i-- {
		//0+1
		cumulativeHeight1 += h1[i]
		stack1.Push(cumulativeHeight1)
	}
	for i := len(h2) - 1; i >= 0; i-- {
		cumulativeHeight2 += h2[i]
		stack2.Push(cumulativeHeight2)
	}
	for i := len(h3) - 1; i >= 0; i-- {
		cumulativeHeight3 += h3[i]
		stack3.Push(cumulativeHeight3)
	}
	for true {
		//check if any stack is empty, then return 0 if true
		if stack1.IsEmpty() || stack2.IsEmpty() || stack3.IsEmpty() {
			return 0
		}
		cumulativeHeight1 = stack1.Peek()
		cumulativeHeight2 = stack2.Peek()
		cumulativeHeight3 = stack3.Peek()
		//checks if the three e stacks are equal when we peek
		if cumulativeHeight1 == cumulativeHeight2 && cumulativeHeight2 == cumulativeHeight3 {
			return cumulativeHeight1
		}
		//keep popping if they are not equal
		if cumulativeHeight1 >= cumulativeHeight2 && cumulativeHeight1 >= cumulativeHeight3 {
			stack1.Pop()
		} else if cumulativeHeight2 >= cumulativeHeight3 && cumulativeHeight2 >= cumulativeHeight1 {
			stack2.Pop()
		} else if cumulativeHeight3 >= cumulativeHeight2 && cumulativeHeight3 >= cumulativeHeight1 {
			stack3.Pop()
		}
	}
	return stack1.Peek()
}

func MyMain() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	result := make([]int, 0)
	reader := bufio.NewReader(os.Stdin)
	inputReader, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	n, err := strconv.Atoi(strings.TrimSpace(strings.Trim(string(inputReader), "\r\n")))
	if err != nil {
		log.Printf("An error occurred: %v", err.Error())
		return
	}
	for i := 1; i <= n; i++ {
		inputReader, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		if inputReader[0] == '1' {
			val, _ := strconv.Atoi(strings.TrimSpace(strings.Trim(string(inputReader[1:]), "\r\n")))
			result = append(result, val)
		} else if inputReader[0] == '2' {
			result = result[1:]
		} else if inputReader[0] == '3' {
			fmt.Printf("%d\n", result[0])
		}
	}
}

func extraLongFactorials1(n int32) {
	// Write your code here
	fac := big.NewInt(1)
	for ; n >= 1; n-- {
		fac = fac.Mul(fac, big.NewInt(int64(n)))
	}
	fmt.Println(fac)
}

func IsIsogram(str string) bool {
	fmt.Println(str)
	str = strings.ToLower(str)
	mapString := map[string]int{}
	for i := 0; i < len(str); i++ {
		mapString[string(str[i])] = i
	}
	if len(mapString) < len(str) {
		return false
	}
	return true
}

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{-1, -1}
}

func Autocorrect(text string) string {
	re := regexp.MustCompile(`(?i)\b(u|you+)\b`)
	return re.ReplaceAllString(text, "your client")
}
