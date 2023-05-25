package goedirede

import (
	"github.com/armando-couto/goutils"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo030CabecalhoDoArquivo ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo030CabecalhoDoArquivo struct {
	TipoDeRegistro                int
	DataDeEmissao                 time.Time
	RedeAdquirente                string
	ExtratoMovimentacaoFinanceira string
	NomeComercial                 string
	SequenciaMovimento            int
	NumeroPvGrupoOuMatriz         int
	TipoProcessamento             string
	VersaoDoArquivo               string
}

func PopularRegistroTipo030CabecalhoDoArquivo(linha string) RegistroTipo030CabecalhoDoArquivo {
	value := RegistroTipo030CabecalhoDoArquivo{}
	if len(linha) < 125 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.DataDeEmissao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[3:11])
	value.RedeAdquirente = linha[11:19]
	value.ExtratoMovimentacaoFinanceira = linha[19:53]
	value.NomeComercial = linha[53:75]
	value.SequenciaMovimento = goutils.ConvertStringToInt(linha[75:81])
	value.NumeroPvGrupoOuMatriz = goutils.ConvertStringToInt(linha[81:90])
	value.TipoProcessamento = linha[90:105]
	value.VersaoDoArquivo = linha[105:125]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////// RegistroTipo032CabecalhoMatriz ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo032CabecalhoMatriz struct {
	TipoDeRegistro        int
	NumeroPvMatriz        string
	NomeComercialDaMatriz string
}

func PopularRegistroTipo032CabecalhoMatriz(linha string) RegistroTipo032CabecalhoMatriz {
	value := RegistroTipo032CabecalhoMatriz{}
	if len(linha) < 34 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvMatriz = linha[3:12]
	value.NomeComercialDaMatriz = linha[12:34]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////////// RegistroTipo034Creditos ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo034Creditos struct {
	Index                 int
	TipoDeRegistro        int
	NumeroPvCentralizador int
	NumeroDoDocumento     int
	DataDoLancamento      time.Time
	ValorDoLancamento     goutils.KeepZero
	Credito               string
	Banco                 int
	Agencia               int
	ContaCorrente         int
	DataDoMovimento       time.Time
	NumeroDoRv            int
	DataDoRv              time.Time
	Bandeira              string
	TipoDeTransacao       int
	ValorBrutoDoRv        goutils.KeepZero
	ValorDaTaxaDeDesconto goutils.KeepZero
	NumeroDaParcelaTotal  string
	StatusDeCredito       string
	NumeroPvOriginal      int
}

func PopularRegistroTipo034Creditos(linha string, index int) RegistroTipo034Creditos {
	value := RegistroTipo034Creditos{}
	if len(linha) < 140 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvCentralizador = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoDocumento = goutils.ConvertStringToInt(linha[12:23])
	value.DataDoLancamento = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[23:31])
	value.ValorDoLancamento = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[31:46]))
	value.Credito = linha[46:47]
	value.Banco = goutils.ConvertStringToInt(linha[47:50])
	value.Agencia = goutils.ConvertStringToInt(linha[50:56])
	value.ContaCorrente = goutils.ConvertStringToInt(linha[56:67])
	value.DataDoMovimento = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[67:75])
	value.NumeroDoRv = goutils.ConvertStringToInt(linha[75:84])
	value.DataDoRv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[84:92])
	value.Bandeira = linha[92:93]
	value.TipoDeTransacao = goutils.ConvertStringToInt(linha[93:94])
	value.ValorBrutoDoRv = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[94:109]))
	value.ValorDaTaxaDeDesconto = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[109:124]))
	value.NumeroDaParcelaTotal = linha[124:129]
	value.StatusDeCredito = linha[129:131]
	value.NumeroPvOriginal = goutils.ConvertStringToInt(linha[131:140])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////// RegistroTipo035AjustesNetEDesagendados ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo035AjustesNetEDesagendados struct {
	Index                              int
	TipoDeRegistro                     int
	NumeroPvAjustado                   int
	NumeroRvAjustado                   int
	DataDoAjuste                       time.Time
	ValorDoAjuste                      goutils.KeepZero
	Debito                             string
	MotivoDoAjusteCodigo               int
	MotivoDoAjuste                     string
	NumeroDoCartao                     string
	DataDaTransacaoCv                  time.Time
	NumeroRvOriginal                   int
	NumeroDeReferenciaDaCartaFax       string
	DataDaCarta                        time.Time
	MesRefenrecia                      int
	NumeroPvOriginal                   int
	DataRvOriginal                     time.Time
	ValorDaTransacao                   goutils.KeepZero
	DesagendamentoOuNet                string
	DataDeCredito                      time.Time
	NovoValorDaParcela                 goutils.KeepZero
	ValorOriginalDaParcela             goutils.KeepZero
	ValorBrutoDoResumoDeVendasOriginal goutils.KeepZero
	ValorDoCancelamentoSolicitado      goutils.KeepZero
	NumeroDoNsu                        string
	NumeroDaAutorizacao                string
	TipoDeDebito                       string
	NumeroDaOrdemdeDebito              int
	ValorDoDebitoTotal                 float64
	ValorPendente                      float64
	BandeiraDoRvDeOrigem               string
	BandeiraDoRvAjustado               string
}

func PopularRegistroTipo035AjustesNetEDesagendados(linha string, index int) RegistroTipo035AjustesNetEDesagendados {
	value := RegistroTipo035AjustesNetEDesagendados{}
	if len(linha) < 300 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvAjustado = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroRvAjustado = goutils.ConvertStringToInt(linha[12:21])
	value.DataDoAjuste = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[21:29])
	value.ValorDoAjuste = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[29:44]))
	value.Debito = linha[44:45]

	value.MotivoDoAjuste = linha[47:75]
	value.NumeroDoCartao = linha[75:91]
	value.DataDaTransacaoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[91:99])
	value.NumeroRvOriginal = goutils.ConvertStringToInt(linha[99:108])
	value.NumeroDeReferenciaDaCartaFax = linha[108:123]
	value.DataDaCarta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[123:131])
	value.MesRefenrecia = goutils.ConvertStringToInt(linha[131:137])
	value.NumeroPvOriginal = goutils.ConvertStringToInt(linha[137:146])
	value.DataRvOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[146:154])
	value.ValorDaTransacao = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[154:169]))
	value.DesagendamentoOuNet = linha[169:170]
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[170:178])
	value.NovoValorDaParcela = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[178:193]))
	value.ValorOriginalDaParcela = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[193:208]))
	value.ValorBrutoDoResumoDeVendasOriginal = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[208:223]))
	value.ValorDoCancelamentoSolicitado = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[223:238]))
	value.NumeroDoNsu = linha[238:250]
	value.NumeroDaAutorizacao = linha[250:256]
	value.TipoDeDebito = linha[256:257]
	value.NumeroDaOrdemdeDebito = goutils.ConvertStringToInt(linha[257:268])
	value.ValorDoDebitoTotal = goutils.ConvertStringToFloatScale2FormatNumber(linha[268:283])
	value.ValorPendente = goutils.ConvertStringToFloatScale2FormatNumber(linha[283:298])
	value.BandeiraDoRvDeOrigem = linha[298:299]
	value.BandeiraDoRvAjustado = linha[299:300]

	if len(linha) > 313 { // Esse é o novo campo
		value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(linha[312:316])
	} else { // campo Antigo
		value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(linha[45:47])
	}

	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////// RegistroTipo036Antecipacoes ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo036Antecipacoes struct {
	Index                    int
	TipoDeRegistro           int
	NumeroDoPv               int
	NumeroDoDocumento        int
	DataDoLancamento         time.Time
	ValorDoLancamento        float64
	Credito                  string
	Banco                    int
	Agencia                  int
	ContaCorrente            int
	NumeroDoRvCorrespondente int
	DataDoRvCorrespondente   time.Time
	ValorDoCreditoOriginal   float64
	DataDoVencimentoOriginal time.Time
	NumeroDaParcelaTotal     string
	ValorBruto               float64
	ValorDaTaxaDeDesconto    float64
	NumeroPvOrginal          int
	Bandeira                 string
}

func PopularRegistroTipo036Antecipacoes(linha string, index int) RegistroTipo036Antecipacoes {
	value := RegistroTipo036Antecipacoes{}
	if len(linha) < 152 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoDocumento = goutils.ConvertStringToInt(linha[12:23])
	value.DataDoLancamento = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[23:31])
	value.ValorDoLancamento = goutils.ConvertStringToFloatScale2FormatNumber(linha[31:46])
	value.Credito = linha[46:47]
	value.Banco = goutils.ConvertStringToInt(linha[47:50])
	value.Agencia = goutils.ConvertStringToInt(linha[50:56])
	value.ContaCorrente = goutils.ConvertStringToInt(linha[56:67])
	value.NumeroDoRvCorrespondente = goutils.ConvertStringToInt(linha[67:76])
	value.DataDoRvCorrespondente = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[76:84])
	value.ValorDoCreditoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[84:99])
	value.DataDoVencimentoOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[99:107])
	value.NumeroDaParcelaTotal = linha[107:112]
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(linha[112:127])
	value.ValorDaTaxaDeDesconto = goutils.ConvertStringToFloatScale2FormatNumber(linha[127:142])
	value.NumeroPvOrginal = goutils.ConvertStringToInt(linha[142:151])
	value.Bandeira = linha[151:152]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// RegistroTipo037TotalizadorDeCreditos //////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo037TotalizadorDeCreditos struct {
	TipoDeRegistro                   int
	NumeroDoPv                       int
	Brancos                          string
	DataDoCredito                    time.Time
	ValorTotalDoCredito              float64
	BrancosCodigo                    string
	NumeroDoBanco                    int
	NumeroDaAgencia                  int
	NumeroDaContaCorrente            int
	DataDaGeracaoDoArquivo           time.Time
	DataDoCreditoAntecipado          time.Time
	ValorTotalDosCreditosAntecipados float64
}

func PopularRegistroTipo037TotalizadorDeCreditos(linha string) RegistroTipo037TotalizadorDeCreditos {
	value := RegistroTipo037TotalizadorDeCreditos{}
	if len(linha) < 94 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.Brancos = linha[12:19]
	value.DataDoCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[19:27])
	value.ValorTotalDoCredito = goutils.ConvertStringToFloatScale2FormatNumber(linha[27:42])
	value.BrancosCodigo = linha[42:43]
	value.NumeroDoBanco = goutils.ConvertStringToInt(linha[43:46])
	value.NumeroDaAgencia = goutils.ConvertStringToInt(linha[46:52])
	value.NumeroDaContaCorrente = goutils.ConvertStringToInt(linha[52:63])
	value.DataDaGeracaoDoArquivo = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[63:71])
	value.DataDoCreditoAntecipado = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[71:79])
	value.ValorTotalDosCreditosAntecipados = goutils.ConvertStringToFloatScale2FormatNumber(linha[79:94])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// RegistroTipo038AjustesADebitoViaBanco /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo038AjustesADebitoViaBanco struct {
	TipoDeRegistro                int
	NumeroDoPv                    int
	NumeroDoDocumento             int
	DataDaEmissao                 time.Time
	ValorDoDebito                 float64
	Debito                        string
	Banco                         int
	Agencia                       int
	ContaCorrente                 int
	NumeroDoRvOriginal            int
	DataRvOriginal                time.Time
	ValorDoCreditoOriginal        float64
	MotivoDoDebitoCodigo          int
	MotivoDoDebito                string
	NumeroDoCartao                string
	NumeroDeReferenciaDaCartaFax  string
	MesReferencia                 int
	DataDaCarta                   time.Time
	ValorDoCancelamentoSolicitado float64
	NumeroDoProcesso              int
	NumeroPvOriginal              int
	DataDaTransacaoCv             time.Time
	NumeroDoNsu                   int
	NumeroDoResumoDeDebito        int
	DataDoDebito                  time.Time
	ValorDaTransacaoOriginal      float64
	NumeroDaAutorizacao           string
	TipoDeDebito                  string
	ValorDoDebitoTotal            float64
	ValorPendente                 float64
	BandeiraDoRvDeOrigem          string
}

func PopularRegistroTipo038AjustesADebitoViaBanco(linha string) RegistroTipo038AjustesADebitoViaBanco {
	value := RegistroTipo038AjustesADebitoViaBanco{}
	if len(linha) < 303 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoDocumento = goutils.ConvertStringToInt(linha[12:23])
	value.DataDaEmissao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[23:31])
	value.ValorDoDebito = goutils.ConvertStringToFloatScale2FormatNumber(linha[31:46])
	value.Debito = linha[46:47]
	value.Banco = goutils.ConvertStringToInt(linha[47:50])
	value.Agencia = goutils.ConvertStringToInt(linha[50:56])
	value.ContaCorrente = goutils.ConvertStringToInt(linha[56:67])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[67:76])
	value.DataRvOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[76:84])
	value.ValorDoCreditoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[84:99])
	value.MotivoDoDebitoCodigo = goutils.ConvertStringToInt(linha[99:101])

	value.NumeroDoCartao = linha[129:145]
	value.NumeroDeReferenciaDaCartaFax = linha[145:160]
	value.MesReferencia = goutils.ConvertStringToInt(linha[160:166])
	value.DataDaCarta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[166:174])
	value.ValorDoCancelamentoSolicitado = goutils.ConvertStringToFloatScale2FormatNumber(linha[174:189])
	value.NumeroDoProcesso = goutils.ConvertStringToInt(linha[189:204])
	value.NumeroPvOriginal = goutils.ConvertStringToInt(linha[204:213])
	value.DataDaTransacaoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[213:221])
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[221:233])
	value.NumeroDoResumoDeDebito = goutils.ConvertStringToInt(linha[233:242])
	value.DataDoDebito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[242:250])
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[250:265])
	value.NumeroDaAutorizacao = linha[265:271]
	value.TipoDeDebito = linha[271:272]
	value.ValorDoDebitoTotal = goutils.ConvertStringToFloatScale2FormatNumber(linha[272:287])
	value.ValorPendente = goutils.ConvertStringToFloatScale2FormatNumber(linha[287:302])
	value.BandeiraDoRvDeOrigem = linha[302:303]

	if len(linha) > 303 { // Esse é o novo campo
		value.MotivoDoDebito = linha[303:307]
	} else { // campo Antigo
		value.MotivoDoDebito = linha[101:129]
	}

	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// RegistroTipo040Serasa /////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo040Serasa struct {
	TipoDeRegistro                           int
	NumeroDoPv                               int
	QuantidadeDeConsultasRealizadasNoPeriodo int
	ValortotalDasConsultasNoPeriodo          float64
	InicioDoPeriodoDaConsulta                time.Time
	FimDoPeriodoDaConsulta                   time.Time
	ValorPorConsultaDestePeriodo             float64
}

func PopularRegistroTipo040Serasa(linha string) RegistroTipo040Serasa {
	value := RegistroTipo040Serasa{}
	if len(linha) < 63 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeDeConsultasRealizadasNoPeriodo = goutils.ConvertStringToInt(linha[12:17])
	value.ValortotalDasConsultasNoPeriodo = goutils.ConvertStringToFloatScale2FormatNumber(linha[17:32])
	value.InicioDoPeriodoDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[32:40])
	value.FimDoPeriodoDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.ValorPorConsultaDestePeriodo = goutils.ConvertStringToFloatScale2FormatNumber(linha[48:63])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// RegistroTipo041AVS ////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo041AVS struct {
	TipoDeRegistro                           int
	NumeroDoPv                               int
	QuantidadeDeConsultasRealizadasNoPeriodo int
	ValorTotalDasConsultasNoPeriodo          float64
	InicioDoPeriodoDaConsulta                time.Time
	FimDoPeriodoDaConsulta                   time.Time
	ValorPorConsultaDestePeriodo             float64
}

func PopularRegistroTipo041AVS(linha string) RegistroTipo041AVS {
	value := RegistroTipo041AVS{}
	if len(linha) < 63 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeDeConsultasRealizadasNoPeriodo = goutils.ConvertStringToInt(linha[12:17])
	value.ValorTotalDasConsultasNoPeriodo = goutils.ConvertStringToFloatScale2FormatNumber(linha[17:32])
	value.InicioDoPeriodoDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[32:40])
	value.FimDoPeriodoDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.ValorPorConsultaDestePeriodo = goutils.ConvertStringToFloatScale2FormatNumber(linha[48:63])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// RegistroTipo042SecureCode /////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo042SecureCode struct {
	TipoDeRegistro                           int
	NumeroDoPv                               int
	QuantidadeDeConsultasRealizadasNoPeriodo int
	ValorTotalDasConsultasNoPeriodo          float64
	InicioDoPeriodoDaConsulta                time.Time
	FimDoPeriodoDaConsulta                   time.Time
	ValorPorConsultaDestePeriodo             float64
	Bandeira                                 string
}

func PopularRegistroTipo042SecureCode(linha string) RegistroTipo042SecureCode {
	value := RegistroTipo042SecureCode{}
	if len(linha) < 64 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeDeConsultasRealizadasNoPeriodo = goutils.ConvertStringToInt(linha[12:17])
	value.ValorTotalDasConsultasNoPeriodo = goutils.ConvertStringToFloatScale2FormatNumber(linha[17:32])
	value.InicioDoPeriodoDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[32:40])
	value.FimDoPeriodoDaConsulta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.ValorPorConsultaDestePeriodo = goutils.ConvertStringToFloatScale2FormatNumber(linha[48:63])
	value.Bandeira = linha[63:64]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////// RegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro ///////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro struct {
	Index                   int
	TipoDeRegistro          int
	NumeroDoPvCreditado     int
	NumeroDoResumoDeCredito int
	NumeroDoDocumento       int
	DataDeEmissao           time.Time
	DataDeCredito           time.Time
	ValorDoCredito          float64
	Credito                 string
	Banco                   int
	Agencia                 int
	ContaCorrente           string
	MotivoDoCreditoCodigo   int
	MotivoDoCredito         string
	Bandeira                string
}

func PopularRegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro(linha string, index int) RegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro {
	value := RegistroTipo043AjustesACreditoExtratoEletronicoFinanceiro{}
	if len(linha) < 115 {
		return value
	}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPvCreditado = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoResumoDeCredito = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDoDocumento = goutils.ConvertStringToInt(linha[21:32])
	value.DataDeEmissao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[32:40])
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[40:48])
	value.ValorDoCredito = goutils.ConvertStringToFloatScale2FormatNumber(linha[48:63])
	value.Credito = linha[63:64]
	value.Banco = goutils.ConvertStringToInt(linha[64:67])
	value.Agencia = goutils.ConvertStringToInt(linha[67:73])
	value.ContaCorrente = linha[73:84]

	value.MotivoDoCredito = linha[86:114]
	value.Bandeira = linha[114:115]

	if len(linha) > 116 { // Esse é o novo campo
		value.MotivoDoCreditoCodigo = goutils.ConvertStringToInt(linha[115:119])
	} else { // campo Antigo
		value.MotivoDoCreditoCodigo = goutils.ConvertStringToInt(linha[84:86])
	}

	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// RegistroTipo044DebitosPendentes ///////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo044DebitosPendentes struct {
	TipoDeRegistro                int
	NumeroDoPv                    int
	NumeroDaOrdemDeDebito         int
	DataDaOD                      time.Time
	ValorDaOD                     float64
	MotivoDoAjusteCodigo          int
	MotivoDoAjuste                string
	NumeroDoCartao                string
	NumeroDoNsu                   int
	DataDoCvOriginalDaTransacao   time.Time
	NumeroDaAutorizacao           string
	ValorDaTransacaoOriginal      float64
	NumeroDoRvOriginal            int
	DataDoRvOriginal              time.Time
	NumeroDoPvOriginal            int
	NumeroDaRefenrenciaDaCartaFax string
	DataDaCarta                   time.Time
	NumeroDoProcessoDeChargeback  int
	MesReferencia                 int
	ValorCompensadoPago           float64
	DataDoPagamento               time.Time
	ValorPendenteDeDebito         float64
	NumeroDoProcessoDeRetencao    int
	MeioDeCompensacaoCodigo       int
	MeioDeCompensacao             string
	Bandeira                      string
}

func PopularRegistroTipo044DebitosPendentes(linha string) RegistroTipo044DebitosPendentes {
	value := RegistroTipo044DebitosPendentes{}
	if len(linha) < 287 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDaOrdemDeDebito = goutils.ConvertStringToInt(linha[12:23])
	value.DataDaOD = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[23:31])
	value.ValorDaOD = goutils.ConvertStringToFloatScale2FormatNumber(linha[31:46])

	value.MotivoDoAjuste = linha[48:76]
	value.NumeroDoCartao = linha[76:92]
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[92:104])
	value.DataDoCvOriginalDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[104:112])
	value.NumeroDaAutorizacao = linha[112:118]
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[118:133])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[133:142])
	value.DataDoRvOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[142:150])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[150:159])
	value.NumeroDaRefenrenciaDaCartaFax = linha[159:174]
	value.DataDaCarta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[174:182])
	value.NumeroDoProcessoDeChargeback = goutils.ConvertStringToInt(linha[182:197])
	value.MesReferencia = goutils.ConvertStringToInt(linha[197:203])
	value.ValorCompensadoPago = goutils.ConvertStringToFloatScale2FormatNumber(linha[203:218])
	value.DataDoPagamento = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[218:226])
	value.ValorPendenteDeDebito = goutils.ConvertStringToFloatScale2FormatNumber(linha[226:241])
	value.NumeroDoProcessoDeRetencao = goutils.ConvertStringToInt(linha[241:256])
	value.MeioDeCompensacaoCodigo = goutils.ConvertStringToInt(linha[256:258])
	value.MeioDeCompensacao = linha[258:286]
	value.Bandeira = linha[286:287]

	if len(linha) > 287 { // Esse é o novo campo
		value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(linha[287:291])
	} else { // campo Antigo
		value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(linha[46:48])
	}

	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////// RegistroTipo045DebitosLiquidados ///////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo045DebitosLiquidados struct {
	TipoDeRegistro               int
	NumeroDoPv                   int
	NumeroDaOrdemDeDebito        int
	DataDaOD                     time.Time
	ValorDaOD                    float64
	MotivoDoAjusteCodigo         int
	MotivoDoAjuste               string
	NumeroDoCartao               string
	NumeroDoNsu                  int
	DataDoCvOriginalDaTransacao  time.Time
	NumeroDaAutorizacao          string
	ValorDaTransacaoOriginal     float64
	NumeroDoRvOriginal           int
	DataDoRvOriginal             time.Time
	NumeroDoPvOriginal           int
	NumeroDeRefenciaDaCartaFax   string
	DataDaCarta                  time.Time
	NumeroDoProcessoDeChargeback int
	MesRefenrecia                int
	ValorLiquidado               float64
	DataDaLiquidacao             time.Time
	NumeroDeProcessoDeRetencao   string
	MeioDeCompensacaoCodigo      int
	MeioDeCompensacao            string
	Bandeira                     string
}

func PopularRegistroTipo045DebitosLiquidados(linha string) RegistroTipo045DebitosLiquidados {
	value := RegistroTipo045DebitosLiquidados{}
	if len(linha) < 272 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPv = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDaOrdemDeDebito = goutils.ConvertStringToInt(linha[12:23])
	value.DataDaOD = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[23:31])
	value.ValorDaOD = goutils.ConvertStringToFloatScale2FormatNumber(linha[31:46])

	value.MotivoDoAjuste = linha[48:76]
	value.NumeroDoCartao = linha[76:92]
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[92:104])
	value.DataDoCvOriginalDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[104:112])
	value.NumeroDaAutorizacao = linha[112:118]
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[118:133])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[133:142])
	value.DataDoRvOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[142:150])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[150:159])
	value.NumeroDeRefenciaDaCartaFax = linha[159:174]
	value.DataDaCarta = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[174:182])
	value.NumeroDoProcessoDeChargeback = goutils.ConvertStringToInt(linha[182:197])
	value.MesRefenrecia = goutils.ConvertStringToInt(linha[197:203])
	value.ValorLiquidado = goutils.ConvertStringToFloatScale2FormatNumber(linha[203:218])
	value.DataDaLiquidacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[218:226])
	value.NumeroDeProcessoDeRetencao = linha[226:241]
	value.MeioDeCompensacaoCodigo = goutils.ConvertStringToInt(linha[241:243])
	value.MeioDeCompensacao = linha[243:271]
	value.Bandeira = linha[271:272]

	if len(linha) > 272 { // Esse é o novo campo
		value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(linha[272:276])
	} else { // campo Antigo
		value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(linha[46:48])
	}

	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////// RegistroTipo049DesagendamentosDeParcelas /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo049DesagendamentosDeParcelas struct {
	TipoDeRegistro                 int
	NumeroDoPvOriginal             int
	NumeroDoRvOriginal             int
	NumeroDeReferencia             string
	DataDeCredito                  time.Time
	NovoValorDaParcela             goutils.KeepZero
	ValorOriginalDaParcelaAlterada goutils.KeepZero
	ValorDoAjuste                  goutils.KeepZero
	DataDoCancelamento             time.Time
	ValorDoRvOriginal              goutils.KeepZero
	ValorDoCancelamentoSolicitado  goutils.KeepZero
	NumeroDoCartao                 string
	DataDaTransacao                time.Time
	NumeroNsu                      string
	TipoDeDebito                   int
	NumeroDaParcela                int
	BandeiraDoRvDeOrigem           string
}

func PopularRegistroTipo049DesagendamentosDeParcelas(linha string) RegistroTipo049DesagendamentosDeParcelas {
	value := RegistroTipo049DesagendamentosDeParcelas{}
	if len(linha) < 167 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[12:21])
	value.NumeroDeReferencia = linha[21:36]
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[36:44])
	value.NovoValorDaParcela = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[44:59]))
	value.ValorOriginalDaParcelaAlterada = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[59:74]))
	value.ValorDoAjuste = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[74:89]))
	value.DataDoCancelamento = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[89:97])
	value.ValorDoRvOriginal = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[97:112]))
	value.ValorDoCancelamentoSolicitado = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[112:127]))
	value.NumeroDoCartao = linha[127:143]
	value.DataDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[143:151])
	value.NumeroNsu = linha[151:163]
	value.TipoDeDebito = goutils.ConvertStringToInt(linha[163:164])
	value.NumeroDaParcela = goutils.ConvertStringToInt(linha[164:166])
	value.BandeiraDoRvDeOrigem = linha[166:167]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////////// RegistroTipo050TotalizadorMatriz ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo050TotalizadorMatriz struct {
	TipoDeRegistro                  int
	NumeroPvMatriz                  int
	QuantidadeTotalDeResumosMatriz  int
	ValorTotalDosCreditosNormais    float64
	QuantidadeDeCreditosAntecipados int
	ValorTotalAntecipado            float64
	QuantidadeDeAjustesACredito     int
	ValorTotalDeAjustesACredito     float64
	QuantidadeDeAjustesADebito      int
	ValorTotalDeAjustesADebito      float64
}

func PopularRegistroTipo050TotalizadorMatriz(linha string) RegistroTipo050TotalizadorMatriz {
	value := RegistroTipo050TotalizadorMatriz{}
	if len(linha) < 94 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroPvMatriz = goutils.ConvertStringToInt(linha[3:12])
	value.QuantidadeTotalDeResumosMatriz = goutils.ConvertStringToInt(linha[12:18])
	value.ValorTotalDosCreditosNormais = goutils.ConvertStringToFloatScale2FormatNumber(linha[18:33])
	value.QuantidadeDeCreditosAntecipados = goutils.ConvertStringToInt(linha[33:39])
	value.ValorTotalAntecipado = goutils.ConvertStringToFloatScale2FormatNumber(linha[39:54])
	value.QuantidadeDeAjustesACredito = goutils.ConvertStringToInt(linha[54:58])
	value.ValorTotalDeAjustesACredito = goutils.ConvertStringToFloatScale2FormatNumber(linha[58:73])
	value.QuantidadeDeAjustesADebito = goutils.ConvertStringToInt(linha[73:79])
	value.ValorTotalDeAjustesADebito = goutils.ConvertStringToFloatScale2FormatNumber(linha[79:94])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////// TrailerDoArquivo //////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo052TrailerDoArquivo struct {
	TipoDeRegistro                  int
	QuantidadeDeMatrizesNoArquivo   int
	QuantidadeDeRegistrosNoArquivo  int
	NumeroPvGrupo                   int
	QuantidadeTotalDeResumosGrupo   int
	ValorTotalDosCreditosNormais    float64
	QuantidadeDeCreditosAntecipados int
	ValorTotalAntecipado            goutils.KeepZero
	QuantidadeDeAjustesACredito     int
	ValorTotalDeAjustesACredito     float64
	QuantidadeDeAjustesADebito      int
	ValorTotalDeAjustesADebito      float64
}

func PopularRegistroTipo052TrailerDoArquivo(linha string) RegistroTipo052TrailerDoArquivo {
	value := RegistroTipo052TrailerDoArquivo{}
	if len(linha) < 100 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.QuantidadeDeMatrizesNoArquivo = goutils.ConvertStringToInt(linha[3:7])
	value.QuantidadeDeRegistrosNoArquivo = goutils.ConvertStringToInt(linha[7:13])
	value.NumeroPvGrupo = goutils.ConvertStringToInt(linha[13:22])
	value.QuantidadeTotalDeResumosGrupo = goutils.ConvertStringToInt(linha[22:26])
	value.ValorTotalDosCreditosNormais = goutils.ConvertStringToFloatScale2FormatNumber(linha[26:41])
	value.QuantidadeDeCreditosAntecipados = goutils.ConvertStringToInt(linha[41:47])
	value.ValorTotalAntecipado = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(linha[47:62]))
	value.QuantidadeDeAjustesACredito = goutils.ConvertStringToInt(linha[62:66])
	value.ValorTotalDeAjustesACredito = goutils.ConvertStringToFloatScale2FormatNumber(linha[66:81])
	value.QuantidadeDeAjustesADebito = goutils.ConvertStringToInt(linha[81:85])
	value.ValorTotalDeAjustesADebito = goutils.ConvertStringToFloatScale2FormatNumber(linha[85:100])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////// RegistroTipo054AjustesADebitoViaBancoECommerce /////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo053AjustesNetEDesagendamentosECommerce struct {
	TipoDeRegistro           int
	NumeroDoCartao           string
	DataDaTransacaoCv        time.Time
	NumeroDoRvOriginal       int
	NumeroDoPvOriginal       int
	ValorDaTransacaoOriginal float64
	NumeroDoNsu              int
	NumeroDaAutorizacao      string
	Tid                      string
	NumeroDoPedido           string
}

func PopularRegistroTipo053AjustesNetEDesagendamentosECommerce(linha string) RegistroTipo053AjustesNetEDesagendamentosECommerce {
	value := RegistroTipo053AjustesNetEDesagendamentosECommerce{}
	if len(linha) < 128 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoCartao = linha[3:19]
	value.DataDaTransacaoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[19:27])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[27:36])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[36:45])
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[45:60])
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[60:72])
	value.NumeroDaAutorizacao = linha[72:78]
	value.Tid = linha[78:98]
	value.NumeroDoPedido = linha[98:128]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////// RegistroTipo054AjustesADebitoViaBancoECommerce /////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo054AjustesADebitoViaBancoECommerce struct {
	TipoDeRegistro           int
	NumeroDoRvOriginal       int
	NumeroDoCartao           string
	NumeroDoPvOriginal       int
	DataDaTransacaoCv        time.Time
	NumeroDoNsu              int
	ValorDaTransacaoOriginal float64
	NumeroDaAutorizacao      string
	Tid                      string
	NumeroDoPedido           string
}

func PopularRegistroTipo054AjustesADebitoViaBancoECommerce(linha string) RegistroTipo054AjustesADebitoViaBancoECommerce {
	value := RegistroTipo054AjustesADebitoViaBancoECommerce{}
	if len(linha) < 128 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoCartao = linha[12:28]
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[28:37])
	value.DataDaTransacaoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[37:45])
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[45:57])
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[57:72])
	value.NumeroDaAutorizacao = linha[72:78]
	value.Tid = linha[78:98]
	value.NumeroDoPedido = linha[98:128]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo055DebitosPendentes /////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo055DebitosPendentes struct {
	TipoDeRegistro              int
	NumeroDoCartao              string
	NumeroDoNsu                 int
	DataDoCvOriginalDaTransacao time.Time
	NumeroDeAutorizacao         string
	ValorDaTransacaoOriginal    float64
	NumeroDoRvOriginal          int
	NumeroDoPvOriginal          int
	Tid                         string
	NumeroDoPedido              string
}

func PopularRegistroTipo055DebitosPendentes(linha string) RegistroTipo055DebitosPendentes {
	value := RegistroTipo055DebitosPendentes{}
	if len(linha) < 128 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoCartao = linha[3:19]
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[19:31])
	value.DataDoCvOriginalDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[31:39])
	value.NumeroDeAutorizacao = linha[39:45]
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[45:60])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[60:69])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[69:78])
	value.Tid = linha[78:98]
	value.NumeroDoPedido = linha[98:128]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////////////// RegistroTipo056DebitosLiquidadosECommerce //////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo056DebitosLiquidadosECommerce struct {
	TipoDeRegistro              int
	NumeroDoCartao              string
	NumeroDoNsu                 int
	DataDoCvOriginalDaTransacao time.Time
	NumeroDaAutorizacao         string
	ValorDaTransacaoOriginal    float64
	NumeroDoRvOriginal          int
	NumeroDoPvOriginal          int
	Tid                         string
	NumeroDoPedido              string
}

func PopularRegistroTipo056DebitosLiquidadosECommerce(linha string) RegistroTipo056DebitosLiquidadosECommerce {
	value := RegistroTipo056DebitosLiquidadosECommerce{}
	if len(linha) < 128 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoCartao = linha[3:19]
	value.NumeroDoNsu = goutils.ConvertStringToInt(linha[19:31])
	value.DataDoCvOriginalDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[31:39])
	value.NumeroDaAutorizacao = linha[39:45]
	value.ValorDaTransacaoOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[45:60])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[60:69])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[69:78])
	value.Tid = linha[78:98]
	value.NumeroDoPedido = linha[98:128]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////// RegistroTipo057DesagendamentosDeParcelasECommerce ////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo057DesagendamentosDeParcelasECommerce struct {
	TipoDeRegistro     int
	NumeroDoPvOriginal int
	NumeroDoRvOriginal int
	ValorDoRvOriginal  float64
	NumeroDoCartao     string
	DataDaTransacao    time.Time
	NumeroNsu          int
	Tid                string
	NumeroDoPedido     string
}

func PopularRegistroTipo057DesagendamentosDeParcelasECommerce(linha string) RegistroTipo057DesagendamentosDeParcelasECommerce {
	value := RegistroTipo057DesagendamentosDeParcelasECommerce{}
	if len(linha) < 125 {
		return value
	}
	value.TipoDeRegistro = goutils.ConvertStringToInt(linha[0:3])
	value.NumeroDoPvOriginal = goutils.ConvertStringToInt(linha[3:12])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(linha[12:21])
	value.ValorDoRvOriginal = goutils.ConvertStringToFloatScale2FormatNumber(linha[21:35])
	value.NumeroDoCartao = linha[35:52]
	value.DataDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(linha[52:60])
	value.NumeroNsu = goutils.ConvertStringToInt(linha[60:75])
	value.Tid = linha[75:95]
	value.NumeroDoPedido = linha[95:125]
	return value
}
