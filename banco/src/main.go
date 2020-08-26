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
	nascimentoZe := time.Date(1900, time.December, 18, 0, 0, 0, 0, time.UTC)
	ze.SetDataDeNascimento(&nascimentoZe)

	//Criando um titular chamado joao
	joao := clientes.Titular{}
	joao.SetNome("Joao")
	joao.SetSobrenome("da Silva")
	joao.SetTelefone("+55 (11) 12555-2345")
	joao.SetEmail("joao@gmail.com")
	joao.SetProfissao("Investidor")
	joao.SetCpf("625.733.383-04")
	nascimentoJoao := time.Date(1900, time.June, 5, 0, 0, 0, 0, time.UTC)
	joao.SetDataDeNascimento(&nascimentoJoao)

	//Criando um titular chamado fulano
	fulano := clientes.Titular{}
	fulano.SetNome("Fulano")
	fulano.SetSobrenome("da Silva")
	fulano.SetTelefone("+55 (61) 98765-4321")
	fulano.SetEmail("fulano@gmail.com")
	fulano.SetProfissao("Desenvolvedor")
	fulano.SetCpf("210.376.590-74")
	nascimentoFulano := time.Date(1930, time.January, 8, 0, 0, 0, 0, time.UTC)
	fulano.SetDataDeNascimento(&nascimentoFulano)

	//Criando um titular chamado ciclano
	ciclano := clientes.Titular{}
	ciclano.SetNome("Ciclano")
	ciclano.SetSobrenome("da Silva")
	ciclano.SetTelefone("+55 (43) 99999-9999")
	ciclano.SetEmail("ciclano@gmail.com")
	ciclano.SetProfissao("Técnico de TI")
	ciclano.SetCpf("000.451.460-20")
	nascimentoCiclano := time.Date(1990, time.October, 13, 0, 0, 0, 0, time.UTC)
	ciclano.SetDataDeNascimento(&nascimentoCiclano)

	//Criando uma conta corrente para o Ze
	contaDoZe := contas.ContaCorrente{}
	contaDoZe.SetTitular(&ze)
	contaDoZe.SetNumeroAgencia(0)
	contaDoZe.SetNumeroConta(0)
	contaDoZe.Depositar(134217)

	//Imprimindo os dados da conta do Ze
	fmt.Println(contaDoZe.ToString())

	//Sacando dinheiro da conta do ze
	sucesso, mensagemDeErro, novoSaldo := contaDoZe.Sacar(1000000)
	fmt.Println("Saque da conta do", ze.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo os dados da conta do Ze
	fmt.Println(contaDoZe.ToString())

	//Depositando dinheiro na conta do ze
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Depositar(500)
	fmt.Println("Deposito na conta do", ze.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo os dados da conta do Ze
	fmt.Println(contaDoZe.ToString())

	//Criando uma conta corrente para o Joao
	contaDoJoao := contas.ContaCorrente{}
	contaDoJoao.SetTitular(&joao)
	contaDoJoao.SetNumeroAgencia(1)
	contaDoJoao.SetNumeroConta(0)
	contaDoJoao.Depositar(1024)

	//Imprimindo os dados da conta do Joao
	fmt.Println(contaDoJoao.ToString())

	//Sacando dinheiro da conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Sacar(100)
	fmt.Println("Saque da conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo os dados da conta do Joao
	fmt.Println(contaDoJoao.ToString())

	//Depositando dinheiro na conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Depositar(-13)
	fmt.Println("Deposito na conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo os dados da conta do Joao
	fmt.Println(contaDoJoao.ToString())

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(500, &contaDoJoao)
	fmt.Println("Transferencia da conta do", ze.GetNome(), "para a conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", ze.GetNome(), ":", novoSaldo)

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(9999999999, &contaDoJoao)
	fmt.Println("Transferencia da conta do", ze.GetNome(), "para a conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", ze.GetNome(), ":", novoSaldo)

	//Realizando transferencia entre a conta do Ze e a do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Transferir(-20, &contaDoJoao)
	fmt.Println("Transferencia da conta do", ze.GetNome(), "para a conta do", joao.GetNome(), "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo da conta do", ze.GetNome(), ":", novoSaldo)

	//Imprimindo os dados da conta do Ze e doJoao
	fmt.Println(contaDoZe.ToString())
	fmt.Println(contaDoJoao.ToString())

	//Criando uma conta para o fulano
	contaDoFulano := contas.ContaCorrente{}
	contaDoFulano.SetTitular(&fulano)
	contaDoFulano.SetNumeroAgencia(2)
	contaDoFulano.SetNumeroConta(0)
	contaDoFulano.Depositar(99999999)

	//Criando uma conta para o ciclano
	contaDoCiclano := contas.ContaPoupanca{}
	contaDoCiclano.SetTitular(&ciclano)
	contaDoCiclano.SetNumeroAgencia(2)
	contaDoCiclano.SetNumeroConta(0)
	contaDoCiclano.Depositar(2000)

	//Imprimindo dados da contas do Fulano e Ciclano
	fmt.Println(contaDoFulano.ToString())
	fmt.Println(contaDoCiclano.ToString())

	//Pagando boleto com a conta do Fulano
	sucesso, mensagemDeErro, novoSaldo = contas.PagarBoleto(&contaDoFulano, 500)
	fmt.Println("Pagamento de boleto na conta do", fulano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo é:", novoSaldo)

	//Imprimindo dados da conta do Fulano
	fmt.Println(contaDoFulano.ToString())

	//Pagando boleto com a conta do Ciclano
	sucesso, mensagemDeErro, novoSaldo = contas.PagarBoleto(&contaDoCiclano, 500)
	fmt.Println("Pagamento de boleto na conta do", ciclano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo é:", novoSaldo)

	//Imprimindo dados da conta do Ciclano
	fmt.Println(contaDoCiclano.ToString())

	//Pagando boleto com a conta do Ciclano
	sucesso, mensagemDeErro, novoSaldo = contas.PagarBoleto(&contaDoCiclano, -5)
	fmt.Println("Pagamento de boleto na conta do", ciclano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo é:", novoSaldo)

	//Imprimindo dados da conta do Ciclano
	fmt.Println(contaDoCiclano.ToString())

	//Pagando boleto com a conta do Ciclano
	sucesso, mensagemDeErro, novoSaldo = contas.PagarBoleto(&contaDoCiclano, 999999999999999999999999999999)
	fmt.Println("Pagamento de boleto na conta do", ciclano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo é:", novoSaldo)

	//Imprimindo dados da conta do Ciclano
	fmt.Println(contaDoCiclano.ToString())

	//Transferindo dinheiro da conta do Fulano para a do Ciclano
	sucesso, mensagemDeErro, novoSaldo = contas.Transferir(5000, &contaDoFulano, &contaDoCiclano)
	fmt.Println("Transferencia de dinheiro da conta do", fulano.GetNome(), "para a do", ciclano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo de", fulano.GetNome(), "é:", novoSaldo)

	//Imprimindo dados da contas do Fulano e Ciclano
	fmt.Println(contaDoFulano.ToString())
	fmt.Println(contaDoCiclano.ToString())

	//Transferindo dinheiro da conta do Fulano para a do Ciclano
	sucesso, mensagemDeErro, novoSaldo = contas.Transferir(-10, &contaDoFulano, &contaDoCiclano)
	fmt.Println("Transferencia de dinheiro da conta do", fulano.GetNome(), "para a do", ciclano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo de", fulano.GetNome(), "é:", novoSaldo)

	//Imprimindo dados da contas do Fulano e Ciclano
	fmt.Println(contaDoFulano.ToString())
	fmt.Println(contaDoCiclano.ToString())

	//Transferindo dinheiro da conta do Fulano para a do Ciclano
	sucesso, mensagemDeErro, novoSaldo = contas.Transferir(99999999999999999999999999999999999999999999999999999999999, &contaDoFulano, &contaDoCiclano)
	fmt.Println("Transferencia de dinheiro da conta do", fulano.GetNome(), "para a do", ciclano.GetNome(), ":", sucesso, mensagemDeErro, ".O novo saldo de", fulano.GetNome(), "é:", novoSaldo)

	//Imprimindo dados da contas do Fulano e Ciclano
	fmt.Println(contaDoFulano.ToString())
	fmt.Println(contaDoCiclano.ToString())

}
