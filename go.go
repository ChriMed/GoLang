package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"

	"strings"
	"time"
)

type Student struct {
	class string
	grade float32
}

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 5

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var waitgroup = sync.WaitGroup{} // Pretty much counters, we should add one whenever we launch a goroutine
var mutex = sync.RWMutex{}
var results = []string{} // Place to store our dbData in main
func main() {

	num := 10

	result2 := CalculateFactorial(num)
	fmt.Printf("\n The factorial of %v  result is %v \n", num, result2)
	// Variables Strings
	intNum := 200
	intNum3 := 12
	fmt.Println(intNum)
	fmt.Println(intNum3)
	Oldresult := intNum3 + intNum
	fmt.Println(Oldresult)
	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)
	var intNum2 int
	fmt.Println(intNum2)
	myVar := "text"
	fmt.Println(myVar)
	const myConst string = "const value - Functions: "

	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 1
	var result, remainder, err = intDivision(numerator, denominator)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v ", result)

	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)

	}
	/*Conditional Switch statement.
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}*/
	//Functions End Array begin
	intArr := [...]int32{1, 2, 3}

	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// Slices .. parts of Arrays

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v wuth caoacity %v", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\n The length is %v wuth caoacity %v", len(intSlice), cap(intSlice))
	var intSlice2 []int32 = []int32{8, 9}

	intSlice = append(intSlice, intSlice2...)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)
	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45, "Jason": 34}
	fmt.Println(myMap2)
	delete(myMap2, "Jason")
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Println("Invalid name")
	}

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}

	//Speed test
	var n int = 100000
	var testslice = []int{}
	var testslice2 = make([]int, 0, n)
	fmt.Printf("The total tme without preallocation: %v", timeLoop(testslice, n))
	fmt.Printf("\nThe total tme with preallocation: %v", timeLoop(testslice2, n))
	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()

	fmt.Printf("\n%v", catStr)

	//Structs and Interfaces
	//Pointers

	//var p *int32 = new(int32)
	//var i int32
	//fmt.Printf("The dereferences memory point is %v with address %v", *p, p) // p and *p values
	//p = &i
	//*p = 1
	//fmt.Printf("\nThe dereferences memory point is %v with address %v", *p, p) //&i and i values
	//Beginning Goroutines : Ways to launch multiple functions and have them execute concurrently

	//Goroutines
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitgroup.Add(1)
		go dbCall(i) //This establishes the goroutine and runs concurrently. requiresa multicore system
	}
	waitgroup.Wait()                                         //Waits until the counter goes down to zero, and all tasks are completed, then rest of code will execute
	fmt.Printf("\nTotal execution time: %v", time.Since(t0)) // Execution is not concurrent when go is not added
	fmt.Printf("\n Results are: %v\n", results)              //

	//Goroutine Channels
	//Start the program. call the goroutine process, wait for value to be used in print. goroutine channel, set channel value, function notices

	//var b = make(chan int, 5) // buffer channel 5 values

	//	var c = make(chan int) // normal channel
	//	go simpleprocess(c)
	//fmt.Println(<-c)

	var classChannel = make(chan Student)

	var classes = []string{"Class A", "Class B", "Class C", "Class D"}

	for i := range classes {
		go checkClassGrades(classes[i], classChannel)
	}

	sendReport(classChannel)

	//Complex program checks chicken sales at walmart whole foods costco

	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)

	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)

	}
	sendMessage(chickenChannel, tofuChannel)
}

// Complex Program
func checkClassGrades(classes string, classChannel chan Student) {
	for {
		time.Sleep(time.Second * 1)
		var curStudent Student
		curStudent.class = classes
		curStudent.grade = rand.Float32() * 0.99
		fmt.Printf("\nTESTING: Student from Class %v, is %v \n", curStudent.class, curStudent.grade)
		if curStudent.grade >= .90 {
			classChannel <- curStudent
			break
		}
	}
}
func sendReport(classChannel chan Student) {
	var peakStudent Student = <-classChannel
	fmt.Printf("Our best student scored a %v from class %v", peakStudent.grade, peakStudent.class)

}

// Factorial
func CalculateFactorial(fact int) int {
	term := fact
	for i := 1; i < term; i++ {
		fact *= i
	}
	return fact
}

// Complex Program
func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

// If we find a value in chickenChannel we update the value in website and execute the text statement
// If we find a value in tofuChannel we update the value in website and execute the text statement
func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("\nText sent: Found deal on chicken at %v", website)

	case website := <-tofuChannel:
		fmt.Printf("\nEmail sent: Found deal on tofu at %v", website)
	}

}

// SImple Channel Processes
func process(b chan int) {
	defer close(b) // to avoid deadlock
	for i := 0; i < 5; i++ {
		b <- i
	}
	fmt.Println("Exiting process")
}

func simpleprocess(c chan int) {
	c <- 123
}

// GOROUTINE FUNCTIONS
func dbCall(i int) {
	//Simulate DB call delay
	var delay float32 = 2000 // 2 second delay added
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	waitgroup.Done() // wait is done, decrement
}
func save(result string) {
	mutex.Lock() // When a goroutine reaches this lock it will check to see if a lock has already been performed, if it has
	//  it will wait here until the lock is released, then it will set the lock itself, execute code, remove lock
	results = append(results, result)
	mutex.Unlock()
}
func log() {
	mutex.RLock()
	fmt.Printf("\nThe current results are: %v", results)
	mutex.RUnlock()
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func printMe(printValue string) {
	fmt.Println(printValue)

}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
