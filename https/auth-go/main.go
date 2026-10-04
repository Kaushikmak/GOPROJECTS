package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/cors"
)

const (
	accessTokenSecret  = "access-secret"
	refreshTokenSecret = "refresh-secret"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{{
	ID:       1,
	Username: "kaushik",
	Password: "1234567890",
}}

var refreshTokens = map[string]bool{}

type AccessClaims struct {
	UserID   int    `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

func createAccessToken(user User) (string, error) {

	claims := AccessClaims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(15 * time.Minute),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(accessTokenSecret),
	)
}

func createRefreshToken(user User) (string, error) {

	claims := RefreshClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(7 * 24 * time.Hour),
			),
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(refreshTokenSecret),
	)
}

func signup(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if body.Username == "" || body.Password == "" {
		http.Error(
			w,
			"Username and password are required",
			http.StatusBadRequest,
		)
		return
	}

	// Check existing user

	for _, user := range users {

		if user.Username == body.Username {

			http.Error(
				w,
				"User already exists",
				http.StatusConflict,
			)

			return
		}
	}

	newUser := User{
		ID:       len(users) + 1,
		Username: body.Username,
		Password: body.Password,
	}

	users = append(users, newUser)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created",
		"user": map[string]interface{}{
			"id":       newUser.ID,
			"username": newUser.Username,
		},
	})
}

func signin(w http.ResponseWriter, r *http.Request) {

	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var user User

	found := false

	for _, u := range users {

		if u.Username == body.Username &&
			u.Password == body.Password {

			user = u
			found = true
			break
		}
	}

	if !found {

		http.Error(
			w,
			"Invalid username or password",
			http.StatusUnauthorized,
		)

		return
	}

	// Create access token

	accessToken, err := createAccessToken(user)

	if err != nil {
		http.Error(
			w,
			"Could not create access token",
			http.StatusInternalServerError,
		)

		return
	}

	// Create refresh token

	refreshToken, err := createRefreshToken(user)

	if err != nil {
		http.Error(
			w,
			"Could not create refresh token",
			http.StatusInternalServerError,
		)

		return
	}

	// Store refresh token

	refreshTokens[refreshToken] = true

	// Put refresh token in HttpOnly cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   false, // true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		Path:     "/",
		MaxAge:   7 * 24 * 60 * 60,
	})

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message":     "Signin successful",
		"accessToken": accessToken,
	})
}

func authMiddleware(
	next http.Handler,
) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {

				http.Error(
					w,
					"No access token provided",
					http.StatusUnauthorized,
				)

				return
			}

			parts := strings.Split(
				authHeader,
				" ",
			)

			if len(parts) != 2 ||
				parts[0] != "Bearer" {

				http.Error(
					w,
					"Invalid authorization header",
					http.StatusUnauthorized,
				)

				return
			}

			tokenString := parts[1]

			// Parse JWT

			token, err := jwt.ParseWithClaims(
				tokenString,
				&AccessClaims{},
				func(token *jwt.Token) (interface{}, error) {

					return []byte(accessTokenSecret), nil
				},
			)

			if err != nil || !token.Valid {

				http.Error(
					w,
					"Invalid or expired access token",
					http.StatusUnauthorized,
				)

				return
			}

			claims, ok := token.Claims.(*AccessClaims)

			if !ok {

				http.Error(
					w,
					"Invalid token claims",
					http.StatusUnauthorized,
				)

				return
			}

			// Add user information to request context

			ctx := r.Context()

			ctx = contextWithUser(
				ctx,
				claims,
			)

			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		},
	)
}

func profile(w http.ResponseWriter, r *http.Request) {

	claims := getUserFromContext(r.Context())

	if claims == nil {

		http.Error(
			w,
			"User not found",
			http.StatusUnauthorized,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "You are authenticated",
		"user": map[string]interface{}{
			"userId":   claims.UserID,
			"username": claims.Username,
		},
	})
}

func refresh(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("refreshToken")

	if err != nil {

		http.Error(
			w,
			"Refresh token required",
			http.StatusUnauthorized,
		)

		return
	}

	refreshToken := cookie.Value

	// Check our stored refresh tokens

	if !refreshTokens[refreshToken] {

		http.Error(
			w,
			"Invalid refresh token",
			http.StatusForbidden,
		)

		return
	}

	// Verify JWT

	token, err := jwt.ParseWithClaims(
		refreshToken,
		&RefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {

			return []byte(refreshTokenSecret), nil
		},
	)

	if err != nil || !token.Valid {

		http.Error(
			w,
			"Invalid or expired refresh token",
			http.StatusForbidden,
		)

		return
	}

	claims, ok := token.Claims.(*RefreshClaims)

	if !ok {

		http.Error(
			w,
			"Invalid refresh claims",
			http.StatusForbidden,
		)

		return
	}

	// Find user

	var user User

	for _, u := range users {

		if u.ID == claims.UserID {

			user = u
			break
		}
	}

	// Create new access token

	accessToken, err := createAccessToken(user)

	if err != nil {

		http.Error(
			w,
			"Could not create access token",
			http.StatusInternalServerError,
		)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"accessToken": accessToken,
	})
}

type contextKey string

const userKey contextKey = "user"

func contextWithUser(
	ctx context.Context,
	user *AccessClaims,
) context.Context {

	return context.WithValue(
		ctx,
		userKey,
		user,
	)
}

func getUserFromContext(
	ctx context.Context,
) *AccessClaims {

	user, ok := ctx.Value(userKey).(*AccessClaims)

	if !ok {
		return nil
	}

	return user
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/signup",
		signup,
	)

	mux.HandleFunc(
		"/signin",
		signin,
	)

	mux.HandleFunc(
		"/refresh",
		refresh,
	)

	// Protected route

	mux.Handle(
		"/profile",
		authMiddleware(
			http.HandlerFunc(profile),
		),
	)

	// CORS

	handler := cors.New(
		cors.Options{
			AllowedOrigins: []string{
				"http://localhost:5173",
			},
			AllowCredentials: true,
		},
	).Handler(mux)

	fmt.Println(
		"Server running on http://localhost:3000",
	)

	http.ListenAndServe(
		":3000",
		handler,
	)
}
