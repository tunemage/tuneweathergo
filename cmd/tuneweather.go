package cmd

import (
	"fmt"
	"github.com/spf13/cobra"

	"encoding/json"
	"io"
    "net/http"
)

var parameters struct {
	city string
}

// tuneweatherCmd represents the tuneweather command
var tuneweatherCmd = &cobra.Command{
	Use:   "tuneweather",
	Short: "tw",
	Long: `Description
改行が使える`,
    //　処理を行うところ
	Run: func(cmd *cobra.Command, args []string) {
		latitude, longitude := getLatitudeAndLongtitude(parameters.city)
		url := "https://api.open-meteo.com/v1/forecast?latitude=" + latitude + "&longitude=" + longitude + "&daily=weathercode&timezone=Asia%2FTokyo"

		req, _ := http.NewRequest(http.MethodGet, url, nil)

		client := new(http.Client)
		res, err := client.Do(req)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			fmt.Println("Error", res.Status)
		}

		body, _ := io.ReadAll(res.Body)

		var weather Weather 

		json.Unmarshal(body, &weather)

		// for in 0...5みたいな事をしたい
		array := []int{0, 1, 2, 3, 4 ,5}
		for i, _ := range array {
			fmt.Println(weather.Daily.Time[i],":",getWeather(weather.Daily.Weathercode[i]))
		}
	},
}

//WMO codeから記号を返す（コード体系を分かってないので間違ってる可能性大。その他は一律「分からん」で返す）
func getWeather(code int) string{
    if 1 <= code && code <= 3 {
        return "☀"
    }else if 60 <= code && code <= 69 {
        return "☂"
    }else if 70 <= code && code <= 79 {
        return "☃"
    }else{
        return "分からん"
    }
}

//都市名から、緯度経度を返す。東京・大阪・名古屋のみ対応
func getLatitudeAndLongtitude(city string)(string, string){
	if city == "tokyo" {
		return "35.6785","139.6823"
	} else if city == "osaka" {
		return "34.6723","135.4848"
	} else if city == "nagoya" {
		return "35.1833","136.8999"
	} else {
		return "35.6785","139.6823" 
	}
}

// 引数の設定とかを行うところ
func init() {
	rootCmd.AddCommand(tuneweatherCmd)

	tuneweatherCmd.Flags().StringVarP(&parameters.city,"city","c","","都市名（tokyo,osaka,nagoyaのみ）")
}

type Weather struct {
	Latitude   float64
    Longitude  float64
	Elevation  float64
    Generationtime_ms float64
	Daily_units DailyUnits
	Daily Daily
}

type DailyUnits struct {
	Weathercode   string
    Time  string
}

type Daily struct {
	Weathercode   []int
    Time  []string
}