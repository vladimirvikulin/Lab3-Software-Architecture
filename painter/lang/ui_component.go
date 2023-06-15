package lang

import (
	"errors"
	"github.com/vladimirvikulin/Lab3-Software-Architecture/painter"
	"strconv"
)

type UIState struct {
	Figures         []*painter.FigureOp
	BackgroundColor painter.OperationFunc
	BackgroundRect  painter.OperationFunc
	UpdateOp        painter.Operation
	MoveOp          []painter.Operation
}

func NewUIState() *UIState {
	return &UIState{}
}

func (s *UIState) GetOperations() []painter.Operation {
	var ops []painter.Operation

	if s.BackgroundColor != nil {
		ops = append(ops, s.BackgroundColor)
	}
	if s.BackgroundRect != nil {
		ops = append(ops, s.BackgroundRect)
	}
	if len(s.MoveOp) != 0 {
		ops = append(ops, s.MoveOp...)
		s.MoveOp = nil
	}
	if len(s.Figures) != 0 {
		for _, figure := range s.Figures {
			ops = append(ops, figure)
		}
	}
	if s.UpdateOp != nil {
		ops = append(ops, s.UpdateOp)
	}

	return ops
}

// WhiteBackground sets the background color to white.
func (s *UIState) WhiteBackground() {
	s.BackgroundColor = painter.OperationFunc(painter.WhiteFill)
}

// GreenBackground sets the background color to green.
func (s *UIState) GreenBackground() {
	s.BackgroundColor = painter.OperationFunc(painter.GreenFill)
}

// SetBackgroundRect устанавливает прямоугольник фона.
func (s *UIState) SetBackgroundRect(args []string) error {
	if len(args) != 4 {
		return errors.New("incorrect backgroundrect number of arguments")
	}

	values := make([]float64, 4)
	for i := 0; i < 4; i++ {
		value, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			return errors.New("incorrect backgroundrect arguments")
		}
		values[i] = value
	}

	x1, y1, x2, y2 := values[0], values[1], values[2], values[3]

	x1Int := int(x1 * 800)
	y1Int := int(y1 * 800)
	x2Int := int(x2 * 800)
	y2Int := int(y2 * 800)

	s.BackgroundRect = painter.BackgroundRect(x1Int, y1Int, x2Int, y2Int)

	return nil
}

// AddFigure добавляет фигуру.
func (s *UIState) AddFigure(args []string) error {
	if len(args) != 2 {
		return errors.New("incorrect figure number of arguments")
	}

	values := make([]float64, 2)
	for i := 0; i < 2; i++ {
		value, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			return errors.New("incorrect figure arguments")
		}
		values[i] = value
	}

	x, y := values[0], values[1]

	xInt := int(x * 800)
	yInt := int(y * 800)

	s.Figures = append(s.Figures, &painter.FigureOp{
		F: painter.Figure{X: xInt, Y: yInt},
	})

	return nil
}

// MoveFigures выполняет операцию перемещения фигур.
func (s *UIState) MoveFigures(args []string) error {
	if len(args) != 2 {
		return errors.New("incorrect move number of arguments")
	}

	values := make([]float64, 2)
	for i := 0; i < 2; i++ {
		value, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			return errors.New("incorrect move arguments")
		}
		values[i] = value
	}

	x, y := values[0], values[1]
	s.MoveOp = append(s.MoveOp, painter.Move(int(x*800), int(y*800), s.Figures))

	return nil
}

func (s *UIState) SetUpdateOperation() {
	s.UpdateOp = painter.UpdateOp
}

func (s *UIState) Reset() {
	s.BackgroundColor = nil
	s.BackgroundRect = nil
	s.Figures = nil
	s.MoveOp = nil
	s.UpdateOp = nil
}

func (s *UIState) ResetOperations() {
	if s.BackgroundColor == nil {
		s.BackgroundColor = painter.Reset()
	}
	if s.UpdateOp != nil {
		s.UpdateOp = nil
	}
}
