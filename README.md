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

<img width="1173" alt="webrtc-go" src="https://user-images.githubusercontent.com/43849911/153179058-0fe6e66c-8319-41d3-9b08-20fe93147e62.png">
<img width="1069" alt="webrtc1" src="https://user-images.githubusercontent.com/43849911/153179125-8b55e5ca-2560-495c-b33f-3bf5b2d05edf.png">
<img width="1084" alt="webrtc2" src="https://user-images.githubusercontent.com/43849911/153179130-3216e3f6-bdd9-4885-abcc-5e171256ae2e.png">
<img width="966" alt="webrtc3" src="https://user-images.githubusercontent.com/43849911/153179160-3e56b2e3-2e69-4522-afe4-4866ab05cac8.png">
<img width="1018" alt="webrtc4" src="https://user-images.githubusercontent.com/43849911/153179171-b3876835-e5b4-41ff-bae4-2c3208bfbc70.png">
<img width="881" alt="webrtc5" src="https://user-images.githubusercontent.com/43849911/153179233-bdd8ea1e-49f9-4ac6-bea1-2e6e133d79fa.png">
<img width="879" alt="webrtc6" src="https://user-images.githubusercontent.com/43849911/153179237-2925d3e9-4935-413d-9f36-c599f04524c1.png">
<img width="1440" alt="go" src="https://user-images.githubusercontent.com/43849911/153179242-1d404afc-3c10-4b79-a8fb-1967c271430f.png">
