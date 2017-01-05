package chatMonitor

import (
  "fmt"
  "time"
)

var currentChatCount int
var lastAverageCount int
var historyOfCounts []int
var historyOfAverages []int
var overallAverage int
var startTimeStamp time.Time
var endTimeStamp  time.Time


func resetChatCount(){
  time.Sleep(5 * time.Second)
  historyOfCounts = append(historyOfCounts, currentChatCount)
  checkAgainstAverage()
  currentChatCount = 0
  go resetChatCount()
}

func checkAgainstAverage(){
  if (overallAverage != 0 && currentChatCount != 0){
    increase := currentChatCount - overallAverage
    changePercentage := (float64(increase) / float64(overallAverage)) * 100
    fmt.Println("Percentage change was")
    fmt.Printf("%.6f\n", changePercentage)
    isChatHyped(changePercentage, 50)
  }
}

func isChatHyped(changePercentage float64, hypePercentage float64){
  if changePercentage > hypePercentage && time.Time(startTimeStamp).IsZero(){
    fmt.Println("Hype has started")
    startTimeStamp = time.Now()
  }else if changePercentage < hypePercentage && !time.Time(startTimeStamp).IsZero(){
    now := time.Now()
    duration := now.Sub(startTimeStamp)
    endTimeStamp = now
    fmt.Println("Hype has ended", startTimeStamp, endTimeStamp)
    fmt.Println("For a duration of", duration.Seconds())
    startTimeStamp = time.Time{}
    endTimeStamp = time.Time{}
  }
}

func calculateAverage (numbers []int) int{
  sum := 0
  average := 0
  for _, number := range numbers {
    sum += number
  }
  if sum != 0 {
    average = sum / len(numbers)
  }
  return average
}

func calculateAveragePerMinute(){
  time.Sleep(60 * time.Second)
  lastAverage := calculateAverage(historyOfCounts)
  historyOfCounts = nil
  historyOfAverages = append(historyOfAverages, lastAverage)
  overallAverage = calculateAverage(historyOfAverages)
  fmt.Println(historyOfAverages)
  fmt.Println("calculating overallAverage", overallAverage)
  go calculateAveragePerMinute()
}

func AddToCount() {
  currentChatCount++
}

func StartMonitoring(){
  go resetChatCount()
  go calculateAveragePerMinute()
}
