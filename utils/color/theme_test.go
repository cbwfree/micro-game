package color

import (
	"fmt"
	"testing"
)

func TestTheme(t *testing.T) {
	fmt.Println(Trace.Text("Trace"))
	fmt.Println(Debug.Text("Debug"))
	fmt.Println(Info.Text("Info"))
	fmt.Println(Warn.Text("Warn"))
	fmt.Println(Error.Text("Error"))
	fmt.Println(Success.Text("Success"))
	fmt.Println(Question.Text("Question"))
	fmt.Println(Fatal.Text("Fatal"))

	fmt.Println("------------------")

	fmt.Println(Note.Text("Note"))
	fmt.Println(Primary.Text("Primary"))
	fmt.Println(Comment.Text("Comment"))
}
