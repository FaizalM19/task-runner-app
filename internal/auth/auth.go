package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"task-runner/models"
	"task-runner/pkg/utils"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

var UserStore = make(map[string]string)
var mu sync.Mutex

// ValidateToken validates the Bearer token using HMAC
func ValidateToken(token string) bool {
	if len(token) < 7 || token[:7] != "Bearer " {
		return false
	}
	token = token[7:]

	mu.Lock()
	defer mu.Unlock()

	for _, storedToken := range UserStore {
		if storedToken == token {
			return true
		}
	}
	return false
}

func GenerateToken(username string) string {
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(username))

	token := base64.StdEncoding.EncodeToString(h.Sum(nil))

	mu.Lock()
	UserStore[username] = token
	mu.Unlock()

	return token
}

type AuthService struct{}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (authService *AuthService) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if loginReq.Username == "" || loginReq.Password == "" {
		utils.ErrorResponse(w, http.StatusBadRequest, "Username and Password are required")
		return
	}

	token := GenerateToken(loginReq.Username)

	utils.JSONResponse(w, http.StatusOK, map[string]string{"Token": token})
}
