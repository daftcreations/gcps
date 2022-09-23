package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/daftcreations/gcps/pkg/cmd"
	"github.com/daftcreations/gcps/pkg/file"
	"github.com/daftcreations/gcps/pkg/gcp"
)

func main() {
	// validation
	if _, err := exec.LookPath("gcloud"); err != nil {
		fmt.Println("gcloud command not found")
		fmt.Println("Refer to this guide for help to setting up a gcloud")
		fmt.Println("https://cloud.google.com/sdk/docs/install")
		return
	}

	// store selected profile
	var selectedProfile string

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
				selectedProfile = lastProfile
			}
		} else {
			isContain := gcp.ContainsProfile(list, os.Args[1])
			if isContain {
				selectedProfile = os.Args[1]
			} else {
				log.Printf("Given profile %s is invalide please select from below!\n", os.Args[1])
			}

		}
	}

	if len(selectedProfile) == 0 {
		selectedProfile, err = cmd.GetProfileFromUser(list)
		if err != nil {
			log.Fatalf("Error while setting up cli: %v", err)
		}
	}

	activeProfile := gcp.GetActiveProfile(list)
	if activeProfile == "" {
		err := file.WriteLastProfile(selectedProfile)
		if err != nil {
			log.Fatalf("Error while writing last activate profile: %v", err)
		}
	} else if activeProfile == selectedProfile {
		fmt.Printf("Activate profile %s. No need to switch again\n", selectedProfile)
		return
	} else {
		err := file.WriteLastProfile(activeProfile)
		if err != nil {
			log.Fatalf("Error while writing last activate profile: %v", err)
		}
	}

	if err := gcp.SetProfile(selectedProfile); err != nil {
		log.Fatalf("Error setting up the profile: %v", err)
	}

	fmt.Printf("Switched profile to %s successfully!!\n", selectedProfile)
}
