package painter

import (
	"golang.org/x/exp/shiny/screen"
	"image"
	"image/color"
)

// Operation змінює вхідну текстуру.
type Operation interface {
	// Do виконує зміну операції, повертаючи true, якщо текстура вважається готовою для відображення.
	Do(t screen.Texture) (ready bool)
}

// OperationList групує список операції в одну.
type OperationList []Operation

func (ol OperationList) Do(t screen.Texture) (ready bool) {
	for _, o := range ol {
		ready = o.Do(t) || ready
	}
	return
}

// UpdateOp операція, яка не змінює текстуру, але сигналізує, що текстуру потрібно розглядати як готову.
var UpdateOp = updateOp{}

type updateOp struct{}

func (op updateOp) Do(t screen.Texture) bool { return true }

// OperationFunc використовується для перетворення функції оновлення текстури в Operation.
type OperationFunc func(t screen.Texture)

func (f OperationFunc) Do(t screen.Texture) bool {
	f(t)
	return false
}

// WhiteFill зафарбовує тестуру у білий колір. Може бути викоистана як Operation через OperationFunc(WhiteFill).
func WhiteFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.White, screen.Src)
}

// GreenFill зафарбовує тестуру у зелений колір. Може бути викоистана як Operation через OperationFunc(GreenFill).
func GreenFill(t screen.Texture) {
	t.Fill(t.Bounds(), color.RGBA{G: 0xff, A: 0xff}, screen.Src)
}

func BackgroundRect(x1, y1, x2, y2 int) OperationFunc {
	return func(t screen.Texture) {
		bounds := image.Rect(x1, y1, x2, y2)
		t.Fill(bounds, color.Black, screen.Src)
	}
}

type Figure struct {
	X, Y int
}

// FigureOp використовується для намалювання фігури з вказаними координатами.
type FigureOp struct {
	F Figure
}

func (op FigureOp) Do(t screen.Texture) bool {
	op.Figure()(t)
	return false
}

func (op *FigureOp) Figure() OperationFunc {
	return func(t screen.Texture) {
		t.Fill(image.Rect(op.F.X-200, op.F.Y+75, op.F.X+200, op.F.Y-75), color.RGBA{R: 255, G: 255, B: 0, A: 1}, screen.Src)
		t.Fill(image.Rect(op.F.X-75, op.F.Y+200, op.F.X+75, op.F.Y-200), color.RGBA{R: 255, G: 255, B: 0, A: 1}, screen.Src)
	}
}
func Move(dx, dy int, figures []*FigureOp) OperationFunc {
	return func(t screen.Texture) {
		for _, op := range figures {
			op.F.X += dx
			op.F.Y += dy
		}
	}
}

func Reset() OperationFunc {
	return func(t screen.Texture) {
		t.Fill(t.Bounds(), color.Black, screen.Src)
	}
}
