package model

//类型， 这个字段是否能为null，这个字段应该设置为可以为nu1l还是设置为空，0
//实际开发过程中尽量设置为不为null
//https://zhuanlan.zhihu.com/p/73997266

type Category struct {
	BaseModel
	Name             string      `gorm:"type:varchar(20);comment:名称;not null" json:"name"`
	ParentCategoryID int32       `json:"parent"`
	ParentCategory   *Category   `json:"-"`
	SubCategory      []*Category `gorm:"foreignKey:ParentCategoryID;references:ID" json:"sub_category"`
	Level            int32       `gorm:"type:int;default:1;not null;comment:等级" json:"level"`
	IsTab            bool        `gorm:"default:false;not null;comment:是否展示在tab" json:"is_tab"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);comment:名称;not null"`
	Logo string `gorm:"type:varchar(200);default:'';comment:名称;not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique'"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;index:idx_category_brand,unique'"`
	Brands     Brands
}

// // 重置表名
func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);comment:图片;not null"`
	Url   string `gorm:"type:varchar(200);comment:跳转链接;not null"`
	Index int32  `gorm:"type:int;default:1;comment:排序;not null"`
}

type Goods struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;not null;comment:分类id;"`
	Category   Category
	BrandsID   int32 `gorm:"type:int;not null;comment:品牌id"`
	Brands     Brands

	OnSale   bool `gorm:"default:false;not null;comment:是否上架;"`
	ShipFree bool `gorm:"default:false;not null;comment:是否免运费"`
	IsNew    bool `gorm:"default:false;not null;comment:是否新品"`
	IsHot    bool `gorm:"default:false;not null;comment:是否热卖商品"`

	Name            string   `gorm:"type:varchar(100);not null;comment:名称"`
	GoodsSn         string   `gorm:"type:varchar(50);not null;comment:商品编号"`
	ClickNum        int32    `gorm:"type:int;default:0;not null;comment:点击量"`
	SoldNum         int32    `gorm:"type:int;default:0;not null;comment:销量"`
	FavNum          int32    `gorm:"type:int;default:0;not null;comment:收藏数"`
	MarketPrice     float32  `gorm:"not null;comment:市场价"`
	ShopPrice       float32  `gorm:"not null;comment:价格"`
	GoodsBrief      string   `gorm:"type:varchar(100);not null;comment:商品简介"`
	Images          GormList `gorm:"type:varchar(1000);not null;comment:图册"`
	DescImages      GormList `gorm:"type:varchar(1000);not null;comment:商品详情图"`
	GoodsFrontImage string   `gorm:"type:varchar(200);not null;comment:商品主图"`
}
