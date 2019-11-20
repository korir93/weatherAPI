package main

import (
	"net/http"
	"encoding/json"	
	"io/ioutil"
	"fmt"
	"weatheropenAPI/config"
	"github.com/gorilla/mux"

)
type weatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int    `json:"type"`
		ID      int    `json:"id"`
		Country string `json:"country"`
		Sunrise int    `json:"sunrise"`
		Sunset  int    `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
} 
 func buildconfiguration()config.Config{
	//var conf config.Config
	conf:=config.Config{
		Apikey: "f5e8f46f77b7bbc61a3dd202a1bdcf24",
	}
	config.SetConfig(conf) 
		return conf
 }

func main()  {
	
	router := mux.NewRouter()
	router.HandleFunc("/api/getbyname", getbycityname).Methods("GET")
	http.ListenAndServe(":5000", router)		
}
func getbycityname(w http.ResponseWriter, r *http.Request){
	city:= r.URL.Query().Get("city")
	data,_:=get(city)
//	if  r!=nil {
//		return
//	}
     w.Header().Add("Content-Type", "application/json") 
	 w.Write(data) 
	 return
fmt.Println(string(data))
}
func  get(city string)([]byte,error){
	client := &http.Client{}
	//city:="nairobi"
	api_key:="f5e8f46f77b7bbc61a3dd202a1bdcf24"
	req, err := http.NewRequest("GET", "http://api.openweathermap.org/data/2.5/weather?"+"appid="+api_key +"&q="+city, nil) 	
	if err != nil {
				 		
	 //return nil 
	}
	
	  res := weatherResponse{}
	  resp, err := client.Do(req) 
		  if err != nil { 		
                 return nil,err
  	} 

			  defer resp.Body.Close()  	
			  d, _ := ioutil.ReadAll(resp.Body)
			  json.Unmarshal(d,&res)
			  return d,nil
}