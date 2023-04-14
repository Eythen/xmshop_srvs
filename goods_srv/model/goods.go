package model

//类型， 这个字段是否能为null，这个字段应该设置为可以为nu1l还是设置为空，0
//实际开发过程中尽量设置为不为null
//https://zhuanlan.zhihu.com/p/73997266

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20);not null comment '名称'"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1 comment '等级'"`
	IsTab            bool  `gorm:"default:false;not null comment '是否展现在tab栏'"`
}
