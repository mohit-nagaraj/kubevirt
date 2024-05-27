/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright The KubeVirt Authors.
 *
 */

package featuregate

const (
	ExpandDisksGate       = "ExpandDisks"
	CPUManager            = "CPUManager"
	IgnitionGate          = "ExperimentalIgnitionSupport"
	HypervStrictCheckGate = "HypervStrictCheck"
	SidecarGate           = "Sidecar"
	HostDevicesGate       = "HostDevices"
	SnapshotGate          = "Snapshot"
	VMExportGate          = "VMExport"
	HotplugVolumesGate    = "HotplugVolumes"
	HostDiskGate          = "HostDisk"

	DownwardMetricsFeatureGate = "DownwardMetrics"
	Root                       = "Root"
	ClusterProfiler            = "ClusterProfiler"
	WorkloadEncryptionSEV      = "WorkloadEncryptionSEV"
	VSOCKGate                  = "VSOCK"
	// KubevirtSeccompProfile indicate that Kubevirt will install its custom profile and
	// user can tell Kubevirt to use it
	KubevirtSeccompProfile = "KubevirtSeccompProfile"
	// DisableMediatedDevicesHandling disables the handling of mediated
	// devices, its creation and deletion
	DisableMediatedDevicesHandling = "DisableMDEVConfiguration"
	// PersistentReservation enables the use of the SCSI persistent reservation with the pr-helper daemon
	PersistentReservation = "PersistentReservation"
	// VMPersistentState enables persisting backend state files of VMs, such as the contents of the vTPM
	VMPersistentState = "VMPersistentState"
	MultiArchitecture = "MultiArchitecture"
	// AutoResourceLimitsGate enables automatic setting of vmi limits if there is a ResourceQuota with limits associated with the vmi namespace.
	AutoResourceLimitsGate = "AutoResourceLimitsGate"

	// AlignCPUsGate allows emulator thread to assign two extra CPUs if needed to complete even parity.
	AlignCPUsGate = "AlignCPUs"

	// Owner: @xpivarc
	// Alpha: v1.3.0
	//
	// NodeRestriction enables Kubelet's like NodeRestriction but for Kubevirt's virt-handler.
	// This feature requires following Kubernetes feature gate "ServiceAccountTokenPodNodeInfo". The feature gate is available
	// in Kubernetes 1.30 as Beta.
	NodeRestrictionGate = "NodeRestriction"
	// Owner: @lyarwood
	// Alpha: v1.4.0
	//
	// InstancetypeReferencePolicy allows a cluster admin to control how a VirtualMachine references instance types and preferences
	// through the kv.spec.configuration.instancetype.referencePolicy configurable.
	InstancetypeReferencePolicy = "InstancetypeReferencePolicy"

	VirtIOFSConfigVolumesGate = "EnableVirtioFsConfigVolumes"
	VirtIOFSStorageVolumeGate = "EnableVirtioFsStorageVolumes"
)
