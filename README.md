# Go Simple Webserver Example

This project is a simple example project for [phoihos/gosim](https://github.com/phoihos/gosim) package. 
- Developed based on **[vscode-remote-try-go](https://github.com/microsoft/vscode-remote-try-go)**

To build or test this project you have to install follows:
- Latest **VSCode** : https://code.visualstudio.com
    - **Remote - Containers** extension : https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers
- 2.0+ **Docker Desktop** : https://www.docker.com/products/docker-desktop
    - Check issue : [Docker 2.3.0.2 - open folder in container - error CreateProcess failed](https://github.com/microsoft/vscode-remote-release/issues/2975)

You can learn more about **[Developing inside a Container with VSCode](https://aka.ms/vscode-remote/containers)**.

## Setting up the development container

Follow these steps to open this project in a container:

1. Clone this repository.
    >
    ```shell
    git clone https://github.com/phoihos/gosim-example.git
    ```

2. Open VSCode.

3. Go to the Command Palette and launch **Remote-Containers: Open Folder in Container...** command.
   - Press <kbd>F1</kbd> key or use shortcut <kbd>Ctrl</kbd>+<kbd>Shift</kbd>+<kbd>P</kbd> on Windows (<kbd>CMD</kbd>+<kbd>Shift</kbd>+<kbd>P</kbd> on macOS) to open Command Palette.

4. Select the repository folder.

5. Wait a while for Docker Image Download and Container Creation.

6. Press <kbd>F5</kbd> key to launch example webserver.

## Documentation
- Guide for [.vscode/settings.json](.vscode/settings.json) : https://github.com/Microsoft/vscode-go/wiki/GOPATH-in-the-VS-Code-Go-extension

## License

Released under the [MIT License](LICENSE)
