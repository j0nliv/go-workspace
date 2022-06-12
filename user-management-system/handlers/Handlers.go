package handlers

import (
	"encoding/json"
	"net/http"

	. "sametemiroglu.com/dataloaders"
	. "sametemiroglu.com/models"
)

func Run() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// PAGE OBJECT
	page := Page{ID: 7, Name: "Users", Description: "User Lists", URI: "/users"}

	// DATA LOADERS
	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadInterestMappings()

	// PROCESS
	var newUsers []User
	for _, user := range users {
		for _, inteinterestMapping := range interestMappings {
			if user.ID == inteinterestMapping.UserID {
				for _, interest := range interests {
					if inteinterestMapping.InterestID == interest.ID {
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers, user)
		//fmt.Println(user.Firstname)
	}
	viewModel := UserViewModel{Page: page, Users: newUsers}
	data, _ := json.Marshal(viewModel)
	w.Write([]byte(data))
}
