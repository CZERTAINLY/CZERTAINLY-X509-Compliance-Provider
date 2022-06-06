package rules

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"strings"
)

type service struct {
	logger log.Logger
}

var rulesDefinitions []RuleDefinition
var groupDefinitions []GroupDefinition
var uuidToRuleMap = map[string]RuleDefinition{}
var nameToUuidMap = map[string]string{}

func init() {
	rulesFile, _ := ioutil.ReadFile("cmd/rules/rules.json")
	_ = json.Unmarshal(rulesFile, &rulesDefinitions)

	groupsFile, _ := ioutil.ReadFile("cmd/rules/groups.json")
	_ = json.Unmarshal(groupsFile, &groupDefinitions)

	generateUuidToNameMap()

}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s service) GetRules(ctx context.Context, kind string, certificateType []string) ([]Response, error) {
	var filteredRules []Response
	for _, s := range rulesDefinitions {
		if s.Kind == kind {
			if len(certificateType) > 0 {
				if !strings.Contains(s.CertificateType, strings.Join(certificateType, "-")) {
					continue
				}
			}
			filteredRules = append(filteredRules, Response{
				UUID:        s.UUID,
				Name:        s.Name,
				Description: s.Description,
				Attributes:  s.Attributes,
			})
		}
	}
	return filteredRules, nil
}

func (s service) GetGroups(ctx context.Context, kind string) ([]GroupResponse, error) {
	var filteredGroups []GroupResponse
	for _, s := range groupDefinitions {
		if s.Kind == kind {
			filteredGroups = append(filteredGroups, GroupResponse{
				UUID:        s.UUID,
				Name:        s.Name,
				Description: s.Description,
			})
		}
	}
	return filteredGroups, nil
}

func (s service) GetGroupDetails(ctx context.Context, uuid string, kind string) ([]Response, error) {
	var filteredRules []Response
	for _, s := range rulesDefinitions {
		if s.GroupUUID == uuid {
			filteredRules = append(filteredRules, Response{
				UUID:        s.UUID,
				Name:        s.Name,
				Description: s.Description,
				Attributes:  s.Attributes,
			})
		}
	}
	return filteredRules, nil
}

func generateUuidToNameMap() {
	for _, s := range rulesDefinitions {
		uuidToRuleMap[s.UUID] = s
	}

	for _, s := range rulesDefinitions {
		nameToUuidMap[s.Name] = s.UUID
	}
}

func GetRuleFromUuid(uuid string) (name RuleDefinition) {
	return uuidToRuleMap[uuid]
}

func GetRuleUuidFromName(name string) (uuid string) {
	return nameToUuidMap[name]
}
