package biz

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/smtp"
)

type User struct {
	Id        int
	Email     string
	Username  string
	Password  string
	IsTeacher bool
}

type UserLogin struct {
	Email    string
	Username string
}

// UserRepo 定义用户相关的数据库操作接口
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindById(ctx context.Context, id int) (*User, error)
	GenerateVerificationCode(ctx context.Context, email string, expiryMinutes int) string
	IsExpired(ctx context.Context, email string, code string) bool
	IsCodeVerified(ctx context.Context, email string, code string) bool
}

type UserUsecase struct {
	ur  UserRepo
	log *log.Helper
}

// NewUserUsecase 创建 UserUsecase 实例
func NewUserUsecase(ur UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{ur: ur, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Login(ctx context.Context, email string, password string) (*UserLogin, error) {
	// 从仓库中根据邮箱获取用户信息
	user, err := uc.ur.FindByEmail(ctx, email)
	if err != nil {
		// 如果查询不到用户，返回错误
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	// 校验密码是否正确
	isValid := checkPassword(user.Password, password)
	if !isValid {
		return nil, ErrPassword
	}

	// 返回登录成功信息
	return &UserLogin{
		Email:    user.Email,
		Username: user.Username,
	}, nil
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	// 如果返回 nil，说明匹配成功
	return err == nil
}

func (uc *UserUsecase) Register(ctx context.Context, username string, email string, password string, confirmpassword string, isteacher bool, code string) (*User, error) {
	// 确认密码和密码不一致
	if password != confirmpassword {
		return nil, ErrConfirmPassword
	}
	// 校验用户是否已存在
	existingUser, err := uc.ur.FindByUsername(ctx, username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, ErrUserAlreadyExists
	}

	// 验证码是否超时
	timeout := uc.ur.IsExpired(ctx, email, code)
	// false 为存在问题
	if !timeout {
		return nil, ErrTimeOut
	}

	// 验证码是否正确
	codeVerified := uc.ur.IsCodeVerified(ctx, email, code)
	if !codeVerified {
		return nil, ErrCodeErrors
	}

	hpassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &User{
		Email:     email,
		Username:  username,
		Password:  hpassword,
		IsTeacher: isteacher,
	}
	if err := uc.ur.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //[]byte默认utf-8
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (uc *UserUsecase) SendEmail(ctx context.Context, email string) error {
	// 校验用户是否已存在
	existingUser, err := uc.ur.FindByEmail(ctx, email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return ErrUserAlreadyExists
	}

	code := uc.ur.GenerateVerificationCode(ctx, email, 15)

	err = sendEmail(email, code)

	if err != nil {
		return err
	}
	return nil
}

// sendEmail 发送邮件函数
func sendEmail(to string, code string) error {
	from := "2493325754@qq.com"
	password := "kfpjhmkeiykmebec" // 邮箱授权码
	smtpServer := "smtp.qq.com:465"

	// 设置 PlainAuth
	auth := smtp.PlainAuth("", from, password, "smtp.qq.com")

	// 创建 tls 配置
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         "smtp.qq.com",
	}

	// 连接到 SMTP 服务器
	conn, err := tls.Dial("tcp", smtpServer, tlsconfig)
	if err != nil {
		return fmt.Errorf("TLS 连接失败: %v", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, "smtp.qq.com")
	if err != nil {
		return fmt.Errorf("SMTP 客户端创建失败: %v", err)
	}
	defer client.Quit()

	// 使用 auth 进行认证
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("认证失败: %v", err)
	}

	// 设置发件人和收件人
	if err = client.Mail(from); err != nil {
		return fmt.Errorf("发件人设置失败: %v", err)
	}
	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("收件人设置失败: %v", err)
	}

	// 写入邮件内容
	wc, err := client.Data()
	if err != nil {
		return fmt.Errorf("数据写入失败: %v", err)
	}
	defer wc.Close()

	subject := "CCNU-EDU-LLM"
	body := `
		<h1>Verification Code</h1>
		<p>Your verification code is: <strong>` + code + `</strong></p>
		<p>This verification code is valid for 15 minutes</p>
		<p>If you are not doing it yourself, please ignore it !</p>
	`
	msg := []byte("From: Sender Name <" + from + ">\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"\r\n" +
		body)

	_, err = wc.Write(msg)
	if err != nil {
		return fmt.Errorf("消息发送失败: %v", err)
	}

	return nil
}
