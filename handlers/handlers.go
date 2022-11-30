package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ElPoderosoLukita/goCRUDbd/models"
	"github.com/gorilla/mux"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

// POST HANDLER
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	body := r.Body
	user := models.CreateUser()

	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(jsonErr))
	}

	err = models.InsertUser(user)
	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusInternalServerError, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(jsonErr))
	}

	jsonResp, _ := json.Marshal(models.NewResponse(http.StatusCreated, "The user was created correctly."))
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, string(jsonResp))
}

// DELETE HANDLER
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)

	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(jsonErr))
	} else {
		models.DeleteUser(idInt)
		jsonResp, _ := json.Marshal(models.NewResponse(http.StatusCreated, "The user was deleted correctly."))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonResp))
	}
}

// UPDATE HANDLER
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)
	user := models.CreateUser()

	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(jsonErr))
	}

	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, jsonErr)
	} else {
		models.UpdateUser(idInt, user)
		jsonResp, _ := json.Marshal(models.NewResponse(http.StatusCreated, "The user was updated correctly."))
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonResp))
	}
}

// GET USER HANDLER
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]
	idInt, err := strconv.Atoi(id)

	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(jsonErr))
	}

	user, err := models.GetUser(idInt)

	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, string(jsonErr))
	} else {

		jsonUser, _ := json.Marshal(user)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonUser))
	}
}

// GET USERS HANDLER
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := models.GetUsers()
	jsonUsers, err := json.Marshal(users)

	if err != nil {
		jsonErr, _ := json.Marshal(models.NewErrorStruct(http.StatusBadRequest, err.Error()))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, jsonErr)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, string(jsonUsers))
	}
}
