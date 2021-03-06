// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type Instance struct {
	// 实例所在的宿主机IP
	HostIP *string `json:"HostIP,omitempty" xml:"HostIP,omitempty"`
	// 实例所在的宿主机名字
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// 实例的内网IP
	InnerIP *string `json:"InnerIP,omitempty" xml:"InnerIP,omitempty"`
	// 实例的名字
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 实例的网络端口
	InstancePort *int32 `json:"InstancePort,omitempty" xml:"InstancePort,omitempty"`
	// 实例上一次退出的状态
	LastState []map[string]interface{} `json:"LastState,omitempty" xml:"LastState,omitempty" type:"Repeated"`
	// 实例的命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 实例已经启动完成的进程数
	ReadyProcesses *int32 `json:"ReadyProcesses,omitempty" xml:"ReadyProcesses,omitempty"`
	// 实例当前状态的标识
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 实例重启次数
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// 实例的启动时间
	StartAt *string `json:"StartAt,omitempty" xml:"StartAt,omitempty"`
	// 实例状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 实例总的进程数
	TotalProcesses *int32 `json:"TotalProcesses,omitempty" xml:"TotalProcesses,omitempty"`
}

func (s Instance) String() string {
	return tea.Prettify(s)
}

func (s Instance) GoString() string {
	return s.String()
}

func (s *Instance) SetHostIP(v string) *Instance {
	s.HostIP = &v
	return s
}

func (s *Instance) SetHostName(v string) *Instance {
	s.HostName = &v
	return s
}

func (s *Instance) SetInnerIP(v string) *Instance {
	s.InnerIP = &v
	return s
}

func (s *Instance) SetInstanceName(v string) *Instance {
	s.InstanceName = &v
	return s
}

func (s *Instance) SetInstancePort(v int32) *Instance {
	s.InstancePort = &v
	return s
}

func (s *Instance) SetLastState(v []map[string]interface{}) *Instance {
	s.LastState = v
	return s
}

func (s *Instance) SetNamespace(v string) *Instance {
	s.Namespace = &v
	return s
}

func (s *Instance) SetReadyProcesses(v int32) *Instance {
	s.ReadyProcesses = &v
	return s
}

func (s *Instance) SetReason(v string) *Instance {
	s.Reason = &v
	return s
}

func (s *Instance) SetRestartCount(v int32) *Instance {
	s.RestartCount = &v
	return s
}

func (s *Instance) SetStartAt(v string) *Instance {
	s.StartAt = &v
	return s
}

func (s *Instance) SetStatus(v string) *Instance {
	s.Status = &v
	return s
}

func (s *Instance) SetTotalProcesses(v int32) *Instance {
	s.TotalProcesses = &v
	return s
}

type Resource struct {
	// 资源组所在的集群
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 资源组CPU数量
	CpuCount *int32 `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	// 资源组创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 资源组自定义数据
	ExtraData map[string]interface{} `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// 资源组GPU个数
	GpuCount *int32 `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	// 资源组实例个数
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// 资源组摘要信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 资源组按量付费实例个数
	PostPaidInstanceCount *int32 `json:"PostPaidInstanceCount,omitempty" xml:"PostPaidInstanceCount,omitempty"`
	// 资源组预付费实例个数
	PrePaidInstanceCount *int32 `json:"PrePaidInstanceCount,omitempty" xml:"PrePaidInstanceCount,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 资源组名字
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// 资源组的状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 资源组更新时间
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s Resource) String() string {
	return tea.Prettify(s)
}

func (s Resource) GoString() string {
	return s.String()
}

func (s *Resource) SetClusterId(v string) *Resource {
	s.ClusterId = &v
	return s
}

func (s *Resource) SetCpuCount(v int32) *Resource {
	s.CpuCount = &v
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

func (s *Resource) SetGpuCount(v int32) *Resource {
	s.GpuCount = &v
	return s
}

func (s *Resource) SetInstanceCount(v int32) *Resource {
	s.InstanceCount = &v
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

func (s *Resource) SetStatus(v string) *Resource {
	s.Status = &v
	return s
}

func (s *Resource) SetUpdateTime(v string) *Resource {
	s.UpdateTime = &v
	return s
}

type ResourceInstance struct {
	// 实例是否自动续费
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// 实例的计费类型
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 实例的创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 实例过期时间
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// 实例的Cpu个数
	InstanceCpuCount *int32 `json:"InstanceCpuCount,omitempty" xml:"InstanceCpuCount,omitempty"`
	// 实例的Gpu个数
	InstanceGpuCount *int32 `json:"InstanceGpuCount,omitempty" xml:"InstanceGpuCount,omitempty"`
	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 实例IP
	InstanceIp *string `json:"InstanceIp,omitempty" xml:"InstanceIp,omitempty"`
	// 实例的内存大小
	InstanceMemory *string `json:"InstanceMemory,omitempty" xml:"InstanceMemory,omitempty"`
	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// 实例状态
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// 实例的机型
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 实例被使用的CPU数量
	InstanceUsedCpu *float32 `json:"InstanceUsedCpu,omitempty" xml:"InstanceUsedCpu,omitempty"`
	// 实例被使用的GPU数量
	InstanceUsedGpu *int32 `json:"InstanceUsedGpu,omitempty" xml:"InstanceUsedGpu,omitempty"`
	// 实例被使用的内存大小
	InstanceUsedMemory *string `json:"InstanceUsedMemory,omitempty" xml:"InstanceUsedMemory,omitempty"`
}

func (s ResourceInstance) String() string {
	return tea.Prettify(s)
}

func (s ResourceInstance) GoString() string {
	return s.String()
}

func (s *ResourceInstance) SetAutoRenewal(v bool) *ResourceInstance {
	s.AutoRenewal = &v
	return s
}

func (s *ResourceInstance) SetChargeType(v string) *ResourceInstance {
	s.ChargeType = &v
	return s
}

func (s *ResourceInstance) SetCreateTime(v string) *ResourceInstance {
	s.CreateTime = &v
	return s
}

func (s *ResourceInstance) SetExpiredTime(v string) *ResourceInstance {
	s.ExpiredTime = &v
	return s
}

func (s *ResourceInstance) SetInstanceCpuCount(v int32) *ResourceInstance {
	s.InstanceCpuCount = &v
	return s
}

func (s *ResourceInstance) SetInstanceGpuCount(v int32) *ResourceInstance {
	s.InstanceGpuCount = &v
	return s
}

func (s *ResourceInstance) SetInstanceId(v string) *ResourceInstance {
	s.InstanceId = &v
	return s
}

func (s *ResourceInstance) SetInstanceIp(v string) *ResourceInstance {
	s.InstanceIp = &v
	return s
}

func (s *ResourceInstance) SetInstanceMemory(v string) *ResourceInstance {
	s.InstanceMemory = &v
	return s
}

func (s *ResourceInstance) SetInstanceName(v string) *ResourceInstance {
	s.InstanceName = &v
	return s
}

func (s *ResourceInstance) SetInstanceStatus(v string) *ResourceInstance {
	s.InstanceStatus = &v
	return s
}

func (s *ResourceInstance) SetInstanceType(v string) *ResourceInstance {
	s.InstanceType = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedCpu(v float32) *ResourceInstance {
	s.InstanceUsedCpu = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedGpu(v int32) *ResourceInstance {
	s.InstanceUsedGpu = &v
	return s
}

func (s *ResourceInstance) SetInstanceUsedMemory(v string) *ResourceInstance {
	s.InstanceUsedMemory = &v
	return s
}

type ResourceInstanceWorker struct {
	// CpuLimit
	CpuLimit *int32 `json:"CpuLimit,omitempty" xml:"CpuLimit,omitempty"`
	// CpuRequest
	CpuRequest *int32 `json:"CpuRequest,omitempty" xml:"CpuRequest,omitempty"`
	// GpuLimit
	GpuLimit *int32 `json:"GpuLimit,omitempty" xml:"GpuLimit,omitempty"`
	// GpuRequest
	GpuRequest *int32 `json:"GpuRequest,omitempty" xml:"GpuRequest,omitempty"`
	// MemoryLimit
	MemoryLimit *int32 `json:"MemoryLimit,omitempty" xml:"MemoryLimit,omitempty"`
	// MemoryRquest
	MemoryRquest *int32 `json:"MemoryRquest,omitempty" xml:"MemoryRquest,omitempty"`
	// pod名
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 是否ready
	Ready *bool `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// RestartCount
	RestartCount *int32 `json:"RestartCount,omitempty" xml:"RestartCount,omitempty"`
	// 服务名
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// StartTime
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// pod状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ResourceInstanceWorker) String() string {
	return tea.Prettify(s)
}

func (s ResourceInstanceWorker) GoString() string {
	return s.String()
}

func (s *ResourceInstanceWorker) SetCpuLimit(v int32) *ResourceInstanceWorker {
	s.CpuLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetCpuRequest(v int32) *ResourceInstanceWorker {
	s.CpuRequest = &v
	return s
}

func (s *ResourceInstanceWorker) SetGpuLimit(v int32) *ResourceInstanceWorker {
	s.GpuLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetGpuRequest(v int32) *ResourceInstanceWorker {
	s.GpuRequest = &v
	return s
}

func (s *ResourceInstanceWorker) SetMemoryLimit(v int32) *ResourceInstanceWorker {
	s.MemoryLimit = &v
	return s
}

func (s *ResourceInstanceWorker) SetMemoryRquest(v int32) *ResourceInstanceWorker {
	s.MemoryRquest = &v
	return s
}

func (s *ResourceInstanceWorker) SetName(v string) *ResourceInstanceWorker {
	s.Name = &v
	return s
}

func (s *ResourceInstanceWorker) SetReady(v bool) *ResourceInstanceWorker {
	s.Ready = &v
	return s
}

func (s *ResourceInstanceWorker) SetRestartCount(v int32) *ResourceInstanceWorker {
	s.RestartCount = &v
	return s
}

func (s *ResourceInstanceWorker) SetServiceName(v string) *ResourceInstanceWorker {
	s.ServiceName = &v
	return s
}

func (s *ResourceInstanceWorker) SetStartTime(v string) *ResourceInstanceWorker {
	s.StartTime = &v
	return s
}

func (s *ResourceInstanceWorker) SetStatus(v string) *ResourceInstanceWorker {
	s.Status = &v
	return s
}

type Service struct {
	// 服务的请求Token
	AccessToken *string `json:"AccessToken,omitempty" xml:"AccessToken,omitempty"`
	// 服务创建账号的UID
	CallerUid *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	// 每个实例申请的cpu
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 服务的创建时间
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 当前运行的模型版本
	CurrentVersion *int32 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// 每个实例申请的gpu
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// 服务的数据镜像
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// 服务的公网endpoint
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	// 服务内网endpoint
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	// 服务最新版本号
	LatestVersion *int32 `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	// 每个worker需要的内存大小，单位为M
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 服务的摘要信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 服务所在的命名空间
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// 服务创建账号的主账号UID
	ParentUid *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	// 被挂起的服务的实例个数
	PendingInstance *int32 `json:"PendingInstance,omitempty" xml:"PendingInstance,omitempty"`
	// 服务的状态信息
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// 服务所在的区域
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务所在的资源组
	Resource *string `json:"Resource,omitempty" xml:"Resource,omitempty"`
	// 正在运行的服务的实例个数
	RunningInstance *int32 `json:"RunningInstance,omitempty" xml:"RunningInstance,omitempty"`
	// 服务的配置信息
	ServiceConfig *string `json:"ServiceConfig,omitempty" xml:"ServiceConfig,omitempty"`
	// 服务ID
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// 服务的名字
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// 服务的部署来源
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// 服务的状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 服务的所有实例总个数
	TotalInstance *int32 `json:"TotalInstance,omitempty" xml:"TotalInstance,omitempty"`
	// 服务的更新时间
	Updatetime *string `json:"Updatetime,omitempty" xml:"Updatetime,omitempty"`
	// 服务灰度发布的权重值
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s Service) String() string {
	return tea.Prettify(s)
}

func (s Service) GoString() string {
	return s.String()
}

func (s *Service) SetAccessToken(v string) *Service {
	s.AccessToken = &v
	return s
}

func (s *Service) SetCallerUid(v string) *Service {
	s.CallerUid = &v
	return s
}

func (s *Service) SetCpu(v int32) *Service {
	s.Cpu = &v
	return s
}

func (s *Service) SetCreateTime(v string) *Service {
	s.CreateTime = &v
	return s
}

func (s *Service) SetCurrentVersion(v int32) *Service {
	s.CurrentVersion = &v
	return s
}

func (s *Service) SetGpu(v int32) *Service {
	s.Gpu = &v
	return s
}

func (s *Service) SetImage(v string) *Service {
	s.Image = &v
	return s
}

func (s *Service) SetInternetEndpoint(v string) *Service {
	s.InternetEndpoint = &v
	return s
}

func (s *Service) SetIntranetEndpoint(v string) *Service {
	s.IntranetEndpoint = &v
	return s
}

func (s *Service) SetLatestVersion(v int32) *Service {
	s.LatestVersion = &v
	return s
}

func (s *Service) SetMemory(v int32) *Service {
	s.Memory = &v
	return s
}

func (s *Service) SetMessage(v string) *Service {
	s.Message = &v
	return s
}

func (s *Service) SetNamespace(v string) *Service {
	s.Namespace = &v
	return s
}

func (s *Service) SetParentUid(v string) *Service {
	s.ParentUid = &v
	return s
}

func (s *Service) SetPendingInstance(v int32) *Service {
	s.PendingInstance = &v
	return s
}

func (s *Service) SetReason(v string) *Service {
	s.Reason = &v
	return s
}

func (s *Service) SetRegion(v string) *Service {
	s.Region = &v
	return s
}

func (s *Service) SetRequestId(v string) *Service {
	s.RequestId = &v
	return s
}

func (s *Service) SetResource(v string) *Service {
	s.Resource = &v
	return s
}

func (s *Service) SetRunningInstance(v int32) *Service {
	s.RunningInstance = &v
	return s
}

func (s *Service) SetServiceConfig(v string) *Service {
	s.ServiceConfig = &v
	return s
}

func (s *Service) SetServiceId(v string) *Service {
	s.ServiceId = &v
	return s
}

func (s *Service) SetServiceName(v string) *Service {
	s.ServiceName = &v
	return s
}

func (s *Service) SetSource(v string) *Service {
	s.Source = &v
	return s
}

func (s *Service) SetStatus(v string) *Service {
	s.Status = &v
	return s
}

func (s *Service) SetTotalInstance(v int32) *Service {
	s.TotalInstance = &v
	return s
}

func (s *Service) SetUpdatetime(v string) *Service {
	s.Updatetime = &v
	return s
}

func (s *Service) SetWeight(v int32) *Service {
	s.Weight = &v
	return s
}

type CreateBenchmarkTaskRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBenchmarkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskRequest) SetBody(v string) *CreateBenchmarkTaskRequest {
	s.Body = &v
	return s
}

type CreateBenchmarkTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region    *string `json:"Region,omitempty" xml:"Region,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskResponseBody) SetMessage(v string) *CreateBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetName(v string) *CreateBenchmarkTaskResponseBody {
	s.Name = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetRegion(v string) *CreateBenchmarkTaskResponseBody {
	s.Region = &v
	return s
}

func (s *CreateBenchmarkTaskResponseBody) SetRequestId(v string) *CreateBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateBenchmarkTaskResponse) SetHeaders(v map[string]*string) *CreateBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateBenchmarkTaskResponse) SetStatusCode(v int32) *CreateBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBenchmarkTaskResponse) SetBody(v *CreateBenchmarkTaskResponseBody) *CreateBenchmarkTaskResponse {
	s.Body = v
	return s
}

type CreateResourceRequest struct {
	// 是否自动续费
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// 付费类型，预付费PrePaid，后付费PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 实例数量
	EcsInstanceCount *int32 `json:"EcsInstanceCount,omitempty" xml:"EcsInstanceCount,omitempty"`
	// 实例机型，对应ecs机型
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
}

func (s CreateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRequest) SetAutoRenewal(v bool) *CreateResourceRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateResourceRequest) SetChargeType(v string) *CreateResourceRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateResourceRequest) SetEcsInstanceCount(v int32) *CreateResourceRequest {
	s.EcsInstanceCount = &v
	return s
}

func (s *CreateResourceRequest) SetEcsInstanceType(v string) *CreateResourceRequest {
	s.EcsInstanceType = &v
	return s
}

type CreateResourceResponseBody struct {
	// 资源组所在集群ID
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 资源组的Owner UID
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 资源组名称
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s CreateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceResponseBody) SetClusterId(v string) *CreateResourceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateResourceResponseBody) SetOwnerUid(v string) *CreateResourceResponseBody {
	s.OwnerUid = &v
	return s
}

func (s *CreateResourceResponseBody) SetRequestId(v string) *CreateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceId(v string) *CreateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *CreateResourceResponseBody) SetResourceName(v string) *CreateResourceResponseBody {
	s.ResourceName = &v
	return s
}

type CreateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceResponse) SetHeaders(v map[string]*string) *CreateResourceResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceResponse) SetStatusCode(v int32) *CreateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceResponse) SetBody(v *CreateResourceResponseBody) *CreateResourceResponse {
	s.Body = v
	return s
}

type CreateResourceInstancesRequest struct {
	// 是否自动续费
	AutoRenewal *bool `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// 付费类型，预付费PrePaid，后付费PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 新创建的实例个数，(0, 100]
	EcsInstanceCount *int32 `json:"EcsInstanceCount,omitempty" xml:"EcsInstanceCount,omitempty"`
	// 实例机型，对应ecs机型
	EcsInstanceType *string `json:"EcsInstanceType,omitempty" xml:"EcsInstanceType,omitempty"`
	// 用户自这义数据，小于 16KB
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesRequest) SetAutoRenewal(v bool) *CreateResourceInstancesRequest {
	s.AutoRenewal = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetChargeType(v string) *CreateResourceInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetEcsInstanceCount(v int32) *CreateResourceInstancesRequest {
	s.EcsInstanceCount = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetEcsInstanceType(v string) *CreateResourceInstancesRequest {
	s.EcsInstanceType = &v
	return s
}

func (s *CreateResourceInstancesRequest) SetUserData(v string) *CreateResourceInstancesRequest {
	s.UserData = &v
	return s
}

type CreateResourceInstancesResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesResponseBody) SetMessage(v string) *CreateResourceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceInstancesResponseBody) SetRequestId(v string) *CreateResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceInstancesResponse) SetHeaders(v map[string]*string) *CreateResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceInstancesResponse) SetStatusCode(v int32) *CreateResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceInstancesResponse) SetBody(v *CreateResourceInstancesResponseBody) *CreateResourceInstancesResponse {
	s.Body = v
	return s
}

type CreateResourceLogRequest struct {
	// sls日志库
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// 资源组对应的sls日志管理项目
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s CreateResourceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceLogRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceLogRequest) SetLogStore(v string) *CreateResourceLogRequest {
	s.LogStore = &v
	return s
}

func (s *CreateResourceLogRequest) SetProjectName(v string) *CreateResourceLogRequest {
	s.ProjectName = &v
	return s
}

type CreateResourceLogResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateResourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateResourceLogResponseBody) SetMessage(v string) *CreateResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *CreateResourceLogResponseBody) SetRequestId(v string) *CreateResourceLogResponseBody {
	s.RequestId = &v
	return s
}

type CreateResourceLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateResourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateResourceLogResponse) GoString() string {
	return s.String()
}

func (s *CreateResourceLogResponse) SetHeaders(v map[string]*string) *CreateResourceLogResponse {
	s.Headers = v
	return s
}

func (s *CreateResourceLogResponse) SetStatusCode(v int32) *CreateResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateResourceLogResponse) SetBody(v *CreateResourceLogResponseBody) *CreateResourceLogResponse {
	s.Body = v
	return s
}

type CreateServiceRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRequest) SetBody(v string) *CreateServiceRequest {
	s.Body = &v
	return s
}

type CreateServiceResponseBody struct {
	InternetEndpoint *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceId   *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceResponseBody) SetInternetEndpoint(v string) *CreateServiceResponseBody {
	s.InternetEndpoint = &v
	return s
}

func (s *CreateServiceResponseBody) SetIntranetEndpoint(v string) *CreateServiceResponseBody {
	s.IntranetEndpoint = &v
	return s
}

func (s *CreateServiceResponseBody) SetRegion(v string) *CreateServiceResponseBody {
	s.Region = &v
	return s
}

func (s *CreateServiceResponseBody) SetRequestId(v string) *CreateServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceId(v string) *CreateServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *CreateServiceResponseBody) SetServiceName(v string) *CreateServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *CreateServiceResponseBody) SetStatus(v string) *CreateServiceResponseBody {
	s.Status = &v
	return s
}

type CreateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceResponse) SetHeaders(v map[string]*string) *CreateServiceResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceResponse) SetStatusCode(v int32) *CreateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceResponse) SetBody(v *CreateServiceResponseBody) *CreateServiceResponse {
	s.Body = v
	return s
}

type CreateServiceAutoScalerRequest struct {
	// 最大 replica 数，需要大于MinReplica
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// 最小 replica 数，需要大于0
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// map 类型的策略定义
	Strategies *CreateServiceAutoScalerRequestStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Struct"`
}

func (s CreateServiceAutoScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequest) SetMax(v int32) *CreateServiceAutoScalerRequest {
	s.Max = &v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetMin(v int32) *CreateServiceAutoScalerRequest {
	s.Min = &v
	return s
}

func (s *CreateServiceAutoScalerRequest) SetStrategies(v *CreateServiceAutoScalerRequestStrategies) *CreateServiceAutoScalerRequest {
	s.Strategies = v
	return s
}

type CreateServiceAutoScalerRequestStrategies struct {
	// 最大 replica 数，需要大于MinReplica
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 每个实例支持的最大qps数，超出即扩容
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
}

func (s CreateServiceAutoScalerRequestStrategies) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerRequestStrategies) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerRequestStrategies) SetCpu(v float32) *CreateServiceAutoScalerRequestStrategies {
	s.Cpu = &v
	return s
}

func (s *CreateServiceAutoScalerRequestStrategies) SetQps(v float32) *CreateServiceAutoScalerRequestStrategies {
	s.Qps = &v
	return s
}

type CreateServiceAutoScalerResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerResponseBody) SetMessage(v string) *CreateServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceAutoScalerResponseBody) SetRequestId(v string) *CreateServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceAutoScalerResponse) SetHeaders(v map[string]*string) *CreateServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceAutoScalerResponse) SetStatusCode(v int32) *CreateServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceAutoScalerResponse) SetBody(v *CreateServiceAutoScalerResponseBody) *CreateServiceAutoScalerResponse {
	s.Body = v
	return s
}

type CreateServiceCronScalerRequest struct {
	// 需要排除的时间点的cron表达式
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// 定时伸缩任务描述
	ScaleJobs []*CreateServiceCronScalerRequestScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
}

func (s CreateServiceCronScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerRequest) SetExcludeDates(v []*string) *CreateServiceCronScalerRequest {
	s.ExcludeDates = v
	return s
}

func (s *CreateServiceCronScalerRequest) SetScaleJobs(v []*CreateServiceCronScalerRequestScaleJobs) *CreateServiceCronScalerRequest {
	s.ScaleJobs = v
	return s
}

type CreateServiceCronScalerRequestScaleJobs struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 要执行伸缩任务的cron表达式
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// 执行伸缩任务的目标replica
	TargetSize *int32 `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s CreateServiceCronScalerRequestScaleJobs) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerRequestScaleJobs) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetName(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.Name = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetSchedule(v string) *CreateServiceCronScalerRequestScaleJobs {
	s.Schedule = &v
	return s
}

func (s *CreateServiceCronScalerRequestScaleJobs) SetTargetSize(v int32) *CreateServiceCronScalerRequestScaleJobs {
	s.TargetSize = &v
	return s
}

type CreateServiceCronScalerResponseBody struct {
	// 操作成功消息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerResponseBody) SetMessage(v string) *CreateServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceCronScalerResponseBody) SetRequestId(v string) *CreateServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceCronScalerResponse) SetHeaders(v map[string]*string) *CreateServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceCronScalerResponse) SetStatusCode(v int32) *CreateServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceCronScalerResponse) SetBody(v *CreateServiceCronScalerResponseBody) *CreateServiceCronScalerResponse {
	s.Body = v
	return s
}

type CreateServiceMirrorRequest struct {
	// 比例 [0, 100]
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// 服务实例列表
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s CreateServiceMirrorRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorRequest) SetRatio(v int32) *CreateServiceMirrorRequest {
	s.Ratio = &v
	return s
}

func (s *CreateServiceMirrorRequest) SetTarget(v []*string) *CreateServiceMirrorRequest {
	s.Target = v
	return s
}

type CreateServiceMirrorResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorResponseBody) SetMessage(v string) *CreateServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceMirrorResponseBody) SetRequestId(v string) *CreateServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

type CreateServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *CreateServiceMirrorResponse) SetHeaders(v map[string]*string) *CreateServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *CreateServiceMirrorResponse) SetStatusCode(v int32) *CreateServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateServiceMirrorResponse) SetBody(v *CreateServiceMirrorResponseBody) *CreateServiceMirrorResponse {
	s.Body = v
	return s
}

type CreateStressRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStressRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStressRequest) GoString() string {
	return s.String()
}

func (s *CreateStressRequest) SetBody(v string) *CreateStressRequest {
	s.Body = &v
	return s
}

type CreateStressResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region  *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStressResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStressResponseBody) SetMessage(v string) *CreateStressResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStressResponseBody) SetName(v string) *CreateStressResponseBody {
	s.Name = &v
	return s
}

func (s *CreateStressResponseBody) SetRegion(v string) *CreateStressResponseBody {
	s.Region = &v
	return s
}

func (s *CreateStressResponseBody) SetRequestId(v string) *CreateStressResponseBody {
	s.RequestId = &v
	return s
}

type CreateStressResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateStressResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStressResponse) GoString() string {
	return s.String()
}

func (s *CreateStressResponse) SetHeaders(v map[string]*string) *CreateStressResponse {
	s.Headers = v
	return s
}

func (s *CreateStressResponse) SetStatusCode(v int32) *CreateStressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStressResponse) SetBody(v *CreateStressResponseBody) *CreateStressResponse {
	s.Body = v
	return s
}

type DeleteBenchmarkTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskResponseBody) SetMessage(v string) *DeleteBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteBenchmarkTaskResponseBody) SetRequestId(v string) *DeleteBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type DeleteBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *DeleteBenchmarkTaskResponse) SetHeaders(v map[string]*string) *DeleteBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *DeleteBenchmarkTaskResponse) SetStatusCode(v int32) *DeleteBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBenchmarkTaskResponse) SetBody(v *DeleteBenchmarkTaskResponseBody) *DeleteBenchmarkTaskResponse {
	s.Body = v
	return s
}

type DeleteResourceResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponseBody) SetMessage(v string) *DeleteResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceResponseBody) SetRequestId(v string) *DeleteResourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceResponse) SetHeaders(v map[string]*string) *DeleteResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceResponse) SetStatusCode(v int32) *DeleteResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceResponse) SetBody(v *DeleteResourceResponseBody) *DeleteResourceResponse {
	s.Body = v
	return s
}

type DeleteResourceDLinkResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceDLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkResponseBody) SetMessage(v string) *DeleteResourceDLinkResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceDLinkResponseBody) SetRequestId(v string) *DeleteResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceDLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceDLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceDLinkResponse) SetHeaders(v map[string]*string) *DeleteResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceDLinkResponse) SetStatusCode(v int32) *DeleteResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceDLinkResponse) SetBody(v *DeleteResourceDLinkResponseBody) *DeleteResourceDLinkResponse {
	s.Body = v
	return s
}

type DeleteResourceInstancesRequest struct {
	AllFailed    *bool   `json:"AllFailed,omitempty" xml:"AllFailed,omitempty"`
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
}

func (s DeleteResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesRequest) SetAllFailed(v bool) *DeleteResourceInstancesRequest {
	s.AllFailed = &v
	return s
}

func (s *DeleteResourceInstancesRequest) SetInstanceList(v string) *DeleteResourceInstancesRequest {
	s.InstanceList = &v
	return s
}

type DeleteResourceInstancesResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesResponseBody) SetMessage(v string) *DeleteResourceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceInstancesResponseBody) SetRequestId(v string) *DeleteResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceInstancesResponse) SetHeaders(v map[string]*string) *DeleteResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceInstancesResponse) SetStatusCode(v int32) *DeleteResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceInstancesResponse) SetBody(v *DeleteResourceInstancesResponseBody) *DeleteResourceInstancesResponse {
	s.Body = v
	return s
}

type DeleteResourceLogResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteResourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogResponseBody) SetMessage(v string) *DeleteResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteResourceLogResponseBody) SetRequestId(v string) *DeleteResourceLogResponseBody {
	s.RequestId = &v
	return s
}

type DeleteResourceLogResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteResourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteResourceLogResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceLogResponse) SetHeaders(v map[string]*string) *DeleteResourceLogResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceLogResponse) SetStatusCode(v int32) *DeleteResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceLogResponse) SetBody(v *DeleteResourceLogResponseBody) *DeleteResourceLogResponse {
	s.Body = v
	return s
}

type DeleteServiceResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponseBody) SetMessage(v string) *DeleteServiceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceResponseBody) SetRequestId(v string) *DeleteServiceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceResponse) SetHeaders(v map[string]*string) *DeleteServiceResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceResponse) SetStatusCode(v int32) *DeleteServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceResponse) SetBody(v *DeleteServiceResponseBody) *DeleteServiceResponse {
	s.Body = v
	return s
}

type DeleteServiceAutoScalerResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerResponseBody) SetMessage(v string) *DeleteServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceAutoScalerResponseBody) SetRequestId(v string) *DeleteServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceAutoScalerResponse) SetHeaders(v map[string]*string) *DeleteServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceAutoScalerResponse) SetStatusCode(v int32) *DeleteServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceAutoScalerResponse) SetBody(v *DeleteServiceAutoScalerResponseBody) *DeleteServiceAutoScalerResponse {
	s.Body = v
	return s
}

type DeleteServiceCronScalerResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerResponseBody) SetMessage(v string) *DeleteServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceCronScalerResponseBody) SetRequestId(v string) *DeleteServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceCronScalerResponse) SetHeaders(v map[string]*string) *DeleteServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceCronScalerResponse) SetStatusCode(v int32) *DeleteServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceCronScalerResponse) SetBody(v *DeleteServiceCronScalerResponseBody) *DeleteServiceCronScalerResponse {
	s.Body = v
	return s
}

type DeleteServiceInstancesRequest struct {
	// 删除的实例列表，多个实例名字之间逗号隔开
	InstanceList *string `json:"InstanceList,omitempty" xml:"InstanceList,omitempty"`
}

func (s DeleteServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesRequest) SetInstanceList(v string) *DeleteServiceInstancesRequest {
	s.InstanceList = &v
	return s
}

type DeleteServiceInstancesResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponseBody) SetMessage(v string) *DeleteServiceInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceInstancesResponseBody) SetRequestId(v string) *DeleteServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceInstancesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceInstancesResponse) SetHeaders(v map[string]*string) *DeleteServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceInstancesResponse) SetStatusCode(v int32) *DeleteServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceInstancesResponse) SetBody(v *DeleteServiceInstancesResponseBody) *DeleteServiceInstancesResponse {
	s.Body = v
	return s
}

type DeleteServiceMirrorResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorResponseBody) SetMessage(v string) *DeleteServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceMirrorResponseBody) SetRequestId(v string) *DeleteServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

type DeleteServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceMirrorResponse) SetHeaders(v map[string]*string) *DeleteServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceMirrorResponse) SetStatusCode(v int32) *DeleteServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceMirrorResponse) SetBody(v *DeleteServiceMirrorResponseBody) *DeleteServiceMirrorResponse {
	s.Body = v
	return s
}

type DeleteStressResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStressResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStressResponseBody) SetMessage(v string) *DeleteStressResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteStressResponseBody) SetRequestId(v string) *DeleteStressResponseBody {
	s.RequestId = &v
	return s
}

type DeleteStressResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteStressResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStressResponse) GoString() string {
	return s.String()
}

func (s *DeleteStressResponse) SetHeaders(v map[string]*string) *DeleteStressResponse {
	s.Headers = v
	return s
}

func (s *DeleteStressResponse) SetStatusCode(v int32) *DeleteStressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStressResponse) SetBody(v *DeleteStressResponseBody) *DeleteStressResponse {
	s.Body = v
	return s
}

type DescribeBenchmarkTaskResponseBody struct {
	AvailableAgent *int64  `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	CallerUid      *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	DesiredAgent   *int64  `json:"DesiredAgent,omitempty" xml:"DesiredAgent,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentUid      *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RequestId      *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName       *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	Token          *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskResponseBody) SetAvailableAgent(v int64) *DescribeBenchmarkTaskResponseBody {
	s.AvailableAgent = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetCallerUid(v string) *DescribeBenchmarkTaskResponseBody {
	s.CallerUid = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetDesiredAgent(v int64) *DescribeBenchmarkTaskResponseBody {
	s.DesiredAgent = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetMessage(v string) *DescribeBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetParentUid(v string) *DescribeBenchmarkTaskResponseBody {
	s.ParentUid = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetReason(v string) *DescribeBenchmarkTaskResponseBody {
	s.Reason = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetRequestId(v string) *DescribeBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetServiceName(v string) *DescribeBenchmarkTaskResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetStatus(v string) *DescribeBenchmarkTaskResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetTaskId(v string) *DescribeBenchmarkTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetTaskName(v string) *DescribeBenchmarkTaskResponseBody {
	s.TaskName = &v
	return s
}

func (s *DescribeBenchmarkTaskResponseBody) SetToken(v string) *DescribeBenchmarkTaskResponseBody {
	s.Token = &v
	return s
}

type DescribeBenchmarkTaskResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskResponse) SetHeaders(v map[string]*string) *DescribeBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeBenchmarkTaskResponse) SetStatusCode(v int32) *DescribeBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBenchmarkTaskResponse) SetBody(v *DescribeBenchmarkTaskResponseBody) *DescribeBenchmarkTaskResponse {
	s.Body = v
	return s
}

type DescribeBenchmarkTaskReportResponseBody struct {
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeBenchmarkTaskReportResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetReportUrl(v string) *DescribeBenchmarkTaskReportResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponseBody) SetRequestId(v string) *DescribeBenchmarkTaskReportResponseBody {
	s.RequestId = &v
	return s
}

type DescribeBenchmarkTaskReportResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeBenchmarkTaskReportResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeBenchmarkTaskReportResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeBenchmarkTaskReportResponse) GoString() string {
	return s.String()
}

func (s *DescribeBenchmarkTaskReportResponse) SetHeaders(v map[string]*string) *DescribeBenchmarkTaskReportResponse {
	s.Headers = v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) SetStatusCode(v int32) *DescribeBenchmarkTaskReportResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBenchmarkTaskReportResponse) SetBody(v *DescribeBenchmarkTaskReportResponseBody) *DescribeBenchmarkTaskReportResponse {
	s.Body = v
	return s
}

type DescribeResourceResponseBody struct {
	ClusterId             *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CpuCount              *int32  `json:"CpuCount,omitempty" xml:"CpuCount,omitempty"`
	CreateTime            *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	ExtraData             *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	GpuCount              *int32  `json:"GpuCount,omitempty" xml:"GpuCount,omitempty"`
	InstanceCount         *int32  `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	Message               *string `json:"Message,omitempty" xml:"Message,omitempty"`
	OwnerUid              *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	PostPaidInstanceCount *int32  `json:"PostPaidInstanceCount,omitempty" xml:"PostPaidInstanceCount,omitempty"`
	PrePaidInstanceCount  *int32  `json:"PrePaidInstanceCount,omitempty" xml:"PrePaidInstanceCount,omitempty"`
	RequestId             *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId            *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName          *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	Status                *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime            *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceResponseBody) SetClusterId(v string) *DescribeResourceResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCpuCount(v int32) *DescribeResourceResponseBody {
	s.CpuCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetCreateTime(v string) *DescribeResourceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeResourceResponseBody) SetExtraData(v string) *DescribeResourceResponseBody {
	s.ExtraData = &v
	return s
}

func (s *DescribeResourceResponseBody) SetGpuCount(v int32) *DescribeResourceResponseBody {
	s.GpuCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetInstanceCount(v int32) *DescribeResourceResponseBody {
	s.InstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetMessage(v string) *DescribeResourceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceResponseBody) SetOwnerUid(v string) *DescribeResourceResponseBody {
	s.OwnerUid = &v
	return s
}

func (s *DescribeResourceResponseBody) SetPostPaidInstanceCount(v int32) *DescribeResourceResponseBody {
	s.PostPaidInstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetPrePaidInstanceCount(v int32) *DescribeResourceResponseBody {
	s.PrePaidInstanceCount = &v
	return s
}

func (s *DescribeResourceResponseBody) SetRequestId(v string) *DescribeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceId(v string) *DescribeResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *DescribeResourceResponseBody) SetResourceName(v string) *DescribeResourceResponseBody {
	s.ResourceName = &v
	return s
}

func (s *DescribeResourceResponseBody) SetStatus(v string) *DescribeResourceResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeResourceResponseBody) SetUpdateTime(v string) *DescribeResourceResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeResourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceResponse) SetHeaders(v map[string]*string) *DescribeResourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceResponse) SetStatusCode(v int32) *DescribeResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceResponse) SetBody(v *DescribeResourceResponseBody) *DescribeResourceResponse {
	s.Body = v
	return s
}

type DescribeResourceDLinkResponseBody struct {
	// 已打通直连的副VSwitch ID
	AuxVSwitchList []*string `json:"AuxVSwitchList,omitempty" xml:"AuxVSwitchList,omitempty" type:"Repeated"`
	// 要打通的客户端的网段信息，会将该网段加入到服务端的回包路由中，与VSwitchIdList可二选一
	DestinationCIDRs *string `json:"DestinationCIDRs,omitempty" xml:"DestinationCIDRs,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 已打通直连的安全组
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 已打通直连的主VSwitch ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 已打通直接的Vpc ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeResourceDLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkResponseBody) SetAuxVSwitchList(v []*string) *DescribeResourceDLinkResponseBody {
	s.AuxVSwitchList = v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetDestinationCIDRs(v string) *DescribeResourceDLinkResponseBody {
	s.DestinationCIDRs = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetRequestId(v string) *DescribeResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetSecurityGroupId(v string) *DescribeResourceDLinkResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetVSwitchId(v string) *DescribeResourceDLinkResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *DescribeResourceDLinkResponseBody) SetVpcId(v string) *DescribeResourceDLinkResponseBody {
	s.VpcId = &v
	return s
}

type DescribeResourceDLinkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceDLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceDLinkResponse) SetHeaders(v map[string]*string) *DescribeResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceDLinkResponse) SetStatusCode(v int32) *DescribeResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceDLinkResponse) SetBody(v *DescribeResourceDLinkResponseBody) *DescribeResourceDLinkResponse {
	s.Body = v
	return s
}

type DescribeResourceLogResponseBody struct {
	// sls日志库
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// sls日志信息
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 资源组对应的sls日志管理项目
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 资源组状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogResponseBody) SetLogStore(v string) *DescribeResourceLogResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetMessage(v string) *DescribeResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetProjectName(v string) *DescribeResourceLogResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetRequestId(v string) *DescribeResourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetStatus(v string) *DescribeResourceLogResponseBody {
	s.Status = &v
	return s
}

type DescribeResourceLogResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeResourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeResourceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeResourceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogResponse) SetHeaders(v map[string]*string) *DescribeResourceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeResourceLogResponse) SetStatusCode(v int32) *DescribeResourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeResourceLogResponse) SetBody(v *DescribeResourceLogResponseBody) *DescribeResourceLogResponse {
	s.Body = v
	return s
}

type DescribeServiceResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Service           `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceResponse) SetHeaders(v map[string]*string) *DescribeServiceResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceResponse) SetStatusCode(v int32) *DescribeServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceResponse) SetBody(v *Service) *DescribeServiceResponse {
	s.Body = v
	return s
}

type DescribeServiceAutoScalerResponseBody struct {
	CurrentValues map[string]interface{} `json:"CurrentValues,omitempty" xml:"CurrentValues,omitempty"`
	// 服务最大实例数
	MaxReplica *int32 `json:"MaxReplica,omitempty" xml:"MaxReplica,omitempty"`
	// 服务最小实例数
	MinReplica *int32 `json:"MinReplica,omitempty" xml:"MinReplica,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务名字
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// 扩缩控制器控制策略
	Strategies map[string]interface{} `json:"Strategies,omitempty" xml:"Strategies,omitempty"`
}

func (s DescribeServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponseBody) SetCurrentValues(v map[string]interface{}) *DescribeServiceAutoScalerResponseBody {
	s.CurrentValues = v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetMaxReplica(v int32) *DescribeServiceAutoScalerResponseBody {
	s.MaxReplica = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetMinReplica(v int32) *DescribeServiceAutoScalerResponseBody {
	s.MinReplica = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetRequestId(v string) *DescribeServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetServiceName(v string) *DescribeServiceAutoScalerResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceAutoScalerResponseBody) SetStrategies(v map[string]interface{}) *DescribeServiceAutoScalerResponseBody {
	s.Strategies = v
	return s
}

type DescribeServiceAutoScalerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceAutoScalerResponse) SetHeaders(v map[string]*string) *DescribeServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceAutoScalerResponse) SetStatusCode(v int32) *DescribeServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceAutoScalerResponse) SetBody(v *DescribeServiceAutoScalerResponseBody) *DescribeServiceAutoScalerResponse {
	s.Body = v
	return s
}

type DescribeServiceCronScalerResponseBody struct {
	ExcludeDates []*string                                         `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	RequestId    *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ScaleJobs    []*DescribeServiceCronScalerResponseBodyScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
	ServiceName  *string                                           `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s DescribeServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponseBody) SetExcludeDates(v []*string) *DescribeServiceCronScalerResponseBody {
	s.ExcludeDates = v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetRequestId(v string) *DescribeServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetScaleJobs(v []*DescribeServiceCronScalerResponseBodyScaleJobs) *DescribeServiceCronScalerResponseBody {
	s.ScaleJobs = v
	return s
}

func (s *DescribeServiceCronScalerResponseBody) SetServiceName(v string) *DescribeServiceCronScalerResponseBody {
	s.ServiceName = &v
	return s
}

type DescribeServiceCronScalerResponseBodyScaleJobs struct {
	CreateTime    *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	LastProbeTime *string `json:"LastProbeTime,omitempty" xml:"LastProbeTime,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Name          *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Schedule      *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	State         *string `json:"State,omitempty" xml:"State,omitempty"`
	TargetSize    *int32  `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s DescribeServiceCronScalerResponseBodyScaleJobs) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceCronScalerResponseBodyScaleJobs) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetCreateTime(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.CreateTime = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetLastProbeTime(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.LastProbeTime = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetMessage(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Message = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetName(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Name = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetSchedule(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.Schedule = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetState(v string) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.State = &v
	return s
}

func (s *DescribeServiceCronScalerResponseBodyScaleJobs) SetTargetSize(v int32) *DescribeServiceCronScalerResponseBodyScaleJobs {
	s.TargetSize = &v
	return s
}

type DescribeServiceCronScalerResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceCronScalerResponse) SetHeaders(v map[string]*string) *DescribeServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceCronScalerResponse) SetStatusCode(v int32) *DescribeServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceCronScalerResponse) SetBody(v *DescribeServiceCronScalerResponseBody) *DescribeServiceCronScalerResponse {
	s.Body = v
	return s
}

type DescribeServiceLogRequest struct {
	// 查询的结束时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 要查询的机器ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// 查询的关键字
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// 请求的页码（默认为1）
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// 每页的大小（默认为500）
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 查询的开始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeServiceLogRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLogRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogRequest) SetEndTime(v string) *DescribeServiceLogRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeServiceLogRequest) SetIp(v string) *DescribeServiceLogRequest {
	s.Ip = &v
	return s
}

func (s *DescribeServiceLogRequest) SetKeyword(v string) *DescribeServiceLogRequest {
	s.Keyword = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPageNum(v int64) *DescribeServiceLogRequest {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceLogRequest) SetPageSize(v int64) *DescribeServiceLogRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeServiceLogRequest) SetStartTime(v string) *DescribeServiceLogRequest {
	s.StartTime = &v
	return s
}

type DescribeServiceLogResponseBody struct {
	// 返回的日志信息
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// 当前页码
	PageNum *int64 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 总计数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 总计页码
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s DescribeServiceLogResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogResponseBody) SetLogs(v []*string) *DescribeServiceLogResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeServiceLogResponseBody) SetPageNum(v int64) *DescribeServiceLogResponseBody {
	s.PageNum = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetRequestId(v string) *DescribeServiceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetTotalCount(v int64) *DescribeServiceLogResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeServiceLogResponseBody) SetTotalPageNum(v int64) *DescribeServiceLogResponseBody {
	s.TotalPageNum = &v
	return s
}

type DescribeServiceLogResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceLogResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceLogResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceLogResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceLogResponse) SetHeaders(v map[string]*string) *DescribeServiceLogResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceLogResponse) SetStatusCode(v int32) *DescribeServiceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceLogResponse) SetBody(v *DescribeServiceLogResponseBody) *DescribeServiceLogResponse {
	s.Body = v
	return s
}

type DescribeServiceMirrorResponseBody struct {
	// 比例[0,100]
	Ratio *string `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 服务名字
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// 设置流量镜像对服务列表
	Target *string `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s DescribeServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorResponseBody) SetRatio(v string) *DescribeServiceMirrorResponseBody {
	s.Ratio = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetRequestId(v string) *DescribeServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetServiceName(v string) *DescribeServiceMirrorResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeServiceMirrorResponseBody) SetTarget(v string) *DescribeServiceMirrorResponseBody {
	s.Target = &v
	return s
}

type DescribeServiceMirrorResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMirrorResponse) SetHeaders(v map[string]*string) *DescribeServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMirrorResponse) SetStatusCode(v int32) *DescribeServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMirrorResponse) SetBody(v *DescribeServiceMirrorResponseBody) *DescribeServiceMirrorResponse {
	s.Body = v
	return s
}

type DescribeStressResponseBody struct {
	AvailableAgent *int64  `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	CallerUid      *string `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	DesiredAgent   *int64  `json:"DesiredAgent,omitempty" xml:"DesiredAgent,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	ParentUid      *string `json:"ParentUid,omitempty" xml:"ParentUid,omitempty"`
	Reason         *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// Id of the request
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StressName  *string `json:"StressName,omitempty" xml:"StressName,omitempty"`
	Token       *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s DescribeStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeStressResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStressResponseBody) SetAvailableAgent(v int64) *DescribeStressResponseBody {
	s.AvailableAgent = &v
	return s
}

func (s *DescribeStressResponseBody) SetCallerUid(v string) *DescribeStressResponseBody {
	s.CallerUid = &v
	return s
}

func (s *DescribeStressResponseBody) SetDesiredAgent(v int64) *DescribeStressResponseBody {
	s.DesiredAgent = &v
	return s
}

func (s *DescribeStressResponseBody) SetMessage(v string) *DescribeStressResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeStressResponseBody) SetParentUid(v string) *DescribeStressResponseBody {
	s.ParentUid = &v
	return s
}

func (s *DescribeStressResponseBody) SetReason(v string) *DescribeStressResponseBody {
	s.Reason = &v
	return s
}

func (s *DescribeStressResponseBody) SetRequestId(v string) *DescribeStressResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStressResponseBody) SetServiceName(v string) *DescribeStressResponseBody {
	s.ServiceName = &v
	return s
}

func (s *DescribeStressResponseBody) SetStatus(v string) *DescribeStressResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeStressResponseBody) SetStressName(v string) *DescribeStressResponseBody {
	s.StressName = &v
	return s
}

func (s *DescribeStressResponseBody) SetToken(v string) *DescribeStressResponseBody {
	s.Token = &v
	return s
}

type DescribeStressResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeStressResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeStressResponse) GoString() string {
	return s.String()
}

func (s *DescribeStressResponse) SetHeaders(v map[string]*string) *DescribeStressResponse {
	s.Headers = v
	return s
}

func (s *DescribeStressResponse) SetStatusCode(v int32) *DescribeStressResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeStressResponse) SetBody(v *DescribeStressResponseBody) *DescribeStressResponse {
	s.Body = v
	return s
}

type ListBenchmarkTaskResponseBody struct {
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Tasks     []*ListBenchmarkTaskResponseBodyTasks `json:"Tasks,omitempty" xml:"Tasks,omitempty" type:"Repeated"`
}

func (s ListBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponseBody) SetRequestId(v string) *ListBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBenchmarkTaskResponseBody) SetTasks(v []*ListBenchmarkTaskResponseBodyTasks) *ListBenchmarkTaskResponseBody {
	s.Tasks = v
	return s
}

type ListBenchmarkTaskResponseBodyTasks struct {
	AvailableAgent *int64  `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskName       *string `json:"TaskName,omitempty" xml:"TaskName,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListBenchmarkTaskResponseBodyTasks) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskResponseBodyTasks) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetAvailableAgent(v int64) *ListBenchmarkTaskResponseBodyTasks {
	s.AvailableAgent = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetCreateTime(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.CreateTime = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetMessage(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Message = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetRegion(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Region = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetServiceName(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.ServiceName = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetStatus(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.Status = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetTaskId(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.TaskId = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetTaskName(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.TaskName = &v
	return s
}

func (s *ListBenchmarkTaskResponseBodyTasks) SetUpdateTime(v string) *ListBenchmarkTaskResponseBodyTasks {
	s.UpdateTime = &v
	return s
}

type ListBenchmarkTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *ListBenchmarkTaskResponse) SetHeaders(v map[string]*string) *ListBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *ListBenchmarkTaskResponse) SetStatusCode(v int32) *ListBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBenchmarkTaskResponse) SetBody(v *ListBenchmarkTaskResponseBody) *ListBenchmarkTaskResponse {
	s.Body = v
	return s
}

type ListResourceInstanceWorkerRequest struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceInstanceWorkerRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstanceWorkerRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerRequest) SetPageNumber(v int32) *ListResourceInstanceWorkerRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerRequest) SetPageSize(v int32) *ListResourceInstanceWorkerRequest {
	s.PageSize = &v
	return s
}

type ListResourceInstanceWorkerResponseBody struct {
	// 当前页码
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// pod列表
	Pods []*ResourceInstanceWorker `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// pod总数
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceInstanceWorkerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstanceWorkerResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerResponseBody) SetPageNumber(v int32) *ListResourceInstanceWorkerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetPageSize(v int32) *ListResourceInstanceWorkerResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetPods(v []*ResourceInstanceWorker) *ListResourceInstanceWorkerResponseBody {
	s.Pods = v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetRequestId(v string) *ListResourceInstanceWorkerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInstanceWorkerResponseBody) SetTotalCount(v int32) *ListResourceInstanceWorkerResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceInstanceWorkerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceInstanceWorkerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceInstanceWorkerResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstanceWorkerResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstanceWorkerResponse) SetHeaders(v map[string]*string) *ListResourceInstanceWorkerResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstanceWorkerResponse) SetStatusCode(v int32) *ListResourceInstanceWorkerResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceInstanceWorkerResponse) SetBody(v *ListResourceInstanceWorkerResponseBody) *ListResourceInstanceWorkerResponse {
	s.Body = v
	return s
}

type ListResourceInstancesRequest struct {
	// 付费类型
	// PrePaid预付费
	// PostPaid后付费
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// 请求的页码（默认为1）
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的大小（默认为100）
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesRequest) SetChargeType(v string) *ListResourceInstancesRequest {
	s.ChargeType = &v
	return s
}

func (s *ListResourceInstancesRequest) SetPageNumber(v int32) *ListResourceInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesRequest) SetPageSize(v int32) *ListResourceInstancesRequest {
	s.PageSize = &v
	return s
}

type ListResourceInstancesResponseBody struct {
	Instances  []*ResourceInstance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageNumber *int32              `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32              `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponseBody) SetInstances(v []*ResourceInstance) *ListResourceInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListResourceInstancesResponseBody) SetPageNumber(v int32) *ListResourceInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetPageSize(v int32) *ListResourceInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetRequestId(v string) *ListResourceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceInstancesResponseBody) SetTotalCount(v int32) *ListResourceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceInstancesResponse) SetHeaders(v map[string]*string) *ListResourceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceInstancesResponse) SetStatusCode(v int32) *ListResourceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceInstancesResponse) SetBody(v *ListResourceInstancesResponseBody) *ListResourceInstancesResponse {
	s.Body = v
	return s
}

type ListResourceServicesRequest struct {
	// 请求的页码（默认为1）
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的大小（默认为100）
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourceServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourceServicesRequest) GoString() string {
	return s.String()
}

func (s *ListResourceServicesRequest) SetPageNumber(v int32) *ListResourceServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourceServicesRequest) SetPageSize(v int32) *ListResourceServicesRequest {
	s.PageSize = &v
	return s
}

type ListResourceServicesResponseBody struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services   []*Service `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	TotalCount *int32     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourceServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceServicesResponseBody) SetPageNumber(v int32) *ListResourceServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetPageSize(v int32) *ListResourceServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetRequestId(v string) *ListResourceServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetServices(v []*Service) *ListResourceServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListResourceServicesResponseBody) SetTotalCount(v int32) *ListResourceServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourceServicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourceServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourceServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourceServicesResponse) GoString() string {
	return s.String()
}

func (s *ListResourceServicesResponse) SetHeaders(v map[string]*string) *ListResourceServicesResponse {
	s.Headers = v
	return s
}

func (s *ListResourceServicesResponse) SetStatusCode(v int32) *ListResourceServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceServicesResponse) SetBody(v *ListResourceServicesResponseBody) *ListResourceServicesResponse {
	s.Body = v
	return s
}

type ListResourcesRequest struct {
	// 请求的页码（默认为1）
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的大小（默认为100）
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) SetPageNumber(v int32) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

type ListResourcesResponseBody struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources  []*Resource `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	TotalCount *int32      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourcesResponseBody) SetPageNumber(v int32) *ListResourcesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesResponseBody) SetPageSize(v int32) *ListResourcesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourcesResponseBody) SetRequestId(v string) *ListResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourcesResponseBody) SetResources(v []*Resource) *ListResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *ListResourcesResponseBody) SetTotalCount(v int32) *ListResourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListResourcesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListResourcesResponse) SetHeaders(v map[string]*string) *ListResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListResourcesResponse) SetStatusCode(v int32) *ListResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourcesResponse) SetBody(v *ListResourcesResponseBody) *ListResourcesResponse {
	s.Body = v
	return s
}

type ListServiceInstancesRequest struct {
	// 请求的页码（默认为1）
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的大小（默认为100）
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListServiceInstancesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesRequest) SetPageNumber(v int32) *ListServiceInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServiceInstancesRequest) SetPageSize(v int32) *ListServiceInstancesRequest {
	s.PageSize = &v
	return s
}

type ListServiceInstancesResponseBody struct {
	Instances  []*Instance `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	PageNumber *int32      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceInstancesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponseBody) SetInstances(v []*Instance) *ListServiceInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListServiceInstancesResponseBody) SetPageNumber(v int32) *ListServiceInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetPageSize(v int32) *ListServiceInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetRequestId(v string) *ListServiceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceInstancesResponseBody) SetTotalCount(v int32) *ListServiceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServiceInstancesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServiceInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServiceInstancesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServiceInstancesResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstancesResponse) SetHeaders(v map[string]*string) *ListServiceInstancesResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstancesResponse) SetStatusCode(v int32) *ListServiceInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstancesResponse) SetBody(v *ListServiceInstancesResponseBody) *ListServiceInstancesResponse {
	s.Body = v
	return s
}

type ListServicesRequest struct {
	// 模糊匹配字段（只支持按服务名字模糊匹配）
	Filter *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	// 排序方式（默认降序）
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 请求的页码（默认为1）
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页的大小（默认为100）
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 排序字段 （时间戳类型默认倒序排序）
	Sort *string `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ListServicesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListServicesRequest) GoString() string {
	return s.String()
}

func (s *ListServicesRequest) SetFilter(v string) *ListServicesRequest {
	s.Filter = &v
	return s
}

func (s *ListServicesRequest) SetOrder(v string) *ListServicesRequest {
	s.Order = &v
	return s
}

func (s *ListServicesRequest) SetPageNumber(v int32) *ListServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListServicesRequest) SetPageSize(v int32) *ListServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListServicesRequest) SetSort(v string) *ListServicesRequest {
	s.Sort = &v
	return s
}

type ListServicesResponseBody struct {
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	RequestId  *string    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services   []*Service `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	TotalCount *int32     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) SetPageNumber(v int32) *ListServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListServicesResponseBody) SetPageSize(v int32) *ListServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) SetServices(v []*Service) *ListServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListServicesResponseBody) SetTotalCount(v int32) *ListServicesResponseBody {
	s.TotalCount = &v
	return s
}

type ListServicesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListServicesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListServicesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListServicesResponse) GoString() string {
	return s.String()
}

func (s *ListServicesResponse) SetHeaders(v map[string]*string) *ListServicesResponse {
	s.Headers = v
	return s
}

func (s *ListServicesResponse) SetStatusCode(v int32) *ListServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServicesResponse) SetBody(v *ListServicesResponseBody) *ListServicesResponse {
	s.Body = v
	return s
}

type ListStressesResponseBody struct {
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Stresses  []*ListStressesResponseBodyStresses `json:"Stresses,omitempty" xml:"Stresses,omitempty" type:"Repeated"`
}

func (s ListStressesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListStressesResponseBody) GoString() string {
	return s.String()
}

func (s *ListStressesResponseBody) SetRequestId(v string) *ListStressesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListStressesResponseBody) SetStresses(v []*ListStressesResponseBodyStresses) *ListStressesResponseBody {
	s.Stresses = v
	return s
}

type ListStressesResponseBodyStresses struct {
	AvailableAgent *int64  `json:"AvailableAgent,omitempty" xml:"AvailableAgent,omitempty"`
	CreateTime     *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Region         *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ServiceName    *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	Status         *string `json:"Status,omitempty" xml:"Status,omitempty"`
	StressName     *string `json:"StressName,omitempty" xml:"StressName,omitempty"`
	UpdateTime     *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListStressesResponseBodyStresses) String() string {
	return tea.Prettify(s)
}

func (s ListStressesResponseBodyStresses) GoString() string {
	return s.String()
}

func (s *ListStressesResponseBodyStresses) SetAvailableAgent(v int64) *ListStressesResponseBodyStresses {
	s.AvailableAgent = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetCreateTime(v string) *ListStressesResponseBodyStresses {
	s.CreateTime = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetMessage(v string) *ListStressesResponseBodyStresses {
	s.Message = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetRegion(v string) *ListStressesResponseBodyStresses {
	s.Region = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetServiceName(v string) *ListStressesResponseBodyStresses {
	s.ServiceName = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetStatus(v string) *ListStressesResponseBodyStresses {
	s.Status = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetStressName(v string) *ListStressesResponseBodyStresses {
	s.StressName = &v
	return s
}

func (s *ListStressesResponseBodyStresses) SetUpdateTime(v string) *ListStressesResponseBodyStresses {
	s.UpdateTime = &v
	return s
}

type ListStressesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListStressesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListStressesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListStressesResponse) GoString() string {
	return s.String()
}

func (s *ListStressesResponse) SetHeaders(v map[string]*string) *ListStressesResponse {
	s.Headers = v
	return s
}

func (s *ListStressesResponse) SetStatusCode(v int32) *ListStressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListStressesResponse) SetBody(v *ListStressesResponseBody) *ListStressesResponse {
	s.Body = v
	return s
}

type ReleaseServiceRequest struct {
	// 灰度权重，范围 [0, 100]
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ReleaseServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseServiceRequest) SetWeight(v int32) *ReleaseServiceRequest {
	s.Weight = &v
	return s
}

type ReleaseServiceResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponseBody) SetMessage(v string) *ReleaseServiceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseServiceResponseBody) SetRequestId(v string) *ReleaseServiceResponseBody {
	s.RequestId = &v
	return s
}

type ReleaseServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReleaseServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReleaseServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s ReleaseServiceResponse) GoString() string {
	return s.String()
}

func (s *ReleaseServiceResponse) SetHeaders(v map[string]*string) *ReleaseServiceResponse {
	s.Headers = v
	return s
}

func (s *ReleaseServiceResponse) SetStatusCode(v int32) *ReleaseServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseServiceResponse) SetBody(v *ReleaseServiceResponseBody) *ReleaseServiceResponse {
	s.Body = v
	return s
}

type ReportStressResponseBody struct {
	ReportUrl *string `json:"ReportUrl,omitempty" xml:"ReportUrl,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReportStressResponseBody) GoString() string {
	return s.String()
}

func (s *ReportStressResponseBody) SetReportUrl(v string) *ReportStressResponseBody {
	s.ReportUrl = &v
	return s
}

func (s *ReportStressResponseBody) SetRequestId(v string) *ReportStressResponseBody {
	s.RequestId = &v
	return s
}

type ReportStressResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ReportStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReportStressResponse) String() string {
	return tea.Prettify(s)
}

func (s ReportStressResponse) GoString() string {
	return s.String()
}

func (s *ReportStressResponse) SetHeaders(v map[string]*string) *ReportStressResponse {
	s.Headers = v
	return s
}

func (s *ReportStressResponse) SetStatusCode(v int32) *ReportStressResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportStressResponse) SetBody(v *ReportStressResponseBody) *ReportStressResponse {
	s.Body = v
	return s
}

type StartBenchmarkTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskResponseBody) SetMessage(v string) *StartBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StartBenchmarkTaskResponseBody) SetRequestId(v string) *StartBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type StartBenchmarkTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StartBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *StartBenchmarkTaskResponse) SetHeaders(v map[string]*string) *StartBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *StartBenchmarkTaskResponse) SetStatusCode(v int32) *StartBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBenchmarkTaskResponse) SetBody(v *StartBenchmarkTaskResponseBody) *StartBenchmarkTaskResponse {
	s.Body = v
	return s
}

type StartServiceResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StartServiceResponseBody) SetMessage(v string) *StartServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StartServiceResponseBody) SetRequestId(v string) *StartServiceResponseBody {
	s.RequestId = &v
	return s
}

type StartServiceResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StartServiceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceResponse) SetHeaders(v map[string]*string) *StartServiceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceResponse) SetStatusCode(v int32) *StartServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceResponse) SetBody(v *StartServiceResponseBody) *StartServiceResponse {
	s.Body = v
	return s
}

type StartStressResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartStressResponseBody) GoString() string {
	return s.String()
}

func (s *StartStressResponseBody) SetMessage(v string) *StartStressResponseBody {
	s.Message = &v
	return s
}

func (s *StartStressResponseBody) SetRequestId(v string) *StartStressResponseBody {
	s.RequestId = &v
	return s
}

type StartStressResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartStressResponse) String() string {
	return tea.Prettify(s)
}

func (s StartStressResponse) GoString() string {
	return s.String()
}

func (s *StartStressResponse) SetHeaders(v map[string]*string) *StartStressResponse {
	s.Headers = v
	return s
}

func (s *StartStressResponse) SetStatusCode(v int32) *StartStressResponse {
	s.StatusCode = &v
	return s
}

func (s *StartStressResponse) SetBody(v *StartStressResponseBody) *StartStressResponse {
	s.Body = v
	return s
}

type StopBenchmarkTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskResponseBody) SetMessage(v string) *StopBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopBenchmarkTaskResponseBody) SetRequestId(v string) *StopBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type StopBenchmarkTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s StopBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *StopBenchmarkTaskResponse) SetHeaders(v map[string]*string) *StopBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *StopBenchmarkTaskResponse) SetStatusCode(v int32) *StopBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StopBenchmarkTaskResponse) SetBody(v *StopBenchmarkTaskResponseBody) *StopBenchmarkTaskResponse {
	s.Body = v
	return s
}

type StopServiceResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponseBody) GoString() string {
	return s.String()
}

func (s *StopServiceResponseBody) SetMessage(v string) *StopServiceResponseBody {
	s.Message = &v
	return s
}

func (s *StopServiceResponseBody) SetRequestId(v string) *StopServiceResponseBody {
	s.RequestId = &v
	return s
}

type StopServiceResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s StopServiceResponse) GoString() string {
	return s.String()
}

func (s *StopServiceResponse) SetHeaders(v map[string]*string) *StopServiceResponse {
	s.Headers = v
	return s
}

func (s *StopServiceResponse) SetStatusCode(v int32) *StopServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *StopServiceResponse) SetBody(v *StopServiceResponseBody) *StopServiceResponse {
	s.Body = v
	return s
}

type StopStressResponseBody struct {
	Code      *int64  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopStressResponseBody) GoString() string {
	return s.String()
}

func (s *StopStressResponseBody) SetCode(v int64) *StopStressResponseBody {
	s.Code = &v
	return s
}

func (s *StopStressResponseBody) SetMessage(v string) *StopStressResponseBody {
	s.Message = &v
	return s
}

func (s *StopStressResponseBody) SetRequestId(v string) *StopStressResponseBody {
	s.RequestId = &v
	return s
}

type StopStressResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopStressResponse) String() string {
	return tea.Prettify(s)
}

func (s StopStressResponse) GoString() string {
	return s.String()
}

func (s *StopStressResponse) SetHeaders(v map[string]*string) *StopStressResponse {
	s.Headers = v
	return s
}

func (s *StopStressResponse) SetStatusCode(v int32) *StopStressResponse {
	s.StatusCode = &v
	return s
}

func (s *StopStressResponse) SetBody(v *StopStressResponseBody) *StopStressResponse {
	s.Body = v
	return s
}

type UpdateBenchmarkTaskRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBenchmarkTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBenchmarkTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskRequest) SetBody(v string) *UpdateBenchmarkTaskRequest {
	s.Body = &v
	return s
}

type UpdateBenchmarkTaskResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateBenchmarkTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateBenchmarkTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskResponseBody) SetMessage(v string) *UpdateBenchmarkTaskResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBenchmarkTaskResponseBody) SetRequestId(v string) *UpdateBenchmarkTaskResponseBody {
	s.RequestId = &v
	return s
}

type UpdateBenchmarkTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateBenchmarkTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateBenchmarkTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBenchmarkTaskResponse) GoString() string {
	return s.String()
}

func (s *UpdateBenchmarkTaskResponse) SetHeaders(v map[string]*string) *UpdateBenchmarkTaskResponse {
	s.Headers = v
	return s
}

func (s *UpdateBenchmarkTaskResponse) SetStatusCode(v int32) *UpdateBenchmarkTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBenchmarkTaskResponse) SetBody(v *UpdateBenchmarkTaskResponseBody) *UpdateBenchmarkTaskResponse {
	s.Body = v
	return s
}

type UpdateResourceRequest struct {
	// 新的资源组名称
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s UpdateResourceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRequest) SetResourceName(v string) *UpdateResourceRequest {
	s.ResourceName = &v
	return s
}

type UpdateResourceResponseBody struct {
	// Id of the request
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceId   *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
}

func (s UpdateResourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponseBody) SetRequestId(v string) *UpdateResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetResourceId(v string) *UpdateResourceResponseBody {
	s.ResourceId = &v
	return s
}

func (s *UpdateResourceResponseBody) SetResourceName(v string) *UpdateResourceResponseBody {
	s.ResourceName = &v
	return s
}

type UpdateResourceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceResponse) SetHeaders(v map[string]*string) *UpdateResourceResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceResponse) SetStatusCode(v int32) *UpdateResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceResponse) SetBody(v *UpdateResourceResponseBody) *UpdateResourceResponse {
	s.Body = v
	return s
}

type UpdateResourceDLinkRequest struct {
	// 要打通的客户端的网段信息，会将该网段加入到服务端的回包路由中，与VSwitchIdList可二选一
	DestinationCIDRs *string `json:"DestinationCIDRs,omitempty" xml:"DestinationCIDRs,omitempty"`
	// 客户端ECS归属的安全组
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 对端的主VSwitchID，会在该vswitch中创建ENI
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 要打通的客户端的vswitch列表，会将这些vswitch对应的网段加入到服务端的回包路由中
	VSwitchIdList []*string `json:"VSwitchIdList,omitempty" xml:"VSwitchIdList,omitempty" type:"Repeated"`
}

func (s UpdateResourceDLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceDLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkRequest) SetDestinationCIDRs(v string) *UpdateResourceDLinkRequest {
	s.DestinationCIDRs = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetSecurityGroupId(v string) *UpdateResourceDLinkRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetVSwitchId(v string) *UpdateResourceDLinkRequest {
	s.VSwitchId = &v
	return s
}

func (s *UpdateResourceDLinkRequest) SetVSwitchIdList(v []*string) *UpdateResourceDLinkRequest {
	s.VSwitchIdList = v
	return s
}

type UpdateResourceDLinkResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateResourceDLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceDLinkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkResponseBody) SetMessage(v string) *UpdateResourceDLinkResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateResourceDLinkResponseBody) SetRequestId(v string) *UpdateResourceDLinkResponseBody {
	s.RequestId = &v
	return s
}

type UpdateResourceDLinkResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateResourceDLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateResourceDLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateResourceDLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateResourceDLinkResponse) SetHeaders(v map[string]*string) *UpdateResourceDLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateResourceDLinkResponse) SetStatusCode(v int32) *UpdateResourceDLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateResourceDLinkResponse) SetBody(v *UpdateResourceDLinkResponseBody) *UpdateResourceDLinkResponse {
	s.Body = v
	return s
}

type UpdateServiceRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRequest) SetBody(v string) *UpdateServiceRequest {
	s.Body = &v
	return s
}

type UpdateServiceResponseBody struct {
	// 请求返回消息。
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID。
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponseBody) SetMessage(v string) *UpdateServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceResponseBody) SetRequestId(v string) *UpdateServiceResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceResponse) SetHeaders(v map[string]*string) *UpdateServiceResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceResponse) SetStatusCode(v int32) *UpdateServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceResponse) SetBody(v *UpdateServiceResponseBody) *UpdateServiceResponse {
	s.Body = v
	return s
}

type UpdateServiceAutoScalerRequest struct {
	// 最大 replica 数，需要大于MinReplica
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// 最小 replica 数，需要大于0
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// map 类型的策略定义
	Strategies *UpdateServiceAutoScalerRequestStrategies `json:"Strategies,omitempty" xml:"Strategies,omitempty" type:"Struct"`
}

func (s UpdateServiceAutoScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequest) SetMax(v int32) *UpdateServiceAutoScalerRequest {
	s.Max = &v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetMin(v int32) *UpdateServiceAutoScalerRequest {
	s.Min = &v
	return s
}

func (s *UpdateServiceAutoScalerRequest) SetStrategies(v *UpdateServiceAutoScalerRequestStrategies) *UpdateServiceAutoScalerRequest {
	s.Strategies = v
	return s
}

type UpdateServiceAutoScalerRequestStrategies struct {
	// 最大 replica 数，需要大于MinReplica
	Cpu *float32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// 每个实例支持的最大qps数，超出即扩容
	Qps *float32 `json:"Qps,omitempty" xml:"Qps,omitempty"`
}

func (s UpdateServiceAutoScalerRequestStrategies) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerRequestStrategies) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerRequestStrategies) SetCpu(v float32) *UpdateServiceAutoScalerRequestStrategies {
	s.Cpu = &v
	return s
}

func (s *UpdateServiceAutoScalerRequestStrategies) SetQps(v float32) *UpdateServiceAutoScalerRequestStrategies {
	s.Qps = &v
	return s
}

type UpdateServiceAutoScalerResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceAutoScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerResponseBody) SetMessage(v string) *UpdateServiceAutoScalerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceAutoScalerResponseBody) SetRequestId(v string) *UpdateServiceAutoScalerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceAutoScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceAutoScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceAutoScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceAutoScalerResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceAutoScalerResponse) SetHeaders(v map[string]*string) *UpdateServiceAutoScalerResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceAutoScalerResponse) SetStatusCode(v int32) *UpdateServiceAutoScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceAutoScalerResponse) SetBody(v *UpdateServiceAutoScalerResponseBody) *UpdateServiceAutoScalerResponse {
	s.Body = v
	return s
}

type UpdateServiceCronScalerRequest struct {
	// 需要排除的时间点的cron表达式
	ExcludeDates []*string `json:"ExcludeDates,omitempty" xml:"ExcludeDates,omitempty" type:"Repeated"`
	// 定时伸缩任务描述
	ScaleJobs []*UpdateServiceCronScalerRequestScaleJobs `json:"ScaleJobs,omitempty" xml:"ScaleJobs,omitempty" type:"Repeated"`
}

func (s UpdateServiceCronScalerRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerRequest) SetExcludeDates(v []*string) *UpdateServiceCronScalerRequest {
	s.ExcludeDates = v
	return s
}

func (s *UpdateServiceCronScalerRequest) SetScaleJobs(v []*UpdateServiceCronScalerRequestScaleJobs) *UpdateServiceCronScalerRequest {
	s.ScaleJobs = v
	return s
}

type UpdateServiceCronScalerRequestScaleJobs struct {
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 要执行伸缩任务的cron表达式
	Schedule *string `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
	// 执行伸缩任务的目标replica
	TargetSize *int32 `json:"TargetSize,omitempty" xml:"TargetSize,omitempty"`
}

func (s UpdateServiceCronScalerRequestScaleJobs) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerRequestScaleJobs) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetName(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.Name = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetSchedule(v string) *UpdateServiceCronScalerRequestScaleJobs {
	s.Schedule = &v
	return s
}

func (s *UpdateServiceCronScalerRequestScaleJobs) SetTargetSize(v int32) *UpdateServiceCronScalerRequestScaleJobs {
	s.TargetSize = &v
	return s
}

type UpdateServiceCronScalerResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceCronScalerResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerResponseBody) SetMessage(v string) *UpdateServiceCronScalerResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceCronScalerResponseBody) SetRequestId(v string) *UpdateServiceCronScalerResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceCronScalerResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceCronScalerResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceCronScalerResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceCronScalerResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceCronScalerResponse) SetHeaders(v map[string]*string) *UpdateServiceCronScalerResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceCronScalerResponse) SetStatusCode(v int32) *UpdateServiceCronScalerResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceCronScalerResponse) SetBody(v *UpdateServiceCronScalerResponseBody) *UpdateServiceCronScalerResponse {
	s.Body = v
	return s
}

type UpdateServiceMirrorRequest struct {
	// 比例 [0, 100]
	Ratio *int32 `json:"Ratio,omitempty" xml:"Ratio,omitempty"`
	// 服务实例列表
	Target []*string `json:"Target,omitempty" xml:"Target,omitempty" type:"Repeated"`
}

func (s UpdateServiceMirrorRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMirrorRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorRequest) SetRatio(v int32) *UpdateServiceMirrorRequest {
	s.Ratio = &v
	return s
}

func (s *UpdateServiceMirrorRequest) SetTarget(v []*string) *UpdateServiceMirrorRequest {
	s.Target = v
	return s
}

type UpdateServiceMirrorResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceMirrorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMirrorResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorResponseBody) SetMessage(v string) *UpdateServiceMirrorResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceMirrorResponseBody) SetRequestId(v string) *UpdateServiceMirrorResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceMirrorResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceMirrorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceMirrorResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceMirrorResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceMirrorResponse) SetHeaders(v map[string]*string) *UpdateServiceMirrorResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceMirrorResponse) SetStatusCode(v int32) *UpdateServiceMirrorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceMirrorResponse) SetBody(v *UpdateServiceMirrorResponseBody) *UpdateServiceMirrorResponse {
	s.Body = v
	return s
}

type UpdateServiceVersionRequest struct {
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UpdateServiceVersionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionRequest) SetVersion(v int32) *UpdateServiceVersionRequest {
	s.Version = &v
	return s
}

type UpdateServiceVersionResponseBody struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceVersionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponseBody) SetMessage(v string) *UpdateServiceVersionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceVersionResponseBody) SetRequestId(v string) *UpdateServiceVersionResponseBody {
	s.RequestId = &v
	return s
}

type UpdateServiceVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateServiceVersionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponse) SetHeaders(v map[string]*string) *UpdateServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceVersionResponse) SetStatusCode(v int32) *UpdateServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceVersionResponse) SetBody(v *UpdateServiceVersionResponseBody) *UpdateServiceVersionResponse {
	s.Body = v
	return s
}

type UpdateStressRequest struct {
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStressRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStressRequest) GoString() string {
	return s.String()
}

func (s *UpdateStressRequest) SetBody(v string) *UpdateStressRequest {
	s.Body = &v
	return s
}

type UpdateStressResponseBody struct {
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateStressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStressResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStressResponseBody) SetMessage(v string) *UpdateStressResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateStressResponseBody) SetRequestId(v string) *UpdateStressResponseBody {
	s.RequestId = &v
	return s
}

type UpdateStressResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateStressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateStressResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStressResponse) GoString() string {
	return s.String()
}

func (s *UpdateStressResponse) SetHeaders(v map[string]*string) *UpdateStressResponse {
	s.Headers = v
	return s
}

func (s *UpdateStressResponse) SetStatusCode(v int32) *UpdateStressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStressResponse) SetBody(v *UpdateStressResponseBody) *UpdateStressResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("regional")
	client.EndpointMap = map[string]*string{
		"cn-beijing":            tea.String("pai-eas.cn-beijing.aliyuncs.com"),
		"cn-zhangjiakou":        tea.String("pai-eas.cn-zhangjiakou.aliyuncs.com"),
		"cn-hangzhou":           tea.String("pai-eas.cn-hangzhou.aliyuncs.com"),
		"cn-shanghai":           tea.String("pai-eas.cn-shanghai.aliyuncs.com"),
		"cn-shenzhen":           tea.String("pai-eas.cn-shenzhen.aliyuncs.com"),
		"cn-hongkong":           tea.String("pai-eas.cn-hongkong.aliyuncs.com"),
		"ap-southeast-1":        tea.String("pai-eas.ap-southeast-1.aliyuncs.com"),
		"ap-southeast-5":        tea.String("pai-eas.ap-southeast-5.aliyuncs.com"),
		"us-east-1":             tea.String("pai-eas.us-east-1.aliyuncs.com"),
		"us-west-1":             tea.String("pai-eas.us-west-1.aliyuncs.com"),
		"eu-central-1":          tea.String("pai-eas.eu-central-1.aliyuncs.com"),
		"ap-south-1":            tea.String("pai-eas.ap-south-1.aliyuncs.com"),
		"cn-shanghai-finance-1": tea.String("pai-eas.cn-shanghai-finance-1.aliyuncs.com"),
		"cn-north-2-gov-1":      tea.String("pai-eas.cn-north-2-gov-1.aliyuncs.com"),
		"cn-chengdu":            tea.String("pai-eas.cn-chengdu.aliyuncs.com"),
	}
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eas"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBenchmarkTask(request *CreateBenchmarkTaskRequest) (_result *CreateBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBenchmarkTaskResponse{}
	_body, _err := client.CreateBenchmarkTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateBenchmarkTaskWithOptions(request *CreateBenchmarkTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBenchmarkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResource(request *CreateResourceRequest) (_result *CreateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceResponse{}
	_body, _err := client.CreateResourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceWithOptions(request *CreateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceCount)) {
		body["EcsInstanceCount"] = request.EcsInstanceCount
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceType)) {
		body["EcsInstanceType"] = request.EcsInstanceType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourceInstances(ClusterId *string, ResourceId *string, request *CreateResourceInstancesRequest) (_result *CreateResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceInstancesResponse{}
	_body, _err := client.CreateResourceInstancesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceInstancesWithOptions(ClusterId *string, ResourceId *string, request *CreateResourceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRenewal)) {
		body["AutoRenewal"] = request.AutoRenewal
	}

	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		body["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceCount)) {
		body["EcsInstanceCount"] = request.EcsInstanceCount
	}

	if !tea.BoolValue(util.IsUnset(request.EcsInstanceType)) {
		body["EcsInstanceType"] = request.EcsInstanceType
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/instances"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateResourceLog(ClusterId *string, ResourceId *string, request *CreateResourceLogRequest) (_result *CreateResourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateResourceLogResponse{}
	_body, _err := client.CreateResourceLogWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateResourceLogWithOptions(ClusterId *string, ResourceId *string, request *CreateResourceLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateResourceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.LogStore)) {
		body["LogStore"] = request.LogStore
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateResourceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/log"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateResourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateService(request *CreateServiceRequest) (_result *CreateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceResponse{}
	_body, _err := client.CreateServiceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceWithOptions(request *CreateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceAutoScaler(ClusterId *string, ServiceName *string, request *CreateServiceAutoScalerRequest) (_result *CreateServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceAutoScalerResponse{}
	_body, _err := client.CreateServiceAutoScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, request *CreateServiceAutoScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceAutoScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Max)) {
		body["Max"] = request.Max
	}

	if !tea.BoolValue(util.IsUnset(request.Min)) {
		body["Min"] = request.Min
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Strategies))) {
		body["Strategies"] = request.Strategies
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/autoscaler"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceCronScaler(ClusterId *string, ServiceName *string, request *CreateServiceCronScalerRequest) (_result *CreateServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceCronScalerResponse{}
	_body, _err := client.CreateServiceCronScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, request *CreateServiceCronScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceCronScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeDates)) {
		body["ExcludeDates"] = request.ExcludeDates
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleJobs)) {
		body["ScaleJobs"] = request.ScaleJobs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/cronscaler"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateServiceMirror(ClusterId *string, ServiceName *string, request *CreateServiceMirrorRequest) (_result *CreateServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateServiceMirrorResponse{}
	_body, _err := client.CreateServiceMirrorWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateServiceMirrorWithOptions(ClusterId *string, ServiceName *string, request *CreateServiceMirrorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateServiceMirrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ratio)) {
		body["Ratio"] = request.Ratio
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		body["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/mirror"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateStress(request *CreateStressRequest) (_result *CreateStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateStressResponse{}
	_body, _err := client.CreateStressWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateStressWithOptions(request *CreateStressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateStressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBenchmarkTask(ClusterId *string, TaskName *string) (_result *DeleteBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBenchmarkTaskResponse{}
	_body, _err := client.DeleteBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBenchmarkTaskResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	TaskName = openapiutil.GetEncodeParam(TaskName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(TaskName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResource(ClusterId *string, ResourceId *string) (_result *DeleteResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceResponse{}
	_body, _err := client.DeleteResourceWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceDLink(ClusterId *string, ResourceId *string) (_result *DeleteResourceDLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceDLinkResponse{}
	_body, _err := client.DeleteResourceDLinkWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceDLinkWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceDLinkResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceDLink"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/dlink"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceDLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceInstances(ClusterId *string, ResourceId *string, request *DeleteResourceInstancesRequest) (_result *DeleteResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceInstancesResponse{}
	_body, _err := client.DeleteResourceInstancesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceInstancesWithOptions(ClusterId *string, ResourceId *string, request *DeleteResourceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AllFailed)) {
		query["AllFailed"] = request.AllFailed
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceList)) {
		query["InstanceList"] = request.InstanceList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/instances"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteResourceLog(ClusterId *string, ResourceId *string) (_result *DeleteResourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteResourceLogResponse{}
	_body, _err := client.DeleteResourceLogWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteResourceLogWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteResourceLogResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteResourceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/log"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteResourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteService(ClusterId *string, ServiceName *string) (_result *DeleteServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceResponse{}
	_body, _err := client.DeleteServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceAutoScaler(ClusterId *string, ServiceName *string) (_result *DeleteServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceAutoScalerResponse{}
	_body, _err := client.DeleteServiceAutoScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceAutoScalerResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/autoscaler"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceCronScaler(ClusterId *string, ServiceName *string) (_result *DeleteServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceCronScalerResponse{}
	_body, _err := client.DeleteServiceCronScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceCronScalerResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/cronscaler"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceInstances(ClusterId *string, ServiceName *string, request *DeleteServiceInstancesRequest) (_result *DeleteServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.DeleteServiceInstancesWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceInstancesWithOptions(ClusterId *string, ServiceName *string, request *DeleteServiceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InstanceList)) {
		query["InstanceList"] = request.InstanceList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/instances"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteServiceMirror(ClusterId *string, ServiceName *string) (_result *DeleteServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteServiceMirrorResponse{}
	_body, _err := client.DeleteServiceMirrorWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteServiceMirrorWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteServiceMirrorResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/mirror"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteStress(ClusterId *string, StressName *string) (_result *DeleteStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteStressResponse{}
	_body, _err := client.DeleteStressWithOptions(ClusterId, StressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteStressWithOptions(ClusterId *string, StressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteStressResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	StressName = openapiutil.GetEncodeParam(StressName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(StressName)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBenchmarkTask(ClusterId *string, TaskName *string) (_result *DescribeBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeBenchmarkTaskResponse{}
	_body, _err := client.DescribeBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeBenchmarkTaskResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	TaskName = openapiutil.GetEncodeParam(TaskName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(TaskName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeBenchmarkTaskReport(ClusterId *string, TaskName *string) (_result *DescribeBenchmarkTaskReportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeBenchmarkTaskReportResponse{}
	_body, _err := client.DescribeBenchmarkTaskReportWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeBenchmarkTaskReportWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeBenchmarkTaskReportResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	TaskName = openapiutil.GetEncodeParam(TaskName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeBenchmarkTaskReport"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(TaskName) + "/report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeBenchmarkTaskReportResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResource(ClusterId *string, ResourceId *string) (_result *DescribeResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceResponse{}
	_body, _err := client.DescribeResourceWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourceResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceDLink(ClusterId *string, ResourceId *string) (_result *DescribeResourceDLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceDLinkResponse{}
	_body, _err := client.DescribeResourceDLinkWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceDLinkWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourceDLinkResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceDLink"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/dlink"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceDLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeResourceLog(ClusterId *string, ResourceId *string) (_result *DescribeResourceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeResourceLogResponse{}
	_body, _err := client.DescribeResourceLogWithOptions(ClusterId, ResourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeResourceLogWithOptions(ClusterId *string, ResourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeResourceLogResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeResourceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/log"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeResourceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeService(ClusterId *string, ServiceName *string) (_result *DescribeServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceResponse{}
	_body, _err := client.DescribeServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceAutoScaler(ClusterId *string, ServiceName *string) (_result *DescribeServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceAutoScalerResponse{}
	_body, _err := client.DescribeServiceAutoScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceAutoScalerResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/autoscaler"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceCronScaler(ClusterId *string, ServiceName *string) (_result *DescribeServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceCronScalerResponse{}
	_body, _err := client.DescribeServiceCronScalerWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceCronScalerResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/cronscaler"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceLog(ClusterId *string, ServiceName *string, request *DescribeServiceLogRequest) (_result *DescribeServiceLogResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceLogResponse{}
	_body, _err := client.DescribeServiceLogWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceLogWithOptions(ClusterId *string, ServiceName *string, request *DescribeServiceLogRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceLogResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.Ip)) {
		query["Ip"] = request.Ip
	}

	if !tea.BoolValue(util.IsUnset(request.Keyword)) {
		query["Keyword"] = request.Keyword
	}

	if !tea.BoolValue(util.IsUnset(request.PageNum)) {
		query["PageNum"] = request.PageNum
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceLog"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceLogResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeServiceMirror(ClusterId *string, ServiceName *string) (_result *DescribeServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeServiceMirrorResponse{}
	_body, _err := client.DescribeServiceMirrorWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeServiceMirrorWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeServiceMirrorResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/mirror"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeStress(ClusterId *string, StressName *string) (_result *DescribeStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeStressResponse{}
	_body, _err := client.DescribeStressWithOptions(ClusterId, StressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeStressWithOptions(ClusterId *string, StressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeStressResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	StressName = openapiutil.GetEncodeParam(StressName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(StressName)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBenchmarkTask() (_result *ListBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBenchmarkTaskResponse{}
	_body, _err := client.ListBenchmarkTaskWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBenchmarkTaskWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBenchmarkTaskResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceInstanceWorker(ClusterId *string, ResourceId *string, InstanceName *string, request *ListResourceInstanceWorkerRequest) (_result *ListResourceInstanceWorkerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceInstanceWorkerResponse{}
	_body, _err := client.ListResourceInstanceWorkerWithOptions(ClusterId, ResourceId, InstanceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceInstanceWorkerWithOptions(ClusterId *string, ResourceId *string, InstanceName *string, request *ListResourceInstanceWorkerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceInstanceWorkerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	InstanceName = openapiutil.GetEncodeParam(InstanceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceInstanceWorker"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/instance/" + tea.StringValue(InstanceName) + "/workers"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceInstanceWorkerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceInstances(ClusterId *string, ResourceId *string, request *ListResourceInstancesRequest) (_result *ListResourceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.ListResourceInstancesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceInstancesWithOptions(ClusterId *string, ResourceId *string, request *ListResourceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ChargeType)) {
		query["ChargeType"] = request.ChargeType
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResourceServices(ClusterId *string, ResourceId *string, request *ListResourceServicesRequest) (_result *ListResourceServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourceServicesResponse{}
	_body, _err := client.ListResourceServicesWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourceServicesWithOptions(ClusterId *string, ResourceId *string, request *ListResourceServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourceServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResourceServices"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourceServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListResources(request *ListResourcesRequest) (_result *ListResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListResourcesResponse{}
	_body, _err := client.ListResourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListResourcesWithOptions(request *ListResourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListResources"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServiceInstances(ClusterId *string, ServiceName *string, request *ListServiceInstancesRequest) (_result *ListServiceInstancesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.ListServiceInstancesWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServiceInstancesWithOptions(ClusterId *string, ServiceName *string, request *ListServiceInstancesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServiceInstancesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServiceInstances"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/instances"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServiceInstancesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListServices(request *ListServicesRequest) (_result *ListServicesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListServicesResponse{}
	_body, _err := client.ListServicesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListServicesWithOptions(request *ListServicesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListServicesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Filter)) {
		query["Filter"] = request.Filter
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		query["Sort"] = request.Sort
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListServices"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListServicesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListStresses() (_result *ListStressesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListStressesResponse{}
	_body, _err := client.ListStressesWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListStressesWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListStressesResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ListStresses"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListStressesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReleaseService(ClusterId *string, ServiceName *string, request *ReleaseServiceRequest) (_result *ReleaseServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReleaseServiceResponse{}
	_body, _err := client.ReleaseServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReleaseServiceWithOptions(ClusterId *string, ServiceName *string, request *ReleaseServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReleaseServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Weight)) {
		body["Weight"] = request.Weight
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReleaseService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/release"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReleaseServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReportStress(ClusterId *string, StressName *string) (_result *ReportStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReportStressResponse{}
	_body, _err := client.ReportStressWithOptions(ClusterId, StressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReportStressWithOptions(ClusterId *string, StressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReportStressResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	StressName = openapiutil.GetEncodeParam(StressName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("ReportStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(StressName) + "/report"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ReportStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartBenchmarkTask(ClusterId *string, TaskName *string) (_result *StartBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartBenchmarkTaskResponse{}
	_body, _err := client.StartBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartBenchmarkTaskResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	TaskName = openapiutil.GetEncodeParam(TaskName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(TaskName) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartService(ClusterId *string, ServiceName *string) (_result *StartServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartServiceResponse{}
	_body, _err := client.StartServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartServiceResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartStress(ClusterId *string, StressName *string) (_result *StartStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartStressResponse{}
	_body, _err := client.StartStressWithOptions(ClusterId, StressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartStressWithOptions(ClusterId *string, StressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartStressResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	StressName = openapiutil.GetEncodeParam(StressName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StartStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(StressName) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopBenchmarkTask(ClusterId *string, TaskName *string) (_result *StopBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopBenchmarkTaskResponse{}
	_body, _err := client.StopBenchmarkTaskWithOptions(ClusterId, TaskName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopBenchmarkTaskResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	TaskName = openapiutil.GetEncodeParam(TaskName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(TaskName) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopService(ClusterId *string, ServiceName *string) (_result *StopServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopServiceResponse{}
	_body, _err := client.StopServiceWithOptions(ClusterId, ServiceName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopServiceWithOptions(ClusterId *string, ServiceName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopServiceResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopStress(ClusterId *string, StressName *string) (_result *StopStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopStressResponse{}
	_body, _err := client.StopStressWithOptions(ClusterId, StressName, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopStressWithOptions(ClusterId *string, StressName *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopStressResponse, _err error) {
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	StressName = openapiutil.GetEncodeParam(StressName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("StopStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(StressName) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateBenchmarkTask(ClusterId *string, TaskName *string, request *UpdateBenchmarkTaskRequest) (_result *UpdateBenchmarkTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBenchmarkTaskResponse{}
	_body, _err := client.UpdateBenchmarkTaskWithOptions(ClusterId, TaskName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateBenchmarkTaskWithOptions(ClusterId *string, TaskName *string, request *UpdateBenchmarkTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBenchmarkTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	TaskName = openapiutil.GetEncodeParam(TaskName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateBenchmarkTask"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/benchmark-tasks/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(TaskName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateBenchmarkTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResource(ClusterId *string, ResourceId *string, request *UpdateResourceRequest) (_result *UpdateResourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceResponse{}
	_body, _err := client.UpdateResourceWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceWithOptions(ClusterId *string, ResourceId *string, request *UpdateResourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceName)) {
		body["ResourceName"] = request.ResourceName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResource"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateResourceDLink(ClusterId *string, ResourceId *string, request *UpdateResourceDLinkRequest) (_result *UpdateResourceDLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateResourceDLinkResponse{}
	_body, _err := client.UpdateResourceDLinkWithOptions(ClusterId, ResourceId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateResourceDLinkWithOptions(ClusterId *string, ResourceId *string, request *UpdateResourceDLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateResourceDLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ResourceId = openapiutil.GetEncodeParam(ResourceId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DestinationCIDRs)) {
		body["DestinationCIDRs"] = request.DestinationCIDRs
	}

	if !tea.BoolValue(util.IsUnset(request.SecurityGroupId)) {
		body["SecurityGroupId"] = request.SecurityGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchId)) {
		body["VSwitchId"] = request.VSwitchId
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchIdList)) {
		body["VSwitchIdList"] = request.VSwitchIdList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateResourceDLink"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/resources/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ResourceId) + "/dlink"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateResourceDLinkResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateService(ClusterId *string, ServiceName *string, request *UpdateServiceRequest) (_result *UpdateServiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceResponse{}
	_body, _err := client.UpdateServiceWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateService"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceAutoScaler(ClusterId *string, ServiceName *string, request *UpdateServiceAutoScalerRequest) (_result *UpdateServiceAutoScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceAutoScalerResponse{}
	_body, _err := client.UpdateServiceAutoScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceAutoScalerWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceAutoScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceAutoScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Max)) {
		body["Max"] = request.Max
	}

	if !tea.BoolValue(util.IsUnset(request.Min)) {
		body["Min"] = request.Min
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Strategies))) {
		body["Strategies"] = request.Strategies
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceAutoScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/autoscaler"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceAutoScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceCronScaler(ClusterId *string, ServiceName *string, request *UpdateServiceCronScalerRequest) (_result *UpdateServiceCronScalerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceCronScalerResponse{}
	_body, _err := client.UpdateServiceCronScalerWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceCronScalerWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceCronScalerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceCronScalerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExcludeDates)) {
		body["ExcludeDates"] = request.ExcludeDates
	}

	if !tea.BoolValue(util.IsUnset(request.ScaleJobs)) {
		body["ScaleJobs"] = request.ScaleJobs
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceCronScaler"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/cronscaler"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceCronScalerResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceMirror(ClusterId *string, ServiceName *string, request *UpdateServiceMirrorRequest) (_result *UpdateServiceMirrorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceMirrorResponse{}
	_body, _err := client.UpdateServiceMirrorWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceMirrorWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceMirrorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceMirrorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Ratio)) {
		body["Ratio"] = request.Ratio
	}

	if !tea.BoolValue(util.IsUnset(request.Target)) {
		body["Target"] = request.Target
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceMirror"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/mirror"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceMirrorResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateServiceVersion(ClusterId *string, ServiceName *string, request *UpdateServiceVersionRequest) (_result *UpdateServiceVersionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.UpdateServiceVersionWithOptions(ClusterId, ServiceName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateServiceVersionWithOptions(ClusterId *string, ServiceName *string, request *UpdateServiceVersionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateServiceVersionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	ServiceName = openapiutil.GetEncodeParam(ServiceName)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Version)) {
		body["Version"] = request.Version
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateServiceVersion"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/services/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(ServiceName) + "/version"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateServiceVersionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateStress(ClusterId *string, StressName *string, request *UpdateStressRequest) (_result *UpdateStressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateStressResponse{}
	_body, _err := client.UpdateStressWithOptions(ClusterId, StressName, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateStressWithOptions(ClusterId *string, StressName *string, request *UpdateStressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateStressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	ClusterId = openapiutil.GetEncodeParam(ClusterId)
	StressName = openapiutil.GetEncodeParam(StressName)
	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStress"),
		Version:     tea.String("2021-07-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v2/stress/" + tea.StringValue(ClusterId) + "/" + tea.StringValue(StressName)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStressResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
