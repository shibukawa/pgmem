package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0]))
	if v10 < int32(200) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(_a_F_LWLockAcquire_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[1]))
	v16 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[1])) = v15 + v16
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[2]))
	v27 = int32(0)
	v29 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L49
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v30
	goto L6
L6:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.B2i32(v53 == int32(0)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	v58 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v33, v54)
	if v33 != v58 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v44 = v33 & int32(_a_F_LWLockAcquire_1)
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v47 = v33 & int32(_a_F_LWLockAcquire_2)
	v53 = v47
	v54 = base.B2i32(int32(base.Ui32(v47)>>(uint(int32(18))%32)) == int32(0)) + v33
	goto L8
L12:
	;
	v45 = v33
	goto L14
L13:
	;
	v45 = v33 | int32(_a_F_LWLockAcquire_2)
	goto L14
L14:
	;
	v53 = v44
	v54 = v45
	goto L8
L15:
	;
	v33 = v58
	goto L6
L16:
	;
	goto L17
L17:
	;
	goto L7
L18:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[3]))
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v150 | int32(16777216)
	v159 = v27
	goto L45
L19:
	;
	F_LWLockQueueSelf(m, l0, l1)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v108 = int32(_a_F_LWLockAcquire_3)
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0]))
	v111 = v109 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_LWLockAcquire[4]))) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0])) = v109 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+uint32(_c_F_LWLockAcquire[5]))) = l1
	if int32(0) < v27 {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	return int32(0)
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v69 = v66
	goto L24
L24:
	;
	if l1 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if base.B2i32(v89 == int32(0)) == int32(0) {
		goto L18
	} else {
		goto L36
	}
L26:
	;
	v94 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v69, v90)
	if v69 != v94 {
		goto L33
	} else {
		goto L34
	}
L27:
	;
	v80 = v69 & int32(_a_F_LWLockAcquire_1)
	if v80 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	v83 = v69 & int32(_a_F_LWLockAcquire_2)
	v89 = v83
	v90 = base.B2i32(int32(base.Ui32(v83)>>(uint(int32(18))%32)) == int32(0)) + v69
	goto L26
L30:
	;
	v81 = v69
	goto L32
L31:
	;
	v81 = v69 | int32(_a_F_LWLockAcquire_2)
	goto L32
L32:
	;
	v89 = v80
	v90 = v81
	goto L26
L33:
	;
	v69 = v94
	goto L24
L34:
	;
	goto L35
L35:
	;
	goto L25
L36:
	;
	F_LWLockDequeueSelf(m, l0)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L22
	} else {
		goto L37
	}
L37:
	;
	goto L21
L38:
	;
	v129 = v27
	goto L41
L39:
	;
	goto L40
L40:
	;
	return v29
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v20)+332))
	F_PGSemaphoreUnlock(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L22
	} else {
		goto L43
	}
L42:
	;
	goto L40
L43:
	;
	v135 = int32(1)
	if base.Ui32(v135) < base.Ui32(v129) {
		v129 = v129 - v135
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v20)+332))
	F_PGSemaphoreLock(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L22
	} else {
		goto L47
	}
L46:
	;
	v170 = base.AtomicRmwAnd32(m, l0, int32(4), int32(-1073741825))
	v171 = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v173))) = v171
	v27 = v159
	v29 = v171
	goto L4
L47:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+344)))
	if v167 != 0 {
		v159 = v159 + int32(1)
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockAcquire_4), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L22
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_LWLockAcquire_5), int32(1182), int32(_a_F_LWLockAcquire_6))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LWLockConditionalAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0]))
	if v8 < int32(200) {
		v11 = int32(_a_F_LWLockConditionalAcquire_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1])) = v13 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = v17
		for {
			if l1 == int32(0) {
				v29 = v20 & int32(_a_F_LWLockConditionalAcquire_1)
				if v29 != 0 {
					v30 = v20
				} else {
					v30 = v20 | int32(_a_F_LWLockConditionalAcquire_2)
				}
				v38 = v29
				v39 = v30
			} else {
				v32 = v20 & int32(_a_F_LWLockConditionalAcquire_2)
				v38 = v32
				v39 = base.B2i32(int32(base.Ui32(v32)>>(uint(int32(18))%32)) == int32(0)) + v20
			}
			v41 = base.B2i32(v38 == int32(0))
			v43 = base.AtomicRmwCmpxchg32(m, l0, int32(4), v20, v39)
			if v20 != v43 {
				v20 = v43
				continue
			} else {
				break
			}
			break
		}
		if v41 == int32(0) {
			v47 = int32(_a_F_LWLockConditionalAcquire_0)
			v49 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1])) = v49 - int32(1)
			return v41
		} else {
			v54 = int32(_a_F_LWLockConditionalAcquire_3)
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0]))
			v57 = v55 << (uint(int32(3)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_LWLockConditionalAcquire[2]))) = l0
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0])) = v55 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v57)+uint32(_c_F_LWLockConditionalAcquire[3]))) = l1
			return v41
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v74 = m.ExcPending
		if v74 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_LWLockConditionalAcquire_4), int32(0))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_LWLockConditionalAcquire_5), int32(1331), int32(_a_F_LWLockConditionalAcquire_6))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
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
func F_LWLockReleaseAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[0]))
	if int32(0) < v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = int32(_a_F_LWLockReleaseAll_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1])) = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(3))%32))+uint32(_c_F_LWLockReleaseAll[2])))
	F_LWLockRelease(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[0]))
	if int32(0) < v21 {
		v6 = v21
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
}
