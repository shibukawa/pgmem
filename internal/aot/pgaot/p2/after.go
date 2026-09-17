package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AfterTriggerFreeQuery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	goto L1
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = int64(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	if v18 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v12
	F_pfree(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	return
L7:
	;
	goto L1
L8:
	;
	F_tuplestore_end(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v23 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L10
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	F_list_free_deep(m, v23)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L6
	} else {
		goto L30
	}
L13:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v33 = int32(0)
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v33<<(uint(int32(2))%32))))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+24)) = int32(0)
	if v39 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L12
L17:
	;
	F_tuplestore_end(m, v39)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L6
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+28)) = int32(0)
	if v44 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_tuplestore_end(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+32))
	if v49 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+32)) = int32(0)
	F_ExecDropSingleTupleTableSlot(m, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v55 = v33 + int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v55 < v56 {
		v33 = v55
		goto L15
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	goto L16
L30:
	;
	return
}
func F_afterTriggerMarkEvents(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v182 int32
	_ = v182
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v291 int32
	_ = v291
	v4 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == v4 {
		v291 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v291
L2:
	;
	v20 = l1
	v21 = l2
	v23 = v16
	v32 = v4
	v33 = v4
	goto L3
L3:
	;
	v35 = v23 + int32(16)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v35) < base.Ui32(v36) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v248 == int32(0) {
		v291 = v249
		goto L1
	} else {
		goto L42
	}
L5:
	;
	v38 = v35
	v51 = v32
	v52 = v33
	goto L8
L6:
	;
	v248 = v32
	v249 = v33
	goto L7
L7:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v250 != 0 {
		v23 = v250
		v32 = v248
		v33 = v249
		goto L3
	} else {
		goto L41
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if base.Ui32(int32(1073741823)) < base.Ui32(v53) {
		v207 = v53
		v217 = v51
		v218 = v52
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v248 = v217
	v249 = v218
	goto L7
L10:
	;
	v221 = v207 & int32(939524096)
	if v221 == int32(134217728) {
		v231 = int32(24)
		goto L34
	} else {
		goto L35
	}
L11:
	;
	v58 = v38 + v53&int32(134217727)
	if v21 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v202
	v207 = v202
	v217 = v200
	v218 = v201
	goto L10
L13:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerMarkEvents[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v182
	v200 = v51
	v201 = int32(1)
	v202 = v53 | int32(1073741824)
	goto L12
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v61&int32(32) == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_afterTriggerMarkEvents[1]))
	if v67 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v20 == int32(0) {
		v207 = v53
		v217 = v51
		v218 = v52
		goto L10
	} else {
		goto L31
	}
L17:
	;
	if v61&int32(64) == int32(0) {
		goto L13
	} else {
		goto L30
	}
L18:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v70 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v118 != int32(1) {
		goto L17
	} else {
		goto L28
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v84 = int32(0)
	goto L21
L21:
	;
	v94 = v67 + int32(12) + v84<<(uint(int32(3))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v73 != v95 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v100 == int32(0) {
		goto L13
	} else {
		goto L27
	}
L23:
	;
	v98 = v84 + int32(1)
	if v70 != v98 {
		v84 = v98
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	goto L22
L26:
	;
	goto L19
L27:
	;
	goto L16
L28:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v121 != 0 {
		goto L16
	} else {
		goto L29
	}
L29:
	;
	goto L13
L30:
	;
	goto L16
L31:
	;
	F_afterTriggerAddEvent(m, v20, v38, v58)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v200 = int32(1)
	v201 = v52
	v202 = v163 | int32(-2147483648)
	goto L12
L34:
	;
	v232 = v231 + v38
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v232) < base.Ui32(v233) {
		v38 = v232
		v51 = v217
		v52 = v218
		goto L8
	} else {
		goto L40
	}
L35:
	;
	if v221 != int32(805306368) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if v221 == int32(268435456) {
		v231 = int32(12)
		goto L34
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v231 = int32(16)
	goto L34
L39:
	;
	v231 = int32(4)
	goto L34
L40:
	;
	goto L9
L41:
	;
	goto L4
L42:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_afterTriggerMarkEvents[2])))
	goto L43
L43:
	;
	if int32(base.Ui32(v254&int32(2))>>(uint(int32(1))%32)) == int32(0) {
		v291 = v249
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L32
	} else {
		goto L45
	}
L45:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L32
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(_a_F_afterTriggerMarkEvents_0), int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L32
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_afterTriggerMarkEvents_1), int32(_a_F_afterTriggerMarkEvents_2), int32(_a_F_afterTriggerMarkEvents_3))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L32
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
