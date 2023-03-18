package plugin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	feishuRobotGo "github.com/saodd/go-feishu-robot"
	"log"
)

func (p *Plugin) Exec(c context.Context) error {
	if err := p.CheckArgs(); err != nil {
		return err
	}

	var content *feishuRobotGo.RobotContent
	if p.BuildFeishuContent == nil {
		content = DefaultBuildFeishuContent(p)
	} else {
		content = p.BuildFeishuContent(p)
	}

	robot := &feishuRobotGo.Robot{
		Secret: p.Config.Secret,
		Hook:   p.Config.Webhook,
	}
	return robot.SendPost(c, content)
}

func (p *Plugin) CheckArgs() error {
	if p.Config.Webhook == "" {
		return errors.New("feishu Webhook is empty")
	}
	if p.Config.Secret == "" {
		return errors.New("feishu Secret is empty")
	}
	return nil
}

func DefaultBuildFeishuContent(p *Plugin) *feishuRobotGo.RobotContent {
	content := &feishuRobotGo.RobotContent{
		Post: feishuRobotGo.RobotPostContent{
			ZhCn: feishuRobotGo.RobotPostContentGroup{
				Title: fmt.Sprintf("<%s>%s/%s: %s", p.BuildInfo.Status, p.RepoInfo.Name, p.BuildInfo.Branch, p.StageInfo.Name),
				Content: [][]feishuRobotGo.RobotPostContentGroupContent{
					{
						{
							Tag:  "a",
							Text: "构建日志",
							Href: p.BuildInfo.Link,
						},
					},
				},
			},
		},
	}
	if p.Config.Message != "" {
		content.Post.ZhCn.Content = append(content.Post.ZhCn.Content, []feishuRobotGo.RobotPostContentGroupContent{
			{
				Tag:  "text",
				Text: p.Config.Message,
			},
		})
	}
	log.Println(content)
	j, _ := json.Marshal(content)
	log.Println(string(j))
	return content
}
