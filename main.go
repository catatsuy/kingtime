package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"time"
)

type Attendance struct {
	WorkDay  string
	ClockIn  string
	ClockOut string
}

func main() {
	attendances := make([]Attendance, 0, 10)
	re := regexp.MustCompile(`kintai: ([0-9]{2}):([0-9]{2}) *[-]? *([0-9]{2}):([0-9]{2})`)
	s := bufio.NewScanner(os.Stdin)
	aTmp := Attendance{}
	for s.Scan() {
		text := s.Text()
		if len(text) < 3 {
			continue
		}

		tt, err := time.Parse("Jan 02", text[:len(text)-2])
		if err == nil {
			aTmp.WorkDay = time.Now().Format("2006") + "/" + tt.Format("01/02")
			continue
		}

		tt, err = time.Parse("Jan 2", text[:len(text)-2])
		if err == nil {
			aTmp.WorkDay = time.Now().Format("2006") + "/" + tt.Format("01/02")
			continue
		}

		sp := re.FindStringSubmatch(text)
		if len(sp) > 3 {
			aTmp.ClockIn = sp[1] + ":" + sp[2]
			aTmp.ClockOut = sp[3] + ":" + sp[4]
		}

		if aTmp.WorkDay != "" && aTmp.ClockIn != "" && aTmp.ClockOut != "" {
			attendances = append(attendances, aTmp)

			aTmp = Attendance{}
		}
	}

	for i := len(attendances) - 1; i >= 0; i-- {
		a := attendances[i]
		s := fmt.Sprintf("%s\t%s\t%s", a.WorkDay, a.ClockIn, a.ClockOut)
		fmt.Println(s)
	}
}
