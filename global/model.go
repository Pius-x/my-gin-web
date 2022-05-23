package global

type GvaModel struct {
	ID        uint64 `gorm:"primarykey"` // 主键ID
	CreatedAt uint64 // 创建时间
	UpdatedAt uint64 // 更新时间
}
