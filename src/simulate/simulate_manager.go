package simulate

import (
	"time"

	"github.com/yuyyi51/packet-simulator/src/core"
)

type SimulateManagerI interface {
	CreateEvent(time time.Time) core.EventI
	AddEvent(event core.EventI) error
	GetCurrentTime() time.Time
	GetCurrentTimeOffset() time.Duration
}

type SimulateManager struct {
	eventManager core.EventManagerI
}

func NewSimulateManager() *SimulateManager {
	return &SimulateManager{eventManager: core.NewEventManager(time.Now())}
}

func (s *SimulateManager) CreateEvent(time time.Time) core.EventI {
	return s.eventManager.CreateEvent(time)
}

func (s *SimulateManager) AddEvent(event core.EventI) error {
	return s.eventManager.AddEvent(event)
}

func (s *SimulateManager) GetCurrentTime() time.Time {
	return s.eventManager.GetCurrentTime()
}

func (s *SimulateManager) GetCurrentTimeOffset() time.Duration {
	return s.eventManager.GetCurrentTimeOffset()
}

func (s *SimulateManager) Run() {
	s.eventManager.Run()
}

func (s *SimulateManager) CreateDevice() *Device {
	d := NewDevice()
	d.manager = s
	return d
}

func (s *SimulateManager) CreateRouter(bandwidth int64, maxPktQueueLen int) *Router {
	r := NewRouter(bandwidth, maxPktQueueLen)
	r.manager = s
	return r
}

func (s *SimulateManager) CreateLink(peer1 DeviceI, peer2 DeviceI) *Link {
	l := NewLink(peer1, peer2)
	l.manager = s
	return l
}

func (s *SimulateManager) ConnectDevice(peer1 DeviceI, peer2 DeviceI) *Link {
	l := s.CreateLink(peer1, peer2)
	peer1.AddLink(l)
	peer2.AddLink(l)
	return l
}

func (s *SimulateManager) CreateApplicationBase(port int) *ApplicationBase {
	app := NewApplicationBase(port)
	app.manager = s
	return app
}
