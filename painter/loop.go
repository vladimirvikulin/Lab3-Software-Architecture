package painter

import (
	"image"
	"sync"

	"golang.org/x/exp/shiny/screen"
)

// Receiver отримує текстуру, яка була підготовлена в результаті виконання команд у циелі подій.
type Receiver interface {
	Update(t screen.Texture)
}

// Loop реалізує цикл подій для формування текстури отриманої через виконання операцій отриманих з внутрішньої черги.
type Loop struct {
	Receiver Receiver

	next screen.Texture // текстура, яка зараз формується
	prev screen.Texture // текстура, яка була відправленя останнього разу у Receiver

	Mq      messageQueue
	stopped chan struct{}
	stopReq bool
}

var size = image.Pt(800, 800)

// Start запускає цикл подій. Цей метод потрібно запустити до того, як викликати на ньому будь-які інші методи.
func (l *Loop) Start(s screen.Screen) {
	l.next, _ = s.NewTexture(size)
	l.prev, _ = s.NewTexture(size)

	l.stopped = make(chan struct{})

	go func() {
		for !l.stopReq || !l.Mq.empty() {
			op := l.Mq.pull()
			update := op.Do(l.next)
			if update {
				l.Receiver.Update(l.next)
				l.next, l.prev = l.prev, l.next
			}
		}
		close(l.stopped)
	}()
}

// Post додає нову операцію у внутрішню чергу.
func (l *Loop) Post(op Operation) {
	// TODO: реалізувати додавання операції в чергу. Поточна імплементація
	if op != nil {
		l.Mq.push(op)
	}
}

// StopAndWait сигналізує
func (l *Loop) StopAndWait() {
	l.Post(OperationFunc(func(screen.Texture) {
		l.stopReq = true
	}))
	<-l.stopped
}

// TODO: реалізувати власну чергу повідомлень.
type messageQueue struct {
	Ops []Operation
	Mu  sync.Mutex

	blocked chan struct{}
}

func (Mq *messageQueue) push(op Operation) {
	Mq.Mu.Lock()
	defer Mq.Mu.Unlock()

	Mq.Ops = append(Mq.Ops, op)

	if Mq.blocked != nil {
		close(Mq.blocked)
		Mq.blocked = nil
	}
}

func (Mq *messageQueue) pull() Operation {
	Mq.Mu.Lock()
	defer Mq.Mu.Unlock()

	for len(Mq.Ops) == 0 {
		Mq.blocked = make(chan struct{})
		Mq.Mu.Unlock()
		<-Mq.blocked
		Mq.Mu.Lock()
	}

	op := Mq.Ops[0]
	Mq.Ops[0] = nil
	Mq.Ops = Mq.Ops[1:]
	return op
}

func (Mq *messageQueue) empty() bool {
	Mq.Mu.Lock()
	defer Mq.Mu.Unlock()
	return len(Mq.Ops) == 0
}
