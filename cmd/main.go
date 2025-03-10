package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	calc "SF-HW-34.6.1/pkg/calculator"
)

var (
	expRe = regexp.MustCompile(`^\s*(-?[0-9]+(?:\.[0-9]+)?)\s*([+\-*/])\s*(-?[0-9]+(?:\.[0-9]+)?)\s*=\s*\?\s*$`)
)

var (
	ErrEmptyPath           = fmt.Errorf("empty path specified")
	ErrCantCreateOutFile   = fmt.Errorf("can't create output file")
	ErrCantParseExpression = fmt.Errorf("error while parsing math expression")
	ErrNoExpressionFound   = fmt.Errorf("no expression found")
)

func main() {
	var inputPath string
	var outputPath string
	flag.StringVar(&inputPath, "in", "", "Specify the file with data to be processed")
	flag.StringVar(&outputPath, "out", "out.txt", "Specify the path to the file to write the result to")
	flag.Parse()

	if len(inputPath) == 0 {
		log.Fatal(fmt.Errorf("input path can't be empty: %w", ErrEmptyPath))
	}

	inFile, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()

	outFile, err := os.Create(outputPath)
	if err != nil {
		log.Fatal(ErrCantCreateOutFile)
	}
	defer outFile.Close()

	inReader := bufio.NewReader(inFile)
	outWriter := bufio.NewWriter(outFile)
	for {
		line, err := inReader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		line = strings.TrimSpace(line)
		exp, err := parseExpression(expRe, line)
		if err != nil {
			log.Printf("%v: %s\n", err, line)
		} else {
			log.Printf("found expression: %v\n", exp)
			res, err := calc.Calc(*exp)
			if err != nil {
				log.Printf("evaluating %s resulted in an error: %v\n", exp, err)
			} else {
				line := fmt.Sprintf("%s=%g\n", exp, res)
				outWriter.WriteString(line)
			}
		}
	}
	outWriter.Flush()
}

// parseExpression parses a string into an Expression struct using the provided regexp.
// It returns an error if the string doesn't match the expected format or if parsing fails.
func parseExpression(re *regexp.Regexp, str string) (*calc.Expression, error) {
	match := re.FindAllStringSubmatch(str, -1)
	if len(match) > 0 && len(match[0]) == 4 {
		op1, err := strconv.ParseFloat(match[0][1], 64)
		if err != nil {
			return nil, fmt.Errorf("%w:, %w", ErrCantParseExpression, err)
		}
		op2, err := strconv.ParseFloat(match[0][3], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %w", ErrCantParseExpression, err)
		}
		operator := match[0][2]
		return calc.NewExpression(op1, op2, operator), nil
	}

	return nil, ErrNoExpressionFound
}
