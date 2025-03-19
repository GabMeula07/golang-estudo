package main

import "fmt"

func main()  {
	
	fmt.Println("\n Calculadora \n")
	for {
		var c = calculadora()
		if c == -1{
			break
		}
		conta(c)
		
	}
	
}

func calculadora()  int {
	var c int 
	fmt.Println("\nQual a operação que se deseja fazer?")
	fmt.Println(" [1] soma")
	fmt.Println(" [2] subtração")
	fmt.Println(" [3] multiplicação")
	fmt.Println(" [4] divisão")

	fmt.Print("\nOpção ou -1 para encerrar: ")
	fmt.Scan(&c)

	return c
}

func collectValue() int {
	var n1 int
	fmt.Print("digite o número: ")
	fmt.Scan(&n1)
	return n1
}


func conta(c int) {
	var t int
	var resultado int
	switch c {
	case 1: 
		fmt.Print("quantos números serão somados? n: ")
		fmt.Scan(&t)
		resultado = soma(t)
	case 2:
		fmt.Print("quantos números serão subtraidos? n: ")
		fmt.Scan(&t)
		resultado = sub(t)
	case 3:
		fmt.Print("quantos números serão multiplicado? n: ")
		fmt.Scan(&t)
		resultado = mult(t)	
	case 4:
		fmt.Print("quantos números serão dividos? n: ")
		fmt.Scan(&t)
		resultado = div(t)	
	}
	fmt.Println("O resultado deu: ", resultado) 
	
}

func soma(t int) int {
	var total int
	for i := 0; i < t; i++ {
		value := collectValue()
		total += value
	}

	return total
}

func sub(t int) int {
	var total int
	value := collectValue()
	total += value
	for i := 1; i < t; i++ {
		value := collectValue()
		total -= value
		
	}
	return total
}

func mult(t int) int{
	var total int
	value := collectValue()
	total += value
	for i := 1; i < t; i++ {
		value := collectValue()
		total *= value
	
	}
	return total
}

func div(t int) int{
	var total int
	value := collectValue()
	total += value
	for i := 1; i < t; i++ {
		value := collectValue()
		total /= value
	
	}
	return total
}