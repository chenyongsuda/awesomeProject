package http

import (
	"net/http"
	"g"
	"fmt"
	"log"
	"encoding/json"
	"github.com/toolkits/file"
)

func RenderJson(w http.ResponseWriter, data interface{}){
	jData,err := json.Marshal(data)
	if nil != err {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(jData)
}

func HandRequest() {
	http.HandleFunc("/version",func (w http.ResponseWriter, r *http.Request) {
		log.Println("version")
	})
	http.HandleFunc("/health",func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.HandleFunc("/workdir",func (w http.ResponseWriter, r *http.Request) {
		RenderJson(w,file.SelfDir())
	})
	http.HandleFunc("/config/infos",func (w http.ResponseWriter, r *http.Request) {
		RenderJson(w,g.Conf())
	})

}

func Start() error{
	if !g.Conf().Http.Enabled {
		return nil
	}

	if "" == g.Conf().Http.Port{
		return nil
	}

	address := g.Conf().Http.Port
	HandRequest()
	err := http.ListenAndServe(address,nil)
	if nil != err {
		return fmt.Errorf("%s",err)
	}
	return nil
}
