//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by defaulter-gen. DO NOT EDIT.

package v1alpha4

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&Configuration{}, func(obj interface{}) { SetObjectDefaults_Configuration(obj.(*Configuration)) })
	scheme.AddTypeDefaultingFunc(&Pool{}, func(obj interface{}) { SetObjectDefaults_Pool(obj.(*Pool)) })
	scheme.AddTypeDefaultingFunc(&VM{}, func(obj interface{}) { SetObjectDefaults_VM(obj.(*VM)) })
	return nil
}

func SetObjectDefaults_Configuration(in *Configuration) {
	SetDefaults_ConfigurationSpec(&in.Spec)
	SetDefaults_VMSpec(&in.Spec.VMDefaults)
	SetDefaults_VMSandboxSpec(&in.Spec.VMDefaults.Sandbox)
	SetDefaults_VMKernelSpec(&in.Spec.VMDefaults.Kernel)
}

func SetObjectDefaults_Pool(in *Pool) {
	SetDefaults_PoolSpec(&in.Spec)
}

func SetObjectDefaults_VM(in *VM) {
	SetDefaults_VMSpec(&in.Spec)
	SetDefaults_VMSandboxSpec(&in.Spec.Sandbox)
	SetDefaults_VMKernelSpec(&in.Spec.Kernel)
	SetDefaults_VMStatus(&in.Status)
}
