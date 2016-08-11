// Code generated by ros-gen-go.
// source: GoalID.msg
// DO NOT EDIT!
package actionlib_msgs

import (
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgGoalID struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgGoalID) Text() string {
	return t.text
}

func (t *_MsgGoalID) Name() string {
	return t.name
}

func (t *_MsgGoalID) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgGoalID) NewMessage() ros.Message {
	m := new(GoalID)

	return m
}

var (
	MsgGoalID = &_MsgGoalID{
		`# The stamp should store the time at which this goal was requested.
# It is used by an action server when it tries to preempt all
# goals that were requested before a certain time
time stamp

# The id provides a way to associate feedback and
# result message with specific goal requests. The id
# specified must be unique.
string id

`,
		"actionlib_msgs/GoalID",
		"80cf13439fb52033034dd028f646e989",
	}
)

type GoalID struct {
	Stamp ros.Time
	ID    string
}

func (m *GoalID) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "time", &m.Stamp); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.ID); err != nil {
		return err
	}

	return
}

func (m *GoalID) Deserialize(r io.Reader) (err error) {
	// Stamp
	if err = ros.DeserializeMessageField(r, "time", &m.Stamp); err != nil {
		return err
	}

	// ID
	if err = ros.DeserializeMessageField(r, "string", &m.ID); err != nil {
		return err
	}

	return
}