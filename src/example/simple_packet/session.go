package simple_packet

import (
	"fmt"
	"net"
	"time"

	"github.com/yuyyi51/packet-simulator/src/core"

	"github.com/yuyyi51/packet-simulator/src/simulate"
)

type Session struct {
	app                  *Application
	sentPacketHandler    *SentPacketHandler
	receivePacketHandler *ReceivePacketHandler
	conv                 int64
	remoteAddr           net.Addr

	writeBuffer   int64
	offset        int64
	receiveSorter *ReceiveDataSorter
	timer         *simulate.Timer

	lastIdleTime time.Time
	closed       bool
	logger       *core.Logger
}

func NewSession(app *Application, conv int64, remoteAddr net.Addr) *Session {
	return &Session{
		app:                  app,
		sentPacketHandler:    NewSentPacketHandler(),
		receivePacketHandler: NewReceivePacketHandler(),
		conv:                 conv,
		remoteAddr:           remoteAddr,
		receiveSorter:        NewReceiveDataSorter(),
		logger:               app.GetManager().GetLogger(),
	}
}

func (sess *Session) String() string {
	return fmt.Sprintf("Session [%08x]", sess.conv)
}

func (sess *Session) Start() {
	sess.timer = sess.app.CreateTimer(10*time.Millisecond, func() {
		sess.run()
	})
	sess.timer.Start()
	sess.lastIdleTime = sess.app.Now()
}

func (sess *Session) onPacketReceive(pkt *SimplePacket) {
	sess.logger.Debugf("%s received %s", sess, pkt)
	sess.lastIdleTime = sess.app.Now()
	sess.receivePacketHandler.onPacketReceive(pkt)
	content := pkt.Content
	switch c := content.(type) {
	case *SimplePacketContentData:
		sess.onReceiveData(c)
	case *SimplePacketContentAck:
		sess.onReceiveAck(c)
	case *SimplePacketContentPing:
		sess.onReceivePing(c)
	case *SimplePacketContentClose:
		sess.onReceiveClose(c)
	}
}

func (sess *Session) SimulateWrite(length int64) {
	sess.writeBuffer += length
}

func (sess *Session) createPacket(pktType SimplePacketType, content Content) *SimplePacket {
	seq := sess.sentPacketHandler.getNextSeq()
	pkt := &SimplePacket{
		PacketBase: &simulate.PacketBase{
			SourceAddr: sess.app.Addr(),
			TargetAddr: sess.remoteAddr,
		},
		Conv:       sess.conv,
		Seq:        seq,
		PacketType: pktType,
		Content:    content,
	}
	return pkt
}

func (sess *Session) sendAckPacket() {
	r := sess.receivePacketHandler.generateAckRanges()
	if len(r) == 0 {
		return
	}
	pkt := sess.createPacket(SimplePacketTypeAck, &SimplePacketContentAck{Ranges: r})
	sess.receivePacketHandler.needAck = false
	sess.sendPacket(pkt)
}

func (sess *Session) run() {
	//sess.logger.Infof("%s run", sess)
	if sess.app.Now().Sub(sess.lastIdleTime) > 10*time.Second {
		// need close
		sess.sendClosePacket()
		sess.closed = true
		sess.timer.Stop()
	}
	sess.sendPackets()
}

func (sess *Session) sendClosePacket() {
	pkt := sess.createPacket(SimplePacketTypeClose, &SimplePacketContentClose{})
	sess.sendPacket(pkt)
}

func (sess *Session) sendPackets() {
	if sess.closed {
		return
	}
	if sess.receivePacketHandler.needSendAck() {
		sess.sendAckPacket()
	}
	var cwnd int64 = 10 * 1400
	var mtu int64 = 1452 - SimplePacketDataOverhead
	for cwnd > 0 {
		data, more := sess.createDataContent(mtu)
		if data.Size == 0 {
			break
		}
		cwnd -= data.Size
		pkt := sess.createPacket(SimplePacketTypeData, data)
		//sess.logger.Infof("Session[%08x] send packet [%d]", sess.conv, pkt.Seq)
		sess.sendPacket(pkt)
		if !more {
			break
		}
	}
	sess.timer.Start()
}

func (sess *Session) sendPacket(pkt *SimplePacket) {
	sess.logger.Infof("%s send %s", sess, pkt)
	sess.sentPacketHandler.onSentPacket(pkt, sess.app.Now())
	sess.app.SendPacket(pkt)
}

func (sess *Session) createDataContent(maxBytes int64) (*SimplePacketContentData, bool /* more data */) {
	var dataCanWrite int64
	if sess.writeBuffer <= maxBytes {
		dataCanWrite = sess.writeBuffer
	} else {
		dataCanWrite = maxBytes
	}
	sess.writeBuffer -= dataCanWrite
	data := &SimplePacketContentData{
		Offset: sess.offset,
		Size:   dataCanWrite,
	}
	sess.offset += dataCanWrite

	return data, sess.writeBuffer != 0
}

func (sess *Session) onReceiveData(content *SimplePacketContentData) {
	//sess.logger.Infof("Session[%08x] received data, offset %d, size %d", sess.conv, content.Offset, content.Size)
	sess.receiveSorter.onReceiveData(content.Offset, content.Size)
	for sess.receiveSorter.hasData() {
		sess.receiveSorter.readData()
	}
}

func (sess *Session) onReceiveAck(content *SimplePacketContentAck) {
	for i, r := range content.Ranges {
		for j := r.left; j <= r.right; j++ {
			updated := sess.sentPacketHandler.onAckPacket(j, sess.app.Now(), j == r.right && i == len(content.Ranges)-1)
			if updated {
				sess.logger.Debugf("new rtt %s, latest rtt %s", sess.sentPacketHandler.rttStats.smoothedRtt, sess.sentPacketHandler.rttStats.latestRtt)
			}
		}
	}
}

func (sess *Session) onReceivePing(content *SimplePacketContentPing) {

}

func (sess *Session) onReceiveClose(content *SimplePacketContentClose) {
	sess.Close()
	sess.logger.Infof("Session[%s] received close and quit", fmt.Sprintf("%08x", sess.conv))
}

func (sess *Session) Close() {
	sess.timer.Stop()
	sess.closed = true
}
