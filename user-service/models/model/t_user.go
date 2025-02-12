package model

import (
	"comm/util"
	"database/sql"
	"fmt"
	pb "proto/user"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = fmt.State.Write
	_ = util.UnixZero
)

type User struct {
	Id             int32  `gorm:"column:FID;primary_key" json:"id"`
	UserId         int64  `gorm:"column:Fuser_id" json:"user_id"`
	UserName       string `gorm:"column:Fuser_name" json:"user_name"`
	NickName       string `gorm:"column:Fnick_name" json:"nick_name"`
	Sex            int32  `gorm:"column:Fsex" json:"sex"`
	Age            int32  `gorm:"column:Fage" json:"age"`
	Phone          int64  `gorm:"column:Fphone" json:"phone"`
	Email          string `gorm:"column:Femail" json:"email"`
	RegTime        int32  `gorm:"column:Freg_time" json:"reg_time"`
	LastLoginTime  int32  `gorm:"column:Flast_login_time" json:"last_login_time"`
	LastLoginIp    string `gorm:"column:Flast_login_ip" json:"last_login_ip"`
	LastUpdateTime int32  `gorm:"column:Flast_update_time" json:"last_update_time"`
	UserBirth      string `gorm:"column:Fuser_birth" json:"user_birth"`
	WechatId       string `gorm:"column:Fwechat_id" json:"wechat_id"`
	UserAddress    string `gorm:"column:Fuser_Address" json:"user_address"`
	UserStatus     int32  `gorm:"column:Fuser_status" json:"user_status"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return fmt.Sprintf("t_user_%03d", u.UserId%1000)
}

func (u *User) Marshal(o *pb.User) error {
	o.Id = u.Id
	o.UserId = u.UserId
	o.UserName = u.UserName
	o.NickName = u.NickName
	o.Sex = u.Sex
	o.Age = u.Age
	o.Email = u.Email
	o.Phone = u.Phone
	o.RegTime = u.RegTime
	o.LastLoginTime = u.LastLoginTime
	o.LastLoginIp = u.LastLoginIp
	o.LastUpdateTime = u.LastUpdateTime
	o.UserBirth = u.UserBirth
	o.UserAddress = u.UserAddress
	o.UserStatus = u.UserStatus

	return nil
}

func (u *User) Unmarshal(o *pb.User) error {
	u.Id = o.Id
	u.UserId = o.UserId
	u.UserName = o.UserName
	u.NickName = o.NickName
	u.Sex = o.Sex
	u.Age = o.Age
	u.Email = o.Email
	u.Phone = o.Phone
	u.RegTime = o.RegTime
	u.LastLoginTime = o.LastLoginTime
	u.LastLoginIp = o.LastLoginIp
	u.LastUpdateTime = o.LastUpdateTime
	u.UserBirth = o.UserBirth
	u.UserAddress = o.UserAddress
	u.UserStatus = o.UserStatus

	return nil
}

func UserMarshalLst(lst []User, pbLst *[]*pb.User) {
	for _, v := range lst {
		var tmp pb.User
		v.Marshal(&tmp)
		*pbLst = append(*pbLst, &tmp)
	}
}

func UserUnmarshalLst(pbLst []*pb.User, lst *[]User) {
	for _, v := range pbLst {
		var tmp User
		tmp.Unmarshal(v)
		*lst = append(*lst, tmp)
	}
}
