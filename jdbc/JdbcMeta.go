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
	meta.Url = "tcp(192.168.211.130:3306)/?charset=utf8"
	return meta
}
