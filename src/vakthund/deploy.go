package vakthund

import (
	"os/exec"
	"log"
)

func executeDeploy(status Status) {
	log.Println("Deploying", status.Repository.FullName, ",commit:", status.Sha, "@", branchVar)
	out, err := exec.
		Command("/bin/sh", scriptVar, status.Sha, branchVar, status.Repository.FullName).
		Output()

	if err != nil {
		log.Println("Exectue error", err)
		return
	}

	log.Println(string(out))
}
