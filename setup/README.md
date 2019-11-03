# Mesh Network Setup

## Usage

The ``run_me.sh`` script sets up all the necessary configuration files to make a Raspberry Pi connect to a mesh network.
It takes the following arguments:

```
1 - WiFi Channel - Number between 1 and 13
2 - SSID - Name of the mesh network
```

## Notes

* Please make sure your run the script as root
* You will have to reboot the Pi after running the script
* Both parameters have be exactly the same on all devices that intend to be connected.

## Example

```sudo ./run_me.sh 1 example-mesh-network```