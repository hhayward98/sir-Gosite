# Æsir Gosite
## Hephaestus-Go
###
#### Hephaestus, the creator of Zeus’ thunderbolt, Poseidon's trident, and more.
#### The Hephaestus-Go app is a program that auto-generates the Golang environment for rapid deployment of web apps.
#### The goal is to minimize setup time on a new machine or virtual machine, such as AWS ec2 instances. 
#### Users will be able to customize the environment to better suit their needs.
#### Personally, I use VIM on Linux, and I waste time setting up the environments for launching web apps.
#### with Hephaestus-Go, I just run one program and all the tedious, redundant code is done!!!
###
## Updates
###
#### As time goes on, I will be updating and adding more complexity and customization.
###
## Using Hephaestus-Go
###
#### Download HephaestusGo.go file and run it like any other go file. (only option 1 works right now)
#### Select option 1 and enter the number of html pages you want.
#### once the program is done you will need to run: $go mod init <name>  
#### you should also remove the HephaestusGo.go from your directory to prevent multiple mains errors in Docker.
#### run the docker file and the rest is history!!!



### For an Example of the final product visit http://52.205.255.136:8080/Page0 (Swap between ports 8080 amd 8088) 
#### pages http://52.205.255.136:8088/Page1, http://52.205.255.136:8080/Page1, http://52.205.255.136:8088/index
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