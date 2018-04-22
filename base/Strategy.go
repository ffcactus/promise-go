package base

// StrategyInterface is the interface that a Strategy have.
type StrategyInterface interface {
	Execute(ctx ContextInterface) (ResponseInterface, *string, []Message)
}
