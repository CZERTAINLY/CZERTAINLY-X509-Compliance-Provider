package rules

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

var logger *zap.Logger
var tService Service
var RULE_FILE_NAME = "../../rules.json"
var GROUP_FILE_NAME = "../../groups.json"

func init() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	tService = NewService(sugar, RULE_FILE_NAME, GROUP_FILE_NAME)
}

func TestGetRuleFromUuid(t *testing.T) {
	got := GetRuleFromUuid("40f084cc-ddc1-11ec-9d7f-34cff65c6ee3").Name
	want := "e_algorithm_identifier_improper_encoding"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestGetRuleUuidFromName(t *testing.T) {
	got := GetRuleUuidFromName("e_algorithm_identifier_improper_encoding")
	want := "40f084cc-ddc1-11ec-9d7f-34cff65c6ee3"
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestService_GetRules(t *testing.T) {
	rules, err := tService.GetRules("x509", []string{"X.509"})
	if err != nil {
		t.Error("Error", err)
	}
	got := len(rules)
	if got <= 0 {
		t.Errorf("Number of Rules: got %q", got)
	}
}

func TestService_GetGroups(t *testing.T) {
	groups, err := tService.GetGroups("x509")
	if err != nil {
		t.Error("Error", err)
	}
	got := len(groups)
	fmt.Println(got)
	if got <= 0 {
		t.Error("Number of Groups is 0")
	}
}

func TestService_GetGroupDetails(t *testing.T) {
	rules, err := tService.GetGroupDetails("5235104e-ddb2-11ec-9d64-0242ac120002", "x509")
	if err != nil {
		t.Error("Error", err)
	}
	got := len(rules)
	if got <= 0 {
		t.Error("Number of Groups is 0")
	}
}
