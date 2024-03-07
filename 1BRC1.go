package main 

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"time"
	"sort"
)

func main(){
    start := time.Now()

    file, err := os.Open("measurements_1b.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    results:=make(map[string]map[string]float64)

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
	    line:=(scanner.Text())
	    values:=strings.Split(line,";")
	    field:=values[0]
	    value, err:=strconv.ParseFloat(values[1], 64)
	    if err != nil {
		fmt.Println("Error:", err)
		return
	        }

	    if item, ok := results[field]; ok {
		    mini:=item["min"]
		    mean:=item["mean"]
		    maxi:=item["max"]
		    total:=item["total"]
		    if value > maxi {
			results[field]["max"] = value
		}
               	    if value < mini {
                       results[field]["min"] = value
                }
                    results[field]["mean"] = mean + value
                    results[field]["total"] = total + 1
	        } else{
		    results[field]=map[string]float64{
		    "max":value,
		    "mean":value,
		    "min":value,
 		    "total":1,
		    }
	        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    final := []string{}
    for key, item := range results {
        mini := strconv.FormatFloat(item["min"], 'f', 1, 64)
        mean := item["mean"] / item["total"]
        maxi := strconv.FormatFloat(item["max"], 'f', 1, 64)
        final = append(final, fmt.Sprintf("%s=%s/%.1f/%s", key, mini, mean, maxi))
    }

    sort.Strings(final)

//    fmt.Println(final)
    fmt.Printf("--- %.6f seconds ---\n", time.Since(start).Seconds())
}
