package plugin

import (
	"github.com/urfave/cli"
	"strings"
)

// ParseAppArgs 用cli.App中读取的环境变量来填充Plugin的数据
func (p *Plugin) ParseAppArgs(c *cli.Context) {
	p.RepoInfo = RepoInfo{
		Owner: c.String("repo.owner"),
		Name:  c.String("repo.name"),
	}

	p.BuildInfo = BuildInfo{
		Tag:    c.String("build.tag"),
		Number: c.Int("build.number"),
		Parent: c.Int("build.parent"),
		Event:  c.String("build.event"),
		Status: c.String("build.status"),
		Commit: c.String("commit.sha"),
		Ref:    c.String("commit.ref"),
		Branch: c.String("commit.branch"),
		// Author...
		Pull: c.String("commit.pull"),
		// Message...
		DeployTo: c.String("build.deployTo"),
		Link:     c.String("build.link"),
		Started:  c.Int64("build.started"),
		Created:  c.Int64("build.created"),
	}

	p.BuildInfo.Author.Username = c.String("commit.author")
	p.BuildInfo.Author.Name = c.String("commit.author.name")
	p.BuildInfo.Author.Email = c.String("commit.author.email")
	p.BuildInfo.Author.Avatar = c.String("commit.author.avatar")

	{
		splitMsg := strings.Split(c.String("commit.message"), "\n")
		p.BuildInfo.Message.Title = strings.TrimSpace(splitMsg[0])
		p.BuildInfo.Message.Body = strings.TrimSpace(strings.Join(splitMsg[1:], "\n"))
	}

	p.StageInfo = StageInfo{
		Started: c.Int64("stage.started"),
		Name:    c.String("stage.name"),
	}

	p.Config = Config{
		Webhook: c.String("webhook"),
		Secret:  c.String("secret"),
		Message: c.String("message"),
	}
}
