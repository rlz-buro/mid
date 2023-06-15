package mid

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Client struct {
	conn                   net.Conn
	feedback               chan []byte
	vinSub                 chan []byte
	tighteningSub          chan []byte
	multiSpindelSub        chan []byte
	powerMACSTighteningSub chan []byte
	semaphore              chan struct{}
	done                   chan struct{}
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
	cln := &Client{
		conn:      conn,
		feedback:  make(chan []byte),
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

func (c *Client) VehicleIDNumberSubscribe() (chan []byte, error) {
	c.vinSub = make(chan []byte)
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
	c.vinSub = make(chan []byte)
	return c.vinSub, nil
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
	c.tighteningSub = make(chan []byte)
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
	return c.tighteningSub, nil
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
	c.multiSpindelSub = make(chan []byte)
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
	return c.multiSpindelSub, nil
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
	c.powerMACSTighteningSub = make(chan []byte)
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
	return c.powerMACSTighteningSub, nil
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
		close(c.vinSub)
		close(c.tighteningSub)
		close(c.multiSpindelSub)
		close(c.powerMACSTighteningSub)
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
			switch string(data[4:8]) {
			case "0052":
				go func() {
					c.vinSub <- data
				}()
			case "0061":
				go func() {
					c.tighteningSub <- data
				}()
			case "0101":
				go func() {
					c.multiSpindelSub <- data
				}()
			case "0106", "0107":
				go func() {
					c.powerMACSTighteningSub <- data
				}()
			default:
				c.feedback <- data
			}
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
