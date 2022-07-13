// Package email to sending mail use smtp server.
package email

import (
	"bytes"
	"github.com/skvdmt/f"
	"log"
	"net/smtp"
	"strconv"
	"strings"
)

// Server is smtp server
type Server struct {
	host     string
	port     uint16
	username string
	password string
}

// Set is configuration server setting host and port
func (serv *Server) Set(name string, value interface{}) {
	switch name {
	case "host":
		serv.host = value.(string)
	case "port":
		serv.port = uint16(value.(int))
	case "username":
		serv.username = value.(string)
	case "password":
		serv.password = value.(string)
	default:
		log.Fatalln(name + " is unknown property")
	}
}

// Conn dial to smtp server without auth
func (serv *Server) Conn() *smtp.Client {
	// default settings
	if serv.host == "" {
		serv.host = "localhost"
	}
	if serv.port == 0 {
		serv.port = 25
	}
	// connection
	client, err := smtp.Dial(serv.host + ":" + strconv.FormatUint(uint64(serv.port), 10))
	f.Check(err)
	return client
}

// Auth on server with username and password
func (serv *Server) Auth() smtp.Auth {
	return smtp.PlainAuth("", serv.username, serv.password, serv.host)
}

// Letter is mail struct
type Letter struct {
	to          string
	contentType string
	subject     string
	body        string
}

// Set letter headers data and body mail
func (let *Letter) Set(name string, value string) {
	switch strings.ToLower(name) {
	case "to":
		let.to = value
	case "content-type":
		let.contentType = value
	case "subject":
		let.subject = value
	case "body":
		let.body = value
	default:
		log.Fatalln(name + " is unknown property")
	}
}

// Send sending email with or without auth
func Send(serv *Server, let *Letter, conn interface{}) {
	switch conn.(type) {
	case smtp.Auth:
		err := smtp.SendMail(
			serv.host+":"+strconv.FormatUint(uint64(serv.port), 10),
			conn.(smtp.Auth),
			serv.username,
			[]string{let.to},
			[]byte("Content-type: "+let.contentType+"\r\n"+
				"To: "+let.to+"\r\n"+
				"Subject: "+let.subject+"\r\n"+
				"\r\n"+
				let.body+"\r\n"))
		f.Check(err)
	case *smtp.Client:
		conn := conn.(*smtp.Client)
		err := conn.Mail(serv.username)
		f.Check(err)
		err = conn.Rcpt(let.to)
		f.Check(err)
		wc, err := conn.Data()
		f.Check(err)
		defer wc.Close()
		buf := bytes.NewBufferString("Content-type: " + let.contentType + "\r\n" +
			"To: " + let.to + "\r\n" +
			"Subject: " + let.subject + "\r\n" +
			"\r\n" +
			let.body + "\r\n")
		_, err = buf.WriteTo(wc)
		f.Check(err)
	}
}
