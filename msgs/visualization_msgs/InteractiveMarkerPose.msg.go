// Code generated by ros-gen-go.
// source: InteractiveMarkerPose.msg
// DO NOT EDIT!
package visualization_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgInteractiveMarkerPose struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgInteractiveMarkerPose) Text() string {
	return t.text
}

func (t *_MsgInteractiveMarkerPose) Name() string {
	return t.name
}

func (t *_MsgInteractiveMarkerPose) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgInteractiveMarkerPose) NewMessage() ros.Message {
	m := new(InteractiveMarkerPose)

	return m
}

var (
	MsgInteractiveMarkerPose = &_MsgInteractiveMarkerPose{
		`# Time/frame info.
Header header

# Initial pose. Also, defines the pivot point for rotations.
geometry_msgs/Pose pose

# Identifying string. Must be globally unique in
# the topic that this message is sent through.
string name
`,
		"visualization_msgs/InteractiveMarkerPose",
		"848ac869b9cb4c4d85af9cbd2ac09bcf",
	}
)

type InteractiveMarkerPose struct {
	Header std_msgs.Header
	Pose   geometry_msgs.Pose
	Name   string
}

func (m *InteractiveMarkerPose) Type() ros.MessageType {
	return MsgInteractiveMarkerPose
}

func (m *InteractiveMarkerPose) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "geometry_msgs/Pose", &m.Pose); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.Name); err != nil {
		return err
	}

	return
}

func (m *InteractiveMarkerPose) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Pose
	if err = ros.DeserializeMessageField(r, "geometry_msgs/Pose", &m.Pose); err != nil {
		return err
	}

	// Name
	if err = ros.DeserializeMessageField(r, "string", &m.Name); err != nil {
		return err
	}

	return
}
