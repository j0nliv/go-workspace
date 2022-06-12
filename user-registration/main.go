package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	username, email, passwd, rpasswd := "", "", "", ""

	// REGISTER

	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username = r.FormValue("username")
		email = r.FormValue("email")
		passwd = r.FormValue("passwd")
		rpasswd = r.FormValue("rpasswd")

		usernameCheck := isEmpty(username)
		emailCheck := isEmpty(email)
		passwd := isEmpty(passwd)
		rpasswd := isEmpty(rpasswd)

		if usernameCheck || emailCheck || passwd || rpasswd {
			fmt.Fprintf(w, "There is empty form data.")
			return
		}

		if passwd == rpasswd {
			fmt.Fprintf(w, "Registration is correct.")
		} else {
			fmt.Fprintf(w, "Password and rpassword must be similar. Please try again!")
		}

	})

	// LOGIN

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		username = r.FormValue("username")
		passwd = r.FormValue("passwd")

		usernameCheck := isEmpty(username)
		passwdCheck := isEmpty(passwd)

		if usernameCheck || passwdCheck {
			fmt.Fprintf(w, "There is empty form data.")
			return
		}

		credsusername := "sametemiroglu"
		credspasswd := "Passwd123!"

		if username == credsusername && passwd == credspasswd {
			fmt.Fprintln(w, "Login succesful!")
		} else {
			fmt.Fprintln(w, "Login error!")
		}

	})

	http.ListenAndServe(":8000", mux)

}

func isEmpty(data string) bool {
	if len(data) == 0 {
		return true
	} else {
		return false
	}
}
