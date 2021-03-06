// Code generated by ros-gen-go.
// source: ByteMultiArray.msg
// DO NOT EDIT!
package std_msgs

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/ppg/rosgo/ros"
)

type _MsgByteMultiArray struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgByteMultiArray) Text() string {
	return t.text
}

func (t *_MsgByteMultiArray) Name() string {
	return t.name
}

func (t *_MsgByteMultiArray) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgByteMultiArray) NewMessage() ros.Message {
	m := new(ByteMultiArray)

	return m
}

var (
	MsgByteMultiArray = &_MsgByteMultiArray{
		`# Please look at the MultiArrayLayout message definition for
# documentation on all multiarrays.

MultiArrayLayout  layout        # specification of data layout
byte[]            data          # array of data

`,
		"std_msgs/ByteMultiArray",
		"3c0a650df9a2e34adaa4288635cd5d87",
	}
)

type ByteMultiArray struct {
	Layout MultiArrayLayout
	Data   []int8
}

func (m *ByteMultiArray) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.Data)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.Data {
		if err = ros.SerializeMessageField(w, "byte", &elem); err != nil {
			return err
		}
	}

	return
}

func (m *ByteMultiArray) Deserialize(r io.Reader) (err error) {
	// Layout
	if err = ros.DeserializeMessageField(r, "MultiArrayLayout", &m.Layout); err != nil {
		return err
	}

	// Data
	{
		// Read size little endian
		var size uint32
		if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
			return fmt.Errorf("cannot read array size for Data: %s", err)
		}
		m.Data = make([]int8, int(size))
		for i := 0; i < int(size); i++ {
			if err = ros.DeserializeMessageField(r, "byte", &m.Data[i]); err != nil {
				return err
			}
		}
	}

	return
}
