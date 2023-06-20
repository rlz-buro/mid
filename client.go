package mid

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
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
	feedback  chan []byte
	chans     sync.Map
	semaphore chan struct{}
	done      chan struct{}
}

func NewClient(host string, port string) (*Client, error) {
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
		feedback:  make(chan []byte),
		chans:     sync.Map{},
		semaphore: make(chan struct{}, 1),
		done:      make(chan struct{}),
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

func (c *Client) JobInfoSubscribe() (chan []byte, error) {
	ch := make(chan []byte)
	c.chans.Store(jobInfoSub, ch)
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
	return ch, nil
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

func (c *Client) VehicleIDNumberSubscribe() (chan []byte, error) {
	ch := make(chan []byte)
	c.chans.Store(vinSub, ch)
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
	return ch, nil
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

func (c *Client) LastTighteningResultDataSubscribe() (chan []byte, error) {
	ch := make(chan []byte)
	c.chans.Store(tighteningSub, ch)
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
	return ch, nil
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

func (c *Client) MultiSpindleResultSubscribe() (chan []byte, error) {
	ch := make(chan []byte)
	c.chans.Store(multiSpindelSub, ch)
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
	return ch, nil
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

func (c *Client) LastPowerMACSTighteningResultDataSubscribe() (chan []byte, error) {
	ch := make(chan []byte)
	c.chans.Store(powerMACSTighteningSub, ch)
	c.chans.Store(powerMACSTighteningBoltSub, ch)
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
	return ch, nil
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
		c.chans.Range(func(key, value any) bool {
			v, _ := c.chans.Load(key)
			ch, ok := v.(chan []byte)
			if ok {
				close(ch)
				c.chans.Delete(key)
			}
			return true
		})
		close(c.feedback)
		c.conn.Close()
	}()
	for {
		select {
		case <-c.done:
			return
		default:
			data, err := bufio.NewReader(c.conn).ReadBytes('\x00')
			if err != nil {
				log.Printf("ERROR: %v\n", err)
				return
			}
			log.Println("RECEIVE:", string(data))
			if len(data) < 20 {
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
					ch, ok := v.(chan []byte)
					if ok {
						ch <- data
					}
				default:
					c.feedback <- data
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
	log.Println("SEND:", string(payload))
	if _, err := c.conn.Write(append(payload, '\x00')); err != nil {
		return nil, err
	}
	data, ok := <-c.feedback
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
	log.Println("SEND:", string(payload))
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
