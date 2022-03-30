package gcp

import (
	"encoding/json"
	"os/exec"

	"github.com/naman2706/gcps/pkg/types"
)

// get list of profile
func GetProfileList() ([]types.List, error) {
	// run cmd to get list of profiles
	out, err := exec.Command("gcloud", "config", "configurations", "list", "--format=json").Output()
	if err != nil {
		return []types.List{}, err
	}

	// list of profiles
	list := []types.List{}

	// unmarshal to get the desired value
	if err := json.Unmarshal(out, &list); err != nil {
		return []types.List{}, err
	}

	return list, nil
}

// set the given profile
func SetProfile(profile string) error {
	if err := exec.Command("gcloud", "config", "configurations", "activate", profile).Run(); err != nil {
		return err
	}
	return nil
}

// Check if the given profile exists in given array
func ContainsProfile(arr []types.List, profile string) bool {
	for _, n := range arr {
		if n.Name == profile {
			return true
		}
	}
	return false
}

func GetActiveProfile(list []types.List) string {
	for _, n := range list {
		if n.IsActive {
			return n.Name
		}
	}
	return ""
}
