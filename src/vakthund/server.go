package vakthund

import (
	"log"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

func payload(w http.ResponseWriter, req *http.Request) {

	//extract the event header
	event := req.Header.Get("X-Github-Event")
	log.Println("Event type:", event)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println("Read body error")
		panic(err)
	}

	//the logic for this request handler..
	if event == "status" {
		var status Status
		err = json.Unmarshal(body, &status)
		if err != nil {
			panic(err)
		}
		log.Println(status)

		if status.State == "success" {
			for _, branch := range status.Branches {
				if branch.Name == branchVar {
					if branch.Commit["sha"] == status.Sha {
						log.Println("We should deploy a new build on the dev server!")
						w.Write([]byte("A new build should be deployed to dev"))

						executeDeploy(status)
						log.Println("Deploy done..")
						return
					}
					w.WriteHeader(406)
					w.Write([]byte("Branch is correct but not commit sha"))
					return
				}
			}
		} else {
			log.Println("Nothing to do here...")
			w.WriteHeader(202)
			w.Write([]byte("State is not sucess, nothing to do"))

			return
		}
	}
	w.WriteHeader(202)
	w.Write([]byte("Only status events are interesting"))
}

var branchVar string
var scriptVar string
var devMode bool

func Start(branch,script string,dev bool) {
	branchVar = branch
	scriptVar = script
	devMode = dev

	payloadHandler := http.HandlerFunc(payload)
	http.Handle("/payload", HmacHandler(payloadHandler))
	http.ListenAndServe(":4545", nil)
}
