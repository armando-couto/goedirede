package dtos

import (
	"github.com/armando-couto/goutils"
	"time"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo00CabecalhoDoArquivo ////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo00CabecalhoDoArquivo struct {
	TipoDeRegistro                           int
	NumeroDaFiliacaoDaMatrizOuGrupoComercial int
	DataDeEmissao                            time.Time
	DataDeMovimentacao                       time.Time
	RedeAdquirente                           string
	TipoDoAqrquivo                           string
	NomeComercial                            string
	SequenciaDoMovimento                     int
	TipoDeProcessamento                      string
	Versao                                   string
}

func PopularRegistroTipo00CabecalhoDoArquivo(coluna []string) RegistroTipo00CabecalhoDoArquivo {
	value := RegistroTipo00CabecalhoDoArquivo{}
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDaMatrizOuGrupoComercial = goutils.ConvertStringToInt(coluna[1])
	value.DataDeEmissao = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[2])
	value.DataDeMovimentacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.RedeAdquirente = coluna[4]
	value.TipoDoAqrquivo = coluna[5]
	value.NomeComercial = coluna[6]
	value.SequenciaDoMovimento = goutils.ConvertStringToInt(coluna[7])
	value.TipoDeProcessamento = coluna[8]
	value.Versao = coluna[9]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo01ResumoDeVendas ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo01ResumoDeVendas struct {
	Index                                    int
	TipoDeRegistro                           int
	NumeroDaFiliacaoDaMatrizOuGrupoComercial int
	DataDeCredito                            time.Time
	DataDoResumoDeVendas                     time.Time
	NumeroDeResumoDeVendas                   int
	QuantidadeDeComprovantesDeVendas         int
	ValorBruto                               goutils.KeepZero
	ValorDesconto                            goutils.KeepZero
	ValorLiquido                             goutils.KeepZero
	TipoDeResumo                             string
	Banco                                    int
	Agencia                                  int
	ContaCorrente                            int
	Bandeira                                 string
}

func PopularRegistroTipo01ResumoDeVendas(coluna []string, index int) RegistroTipo01ResumoDeVendas {
	value := RegistroTipo01ResumoDeVendas{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDaMatrizOuGrupoComercial = goutils.ConvertStringToInt(coluna[1])
	value.DataDeCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[2])
	value.DataDoResumoDeVendas = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.NumeroDeResumoDeVendas = goutils.ConvertStringToInt(coluna[4])
	value.QuantidadeDeComprovantesDeVendas = goutils.ConvertStringToInt(coluna[5])
	value.ValorBruto = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[6]))
	value.ValorDesconto = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[7]))
	value.ValorLiquido = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[8]))
	value.TipoDeResumo = coluna[9]
	value.Banco = goutils.ConvertStringToInt(coluna[10])
	value.Agencia = goutils.ConvertStringToInt(coluna[11])
	value.ContaCorrente = goutils.ConvertStringToInt(coluna[12])
	value.Bandeira = coluna[13]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo02TotalDoPontoDeVenda ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo02TotalDoPontoDeVenda struct {
	TipoDeRegistro                        int
	NumeroDaFiliacaoDaMatriz              int
	QuantidadeDereResumosDeVendasAtacados int
	QuantidadeDeComprovantesDeVendas      int
	ValorBruto                            float64
	ValorDesconto                         float64
	ValorLiquido                          float64
	ValorBrutoPreDatado                   float64
	ValorDescontoPreDatado                float64
	ValorLiquidoPreDatado                 float64
}

func PopularRegistroTipo02TotalDoPontoDeVenda(coluna []string) RegistroTipo02TotalDoPontoDeVenda {
	value := RegistroTipo02TotalDoPontoDeVenda{}
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDaMatriz = goutils.ConvertStringToInt(coluna[1])
	value.QuantidadeDereResumosDeVendasAtacados = goutils.ConvertStringToInt(coluna[2])
	value.QuantidadeDeComprovantesDeVendas = goutils.ConvertStringToInt(coluna[3])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[4])
	value.ValorDesconto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[5])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(coluna[6])
	value.ValorBrutoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[7])
	value.ValorDescontoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[8])
	value.ValorLiquidoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[9])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////// RegistroTipo03TotalDaMatriz ///////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo03TotalDaMatriz struct {
	TipoDeRegistro                        int
	NumeroDaFiliacaoDaMatriz              int
	QuantidadeDereResumosDeVendasAtacados int
	QuantidadeDeComprovantesDeVendas      int
	ValorBruto                            float64
	ValorDesconto                         float64
	ValorLiquido                          float64
	ValorBrutoPreDatado                   float64
	ValorDescontoPreDatado                float64
	ValorLiquidoPreDatado                 float64
}

func PopularRegistroTipo03TotalDaMatriz(coluna []string) RegistroTipo03TotalDaMatriz {
	value := RegistroTipo03TotalDaMatriz{}
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDaMatriz = goutils.ConvertStringToInt(coluna[1])
	value.QuantidadeDereResumosDeVendasAtacados = goutils.ConvertStringToInt(coluna[2])
	value.QuantidadeDeComprovantesDeVendas = goutils.ConvertStringToInt(coluna[3])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[4])
	value.ValorDesconto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[5])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(coluna[6])
	value.ValorBrutoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[7])
	value.ValorDescontoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[8])
	value.ValorLiquidoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[9])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////// RegistroTipo04TotalDoArquivo ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo04TotalDoArquivo struct {
	TipoDeRegistro                           int
	NumeroDaFiliacaoDaMatrizOuGrupoComercial int
	QuantidadeDereResumosDeVendasAtacados    int
	QuantidadeDeComprovantesDeVendas         int
	ValorBruto                               float64
	ValorDesconto                            float64
	ValorLiquido                             float64
	ValorBrutoPreDatado                      float64
	ValorDescontoPreDatado                   float64
	ValorLiquidoPreDatado                    float64
	TotalDeRegistrosNoArquivo                int
}

func PopularRegistroTipo04TotalDoArquivo(coluna []string) RegistroTipo04TotalDoArquivo {
	value := RegistroTipo04TotalDoArquivo{}
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDaMatrizOuGrupoComercial = goutils.ConvertStringToInt(coluna[1])
	value.QuantidadeDereResumosDeVendasAtacados = goutils.ConvertStringToInt(coluna[2])
	value.QuantidadeDeComprovantesDeVendas = goutils.ConvertStringToInt(coluna[3])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[4])
	value.ValorDesconto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[5])
	value.ValorLiquido = goutils.ConvertStringToFloatScale2FormatNumber(coluna[6])
	value.ValorBrutoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[7])
	value.ValorDescontoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[8])
	value.ValorLiquidoPreDatado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[9])
	value.TotalDeRegistrosNoArquivo = goutils.ConvertStringToInt(coluna[10])
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////// RegistroTipo05DetalhamentoDosComprovantesDeVendas /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo05DetalhamentoDosComprovantesDeVendas struct {
	Index                           int
	TipoDeRegistro                  int
	NumeroDaFiliacaoDoPontoDeVendas int
	NumeroDoResumoDeVendas          int
	DataDoCv                        time.Time
	ValorBruto                      goutils.KeepZero
	ValorDesconto                   goutils.KeepZero
	ValorLiquido                    goutils.KeepZero
	NumeroDoCartao                  string
	TipoDeTransacao                 string
	NumeroDoCv                      string
	DataDoCredito                   time.Time
	StatusDaTransacao               int
	HoraDaTransacao                 time.Time
	NumeroDoTerminal                string
	TipoDeCaptura                   int
	Reservado                       int
	ValorDaCompra                   float64
	ValorDoSaque                    float64
	Bandeira                        string
}

func PopularRegistroTipo05DetalhamentoDosComprovantesDeVendas(coluna []string, index int) RegistroTipo05DetalhamentoDosComprovantesDeVendas {
	value := RegistroTipo05DetalhamentoDosComprovantesDeVendas{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDoPontoDeVendas = goutils.ConvertStringToInt(coluna[1])
	value.NumeroDoResumoDeVendas = goutils.ConvertStringToInt(coluna[2])
	value.DataDoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.ValorBruto = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[4]))
	value.ValorDesconto = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[5]))
	value.ValorLiquido = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[6]))
	value.NumeroDoCartao = coluna[7]
	value.TipoDeTransacao = coluna[8]
	value.NumeroDoCv = coluna[9]
	value.DataDoCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[10])
	value.StatusDaTransacao = goutils.ConvertStringToInt(coluna[11])
	value.HoraDaTransacao = goutils.ConvertStringToTimeLayoutHHMMSS(coluna[12])
	value.NumeroDoTerminal = coluna[13]
	value.TipoDeCaptura = goutils.ConvertStringToInt(coluna[14])
	value.Reservado = goutils.ConvertStringToInt(coluna[15])
	value.ValorDaCompra = goutils.ConvertStringToFloatScale2FormatNumber(coluna[16])
	value.ValorDaCompra = goutils.ConvertStringToFloatScale2FormatNumber(coluna[17])
	value.Bandeira = coluna[18]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////// RegistroTipo08DesagendamentoDeVendasPreDatadas /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo08DesagendamentoDeVendasPreDatadas struct {
	Index                          int
	TipoDeRegistro                 int
	NumeroDaFiliacaoDoPontoDeVenda int
	NumeroDoResumoDeVendas         string
	DataDoCv                       time.Time
	NumeroDoCv                     string
	ValorBrutoDoCv                 goutils.KeepZero
	ValorDoCancelamento            goutils.KeepZero
	MotivoDoCancelamento           string
	DataDoCredito                  time.Time
	NovoValorDoCredito             float64
	TipoDeTransacao                string
}

func PopularRegistroTipo08DesagendamentoDeVendasPreDatadas(coluna []string, index int) RegistroTipo08DesagendamentoDeVendasPreDatadas {
	value := RegistroTipo08DesagendamentoDeVendasPreDatadas{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDoPontoDeVenda = goutils.ConvertStringToInt(coluna[1])
	value.NumeroDoResumoDeVendas = coluna[2]
	value.DataDoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.NumeroDoCv = coluna[4]
	value.ValorBrutoDoCv = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[5]))
	value.ValorDoCancelamento = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[6]))
	value.MotivoDoCancelamento = coluna[7]
	value.DataDoCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[8])
	value.NovoValorDoCredito = goutils.ConvertStringToFloatScale2FormatNumber(coluna[9])
	value.TipoDeTransacao = coluna[10]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
///////////////////// RegistroTipo09TransacoesPreDatadasLiquidadas /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo09TransacoesPreDatadasLiquidadas struct {
	Index                          int
	TipoDeRegistro                 int
	NumeroDaFiliacaoDoPontoDeVenda int
	DataDaLiquidacao               time.Time
	ValorLiquidado                 float64
	NumeroDoNsuDaTransacao         int
	DataDaTransacao                time.Time
	DataDoVencimento               time.Time
	ValorDesconto                  float64
	ValorBruto                     float64
	Bandeira                       string
}

func PopularRegistroTipo09TransacoesPreDatadasLiquidadas(coluna []string, index int) RegistroTipo09TransacoesPreDatadasLiquidadas {
	value := RegistroTipo09TransacoesPreDatadasLiquidadas{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDoPontoDeVenda = goutils.ConvertStringToInt(coluna[1])
	value.DataDaLiquidacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[2])
	value.ValorLiquidado = goutils.ConvertStringToFloatScale2FormatNumber(coluna[3])
	value.NumeroDoNsuDaTransacao = goutils.ConvertStringToInt(coluna[4])
	value.DataDaTransacao = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[5])
	value.DataDoVencimento = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[6])
	value.ValorDesconto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[7])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[8])
	value.Bandeira = coluna[9]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
/////////////////////////////////// RegistroTipo11AjustesNet /////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo11AjustesNet struct {
	Index                              int
	TipoDeRegistro                     int
	NumeroDoPvAjustado                 int
	NumeroDoRvAjustado                 int
	DataDoAjuste                       time.Time
	ValorDoAjuste                      goutils.KeepZero
	Debito                             string
	MotivoDoAjusteCodigo               int
	MotivoDoAjuste                     string
	NumeroDoCartao                     int
	DataDaTransacaoCv                  time.Time
	NumeroDoRvOriginal                 int
	NumeroDeReferenciaDaCartaFax       int
	DataDaCarta                        time.Time
	MesDeReferencia                    int
	NumeroPvOriginal                   int
	DataRvOriginal                     time.Time
	ValorDaTransacao                   goutils.KeepZero
	Net                                string
	DataDoCredito                      time.Time
	ValorBrutoDoResumoDeVendasOriginal goutils.KeepZero
	ValorDoCancelamentoSolicitado      goutils.KeepZero
	NumeroDoNsu                        string
	NumeroDaAutorizacao                string
	TipoDeBebito                       string
	NumeroDaordemDeDebito              int
	ValorDoDebitoTotal                 goutils.KeepZero
	ValorPendente                      goutils.KeepZero
	BandeiraDoRvDeOrigem               string
	BandeiraDoRvAjustado               string
}

func PopularRegistroTipo11AjustesNet(coluna []string, index int) RegistroTipo11AjustesNet {
	value := RegistroTipo11AjustesNet{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDoPvAjustado = goutils.ConvertStringToInt(coluna[1])
	value.NumeroDoRvAjustado = goutils.ConvertStringToInt(coluna[2])
	value.DataDoAjuste = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.ValorDoAjuste = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[4]))
	value.Debito = coluna[5]
	value.MotivoDoAjusteCodigo = goutils.ConvertStringToInt(coluna[6])
	value.MotivoDoAjuste = coluna[7]
	value.NumeroDoCartao = goutils.ConvertStringToInt(coluna[8])
	value.DataDaTransacaoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[9])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(coluna[10])
	value.NumeroDeReferenciaDaCartaFax = goutils.ConvertStringToInt(coluna[11])
	value.DataDaCarta = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[12])
	value.MesDeReferencia = goutils.ConvertStringToInt(coluna[13])
	value.NumeroPvOriginal = goutils.ConvertStringToInt(coluna[14])
	value.DataRvOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[15])
	value.ValorDaTransacao = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[16]))
	value.Net = coluna[17]
	value.DataDoCredito = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[18])
	value.ValorBrutoDoResumoDeVendasOriginal = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[19]))
	value.ValorDoCancelamentoSolicitado = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[20]))
	value.NumeroDoNsu = coluna[21]
	value.NumeroDaAutorizacao = coluna[22]
	value.TipoDeBebito = coluna[23]
	value.NumeroDaordemDeDebito = goutils.ConvertStringToInt(coluna[24])
	value.ValorDoDebitoTotal = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[25]))
	value.ValorPendente = goutils.KeepZero(goutils.ConvertStringToFloatScale2FormatNumber(coluna[26]))
	value.BandeiraDoRvDeOrigem = coluna[27]
	value.BandeiraDoRvAjustado = coluna[28]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////// RegistroTipo13DetalhamentoDosComprovantesDeVendasEcommerce ////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo13DetalhamentoDosComprovantesDeVendasEcommerce struct {
	Index                           int
	TipoDeRegistro                  int
	NumeroDaFiliacaoDoPontoDeVendas int
	NumeroDoResumoDeVendas          int
	DataDoCv                        time.Time
	ValorBruto                      float64
	NumeroDoCartao                  string
	NumeroDoCv                      int
	Tid                             string
	NumeroPedido                    string
}

func PopularRegistroTipo13DetalhamentoDosComprovantesDeVendasEcommerce(coluna []string, index int) RegistroTipo13DetalhamentoDosComprovantesDeVendasEcommerce {
	value := RegistroTipo13DetalhamentoDosComprovantesDeVendasEcommerce{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDoPontoDeVendas = goutils.ConvertStringToInt(coluna[1])
	value.NumeroDoResumoDeVendas = goutils.ConvertStringToInt(coluna[2])
	value.DataDoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.ValorBruto = goutils.ConvertStringToFloatScale2FormatNumber(coluna[4])
	value.NumeroDoCartao = coluna[5]
	value.NumeroDoCv = goutils.ConvertStringToInt(coluna[6])
	value.Tid = coluna[7]
	value.NumeroPedido = coluna[8]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////// RegistroTipo17AjustesNetEcommerce /////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo17AjustesNetEcommerce struct {
	Index               int
	TipoDeRegistro      int
	NumeroDoCartao      int
	DataDaTransacaoCv   time.Time
	NumeroDoRvOriginal  int
	NumeroPvOriginal    int
	DataRvOriginal      time.Time
	ValorDaTransacao    float64
	NumeroDoNsu         int
	NumeroDaAutorizacao string
	Tid                 string
	NumeroDoPedido      string
}

func PopularRegistroTipo17AjustesNetEcommerce(coluna []string, index int) RegistroTipo17AjustesNetEcommerce {
	value := RegistroTipo17AjustesNetEcommerce{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDoCartao = goutils.ConvertStringToInt(coluna[1])
	value.DataDaTransacaoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[2])
	value.NumeroDoRvOriginal = goutils.ConvertStringToInt(coluna[3])
	value.NumeroPvOriginal = goutils.ConvertStringToInt(coluna[4])
	value.DataRvOriginal = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[5])
	value.ValorDaTransacao = goutils.ConvertStringToFloatScale2FormatNumber(coluna[6])
	value.NumeroDoNsu = goutils.ConvertStringToInt(coluna[7])
	value.NumeroDaAutorizacao = coluna[8]
	value.Tid = coluna[9]
	value.NumeroDoPedido = coluna[10]
	return value
}

//////////////////////////////////////////////////////////////////////////////////////////////////////
////////////// RegistroTipo20DetalhamentoDosComprovantesDeVendasOperadorasDeCelular //////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////////
type RegistroTipo20DetalhamentoDosComprovantesDeVendasOperadorasDeCelular struct {
	Index                                    int
	TipoDeRegistro                           int
	NumeroDaFiliacaoDaMatrizOuGrupoComercial int
	NumeroDoResumoDeVendas                   int
	DataDoCv                                 time.Time
	ValorRecarga                             float64
	NumeroDoComprovantes                     int
	NumeroDoTelefone                         int
	Bandeira                                 string
	CodigoDeAutorizacao                      string
}

func PopularRegistroTipo20DetalhamentoDosComprovantesDeVendasOperadorasDeCelular(coluna []string, index int) RegistroTipo20DetalhamentoDosComprovantesDeVendasOperadorasDeCelular {
	value := RegistroTipo20DetalhamentoDosComprovantesDeVendasOperadorasDeCelular{}
	value.Index = index
	value.TipoDeRegistro = goutils.ConvertStringToInt(coluna[0])
	value.NumeroDaFiliacaoDaMatrizOuGrupoComercial = goutils.ConvertStringToInt(coluna[1])
	value.NumeroDoResumoDeVendas = goutils.ConvertStringToInt(coluna[2])
	value.DataDoCv = goutils.ConvertStringToTimeLayoutDDMMYYYY(coluna[3])
	value.ValorRecarga = goutils.ConvertStringToFloatScale2FormatNumber(coluna[4])
	value.NumeroDoComprovantes = goutils.ConvertStringToInt(coluna[5])
	value.NumeroDoTelefone = goutils.ConvertStringToInt(coluna[6])
	value.Bandeira = coluna[7]
	value.CodigoDeAutorizacao = coluna[8]
	return value
}
