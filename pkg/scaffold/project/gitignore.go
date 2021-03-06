/*
Copyright 2018 The Kubernetes Authors.

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

package project

import (
	"sigs.k8s.io/kubebuilder/pkg/model/file"
)

var _ file.Template = &GitIgnore{}

// GitIgnore scaffolds the .gitignore file
type GitIgnore struct {
	file.Input
}

// GetInput implements input.Template
func (f *GitIgnore) GetInput() (file.Input, error) {
	if f.Path == "" {
		f.Path = ".gitignore"
	}
	f.TemplateBody = gitignoreTemplate
	return f.Input, nil
}

const gitignoreTemplate = `
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
bin

# Test binary, build with ` + "`go test -c`" + `
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Kubernetes Generated files - skip generated files, except for vendored files

!vendor/**/zz_generated.*

# editor and IDE paraphernalia
.idea
*.swp
*.swo
*~
`
