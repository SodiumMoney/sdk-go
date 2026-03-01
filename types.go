package sodium

type Platform string

const (
	PlatformTwitter  Platform = "twitter"
	PlatformTelegram Platform = "telegram"
)

type Network string

const (
	Devnet  Network = "devnet"
	Mainnet Network = "mainnet"
)
