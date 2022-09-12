# Æsir Gosite
## Hephaestus-Go
###
#### Hephaestus, the creator of Zeus’ thunderbolt, Poseidon's trident, and more.
#### The Hephaestus-Go app is a program that auto-generates the Golang environment for rapid deployment of web apps.
#### The goal is to minimize setup time on a new machine or virtual machine, such as AWS ec2 instances. 
###

## Using Hephaestus-Go
###
#### Download HephaestusGo.go file and run it like any other go file.
#### once the program is done you will need to run: $go mod init <name>  
#### you should also remove the HephaestusGo.go from your directory to prevent multiple mains errors in Docker.
#### run the docker file and the rest is history!!!



####
####
## AWS EC2 Instance Configuration 
###
### Updating and installing 
###
#### $sudo yum update -y
#### $sudo yum install -y golang
#### $sudo yum install -y docker
#### $sudo yum install httpd -y
#### $sudo yum install ca-certificates
###
### Configuration
#### $sudo systemctl start httpd
#### $sudo systemctl enable httpd
#### $sudo usermod -a -G apache ec2-user
#### $sudo update-ca-trust force-enable
#### $sudo update-ca-trust
###
### Docker 
#### $sudo systemctl start docker
#### $sudo docker build -t my-golang-app .
#### $sudo docker run -d -p 'AWSport:applicationport' -it --rm --name MywebsiteName my-golang-app
### Docker-compose

#### $sudo curl -L https://github.com/docker/compose/releases/download/1.21.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
#### $sudo chmod +x /usr/local/bin/docker-compose
#### $sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
#### $docker-compose --version