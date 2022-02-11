package main

import (
	"errors"

	"github.com/opensourceways/community-robot-lib/config"
	framework "github.com/opensourceways/community-robot-lib/robot-gitee-framework"
	sdk "github.com/opensourceways/go-gitee/gitee"
	"github.com/sirupsen/logrus"
)

const botName = "scavenger"

type iClient interface{}

func newRobot(cli iClient) *robot {
	return &robot{cli: cli}
}

type robot struct {
	cli iClient
}

func (bot *robot) NewConfig() config.Config {
	return &configuration{}
}

func (bot *robot) getConfig(cfg config.Config) (*configuration, error) {
	if c, ok := cfg.(*configuration); ok {
		return c, nil
	}

	return nil, errors.New("can't convert to configuration")
}

func (bot *robot) RegisterEventHandler(f framework.HandlerRegitster) {
	f.RegisterIssueHandler(bot.handleIssueEvent)
	f.RegisterPullRequestHandler(bot.handlePREvent)
	f.RegisterNoteEventHandler(bot.handleNoteEvent)
	f.RegisterPushEventHandler(bot.handlePushEvent)
}

func (bot *robot) handlePREvent(e *sdk.PullRequestEvent, c config.Config, log *logrus.Entry) error {
	return nil
}

func (bot *robot) handleIssueEvent(e *sdk.IssueEvent, c config.Config, log *logrus.Entry) error {
	return nil
}

func (bot *robot) handlePushEvent(e *sdk.PushEvent, c config.Config, log *logrus.Entry) error {
	return nil
}

func (bot *robot) handleNoteEvent(e *sdk.NoteEvent, c config.Config, log *logrus.Entry) error {
	return nil
}
