# Paste-Secret

Embed your GitHub Secrets in files

# Usage

## Inputs

**Required** `secrets` : Secrets ise JSON object array. Holds filename, keys and values which will be replaced

Structure of object array is `[{"filename":"TARGETFILE","keys"[{"key":"PARAMETER","value":"YOURSECRETVALUE"} ]}]`

Put your keys in curly braces `{{PARAMETER}}` in target file

## Examples

## Singleline Example

```txt
    Hello {{name}}
```

```yml
steps:
- uses:bariscanyilmaz/paste-secret@v1
    with:
        secrets:'[{"filename":"example.txt","keys":[{"key":"name","value":"baris"}]}]'

```

## Multiline Example

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

## Usage In Action

```yml
steps:
- uses: bariscanyilmaz/paste-secret@v1
    with:
        secrets: |
        [
            {"filename": "setting.json","keys":[{ "key": "api","value":"192.168.2.1"}]},
            {"filename": "config.json","keys":[{ "key": "user","value":"admin"},{ "key": "passwd","value":"admin123456"}]}
        ]
```