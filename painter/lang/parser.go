package lang

import (
	"io"
	"bufio"
	"errors"
	"strconv"
	"strings"
	"github.com/vladimirvikulin/Lab3-Software-Architecture/painter"
)

// Parser уміє прочитати дані з вхідного io.Reader та повернути список операцій представлені вхідним скриптом.
type Parser struct {
	Figures []*painter.FigureOp
	BackgroundColor painter.OperationFunc
	BackgroundRect painter.OperationFunc
}

func (p *Parser) Parse(in io.Reader) ([]painter.Operation, error) {
	scanner := bufio.NewScanner(in)
	var res []painter.Operation
	for scanner.Scan() {
		commandLine := scanner.Text()
		commands := strings.Split(commandLine, ",")
		if len(commands) == 0 {
			continue
		}
		for _, val := range commands {
			part := strings.Fields(val)
			
			switch part[0] {
			case "white":
				p.BackgroundColor = painter.OperationFunc(painter.WhiteFill)
			
			case "green":
				p.BackgroundColor = painter.OperationFunc(painter.GreenFill)

			case "update":
				if p.BackgroundColor != nil {
					res = append(res, p.BackgroundColor)
				}
				if p.BackgroundRect != nil {
					res = append(res, p.BackgroundRect)
				}
				for _, figureInstance := range p.Figures {
					res = append(res, figureInstance.Figure())
				}
				res = append(res, painter.UpdateOp)

			case "bgrect":
				if len(part) != 5 {
					return nil, errors.New("incorrect backgroundrect number of arguments")
				}

				values := make([]float64, 4)
				for i := 1; i <= 4; i++ {
					value, err := strconv.ParseFloat(part[i], 64)
					if err != nil {
					return nil, errors.New("incorrect backgroundrect arguments")
					}
					values[i-1] = value
				}

				x1, y1, x2, y2 := values[0], values[1], values[2], values[3]

				x1Int := int(x1 * 800)
				y1Int := int(y1 * 800)
				x2Int := int(x2 * 800)
				y2Int := int(y2 * 800)

				p.BackgroundRect = painter.BackgroundRect(x1Int, y1Int, x2Int, y2Int)
			case "figure":
				if len(part) != 3 {
					return nil, errors.New("incorrect figure number of arguments")
				}

				values := make([]float64, 2)
				for i := 1; i <= 2; i++ {
					value, err := strconv.ParseFloat(part[i], 64)
					if err != nil {
						return nil, errors.New("incorrect figure arguments")
					}
					values[i-1] = value
				}

				x, y := values[0], values[1]

				xInt := int(x * 800)
				yInt := int(y * 800)

				p.Figures = append(p.Figures, &painter.FigureOp {
					F: painter.Figure{X: xInt, Y: yInt}})
			case "move":
				if len(part) != 3 {
					return nil, errors.New("incorrect move number of arguments")
				}
				values := make([]float64, 2)
				for i := 1; i <= 2; i++ {
					value, err := strconv.ParseFloat(part[i], 64)
					if err != nil {
					return nil, errors.New("incorrect move arguments")
					}
					values[i-1] = value
				}
				x, y := values[0], values[1]
				res = append(res, painter.Move(int(x * 800), int(y * 800), p.Figures))

			case "reset":
				res = append(res, painter.Reset())
			}
		}
	}
	return res, nil
}
