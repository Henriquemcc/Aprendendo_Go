package main

import "fmt"

//Essa estrutura serve para guardar os dados da conta corrente do usuario.
type contaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

//Este metodo serve para realizar o saque na conta corrente do usuario.
//Parametro: valorDoSaque: Valor da quantidade de dinheiro que sera sacada.
//Retorno: bool: Valor booleano indicando se o saque ocorreu com sucesso.
func (c *contaCorrente) Sacar(valorDoSaque float64) bool {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	sucessoSaque := false

	if podeSacar {
		c.saldo -= valorDoSaque
		sucessoSaque = true
	}

	return sucessoSaque
}

//Este metodo eh o metodo principal, onde o programa comeca a ser executado
func main() {

	//Criando uma conta corrente para o personagem ficticio chamado Ze
	contaDoZe := contaCorrente{}
	contaDoZe.titular = "Zé"
	contaDoZe.saldo = 134217

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.titular, ":", contaDoZe.saldo)

	//Sacando dinheiro da conta do ze
	fmt.Println("Saque da conta da ze realizado com sucesso:", contaDoZe.Sacar(1000000))

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.titular, ":", contaDoZe.saldo)

	//Criando uma conta corrente para o personagem ficticio chamado Joao da Silva
	contaDoJoao := contaCorrente{}
	contaDoJoao.titular = "Joao da Silva"
	contaDoJoao.saldo = 1024

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.titular, ":", contaDoJoao.saldo)

	//Sacando dinheiro da conta do Joao
	fmt.Println("Saque da conta da ze realizado com sucesso:", contaDoJoao.Sacar(1000000))

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.titular, ":", contaDoJoao.saldo)

	//Imprimindo as contas do Ze e do Joao
	fmt.Println(contaDoZe)
	fmt.Println(contaDoJoao)
}
