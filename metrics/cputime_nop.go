// Copyright 2018 The go-embreum  Authors
// This file is part of the go-embreum  library.
//
// The go-embreum  library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-embreum  library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-embreum  library. If not, see <http://www.gnu.org/licenses/>.

//go:build windows || js
// +build windows js

package metrics

// getProcessCPUTime returns 0 on Windows as there is no system call to resolve
// the actual process' CPU time.
func getProcessCPUTime() int64 {
	return 0
}
