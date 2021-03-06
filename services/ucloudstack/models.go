// Code is generated by ucloud-model, DO NOT EDIT IT.

package ucloudstack

/*
BindVSInfo - 证书绑定的vs信息
*/
type BindVSInfo struct {

	// LB ID
	LBID string

	// LB名称
	LBName string

	// VS的端口
	Port int

	// VS的协议
	Protocol string

	// VS ID
	VSID string
}

/*
CertificateInfo - 证书信息
*/
type CertificateInfo struct {

	// 证书内容
	CertificateContent string

	// 证书ID
	CertificateID string

	// 证书类型，枚举值["ServerCrt","CACrt"]
	CertificateType string

	// 主域名
	CommonName string

	// 创建时间（平台创建时间）
	CreateTime int

	// 证书内容的过期时间
	ExpireTime int

	// 证书指纹
	Fingerprint string

	// 证书名
	Name string

	// 私钥内容
	Privatekey string

	// 地域
	Region string

	// 证书描述
	Remark string

	// 备域名
	SubjectAlternativeNames []string

	// 关联的vs的信息
	VSInfos []BindVSInfo

	// 可用区
	Zone string
}

/*
DiskInfo - 磁盘信息
*/
type DiskInfo struct {

	// 绑定资源ID
	AttachResourceID string

	// 硬盘计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 创建时间。时间戳
	CreateTime int

	// 硬盘ID
	DiskID string

	// 硬盘状态。Creating：创建中,BeingCloned：正在被克隆中,Unbound：已解绑,Unbounding：解绑中,Bounding：绑定中,Bound：已绑定,Upgrading：升级中,Deleting：删除中,Deleted：已删除,Releasing：销毁中,Released：已销毁；Snapshoting（快照中）；Rollbacking（回滚中）
	DiskStatus string

	// 硬盘用途类型，Boot（系统盘）、Data（数据盘）
	DiskType string

	// 过期时间。时间戳
	ExpireTime int

	// 名称
	Name string

	// 地域
	Region string

	// 备注
	Remark string

	// 磁盘类型。例如：Normal,SSD
	SetType string

	// 大小。单位GB
	Size int

	// 可用区
	Zone string
}

/*
EIPInfo - 外网IP信息
*/
type EIPInfo struct {

	// 带宽大小
	Bandwidth int

	// 绑定资源ID
	BindResourceID string

	// 绑定资源类型
	BindResourceType string

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 创建时间。时间戳
	CreateTime int

	// ID
	EIPID string

	// 过期时间。时间戳
	ExpireTime int

	// 外网IP
	IP string

	// IP版本,支持值：IPv4\IPv6
	IPVersion string

	// 名称
	Name string

	// 线路
	OperatorName string

	// 地域
	Region string

	// 备注
	Remark string

	// 状态。Allocating：申请中,Free：未绑定,Bounding：绑定中,Bound：已绑定,Unbounding：解绑中,Deleted：已删除,Releasing：销毁中,Released：已销毁,BandwidthChanging：带宽修改中
	Status string

	// 可用区
	Zone string
}

/*
ImageInfo - 镜像信息
*/
type ImageInfo struct {

	// 创建时间。时间戳。
	CreateTime int

	// 镜像ID
	ImageID string

	// 镜像状态。枚举类型：Making（创建中）,Available（可用）,Unavailable（不可用）,Terminating（销毁中）,Used（被使用中）,Deleting（删除中）,Deleted（已删除）, Uploading（导入中）
	ImageStatus string

	// 镜像类型。枚举类型：Base(基础镜像),Custom（自制镜像）。
	ImageType string

	// 镜像名称
	Name string

	// 镜像系统发行版本。例如：Centos, Ubuntu, Windows等
	OSDistribution string

	// 系统名称。例如：CentOS 7.4 x86_64
	OSName string

	// 系统类型。例如：Linux, Windows，Kylin
	OSType string

	// 地域
	Region string

	// 架构名称。例如：x86_64
	SetArch string

	// 可用区
	Zone string
}

/*
LBInfo - 负载均衡信息
*/
type LBInfo struct {

	// 告警模板ID
	AlarmTemplateID string

	// 虚拟机计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 创建时间，时间戳
	CreateTime int

	// 过期时间，时间戳
	ExpireTime int

	// 负载均衡ID
	LBID string

	// 状态。Creating:创建中,Running:运行中,Deleting:删除中,Deleted:已删除
	LBStatus string

	// 负载均衡类型，枚举值，WAN:外网负载均衡，LAN:内网负载均衡。
	LBType string

	// 名称
	Name string

	// 负载均衡的内网 IP 地址，当LB为外网类型时，该值为空。
	PrivateIP string

	// 负载均衡的外网 IP 地址，当LB为内网类型时，该值为空。
	PublicIP string

	// 地域
	Region string

	// 描述
	Remark string

	// 安全组 ID ，当LB为内网类型时，该值为空。
	SGID string

	// 子网ID
	SubnetID string

	// VPCID
	VPCID string

	// VServer的数量
	VSCount int

	// 可用区
	Zone string
}

/*
MetricSet - 监控值
*/
type MetricSet struct {

	// 监控时间
	Timestamp int

	// 监控值
	Value float64
}

/*
MetricInfo - 监控信息
*/
type MetricInfo struct {

	// 监控值信息
	Infos []MetricSet

	// 监控指标。虚拟机的监控指标枚举值为：BlockProcessCount，表示阻塞进程数；CPUUtilization，表示CPU使用率；DiskReadOps，表示磁盘读次数；DiskWriteOps，表示磁盘写次数；IORead，表示磁盘读吞吐；IOWrite，表示磁盘写吞吐；LoadAvg，表示平均负载1分钟；MemUsage，表示内存使用率；NetPacketIn，表示网卡入包量；NetPacketOut，表示网卡出包量；NICIn，表示网卡入带宽；NICOut，表示网卡出带宽；SpaceUsage，表示空间使用率；TCPConnectCount，表示TCP连接数；
	MetricName string
}

/*
NATGWInfo - NAT网关信息
*/
type NATGWInfo struct {

	// 告警模板ID
	AlarmTemplateID string

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 创建时间，时间戳
	CreateTime int

	// 虚拟IP
	EIP string

	// 过期时间，时间戳
	ExpireTime int

	// NAT网关ID
	NATGWID string

	// 状态。Creating:创建中, Running:运行中, Deleting:删除中, Deleted:已删除
	NATGWStatus string

	// 名称
	Name string

	// 地域
	Region string

	// 备注
	Remark string

	// NAT网关绑定的安全组ID
	SGID string

	// NAT网关实例所在的子网 ID
	SubnetID string

	// NAT网关实例所在的 VPC ID
	VPCID string

	// 可用区
	Zone string
}

/*
NATGWRuleInfo - NAT网关关联的白名单资源信息
*/
type NATGWRuleInfo struct {

	// 绑定的资源ID
	BindResourceID string

	// 绑定资源的类型
	BindResourceType string

	// 创建时间，时间戳。
	CreateTime int

	// 白名单资源的内网IP地址
	IP string

	// NAT网关ID
	NATGWID string

	// nat网关类型
	NATGWType string

	// 添加的白名单资源名称
	Name string

	// 白名单ID
	RuleID string

	// 状态。Bounding:绑定中,Bound:已绑定,Unbounding:解绑中,Unbound：已解绑
	RuleStatus string
}

/*
NICInfo - 网卡信息
*/
type NICInfo struct {

	// 绑定资源ID
	BindResourceID string

	// 创建时间。时间戳
	CreateTime int

	// IP
	IP string

	// mac 地址
	MAC string

	// 网卡ID
	NICID string

	// 网卡状态。枚举值。Creating：创建中,Free：未绑定,Unbounding：解绑中,Bounding：绑定中,Bound：已绑定,BindSGing：绑定安全组中,UnbindSGing：解绑安全组中,UpdateSGing：更新安全组中,Deleting：删除中,Deleted：已删除,Releasing：销毁中,Released：已销毁
	NICStatus string

	// 名称
	Name string

	// 地域
	Region string

	// 备注
	Remark string

	// 安全组ID
	SGID string

	// Subnet ID
	SubnetID string

	// VPC ID
	VPCID string

	// 可用区
	Zone string
}

/*
OPLogInfo - 操作日志
*/
type OPLogInfo struct {

	// 创建时间
	CreateTime int

	// 是否操作成功， Yes, No
	IsSuccess string

	// 日志ID
	OPLogsID string

	// API
	OPName string

	// 操作时间
	OPTime int

	// 错误信息
	OpMessage string

	//
	Region string

	// 资源ID
	ResourceID string

	// 资源类型
	ResourceType int

	// 状态码
	RetCode int

	// 账号邮箱
	UserEmail string

	//
	Zone string
}

/*
PhysicalIPInfo - 物理IP信息
*/
type PhysicalIPInfo struct {

	// 绑定资源ID
	BindResourceID string

	// 绑定资源类型
	BindResourceType string

	// 创建时间。时间戳
	CreateTime int

	// 物理IP
	IP string

	// 名称
	Name string

	// 线路
	OperatorName string

	// 物理IP的ID
	PhysicalIPID string

	// 地域
	Region string

	// 备注
	Remark string

	// 状态。Allocating：申请中,Free：未绑定,Bounding：绑定中,Bound：已绑定,Unbounding：解绑中,Deleted：已删除,Releasing：销毁中,Released：已销毁
	Status string

	// 过期时间。时间戳
	UpdateTime int

	// 可用区
	Zone string
}

/*
RSInfo - 转发规则关联的服务节点信息
*/
type RSInfo struct {

	// 绑定的资源ID
	BindResourceID string

	// 创建时间，时间戳
	CreateTime int

	// 服务节点的内网 IP 地址
	IP string

	// 服务节点所属的负载均衡 ID
	LBID string

	// 服务节点的资源名称
	Name string

	// 服务节点暴露的服务端口号
	Port int

	// 服务节点的 ID
	RSID string

	// 节点模式。枚举值，Enabling:开启中,Enable:已启用,Disabling:禁用中,Disable:已禁用
	RSMode string

	// RSStatus 的描述修改为：状态，枚举值，Creating:创建中,Inactive:无效,Active:有效,Updating:更新中,Deleting:删除中,Deleted:已删除。其中有效代表节点服务健康，无效代表节点服务异常。
	RSStatus string

	// 更新时间，时间戳
	UpdateTime int

	// 服务节点所属的 VServer ID
	VSID string

	// 服务节点的权重
	Weight int
}

/*
RecycledResourceInfo - 回收站资源信息
*/
type RecycledResourceInfo struct {

	// 创建时间
	CreateTime int

	// 删除时间
	DeleteTime int

	// 描述
	Description string

	// 过期时间
	ExpireTime int

	// 是否自动销户
	IsAutoTerminated bool

	// 名称
	Name string

	// 地域
	Region string

	// 资源ID
	ResourceID string

	// 资源类型：VM:虚拟机，Disk:硬盘，EIP:外网IP，PIP:物理IP，MySQL:数据库，Redis:缓存
	ResourceType string

	// 资源状态
	Status string

	// 销毁时间
	WillTerminateTime int

	// 可用区
	Zone string
}

/*
SGRuleInfo - 安全组规则信息
*/
type SGRuleInfo struct {

	// 端口号
	DstPort string

	// 方向。1：入，0：出
	IsIn string

	// 优先级。HIGH:高，MEDIUM:中，LOW:低
	Priority string

	// 协议
	ProtocolType string

	// 动作。ACCEPT：接受，DROP：拒绝
	RuleAction string

	// 规则ID
	RuleID string

	// IP或者掩码/段形式。10.0.0.2,10.0.10.10/16
	SrcIP string
}

/*
SGInfo - 安全组信息
*/
type SGInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 地域
	Region string

	// 描述
	Remark string

	// 资源绑定数量
	ResourceCount int

	// 安全组规则。
	Rule []SGRuleInfo

	// 规则数量
	RuleCount int

	// 安全组ID
	SGID string

	// 状态。Creating：创建中,Updating：更新中,Available：有效,Deleted：已删除,Terminating：销毁中,Terminated：已销毁
	Status string

	// 更新时间，时间戳
	UpdateTime int

	// 可用区
	Zone string
}

/*
StorageTypeInfo - 存储类型信息
*/
type StorageTypeInfo struct {

	// 地域
	Region string

	// 架构
	SetArch string

	// 存储类型
	StorageType string

	// 存储类型别名
	StorageTypeAlias string

	// 可用区
	Zone string
}

/*
SubnetInfo - 子网信息
*/
type SubnetInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 网段
	Network string

	// 地域
	Region string

	// 描述
	Remark string

	// 状态；Allocating：申请中,Available：有效,Deleting：删除中,Deleted：已删除
	State string

	// ID
	SubnetID string

	// 更新时间，时间戳
	UpdateTime int

	// 可用区
	Zone string
}

/*
UserInfo - 租户信息
*/
type UserInfo struct {

	// 账户余额
	Amount float64

	// 账户创建时间。时间戳
	CreateTime int

	// 租户名称
	Email string

	// 私钥
	PrivateKey string

	// 公钥
	PublicKey string

	// 用户状态。USER_STATUS_AVAILABLE：正常，USER_STATUS_FREEZE：冻结
	Status string

	// 更新时间。时间戳
	UpdateTime int

	// 租户ID.
	UserID int
}

/*
VMIPInfo - UCloudStack虚拟机IP信息
*/
type VMIPInfo struct {

	// IP 值
	IP string

	// IP版本,支持值：IPv4\IPv6
	IPVersion string

	// 网卡 ID
	InterfaceID string

	// 是否是弹性网卡。枚举值：Y，表示是；N，表示否；
	IsElastic string

	// MAC 地址值
	MAC string

	// 安全组 ID
	SGID string

	// 安全组名称
	SGName string

	// 子网 ID
	SubnetID string

	// 子网名称
	SubnetName string

	// IP 类型。枚举值：Private，表示内网；Public，表示外网；Physical，表示物理网；
	Type string

	// VPC ID
	VPCID string

	// VPC 名称
	VPCName string
}

/*
VMDiskInfo - UCloudStack虚拟机磁盘信息
*/
type VMDiskInfo struct {

	// 磁盘 ID
	DiskID string

	// 磁盘盘符
	Drive string

	// 是否是弹性磁盘。枚举值为：Y，表示是；N，表示否；
	IsElastic string

	// 磁盘名称
	Name string

	// 磁盘大小，单位 GB
	Size int

	// 磁盘类型。枚举值：Boot，表示系统盘；Data，表示数据盘；
	Type string
}

/*
VMInstanceInfo - UCloudStack虚拟机信息
*/
type VMInstanceInfo struct {

	// CPU 个数
	CPU int

	// 虚拟机计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 虚拟机创建时间
	CreateTime int

	// 磁盘信息
	DiskInfos []VMDiskInfo

	// 虚拟机过期时间
	ExpireTime int

	// IP 信息
	IPInfos []VMIPInfo

	// 镜像 ID
	ImageID string

	// 内存大小，单位 M
	Memory int

	// 虚拟机名称
	Name string

	// 操作系统名称
	OSName string

	// 操作系统类型
	OSType string

	// Region
	Region string

	// Region 别名
	RegionAlias string

	// 备注
	Remark string

	// 虚拟机状态。枚举值：Initializing，表示初始化；Starting，表示启动中；Restarting，表示重启中；Running，表示运行；Stopping，表示关机中；Stopped，表示关机；Deleted，表示已删除；Resizing，表示修改配置中；Terminating，表示销毁中；Terminated，表示已销毁；Migrating，表示迁移中；WaitReinstall，表示等待重装系统；Reinstalling，表示重装中；Poweroffing，表示断电中；ChangeSGing，表示修改防火墙中；
	State string

	// 子网 ID
	SubnetID string

	// 子网 名称
	SubnetName string

	// 虚拟机 ID
	VMID string

	// 虚拟机类型
	VMType string

	// 虚拟机类型别名
	VMTypeAlias string

	// VPC ID
	VPCID string

	// VPC 名称
	VPCName string

	// Zone
	Zone string

	// Zone 别名
	ZoneAlias string
}

/*
VMTypeInfo - 主机机型信息
*/
type VMTypeInfo struct {

	// 地域
	Region string

	// 架构
	SetArch string

	// 机型
	VMType string

	// 机型别名
	VMTypeAlias string

	// 可用区
	Zone string
}

/*
VPCInfo - VPC信息
*/
type VPCInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 名称
	Name string

	// 网段，比如10.0.0.0/16
	Network string

	// 地域。
	Region string

	// 描述
	Remark string

	// 状态；Allocating：申请中,Available：有效,Terminating：销毁中,Terminated：已销毁
	State string

	// 该VPC下拥有的子网数目
	SubnetCount int

	// 该VPC下子网信息。
	SubnetInfos []SubnetInfo

	// 修改时间，时间戳
	UpdateTime int

	// VPC的ID
	VPCID string

	// 可用区
	Zone string
}

/*
VSPolicyInfo - 内容转发规则信息
*/
type VSPolicyInfo struct {

	// 创建时间，时间戳
	CreateTime int

	// 内容转发规则关联的请求域名，值可为空，即代表仅匹配路径。
	Domain string

	// 负载均衡ID
	LBID string

	// 内容转发规则关联的请求访问路径，如 "/" 。
	Path string

	// 内容转发规则ID
	PolicyID string

	// 状态，枚举值，Available:有效,Deleted:已删除
	PolicyStatus string

	// 转发规则关联的服务节点信息
	RSInfos []RSInfo

	// 更新时间，时间戳
	UpdateTime int

	// VServerID
	VSID string
}

/*
VSInfo - RServer信息
*/
type VSInfo struct {

	// 告警模板ID
	AlarmTemplateID string

	// 创建时间，时间戳
	CreateTime int

	// HTTP 健康检查时校验请求的 HOST 字段中的域名。当健康检查类型为端口检查时，该值为空。
	Domain string

	// 负载均衡的健康检查类型。枚举值：Port:端口检查；Path: HTTP检查 。
	HealthcheckType string

	// 负载均衡的连接空闲超时时间，单位为秒，默认值为 60s 。当 VServer 协议为 UDP 时，该值为空。
	KeepaliveTimeout int

	// VServer 所属的负载均衡 ID
	LBID string

	// HTTP 健康检查的路径。当健康检查类型为端口检查时，该值为空。
	Path string

	// 会话保持KEY，仅当 VServer 协议为 HTTP 且会话保持为手动时有效。
	PersistenceKey string

	// 会话保持类型。枚举值：None:关闭；Auto:自动生成；Manual:手动生成 。当协议为 TCP 时，该值为空；当协议为 UDP 时 Auto 表示开启会话保持 。
	PersistenceType string

	// 端口
	Port int

	// 协议
	Protocol string

	// 健康检查状态，枚举值，Empty:全部异常,Parts:部分异常,All:正常
	RSHealthStatus string

	// 前 VServer 中已添加的服务节点信息。
	RSInfos []RSInfo

	// 更新时间，时间戳
	UpdateTime int

	// VServer的ID
	VSID string

	// 转发规则 Policy 的规则信息
	VSPolicyInfos []VSPolicyInfo

	// VServer 的资源状态。枚举值，Available:可用,Updating:更新中,Deleted:已删除 。
	VSStatus string
}

/*
PriceInfo - 价格信息
*/
type PriceInfo struct {

	// 计费模式。枚举值：Dynamic，表示小时；Month，表示月；Year，表示年；
	ChargeType string

	// 价格
	Price float64
}
