package main

import (
  "net/http"
  "fmt"
  "html/template"
  "log"
  "time"
  "github.com/shirou/gopsutil/cpu"
  "math"
  "encoding/json"
  "github.com/shirou/gopsutil/mem"
)
type JsonMemoria struct {
  Total uint64
  PorcentajeMEM float64
  PorcentajeCPU float64
}
type PageVariables struct {
  Date         string
  Time         string
}
func HomePage(w http.ResponseWriter, r *http.Request){

    now := time.Now() // find the time right now
    HomePageVars := PageVariables{ //store the date and time in a struct
      Date: now.Format("02-01-2006"),
      Time: now.Format("15:04:05"),
    }

    t, err := template.ParseFiles("grafica.html") //parse the html file homepage.html
    if err != nil { // if there is an error
      log.Print("template parsing error: ", err) // log it
    }
    err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
    if err != nil { // if there is an error
      log.Print("template executing error: ", err) //log it
    }
}
func totalCpu(w http.ResponseWriter, r *http.Request) {  
  vmStat,err := cpu.Percent(0,false);
  vmStat2,err := mem.VirtualMemory()
  jsonMemoria := JsonMemoria{bToMb(vmStat2.Total),math.Floor(vmStat2.UsedPercent*100)/100,math.Floor(vmStat[0]*100)/100 }
  js, err := json.Marshal(jsonMemoria) 
  _=err
  w.Write(js)
}
func dealwithErr(err error) {
  if err != nil {
      fmt.Println(err)
  }
}
func main() {
  http.HandleFunc("/hola", receiveAjax)
  http.HandleFunc("/", HomePage)
  http.HandleFunc("/ObtenerCPU", totalCpu)
  http.HandleFunc("/smoothie.js", serveCss)
  http.HandleFunc("/bootstrap.css", serveCss2)
  http.HandleFunc("/bootstrap.min.css", serveCss3)
  if err := http.ListenAndServe(":8000", nil); err != nil {
    panic(err)
  }
}
func serveCss(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "smoothie.js")
}
func serveCss2(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "bootstrap.css")
}
func serveCss3(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "bootstrap.min.css")
}
func receiveAjax(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
  w.Write([]byte("Calculos()")) // CALCULOS era un string con los valores necesarios separados por un espacio, para obtenerlos en el html con un split de espacios
  }
}
func bToMb(b uint64) uint64 {
    return b / 1024 / 1024
}
