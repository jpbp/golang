/*
The Benefits of Named Returns
Good for Documentation (Understanding)
Named return parameters are great for documenting a function. We know what the function is returning directly from its signature, no need for a comment.

Named return parameters are particularly important in longer functions with many return values.

	func calculator(a, b int) (mul, div int, err error) {
	    if b == 0 {
	      return 0, 0, errors.New("can't divide by zero")
	    }
	    mul = a * b
	    div = a / b
	    return mul, div, nil
	}

Which is easier to understand than:

	func calculator(a, b int) (int, int, error) {
	    if b == 0 {
	      return 0, 0, errors.New("can't divide by zero")
	    }
	    mul := a * b
	    div := a / b
	    return mul, div, nil
	}

We know the meaning of each return value just by looking at the function signature: func calculator(a, b int) (mul, div int, err error)

nil is the zero value of an error. More on this later.

Less Code (Sometimes)
If there are multiple return statements in a function, you don’t need to write all the return values each time, though you probably should.

When you choose to omit return values, it's called a naked return. Naked returns should only be used in short and simple functions.
*/
package main

import (
	"errors"
	"fmt"
)

func calculator1(a, b int) (mul, div int, err error) {
    if b == 0 {
      return 0, 0, errors.New("can't divide by zero")
    }
    mul = a * b
    div = a / b
    return mul, div, nil
}
func calculator2(a, b int) (int, int, error) {
    if b == 0 {
      return 0, 0, errors.New("can't divide by zero")
    }
    mul := a * b
    div := a / b
    return mul, div, nil
}

func main(){
	a1,a2,a3 := calculator1(2,0)
	b1,b2,b3 := calculator2(2,0)
	fmt.Println(a1,a2,a3)
	fmt.Println(b1,b2,b3)
}
