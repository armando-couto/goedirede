package goedirede

import (
	"github.com/armando-couto/goutils"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo060CabecalhoDoArquivo ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo060CabecalhoDoArquivo struct {
	TipoDeRegistro            int
	DataDeEmissao             time.Time
	RedeAdquirente            string
	ExtratoEletronicoDeVendas string
	NomeComercial             string
	SequenciaDoMovimento      int
	NumeroPvGrupoOuMatriz     int
	DiarioOuReprocessamento   string
	VersaoDoArquivo           string
}

func PopularRegistroTipo060CabecalhoDoArquivo(linha string) RegistroTipo060CabecalhoDoArquivo {
	value := RegistroTipo060CabecalhoDoArquivo{}
	if len(linha) < 119 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.DataDeEmissao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[3:11])
	value.RedeAdquirente = linha[11:19]
	value.ExtratoEletronicoDeVendas = linha[19:59]
	value.NomeComercial = linha[59:81]
	value.SequenciaDoMovimento = goutils.ConvertStringToInt(linha[81:87])
	value.NumeroPvGrupoOuMatriz = goutils.ConvertStringToInt(linha[87:96])
	value.DiarioOuReprocessamento = linha[96:111]
	value.VersaoDoArquivo = linha[111:119]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo061CabecalhoDoArquivo ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo061CabecalhoDoArquivo struct {
	TipoDeRegistro        int
	NumeroPvMatriz        int
	NomeComercialDaMatriz string
}

func PopularRegistroTipo061CabecalhoDoArquivo(linha string) RegistroTipo061CabecalhoDoArquivo {
	value := RegistroTipo061CabecalhoDoArquivo{}
	if len(linha) < 34 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvMatriz = goutils.ConvertStringToInt(linha[3:12])
	value.NomeComercialDaMatriz = linha[12:34]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo062CabecalhoDoArquivo ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo062CabecalhoDoArquivo struct {
	TipoDeRegistro           int
	NumeroDaOrdemDeCredito   string
	Tipo                     string
	Banco                    int
	Agencia                  int
	ContaCorrente            string
	DataVencimentoDaOC       time.Time
	NumeroEstabelecimento    int
	MicrofilmeDeOrigem       string
	NumeroRv                 int
	DataRv                   time.Time
	Tipo2                    int
	ValorDaVenda             float64
	ValorDoDesconto          float64
	ValorDaGorjeta           float64
	Oc                       float64
	NumeroPv                 int
	NumeroDaParcelaCreditada int
	Bandeira                 string
}

func PopularRegistroTipo062CabecalhoDoArquivo(linha string) RegistroTipo062CabecalhoDoArquivo {
	value := RegistroTipo062CabecalhoDoArquivo{}
	if len(linha) < 162 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDaOrdemDeCredito = linha[3:18]
	value.Tipo = linha[18:19]
	value.Banco = goutils.ConvertStringToInt(linha[19:22])
	value.Agencia = goutils.ConvertStringToInt(linha[22:31])
	value.ContaCorrente = linha[31:42]
	value.DataVencimentoDaOC = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[42:50])
	value.NumeroEstabelecimento = goutils.ConvertStringToInt(linha[50:59])
	value.MicrofilmeDeOrigem = linha[59:72]
	value.NumeroRv = goutils.ConvertStringToInt(linha[72:81])
	value.DataRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[81:89])
	value.Tipo2 = goutils.ConvertStringToInt(linha[89:90])
	value.ValorDaVenda = goutils.ConvertStringToFloatScale2FormatNumber(linha[90:105])
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[105:120])
	value.ValorDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[120:135])
	value.Oc = goutils.ConvertStringToFloatScale2FormatNumber(linha[135:150])
	value.NumeroPv = goutils.ConvertStringToInt(linha[150:159])
	value.NumeroDaParcelaCreditada = goutils.ConvertStringToInt(linha[159:161])
	value.Bandeira = linha[161:162]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo066TotalizadorMatriz ////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo066TotalizadorMatriz struct {
	TipoDeRegistro                 int
	NumeroPvMatriz                 int
	QuantidadeTotalDeResumosMatriz int
	ValorTotalLiquidoASerCreditado float64
}

func PopularRegistroTipo066TotalizadorMatriz(linha string) RegistroTipo066TotalizadorMatriz {
	value := RegistroTipo066TotalizadorMatriz{}
	if len(linha) < 77 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvMatriz = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeTotalDeResumosMatriz = goutils.ConvertStringToInt(linha[12:17])
	value.ValorTotalLiquidoASerCreditado = goutils.ConvertStringToFloatScale2FormatNumber(linha[62:77])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo068TrailerDoArquivo /////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo068TrailerDoArquivo struct {
	TipoDeRegistro                 int
	QuantidadeDeMatrizesNoArquivo  int
	QuantidadeDeRegistrosNoArquivo int
	NumeroPvGrupo                  int
	ValorTotalLiquido              float64
}

func PopularRegistroTipo068TrailerDoArquivo(linha string) RegistroTipo068TrailerDoArquivo {
	value := RegistroTipo068TrailerDoArquivo{}
	if len(linha) < 36 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.QuantidadeDeMatrizesNoArquivo = goutils.ConvertStringToInt(linha[3:7])
	value.QuantidadeDeRegistrosNoArquivo = goutils.ConvertStringToInt(linha[7:12])
	value.NumeroPvGrupo = goutils.ConvertStringToInt(linha[12:21])
	value.ValorTotalLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[21:36])
	return value
}
