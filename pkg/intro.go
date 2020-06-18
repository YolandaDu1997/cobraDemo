package pkg

import "fmt"

func Show(name string, age int) {
	fmt.Printf("My name is %s. I am %d years old.\n", name, age)
}

func ShowSub(name string) {
	fmt.Printf("Sub command description:%s\n", name)
}
