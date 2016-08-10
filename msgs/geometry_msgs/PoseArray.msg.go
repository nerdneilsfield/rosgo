// Code generated by ros-gen-go.
// source: PoseArray.msg
// DO NOT EDIT!
package geometry_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgPoseArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgPoseArray) Text() string {
	return t.text
}

func (t *_MsgPoseArray) Name() string {
	return t.name
}

func (t *_MsgPoseArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgPoseArray) NewMessage() ros.Message {
	m := new(PoseArray)

	return m
}

var (
	MsgPoseArray = &_MsgPoseArray{
		`# An array of poses with a header for global reference.

Header header

Pose[] poses
`,
		"geometry_msgs/PoseArray",
		"5f3f794301c7af61b3beab5b9997bb64",
	}
)

type PoseArray struct {
	Header std_msgs.Header
	Poses  []Pose
}

func (m *PoseArray) Type() ros.MessageType {
	return MsgPoseArray
}

func (m *PoseArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Poses)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Poses {
		if err = ros.SerializeMessageField(w, "Pose", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *PoseArray) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Poses
	// Read size little endian
	var size uint32
	if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("cannot read array size for Poses: %s", err)
	}
	m.Poses = make([]Pose, int(size))
	for i := 0; i < int(size); i++ {
		if err = ros.DeserializeMessageField(r, "Pose", &m.Poses[i]); err != nil {
			return err
		}
	}

	return
}
