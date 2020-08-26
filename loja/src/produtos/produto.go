package produtos

import "strings"

//Produto serve para guardar os dados de um produto.
type Produto struct {
	nome       string
	descricao  string
	preco      float64
	quantidade int
}

//GetNome serve para obter o valor do nome do produto.
//Retorno: string: Valor do nome do produto.
func (p *Produto) GetNome() string {
	return p.nome
}

//SetNome serve para alterar o valor do nome do produto.
//Parametro: nome: Novo valor para o nome do produto.
func (p *Produto) SetNome(nome string) {
	p.nome = strings.Title(nome)
}

//GetDescricao serve para obter o valor da descricao do produto.
//Retorno: string: Valor da descricao do produto.
func (p *Produto) GetDescricao() string {
	return p.descricao
}

//SetDescricao serve para alterar o valor da descricao do produto.
//Parametro: descricao: Novo valor para a descricao do produto.
func (p *Produto) SetDescricao(descricao string) {
	p.descricao = descricao
}

//GetPreco serve para obter o valor do preco do produto.
//Retorno: float64: Valor do preco do produto.
func (p *Produto) GetPreco() float64 {
	return p.preco
}

//SetPreco serve para alterar o valor do preco do produto.
//Parametro: preco: Novo valor para o preco do produto.
//Retorno: bool: Valor booleano indicando se foi possivel alterar o preco do produto.
//Retorno: string: Mensagem de erro caso nao seja possivel alterar o preco do produto.
func (p *Produto) SetPreco(preco float64) (bool, string) {
	var mensagemDeErro string

	//Verificando se o preco eh um valor valido
	precoEhValido := preco >= 0

	//Adicionando o novo valor ao preco caso ele seja valido
	if precoEhValido {
		p.preco = preco

		//Gerando mesagem de erro caso o preco seja invalido
	} else {
		mensagemDeErro = "O preço do produto não pode ser menor do que zero."
	}

	return precoEhValido, mensagemDeErro
}

//GetQuantidade serve para obter o valor da quantidade do produto que esta disponivel.
//Retorno: int: Valor da quantidade do produto que esta disponivel.
func (p *Produto) GetQuantidade() int {
	return p.quantidade
}

//SetQuantidade serve para alterar o valor da quantidade do produto que esta disponivel.
//Parametro: quantidade: Novo valor para a quantidade do produto.
//Retorno: bool: Valor booleano indicando se foi possivel alterar a quantidade do produto.
//Retorno: string: Mensagem de erro caso nao seja possivel alterar a quantidade do produto.
func (p *Produto) SetQuantidade(quantidade int) (bool, string) {
	var mensagemDeErro string

	//Verificando se a quantidade eh valida
	quantidadeEhValida := quantidade >= 0

	//Adicionando o novo valor caso ele seja valido
	if quantidadeEhValida {
		p.quantidade = quantidade

		//Gerando mensagem de erro caso a quantidade seja invalido
	} else {
		mensagemDeErro = "A quantidade não pode ser menor que zero."
	}

	return quantidadeEhValida, mensagemDeErro
}
