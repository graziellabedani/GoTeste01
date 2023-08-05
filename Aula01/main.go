package main

type Animal interface{
 Sound() string 
 Move() string 
}

type Dog struct{
	Name string 
}

func (d Dog) Sound() string{
	return d.Name + "Say, woof"
}

func (d Dog) Move() string{
	return d.name + "Running"
}

func PerfomActions (a Animal){
	fmt.Println(a.Sound())
	fmt.Println(a.Move())

}

func main() {
	dog := Dog{
		Name: "yann"
	}
	PerformActions(dog)

}
