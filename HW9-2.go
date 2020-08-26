package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"strconv"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type user struct {
	name string
	age int
	password string
}

var users=make(map[string]user)

var mySigningKey = []byte("my_secret_key")

func GenerateJWT(u user) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userlogin"] = u.name
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func registration(w http.ResponseWriter, r *http.Request){
	if(r.Method=="GET") {
		http.ServeFile(w, r, "D:\\Go\\HW9\\reg.html")
	} else if(r.Method=="POST"){
		if err := r.ParseForm(); err != nil {
			fmt.Println(err.Error())
			return
		}
		login := r.Form.Get("userlogin")
		name := r.Form.Get("username")
		age,_ := strconv.Atoi(r.Form.Get("userage"))
		password := r.Form.Get("userpassword")
		hash, _ := hashPassword(password)
		users[login]=user{
			name:     name,
			age:      age,
			password: hash,
		}
		fmt.Fprintf(w, "Login: %s Name: %s Age: %d Password: %s", login, name, age, hash)
	}
}

func checkUser(login,password string)  bool{
	for k, v := range users {
		if(k==login&&checkPasswordHash(password,v.password)) {
			return true
		}
	}
	return false
}

func findUser(login string)  user{
	for k, v := range users {
		if(k==login) {
			return v
		}
	}
	return user{}
}

func login(w http.ResponseWriter, r *http.Request){
	if(r.Method=="GET") {
		http.ServeFile(w, r, "D:\\Go\\HW9\\login.html")
	} else if(r.Method=="POST"){
		if err := r.ParseForm(); err != nil {
			fmt.Println(err.Error())
			return
		}
		if(r.Header.Get("content-type")=="application/x-www-form-urlencoded") {
			login := r.Form.Get("userlogin")
			password := r.Form.Get("userpassword")
			if (checkUser(login, password)) {
				fmt.Fprintf(w, "Login: %s successful", login)
				curentUser:=findUser(login)
				validToken, err := GenerateJWT(curentUser)
				if(err==nil) {
					fmt.Fprintf(w,"Token: %s", validToken)

				}


			} else {
				fmt.Fprintf(w, "Login or password is incorrect")
			}
		} else if(r.Header.Get("content-type")=="application/json"){
			mymap := make(map[string]string)
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			err = json.Unmarshal(body, &mymap)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			login:=mymap["login"]
			password:=mymap["password"]
			if (checkUser(login, password)) {
				fmt.Fprintf(w, "Login: %s successful", login)
			} else {
				fmt.Fprintf(w, "Login or password is incorrect")
			}
		}
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "D:\\Go\\HW9\\main.html")
	})
	http.HandleFunc("/reg",registration)
	http.HandleFunc("/log",login)
	fmt.Println("Server is listening...")
	err:=http.ListenAndServe(":8080", nil)
	if err!=nil{
		fmt.Println(err.Error())
	}
}