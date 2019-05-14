# TGMSG
Simple application that able to send a Telegram message from name of a bot
## Usage
```
Usage:
  tgmsg [OPTIONS] Message...

Application Options:
  -p, --proxy=                     Proxy server address
  -b, --bot_id=                    Telegram bot token
  -c, --chat_id=                   Chat ID message should be sent to
  -m, --parse_mode=[Markdown|HTML] Parse mode. Could be 'Markdown' or 'HTML' (default: Markdown)
  -v, --verbose                    Show verbose debug information

Help Options:
  -h, --help                       Show this help message
```

## Build status
| Linux x64 | Windows x64 |
| :---:     | :---:       |
|[![Build Status](https://travis-ci.com/pshurgal/tgmsg.svg?branch=master)](https://travis-ci.com/pshurgal/tgmsg) | [![Build status](https://ci.appveyor.com/api/projects/status/1rj1an0f5uv8djpl/branch/master?svg=true)](https://ci.appveyor.com/project/pshurgal/tgmsg/branch/master) |
| [Latest release binary](https://github.com/pshurgal/tgmsg/releases/latest/download/tgmsg ) | [Latest release binary](https://github.com/pshurgal/tgmsg/releases/latest/download/tgmsg.exe ) |
