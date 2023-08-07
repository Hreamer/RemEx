# What is RemEx? 
RemEx is a Remote(Rem) Code Compilation and Execution(Ex) Tool  
This tool will compile code you send it and than return the output  
To ensure safety of the machine that is running the server side of the tool, as remote code execution is one of the holy grails of security flaws, the Server.go file runs inside a docker container(not yet implemented)  
The Client.go file just sends the specified file over a HTTP POST request.  

# How to use 
Clone the repo and compile the Client.go file and the Server.go file. Run the server however you want but it listens on port 8080. Than run the Client.go file with the correct flags set and it should return the output of stdout as if you built the file locally. 

# Limitations 
Currently RemEx supports sending a single file

# Languages Supported 
Go. use "-l go" to send go content


## Applications 
This project has specific applications in the embedded world where you want to offload the compilation and running of a program. I can see a world where this acts as a microservice for a website like LeetCode as well. 

## What Inspired Me?  
Just wanted to make a personal project that involved Golang as I wanted to learn the language and I had the idea for this project ever since I heard about the Log4J Exploit. Name of the project is in honor of one of my best friends Remo. 
