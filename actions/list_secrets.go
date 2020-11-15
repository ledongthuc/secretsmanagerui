package actions

import (
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

type SortType string

const (
	SortTypeAsc  SortType = "ASC"
	SortTypeDesc SortType = "DESC"
)

type SecretSort string

const (
	SecretSortName       SecretSort = "Name"
	SecretSortAccessedAt SecretSort = "Accessed at"
	SecretSortCreatedAt  SecretSort = "Created at"
	SecretSortUpdatedAt  SecretSort = "Updated at"
	SecretSortRotatedAt  SecretSort = "Rotated at"
)

type SecretSortOption struct {
	Sort        SecretSort
	SortType    SortType
	NiceText    string
	Description string
}

var (
	SecretSortNameAsc        = SecretSortOption{SecretSortName, SortTypeAsc, "Name A->Z", "Alphabet"}
	SecretSortNameDesc       = SecretSortOption{SecretSortName, SortTypeDesc, "Name Z->A", "Alphabet"}
	SecretSortAccessedAtAsc  = SecretSortOption{SecretSortAccessedAt, SortTypeAsc, "Show recent 'Accessed at'", "Newest to oldest"}
	SecretSortAccessedAtDesc = SecretSortOption{SecretSortAccessedAt, SortTypeDesc, "Show oldest 'Accessed at'", "Oldest to newest"}
	SecretSortCreatedAtAsc   = SecretSortOption{SecretSortCreatedAt, SortTypeAsc, "Show recent 'Created at'", "Newest to oldest"}
	SecretSortCreatedAtDesc  = SecretSortOption{SecretSortCreatedAt, SortTypeDesc, "Show oldest 'Created at'", "Oldest to newest"}
	SecretSortUpdatedAtAsc   = SecretSortOption{SecretSortUpdatedAt, SortTypeAsc, "Show recent 'Updated at'", "Newest to oldest"}
	SecretSortUpdatedAtDesc  = SecretSortOption{SecretSortUpdatedAt, SortTypeDesc, "Show oldest 'Updated at'", "Oldest to newest"}
	SecretSortRotatedAtAsc   = SecretSortOption{SecretSortRotatedAt, SortTypeAsc, "Show recent 'Rotated at'", "Newest to oldest"}
	SecretSortRotatedAtDesc  = SecretSortOption{SecretSortRotatedAt, SortTypeDesc, "Show oldest 'Rotated at'", "Oldest to newest"}
)

var SecretSortPossibleValues = []SecretSortOption{
	SecretSortNameAsc,
	SecretSortNameDesc,
	SecretSortAccessedAtAsc,
	SecretSortAccessedAtDesc,
	SecretSortCreatedAtAsc,
	SecretSortCreatedAtDesc,
	SecretSortUpdatedAtAsc,
	SecretSortUpdatedAtDesc,
	SecretSortRotatedAtAsc,
	SecretSortRotatedAtDesc,
}

func GetListSecrets(sort SecretSortOption) ([]*secretsmanager.SecretListEntry, error) {
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
	SortSecrets(secrets, sort)
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

func SortSecrets(secrets []*secretsmanager.SecretListEntry, sortOption SecretSortOption) {
	if sortOption.Sort == SecretSortName {
		sort.Slice(secrets, func(i, j int) bool {
			if secrets[i] == nil || secrets[i].Name == nil {
				return false
			}
			if secrets[j] == nil || secrets[j].Name == nil {
				return true
			}
			if sortOption.SortType == SortTypeAsc {
				return strings.ToLower(*secrets[i].Name) < strings.ToLower(*secrets[j].Name)
			}
			return strings.ToLower(*secrets[i].Name) >= strings.ToLower(*secrets[j].Name)
		})
	}
	if sortOption.Sort == SecretSortAccessedAt {
		sort.Slice(secrets, func(i, j int) bool {
			if secrets[i] == nil || secrets[i].LastAccessedDate == nil {
				return false
			}
			if secrets[j] == nil || secrets[j].LastAccessedDate == nil {
				return true
			}
			if sortOption.SortType == SortTypeAsc {
				return secrets[i].LastAccessedDate.Before(*secrets[j].LastAccessedDate)
			}
			return secrets[i].LastAccessedDate.Equal(*secrets[j].LastAccessedDate) || secrets[i].LastAccessedDate.After(*secrets[j].LastAccessedDate)
		})
	}
	if sortOption.Sort == SecretSortCreatedAt {
		sort.Slice(secrets, func(i, j int) bool {
			if secrets[i] == nil || secrets[i].CreatedDate == nil {
				return false
			}
			if secrets[j] == nil || secrets[j].CreatedDate == nil {
				return true
			}
			if sortOption.SortType == SortTypeAsc {
				return secrets[i].CreatedDate.Before(*secrets[j].CreatedDate)
			}
			return secrets[i].CreatedDate.Equal(*secrets[j].CreatedDate) || secrets[i].CreatedDate.After(*secrets[j].CreatedDate)
		})
	}
	if sortOption.Sort == SecretSortUpdatedAt {
		sort.Slice(secrets, func(i, j int) bool {
			if secrets[i] == nil || secrets[i].LastChangedDate == nil {
				return false
			}
			if secrets[j] == nil || secrets[j].LastChangedDate == nil {
				return true
			}
			if sortOption.SortType == SortTypeAsc {
				return secrets[i].LastChangedDate.Before(*secrets[j].LastChangedDate)
			}
			return secrets[i].LastChangedDate.Equal(*secrets[j].LastChangedDate) || secrets[i].LastChangedDate.After(*secrets[j].LastChangedDate)
		})
	}
	if sortOption.Sort == SecretSortRotatedAt {
		sort.Slice(secrets, func(i, j int) bool {
			if secrets[i] == nil || secrets[i].LastRotatedDate == nil {
				return false
			}
			if secrets[j] == nil || secrets[j].LastRotatedDate == nil {
				return true
			}
			if sortOption.SortType == SortTypeAsc {
				return secrets[i].LastRotatedDate.Before(*secrets[j].LastRotatedDate)
			}
			return secrets[i].LastRotatedDate.Equal(*secrets[j].LastRotatedDate) || secrets[i].LastRotatedDate.After(*secrets[j].LastRotatedDate)
		})
	}
}
