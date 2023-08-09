package repository

//为什么有 repository 之后，还要有 dao？repository 是一个整体抽象，它里面既可以考虑用 ElasticSearch，也可以考虑使用 MySQL，还可以考虑用 MongoDB。所以它只代表数据存储，但是不 代表数据库。
//• service 是拿来干嘛的？简单来说，就是组合各种 repository、domain，偶尔也会组合别的 service， 来共同完成一个业务功能。
//• domain 又是什么？它被认为是业务在系统中的直接反应，或者你直接理解为一个业务对象，又或者就 是一个现实对象在程序中的反应
import (
	"GeekBasicGo/gin/webook/internal/domain"
	"GeekBasicGo/gin/webook/internal/repository/dao"
	"context"
	"errors"
	"gorm.io/gorm"
)

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	// SELECT * FROM `users` WHERE `email`=?
	u, err := r.dao.FindByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
}

func (r *UserRepository) FindById(int64) {
	// 先从 cache 里面找
	// 再从 dao 里面找
	// 找到了回写 cache
}
