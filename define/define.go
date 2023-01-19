package define

const Address = "192.168.0.100" //服务地址

var (
	DataTime        = "2006-01-02 15"
	FTPUserName     = ""                               // ftp 姓名
	FTPPassword     = ""                               // ftp 密码
	FTPAddress      = Address + ":21"                  // ftp 地址
	PlayUrlAddress  = "http://" + Address + "/video/"  // ftp视频路径
	CoverUrlAddress = "http://" + Address + "/images/" // ftp图片路径
)
