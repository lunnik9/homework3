package main

import (
	"fmt"
	"strconv"
)

func main() {
	conp := ConnectionPool{}

	conp.init(5)

	userid1 := "111"
	userid2 := "222"
	userid3 := "333"

	conp.getConnection(userid1)
	conp.getConnection(userid2)
	conp.getConnection(userid3)

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
	for i := 0; i < size; i++ {
		con := Connection{}
		con.id = strconv.Itoa(i)
		con.timeout = "50"
		c.connections[con] = ""
	}
}
