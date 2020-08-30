package models

import (
	"fmt"
	"loja/db"
	"strconv"
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
//Retorno: bool: Valor booleano indicando se foi possivel alterar o id.
//Retorno: string: Mensagem de erro caso nao seja possivel alterar o id do produto.
func (p *Produto) SetID(id int) (bool, string) {
	var mensagemDeErro string

	//Verificando se o id eh valido
	idEhValido := id >= 0

	//Alterando o novo valor para o id
	if idEhValido {
		p.id = id

		//Gerando mensagem de erro
	} else {
		mensagemDeErro = "O id não pode ser menor que zero."
	}

	return idEhValido, mensagemDeErro
}

//ToString serve para gerar uma string de uma instancia da struct produto.
//Retorno: string: String contendo todos os dados da instancia da struct produto.
func (p *Produto) ToString() string {

	mensagem := "Produto:\n"
	mensagem += "ID: " + strconv.Itoa(p.GetID()) + "\n"
	mensagem += "Nome: " + p.GetNome() + "\n"
	mensagem += "Descrição: " + p.GetDescricao() + "\n"
	mensagem += "Preço: " + strconv.FormatFloat(p.GetPreco(), 'G', -1, 64) + "\n"
	mensagem += "Quantidade: " + strconv.Itoa(p.GetQuantidade()) + "\n"

	return mensagem
}

//BuscarTodosProdutos serve para buscar todos os produtos armazenados no banco de dados.
//Retorno: []models.Produto: Lista de produtos armazenados no banco de dados.
func BuscarTodosProdutos() []Produto {

	//Conectando com o banco de dados
	db := db.ConectarComBancoDeDados(db.ObterCredenciaisDeAcessoAoBancoDeDados(db.NomeArquivoCredencialBancoDeDados))

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

		//Caso algum erro acontecer o erro sera exibido
		if erro != nil {
			panic(erro.Error())
		}

		//Criando uma instancia de produto
		produto := Produto{}

		//Adicionando os dados lidos do banco de dados a instancia da struct Produto
		sucesso, mensagemDeErro := produto.SetID(id)
		//Caso algum erro ocorra ao adicionar o id do produto, sera impresso uma mensagem de erro
		if !sucesso {
			fmt.Println(mensagemDeErro)
		}

		produto.SetNome(nome)
		produto.SetDescricao(descricao)

		sucesso, mensagemDeErro = produto.SetPreco(preco)
		//Caso algum erro ocorra ao adicionar o preco do produto, sera impresso a mensagem de erro
		if !sucesso {
			fmt.Println(mensagemDeErro)
		}

		sucesso, mensagemDeErro = produto.SetQuantidade(quantidade)
		//Caso algum erro ocorra ao adicionar a quantidade do produto, sera impresso a mensagem de erro
		if !sucesso {
			fmt.Println(mensagemDeErro)
		}

		//Adicionando a instancia da struct Produto a slice de produtos
		listaDeProdutos = append(listaDeProdutos, produto)
	}

	//Fechando a conexao com o banco de dados
	erro = db.Close()

	//Caso algum erro ocorra ao fechar a conexao, o erro sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	return listaDeProdutos
}

//CriarNovoProduto serve para conectar ao banco de dados e adicionar um produto no banco de dados.
//Parametro: produto: Instancia da struct Produto para ser adicionada no banco de dados.
func CriarNovoProduto(produto Produto) {

	//Conectndo com o banco de dados
	db := db.ConectarComBancoDeDados(db.ObterCredenciaisDeAcessoAoBancoDeDados(db.NomeArquivoCredencialBancoDeDados))

	//Preparando a query sql
	inserirDadosNoBancoDeDados, erro := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	//Inserindo no banco de dados os valores do novo produto
	_, erro = inserirDadosNoBancoDeDados.Exec(produto.GetNome(), produto.GetDescricao(), produto.GetPreco(), produto.GetQuantidade())

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	//Fechando a conexao com o banco de dados
	erro = db.Close()

	//Caso algum erro ocorra ao fechar a conexao, o erro sera exibido
	if erro != nil {
		panic(erro.Error())
	}
}

//DeletarProduto serve para conectar ao banco de dados e deletar um produto.
//Parametro: id: Id do produto a ser removido.
func DeletarProduto(id string) {

	//Conectando com o banco de dados
	db := db.ConectarComBancoDeDados(db.ObterCredenciaisDeAcessoAoBancoDeDados(db.NomeArquivoCredencialBancoDeDados))

	//Preparando a query sql
	deletarProduto, erro := db.Prepare("delete from produtos where id=$1")

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	//Deletando produto do banco de dados
	_, erro = deletarProduto.Exec(id)

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	//Fechando a conexao com o banco de dados
	erro = db.Close()

	//Caso algum erro ocorra ao fechar a conexao, o erro sera exibido
	if erro != nil {
		panic(erro.Error())
	}

}

//EditarProduto serve para conectar ao banco de dados e obter o produto para ser editado.
//Parametro: id: Id do produto que sera obtido para ser atualizado.
//Retorno: Produto: Instancia de produto que foi retornado do banco de dados.
func EditarProduto(id string) Produto {

	//Conectando com o banco de dados
	db := db.ConectarComBancoDeDados(db.ObterCredenciaisDeAcessoAoBancoDeDados(db.NomeArquivoCredencialBancoDeDados))

	//Obtendo os dados do produto do banco de dados
	produtoDoBancoDeDados, erro := db.Query("select * from produtos where id=$1", id)

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	//Criando uma nova instancia de produto para armazenar os dados obtidos do banco de dados
	produtoParaAtualizar := Produto{}

	//Lendo os dados do banco de dados
	for produtoDoBancoDeDados.Next() {

		//Declarando as variaveis que receberao os dados do banco de dados
		var id, quantidade int
		var nome, descricao string
		var preco float64

		//Obtendo os dados do produto do banco de dados
		erro = produtoDoBancoDeDados.Scan(&id, &nome, &descricao, &preco, &quantidade)

		//Caso algum erro ocorra, ele sera exibido
		if erro != nil {
			panic(erro.Error())
		}

		//Adicionando dados obtidos do banco de dados na instancia de produto
		produtoParaAtualizar.SetID(id)
		produtoParaAtualizar.SetNome(nome)
		produtoParaAtualizar.SetDescricao(descricao)

		sucesso, mensagemDeErro := produtoParaAtualizar.SetPreco(preco)
		//Caso algum erro ocorra ao adicionar o preco do produto, sera impresso a mensagem de erro
		if !sucesso {
			fmt.Println(mensagemDeErro)
		}

		sucesso, mensagemDeErro = produtoParaAtualizar.SetQuantidade(quantidade)
		//Caso algum erro ocorra ao adicionar a quantidade do produto, sera impresso a mensagem de erro
		if !sucesso {
			fmt.Println(mensagemDeErro)
		}
	}

	//Fechando a conexao com o banco de dados
	erro = db.Close()

	//Caso algum erro ocorra ao fechar a conexao, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	return produtoParaAtualizar
}

//AtualizarProduto serve para conectar ao banco de dados e atualizar um produto
//Parametro: produto: Instancia da struct Produto que sera atualizada no banco de dados.
func AtualizarProduto(produto Produto) {

	//Conectndo com o banco de dados
	db := db.ConectarComBancoDeDados(db.ObterCredenciaisDeAcessoAoBancoDeDados(db.NomeArquivoCredencialBancoDeDados))

	//Preparando query sql
	AtualizarProduto, erro := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}

	_, erro = AtualizarProduto.Exec(produto.GetNome(), produto.GetDescricao(), produto.GetPreco(), produto.GetQuantidade(), produto.GetID())

	//Caso algum erro ocorra, ele sera exibido
	if erro != nil {
		panic(erro.Error())
	}
}
