package bean

import (
	"github.com/GodSlave/MyGoServer/module"
	"github.com/go-xorm/xorm"
)

// model



const WorShip_Inter = 60



//EDuty
type EDuty int32
const EDuty_health = 1 // 健康
const EDuty_marriage = 2 // 姻缘
const EDuty_cause = 3 // 事业


func EnableDBCache(app module.App) {
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)

}

func ClearDBChache(app module.App) {


}

func DisableDBCache(app module.App) {


}

