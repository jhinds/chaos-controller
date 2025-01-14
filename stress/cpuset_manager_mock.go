// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2022 Datadog, Inc.

package stress

import (
	"github.com/DataDog/chaos-controller/cgroup"
	"github.com/stretchr/testify/mock"
	"k8s.io/apimachinery/pkg/util/intstr"
)
import (
	"github.com/DataDog/chaos-controller/cpuset"
)

type StresserManagerMock struct {
	mock.Mock
}

//nolint:golint
func (f *StresserManagerMock) IsCoreAlreadyStressed(core int) bool {
	args := f.Called(core)
	return args.Bool(0)
}

//nolint:golint
func (f *StresserManagerMock) TrackCoreAlreadyStressed(core int, stresserPID int) {
	f.Called(core, stresserPID)
}

//nolint:golint
func (f *StresserManagerMock) StresserPIDs() map[int]int {
	args := f.Called()
	return args.Get(0).(map[int]int)
}

//nolint:golint
func (f *StresserManagerMock) TrackInjectorCores(cgroup cgroup.Manager, userRequestCount *intstr.IntOrString) (cpuset.CPUSet, error) {
	args := f.Called(cgroup, userRequestCount)
	return args.Get(0).(cpuset.CPUSet), args.Error(1)
}
