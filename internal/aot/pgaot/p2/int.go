package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__int_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_DirectFunctionCall2Coll(m, int32(6357), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_readIntCols(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L10
	} else {
		goto L48
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L10
	} else {
		goto L45
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L10
	} else {
		goto L42
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	switch v13 {
	case 0:
		v113 = int32(0)
		goto L7
	case 1:
		goto L8
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L39
	}
L7:
	;
	m.G0 = v8 + int32(16)
	return v113
L8:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v14 != int32(40) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v20 = F_palloc(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if int32(0) < l0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v28 = int32(0)
	goto L15
L13:
	;
	goto L14
L14:
	;
	v101 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v101 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L15:
	;
	v33 = F_pg_strtok(m, v8+int32(12))
	mBase = m.M
	if v33 == int32(0) {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v36 == int32(41) {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v45 = v33
	goto L20
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20+v28<<(uint(int32(2))%32)))) = v89
	v92 = v28 + int32(1)
	if v92 != l0 {
		v28 = v92
		goto L15
	} else {
		goto L35
	}
L20:
	;
	v50 = v45 + int32(1)
	v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v45))))
	v52 = F___isspace(m, v51)
	mBase = m.M
	if v52 != 0 {
		v45 = v50
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v53 = int32(1)
	switch v51&int32(255) - int32(43) {
	case 0:
		v59 = v53
		goto L24
	default:
		v61 = v51
		v62 = v45
		v63 = v53
		goto L23
	case 2:
		goto L25
	}
L22:
	;
	goto L21
L23:
	;
	v64 = int32(0)
	v66 = v61 - int32(48)
	if base.Ui32(v66) <= base.Ui32(int32(9)) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v60 = int32(*(*int8)(unsafe.Add(mBase, uint32(v50))))
	v61 = v60
	v62 = v50
	v63 = v59
	goto L23
L25:
	;
	v59 = int32(0)
	goto L24
L26:
	;
	v69 = v64
	v70 = v66
	v71 = v62
	goto L29
L27:
	;
	v83 = v64
	goto L28
L28:
	;
	if v63 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v73 = int32(10)
	v75 = v69*v73 - v70
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v71)+1)))
	v80 = v76 - int32(48)
	if base.Ui32(v80) < base.Ui32(v73) {
		v69 = v75
		v70 = v80
		v71 = v71 + int32(1)
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v83 = v75
	goto L28
L31:
	;
	goto L30
L32:
	;
	v89 = int32(0) - v83
	goto L34
L33:
	;
	v89 = v83
	goto L34
L34:
	;
	goto L19
L35:
	;
	goto L16
L36:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v104 != int32(1) {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v107 != int32(41) {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v113 = v20
	goto L7
L39:
	;
	F_errmsg_internal(m, int32(24952), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(491759), int32(696), int32(151034))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v137
	F_errmsg_internal(m, int32(671860), v8)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L10
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(491759), int32(696), int32(151034))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	F_errmsg_internal(m, int32(24952), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(491759), int32(696), int32(151034))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errmsg_internal(m, int32(24952), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(491759), int32(696), int32(151034))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
