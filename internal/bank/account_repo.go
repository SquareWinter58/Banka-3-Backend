package bank

import "errors"

func (s *Server) AccountNameExistsForOwner(ownerID int64, name string, excludeAccountNumber string) (bool, error) {
	var count int64

	err := s.db_gorm.
		Model(&Account{}).
		Where("owner = ? AND name = ? AND number <> ?", ownerID, name, excludeAccountNumber).
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *Server) UpdateAccountNameRecord(accountNumber string, name string) error {
	result := s.db_gorm.
		Model(&Account{}).
		Where("number = ?", accountNumber).
		Update("name", name)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found")
	}

	return nil
}

func (s *Server) UpdateAccountLimitsRecord(accountNumber string, dailyLimit *int64, monthlyLimit *int64) error {
	updates := map[string]any{}

	if dailyLimit != nil {
		updates["daily_limit"] = *dailyLimit
	}
	if monthlyLimit != nil {
		updates["monthly_limit"] = *monthlyLimit
	}

	if len(updates) == 0 {
		return errors.New("no limits provided")
	}

	result := s.db_gorm.
		Model(&Account{}).
		Where("number = ?", accountNumber).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("account not found")
	}

	return nil
}
