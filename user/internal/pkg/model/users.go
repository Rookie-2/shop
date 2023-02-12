package model

type BaseModel struct {
	Id        int64 `gorm:"column:id;bigint(20);primary_key;comment:主键ID" json:"id"`
	IsDeleted uint  `gorm:"column:is_deleted;type:tinyint(3) unsigned;default:0;comment:状态, 0-可用,1-不可用;NOT NULL" json:"is_deleted"`
	CreatedOn int64 `gorm:"column:created_on;type:bigint(20);comment:创建时间" json:"created_on"`
	UpdatedOn int64 `gorm:"column:updated_on;type:bigint(20);comment:更新时间" json:"updated_on"`
	DeletedOn int64 `gorm:"column:deleted_on;type:bigint(20);comment:删除时间" json:"deleted_on"`
}

type UserM struct {
	BaseModel
	UserId   string `gorm:"column:user_id;type:char(32);comment:业务唯一ID;NOT NULL" json:"user_id"`
	Mobile   string `gorm:"column:mobile;type:varchar(20);comment:用户手机号;NOT NULL" json:"mobile"`
	Password string `gorm:"column:password;type:varchar(100);comment:密码;NOT NULL" json:"password"`
	NikeName string `gorm:"column:nike_name;type:char(20);comment:用户名" json:"nike_name"`
	Birthday int64  `gorm:"column:birthday;type:bigint(20);comment:生日" json:"birthday"`
	Gender   string `gorm:"column:gender;type:varchar(6);default:male;comment:性别 female 为女，male 为男" json:"gender"`
	Role     int    `gorm:"column:role;type:tinyint(3);default:1;comment:角色1 为普通用户,2为管理员" json:"role"`
}

func (u UserM) TableName() string {
	return ""
}
