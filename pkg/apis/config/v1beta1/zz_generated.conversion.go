// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1beta1

import (
	unsafe "unsafe"

	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	configv1 "k8s.io/kube-scheduler/config/v1"
	config "sigs.k8s.io/scheduler-plugins/pkg/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CoschedulingArgs)(nil), (*config.CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(a.(*CoschedulingArgs), b.(*config.CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CoschedulingArgs)(nil), (*CoschedulingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(a.(*config.CoschedulingArgs), b.(*CoschedulingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*LoadVariationRiskBalancingArgs)(nil), (*config.LoadVariationRiskBalancingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(a.(*LoadVariationRiskBalancingArgs), b.(*config.LoadVariationRiskBalancingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.LoadVariationRiskBalancingArgs)(nil), (*LoadVariationRiskBalancingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_LoadVariationRiskBalancingArgs_To_v1beta1_LoadVariationRiskBalancingArgs(a.(*config.LoadVariationRiskBalancingArgs), b.(*LoadVariationRiskBalancingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricProviderSpec)(nil), (*config.MetricProviderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec(a.(*MetricProviderSpec), b.(*config.MetricProviderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.MetricProviderSpec)(nil), (*MetricProviderSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec(a.(*config.MetricProviderSpec), b.(*MetricProviderSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkOverheadArgs)(nil), (*config.NetworkOverheadArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(a.(*NetworkOverheadArgs), b.(*config.NetworkOverheadArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NetworkOverheadArgs)(nil), (*NetworkOverheadArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NetworkOverheadArgs_To_v1beta1_NetworkOverheadArgs(a.(*config.NetworkOverheadArgs), b.(*NetworkOverheadArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeResourceTopologyMatchArgs)(nil), (*config.NodeResourceTopologyMatchArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs(a.(*NodeResourceTopologyMatchArgs), b.(*config.NodeResourceTopologyMatchArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeResourceTopologyMatchArgs)(nil), (*NodeResourceTopologyMatchArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeResourceTopologyMatchArgs_To_v1beta1_NodeResourceTopologyMatchArgs(a.(*config.NodeResourceTopologyMatchArgs), b.(*NodeResourceTopologyMatchArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NodeResourcesAllocatableArgs)(nil), (*config.NodeResourcesAllocatableArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(a.(*NodeResourcesAllocatableArgs), b.(*config.NodeResourcesAllocatableArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.NodeResourcesAllocatableArgs)(nil), (*NodeResourcesAllocatableArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(a.(*config.NodeResourcesAllocatableArgs), b.(*NodeResourcesAllocatableArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ScoringStrategy)(nil), (*config.ScoringStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ScoringStrategy_To_config_ScoringStrategy(a.(*ScoringStrategy), b.(*config.ScoringStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ScoringStrategy)(nil), (*ScoringStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ScoringStrategy_To_v1beta1_ScoringStrategy(a.(*config.ScoringStrategy), b.(*ScoringStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TargetLoadPackingArgs)(nil), (*config.TargetLoadPackingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(a.(*TargetLoadPackingArgs), b.(*config.TargetLoadPackingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TargetLoadPackingArgs)(nil), (*TargetLoadPackingArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TargetLoadPackingArgs_To_v1beta1_TargetLoadPackingArgs(a.(*config.TargetLoadPackingArgs), b.(*TargetLoadPackingArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TopologicalSortArgs)(nil), (*config.TopologicalSortArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_TopologicalSortArgs_To_config_TopologicalSortArgs(a.(*TopologicalSortArgs), b.(*config.TopologicalSortArgs), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.TopologicalSortArgs)(nil), (*TopologicalSortArgs)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_TopologicalSortArgs_To_v1beta1_TopologicalSortArgs(a.(*config.TopologicalSortArgs), b.(*TopologicalSortArgs), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_int64_To_int64(&in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int64_To_int64(&in.DeniedPGExpirationTimeSeconds, &out.DeniedPGExpirationTimeSeconds, s); err != nil {
		return err
	}
	// WARNING: in.KubeMaster requires manual conversion: does not exist in peer-type
	// WARNING: in.KubeConfigPath requires manual conversion: does not exist in peer-type
	return nil
}

// Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs is an autogenerated conversion function.
func Convert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in *CoschedulingArgs, out *config.CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_CoschedulingArgs_To_config_CoschedulingArgs(in, out, s)
}

func autoConvert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	if err := v1.Convert_int64_To_Pointer_int64(&in.PermitWaitingTimeSeconds, &out.PermitWaitingTimeSeconds, s); err != nil {
		return err
	}
	if err := v1.Convert_int64_To_Pointer_int64(&in.DeniedPGExpirationTimeSeconds, &out.DeniedPGExpirationTimeSeconds, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs is an autogenerated conversion function.
func Convert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in *config.CoschedulingArgs, out *CoschedulingArgs, s conversion.Scope) error {
	return autoConvert_config_CoschedulingArgs_To_v1beta1_CoschedulingArgs(in, out, s)
}

func autoConvert_v1beta1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(in *LoadVariationRiskBalancingArgs, out *config.LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	if err := Convert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_float64_To_float64(&in.SafeVarianceMargin, &out.SafeVarianceMargin, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_float64_To_float64(&in.SafeVarianceSensitivity, &out.SafeVarianceSensitivity, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs is an autogenerated conversion function.
func Convert_v1beta1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(in *LoadVariationRiskBalancingArgs, out *config.LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_LoadVariationRiskBalancingArgs_To_config_LoadVariationRiskBalancingArgs(in, out, s)
}

func autoConvert_config_LoadVariationRiskBalancingArgs_To_v1beta1_LoadVariationRiskBalancingArgs(in *config.LoadVariationRiskBalancingArgs, out *LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	if err := Convert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	if err := v1.Convert_float64_To_Pointer_float64(&in.SafeVarianceMargin, &out.SafeVarianceMargin, s); err != nil {
		return err
	}
	if err := v1.Convert_float64_To_Pointer_float64(&in.SafeVarianceSensitivity, &out.SafeVarianceSensitivity, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_LoadVariationRiskBalancingArgs_To_v1beta1_LoadVariationRiskBalancingArgs is an autogenerated conversion function.
func Convert_config_LoadVariationRiskBalancingArgs_To_v1beta1_LoadVariationRiskBalancingArgs(in *config.LoadVariationRiskBalancingArgs, out *LoadVariationRiskBalancingArgs, s conversion.Scope) error {
	return autoConvert_config_LoadVariationRiskBalancingArgs_To_v1beta1_LoadVariationRiskBalancingArgs(in, out, s)
}

func autoConvert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec(in *MetricProviderSpec, out *config.MetricProviderSpec, s conversion.Scope) error {
	out.Type = config.MetricProviderType(in.Type)
	if err := v1.Convert_Pointer_string_To_string(&in.Address, &out.Address, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.Token, &out.Token, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec is an autogenerated conversion function.
func Convert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec(in *MetricProviderSpec, out *config.MetricProviderSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec(in, out, s)
}

func autoConvert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec(in *config.MetricProviderSpec, out *MetricProviderSpec, s conversion.Scope) error {
	out.Type = MetricProviderType(in.Type)
	if err := v1.Convert_string_To_Pointer_string(&in.Address, &out.Address, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.Token, &out.Token, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec is an autogenerated conversion function.
func Convert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec(in *config.MetricProviderSpec, out *MetricProviderSpec, s conversion.Scope) error {
	return autoConvert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec(in, out, s)
}

func autoConvert_v1beta1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(in *NetworkOverheadArgs, out *config.NetworkOverheadArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_string_To_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.MasterOverride, &out.MasterOverride, s); err != nil {
		return err
	}
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	if err := v1.Convert_Pointer_string_To_string(&in.WeightsName, &out.WeightsName, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.NetworkTopologyName, &out.NetworkTopologyName, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_NetworkOverheadArgs_To_config_NetworkOverheadArgs is an autogenerated conversion function.
func Convert_v1beta1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(in *NetworkOverheadArgs, out *config.NetworkOverheadArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_NetworkOverheadArgs_To_config_NetworkOverheadArgs(in, out, s)
}

func autoConvert_config_NetworkOverheadArgs_To_v1beta1_NetworkOverheadArgs(in *config.NetworkOverheadArgs, out *NetworkOverheadArgs, s conversion.Scope) error {
	if err := v1.Convert_string_To_Pointer_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.MasterOverride, &out.MasterOverride, s); err != nil {
		return err
	}
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	if err := v1.Convert_string_To_Pointer_string(&in.WeightsName, &out.WeightsName, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.NetworkTopologyName, &out.NetworkTopologyName, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_NetworkOverheadArgs_To_v1beta1_NetworkOverheadArgs is an autogenerated conversion function.
func Convert_config_NetworkOverheadArgs_To_v1beta1_NetworkOverheadArgs(in *config.NetworkOverheadArgs, out *NetworkOverheadArgs, s conversion.Scope) error {
	return autoConvert_config_NetworkOverheadArgs_To_v1beta1_NetworkOverheadArgs(in, out, s)
}

func autoConvert_v1beta1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs(in *NodeResourceTopologyMatchArgs, out *config.NodeResourceTopologyMatchArgs, s conversion.Scope) error {
	// WARNING: in.KubeConfigPath requires manual conversion: does not exist in peer-type
	// WARNING: in.MasterOverride requires manual conversion: does not exist in peer-type
	// WARNING: in.ScoringStrategy requires manual conversion: inconvertible types (*sigs.k8s.io/scheduler-plugins/pkg/apis/config/v1beta1.ScoringStrategy vs sigs.k8s.io/scheduler-plugins/pkg/apis/config.ScoringStrategy)
	// Added manually
	out.ScoringStrategy = *(*config.ScoringStrategy)(unsafe.Pointer(in.ScoringStrategy))
	return nil
}

// Convert_v1beta1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs is an autogenerated conversion function.
func Convert_v1beta1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs(in *NodeResourceTopologyMatchArgs, out *config.NodeResourceTopologyMatchArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_NodeResourceTopologyMatchArgs_To_config_NodeResourceTopologyMatchArgs(in, out, s)
}

func autoConvert_config_NodeResourceTopologyMatchArgs_To_v1beta1_NodeResourceTopologyMatchArgs(in *config.NodeResourceTopologyMatchArgs, out *NodeResourceTopologyMatchArgs, s conversion.Scope) error {
	// WARNING: in.ScoringStrategy requires manual conversion: inconvertible types (*sigs.k8s.io/scheduler-plugins/pkg/apis/config/v1beta1.ScoringStrategy vs sigs.k8s.io/scheduler-plugins/pkg/apis/config.ScoringStrategy)
	// Added manually
	out.ScoringStrategy = (*ScoringStrategy)(unsafe.Pointer(&in.ScoringStrategy))
	return nil
}

// Convert_config_NodeResourceTopologyMatchArgs_To_v1beta1_NodeResourceTopologyMatchArgs is an autogenerated conversion function.
func Convert_config_NodeResourceTopologyMatchArgs_To_v1beta1_NodeResourceTopologyMatchArgs(in *config.NodeResourceTopologyMatchArgs, out *NodeResourceTopologyMatchArgs, s conversion.Scope) error {
	return autoConvert_config_NodeResourceTopologyMatchArgs_To_v1beta1_NodeResourceTopologyMatchArgs(in, out, s)
}


func autoConvert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs, out *config.NodeResourcesAllocatableArgs, s conversion.Scope) error {
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	out.Mode = config.ModeType(in.Mode)
	return nil
}

// Convert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs is an autogenerated conversion function.
func Convert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in *NodeResourcesAllocatableArgs, out *config.NodeResourcesAllocatableArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_NodeResourcesAllocatableArgs_To_config_NodeResourcesAllocatableArgs(in, out, s)
}

func autoConvert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(in *config.NodeResourcesAllocatableArgs, out *NodeResourcesAllocatableArgs, s conversion.Scope) error {
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	out.Mode = ModeType(in.Mode)
	return nil
}

// Convert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs is an autogenerated conversion function.
func Convert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(in *config.NodeResourcesAllocatableArgs, out *NodeResourcesAllocatableArgs, s conversion.Scope) error {
	return autoConvert_config_NodeResourcesAllocatableArgs_To_v1beta1_NodeResourcesAllocatableArgs(in, out, s)
}

func autoConvert_v1beta1_ScoringStrategy_To_config_ScoringStrategy(in *ScoringStrategy, out *config.ScoringStrategy, s conversion.Scope) error {
	out.Type = config.ScoringStrategyType(in.Type)
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1beta1_ScoringStrategy_To_config_ScoringStrategy is an autogenerated conversion function.
func Convert_v1beta1_ScoringStrategy_To_config_ScoringStrategy(in *ScoringStrategy, out *config.ScoringStrategy, s conversion.Scope) error {
	return autoConvert_v1beta1_ScoringStrategy_To_config_ScoringStrategy(in, out, s)
}

func autoConvert_config_ScoringStrategy_To_v1beta1_ScoringStrategy(in *config.ScoringStrategy, out *ScoringStrategy, s conversion.Scope) error {
	out.Type = ScoringStrategyType(in.Type)
	out.Resources = *(*[]configv1.ResourceSpec)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_config_ScoringStrategy_To_v1beta1_ScoringStrategy is an autogenerated conversion function.
func Convert_config_ScoringStrategy_To_v1beta1_ScoringStrategy(in *config.ScoringStrategy, out *ScoringStrategy, s conversion.Scope) error {
	return autoConvert_config_ScoringStrategy_To_v1beta1_ScoringStrategy(in, out, s)
}

func autoConvert_v1beta1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(in *TargetLoadPackingArgs, out *config.TargetLoadPackingArgs, s conversion.Scope) error {
	out.DefaultRequests = *(*corev1.ResourceList)(unsafe.Pointer(&in.DefaultRequests))
	if err := v1.Convert_Pointer_string_To_string(&in.DefaultRequestsMultiplier, &out.DefaultRequestsMultiplier, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_int64_To_int64(&in.TargetUtilization, &out.TargetUtilization, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_MetricProviderSpec_To_config_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs is an autogenerated conversion function.
func Convert_v1beta1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(in *TargetLoadPackingArgs, out *config.TargetLoadPackingArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_TargetLoadPackingArgs_To_config_TargetLoadPackingArgs(in, out, s)
}

func autoConvert_config_TargetLoadPackingArgs_To_v1beta1_TargetLoadPackingArgs(in *config.TargetLoadPackingArgs, out *TargetLoadPackingArgs, s conversion.Scope) error {
	out.DefaultRequests = *(*corev1.ResourceList)(unsafe.Pointer(&in.DefaultRequests))
	if err := v1.Convert_string_To_Pointer_string(&in.DefaultRequestsMultiplier, &out.DefaultRequestsMultiplier, s); err != nil {
		return err
	}
	if err := v1.Convert_int64_To_Pointer_int64(&in.TargetUtilization, &out.TargetUtilization, s); err != nil {
		return err
	}
	if err := Convert_config_MetricProviderSpec_To_v1beta1_MetricProviderSpec(&in.MetricProvider, &out.MetricProvider, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.WatcherAddress, &out.WatcherAddress, s); err != nil {
		return err
	}
	return nil
}

// Convert_config_TargetLoadPackingArgs_To_v1beta1_TargetLoadPackingArgs is an autogenerated conversion function.
func Convert_config_TargetLoadPackingArgs_To_v1beta1_TargetLoadPackingArgs(in *config.TargetLoadPackingArgs, out *TargetLoadPackingArgs, s conversion.Scope) error {
	return autoConvert_config_TargetLoadPackingArgs_To_v1beta1_TargetLoadPackingArgs(in, out, s)
}

func autoConvert_v1beta1_TopologicalSortArgs_To_config_TopologicalSortArgs(in *TopologicalSortArgs, out *config.TopologicalSortArgs, s conversion.Scope) error {
	if err := v1.Convert_Pointer_string_To_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	if err := v1.Convert_Pointer_string_To_string(&in.MasterOverride, &out.MasterOverride, s); err != nil {
		return err
	}
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	return nil
}

// Convert_v1beta1_TopologicalSortArgs_To_config_TopologicalSortArgs is an autogenerated conversion function.
func Convert_v1beta1_TopologicalSortArgs_To_config_TopologicalSortArgs(in *TopologicalSortArgs, out *config.TopologicalSortArgs, s conversion.Scope) error {
	return autoConvert_v1beta1_TopologicalSortArgs_To_config_TopologicalSortArgs(in, out, s)
}

func autoConvert_config_TopologicalSortArgs_To_v1beta1_TopologicalSortArgs(in *config.TopologicalSortArgs, out *TopologicalSortArgs, s conversion.Scope) error {
	if err := v1.Convert_string_To_Pointer_string(&in.KubeConfigPath, &out.KubeConfigPath, s); err != nil {
		return err
	}
	if err := v1.Convert_string_To_Pointer_string(&in.MasterOverride, &out.MasterOverride, s); err != nil {
		return err
	}
	out.Namespaces = *(*[]string)(unsafe.Pointer(&in.Namespaces))
	return nil
}

// Convert_config_TopologicalSortArgs_To_v1beta1_TopologicalSortArgs is an autogenerated conversion function.
func Convert_config_TopologicalSortArgs_To_v1beta1_TopologicalSortArgs(in *config.TopologicalSortArgs, out *TopologicalSortArgs, s conversion.Scope) error {
	return autoConvert_config_TopologicalSortArgs_To_v1beta1_TopologicalSortArgs(in, out, s)
}
