package mobpush

type ClientConfig struct {
	KeyPrefix string
}

type PushObject struct {
	Source           string            `json:"source"`
	AppKey           string            `json:"appkey"`
	PushTarget       *PushTarget       `json:"pushTarget"`
	PushNotify       *PushNotify       `json:"pushNotify"`
	PushFactoryExtra *PushFactoryExtra `json:"pushFactoryExtra"`
}
type PushTarget struct {
	//推送目标类型
	//- 1：广播
	//- 2：别名（alias）
	//- 3：标签（tags）
	//- 4：rid
	//- 5：地理位置（city）
	//- 6：用户分群
	//- 9：复杂地理位置推送（pushAreas）
	//- 14：fileid推送
	Target int `json:"target"`
	//推送别名集合，集合元素限制1000个以内。
	//例：["alias1","alias2"]
	//target等于2时必填
	Alias []string `json:"alias"`
	//推送标签集合，集合元素限制1000个以内。
	//例：["tag1","tag2"]
	//target等于3时必填
	Tags []string `json:"tags"`
	//标签组合方式，target等于3时可用
	//- 1：并集
	//- 2：交集
	//- 3：补集(暂未实现)
	//不填默认值为
	//- 1：并集
	TagsType int `json:"tagsType"`
	//推送rid集合，集合元素限制1000个以内。
	//例：["id1","id2"]
	//target等于4时必填
	Rids []string `json:"rids"`
}
type PushNotify struct {
	//推送生效渠道
	//- 1：android
	//- 2：iOS
	//- 8：harmony
	//例：[1, 2, 8]
	Plats []int `json:"plats"`
	//推送内容
	//注1：内容长度超过厂商限制会被截断。
	//注2：vivo不支持纯表情。
	Content string `json:"content"`
	//通知标题
	//注1：默认通知标题为应用名称
	//注2：标题长度超过厂商限制会被截断
	//注3：vivo不允许纯表情
	Title string `json:"title"`
	//推送类型
	//- 1：通知
	//- 2：自定义
	Type int `json:"type"`
	//是否开启个性化参数，默认为false
	//注：开启后，pushNotify.content参数支持${{}}个性化通知占位符
	CustomParamsSwitch bool `json:"customParamsSwitch"`
	//个性化参数默认内容，${{}}个性化通知占位符无法替换时会使用。
	//注：开启个性化参数时必填
	CustomParamDefaultContent string `json:"customParamDefaultContent"`
	//Android通知消息对象
	AndroidNotify AndroidNotify `json:"androidNotify"`
	//iOS通知消息对象
	IosNotify IosNotify `json:"iosNotify"`
	//鸿蒙通知消息对象
	HarmonyNotify HarmonyNotify `json:"harmonyNotify"`
	//是否是定时消息
	//- 0：否（默认）
	//- 1：是
	TaskCron int `json:"taskCron"`
	//taskCron=1时必填
	//定时消息发送时间，单位：毫秒时间戳
	//例：1594277916000
	TaskTime int64 `json:"taskTime"`
	//每秒推送速率的趋势，默认为0（代表不开启）
	Speed int `json:"speed"`
	//推送策略
	//- 1：先走tcp，再走厂商
	//- 2：先走厂商，再走tcp
	//- 3：只走厂商
	//- 4：只走tcp
	//- 5：设备亮屏推送 注：厂商透传只支持策略3或4
	Policy int `json:"policy"`
	//附加参数列表
	//例：{"key1":"value1","key2":"value2",…}
	//注：FCM厂商通道要求key不可以包含"from”、“google.”字符
	ExtrasMapList []ExtrasMap `json:"extrasMapList"`
}
type AndroidNotify struct {
	//角标类型
	//- 1：角标数值取androidBadge值
	//- 2：角标数值为androidBadge当前值加1
	//注：透传消息不支持
	AndroidBadgeType int `json:"androidBadgeType"`
	//角标数值
	AndroidBadge int `json:"androidBadge"`
	//TCP消息类别，当前仅华为机型支持可选枚举值：
	//promo 营销推广
	//recommendation 内容推荐
	//social 社交动态
	//call 通话
	//email 邮件
	//msg 即时聊天
	//navigation 导航
	//reminder 事项提醒
	//service 财务
	//alarm 闹钟/计时器
	//stopwatch 秒表
	//progress 进度
	//location_sharing 位置共享
	//注：参数为空时，默认赋值为：promo
	NativeCategory string `json:"nativeCategory"`
}
type IosNotify struct {
	Badge int `json:"badge"` // 角标
	// 角标类型
	//- 1:绝对值 不能为负数
	//- 2增减(正数增加，负数减少)，减到0以下会设置为0
	BadgeType int `json:"badgeType"`
	//APNs的category字段，只有IOS8及以上系统才支持此参数推送
	Category string `json:"category"`
	//slientpush =1 ,则为静默模式，不会携带任何badge，sound 和消息内容等参数给到用户，可以在不打扰用户的情况下进行内容更新等操作；
	//若slientpush和contentAvailable 都为1，则优先静默模式；
	SlientPush int `json:"slientPush"`
	//否contentAvailable =1，则为Background Remote Notification，表示应用程序在收到通知时会在后台唤醒，并可以执行一些操作；若为0，则是普通的 Remote Notification。
	//若slientpush和contentAvailable 都为1，则优先静默模式；
	ContentAvailable int `json:"contentAvailable"`
}
type ExtrasMap struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
type Response struct {
	Status int    `json:"status"` // 200 为成功
	Res    Result `json:"res"`
	Error  string `json:"error"`
}
type Result struct {
	BatchId string `json:"batchId"`
}

// HarmonyNotify 鸿蒙通知消息对象
// 渠道类型
// - 0：未知类型
// - 1：社交类型
// - 2：服务类型
// - 3：内容类型
// - 4：实况类型
// - 5：客户服务类型
// - 65535：其他类型
type HarmonyNotify struct {
	SlotType int `json:"slotType"` // 通知栏消息样式
}

// PushFactoryExtra 厂商通道扩展参数
type PushFactoryExtra struct {
	HuaweiExtra *HuaweiExtra `json:"huaweiExtra"`
}
type HuaweiExtra struct {
	//消息类型
	//- LOW：资讯营销类
	//- NORMAL：服务与通讯类
	//注：资讯营销类的消息提醒方式为静默通知，仅在下拉通知栏展示。 服务与通讯类的消息提醒方式为锁屏+铃声+震动
	Importance string `json:"importance"`
	//作用一：完成自分类权益申请后，用于标识消息类型，确定消息提醒方式，对特定类型消息加快发送，取值如下：
	//IM：即时聊天
	//VOIP：音视频通话
	//SUBSCRIPTION：订阅
	//TRAVEL：出行
	//HEALTH：健康
	//WORK：工作事项提醒
	//ACCOUNT：帐号动态
	//EXPRESS：订单&物流
	//FINANCE：财务
	//DEVICE_REMINDER：设备提醒
	//SYSTEM_REMINDER：系统提示
	//MAIL：邮件
	//PLAY_VOICE：语音播报（仅透传消息支持）
	//MARKETING：内容推荐、新闻、财经动态、生活资讯、社交动态、调研、产品促销、功能推荐、运营活动（仅对内容进行标识，不会加快消息发送）
	//作用二：申请特殊权限后，用于标识高优先级透传场景，取值如下：
	//VOIP：音视频通话
	//PLAY_VOICE：语音播报
	Category string `json:"category"`
}
type XiaomiExtra struct {
	//小米渠道Id 适配定制化渠道
	ChannelId string `json:"channelId"`
}
type OppoExtra struct {
	//OPPO渠道Id 适配定制化渠道
	ChannelId string `json:"channelId"`
}
type VivoExtra struct {
	//	VIVO消息类型
	//- 0：运营类型消息
	//- 1：系统类型消息
	Classification int `json:"classification"`
	//二级分类，传值参见二级分类标准中category说明
	//1、填写category后，可以不填写classification、messageSort，但若填写classification、messageSort，请保证category与messageSort或classification是正确对应关系，否则返回错误码10097；
	//2、赋值请按照消息分类规则填写，且必须大写；若传入错误无效的值，否则返回错误码10096；
	Category string `json:"category"`
}
