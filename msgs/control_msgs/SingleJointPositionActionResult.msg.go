// Code generated by ros-gen-go.
// source: SingleJointPositionActionResult.msg
// DO NOT EDIT!
package control_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgSingleJointPositionActionResult struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgSingleJointPositionActionResult) Text() string {
	return t.text
}

func (t *_MsgSingleJointPositionActionResult) Name() string {
	return t.name
}

func (t *_MsgSingleJointPositionActionResult) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgSingleJointPositionActionResult) NewMessage() ros.Message {
	m := new(SingleJointPositionActionResult)

	return m
}

var (
	MsgSingleJointPositionActionResult = &_MsgSingleJointPositionActionResult{
		`# ====== DO NOT MODIFY! AUTOGENERATED FROM AN ACTION DEFINITION ======

Header header
actionlib_msgs/GoalStatus status
SingleJointPositionResult result
`,
		"control_msgs/SingleJointPositionActionResult",
		"721bbe2e50a90049a415db83086873b8",
	}
)

type SingleJointPositionActionResult struct {
	Header std_msgs.Header
	Status actionlib_msgs.GoalStatus
	Result SingleJointPositionResult
}

func (m *SingleJointPositionActionResult) Type() ros.MessageType {
	return MsgSingleJointPositionActionResult
}

func (m *SingleJointPositionActionResult) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "SingleJointPositionResult", &m.Result); err != nil {
		return err
	}

	return
}

func (m *SingleJointPositionActionResult) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Status
	if err = ros.DeserializeMessageField(r, "actionlib_msgs/GoalStatus", &m.Status); err != nil {
		return err
	}

	// Result
	if err = ros.DeserializeMessageField(r, "SingleJointPositionResult", &m.Result); err != nil {
		return err
	}

	return
}
