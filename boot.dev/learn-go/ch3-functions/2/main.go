/*
Multiple Parameters
When multiple arguments are of the same type, and are next to each other in the function signature, the type only needs to be declared after the last argument.

Here are some examples:

func addToDatabase(hp, damage int) {
  // ...
}

func addToDatabase(hp, damage int, name string) {
  // ?
}

func addToDatabase(hp, damage int, name string, level int) {
  // ?
}
*/

package main

import "fmt"

func addToDatabase1(hp, damage int) {
	fmt.Printf("O valor do seu hp é: %v, seu damage atual é de: %v\n",hp,damage)
}

func addToDatabase2(hp, damage int, name string) {
	fmt.Printf("O valor do seu hp é: %v, seu damage atual é de: %v, seu nome é de: %s\n",hp,damage,name)
}

func addToDatabase3(hp, damage int, name string, level int) {
	fmt.Printf("O valor do seu hp é: %v, seu damage atual é de: %v, seu nome é de: %s, e seu level é: %v\n",hp,damage,name,level)
}

func main(){
	hp := 100
	damage := 50
	name := "JP"
	level := 5
	addToDatabase1(hp,damage)
	addToDatabase2(hp,damage,name)
	addToDatabase3(hp,damage,name,level)
}
