package timer

import (
	"math/rand"
	"sync"
	"time"
)

type Timer interface {
	// Start the timer
	// Timer must be stopped, which is the state of a new timer.
	Start()
	// Pause the timer.
	// Timer must be started.
	Pause()
	// Resumed the timer.
	// Timer must be paused.
	Resume()
	// Stop the timer.
	// Timer must be started.
	Stop()
}

type timerState int

// A variable that is settable.
// The use of this interface allows for control
// on how the averaged timed value accessed.
type Setter interface {
	Set(int64)
}

const (
	Stopped timerState = iota
	Started
	Paused
)

// Perform basic timings of sections of code.
// Keeps a running average of timing values.
type timer struct {
	sampleRate float64
	i          int64

	start   time.Time
	current time.Duration
	avg     *movavg
	state   timerState

	avgVar Setter
	random *rand.Rand
}

func New(sampleRate float64, movingAverageSize int, avgVar Setter) Timer {

	return &timer{
		sampleRate: sampleRate,
		avg:        newMovAvg(movingAverageSize),
		avgVar:     avgVar,
		// Each timer gets its own random source or else
		// all timers would be locking on the global source.
		random: rand.New(rand.NewSource(rand.Int63())),
	}
}

// Start timer.
func (t *timer) Start() {
	if t.state != Stopped {
		panic("invalid timer state")
	}
	if t.random.Float64() < t.sampleRate {
		t.state = Started
		t.start = time.Now()
	}
}

// Pause current timing event.
func (t *timer) Pause() {
	if t.state != Started {
		return
	}
	t.current += time.Now().Sub(t.start)
	t.state = Paused
}

// Resumed paused timer.
func (t *timer) Resume() {
	if t.state != Paused {
		return
	}
	t.start = time.Now()
	t.state = Started
}

// Stop and record time of event.
// The moving average is updated at this point.
func (t *timer) Stop() {
	if t.state != Started {
		return
	}
	t.current += time.Now().Sub(t.start)
	// Use float64 precision when performing movavg calculations.
	avg := t.avg.update(float64(t.current))
	t.current = 0
	t.state = Stopped
	// Truncate to int64 now that we have a final value.
	t.avgVar.Set(int64(avg))
}

// Maintains a moving average of values
type movavg struct {
	size    int
	history []float64
	idx     int
	count   int
	avg     float64
}

func newMovAvg(size int) *movavg {
	return &movavg{
		size:    size,
		history: make([]float64, size),
		idx:     -1,
	}
}

func (m *movavg) update(value float64) float64 {
	m.count++
	n := float64(m.count)
	m.avg += (value - m.avg) / n
	m.idx = (m.idx + 1) % m.size

	if m.count == m.size+1 {
		old := m.history[m.idx]
		m.avg = (n*m.avg - old) / (n - 1)
		m.count--
	}
	m.history[m.idx] = value
	return m.avg
}

// A setter that can have distinct part or sections.
// By using a MultiPartSetter one can time distinct sections
// of code and have their individuals times summed
// to form a total timed value.
type MultiPartSetter struct {
	mu       sync.Mutex
	wg       sync.WaitGroup
	stopping chan struct{}

	setter     Setter
	partValues []int64
	updates    chan update
}

func NewMultiPartSetter(setter Setter) *MultiPartSetter {
	mp := &MultiPartSetter{
		setter:   setter,
		updates:  make(chan update),
		stopping: make(chan struct{}),
	}
	mp.wg.Add(1)
	go mp.run()
	return mp
}

// Add a new distinct part. As new timings are set
// for this part they will contribute to the total time.
func (mp *MultiPartSetter) NewPart() Setter {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	p := part{
		id:       len(mp.partValues),
		updates:  mp.updates,
		stopping: mp.stopping,
	}
	mp.partValues = append(mp.partValues, 0)

	return p
}

func (mp *MultiPartSetter) Stop() {
	close(mp.stopping)
	mp.wg.Wait()
}

func (mp *MultiPartSetter) run() {
	defer mp.wg.Done()
	for {
		select {
		case <-mp.stopping:
			return
		case update := <-mp.updates:
			mp.partValues[update.part] = update.value
			var sum int64
			for _, v := range mp.partValues {
				sum += v
			}
			mp.setter.Set(sum)
		}
	}
}

type update struct {
	part  int
	value int64
}

type part struct {
	id       int
	updates  chan<- update
	stopping chan struct{}
}

func (p part) Set(v int64) {
	select {
	case <-p.stopping:
	case p.updates <- update{part: p.id, value: v}:
	}
}
