package rules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"go.uber.org/zap"
)

type service struct {
	logger *zap.SugaredLogger
}

var rulesDefinitions []RuleDefinition
var groupDefinitions []GroupDefinition
var uuidToRuleMap = map[string]RuleDefinition{}
var nameToUuidMap = map[string]string{}

func NewService(logger *zap.SugaredLogger, ruleFileName string, groupFileName string) Service {

	ruleFilePath, _ := filepath.Abs(ruleFileName)
	fmt.Println("Rule File Path: ", ruleFilePath)
	rulesFile, _ := ioutil.ReadFile(ruleFilePath)
	_ = json.Unmarshal(rulesFile, &rulesDefinitions)

	groupFilePath, _ := filepath.Abs(groupFileName)
	fmt.Println("Group File Path: ", groupFilePath)
	groupsFile, _ := ioutil.ReadFile(groupFilePath)
	_ = json.Unmarshal(groupsFile, &groupDefinitions)

	generateUuidToNameMap()
	return &service{
		logger: logger,
	}
}

func (s service) GetRules(kind string, certificateType []string) ([]Response, error) {
	s.logger.Info("Entering method: GetRules with Kind ", kind, " CertificateTypes ", certificateType)
	var filteredRules []Response
	for _, s := range rulesDefinitions {
		if s.Kind == kind {
			if len(certificateType) > 0 {
				if !strings.Contains(s.CertificateType, strings.Join(certificateType, "-")) {
					continue
				}
			}
			filteredRules = append(filteredRules, Response{
				UUID:            s.UUID,
				Name:            s.Name,
				CertificateType: s.CertificateType,
				Description:     s.Description,
				GroupUuid:       s.GroupUUID,
				Attributes:      s.Attributes,
			})
		}
	}
	s.logger.Info("Total of ", len(filteredRules), " found for the request")
	s.logger.Debug("Rules for the request ", filteredRules)
	return filteredRules, nil
}

func (s service) GetGroups(kind string) ([]GroupResponse, error) {
	s.logger.Info("Entering method: GetGroups with Kind ", kind)
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
	s.logger.Info("Total of ", len(filteredGroups), " found for the request")
	s.logger.Debug("Groups for the request ", filteredGroups)
	return filteredGroups, nil
}

func (s service) GetGroupDetails(uuid string, kind string) ([]Response, error) {
	s.logger.Info("Entering method: GetGroupDetails with Kind", kind, "Group UUID", uuid)
	var filteredRules []Response
	for _, s := range rulesDefinitions {
		if s.GroupUUID == uuid {
			filteredRules = append(filteredRules, Response{
				UUID:            s.UUID,
				Name:            s.Name,
				CertificateType: s.CertificateType,
				Description:     s.Description,
				GroupUuid:       s.GroupUUID,
				Attributes:      s.Attributes,
			})
		}
	}
	s.logger.Info("Total of ", len(filteredRules), " found for the request")
	s.logger.Debug("Rules for the request ", filteredRules)
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
