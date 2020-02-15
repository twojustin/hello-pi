# Hello world for Golang on Raspberry Pi

1. Prepare SSH on Raspberry Pi via `sudo raspi-config`
2. compile go source for Pi's CPU
    * Pi 3: `env GOOS=linux GOARCH=arm GOARM=5 go build -o hello` 
3. copy compiled binary to pi: `scp hello pi@[rpi ip]:/home/pi/.`
4. ssh into pi `ssh pi@[raspberry pi host ip address]`
5. run it on pi with `./hello`
