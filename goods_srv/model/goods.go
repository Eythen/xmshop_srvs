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

type Goods struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;not null comment '分类id'"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null comment '品牌id'"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null comment '是否上架'"`
	ShipFree bool `gorm:"default:false;not null comment '是否免运费'"`
	IsNew    bool `gorm:"default:false;not null comment '是否新品'"`
	IsHot    bool `gorm:"default:false;not null comment '是否热卖商品'"`

	Name            string   `gorm:"type:varchar(50);not null comment '名称'"`
	GoodsSn         string   `gorm:"type:varchar(50);not null comment '商品编号'"`
	ClickNum        int32    `gorm:"type:int;default:0;not null comment '点击量'"`
	SoldNum         int32    `gorm:"type:int;default:0;not null comment '销量'"`
	FavNum          int32    `gorm:"type:int;default:0;not null comment '收藏数'"`
	MarketPrice     float32  `gorm:"not null comment '市场价'"`
	ShopPrice       float32  `gorm:"not null comment '价格'"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null comment '商品简介'"`
	Images          GormList `gorm:"type:varchar(1000);not null comment '图册'"`
	DescImages      GormList `gorm:"type:varchar(1000);not null comment '商品详情图'"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null comment '商品主图'"`
}
