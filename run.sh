#!/bin/bash
#add conf
rm -f /etc/cni/net.d/10-corenet.conf
cp /opt/corenet/10-corenet.conf /etc/cni/net.d/10-corenet.conf
#add exec bin
rm -f /opt/cni/bin/corenet
cp /opt/corenet/corenet /opt/cni/bin/corenet
chmod +x /opt/corenet/corenet

