package bean



// model

type Buddha struct {
   Id     int32//佛id
   Name     string//佛像名
   Duty     EDuty//职能
   FoModel     string//模型
   MapName     string//道场名
   MapRes     string//地图

}
type BuddhaLevel struct {
   Id     int32//序号
   BuddhaId     int32//佛id
   Level     int32//等级
   Cost     []ItemBase//需求
   Produce     int32//产出的愿力值
   DropId     int32//掉落id

}
type Birth struct {
   Id     int32//序号
   AwardItem     []ItemBase//初始道具

}
type Dropitem struct {
   Id     int32//掉落Id
   DropItems     []ItemDrop//道具id

}
type DropItemInfo struct {

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
type Item struct {
   Id     int32//道具id
   Name     string//道具名
   Icon     string//icon
   Deposit     int32//可堆叠数

}
type SacredLand struct {
   Id     int32//编号
   TermId     int32//几号道场做为条件
   TermLevel     int32//需求等级

}
type Attendant struct {
   Id     int32//随从id
   Name     string//随从名
   BuddhaId     int32//佛id
   BuddhaLevel     int32//佛等级
   Model     string//模型
   Seat     int32//位置

}


const WorShip_Inter = 60



//UserLoginResponseState
type UserLoginResponseState int32
const UserLoginResponseState_LoginSuccess = 0 // 登陆成功
const UserLoginResponseState_LoginFailure = 1 // 登陆失败

//ProtocolModule
type ProtocolModule int32
const ProtocolModule_User = 2 // 用户模块
const ProtocolModule_Role = 3 // 角色
const ProtocolModule_Buddha = 4 // 佛像
const ProtocolModule_Storage = 5 // 背包

//UserProtocol
type UserProtocol int32
const UserProtocol_Login = 1 // 用户登陆
const UserProtocol_Register = 2 // 注册
const UserProtocol_GetVerifyCode = 3 // 获取验证码
const UserProtocol_GetSelfInfo = 4 // 获取自己的信息

//RoleProtocol
type RoleProtocol int32
const RoleProtocol_CreateRole = 1 // 创建角色
const RoleProtocol_GetRoleInfo = 2 // 获取角色信息

//BuddhaProtocol
type BuddhaProtocol int32
const BuddhaProtocol_GetBuddhaInfo = 1 // 获取信息
const BuddhaProtocol_ActiveBuddha = 2 // 启用佛像
const BuddhaProtocol_UpGradeBuddha = 3 // 升级佛像
const BuddhaProtocol_WorshipBuddha = 4 // 拜佛
const BuddhaProtocol_CollectDesire = 5 // 收集愿力值
const BuddhaProtocol_SwitchBuddha = 6 // 切换佛

//StorageProtocol
type StorageProtocol int32
const StorageProtocol_GetAllStore = 1 // 获取背包物品信息

//EDuty
type EDuty int32
const EDuty_health = 1 // 健康
const EDuty_marriage = 2 // 姻缘
const EDuty_cause = 3 // 事业


var Buddhas  []Buddha
var BuddhaLevels  map[string]BuddhaLevel
var Births  map[string]Birth
var Dropitems  map[string]Dropitem
var DropItemInfos  []DropItemInfo
var ItemBases  []ItemBase
var ItemDrops  []ItemDrop
var Items  map[string]Item
var SacredLands  map[string]SacredLand
var Attendants  map[string]Attendant


func init()  {

   
   
   Buddhas = append(Buddhas , Buddha{  })
   Buddhas = append(Buddhas , Buddha{  })
   Buddhas = append(Buddhas , Buddha{  })
   
   
   
   BuddhaLevels["1"]= BuddhaLevel{ 1,1,1,[]ItemBase{{101,3,},{1,10,},},10,1, } 
   BuddhaLevels["2"]= BuddhaLevel{ 2,1,2,[]ItemBase{{101,3,},{102,3,},{2,20,},},20,1, } 
   BuddhaLevels["3"]= BuddhaLevel{ 3,1,3,[]ItemBase{{101,5,},{102,6,},{3,50,},},40,2, } 
   BuddhaLevels["4"]= BuddhaLevel{ 4,1,4,[]ItemBase{{101,8,},{102,7,},{103,3,},{4,99,},},80,2, } 
   BuddhaLevels["5"]= BuddhaLevel{ 5,1,5,[]ItemBase{{101,10,},{102,9,},{103,6,},},160,3, } 
   BuddhaLevels["6"]= BuddhaLevel{ 6,1,6,[]ItemBase{{101,15,},{102,10,},{103,8,},},320,4, } 
   BuddhaLevels["7"]= BuddhaLevel{ 7,1,7,[]ItemBase{{101,20,},{102,12,},{103,10,},},700,5, } 
   BuddhaLevels["8"]= BuddhaLevel{ 8,2,1,[]ItemBase{{101,3,},{1,10,},},10,1, } 
   BuddhaLevels["9"]= BuddhaLevel{ 9,2,2,[]ItemBase{{101,3,},{102,3,},{2,20,},},20,1, } 
   BuddhaLevels["10"]= BuddhaLevel{ 10,2,3,[]ItemBase{{101,5,},{102,6,},{3,50,},},40,2, } 
   BuddhaLevels["11"]= BuddhaLevel{ 11,2,4,[]ItemBase{{101,8,},{102,7,},{103,3,},{4,99,},},80,2, } 
   BuddhaLevels["12"]= BuddhaLevel{ 12,2,5,[]ItemBase{{101,10,},{102,9,},{103,6,},},160,3, } 
   BuddhaLevels["13"]= BuddhaLevel{ 13,2,6,[]ItemBase{{101,15,},{102,10,},{103,8,},},320,4, } 
   BuddhaLevels["14"]= BuddhaLevel{ 14,2,7,[]ItemBase{{101,20,},{102,12,},{103,10,},},700,5, } 
   
   
   
   Births["1"]= Birth{ 1,[]ItemBase{{101,99,},{102,99,},{103,99,},}, } 
   
   
   
   Dropitems["1"]= Dropitem{ 1,[]ItemDrop{{1001,1,7000,},{1002,1,3000,},}, } 
   Dropitems["2"]= Dropitem{ 2,[]ItemDrop{{1002,1,7000,},{1003,1,3000,},}, } 
   Dropitems["3"]= Dropitem{ 3,[]ItemDrop{{1002,1,5000,},{1003,1,4000,},{1004,1,1000,},}, } 
   Dropitems["4"]= Dropitem{ 4,[]ItemDrop{{1003,1,8000,},{1004,1,2000,},}, } 
   Dropitems["5"]= Dropitem{ 5,[]ItemDrop{{1003,1,5000,},{1004,1,4000,},{1005,1,1000,},}, } 
   
   
   
   DropItemInfos = append(DropItemInfos , DropItemInfo{  })
   
   
   
   
   
   
   
   
   
   Items["1"]= Item{ 1,"功德币","UI_gongdebi_icon",99999999, } 
   Items["2"]= Item{ 2,"愿力值","UI_yuanlizhi_icon",99999999, } 
   Items["3"]= Item{ 3,"修行值","UI_xiuxingzhi_icon",99999999, } 
   Items["4"]= Item{ 4,"福报值","UI_fubaozhi_icon",99999999, } 
   Items["101"]= Item{ 101,"玉净瓶","UI_yujingping_icon",99, } 
   Items["102"]= Item{ 102,"玲珑宝塔","UI_baota_icon",99, } 
   Items["103"]= Item{ 103,"莲花","UI_lianhua_icon",99, } 
   Items["1001"]= Item{ 1001,"健康福牌(白)","UI_fupai1_icon",99, } 
   Items["1002"]= Item{ 1002,"健康福牌(绿)","UI_fupai2_icon",99, } 
   Items["1003"]= Item{ 1003,"健康福牌(蓝)","UI_fupai3_icon",99, } 
   Items["1004"]= Item{ 1004,"健康福牌(紫)","UI_fupai4_icon",99, } 
   Items["1005"]= Item{ 1005,"健康福牌(橙)","UI_fupai5_icon",99, } 
   Items["1011"]= Item{ 1011,"姻缘福牌(白)","UI_fupai6_icon",99, } 
   Items["1012"]= Item{ 1012,"姻缘福牌(绿)","UI_fupai7_icon",99, } 
   Items["1013"]= Item{ 1013,"姻缘福牌(蓝)","UI_fupai8_icon",99, } 
   Items["1014"]= Item{ 1014,"姻缘福牌(紫)","UI_fupai9_icon",99, } 
   Items["1015"]= Item{ 1015,"姻缘福牌(橙)","UI_fupai10_icon",99, } 
   Items["1021"]= Item{ 1021,"事业福牌(白)","UI_fupai11_icon",99, } 
   Items["1022"]= Item{ 1022,"事业福牌(绿)","UI_fupai12_icon",99, } 
   Items["1023"]= Item{ 1023,"事业福牌(蓝)","UI_fupai13_icon",99, } 
   Items["1024"]= Item{ 1024,"事业福牌(紫)","UI_fupai14_icon",99, } 
   Items["1025"]= Item{ 1025,"事业福牌(橙)","UI_fupai15_icon",99, } 
   Items["10001"]= Item{ 10001,"低级细香","UI_lifo1_icon",99, } 
   Items["10002"]= Item{ 10002,"中级细香","UI_lifo2_icon",99, } 
   Items["10003"]= Item{ 10003,"高级细香","UI_lifo3_icon",99, } 
   Items["10004"]= Item{ 10004,"低级粗香","UI_lifo4_icon",99, } 
   Items["10005"]= Item{ 10005,"中级粗香","UI_lifo5_icon",99, } 
   Items["10006"]= Item{ 10006,"高级粗香","UI_lifo6_icon",99, } 
   Items["10101"]= Item{ 10101,"苹果","UI_lifo7_icon",99, } 
   Items["10102"]= Item{ 10102,"橙子","UI_lifo8_icon",99, } 
   Items["10103"]= Item{ 10103,"金桔","UI_lifo9_icon",99, } 
   Items["10104"]= Item{ 10104,"桃子","UI_lifo10_icon",99, } 
   Items["10105"]= Item{ 10105,"人参果","UI_lifo11_icon",99, } 
   Items["10106"]= Item{ 10106,"释迦果","UI_lifo12_icon",99, } 
   Items["10201"]= Item{ 10201,"黄蜡烛","UI_lifo13_icon",99, } 
   Items["10202"]= Item{ 10202,"红蜡烛","UI_lifo14_icon",99, } 
   Items["10203"]= Item{ 10203,"龙凤烛","UI_lifo15_icon",99, } 
   Items["10204"]= Item{ 10204,"酥油灯烛","UI_lifo16_icon",99, } 
   Items["10205"]= Item{ 10205,"如意莲花烛","UI_lifo17_icon",99, } 
   Items["10206"]= Item{ 10206,"佛光普照","UI_lifo18_icon",99, } 
   
   
   
   SacredLands["1"]= SacredLand{ 1,0,0, } 
   SacredLands["2"]= SacredLand{ 2,1,3, } 
   SacredLands["3"]= SacredLand{ 3,2,3, } 
   SacredLands["4"]= SacredLand{ 4,3,5, } 
   SacredLands["5"]= SacredLand{ 5,4,5, } 
   SacredLands["6"]= SacredLand{ 6,5,6, } 
   SacredLands["7"]= SacredLand{ 7,6,7, } 
   
   
   
   Attendants["1"]= Attendant{ 1,"日光菩萨",1,2,"M_riguang_att",1, } 
   Attendants["2"]= Attendant{ 2,"月光菩萨",1,3,"M_yueguang_att",2, } 
   Attendants["3"]= Attendant{ 3,"神将1",1,4,"M_shenjiang1_att",3, } 
   Attendants["4"]= Attendant{ 4,"神将2",1,4,"M_shenjiang2_att",4, } 
   Attendants["5"]= Attendant{ 5,"神将3",1,4,"M_shenjiang3_att",5, } 
   Attendants["6"]= Attendant{ 6,"神将4",1,5,"M_shenjiang4_att",6, } 
   Attendants["7"]= Attendant{ 7,"神将5",1,5,"M_shenjiang5_att",7, } 
   Attendants["8"]= Attendant{ 8,"神将6",1,5,"M_shenjiang6_att",8, } 
   Attendants["9"]= Attendant{ 9,"神将7",1,6,"M_shenjiang7_att",9, } 
   Attendants["10"]= Attendant{ 10,"神将8",1,6,"M_shenjiang8_att",10, } 
   Attendants["11"]= Attendant{ 11,"神将9",1,6,"M_shenjiang9_att",11, } 
   Attendants["12"]= Attendant{ 12,"神将10",1,7,"M_shenjiang10_att",12, } 
   Attendants["13"]= Attendant{ 13,"神将11",1,7,"M_shenjiang11_att",13, } 
   Attendants["14"]= Attendant{ 14,"神将12",1,7,"M_shenjiang12_att",14, } 
   
}
