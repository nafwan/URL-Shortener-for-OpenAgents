# URL Shortener Plugin For OpenAgents

This plugin is designed to shorten long URLs using the [CleanURI API](https://cleanuri.com/). It takes a long URL as input and returns a shortened version of the URL as output.

## Usage

1. Install the Extism CLI and the URL Shortener Plugin.
2. Run the plugin with the long URL as input:

```bash
$ extism call plugin.wasm run --input 'https://openagents.com/' --wasi --allow-host='cleanuri.com'
https://cleanuri.com/kd9P71
```

## How it Works

1. The plugin reads the input URL and validates it to ensure it doesn't contain any spaces, tabs, or newlines.
2. The long URL is encoded into a JSON payload.
3. An HTTP POST request is sent to the CleanURI API endpoint with the JSON payload as the request body.
4. The API response is parsed, and any errors are handled appropriately.
5. If the API response is successful, the shortened URL is set as the plugin output.

## Error Handling

The plugin handles the following types of errors:

- Invalid input URL (contains spaces, tabs, or newlines)
- Failure to create the request payload
- API request failure (non-200 status code)
- Failure to parse the API response
- Errors returned by the CleanURI API

If any of these errors occur, the plugin will print an error message and exit with a non-zero status code.
