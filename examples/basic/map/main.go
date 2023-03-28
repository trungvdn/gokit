package main

import (
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

type Learner struct {
	lessonID string
	userID   string
}

type LessonLearners []*Learner

func main() {

	// lessonLearnerMap := make(map[string]int)
	// lessonLearnerMap["s"] = 1

	// for k, v := range lessonLearnerMap {
	// 	fmt.Println(k, v)
	// }
	// // member, ok := lessonLearnerMap["not-exits"]
	// // fmt.Printf("type %T\n", member)
	// // fmt.Println(member, ok)
	// var myMap map[string]int
	// fmt.Println(myMap)
	// for s := range myMap {
	// 	fmt.Println("start")
	// 	fmt.Println(s)
	// }

	// now := time.Now()
	// var delete *time.Time
	// delete = &now
	// var t pgtype.Timestamptz
	// t.Set(delete)
	// fmt.Println(t)
	// fmt.Println(timestamppb.New(*delete))
	fmt.Println(validateAndParseDate("2022-02-22", "1998-02-23"))
	var spotlightedUser pgtype.Text
	fmt.Println(spotlightedUser.Status, spotlightedUser.String)
	fmt.Println(pgtype.Null)

}

const layout string = "2006-01-02"

func validateAndParseDate(fromDate, toDate string) (err error) {
	if fromDate == "" || toDate == "" {
		return fmt.Errorf("from_date and to_date is required")
	}
	start, err := time.Parse(layout, fromDate)
	if err != nil {
		return err
	}
	end, err := time.Parse(layout, toDate)
	if err != nil {
		return err
	}

	if end.Sub(start).Hours() < 0 {
		return fmt.Errorf("Start date must come before End date")
	}

	return nil
}
