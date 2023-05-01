package painter

import (
	"image"
	"image/color"
	"image/draw"
	"testing"

	"golang.org/x/exp/shiny/screen"
)

func TestLoop_Start_Post(t *testing.T) {
	var (
		l  Loop
		tr testReceiver
	)

	l.Receiver = &tr
	l.Start(mockScreen{})

	l.Post(OperationFunc(WhiteFill))
	l.Post(OperationFunc(GreenFill))
	l.Post(UpdateOp)

	if tr.LastTexture != nil {
		t.Error("Reciever got the texture too early")
	}

	l.StopAndWait()

	tx, ok := tr.LastTexture.(*mockTexture)
	if !ok {
		t.Error("No textures")
	}
	if tx.FillCnt != 2 {
		t.Error("Incorrect number of calls:", tx.FillCnt)
	}
}

func TestLoop_Push_Pull(t *testing.T) {

	var (
		mq messageQueue
		op Operation
	)

	mq.push(op)

	if len(mq.Ops) != 1 {
		t.Error("Failed to push")
	}

	pulledOp := mq.pull()
	if pulledOp != op {
		t.Error("Failed to pull")
	}

	if !mq.empty() {
		t.Error("Queue is not empty")
	}
}

type testReceiver struct {
	LastTexture screen.Texture
}

func (tr *testReceiver) Update(t screen.Texture) {
	tr.LastTexture = t
}

type mockScreen struct{}

func (m mockScreen) NewBuffer(size image.Point) (screen.Buffer, error) { return nil, nil }

func (m mockScreen) NewTexture(size image.Point) (screen.Texture, error) {
	return new(mockTexture), nil
}

func (m mockScreen) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) { return nil, nil }

type mockTexture struct {
	FillCnt int
}

func (m *mockTexture) Release() {}

func (m *mockTexture) Size() image.Point { return size }

func (m *mockTexture) Bounds() image.Rectangle { return image.Rectangle{Max: size} }

func (m *mockTexture) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {}

func (m *mockTexture) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	m.FillCnt++
}
