Toy app for controlling a vpn running from the pi

Compile the go app for arm.

```
GOARCH="arm" go build
```

This gives you a file vpnstate, which can be copied to the pi. Statically linked, so you don't need to worry about libraries etc. You can just run the executable on the pi, and it'll start listening on port 3001.

Also copy the script, vpn-control.sh to the pi. Run this script, pointing it to the webapp and the ovpn file for your vpn.

```
bash vpn-control.sh -u http://localhost:3001/poll -v ireland.ovpn
```

Both the webapp and the script should be run in a tmux so you don't need to stay logged in. Someday I may do systemd files.
