package main

import (
	"fmt"
	"strconv"
)

func main() {
	conp := ConnectionPool{}

	userid1 := "111"
	userid2 := "222"
	userid3 := "333"
	userid4 := "444"
	userid5 := "555"
	userid6 := "666"

	conp.getConnection(userid1)
	fmt.Println(conp)
	conp.getConnection(userid2)
	conp.getConnection(userid3)
	conp.getConnection(userid4)
	conp.getConnection(userid5)
	fmt.Println(conp)
	conp.getConnection(userid6)
	fmt.Println(conp)
	fmt.Println(conp.returnConnection(userid1))
	fmt.Println(conp)

}

type ConnectionPool struct {
	connections map[Connection]string
}

type Connection struct {
	id      string
	timeout string
}

func (c *ConnectionPool) getConnection(userid string) Connection {

	if len(c.connections) == 0 {
		c.init(1)
		c.getConnection(userid)
	}
	trigger := 0

	for _, uid := range c.connections {
		if uid != "" {
			trigger++
		}
	}

	if trigger == len(c.connections) {
		c.add(1)
		c.getConnection(userid)
	}

	con := Connection{}
	for conn, uid := range c.connections {
		if uid == "" {
			c.connections[conn] = userid
			con = conn
			break
		}
	}
	return con
}

func (c *ConnectionPool) returnConnection(userid string) Connection {
	con := Connection{}
	for conn, uid := range c.connections {
		if uid == userid {
			c.connections[conn] = ""
			con = conn
			break
		}
	}
	return con
}

func (c *ConnectionPool) init(size int) {
	c.connections = make(map[Connection]string)
	c.add(size)
}

func (c *ConnectionPool) add(size int) {
	for i := 0; i < size; i++ {
		con := Connection{}
		con.id = strconv.Itoa(i + len(c.connections))
		con.timeout = "50"
		c.connections[con] = ""
	}
}
