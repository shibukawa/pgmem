package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetOldestMultiXactId(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	v1 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[0]))
	v15 = F_LWLockAcquire(m, v11+int32(1664), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[1]))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		v23 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[2]))
		v25 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[3]))
		v26 = v23 + v25
		if v26 <= int32(0) {
			v86 = v21
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[4]))
			if v26 != int32(1) {
				v37 = v21
				v38 = v1
				v43 = v1
				for {
					v48 = v30 + v38<<(uint(int32(2))%32)
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
					if v50-v37 < int32(0) {
						v54 = v50
					} else {
						v54 = v37
					}
					if v50 != 0 {
						v55 = v54
					} else {
						v55 = v37
					}
					if v49-v55 < int32(0) {
						v59 = v49
					} else {
						v59 = v55
					}
					if v49 != 0 {
						v60 = v59
					} else {
						v60 = v55
					}
					v61 = int32(2)
					v62 = v38 + v61
					v64 = v43 + v61
					if v64 != v26&int32(2147483646) {
						v37 = v60
						v38 = v62
						v43 = v64
						continue
					} else {
						break
					}
					break
				}
				if v26&int32(1) == int32(0) {
					v86 = v60
				} else {
					v68 = v60
					v69 = v62
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v30+v69<<(uint(int32(2))%32))))
					if v80-v68 < int32(0) {
						v84 = v80
					} else {
						v84 = v68
					}
					if v80 != 0 {
						v85 = v84
					} else {
						v85 = v68
					}
					v86 = v85
				}
			} else {
				v68 = v21
				v69 = v1
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v30+v69<<(uint(int32(2))%32))))
				if v80-v68 < int32(0) {
					v84 = v80
				} else {
					v84 = v68
				}
				if v80 != 0 {
					v85 = v84
				} else {
					v85 = v68
				}
				v86 = v85
			}
		}
		if v23 <= int32(0) {
			v157 = v86
		} else {
			v98 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[5]))
			if v23 == int32(1) {
				v139 = v86
				v140 = int32(0)
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v98+v140<<(uint(int32(2))%32))))
				if v151-v139 < int32(0) {
					v155 = v151
				} else {
					v155 = v139
				}
				if v151 != 0 {
					v156 = v155
				} else {
					v156 = v139
				}
				v157 = v156
			} else {
				v106 = int32(0)
				v108 = v86
				v109 = v106
				v114 = v106
				for {
					v119 = v98 + v109<<(uint(int32(2))%32)
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+4))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
					if v121-v108 < int32(0) {
						v125 = v121
					} else {
						v125 = v108
					}
					if v121 != 0 {
						v126 = v125
					} else {
						v126 = v108
					}
					if v120-v126 < int32(0) {
						v130 = v120
					} else {
						v130 = v126
					}
					if v120 != 0 {
						v131 = v130
					} else {
						v131 = v126
					}
					v132 = int32(2)
					v133 = v109 + v132
					v135 = v114 + v132
					if v135 != v23&int32(2147483646) {
						v108 = v131
						v109 = v133
						v114 = v135
						continue
					} else {
						break
					}
					break
				}
				if v23&int32(1) == int32(0) {
					v157 = v131
				} else {
					v139 = v131
					v140 = v133
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v98+v140<<(uint(int32(2))%32))))
					if v151-v139 < int32(0) {
						v155 = v151
					} else {
						v155 = v139
					}
					if v151 != 0 {
						v156 = v155
					} else {
						v156 = v139
					}
					v157 = v156
				}
			}
		}
		v167 = *(*int32)(unsafe.Add(mBase, _c_F_GetOldestMultiXactId[0]))
		F_LWLockRelease(m, v167+int32(1664))
		mBase = m.M
		v171 = m.ExcPending
		if v171 != 0 {
			return int32(0)
		} else {
			return v157
		}
	}
}
func F_GetOldestTransactionIdConsideredRunning(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	F_ComputeXidHorizons(m, v5+int32(8))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v5)+24))
		m.G0 = v5 + int32(48)
		return v13
	}
}
