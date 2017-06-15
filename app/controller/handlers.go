package controller

import (

	"net/http"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/Fezerik/app/config"
)


func Index(w http.ResponseWriter, r *http.Request) {
	// fs := http.FileServer(http.Dir("static"))
	//Call to ParseForm makes form fields available.
	err := r.ParseForm()
	if err != nil {
		// Handle error here via logging and then return
	}
	config.TPL.ExecuteTemplate(w, "index.gohtml", nil)

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		fmt.Println("method:", r.Method)
		config.TPL.ExecuteTemplate(w, "signup.gohtml", nil)
		//config.TPL.ParseFiles("signup.gohtml")
		//t.config.TPL.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in

		password := r.PostFormValue("password")
		email := r.PostFormValue("email")

		fmt.Println(password)
		fmt.Println(email)

		config := &aws.Config{
			Region: aws.String("ap-northeast-1"),

		}

		sess := session.Must(session.NewSession(config))

		svc := cognitoidentityprovider.New(sess)

		params := &cognitoidentityprovider.SignUpInput{
			ClientId:   aws.String("1b55tttohaokjfj2dn1q9bn63b"), // Required
			Password:   aws.String(password), // Required
			Username:   aws.String(email), // Required


			UserAttributes: []*cognitoidentityprovider.AttributeType{
				{ // Required
					Name:  aws.String("email"), // Require
					Value: aws.String(email),

				},
				// More values...

			},

		}
		resp, err := svc.SignUp(params)
			//need to add logic to display err on page
		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			return
		}

		// Pretty-print the response data.
		fmt.Println(resp)

		http.Redirect(w, r, "/activate", http.StatusSeeOther)
	}

}

func Activate(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		fmt.Println("method:", r.Method)
		config.TPL.ExecuteTemplate(w, "activate.gohtml", nil)
		//config.TPL.ParseFiles("signup.gohtml")
		//t.config.TPL.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in

		code := r.PostFormValue("code")
		email := r.PostFormValue("email")

		fmt.Println(code)
		fmt.Println(email)

		config := &aws.Config{
			Region: aws.String("ap-northeast-1"),

		}

		sess := session.Must(session.NewSession(config))

		svc := cognitoidentityprovider.New(sess)

		params := &cognitoidentityprovider.GetUserAttributeVerificationCodeInput{
			AttributeName: aws.String("email"), // Required
			//AccessToken:   aws.String("code"),
		}
		resp, err := svc.GetUserAttributeVerificationCode(params)

		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			return
		}

		// Pretty-print the response data.
		fmt.Println(resp)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}





}

func Login(w http.ResponseWriter, r *http.Request) {
	//Call to ParseForm makes form fields available.
	//err := r.ParseForm()
	//if err != nil {
	//	// Handle error here via logging and then return
	//}
	config.TPL.ExecuteTemplate(w, "login.gohtml", nil)

}

