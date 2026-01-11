package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Printer interface {
	PrintTable(headers []string, rows [][]string) error
	PrintJSON(v any) error
}

type StdPrinter struct {
	out io.Writer
}

func NewStdPrinter(out io.Writer) *StdPrinter {
	if out == nil {
		out = os.Stdout
	}
	return &StdPrinter{out: out}
}

func (p *StdPrinter) PrintTable(headers []string, rows [][]string) error {
	cols := len(headers)
	for _, row := range rows {
		if len(row) > cols {
			cols = len(row)
		}
	}
	if cols == 0 {
		return nil
	}

	widths := make([]int, cols)
	for i := 0; i < cols; i++ {
		if i < len(headers) && len(headers[i]) > widths[i] {
			widths[i] = len(headers[i])
		}
	}
	for _, row := range rows {
		for i := 0; i < cols; i++ {
			cell := ""
			if i < len(row) {
				cell = row[i]
			}
			if len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	if len(headers) > 0 {
		for i := 0; i < cols; i++ {
			if i > 0 {
				if _, err := fmt.Fprint(p.out, "  "); err != nil {
					return err
				}
			}
			header := ""
			if i < len(headers) {
				header = headers[i]
			}
			if _, err := fmt.Fprintf(p.out, "%-*s", widths[i], header); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(p.out); err != nil {
			return err
		}
		for i := 0; i < cols; i++ {
			if i > 0 {
				if _, err := fmt.Fprint(p.out, "  "); err != nil {
					return err
				}
			}
			if _, err := fmt.Fprint(p.out, strings.Repeat("-", widths[i])); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(p.out); err != nil {
			return err
		}
	}

	for _, row := range rows {
		for i := 0; i < cols; i++ {
			if i > 0 {
				if _, err := fmt.Fprint(p.out, "  "); err != nil {
					return err
				}
			}
			cell := ""
			if i < len(row) {
				cell = row[i]
			}
			if _, err := fmt.Fprintf(p.out, "%-*s", widths[i], cell); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprintln(p.out); err != nil {
			return err
		}
	}
	return nil
}

func (p *StdPrinter) PrintJSON(v any) error {
	enc := json.NewEncoder(p.out)
	enc.SetIndent("", "  ")
	return enc.Encode(v)
}
