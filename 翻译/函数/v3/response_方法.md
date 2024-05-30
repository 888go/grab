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

[func (c *Response) IsComplete() bool {]
ff=是否已完成

[func (c *Response) Cancel() error {]
ff=取消

[func (c *Response) Wait() {]
ff=等待完成

[func (c *Response) Err() error {]
ff=等待错误

[func (c *Response) Size() int64 {]
ff=取总字节

[func (c *Response) BytesComplete() int64 {]
ff=已完成字节

[func (c *Response) BytesPerSecond() float64 {]
ff=取每秒字节

[func (c *Response) Progress() float64 {]
ff=取进度

[func (c *Response) Duration() time.Duration {]
ff=取下载已持续时间

[func (c *Response) ETA() time.Time {]
ff=取估计完成时间

[func (c *Response) Open() (io.ReadCloser, error) {]
ff=等待完成后打开文件
[func (c *Response) Bytes() (#左中括号##右中括号#byte, error) {]
ff=等待完成后取字节集
