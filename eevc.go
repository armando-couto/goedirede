package goedirede

import (
	"github.com/armando-couto/goutils"
	"time"
)

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////// RegistroTipo002HeaderDoArquivo //////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo002HeaderDoArquivo struct {
	TipoDeRegistro            int
	DataDeEmissao             time.Time
	RedeAdquirente            string
	ExtratoEletronicoDeVendas string
	NomeComercial             string
	SequenciaDoMovimento      int
	NumeroPvGrupoOuMatriz     int
	TipoDeMovimento           string
	VersaoDoArquivo           string
	Livre                     string
}

func PopularRegistroTipo002HeaderDoArquivo(linha string) RegistroTipo002HeaderDoArquivo {
	value := RegistroTipo002HeaderDoArquivo{}
	if len(linha) < 121 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.DataDeEmissao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[3:11])
	value.RedeAdquirente = linha[11:19]
	value.ExtratoEletronicoDeVendas = linha[19:49]
	value.NomeComercial = linha[49:71]
	value.SequenciaDoMovimento = goutils.ConvertStringToInt(linha[71:77])
	value.NumeroPvGrupoOuMatriz = goutils.ConvertStringToInt(linha[77:86])
	value.TipoDeMovimento = linha[86:101]
	value.VersaoDoArquivo = linha[101:121]
	if len(linha) > 121 {
		value.Livre = linha[121:]
	}
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////// RegistroTipo004HeaderMatriz //////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo004HeaderMatriz struct {
	TipoDeRegistro        int
	NumeroPvMatriz        int
	NomeComercialDaMatriz string
}

func PopularRegistroTipo004HeaderMatriz(linha string) RegistroTipo004HeaderMatriz {
	value := RegistroTipo004HeaderMatriz{}
	if len(linha) < 12 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvMatriz = goutils.ConvertStringToInt(linha[3:12])
	value.NomeComercialDaMatriz = linha[12:]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////// RegistroTipo005Request /////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo005Request struct {
	TipoDeRegistro             int
	NumeroPv                   int
	NumeroRv                   int
	NumeroDoCartao             string
	ValorDaTransacao           float64
	DataDaTransacao            time.Time
	NumeroDeReferencia         int
	NumeroDoProcesso           int
	NumeroDoCvNsu              int
	NumeroDeAutorizacao        string
	CodigoRequest              string
	LimiteDeEnvioDosDocumentos time.Time
	Bandeira                   string
	Livre                      string
}

func PopularRegistroTipo005Request(linha string) RegistroTipo005Request {
	value := RegistroTipo005Request{}
	if len(linha) < 121 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoCartao = linha[21:37]
	value.ValorDaTransacao = goutils.ConvertStringToFloatScale2FormatNumber(linha[37:52])
	value.DataDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[52:60])
	value.NumeroDeReferencia = goutils.ConvertStringToInt(linha[60:75])
	value.NumeroDoProcesso = goutils.ConvertStringToInt(linha[75:90])
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[90:102])
	value.NumeroDeAutorizacao = linha[102:108]
	value.CodigoRequest = linha[108:112]
	value.LimiteDeEnvioDosDocumentos = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[112:120])
	value.Bandeira = linha[120:121]
	if len(linha) > 121 {
		value.Livre = linha[121:]
	}
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////// RegistroTipo033RequestEcommerce ////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo033RequestEcommerce struct {
	TipoDeRegistro      int
	NumeroPv            int
	NumeroRv            int
	NumeroDoCartao      string
	DataDaTransacao     time.Time
	NumeroDoCvNsu       int
	NumeroDeAutorizacao string
	Tid                 string
	NumeroPedido        string
}

func PopularRegistroTipo033RequestEcommerce(linha string) RegistroTipo033RequestEcommerce {
	value := RegistroTipo033RequestEcommerce{}
	if len(linha) < 113 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoCartao = linha[21:37]
	value.DataDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[37:45])
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[45:57])
	value.NumeroDeAutorizacao = linha[57:63]
	value.Tid = linha[63:83]
	value.NumeroPedido = linha[83:113]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////// RegistroTipo006RvRotativo ///////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo006RvRotativo struct {
	Index                      int
	TipoDeRegistro             int
	NumeroPv                   string
	NumeroRv                   int
	NumeroDoBanco              int
	NumeroAgencia              int
	NumeroContaCorrente        int
	DataDoRv                   time.Time
	QuantidadeDeCvNsusAcatados int
	ValorBruto                 float64
	ValorDaGorjeta             float64
	ValorRejeitado             float64
	ValorDesconto              float64
	ValorLiquido               float64
	DataDeCredito              time.Time
	Bandeira                   string
}

func PopularRegistroTipo006RvRotativo(linha string, index int) RegistroTipo006RvRotativo {
	value := RegistroTipo006RvRotativo{}
	if len(linha) < 137 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = linha[3:12]
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoBanco = goutils.ConvertStringToInt(linha[21:24])
	value.NumeroAgencia = goutils.ConvertStringToInt(linha[24:29])
	value.NumeroContaCorrente = goutils.ConvertStringToInt(linha[29:40])
	value.DataDoRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.QuantidadeDeCvNsusAcatados = goutils.ConvertStringToInt(linha[48:53])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[53:68])
	value.ValorDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[68:83])
	value.ValorRejeitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[83:98])
	value.ValorDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[98:113])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[113:128])
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[128:136])
	value.Bandeira = linha[136:137]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////// RegistroTipo008CvNsuRotativo ///////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo008CvNsuRotativo struct {
	Index              int
	TipoDeRegistro     int
	NumeroPv           int
	NumeroRv           int
	DataDoCvNsu        time.Time
	Zeros              int
	ValorDoCvNsu       goutils.KeepZero
	ValorDaGorjeta     goutils.KeepZero
	NumeroDoCartao     string
	StatusDoCvNsu      string
	NumeroDoCvNsu      string
	NumeroDeReferencia string
	ValorDesconto      float64
	NumeroAutorizacao  string
	HoraDaTransacao    time.Time
	NumeroDoBilhete1   string
	NumeroDoBilhete2   string
	NumeroDoBilhete3   string
	NumeroDoBilhete4   string
	TipoDeCaptura      int
	ValorLiquido       goutils.KeepZero
	NumeroDoTerminal   string
	SiglaDoPais        string
	Bandeira           string
}

func PopularRegistroTipo008CvNsuRotativo(linha string, index int) RegistroTipo008CvNsuRotativo {
	value := RegistroTipo008CvNsuRotativo{}
	if len(linha) < 230 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.Zeros = goutils.ConvertStringToInt(linha[29:37])
	value.ValorDoCvNsu = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[37:52]))
	value.ValorDaGorjeta = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[52:67]))
	value.NumeroDoCartao = linha[67:83]
	value.StatusDoCvNsu = linha[83:86]
	value.NumeroDoCvNsu = linha[86:98]
	value.NumeroDeReferencia = linha[98:111]
	value.ValorDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[111:126])
	value.NumeroAutorizacao = linha[126:132]
	value.HoraDaTransacao = goutils.ConvertStringToTimeLayoutHHMMSS(linha[132:138])
	value.NumeroDoBilhete1 = linha[138:154]
	value.NumeroDoBilhete2 = linha[154:170]
	value.NumeroDoBilhete3 = linha[170:186]
	value.NumeroDoBilhete4 = linha[186:202]
	value.TipoDeCaptura = goutils.ConvertStringToInt(linha[202:203])
	value.ValorLiquido = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[203:218]))
	value.NumeroDoTerminal = linha[218:226]
	value.SiglaDoPais = linha[226:229]
	value.Bandeira = linha[229:230]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////// RegistroTipo034CvNsuRotativoEcommerce ///////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo034CvNsuRotativoEcommerce struct {
	Index               int
	TipoDeRegistro      int
	NumeroPv            int
	NumeroRv            int
	DataDoCvNsu         time.Time
	ValorDoCvNsu        float64
	NumeroDoCartao      string
	NumeroDoCvNsu       int
	NumeroDeAutorizacao string
	Tid                 string
	NumeroPedido        string
}

func PopularRegistroTipo034CvNsuRotativoEcommerce(linha string, index int) RegistroTipo034CvNsuRotativoEcommerce {
	value := RegistroTipo034CvNsuRotativoEcommerce{}
	if len(linha) < 98 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.ValorDoCvNsu = goutils.ConvertStringToFloatScale2FormatNumber(linha[29:44])
	value.NumeroDoCartao = linha[44:60]
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[60:72])
	value.NumeroDeAutorizacao = linha[72:78]
	value.Tid = linha[78:98]
	value.NumeroPedido = linha[98:]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////////////// RegistroTipo040CvNsuRecarga ///////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo040CvNsuRecarga struct {
	TipoDeRegistro      int
	NumeroPv            int
	NumeroRv            int
	DataDoCvNsu         time.Time
	NumeroDoCvNsu       int
	ValorDaRecarga      float64
	NumeroDeAutorizacao string
	NumeroDoCelular     string
	Bandeira            string
}

func PopularRegistroTipo040CvNsuRecarga(linha string) RegistroTipo040CvNsuRecarga {
	value := RegistroTipo040CvNsuRecarga{}
	if len(linha) < 78 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[29:41])
	value.ValorDaRecarga = goutils.ConvertStringToFloatScale2FormatNumber(linha[41:56])
	value.NumeroDeAutorizacao = linha[56:62]
	value.NumeroDoCelular = linha[62:77]
	value.Bandeira = linha[77:78]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////// RegistroTipo010RvParceladoSemJuros ///////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo010RvParceladoSemJuros struct {
	Index                   int
	TipoDeRegistro          int
	NumeroPv                int
	NumeroRv                int
	NumeroDoBanco           int
	NumeroDaAgencia         int
	NumeroDaContaCorrente   int
	DataDoRv                time.Time
	QuantidadeDeCvNsus      int
	ValorBruto              float64
	ValorDaGorjeta          float64
	ValorRejeitado          float64
	ValorDoDesconto         float64
	ValorLiquido            float64
	DataDeCreditoDa1Parcela time.Time
	Bandeira                string
}

func PopularRegistroTipo010RvParceladoSemJuros(linha string, index int) RegistroTipo010RvParceladoSemJuros {
	value := RegistroTipo010RvParceladoSemJuros{}
	if len(linha) < 137 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRv = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoBanco = goutils.ConvertStringToInt(linha[21:24])
	value.NumeroDaAgencia = goutils.ConvertStringToInt(linha[24:29])
	value.NumeroDaContaCorrente = goutils.ConvertStringToInt(linha[29:40])
	value.DataDoRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.QuantidadeDeCvNsus = goutils.ConvertStringToInt(linha[48:53])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[53:68])
	value.ValorDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[68:83])
	value.ValorRejeitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[83:98])
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[98:113])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[113:128])
	value.DataDeCreditoDa1Parcela = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[128:136])
	value.Bandeira = linha[136:137]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////// RegistroTipo011AjustesACreditoExtratoEletronicoDeVendas ////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo011AjustesACreditoExtratoEletronicoDeVendas struct {
	Index                   int
	TipoDeRegistro          int
	NumeroDoPvCredito       string
	NumeroDoResumoDoCredito string
	DataDoAjuste            time.Time
	ValorDoAjuste           goutils.KeepZero
	DataDoCredito           time.Time
	ValorDoCredito          goutils.KeepZero
	TipoCredito             string
	NumeroDoBanco           int
	NumeroDaAgencia         int
	NumeroDaContaCorrente   int
	ModeloDoAjuste          int
	DescricaoDoAjuste       string
	Bandeira                string
}

func PopularRegistroTipo011AjustesACreditoExtratoEletronicoDeVendas(linha string) RegistroTipo011AjustesACreditoExtratoEletronicoDeVendas {
	value := RegistroTipo011AjustesACreditoExtratoEletronicoDeVendas{}
	if len(linha) < 119 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPvCredito = linha[3:12]
	value.NumeroDoResumoDoCredito = linha[12:21]
	value.DataDoAjuste = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.ValorDoAjuste = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[29:44]))
	value.DataDoCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[44:52])
	value.ValorDoCredito = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[52:67]))
	value.TipoCredito = linha[67:68]
	value.NumeroDoBanco = goutils.ConvertStringToInt(linha[68:71])
	value.NumeroDaAgencia = goutils.ConvertStringToInt(linha[71:77])
	value.NumeroDaContaCorrente = goutils.ConvertStringToInt(linha[77:88])

	value.DescricaoDoAjuste = linha[90:118]
	value.Bandeira = linha[118:119]

	if len(linha) > 119 { // Esse é o novo campo
		value.ModeloDoAjuste = goutils.ConvertStringToInt(linha[119:123])
	} else { // campo Antigo
		value.ModeloDoAjuste = goutils.ConvertStringToInt(linha[88:90])
	}
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////// RegistroTipo012CvNsuParceladoSemJuros //////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo012CvNsuParceladoSemJuros struct {
	Index                         int
	TipoDeRegistro                int
	NumeroDoPv                    int
	NumeroDoRv                    int
	DataDoCvNsu                   time.Time
	Zeros                         int
	ValorDoCvNsu                  goutils.KeepZero
	ValorDaGorjeta                goutils.KeepZero
	NumeroDoCartao                string
	StatusDoCvNsu                 string
	NumeroDeParcelas              int
	NumeroDoCvNsu                 string
	NumeroDeReferencia            string
	ValorDoDesconto               float64
	NumeroDaAutorizacao           string
	HoraDaTransacao               time.Time
	NumeroDoBilhete1              string
	NumeroDoBilhete2              string
	NumeroDoBilhete3              string
	NumeroDoBilhete4              string
	TipoDeCaptura                 int
	ValorLiquidoDoCvNsu           goutils.KeepZero
	ValorLiquidoDaPrimeiraParcela goutils.KeepZero
	ValorLiquidoDasDemaisParcelas goutils.KeepZero
	NumeroDoTerminal              string
	SiglaDoPais                   string
	Bandeira                      string
}

func PopularRegistroTipo012CvNsuParceladoSemJuros(linha string, index int) RegistroTipo012CvNsuParceladoSemJuros {
	value := RegistroTipo012CvNsuParceladoSemJuros{}
	if len(linha) < 262 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.Zeros = goutils.ConvertStringToInt(linha[29:37])
	value.ValorDoCvNsu = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[37:52]))
	value.ValorDaGorjeta = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[52:67]))
	value.NumeroDoCartao = linha[67:83]
	value.StatusDoCvNsu = linha[83:86]
	value.NumeroDeParcelas = goutils.ConvertStringToInt(linha[86:88])
	value.NumeroDoCvNsu = linha[88:100]
	value.NumeroDeReferencia = linha[100:113]
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[113:128])
	value.NumeroDaAutorizacao = linha[128:134]
	value.HoraDaTransacao = goutils.ConvertStringToTimeLayoutHHMMSS(linha[134:140])
	value.NumeroDoBilhete1 = linha[140:156]
	value.NumeroDoBilhete2 = linha[156:172]
	value.NumeroDoBilhete3 = linha[172:188]
	value.NumeroDoBilhete4 = linha[188:204]
	value.TipoDeCaptura = goutils.ConvertStringToInt(linha[204:205])
	value.ValorLiquidoDoCvNsu = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[205:220]))
	value.ValorLiquidoDaPrimeiraParcela = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[221:235]))
	value.ValorLiquidoDasDemaisParcelas = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[235:250]))
	value.NumeroDoTerminal = linha[250:258]
	value.SiglaDoPais = linha[258:261]
	value.Bandeira = linha[261:262]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////// RegistroTipo035CvNsuParceladoSemJurosEcommerce ///////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo035CvNsuParceladoSemJurosEcommerce struct {
	Index             int
	TipoDeRegistro    int
	NumeroDoPv        int
	NumeroDoRv        int
	DataDoCvNsu       time.Time
	ValorDoCvNsu      float64
	NumeroDoCartao    string
	NumeroDoCvNsu     int
	NumeroAutorizacao string
	Tid               string
	NumeroPedido      string
}

func PopularRegistroTipo035CvNsuParceladoSemJurosEcommerce(linha string, index int) RegistroTipo035CvNsuParceladoSemJurosEcommerce {
	value := RegistroTipo035CvNsuParceladoSemJurosEcommerce{}
	if len(linha) < 98 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.ValorDoCvNsu = goutils.ConvertStringToFloatScale2FormatNumber(linha[29:44])
	value.NumeroDoCartao = linha[44:60]
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[60:72])
	value.NumeroAutorizacao = linha[72:78]
	value.Tid = linha[78:98]
	value.NumeroPedido = linha[98:]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////// RegistroTipo014ParcelasParceladoSemJuros /////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo014ParcelasParceladoSemJuros struct {
	Index                        int
	TipoDeRegistro               int
	NumeroDoPv                   int
	NumeroDoRv                   int
	DataDoCvNsu                  time.Time
	Brancos                      int
	NumeroDaParcela              int
	ValorDaParcelaBruto          float64
	ValorDoDescontoSobreAParcela float64
	ValorDaParcelaLiquida        float64
	DataDeCredito                time.Time
	Livre                        string
}

func PopularRegistroTipo014ParcelasParceladoSemJuros(linha string, index int) RegistroTipo014ParcelasParceladoSemJuros {
	value := RegistroTipo014ParcelasParceladoSemJuros{}
	if len(linha) < 92 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.Brancos = goutils.ConvertStringToInt(linha[29:37])
	value.NumeroDaParcela = goutils.ConvertStringToInt(linha[37:39])
	value.ValorDaParcelaBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[39:54])
	value.ValorDoDescontoSobreAParcela = goutils.ConvertStringToFloatScale2FormatNumber(linha[54:69])
	value.ValorDaParcelaLiquida = goutils.ConvertStringToFloatScale2FormatNumber(linha[69:84])
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[84:92])
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////// RegistroTipo016RvIata ////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo016RvIata struct {
	TipoDeRegistro                       int
	NumeroDoPv                           int
	NumeroDoRv                           int
	NumeroDoBanco                        int
	NumeroDaAgencia                      int
	NumeroDaContaCorrente                int
	DataDoRv                             time.Time
	QuantidadeDeCvNsus                   int
	ValorBruto                           float64
	ValorDaTaxaDeEmbarque                float64
	ValorRejeitado                       float64
	ValorDoDesconto                      float64
	ValorLiquido                         float64
	DataDeCredito                        time.Time
	IdentificaABandeiraDoCartaoUtilizado string
}

func PopularRegistroTipo016RvIata(linha string) RegistroTipo016RvIata {
	value := RegistroTipo016RvIata{}
	if len(linha) < 137 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoBanco = goutils.ConvertStringToInt(linha[21:24])
	value.NumeroDaAgencia = goutils.ConvertStringToInt(linha[24:29])
	value.NumeroDaContaCorrente = goutils.ConvertStringToInt(linha[29:40])
	value.DataDoRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.QuantidadeDeCvNsus = goutils.ConvertStringToInt(linha[48:53])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[53:68])
	value.ValorDaTaxaDeEmbarque = goutils.ConvertStringToFloatScale2FormatNumber(linha[68:83])
	value.ValorRejeitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[83:98])
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[98:113])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[113:128])
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[128:136])
	value.IdentificaABandeiraDoCartaoUtilizado = linha[136:137]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////////// RegistroTipo017Avs ///////////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo017Avs struct {
	TipoDeRegistro                  int
	NumeroDoPv                      int
	QuantidadeDeConsultasRealizadas int
	DataDaConsulta                  time.Time
}

func PopularRegistroTipo017Avs(linha string) RegistroTipo017Avs {
	value := RegistroTipo017Avs{}
	if len(linha) < 25 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeDeConsultasRealizadas = goutils.ConvertStringToInt(linha[12:17])
	value.DataDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[17:25])
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////// RegistroTipo018CvNsuParceladoIata ///////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo018CvNsuParceladoIata struct {
	TipoDeRegistro                int
	NumeroDoPv                    int
	NumeroDoRv                    int
	DataDoCvNsu                   time.Time
	Zeros                         int
	ValorDoCvNsu                  float64
	ValorDaTaxaDeEmbarque         float64
	NumeroDoCartao                string
	StatusDoCvNsu1                string
	NumeroDeParcelas              int
	NumeroDoCvNsu                 int
	NumeroDeReferencia            string
	ValorDoDesconto               float64
	NumeroAutorizacao             string
	HoraDaTransacao               time.Time
	NumeroDoBilhete1              string
	NumeroDoBilhete2              string
	NumeroDoBilhete3              string
	NumeroDoBilhete4              string
	TipoDeCaptura                 int
	ValorLiquido                  float64
	ValorLiquidoDaPrimeiraParcela float64
	ValorLiquidoDasDemaisParcelas float64
	NumeroDoTerminal              string
	SiglaDoPais                   string
	Bandeira                      string
}

func PopularRegistroTipo018CvNsuParceladoIata(linha string) RegistroTipo018CvNsuParceladoIata {
	value := RegistroTipo018CvNsuParceladoIata{}
	if len(linha) < 262 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.Zeros = goutils.ConvertStringToInt(linha[29:37])
	value.ValorDoCvNsu = goutils.ConvertStringToFloatScale2FormatNumber(linha[37:52])
	value.ValorDaTaxaDeEmbarque = goutils.ConvertStringToFloatScale2FormatNumber(linha[52:67])
	value.NumeroDoCartao = linha[68:83]
	value.StatusDoCvNsu1 = linha[83:86]
	value.NumeroDeParcelas = goutils.ConvertStringToInt(linha[86:88])
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[88:100])
	value.NumeroDeReferencia = linha[100:113]
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[113:128])
	value.NumeroAutorizacao = linha[128:134]
	value.HoraDaTransacao = goutils.ConvertStringToTimeLayoutHHMMSS(linha[134:140])
	value.NumeroDoBilhete1 = linha[140:156]
	value.NumeroDoBilhete2 = linha[156:172]
	value.NumeroDoBilhete3 = linha[172:188]
	value.NumeroDoBilhete4 = linha[188:204]
	value.TipoDeCaptura = goutils.ConvertStringToInt(linha[204:205])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[205:220])
	value.ValorLiquidoDaPrimeiraParcela = goutils.ConvertStringToFloatScale2FormatNumber(linha[220:235])
	value.ValorLiquidoDasDemaisParcelas = goutils.ConvertStringToFloatScale2FormatNumber(linha[235:250])
	value.NumeroDoTerminal = linha[250:258]
	value.SiglaDoPais = linha[258:261]
	value.Bandeira = linha[261:262]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////// RegistroTipo036CvNsuParceladoIataEcommerce //////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo036CvNsuParceladoIataEcommerce struct {
	TipoDeRegistro    int
	NumeroDoPv        int
	NumeroDoRv        int
	DataDoCvNsu       time.Time
	ValorDoCvNsu      float64
	NumeroDoCartao    string
	NumeroDoCvNsu     int
	NumeroAutorizacao string
	Tid               string
	NumeroDoPedido    string
}

func PopularRegistroTipo036CvNsuParceladoIataEcommerce(linha string) RegistroTipo036CvNsuParceladoIataEcommerce {
	value := RegistroTipo036CvNsuParceladoIataEcommerce{}
	if len(linha) < 128 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[22:29])
	value.ValorDoCvNsu = goutils.ConvertStringToFloatScale2FormatNumber(linha[29:44])
	value.NumeroDoCartao = linha[44:60]
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[60:72])
	value.NumeroAutorizacao = linha[72:78]
	value.Tid = linha[78:98]
	value.NumeroDoPedido = linha[98:128]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ////////////////////////////////////// RegistroTipo019Serasa ///////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo019Serasa struct {
	TipoDeRegistro                  int
	NumeroDoPv                      int
	QuantidadeDeConsultasRealizadas int
	DataDaConsulta                  time.Time
}

func PopularRegistroTipo019Serasa(linha string) RegistroTipo019Serasa {
	value := RegistroTipo019Serasa{}
	if len(linha) < 25 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeDeConsultasRealizadas = goutils.ConvertStringToInt(linha[12:17])
	value.DataDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[17:25])
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// //////////////////////////// RegistroTipo020ParcelasParcelandoIata /////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo020ParcelasParcelandoIata struct {
	TipoDeRegistro               int
	NumeroDoPv                   int
	NumeroDoRv                   int
	DataDoRv                     time.Time
	Brancos                      int
	NumeroDaPArcelaBruta         int
	ValorDaParcelaBruto          float64
	ValorDoDescontoSobreAParcela float64
	ValorDaParcelaLiquida        float64
	DataDoCredito                time.Time
	Livre                        string
}

func PopularRegistroTipo020ParcelasParcelandoIata(linha string) RegistroTipo020ParcelasParcelandoIata {
	value := RegistroTipo020ParcelasParcelandoIata{}
	if len(linha) < 1024 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.Brancos = goutils.ConvertStringToInt(linha[29:37])
	value.NumeroDaPArcelaBruta = goutils.ConvertStringToInt(linha[37:39])
	value.ValorDaParcelaBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[39:54])
	value.ValorDoDescontoSobreAParcela = goutils.ConvertStringToFloatScale2FormatNumber(linha[54:69])
	value.ValorDaParcelaLiquida = goutils.ConvertStringToFloatScale2FormatNumber(linha[69:84])
	value.DataDoCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[84:92])
	value.Livre = linha[92:1024]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////// RegistroTipo021SecureCode //////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo021SecureCode struct {
	TipoDeRegistro                  int
	NumeroDoPv                      int
	QuantidadeDeConsultasRealizadas int
	DataDaConsulta                  time.Time
	Bandeira                        string
}

func PopularRegistroTipo021SecureCode(linha string) RegistroTipo021SecureCode {
	value := RegistroTipo021SecureCode{}
	if len(linha) < 26 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeDeConsultasRealizadas = goutils.ConvertStringToInt(linha[12:17])
	value.DataDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[17:25])
	value.Bandeira = linha[25:26]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////// RegistroTipo022RvDolar /////////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo022RvDolar struct {
	TipoDeRegistro        int
	NumeroDoPv            int
	NumeroDoRv            int
	NumeroDoBanco         int
	NumeroDoAgencia       int
	NumeroDaContaCorrente int
	DataDoRv              time.Time
	QuantidadeDeCvNsus    int
	ValorBruto            float64
	ValorDaGorjeta        float64
	ValorRejeitado        float64
	ValorDoDesconto       float64
	ValorLiquido          float64
	DataDeCredito         time.Time
	Bandeira              string
}

func PopularRegistroTipo022RvDolar(linha string) RegistroTipo022RvDolar {
	value := RegistroTipo022RvDolar{}
	if len(linha) < 137 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoBanco = goutils.ConvertStringToInt(linha[21:24])
	value.NumeroDoAgencia = goutils.ConvertStringToInt(linha[24:29])
	value.NumeroDaContaCorrente = goutils.ConvertStringToInt(linha[29:40])
	value.DataDoRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.QuantidadeDeCvNsus = goutils.ConvertStringToInt(linha[48:53])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[53:68])
	value.ValorDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[68:83])
	value.ValorRejeitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[83:98])
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[98:113])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[113:128])
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[128:136])
	value.Bandeira = linha[136:137]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// /////////////////////////////////// RegistroTipo024DvNsuDolar //////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo024DvNsuDolar struct {
	TipoDeRegistro     int
	NumeroDoPv         int
	NumeroDoRv         int
	DataDoCvNsu        time.Time
	Zeros              int
	ValorDoCvNsu       float64
	ValorDaGorjeta     float64
	NumeroDoCartao     string
	StatusDoCvNsu      string
	CotacaoDoDolar     float64
	DataDaCotacao      time.Time
	NumeroDoCvNsu      int
	NumeroDeReferencia string
	ValorDoDesconto    float64
	NumeroAutorizaao   string
	HoraDaTransacao    time.Time
	NumeroDoTerminal   string
	TipoDeCaptura      int
	SiglaDoPais        string
	Bandeira           string
}

func PopularRegistroTipo024DvNsuDolar(linha string) RegistroTipo024DvNsuDolar {
	value := RegistroTipo024DvNsuDolar{}
	if len(linha) < 169 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoCvNsu = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.Zeros = goutils.ConvertStringToInt(linha[29:37])
	value.ValorDoCvNsu = goutils.ConvertStringToFloatScale2FormatNumber(linha[37:52])
	value.ValorDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[52:67])
	value.NumeroDoCartao = linha[67:83]
	value.StatusDoCvNsu = linha[83:86]
	value.CotacaoDoDolar = goutils.ConvertStringToFloatScale2FormatNumber(linha[86:95])
	value.DataDaCotacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[95:103])
	value.NumeroDoCvNsu = goutils.ConvertStringToInt(linha[103:115])
	value.NumeroDeReferencia = linha[115:128]
	value.ValorDoDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[128:143])
	value.NumeroAutorizaao = linha[143:149]
	value.HoraDaTransacao = goutils.ConvertStringToTimeLayoutHHMMSS(linha[149:155])
	value.NumeroDoTerminal = linha[155:163]
	value.TipoDeCaptura = goutils.ConvertStringToInt(linha[163:165])
	value.SiglaDoPais = linha[165:168]
	value.Bandeira = linha[168:169]
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////// RegistroTipo026TotalizadorMatriz /////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo026TotalizadorMatriz struct {
	TipoDeRegistro               int
	NumeroPvMatriz               int
	ValoBruto                    float64
	QuantidadeDeCvNsusRejeitados int
	ValorTotalRejeitado          float64
	ValorTotalRotativo           float64
	ValorTotalParceladoSemJuros  float64
	ValorTotalParceladoIata      float64
	ValorTotalDolar              float64
	ValorTotalDesconto           float64
	ValorTotalLiquido            float64
	ValorTotalDaGorjeta          float64
	ValorTotalDaTaxaDeEmbarque   float64
	QuantidadeCvNsuAcatados      int
}

func PopularRegistroTipo026TotalizadorMatriz(linha string) RegistroTipo026TotalizadorMatriz {
	value := RegistroTipo026TotalizadorMatriz{}
	if len(linha) < 174 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvMatriz = goutils.ConvertStringToInt(linha[3:12])
	value.ValoBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[12:27])
	value.QuantidadeDeCvNsusRejeitados = goutils.ConvertStringToInt(linha[27:33])
	value.ValorTotalRejeitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[33:48])
	value.ValorTotalRotativo = goutils.ConvertStringToFloatScale2FormatNumber(linha[48:63])
	value.ValorTotalParceladoSemJuros = goutils.ConvertStringToFloatScale2FormatNumber(linha[63:78])
	value.ValorTotalParceladoIata = goutils.ConvertStringToFloatScale2FormatNumber(linha[78:93])
	value.ValorTotalDolar = goutils.ConvertStringToFloatScale2FormatNumber(linha[93:108])
	value.ValorTotalDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[108:123])
	value.ValorTotalLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[123:138])
	value.ValorTotalDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[138:153])
	value.ValorTotalDaTaxaDeEmbarque = goutils.ConvertStringToFloatScale2FormatNumber(linha[153:168])
	value.QuantidadeCvNsuAcatados = goutils.ConvertStringToInt(linha[168:174])
	return value
}

// ////////////////////////////////////////////////////////////////////////////////////////////////////
// ///////////////////////////// RegistroTipo028TrailerDeArquivo //////////////////////////////////////
// ////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo028TrailerDeArquivo struct {
	TipoDeRegistro               int
	QuantidadeDeMatrizes         int
	QuantidadeDeRegistros        int
	NumeroPvGrupo                int
	ValoTotalBruto               float64
	QuantidadeDeCvNsusRejeitados int
	ValorTotalRejeitado          float64
	ValorTotalRotativo           float64
	ValorTotalParceladoSemJuros  float64
	ValorTotalParceladoIata      float64
	ValorTotalDolar              float64
	ValorTotalDesconto           float64
	ValorTotalLiquido            float64
	ValorTotalDaGorjeta          float64
	ValorTotalDaTaxaDeEmbarque   float64
	QuantidadeCvNsuAcatados      int
}

func PopularRegistroTipo028TrailerDeArquivo(linha string) RegistroTipo028TrailerDeArquivo {
	value := RegistroTipo028TrailerDeArquivo{}
	if len(linha) < 184 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.QuantidadeDeMatrizes = goutils.ConvertStringToInt(linha[3:7])
	value.QuantidadeDeRegistros = goutils.ConvertStringToInt(linha[7:13])
	value.NumeroPvGrupo = goutils.ConvertStringToInt(linha[13:22])
	value.ValoTotalBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[22:37])
	value.QuantidadeDeCvNsusRejeitados = goutils.ConvertStringToInt(linha[37:43])
	value.ValorTotalRejeitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[43:58])
	value.ValorTotalRotativo = goutils.ConvertStringToFloatScale2FormatNumber(linha[58:73])
	value.ValorTotalParceladoSemJuros = goutils.ConvertStringToFloatScale2FormatNumber(linha[73:88])
	value.ValorTotalParceladoIata = goutils.ConvertStringToFloatScale2FormatNumber(linha[88:103])
	value.ValorTotalDolar = goutils.ConvertStringToFloatScale2FormatNumber(linha[103:118])
	value.ValorTotalDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[118:133])
	value.ValorTotalLiquido = goutils.ConvertStringToFloatScale2FormatNumber(linha[133:148])
	value.ValorTotalDaGorjeta = goutils.ConvertStringToFloatScale2FormatNumber(linha[148:163])
	value.ValorTotalDaTaxaDeEmbarque = goutils.ConvertStringToFloatScale2FormatNumber(linha[163:178])
	value.QuantidadeCvNsuAcatados = goutils.ConvertStringToInt(linha[178:184])
	return value
}
