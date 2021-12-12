package day06

import (
	"strconv"
	"strings"
)

type Solution struct {}

func (s Solution) Part1(lines []string) (string, error) {
	input := strings.Split(lines[0], ",")
	nums := make([]int, len(input))
	for i, s := range input {
		nums[i], _ = strconv.Atoi(s)
	}
	cohort := newCohorts(nums)
	for i := 0; i < 80; i++ {
		cohort.update()
	}
	return strconv.Itoa(cohort.totalFish()), nil
}

func (s Solution) Part2(lines []string) (string, error) {
	input := strings.Split(lines[0], ",")
	nums := make([]int, len(input))
	for i, s := range input {
		nums[i], _ = strconv.Atoi(s)
	}
	cohort := newCohorts(nums)
	for i := 0; i < 256; i++ {
		cohort.update()
	}
	return strconv.Itoa(cohort.totalFish()), nil
}

const (
	numDays = 9
	daysToNextFishAfterFirstBirth = 6
)

// We will use the following approach:
// We will only count the number of fish within a given cohort. A cohort is a group of fish
// that will have children at the same time.


type cohorts struct {
	// The number of days
	numTicks int
	// counts per cohort. The i-th index is the days left for that cohort
	cohortCounts [numDays]int
}

func newCohorts(fishDueDates []int) cohorts {
	countPerDueDate := map[int]int{}
	for _, dueDate := range fishDueDates {
		val, _ := countPerDueDate[dueDate]
		countPerDueDate[dueDate] = val + 1
	}
	c := cohorts{cohortCounts: [numDays]int{}}
	for date, count := range countPerDueDate {
		c.cohortCounts[date] = count
	}
	return c
}

func (c *cohorts) update() {
	c.numTicks++
	newCohorts := [numDays]int{}
	for day := numDays - 1; day >= 0; day-- {
		fishForDay := c.cohortCounts[day]
		if day == 0 {
			// The current fishes on day zero will now be moved to the cohort that gives birth in 6 days
			newCohorts[daysToNextFishAfterFirstBirth] += fishForDay
			// The newly spawned fish will now live in the cohort that gives birth in 8 days
			newCohorts[numDays - 1] = fishForDay
		} else {
			// shuffle the current generation to the next day.
			newCohorts[day - 1] = c.cohortCounts[day]
		}
	}
	c.cohortCounts = newCohorts
}

func (c *cohorts) totalFish() int{
	total := 0
	for _, v := range c.cohortCounts {
		total += v
	}
	return total
}

