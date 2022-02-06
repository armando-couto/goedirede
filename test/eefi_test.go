package test

import (
	"github.com/armando-couto/leitor_edi_rede"
	"testing"
)

func TestEEFI(t *testing.T) {
	linha := "03005022022RedecardExtrato de Movimentacao FinanceiraTESTE                 000634080956998DIARIO         V3.01 - 09/06 - EEFI"
	if retorno := leitor_edi_rede.PopularRegistroTipo030CabecalhoDoArquivo(linha); retorno.TipoDeRegistro > 0 {
		t.Errorf("Erro no teste PopularRegistroTipo030CabecalhoDoArquivo(linha)")
	}

	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo032CabecalhoMatriz(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo034Creditos(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo035AjustesNetEDesagendados(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo036Antecipacoes(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo037TotalizadorDeCreditos(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo038AjustesADebitoViaBanco(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo040Serasa(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo041AVS(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo042SecureCode(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo044DebitosPendentes(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo045DebitosLiquidados(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo049DesagendamentosDeParcelas(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo050TotalizadorMatriz(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo052TrailerDoArquivo(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo054AjustesADebitoViaBancoECommerce(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo055DebitosPendentes(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo056DebitosLiquidadosECommerce(linha)
	//
	//linha = ""
	//leitor_edi_rede.PopularRegistroTipo057DesagendamentosDeParcelasECommerce(linha)

	//got := Add(4, 6)
	//want := 10
}
