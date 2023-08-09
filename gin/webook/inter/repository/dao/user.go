package dao

//这一层是代表数据库操作，外层的 reopsitory.user是 mysql,mongodb，ES等的抽象

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

// User 0、定义User表， 直接对应数据库表结构
// 有些人叫做 entity，有些人叫做 model，有些人叫做 PO(persistent object)
type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 全部用户唯一
	Email    string `gorm:"unique"`
	Password string

	// 往这面加

	// 创建时间，毫秒数
	Ctime int64
	// 更新时间，毫秒数
	Utime int64
}

// UserDAO 1、定义一个迁移表的结构体，这样连接好数据库后，连接方法的返回值就有可接收的实体对象
type UserDAO struct {
	db *gorm.DB
}

// NewUserDAO 2、新建一个数据库表对应的地址，方便做增删改查等操作。
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

// FindByEmail 作为UserDao的方法，传入请求上下文，和要查找的 email,便可以得到 user或者 err
func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("email = ?", email).First(&u).Error
	//err := dao.db.WithContext(ctx).First(&u, "email = ?", email).Error
	return u, err
}

// Insert 作为UserDao的方法，传入请求上下文，和要插入的 user,便可以得到 user或者 err
func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	// 存毫秒数
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 邮箱冲突
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) DeleteByUserID(ctx context.Context, userid string) error {
	err := dao.db.WithContext(ctx).Delete(userid).Error
	return err
}
