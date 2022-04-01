package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/naman2706/gcps/pkg/cmd"
	"github.com/naman2706/gcps/pkg/file"
	"github.com/naman2706/gcps/pkg/gcp"
)

func main() {
	// validation
	if _, err := exec.LookPath("gcloud");  err != nil {
		fmt.Println("gcloud command not found")
		fmt.Println("Refer to this guide for help to setting up a gcloud")
		fmt.Println("https://cloud.google.com/sdk/docs/install")
		return
	}

	// store seleted profile
	var seletedProfile string

	list, err := gcp.GetProfileList()
	if err != nil {
		log.Fatalf("Error getting profile list: %v", err)
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "--help" {
			fmt.Println("usage: gcps [arguments] or flags")
			fmt.Print("\t Switch google cloud profile\n\n")
			fmt.Println("\t Arguments:")
			fmt.Println("\t\t[name]\tswitch profile to given name.")
			fmt.Print("\t Flags:\n")
			fmt.Println("\t\t-\tswitch to previous profile.")
			fmt.Print("\t\t--help\tto get help about gcps.\n\n")
			return
		} else if os.Args[1] == "-" {
			lastProfile, err := file.ReadLastProfile()
			if err != nil {
				log.Fatalf("Error while reading file: %v", err)
			} else if lastProfile == "" {
				fmt.Println("No previous profile found. Please select profile from below")
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
		err := file.WriteLastProfile(seletedProfile)
		if err != nil {
			log.Fatalf("Error while writing last activate profile: %v", err)
		}
	} else if activeProfile == seletedProfile {
		fmt.Printf("Activate profile %s. No need to switch again\n", seletedProfile)
		return
	} else {
		err := file.WriteLastProfile(activeProfile)
		if err != nil {
			log.Fatalf("Error while writing last activate profile: %v", err)
		}
	}

	if err := gcp.SetProfile(seletedProfile); err != nil {
		log.Fatalf("Error setting up the profile: %v", err)
	}

	fmt.Printf("Switched profile to %s successfully!!\n", seletedProfile)
}
