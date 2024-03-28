package boot

import (
	"simple-skeleton/conf"
)

// StaterContext 启动器参数
type StaterContext map[string]interface{}

const (
	ConfKey = "conf"
)

func (s StaterContext) Conf() *conf.Config {
	v := s[ConfKey]

	if v == nil {
		panic("配置文件未被初始化")
	}

	return v.(*conf.Config)
}

// Starter 资源启动器
type Starter interface {
	// Init 初始化
	Init(ctx StaterContext)
	// Setup 安装
	Setup(ctx StaterContext)
	// Start 启动
	Start(ctx StaterContext)
	// StarterBlocking 是否阻塞
	StarterBlocking() bool
	// Stop 资源停止和销毁
	Stop(ctx StaterContext)
}

// BaseStarter 基础starter
type BaseStarter struct{}

func (b *BaseStarter) Init(ctx StaterContext) {}

func (b *BaseStarter) Setup(ctx StaterContext) {}

func (b *BaseStarter) Start(ctx StaterContext) {}

func (b *BaseStarter) StarterBlocking() bool { return false }

func (b *BaseStarter) Stop(ctx StaterContext) {}

// starterRegister 资源注册
type starterRegister struct {
	starters []Starter
}

func (s *starterRegister) Register(starter Starter) {
	s.starters = append(s.starters, starter)
}

func (s *starterRegister) AllStarter() []Starter {
	return s.starters
}

var StarterRegister *starterRegister = new(starterRegister)

func Register(s Starter) {
	StarterRegister.Register(s)
}
