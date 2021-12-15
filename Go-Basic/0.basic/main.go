package main

import (
	"fmt"
	"math"
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
	for i := 0; i < 5; i++ {
		fmt.Println((i))
	}

	fmt.Println("********** Exiting early **********")

	fmt.Println("********** Loop through collections **********")
}
