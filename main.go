package main

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

// to unmarshal gcloud configurations json
type List struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func main() {
	// seleted profile by user
	var seletedProfile string
	// to set activated profile as default 
	var activeProfile string
	// list of profiles
	profileList := []string{}

	// run cmd to get list of profiles
	out, err := exec.Command("gcloud", "config", "configurations", "list", "--format=json").Output()
	if err != nil {
		log.Fatal(err)
	}

	// temporary list of profiles
	l := []List{}

	// unmarshal to get the desired value
	if json.Unmarshal(out, &l); err != nil {
		log.Fatal(err)
	}

	for _, e := range l {
		if e.IsActive {
			activeProfile = e.Name
		}
		profileList = append(profileList, e.Name)
	}

	// survey
	p := []*survey.Question{
		{
			Name: "profiles",
			Prompt: &survey.Select{
				Message: "Select Profile:",
				Options: profileList,
				Default: activeProfile,
			},
			Validate: survey.Required,
		},
	}

	// ask 
	err = survey.Ask(p, &seletedProfile)
	if err != nil {
		log.Fatal(err)
	}

	// set the given profile
	if exec.Command("gcloud", "config", "configurations", "activate", seletedProfile).Run(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Switched profile to %s successfully!!", seletedProfile)
}
