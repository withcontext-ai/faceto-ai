package liveGPT

import (
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"sync"
	"sync/atomic"
	"time"

	"faceto-ai/internal/pkg/utils"
	"github.com/livekit/protocol/logger"
	lksdk "github.com/livekit/server-sdk-go"
	"github.com/pion/webrtc/v3"
	"github.com/pion/webrtc/v3/pkg/media"
)

var (
	ErrMuted         = errors.New("the track is muted")
	ErrInvalidFormat = errors.New("invalid format")

	OpusSilenceFrame = []byte{
		0xf8, 0xff, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	OpusSilenceFrameDuration = 20 * time.Millisecond
)

type GPTTrack struct {
	sampleTrack *lksdk.LocalSampleTrack
	provider    *provider
	log         *log.Helper

	doneChan   chan struct{}
	closedChan chan struct{}
}

func NewGPTTrack(log *log.Helper) (*GPTTrack, error) {
	cap := webrtc.RTPCodecCapability{
		Channels:  1,
		MimeType:  webrtc.MimeTypeOpus,
		ClockRate: 48000,
	}

	track, err := lksdk.NewLocalSampleTrack(cap)
	if err != nil {
		return nil, err
	}

	provider := &provider{log: log}
	err = track.StartWrite(provider, func() {})
	if err != nil {
		return nil, err
	}

	return &GPTTrack{
		log:         log,
		sampleTrack: track,
		provider:    provider,
		doneChan:    make(chan struct{}),
		closedChan:  make(chan struct{}),
	}, nil
}

func (t *GPTTrack) Publish(lp *lksdk.LocalParticipant) (pub *lksdk.LocalTrackPublication, err error) {
	pub, err = lp.PublishTrack(t.sampleTrack, &lksdk.TrackPublicationOptions{})
	return
}

func (t *GPTTrack) OnConsumer(f func(text string)) {
	t.provider.OnConsumer(f)
}

func (t *GPTTrack) OnSpeaking(f func(data []byte, err error)) {
	t.provider.OnSpeaking(f)
}

func (t *GPTTrack) OnStop(flag bool) {
	t.provider.OnStop(flag)
}

func (t *GPTTrack) OnBreak(f func(text string)) {
	t.provider.OnBreak(f)
}

// Called when the last oggReader in the queue finished being read
func (t *GPTTrack) OnComplete(f func(err error)) {
	t.provider.OnComplete(f)
}

func (t *GPTTrack) QueueReader(reader io.Reader) error {
	oggReader, oggHeader, err := utils.NewOggReader(reader)
	if err != nil {
		return err
	}

	// oggHeader.SampleRate is _not_ the sample rate to use for playback.
	// see https://www.rfc-editor.org/rfc/rfc7845.html#section-3
	if oggHeader.Channels != 1 /*|| oggHeader.SampleRate != 48000*/ {
		return ErrInvalidFormat
	}

	t.provider.QueueReader(oggReader)
	return nil
}

func (t *GPTTrack) QueueReaderString(text string) error {
	t.provider.QueueString(text)
	return nil
}

type provider struct {
	log         *log.Helper
	reader      *utils.OggReader
	lastGranule uint64

	queue       []*utils.OggReader
	queueString []string
	lock        sync.Mutex
	onStop      atomic.Bool
	onComplete  func(err error)
	onSpeaking  func(data []byte, err error)
	onConsumer  func(text string)
	onBreak     func(text string)
	currentText string
}

func (p *provider) NextSample() (media.Sample, error) {
	p.lock.Lock()
	onComplete := p.onComplete
	if p.reader == nil && len(p.queue) > 0 {
		p.lastGranule = 0
		p.reader = p.queue[0]
		p.queue = p.queue[1:]

		if len(p.queueString) > 0 {
			p.currentText = p.queueString[0]
			p.queueString = p.queueString[1:]
			if p.onConsumer != nil && !p.onStop.Load() {
				p.onConsumer(p.currentText)
			}
		}
	}
	p.lock.Unlock()

	if p.reader != nil {
		data, err := p.reader.ReadPacket()

		// stop
		if p.onStop.Load() {
			p.log.Debugf("[[[[[[[[p.onStop]]]]]]]], len(p.queue):%d, currentText:%s", len(p.queue), p.currentText)
			p.queue = make([]*utils.OggReader, 0)
			p.onBreak(p.currentText)
			p.reader = nil
			return p.NextSample()
		}

		if p.onSpeaking != nil {
			p.onSpeaking(data, err)
		}

		if err != nil {
			if onComplete != nil {
				onComplete(err)
			}

			if err == io.EOF {
				p.reader = nil
				return p.NextSample()
			} else {
				logger.Errorw("failed to parse next page", err)
				return media.Sample{}, err
			}
		}

		duration, err := utils.ParsePacketDuration(data)
		if err != nil {
			return media.Sample{}, err
		}

		return media.Sample{
			Data:     data,
			Duration: duration,
		}, nil
	}

	// Otherwise send empty Opus frames
	return media.Sample{
		Data:     OpusSilenceFrame,
		Duration: OpusSilenceFrameDuration,
	}, nil
}

func (p *provider) OnBind() error {
	return nil
}

func (p *provider) OnUnbind() error {
	return nil
}

func (t *provider) OnConsumer(f func(text string)) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.onConsumer = f
}

func (t *provider) OnSpeaking(f func(data []byte, err error)) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.onSpeaking = f
}

func (t *provider) OnStop(flag bool) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.onStop.Swap(flag)
}

func (t *provider) OnBreak(f func(text string)) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.onBreak = f
}

// Called when the *one* oggReader finished reading
func (t *provider) OnComplete(f func(err error)) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.onComplete = f
}

func (p *provider) QueueReader(reader *utils.OggReader) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.queue = append(p.queue, reader)
}

func (p *provider) QueueString(text string) {
	p.lock.Lock()
	defer p.lock.Unlock()

	p.queueString = append(p.queueString, text)
}
