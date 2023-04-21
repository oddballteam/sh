package main

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var qs = []*survey.Question{
	// {
	// 	Name: "access",
	// 	Prompt: &survey.Select{
	// 		Message: "Would you like to Web or CLI access:",
	// 		Options: []string{"Web", "CLI"},
	// 		Default: "CLI",
	// 	},
	// },
	{
		Name: "account",
		Prompt: &survey.Select{
			Message: "What account would you like access for:",
			Options: []string{"wdsops-dev"},
			Default: "wdsops-dev",
		},
	},
}

func retrieveData(input string) (string, string) {
	accounts := map[string][2]string{
		// "flh-dev":    {"405622158906", "flh-developer-admin"},
		// "learn-dev":  {"329043178936", "learn-application-admin"},
		// "learn-prod": {"840218319247", "learn-application-admin"},
		"wdsops-dev": {"879613780019", "wdsops-developer-admin"},
	}
	return accounts[input][0], accounts[input][1]
}

// type Data struct {
// 	Key    string `json:"access_key"`
// 	Secret string `json:"secret_access_key"`
// 	Token  string `json:"session_token"`
// }

func main() {
	answers := struct {
		Access  string
		Account string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	id, role := retrieveData(answers.Account)
	// answers.Access == "web"
	if true {
		cmd := exec.Command("kion", "console", "--account-id", id, "--cloud-access-role", role)
		_, err := cmd.Output()
		if err != nil {
			fmt.Println("could not run command: ", err)
		}
	} else {
		fmt.Println(fmt.Sprintf("source <(kion credentials --account-id %s --cloud-access-role %s --format export)", id, role))
		// source <(kion credentials --account-id 879613780019 --cloud-access-role wdsops-developer-admin --format export)
		// cmd := exec.Command("kion", "credentials", "--account-id", id, "--cloud-access-role", role, "--format", "export")
		// out, err := cmd.Output()
		// if err != nil {
		// 	fmt.Println("could not run command: ", err)
		// }
		// fmt.Println(string(out))
	}

	// var data Data
	// err2 := json.Unmarshal(out, &data)
	// if err2 != nil {
	// 	fmt.Println("could not interpret json: ", err2)
	// }
	// os.Setenv("AWS_SECRET_ACCESS_KEY", data.Key)
	// os.Setenv("AWS_SECRET_ACCESS_KEY", data.Secret)
	// os.Setenv("AWS_SESSION_TOKEN", data.Token)

}
