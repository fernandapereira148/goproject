package models

import "br.com.industrial/loja/db"

type Produto struct {
	Id              int64
	Nome, Descricao string
	Preco           float64
	Quantidade      int64
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	selectProduto, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}
	defer selectProduto.Close()
	p := Produto{}
	produtos := []Produto{}
	for selectProduto.Next() {
		var id, quantidade int64
		var nome, descricao string
		var preco float64

		err = selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	return produtos
}

func CriarNovoProduto(nome string, descricao string, preco float64, quantidade int64) {

	db := db.ConectaComBancoDeDados()
	defer db.Close()

	insereDadosBanco, err := db.Prepare("inser into produtos(nome,descricao,preco,quantidade) values(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	defer insereDadosBanco.Close()

	insereDadosBanco.Exec(nome, descricao, preco, quantidade)
}

func DeletaProduto(id int64) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	deletaProduto, err := db.Prepare("deleta from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer deletaProduto.Close()

	deletaProduto.Exec(id)
}

func EditaProduto(id int64) Produto {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	produtoDoBanco, err := db.Query("select * from produtos where id =?")
	if err != nil {
		panic(err.Error())
	}
	defer produtoDoBanco.Close()

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int64
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	return produtoParaAtualizar
}

func AtualizaProduto(id int64, nome string, descricao string, preco float64, quantidade int64) {
	db := db.ConectaComBancoDeDados()
	defer db.Close()

	atualizaProduto, err := db.Prepare("uodate produtos set nome=?, descricao=?, preco =? quantidade=? where id=?")
	if err != nil {
		panic(err.Error())
	}
	defer atualizaProduto.Close()
	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
}
