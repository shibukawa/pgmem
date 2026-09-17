package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pg_clear_relation_stats(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int64
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(80)
	m.G0 = v5
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+24)) = uint8(v2)
	v9 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v9
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v9
	v13 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+26)) = uint16(v13)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = v15
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+32)) = uint8(v17)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v19
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+72)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+68)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+64)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+60)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+56)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+52)) = int32(-1082130432)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+48)) = uint8(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+44)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+40)) = uint8(v21)
	v41 = F_relation_statistics_update(m, v5+int32(8))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(80)
		return int32(0)
	}
}
func F_pg_relation_filepath(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_SearchSysCache1(m, int32(57), v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v216
L2:
	;
	return int32(0)
L3:
	;
	if v16 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
	v216 = int32(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
	v27 = v25 + v26
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+119)))
	switch v28 - int32(83) {
	case 0, 22, 26, 31, 33:
		goto L9
	default:
		goto L8
	}
L7:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+118)))
	switch v152 - int32(112) {
	case 0, 5:
		v194 = int32(-1)
		goto L51
	default:
		goto L53
	case 4:
		goto L54
	}
L8:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L50
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[1]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	if v36 != 0 {
		v150 = v36
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+117)))
	v38 = int32(0)
	if v37 == v38 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	if v140 != 0 {
		v150 = v140
		goto L7
	} else {
		goto L49
	}
L12:
	;
	goto L11
L13:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	v140 = v135
	goto L12
L14:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[2]))
	if v43 < v45 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v86 = int32(0)
	v88 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[3]))
	if v86 < v88 {
		goto L35
	} else {
		goto L36
	}
L17:
	;
	v49 = v43
	goto L20
L18:
	;
	goto L19
L19:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[4]))
	if v68 <= int32(0) {
		v140 = v38
		goto L12
	} else {
		goto L26
	}
L20:
	;
	v54 = v49 << (uint(int32(3)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+uint32(_c_F_pg_relation_filepath[5])))
	if v55 == v15 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v134 = v54 + int32(_a_F_pg_relation_filepath_0)
	goto L13
L23:
	;
	goto L24
L24:
	;
	v60 = v49 + int32(1)
	if v60 != v45 {
		v49 = v60
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L21
L26:
	;
	v73 = int32(0)
	goto L27
L27:
	;
	v78 = v73 << (uint(int32(3)) % 32)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+uint32(_c_F_pg_relation_filepath[6])))
	if v79 != v15 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v134 = v78 + int32(_a_F_pg_relation_filepath_1)
	goto L13
L29:
	;
	v82 = v73 + int32(1)
	if v68 != v82 {
		v73 = v82
		goto L27
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	v140 = v38
	goto L12
L33:
	;
	v116 = int32(0)
	goto L43
L34:
	;
	v134 = v97 + int32(_a_F_pg_relation_filepath_2)
	goto L13
L35:
	;
	v92 = v86
	goto L38
L36:
	;
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[7]))
	if v109 <= int32(0) {
		v140 = v38
		goto L12
	} else {
		goto L42
	}
L38:
	;
	v97 = v92 << (uint(int32(3)) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+uint32(_c_F_pg_relation_filepath[8])))
	if v15 == v98 {
		goto L34
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	v101 = v92 + int32(1)
	if v101 != v88 {
		v92 = v101
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L33
L43:
	;
	v121 = v116 << (uint(int32(3)) % 32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+uint32(_c_F_pg_relation_filepath[9])))
	if v122 != v15 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v134 = v121 + int32(_a_F_pg_relation_filepath_3)
	goto L13
L45:
	;
	v125 = v116 + int32(1)
	if v109 != v125 {
		v116 = v125
		goto L43
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	goto L44
L48:
	;
	v140 = v38
	goto L12
L49:
	;
	goto L8
L50:
	;
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v147)
	v216 = int32(0)
	goto L1
L51:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L70
	}
L52:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v192 = F_GetTempNamespaceProcNumber(m, v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L2
	} else {
		goto L69
	}
L53:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L2
	} else {
		goto L66
	}
L54:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[10]))
	if v159 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	if v167 == int32(0) {
		goto L52
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v160 = int32(1)
	if v155 == v159 {
		v167 = v160
		goto L56
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v167 = int32(0)
	goto L56
L60:
	;
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[11]))
	if v163 == v155 {
		v167 = v160
		goto L56
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[12]))
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_pg_relation_filepath[13]))
	if v173 == int32(-1) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v176 = v171
	goto L65
L64:
	;
	v176 = v173
	goto L65
L65:
	;
	v194 = v176
	goto L51
L66:
	;
	v181 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v181
	F_errmsg_internal(m, int32(_a_F_pg_relation_filepath_4), v12)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	F_errfinish(m, int32(_a_F_pg_relation_filepath_5), int32(1036), int32(_a_F_pg_relation_filepath_6))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L68
	}
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v194 = v192
	goto L51
L70:
	;
	v198 = v12 + int32(8)
	if v33 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v200 = v33
	goto L73
L72:
	;
	v200 = v32
	goto L73
L73:
	;
	if v200 != int32(1664) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v203 = v35
	goto L76
L75:
	;
	v203 = int32(0)
	goto L76
L76:
	;
	F_GetRelationPath(m, v198, v203, v200, v150, v194, int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L2
	} else {
		goto L77
	}
L77:
	;
	v207 = F_cstring_to_text(m, v198)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v216 = v207
	goto L1
}
