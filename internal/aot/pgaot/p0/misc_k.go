package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KnownAssignedTransactionIdsIdleMaintenance(m *base.Module) {
	var v4 int32
	_ = v4
	F_KnownAssignedXidsCompress(m, int32(3), int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_KnownAssignedXidsAdd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[0]))
	if base.Ui32(l1) < base.Ui32(l0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	if v54 <= v55 {
		goto L17
	} else {
		goto L18
	}
L2:
	;
	v19 = l0
	v20 = int32(1)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v49 = l1 - l0 + int32(1)
	goto L1
L5:
	;
	if base.B2i32(base.Ui32(l1) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v19) < base.Ui32(int32(3))) == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v34 = int32(3)
	v36 = v19 + int32(1)
	if base.Ui32(v36) <= base.Ui32(v34) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	if v19-l1 < int32(0) {
		goto L7
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	if base.Ui32(l1) <= base.Ui32(v19) {
		v49 = v20
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v49 = v20
	goto L1
L12:
	;
	goto L7
L13:
	;
	v39 = v34
	goto L15
L14:
	;
	v39 = v36
	goto L15
L15:
	;
	v19 = v39
	v20 = v20 + int32(1)
	goto L5
L16:
	;
	F_KnownAssignedXidsDisplay(m, int32(15))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L28
	} else {
		goto L52
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	if v77 < v54+v49 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v57 = int32(3)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v60+v54<<(uint(int32(2))%32)-int32(4))))
	if base.B2i32(base.Ui32(l0) < base.Ui32(v57))|base.B2i32(base.Ui32(v66) < base.Ui32(v57)) == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v66-l0 < int32(0) {
		goto L17
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if base.Ui32(l0) <= base.Ui32(v66) {
		goto L16
	} else {
		goto L23
	}
L22:
	;
	goto L16
L23:
	;
	goto L17
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L28
	} else {
		goto L49
	}
L25:
	;
	F_KnownAssignedXidsCompress(m, int32(0), l2)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v87 = v54
	goto L27
L27:
	;
	if v49 <= int32(0) {
		v174 = v87
		goto L31
	} else {
		goto L32
	}
L28:
	;
	return
L29:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	if v83 < v84+v49 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v87 = v84
	goto L27
L31:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v180 + v49
	if l2 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L32:
	;
	if v49 != int32(1) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v97 = l0
	v98 = int32(0)
	v100 = v87
	goto L36
L34:
	;
	v149 = l0
	v152 = v87
	goto L35
L35:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v159+v152<<(uint(int32(2))%32)))) = v149
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
	v167 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165+v152))) = uint8(v167)
	v174 = v152 + v167
	goto L31
L36:
	;
	v106 = int32(_a_F_KnownAssignedXidsAdd_0)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
	v108 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v107+v100<<(uint(v108)%32)))) = v97
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
	v115 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v113+v100))) = uint8(v115)
	v118 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[1]))
	v120 = v100 + v115
	v124 = int32(3)
	v126 = v97 + v115
	if base.Ui32(v126) <= base.Ui32(v124) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v49&int32(1) == int32(0) {
		v174 = v143
		goto L31
	} else {
		goto L45
	}
L38:
	;
	v129 = v124
	goto L40
L39:
	;
	v129 = v126
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118+v120<<(uint(v108)%32)))) = v129
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_KnownAssignedXidsAdd[2]))
	v134 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v132+v120))) = uint8(v134)
	v136 = int32(3)
	v138 = v129 + v134
	if base.Ui32(v138) <= base.Ui32(v136) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = v136
	goto L43
L42:
	;
	v141 = v138
	goto L43
L43:
	;
	v142 = int32(2)
	v143 = v100 + v142
	v145 = v98 + v142
	if v145 != v49&int32(2147483646) {
		v97 = v141
		v98 = v145
		v100 = v143
		goto L36
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	v149 = v141
	v152 = v143
	goto L35
L46:
	;
	v185 = int32(0)
	v188 = base.AtomicRmwOr32(m, v185, int32(_a_F_KnownAssignedXidsAdd_1), v185)
	goto L48
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v174
	return
L49:
	;
	F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsAdd_2), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L28
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(_a_F_KnownAssignedXidsAdd_3), int32(_a_F_KnownAssignedXidsAdd_4), int32(_a_F_KnownAssignedXidsAdd_5))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L28
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L28
	} else {
		goto L53
	}
L53:
	;
	F_errmsg_internal(m, int32(_a_F_KnownAssignedXidsAdd_6), int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L28
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_KnownAssignedXidsAdd_3), int32(_a_F_KnownAssignedXidsAdd_7), int32(_a_F_KnownAssignedXidsAdd_5))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L28
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
