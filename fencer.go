// Copyright 2012 Darren Elwood <darren@textnode.com>
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

// This Go code acts as a fallback for processor architectures other than x86.
//

// +build 386,!gccgo amd64,!gccgo

package fencer

var lfence func()
var mfence func()
var sfence func()

func init() {
	initCPU()
}

func LFence() {
	lfence()
}
func MFence() {
	mfence()
}
func SFence() {
	sfence()
}
