package models

import (
	"database/sql"
	"fmt"
	"loja/db"
	"os"
	"strings"
)

//Produto serve para guardar os dados de um produto.
type Produto struct {
	id         int
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

//GetID serve para obter o valor do id do produto.
//Retorno: int: Valor do id do produto.s
func (p *Produto) GetID() int {
	return p.id
}

//SetID serve para alterar  valor do id do produto.
//Parametro: id: Novo valor para o id do produto.
func (p *Produto) SetID(id int) {
	p.id = id
}

//BuscarTodosProdutosDB serve para buscar todos os produtos armazenados no banco de dados.
//Parametro: db: Um ponteiro para a conexao com o banco de dados.
//Retorno: []models.Produto: Lista de produtos armazenados no banco de dados.
func BuscarTodosProdutosDB(db *sql.DB) []Produto {

	//Criando uma slice de produtos
	listaDeProdutos := []Produto{}

	//Obtendo todos os produtos do banco de dados
	todosProdutos, erro := db.Query("select * from produtos")

	//Adicionando produto por produto
	for todosProdutos.Next() {

		//Criando as variaveis que irao receber os valores do banco de dados
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//Lendo os dados do banco de dados
		erro = todosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		//Caso algum erro acontecer o programa ira abortar
		if erro != nil {
			fmt.Println("Erro ao obter os dados do banco de dados:", erro)
			os.Exit(-1)
		}

		//Criando uma instancia de produto
		produto := Produto{}

		//Adicionando os dados lidos do banco de dados a instancia da struct Produto
		produto.SetID(id)
		produto.SetNome(nome)
		produto.SetDescricao(descricao)
		erroBool, mensagemDeErro := produto.SetPreco(preco)
		if !erroBool {
			fmt.Println(mensagemDeErro)
		}
		erroBool, mensagemDeErro = produto.SetQuantidade(quantidade)
		if !erroBool {
			fmt.Println(mensagemDeErro)
		}

		//Adicionando a instancia da struct Produto a slice de produtos
		listaDeProdutos = append(listaDeProdutos, produto)
	}

	defer db.Close()

	return listaDeProdutos
}

//BuscarTodosProdutos serve para conectar com o banco de dados e obter uma lista com todos os produtos.
//Retorno: []Produto: Lista com todos os produtos recuperados do banco de dados.
func BuscarTodosProdutos() []Produto {

	//Conectando com o banco de dados
	db := db.ConectarComBancoDeDados(db.ObterCredenciaisDeAcessoAoBancoDeDados(db.NomeArquivoCredencialBancoDeDados))

	//Buscando todos os produtos
	listaDeProdutos := BuscarTodosProdutosDB(db)

	//Fechando a conexao com o banco de dados
	erro := db.Close()

	//Caso algum erro ocorra ao fechar a conexao, o programa sera abortado
	if erro != nil {
		fmt.Println("Um erro ocorreu ao tentar fechar a conexao com o banco de dados:", erro)
		os.Exit(-1)
	}

	return listaDeProdutos
}
