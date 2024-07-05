//const identifier [type] = value
//const PI = 3.14159
//const B = "hello"    equal to  const B string = "hello"
//
//Remark: There is a convention to name constant identifiers with all uppercase letters,
//e.g., const INCHTOCM = 2.54. This improves readability.

//var n int
//f(n + 5)   // untyped numeric constant "5" becomes typed as int, because n was int.

//const C1 = 2/3 //okay
//const C2 = getNumber() //not okay

//Numeric constants have no size or sign.
//They can be of arbitrarily high precision and do not overflow:

// const Ln2= 0.693147180559945309417232121458\
// 176568075500134360255254120680009
// const Log2E= 1/Ln2 // this is a precise reciprocal
// const BILLION = 1e9 // float constant
// const HARD_EIGHT = (1 << 100) >> 97
//We used \ (backslash) in declaring constant Ln2.
//It can be used as a continuation character in a constant.

//const BEEF, TWO, C = "meat", 2, "veg"
//const MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY int= 1, 2, 3, 4, 5, 6

// const (
// 	UNKNOWN = 0
// 	FEMALE = 1
// 	MALE = 2
//   )

//UNKNOWN, FEMALE and MALE are now aliases for 0, 1 and 2.
// Interestingly value iota can be used to enumerate the values.
// Let’s enumerate the above example with iota:

// const (
// 	UNKNOWN = iota
// 	FEMALE = iota
// 	MALE = iota
//   )

// The first use of iota gives 0. Whenever iota is used again on a new line, its value is incremented by 1;
// so UNKNOWN gets 0, FEMALE gets 1 and MALE gets 2.
// Remember that a new const block or declaration initializes iota back to 0.
// The above notation can be shortened, making no difference as:

// const (
// 	UNKNOWN = iota
// 	FEMALE
// 	MALE
//   )

package main

import "fmt"

const PI = 3.14159
const B = "hello"
const A = 2 / 3
const a, b, c = 1, 2, 3

const (
	UNKNOWN = iota
	FEMALE  = iota
	MALE    = iota
)

func main() {
	fmt.Println(PI)
	fmt.Println(B)
	fmt.Println(A)
	fmt.Println(a, b, c)
	fmt.Println(UNKNOWN)
	fmt.Println(FEMALE)
	fmt.Println(MALE)
}
