package actions

import (
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SecretSort string

const (
	SecretSortName       SecretSort = "Name"
	SecretSortAccessedAt SecretSort = "Accessed at"
	SecretSortCreatedAt  SecretSort = "Created at"
	SecretSortUpdatedAt  SecretSort = "Updated at"
	SecretSortRotatedAt  SecretSort = "Rotated at"
)

var SecretSortPossibleValues = []SecretSort{
	SecretSortName,
	SecretSortAccessedAt,
	SecretSortCreatedAt,
	SecretSortUpdatedAt,
	SecretSortRotatedAt,
}

func GetListSecrets() ([]*secretsmanager.SecretListEntry, error) {
	setting := GetAWSSetting()
	svc := secretsmanager.New(session.New(&aws.Config{
		Region: aws.String(setting.Region),
	}))
	maxResult := int64(100)

	index := 0
	var token *string
	var secrets []*secretsmanager.SecretListEntry
	for ; token != nil || index == 0; index++ {
		result, err := GetAPageSecrets(svc, token, maxResult)
		if err != nil {
			return nil, err
		}
		secrets = append(secrets, result.SecretList...)
		token = result.NextToken
	}
	SortSecrets(secrets)
	return secrets, nil
}

func GetAPageSecrets(svc *secretsmanager.SecretsManager, token *string, maxResult int64) (*secretsmanager.ListSecretsOutput, error) {
	input := &secretsmanager.ListSecretsInput{
		MaxResults: &maxResult,
		NextToken:  token,
	}
	result, err := svc.ListSecrets(input)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SortSecrets(secrets []*secretsmanager.SecretListEntry) {
	sort.Slice(secrets, func(i, j int) bool {
		if secrets[i] == nil || secrets[i].Name == nil {
			return false
		}
		if secrets[j] == nil || secrets[j].Name == nil {
			return true
		}
		return strings.ToLower(*secrets[i].Name) < strings.ToLower(*secrets[j].Name)
	})
}
