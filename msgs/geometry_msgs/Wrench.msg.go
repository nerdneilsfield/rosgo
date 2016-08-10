// Code generated by ros-gen-go.
// source: Wrench.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgWrench struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgWrench) Text() string {
	return t.text
}

func (t *_MsgWrench) Name() string {
	return t.name
}

func (t *_MsgWrench) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgWrench) NewMessage() ros.Message {
	m := new(Wrench)

	return m
}

var (
	MsgWrench = &_MsgWrench{
		`# This represents force in free space, separated into
# its linear and angular parts.
Vector3  force
Vector3  torque
`,
		"geometry_msgs/Wrench",
		"2aae87faaa553ae28e07e684016d765c",
	}
)

type Wrench struct {
	Force  Vector3
	Torque Vector3
}

func (m *Wrench) Type() ros.MessageType {
	return MsgWrench
}

func (m *Wrench) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Vector3", &m.Force); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "Vector3", &m.Torque); err != nil {
		return err
	}

	return
}

func (m *Wrench) Deserialize(r io.Reader) (err error) {
	// Force
	if err = ros.DeserializeMessageField(r, "Vector3", &m.Force); err != nil {
		return err
	}

	// Torque
	if err = ros.DeserializeMessageField(r, "Vector3", &m.Torque); err != nil {
		return err
	}

	return
}
