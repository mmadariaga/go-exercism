package birdwatcher

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	result := 0
	itemNum := len(birdsPerDay)
	for i := 0; i < itemNum; i++ {
		result += birdsPerDay[i]
	}

	return result
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	result := 0

	startPosition := (week - 1) * 7
	endPosition := startPosition + 7
	maxPosition := minInt(endPosition, len(birdsPerDay))

	for i := startPosition; i < maxPosition; i++ {
		result += birdsPerDay[i]
	}

	return result
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	itemNum := len(birdsPerDay)
	for i := 0; i < itemNum; i++ {

		if i%2 != 0 {
			continue
		}

		birdsPerDay[i] += 1
	}

	return birdsPerDay
}
