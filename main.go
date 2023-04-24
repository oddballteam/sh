package main

import (
	"fmt"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

var qs = []*survey.Question{

	// TODO: assuming web access since I will be putting all cli auth into aliases

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
			Options: []string{"wdsops-dev", "wdsops-prod", "ECWS-v3", "flh-dev", "flh-test", "learn-dev", "learn-prod", "mapi-dev", "mapi-test", "pc2-dev", "pc2-test", "scout-dev", "scout-impl", "scout-prod", "scout-test"},
			Default: "wdsops-dev",
		},
	},
}

func retrieveData(input string) (string, string) {
	accounts := map[string][2]string{
		"wdsops-dev":  {"879613780019", "wdsops-developer-admin"},
		"ECWS-v3":     {"521784486762", "learn-application-admin"},
		"flh-dev":     {"405622158906", "flh-developer-admin"},
		"flh-test":    {"504465124926", "flh-developer-admin"},
		"learn-dev":   {"329043178936", "learn-application-admin"},
		"learn-prod":  {"840218319247", "learn-application-admin"},
		"mapi-dev":    {"204718834091", "mapi-developer-admin"},
		"mapi-test":   {"395677180642", "mapi-developer-admin"},
		"pc2-dev":     {"155987839255", "pc2-developer-admin"},
		"pc2-test":    {"115900149718", "pc2-developer-admin"},
		"scout-dev":   {"198854277681", "scout-application-admin"},
		"scout-impl":  {"394953402894", "scout-application-admin"},
		"scout-prod":  {"263813120251", "scout-application-admin"},
		"scout-test":  {"001265600887", "scout-application-admin"},
		"wdsops-prod": {"251875634914", "wdsops-application-admin"},
	}
	return accounts[input][0], accounts[input][1]
}

type Data struct {
	Key    string `json:"access_key"`
	Secret string `json:"secret_access_key"`
	Token  string `json:"session_token"`
}

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
		fmt.Println(fmt.Sprintf("Run the following command:\n\nsource <(kion credentials --account-id %s --cloud-access-role %s --format export)", id, role))
		// cmd := exec.Command("kion", "credentials", "--account-id", id, "--cloud-access-role", role, "--format", "export")
		// out, err := cmd.Output()
		// if err != nil {
		// 	fmt.Println("could not run command: ", err)
		// }
		// fmt.Println(string(out))
		// var data Data
		// err2 := json.Unmarshal(out, &data)
		// if err2 != nil {
		// 	fmt.Println("could not interpret json: ", err2)
		// }
		// os.Setenv("AWS_SECRET_ACCESS_KEY", data.Key)
		// os.Setenv("AWS_SECRET_ACCESS_KEY", data.Secret)
		// os.Setenv("AWS_SESSION_TOKEN", data.Token)

		// fmt.Println(os.Getenv("AWS_SECRET_ACCESS_KEY"))
	}

}
