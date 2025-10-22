package repository

import (
	"context"
	"errors"

	"github.com/Mroxny/slamIt/internal/model"
	"gorm.io/gorm"
)

type PerformanceRepository struct {
	*Repository[model.Performance]
}

func NewPerformanceRepository(db *gorm.DB) *PerformanceRepository {
	return &PerformanceRepository{
		Repository: NewRepository[model.Performance](db),
	}
}

func (r *PerformanceRepository) FindByStageId(ctx context.Context, stageId string) ([]model.Performance, error) {
	var stageCheck model.Stage
	if err := r.db.WithContext(ctx).Select("id").First(&stageCheck, "id = ?", stageId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("stage not found")
		}
		return nil, err
	}

	var performances []model.Performance
	err := r.db.WithContext(ctx).Find(&performances, "stage_id = ?", stageId).Error
	return performances, err
}

func (r *PerformanceRepository) FindByStageAndParticipation(ctx context.Context, stageId, participationId string) (*model.Performance, error) {
	var p model.Performance
	err := r.db.WithContext(ctx).Where("stage_id = ? AND participation_id = ?", stageId, participationId).First(&p).Error
	return &p, err
}

func (r *PerformanceRepository) FindByID(ctx context.Context, performanceId string) (*model.Performance, error) {
	var perf model.Performance
	err := r.db.WithContext(ctx).
		Preload("OpponentPerformance").
		First(&perf, "id = ?", performanceId).Error
	return &perf, err
}

// FindAndSortByStageID retrieves all performances for a stage and sorts them
// based on the linked-list order defined by OpponentPerformanceId.
func (r *PerformanceRepository) FindSortedByStageId(ctx context.Context, stageId string) ([]model.Performance, error) {
	var stageCheck model.Stage
	if err := r.db.WithContext(ctx).Select("id").First(&stageCheck, "id = ?", stageId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("stage not found")
		}
		return nil, err
	}

	var allPerformances []model.Performance
	if err := r.db.WithContext(ctx).Preload("Participation").Preload("Participation.User").Where("stage_id = ?", stageId).Find(&allPerformances).Error; err != nil {
		return nil, err
	}

	if len(allPerformances) == 0 {
		return []model.Performance{}, nil
	}

	// --- Sorting Logic ---
	// Build a map for quick lookups and find the head of the list.
	performanceMap := make(map[string]model.Performance)
	isPointedTo := make(map[string]bool) // Tracks which IDs are pointed to by an OpponentPerformanceId

	for _, p := range allPerformances {
		performanceMap[p.Id] = p
		if p.OpponentPerformanceId != nil {
			isPointedTo[*p.OpponentPerformanceId] = true
		}
	}

	// Find the head of the performance chain (the one no other performance points to).
	var head *model.Performance
	for _, p := range allPerformances {
		if _, ok := isPointedTo[p.Id]; !ok {
			// This ID is not pointed to by any other, so it's a potential head.
			// For simplicity, we'll just pick the first one we find.
			// A more complex scenario could have multiple chains.
			temp := p
			head = &temp
			break
		}
	}

	sortedPerformances := make([]model.Performance, 0, len(allPerformances))
	inSortedList := make(map[string]bool)

	// Traverse the linked list starting from the head.
	current := head
	for current != nil {
		sortedPerformances = append(sortedPerformances, *current)
		inSortedList[current.Id] = true

		if current.OpponentPerformanceId == nil {
			break
		}
		next, ok := performanceMap[*current.OpponentPerformanceId]
		if !ok {
			break // Chain is broken
		}
		current = &next
	}

	// Add any remaining performances (those with NULL opponents not in the chain) to the end.
	for _, p := range allPerformances {
		if _, ok := inSortedList[p.Id]; !ok {
			sortedPerformances = append(sortedPerformances, p)
		}
	}

	return sortedPerformances, nil
}

// UpdateOrderInTransaction rearranges the opponent chain for a given stage within a single DB transaction.
// orderedIDs should be a slice of performance IDs in the desired new order.
func (r *PerformanceRepository) UpdateOrderTx(ctx context.Context, stageId string, orderedIDs []string) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Step 1: Detach all performances in the stage by setting their opponent ID to NULL.
		// This prevents old chains from interfering and handles performers not in the new order.
		if err := tx.Model(&model.Performance{}).Where("stage_id = ?", stageId).Update("opponent_performance_id", nil).Error; err != nil {
			return err
		}

		if len(orderedIDs) < 2 {
			return nil
		}

		// Step 2: Iterate through the new order and link them together.
		for i := 0; i < len(orderedIDs)-1; i++ {
			currentID := orderedIDs[i]
			nextID := orderedIDs[i+1]

			if err := tx.Model(&model.Performance{}).Where("id = ?", currentID).Update("opponent_performance_id", nextID).Error; err != nil {
				return err
			}
		}

		return nil
	})
}
