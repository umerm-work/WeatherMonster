package service

import (
	"WeatherMonster/pkg/db"
	"WeatherMonster/pkg/io"
	"context"
)

// CityService stores function signature of business operations
type WeatherService interface {
	CityCreate(ctx context.Context, u io.City) (res io.Response)
	CityUpdate(ctx context.Context, u io.City) (res io.Response)
	CityDelete(ctx context.Context, u io.City) (res io.Response)
	TemperatureCreate(ctx context.Context, u io.Temperatures) (res io.Response)
	Forecast(ctx context.Context, u io.Forecast) (res io.Response)
	WebHookCreate(ctx context.Context, u io.WebHook) (res io.Response)
	WebHookDelete(ctx context.Context, u io.WebHook) (res io.Response)
}

type weatherService struct {
	DbRepo db.Repository
	//logger     log.Logger
}

// NewBasiCityService binds cityservice struct to CityService interface
func NewBasiWeatherService(DbRepo db.Repository) WeatherService {
	return &weatherService{
		DbRepo: DbRepo,
	}
}
func (b *weatherService) TemperatureCreate(ctx context.Context, c io.Temperatures) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.TemperatureCreate(ctx, c)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}
func (b *weatherService) CityCreate(ctx context.Context, c io.City) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.CityCreate(ctx, c)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}
func (b *weatherService) CityUpdate(ctx context.Context, c io.City) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.CityUpdate(ctx, c)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}

func (b *weatherService) CityDelete(ctx context.Context, c io.City) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.CityDelete(ctx, c)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}
func (b *weatherService) Forecast(ctx context.Context, c io.Forecast) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.ForecastRead(ctx, c.CityID)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}
func (b *weatherService) WebHookCreate(ctx context.Context, c io.WebHook) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.WebHooksCreate(ctx, c)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}

func (b *weatherService) WebHookDelete(ctx context.Context, c io.WebHook) (res io.Response) {

	var err error
	res.Data ,err = b.DbRepo.WebHooksDelete(ctx, c)
	if err != nil {
		res.Error = err.Error()
		res = io.FailureMessage(res.Error)
		return
	}
	res = io.SuccessMessage(res.Data)
	return
}