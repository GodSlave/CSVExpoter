
package bean

import (
	"github.com/GodSlave/MyGoServer/db"
	"os/exec"
	"os"
	"path/filepath"
	"fmt"
	"flag"
	"github.com/GodSlave/MyGoServer/conf"
)

func Init()  {
	file, _ := exec.LookPath(os.Args[0])
	ApplicationPath, _ := filepath.Abs(file)
	ApplicationDir, _ := filepath.Split(ApplicationPath)
	defaultPath := fmt.Sprintf("%sconf"+string(filepath.Separator)+"server.json", ApplicationDir)
	confPath := flag.String("conf", defaultPath, "Server configuration file path")
	flag.Parse() //解析输入的参数
	f, err := os.Open(*confPath)
	if err != nil {
		panic(err)
	}
	conf.LoadConfig(f.Name()) //加载配置文件
	//sql
	sql := db.BaseSql{
	}
	sql.Url = conf.Conf.DB.SQL
	sql.InitDB()
	sql.CheckMigrate()
	defer sql.Engine.Close()


	sql.Engine.DropTables(&Buddha{})
	sql.Engine.Sync2(&Buddha{})
   
 	sql.Engine.Insert( Buddha{ 1,"药师",true,[]int32{1,2},"M_yaoshi_fo","普陀山","Map_putuo","","",4, })
   
 	sql.Engine.Insert( Buddha{ 2,"观音",true,[]int32{2,3},"M_guanyin_fo","南海","Map_nanhai","","",5, })
   
 	sql.Engine.Insert( Buddha{ 3,"释迦摩尼",true,[]int32{1,3},"M_shijiamoni_fo","灵山","Map_lingshan","","",6, })
   
	sql.Engine.DropTables(&BuddhaLevel{})
	sql.Engine.Sync2(&BuddhaLevel{})
   
 	sql.Engine.Insert( BuddhaLevel{ 1,1,1,[]ItemBase{{101,3,},{1,10,},},10,100,50,20,1,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 2,1,2,[]ItemBase{{101,3,},{102,3,},{2,20,},},20,200,100,40,1,"",[]int32{1,2}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 3,1,3,[]ItemBase{{101,5,},{102,6,},{3,50,},},40,400,200,80,2,"",[]int32{1,2,3,4,5}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 4,1,4,[]ItemBase{{101,8,},{102,7,},{103,3,},{4,99,},},80,800,400,160,2,"",[]int32{1,2,3,4,5,6,7,8}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 5,1,5,[]ItemBase{{101,10,},{102,9,},{103,6,},},160,1600,800,320,3,"",[]int32{1,2,3,4,5,6,7,8,9,10,11}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 6,1,6,[]ItemBase{{101,15,},{102,10,},{103,8,},},320,3200,1600,640,4,"",[]int32{1,2,3,4,5,6,7,8,9,10,11,12,13,14}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 7,1,7,[]ItemBase{{101,20,},{102,12,},{103,10,},},700,6400,3200,1280,5,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 8,2,1,[]ItemBase{{101,3,},{1,10,},},10,100,50,20,6,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 9,2,2,[]ItemBase{{101,3,},{102,3,},{2,20,},},20,200,100,40,6,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 10,2,3,[]ItemBase{{101,5,},{102,6,},{3,50,},},40,400,200,80,7,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 11,2,4,[]ItemBase{{101,8,},{102,7,},{103,3,},{4,99,},},80,800,400,160,7,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 12,2,5,[]ItemBase{{101,10,},{102,9,},{103,6,},},160,1600,800,320,8,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 13,2,6,[]ItemBase{{101,15,},{102,10,},{103,8,},},320,3200,1600,640,9,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 14,2,7,[]ItemBase{{101,20,},{102,12,},{103,10,},},700,6400,3200,1280,10,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 15,3,1,[]ItemBase{{101,3,},{1,10,},},10,100,50,20,11,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 16,3,2,[]ItemBase{{101,3,},{102,3,},{2,20,},},20,200,100,40,11,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 17,3,3,[]ItemBase{{101,5,},{102,6,},{3,50,},},40,400,200,80,12,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 18,3,4,[]ItemBase{{101,8,},{102,7,},{103,3,},{4,99,},},80,800,400,160,12,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 19,3,5,[]ItemBase{{101,10,},{102,9,},{103,6,},},160,1600,800,320,13,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 20,3,6,[]ItemBase{{101,15,},{102,10,},{103,8,},},320,3200,1600,640,14,"",[]int32{}, })
   
 	sql.Engine.Insert( BuddhaLevel{ 21,3,7,[]ItemBase{{101,20,},{102,12,},{103,10,},},700,6400,3200,1280,15,"",[]int32{}, })
   
	sql.Engine.DropTables(&Birth{})
	sql.Engine.Sync2(&Birth{})
   
 	sql.Engine.Insert( Birth{ 1,[]ItemBase{{101,99,},{102,99,},{103,99,},}, })
   
	sql.Engine.DropTables(&Quality{})
	sql.Engine.Sync2(&Quality{})
   
 	sql.Engine.Insert( Quality{ 1,"杉木", })
   
 	sql.Engine.Insert( Quality{ 2,"松木", })
   
 	sql.Engine.Insert( Quality{ 3,"柏木", })
   
 	sql.Engine.Insert( Quality{ 4,"檀木", })
   
 	sql.Engine.Insert( Quality{ 5,"沉香木", })
   
	sql.Engine.DropTables(&WishName{})
	sql.Engine.Sync2(&WishName{})
   
	sql.Engine.DropTables(&Dropitem{})
	sql.Engine.Sync2(&Dropitem{})
   
 	sql.Engine.Insert( Dropitem{ 1,[]ItemDrop{{1001,1,7000,},{1002,1,3000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 2,[]ItemDrop{{1002,1,7000,},{1003,1,3000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 3,[]ItemDrop{{1002,1,5000,},{1003,1,4000,},{1004,1,1000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 4,[]ItemDrop{{1003,1,8000,},{1004,1,2000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 5,[]ItemDrop{{1003,1,5000,},{1004,1,4000,},{1005,1,1000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 6,[]ItemDrop{{2001,1,7000,},{2002,1,3000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 7,[]ItemDrop{{2002,1,7000,},{2003,1,3000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 8,[]ItemDrop{{2002,1,5000,},{2003,1,4000,},{2004,1,1000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 9,[]ItemDrop{{2003,1,8000,},{2004,1,2000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 10,[]ItemDrop{{2003,1,5000,},{2004,1,4000,},{2005,1,1000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 11,[]ItemDrop{{2001,1,7000,},{2002,1,3000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 12,[]ItemDrop{{3002,1,7000,},{3003,1,3000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 13,[]ItemDrop{{3002,1,5000,},{3003,1,4000,},{3004,1,1000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 14,[]ItemDrop{{3003,1,8000,},{3004,1,2000,},}, })
   
 	sql.Engine.Insert( Dropitem{ 15,[]ItemDrop{{3003,1,5000,},{3004,1,4000,},{3005,1,1000,},}, })
   
	sql.Engine.DropTables(&DropItemInfo{})
	sql.Engine.Sync2(&DropItemInfo{})
   
 	sql.Engine.Insert( DropItemInfo{  })
   
	sql.Engine.DropTables(&Poem{})
	sql.Engine.Sync2(&Poem{})
   
 	sql.Engine.Insert( Poem{ 1000,[]string{"百岁为上寿，","一言乃千金，","盛世常青树，","百年不老松。"},2,"祝福健康长寿，寄托对长辈的美好愿望",3, })
   
 	sql.Engine.Insert( Poem{ 2000,[]string{"夕阳无限好，","只是近黄昏，","莫道桑榆晚，","为霞尚满天。"},2,"祝福长辈身体健康，生活快乐",3, })
   
 	sql.Engine.Insert( Poem{ 3000,[]string{"桃之夭夭，","有蕡其实，","之子于归，","宜其家室。"},2,"祝福家庭和睦，生活幸福，早生贵子",3, })
   
 	sql.Engine.Insert( Poem{ 4000,[]string{"关关雎鸠，","在河之洲，","窈窕淑女，","君子好逑。"},2,"祝新婚快乐，生活幸福",3, })
   
 	sql.Engine.Insert( Poem{ 5000,[]string{"白日依山尽，","黄河入海流，","欲穷千里目，","更上一层楼。"},2,"祝福事业顺利，大吉大利，今晚吃鸡",3, })
   
 	sql.Engine.Insert( Poem{ 6000,[]string{"富在术数，","不在劳身，","利在势居，","不在力耕。"},2,"祝福事少钱多离家近，年薪百万",3, })
   
	sql.Engine.DropTables(&WorShip{})
	sql.Engine.Sync2(&WorShip{})
   
 	sql.Engine.Insert( WorShip{ 10001,1,5,300, })
   
 	sql.Engine.Insert( WorShip{ 10002,1,10,1800, })
   
 	sql.Engine.Insert( WorShip{ 10003,1,20,3600, })
   
 	sql.Engine.Insert( WorShip{ 10101,2,5,300, })
   
 	sql.Engine.Insert( WorShip{ 10102,2,10,1800, })
   
 	sql.Engine.Insert( WorShip{ 10103,2,20,3600, })
   
 	sql.Engine.Insert( WorShip{ 10201,3,5,300, })
   
 	sql.Engine.Insert( WorShip{ 10202,3,10,1800, })
   
 	sql.Engine.Insert( WorShip{ 10203,3,20,3600, })
   
	sql.Engine.DropTables(&ItemBase{})
	sql.Engine.Sync2(&ItemBase{})
   
	sql.Engine.DropTables(&ItemDrop{})
	sql.Engine.Sync2(&ItemDrop{})
   
	sql.Engine.DropTables(&Duty{})
	sql.Engine.Sync2(&Duty{})
   
 	sql.Engine.Insert( Duty{ 1,"health","健康",1, })
   
 	sql.Engine.Insert( Duty{ 2,"marriage","姻缘",2, })
   
 	sql.Engine.Insert( Duty{ 3,"career","事业",3, })
   
	sql.Engine.DropTables(&WishCard{})
	sql.Engine.Sync2(&WishCard{})
   
 	sql.Engine.Insert( WishCard{ 1001,"健康牌","UI_HealthCard_Icon",1,[]string{"福","如","东","海"},1000,10000,100,1,4,100, })
   
 	sql.Engine.Insert( WishCard{ 1002,"健康牌","UI_HealthCard_Icon",2,[]string{"福","如","东","海"},1000,20000,200,1,4,200, })
   
 	sql.Engine.Insert( WishCard{ 1003,"健康牌","UI_HealthCard_Icon",3,[]string{"福","如","东","海"},1000,30000,300,1,4,400, })
   
 	sql.Engine.Insert( WishCard{ 1004,"健康牌","UI_HealthCard_Icon",4,[]string{"福","如","东","海"},1000,40000,400,1,4,800, })
   
 	sql.Engine.Insert( WishCard{ 1005,"健康牌","UI_HealthCard_Icon",5,[]string{"福","如","东","海"},1000,50000,500,1,4,1600, })
   
 	sql.Engine.Insert( WishCard{ 2001,"健康牌","UI_HealthCard_Icon",1,[]string{"寿","比","南","山"},2000,10000,100,1,6,100, })
   
 	sql.Engine.Insert( WishCard{ 2001,"健康牌","UI_HealthCard_Icon",2,[]string{"寿","比","南","山"},2000,20000,200,1,6,200, })
   
 	sql.Engine.Insert( WishCard{ 2003,"健康牌","UI_HealthCard_Icon",3,[]string{"寿","比","南","山"},2000,30000,300,1,6,400, })
   
 	sql.Engine.Insert( WishCard{ 2004,"健康牌","UI_HealthCard_Icon",4,[]string{"寿","比","南","山"},2000,40000,400,1,6,800, })
   
 	sql.Engine.Insert( WishCard{ 2005,"健康牌","UI_HealthCard_Icon",5,[]string{"寿","比","南","山"},2000,50000,500,1,6,1600, })
   
 	sql.Engine.Insert( WishCard{ 3001,"姻缘牌","UI_MerriageCard_Icon",1,[]string{"早","生","贵","子"},3000,10000,100,1,4,100, })
   
 	sql.Engine.Insert( WishCard{ 3002,"姻缘牌","UI_MerriageCard_Icon",2,[]string{"早","生","贵","子"},3000,20000,200,2,4,200, })
   
 	sql.Engine.Insert( WishCard{ 3003,"姻缘牌","UI_MerriageCard_Icon",3,[]string{"早","生","贵","子"},3000,30000,300,2,4,400, })
   
 	sql.Engine.Insert( WishCard{ 3004,"姻缘牌","UI_MerriageCard_Icon",4,[]string{"早","生","贵","子"},3000,40000,400,2,4,800, })
   
 	sql.Engine.Insert( WishCard{ 3005,"姻缘牌","UI_MerriageCard_Icon",5,[]string{"早","生","贵","子"},3000,50000,500,2,4,1600, })
   
 	sql.Engine.Insert( WishCard{ 4001,"姻缘牌","UI_MerriageCard_Icon",1,[]string{"百","年","好","合"},4000,10000,100,2,5,100, })
   
 	sql.Engine.Insert( WishCard{ 4002,"姻缘牌","UI_MerriageCard_Icon",2,[]string{"百","年","好","合"},4000,20000,200,2,5,200, })
   
 	sql.Engine.Insert( WishCard{ 4003,"姻缘牌","UI_MerriageCard_Icon",3,[]string{"百","年","好","合"},4000,30000,300,2,5,400, })
   
 	sql.Engine.Insert( WishCard{ 4004,"姻缘牌","UI_MerriageCard_Icon",4,[]string{"百","年","好","合"},4000,40000,400,2,5,800, })
   
 	sql.Engine.Insert( WishCard{ 4005,"姻缘牌","UI_MerriageCard_Icon",5,[]string{"百","年","好","合"},4000,50000,500,2,5,1600, })
   
 	sql.Engine.Insert( WishCard{ 5001,"事业牌","UI_CareerCard_Icon",1,[]string{"一","步","登","天"},5000,10000,100,3,5,100, })
   
 	sql.Engine.Insert( WishCard{ 5002,"事业牌","UI_CareerCard_Icon",2,[]string{"一","步","登","天"},5000,20000,200,3,5,200, })
   
 	sql.Engine.Insert( WishCard{ 5003,"事业牌","UI_CareerCard_Icon",3,[]string{"一","步","登","天"},5000,30000,300,3,5,400, })
   
 	sql.Engine.Insert( WishCard{ 5004,"事业牌","UI_CareerCard_Icon",4,[]string{"一","步","登","天"},5000,40000,400,3,5,800, })
   
 	sql.Engine.Insert( WishCard{ 5005,"事业牌","UI_CareerCard_Icon",5,[]string{"一","步","登","天"},5000,50000,500,3,5,1600, })
   
 	sql.Engine.Insert( WishCard{ 6001,"事业牌","UI_CareerCard_Icon",1,[]string{"财","源","广","进"},6000,10000,100,3,6,100, })
   
 	sql.Engine.Insert( WishCard{ 6002,"事业牌","UI_CareerCard_Icon",2,[]string{"财","源","广","进"},6000,20000,200,3,6,200, })
   
 	sql.Engine.Insert( WishCard{ 6003,"事业牌","UI_CareerCard_Icon",3,[]string{"财","源","广","进"},6000,30000,300,3,6,400, })
   
 	sql.Engine.Insert( WishCard{ 6004,"事业牌","UI_CareerCard_Icon",4,[]string{"财","源","广","进"},6000,40000,400,3,6,800, })
   
 	sql.Engine.Insert( WishCard{ 6005,"事业牌","UI_CareerCard_Icon",5,[]string{"财","源","广","进"},6000,50000,500,3,6,1600, })
   
	sql.Engine.DropTables(&Item{})
	sql.Engine.Sync2(&Item{})
   
 	sql.Engine.Insert( Item{ 1,"功德币","UI_gongdebi_icon",99999999, })
   
 	sql.Engine.Insert( Item{ 2,"修行值","UI_xiuxingzhi_icon",99999999, })
   
 	sql.Engine.Insert( Item{ 3,"福报值","UI_fubaozhi_icon",99999999, })
   
 	sql.Engine.Insert( Item{ 4,"药师如来愿力值","UI_yuanlizhi1_icon",99999999, })
   
 	sql.Engine.Insert( Item{ 5,"观音愿力值","UI_yuanlizhi2_icon",99999999, })
   
 	sql.Engine.Insert( Item{ 6,"释迦摩尼愿力值","UI_yuanlizhi3_icon",99999999, })
   
 	sql.Engine.Insert( Item{ 101,"玉净瓶","UI_yujingping_icon",99, })
   
 	sql.Engine.Insert( Item{ 102,"玲珑宝塔","UI_baota_icon",99, })
   
 	sql.Engine.Insert( Item{ 103,"莲花","UI_lianhua_icon",99, })
   
 	sql.Engine.Insert( Item{ 1001,"杉木健康牌-福如东海","UI_fupai1_icon",99, })
   
 	sql.Engine.Insert( Item{ 1002,"松木健康牌-福如东海","UI_fupai2_icon",99, })
   
 	sql.Engine.Insert( Item{ 1003,"柏木健康牌-福如东海","UI_fupai3_icon",99, })
   
 	sql.Engine.Insert( Item{ 1004,"檀木健康牌-福如东海","UI_fupai4_icon",99, })
   
 	sql.Engine.Insert( Item{ 1005,"沉香木健康牌-福如东海","UI_fupai5_icon",99, })
   
 	sql.Engine.Insert( Item{ 2001,"杉木健康牌-寿比南山","UI_fupai6_icon",99, })
   
 	sql.Engine.Insert( Item{ 2001,"松木健康牌-寿比南山","UI_fupai7_icon",99, })
   
 	sql.Engine.Insert( Item{ 2003,"柏木健康牌-寿比南山","UI_fupai8_icon",99, })
   
 	sql.Engine.Insert( Item{ 2004,"檀木健康牌-寿比南山","UI_fupai9_icon",99, })
   
 	sql.Engine.Insert( Item{ 2005,"沉香木健康牌-寿比南山","UI_fupai10_icon",99, })
   
 	sql.Engine.Insert( Item{ 3001,"杉木姻缘牌-早生贵子","UI_fupai11_icon",99, })
   
 	sql.Engine.Insert( Item{ 3002,"松木姻缘牌-早生贵子","UI_fupai12_icon",99, })
   
 	sql.Engine.Insert( Item{ 3003,"柏木姻缘牌-早生贵子","UI_fupai13_icon",99, })
   
 	sql.Engine.Insert( Item{ 3004,"檀木姻缘牌-早生贵子","UI_fupai14_icon",99, })
   
 	sql.Engine.Insert( Item{ 3005,"沉香木姻缘牌-早生贵子","UI_fupai15_icon",99, })
   
 	sql.Engine.Insert( Item{ 4001,"杉木姻缘牌-百年好合","UI_fupai16_icon",99, })
   
 	sql.Engine.Insert( Item{ 4002,"松木姻缘牌-百年好合","UI_fupai17_icon",99, })
   
 	sql.Engine.Insert( Item{ 4003,"柏木姻缘牌-百年好合","UI_fupai18_icon",99, })
   
 	sql.Engine.Insert( Item{ 4004,"檀木姻缘牌-百年好合","UI_fupai19_icon",99, })
   
 	sql.Engine.Insert( Item{ 4005,"沉香木姻缘牌-百年好合","UI_fupai20_icon",99, })
   
 	sql.Engine.Insert( Item{ 5001,"杉木事业牌-一步登天","UI_fupai21_icon",99, })
   
 	sql.Engine.Insert( Item{ 5002,"松木事业牌-一步登天","UI_fupai22_icon",99, })
   
 	sql.Engine.Insert( Item{ 5003,"柏木事业牌-一步登天","UI_fupai23_icon",99, })
   
 	sql.Engine.Insert( Item{ 5004,"檀木事业牌-一步登天","UI_fupai24_icon",99, })
   
 	sql.Engine.Insert( Item{ 5005,"沉香木事业牌-财源广进","UI_fupai25_icon",99, })
   
 	sql.Engine.Insert( Item{ 6001,"杉木事业牌-财源广进","UI_fupai26_icon",99, })
   
 	sql.Engine.Insert( Item{ 6002,"松木事业牌-财源广进","UI_fupai27_icon",99, })
   
 	sql.Engine.Insert( Item{ 6003,"柏木事业牌-财源广进","UI_fupai28_icon",99, })
   
 	sql.Engine.Insert( Item{ 6004,"檀木事业牌-财源广进","UI_fupai29_icon",99, })
   
 	sql.Engine.Insert( Item{ 6005,"沉香木事业牌-财源广进","UI_fupai30_icon",99, })
   
 	sql.Engine.Insert( Item{ 10001,"低级细香","UI_lifo1_icon",99, })
   
 	sql.Engine.Insert( Item{ 10002,"中级细香","UI_lifo2_icon",99, })
   
 	sql.Engine.Insert( Item{ 10003,"高级细香","UI_lifo3_icon",99, })
   
 	sql.Engine.Insert( Item{ 10101,"菊花","UI_lifo7_icon",99, })
   
 	sql.Engine.Insert( Item{ 10102,"桃花","UI_lifo8_icon",99, })
   
 	sql.Engine.Insert( Item{ 10103,"喇叭花","UI_lifo9_icon",99, })
   
 	sql.Engine.Insert( Item{ 10201,"黄蜡烛","UI_lifo13_icon",99, })
   
 	sql.Engine.Insert( Item{ 10202,"红蜡烛","UI_lifo14_icon",99, })
   
 	sql.Engine.Insert( Item{ 10203,"龙凤烛","UI_lifo15_icon",99, })
   
	sql.Engine.DropTables(&SacredLand{})
	sql.Engine.Sync2(&SacredLand{})
   
 	sql.Engine.Insert( SacredLand{ 1,0,0, })
   
 	sql.Engine.Insert( SacredLand{ 2,1,3, })
   
 	sql.Engine.Insert( SacredLand{ 3,2,3, })
   
 	sql.Engine.Insert( SacredLand{ 4,3,5, })
   
 	sql.Engine.Insert( SacredLand{ 5,4,5, })
   
 	sql.Engine.Insert( SacredLand{ 6,5,6, })
   
 	sql.Engine.Insert( SacredLand{ 7,6,7, })
   
	sql.Engine.DropTables(&Attendant{})
	sql.Engine.Sync2(&Attendant{})
   
 	sql.Engine.Insert( Attendant{ 1,"日光菩萨",1,2,"M_riguang_att",1, })
   
 	sql.Engine.Insert( Attendant{ 2,"月光菩萨",1,3,"M_yueguang_att",2, })
   
 	sql.Engine.Insert( Attendant{ 3,"神将1",1,4,"M_shenjiang1_att",3, })
   
 	sql.Engine.Insert( Attendant{ 4,"神将2",1,4,"M_shenjiang2_att",4, })
   
 	sql.Engine.Insert( Attendant{ 5,"神将3",1,4,"M_shenjiang3_att",5, })
   
 	sql.Engine.Insert( Attendant{ 6,"神将4",1,5,"M_shenjiang4_att",6, })
   
 	sql.Engine.Insert( Attendant{ 7,"神将5",1,5,"M_shenjiang5_att",7, })
   
 	sql.Engine.Insert( Attendant{ 8,"神将6",1,5,"M_shenjiang6_att",8, })
   
 	sql.Engine.Insert( Attendant{ 9,"神将7",1,6,"M_shenjiang7_att",9, })
   
 	sql.Engine.Insert( Attendant{ 10,"神将8",1,6,"M_shenjiang8_att",10, })
   
 	sql.Engine.Insert( Attendant{ 11,"神将9",1,6,"M_shenjiang9_att",11, })
   
 	sql.Engine.Insert( Attendant{ 12,"神将10",1,7,"M_shenjiang10_att",12, })
   
 	sql.Engine.Insert( Attendant{ 13,"神将11",1,7,"M_shenjiang11_att",13, })
   
 	sql.Engine.Insert( Attendant{ 14,"神将12",1,7,"M_shenjiang12_att",14, })
   
}
