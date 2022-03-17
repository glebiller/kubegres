/*
Copyright 2021 Reactive Tech Limited.
"Reactive Tech Limited" is a company located in England, United Kingdom.
https://www.reactive-tech.io

Lead Developer: Alex Arica

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

package statefulset

type SortByInstance []StatefulSetWrapper
type ReverseSortByInstance []StatefulSetWrapper

func (f SortByInstance) Len() int {
	return len(f)
}

func (f SortByInstance) Less(i, j int) bool {
	return f[i].Instance() < f[j].Instance()
}

func (f SortByInstance) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}

func (f ReverseSortByInstance) Len() int {
	return len(f)
}

func (f ReverseSortByInstance) Less(i, j int) bool {
	return f[i].Instance() > f[j].Instance()
}

func (f ReverseSortByInstance) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
