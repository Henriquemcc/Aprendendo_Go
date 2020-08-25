package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
	"time"
)

//Este metodo eh o metodo principal, onde o programa comeca a ser executado
func main() {

	//Criando um titular chamado ze
	ze := clientes.Titular{}
	ze.SetNome("Zé")
	ze.SetSobrenome("da Silva")
	ze.SetTelefone("+55 (11) 16234-1245")
	ze.SetEmail("ze@gmail.com")
	ze.SetProfissao("Empresário")
	ze.SetCpf("450.040.015-02")
	nascimento := time.Date(1900, time.December, 18, 0, 0, 0, 0, time.UTC)
	ze.SetDataDeNascimento(&nascimento)

	//Criando um titular chamado joao
	joao := clientes.Titular{}
	joao.SetNome("Joao")
	joao.SetSobrenome("da Silva")
	joao.SetTelefone("+55 (11) 12555-2345")
	joao.SetEmail("joao@gmail.com")
	joao.SetProfissao("Investidor")
	joao.SetCpf("625.733.383-04")
	nascimento = time.Date(1900, time.June, 5, 0, 0, 0, 0, time.UTC)
	joao.SetDataDeNascimento(&nascimento)

	//Criando uma conta corrente para o Ze
	contaDoZe := contas.ContaCorrente{}
	contaDoZe.SetTitular(&ze)
	contaDoZe.Depositar(134217)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", ze.GetNome(), ":", contaDoZe.GetSaldo())

	//Sacando dinheiro da conta do ze
	sucesso, mensagemDeErro, novoSaldo := contaDoZe.Sacar(1000000)
	fmt.Println("Saque da conta do", ze.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", ze.GetNome(), ":", contaDoZe.GetSaldo())

	//Depositando dinheiro na conta do ze
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Depositar(500)
	fmt.Println("Deposito na conta do", ze.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", ze.GetNome(), ":", contaDoZe.GetSaldo())

	//Criando uma conta corrente para o personagem ficticio chamado Joao da Silva
	contaDoJoao := contas.ContaCorrente{}
	contaDoJoao.SetTitular(&joao)
	contaDoJoao.Depositar(1024)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", joao.GetNome(), ":", contaDoJoao.GetSaldo())

	//Sacando dinheiro da conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Sacar(100)
	fmt.Println("Saque da conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", joao.GetNome(), ":", contaDoJoao.GetSaldo())

	//Depositando dinheiro na conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Depositar(-13)
	fmt.Println("Deposito na conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", joao.GetNome(), ":", contaDoJoao.GetSaldo())

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(500, &contaDoJoao)
	fmt.Println("Transferencia da conta do", ze.GetNome(), "para a conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", ze.GetNome(), ":", novoSaldo)

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(9999999999, &contaDoJoao)
	fmt.Println("Transferencia da conta do", ze.GetNome(), "para a conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", ze.GetNome(), ":", novoSaldo)

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(-20, &contaDoJoao)
	fmt.Println("Transferencia da conta do", ze.GetNome(), "para a conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", ze.GetNome(), ":", novoSaldo)

	//Imprimindo as contas do Ze e do Joao
	fmt.Println("Conta do Zé:")
	fmt.Println(contaDoZe.GetTitular())
	fmt.Println("Saldo:", contaDoZe.GetSaldo())
	fmt.Println("")

	fmt.Println("Conta do João:")
	fmt.Println(contaDoJoao.GetTitular())
	fmt.Println("Saldo:", contaDoJoao.GetSaldo())
	fmt.Println("")
}
