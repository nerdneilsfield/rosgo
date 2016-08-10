// Code generated by ros-gen-go.
// source: BatteryState.msg
// DO NOT EDIT!
package sensor_msgs

import (
	"io"
	"encoding/binary"
	"fmt"

	"github.com/ppg/rosgo/ros"
	"github.com/ppg/rosgo/msgs/std_msgs"
)

type _MsgBatteryState struct {
	text   string
	name   string
	md5sum string
}

func (t *_MsgBatteryState) Text() string {
	return t.text
}

func (t *_MsgBatteryState) Name() string {
	return t.name
}

func (t *_MsgBatteryState) MD5Sum() string {
	return t.md5sum
}

func (t *_MsgBatteryState) NewMessage() ros.Message {
	m := new(BatteryState)

	return m
}

var (
	MsgBatteryState = &_MsgBatteryState{
		`
# Constants are chosen to match the enums in the linux kernel
# defined in include/linux/power_supply.h as of version 3.7
# The one difference is for style reasons the constants are
# all uppercase not mixed case.

# Power supply status constants
uint8 POWER_SUPPLY_STATUS_UNKNOWN = 0
uint8 POWER_SUPPLY_STATUS_CHARGING = 1
uint8 POWER_SUPPLY_STATUS_DISCHARGING = 2
uint8 POWER_SUPPLY_STATUS_NOT_CHARGING = 3
uint8 POWER_SUPPLY_STATUS_FULL = 4

# Power supply health constants
uint8 POWER_SUPPLY_HEALTH_UNKNOWN = 0
uint8 POWER_SUPPLY_HEALTH_GOOD = 1
uint8 POWER_SUPPLY_HEALTH_OVERHEAT = 2
uint8 POWER_SUPPLY_HEALTH_DEAD = 3
uint8 POWER_SUPPLY_HEALTH_OVERVOLTAGE = 4
uint8 POWER_SUPPLY_HEALTH_UNSPEC_FAILURE = 5
uint8 POWER_SUPPLY_HEALTH_COLD = 6
uint8 POWER_SUPPLY_HEALTH_WATCHDOG_TIMER_EXPIRE = 7
uint8 POWER_SUPPLY_HEALTH_SAFETY_TIMER_EXPIRE = 8

# Power supply technology (chemistry) constants
uint8 POWER_SUPPLY_TECHNOLOGY_UNKNOWN = 0
uint8 POWER_SUPPLY_TECHNOLOGY_NIMH = 1
uint8 POWER_SUPPLY_TECHNOLOGY_LION = 2
uint8 POWER_SUPPLY_TECHNOLOGY_LIPO = 3
uint8 POWER_SUPPLY_TECHNOLOGY_LIFE = 4
uint8 POWER_SUPPLY_TECHNOLOGY_NICD = 5
uint8 POWER_SUPPLY_TECHNOLOGY_LIMN = 6

Header  header
float32 voltage          # Voltage in Volts (Mandatory)
float32 current          # Negative when discharging (A)  (If unmeasured NaN)
float32 charge           # Current charge in Ah  (If unmeasured NaN)
float32 capacity         # Capacity in Ah (last full capacity)  (If unmeasured NaN)
float32 design_capacity  # Capacity in Ah (design capacity)  (If unmeasured NaN)
float32 percentage       # Charge percentage on 0 to 1 range  (If unmeasured NaN)
uint8   power_supply_status     # The charging status as reported. Values defined above
uint8   power_supply_health     # The battery health metric. Values defined above
uint8   power_supply_technology # The battery chemistry. Values defined above
bool    present          # True if the battery is present

float32[] cell_voltage   # An array of individual cell voltages for each cell in the pack
                         # If individual voltages unknown but number of cells known set each to NaN
string location          # The location into which the battery is inserted. (slot number or plug)
string serial_number     # The best approximation of the battery serial number
`,
		"sensor_msgs/BatteryState",
		"56ecc2d1da33cbae60aa83fa76d21ae1",
	}
)

type BatteryState struct {
	Header                std_msgs.Header
	Voltage               float32
	Current               float32
	Charge                float32
	Capacity              float32
	DesignCapacity        float32
	Percentage            float32
	PowerSupplyStatus     uint8
	PowerSupplyHealth     uint8
	PowerSupplyTechnology uint8
	Present               bool
	CellVoltage           []float32
	Location              string
	SerialNumber          string
}

func (m *BatteryState) Type() ros.MessageType {
	return MsgBatteryState
}

func (m *BatteryState) Serialize(w io.Writer) (err error) {
	if err = ros.SerializeMessageField(w, "Header", &m.Header); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Voltage); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Current); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Charge); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Capacity); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.DesignCapacity); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "float32", &m.Percentage); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.PowerSupplyStatus); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.PowerSupplyHealth); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "uint8", &m.PowerSupplyTechnology); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "bool", &m.Present); err != nil {
		return err
	}

	// Write size little endian
	err = binary.Write(w, binary.LittleEndian, uint32(len(m.CellVoltage)))
	if err != nil {
		return fmt.Errorf("could not write array length: %s", err)
	}
	for _, elem := range m.CellVoltage {
		if err = ros.SerializeMessageField(w, "float32", &elem); err != nil {
			return err
		}
	}

	if err = ros.SerializeMessageField(w, "string", &m.Location); err != nil {
		return err
	}

	if err = ros.SerializeMessageField(w, "string", &m.SerialNumber); err != nil {
		return err
	}

	return
}

func (m *BatteryState) Deserialize(r io.Reader) (err error) {
	// Header
	if err = ros.DeserializeMessageField(r, "Header", &m.Header); err != nil {
		return err
	}

	// Voltage
	if err = ros.DeserializeMessageField(r, "float32", &m.Voltage); err != nil {
		return err
	}

	// Current
	if err = ros.DeserializeMessageField(r, "float32", &m.Current); err != nil {
		return err
	}

	// Charge
	if err = ros.DeserializeMessageField(r, "float32", &m.Charge); err != nil {
		return err
	}

	// Capacity
	if err = ros.DeserializeMessageField(r, "float32", &m.Capacity); err != nil {
		return err
	}

	// DesignCapacity
	if err = ros.DeserializeMessageField(r, "float32", &m.DesignCapacity); err != nil {
		return err
	}

	// Percentage
	if err = ros.DeserializeMessageField(r, "float32", &m.Percentage); err != nil {
		return err
	}

	// PowerSupplyStatus
	if err = ros.DeserializeMessageField(r, "uint8", &m.PowerSupplyStatus); err != nil {
		return err
	}

	// PowerSupplyHealth
	if err = ros.DeserializeMessageField(r, "uint8", &m.PowerSupplyHealth); err != nil {
		return err
	}

	// PowerSupplyTechnology
	if err = ros.DeserializeMessageField(r, "uint8", &m.PowerSupplyTechnology); err != nil {
		return err
	}

	// Present
	if err = ros.DeserializeMessageField(r, "bool", &m.Present); err != nil {
		return err
	}

	// CellVoltage
	// Read size little endian
	var size uint32
	if err = binary.Read(r, binary.LittleEndian, &size); err != nil {
		return fmt.Errorf("cannot read array size for CellVoltage: %s", err)
	}
	m.CellVoltage = make([]float32, int(size))
	for i := 0; i < int(size); i++ {
		if err = ros.DeserializeMessageField(r, "float32", &m.CellVoltage[i]); err != nil {
			return err
		}
	}

	// Location
	if err = ros.DeserializeMessageField(r, "string", &m.Location); err != nil {
		return err
	}

	// SerialNumber
	if err = ros.DeserializeMessageField(r, "string", &m.SerialNumber); err != nil {
		return err
	}

	return
}
