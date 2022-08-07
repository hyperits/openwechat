package openwechat

// ConfInfo 配置struct
type ConfInfo struct {
	User     string
	Passw    string
	IP       string
	Port     int
	DataBase string
	Table    string
}

var Configer *ConfInfo

func init() {
	Configer = &ConfInfo{
		User:  "root",
		Passw: "hyperits",
		IP:    "114.132.166.120",
		Port:  3306,
	}
}
