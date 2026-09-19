package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockDisownInternal(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v96 int64
	_ = v96
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDisownInternal[0]))
	v15 = v13
	v16 = int32(0)
	goto L2
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v22 = v15 - int32(1)
	if v22 < int32(0) {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LWLockDisownInternal[1])))
	v37 = v13 - int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockDisownInternal[0])) = v37
	if v37 <= v22 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v28 = v22 << (uint(int32(3)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+uint32(_c_F_LWLockDisownInternal[2])))
	if l0 != v29 {
		v15 = v22
		v16 = v16 + int32(1)
		goto L2
	} else {
		goto L5
	}
L5:
	;
	goto L3
L6:
	;
	m.G0 = v10 + int32(16)
	return v33
L7:
	;
	v41 = v16 & int32(3)
	if v41 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v43 = v22
	v44 = int32(0)
	goto L11
L9:
	;
	v61 = v22
	goto L10
L10:
	;
	if base.Ui32(v16-int32(1)) < base.Ui32(int32(3)) {
		goto L6
	} else {
		goto L14
	}
L11:
	;
	v49 = int32(3)
	v51 = int32(1)
	v52 = v43 + v51
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v52<<(uint(v49)%32))+uint32(_c_F_LWLockDisownInternal[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v43<<(uint(v49)%32))+uint32(_c_F_LWLockDisownInternal[2]))) = v55
	v58 = v44 + v51
	if v58 != v41 {
		v43 = v52
		v44 = v58
		goto L11
	} else {
		goto L13
	}
L12:
	;
	v61 = v52
	goto L10
L13:
	;
	goto L12
L14:
	;
	v72 = v61
	goto L15
L15:
	;
	v78 = int32(3)
	v79 = v72 << (uint(v78) % 32)
	v82 = *(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[3])))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[2]))) = v82
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[4])))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[3]))) = v86
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[5])))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[4]))) = v90
	v93 = v72 + int32(4)
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v93<<(uint(v78)%32))+uint32(_c_F_LWLockDisownInternal[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v79)+uint32(_c_F_LWLockDisownInternal[5]))) = v96
	if v93 != v37 {
		v72 = v93
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L6
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v116) <= base.Ui32(int32(94)) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v141
	F_errmsg_internal(m, int32(_a_F_LWLockDisownInternal_0), v10)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L30
	}
L21:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116<<(uint(int32(2))%32))+uint32(_c_F_LWLockDisownInternal[6])))
	v141 = v121
	goto L20
L22:
	;
	goto L23
L23:
	;
	v125 = (v116 - int32(95)) & int32(_a_F_LWLockDisownInternal_1)
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDisownInternal[7]))
	if v125 < v127 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockDisownInternal[8]))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v130+v125<<(uint(int32(2))%32))))
	if v134 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v139 = int32(_a_F_LWLockDisownInternal_2)
	goto L26
L26:
	;
	v141 = v139
	goto L20
L27:
	;
	v136 = v134
	goto L29
L28:
	;
	v136 = int32(_a_F_LWLockDisownInternal_2)
	goto L29
L29:
	;
	v139 = v136
	goto L26
L30:
	;
	F_errfinish(m, int32(_a_F_LWLockDisownInternal_3), int32(1820), int32(_a_F_LWLockDisownInternal_4))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LWLockReleaseClearVar(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v5 = base.AtomicRmwXchg64(m, l1, int32(0), int64(0))
	v6 = F_LWLockDisownInternal(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		F_LWLockReleaseInternal(m, l0, v6)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = int32(_a_F_LWLockReleaseClearVar_0)
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseClearVar[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseClearVar[0])) = v12 - int32(1)
			return
		}
	}
}
