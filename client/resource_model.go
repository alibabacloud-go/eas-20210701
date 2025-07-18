// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResource interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *Resource
	GetClusterId() *string
	SetCpuCount(v int32) *Resource
	GetCpuCount() *int32
	SetCpuUsed(v int32) *Resource
	GetCpuUsed() *int32
	SetCreateTime(v string) *Resource
	GetCreateTime() *string
	SetExtraData(v map[string]interface{}) *Resource
	GetExtraData() map[string]interface{}
	SetFeatures(v []*string) *Resource
	GetFeatures() []*string
	SetGpuCount(v int32) *Resource
	GetGpuCount() *int32
	SetGpuUsed(v float32) *Resource
	GetGpuUsed() *float32
	SetInstanceCount(v int32) *Resource
	GetInstanceCount() *int32
	SetInstanceMaxAllocatableCPU(v int32) *Resource
	GetInstanceMaxAllocatableCPU() *int32
	SetInstanceMaxAllocatableGPU(v float32) *Resource
	GetInstanceMaxAllocatableGPU() *float32
	SetInstanceMaxAllocatableMemory(v int32) *Resource
	GetInstanceMaxAllocatableMemory() *int32
	SetMemory(v int32) *Resource
	GetMemory() *int32
	SetMemoryUsed(v int32) *Resource
	GetMemoryUsed() *int32
	SetMessage(v string) *Resource
	GetMessage() *string
	SetPostPaidInstanceCount(v int32) *Resource
	GetPostPaidInstanceCount() *int32
	SetPrePaidInstanceCount(v int32) *Resource
	GetPrePaidInstanceCount() *int32
	SetResourceId(v string) *Resource
	GetResourceId() *string
	SetResourceName(v string) *Resource
	GetResourceName() *string
	SetResourceType(v string) *Resource
	GetResourceType() *string
	SetStatus(v string) *Resource
	GetStatus() *string
	SetUpdateTime(v string) *Resource
	GetUpdateTime() *string
	SetVendor(v string) *Resource
	GetVendor() *string
}

type Resource struct {
	ClusterId                    *string                `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CpuCount                     *int32                 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	CpuUsed                      *int32                 `json:"CpuUsed,omitempty" xml:"CpuUsed,omitempty"`
	CreateTime                   *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtraData                    map[string]interface{} `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	Features                     []*string              `json:"Features,omitempty" xml:"Features,omitempty" type:"Repeated"`
	GpuCount                     *int32                 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	GpuUsed                      *float32               `json:"GpuUsed,omitempty" xml:"GpuUsed,omitempty"`
	InstanceCount                *int32                 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	InstanceMaxAllocatableCPU    *int32                 `json:"InstanceMaxAllocatableCPU,omitempty" xml:"InstanceMaxAllocatableCPU,omitempty"`
	InstanceMaxAllocatableGPU    *float32               `json:"InstanceMaxAllocatableGPU,omitempty" xml:"InstanceMaxAllocatableGPU,omitempty"`
	InstanceMaxAllocatableMemory *int32                 `json:"InstanceMaxAllocatableMemory,omitempty" xml:"InstanceMaxAllocatableMemory,omitempty"`
	Memory                       *int32                 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	MemoryUsed                   *int32                 `json:"MemoryUsed,omitempty" xml:"MemoryUsed,omitempty"`
	Message                      *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	PostPaidInstanceCount        *int32                 `json:"PostPaidInstanceCount,omitempty" xml:"PostPaidInstanceCount,omitempty"`
	PrePaidInstanceCount         *int32                 `json:"PrePaidInstanceCount,omitempty" xml:"PrePaidInstanceCount,omitempty"`
	ResourceId                   *string                `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName                 *string                `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	ResourceType                 *string                `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	Status                       *string                `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime                   *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Vendor                       *string                `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s Resource) String() string {
	return dara.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) GetClusterId() *string {
	return s.ClusterId
}

func (s *Resource) GetCpuCount() *int32 {
	return s.CpuCount
}

func (s *Resource) GetCpuUsed() *int32 {
	return s.CpuUsed
}

func (s *Resource) GetCreateTime() *string {
	return s.CreateTime
}

func (s *Resource) GetExtraData() map[string]interface{} {
	return s.ExtraData
}

func (s *Resource) GetFeatures() []*string {
	return s.Features
}

func (s *Resource) GetGpuCount() *int32 {
	return s.GpuCount
}

func (s *Resource) GetGpuUsed() *float32 {
	return s.GpuUsed
}

func (s *Resource) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *Resource) GetInstanceMaxAllocatableCPU() *int32 {
	return s.InstanceMaxAllocatableCPU
}

func (s *Resource) GetInstanceMaxAllocatableGPU() *float32 {
	return s.InstanceMaxAllocatableGPU
}

func (s *Resource) GetInstanceMaxAllocatableMemory() *int32 {
	return s.InstanceMaxAllocatableMemory
}

func (s *Resource) GetMemory() *int32 {
	return s.Memory
}

func (s *Resource) GetMemoryUsed() *int32 {
	return s.MemoryUsed
}

func (s *Resource) GetMessage() *string {
	return s.Message
}

func (s *Resource) GetPostPaidInstanceCount() *int32 {
	return s.PostPaidInstanceCount
}

func (s *Resource) GetPrePaidInstanceCount() *int32 {
	return s.PrePaidInstanceCount
}

func (s *Resource) GetResourceId() *string {
	return s.ResourceId
}

func (s *Resource) GetResourceName() *string {
	return s.ResourceName
}

func (s *Resource) GetResourceType() *string {
	return s.ResourceType
}

func (s *Resource) GetStatus() *string {
	return s.Status
}

func (s *Resource) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *Resource) GetVendor() *string {
	return s.Vendor
}

func (s *Resource) SetClusterId(v string) *Resource {
	s.ClusterId = &v
	return s
}

func (s *Resource) SetCpuCount(v int32) *Resource {
	s.CpuCount = &v
	return s
}

func (s *Resource) SetCpuUsed(v int32) *Resource {
	s.CpuUsed = &v
	return s
}

func (s *Resource) SetCreateTime(v string) *Resource {
	s.CreateTime = &v
	return s
}

func (s *Resource) SetExtraData(v map[string]interface{}) *Resource {
	s.ExtraData = v
	return s
}

func (s *Resource) SetFeatures(v []*string) *Resource {
	s.Features = v
	return s
}

func (s *Resource) SetGpuCount(v int32) *Resource {
	s.GpuCount = &v
	return s
}

func (s *Resource) SetGpuUsed(v float32) *Resource {
	s.GpuUsed = &v
	return s
}

func (s *Resource) SetInstanceCount(v int32) *Resource {
	s.InstanceCount = &v
	return s
}

func (s *Resource) SetInstanceMaxAllocatableCPU(v int32) *Resource {
	s.InstanceMaxAllocatableCPU = &v
	return s
}

func (s *Resource) SetInstanceMaxAllocatableGPU(v float32) *Resource {
	s.InstanceMaxAllocatableGPU = &v
	return s
}

func (s *Resource) SetInstanceMaxAllocatableMemory(v int32) *Resource {
	s.InstanceMaxAllocatableMemory = &v
	return s
}

func (s *Resource) SetMemory(v int32) *Resource {
	s.Memory = &v
	return s
}

func (s *Resource) SetMemoryUsed(v int32) *Resource {
	s.MemoryUsed = &v
	return s
}

func (s *Resource) SetMessage(v string) *Resource {
	s.Message = &v
	return s
}

func (s *Resource) SetPostPaidInstanceCount(v int32) *Resource {
	s.PostPaidInstanceCount = &v
	return s
}

func (s *Resource) SetPrePaidInstanceCount(v int32) *Resource {
	s.PrePaidInstanceCount = &v
	return s
}

func (s *Resource) SetResourceId(v string) *Resource {
	s.ResourceId = &v
	return s
}

func (s *Resource) SetResourceName(v string) *Resource {
	s.ResourceName = &v
	return s
}

func (s *Resource) SetResourceType(v string) *Resource {
	s.ResourceType = &v
	return s
}

func (s *Resource) SetStatus(v string) *Resource {
	s.Status = &v
	return s
}

func (s *Resource) SetUpdateTime(v string) *Resource {
	s.UpdateTime = &v
	return s
}

func (s *Resource) SetVendor(v string) *Resource {
	s.Vendor = &v
	return s
}

func (s *Resource) Validate() error {
	return dara.Validate(s)
}
