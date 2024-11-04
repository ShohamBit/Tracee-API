package formatter

import (
	pb "github.com/aquasecurity/tracee/api/v1beta1"
)

// PrintJSON prints an event in JSON format
func (f *Formatter) PrintJSON(event *pb.Event) {
	//TODO: add more output formats
	f.CMD.Printf("%s\n", event.String())
}
