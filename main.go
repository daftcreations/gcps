package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/naman2706/gcps/pkg/cmd"
	"github.com/naman2706/gcps/pkg/file"
	"github.com/naman2706/gcps/pkg/gcp"
)

// to unmarshal gcloud configurations json
type List struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func main() {
	// validation
	if _, err := exec.LookPath("gcloud"); err != nil {
		log.Println("gcloud command not found")
		log.Println("Refer to this guide for help to setting up a gcloud")
		log.Println("https://cloud.google.com/sdk/docs/install")
		return
	}

	// store seleted profile
	var seletedProfile string

	list, err := gcp.GetProfileList()
	if err != nil {
		log.Fatalf("Error getting profile list: %v", err)
	}

	path, err := file.PrevProfileFile()
	if err != nil {
		log.Fatalf("Error while getting file: %v", err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "-" {
			lastProfile, err := file.ReadLastProfile(path)
			if err != nil {
				log.Fatalf("Error while reading file: %v", err)
			} else if lastProfile == "" {
				log.Println("No previous profile found. Please select profile from below")
			} else {
				seletedProfile = lastProfile
			}
		} else {
			isContain := gcp.ContainsProfile(list, os.Args[1])
			if isContain {
				seletedProfile = os.Args[1]
			} else {
				log.Printf("Given profile %s is invalide please select from below!\n", os.Args[1])
			}

		}
	}

	if len(seletedProfile) == 0 {
		seletedProfile, err = cmd.GetProfileFromUser(list)
		if err != nil {
			log.Fatalf("Error while setting up cli: %v", err)
		}
	}

	activeProfile := gcp.GetActiveProfile(list)
	if activeProfile == "" {
		file.WriteLastProfile(path, seletedProfile)
	} else if activeProfile == seletedProfile {
		log.Printf("Activate profile %s. No need to switch again\n", seletedProfile)
		return
	} else {
		file.WriteLastProfile(path, activeProfile)
	}

	if err := gcp.SetProfile(seletedProfile); err != nil {
		log.Fatalf("Error setting up the profile: %v", err)
	}

	log.Printf("Switched profile to %s successfully!!", seletedProfile)
}
