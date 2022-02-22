# Paste-Secret

Paste your GitHub Secrets in files

# Usage

## Inputs

**Required** `secrets` : Secrets ise JSON object array. Holds filename, keys and values which will be replaced

Structure of object array is `[{"filename":"TARGETFILE","keys"[{"key":"PARAMETER","value":"YOURSECRETVALUE"} ]}]`

Put your keys in curly braces `{{PARAMETER}}` in target file

## Examples

## Single line example
`example.txt`
```txt
    Hello {{name}}
```
`Usage In Action`
```yml
steps:
- uses:bariscanyilmaz/paste-secret@v1
    with:
        secrets:'[{"filename":"example.txt","keys":[{"key":"name","value":"baris"}]}]'

```

## Multi line example

`setting.json`
```json
{
	"api-url": "{{api}}",
}
```

`config.json`
```
{
    "Username": "{{user}}",
    "Password": "{{passwd}}",
}
```

`Usage In Action`

```yml
steps:
- uses: bariscanyilmaz/paste-secret@v1
    with:
        secrets: >
        [
            {"filename": "setting.json","keys":[{ "key": "api","value":"192.168.2.1"}]},
            {"filename": "config.json","keys":[{ "key": "user","value":"admin"},{ "key": "passwd","value":"admin123456"}]}
        ]
```

# Which Projects Using

[YTU Dining Hall Telegrom Bot](https://github.com/bariscanyilmaz/ytuyemekhane-telegram-bot) 
