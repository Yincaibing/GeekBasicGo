package web

import (
	"github.com/Yincaibing/GeekBasicGo/gin/webook/inter/domain"
	"github.com/Yincaibing/GeekBasicGo/gin/webook/inter/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserHandler 我准备在它上面定义跟用户有关的路由
type UserHandler struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

// NewUserHandler 为什么要有这个函数NewUserHandler？因为要面向对象编程，如果后面要使用一个结构体方法（比如基于UserHandler的登录注册方法），必须先定义一个创建结构体的函数（方法有接受体，即结构体对象，但是函数不需要），类似于 java里面的结构体方法，先构造对象
func NewUserHandler(svc *service.UserService) *UserHandler {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandler{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (u *UserHandler) RegisterRoutesV1(ug *gin.RouterGroup) {
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) Profile(context *gin.Context) {

}

func (u *UserHandler) SignUp(ctx *gin.Context) {
	//1、定义入参结构体
	type SignUpReq struct {
		Email           string `json:"email"`
		ConfirmPassword string `json:"confirmPassword"`
		Password        string `json:"password"`
	}

	//2、声明入参结构体
	var req SignUpReq

	//3、bind解析入参
	// Bind函数会根据请求的方法和Content-Type来自动选择绑定引擎。根据"Content-Type"头部的不同，会使用不同的绑定方式，例如：
	//
	//     "application/json" --> JSON绑定
	//     "application/xml"  --> XML绑定
	//
	// 如果Content-Type等于"application/json"，则会将请求体解析为JSON，并使用JSON或XML作为JSON输入。
	// 它会将JSON负载解码为指定为指针的结构体。
	// 如果输入无效，它会在响应中写入400错误，并设置Content-Type头部为"text/plain"。
	if err := ctx.Bind(&req); err != nil {
		return
	}

	//4、接下来校验入参
	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}
	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不一致")
		return
	}
	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		// 记录日志
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须大于8位，包含数字、特殊字符")
		return
	}

	// 5、调用svc 的方法，u为Userhandler指针
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	//6、存的过程如果有错误，看是否是从底层dao传上来的错误
	//错误传导，使用别名的机制，层层传导，让我们在 Handler 里面依旧保持只依赖 service，避免了跨层依赖的问题。
	if err == service.ErrUserDuplicateEmail {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}

	ctx.String(http.StatusOK, "注册成功")
}

func (u *UserHandler) Login(context *gin.Context) {
	//1、先定义一个请求结构体
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//2、声明一个请求体变量
	var req LoginReq

	// 3、Bind 方法会根据 Content-Type 来解析你的数据到 req 里面
	// 解析错了，就会直接写回一个 400 的错误
	if err := context.Bind(&req); err != nil {
		return
	}

	//4、调用 service的Login
	user, err := u.svc.Login(context, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		context.String(http.StatusOK, "用户名或密码不对")
		return
	}
	if err != nil {
		context.String(http.StatusOK, "系统错误")
		return
	}

	//5、 在这里登录成功了，登录态需要保持
	// 设置 session
	sess := sessions.Default(context)
	// 我可以随便设置值了
	// 你要放在 session 里面的值
	sess.Set("userId", user.Id)
	sess.Save()
	context.String(http.StatusOK, "登录成功")
	return

}

/*
你需要完善 /users/edit 对应的接口。要求：

允许用户补充基本个人信息，包括：
昵称：字符串，你需要考虑允许的长度。
生日：前端输入为 1992-01-01 这种字符串。
个人简介：一段文本，你需要考虑允许的长度。
尝试校验这些输入，并且返回准确的信息。
修改 /users/profile 接口，确保这些信息也能输出到前端。
不要求你开发前端页面。提交作业的时候，顺便提交 postman 响应截图。加一个 README 文件，里面贴个图。

就是补充 live 分支上的 Edit 和 Profile 接口。

PS：暂时不要求上传头像，后面我们讲到 OSS 之后直接用 OSS。
*/
func (u *UserHandler) Edit(context *gin.Context) {

}
