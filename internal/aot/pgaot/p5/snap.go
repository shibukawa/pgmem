package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SnapBuildWaitSnapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	v3 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v3 < v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L55
	} else {
		goto L66
	}
L2:
	;
	v17 = v3
	goto L5
L3:
	;
	goto L4
L4:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[0])))
	if v182 == int32(1) {
		goto L59
	} else {
		goto L60
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(24)+v17<<(uint(int32(2))%32))))
	if base.Ui32(v23) < base.Ui32(int32(3)) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L4
L7:
	;
	if v155 != 0 {
		goto L1
	} else {
		goto L47
	}
L8:
	;
	v155 = int32(0)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[1]))
	if v35 == v23 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v155 = int32(1)
	goto L7
L12:
	;
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[2]))
	if v39 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v155 = v145
	goto L7
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[3]))
	if v43 == int32(0) {
		v145 = int32(0)
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[4]))
	v115 = int32(0)
	v118 = v39 - int32(1)
	goto L37
L18:
	;
	v48 = v43
	goto L19
L19:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v54 == int32(4) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v145 = int32(0)
	goto L14
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v48)+80))
	if v108 != 0 {
		v48 = v108
		goto L19
	} else {
		goto L36
	}
L22:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v57 == int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v60 = int32(1)
	if v23 == v57 {
		v145 = v60
		goto L14
	} else {
		goto L24
	}
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v48)+52))
	v64 = v62 - int32(1)
	if v64 < int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v48)+48))
	v70 = int32(0)
	v73 = v64
	goto L26
L26:
	;
	v78 = int32(2)
	v79 = base.I32_div_s(v73-v70, v78)
	v80 = v79 + v70
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v67+v80<<(uint(v78)%32))))
	if v84 == v23 {
		v145 = v60
		goto L14
	} else {
		goto L28
	}
L27:
	;
	goto L21
L28:
	;
	v93 = base.B2i32(v84-v23 < int32(0)) | base.B2i32(base.Ui32(v84) < base.Ui32(int32(3)))
	if v93 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v94 = v80 + int32(1)
	goto L31
L30:
	;
	v94 = v70
	goto L31
L31:
	;
	if v93 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = v73
	goto L34
L33:
	;
	v97 = v80 - int32(1)
	goto L34
L34:
	;
	if v94 <= v97 {
		v70 = v94
		v73 = v97
		goto L26
	} else {
		goto L35
	}
L35:
	;
	goto L27
L36:
	;
	goto L20
L37:
	;
	v123 = int32(2)
	v124 = base.I32_div_s(v118-v115, v123)
	v125 = v124 + v115
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v113+v125<<(uint(v123)%32))))
	v130 = base.B2i32(v129 == v23)
	if v129 == v23 {
		v145 = v130
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v145 = v130
	goto L14
L39:
	;
	v133 = base.B2i32(base.Ui32(v129) < base.Ui32(v23))
	if base.Ui32(v129) < base.Ui32(v23) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v134 = v125 + int32(1)
	goto L42
L41:
	;
	v134 = v115
	goto L42
L42:
	;
	if base.Ui32(v129) < base.Ui32(v23) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v137 = v118
	goto L45
L44:
	;
	v137 = v125 - int32(1)
	goto L45
L45:
	;
	if v134 <= v137 {
		v115 = v134
		v118 = v137
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L38
L47:
	;
	if base.B2i32(base.Ui32(l1) < base.Ui32(int32(3)))|base.B2i32(base.Ui32(v23) < base.Ui32(int32(3))) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v171 = v17 + int32(1)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v171 < v172 {
		v17 = v171
		goto L5
	} else {
		goto L57
	}
L49:
	;
	v165 = int32(0)
	F_XactLockTableWait(m, v23, v165, v165, v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L55
	} else {
		goto L56
	}
L50:
	;
	if v23-l1 <= int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(l1) < base.Ui32(v23) {
		goto L48
	} else {
		goto L54
	}
L53:
	;
	goto L48
L54:
	;
	goto L49
L55:
	;
	return
L56:
	;
	goto L48
L57:
	;
	goto L6
L58:
	;
	if v192 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[5]))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v187)+308))
	v190 = base.B2i32(v188 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_SnapBuildWaitSnapshot[0])) = uint8(v190)
	v192 = v190
	goto L61
L60:
	;
	v192 = int32(0)
	goto L61
L61:
	;
	goto L58
L62:
	;
	v195 = F_LogStandbySnapshot(m)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L55
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	return
L65:
	;
	goto L64
L66:
	;
	F_errmsg_internal(m, int32(_a_F_SnapBuildWaitSnapshot_0), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L55
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_SnapBuildWaitSnapshot_1), int32(1453), int32(_a_F_SnapBuildWaitSnapshot_2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L55
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
