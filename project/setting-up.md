# Configuring the new server
adduser tcharles
usermod -aG sudo tcharles

# Configuring the firewall
ufw allow ssh
ufw allow http
ufw allow https

ufw allow 2377/tcp
ufw allow 7946/tcp
ufw allow 7946/udp
ufw allow 4789/udp
ufw allow 8025/tcp

ufw enable

ufw status

>>>  Exit and login with new user

# Install docker
sudo apt-get update

sudo apt upgrade

sudo apt-get install \
ca-certificates \
curl \
gnupg \
lsb-release

sudo mkdir -p /etc/apt/keyrings

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg

echo \
"deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin


# Define hostname
sudo hostnamectl set-hostname node-1

# Define /etc/hosts with all nodes and ips
node-1 IP
node-2 IP

# Define DNS of a domain
A node-1 IP
A node-2 IP
A swarm  IP

# Starting the swarm inside server
sudo docker swarm init --advertise-addr 172.104.195.230


# Inside each server

cd /
sudo mkdir swarm
sudo chown tcharles:tcharles swarm/
cd swarm/
mkdir caddy_data
mkdir caddy_config
vi swarm.yml

>>> Paste the caddy.production.dockerfile there
