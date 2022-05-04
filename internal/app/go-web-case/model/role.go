package model

// Role omitempty 表示如果在序列化的时候、如果不存在则不进行默认值填充。结果就不包含这个字段
//参考 https://old-panda.com/2019/12/11/golang-omitempty/
//类型如果是指针、默认值是nil；如果不是就按照各类型的默认值填入表中
type Role struct {
	//gorm的嵌入方式
	Model
	RoleName string `json:"role_name"`
}
