package main

import (
	"net/http"
        "io"
	"io/ioutil"
        //"net"
	"github.com/go-chi/chi"
        "github.com/go-chi/chi/middleware"
        //"github.com/tidwall/sjson"
        //"github.com/tidwall/gjson"
	"gopkg.in/natefinch/lumberjack.v2"
         "os"
         "encoding/json"
         "fmt"
         "log"
         "time"
)

var builddate string
var gitversion string 
var TheLog *lumberjack.Logger

var router *chi.Mux

func routers() *chi.Mux {
     router.Get("/redfish/v1/Systems",GetSystems)
     router.Get("/redfish/v1/Systems/{systemid}",GetSystemInfo)
     router.Get("/redfish/v1/Systems/{systemid}/BIOS",GetSystemBios)
     router.Get("/redfish/v1/Systems/{systemid}/EthernetInterfaces",GetSystemEthernet)
     router.Get("/redfish/v1/Systems/{systemid}/LogServices",GetSystemLogServices)
     router.Get("/redfish/v1/Systems/{systemid}/Memory",GetSystemMemory)
     router.Get("/redfish/v1/Systems/{systemid}/Processors",GetSystemProc)
     router.Get("/redfish/v1/Systems/{systemid}/SimpleStorage",GetSystemStorage)
     
     router.Get("/healthz",ReadyCheck)
     router.Get("/alivez", AliveCheck)
     return(router)
}

func init() { 
    router = chi.NewRouter() 
    router.Use(middleware.Recoverer)  
    router.Use(middleware.RequestID)
    router.Use(middleware.Logger)
    router.Use(middleware.Recoverer)
    //router.Use(middleware.URLFormat)
}

func EnableRestServices() {
        r := routers()
        http.ListenAndServe(":8085", r)
}


func GetSystems(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystems %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemInfo(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemInfo %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemBios(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemBios %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemEthernet(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemEthernet %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemLogServices(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemLogServices %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemMemory(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemMemory %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemProc(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemProc %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func GetSystemStorage(w http.ResponseWriter, r *http.Request) { 
    log.Printf("GetSystemStorage %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func main() {
        os.MkdirAll("/Program Files/WindowsNodeManager/logs",0755)
        TheLog := &lumberjack.Logger{
                Filename:   "/data/redfish.log",
                MaxSize:    1, // megabytes
                MaxBackups: 6,
                MaxAge:     1, // days
                Compress:   true, // disabled by default
                }
        log.SetOutput(TheLog)
        log.Printf("RedFish2Esi Restarted - version: %s - build Data: %s",gitversion,builddate)
        EnableRestServices()
}

func ReadyCheck(w http.ResponseWriter, r *http.Request) { 
    log.Printf("ReadyCheck %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "ready"})
}

func AliveCheck(w http.ResponseWriter, r *http.Request) { 
    log.Printf("ReadyCheck %s\n", r.Body)
    respondwithJSON(w, http.StatusOK, map[string]string{"message": "alive"})
}


func DownloadFile(theurl string, filepath string) error{
// Create the file
  log.Printf("Download File: %s -> %s\n",theurl,filepath)
  out, err := os.Create(filepath)
  if err != nil  {
    log.Printf("Cannot Create File: %v\n",err)
    return err
  }
  defer out.Close()

  // Get the data
  resp, err := http.Get(theurl)
  if err != nil {
    log.Printf("Cannot Get File: %v\n",err)
    return err
  }
  defer resp.Body.Close()

  // Check server response
  if resp.StatusCode != http.StatusOK {
    log.Printf("Bad Response downloading file: %s\n",resp.Status)
    return fmt.Errorf("bad status: %s", resp.Status)
  }

  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    log.Printf("Cannot Copy File: %v\n",err)
    return err
  }
  return nil

}

func ReadFile(thepath string) string {
    b, err := ioutil.ReadFile(thepath) // just pass the file name
    if err != nil {
        //log.Print(err)
        return ""
    }
    str := string(b)
   return str
}

func WriteFile(thepath string,data string){
     ioutil.WriteFile(thepath,[]byte(data), 0600)
}


func wait_for_file(filename string){

        total_time := 0;
        time_limit := 60 * 15 // 15 Minutes
	for {
          if (fileExists(filename)){
             return
             }
         time.Sleep(2 * time.Second)
         total_time = total_time + 2
         if (total_time > time_limit){
            log.Printf("Timout waiting for done file %s\n",filename)
            return
            }
         }



}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}


// respondwithError return error message
func respondWithError(w http.ResponseWriter, code int, msg string) {
    respondwithJSON(w, code, map[string]string{"message": msg})
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    fmt.Println(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}

