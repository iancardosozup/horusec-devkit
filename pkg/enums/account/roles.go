// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package account

//Role is the role applicable to an account
type Role string

const (
	ApplicationAdmin Role = "applicationAdmin"
	Admin            Role = "admin"
	Member           Role = "member"
	Supervisor       Role = "supervisor"
)

//IsValid checks if a given Role is in possible Values slice
func (r Role) IsValid() bool {
	for _, v := range r.Values() {
		if v == r {
			return true
		}
	}

	return false
}

//Values returns a slice of possible Role values
func (r Role) Values() []Role {
	return []Role{
		ApplicationAdmin,
		Admin,
		Member,
		Supervisor,
	}
}
