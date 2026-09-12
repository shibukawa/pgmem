package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_StartupProcTriggerHandler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	*(*int32)(unsafe.Add(mBase, _consts[632])) = int32(1)
	F_WakeupRecovery(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_begin_startup_progress_phase(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	var v28 int64
	_ = v28
	var v36 int64
	_ = v36
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, _consts[633]))
	if v5 == int32(0) {
		return
	} else {
		F_disable_timeout(m, int32(12))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v12 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[634])) = v12
			v15 = *(*int32)(unsafe.Add(mBase, _consts[633]))
			if v15 == v12 {
				return
			} else {
				v22 = m.G0
				v23 = int32(16)
				v24 = v22 - v23
				m.G0 = v24
				F___gettimeofday(m, v24)
				mBase = m.M
				v27 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
				v28 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
				m.G0 = v24 + v23
				v36 = v28 + v27*int64(1000000) - int64(946684800000000)
				*(*int64)(unsafe.Add(mBase, _consts[635])) = v36
				v39 = *(*int32)(unsafe.Add(mBase, _consts[633]))
				*(*int32)(unsafe.Add(mBase, _consts[636])) = int32(0)
				v51 = m.G0
				v52 = int32(16)
				v53 = v51 - v52
				m.G0 = v53
				F___gettimeofday(m, v53)
				mBase = m.M
				v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
				v57 = int64(*(*int32)(unsafe.Add(mBase, uint32(v53)+8)))
				m.G0 = v53 + v52
				v65 = v57 + v56*int64(1000000) - int64(946684800000000)
				F_enable_timeout(m, int32(12), v65, v36+base.I64_extend_i32_s(v39)*int64(1000), v39)
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_schedule_alarm(m, v65)
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_process_startup_packet_die(m *base.Module, l0 int32) {
	F__Exit(m, int32(1))
	base.Wasm_trap_unreachable()
	for {
	}
}
