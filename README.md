# url-shortener
This is other of my personal projects, this is just a pretty basic url shortener.

It works with base64 and sha256 encoding to generate the shorten URL. It works only locally so you can't share the link to others. But hey, it still works.

## Usage
Run `go run . -url=<YOUR_URL>` to execute the program. Once you do this, the program will show you the shorten URL in the output so you can copy or just write it in your browser.

> [!IMPORTANT]
> The shorten URLs are saved on a JSON file. You have to create it manually. Name it as `shortened-urls.json`.