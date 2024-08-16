// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package config

import (
	"encoding/json"
	"sort"

	"github.com/daytonaio/daytona/pkg/workspace/project/buildconfig"
)

type ProjectConfig struct {
	Name          string                   `json:"name" validate:"required"`
	Image         string                   `json:"image" validate:"required"`
	User          string                   `json:"user" validate:"required"`
	BuildConfig   *buildconfig.BuildConfig `json:"buildConfig,omitempty" validate:"optional"`
	RepositoryUrl string                   `json:"repositoryUrl" validate:"required"`
	EnvVars       map[string]string        `json:"envVars" validate:"required"`
	IsDefault     bool                     `json:"default" validate:"required"`
	Prebuilds     []*PrebuildConfig        `json:"prebuilds" validate:"optional"`
	WebhookId     *string                  `json:"webhookId,omitempty"`
} // @name ProjectConfig

func (pc *ProjectConfig) SetPrebuild(p *PrebuildConfig) error {
	newPrebuild := PrebuildConfig{
		Id:             p.Id,
		Branch:         p.Branch,
		CommitInterval: p.CommitInterval,
		TriggerFiles:   p.TriggerFiles,
	}

	for _, pb := range pc.Prebuilds {
		if pb.Id == p.Id {
			pb = &newPrebuild
			return nil
		}
	}

	pc.Prebuilds = append(pc.Prebuilds, &newPrebuild)
	return nil
}

func (pc *ProjectConfig) FindPrebuild(filter *PrebuildFilter) (*PrebuildConfig, error) {
	for _, pb := range pc.Prebuilds {
		filteredPrebuild := filterPrebuild(pb, filter)
		if filteredPrebuild != nil {
			return filteredPrebuild, nil
		}
	}

	return nil, nil
}

func (pc *ProjectConfig) ListPrebuilds(filter *PrebuildFilter) ([]*PrebuildConfig, error) {
	if filter == nil {
		return pc.Prebuilds, nil
	}

	prebuilds := []*PrebuildConfig{}

	for _, pb := range pc.Prebuilds {
		filteredPrebuild := filterPrebuild(pb, filter)
		if filteredPrebuild != nil {
			prebuilds = append(prebuilds, filteredPrebuild)
		}
	}

	return prebuilds, nil
}

func (pc *ProjectConfig) RemovePrebuild(id string) error {
	newPrebuilds := []*PrebuildConfig{}

	for _, pb := range pc.Prebuilds {
		if pb.Id != id {
			newPrebuilds = append(newPrebuilds, pb)
		}
	}

	pc.Prebuilds = newPrebuilds
	return nil
}

func filterPrebuild(pb *PrebuildConfig, filter *PrebuildFilter) *PrebuildConfig {
	if filter.Id != nil && *filter.Id != pb.Id {
		return nil
	}

	if filter.Branch != nil && *filter.Branch != pb.Branch {
		return nil
	}

	if filter.CommitInterval != nil && *filter.CommitInterval != pb.CommitInterval {
		return nil
	}

	if filter.TriggerFiles != nil {
		// Sort the trigger files before checking if same
		sort.Strings(pb.TriggerFiles)
		sort.Strings(*filter.TriggerFiles)
		triggerFilesJson, err := json.Marshal(pb.TriggerFiles)
		if err != nil {
			return nil
		}
		filterFilesJson, err := json.Marshal(*filter.TriggerFiles)
		if err != nil {
			return nil
		}
		if string(triggerFilesJson) != string(filterFilesJson) {
			return nil
		}

	}

	return pb
}
