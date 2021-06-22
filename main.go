package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/common/model"
)

type targetgroup struct {
	Targets []string
	Labels  model.LabelSet
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/job/{job_id}/targets", ProcessTargets)
	http.HandleFunc("/targets", ProcessTargets)
	http.ListenAndServe(":8080", router)
}
func ProcessTargets(w http.ResponseWriter, r *http.Request) {
	collector := r.URL.Query()["collector"]
	vars := mux.Vars(r)
	jobName := vars["job_id"]
	fmt.Println(collector)
	fmt.Println("Request made at collector:", collector[0], "for job_id:", jobName)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	tgs := []targetgroup{
		{
			Targets: []string{"127.0.0.1:9090", "127.0.0.2:9090", "127.0.0.3:9090"},
			Labels: model.LabelSet{
				"__meta_label_datacenter": "dc3",
			},
		},
		{
			Targets: []string{"127.0.1.1:9090", "127.0.1.2:9090", "127.0.1.3:9090"},
			Labels: model.LabelSet{
				"__meta_label_datacenter": "us2",
			},
		},
	}
	json.NewEncoder(w).Encode(tgs)
}
