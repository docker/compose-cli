/*
   Copyright 2020 Docker Compose CLI authors

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

package etchosts

import (
	"os"
	"path/filepath"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
	"gotest.tools/v3/golden"
)

func TestSetDomain(t *testing.T) {
	dir := fs.NewDir(t, "resolv").Path()
	f := filepath.Join(dir, "hosts")
	touch(t, f)

	err := SetHostNames(f, "foo", "bar", "zot")
	assert.NilError(t, err)

	got, err := os.ReadFile(f)
	assert.NilError(t, err)
	golden.Assert(t, string(got), "etchosts.golden")
}

func touch(t *testing.T, f string) {
	file, err := os.Create(f)
	assert.NilError(t, err)
	err = file.Close()
	assert.NilError(t, err)
}
