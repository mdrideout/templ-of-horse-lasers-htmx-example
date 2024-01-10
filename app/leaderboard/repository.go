package leaderboard

import (
	"main/app/leaderboard/models"
	"main/db"
	"sort"
)

// Get the list of items from the database
func getLeaderboardTableData() ([]models.LeaderboardItemStore, error) {
	items := []models.LeaderboardItemStore{}
	findResult := db.DB.Find(&items)
	if findResult.Error != nil {
		return nil, findResult.Error
	}

	// Sort the items
	sort.Slice(items, func(i, j int) bool {
		return items[i].Score > items[j].Score
	})

	return items, nil
}
