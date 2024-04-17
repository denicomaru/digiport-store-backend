package model

type Carrinho struct {
	UserID       string
	valorTotal   float64
	ID           string
	Infosproduto []InfosProduto
}

type InfosProduto struct {
	produtoId           string
	QuantidadeDeProduto int
}
