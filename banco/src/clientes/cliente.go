package clientes

import (
	"strconv"
	"strings"
	"time"
)

//Titular serve para guardar os dados pessoais do cliente
type Titular struct {
	nome             string
	sobrenome        string
	cpf              string
	profissao        string
	dataDeNascimento *time.Time
	telefone         string
	email            string
}

//GetNome serve para obter o valor do nome do titular.
//Retorno: string: Nome do titular.
func (t *Titular) GetNome() string {
	return t.nome
}

//SetNome serve para alterar o valor do nome do titular.
//Parametro: nome: Novo valor para o nome do titular.
func (t *Titular) SetNome(nome string) {
	t.nome = strings.ToUpper(nome)
}

//GetSobrenome serve para obter o valor do sobrenome do titular.
//Retorno: string: Sobrenome do titular.
func (t *Titular) GetSobrenome() string {
	return t.sobrenome
}

//SetSobrenome serve para alterar o valor do sobrenome do titular.
//Parametro: sobrenome: Novo valor para o sobrenome do titular.
func (t *Titular) SetSobrenome(sobrenome string) {
	t.sobrenome = strings.ToUpper(sobrenome)
}

//GetCpf serve para obter o valor do cpf do titular.
//Retorno: string: cpf do titular.
func (t *Titular) GetCpf() string {
	return t.cpf[0:3] + "." + t.cpf[3:6] + "." + t.cpf[6:9] + "-" + t.cpf[9:11]
}

//SetCpf serve para alterar o valor do cpf do titular.
//Parametro: cpf: Novo valor para o cpf do titular.
//Retorno: bool: Valor booleano indicando se a alteracao do valor do cpf foi bem sucedida.
//Retorno: string: Mensagem de erro caso a alteracao do valor do cpf seja invalida.
func (t *Titular) SetCpf(cpf string) (bool, string) {

	//Removendo todos os . e - do cpf
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	ehValido, mensagemDeErro := cpfEhValido(cpf)

	if ehValido {
		t.cpf = cpf
	}

	return ehValido, mensagemDeErro
}

//GetProfissao serve para obter o valor de profissao do titular.
//Retorno: string: Profissao do titular.
func (t *Titular) GetProfissao() string {
	return t.profissao
}

//SetProfissao serve para alterar o valor de profissao do titular.
//Parametro: profissao: Novo valor para a profissao do titular.
func (t *Titular) SetProfissao(profissao string) {
	t.profissao = strings.ToUpper(profissao)
}

//GetDataDeNascimento serve para obter o valor de data de nascimento do titular.
//Retorno: Time: Data de nascimento.
func (t *Titular) GetDataDeNascimento() *time.Time {
	return t.dataDeNascimento
}

//SetDataDeNascimento serve para alterar o valor da data de nascimento do titular.
//Parametro: dataDeNascimento: Novo valor para a data de nascimento.
func (t *Titular) SetDataDeNascimento(dataDeNascimento *time.Time) {
	t.dataDeNascimento = dataDeNascimento
}

//GetTelefone serve para obter o valor do telefone do titular
//Retorno: string: Numero de telefone.
func (t *Titular) GetTelefone() string {
	var numeroTelefoneExibido string
	if len(t.telefone) == 12 {
		numeroTelefoneExibido = "+" + t.telefone[0:2] + " (" + t.telefone[2:4] + ") " + t.telefone[4:8] + "-" + t.telefone[8:12]
	} else if len(t.telefone) == 13 {
		numeroTelefoneExibido = "+" + t.telefone[0:2] + " (" + t.telefone[2:4] + ") " + t.telefone[4:9] + "-" + t.telefone[9:13]
	}
	return numeroTelefoneExibido
}

//SetTelefone serve para alterar o valor do telefone do titular
//Parametro: telefone: Novo valor para o numero de telefone.
//Retorno: bool: Valor booleano indicando se a alteracao do telefone foi bem sucedida.
//Retorno: string: Mensagem de erro caso a alteracao do telefone nao seja bem sucedida.
func (t *Titular) SetTelefone(telefone string) (bool, string) {
	telefone = strings.ReplaceAll(telefone, "(", "")
	telefone = strings.ReplaceAll(telefone, ")", "")
	telefone = strings.ReplaceAll(telefone, " ", "")
	telefone = strings.ReplaceAll(telefone, "+", "")
	telefone = strings.ReplaceAll(telefone, "-", "")

	//Verificando se o telefone eh valido
	ehValido, mensagemDeInvalido := telefoneEhValido(telefone)

	//Alterando o numero de telefone caso ele seja valido
	if ehValido {
		t.telefone = telefone
	}

	return ehValido, mensagemDeInvalido

}

//GetEmail serve para obter o valor de email do titular.
//Retorno: string: Endereco de email do titular.
func (t *Titular) GetEmail() string {
	return t.email
}

//SetEmail serve para alterar o valor do email do titular.
//Parametro: email: Novo valor para o endereco de email.
//Retorno: bool: Valor booleano indicando se a alteracao do endereco de email foi bem sucedida.
//Retorno: string: Mensagem de erro, caso endereco de email seja invalido.
func (t *Titular) SetEmail(email string) (bool, string) {

	//Convertendo para minusculo
	email = strings.ToLower(email)

	//Verificando se o email eh valido
	ehValido, mensagemDeInvalido := emailEhValido(email)

	//Alterando o endereco de email caso ele seja valido
	if ehValido {
		t.email = email
	}

	return ehValido, mensagemDeInvalido

}

//ToString serve para gerar uma string com os dados do Titular.
//Retorno: string: String com os dados do titular.
func (t *Titular) ToString() string {
	mensagem := "Titular:" + "\n"
	mensagem += "Nome: " + t.GetNome() + "\n"
	mensagem += "Sobrenome: " + t.GetSobrenome() + "\n"
	mensagem += "CPF: " + t.GetCpf() + "\n"
	mensagem += "Profissão: " + t.GetProfissao() + "\n"
	dataDeNascimento := t.GetDataDeNascimento()
	mensagem += "Data de nascimento: " + dataDeNascimento.Format("02/01/2006") + "\n"
	mensagem += "Telefone: " + t.GetTelefone() + "\n"
	mensagem += "Email: " + t.GetEmail() + "\n"

	return mensagem
}

//cpfEhValido serve para verificar se um cpf eh valido.
//Parametro: cpf: cpf (somente numero) que sera verificado se eh valido.
//Retorno: bool: Valor booleano indicando se o cpf eh valido.
//Retorno: string: Mensagem de erro explicando o porque do cpf ser invalido.
func cpfEhValido(cpf string) (bool, string) {
	ehValido := false
	var mensagemDeInvalido string

	if len(cpf) == 11 {

		//Calculando o primeiro digito verificador
		resto := stringToInt(string(cpf[0])) * 10
		resto += stringToInt(string(cpf[1])) * 9
		resto += stringToInt(string(cpf[2])) * 8
		resto += stringToInt(string(cpf[3])) * 7
		resto += stringToInt(string(cpf[4])) * 6
		resto += stringToInt(string(cpf[5])) * 5
		resto += stringToInt(string(cpf[6])) * 4
		resto += stringToInt(string(cpf[7])) * 3
		resto += stringToInt(string(cpf[8])) * 2
		resto %= 11

		primeiroDigitoVerificador := 0
		if resto != 0 && resto != 1 {
			primeiroDigitoVerificador = 11 - resto
		}

		//Verificando se o primeiro digito verificador esta correto
		if primeiroDigitoVerificador == stringToInt(string(cpf[9])) {

			//Calculando o segundo digito verificador
			resto = stringToInt(string(cpf[0])) * 11
			resto += stringToInt(string(cpf[1])) * 10
			resto += stringToInt(string(cpf[2])) * 9
			resto += stringToInt(string(cpf[3])) * 8
			resto += stringToInt(string(cpf[4])) * 7
			resto += stringToInt(string(cpf[5])) * 6
			resto += stringToInt(string(cpf[6])) * 5
			resto += stringToInt(string(cpf[7])) * 4
			resto += stringToInt(string(cpf[8])) * 3
			resto += stringToInt(string(cpf[9])) * 2
			resto %= 11

			segundoDigitoVerificador := 0
			if resto != 0 && resto != 1 {
				segundoDigitoVerificador = 11 - resto
			}

			//Verificando se o segundo digito verificador esta correto
			if segundoDigitoVerificador == stringToInt(string(cpf[10])) {
				ehValido = true

				//Gerando mensagem de erro caso o segundo digito verificador seja invalido
			} else {
				mensagemDeInvalido = "O segundo digito verificador é inválido."
			}

			//Gerando mensagem de erro caso o primeiro digito verificador seja invalido
		} else {
			mensagemDeInvalido = "O primeiro digito verificador é inválido."
		}

		//Gerando mensagem de erro caso o cpf contenha mais de 11 caracteres
	} else if len(cpf) > 11 {
		mensagemDeInvalido = "O CPF contém mais de 11 caracteres numéricos."

		//Gerando mensagem de erro caso o cpf contenha menos de 11 caracteres
	} else {
		mensagemDeInvalido = "O CPF contém menos de 11 caracteres numéricos."
	}

	return ehValido, mensagemDeInvalido
}

//stringToInt serve para converter uma string contendo um numero inteiro em um numero inteiro
//Retorno: int: Numero inteiro resultado da conversao.
func stringToInt(str string) int {
	numeroInteiro, _ := strconv.Atoi(str)
	return numeroInteiro
}

//telefoneEhValido serve para verificar se um numero de telefone eh valido.
//Parametro: telefone: Numero de telefone que sera verificado.
//Retorno: bool: Valor booleano indicando se o numero de telefone eh valido.
//Retorno: string: Mensagem de erro caso o numero de telefone seja invalido.
func telefoneEhValido(telefone string) (bool, string) {
	ehValido := false
	var mensagemDeInvalido string

	//Verificando o tamanho do numero de telefone
	if len(telefone) == 12 || len(telefone) == 13 {
		ehValido = true

		//Gerando mensagem de erro caso o tamanho do numero de telefone seja inferior a 12 caracteres.
	} else if len(telefone) < 12 {
		mensagemDeInvalido = "O numero de telefone tem menos de 12 caracteres"

		//Gerando mensagem de erro caso o tamanho do numero de telefone seja superior a 13 caracteres.
	} else {
		mensagemDeInvalido = "O numero de telefone tem mais de 13 caracteres."
	}

	return ehValido, mensagemDeInvalido
}

//emailEhValido serve para verificar se um endereco de email eh valido.
//Parametro: email: Endereco de email que sera verificado.
//Retorno: bool: Valor booleano indicando se o endereco de email eh valido.
//Retorno: string: Mensagem de erro caso o endereco de email seja invalido.
func emailEhValido(email string) (bool, string) {

	ehValido := false
	var mensagemDeInvalido string

	//Verificando se o endereco de email contem o @
	if strings.Contains(email, "@") {
		ehValido = true

		//Gerando mensagem de erro caso o endereco de email nao contenha o @.
	} else {
		mensagemDeInvalido = "O endereço de email não contem o @."
	}

	return ehValido, mensagemDeInvalido
}
