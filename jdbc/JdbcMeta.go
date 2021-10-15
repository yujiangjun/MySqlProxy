package jdbc

type MetaService interface {
	getMeta()
}

type Meta struct {
	username string
	password string
	url      string
}

func (meta Meta) call() Meta {
	meta.username = "root"
	meta.password = "123456"
	meta.url = "tcp(192.168.136.136:3306)/test?charset=utf8"
	return meta
}
