package bean

import (
	"github.com/GodSlave/MyGoServer/module"
	"github.com/go-xorm/xorm"
)

// model

type Buddha struct {
   Id     int32 "xorm:pk"//佛id
   Name     string//佛像名
   InitialTemple     bool//是否是初始道场
   dutes     []int32  "xorm:extends"//职责
   FoModel     string//模型
   MapName     string//道场名
   MapRes     string//地图
   Thumbnail     string//缩略图
   Desc     string//道场介绍
   ItemId     int32//愿力值类型

}
type BuddhaLevel struct {
   Id     int32//序号
   BuddhaId     int32 "xorm:pk"//佛id
   Level     int32 "xorm:pk"//等级
   Cost     []ItemBase  "xorm:extends"//需求
   Produce     int32//每单位时间产出的愿力值
   CandlePower     int32//点灯获得基础愿力
   FlowerPower     int32//献花获得基础愿力
   IncensePower     int32//上香获得基础愿力
   DropId     int32//掉落id
   Thumbnail     string//缩略图
   Attendents     []int32  "xorm:extends"//随从

}
type Birth struct {
   Id     int32 "xorm:pk"//序号
   AwardItem     []ItemBase  "xorm:extends"//初始道具

}
type Quality struct {
   Id     int32 "xorm:pk"//序号
   QualityText     string//品质类型

}
type WishName struct {
   Id     int32 "xorm:pk"//许愿牌id
   WishName     []string  "xorm:extends"//愿望名称

}
type Dropitem struct {
   Id     int32 "xorm:pk"//掉落Id
   DropItems     []ItemDrop  "xorm:extends"//道具id

}
type DropItemInfo struct {

}
type Poem struct {
   Id     int32 "xorm:pk"//法师偈语id
   Poem     []string  "xorm:extends"//法师偈语内容
   PoemQuality     int32//法师偈语开启条件
   Explain     string//法师偈语解释
   ExplainQuality     int32//解释开启条件

}
type WorShip struct {
   Id     int32 "xorm:pk"//道具Id
   worshipType     EWorshipType//类型
   Award     int32//每单位时间获得奖励
   Last     int32//持续时间

}
type ItemBase struct {
   Id     int32//物品Id
   Num     int32//物品数量

}
type ItemDrop struct {
   Id     int32//物品Id
   Num     int32//物品数量
   Chance     int32//掉落概率

}
type Duty struct {
   Id     int32 "xorm:pk"//id
   Name     string//名称
   Desc     string//描述
   cardtype     EWishCardType//许愿牌类型

}
type WishCard struct {
   Id     int32 "xorm:pk"//许愿牌id
   CardName     string//许愿牌名称
   Icon     string//许愿牌icon
   Quality     int32//许愿牌品质
   WishName     []string  "xorm:extends"//愿望名称
   poemID     int32//法师偈语
   PracticeBase     int32//可兑换基础修行值
   PracticePerBlessed     int32//每次被祝福增加修行值
   cardtype     EWishCardType//许愿牌类型
   ItemId     int32//放飞所需愿力值ID
   PowerNeeded     int32//放飞所需愿力值数量

}
type Item struct {
   Id     int32 "xorm:pk"//道具id
   Name     string//道具名
   Icon     string//icon
   Deposit     int32//可堆叠数

}
type SacredLand struct {
   Id     int32 "xorm:pk"//编号
   TermId     int32//几号道场做为条件
   TermLevel     int32//需求等级

}
type Attendant struct {
   Id     int32 "xorm:pk"//随从id
   Name     string//随从名
   BuddhaId     int32//佛id
   BuddhaLevel     int32//佛等级
   Model     string//模型
   Seat     int32//位置

}


const WorShip_Inter = 60

const Born_Gift = `{"Role":{"roleType":1,"roleName":"TestName"},"StoreItem":[{"itemId":10001,"itemNumber":3},{"itemId":10101,"itemNumber":3},{"itemId":10201,"itemNumber":3},{"itemId":1,"itemNumber":100}]}`

const Born_Fo = `{"Buddha":{[1，2，3]}`

const MaxWishSize = 6

const MaxBlessSize = 20

const WishCostPower = 100



//EProtocolModule
type EProtocolModule int32
const EProtocolModule_User = 2 // 用户模块
const EProtocolModule_Role = 11 // 角色
const EProtocolModule_Buddha = 12 // 佛像
const EProtocolModule_Inventory = 13 // 背包
const EProtocolModule_Base = 14 // 基础数据

//EUserProtocol
type EUserProtocol int32
const EUserProtocol_Login = 1 // 用户登陆
const EUserProtocol_Register = 2 // 注册
const EUserProtocol_GetVerifyCode = 3 // 获取验证码
const EUserProtocol_GetSelfInfo = 4 // 获取自己的信息
const EUserProtocol_Logout = 5 // 登出

//ERoleProtocol
type ERoleProtocol int32
const ERoleProtocol_CreateRole = 1 // 创建角色
const ERoleProtocol_GetRoleInfo = 2 // 获取角色信息

//EBuddhaProtocol
type EBuddhaProtocol int32
const EBuddhaProtocol_GetBuddhaInfo = 1 // 获取信息
const EBuddhaProtocol_ActiveBuddha = 2 // 启用佛像
const EBuddhaProtocol_UpGradeBuddha = 3 // 升级佛像
const EBuddhaProtocol_WorshipBuddha = 4 // 拜佛
const EBuddhaProtocol_CollectDesire = 5 // 收集愿力值
const EBuddhaProtocol_SwitchBuddha = 6 // 切换佛

//EInventoryProtocol
type EInventoryProtocol int32
const EInventoryProtocol_GetAllItem = 1 // 获取背包物品信息
const EInventoryProtocol_GetAllMoney = 2 // 获取用户金钱信息
const EInventoryProtocol_GetAllMateria = 3 // 获取除了金钱的物品信息

//EBaseProtocol
type EBaseProtocol int32
const EBaseProtocol_GetBaseInfo = 1 // 获取基础信息

//EDuty
type EDuty int32
const EDuty_health = 1 // 健康
const EDuty_marriage = 2 // 姻缘
const EDuty_cause = 3 // 事业

//EWishCardType
type EWishCardType int32
const EWishCardType_health = 1 // 健康牌
const EWishCardType_marriage = 2 // 姻缘牌
const EWishCardType_career = 3 // 事业牌

//EWorshipType
type EWorshipType int32
const EWorshipType_incense = 1 // 上香
const EWorshipType_candle = 2 // 点灯
const EWorshipType_flower = 3 // 献花


func EnableDBCache(app module.App) {
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)

	app.GetSqlEngine().MapCacher(&Buddha{}, cacher)
	app.GetSqlEngine().MapCacher(&BuddhaLevel{}, cacher)
	app.GetSqlEngine().MapCacher(&Birth{}, cacher)
	app.GetSqlEngine().MapCacher(&Quality{}, cacher)
	app.GetSqlEngine().MapCacher(&WishName{}, cacher)
	app.GetSqlEngine().MapCacher(&Dropitem{}, cacher)
	app.GetSqlEngine().MapCacher(&DropItemInfo{}, cacher)
	app.GetSqlEngine().MapCacher(&Poem{}, cacher)
	app.GetSqlEngine().MapCacher(&WorShip{}, cacher)
	app.GetSqlEngine().MapCacher(&ItemBase{}, cacher)
	app.GetSqlEngine().MapCacher(&ItemDrop{}, cacher)
	app.GetSqlEngine().MapCacher(&Duty{}, cacher)
	app.GetSqlEngine().MapCacher(&WishCard{}, cacher)
	app.GetSqlEngine().MapCacher(&Item{}, cacher)
	app.GetSqlEngine().MapCacher(&SacredLand{}, cacher)
	app.GetSqlEngine().MapCacher(&Attendant{}, cacher)
}

func ClearDBChache(app module.App) {

	app.GetSqlEngine().ClearCache(&Buddha{})
	app.GetSqlEngine().ClearCache(&BuddhaLevel{})
	app.GetSqlEngine().ClearCache(&Birth{})
	app.GetSqlEngine().ClearCache(&Quality{})
	app.GetSqlEngine().ClearCache(&WishName{})
	app.GetSqlEngine().ClearCache(&Dropitem{})
	app.GetSqlEngine().ClearCache(&DropItemInfo{})
	app.GetSqlEngine().ClearCache(&Poem{})
	app.GetSqlEngine().ClearCache(&WorShip{})
	app.GetSqlEngine().ClearCache(&ItemBase{})
	app.GetSqlEngine().ClearCache(&ItemDrop{})
	app.GetSqlEngine().ClearCache(&Duty{})
	app.GetSqlEngine().ClearCache(&WishCard{})
	app.GetSqlEngine().ClearCache(&Item{})
	app.GetSqlEngine().ClearCache(&SacredLand{})
	app.GetSqlEngine().ClearCache(&Attendant{})
}

func DisableDBCache(app module.App) {

	app.GetSqlEngine().MapCacher(&Buddha{}, nil)
	app.GetSqlEngine().MapCacher(&BuddhaLevel{}, nil)
	app.GetSqlEngine().MapCacher(&Birth{}, nil)
	app.GetSqlEngine().MapCacher(&Quality{}, nil)
	app.GetSqlEngine().MapCacher(&WishName{}, nil)
	app.GetSqlEngine().MapCacher(&Dropitem{}, nil)
	app.GetSqlEngine().MapCacher(&DropItemInfo{}, nil)
	app.GetSqlEngine().MapCacher(&Poem{}, nil)
	app.GetSqlEngine().MapCacher(&WorShip{}, nil)
	app.GetSqlEngine().MapCacher(&ItemBase{}, nil)
	app.GetSqlEngine().MapCacher(&ItemDrop{}, nil)
	app.GetSqlEngine().MapCacher(&Duty{}, nil)
	app.GetSqlEngine().MapCacher(&WishCard{}, nil)
	app.GetSqlEngine().MapCacher(&Item{}, nil)
	app.GetSqlEngine().MapCacher(&SacredLand{}, nil)
	app.GetSqlEngine().MapCacher(&Attendant{}, nil)
}

