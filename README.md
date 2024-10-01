# terraform-prettyplan
Simple AI tool to convert terraform plans into human readable texts

## Installation

1. **Download the latest release**:

   Go to the [releases page](https://github.com/pasali/terraform-prettyplan/releases) and download the latest binary for your operating system.

2. **Extract the binary**:

    ```sh
    tar -xzf terraform-prettyplan-<version>-<os>-<arch>.tar.gz
    ```

3. **Move the binary to a directory in your PATH**:

    ```sh
    mv tfpp /usr/local/bin/
    ```

## Usage

1. **Export your OpenAI API key:**

    ```sh
    export TFPP_OPENAI_API_KEY="your_openai_api_key"
    ```

2. **Pipe terraform plan output directly:**
    ```sh
    terraform plan | tfpp
    ```