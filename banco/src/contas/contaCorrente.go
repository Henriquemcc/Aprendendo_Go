package contas

import (
	"banco/clientes"
	"strconv"
)

//ContaCorrente serve para guardar os dados da conta corrente do usuario.
type ContaCorrente struct {
	titular       *clientes.Titular
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

//Sacar serve para realizar o saque na conta corrente do usuario.
//Parametro: valorDoSaque: Valor da quantidade de dinheiro que sera sacada.
//Retorno: bool: Valor booleano indicando se o saque ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante o saque.
//Retorno: float64: Novo valor do saldo da conta.
func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, string, float64) {
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

//Depositar serve para realizar o deposito na conta corrente do usuario.
//Parametro: valorDoDeposito: Valor da quantidade de dinheiro que sera depositada.
//Retorno: bool: Valor booleano indicando se o deposito ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante o deposito.
//Retorno: float64: Novo valor do saldo da conta.
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, string, float64) {
	podeDepositar := valorDoDeposito > 0
	var mensagemDeErro string

	if podeDepositar {
		c.saldo += valorDoDeposito
	} else {
		mensagemDeErro = "O valor do deposito não pode ser menor que zero."
	}

	return podeDepositar, mensagemDeErro, c.saldo
}

//Transferir serve para realizar a transferencia entre contas correntes.
//Parametro: valorDaTransferencia: Valor da quantidade de dinheiro que sera transferida de uma conta para outra.
//Retorno: bool: Valor booleano indicando se o transferencia ocorreu com sucesso.
//Retorno: string: Mensagem de erro, caso tenha ocorrido algum erro durante a transferencia.
//Retorno: float64: Novo valor da saldo da conta de quem transferiu dinheiro.
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) (bool, string, float64) {
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

//GetTitular serve para obter a instancia da struct titular da conta corrente.
//Retorno: Titular: Uma instancia da struct titular.
func (c *ContaCorrente) GetTitular() *clientes.Titular {
	return c.titular
}

//SetTitular serve para alterar a instancia da struct titular da conta corrente.
//Parametro: titular: uma outra instancia da struct titular.
func (c *ContaCorrente) SetTitular(titular *clientes.Titular) {
	c.titular = titular
}

//GetNumeroAgencia serve para obter o numero da agencia da conta corrente.
//Retorno: int: Numero da agencia.
func (c *ContaCorrente) GetNumeroAgencia() int {
	return c.numeroAgencia
}

//SetNumeroAgencia serve para alterar o valor do numero da agencia da conta corrente.
//Parametro: numeroAgencia: Novo valor para o numero da agencia da conta.
func (c *ContaCorrente) SetNumeroAgencia(numeroAgencia int) {
	c.numeroAgencia = numeroAgencia
}

//GetNumeroConta serve para obter o numero da conta corrente
//Retorno: int: Numero da conta corrente
func (c *ContaCorrente) GetNumeroConta() int {
	return c.numeroConta
}

//SetNumeroConta serve para alterar o valor do numero da conta corrente.
//Parametro: numeroConta: Novo valor para o numero da conta corrente.
func (c *ContaCorrente) SetNumeroConta(numeroConta int) {
	c.numeroConta = numeroConta
}

//GetSaldo serve para obter o valor do saldo da conta corrente.
//Retorno: float64: Valor do saldo da conta.
func (c *ContaCorrente) GetSaldo() float64 {
	return c.saldo
}

//ToString serve para gerar uma string com os dados de Conta Corrente.
//Retorno: string: String com os dados de Conta Corrente.
func (c *ContaCorrente) ToString() string {
	titular := c.GetTitular()
	mensagem := "Conta Corrente:" + "\n"
	mensagem += titular.ToString() + "\n"
	mensagem += "Número da Agencia: " + strconv.Itoa(c.GetNumeroAgencia()) + "\n"
	mensagem += "Número da Conta: " + strconv.Itoa(c.GetNumeroConta()) + "\n"
	mensagem += "Saldo: " + strconv.FormatFloat(c.GetSaldo(), 'G', -1, 64) + "\n"
	return mensagem
}
