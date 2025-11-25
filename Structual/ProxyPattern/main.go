package main

import (
	"fmt"
	"time"
)

type Subject interface {
	Request()
}

type WetherService struct {
	apiKey  string
	baseURL string
}

func (ws *WetherService) Request() {
	fmt.Println("Retrieving weather data from the API...")
}

type WetherServiceProxy struct {
	wetherService WetherService
	cache         map[string]bool
}

func NewWetherServiceProxy(apiKey, baseURL string) *WetherServiceProxy {
	return &WetherServiceProxy{
		wetherService: WetherService{apiKey: apiKey, baseURL: baseURL},
		cache:         make(map[string]bool),
	}
}

func (wsp *WetherServiceProxy) generateCacheKey() string {
	return time.Now().Format("2006-01-02 15:04")
}

func (wsp *WetherServiceProxy) Request() {
	cacheKey := wsp.generateCacheKey()
	if _, found := wsp.cache[cacheKey]; found {
		fmt.Println("Returning cached weather data")
		return
	}

	wsp.wetherService.Request()
	wsp.cache[cacheKey] = true
}

func main() {
	WetherService := &WetherService{apiKey: "your_api_key", baseURL: "https://api.weather.com"}
	WetherService.Request()

	WetherServiceProxy := NewWetherServiceProxy("your_api_key", "https://api.weather.com")
	WetherServiceProxy.Request()

	WetherServiceProxy.Request()
}