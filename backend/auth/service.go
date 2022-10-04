package auth

import (
	"net/http"

	"github.com/bagasalim/simas/model"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(data LoginRequest) (model.User, int, error)
	CreateAccount(data RegisterRequest) (model.User, int, error)
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
		if err.Error() == "Username or Password is wrong" {
			return model.User{}, http.StatusUnauthorized, err
		}
		return model.User{}, http.StatusInternalServerError, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(data.Password))
	if err != nil {
		return model.User{}, http.StatusUnprocessableEntity, err
	}
	User.Password = ""
	return User, http.StatusOK, nil
}
func (s *service) CreateAccount(data RegisterRequest) (model.User, int, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	User := model.User{
		Username: data.Username,
		Password: string(passwordHash),
		Name:     data.Name,
	}
	res, err := s.repo.addUser(User)
	if err != nil {
		return model.User{}, http.StatusInternalServerError, err
	}
	return res, http.StatusOK, nil
}

// func generateJWT(username string, name string) (string, error) {
// 	claim := custom.DataJWT{
// 		Username: username,
// 		Name:     name,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
// 	tokenString, err := token.SignedString([]byte("kampang"))
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }
// func claimToken(tokenString string) (custom.DataJWT, error) {
// 	claims := custom.DataJWT{}
// 	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
// 		// since we only use the one private key to sign the tokens,
// 		// we also only use its public counter part to verify
// 		return []byte("kampang"), nil
// 	})
// 	if err != nil {
// 		return DataJWT{}, err
// 	}
// 	if !token.Valid {
// 		return DataJWT{}, nil
// 	}
// 	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 5*time.Minute {
// 		return DataJWT{}, nil
// 	}
// 	return claims, nil
// }
