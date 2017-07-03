#!/usr/bin/env bash

while getopts "u:v:" opt; do
    case $opt in
    u)
	POLL_URL=$OPTARG
	;;
    v)
	VPN_CONF=$OPTARG
	;;
    esac
done

usage() {
    echo "$0 -u <POLL_URL> -v <VPN_CONF>"
    exit 1
}

if [ -z "$POLL_URL" ]; then
    echo "Polling URL must be specified"
    usage
fi
if [ -z "$VPN_CONF" -o ! -f "$VPN_CONF" ]; then
    echo "Path to vpn conf must be specified and must exist"
    usage
fi

echo "Using URL $POLL_URL"

start_vpn() {
    echo "Starting vpn"
    sudo openvpn $VPN_CONF &
}

stop_vpn() {
    echo "Stopping vpn"
    sudo kill "$(jobs -p)"
}

NEW=$(curl "$POLL_URL?expected-current=OFF&timeout=0")
CURRENT=""

while true; do
    if [ "$NEW" != "$CURRENT" ]; then
	case $NEW in
	"ON")
	    start_vpn
	    ;;
	*)
	    stop_vpn
	    ;;
	esac
	CURRENT=$NEW
    fi

    NEW=$(curl "$POLL_URL?expected-current=$CURRENT&timeout=60")
done
