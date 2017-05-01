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

package dep

// Dependency is the type for the dependency.
type Dependency interface {
	// IsInstalled check where the dependency is installed.
	IsInstalled() (bool, error)
	// Install installs the specific dependency and returns error if fails.
	Install() error
	// Name return the name of the dependency.
	Name() string
}
