package model

type Command string

func (c Command) String() string {
	return string(c)
}

type Query string

func (q Query) String() string {
	return string(q)
}
