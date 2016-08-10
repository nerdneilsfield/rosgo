// Code generated by ros-gen-go.
// source: TFMessage.msg
// DO NOT EDIT!
package tf2_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
)

type _MsgTFMessage struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgTFMessage) Text() string {
	return t.text
}

func (t *_MsgTFMessage) Name() string {
	return t.name
}

func (t *_MsgTFMessage) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgTFMessage) NewMessage() ros.Message {
	m := new(TFMessage)

	return m
}

var (
	MsgTFMessage = &_MsgTFMessage{
		`geometry_msgs/TransformStamped[] transforms
`,
		"tf2_msgs/TFMessage",
		"996e2194fbbcbb59e036b104dcc9f6fb",
	}
)

type TFMessage struct {
	Transforms []geometry_msgs.TransformStamped
}

func (m *TFMessage) Type() ros.MessageType {
	return MsgTFMessage
}

func (m *TFMessage) Serialize(w io.Writer) (err error) {
	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Transforms)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Transforms {
		if err = ros.SerializeMessageField(w, "geometry_msgs/TransformStamped", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *TFMessage) Deserialize(r io.Reader) (err error) {
	// Transforms
	// Read size little endian
	var size uint32
	if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("cannot read array size for Transforms: %s", err)
	}
	m.Transforms = make([]geometry_msgs.TransformStamped, int(size))
	for i := 0; i < int(size); i++ {
		if err = ros.DeserializeMessageField(r, "geometry_msgs/TransformStamped", &m.Transforms[i]); err != nil {
			return err
		}
	}

	return
}
