## Configuration Steps:

1. Go to golang.org

2. Download arm based version for m1 mac

3. Execute following statements:

export GOROOT=/usr/local/go
export GOPATH=$HOME/Desktop/folders/projects/go-workspace
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

5. Afterwards go to golang workspace (i.e., $HOME/Desktop/folders/projects/go-workspace)

6. Execute Following Commands:

mkdir video-chat
cd video-chat
go mod init video-chat-app

7. Place server directory and main.go except client directory

8. Then , excute go build main.go and then go run main.go

9. execute yarn create react-app client at root level. Make use of existing files and execute yarn start.

10. Open localhost:3000 and click at the button

11. Copy and paste the link in a new tab. We'll be able to establish video chat

## Screenshots
