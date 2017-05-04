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

// Badge is the type for a badge.
type Badge struct {
	// Location is the URL for this badge.
	// For example, https://travis-ci.org/gaocegege/maintainer.svg?branch=master
	Location Link
	// Refer is the URL which the badge is linked to.
	// For example, https://travis-ci.org/gaocegege/maintainer
	Refer Link
}

// Link is the type for a link.
type Link struct {
	// Name is the name of this link.
	Name string
	// URL is the URL for the link.
	URL string
}
