package db

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const dateLayout = "20060102"

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("empty repeat rule")
	}

	startDate, err := time.Parse(dateLayout, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Split(repeat, " ")
	rule := parts[0]

	switch rule {
	case "d":
		if len(parts) < 2 {
			return "", errors.New("missing interval for rule 'd'")
		}
		days, err := strconv.Atoi(parts[1])
		if err != nil || days < 1 || days > 400 {
			return "", errors.New("invalid interval for rule 'd'")
		}
		return nextDateDaily(startDate, now, days), nil

	case "y":
		return nextDateYearly(startDate, now), nil

	case "w":
		if len(parts) < 2 {
			return "", errors.New("missing days for rule 'w'")
		}
		weekdays, err := parseWeekdays(parts[1])
		if err != nil {
			return "", err
		}
		return nextDateWeekly(startDate, now, weekdays), nil

	case "m":
		if len(parts) < 2 {
			return "", errors.New("missing days for rule 'm'")
		}
		days, err := parseDays(parts[1])
		if err != nil {
			return "", err
		}
		var months []int
		if len(parts) >= 3 {
			months, err = parseMonths(parts[2])
			if err != nil {
				return "", err
			}
		}
		return nextDateMonthly(startDate, now, days, months)

	default:
		return "", errors.New("unsupported repeat rule")
	}
}

func truncateToDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func nextDateDaily(startDate, now time.Time, interval int) string {
	now = truncateToDay(now)
	date := truncateToDay(startDate)

	if date.After(now) {
		return date.AddDate(0, 0, interval).Format(dateLayout)
	}

	for !date.After(now) {
		date = date.AddDate(0, 0, interval)
	}

	return date.Format(dateLayout)
}

func nextDateYearly(startDate, now time.Time) string {
	now = truncateToDay(now)
	date := truncateToDay(startDate)

	if date.After(now) {
		return date.AddDate(1, 0, 0).Format(dateLayout)
	}

	for !date.After(now) {
		date = date.AddDate(1, 0, 0)
	}

	return date.Format(dateLayout)
}

func nextDateWeekly(startDate, now time.Time, weekdays []int) string {
	now = truncateToDay(now)
	date := now.AddDate(0, 0, 1)

	for {
		wd := int(date.Weekday())
		if wd == 0 {
			wd = 7
		}
		for _, target := range weekdays {
			if wd == target {
				return date.Format(dateLayout)
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func nextDateMonthly(startDate, now time.Time, days, months []int) (string, error) {
	now = truncateToDay(now)
	date := now.AddDate(0, 0, 1)

	for {
		month := int(date.Month())
		if len(months) == 0 || contains(months, month) {
			day := date.Day()
			lastDay := lastDayOfMonth(date)

			for _, targetDay := range days {
				actualDay := targetDay
				if targetDay < 0 {
					actualDay = lastDay + targetDay + 1
				}
				if actualDay == day {
					return date.Format(dateLayout), nil
				}
			}
		}
		date = date.AddDate(0, 0, 1)
	}
}

func parseWeekdays(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	weekdays := make([]int, 0, len(parts))

	for _, p := range parts {
		day, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil || day < 1 || day > 7 {
			return nil, errors.New("invalid weekday")
		}
		weekdays = append(weekdays, day)
	}
	return weekdays, nil
}

func parseDays(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	days := make([]int, 0, len(parts))

	for _, p := range parts {
		day, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil || (day < -2 || day == 0 || day > 31) {
			return nil, errors.New("invalid day of month")
		}
		days = append(days, day)
	}
	return days, nil
}

func parseMonths(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	months := make([]int, 0, len(parts))

	for _, p := range parts {
		month, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil || month < 1 || month > 12 {
			return nil, errors.New("invalid month")
		}
		months = append(months, month)
	}
	return months, nil
}

func lastDayOfMonth(date time.Time) int {
	nextMonth := date.AddDate(0, 1, -date.Day()+1)
	lastDay := nextMonth.AddDate(0, 0, -1)
	return lastDay.Day()
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
