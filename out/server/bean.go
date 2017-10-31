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
   Produce     int32//每单位时间产出的愿力值
   DropId     int32//掉落id

}


const WorShip_Inter = 60



//EDuty
type EDuty int32
const EDuty_health = 1 // 健康
const EDuty_marriage = 2 // 姻缘
const EDuty_cause = 3 // 事业


var Buddhas  []Buddha
var BuddhaLevels  map[string]BuddhaLevel


func init()  {

   
   
   Buddhas = append(Buddhas , Buddha{ 1,"药师",1,"M_yaoshi_fo","普陀山","Map_putuo", })
   Buddhas = append(Buddhas , Buddha{ 2,"观音",2,"M_guanyin_fo","南海","Map_nanhai", })
   Buddhas = append(Buddhas , Buddha{ 3,"释迦摩尼",3,"M_shijiamoni_fo","灵山","Map_lingshan", })
   
   
   BuddhaLevels =  map[string]BuddhaLevel{}
   
   BuddhaLevels["1"]= BuddhaLevel{ 1,1,1,[]ItemBase{101,3,1,10},10,1, } 
   BuddhaLevels["2"]= BuddhaLevel{ 2,1,2,[]ItemBase{101,3,102,3,2,20},20,1, } 
   BuddhaLevels["3"]= BuddhaLevel{ 3,1,3,[]ItemBase{101,5,102,6,3,50},40,2, } 
   BuddhaLevels["4"]= BuddhaLevel{ 4,1,4,[]ItemBase{101,8,102,7,103,3,4,99},80,2, } 
   BuddhaLevels["5"]= BuddhaLevel{ 5,1,5,[]ItemBase{101,10,102,9,103,6},160,3, } 
   BuddhaLevels["6"]= BuddhaLevel{ 6,1,6,[]ItemBase{101,15,102,10,103,8},320,4, } 
   BuddhaLevels["7"]= BuddhaLevel{ 7,1,7,[]ItemBase{101,20,102,12,103,10},700,5, } 
   BuddhaLevels["8"]= BuddhaLevel{ 8,2,1,[]ItemBase{101,3,1,10},10,1, } 
   BuddhaLevels["9"]= BuddhaLevel{ 9,2,2,[]ItemBase{101,3,102,3,2,20},20,1, } 
   BuddhaLevels["10"]= BuddhaLevel{ 10,2,3,[]ItemBase{101,5,102,6,3,50},40,2, } 
   BuddhaLevels["11"]= BuddhaLevel{ 11,2,4,[]ItemBase{101,8,102,7,103,3,4,99},80,2, } 
   BuddhaLevels["12"]= BuddhaLevel{ 12,2,5,[]ItemBase{101,10,102,9,103,6},160,3, } 
   BuddhaLevels["13"]= BuddhaLevel{ 13,2,6,[]ItemBase{101,15,102,10,103,8},320,4, } 
   BuddhaLevels["14"]= BuddhaLevel{ 14,2,7,[]ItemBase{101,20,102,12,103,10},700,5, } 
   
}
