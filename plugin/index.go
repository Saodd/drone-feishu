package plugin

import feishuRobotGo "github.com/saodd/go-feishu-robot"

func NewPlugin(contentBuilder func(p *Plugin) *feishuRobotGo.RobotContent) *Plugin {
	return &Plugin{BuildFeishuContent: contentBuilder}
}

type Plugin struct {
	RepoInfo           RepoInfo
	BuildInfo          BuildInfo
	StageInfo          StageInfo
	Config             Config
	BuildFeishuContent func(p *Plugin) *feishuRobotGo.RobotContent
}

type (
	RepoInfo struct {
		Owner string
		Name  string
	}

	// BuildInfo 针对的是step层级
	BuildInfo struct {
		Tag    string
		Event  string
		Number int
		Parent int
		Commit string
		Ref    string
		Branch string
		Author struct {
			Username string
			Name     string
			Email    string
			Avatar   string
		}
		Pull    string
		Message struct {
			Title string
			Body  string
		}
		DeployTo string
		Status   string
		Link     string
		Started  int64
		Created  int64
	}

	// StageInfo 针对的是pipeline层级
	StageInfo struct {
		Started int64
		Name    string
	}
)

type Config struct {
	Webhook string
	Secret  string
	Message string
}
