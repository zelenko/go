// Package sess manages user registrations, records, and sessions.
// Copyright © 2018 Zelenko. All rights reserved.
package sess

import (
	"bufio"
	b64 "encoding/base64"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type session struct {
	username   string
	browser    string
	ip         string
	lastActive time.Time
}

type user struct {
	password       string
	dateRegistered time.Time
}

// manage unexported type holds sessions and user data
type manage struct {
	activeSessions  map[string]session
	registeredUsers map[string]user
}

var mgr manage

// SManager exported type is used by methods
type SManager string

// NewManager creates new session Manager
func NewManager() *SManager {
	mgr = manage{
		activeSessions:  map[string]session{},
		registeredUsers: map[string]user{},
	}
	//sm := new(SManager)
	sm := SManager("")
	sm.open()
	return &sm
}

// RegisterUser creates a new user
func (sm *SManager) RegisterUser(un, pw string) {
	newUser := user{password: pw, dateRegistered: time.Now()}
	mgr.registeredUsers[un] = newUser
}

// ValidateUser checks if user is registered
func (sm *SManager) ValidateUser(id string) bool {
	_, ok := mgr.registeredUsers[id]
	return ok
}

// LoginSession return session id
func (sm *SManager) LoginSession(un, br, ipAddr string) string {
	s := &session{username: un, browser: br, ip: ipAddr, lastActive: time.Now()}
	id := ("random session id") // need to generate this
	mgr.activeSessions[id] = *s
	return id
}

// UpdateSession sets lastActive date and time
func (sm *SManager) UpdateSession(id string) {
	s := mgr.activeSessions[id]
	s.lastActive = time.Now()
	mgr.activeSessions[id] = s
}

// ShowSessions returns slice of sessions
func (sm *SManager) ShowSessions(userID string) []string {
	var str []string
	for _, s := range mgr.activeSessions {
		if s.username == userID {
			str = append(str, "Name: "+s.username+" IP: "+s.ip+" Time: "+s.lastActive.String())
		}
	}
	return str
}

// ShowUsers returns slice of users
func (sm *SManager) ShowUsers() []string {
	var str []string
	for k, s := range mgr.registeredUsers {
		str = append(str, k+"\tRegistered: "+s.dateRegistered.Format("2006-01-02 03:04pm"))
	}
	return str
}

// Close method saves all user and session records into text file.
// These saved records are loaded into memory next time server runs.
func (sm *SManager) Close() {
	// save session data into sessions.txt file
	// save user data into users.txt file
	usersFile, err := os.Create("users.txt")
	if err != nil {
		log.Println("Cannot create new users.txt file", err)
		return
	}
	defer usersFile.Close()

	usersFile.WriteString("ID\tDate Registered\tHash\n") // First line is Title line
	// write each line to file
	for k, s := range mgr.registeredUsers {
		line := k + "\t" + s.dateRegistered.Format("2006-01-02_03-04-05pm") + "\t" + s.password + "\n"
		usersFile.WriteString(line)
	}
}

// open reads records from text file into memory
func (sm *SManager) open() {
	// open file and read contents to mgr
	// if file does not exist, do nothing
	filename := "users.txt"

	// check if file exists.  If not, then stop.
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// log.Println("file DOES NOT exist")
		return // file DOES NOT exist, so do NOT load anything
	}

	// open file
	f, err := os.Open(filename)
	if err != nil {
		log.Println("error opening file", err)
		return
	}
	defer f.Close()

	r := bufio.NewReaderSize(f, 4*1024)
	// skip first title line
	r.ReadLine()
	line, isPrefix, err := r.ReadLine()
	// read each additional line
	for err == nil && !isPrefix {
		s := string(line)
		if len(s) < 1 { // if line is blank, then stop reading file
			break
		}

		t := strings.Split(s, "\t")
		userName, dateRaw, pass := t[0], t[1], t[2]

		layout := "2006-01-02_15-04-05pm"
		date, err := time.Parse(layout, dateRaw)
		if err != nil {
			log.Println("error converting date: " + err.Error())
		}

		// create a user record and add to map
		newUser := user{password: pass, dateRegistered: date}
		mgr.registeredUsers[userName] = newUser

		// read next line
		line, isPrefix, err = r.ReadLine()
	}

	if isPrefix {
		log.Println("buffer size to small")
		return
	}
	if err != io.EOF && err != nil {
		log.Println("not EOF", err)
		return
	}
}

// RandomSession generates new random session ID
func (sm *SManager) RandomSession() string {
	out, _ := generatePasswordHash(getSessionID())
	return out
}

func getSessionID() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 20)
	for i := range b {
		b[i] = rune(rand.Intn(99))
	}
	return string(b)
}

func generatePasswordHash(in string) (string, error) {

	// generate password hash
	password := []byte(in + "generate random password hash")
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	pass64 := b64.StdEncoding.EncodeToString([]byte(hashedPassword))
	// decodedString, _ := b64.StdEncoding.DecodeString(encodedString)
	return pass64, nil
}
