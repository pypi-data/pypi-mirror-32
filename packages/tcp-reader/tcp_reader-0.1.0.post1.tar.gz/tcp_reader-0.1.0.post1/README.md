# TCP Stream Reader

## Description

Reads raw ethernet data, splits them by TCP stream and use a TCPReader per stream.

## Example

Example usage of TCPManager with win_pcap:

```python
from tcp_reader import TCPManager
from winpcapy import WinPcapUtils

def print_callback(host, port, data):
    print('Received data from', host, 'on Port', port + ':', str(data))

tcp_manager = TCPManager(callback=print_callback, host=['93.184.216.34/32'], client=['192.168.0.0/16'])
WinPcapUtils.capture_on('*Ethernet*', tcp_manager.receive_raw_win_pcap)
```