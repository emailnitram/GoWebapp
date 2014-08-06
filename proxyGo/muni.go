package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RealTime struct {
	//RTT xml.Name `xml:"RTT"`
	AgencyList Agency `xml:"AgencyList>Agency"`
}

type Agency struct {
	Name      string  `xml:"Name,attr"`
	Mode      string  `xml:"Mode,attr"`
	RouteList []Route `xml:"RouteList>Route"`
}

type Route struct {
	Name string `xml:"Name,attr"`
}

type DepartureTime struct {
	DepartureTime string `xml:"DepartureTime"`
}

func main() {
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

}
