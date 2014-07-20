package main

import (
  "encoding/json"
  "fmt"
  "html/template"
  "io/ioutil"
  "log"
  "net/http"
  "net/url"
  "os"
)

func main() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/showimage", showimage)
  fmt.Println("listening...")
  err := http.ListenAndServe(GetPort(), nil)
  if err != nil {
    log.Fatal("ListenAndServe: ", err)
    return
  }
}

// Get the Port form the environment so we can run on Heroku
func GetPort() string {
  var port = os.Getenv("PORT")
  // Set a default port if there is nothing in the environment
  if port == ""{
    port = "4747"
    fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
  }
  return ":" + port
}

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!doctype html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Accept address</title>
      </head>
      <body>
        <h1>Accept Address</h1>
        <p>Please enter your address:</p>
        <form action="/showimage" method="post" accept-charset="utf-8">
          <input type="text" name="str" placeholder="Type address..." id="str">
          <input type="submit" value="... and see the image!">
        </form>
      </body>
    <html>
`

var upperTemplate = template.Must(template.New("showimage").Parse(upperTemplateHTML))

func showimage(w http.ResponseWriter, r *http.Request) {
  // Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
  addr := r.FormValue("str")

  // QueryEscape escapes the addr string so
  // it can be safely places inside a URL query
  safeAddr := url.QueryEscape(addr)
  fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?sensor=false&address=%s", safeAddr)

  // For control over HTTP client headers,
  // redirect policy, and other settings,
  // create a Client
  // A Client is an HTTP client
  client := &http.Client{}

  // Build the request
  req, err := http.NewRequest("GET", fullUrl, nil)
  if err != nil {
    log.Fatal("NewRequest: ", err)
    return
  }

  // Send the request via a client
  // Do sends an HTTP request and 
  // returns an HTTP response
  resp, requestErr := client.Do(req)
  if requestErr != nil {
    log.Fatal("Do: ", requestErr)
    return
  }

  // Callers should close resp.Body
  // when done reading from it
  // Defer the closing of the body
  defer resp.Body.Close()

  // Read the content into a byte array
  body, dataReadErr := ioutil.ReadAll(resp.Body)
  if dataReadErr != nil {
    log.Fatal("ReadAll: ", dataReadErr)
    return
  }

  res := make(map[string][]map[string]map[string]map[string]interface{},0)

  // We will be using the Unmarshal function
  // to transform our JSON bytes into the
  // appropriate structure
  // The Unmarshal function accepts a byte array
  // and a reference to the object which shall be
  // filled with the JSON data (this is simplifying,
  // it actually accepts an interface)
  json.Unmarshal(body, &res)

  lat, _ := res["results"][0]["geometry"]["location"]["lat"]
  lng, _ := res["results"][0]["geometry"]["location"]["lng"]

  // %.13f is used to convert float64 to a string
  queryUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/streetview?sensor=false&size=600x300&location=%.13f,%.13f", lat, lng)

  tempErr := upperTemplate.Execute(w, queryUrl)
  if tempErr != nil {
    http.Error(w, tempErr.Error(), http.StatusInternalServerError)
  }
}

const upperTemplateHTML = `
<!doctype html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Display Image</title>
    </head>
    <body>
      <h3>Image at your Address</h3>
      <img src="{{html .}}" alt="Image" />
    </body>
  </html>
`
