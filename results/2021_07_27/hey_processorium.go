package main



import (
  "fmt"
  "io/ioutil"
  "os"
  "strings"
  "strconv"
)


var fp = fmt.Fprintf



func get_lines ( target string, lines []string ) ( []string ) {
  
  results := make([]string, 0)
  for i := 0; i < len(lines); i++ {
    if strings.Contains ( lines[i], target ) {
      results = append ( results, lines[i] )
    }
  }
  return results
}




func get_values ( target string, lines []string, field int ) ( float64, float64 ) {

  var average_1, average_2 float64

  //if target == "99" {
    //target += string('%')
  //}

  result := get_lines ( target, lines )
  n_lines_with_average := len(result)

  if n_lines_with_average != 10 {
    fp ( os.Stdout, "%d lines with target |%s|\n", n_lines_with_average, target)
    os.Exit ( 0 )
  }
//fmt.Println ( result )

  // The first 5 are the values for 1-sender.
  total := 0.0
  for i := 0; i < 5; i ++ {
    words := strings.Fields ( result[i] )
    val, err := strconv.ParseFloat ( words[field], 64 )
    if err != nil {
      fp ( os.Stdout, "error converting |%s|\n", words[i] )
      os.Exit ( 1 )
    }
    total += val
  }
  average_1 = total / 5


  // The second 5 are the values for 10-senders.
  total = 0.0
  for i := 5; i < 10; i ++ {
    words := strings.Fields ( result[i] )
    val, err := strconv.ParseFloat ( words[field], 64 )
    if err != nil {
      fp ( os.Stdout, "error converting |%s|\n", words[i] )
      os.Exit ( 1 )
    }
    total += val
  }
  average_2 = total / 5
  
  return average_1, average_2
}




func main() {
  filename := "./raw"

  content, err := ioutil.ReadFile ( filename )
  if err != nil {
    fmt.Println ( "dang." );
    os.Exit ( 1 );
  }
  lines := strings.Split ( string(content), "\n" )

  v1, v2 := get_values ( "Average", lines, 1 )
  fp ( os.Stdout, "Average: %f %f\n", v1, v2 )

  v1, v2 = get_values ( "99%", lines, 2 )
  fp ( os.Stdout, "99%%: %f %f\n", v1, v2 )

  v1, v2 = get_values ( "Slowest", lines, 1 )
  fp ( os.Stdout, "100%%: %f %f\n", v1, v2 )

  v1, v2 = get_values ( "Requests", lines, 1 )
  fp ( os.Stdout, "Requests: %f %f\n", v1, v2 )


}




