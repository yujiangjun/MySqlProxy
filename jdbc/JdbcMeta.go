package jdbc

type MetaService interface {
	GetMeta() Meta
}

type Meta struct {
	Username string
	Password string
	Url      string
}

func (meta Meta) GetMeta() Meta {
	meta.Username = "root"
	meta.Password = "123456"
	meta.Url = "tcp(127.0.0.1:3306)/raytine?charset=utf8&parseTime=true"
	return meta
}
