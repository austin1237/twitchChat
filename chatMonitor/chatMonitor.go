package chatMonitor

import (
  "fmt"
  "time"
)

var currentChatCount int
var lastAverageCount int
var historyOfCounts []int

func resetChatCount(){
  time.Sleep(5 * time.Second)
  fmt.Println("Adding the following to history", currentChatCount)
  historyOfCounts = append(historyOfCounts, currentChatCount)
  checkAgainstLastAverage()
  currentChatCount = 0
  go resetChatCount()
}

func checkAgainstLastAverage(){
  if (lastAverageCount != 0 && currentChatCount != 0){
    increase := currentChatCount - lastAverageCount
    changePercentage := (float64(increase) / float64(lastAverageCount)) * 100
    fmt.Println("Percentage change was")
    fmt.Printf("%.6f\n", changePercentage)
  }
}

func calculateAveragePerMinute(){
  time.Sleep(60 * time.Second)
  sumOfCounts := 0
  for _, count := range historyOfCounts {
    sumOfCounts += count
  }
  if sumOfCounts != 0 {
    lastAverageCount = sumOfCounts / len(historyOfCounts)
  }else{
    lastAverageCount = 0
  }
  fmt.Println("Calculating lastAverageCount", lastAverageCount)
  go calculateAveragePerMinute()
}

func AddToCount() {
  currentChatCount++
}

func StartMonitoring(){
  go resetChatCount()
  go calculateAveragePerMinute()
}
