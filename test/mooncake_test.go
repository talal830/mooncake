package test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mooncake/executor"
	"testing"
)

func TestBasicRules(t *testing.T) {
	content, err := ioutil.ReadFile("../test/rules_v1.1.mck")

	if err != nil {
		panic(err)
	}

	rules := string(content)
	jsonFile := "../sample/sample.json"

	// sample context struct
	ctx := struct {
		name      string
		threshold int
		tempValue int
	}{
		"xyz",
		2,
		500,
	}

	result := executor.Execute(rules, jsonFile, ctx)

	prettyResult, err := json.MarshalIndent(result, "", "  ");
	log.Printf("\n\n====Result===== \n\n %s\n\n", prettyResult)
}
