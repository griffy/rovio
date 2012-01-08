package rovio

/*
#cgo LDFLAGS: -lrobot_if
#include <robot_if.h>
*/
import "C"
import "unsafe"

// TODO: implement Camera Functions and Rovio-Man API

const (
	BATTERY_OFF = iota
	BATTERY_HOME
	BATTERY_MAX
	WHEEL_FORWARD
	WHEEL_BACKWARD
	HEAD_LOW
	HEAD_MID
	HEAD_HIGH
	NO_SIGNAL
	SIGNAL_WEAK
	SIGNAL_MID
	SIGNAL_STRONG
)

type Robot struct {
	instance *C.robot_if_t
}

func NewRobot(addr string, id C.int) *Robot {
	var instance C.robot_if_t
	caddr := C.CString(addr)
	defer C.free(unsafe.Pointer(caddr))
    if (C.ri_setup(&instance, caddr, id) == C.RI_RESP_FAILURE) {
        return nil
    }
    return &Robot{&instance}
}

func asBool(status C.int) bool {
	if status == C.RI_RESP_SUCCESS {
		return true
	}
	return false
}

func (r *Robot) move(movement C.int, speed C.int) bool {
	return asBool(C.ri_move(r.instance, movement, speed))	
}

func (r *Robot) GoHome() bool {
	return asBool(C.ri_go_home(r.instance))
}

func (r *Robot) Stop() bool {
	return r.move(C.RI_STOP, C.RI_SLOWEST)
}

func (r *Robot) MoveNorth(speed C.int) bool {
	return r.move(C.RI_MOVE_FORWARD, speed)
}

func (r *Robot) MoveSouth(speed C.int) bool {
	return r.move(C.RI_MOVE_BACKWARD, speed)
}

func (r *Robot) MoveWest(speed C.int) bool {
	return r.move(C.RI_MOVE_LEFT, speed)
}

func (r *Robot) MoveEast(speed C.int) bool {
	return r.move(C.RI_MOVE_RIGHT, speed)
}

func (r *Robot) MoveNorthwest(speed C.int) bool {
	return r.move(C.RI_MOVE_FWD_LEFT, speed)
}

func (r *Robot) MoveNortheast(speed C.int) bool {
	return r.move(C.RI_MOVE_FWD_RIGHT, speed)
}

func (r *Robot) MoveSouthwest(speed C.int) bool {
	return r.move(C.RI_MOVE_BACK_LEFT, speed)
}

func (r *Robot) MoveSoutheast(speed C.int) bool {
	return r.move(C.RI_MOVE_BACK_RIGHT, speed)
}

func (r *Robot) TurnLeft(speed C.int) bool {
	return r.move(C.RI_TURN_LEFT, speed)
}

func (r *Robot) TurnRight(speed C.int) bool {
	return r.move(C.RI_TURN_RIGHT, speed)
}

func (r *Robot) TurnLeft20(speed C.int) bool {
	return r.move(C.RI_TURN_LEFT_20DEG, speed)
}

func (r *Robot) TurnRight20(speed C.int) bool {
	return r.move(C.RI_TURN_RIGHT_20DEG, speed)
}

func (r *Robot) LiftHead() bool {
	return r.move(C.RI_HEAD_UP, C.RI_SLOWEST)
}

func (r *Robot) RestHead() bool {
	return r.move(C.RI_HEAD_MIDDLE, C.RI_SLOWEST)
}

func (r *Robot) LowerHead() bool {
	return r.move(C.RI_HEAD_DOWN, C.RI_SLOWEST)
}

func (r *Robot) UpdateSensorCache() bool {
	return asBool(C.ri_update(r.instance))
}

func (r *Robot) ResetEncoderTotals() bool {
	// FIXME: function is incorrectly documented as returning int
	C.ri_reset_state(r.instance)
	return true
}

func (r *Robot) BatteryLife() int {
	life := C.ri_getBattery(r.instance)
	switch life {
	case C.RI_ROBOT_BATTERY_OFF:
		return BATTERY_OFF
	case C.RI_ROBOT_BATTERY_HOME:
		return BATTERY_HOME
	case C.RI_ROBOT_BATTERY_MAX:
		return BATTERY_MAX
	}
	return BATTERY_OFF
}

func (r *Robot) RawWifiStrength() int {
	return int(C.ri_getWifiStrengthRaw(r.instance))
}

func (r *Robot) HeadPosition() int {
	pos := C.ri_getHeadPosition(r.instance)
	switch pos {
	case C.RI_ROBOT_HEAD_LOW:
		return HEAD_LOW
	case C.RI_ROBOT_HEAD_MID:
		return HEAD_MID
	case C.RI_ROBOT_HEAD_HIGH:
		return HEAD_HIGH
	}
	return HEAD_MID
}

func (r *Robot) wheelDirection(wheel C.int) int {
	direction := C.ri_getWheelDirection(r.instance, wheel)
	switch direction {
	case C.RI_WHEEL_FORWARD:
		return WHEEL_FORWARD
	case C.RI_WHEEL_BACKWARD:
		return WHEEL_BACKWARD
	}
	return WHEEL_FORWARD
}

func (r *Robot) wheelMovement(wheel C.int) int {
	return int(C.ri_getWheelEncoder(r.instance, wheel))
}

func (r *Robot) wheelTotal(wheel C.int) int {
	return int(C.ri_getWheelEncoderTotals(r.instance, wheel))
}

func (r *Robot) LeftWheelDirection() int {
	return r.wheelDirection(C.RI_WHEEL_LEFT)
}

func (r *Robot) RightWheelDirection() int {
	return r.wheelDirection(C.RI_WHEEL_RIGHT)
}

func (r *Robot) RearWheelDirection() int {
	return r.wheelDirection(C.RI_WHEEL_REAR)
}

func (r *Robot) LeftWheelMovement() int {
	return r.wheelMovement(C.RI_WHEEL_LEFT)
}

func (r *Robot) RightWheelMovement() int {
	return r.wheelMovement(C.RI_WHEEL_RIGHT)
}

func (r *Robot) RearWheelMovement() int {
	return r.wheelMovement(C.RI_WHEEL_REAR)
}

func (r *Robot) LeftWheelTotal() int {
	return r.wheelTotal(C.RI_WHEEL_LEFT)
}

func (r *Robot) RightWheelTotal() int {
	return r.wheelTotal(C.RI_WHEEL_RIGHT)
}

func (r *Robot) RearWheelTotal() int {
	return r.wheelTotal(C.RI_WHEEL_REAR)
}

func (r *Robot) X() int {
	return int(C.ri_getX(r.instance))
}

func (r *Robot) Y() int {
	return int(C.ri_getY(r.instance))
}

func (r *Robot) Theta() float64 {
	return float64(C.ri_getTheta(r.instance))
}

func (r *Robot) RoomID() int {
	return int(C.ri_getRoomID(r.instance))
}

func (r *Robot) RawSignalStrength() int {
	return int(C.ri_getNavStrengthRaw(r.instance))
}

func (r *Robot) SignalStrength() int {
	strength := C.ri_getNavStrength(r.instance)
	switch strength {
	case C.RI_ROBOT_NAV_SIGNAL_STRONG:
		return SIGNAL_STRONG
	case C.RI_ROBOT_NAV_SIGNAL_MID:
		return SIGNAL_MID
	case C.RI_ROBOT_NAV_SIGNAL_WEAK:
		return SIGNAL_WEAK
	case C.RI_ROBOT_NAV_SIGNAL_NO_SIGNAL:
		return NO_SIGNAL
	}
	return NO_SIGNAL
}

func (r *Robot) TurnOnHeadlight() bool {
	return asBool(C.ri_headlight(r.instance, C.RI_LIGHT_ON))
}

func (r *Robot) TurnOffHeadlight() bool {
	return asBool(C.ri_headlight(r.instance, C.RI_LIGHT_OFF))
}

func (r *Robot) TurnOnInfrared() bool {
	return asBool(C.ri_IR(r.instance, C.RI_LIGHT_ON))
}

func (r *Robot) TurnOffInfrared() bool {
	return asBool(C.ri_IR(r.instance, C.RI_LIGHT_OFF))
}

func (r *Robot) Blocked() bool {
	return bool(C.ri_IR_Detected(r.instance))
}


