package db

import (
	"WeatherMonster/pkg/io"
	"WeatherMonster/pkg/utile"
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
	"time"
)

// Repository for city data operations
type Repository interface {
	TemperatureCreate(ctx context.Context, data io.Temperatures) (t io.Temperatures, err error)
	ForecastRead(ctx context.Context, cityId int) (t io.Forecast, err error)
	CityCreate(ctx context.Context, data io.City) (city io.City, err error)
	CityUpdate(ctx context.Context, data io.City) (city io.City, err error)
	CityDelete(ctx context.Context, data io.City) (city io.City, err error)
	WebHooksCreate(ctx context.Context, data io.WebHook) (w io.WebHook, err error)
	WebHooksDelete(ctx context.Context, data io.WebHook) (w io.WebHook, err error)
	IsCityExist(ctx context.Context, cityId int) (isExist bool, err error)
}

type repository struct {
	db *gorm.DB
	//logger log.Logger
}

// Connect use to connect the database
func Connect() *gorm.DB {
	arg := fmt.Sprintf("host=%s port=%s sslmode=disable dbname=%s user=%s password=%s",
		os.Getenv(utile.DB_Host),
		os.Getenv(utile.DB_Port),
		os.Getenv(utile.DB_Db),
		os.Getenv(utile.DB_User),
		os.Getenv(utile.DB_Password))
	db, err := gorm.Open("postgres", arg)
	if err != nil {
		log.Print(arg)
		panic(err.Error())
	}
	//Migrate the schema
	db.AutoMigrate(&io.City{})
	db.AutoMigrate(&io.Temperatures{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&io.WebHook{}).AddForeignKey("city_id", "cities(id)", "RESTRICT", "RESTRICT")
	db.LogMode(true)
	return db
}

// New binds cityservice struct to Repository interface
func New(db *gorm.DB) Repository {
	// return  repository
	return &repository{
		db: db,
	}
}
func (b *repository) IsCityExist(ctx context.Context, cityId int) (isExist bool, err error) {
	if cityId == 0 {
		isExist = false
		err = fmt.Errorf(`data doesn't exist`)
		return
	}
	var u io.City
	err = b.db.Where(io.City{ID: cityId}).Find(&u).Error
	if err != nil {
		isExist = false
		err = fmt.Errorf(`data doesn't exist`)
		return
	}
	isExist = true
	return
}
func (b *repository) TemperatureCreate(ctx context.Context, data io.Temperatures) (t io.Temperatures, err error) {

	var isExist bool
	if isExist, err = b.IsCityExist(ctx, data.CityID); !isExist {
		err = fmt.Errorf(err.Error())
		return
	}
	data.ID = 0
	now := time.Now()
	data.Timestamp = now.Unix()
	err = b.db.Save(&data).Scan(&t).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
		err = fmt.Errorf(err.Error())
	}
	return
}
func (b *repository) CityCreate(ctx context.Context, data io.City) (city io.City, err error) {

	data.ID = 0
	err = b.db.Save(&data).Scan(&city).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
		err = fmt.Errorf(err.Error())
	}
	return
}
func (b *repository) ForecastRead(ctx context.Context, cityId int) (t io.Forecast, err error) {

	var u io.Temperatures
	var count int
	err = b.db.Select("avg(max) as max,avg(min) as min").
		Where(`"city_id"= ?`, cityId).
		Find(&u).
		Count(&count).
		Error
	if err != nil {
		err = fmt.Errorf(`data doesn't exist`)
		return
	}
	t.CityID = cityId
	t.Min = u.Min
	t.Max = u.Max
	t.Sample = count

	return
}

func (b *repository) CityUpdate(ctx context.Context, data io.City) (city io.City, err error) {

	var u io.City
	err = b.db.Where(io.City{ID: data.ID}).Find(&u).Error
	if err != nil {
		err = fmt.Errorf(`data doesn't exist`)
		return
	}

	data.CreatedAt = u.CreatedAt
	city = data
	err = b.db.Save(&data).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return
}

func (b *repository) CityDelete(ctx context.Context, data io.City) (city io.City, err error) {

	var u io.City
	err = b.db.Where(io.City{ID: data.ID}).Find(&u).Error
	if err != nil {
		err = fmt.Errorf(`data doesn't exist`)
		return
	}

	city = u
	err = b.db.Delete(&u).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
		err = fmt.Errorf(err.Error())
		return
	}
	b.db.Delete(io.Temperatures{}).Where(`"city_id" = ?`, city.ID)
	b.db.Delete(io.WebHook{}).Where(`"city_id" = ?`, city.ID)
	return
}
func (b *repository) WebHooksCreate(ctx context.Context, data io.WebHook) (w io.WebHook, err error) {

	var isExist bool
	if isExist, err = b.IsCityExist(ctx, data.CityID); !isExist {
		err = fmt.Errorf(err.Error())
		return
	}
	data.ID = 0
	err = b.db.Save(&data).Scan(&w).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
		err = fmt.Errorf(err.Error())
	}
	return
}
func (b *repository) WebHooksDelete(ctx context.Context, data io.WebHook) (w io.WebHook, err error) {

	err = b.db.Where(io.WebHook{ID: data.ID}).Find(&w).Error
	if err != nil {
		err = fmt.Errorf(`data doesn't exist`)
		return
	}

	err = b.db.Delete(&w).Error
	if err != nil {
		log.Printf("Failed to save error: %v", err)
	}
	return
}
