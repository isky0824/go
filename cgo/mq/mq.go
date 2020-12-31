package main

/*
  #include <stdlib.h>
  typedef void (*HelloCB)(char *);

  void Wrapper( HelloCB cb, char*msg);
*/
import "C"
import (
  "log"
  "unsafe"

	"github.com/streadway/amqp"
)

var(
  conn        *amqp.Connection
	ch_send     *amqp.Channel
	ch_recv     *amqp.Channel
)

//export Init
func Init() bool {
	var err error
	conn, err = amqp.Dial("amqp://test:test123456@47.99.203.79:5672/")
	if err != nil {
		return false
	}

	ch_send, err = conn.Channel()
	if err != nil {
		return false
  }
  
	ch_recv, err = conn.Channel()
	if err != nil {
		return false
	}
  
  return true
}

//export Send
func Send(n *C.char, m *C.char) bool {
  name:= C.GoString( n )
  msg := C.GoString( m )

	//定义并发布消息
	err := ch_send.Publish(
		"",    // exchange
		name,  // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
  )
	if err != nil {
		return false
  } else {
    return true
  }
}

//export Recv
func Recv( n *C.char, cb C.HelloCB ) {
  name := C.GoString( n )

  //声明消息队列
  q, err := ch_recv.QueueDeclare(
    name,  // 队列的名字 name
    false, // 是否持久化durable
    false, // 是否自动删除 delete when unused
    false, // 是否排他性 exclusive
    false, // 是否阻塞 no wait
    nil,   // 额外参数 arguments
  )
  if err != nil {
    log.Println("声明消息队列失败: ", name, "/", err)
  }

  msgs, err := ch_recv.Consume(
    q.Name, // queue
    "",     // consumer
    true,   // auto-ack
    false,  // exclusive
    false,  // no-local
    false,  // no-wait
    nil,    // arguments
  )
  if err != nil {
    log.Println("消费消息失败: ", name, "/", err)
  }

  for d := range msgs {
    cs := C.CString( string( d.Body ) )
    C.Wrapper( cb, cs )
    C.free( unsafe.Pointer( cs ) )
  }
}

//export Release
func Release() {
  if ch_send != nil {
    defer ch_send.Close()
  }

  if ch_recv != nil {
    defer ch_recv.Close()
  }

  if conn != nil {
    defer conn.Close()
  }
}

func main() {

}