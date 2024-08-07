package handlers

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"strconv"

	"github.com/javier-elizaga/go-api/handlers/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := getUsers()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	json.NewEncoder(w).Encode(users)
}

func NearbyUser(w http.ResponseWriter, r *http.Request) {
	lat, err := utils.GetFloat64Param(r, "lat")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}
	lng, err := utils.GetFloat64Param(r, "lng")
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	users, err := getUsers()
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	if len(users) == 0 {
		utils.Error(w, http.StatusNotFound, "Users not found")
		return
	}

	user := users[0]
	minDistance, _ := user.distance(lat, lng)

	for i := 1; i < len(users); i++ {
		distance, _ := users[i].distance(lat, lng)
		if distance < minDistance {
			user = users[i]
			minDistance = distance
		}
	}

	json.NewEncoder(w).Encode(user)
}

func getUsers() ([]User, error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, errors.New("call users api")
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	var data []User
	err = decoder.Decode(&data)
	if err != nil {
		return nil, errors.New("decode users response")
	}
	return data, nil
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  struct {
		Street  string `json:"street"`
		Suite   string `json:"suite"`
		City    string `json:"city"`
		Zipcode string `json:"zipcode"`
		Geo     struct {
			Lat string `json:"lat"`
			Lng string `json:"lng"`
		} `json:"geo"`
	} `json:"address"`
	Phone   string `json:"phone"`
	Website string `json:"website"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func (u User) location() (lat, lng float64, err error) {
	lat, err = strconv.ParseFloat(u.Address.Geo.Lat, 64)
	if err != nil {
		return 0, 0, err
	}
	lng, err = strconv.ParseFloat(u.Address.Geo.Lng, 64)
	if err != nil {
		return 0, 0, err
	}
	return lat, lng, nil
}

func (u User) distance(lat, lng float64) (distance float64, err error) {
	lat1, lng1, err := u.location()
	if err != nil {
		return 0, err
	}
	// euclidean formula
	return math.Sqrt(math.Pow(lat-lat1, 2) + math.Pow(lng-lng1, 2)), nil
}
