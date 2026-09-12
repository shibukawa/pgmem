package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AddEventToPendingNotifies(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+60)) = l0
	v13 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v14 == v2 {
		v88 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v89 = F_lappend(m, v88, l0)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L17
	}
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v17 < int32(16) {
		v88 = v14
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v20 != 0 {
		v88 = v14
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(512)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(513)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = int64(17179869188)
	v28 = *(*int32)(unsafe.Add(mBase, _consts[212]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = v28
	v35 = F_hash_create(m, int32(168936), int32(256), v6+int32(-52), int32(1224))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return
L6:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v35
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	if v40 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if int32(0) < v41 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v81 = v38
	goto L9
L9:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v88 = v82
	goto L1
L10:
	;
	v47 = v2
	goto L13
L11:
	;
	goto L12
L12:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v81 = v75
	goto L9
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49+v47<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v53
	v56 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v63 = F_hash_search(m, v57, v6+int32(-56), int32(1), v6+int32(-57))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L5
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v66 = v47 + int32(1)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v66 < v67 {
		v47 = v66
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	*(*int32)(unsafe.Add(mBase, uint32(v92)+4)) = v89
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v92)+8))
	if v94 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v100 = F_hash_search(m, v94, v6+int32(-4), int32(1), v6+int32(-52))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	m.G0 = v8 - int32(-64)
	return
L21:
	;
	goto L20
}
func F_EventTriggerAlterTableStart(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v5 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+20)))
		if v8 != 0 {
			return
		} else {
			v9 = int32(4520560)
			v10 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v12
			v15 = F_palloc(m, int32(40))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(1)
				v20 = int32(*(*uint8)(unsafe.Add(mBase, _consts[430])))
				*(*uint8)(unsafe.Add(mBase, uint32(v15)+4)) = uint8(v20)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(5407363825664)
				v26 = F_copyObjectImpl(m, l0)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v26
					v30 = *(*int32)(unsafe.Add(mBase, _consts[427]))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v15
					*(*int32)(unsafe.Add(mBase, _consts[28])) = v10
					return
				}
			}
		}
	}
}
func F_EventTriggerTableRewrite(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v105 int32
	_ = v105
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v22 = v4
	v23 = v4
	v24 = v4
	v25 = v4
	v26 = v4
	v27 = v15
	v28 = v4
	v29 = int32(-1)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v29 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	goto L2
L5:
	;
	v145 = int32(m.ExcTag)
	v146 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v145 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v77
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v79
	v128 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v78
	F_pg_re_throw(m)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		v143 = v80
		goto L5
	} else {
		goto L24
	}
L7:
	;
	m.G0 = v15 + int32(32)
	return
L8:
	;
	v33 = v27 - int32(16)
	m.G0 = v33
	v36 = v33 - int32(160)
	m.G0 = v36
	v39 = int32(*(*uint8)(unsafe.Add(mBase, _consts[428])))
	if v39 != int32(1) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	v75 = v22
	v76 = v23
	v77 = v24
	v78 = v25
	v79 = v26
	v80 = v27
	v81 = v28
	goto L10
L10:
	;
	if v81 != 0 {
		goto L6
	} else {
		goto L20
	}
L11:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, _consts[429])))
	if v43 != int32(1) {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	if v47 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v33
	v57 = F_EventTriggerCommonSetup(m, l0, int32(3), int32(350449), v33)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		v143 = v36
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v57 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L15:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	*(*int32)(unsafe.Add(mBase, uint32(v63)+16)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v63)+12)) = l1
	v67 = *(*int32)(unsafe.Add(mBase, _consts[394]))
	v69 = *(*int32)(unsafe.Add(mBase, _consts[426]))
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v15 + int32(8)
	goto L19
L17:
	;
	v75 = v57
	v76 = v36
	v77 = v67
	v78 = v33
	v79 = v69
	v80 = v36
	v81 = int32(0)
	goto L10
L19:
	;
	goto L17
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v78
	F_EventTriggerInvoke(m, v75, v78)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		v143 = v80
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[427]))
	*(*int64)(unsafe.Add(mBase, uint32(v92)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, _consts[394])) = v77
	*(*int32)(unsafe.Add(mBase, _consts[426])) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v78
	F_list_free(m, v75)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		v143 = v80
		goto L5
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v78
	F_CommandCounterIncrement(m)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		v143 = v80
		goto L5
	} else {
		goto L23
	}
L23:
	;
	goto L7
L24:
	;
	goto L4
L25:
	;
	v150 = int32(v146)
	m.G0 = v143
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	if v15+int32(8) == v157 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	m.ExcPending = 1
	goto L34
L27:
	;
	if v160 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v160 = v159
	goto L30
L29:
	;
	v160 = int32(0)
	goto L30
L30:
	;
	goto L27
L31:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v15)+24))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v22 = v163
	v23 = v162
	v24 = v165
	v25 = v161
	v26 = v164
	v27 = v143
	v28 = v152
	v29 = v160
	goto L1
L32:
	;
	goto L33
L33:
	;
	F___wasm_longjmp(m, v153, v152)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
