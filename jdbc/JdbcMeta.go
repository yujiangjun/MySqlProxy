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
	meta.Password = "83baad54"
	meta.Url = "tcp(49.234.25.117:33060)/daxin?charset=utf8&parseTime=true"
	return meta
}
