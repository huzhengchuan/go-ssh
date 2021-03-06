package elfgate


const APP_VERSION = "1.4.0"
const APP_NAME    = "elfgate"


type CpFile struct {
    Filename    string
    Destination string
    Content     []byte
}

type CmdOutput struct {
    Host        string
    Output      []string
    Error       error
}

type ConfigStruct struct {
    Username    string
    Password    string
    PublicKey   string `yaml:"public_key"`
    Groups      map[string][]string
}

/* vim: set expandtab ts=4 sw=4 sts=4 tw=100: */
