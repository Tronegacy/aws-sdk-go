// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package opsworksiface provides an interface for the AWS OpsWorks.
package opsworksiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/opsworks"
)

// OpsWorksAPI is the interface type for opsworks.OpsWorks.
type OpsWorksAPI interface {
	AssignInstanceRequest(*opsworks.AssignInstanceInput) (*service.Request, *opsworks.AssignInstanceOutput)

	AssignInstance(*opsworks.AssignInstanceInput) (*opsworks.AssignInstanceOutput, error)

	AssignVolumeRequest(*opsworks.AssignVolumeInput) (*service.Request, *opsworks.AssignVolumeOutput)

	AssignVolume(*opsworks.AssignVolumeInput) (*opsworks.AssignVolumeOutput, error)

	AssociateElasticIpRequest(*opsworks.AssociateElasticIpInput) (*service.Request, *opsworks.AssociateElasticIpOutput)

	AssociateElasticIp(*opsworks.AssociateElasticIpInput) (*opsworks.AssociateElasticIpOutput, error)

	AttachElasticLoadBalancerRequest(*opsworks.AttachElasticLoadBalancerInput) (*service.Request, *opsworks.AttachElasticLoadBalancerOutput)

	AttachElasticLoadBalancer(*opsworks.AttachElasticLoadBalancerInput) (*opsworks.AttachElasticLoadBalancerOutput, error)

	CloneStackRequest(*opsworks.CloneStackInput) (*service.Request, *opsworks.CloneStackOutput)

	CloneStack(*opsworks.CloneStackInput) (*opsworks.CloneStackOutput, error)

	CreateAppRequest(*opsworks.CreateAppInput) (*service.Request, *opsworks.CreateAppOutput)

	CreateApp(*opsworks.CreateAppInput) (*opsworks.CreateAppOutput, error)

	CreateDeploymentRequest(*opsworks.CreateDeploymentInput) (*service.Request, *opsworks.CreateDeploymentOutput)

	CreateDeployment(*opsworks.CreateDeploymentInput) (*opsworks.CreateDeploymentOutput, error)

	CreateInstanceRequest(*opsworks.CreateInstanceInput) (*service.Request, *opsworks.CreateInstanceOutput)

	CreateInstance(*opsworks.CreateInstanceInput) (*opsworks.CreateInstanceOutput, error)

	CreateLayerRequest(*opsworks.CreateLayerInput) (*service.Request, *opsworks.CreateLayerOutput)

	CreateLayer(*opsworks.CreateLayerInput) (*opsworks.CreateLayerOutput, error)

	CreateStackRequest(*opsworks.CreateStackInput) (*service.Request, *opsworks.CreateStackOutput)

	CreateStack(*opsworks.CreateStackInput) (*opsworks.CreateStackOutput, error)

	CreateUserProfileRequest(*opsworks.CreateUserProfileInput) (*service.Request, *opsworks.CreateUserProfileOutput)

	CreateUserProfile(*opsworks.CreateUserProfileInput) (*opsworks.CreateUserProfileOutput, error)

	DeleteAppRequest(*opsworks.DeleteAppInput) (*service.Request, *opsworks.DeleteAppOutput)

	DeleteApp(*opsworks.DeleteAppInput) (*opsworks.DeleteAppOutput, error)

	DeleteInstanceRequest(*opsworks.DeleteInstanceInput) (*service.Request, *opsworks.DeleteInstanceOutput)

	DeleteInstance(*opsworks.DeleteInstanceInput) (*opsworks.DeleteInstanceOutput, error)

	DeleteLayerRequest(*opsworks.DeleteLayerInput) (*service.Request, *opsworks.DeleteLayerOutput)

	DeleteLayer(*opsworks.DeleteLayerInput) (*opsworks.DeleteLayerOutput, error)

	DeleteStackRequest(*opsworks.DeleteStackInput) (*service.Request, *opsworks.DeleteStackOutput)

	DeleteStack(*opsworks.DeleteStackInput) (*opsworks.DeleteStackOutput, error)

	DeleteUserProfileRequest(*opsworks.DeleteUserProfileInput) (*service.Request, *opsworks.DeleteUserProfileOutput)

	DeleteUserProfile(*opsworks.DeleteUserProfileInput) (*opsworks.DeleteUserProfileOutput, error)

	DeregisterEcsClusterRequest(*opsworks.DeregisterEcsClusterInput) (*service.Request, *opsworks.DeregisterEcsClusterOutput)

	DeregisterEcsCluster(*opsworks.DeregisterEcsClusterInput) (*opsworks.DeregisterEcsClusterOutput, error)

	DeregisterElasticIpRequest(*opsworks.DeregisterElasticIpInput) (*service.Request, *opsworks.DeregisterElasticIpOutput)

	DeregisterElasticIp(*opsworks.DeregisterElasticIpInput) (*opsworks.DeregisterElasticIpOutput, error)

	DeregisterInstanceRequest(*opsworks.DeregisterInstanceInput) (*service.Request, *opsworks.DeregisterInstanceOutput)

	DeregisterInstance(*opsworks.DeregisterInstanceInput) (*opsworks.DeregisterInstanceOutput, error)

	DeregisterRdsDbInstanceRequest(*opsworks.DeregisterRdsDbInstanceInput) (*service.Request, *opsworks.DeregisterRdsDbInstanceOutput)

	DeregisterRdsDbInstance(*opsworks.DeregisterRdsDbInstanceInput) (*opsworks.DeregisterRdsDbInstanceOutput, error)

	DeregisterVolumeRequest(*opsworks.DeregisterVolumeInput) (*service.Request, *opsworks.DeregisterVolumeOutput)

	DeregisterVolume(*opsworks.DeregisterVolumeInput) (*opsworks.DeregisterVolumeOutput, error)

	DescribeAgentVersionsRequest(*opsworks.DescribeAgentVersionsInput) (*service.Request, *opsworks.DescribeAgentVersionsOutput)

	DescribeAgentVersions(*opsworks.DescribeAgentVersionsInput) (*opsworks.DescribeAgentVersionsOutput, error)

	DescribeAppsRequest(*opsworks.DescribeAppsInput) (*service.Request, *opsworks.DescribeAppsOutput)

	DescribeApps(*opsworks.DescribeAppsInput) (*opsworks.DescribeAppsOutput, error)

	DescribeCommandsRequest(*opsworks.DescribeCommandsInput) (*service.Request, *opsworks.DescribeCommandsOutput)

	DescribeCommands(*opsworks.DescribeCommandsInput) (*opsworks.DescribeCommandsOutput, error)

	DescribeDeploymentsRequest(*opsworks.DescribeDeploymentsInput) (*service.Request, *opsworks.DescribeDeploymentsOutput)

	DescribeDeployments(*opsworks.DescribeDeploymentsInput) (*opsworks.DescribeDeploymentsOutput, error)

	DescribeEcsClustersRequest(*opsworks.DescribeEcsClustersInput) (*service.Request, *opsworks.DescribeEcsClustersOutput)

	DescribeEcsClusters(*opsworks.DescribeEcsClustersInput) (*opsworks.DescribeEcsClustersOutput, error)

	DescribeElasticIpsRequest(*opsworks.DescribeElasticIpsInput) (*service.Request, *opsworks.DescribeElasticIpsOutput)

	DescribeElasticIps(*opsworks.DescribeElasticIpsInput) (*opsworks.DescribeElasticIpsOutput, error)

	DescribeElasticLoadBalancersRequest(*opsworks.DescribeElasticLoadBalancersInput) (*service.Request, *opsworks.DescribeElasticLoadBalancersOutput)

	DescribeElasticLoadBalancers(*opsworks.DescribeElasticLoadBalancersInput) (*opsworks.DescribeElasticLoadBalancersOutput, error)

	DescribeInstancesRequest(*opsworks.DescribeInstancesInput) (*service.Request, *opsworks.DescribeInstancesOutput)

	DescribeInstances(*opsworks.DescribeInstancesInput) (*opsworks.DescribeInstancesOutput, error)

	DescribeLayersRequest(*opsworks.DescribeLayersInput) (*service.Request, *opsworks.DescribeLayersOutput)

	DescribeLayers(*opsworks.DescribeLayersInput) (*opsworks.DescribeLayersOutput, error)

	DescribeLoadBasedAutoScalingRequest(*opsworks.DescribeLoadBasedAutoScalingInput) (*service.Request, *opsworks.DescribeLoadBasedAutoScalingOutput)

	DescribeLoadBasedAutoScaling(*opsworks.DescribeLoadBasedAutoScalingInput) (*opsworks.DescribeLoadBasedAutoScalingOutput, error)

	DescribeMyUserProfileRequest(*opsworks.DescribeMyUserProfileInput) (*service.Request, *opsworks.DescribeMyUserProfileOutput)

	DescribeMyUserProfile(*opsworks.DescribeMyUserProfileInput) (*opsworks.DescribeMyUserProfileOutput, error)

	DescribePermissionsRequest(*opsworks.DescribePermissionsInput) (*service.Request, *opsworks.DescribePermissionsOutput)

	DescribePermissions(*opsworks.DescribePermissionsInput) (*opsworks.DescribePermissionsOutput, error)

	DescribeRaidArraysRequest(*opsworks.DescribeRaidArraysInput) (*service.Request, *opsworks.DescribeRaidArraysOutput)

	DescribeRaidArrays(*opsworks.DescribeRaidArraysInput) (*opsworks.DescribeRaidArraysOutput, error)

	DescribeRdsDbInstancesRequest(*opsworks.DescribeRdsDbInstancesInput) (*service.Request, *opsworks.DescribeRdsDbInstancesOutput)

	DescribeRdsDbInstances(*opsworks.DescribeRdsDbInstancesInput) (*opsworks.DescribeRdsDbInstancesOutput, error)

	DescribeServiceErrorsRequest(*opsworks.DescribeServiceErrorsInput) (*service.Request, *opsworks.DescribeServiceErrorsOutput)

	DescribeServiceErrors(*opsworks.DescribeServiceErrorsInput) (*opsworks.DescribeServiceErrorsOutput, error)

	DescribeStackProvisioningParametersRequest(*opsworks.DescribeStackProvisioningParametersInput) (*service.Request, *opsworks.DescribeStackProvisioningParametersOutput)

	DescribeStackProvisioningParameters(*opsworks.DescribeStackProvisioningParametersInput) (*opsworks.DescribeStackProvisioningParametersOutput, error)

	DescribeStackSummaryRequest(*opsworks.DescribeStackSummaryInput) (*service.Request, *opsworks.DescribeStackSummaryOutput)

	DescribeStackSummary(*opsworks.DescribeStackSummaryInput) (*opsworks.DescribeStackSummaryOutput, error)

	DescribeStacksRequest(*opsworks.DescribeStacksInput) (*service.Request, *opsworks.DescribeStacksOutput)

	DescribeStacks(*opsworks.DescribeStacksInput) (*opsworks.DescribeStacksOutput, error)

	DescribeTimeBasedAutoScalingRequest(*opsworks.DescribeTimeBasedAutoScalingInput) (*service.Request, *opsworks.DescribeTimeBasedAutoScalingOutput)

	DescribeTimeBasedAutoScaling(*opsworks.DescribeTimeBasedAutoScalingInput) (*opsworks.DescribeTimeBasedAutoScalingOutput, error)

	DescribeUserProfilesRequest(*opsworks.DescribeUserProfilesInput) (*service.Request, *opsworks.DescribeUserProfilesOutput)

	DescribeUserProfiles(*opsworks.DescribeUserProfilesInput) (*opsworks.DescribeUserProfilesOutput, error)

	DescribeVolumesRequest(*opsworks.DescribeVolumesInput) (*service.Request, *opsworks.DescribeVolumesOutput)

	DescribeVolumes(*opsworks.DescribeVolumesInput) (*opsworks.DescribeVolumesOutput, error)

	DetachElasticLoadBalancerRequest(*opsworks.DetachElasticLoadBalancerInput) (*service.Request, *opsworks.DetachElasticLoadBalancerOutput)

	DetachElasticLoadBalancer(*opsworks.DetachElasticLoadBalancerInput) (*opsworks.DetachElasticLoadBalancerOutput, error)

	DisassociateElasticIpRequest(*opsworks.DisassociateElasticIpInput) (*service.Request, *opsworks.DisassociateElasticIpOutput)

	DisassociateElasticIp(*opsworks.DisassociateElasticIpInput) (*opsworks.DisassociateElasticIpOutput, error)

	GetHostnameSuggestionRequest(*opsworks.GetHostnameSuggestionInput) (*service.Request, *opsworks.GetHostnameSuggestionOutput)

	GetHostnameSuggestion(*opsworks.GetHostnameSuggestionInput) (*opsworks.GetHostnameSuggestionOutput, error)

	GrantAccessRequest(*opsworks.GrantAccessInput) (*service.Request, *opsworks.GrantAccessOutput)

	GrantAccess(*opsworks.GrantAccessInput) (*opsworks.GrantAccessOutput, error)

	RebootInstanceRequest(*opsworks.RebootInstanceInput) (*service.Request, *opsworks.RebootInstanceOutput)

	RebootInstance(*opsworks.RebootInstanceInput) (*opsworks.RebootInstanceOutput, error)

	RegisterEcsClusterRequest(*opsworks.RegisterEcsClusterInput) (*service.Request, *opsworks.RegisterEcsClusterOutput)

	RegisterEcsCluster(*opsworks.RegisterEcsClusterInput) (*opsworks.RegisterEcsClusterOutput, error)

	RegisterElasticIpRequest(*opsworks.RegisterElasticIpInput) (*service.Request, *opsworks.RegisterElasticIpOutput)

	RegisterElasticIp(*opsworks.RegisterElasticIpInput) (*opsworks.RegisterElasticIpOutput, error)

	RegisterInstanceRequest(*opsworks.RegisterInstanceInput) (*service.Request, *opsworks.RegisterInstanceOutput)

	RegisterInstance(*opsworks.RegisterInstanceInput) (*opsworks.RegisterInstanceOutput, error)

	RegisterRdsDbInstanceRequest(*opsworks.RegisterRdsDbInstanceInput) (*service.Request, *opsworks.RegisterRdsDbInstanceOutput)

	RegisterRdsDbInstance(*opsworks.RegisterRdsDbInstanceInput) (*opsworks.RegisterRdsDbInstanceOutput, error)

	RegisterVolumeRequest(*opsworks.RegisterVolumeInput) (*service.Request, *opsworks.RegisterVolumeOutput)

	RegisterVolume(*opsworks.RegisterVolumeInput) (*opsworks.RegisterVolumeOutput, error)

	SetLoadBasedAutoScalingRequest(*opsworks.SetLoadBasedAutoScalingInput) (*service.Request, *opsworks.SetLoadBasedAutoScalingOutput)

	SetLoadBasedAutoScaling(*opsworks.SetLoadBasedAutoScalingInput) (*opsworks.SetLoadBasedAutoScalingOutput, error)

	SetPermissionRequest(*opsworks.SetPermissionInput) (*service.Request, *opsworks.SetPermissionOutput)

	SetPermission(*opsworks.SetPermissionInput) (*opsworks.SetPermissionOutput, error)

	SetTimeBasedAutoScalingRequest(*opsworks.SetTimeBasedAutoScalingInput) (*service.Request, *opsworks.SetTimeBasedAutoScalingOutput)

	SetTimeBasedAutoScaling(*opsworks.SetTimeBasedAutoScalingInput) (*opsworks.SetTimeBasedAutoScalingOutput, error)

	StartInstanceRequest(*opsworks.StartInstanceInput) (*service.Request, *opsworks.StartInstanceOutput)

	StartInstance(*opsworks.StartInstanceInput) (*opsworks.StartInstanceOutput, error)

	StartStackRequest(*opsworks.StartStackInput) (*service.Request, *opsworks.StartStackOutput)

	StartStack(*opsworks.StartStackInput) (*opsworks.StartStackOutput, error)

	StopInstanceRequest(*opsworks.StopInstanceInput) (*service.Request, *opsworks.StopInstanceOutput)

	StopInstance(*opsworks.StopInstanceInput) (*opsworks.StopInstanceOutput, error)

	StopStackRequest(*opsworks.StopStackInput) (*service.Request, *opsworks.StopStackOutput)

	StopStack(*opsworks.StopStackInput) (*opsworks.StopStackOutput, error)

	UnassignInstanceRequest(*opsworks.UnassignInstanceInput) (*service.Request, *opsworks.UnassignInstanceOutput)

	UnassignInstance(*opsworks.UnassignInstanceInput) (*opsworks.UnassignInstanceOutput, error)

	UnassignVolumeRequest(*opsworks.UnassignVolumeInput) (*service.Request, *opsworks.UnassignVolumeOutput)

	UnassignVolume(*opsworks.UnassignVolumeInput) (*opsworks.UnassignVolumeOutput, error)

	UpdateAppRequest(*opsworks.UpdateAppInput) (*service.Request, *opsworks.UpdateAppOutput)

	UpdateApp(*opsworks.UpdateAppInput) (*opsworks.UpdateAppOutput, error)

	UpdateElasticIpRequest(*opsworks.UpdateElasticIpInput) (*service.Request, *opsworks.UpdateElasticIpOutput)

	UpdateElasticIp(*opsworks.UpdateElasticIpInput) (*opsworks.UpdateElasticIpOutput, error)

	UpdateInstanceRequest(*opsworks.UpdateInstanceInput) (*service.Request, *opsworks.UpdateInstanceOutput)

	UpdateInstance(*opsworks.UpdateInstanceInput) (*opsworks.UpdateInstanceOutput, error)

	UpdateLayerRequest(*opsworks.UpdateLayerInput) (*service.Request, *opsworks.UpdateLayerOutput)

	UpdateLayer(*opsworks.UpdateLayerInput) (*opsworks.UpdateLayerOutput, error)

	UpdateMyUserProfileRequest(*opsworks.UpdateMyUserProfileInput) (*service.Request, *opsworks.UpdateMyUserProfileOutput)

	UpdateMyUserProfile(*opsworks.UpdateMyUserProfileInput) (*opsworks.UpdateMyUserProfileOutput, error)

	UpdateRdsDbInstanceRequest(*opsworks.UpdateRdsDbInstanceInput) (*service.Request, *opsworks.UpdateRdsDbInstanceOutput)

	UpdateRdsDbInstance(*opsworks.UpdateRdsDbInstanceInput) (*opsworks.UpdateRdsDbInstanceOutput, error)

	UpdateStackRequest(*opsworks.UpdateStackInput) (*service.Request, *opsworks.UpdateStackOutput)

	UpdateStack(*opsworks.UpdateStackInput) (*opsworks.UpdateStackOutput, error)

	UpdateUserProfileRequest(*opsworks.UpdateUserProfileInput) (*service.Request, *opsworks.UpdateUserProfileOutput)

	UpdateUserProfile(*opsworks.UpdateUserProfileInput) (*opsworks.UpdateUserProfileOutput, error)

	UpdateVolumeRequest(*opsworks.UpdateVolumeInput) (*service.Request, *opsworks.UpdateVolumeOutput)

	UpdateVolume(*opsworks.UpdateVolumeInput) (*opsworks.UpdateVolumeOutput, error)
}
