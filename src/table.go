package main

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

func makeTableWriter(fl []float64) table.Writer {
	t := table.NewWriter()
	t.SetStyle(table.Style{
		Name: "tabstyle",
		Box: table.BoxStyle{
			BottomLeft:       "\\",
			BottomRight:      "/",
			BottomSeparator:  "v",
			Left:             "[",
			LeftSeparator:    "{",
			MiddleHorizontal: "-",
			MiddleSeparator:  "+",
			MiddleVertical:   "|",
			PaddingLeft:      " ",
			PaddingRight:     " ",
			Right:            " ]",
			RightSeparator:   "}",
			TopLeft:          "(",
			TopRight:         ")",
			TopSeparator:     "^",
			UnfinishedRow:    " ~~~",
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateFooter:  true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Date", "SunAzimuth",
		fmt.Sprintf("%.2f", fl[0]),
		fmt.Sprintf("%.2f", fl[1]),
		fmt.Sprintf("%.2f", fl[2]),
		fmt.Sprintf("%.2f", fl[3]),
		fmt.Sprintf("%.2f", fl[4]),
	})
	return t
}
