# A Restful Api Server For MongoDB By Golang Language

##Summary

At first, I'm only want to have a ejabberd+riak IM server. So I set up it on my vps.
And next, I want to modify some models of ejabberd. Like user info, relationship and so on. So, this api is a DB service. If you like it or you have some questions, Feel free to contact me

##Include Package

* [beego](http://beego.me/) -- A Restful framework
* [mgo](http://labix.org/mgo) -- MongoDB drive

##Application & Package Version

* Golang 1.6
* bee 1.4.1
* beego 1.6.1
* mongodb 3.0
* mgo v2

##Framework Design

Follow Beego. I'm just add modules for business logic layer. Package Models is only for DB models. And I'm writing package tools.Have some tools classes in it. Of course, You can use any one of them alone.In mongohelpes, there have show one, show all, inster, pull, push, update & delete. It only deals with the interface or map type. So you can deals object & json in package modules.So that it maintains the reuse function.

* Beego design: request --> routers --> controllers --> models
* My design: request --> controllers --> modules  --> models

##Usage
`
1. cd YOUR_CODE_PWD</br>
`
`
2. export GOPATH=$pwd #Set this path of your gopath
`
3. go get github.com/astaxie/beego \<br>  
4. go get gopkg.in/mgo.v2\<br>  
