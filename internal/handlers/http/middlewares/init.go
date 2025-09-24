package middlewares

type Middlewares struct {
}

type DataInitMiddlewares struct {
}

func NewMiddlewares(data DataInitMiddlewares) *Middlewares {
	return &Middlewares{}
}
