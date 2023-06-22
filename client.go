package mid

import (
	"bufio"
	"fmt"
	"net"
	"sync"

	"github.com/rs/zerolog"
)

const (
	jobInfoSub                 = "0035"
	vinSub                     = "0052"
	tighteningSub              = "0061"
	multiSpindelSub            = "0101"
	powerMACSTighteningSub     = "0106"
	powerMACSTighteningBoltSub = "0107"
)

type Client struct {
	conn      net.Conn
	feedback  *Publisher
	chans     sync.Map
	semaphore chan struct{}
	done      chan struct{}
	logger    zerolog.Logger
}

func NewClient(host string, port string, logger zerolog.Logger) (*Client, error) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}
	err = conn.SetKeepAlive(true)
	if err != nil {
		return nil, err
	}
	cln := &Client{
		conn:      conn,
		feedback:  NewPublisher(),
		chans:     sync.Map{},
		semaphore: make(chan struct{}, 1),
		done:      make(chan struct{}),
		logger:    logger,
	}
	go cln.read()
	return cln, nil
}

func (c *Client) Close() {
	close(c.done)
}

func (c *Client) ApplicationCommunicationStart() error {
	mid0001 := MID{
		Header: Header{
			Length:   20,
			MID:      1,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0001, func(mid MID) error {
		if mid.Header.MID == 4 {
			return midErr(mid)
		}
		if mid.Header.MID != 2 {
			return fmt.Errorf("invalid mid: %d", mid.Header.MID)
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) ApplicationCommunicationStop() error {
	mid0003 := MID{
		Header: Header{
			Length:   20,
			MID:      3,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0003, func(mid MID) error {
		if mid.Header.MID != 5 {
			return fmt.Errorf("invalid mid: %d", mid.Header.MID)
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) JobInfoSubscribe() (<-chan []byte, error) {
	p := NewPublisher()
	c.chans.Store(jobInfoSub, p)
	mid0034 := MID{
		Header: Header{
			Length:   20,
			MID:      34,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0034, standartHandler); err != nil {
		return nil, err
	}
	return p.Read(), nil
}

func (c *Client) JobInfoAcknowledge() error {
	mid0036 := MID{
		Header: Header{
			Length:   20,
			MID:      36,
			Revision: 1,
		},
	}
	return c.acknowledge(mid0036)
}

func (c *Client) JobInfoUnsubscribe() error {
	mid0037 := MID{
		Header: Header{
			Length:   20,
			MID:      37,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0037, standartHandler); err != nil {
		return err
	}
	return nil
}

func (c *Client) VehicleIDNumberSubscribe() (<-chan []byte, error) {
	p := NewPublisher()
	c.chans.Store(vinSub, p)
	mid0051 := MID{
		Header: Header{
			Length:   20,
			MID:      51,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0051, standartHandler); err != nil {
		return nil, err
	}
	return p.Read(), nil
}

func (c *Client) VehicleIDNumberAcknowledge() error {
	mid0053 := MID{
		Header: Header{
			Length:   20,
			MID:      53,
			Revision: 1,
		},
	}
	return c.acknowledge(mid0053)
}

func (c *Client) VehicleIDNumberUnsubscribe() error {
	mid0054 := MID{
		Header: Header{
			Length:   20,
			MID:      54,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0054, standartHandler); err != nil {
		return err
	}
	return nil
}

func (c *Client) LastTighteningResultDataSubscribe() (<-chan []byte, error) {
	p := NewPublisher()
	c.chans.Store(tighteningSub, p)
	mid0060 := MID{
		Header: Header{
			Length:   20,
			MID:      60,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0060, standartHandler); err != nil {
		return nil, err
	}
	return p.Read(), nil
}

func (c *Client) LastTighteningResultDataAcknowledge() error {
	mid0062 := MID{
		Header: Header{
			Length:   20,
			MID:      62,
			Revision: 1,
		},
	}
	return c.acknowledge(mid0062)
}

func (c *Client) LastTighteningResultDataUnsubscribe() error {
	mid0063 := MID{
		Header: Header{
			Length:   20,
			MID:      63,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0063, standartHandler); err != nil {
		return err
	}
	return nil
}

func (c *Client) MultiSpindleResultSubscribe() (<-chan []byte, error) {
	p := NewPublisher()
	c.chans.Store(multiSpindelSub, p)
	mid0100 := MID{
		Header: Header{
			Length:   20,
			MID:      100,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0100, standartHandler); err != nil {
		return nil, err
	}
	return p.Read(), nil
}

func (c *Client) MultiSpindleResultAcknowledge() error {
	mid0102 := MID{
		Header: Header{
			Length:   20,
			MID:      102,
			Revision: 1,
		},
	}
	return c.acknowledge(mid0102)
}

func (c *Client) MultiSpindleResultUnsubscribe() error {
	mid0103 := MID{
		Header: Header{
			Length:   20,
			MID:      103,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0103, standartHandler); err != nil {
		return err
	}
	return nil
}

func (c *Client) LastPowerMACSTighteningResultDataSubscribe() (<-chan []byte, error) {
	p := NewPublisher()
	c.chans.Store(powerMACSTighteningSub, p)
	c.chans.Store(powerMACSTighteningBoltSub, p)
	mid0105 := MID{
		Header: Header{
			Length:   20,
			MID:      105,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0105, standartHandler); err != nil {
		return nil, err
	}
	return p.Read(), nil
}

func (c *Client) LastPowerMACSTighteningResultDataAcknowledge(withBoltData bool) error {
	mid0108 := MID{
		Header: Header{
			Length:   21,
			MID:      108,
			Revision: 1,
		},
		Data: []byte("0"),
	}
	if withBoltData {
		mid0108.Data = []byte("1")
	}
	return c.acknowledge(mid0108)
}

func (c *Client) LastPowerMACSTighteningResultDataUnsubscribe() error {
	mid0109 := MID{
		Header: Header{
			Length:   20,
			MID:      109,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid0109, standartHandler); err != nil {
		return err
	}
	return nil
}

func (c *Client) KeepAliveMessage() error {
	mid9999 := MID{
		Header: Header{
			Length:   20,
			MID:      9999,
			Revision: 1,
		},
	}
	if err := c.execCMD(mid9999, func(mid MID) error {
		if mid.Header.MID != 9999 {
			return fmt.Errorf("invalid mid: %d", mid.Header.MID)
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) read() {
	defer func() {
		if r := recover(); r != nil {
			c.logger.Error().Interface("panic", r).Msg("Recover from panic")
			return
		}
		c.chans.Range(func(key, value any) bool {
			p, ok := value.(*Publisher)
			if ok {
				p.Close()
				c.chans.Delete(key)
			}
			return true
		})
		c.feedback.Close()
		c.conn.Close()
	}()
	for {
		select {
		case <-c.done:
			return
		default:
			data, err := bufio.NewReader(c.conn).ReadBytes('\x00')
			if err != nil {
				c.logger.Error().Err(err).Msg("Failed to read from connection")
				return
			}
			c.logger.Info().Bytes("data", data).Msg("Receive mid message")
			if len(data) < 20 {
				c.logger.Error().Msg("Invalid mid header lenght")
				return
			}
			go func(key string) {
				switch key {
				case
					jobInfoSub,
					vinSub,
					tighteningSub,
					multiSpindelSub,
					powerMACSTighteningSub,
					powerMACSTighteningBoltSub:
					v, _ := c.chans.Load(key)
					p, ok := v.(*Publisher)
					if ok {
						p.Write(data)
					}
				default:
					c.feedback.Write(data)
				}
			}(string(data[4:8]))
		}
	}
}

func (c *Client) execCMD(mid MID, f func(mid MID) error) error {
	if f == nil {
		return fmt.Errorf("nil feedback handler func")
	}
	payload, err := MarshalMID(mid)
	if err != nil {
		return err
	}
	raw, err := c.do(payload)
	if err != nil {
		return err
	}
	mid = MID{}
	err = UnmarshalMID(raw, &mid)
	if err != nil {
		return err
	}
	return f(mid)
}

func (c *Client) do(payload []byte) ([]byte, error) {
	c.semaphore <- struct{}{}
	defer func() { <-c.semaphore }()
	c.logger.Info().Bytes("data", payload).Msg("Send mid message")
	if _, err := c.conn.Write(append(payload, '\x00')); err != nil {
		return nil, err
	}
	data, ok := <-c.feedback.Read()
	if !ok {
		return nil, fmt.Errorf("error feedback")
	}
	return data, nil
}

func (c *Client) acknowledge(mid MID) error {
	payload, err := MarshalMID(mid)
	if err != nil {
		return err
	}
	c.logger.Info().Bytes("data", payload).Msg("Send mid message")
	if _, err := c.conn.Write(append(payload, '\x00')); err != nil {
		return err
	}
	return nil
}

func standartHandler(mid MID) error {
	if mid.Header.MID == 4 {
		return midErr(mid)
	}
	if mid.Header.MID != 5 {
		return fmt.Errorf("invalid mid: %d", mid.Header.MID)
	}
	return nil
}

func midErr(mid MID) error {
	mid0004 := &MID0004REV001{}
	if err := Unmarshal(append(make([]byte, 20), mid.Data...), mid0004); err != nil {
		return err
	}
	return mid0004
}
