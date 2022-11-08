package handler

import (
	api "com/josh/asset/api"
	"com/josh/asset/db"
	"com/josh/asset/service"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var whiteListTokens = make([]string, 5, 5)

func Login(w http.ResponseWriter, r *http.Request) {
	user := &service.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	fmt.Println(*user)
	resp := FindOne(user.Username, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOne(email, password string) map[string]interface{} {
	user := &service.User{}
	fmt.Println(email, password)
	db := db.GetDB()

	if err := db.Find(user, "username = ?", email).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}
	fmt.Println("after where", *&user.Username)

	expiresAt := time.Now().Add(time.Minute * 10).Unix()

	// errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.Password != password { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &service.Token{
		UserID: user.EmpId,
		Name:   user.Username,
		// Email:  user.Email,
		Role: user.Role,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}
	//add into whitelist
	whiteListTokens = append(whiteListTokens, tokenString)

	var resp = map[string]interface{}{}
	resp["token"] = tokenString //Store the token in the response
	// resp["user"] = user
	return resp
}

// JwtVerify Middleware function
func JwtVerify(next http.Handler, role ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var header = r.Header.Get("Authorization") //Grab the token from the header
		fmt.Println(r.Header)
		header = strings.TrimSpace(header)

		if header == "" {
			//Token is missing, returns with error code 403 Unauthorized
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Missing auth token")
			return
		}
		if !contains(whiteListTokens, header) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("Invalid token")
			return
		}

		tk := &service.Token{}
		_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})

		fmt.Println("tk=>>>", *tk)

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		if !contains(role, tk.Role) {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode("You don't has enough privileges to perform an action ")
			return
		}
		ctx := context.WithValue(r.Context(), "user", tk)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func contains(list []string, ele string) bool {
	for _, role := range list {
		if role == ele {
			return true
		}
	}
	return false
}

func Logout(w http.ResponseWriter, r *http.Request) {

	var header = r.Header.Get("Authorization") //Grab the token from the header
	fmt.Println(r.Header)
	header = strings.TrimSpace(header)

	if header == "" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode("Missing auth token")
		return
	}
	tk := &service.Token{}
	_, err := jwt.ParseWithClaims(header, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	fmt.Println("tk=>>>", *tk)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	for i, v := range whiteListTokens {
		if v == header {
			whiteListTokens = append(whiteListTokens[:i], whiteListTokens[i+1:]...)
		}
	}

	if !contains(whiteListTokens, header) {
		api.Response(http.StatusOK, "Logout successfully", w)

	}

}

///

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func Generate(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func CreateUser(user service.User) error {
	db := db.GetDB()
	err := db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
