package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	town = "Ekaterinburg"
	api  = "42b582831e095766d825429ec639a096" //Регаемся на openweathermap.org и получаем ключ
)

func main() {

	urlString := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s", town, api) // формируем http запрос
	ans, err := http.Get(urlString)

	if err != nil { //Проверяем наличие ошибок
		fmt.Println(err)
		return
	}

	body, err := ioutil.ReadAll(ans.Body) //Читаем тело ответа

	if err != nil {
		fmt.Println(err)
		return
	}

	var data map[string]interface{} //создаем карту с пустым интерфейсом

	err = json.Unmarshal(body, &data)

	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(data) // json

	//Достаем нужные данные вручную (потом автоматизируем)
	log.Printf("\n") //Дата и время
	fmt.Printf("Погода: %s \n", data["weather"].([]interface{})[0].(map[string]interface{})["description"])
	fmt.Printf("Температура: %.0f *C\n", data["main"].(map[string]interface{})["temp"])
	fmt.Printf("Ощущается: %.0f *C\n", data["main"].(map[string]interface{})["feels_like"])
	fmt.Printf("Давление: %.0f hPa\n", data["main"].(map[string]interface{})["pressure"])
	fmt.Printf("Влажность: %.0f %% \n", data["main"].(map[string]interface{})["humidity"])
	fmt.Printf("Скорость ветра: %.0f м/с \n", data["wind"].(map[string]interface{})["speed"])
	fmt.Printf("Облачность: %.0f %% \n", data["clouds"].(map[string]interface{})["all"])

	/* Получаем:
	    2022/01/23 20:15:02
            Погода: mist
	    Температура: -17 *C
	    Ощущается: -17 *C
	    Давление: 1035 hPa
	    Влажность: 92 %
	    Скорость ветра: 0 м/с
	    Облачность: 100 %
	*/
}
