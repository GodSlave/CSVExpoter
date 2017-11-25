
package bean

import "github.com/GodSlave/MyGoServer/base"

var (

	
	UserLogin_LoginSuccess     = base.NewError(0, "登陆成功")
	
	UserLogin_LoginFailure     = base.NewError(1, "登陆失败")
	

	
	Register_Success     = base.NewError(0, "成功")
	
	Register_NameOrPasswordFormatError     = base.NewError(1, "用户名或密码格式错误")
	
	Register_NameHasBeenTaken     = base.NewError(2, "用户名已经被占用")
	

	
	GetVerifyCode_Success     = base.NewError(0, "成功")
	
	GetVerifyCode_NotReady     = base.NewError(1, "发送太频繁")
	

	
	GetSelfInfo_Success     = base.NewError(0, "成功")
	

	
	CreateRole_Success     = base.NewError(0, "成功")
	
	CreateRole_AlreadyHaveRole     = base.NewError(1, "已经创建过角色了")
	
	CreateRole_NameError     = base.NewError(2, "角色名称不合格")
	

	
	GetRoleInfo_Success     = base.NewError(0, "成功")
	
	GetRoleInfo_DoNotHaveRole     = base.NewError(1, "还没有角色")
	

	
	GetBuddhaInfo_Success     = base.NewError(0, "成功")
	
	GetBuddhaInfo_DoNotHaveBuddha     = base.NewError(1, "没有指定佛像")
	

	
	ActiveBuddha_Success     = base.NewError(0, "成功")
	
	ActiveBuddha_LevelNotSatisfied     = base.NewError(1, "等级条件不满足")
	
	ActiveBuddha_MateriaNotSatified     = base.NewError(2, "材料不满足")
	
	ActiveBuddha_AlreadyActived     = base.NewError(3, "已经激活过佛了")
	
	ActiveBuddha_BuddhaIdError     = base.NewError(4, "佛像ID错误")
	

	
	UpGradeBuddha_Success     = base.NewError(0, "成功")
	
	UpGradeBuddha_LevelNotSatisfied     = base.NewError(1, "等级条件不满足")
	
	UpGradeBuddha_MateriaNotSatified     = base.NewError(2, "材料不满足")
	
	UpGradeBuddha_NotActiveBuddha     = base.NewError(3, "当前佛像没有激活")
	

	
	WorshipBuddha_Success     = base.NewError(0, "成功")
	
	WorshipBuddha_MatiaNotSatisFied     = base.NewError(1, "材料不足")
	
	WorshipBuddha_Processing     = base.NewError(2, "正在拜佛中。。。")
	
	WorshipBuddha_NotActiveBuddha     = base.NewError(3, "当前佛像没有激活")
	

	
	CollectDesire_Success     = base.NewError(0, "成功")
	

	
	SwitchBuddha_Success     = base.NewError(0, "成功")
	
	SwitchBuddha_NotActiveThisBuddha     = base.NewError(1, "没有激活佛")
	

	
	GetAllStore_Success     = base.NewError(0, "成功")
	

	
	GetBaseInfo_Success     = base.NewError(0, "成功")
	

)