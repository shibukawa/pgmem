package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSeqScanEPQ(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_ExecScan(m, l0, int32(753), int32(754))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_read_seq_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v12
		F_LockBuffer(m, v12, int32(2))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v20 < int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v24+(v20^int32(-1))<<(uint(int32(2))%32))))
				v38 = v30
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _consts[1]))
				v38 = v32 + v20<<(uint(int32(13))%32) + int32(-8192)
			}
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+16)))
			v40 = v39 + v38
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
			if v41 == int32(5911) {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
				v47 = v38 + v44&int32(32767)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v47
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(base.Ui32(v49) >> (uint(int32(17)) % 32))
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
				if v53 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = int32(0)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+20)))
					v59 = v57 & int32(64511)
					*(*uint16)(unsafe.Add(mBase, uint32(v56)+20)) = uint16(v59)
					v61 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+20)))
					v64 = v62 | int32(2048)
					*(*uint16)(unsafe.Add(mBase, uint32(v61)+20)) = uint16(v64)
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					F_MarkBufferDirtyHint(m, v66, int32(1))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						v71 = v70
						v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
						m.G0 = v9 + int32(16)
						return v71 + v73
					}
				} else {
					v71 = v47
					v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
					m.G0 = v9 + int32(16)
					return v71 + v73
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v84
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v83 + int32(4)
					F_errmsg_internal(m, int32(498666), v9)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(488282), int32(1205), int32(374804))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
