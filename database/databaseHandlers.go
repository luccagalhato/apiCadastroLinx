package database

import "apiCadastro/server"

// Handlers ...
func (c *ConectionSqlConfig) Handlers() map[string]server.Handler {
	return map[string]server.Handler{
		// "/cadastro_page": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.CadastroPage,
		// },
		// "/cadastro_produto": {
		// 	Method: http.MethodPost,
		// 	Fn:     c.CadastroProduto,
		// },
		// "/consulta_linx/{id}": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaLinx,
		// },
		// "/consulta_produtos": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutos,
		// },
		// "/photos/{file}": {
		// 	Method: http.MethodGet,
		// 	Fn:     http.StripPrefix("/photos/", http.FileServer(http.Dir("photos/"))).ServeHTTP,
		// },
		// "/pedidos_page": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.PedidosPage,
		// },
		// "/cadastro_pedido": {
		// 	Method: http.MethodPost,
		// 	Fn:     c.CadastroPedido,
		// },
		// "/consulta_pedidos": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaPedidos,
		// },
		// "/consulta_produtos_ecom": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutosEcom,
		// },
		// "/consulta_produtos_ecom_colecao": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutosEcomColecao,
		// },
		// "/consulta_item_pedido": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaItemPedido,
		// },
		// "/consulta_produtos_look": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ConsultaProdutosEcom,
		// },
		// "/adicionar_item_pedido": {
		// 	Method: http.MethodPost,
		// 	Fn:     c.AdicionarItemPedido,
		// },

		// "/resumo_pedido": {
		// 	Method: http.MethodGet,
		// 	Fn:     c.ResumoPedido,
		// },
		// "/itemsCadastrado": {
		// 	http.MethodGet,
		// 	apihandler.Getall,
		// },
		// "/grade": {
		// 	http.MethodGet,
		// 	apihandler.GetAllGrade,
		// },
		// "/marca": {
		// 	http.MethodGet,
		// 	apihandler.GetMarca,
		// },
		// "/classificacao": {
		// 	http.MethodGet,
		// 	apihandler.GetAllClassificacao,
		// },
		// "/subclassificacao": {
		// 	http.MethodGet,
		// 	apihandler.GetSubClassificacao,
		// },
		// "/colecao": {
		// 	http.MethodGet,
		// 	apihandler.GetAllColecao,
		// },
		// "/subcolecao": {
		// 	http.MethodGet,
		// 	apihandler.GetSubColecao,
		// },
		// "/novoItem": {
		// 	http.MethodPost,
		// 	apihandler.NewItem,
		// },
		// "/item/{id}": {
		// 	http.MethodPost,
		// 	apihandler.GetItem,
		// },
	}
}
