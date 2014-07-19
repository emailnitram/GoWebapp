package main

import (
  "log"
  "os"
  "text/template"
)

type Person struct {
  Name string
  Emails []string
}

const tmpl = `{{$name := .Name}}
{{range .Emails}}
  Name is {{$name}}, email is {{.}}
{{end}}
`

func main() {
  person := Person{
    Name: "Martin",
    Emails: []string{"emailnitram@gmail.com","emailnitram2@gmail.com"},
  }

  t := template.New("Person template")

  t, err := t.Parse(tmpl)
  if err != nil {
    log.Fatal("Parse: ", err)
    return
  }

  err = t.Execute(os.Stdout, person)
  if err != nil {
    log.Fatal("Execute: ", err)
    return
  }

}

