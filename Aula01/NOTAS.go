//declarar variável local 
var (
	name string 
	idade int 

)
func main(){
	name:="Graziella Linda"
	idade:= 17
	println(nome, " ", idade)
}
 
//variavel global 
var (
	nome string 
	idade int 
)
func main(){
	nome = "Grazu"
	idade = 17
	println(nome, " ", idade)
}

//Função
func Soma (a,b int) int{
	return a + b 
}

func main(){
	println: Soma(25,30)
}

//Lista
func main(){
	notas := []int()

	println(len(notas)), cap(notas))

	notas = append(notas, 35)

	println(len(notas)), cap(notas))
}

//Desasafio - substituir o indice do vetor:

var (
	slice []int 
	index uint8
	value int
)

func main() {
	slice = []int {1,2,3,4,5,6}

	fmt.Print("Insira o indice que sera modificado:")
	fmt.Scanf("$d", &index)
	fmt.Print("Insira o indice para modificacao:")
	fmt.Scanf("$d", &value)

	slice = append(slice)[:index], append ([]int(value)), slice[index+1:]...)...)

	fmt.Println(slice)
}

//Estrutura - Struct 

import(
	"fmt"
)
type Client struct{
	Nome string 
	Email string 
}

func main(){
	cliente :=Client{
		Nome:"Grazi",
		Email:"Aaaaaa@.com",
	}
	fmt.Println(cliente.Nome, cliente.Email)
}

