package httphandler

import (
	db2 "WeatherMonster/pkg/db"
	"WeatherMonster/pkg/io"
	"WeatherMonster/pkg/service"
	"WeatherMonster/pkg/utile"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var handler WeatherHandler

func InitConfig() {

	utile.GetConf()
	db := db2.New(db2.Connect())
	weatherService := service.NewBasiWeatherService(db)
	handler = NewWatherHandler(weatherService)

}

/*
// TRUE SUCCESS SCENARIO
func TestCreateCityHandler(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.City{
		Name:      "xyz",
		Latitude:  22,
		Longitude: 32,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/cities", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateCityHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
		log.Print("Response", string(rr.Body.Bytes()))
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	tempData := m["data"].(map[string]interface{})

	log.Print("Response ", string(rr.Body.Bytes()))
	if tempData["name"].(string) != reqData.Name {
		t.Errorf("handler returned unexpected city: got %v want %v",
			tempData["name"].(string), reqData.Name)
	}
	if tempData["longitude"].(float64) != reqData.Longitude {
		t.Errorf("handler returned unexpected max: got %v want %v",
			tempData["longitude"].(float64), reqData.Longitude)
	}
	if tempData["latitude"].(float64) != reqData.Latitude {
		t.Errorf("handler returned unexpected min: got %v want %v",
			tempData["latitude"].(float64), reqData.Latitude)
	}
}
*/
// FALSE SUCCESS SCENARIO
func TestCreateCityHandlerCheckName(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.City{
		Name:      "",
		Latitude:  22,
		Longitude: 32,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/cities", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateCityHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "name required" {
		t.Errorf("handler returned unexpected city: got %v want %v",
			m["error"].(string), "name required")
		log.Print("error", m["error"].(string))
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["longitude"].(float64), reqData.Longitude)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["data"].(float64), "")
	}

}

// FALSE SUCCESS SCENARIO
func TestCreateCityHandlerCheckLongitude(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.City{
		Name:      "Berlin",
		Latitude:  22,
		Longitude: 0,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/cities", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateCityHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "longitude required" {
		t.Errorf("handler returned unexpected city: got %v want %v",
			m["error"].(string), "name required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["longitude"].(float64), reqData.Longitude)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["data"].(float64), "")
	}
}

// FALSE SUCCESS SCENARIO
func TestCreateCityHandlerCheckLatitude(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.City{
		Name:      "Berlin",
		Latitude:  0,
		Longitude: 5.5646464,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/cities", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateCityHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "latitude required" {
		t.Errorf("handler returned unexpected city: got %v want %v",
			m["error"].(string), "name required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["success"].(bool), false)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["data"].(float64), "")
	}
}

// TRUE SUCCESS SCENARIO
func TestCreateTemperatureHandler(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	tempReq := io.Temperatures{
		CityID: 1,
		Min:    22,
		Max:    32,
	}
	bytData, _ := json.Marshal(&tempReq)
	req, err := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateTemperatureHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	tempData := m["data"].(map[string]interface{})

	log.Print("Response ", string(rr.Body.Bytes()))
	if int(tempData["city_id"].(float64)) != tempReq.CityID {
		t.Errorf("handler returned unexpected city: got %v want %v",
			tempData["city_id"].(int), tempReq.CityID)
	}
	if tempData["max"].(float64) != tempReq.Max {
		t.Errorf("handler returned unexpected max: got %v want %v",
			tempData["max"].(float64), tempReq.Max)
	}
	if tempData["min"].(float64) != tempReq.Min {
		t.Errorf("handler returned unexpected min: got %v want %v",
			tempData["min"].(float64), tempReq.Min)
	}
}

// FALSE SUCCESS SCENARIO
func TestCreateTemperatureHandlerCheckCityID(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.Temperatures{
		CityID: 0,
		Min:    22,
		Max:    32,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateTemperatureHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "city id required" {
		t.Errorf("handler returned unexpected city id: got %v want %v",
			m["error"].(string), "city id required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected success: got %v want %v",
			m["success"].(bool), false)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected data: got %v want %v",
			m["data"].(float64), "")
	}
}

// FALSE SUCCESS SCENARIO
func TestCreateTemperatureHandlerCheckMin(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.Temperatures{
		CityID: 1,
		Min:    0,
		Max:    32,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateTemperatureHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "min required" {
		t.Errorf("handler returned unexpected min: got %v want %v",
			m["error"].(string), "min required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected success: got %v want %v",
			m["success"].(bool), false)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected data: got %v want %v",
			m["data"].(float64), "")
	}
}

// FALSE SUCCESS SCENARIO
func TestCreateTemperatureHandlerCheckMax(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.Temperatures{
		CityID: 1,
		Min:    0,
		Max:    32,
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/temperatures", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateTemperatureHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "min required" {
		t.Errorf("handler returned unexpected max: got %v want %v",
			m["error"].(string), "max required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected success: got %v want %v",
			m["success"].(bool), false)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected data: got %v want %v",
			m["data"].(float64), "")
	}
}

// TRUE SUCCESS SCENARIO
func TestCreateWebHookHandler(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	tempReq := io.WebHook{
		CityID:      1,
		CallbackUrl: "localhost:3000",
	}
	bytData, _ := json.Marshal(&tempReq)
	req, err := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateWebHookHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)
	tempData := m["data"].(map[string]interface{})

	log.Print("Response ", string(rr.Body.Bytes()))
	if int(tempData["city_id"].(float64)) != tempReq.CityID {
		t.Errorf("handler returned unexpected city: got %v want %v",
			tempData["city_id"].(int), tempReq.CityID)
	}
	if tempData["callback_url"].(string) != tempReq.CallbackUrl {
		t.Errorf("handler returned unexpected callback_url: got %v want %v",
			tempData["max"].(float64), tempReq.CallbackUrl)
	}
}

// FALSE SUCCESS SCENARIO
func TestCreateWebHookHandlerCheckCityId(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.WebHook{
		CityID:      0,
		CallbackUrl: "localhost:3000",
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateWebHookHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "city id required" {
		t.Errorf("handler returned unexpected city id: got %v want %v",
			m["error"].(string), "city id required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected success: got %v want %v",
			m["success"].(bool), false)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected data: got %v want %v",
			m["data"].(float64), "")
	}
}

// FALSE SUCCESS SCENARIO
func TestCreateWebHookHandlerCheckCallbackUrl(t *testing.T) {
	InitConfig()
	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter.
	reqData := io.WebHook{
		CityID:      1,
		CallbackUrl: "",
	}
	bytData, _ := json.Marshal(&reqData)
	req, err := http.NewRequest("POST", "/webhooks", bytes.NewBuffer(bytData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-Type", "application/json")
	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.CreateWebHookHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	// Check the response body is what we expect.

	var m map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &m)

	log.Print("Response ", string(rr.Body.Bytes()))
	if m["error"].(string) != "call back url required" {
		t.Errorf("handler returned unexpected callback url: got %v want %v",
			m["error"].(string), "call back url required")
	}
	if m["success"].(bool) != false {
		t.Errorf("handler returned unexpected success: got %v want %v",
			m["success"].(bool), false)
	}
	if m["data"] != nil {
		t.Errorf("handler returned unexpected data: got %v want %v",
			m["data"].(float64), "")
	}
}
