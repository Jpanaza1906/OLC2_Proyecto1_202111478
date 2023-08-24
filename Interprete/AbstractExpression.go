package interprete

type AbstractExpression interface {
	Interpretar(ctx *Contexto) *Resultado
}
