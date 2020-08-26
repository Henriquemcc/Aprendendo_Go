package contas

import (
	"banco/clientes"
	"strconv"
)

//ContaPoupanca serve para guardar dados da conta poupanca do usuario
type ContaPoupanca struct {
	titular                              *clientes.Titular
	numeroAgencia, numeroConta, operacao int
	saldo                                float64
}

//Sacar serve para realizar o saque na conta poupanca do usuario.
//Parametro: valorDoSaque: Valor da quantidade de dinheiro que sera sacada.
//Retorno: bool: Valor booleano indicando se o saque ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante o saque.
//Retorno: float64: Novo valor do saldo da conta.
func (c *ContaPoupanca) Sacar(valorDoSaque float64) (bool, string, float64) {
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

//Depositar serve para realizar o deposito na conta poupanca do usuario.
//Parametro: valorDoDeposito: Valor da quantidade de dinheiro que sera depositada.
//Retorno: bool: Valor booleano indicando se o deposito ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante o deposito.
//Retorno: float64: Novo valor do saldo da conta.
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (bool, string, float64) {
	podeDepositar := valorDoDeposito > 0
	var mensagemDeErro string

	if podeDepositar {
		c.saldo += valorDoDeposito
	} else {
		mensagemDeErro = "O valor do deposito não pode ser menor que zero."
	}

	return podeDepositar, mensagemDeErro, c.saldo
}

//Transferir serve para realizar a transferencia entre contas poupancas.
//Parametro: valorDaTransferencia: Valor da quantidade de dinheiro que sera transferida de uma conta para outra.
//Retorno: bool: Valor booleano indicando se o transferencia ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante a transferencia.
//Retorno: float64: Novo valor da saldo da conta de quem transferiu dinheiro.
func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDestino *ContaPoupanca) (bool, string, float64) {
	podeTransferir, _, _ := c.Sacar(valorDaTransferencia)
	var mensagemDeErro string

	if podeTransferir {
		contaDestino.Depositar(valorDaTransferencia)
	} else if valorDaTransferencia < 0 {
		mensagemDeErro = "O valor da transferencia não pode ser menor que zero."
	} else if valorDaTransferencia > c.saldo {
		mensagemDeErro = "O valor da transferencia não pode ser maior que o saldo da conta."
	}

	return podeTransferir, mensagemDeErro, c.saldo
}

//GetTitular serve para obter a instancia da struct titular da conta poupanca.
//Retorno: Titular: Uma instancia da struct titular.
func (c *ContaPoupanca) GetTitular() *clientes.Titular {
	return c.titular
}

//SetTitular serve para alterar a instancia da struct titular da conta poupanca.
//Parametro: titular: uma outra instancia da struct titular.
func (c *ContaPoupanca) SetTitular(titular *clientes.Titular) {
	c.titular = titular
}

//GetNumeroAgencia serve para obter o numero da agencia da conta poupanca.
//Retorno: int: Numero da agencia.
func (c *ContaPoupanca) GetNumeroAgencia() int {
	return c.numeroAgencia
}

//SetNumeroAgencia serve para alterar o valor do numero da agencia da conta poupanca.
//Parametro: numeroAgencia: Novo valor para o numero da agencia da conta.
func (c *ContaPoupanca) SetNumeroAgencia(numeroAgencia int) {
	c.numeroAgencia = numeroAgencia
}

//GetNumeroConta serve para obter o numero da conta poupanca.
//Retorno: int: Numero da conta poupanca.
func (c *ContaPoupanca) GetNumeroConta() int {
	return c.numeroConta
}

//SetNumeroConta serve para alterar o valor do numero da conta poupanca.
//Parametro: numeroConta: Novo valor para o numero da conta poupanca.
func (c *ContaPoupanca) SetNumeroConta(numeroConta int) {
	c.numeroConta = numeroConta
}

//GetSaldo serve para obter o valor do saldo da conta poupanca.
//Retorno: float64: Valor do saldo da conta.
func (c *ContaPoupanca) GetSaldo() float64 {
	return c.saldo
}

//GetOperacao serve para obter o valor da operacao da conta poupanca.
//Retorno: int: Valor da operacao
func (c *ContaPoupanca) GetOperacao() int {
	return c.operacao
}

//SetOperacao serve para alterar o valor da operacao da conta poupanca.
//Parametro: operacao: Novo valor para a operacao da conta poupanca.
func (c *ContaPoupanca) SetOperacao(operacao int) {
	c.operacao = operacao
}

//ToString serve para gerar uma string com os dados de Conta Poupanca.
//Retorno: string: String com os dados de Conta Poupanca.
func (c *ContaPoupanca) ToString() string {
	titular := c.GetTitular()
	mensagem := "Conta Corrente:" + "\n"
	mensagem += titular.ToString() + "\n"
	mensagem += "Número da Agencia: " + strconv.Itoa(c.GetNumeroAgencia()) + "\n"
	mensagem += "Número da Conta: " + strconv.Itoa(c.GetNumeroConta()) + "\n"
	mensagem += "Operação: " + strconv.Itoa(c.GetOperacao()) + "\n"
	mensagem += "Saldo: " + strconv.FormatFloat(c.GetSaldo(), 'G', -1, 64) + "\n"
	return mensagem
}
