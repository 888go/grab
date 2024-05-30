# 备注开始
# **_方法.md 文件备注:
# ff= 方法,重命名方法名称
# 如://ff:取文本
#
# yx=true,此方法优先翻译
# 如: //yx=true

# **_package.md 文件备注:
# bm= 包名,更换新的包名称 
# 如: package gin //bm:gin类

# **_其他.md 文件备注:
# qm= 前面,跳转到前面进行重命名.文档内如果有多个相同的,会一起重命名.
# hm= 后面,跳转到后面进行重命名.文档内如果有多个相同的,会一起重命名.
# cz= 查找,配合前面/后面使用,
# 如: type Regexp struct {//qm:正则 cz:Regexp struct
#
# th= 替换,用于替换文本,文档内如果有多个相同的,会一起替换
# 如:
# type Regexp struct {//th:type Regexp222 struct
#
# cf= 重复,用于重命名多次,
# 如: 
# 一个文档内有2个"One(result interface{}) error"需要重命名.
# 但是要注意,多个新名称要保持一致. 如:"X取一条(result interface{})"

# **_追加.md 文件备注:
# 在代码内追加代码,如:
# //zj:
# func (re *Regexp) X取文本() string { 
# re.F.String()
# }
# //zj:
# 备注结束

[Label string]
qm=名称
hm=
cz=Label string

[HTTPRequest *http.Request]
qm=Http协议头
hm=
cz=HTTPRequest *http.Request

[Filename string]
qm=文件名
hm=
cz=Filename string

[SkipExisting bool]
qm=跳过已存在文件
hm=
cz=SkipExisting bool

[NoResume bool]
qm=不续传
hm=
cz=NoResume bool

[NoStore bool]
qm=不写入本地文件系统
hm=
cz=NoStore bool

[NoCreateDirectories bool]
qm=不自动创建目录
hm=
cz=NoCreateDirectories bool

[IgnoreBadStatusCodes bool]
qm=忽略错误状态码
hm=
cz=IgnoreBadStatusCodes bool

[IgnoreRemoteTime bool]
qm=忽略远程时间
hm=
cz=IgnoreRemoteTime bool

[Size int64]
qm=预期文件大小
hm=
cz=Size int64

[BufferSize int]
qm=缓冲区大小
hm=
cz=BufferSize int

[RateLimiter RateLimiter]
qm=速率限制器
hm=
cz=RateLimiter RateLimiter

[BeforeCopy Hook]
qm=传输开始之前回调
hm=
cz=BeforeCopy Hook

[AfterCopy Hook]
qm=传输完成之后回调
hm=
cz=AfterCopy Hook
