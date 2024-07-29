package gopkg

import (
	original_fmt "fmt"

	"github.com/issueye/build_magic/backend/code_engine/stand"
)

var (
	fmtModule = stand.NewGojaModule("go/fmt")
)

func InitFmt() {
	fmtModule.Set(
		stand.Objects{
			// Functions
			"append":       original_fmt.Append,
			"appendf":      original_fmt.Appendf,
			"appendln":     original_fmt.Appendln,
			"errorf":       original_fmt.Errorf,
			"formatString": original_fmt.FormatString,
			"fprint":       original_fmt.Fprint,
			"fprintf":      original_fmt.Fprintf,
			"fprintln":     original_fmt.Fprintln,
			"fscan":        original_fmt.Fscan,
			"fscanf":       original_fmt.Fscanf,
			"fscanln":      original_fmt.Fscanln,
			"print":        original_fmt.Print,
			"printf":       original_fmt.Printf,
			"println":      original_fmt.Println,
			"scan":         original_fmt.Scan,
			"scanf":        original_fmt.Scanf,
			"scanln":       original_fmt.Scanln,
			"sprint":       original_fmt.Sprint,
			"sprintf":      original_fmt.Sprintf,
			"sprintln":     original_fmt.Sprintln,
			"sscan":        original_fmt.Sscan,
			"sscanf":       original_fmt.Sscanf,
			"sscanln":      original_fmt.Sscanln,

			// Var and consts

			// Types (value type)

			// Types (pointer type)
		},
	).Register()
}

// func Enable(runtime *goja.Runtime) {
// 	fmtModule.Enable(runtime)
// }
