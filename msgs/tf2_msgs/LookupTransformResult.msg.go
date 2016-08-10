// Code generated by ros-gen-go.
// source: LookupTransformResult.msg
// DO NOT EDIT!
package tf2_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgLookupTransformResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgLookupTransformResult) Text() string {
	return t.text
}

func (t *_MsgLookupTransformResult) Name() string {
	return t.name
}

func (t *_MsgLookupTransformResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgLookupTransformResult) NewMessage() ros.Message {
	m := new(LookupTransformResult)

	return m
}

var (
	MsgLookupTransformResult = &_MsgLookupTransformResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======
geometry_msgs/TransformStamped transform
tf2_msgs/TF2Error error
`,
		"tf2_msgs/LookupTransformResult",
		"1b424b9cebdb0bd3b355dd47a6c2cdca",
	}
)

type LookupTransformResult struct {
	Transform geometry_msgs.TransformStamped
	Error     TF2Error
}

func (m *LookupTransformResult) Type() ros.MessageType {
	return MsgLookupTransformResult
}

func (m *LookupTransformResult) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "geometry_msgs/TransformStamped", &m.Transform); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "tf2_msgs/TF2Error", &m.Error); err != nil {
		return err
	}

	return
}

func (m *LookupTransformResult) Deserialize(r io.Reader) (err error) {
	// Transform
	if err = ros.DeserializeMessageField(r, "geometry_msgs/TransformStamped", &m.Transform); err != nil {
		return err
	}

	// Error
	if err = ros.DeserializeMessageField(r, "tf2_msgs/TF2Error", &m.Error); err != nil {
		return err
	}

	return
}
