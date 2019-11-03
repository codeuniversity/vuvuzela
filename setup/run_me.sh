# Check if both parameters exist

if [ -z "$1" ] || [ -z "$2" ]
then
  echo "Please provide two parameters (see README)."
  exit 1
fi

# Create and fill starting script for network interfaces
echo "Configuring batman-adv..."

cd ~ || exit
printf '#!/bin/bash\nsudo batctl if add wlan0\nsudo ifconfig wlan0 up\nsudo ifconfig bat0 up' > start-batman-adv.sh
chmod +x start-batman-adv.sh

# Create config files for network interfaces (TODO: Adjust for Debian Buster)

printf 'auto bat0\niface bat0 inet auto\n\tpre-up /usr/sbin/batctl if add wlan0' > /etc/network/interfaces.d/bat0
printf "auto wlan0\niface wlan0 inet manual\n\tmtu 1532\n\twireless-channel %s\n\twireless-essid %s\n\twireless-mode ad-hoc\n\twireless-ap 02:12:34:56:78:9A" "$1" "$2" > /etc/network/interfaces.d/wlan0

# Install batctl and configure
echo "Installing batctl..."

sudo apt-get install -y batctl
echo 'batman-adv' | sudo tee --append /etc/modules
echo 'denyinterfaces wlan0' | sudo tee --append /etc/dhcpcd.conf
echo "$(pwd)/start-batman-adv.sh" >> ~/.bashrc

# Install Alfred
echo "Installing Alfred..."

sudo apt-get install libcap-dev gpsd libgps-dev libnl-3-dev libnl-genl-3-dev # install dependencies
wget https://github.com/open-mesh-mirror/alfred/archive/v2017.3.zip
unzip v2017.3.zip
rm v2017.3.zip
cd alfred-2017.3 || exit
make
sudo make install
cd ..
rm -rf alfred-2017.3

echo "Configuration done. Reboot to complete."