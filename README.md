The [Z-Cam](https://www.z-cam.com/) flagship line has [an API of sorts](https://github.com/imaginevision/Z-Camera-Doc/blob/master/E2/protocol/http.md). This can be used to control the camera--via a StreamDeck, say. 

This seems like a good enough reason to me to learn Go. 

Ideally this is useful for someone besides me. 

## Usage


`./zcam-controller [options] start` starts the camera recording
`./zcam-controller [options] stop` starts the camera recording
`./zcam-controller [options] download` downloads the files currently on the device's disk to a location you specify with the `-o` flag. An optional `--delete` at the end will then delete the files from the camera post-copy.

A `-u` or `--url` flag is required, and it expects either an IP address or a resolvable hostname for the z-cam. 

Putting it all together, `./zcam-controller -u 192.168.1.128 -o "/tmp" download --delete` is what a command looks like.