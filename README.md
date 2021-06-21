# another-example-go-app

This app was to see how many bytes of memory channels of gopacket.Packet would take up at different capacities.

To create this go module, I made a directory in my GOPATH (~/go). On my local machine it is ~/go/src/size-gopacket.

Then I wrote the main.go file and ran `go mod init` to create the go.mod file. Then I ran `go mod tidy` to download the dependencies and create the go.sum file.

After that you can create the binary like usual with `go build .`

You can change the go.mod file manually and then rerun `go mod tidy` to update the dependencies if you want. 
For example if you delete a line in the 'require' section, `go mod tidy` will pull the latest version of it and list it in the mod and sum file.

To download the source files of a library you can run go get, like so: `go get -u github.com/trstruth/beacon`. 
If the repo is private, like my beacon fork, you need to have an SSH key setup from your dev environment. Then you can run `GOPRIVATE=github.com go get -u github.com/amyhuan/beacon`

You can also replace a dependency repo with your own fork of it, like I did for beacon.
I put this line in go.mod before the `require` section:
`replace github.com/trstruth/beacon => github.com/amyhuan/beacon latest`
and then I ran `go mod tidy` and this is what it ends up showing in the mod file: 
`replace github.com/trstruth/beacon => github.com/amyhuan/beacon v0.0.0-20210618180040-4a8ac605c0a8`

You can even replace a dependency with a local repo you have on your machine. Instead of a fork version, use an absolute path to the repo on your machine
`replace github.com/trstruth/beacon => /home/amyhuan/go/src/github.com/amyhuan/beacon`
