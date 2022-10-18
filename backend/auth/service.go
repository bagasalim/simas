package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/bagasalim/simas/custom"
	"github.com/bagasalim/simas/model"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(data LoginRequest) (model.User, int, error)
	CreateAccount(data RegisterRequest) (model.User, int, error)
	SetOtp(Username string) (string,string, int, error)
}

type service struct {
	repo AuthRepository
}

func NewService(repo AuthRepository) *service {
	return &service{repo}
}
func (s *service) Login(data LoginRequest) (model.User, int, error) {
	username := data.Username
	User, err := s.repo.FindUser(username)

	if err != nil {
		if err.Error() == "Not found" {
			return model.User{}, http.StatusUnauthorized, errors.New("Username or Password is wrong")
		}
		return model.User{}, http.StatusInternalServerError, err
	}
	UserOtp, err := s.repo.FindOTP(User.ID)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	// fmt.Println("data",UserOtp.Expire.In(loc),"now", time.Now().In(loc), UserOtp.Expire.In(loc).After(time.Now().In(loc)))
	if err != nil || UserOtp.Code == ""{
		return model.User{}, http.StatusUnauthorized, errors.New("OTP is wrong")
	}
	if UserOtp.Code != data.Code{
		return model.User{}, http.StatusUnauthorized, errors.New("OTP is wrong")
	}
	if UserOtp.Expire.In(loc).Before(time.Now().In(loc)) || UserOtp.Used{
		return model.User{}, http.StatusUnauthorized, errors.New("OTP is expire")
	}
	if err != nil || UserOtp.Code == "" || UserOtp.Code != data.Code || UserOtp.Expire.In(loc).Before(time.Now().In(loc)){
		return model.User{}, http.StatusUnauthorized, errors.New("Token is wrong")
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, http.StatusUnauthorized, errors.New("Password is wrong")
	}
	err = s.repo.UpdateOTPExpire(UserOtp.ID)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User.Password = ""
	return User, http.StatusOK, nil
}
func (s *service) CreateAccount(data RegisterRequest) (model.User, int, error) {
	found, err := s.repo.FindUser(data.Username)
	if err == nil && found.Name != "" {
		return model.User{}, http.StatusBadRequest, errors.New("Duplicate Data")
	}
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User := model.User{
		Username: data.Username,
		Password: string(passwordHash),
		Name:     data.Name,
		Role:     2,
	}
	res, err := s.repo.AddUser(User)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}
func(s *service) SetOtp(Username string) (string, string, int, error){
	User, err := s.repo.FindUser(Username)

	if err != nil {
		if err.Error() == "Not found" {
			return  "","",http.StatusNotFound, errors.New("Username not found")
		}
		return  "","", http.StatusInternalServerError, err
	}
	data, err := s.repo.FindOTP(User.ID)
	if err != nil {
		return  "","", http.StatusInternalServerError, err
	}
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if data.Code != "" && data.Expire.Add(-2 * time.Minute).In(loc).After(time.Now().In(loc)) && data.Used == false{
		return data.Code, User.Email, http.StatusOK, nil
	}
	loc, _ = time.LoadLocation("UTC")
	userLog := model.UserOTP{
		UserID: User.ID,
		Code: custom.RandStringBytes(6),
		Expire: time.Now().Add(5 * time.Minute).In(loc),
	}
	err = s.repo.AddOTP(userLog)
	if err != nil {
		return  "","", http.StatusInternalServerError, err
	}
	return userLog.Code, User.Email, http.StatusOK, nil
}
