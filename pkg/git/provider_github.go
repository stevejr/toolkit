/*
Copyright 2020 The Flux CD contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package git

import (
	"context"
	"fmt"
	"github.com/google/go-github/v32/github"
	"strings"
)

// GithubProvider represents a GitHub API wrapper
type GithubProvider struct {
	IsPrivate  bool
	IsPersonal bool
}

const (
	GitHubTokenName       = "GITHUB_TOKEN"
	GitHubDefaultHostname = "github.com"
)

func (p *GithubProvider) newClient(r *Repository) (*github.Client, error) {
	auth := github.BasicAuthTransport{
		Username: "git",
		Password: r.Token,
	}

	gh := github.NewClient(auth.Client())
	if r.Host != GitHubDefaultHostname {
		baseURL := fmt.Sprintf("https://%s/api/v3/", r.Host)
		uploadURL := fmt.Sprintf("https://%s/api/uploads/", r.Host)
		if g, err := github.NewEnterpriseClient(baseURL, uploadURL, auth.Client()); err == nil {
			gh = g
		} else {
			return nil, err
		}
	}

	return gh, nil
}

// CreateRepository returns false if the repository exists
func (p *GithubProvider) CreateRepository(ctx context.Context, r *Repository) (bool, error) {
	gh, err := p.newClient(r)
	if err != nil {
		return false, fmt.Errorf("client error: %w", err)
	}
	org := ""
	if !p.IsPersonal {
		org = r.Owner
	}

	if _, _, err := gh.Repositories.Get(ctx, org, r.Name); err == nil {
		return false, nil
	}

	autoInit := true
	_, _, err = gh.Repositories.Create(ctx, org, &github.Repository{
		AutoInit: &autoInit,
		Name:     &r.Name,
		Private:  &p.IsPrivate,
	})
	if err != nil {
		if !strings.Contains(err.Error(), "name already exists on this account") {
			return false, fmt.Errorf("failed to create repository, error: %w", err)
		}
	} else {
		return true, nil
	}
	return false, nil
}

// AddTeam returns false if the team is already assigned to the repository
func (p *GithubProvider) AddTeam(ctx context.Context, r *Repository, name, permission string) (bool, error) {
	gh, err := p.newClient(r)
	if err != nil {
		return false, fmt.Errorf("client error: %w", err)
	}

	// check team exists
	_, _, err = gh.Teams.GetTeamBySlug(ctx, r.Owner, name)
	if err != nil {
		return false, fmt.Errorf("failed to retrieve team '%s', error: %w", name, err)
	}

	// check if team is assigned to the repo
	_, resp, err := gh.Teams.IsTeamRepoBySlug(ctx, r.Owner, name, r.Owner, r.Name)
	if resp == nil && err != nil {
		return false, fmt.Errorf("failed to determine if team '%s' is assigned to the repository, error: %w", name, err)
	}

	// add team to the repo
	if resp.StatusCode == 404 {
		_, err = gh.Teams.AddTeamRepoBySlug(ctx, r.Owner, name, r.Owner, r.Name, &github.TeamAddTeamRepoOptions{
			Permission: permission,
		})
		if err != nil {
			return false, fmt.Errorf("failed to add team '%s' to the repository, error: %w", name, err)
		}
		return true, nil
	}

	return false, nil
}

// AddDeployKey returns false if the key exists and the content is the same
func (p *GithubProvider) AddDeployKey(ctx context.Context, r *Repository, key, keyName string) (bool, error) {
	gh, err := p.newClient(r)
	if err != nil {
		return false, fmt.Errorf("client error: %w", err)
	}

	// list deploy keys
	keys, resp, err := gh.Repositories.ListKeys(ctx, r.Owner, r.Name, nil)
	if err != nil {
		return false, fmt.Errorf("failed to list deploy keys, error: %w", err)
	}
	if resp.StatusCode >= 300 {
		return false, fmt.Errorf("failed to list deploy keys (status code: %s)", resp.Status)
	}

	// check if the key exists
	shouldCreateKey := true
	var existingKey *github.Key
	for _, k := range keys {
		if k.Title != nil && k.Key != nil && *k.Title == keyName {
			if *k.Key != key {
				existingKey = k
			} else {
				shouldCreateKey = false
			}
			break
		}
	}

	// delete existing key if the value differs
	if existingKey != nil {
		resp, err := gh.Repositories.DeleteKey(ctx, r.Owner, r.Name, *existingKey.ID)
		if err != nil {
			return false, fmt.Errorf("failed to delete deploy key '%s', error: %w", keyName, err)
		}
		if resp.StatusCode >= 300 {
			return false, fmt.Errorf("failed to delete deploy key '%s' (status code: %s)", keyName, resp.Status)
		}
	}

	// create key
	if shouldCreateKey {
		isReadOnly := true
		_, _, err = gh.Repositories.CreateKey(ctx, r.Owner, r.Name, &github.Key{
			Title:    &keyName,
			Key:      &key,
			ReadOnly: &isReadOnly,
		})
		if err != nil {
			return false, fmt.Errorf("failed to create deploy key '%s', error: %w", keyName, err)
		}
		return true, nil
	}

	return false, nil
}
