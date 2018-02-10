package chatRoom

// SOURCE: https://github.com/coolspeed/century/blob/master/century.go

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

// Session -
type Session struct {
	sid                       int
	chatRoom                  *ChatRoom
	connection                net.Conn
	incoming                  chan string
	outgoing                  chan string
	reader                    *bufio.Reader
	writer                    *bufio.Writer
	killRoomConnGoroutine     chan bool
	killSocketReaderGoroutine chan bool
	killSocketWriterGoroutine chan bool
	sessionMutex              *sync.Mutex
}

// NewSession -
func NewSession(sid int, chatRoom *ChatRoom, connection net.Conn) *Session {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	session := &Session{
		sid:                       sid,
		chatRoom:                  chatRoom,
		connection:                connection,
		incoming:                  make(chan string),
		outgoing:                  make(chan string),
		reader:                    reader,
		writer:                    writer,
		killRoomConnGoroutine:     make(chan bool),
		killSocketReaderGoroutine: make(chan bool),
		killSocketWriterGoroutine: make(chan bool),
		sessionMutex:              &sync.Mutex{},
	}
	fmt.Println("A new session created. sid=", sid)
	return session
}

// Read -
func (session *Session) Read() {
	for {
		select {
		case <-session.killSocketReaderGoroutine:
			return
		default:
			line, err := session.reader.ReadString('\n')
			if err != nil {
				// EOF? yes: disconnected
				// Judge. if true: LeaveAndDelete -- pop session from chatRoom, and delete session
				if err == io.EOF {
					fmt.Println("Client disconnected. Destroy session, sid=", session.sid)
					session.LeaveAndDelete()
				}

				// else:
				fmt.Println("bufio.reader.ReadString failed.")
				fmt.Println(err)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			session.incoming <- line
		}
	}
}

// Write -
func (session *Session) Write() {
	for {
		select {
		case <-session.killSocketWriterGoroutine:
			return
		case data := <-session.outgoing:
			session.writer.WriteString(data)
			session.writer.Flush()
		}
	}
}

// Listen -
func (session *Session) Listen() {
	go session.Read()
	go session.Write()
}

// LeaveAndDelete -
func (session *Session) LeaveAndDelete() {
	// leave

	// https://github.com/golang/go/issues/13675
	chatRoom := *session.chatRoom
	sid := session.sid
	chatRoom.roomMutex.Lock()
	defer chatRoom.roomMutex.Unlock()
	delete(chatRoom.sessions, sid)

	// delete

	session.sessionMutex.Lock()
	defer session.sessionMutex.Unlock()

	// release resources

	// resouce: socket reader goroutine & socket writer goroutine
	session.killSocketReaderGoroutine <- true
	session.killSocketWriterGoroutine <- true

	// resource: reader & writer
	session.reader = nil
	session.writer = nil

	// resource: socket conection
	session.connection.Close()
	session.connection = nil

	// resource: RoomConnGoroutine
	session.killRoomConnGoroutine <- true

	// resource: connection to chatRoom
	// "many in, one out" 형태의 'chatRoom'의(!) channel 이기에 지워야할 채널(RoomConn)이 사실은 존재하지 않는다. (위에서) 해당 goroutine 지워줬으니 끝.
}

// ChatRoom -
type ChatRoom struct {
	sessions  map[int]*Session
	lastSid   int
	entrance  chan net.Conn
	incoming  chan string
	outgoing  chan string
	roomMutex sync.Mutex
}

// NewChatRoom -
func NewChatRoom() *ChatRoom {
	chatRoom := &ChatRoom{
		sessions: make(map[int]*Session),
		lastSid:  -1,
		entrance: make(chan net.Conn),
		incoming: make(chan string),
		outgoing: make(chan string),
		//		roomMutex:
	}
	fmt.Println("A new ChatRoom created.")
	return chatRoom
}

// Broadcast -
func (chatRoom *ChatRoom) Broadcast(data string) {
	for _, session := range chatRoom.sessions {
		session.outgoing <- data
	}
}

// Join -
func (chatRoom *ChatRoom) Join(connection net.Conn) {
	chatRoom.roomMutex.Lock()
	defer chatRoom.roomMutex.Unlock()
	newSessionID := chatRoom.lastSid + 1
	chatRoom.lastSid = newSessionID

	session := NewSession(newSessionID, chatRoom, connection)
	session.Listen()
	fmt.Println("session started listening.")

	_, keyExist := chatRoom.sessions[newSessionID]
	if !keyExist {
		chatRoom.sessions[newSessionID] = session
	}

	go func() { // goroutine for roomConn writer
		for {
			select {
			case <-session.killRoomConnGoroutine:
				return
			case data := <-session.incoming:
				chatRoom.incoming <- data
			}
		}
	}()
}
