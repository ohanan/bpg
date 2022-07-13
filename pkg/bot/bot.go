package bot

type Bot interface {
	Init(ctx Context, config *Config) error
}

func Run(bot Bot) {

}
