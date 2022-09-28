package application

import (
	"context"
	"fmt"

	"github.com/x-foby/kakafka/internal/kafka"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Application struct {
	ctx context.Context
	cfg Config

	conns map[string]*kafka.Conn
}

func New() *Application {
	return &Application{
		conns: make(map[string]*kafka.Conn),
	}
}

func (a *Application) Startup(ctx context.Context) {
	a.ctx = ctx

	cfg, err := loadConfig(a.ctx)
	if err != nil {
		runtime.LogFatalf(a.ctx, "failed to load config: %s", err.Error())
	}

	a.cfg = *cfg
}

func (a *Application) GetConfigs() Config {
	return a.cfg
}

func (a *Application) Connect(profileName string) error {
	err := a.connect(profileName)
	if err != nil {
		runtime.MessageDialog(
			a.ctx,
			runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Message: err.Error(),
			},
		)
	}

	return err
}

func (a *Application) connect(profileName string) error {
	runtime.LogDebugf(a.ctx, "lookup for opened connects: %s", profileName)

	if _, ok := a.conns[profileName]; ok {
		return nil
	}

	runtime.LogDebugf(a.ctx, "lookup for profiles: %s", profileName)

	profile, ok := a.getProfileByName(profileName)
	if !ok {
		return fmt.Errorf("profile %q not found", profileName)
	}

	runtime.LogDebugf(a.ctx, "connecting for profiles: %s", profileName)

	conn, err := kafka.Connect(profile.BootstrapServers...)
	if err != nil {
		return err
	}

	runtime.LogDebug(a.ctx, "successifully connected")

	a.conns[profileName] = conn

	return nil
}

func (a *Application) GetTopics(profileName string) ([]kafka.Topic, error) {
	topics, err := a.getTopics(profileName)
	if err != nil {
		runtime.MessageDialog(
			a.ctx,
			runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Message: err.Error(),
			},
		)
	}

	return topics, err
}

func (a *Application) getTopics(profileName string) ([]kafka.Topic, error) {
	runtime.LogDebugf(a.ctx, "lookup for opened connects: %s", profileName)

	conn, ok := a.conns[profileName]
	if !ok {
		return nil, fmt.Errorf("profile %q not connected", profileName)
	}

	return conn.GetTopics()
}

func (a *Application) CreateTopic(profileName, topicName string, partitions, replicas int) (*kafka.Topic, error) {
	topics, err := a.createTopic(profileName, topicName, partitions, replicas)
	if err != nil {
		runtime.MessageDialog(
			a.ctx,
			runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Message: err.Error(),
			},
		)
	}

	return topics, err
}

func (a *Application) createTopic(profileName, topicName string, partitions, replicas int) (*kafka.Topic, error) {
	runtime.LogDebugf(a.ctx, "lookup for opened connects: %s", profileName)

	conn, ok := a.conns[profileName]
	if !ok {
		return nil, fmt.Errorf("profile %q not connected", profileName)
	}

	runtime.LogDebugf(a.ctx, "creating topic %s with %d partitions and %d replicas", topicName, partitions, replicas)

	return conn.CreateTopic(topicName, partitions, replicas)
}

func (a *Application) DeleteTopic(profileName, topicName string) error {
	if err := a.deleteTopic(profileName, topicName); err != nil {
		runtime.MessageDialog(
			a.ctx,
			runtime.MessageDialogOptions{
				Type:    runtime.ErrorDialog,
				Message: err.Error(),
			},
		)
	}

	return nil
}

func (a *Application) deleteTopic(profileName, topicName string) error {
	runtime.LogDebugf(a.ctx, "lookup for opened connects: %s", profileName)

	conn, ok := a.conns[profileName]
	if !ok {
		return fmt.Errorf("profile %q not connected", profileName)
	}

	runtime.LogDebugf(a.ctx, "deleting topic %s", topicName)

	return conn.DeleteTopic(topicName)
}

func (a *Application) getProfileByName(profileName string) (Profile, bool) {
	for _, p := range a.cfg.Profiles {
		if p.Name == profileName {
			return p, true
		}
	}

	return Profile{}, false
}
