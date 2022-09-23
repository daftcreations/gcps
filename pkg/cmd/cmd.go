package cmd

import (
	"strings"

	"github.com/daftcreations/gcps/pkg/types"
	"github.com/manifoldco/promptui"
)

// GetProfileFromUser prompts user to select desires profile from provided Google Cloud profiles list.
func GetProfileFromUser(list []types.List) (string, error) {
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001f449 {{ .Name | cyan }} ({{ .IsActive | red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .IsActive | red }})",
		Selected: "\u2705 {{ .Name | green }}",
	}

	prompt := promptui.Select{
		Label:     "Select Profile",
		Items:     list,
		Size:      10,
		Templates: templates,
		HideHelp:  true,
	}

	_, result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return strings.Split(strings.Trim(result, "{}"), " ")[0], nil
}
