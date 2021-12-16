package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"reflect"
)

const myConst1 int16 = 27 // constant can be shadowed (compare with myConst defined in main())

// enumerate constant
const (
	myConsta = iota
	myConstb //= iota
	myConstc //= iota
)

const (
	myConstd = iota
)

// enumerate constant usecase 1
const (
	_ = iota + 5 //specialListType
	dogSpecialList
	catSpecialList
	snakeSpecialList
)

// enumerate constant usecase 2
const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

// enumerate constant usecase 3

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // Note +/-/*// and bitwise operations are allowed (not function!)
	MB                    //bit shift 10 bit to left (noteL iota increment)
	GB
	TB
	PB
	EB
	ZB
	YB
)

// struct example:
// Naming convention: if PascalFormat is exported outside the package; if smallForm is declare internally in tha package
type Doctor struct {
	number     int //field name : type
	actorName  string
	companions []string
}

// struct example: Embedding (composition) (similar to inheretent)
type Animal struct {
	Name   string `required max: "100"` //`` is a tag
	Origin string
}

type Bird struct {
	Animal   // important here! embedding the strcut Animal !!!!!
	SpeedKPH float32
	CanFly   bool
}

func main() {
	fmt.Println("============primitive type============")
	// primitive type
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", c, c)

	e := float32(a) + c
	fmt.Printf("%v, %T\n", e, e)

	fmt.Println("============constant============")
	// constant
	/*
		Name convention
			global: MyConst
			local: myConst
	*/
	const myConst1 int = 42
	//myConst = 100 // change value of constant is not allowed!
	fmt.Printf("%v, %T\n", myConst1, myConst1)
	//const myConst float64 = math.Sin(1,5) // Not allowed to assign the constant with a value that complied in runtime! to compute the math.Sin need function to execute => not allowed at compile time for constant

	var myVar1 int = myConst1
	fmt.Printf("%v, %T\n", myVar1+myConst1, myVar1+myConst1)

	const myConst2 = 42
	var myVar2 float32 = 3.14
	fmt.Printf("%v, %T\n", myVar2+myConst2, myVar2+myConst2)

	// enumerate constant
	fmt.Printf("%v, %T\n", myConsta, myConsta)
	fmt.Printf("%v, %T\n", myConstb, myConstb)
	fmt.Printf("%v, %T\n", myConstc, myConstc)
	fmt.Printf("%v, %T\n", myConstd, myConstd)

	// enumerate constant usecase 1
	fmt.Println("******enumerate constant usecase 1******")
	var specialListType int
	fmt.Printf("%v\n", specialListType == dogSpecialList)
	specialListType = dogSpecialList
	fmt.Printf("%v\n", specialListType == dogSpecialList)
	fmt.Printf("%v\n", specialListType == catSpecialList)
	fmt.Printf("%v\n", dogSpecialList)

	// enumerate constant usecase 2
	fmt.Println("******enumerate constant usecase 2******")
	fileSize := 40000000000.
	fmt.Printf("the file size is %.2fGB\n", fileSize/GB)

	// enumerate constant usecase 2
	fmt.Println("******enumerate constant usecase 3******")
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("the role assess: %b\n", roles)
	fmt.Printf("Is Admin?: %v\n", isAdmin&roles == isAdmin) //isAdmin&roles bitmask
	fmt.Printf("Is HQ?: %v\n", isHeadquarters&roles == isHeadquarters)

	fmt.Println("============collection type: Array============")
	/*
	 Array needs delcaring the size of the array!
	*/
	grades := [3]int{97, 85, 93} //grades array holds 3 int element
	//grades := [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Lisa"
	students[1] = "Ahmed"
	students[2] = "Arnold"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student #1: %v\n", students[1]) // dereference the array
	fmt.Printf("Number of Student: %v\n", len(students))

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println(identityMatrix)

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix2)

	aa := [...]int{1, 2, 3}
	bb := aa  // array assignment not only just pointing to same array(addr), but a literal copy to different place
	bb[1] = 5 //this only change the element of the bb but not the aa (aa not exact same as bb (addr))
	fmt.Printf("aa = %v\n", aa)
	fmt.Printf("bb = %v\n", bb)

	// Pointer: to avoid above behavior: if a long array => deep copy => slow
	cc := [...]int{1, 2, 3}
	dd := &cc // & is important, indicates dd and cc are point to same address now
	dd[1] = 5 // dd will assess the contain as same addr as cc
	fmt.Printf("cc = %v\n", cc)
	fmt.Printf("dd = %v\n", dd)

	fmt.Println("============collection type: Slice============")
	/*
	 Slice dosent need delcaring the size of the array!
	*/
	saa := []int{1, 2, 3}
	fmt.Printf("saa = %v\n", saa)
	fmt.Printf("Length = %v\n", len(saa))
	fmt.Printf("Capacity = %v\n", cap(saa))
	// different from array: slice is a reference type => assignment will reference(point) to the same address
	sbb := saa
	sbb[1] = 5
	fmt.Printf("saa = %v\n", saa)
	fmt.Printf("sbb = %v\n", sbb)

	sa := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sb := sa[:]   // slice of all elements
	sc := sa[3:]  // slice from 4th elements to end
	sd := sa[:6]  // slice first 6 elements
	se := sa[3:6] // slice teh 4th, 5th and 6th elements
	fmt.Printf("sa = %v\n", sa)
	fmt.Printf("sb = %v\n", sb)
	fmt.Printf("sc = %v\n", sc)
	fmt.Printf("sd = %v\n", sd)
	fmt.Printf("se = %v\n", se)
	// Note: all slice will point to same array!
	sc[0] = 100 // change value @index 4
	fmt.Printf("sa = %v\n", sa)
	fmt.Printf("sb = %v\n", sb)
	fmt.Printf("sc = %v\n", sc)
	fmt.Printf("sd = %v\n", sd)
	fmt.Printf("se = %v\n", se)

	// use make function to create slices
	sma := make([]int, 3, 10) // type: []int; size: 3; capacity: 10
	fmt.Printf("sma = %v\n", sma)
	fmt.Printf("Length = %v\n", len(sma))
	fmt.Printf("Capacity = %v\n", cap(sma))

	smb := []int{}
	fmt.Printf("smb = %v\n", smb)
	fmt.Printf("Length = %v\n", len(smb))
	fmt.Printf("Capacity = %v\n", cap(smb))
	smb = append(smb, 1)
	fmt.Printf("smb = %v\n", smb)
	fmt.Printf("Length = %v\n", len(smb))
	fmt.Printf("Capacity = %v\n", cap(smb))
	smb = append(smb, 2, 3, 4, 5) // this will resize the array (doubled)
	fmt.Printf("smb = %v\n", smb)
	fmt.Printf("Length = %v\n", len(smb))
	fmt.Printf("Capacity = %v\n", cap(smb))

	ssa := []int{1, 2, 3, 4, 5}
	fmt.Printf("ssa = %v\n", ssa)
	// remove element in front
	fmt.Printf("remove element in front = %v\n", ssa[1:])
	// remove element behind
	fmt.Printf("remove element in front = %v\n", ssa[:len(ssa)-1])
	// remove element in middle!
	ssb := append(ssa[:2], ssa[3:]...) // ... is spread operation! important
	fmt.Printf("remove element in middle = %v\n", ssb)
	fmt.Printf("ssa = %v\n", ssa) // warning! the result is [1 2 4 5 5] 5 is duplicated! solution for go: have to make a loop to create a COPY of it to further work with it

	fmt.Println("============Map and Struct============")
	fmt.Println("============ Map ============")
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Florida"])
	statePopulations["Georgia"] = 1131431
	fmt.Println(statePopulations) // note: the return of the key order is not guarantee!!!
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Georgia"]) // return 0 !
	pop, ok := statePopulations["oho"]       // ok shows that if the key is in the map or not
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))

	sp := statePopulations // Its a reference type! change one will effect all!!
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	//m := map[[3]int]string{} error : slice is not a key type
	m := map[[3]int]string{} // array is a key type
	fmt.Println(m)

	fmt.Println("============ Struct ============")
	aDoctor := Doctor{
		number:    3, // suggest to add the field name instead of using index: in case of that adding field in middele will break the sequence (index of later fields)
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Printf("aDoctor with Doctor Struct: %v\n", aDoctor)
	fmt.Printf("aDoctor name is : %v\n", aDoctor.actorName)
	fmt.Printf("aDoctor's companions are : %v\n", aDoctor.companions)

	//Anonymous struct (rarely use) a temperory
	bDoctor := struct{ name string }{name: "Liz Shaw"}
	fmt.Printf("bDoctor with Anonymous Struct: %v\n", bDoctor)

	anotherDoctor := bDoctor // assigned in COPY
	anotherDoctor.name = "Tom Backer"
	fmt.Printf("anotherDoctor assigned from bDoctor with new Name: %v\n", anotherDoctor)
	fmt.Printf("bDoctor with Anonymous Struct: %v\n", bDoctor)

	anotherDoctor2 := &bDoctor // assigned to share address
	anotherDoctor2.name = "Sarah"
	fmt.Printf("anotherDoctor2 assigned to share adress from bDoctor with new Name: %v\n", anotherDoctor2)
	fmt.Printf("bDoctor with Anonymous Struct: %v\n", bDoctor)

	// struct example: Embedding (composition) (similar to inheretent)
	/* Go has no concept OOP and inheretent
	 * but use "composition" which may have the similar behavior as inherent
	 * In Go, it calls Embedding (composition)

	 declared on top:
		type Animal struct {
			Name   string
			Origin string
		}

		type Bird struct {
			Animal   // important here! embedding the strcut Animal !!!!!
			SpeedKPH float32
			CanFly   bool
		}
	 Note: the Bird and Animal not inherent relationship but the composition relationship
			which says that the ** "Bird" is not an "Animal" ** But ** "Bird" has the the "Animal" characteristic** !!!!!

	 Note -ve: we cant use the embedding interchangeably!
	*/
	abird := Bird{}
	abird.Name = "Emu"         // field "Name" from embedding Animal
	abird.Origin = "Australia" // field "Origin" from embedding Animal
	abird.SpeedKPH = 48        // field "SpeedKPH" from Bird itself
	abird.CanFly = false       // field "CanFly" from Bird itself
	fmt.Printf("abird all information: %v\n", abird)
	fmt.Printf("abird name: %v\n", abird.Name)

	bbird := Bird{
		Animal:   Animal{Name: "Bubba", Origin: "Africa"},
		SpeedKPH: 60,
		CanFly:   true,
	}

	fmt.Printf("bbird all information: %v\n", bbird)
	fmt.Printf("bbird name: %v\n", bbird.Name)

	//struct: tag
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name") // _ holds the ok (bool) which return wether the "Name" is one of teh field of t or not, here we dont care use _ to throw away of this variable
	fmt.Printf("The tag for field in Name is: %v\n", field.Tag)

	fmt.Println("============Control flow: if and switch============")
	fmt.Println("============if statement============")
	/*
		defined above

		statePopulations := make(map[string]int)
		statePopulations = map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}
	*/
	// initializer syntax
	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	// fmt.Println(pop) // error: the pop is not defined

	number := 50
	guess := 30
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100")
	} else {
		// if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Too low")
		} else if guess > number {
			fmt.Println("Too hight")
		} else if guess == number {
			fmt.Println("You got it!")
		}
	}
	fmt.Println(false)
	fmt.Println(!false)

	// Note: special case to take care
	// Compare the decimal number
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different") // will go to this case; the decimal place may not be the same
	}
	// Correct way to do it
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different") // will go to this case; the decimal place may not be the same
	}

	fmt.Println("============switch statement============")
	switch 1 { // 1 is the tag
	case 1, 5, 10:
		fmt.Println("one, five or ten")
	case 2, 4, 6:
		fmt.Println("two, four or six")
	default:
		fmt.Println("unknown")
	}

	i := 10
	switch { // tagless
	case i <= 10:
		fmt.Println("less than or equal to ten")
		// break by default
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	switch { // tagless
	case i <= 10:
		fmt.Println("less than or equal to ten")
		fallthrough
	case i <= 20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}

	// Type Switch!
	var Interface interface{} = 1
	switch Interface.(type) {
	case int:
		fmt.Println("Interface is an int")
		break
		fmt.Println("This will print too")
	case float64:
		fmt.Println("Interface is an float64")
	case string:
		fmt.Println("Interface is an string")
	default:
		fmt.Println("Interface is another type")
	}

	fmt.Println("============ Looping! ============")
	fmt.Println("********** Simple loops **********")
	fmt.Println("loop 1: ")
	for i := 0; i < 5; i++ {
		fmt.Println((i))
	}
	fmt.Println("loop 2: ")
	for i := 0; i < 5; i += 2 {
		fmt.Println((i))
	}
	fmt.Println("loop 3: ")
	for i, j := 0, 0; i < 5; i, j = i+1, j+2 {
		fmt.Printf("i = %d; j = %d \n", i, j)
	}
	fmt.Println("loop 4: ")
	for i := 0; i < 5; i++ {
		fmt.Println((i))
		// not recommend to mess around the counter lol
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*i + 1
		}
	}
	fmt.Println("loop 5: ")
	j := 0
	for ; j < 5; j++ {
		fmt.Println(j)
	}
	fmt.Println(j)
	fmt.Println("loop 6:  while loop in go")
	j = 0
	for j < 5 {
		fmt.Println(j)
		j++
	}
	fmt.Println(j)

	fmt.Println("********** Exiting early **********")

	fmt.Println("loop 7:  infinite loop (with conditional break)")
	j = 0
	for {
		fmt.Println(j)
		j++
		if j == 5 {
			break
		}
	}

	fmt.Println("loop 8:  continue")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("loop 9:  nested loop")
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
		}
	}

	fmt.Println("loop 10:  nested loop break failed")
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println((i * j))
			if i*j >= 3 {
				break // only break inner loop lol
			}
		}
	}

	fmt.Println("loop 10:  nested loop break success! with Label")
Loop:
	for i := 0; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println((i * j))
			if i*j >= 3 {
				break Loop
			}
		}
	}

	fmt.Println("********** Loop through collections **********") // only syntax for this lol
	fmt.Println("loop 11: loop over slice")
	loopSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(loopSlice)
	for k, v := range loopSlice {
		fmt.Printf("k = %d;  v = %d \n", k, v)
	}

	fmt.Println("loop 12: loop over map")
	/*
		defined above

		statePopulations := make(map[string]int)
		statePopulations = map[string]int{
			"California":   39250017,
			"Texas":        27862596,
			"Florida":      20612439,
			"New York":     19745289,
			"Pennsylvania": 12802503,
			"Illinois":     12801539,
			"Ohio":         11614373,
		}
	*/
	for state, population := range statePopulations {
		fmt.Printf("state: %s;  population = %d \n", state, population)
	}

	fmt.Println("loop 13: loop over string")
	loopString := "hello Go!"
	for k, v := range loopString {
		fmt.Printf("k = %v;  v = %v; string(v) = %v\n", k, v, string(v)) // v value is the unicode of each byte(char, note there is no data type "Char" in Go!)
	}

	fmt.Println("loop 14: loop over channel : TBD lol")

	fmt.Println("loop 15: loop over map but only want key or value")
	for state, _ := range statePopulations {
		fmt.Printf("state = %s \n", state)
	}
	for _, population := range statePopulations {
		fmt.Printf("population = %d \n", population)
	}

	fmt.Println("============ Flow Control: Defer, Panic, Recover ============")
	fmt.Println("********** Defer **********")
	// Note the defer follows the Lifo (last in first out); below will printed in the end and opposite order
	defer fmt.Println("start")
	defer fmt.Println("middle") // executed when the main has exit but before the main function return!
	defer fmt.Println("end")
	defer fmt.Println("Defer: example 1:")

	fmt.Println("Defer: practical example:")
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	// very common pattern; be careful with loop!
	defer res.Body.Close() //  i want to close but later (end of the main) , in case that i forget
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("robots: %s", robots)

	start := "start"
	defer fmt.Println(start) // will print start; the derfer will see the place to call but not the time when execute it! so it wont treated as execution after start = "end"
	defer fmt.Println("Defer: example 3:")
	start = "end"

	fmt.Println("********** Panic **********")
	fmt.Println("Panic case 1: ")
	seePanic := false
	// Error handling: GO dosent have exception!
	if seePanic {
		panicA, panicB := 1, 0
		ans := panicA / panicB // this line will return a panic - panic: runtime error: integer divide by zero
		fmt.Println(ans)
	} else {
		fmt.Println("set seePanic = true to see panic")
	}

	fmt.Println("Panic case 2: ")
	seePanic = false
	if seePanic {
		fmt.Println("start")
		panic("something bad happended")
		fmt.Println("end")
	} else {
		fmt.Println("set seePanic = true to see panic")
	}

	fmt.Println("Panic case 3 practical: ")
	seePanic = false
	if seePanic {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello Go!"))
		})
		errCase3 := http.ListenAndServe(":8080", nil)

		if errCase3 != nil {
			panic(errCase3.Error())
		}
	} else {
		fmt.Println("set seePanic = true to see panic")
	}

	fmt.Println("Panic case 4 with defer: ")
	seePanic = false
	if seePanic {
		fmt.Println("start")
		defer fmt.Println("******Panic case 4******")
		defer fmt.Println("this was derferred")
		defer fmt.Println("******Panic case 4******")
		panic("something bad happended") // Note of execute order: main function => defer statement => panic statement => handle return
		fmt.Println("end")
	} else {
		fmt.Println("set seePanic = true to see panic")
	}

	fmt.Println("********** Recover **********")
	fmt.Println("Panic and recover  case 4 with defer: ")
	seePanic = true
	if seePanic {
		fmt.Println("start")
		panicker()
		fmt.Println("end")
	} else {
		fmt.Println("set seePanic = true to see panic")
	}

	fmt.Println("============ Pointers! ============")
	fmt.Println("********** Creating pointers **********")
	fmt.Println("********** Dereferencing pointers **********")
	fmt.Println("example 1 address explain:")
	var pa int = 42
	var pb *int = &pa // * declare that pb is hold the address ;  & address of
	fmt.Println(&pa, pb)
	fmt.Println(pa, *pb) // * : used to dereference the address to get the content: here we dereference the address pd to see what the content of that address in pb
	pa = 27
	fmt.Println(pa, *pb)
	*pb = 14
	fmt.Println(pa, *pb)

	fmt.Println("example 2 work with pointer as variables (pointer arithmetic is no go for go (simplicity)) :")
	pa2 := [3]int{1, 2, 3}
	pb2 := &pa2[0] //&pa2[0] give the address where the 1st element of pa2 stroed
	pc2 := &pa2[1] //&pa2[1] give the address where the 2nd element of pa2 stroed
	//pc2 := &pa2[1]  - 4 to access next array element address is not a default allowable in Go (for simplicity)
	fmt.Printf("pa2: %v; 1st element=%v @ %p; 2nd element=%v @ %p; type of pc2 is %T\n", pa2, *pb2, pb2, *pc2, pc2, pc2) // notice that the address of pc2 is highier than pb2 by 4 bytes; think about that: int32 => 4 bytes

	fmt.Println("********** The new function **********")

	fmt.Println("********** Working with nil **********")
	fmt.Println("example 1:")
	// var ms myStruct
	// ms = myStruct{foo: 42}
	var ms *myStruct
	ms = new(myStruct) // or ms = &myStruct{foo: 42}
	// (*ms).foo = 42     // dereference the ms then access the field
	// fmt.Println(*ms)
	ms.foo = 42 // the syntax above is the complete form, the ms.foo is working as same as (*ms).foo is because the complier help to handle this
	fmt.Println(ms.foo)

	fmt.Println("********** Types with internal pointers **********")
	fmt.Println("example: array vs slice/map! literal vs reference type.\n  slice/map only holds/copy pointers to the same underlying data (will change the content that share all slice/map that point to the same reference/address)")

	fmt.Println("============ Functions! ============")
	fmt.Println("********** Basic syntax **********")
	fmt.Println("********** Parameters **********")
	fmt.Println("example 1:")
	sayMessage(`this function is to print paragraph like this
	hello Go!`)

	fmt.Println("example 2: Notice the difference between passing variable and passing in pointer! if we changen reassign the variable name")
	greeting := "Hello"
	name := "stacey"
	fmt.Println("call func sayGreetingPassVar(greeting, name string)")
	sayGreetingPassVar(greeting, name)
	fmt.Printf("name after the function call: %s\n", name)

	fmt.Println("call func sayGreetingPassPointer(greeting, name *string)")
	sayGreetingPassPointer(&greeting, &name)
	fmt.Printf("name after the function call: %s\n", name)

	fmt.Println("example 3 variadic parameters:")
	sum("i am msg!", 1, 2, 3, 4, 5)

	fmt.Println("********** Return values **********")
	fmt.Println("example 1 function with return:")
	s := sumWithReturn("i am msg!", 1, 2, 3, 4, 5) // Note the s here is a COPY of return result
	fmt.Println("Then sum is ", s)

	fmt.Println("example 3 function with return Address (Pointer):")
	swp := sumWithReturnPointer("i am msg!", 1, 2, 3, 4, 5) // Note the swp here is a pointer that referencing to the same address as result
	fmt.Println("Then sum is ", swp)

	fmt.Println("example 3 function with Name return syntax form Go:")
	s = sumWithNameReturn("i am msg!", 1, 2, 3, 4, 5) // Note the swp here is a pointer that referencing to the same address as result
	fmt.Println("Then sum is ", s)

	fmt.Println("example 4 function with multiple returns:")
	fm, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(fm)

	fmt.Println("********** Anonymous functions **********")
	// declare the function on fly
	fmt.Println("example 1 simplest case:")
	func() {
		fmt.Println("Hello Go!")
	}()

	fmt.Println("example 2 usage case for loop (Async):")
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("********** Functions as types **********")

	fmt.Println("example 3 Assign variable as function:")
	f1 := func() {
		fmt.Println("Hello Go!")
	}
	f1()

	var f2 func() = func() {
		fmt.Println("Hello Go!")
	}
	f2()

	fmt.Println("example 4 more complex case:")
	var divideFunc func(float64, float64) (float64, error)
	divideFunc = func(a, b float64) (float64, error) {
		if b == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		} else {
			return a / b, nil
		}
	}
	dResult, err := divideFunc(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		// return
	}
	fmt.Println(dResult)

	fmt.Println("********** Methods **********")
	fmt.Println("example 1 (value receiver):")
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("The new name is", g.name)

	fmt.Println("example 2 Pointer receiver:")
	g2 := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g2.greetPointerReceiver()
	fmt.Println("The new name is", g2.name)

}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() { // note the g greeter; the greet is a function that added to a type (g) that allows the g to invoke (g.greet) <= method
	// g greet is the value receiver
	fmt.Println(g.greeting, g.name)
	g.name = "" // the g here is just a copy of the greeter type, wont effect the the value outside this method
}

func (g *greeter) greetPointerReceiver() { // note the g greeter; the greet is a function that added to a type (g) that allows the g to invoke (g.greet) <= method
	// g greet is the value receiver
	fmt.Println(g.greeting, g.name)
	g.name = "" // the g here is just a copy of the greeter type, wont effect the the value outside this method
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

//variadic parameters
// Note: the variadic parameters have to be the last parameters that be passed in!
func sum(msg string, values ...int) { //  Take note the arguments parsers (parameters) here! ...int is the variadic parameter, which read all teh last argument that passed inm and wrap up to a **slice** that has name of variable we defined as "values"
	fmt.Println(msg)
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("Then sum is ", result)
}

func sumWithReturn(msg string, values ...int) int {
	fmt.Println(msg)
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("Then sum is ", result)
	return result
}

func sumWithReturnPointer(msg string, values ...int) *int {
	fmt.Println(msg)
	fmt.Println(values)
	result := 0
	for _, v := range values {
		result += v
	}
	fmt.Println("Then sum is ", result)
	return &result // teh go wont discard the return but promote it as a variable into heap memory
}

func sumWithNameReturn(msg string, values ...int) (result int) {
	fmt.Println(msg)
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	fmt.Println("Then sum is ", result)
	return
}

func sayGreetingPassVar(greeting, name string) { // note, here greeting and name are sharing same type, so w declare tgt. we also also can define separate: func sayGreetingPassVar(greeting string, name string){}
	fmt.Println(greeting, name)
	name = "Ted"
	fmt.Printf("name inside the function after re-assignment: %s\n", name)
}

func sayGreetingPassPointer(greeting, name *string) { // here teh greeting and name are pointer to string, where we pass the address that pointing to those to location
	fmt.Println(*greeting, *name) // to access the content we dereferencing them
	*name = "Ted"
	fmt.Printf("name inside the function after re-assignment: %s\n", *name)
}

func sayMessage(msg string) {
	fmt.Printf("%s\n", msg)
}

type myStruct struct {
	foo int
}

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil { // the  recover() said that teh application can be carried on the rest of the code outside of the func => it will continue up to panic and continue in main the rest
			log.Println("Error: ", err) //  log will be print @ end of func panicker
			//panic(err) // re-panic if we think that the err cant be handled
		}
	}() // () is to call the temporal defined anonymous function

	panic("something bad happended") // the panic triggered
	fmt.Println("done panicking")    // this one wont be called because the panic triggered to think the rest of lines in this function panicker shouldnet be run (Note but not after return => which are the lines outside the func in main)
}
