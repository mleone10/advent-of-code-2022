package newday

import (
	_ "embed"
)

//go:embed day.go.template
var dayTemplate string

//go:embed day_test.go.template
var dayTestTemplate string

// TODO: Implement tests for helper client
// TODO: Order should be 1) create directory, 2) create files, 3) write files
func InitializeDay(dayNum int) error {
	// err := os.Mkdir(fmt.Sprintf("src/day%02d", dayNum), 0755)
	// if err != nil {
	// 	log.Fatalf("failed to create new day directory: %v", err)
	// }

	// t := template.Must(template.New("day").Parse(dayTemplate))
	// t.Execute(os.Stdout, struct {
	// 	DayNum int
	// }{
	// 	DayNum: dayNum,
	// })
	return nil
}
