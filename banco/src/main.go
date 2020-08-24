package main

import (
	"banco/contas"
	"fmt"
)

//Este metodo eh o metodo principal, onde o programa comeca a ser executado
func main() {

	//Criando uma conta corrente para o personagem ficticio chamado Ze
	contaDoZe := contas.ContaCorrente{}
	contaDoZe.SetTitular("Zé")
	contaDoZe.SetSaldo(134217)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.GetTitular(), ":", contaDoZe.GetSaldo())

	//Sacando dinheiro da conta do ze
	sucesso, mensagemDeErro, novoSaldo := contaDoZe.Sacar(1000000)
	fmt.Println("Saque da conta do", contaDoZe.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.GetTitular(), ":", contaDoZe.GetSaldo())

	//Depositando dinheiro na conta do ze
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Depositar(500)
	fmt.Println("Deposito na conta do", contaDoZe.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.GetTitular(), ":", contaDoZe.GetSaldo())

	//Criando uma conta corrente para o personagem ficticio chamado Joao da Silva
	contaDoJoao := contas.ContaCorrente{}
	contaDoJoao.SetTitular("Joao da Silva")
	contaDoJoao.SetSaldo(1024)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.GetTitular(), ":", contaDoJoao.GetSaldo())

	//Sacando dinheiro da conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Sacar(100)
	fmt.Println("Saque da conta do", contaDoJoao.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.GetTitular(), ":", contaDoJoao.GetSaldo())

	//Depositando dinheiro na conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Depositar(-13)
	fmt.Println("Deposito na conta do", contaDoJoao.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.GetTitular(), ":", contaDoJoao.GetSaldo())

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(500, &contaDoJoao)
	fmt.Println("Transferencia da conta do", contaDoZe.GetTitular(), "para a conta do", contaDoJoao.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", contaDoZe.GetTitular(), ":", novoSaldo)

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(9999999999, &contaDoJoao)
	fmt.Println("Transferencia da conta do", contaDoZe.GetTitular(), "para a conta do", contaDoJoao.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", contaDoZe.GetTitular(), ":", novoSaldo)

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(-20, &contaDoJoao)
	fmt.Println("Transferencia da conta do", contaDoZe.GetTitular(), "para a conta do", contaDoJoao.GetTitular(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", contaDoZe.GetTitular(), ":", novoSaldo)

	//Imprimindo as contas do Ze e do Joao
	fmt.Println(contaDoZe)
	fmt.Println(contaDoJoao)
}
