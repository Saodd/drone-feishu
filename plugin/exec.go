package plugin

import (
	"context"
	"errors"
	"fmt"
	"github.com/saodd/alog"
	feishuRobotGo "github.com/saodd/go-feishu-robot"
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
	if err := robot.SendPost(c, content); err != nil {
		alog.CE(c, err, alog.V{"content": content})
		return err
	}
	return nil
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
				Title: fmt.Sprintf("%s/%s: %s (%s)", p.RepoInfo.Name, p.BuildInfo.Branch, p.StageInfo.Name, p.BuildInfo.Status),
			},
		},
	}

	if p.BuildInfo.Status == "success" {
		content.Post.ZhCn.Content = append(content.Post.ZhCn.Content, []feishuRobotGo.RobotPostContentGroupContent{
			{
				Tag:  "text",
				Text: "构建成功",
			},
		})
	} else {
		content.Post.ZhCn.Content = append(content.Post.ZhCn.Content, []feishuRobotGo.RobotPostContentGroupContent{
			{
				Tag:  "text",
				Text: "构建异常！！请查看：",
			},
			{
				Tag:  "a",
				Text: "日志",
				Href: p.BuildInfo.Link,
			},
		})
	}

	if p.Config.Message != "" {
		content.Post.ZhCn.Content = append(content.Post.ZhCn.Content, []feishuRobotGo.RobotPostContentGroupContent{
			{
				Tag:  "text",
				Text: p.Config.Message,
			},
		})
	}
	return content
}

// 不再默认使用表情标签，因为接口响应：[10002]not support emotion tag.
func generateEmojiType(p *Plugin) string {
	switch p.BuildInfo.Status {
	case "success":
		return "CheckMark"
	case "failure":
		return "CrossMark"
	default:
		return "WHAT"
	}
}
