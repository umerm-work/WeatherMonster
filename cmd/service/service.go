package service

import (
	db2 "WeatherMonster/pkg/db"
	"WeatherMonster/pkg/http"
	"WeatherMonster/pkg/service"
	"WeatherMonster/pkg/utile"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func Run() {

	inConfig()
	db := db2.New(db2.Connect())
	weatherService := service.NewBasiWeatherService(db)
	handler := httphandler.NewWatherHandler(weatherService)
	intHttp(handler)
}

func inConfig() {
	utile.GetConf()
}
func intHttp(w httphandler.WeatherHandler) {

	r := mux.NewRouter()
	addr :=  "127.0.0.1:3000"
	r.HandleFunc("/cities", w.CreateCityHandler).Methods("POST")
	r.HandleFunc("/cities/{id}", w.UpdateCityHandler).Methods("PATCH")
	r.HandleFunc("/cities/{id}", w.DeleteCityHandler).Methods("DELETE")
	r.HandleFunc("/temperatures", w.CreateTemperatureHandler).Methods("POST")
	r.HandleFunc("/forecasts/{city_id}", w.ForecastHandler).Methods("GET")
	r.HandleFunc("/webhooks", w.CreateWebHookHandler).Methods("POST")
	r.HandleFunc("/webhooks/{id}", w.DeleteWebHookHandler).Methods("DELETE")
	//r.HandleFunc("/{name}",w.ForecastHandler).Methods("GET")

	http.Handle("/", r)
	srv := &http.Server{
		Handler: r,
		Addr:   addr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Print("Running on ",addr)
	srv.ListenAndServe()

	log.Print("Exit")

}
