# GGPT (Golang GPT) - Simple Terminal GPT Chat Client

GGPT is a super simple GPT (Generative Pre-trained Transformer) chat client implemented in Golang, which is implemented to work with OpenAI models. It allows you to have text-based conversations with the GPT model via the terminal.

**Please note:** This is a basic implementation and may have limitations such as lack of history navigation and potential issues with sending complex text.

## Configuration

The configuration for GGPT is stored in the `.config.yaml` file located in the root of the repository. The configuration options are as follows:

- `gpt_role`: The default role to be used in the conversation.
- `api_url`: The API endpoint for the OpenAI GPT API.
- `api_key`: Your OpenAI API key (Leave empty if you prefer to provide it via other means).
- `max_tokens`: The maximum number of tokens in the response.
- `model`: The GPT model to use for generating responses.
- `roles`: A map of predefined roles and their corresponding descriptions.

## Usage

1. Clone this repository to your local machine.

2. Copy the `.ggpt_config.yaml` file in the root of the repository to your home, and configure it according to your needs. Refer to the configuration options mentioned above.

3. Build the project by running the following command in the repository's root directory:

   ```sh
   go build -o ggpt cmd/terminal/main.go 
   ```

4. Move the built executable to a directory in your system's PATH:

   ```sh
   sudo mv ggpt /usr/local/bin/
   ```

5. Run GGPT by simply typing `ggpt -role <your role>` in your terminal.

## Limitations

- Currently, GGPT does not support history navigation.
- Special characters and complex text may not be handled perfectly due to the simplicity of the implementation.

## Contributing

Contributions to this project are welcome! Feel free to open issues and submit pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
