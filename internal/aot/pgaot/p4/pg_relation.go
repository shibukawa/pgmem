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
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
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
	return v215
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
	v215 = int32(0)
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
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+118)))
	switch v149 - int32(112) {
	case 0, 5:
		v191 = int32(-1)
		goto L44
	default:
		goto L46
	case 4:
		goto L47
	}
L8:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L2
	} else {
		goto L43
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	v35 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v27)+88))
	if v36 != 0 {
		v147 = v36
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+117)))
	v38 = int32(0)
	if v37 == v38 {
		goto L15
	} else {
		goto L16
	}
L11:
	;
	if v137 != 0 {
		v147 = v137
		goto L7
	} else {
		goto L42
	}
L12:
	;
	goto L11
L13:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+4))
	v137 = v132
	goto L12
L14:
	;
	v115 = int32(0)
	goto L38
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if int32(0) < v44 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[1260]))
	if int32(0) < v72 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v49 = v38
	goto L21
L19:
	;
	goto L20
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[1261]))
	if v67 <= int32(0) {
		v137 = v38
		goto L12
	} else {
		goto L25
	}
L21:
	;
	v53 = v49 << (uint(int32(3)) % 32)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53)+uint32(_consts[398])))
	if v15 == v56 {
		v128 = v53 + int32(4504448)
		goto L13
	} else {
		goto L23
	}
L22:
	;
	goto L20
L23:
	;
	v59 = v49 + int32(1)
	if v59 != v44 {
		v49 = v59
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	goto L14
L26:
	;
	v77 = v38
	goto L29
L27:
	;
	goto L28
L28:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[1262]))
	if v95 <= int32(0) {
		v137 = v38
		goto L12
	} else {
		goto L33
	}
L29:
	;
	v81 = v77 << (uint(int32(3)) % 32)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v81)+uint32(_consts[1263])))
	if v15 == v84 {
		v128 = v81 + int32(4503400)
		goto L13
	} else {
		goto L31
	}
L30:
	;
	goto L28
L31:
	;
	v87 = v77 + int32(1)
	if v87 != v72 {
		v77 = v87
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v101 = int32(0)
	goto L34
L34:
	;
	v105 = v101 << (uint(int32(3)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+uint32(_consts[1264])))
	if v15 == v108 {
		v128 = v105 + int32(4503924)
		goto L13
	} else {
		goto L36
	}
L35:
	;
	v137 = v38
	goto L12
L36:
	;
	v111 = v101 + int32(1)
	if v95 != v111 {
		v101 = v111
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v119 = v115 << (uint(int32(3)) % 32)
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v119)+uint32(_consts[1265])))
	if v15 == v122 {
		v128 = v119 + int32(4504972)
		goto L13
	} else {
		goto L40
	}
L39:
	;
	v137 = v38
	goto L12
L40:
	;
	v125 = v115 + int32(1)
	if v67 != v125 {
		v115 = v125
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	goto L8
L43:
	;
	v144 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v144)
	v215 = int32(0)
	goto L1
L44:
	;
	F_ReleaseCatCache(m, v16)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L2
	} else {
		goto L63
	}
L45:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v189 = F_GetTempNamespaceProcNumber(m, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L2
	} else {
		goto L62
	}
L46:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L2
	} else {
		goto L59
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v27)+68))
	v156 = *(*int32)(unsafe.Add(mBase, _consts[260]))
	if v156 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	if v164 == int32(0) {
		goto L45
	} else {
		goto L55
	}
L49:
	;
	goto L48
L50:
	;
	v157 = int32(1)
	if v152 == v156 {
		v164 = v157
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v164 = int32(0)
	goto L49
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[261]))
	if v160 == v152 {
		v164 = v157
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[166]))
	v170 = *(*int32)(unsafe.Add(mBase, _consts[380]))
	if v170 == int32(-1) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v173 = v168
	goto L58
L57:
	;
	v173 = v170
	goto L58
L58:
	;
	v191 = v173
	goto L44
L59:
	;
	v178 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+118)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v178
	F_errmsg_internal(m, int32(502242), v12)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L2
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(498105), int32(1036), int32(321145))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	v191 = v189
	goto L44
L63:
	;
	if v33 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v198 = v33
	goto L66
L65:
	;
	v198 = v32
	goto L66
L66:
	;
	if v198 != int32(1664) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v201 = v35
	goto L69
L68:
	;
	v201 = int32(0)
	goto L69
L69:
	;
	F_GetRelationPath(m, v12+int32(8), v201, v198, v147, v191, int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v207 = F_cstring_to_text(m, v12+int32(8))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L2
	} else {
		goto L71
	}
L71:
	;
	v215 = v207
	goto L1
}
