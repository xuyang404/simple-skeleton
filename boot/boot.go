package boot

import (
	"simple-skeleton/conf"
)

type BootApplication struct {
	ctx  StaterContext
	conf *conf.Config
}

func NewBootApplication(conf *conf.Config) *BootApplication {
	b := &BootApplication{conf: conf, ctx: StaterContext{}}
	b.ctx[ConfKey] = b.conf
	return b
}

func (b *BootApplication) Start() {
	// 2、初始化starter
	b.init()
	// 3、安装starter
	b.setup()
	// 4、启动starter
	b.start()
}

func (b *BootApplication) init() {
	for _, s := range StarterRegister.AllStarter() {
		s.Init(b.ctx)
	}
}

func (b *BootApplication) setup() {
	for _, s := range StarterRegister.AllStarter() {
		s.Setup(b.ctx)
	}
}

func (b *BootApplication) start() {
	for i, s := range StarterRegister.AllStarter() {
		// 如果是最后一个就直接启动
		if i+1 == len(StarterRegister.AllStarter()) {
			s.Start(b.ctx)
		} else {
			// 否则的话判断一下是不是会阻塞，如果会阻塞，就开协程跑
			if s.StarterBlocking() {
				go s.Start(b.ctx)
			} else {
				s.Start(b.ctx)
			}
		}
	}
}
