package f

// 小写（未公开类型） 大写（公开类型）
type alertCounter int

// go 的工厂方法的约定写法 new 关键字，返回未公开实例
func New(value int) alertCounter {
	return alertCounter(value)
}

// 公开结构 User
type User struct {
	Name  string // 公开字段
	email string // 未公开字段
}

// 未公开结构 user
type user struct {
	Name  string // 结构体未公开 字段公开
	Email string // 结构体未公开 字段公开
}

// 公开类型 Admin
type Admin struct {
	user // 嵌入未公开类型 user
	Rights int
}
