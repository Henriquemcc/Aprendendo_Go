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
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante o saque.
//Retorno: float64: Novo valor do saldo da conta.
func (c *contaCorrente) Sacar(valorDoSaque float64) (bool, string, float64) {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	var mensagemDeErro string

	if podeSacar {
		c.saldo -= valorDoSaque
	} else if valorDoSaque <= 0 {
		mensagemDeErro = "O valor do saque não pode ser menor ou igual a zero."
	} else if valorDoSaque > c.saldo {
		mensagemDeErro = "O valor do saque não pode ser maior que o saldo da conta."
	}

	return podeSacar, mensagemDeErro, c.saldo
}

//Este metodo serve para realizar o deposito na conta corrente do usuario.
//Parametro: valorDoDeposito: Valor da quantidade de dinheiro que sera depositada.
//Retorno: bool: Valor booleano indicando se o deposito ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante o deposito.
//Retorno: float64: Novo valor do saldo da conta.
func (c *contaCorrente) Depositar(valorDoDeposito float64) (bool, string, float64) {
	podeDepositar := valorDoDeposito > 0
	var mensagemDeErro string

	if podeDepositar {
		c.saldo += valorDoDeposito
	} else {
		mensagemDeErro = "O valor do deposito não pode ser menor que zero."
	}

	return podeDepositar, mensagemDeErro, c.saldo
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
	sucesso, mensagemDeErro, novoSaldo := contaDoZe.Sacar(1000000)
	fmt.Println("Saque da conta do", contaDoZe.titular, "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.titular, ":", contaDoZe.saldo)

	//Depositando dinheiro na conta do ze
	sucesso, mensagemDeErro, novoSaldo = contaDoZe.Depositar(500)
	fmt.Println("Deposito na conta do", contaDoZe.titular, "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do ze
	fmt.Println("Saldo da conta do", contaDoZe.titular, ":", contaDoZe.saldo)

	//Criando uma conta corrente para o personagem ficticio chamado Joao da Silva
	contaDoJoao := contaCorrente{}
	contaDoJoao.titular = "Joao da Silva"
	contaDoJoao.saldo = 1024

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.titular, ":", contaDoJoao.saldo)

	//Sacando dinheiro da conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Sacar(100)
	fmt.Println("Saque da conta do", contaDoJoao.titular, "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.titular, ":", contaDoJoao.saldo)

	//Depositando dinheiro na conta do Joao
	sucesso, mensagemDeErro, novoSaldo = contaDoJoao.Depositar(-13)
	fmt.Println("Deposito na conta do", contaDoJoao.titular, "realizado com sucesso:", sucesso, mensagemDeErro, ". Novo valor do saldo:", novoSaldo)

	//Imprimindo o saldo da conta do Joao
	fmt.Println("Saldo da conta do", contaDoJoao.titular, ":", contaDoJoao.saldo)

	//Imprimindo as contas do Ze e do Joao
	fmt.Println(contaDoZe)
	fmt.Println(contaDoJoao)
}
