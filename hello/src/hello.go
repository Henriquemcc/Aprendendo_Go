package main

import (
	"fmt"
)

func main() {

	//Exibindo mensagem de bem vindo
	fmt.Println("Bem vindo ao monitor de websites.")

	//Exibindo mensagem pedindo para o usuario escolher o que deseja fazer
	fmt.Println("O que deseja fazer: ")
	fmt.Println("0 - Sair do programa")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir os logs")

	//Lendo a entrada do usuario
	var comando int8
	//fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo os logs...")
	} else if comando != 0 {
		fmt.Println("Comando Inválido!")
	}

}
