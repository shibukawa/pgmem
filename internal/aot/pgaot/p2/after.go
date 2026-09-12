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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
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
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v155 int32
	_ = v155
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	v4 = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v16 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v283
L2:
	;
	v283 = v4
	goto L1
L3:
	;
	goto L4
L4:
	;
	v20 = l1
	v21 = l2
	v23 = v16
	v26 = v4
	v28 = v4
	goto L5
L5:
	;
	v35 = v23 + int32(16)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v35) < base.Ui32(v36) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v243 == int32(0) {
		v283 = v241
		goto L1
	} else {
		goto L44
	}
L7:
	;
	v38 = v35
	v45 = v26
	v47 = v28
	goto L10
L8:
	;
	v241 = v26
	v243 = v28
	goto L9
L9:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v249 != 0 {
		v23 = v249
		v26 = v241
		v28 = v243
		goto L5
	} else {
		goto L43
	}
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if base.Ui32(int32(1073741823)) < base.Ui32(v53) {
		v205 = v53
		v209 = v45
		v211 = v47
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v241 = v209
	v243 = v211
	goto L9
L12:
	;
	v219 = v205 & int32(939524096)
	if v219 == int32(134217728) {
		v230 = int32(24)
		goto L36
	} else {
		goto L37
	}
L13:
	;
	v58 = v38 + v53&int32(134217727)
	if v21 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v200
	v205 = v200
	v209 = v192
	v211 = v194
	goto L12
L15:
	;
	if v20 == int32(0) {
		v205 = v53
		v209 = v45
		v211 = v47
		goto L12
	} else {
		goto L33
	}
L16:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _consts[414]))
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v155
	v192 = int32(1)
	v194 = v47
	v200 = v53 | int32(1073741824)
	goto L14
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v61&int32(32) == int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	if v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v61&int32(64) != 0 {
		goto L15
	} else {
		goto L32
	}
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v70 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v116 != int32(1) {
		goto L19
	} else {
		goto L30
	}
L22:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v85 = int32(0)
	goto L23
L23:
	;
	v94 = v67 + int32(12) + v85<<(uint(int32(3))%32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	if v73 != v95 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v100 != 0 {
		goto L15
	} else {
		goto L29
	}
L25:
	;
	v98 = v85 + int32(1)
	if v70 != v98 {
		v85 = v98
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	goto L24
L28:
	;
	goto L21
L29:
	;
	goto L16
L30:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v119 == int32(0) {
		goto L16
	} else {
		goto L31
	}
L31:
	;
	goto L15
L32:
	;
	goto L16
L33:
	;
	F_afterTriggerAddEvent(m, v20, v38, v58)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	return int32(0)
L35:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v192 = v45
	v194 = int32(1)
	v200 = v182 | int32(-2147483648)
	goto L14
L36:
	;
	v231 = v230 + v38
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if base.Ui32(v231) < base.Ui32(v232) {
		v38 = v231
		v45 = v209
		v47 = v211
		goto L10
	} else {
		goto L42
	}
L37:
	;
	if v219 == int32(268435456) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v226 = int32(12)
	goto L40
L39:
	;
	v226 = int32(4)
	goto L40
L40:
	;
	if v219 != int32(805306368) {
		v230 = v226
		goto L36
	} else {
		goto L41
	}
L41:
	;
	v230 = int32(16)
	goto L36
L42:
	;
	goto L11
L43:
	;
	goto L6
L44:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, _consts[277])))
	goto L45
L45:
	;
	if int32(base.Ui32(v253&int32(2))>>(uint(int32(1))%32)) == int32(0) {
		v283 = v241
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L34
	} else {
		goto L47
	}
L47:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L34
	} else {
		goto L48
	}
L48:
	;
	F_errmsg(m, int32(259378), int32(0))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L34
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(492334), int32(4671), int32(119778))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L34
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
