# A Restful Api Server By Golang Language

##Summary

At first, I'm only want to have a ejabberd+riak IM server. So I set up it on my vps.
And next, I want to modify some models of ejabberd. Like user info, relationship and so on. So, this api is a DB service. If you like it or you have some questions, Feel free to contact me

##Include Package

* [beego](http://beego.me/) -- A Restful framework
* [mgo](http://labix.org/mgo) -- MongoDB drive
* [redigo](https://github.com/garyburd/redigo) -- Redis drive
* [neoism](https://github.com/jmcvetta/neoism) -- Neo4j drive

##Application & Package Version

* Golang 1.6
* bee 1.4.1
* beego 1.6.1
* mongodb 3.0
* mgo v2
* redigo
* neoism v1.3.1

##Framework Design

Follow Beego. I'm just add modules for business logic layer. Package Models is only for DB models. And I'm writing package tools.Have some tools classes in it. Of course, You can use any one of them alone.I'm store user login infomation & file chunks by redis. And store other infomation in MongoDB.
It only deals with the interface or map type. So you can deals object & json in package modules.And that it maintains the reuse function.

* Beego design: request → routers → controllers → models
* My design: request → controllers → modules  → models & tools

##Configuration instructions

	appname : your app name
	httpport : http server port
	runmode : run modes like dev & pro etc
	autorender : close the template rendering, in this is false
	copyrequestbody : Read body information, in this is true
	EnableDocs : enable Docs or not
	mongohost = your mongodb host
	mongoport = your mongodb port, default is 27017
	mongodbname = your mongodb dbname
	filecache : IF you need cache file on desk, you need write directory path in here
	redisnetwork : redis net work type TCP/UDP
	redishost = your redis host
	redisport = your redis port, default is 6379
	redispwd = your redis authentication password

##Usage

	cd YOUR_CODE_PWD
	export GOPATH=$pwd #Set this path of your gopath
	go get github.com/astaxie/beego #need build it
	go get gopkg.in/mgo.v2 #need build it
	cd src & git clone https://github.com/littletwolee/mongoapi.git
	cd conf & modify your configuration
	cd ../ & bee run
