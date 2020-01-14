package httphandler

import (
	"WeatherMonster/pkg/io"
	"WeatherMonster/pkg/service"
	"WeatherMonster/pkg/utile"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type WeatherHandler interface {
	ForecastHandler(w http.ResponseWriter, r *http.Request)
	CreateCityHandler(w http.ResponseWriter, r *http.Request)
	CreateTemperatureHandler(w http.ResponseWriter, r *http.Request)
	UpdateCityHandler(w http.ResponseWriter, r *http.Request)
	DeleteCityHandler(w http.ResponseWriter, r *http.Request)
	CreateWebHookHandler(w http.ResponseWriter, r *http.Request)
	DeleteWebHookHandler(w http.ResponseWriter, r *http.Request)
}
type weather struct {
	weather service.WeatherService
}

//NewCouchRepository create new repository
func NewWatherHandler(ub service.WeatherService) WeatherHandler {
	return &weather{
		weather: ub,
	}
}

func (w *weather) CreateCityHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	rw.Header().Set("Content-Type", "application/json")
	log.Print(r.Body)
	ctx := r.Context()
	var request io.City
	if r.Header.Get("Content-Type") == "application/json" {

		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(err.Error())))
			return
		}

	} else if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		r.ParseForm()
		request.Name = r.FormValue("name")
		request.Latitude, _ = strconv.ParseFloat(r.FormValue("latitude"), 64)
		request.Longitude, _ = strconv.ParseFloat(r.FormValue("longitude"), 64)
	}
	if request.Longitude <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("longitude required")))
		return
	}
	if request.Latitude <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("latitude required")))
		return
	}
	if len(request.Name) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("name required")))
		return
	}
	resp := w.weather.CityCreate(ctx, request)
	rw.Write(utile.ReturnResponseWithJson(resp))
}
func (w *weather) CreateTemperatureHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	log.Print(r.Body)
	ctx := r.Context()
	var request io.Temperatures
	if r.Header.Get("Content-Type") == "application/json" {

		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(err.Error())))
			return

		}
	} else if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		r.ParseForm()
		request.CityID, _ = strconv.Atoi(r.FormValue("city_id"))
		request.Min, _ = strconv.ParseFloat(r.FormValue("min"), 64)
		request.Max, _ = strconv.ParseFloat(r.FormValue("max"), 64)
	}
	if request.CityID <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("city id required")))
		return
	}
	if request.Max <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("max required")))
		return
	}
	if request.Min <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("min required")))
		return
	}
	resp := w.weather.TemperatureCreate(ctx, request)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(utile.ReturnResponseWithJson(resp))
}
func (w *weather) UpdateCityHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := mux.Vars(r)
	rw.Header().Set("Content-Type", "application/json")
	if len(params["id"]) < 1 {
		log.Println("Url Param 'key' is missing")
		keyErr := "Url Param 'key' is missing"
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(keyErr)))
		return
	}
	log.Print(r.Body)
	ctx := r.Context()
	var request io.City
	if r.Header.Get("Content-Type") == "application/json" {

		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(err.Error())))
			return
		}
	} else if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		r.ParseForm()
		request.Name = r.FormValue("name")
		request.Latitude, _ = strconv.ParseFloat(r.FormValue("latitude"), 64)
		request.Longitude, _ = strconv.ParseFloat(r.FormValue("longitude"), 64)
	}
	if request.Longitude <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("longitude required")))
		return
	}
	if request.Latitude <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("latitude required")))
		return
	}
	if len(request.Name) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("name required")))
		return
	}
	request.ID, _ = strconv.Atoi(params["id"])
	resp := w.weather.CityUpdate(ctx, request)
	rw.Write(utile.ReturnResponseWithJson(resp))
}

func (w *weather) DeleteCityHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	params := mux.Vars(r)
	rw.Header().Set("Content-Type", "application/json")
	if len(params["id"]) < 1 {
		log.Println("Url Param 'key' is missing")
		keyErr := "Url Param 'key' is missing"
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(keyErr)))
		return
	}
	log.Print(r.Body)
	ctx := r.Context()
	var request io.City
	err := decoder.Decode(&request)
	if err != nil {
		fmt.Println(err)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(err.Error())))
		return
	}
	request.ID, _ = strconv.Atoi(params["id"])
	resp := w.weather.CityDelete(ctx, request)
	rw.Write(utile.ReturnResponseWithJson(resp))
}
func (w *weather) ForecastHandler(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Print("URL",r.URL.Hostname())
	log.Print("HOST",r.URL.Path)
	rw.Header().Set("Content-Type", "application/json")
	if len(params["city_id"]) < 1 {
		log.Println("Url Param 'key' is missing")
		keyErr := "Url Param 'key' is missing"
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(keyErr)))
		return
	}
	ctx := r.Context()
	var request io.Forecast
	request.CityID, _ = strconv.Atoi(params["city_id"])
	resp := w.weather.Forecast(ctx, request)
	rw.Write(utile.ReturnResponseWithJson(resp))
}
func (w *weather) CreateWebHookHandler(rw http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	rw.Header().Set("Content-Type", "application/json")
	var request io.WebHook
	if r.Header.Get("Content-Type") == "application/json" {

		err := decoder.Decode(&request)
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(err.Error())))
			return
		}

	} else if r.Header.Get("Content-Type") == "application/x-www-form-urlencoded" {
		r.ParseForm()
		request.CityID, _ = strconv.Atoi(r.FormValue("city_id"))
		request.CallbackUrl = r.FormValue("callback_url")
	}
	if request.CityID <= 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("city id required")))
		return
	}
	if len(request.CallbackUrl) == 0 {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage("call back url required")))
		return
	}
	ctx := r.Context()
	resp := w.weather.WebHookCreate(ctx, request)
	rw.Write(utile.ReturnResponseWithJson(resp))
}
func (w *weather) DeleteWebHookHandler(rw http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	rw.Header().Set("Content-Type", "application/json")
	if len(params["id"]) < 1 {
		keyErr := "Url Param 'key' is missing"
		log.Println(keyErr)
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write(utile.ReturnResponseWithJson(io.FailureMessage(keyErr)))
		return
	}
	log.Print(r.Body)
	ctx := r.Context()
	var request io.WebHook
	request.ID, _ = strconv.Atoi(params["id"])
	resp := w.weather.WebHookDelete(ctx, request)
	rw.Write(utile.ReturnResponseWithJson(resp))
}
