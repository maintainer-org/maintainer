// Copyright © 2017 Maintainer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package badge

import (
	"log"

	"github.com/gaocegege/maintainer/repo"
	"github.com/gaocegege/maintainer/util"

	"github.com/google/go-github/github"
)

// Recommender is the type to recommend badges.
type Recommender struct {
	Repo *github.Repository
}

// NewRecommender returns a recommender.
func NewRecommender() (*Recommender, error) {
	repo, err := repo.NewRepository()
	if err != nil {
		log.Panicf("Error when read the information from local repository: %s\n", err)
	}
	client := github.NewClient(nil)
	remoteRepo, _, err := client.Repositories.Get(repo.Owner, repo.Name)
	if err != nil {
		return nil, err
	}

	return &Recommender{
		Repo: remoteRepo,
	}, nil
}

// Recommend recommends the badges.
func (b *Recommender) Recommend() error {
	if err := b.recommendLanguageSpecificBadges(); err != nil {
		return err
	}
	return nil
}

func (b *Recommender) recommendLanguageSpecificBadges() error {
	switch *b.Repo.Language {
	case util.LanguageGo:
		return util.ErrNotImplemented
	case util.LanguageJava:
		return util.ErrNotImplemented
	case util.LanguageJavaScript:
		return util.ErrNotImplemented
	case util.LanguageScala:
		return util.ErrNotImplemented
	case util.LanguageShell:
		return util.ErrNotImplemented
	case util.LanguageCpp:
		return util.ErrNotImplemented
	case util.LanguagePython:
		return util.ErrNotImplemented
	case util.LanguagePHP:
		return util.ErrNotImplemented
	case util.LanguageRuby:
		return util.ErrNotImplemented
	case util.LanguageSwift:
		return util.ErrNotImplemented
	case util.LanguageR:
		return util.ErrNotImplemented
	default:
		log.Printf("%s is not supported now to generate coding style guide.", b.Repo.Language)
		return util.ErrNotSupported
	}
}

func (b *Recommender) badgesGo() {

}
