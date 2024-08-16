// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package projectconfig

import (
	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server/builds"
	"github.com/daytonaio/daytona/pkg/server/gitproviders"
	"github.com/daytonaio/daytona/pkg/server/projectconfig/dto"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
)

type IProjectConfigService interface {
	Save(projectConfig *config.ProjectConfig) error
	Find(filter *config.Filter) (*config.ProjectConfig, error)
	List(filter *config.Filter) ([]*config.ProjectConfig, error)
	SetDefault(projectConfigName string) error
	Delete(projectConfigName string) error
	SetPrebuild(dto.CreatePrebuildDTO) (*dto.PrebuildDTO, error)
	FindPrebuild(projectConfigFilter *config.Filter, prebuildFilter *config.PrebuildFilter) (*dto.PrebuildDTO, error)
	ListPrebuilds(projectConfigFilter *config.Filter, prebuildFilter *config.PrebuildFilter) ([]*dto.PrebuildDTO, error)
	DeletePrebuild(projectConfigName string, id string) error
	ProcessGitEvent(gitprovider.GitEventData) error
}

type ProjectConfigServiceConfig struct {
	PrebuildWebhookEndpoint string
	ConfigStore             config.Store
	BuildService            builds.IBuildService
	GitProviderService      gitproviders.IGitProviderService
}

type ProjectConfigService struct {
	prebuildWebhookEndpoint string
	configStore             config.Store
	buildService            builds.IBuildService
	gitProviderService      gitproviders.IGitProviderService
}

func NewProjectConfigService(config ProjectConfigServiceConfig) IProjectConfigService {
	return &ProjectConfigService{
		prebuildWebhookEndpoint: config.PrebuildWebhookEndpoint,
		configStore:             config.ConfigStore,
		buildService:            config.BuildService,
		gitProviderService:      config.GitProviderService,
	}
}

func (s *ProjectConfigService) List(filter *config.Filter) ([]*config.ProjectConfig, error) {
	return s.configStore.List(filter)
}

func (s *ProjectConfigService) SetDefault(projectConfigName string) error {
	projectConfig, err := s.Find(&config.Filter{
		Name: &projectConfigName,
	})
	if err != nil {
		return err
	}

	defaultProjectConfig, err := s.Find(&config.Filter{
		Url:     &projectConfig.RepositoryUrl,
		Default: util.Pointer(true),
	})
	if err != nil && err != config.ErrProjectConfigNotFound {
		return err
	}

	if defaultProjectConfig != nil {
		defaultProjectConfig.IsDefault = false
		err := s.configStore.Save(defaultProjectConfig)
		if err != nil {
			return err
		}
	}

	projectConfig.IsDefault = true
	return s.configStore.Save(projectConfig)
}

func (s *ProjectConfigService) Find(filter *config.Filter) (*config.ProjectConfig, error) {
	return s.configStore.Find(filter)
}

func (s *ProjectConfigService) Save(projectConfig *config.ProjectConfig) error {
	projectConfig.RepositoryUrl = util.CleanUpRepositoryUrl(projectConfig.RepositoryUrl)

	err := s.configStore.Save(projectConfig)
	if err != nil {
		return err
	}

	return s.SetDefault(projectConfig.Name)
}

func (s *ProjectConfigService) Delete(projectConfigName string) error {
	pc, err := s.Find(&config.Filter{
		Name: &projectConfigName,
	})
	if err != nil {
		return err
	}
	return s.configStore.Delete(pc)
}
