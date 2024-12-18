package utils

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"

	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// 验证邮箱格式
func IsValidEmail(email string) bool {
	// 邮箱正则表达式
	var re = regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

func HashPassword(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(hash), err
}

func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte("secret"))
	return "Bearer " + signedToken, err
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ParseJWT(token string) (string, error) {
	if len(token) <= 7 || token[:7] != "Bearer " {
		return "", errors.New("jwt format errors")
	}
	token = token[7:]

	rawMessage, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte("secret"), nil
	})

	if claims, ok := rawMessage.Claims.(jwt.MapClaims); ok && rawMessage.Valid {
		username, ok := claims["username"].(string)
		if !ok {
			return "", errors.New("username claim is not a string")
		}
		return username, nil
	} 
 
	return "", err
}

// 生成指定长度的随机数字字符串
func generateRandomNumberString(length int) string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	digits := make([]byte, length)
	for i := 0; i < length; i++ {
		digits[i] = byte(r.Intn(10)) + '0'
	}
	return string(digits)
}

// 计算 Luhn 校验位
func calculateLuhnCheckDigit(number string) (int, error) {
	sum := 0
	double := false
	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return 0, err
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return (10 - (sum % 10)) % 10, nil
}

// 生成符合 Luhn 校验规则的银行账号
func GenerateBankAccountNumber(length int) (string, error) {
	if length <= 1 {
		return "", fmt.Errorf("长度必须大于1")
	}
	accountNumber := generateRandomNumberString(length - 1)
	checkDigit, err := calculateLuhnCheckDigit(accountNumber)
	if err != nil {
		return "", err
	}
	return accountNumber + strconv.Itoa(checkDigit), nil
}
