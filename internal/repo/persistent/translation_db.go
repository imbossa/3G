package persistent

import (
	"context"
	"fmt"
	"github.com/imbossa/3G/internal/entity"
	"github.com/imbossa/3G/pkg/db"
)

const _defaultEntityCap = 64

// TranslationRepo -.
type TranslationRepo struct {
	db *db.DB // 使用新的 db 包
}

// New -.
func New(db *db.DB) *TranslationRepo {
	return &TranslationRepo{db}
}

// GetHistory -.
func (r *TranslationRepo) GetHistory(ctx context.Context) ([]entity.Translation, error) {
	var entities []entity.Translation

	err := r.db.Conn.WithContext(ctx).
		Table("history").
		Select("source, destination, original, translation").
		Find(&entities).Error
	if err != nil {
		return nil, fmt.Errorf("TranslationRepo - GetHistory - gorm.Find: %w", err)
	}

	// 确保不会返回 nil 切片
	if entities == nil {
		entities = make([]entity.Translation, 0)
	}

	return entities, nil
}

// Store -.
func (r *TranslationRepo) Store(ctx context.Context, t entity.Translation) error {
	result := r.db.Conn.WithContext(ctx).
		Table("history").
		Create(&t)

	if result.Error != nil {
		return fmt.Errorf("TranslationRepo - Store - gorm.Create: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("TranslationRepo - Store - no rows affected")
	}

	return nil
}
