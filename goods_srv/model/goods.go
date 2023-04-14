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

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null comment '名称'"`
	Logo string `gorm:"type:varchar(200);default:'';not null comment '名称'"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique'"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique'"`
	Brands     Brands
}

//// 重置表名
//func (GoodsCategoryBrand) TableName() string {
//	return "goodscategorybrand"
//}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null comment '图片'"`
	Url   string `gorm:"type:varchar(200);not null comment '跳转链接'"`
	Index int32  `gorm:"type:int;default:1;not null comment '排序'"`
}
