package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
)

type User struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}
type ServiceResponse struct {
	Code     int
	Response interface{}
}

var InMemoryUserData = map[string]User{}

func user(w http.ResponseWriter, req *http.Request) {
	var srvcRes ServiceResponse
	defer func() {
		w.WriteHeader(srvcRes.Code)

		w.Header().Add("Content-Type", "application/json")
		data, _ := json.Marshal(srvcRes.Response)
		fmt.Fprint(w, string(data))
	}()
	switch req.Method {
	case http.MethodGet:
		resp := []User{}
		i := 0
		for _, value := range InMemoryUserData {
			resp = append(resp, value)
			i++
			if i == 10 {
				break
			}
		}

		srvcRes = ServiceResponse{http.StatusOK, resp}

		return

	case http.MethodPost:
		body, _ := ioutil.ReadAll(req.Body)
		user := User{}
		json.Unmarshal(body, &user)
		fmt.Printf("in create user by as %v\n", user)
		_, ok := InMemoryUserData[user.ID]
		if user.ID != "" && ok {
			srvcRes = ServiceResponse{http.StatusBadRequest, "user already exists"}
			return
		}
		if user.ID == "" {
			user.ID = uuid.New().String()
		}
		InMemoryUserData[user.ID] = user
		srvcRes = ServiceResponse{http.StatusOK, user}
		return

	case http.MethodPut:
		id := req.URL.Query().Get("id")
		_, ok := InMemoryUserData[id]
		if !ok {
			srvcRes = ServiceResponse{http.StatusNotFound, "user not found for id " + id}
			return
		}

		body, _ := ioutil.ReadAll(req.Body)
		data := User{}
		json.Unmarshal(body, &data)
		fmt.Println(data)

		if data.Name != "" {
			if entry, ok := InMemoryUserData[id]; ok {

				// Then we modify the copy
				entry.Name = data.Name

				// Then we reassign map entry
				InMemoryUserData[id] = entry
			}
		}
		if data.Email != "" {
			if entry, ok := InMemoryUserData[id]; ok {

				// Then we modify the copy
				entry.Email = data.Email

				// Then we reassign map entry
				InMemoryUserData[id] = entry
			}
		}

		srvcRes = ServiceResponse{http.StatusOK, InMemoryUserData[id]}
		return
	}

}

func main() {
	http.HandleFunc("/user", user)
	fmt.Println("strting http server")
	http.ListenAndServe(":8081", nil)
}
