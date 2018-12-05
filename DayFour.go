package adventofcode2018

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type logEntry struct {
	year    int
	month   int
	day     int
	hour    int
	minute  int
	logText string
}

type logEntryList []logEntry

func (logList logEntryList) Len() int {
	return len(logList)
}

func (logList logEntryList) Swap(i, j int) {
	logList[i], logList[j] = logList[j], logList[i]
}

func (logList logEntryList) Less(i, j int) bool {

	logEntry1 := logList[i]
	logEntry2 := logList[j]

	if logEntry1.year == logEntry2.year {

		if logEntry1.month == logEntry2.month {
			if logEntry1.day == logEntry2.day {
				if logEntry1.hour == logEntry2.hour {
					return logEntry1.minute < logEntry2.minute
				} else {
					return logEntry1.hour < logEntry2.hour
				}
			} else {
				return logEntry1.day < logEntry2.day
			}
		} else {
			return logEntry1.month < logEntry2.month
		}
	} else {
		return logEntry1.year < logEntry2.year
	}
}

func DayFourPartOne() {
	fmt.Println("Day Four - Part One")

	input := ReadFile("day4-input.txt")
	inputList := strings.Split(input, "\n")

	logList := make([]logEntry, len(inputList))

	for i := 0; i < len(inputList); i++ {
		logList[i] = parseLogEntry(inputList[i])
	}

	sort.Sort(logEntryList(logList))

	sleepyGuard := findGuardWhoSleepsMost(logList)
	sleepyMinute := findMostMinuteSlept(sleepyGuard, logList)

	fmt.Println("Sleepy Guard: ", sleepyGuard)
	fmt.Println("Sleepy Minute: ", sleepyMinute)
}

func parseLogEntry(rawLog string) logEntry {
	date := rawLog[1:11]
	year, _ := strconv.Atoi(date[0:4])
	month, _ := strconv.Atoi(date[5:7])
	day, _ := strconv.Atoi(date[8:10])

	time := rawLog[12:17]
	hour, _ := strconv.Atoi(time[0:2])
	minute, _ := strconv.Atoi(time[3:])

	logText := rawLog[19:]

	logEntry := logEntry{}
	logEntry.year = year
	logEntry.month = month
	logEntry.day = day
	logEntry.hour = hour
	logEntry.minute = minute
	logEntry.logText = logText

	return logEntry
}

func findGuardWhoSleepsMost(logList []logEntry) string {

	currentGuard := ""
	sleepStartHour := 0
	sleepStartMinute := 0

	guardSleepTotals := make(map[string]int)

	for i := 0; i < len(logList); i++ {

		logEntry := logList[i]

		if strings.HasPrefix(logEntry.logText, "Guard") {
			currentGuard = logEntry.logText
		} else if strings.HasPrefix(logEntry.logText, "falls") {
			sleepStartHour = logEntry.hour
			sleepStartMinute = logEntry.minute
		} else {
			if logEntry.hour == sleepStartHour {
				sleepMinutes := logEntry.minute - sleepStartMinute
				guardSleepTotals[currentGuard] = guardSleepTotals[currentGuard] + sleepMinutes
			} else {
				//hours are only increasing
				sleepMinutes := logEntry.minute
				//TODO missing hour minutes
				sleepMinutes = sleepMinutes + (60 - sleepStartMinute)
				guardSleepTotals[currentGuard] = guardSleepTotals[currentGuard] + sleepMinutes
			}
		}
	}

	maxSleep := 0
	ourGuy := ""

	fmt.Println(guardSleepTotals)
	for key, value := range guardSleepTotals {

		if value > maxSleep {
			ourGuy = key
			maxSleep = value
		}
	}

	return ourGuy
}

func findMostMinuteSlept(guard string, logList []logEntry) int {

	sleepMinutes := make([]int, 60)

	for i := 0; i < len(logList); i++ {
		logEntry := logList[i]

		if logEntry.logText == guard {
			fmt.Println("Found our guard")
			fmt.Println(logList[i+1].minute)
			fmt.Println(logList[i+2].minute)
			for j := logList[i+1].minute; j < logList[i+2].minute; j++ {
				fmt.Println("Sleep on minute: ", j)
				sleepMinutes[j] = sleepMinutes[j] + 1
			}
		}
	}

	fmt.Println(sleepMinutes)
	maxMinute := 0
	maxSleep := 0
	for i := 0; i < len(sleepMinutes); i++ {
		if sleepMinutes[i] >= maxSleep {
			maxMinute = i
			maxSleep = sleepMinutes[i]
		}
	}

	return maxMinute
}
