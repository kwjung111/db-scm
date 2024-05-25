package object

type object interface {
}

type trigger struct {
	Name string
	Db   string
	Def  string
}

type procedure struct {
	Name string
	Db   string
	Def  string
}
