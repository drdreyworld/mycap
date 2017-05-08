# mycap
The MySql network traffic analyzer

## Prepare to install

### Ubuntu
```
sudo apt-get install libpcap-dev
```

## Build project
```
cd $GOPATH/src/

git clone https://github.com/kshvakov/mycap.git
cd $GOPATH/src/mycap/

chmod +x ./build.sh && ./build.sh
```

## Run project

### Start daemons
```
cd $GOPATH/src/mycap/bin/

sudo ./tool -process=agent -command start
sudo ./tool -process=server -command start
sudo ./tool -process=web -command start
```

### Stop daemons
```
cd $GOPATH/src/mycap/bin/

sudo ./tool -process=agent -command stop
sudo ./tool -process=server -command stop
sudo ./tool -process=web -command stop
```

### Restart daemons
```
cd $GOPATH/src/mycap/bin/

sudo ./tool -process=agent -command restart
sudo ./tool -process=server -command restart
sudo ./tool -process=web -command restart
```

## Configuration
All configuration located in ./etc/ folder

## Project structure

### Agent
Agent app parse traffic on network device and collect mysql queries.

### Server
Server collect queries from agents by json-rpc protocol.

### Web
Web app gets queries from server and draw them.
By default web it's possible to open web interface at http://localhost:9700/
