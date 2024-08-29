package mobpush

type ClientConfig struct {
	KeyPrefix string
}

type PushObject struct {
	Source     string      `json:"source"`
	AppKey     string      `json:"appkey"`
	PushTarget *PushTarget `json:"pushTarget"`
	PushNotify *PushNotify `json:"pushNotify"`
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
