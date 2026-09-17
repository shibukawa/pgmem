package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetTempTablespaces(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v21 int64
	_ = v21
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v27 int64
	_ = v27
	var v29 int64
	_ = v29
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	*(*int32)(unsafe.Add(mBase, _c_F_SetTempTablespaces[0])) = l1
	*(*int32)(unsafe.Add(mBase, _c_F_SetTempTablespaces[1])) = l0
	if int32(2) <= l1 {
		v10 = int32(_a_F_SetTempTablespaces_0)
		v11 = int64(0)
		v14 = base.I64_extend_i32_u(l1 - int32(1))
		if base.Ui64(v14) <= base.Ui64(v11) {
			v61 = v11
		} else {
			v21 = v14 - v11
			v23 = *(*int64)(unsafe.Add(mBase, _c_F_SetTempTablespaces[2]))
			v24 = *(*int64)(unsafe.Add(mBase, _c_F_SetTempTablespaces[3]))
			v27 = v24
			v29 = v23
			for {
				v33 = v27 ^ v29
				v35 = base.I64_rotl(v33, int64(37))
				v43 = v33 ^ (v33<<(uint(int64(16))%64) ^ base.I64_rotl(v27, int64(24)))
				v48 = int64(base.Ui64(base.I64_rotl(v27*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v21)) % 64))
				if base.Ui64(v21) < base.Ui64(v48) {
					v27 = v43
					v29 = v35
					continue
				} else {
					break
				}
				break
			}
			*(*int64)(unsafe.Add(mBase, _c_F_SetTempTablespaces[2])) = v35
			*(*int64)(unsafe.Add(mBase, _c_F_SetTempTablespaces[3])) = v43
			v61 = v11 + v48
		}
		v64 = base.I32_wrap_i64(v61)
	} else {
		v64 = int32(0)
	}
	*(*int32)(unsafe.Add(mBase, _c_F_SetTempTablespaces[4])) = v64
	return
}
func F_checkTempNamespaceStatus(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	v2 = int32(0)
	v5 = F_get_namespace_name(m, l0)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v202
L2:
	;
	return int32(0)
L3:
	;
	if v5 == int32(0) {
		v202 = v2
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v11 = int32(8)
	v12 = int32(_a_F_checkTempNamespaceStatus_0)
	goto L8
L5:
	;
	v120 = v5 + v115
	goto L36
L6:
	;
	if v50-v51 == int32(0) {
		v115 = v11
		goto L5
	} else {
		goto L19
	}
L8:
	;
	goto L9
L9:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v19 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v20 = v5
	v21 = v12
	v22 = v11
	v23 = v19
	goto L14
L11:
	;
	v46 = v12
	v50 = int32(0)
	goto L12
L12:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	goto L6
L13:
	;
	v46 = v41
	v50 = v43
	goto L12
L14:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if base.B2i32(v23 != v25)|base.B2i32(v25 == int32(0)) != 0 {
		v41 = v21
		v43 = v23
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v41 = v35
	v43 = int32(0)
	goto L13
L16:
	;
	v31 = v22 - int32(1)
	if v31 == int32(0) {
		v41 = v21
		v43 = v23
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v34 = int32(1)
	v35 = v21 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+1)))
	if v36 != 0 {
		v20 = v20 + v34
		v21 = v35
		v22 = v31
		v23 = v36
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	v61 = int32(14)
	v62 = int32(_a_F_checkTempNamespaceStatus_1)
	goto L22
L20:
	;
	if v100-v101 == int32(0) {
		v115 = v61
		goto L5
	} else {
		goto L33
	}
L22:
	;
	goto L23
L23:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v69 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v70 = v5
	v71 = v62
	v72 = v61
	v73 = v69
	goto L28
L25:
	;
	v96 = v62
	v100 = int32(0)
	goto L26
L26:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	goto L20
L27:
	;
	v96 = v91
	v100 = v93
	goto L26
L28:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if base.B2i32(v73 != v75)|base.B2i32(v75 == int32(0)) != 0 {
		v91 = v71
		v93 = v73
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v91 = v85
	v93 = int32(0)
	goto L27
L30:
	;
	v81 = v72 - int32(1)
	if v81 == int32(0) {
		v91 = v71
		v93 = v73
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v84 = int32(1)
	v85 = v71 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v86 != 0 {
		v70 = v70 + v84
		v71 = v85
		v72 = v81
		v73 = v86
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	F_pfree(m, v5)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L2
	} else {
		goto L34
	}
L34:
	;
	return int32(0)
L35:
	;
	F_pfree(m, v5)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L2
	} else {
		goto L51
	}
L36:
	;
	v125 = v120 + int32(1)
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v120))))
	v127 = F___isspace(m, v126)
	mBase = m.M
	if v127 != 0 {
		v120 = v125
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v128 = int32(1)
	switch v126&int32(255) - int32(43) {
	case 0:
		v134 = v128
		goto L40
	default:
		v136 = v126
		v137 = v120
		v138 = v128
		goto L39
	case 2:
		goto L41
	}
L38:
	;
	goto L37
L39:
	;
	v139 = int32(0)
	v141 = v136 - int32(48)
	if base.Ui32(v141) <= base.Ui32(int32(9)) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v135 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125))))
	v136 = v135
	v137 = v125
	v138 = v134
	goto L39
L41:
	;
	v134 = int32(0)
	goto L40
L42:
	;
	v144 = v139
	v145 = v141
	v146 = v137
	goto L45
L43:
	;
	v158 = v139
	goto L44
L44:
	;
	if v138 != 0 {
		goto L48
	} else {
		goto L49
	}
L45:
	;
	v148 = int32(10)
	v150 = v144*v148 - v145
	v151 = int32(*(*int8)(unsafe.Add(mBase, uint32(v146)+1)))
	v155 = v151 - int32(48)
	if base.Ui32(v155) < base.Ui32(v148) {
		v144 = v150
		v145 = v155
		v146 = v146 + int32(1)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v158 = v150
	goto L44
L47:
	;
	goto L46
L48:
	;
	v164 = int32(0) - v158
	goto L50
L49:
	;
	v164 = v158
	goto L50
L50:
	;
	goto L35
L51:
	;
	if v164 == int32(-1) {
		v202 = v2
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v169 = int32(1)
	v170 = int32(0)
	if v164 < v170 {
		v188 = v170
		goto L54
	} else {
		goto L55
	}
L53:
	;
	if v188 == int32(0) {
		v202 = v169
		goto L1
	} else {
		goto L60
	}
L54:
	;
	goto L53
L55:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_checkTempNamespaceStatus[0]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	if base.Ui32(v177) <= base.Ui32(v164) {
		v188 = int32(0)
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v176)))
	v182 = v179 + v164*int32(640)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v182)+44))
	if v184 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v185 = v182
	goto L59
L58:
	;
	v185 = int32(0)
	goto L59
L59:
	;
	v188 = v185
	goto L54
L60:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v188)+60))
	v193 = *(*int32)(unsafe.Add(mBase, _c_F_checkTempNamespaceStatus[1]))
	if v191 != v193 {
		v202 = v169
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v188)+68))
	if v197 == l0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v199 = int32(2)
	goto L64
L63:
	;
	v199 = int32(1)
	goto L64
L64:
	;
	v202 = v199
	goto L1
}
