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

To redirect nginx, use this conf...
```
server {
    listen 80;
    listen [::]:80;

    server_name vpn;

    location / {
        proxy_pass http://127.0.0.1:3001;
    }
}
```

The pi is configured with an IP on both 192.168.1.0/24 and 192.168.2.0/24. 192.168.1.0/24 is the main home network.

Hosts on 192.168.2.0/24 have the pi address set as the gateway. The pi masquerades all this traffic. IP forward is also enabled.

```
iptables -t nat -A POSTROUTING -s 192.168.2.0/24 -j MASQUERADE
```
