package main

import (
  "bufio"
  "fmt"
  "os"
  "sort"
  "strconv"
  "strings"
)

func main() {
  file_path := "input.txt"

  part1(file_path)
  part2(file_path)
}

func part1(file_path string) {
  file, scanner := initScanner(file_path)
  if scanner == nil {
    return
  }

  defer file.Close()

  var column1, column2 []int

  for scanner.Scan() {
    line := scanner.Text()
    fields := strings.Fields(line)

    value1, err1 := strconv.Atoi(fields[0])
    if err1 != nil {
      fmt.Println("Conversion error:", err1)
      return
    }
    column1 = append(column1, value1)

    value2, err2 := strconv.Atoi(fields[1])
    if err2 != nil {
      fmt.Println("Conversion error:", err2)
      return
    }
    column2 = append(column2, value2)
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  sort.Ints(column1)
  sort.Ints(column2)

  total := 0

  for i := 0; i < len(column1); i++ {
    diff := 0

    if column1[i] > column2[i] {
      diff = column1[i] - column2[i]
    } else {
      diff = column2[i] - column1[i]
    }
    total += diff
  }

  fmt.Println(total)
}

func part2(file_path string) {
  file, scanner := initScanner(file_path)
  if scanner == nil {
    return
  }

  defer file.Close()

  var column1 []int
  value_counts_map := make(map[int]int)

  for scanner.Scan() {
    line := scanner.Text()
    fields := strings.Fields(line)

    value1, err1 := strconv.Atoi(fields[0])
    if err1 != nil {
      fmt.Println("Conversion error:", err1)
      return
    }
    column1 = append(column1, value1)

    value2, err2 := strconv.Atoi(fields[1])
    if err2 != nil {
      fmt.Println("Conversion error:", err2)
      return
    }

    value_counts_map[value2]++
  }

  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  total := 0

  for _, column1_value := range column1 {
    column2_count := value_counts_map[column1_value]
    total += (column1_value * column2_count)
  }

  fmt.Println(total)
}

func initScanner(file_path string) (*os.File, *bufio.Scanner) {
  file, err := os.Open(file_path)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return nil, nil
  }

  return file, bufio.NewScanner(file)
}
