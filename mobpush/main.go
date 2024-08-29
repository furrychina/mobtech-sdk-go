package mobpush

//const (
//	RedisKeyPrefix = "user:device:rid:"
//)

// 1、前端会传递RegistrationID给后端，后端将RegistrationID保存到Redis中，以便后续推送消息时使用

// 保存到Redis
//err := rds.Main.ZAdd(c, RedisKeyPrefix+uid, &redis.Z{Member: request.RegistrationId, Score: float64(time.Now().Unix())}).Err()
//if err != nil {
//	utils.ResponseError(c, http.StatusInternalServerError, "推送服务注册失败", err.Error())
//	return
//}

// 2、当后端需要推送消息时，会从Redis中获取RegistrationID，然后调用MobPush的推送接口进行消息推送

// GetUserRID 获取用户下的RID
//func GetUserRID(userId string) []string {
//	return redis.Main.ZRange(redis.Context, RedisKeyPrefix+userId, 0, -1).Val()
//}
//
//// GetMUserRids 获取多个用户下的RID
//func GetMUserRids(userIds []string) []string {
//	rids := make([]string, 0)
//	for _, userId := range userIds {
//		rids = append(rids, GetUserRID(userId)...)
//	}
//	return rids
//}

// 3、推送示例
//	// 设置推送目标
//	pushTarget := PushTarget{
//		Target: TargetRid,
//		Rids:   GetUserRID(o.Seller.Uid),
//	}
//	// 推送消息
//	pushNotify := NewNotify("有新订单来啦", "XXXX拍下了你的商品", nil)
//	pushObj := NewMessage(&pushTarget, pushNotify)
//	_ = rocketMQ.ProduceMessage(rocketMQ.MQ_TOPIC_APP_PUSH, pushObj) // 推送到RocketMQ
