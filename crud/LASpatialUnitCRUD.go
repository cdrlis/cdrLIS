package crud

import (
	"errors"
	"github.com/cdrlis/cdrLIS/ladm"
	"github.com/jinzhu/gorm"
	"time"
)

type LASpatialUnitCRUD struct {
	DB *gorm.DB
}

func (crud LASpatialUnitCRUD) Read(where ...interface{}) (interface{}, error) {
	var spatialUnit ladm.LASpatialUnit
	if where != nil {
		reader := crud.DB.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", where).
			Preload("BuildingUnit", "endlifespanversion IS NULL").
			Preload("Level", "endlifespanversion IS NULL").
			Preload("Baunit", "endlifespanversion IS NULL").
			Preload("Baunit.BaUnit", "endlifespanversion IS NULL").
			Preload("PlusBfs", "endlifespanversion IS NULL").
			Preload("PlusBfs.Bfs", "endlifespanversion IS NULL").
			Preload("PlusBfs.Bfs.Point", "endlifespanversion IS NULL").
			Preload("PlusBfs.Bfs.Point.Point", "endlifespanversion IS NULL").
			Preload("MinusBfs", "endlifespanversion IS NULL").
			Preload("MinusBfs.Bfs", "endlifespanversion IS NULL").
			Preload("MinusBfs.Bfs.Point", "endlifespanversion IS NULL").
			Preload("MinusBfs.Bfs.Point.Point", "endlifespanversion IS NULL").
			Preload("SpatialUnitGroups", "endlifespanversion IS NULL").
			Preload("SpatialUnitGroups.Whole", "endlifespanversion IS NULL").
			Preload("SuHierarchyParent", "endlifespanversion IS NULL").
			Preload("SuHierarchyParent.Parent", "endlifespanversion IS NULL").
			Preload("SuHierarchyChildren", "endlifespanversion IS NULL").
			Preload("SuHierarchyChildren.Child", "endlifespanversion IS NULL").
			Preload("RelationSu1", "endlifespanversion IS NULL").
			Preload("RelationSu1.Su1", "endlifespanversion IS NULL").
			Preload("RelationSu2", "endlifespanversion IS NULL").
			Preload("RelationSu2.Su2", "endlifespanversion IS NULL").
			First(&spatialUnit)
		if reader.RowsAffected == 0 {
			return nil, errors.New("Entity not found")
		}
		return spatialUnit, nil
	}
	return nil, nil
}

func (crud LASpatialUnitCRUD) Create(spatialUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	spatialUnit := spatialUnitIn.(ladm.LASpatialUnit)
	existing := 0
	reader := tx.Model(&ladm.LASpatialUnit{}).Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).
		Count(&existing)
	if reader.Error != nil {
		tx.Rollback()
		return nil, reader.Error
	}
	if existing != 0 {
		tx.Rollback()
		return nil, errors.New("Entity already exists")
	}
	currentTime := time.Now()
	spatialUnit.ID = spatialUnit.SuID.String()
	spatialUnit.BeginLifespanVersion = currentTime
	spatialUnit.EndLifespanVersion = nil
	writer := tx.Set("gorm:save_associations", false).Create(&spatialUnit)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	if building := spatialUnit.BuildingUnit; building != nil {
		building.ID = spatialUnit.ID
		building.BeginLifespanVersion = spatialUnit.BeginLifespanVersion
		writer = tx.Set("gorm:save_associations", false).Create(building)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return &spatialUnit, nil
}

func (crud LASpatialUnitCRUD) ReadAll(where ...interface{}) (interface{}, error) {
	var spatialUnits []ladm.LASpatialUnit
	if crud.DB.Where("endlifespanversion IS NULL").Find(&spatialUnits).RowsAffected == 0 {
		return nil, errors.New("Entity not found")
	}
	return &spatialUnits, nil
}

func (crud LASpatialUnitCRUD) Update(spatialUnitIn interface{}) (interface{}, error) {
	tx := crud.DB.Begin()
	spatialUnit := spatialUnitIn.(*ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSUnit ladm.LASpatialUnit
	reader := tx.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).
		Preload("BuildingUnit", "endlifespanversion IS NULL").
		Preload("Baunit", "endlifespanversion IS NULL").
		Preload("PlusBfs", "endlifespanversion IS NULL").
		Preload("MinusBfs", "endlifespanversion IS NULL").
		Preload("SpatialUnitGroups", "endlifespanversion IS NULL").
		Preload("SuHierarchyParent", "endlifespanversion IS NULL").
		Preload("SuHierarchyChildren", "endlifespanversion IS NULL").
		Preload("RelationSu1", "endlifespanversion IS NULL").
		Preload("RelationSu2", "endlifespanversion IS NULL").
		First(&oldSUnit)
	if reader.Error != nil {
		tx.Rollback()
		return nil, reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	oldSUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSUnit)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return nil, errors.New("Entity not found")
	}
	spatialUnit.ID = spatialUnit.SuID.String()
	spatialUnit.BeginLifespanVersion = currentTime
	spatialUnit.EndLifespanVersion = nil
	spatialUnit.LevelID = oldSUnit.LevelID
	spatialUnit.LevelBeginLifespanVersion = oldSUnit.LevelBeginLifespanVersion
	writer = tx.Set("gorm:save_associations", false).Create(&spatialUnit)
	if writer.Error != nil {
		tx.Rollback()
		return nil, writer.Error
	}
	if building := oldSUnit.BuildingUnit; building != nil {
		building.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(building)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		if newBuilding := spatialUnit.BuildingUnit; newBuilding != nil {
			newBuilding.ID = spatialUnit.ID
			newBuilding.BeginLifespanVersion = spatialUnit.BeginLifespanVersion
			newBuilding.EndLifespanVersion = spatialUnit.EndLifespanVersion
			writer = tx.Set("gorm:save_associations", false).Create(newBuilding)
			if writer.Error != nil {
				tx.Rollback()
				return nil, writer.Error
			}
		}
	}

	for _, baUnit := range oldSUnit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&baUnit)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		baUnit.BeginLifespanVersion = currentTime
		baUnit.EndLifespanVersion = nil
		baUnit.SUBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&baUnit)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, plusBfs := range oldSUnit.PlusBfs {
		plusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&plusBfs)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		plusBfs.BeginLifespanVersion = currentTime
		plusBfs.EndLifespanVersion = nil
		plusBfs.SuBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&plusBfs)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, minusBfs := range oldSUnit.MinusBfs {
		minusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&minusBfs)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		minusBfs.BeginLifespanVersion = currentTime
		minusBfs.EndLifespanVersion = nil
		minusBfs.SuBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&minusBfs)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, group := range oldSUnit.SpatialUnitGroups {
		group.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&group)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		group.BeginLifespanVersion = currentTime
		group.EndLifespanVersion = nil
		group.PartBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&group)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	if hierarchyParent := oldSUnit.SuHierarchyParent; hierarchyParent != nil {
		hierarchyParent.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(hierarchyParent)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		hierarchyParent.BeginLifespanVersion = currentTime
		hierarchyParent.EndLifespanVersion = nil
		hierarchyParent.ChildBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(hierarchyParent)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, child := range oldSUnit.SuHierarchyChildren {
		child.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&child)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		child.BeginLifespanVersion = currentTime
		child.EndLifespanVersion = nil
		child.ParentBeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&child)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, relation1 := range oldSUnit.RelationSu1 {
		relation1.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&relation1)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		relation1.BeginLifespanVersion = currentTime
		relation1.EndLifespanVersion = nil
		relation1.Su2BeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&relation1)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	for _, relation2 := range oldSUnit.RelationSu2 {
		relation2.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&relation2)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return nil, errors.New("Entity not found")
		}
		relation2.BeginLifespanVersion = currentTime
		relation2.EndLifespanVersion = nil
		relation2.Su1BeginLifespanVersion = currentTime
		writer = tx.Set("gorm:save_associations", false).Create(&relation2)
		if writer.Error != nil {
			tx.Rollback()
			return nil, writer.Error
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return nil, commit.Error
	}
	return spatialUnit, nil
}

func (crud LASpatialUnitCRUD) Delete(spatialUnitIn interface{}) error {
	tx := crud.DB.Begin()
	spatialUnit := spatialUnitIn.(ladm.LASpatialUnit)
	currentTime := time.Now()
	var oldSpatialUnit ladm.LASpatialUnit
	reader := tx.Where("suid = ?::\"Oid\" AND endlifespanversion IS NULL", spatialUnit.SuID).
		Preload("BuildingUnit", "endlifespanversion IS NULL").
		Preload("Baunit", "endlifespanversion IS NULL").
		Preload("PlusBfs", "endlifespanversion IS NULL").
		Preload("MinusBfs", "endlifespanversion IS NULL").
		Preload("SpatialUnitGroups", "endlifespanversion IS NULL").
		Preload("SuHierarchyParent", "endlifespanversion IS NULL").
		Preload("SuHierarchyChildren", "endlifespanversion IS NULL").
		Preload("RelationSu1", "endlifespanversion IS NULL").
		Preload("RelationSu2", "endlifespanversion IS NULL").
		First(&oldSpatialUnit)
	if reader.Error != nil {
		tx.Rollback()
		return reader.Error
	}
	if reader.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	oldSpatialUnit.EndLifespanVersion = &currentTime
	writer := tx.Set("gorm:save_associations", false).Save(&oldSpatialUnit)
	if writer.Error != nil {
		tx.Rollback()
		return writer.Error
	}
	if writer.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("Entity not found")
	}
	if building := oldSpatialUnit.BuildingUnit; building != nil {
		building.EndLifespanVersion = oldSpatialUnit.EndLifespanVersion
		writer = tx.Set("gorm:save_associations", false).Save(building)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, baUnit := range oldSpatialUnit.Baunit {
		baUnit.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&baUnit)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, plusBfs := range oldSpatialUnit.PlusBfs {
		plusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&plusBfs)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, minusBfs := range oldSpatialUnit.MinusBfs {
		minusBfs.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&minusBfs)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, group := range oldSpatialUnit.SpatialUnitGroups {
		group.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&group)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	if hierarchyParent := oldSpatialUnit.SuHierarchyParent; hierarchyParent != nil {
		hierarchyParent.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(hierarchyParent)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, child := range oldSpatialUnit.SuHierarchyChildren {
		child.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&child)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, relation1 := range oldSpatialUnit.RelationSu1 {
		relation1.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&relation1)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	for _, relation2 := range oldSpatialUnit.RelationSu2 {
		relation2.EndLifespanVersion = &currentTime
		writer = tx.Set("gorm:save_associations", false).Save(&relation2)
		if writer.Error != nil {
			tx.Rollback()
			return writer.Error
		}
		if writer.RowsAffected == 0 {
			tx.Rollback()
			return errors.New("Entity not found")
		}
	}
	commit := tx.Commit()
	if commit.Error != nil {
		return commit.Error
	}
	return nil
}
