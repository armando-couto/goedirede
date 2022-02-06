package test

import (
	"github.com/armando-couto/goleitoredirede"
	"testing"
)

func TestEEFI(t *testing.T) {
	linha := "03005022022RedecardExtrato de Movimentacao FinanceiraTESTE                 000634080956998DIARIO         V3.01 - 09/06 - EEFI"
	if retorno := goleitoredirede.PopularRegistroTipo030CabecalhoDoArquivo(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo030CabecalhoDoArquivo(linha)")
	}

	linha = "032080956998TESTE                 "
	if retorno := goleitoredirede.PopularRegistroTipo032CabecalhoMatriz(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo032CabecalhoMatriz(linha)")
	}

	linha = "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo034Creditos(linha, 0); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo034Creditos(linha)")
	}

	linha = "03508095699805114922028012022000000000011158D18CANCEL.DE VENDAS            341445XXXXXX8018190120220511492203220281626473162801202200000008095699819012022000000000068000D19042022000000002408818000000002419976000000000000000000000000068000000257409999023166T0201976766700000000000000000000000000000011"
	if retorno := goleitoredirede.PopularRegistroTipo035AjustesNetEDesagendados(linha, 0); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo035AjustesNetEDesagendados(linha)")
	}

	linha = "0360809569980203405284204022022000000003922450C34100164600000508679031649374030220220000000039618110703202201/060000000129158200000000002001980809569981"
	if retorno := goleitoredirede.PopularRegistroTipo036Antecipacoes(linha, 0); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo036Antecipacoes(linha)")
	}

	linha = "037081438230       00000000000000000000000 341001646000005086790502202204022022000000000004971"
	if retorno := goleitoredirede.PopularRegistroTipo037TotalizadorDeCreditos(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo037TotalizadorDeCreditos(linha)")
	}

	linha = ""
	if retorno := goleitoredirede.PopularRegistroTipo038AjustesADebitoViaBanco(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo038AjustesADebitoViaBanco(linha)")
	}

	linha = "040081438362000000000000000000000101202231012022000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo040Serasa(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo040Serasa(linha)")
	}

	linha = "041081438362000000000000000000000101202231012022000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo041AVS(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo041AVS(linha)")
	}

	linha = "0420814383620000000000000000000001012022310120220000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo042SecureCode(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo042SecureCode(linha)")
	}

	linha = "042081438362000000000000000000000101202231012022000000000000000000000000000000000000000000000000000000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro(linha, 0); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro(linha)")
	}

	linha = "044080956998220350841400402202200000000002731418CANCEL.DE VENDAS            544731XXXXXX6409000017639197190120220532700000000005600000511492181901202208095699874361424       0402202200000007436142402202200000000000000000000000000000000027314               00                            1"
	if retorno := goleitoredirede.PopularRegistroTipo044DebitosPendentes(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo044DebitosPendentes(linha)")
	}

	linha = ""
	if retorno := goleitoredirede.PopularRegistroTipo045DebitosLiquidados(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo045DebitosLiquidados(linha)")
	}

	linha = "0490809569980511492180000000743614241905202200000000043272400000000048735200000000005462804022022000000000000000000000000560000544731XXXXXX6409190120220000176391971041"
	if retorno := goleitoredirede.PopularRegistroTipo049DesagendamentosDeParcelas(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo049DesagendamentosDeParcelas(linha)")
	}

	linha = "0500809569980000010000000000000000003810000001708672580000000000000000000000000000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo050TotalizadorMatriz(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo050TotalizadorMatriz(linha)")
	}

	linha = "0520001000497080956998000100000000000000000038100000017086725800000000000000000000000000000000000000"
	if retorno := goleitoredirede.PopularRegistroTipo052TrailerDoArquivo(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo052TrailerDoArquivo(linha)")
	}

	linha = "053606282XXXXXX11870202202209384944908463364600000000001459000037691617108503110032202020751132215QUTQ162W36RM                  "
	if retorno := goleitoredirede.PopularRegistroTipo053AjustesNetEDesagendamentosECommerce(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste RegistroTipo053AjustesNetEDesagendamentosECommerce(linha)")
	}

	linha = ""
	if retorno := goleitoredirede.PopularRegistroTipo054AjustesADebitoViaBancoECommerce(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo054AjustesADebitoViaBancoECommerce(linha)")
	}

	linha = ""
	if retorno := goleitoredirede.PopularRegistroTipo055DebitosPendentes(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo055DebitosPendentes(linha)")
	}

	linha = "053550209XXXXXX05882901202201044935508463364600000000005143000063563731722744111052201281922536255A5C5IV3V15KQ                  "
	if retorno := goleitoredirede.PopularRegistroTipo056DebitosLiquidadosECommerce(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo056DebitosLiquidadosECommerce(linha)")
	}

	linha = ""
	if retorno := goleitoredirede.PopularRegistroTipo057DesagendamentosDeParcelasECommerce(linha); retorno.TipoDeRegistro == 0 {
		t.Errorf("Erro no teste PopularRegistroTipo057DesagendamentosDeParcelasECommerce(linha)")
	}
}
