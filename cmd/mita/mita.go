// Copyright (C) 2021  mieru authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

// Binary mita is the server of mieru proxy.
package main

import (
	"github.com/enfein/mieru/pkg/appctl"
	"github.com/enfein/mieru/pkg/cli"
	"github.com/enfein/mieru/pkg/log"
	"github.com/enfein/mieru/pkg/rng"
)

func main() {
	appctl.SetAppType(appctl.SERVER_APP)
	rng.InitSeed()
	cli.RegisterServerCommands()
	err := cli.ParseAndExecute()
	if err != nil {
		log.Fatalf("%v", err)
	}
}