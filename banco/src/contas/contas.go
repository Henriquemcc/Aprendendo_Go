package contas

//verificarConta serve para que tanto as contas corrente e poupanca possam realizar o pagamento de boletos e a transferencia entre contas.
type verificarConta interface {
	Sacar(valorDoSaque float64) (bool, string, float64)
	Depositar(valorDoDeposito float64) (bool, string, float64)
	GetSaldo() float64
}

//PagarBoleto serve para realizar o pagamento de um boleto por meio do saldo de uma conta.
//Parametro: conta: Conta que sera usada para pagar o boleto.
//Parametro: valorDoBoleto: Valor do boleto que sera pago.
//Retorno: bool: Valor booleano indicando se foi possivel pagar o boleto
//Retorno: string: Mensagem de erro caso nao seja possivel pagar o boleto.
//Retorno: float64: Novo valor do saldo da conta que foi usada para pagar o boleto.
func PagarBoleto(conta verificarConta, valorDoBoleto float64) (bool, string, float64) {
	podePagarBoleto, _, _ := conta.Sacar(valorDoBoleto)

	var mensagemDeErro string

	//Gerando mensagem de erro caso nao seja possivel pagar o boleto
	if !podePagarBoleto {

		//Gerando mensagem de erro caso o valor do boleto seja menor que zero
		if valorDoBoleto < 0 {
			mensagemDeErro = "O valor do boleto não pode ser menor que zero."

			//Gerando mensagem de erro caso o valor do boleto seja maior que o saldo.
		} else if valorDoBoleto > conta.GetSaldo() {
			mensagemDeErro = "O valor do boleto não pode ser maior que o saldo da conta."
		}
	}

	return podePagarBoleto, mensagemDeErro, conta.GetSaldo()
}

//Transferir serve para realizar a transferencia entre duas contas.
//Parametro: valorDaTransferencia: Valor de dinheiro que sera transferido entre as contas.
//Parametro: contaOrigem: Conta que enviara o dinheiro a conta destino.
//Parametro: contaDestino: Conta que recebera o dinhero da conta origem.
//Retorno: bool: Valor booleano indicando se foi possivel transferir dinheiro entre as contas.
//Retorno: string: Mensagem de erro caso nao seja possivel transferir dinheiro entre as contas.
//Retorno: float64: Novo valor do saldo da conta de origem.
func Transferir(valorDaTransferencia float64, contaOrigem verificarConta, contaDestino verificarConta) (bool, string, float64) {
	podeTransferir, _, _ := contaOrigem.Sacar(valorDaTransferencia)

	var mensagemDeErro string

	//Verificando se eh possivel transferir
	if podeTransferir {
		contaDestino.Depositar(valorDaTransferencia)

		//Gerando mensagem de erro caso o valor da transferencia seja menor que zero
	} else if valorDaTransferencia < 0 {
		mensagemDeErro = "O valor da transferencia não pode ser menor que zero."

		//Gerando mensagem de erro caso o valor de transferencia seja maior que o saldo da conta de origem
	} else if valorDaTransferencia > contaOrigem.GetSaldo() {
		mensagemDeErro = "O valor da transferencia não pode ser maior que o saldo da conta de origem"
	}

	return podeTransferir, mensagemDeErro, contaOrigem.GetSaldo()
}
