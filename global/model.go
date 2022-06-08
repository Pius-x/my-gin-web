package global

type GvaModel struct {
	ID        int64 `db:"id" from:"id" json:"id"`                         // 主键ID
	CreatedAt int64 `db:"created_at" from:"created_at" json:"created_at"` // 创建时间
	UpdatedAt int64 `db:"updated_at" from:"updated_at" json:"updated_at"` // 更新时间
}
