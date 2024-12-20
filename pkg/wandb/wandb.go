package wandb



func Finish() {
}

func SetConfig(config map[string]any) {
}

// TODO: have a messenger object that groups requests etc.
// TODO: register signal handlers for SIGTERM and SIGINT to call messenger.finish()
