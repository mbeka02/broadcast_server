# BROADCAST SERVER

This is a simple CLI based application that can be used to either start a server or connect to the server as a client. The server utilizes websockets to broadcast messagesto all the connected clients - https://roadmap.sh/projects/broadcast-server

When the server is started using the broadcast-server start command, it listens for client connections on a specified port (can be configured using command options otherwise port 3000 is used). When a client connects and sends a message, the server broadcasts this message to all connected clients using websockets.

## Usage

1. Build the application using this command `go build -o broadcast-server ./cmd`;
2. Run the server using `broadcast-server start`;
3. Run the client using `broadcast-server connect`;

## Dependencies

### Go packages

- [spf13/cobra](https://github.com/spf13/cobra)
