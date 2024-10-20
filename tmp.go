package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"
// 	"text/template"
// 	"time"
// )

// type Event struct {
// 	Timestamp   int64       `json:"timestamp"`
// 	UID         int         `json:"uid"`
// 	ProcessName string      `json:"comm"`
// 	ProcessID   int         `json:"pid"`
// 	ThreadID    int         `json:"tid"`
// 	ReturnValue int         `json:"ret"`
// 	EventName   string      `json:"event"`
// 	Args        []ArgStruct `json:"args"`
// }

// type ArgStruct struct {
// 	Name  string      `json:"name"`
// 	Value interface{} `json:"value"`
// }

// // PrintTable prints the event in table format
// func PrintTable(event Event) {
// 	ut := time.Unix(0, event.Timestamp)
// 	timestamp := fmt.Sprintf("%02d:%02d:%02d:%06d", ut.Hour(), ut.Minute(), ut.Second(), ut.Nanosecond()/1000)
// 	fmt.Printf("%-16s %-6d %-16s %-7d %-7d %-16d %-25s ", timestamp, event.UID, event.ProcessName, event.ProcessID, event.ThreadID, event.ReturnValue, event.EventName)
// 	for i, arg := range event.Args {
// 		if i == 0 {
// 			fmt.Printf("%s: %v", arg.Name, arg.Value)
// 		} else {
// 			fmt.Printf(", %s: %v", arg.Name, arg.Value)
// 		}
// 	}
// 	fmt.Println()
// }

// // PrintJSON prints the event in JSON format
// func PrintJSON(event Event) {
// 	eBytes, err := json.Marshal(event)
// 	if err != nil {
// 		fmt.Println("Error marshaling event to JSON:", err)
// 		return
// 	}
// 	fmt.Println(string(eBytes))
// }

// // PrintTemplate prints the event using a provided Go template
// func PrintTemplate(event Event, tmplPath string) {
// 	tmpl, err := template.ParseFiles(tmplPath)
// 	if err != nil {
// 		fmt.Println("Error parsing template:", err)
// 		return
// 	}
// 	err = tmpl.Execute(os.Stdout, event)
// 	if err != nil {
// 		fmt.Println("Error executing template:", err)
// 	}
// }

// func main() {
// 	// Example event data
// 	event := Event{
// 		Timestamp:   time.Now().UnixNano(),
// 		UID:         1000,
// 		ProcessName: "language_server",
// 		ProcessID:   4377,
// 		ThreadID:    4386,
// 		ReturnValue: 0,
// 		EventName:   "security_file_open",
// 		Args: []ArgStruct{
// 			{Name: "pathname", Value: "/home/shoham/TraceeClient/testTraceeClient/dam1"},
// 			{Name: "flags", Value: "O_RDONLY|O_LARGEFILE"},
// 			{Name: "dev", Value: 265289728},
// 			{Name: "inode", Value: 4054},
// 		},
// 	}

// 	// Example usage
// 	fmt.Println("Table Format:")
// 	PrintTable(event)

// 	fmt.Println("\nJSON Format:")
// 	PrintJSON(event)

// 	fmt.Println("\nTemplate Format:")
// 	PrintTemplate(event, "event_template.txt") // assuming a template file exists
// }