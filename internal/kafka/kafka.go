package kafka

import (
	"context"
	"errors"
	"net"
	"sort"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

type Config struct {
	BootstrapServers []string
}

type Conn struct {
	bootstrapServers []string
	// writer *kafka.Writer
	// reader *kafka.Reader
	leader *kafka.Conn
}

var ErrNoBrokers = errors.New("no brokers")

func Connect(bootstrapServers ...string) (*Conn, error) {
	conn, err := connect(bootstrapServers...)
	if err != nil {
		return nil, err
	}

	return &Conn{
		bootstrapServers: bootstrapServers,
		leader:           conn,
	}, nil
}

func connect(bootstrapServers ...string) (*kafka.Conn, error) {
	if len(bootstrapServers) == 0 {
		return nil, ErrNoBrokers
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := kafka.DialContext(ctx, "tcp", bootstrapServers[0])
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return nil, err
	}

	return kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
}

type Topic struct {
	Name       string `json:"name"`
	Partitions int    `json:"partitions"`
}

type SortByName []Topic

func (a SortByName) Len() int           { return len(a) }
func (a SortByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func (c *Conn) GetTopics() ([]Topic, error) {
	if err := c.checkAndRestoreConn(); err != nil {
		return nil, err
	}

	return c.getTopics()
}

func (c *Conn) getTopics() ([]Topic, error) {
	partitions, err := c.leader.ReadPartitions()
	if err != nil {
		return nil, err
	}

	idx := map[string]int{}

	for _, p := range partitions {
		idx[p.Topic]++
	}

	var topics []Topic

	for topic, partitions := range idx {
		topics = append(topics, Topic{
			Name:       topic,
			Partitions: partitions,
		})
	}

	sort.Sort(SortByName(topics))

	return topics, nil
}

func (c *Conn) CreateTopic(name string, partitions, replicas int) (*Topic, error) {
	if err := c.checkAndRestoreConn(); err != nil {
		return nil, err
	}

	return c.createTopic(name, partitions, replicas)
}

func (c *Conn) createTopic(name string, partitions, replicas int) (*Topic, error) {
	if err := c.leader.CreateTopics(kafka.TopicConfig{
		Topic:             name,
		NumPartitions:     partitions,
		ReplicationFactor: replicas,
	}); err != nil {
		return nil, err
	}

	return &Topic{
		Name:       name,
		Partitions: partitions,
	}, nil
}

func (c *Conn) DeleteTopic(name string) error {
	if err := c.checkAndRestoreConn(); err != nil {
		return err
	}

	return c.deleteTopic(name)
}

func (c *Conn) deleteTopic(name string) error {
	return c.leader.DeleteTopics(name)
}

func (c *Conn) ping() error {
	_, err := c.leader.ApiVersions()
	return err
}

func (c *Conn) checkAndRestoreConn() error {
	if err := c.ping(); err == nil {
		return nil
	}

	conn, err := connect(c.bootstrapServers...)
	if err != nil {
		return err
	}

	c.leader = conn

	return nil
}
