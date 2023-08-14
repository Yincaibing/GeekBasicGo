package domain

import "time"

// User 领域对象，是 DDD 中的 entity
// BO(business object)
type User struct {
	Id       int64
	Email    string
	Password string
	Ctime    time.Time
}

//type Address struct {
//}

// 昵称：字符串，你需要考虑允许的长度。
// 生日：前端输入为 1992-01-01 这种字符串。
// 个人简介：一段文本，你需要考虑允许的长度。
type Profile struct {
	Id         int64
	NickName   string
	BirthDate  string
	Ctime      time.Time
	PersonDesc string
}
