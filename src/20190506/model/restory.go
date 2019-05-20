package model

type Student struct {
	Name string
	Score float64
}

type clazz struct {
	Name string
	Score float64
}

func NewClazz(n string, s float64) *clazz {
	return &clazz{
		Name: n,
		Score: s,
	}
}