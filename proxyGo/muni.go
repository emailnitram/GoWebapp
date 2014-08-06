package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type RealTime struct {
	AgencyList Agency `xml:"AgencyList>Agency"`
}

type Agency struct {
	Name      string  `xml:"Name,attr"`
	Mode      string  `xml:"Mode,attr"`
	RouteList []Route `xml:"RouteList>Route"`
}

type Route struct {
	Name           string             `xml:"Name,attr"`
	RouteDirection RouteDirectionCode `xml:"RouteDirectionList>RouteDirection"`
}

type RouteDirectionCode struct {
	Code     string `xml:"Code,attr"`
	Name     string `xml:"Name,attr"`
	StopList Stop   `xml:"StopList>Stop"`
}

type Stop struct {
	Name              string        `xml:"name,attr"`
	StopCode          string        `xml:"StopCode,attr"`
	DepartureTimeList DepartureTime `xml:DepartureTimeList>DepartureTime"`
}

type DepartureTime struct {
	DepartureTime []string `xml:"DepartureTime"`
}

func main() {
	http.HandleFunc("/muni", func(w http.ResponseWriter, r *http.Request) {
		MUNI_TOKEN := os.Getenv("MUNI_TOKEN")
		resp, err := http.Get("http://services.my511.org/Transit2.0/GetNextDeparturesByStopCode.aspx?token=" + MUNI_TOKEN + "&stopcode=13915")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		//fmt.Println(string(body))
		var f RealTime
		err = xml.Unmarshal(body, &f)
		if err != nil {
			panic(err)
		}
		fmt.Println(f)

		w.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal(f)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":4747", nil))

}
