package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DecodeNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 float64
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 float64
	_ = v152
	var v163 int32
	_ = v163
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v9
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeNumber[0])) = v9
	v27 = F_strtol(m, l1, v17+int32(8), int32(10))
	mBase = m.M
	goto L1
L1:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeNumber[0]))
	if v30 == int32(68) {
		v247 = int32(-2)
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v247
L3:
	;
	v33 = int32(-1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v34 == l1 {
		v247 = v33
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v36 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if v36 != int32(46) {
		v247 = v33
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v163 = l3 & int32(14)
	if base.B2i32(l0 != int32(3))|base.B2i32(v163 != int32(4))|base.B2i32(base.Ui32(int32(365)) < base.Ui32(v27-int32(1))) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L8:
	;
	if int32(3) <= v34-l1 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v44 = F_DecodeNumberField(m, l0, l1, l3|int32(14), l4, l5, l6, l7)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v34
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v51 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	return int32(0)
L13:
	;
	v247 = v44 >> (uint(int32(31)) % 32)
	goto L2
L14:
	;
	v53 = v34 + int32(1)
	v54 = int32(_a_F_DecodeNumber_0)
	v58 = m.G0
	v60 = v58 - int32(32)
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v61
	v69 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeNumber[1])))
	if v69 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v152 = float64(0)
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = base.I32_trunc_sat_f64_s(base.F64_nearest(base.F64_mul(v152, float64(1e+06))))
	goto L7
L17:
	;
	v138 = F_strlen(m, v53)
	mBase = m.M
	if v137 != v138 {
		v247 = v33
		goto L2
	} else {
		goto L36
	}
L18:
	;
	v137 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_DecodeNumber[2])))
	if v73 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v77 = v53
	goto L24
L22:
	;
	goto L23
L23:
	;
	v87 = v54
	v88 = v69
	goto L27
L24:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v83 == v69 {
		v77 = v77 + int32(1)
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v137 = v77 - v53
	goto L17
L26:
	;
	goto L25
L27:
	;
	v95 = v60 + int32(base.Ui32(v88)>>(uint(int32(3))%32))&int32(28)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v97 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v96 | v97<<(uint(v88)%32)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v101 != 0 {
		v87 = v87 + v97
		v88 = v101
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v104 == int32(0) {
		v127 = v53
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	v137 = v127 - v53
	goto L17
L31:
	;
	v108 = v53
	v109 = v104
	goto L32
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v60+int32(base.Ui32(v109)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v117)>>(uint(v109)%32))&int32(1) == int32(0) {
		v127 = v108
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v127 = v125
	goto L30
L34:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+1)))
	v125 = v108 + int32(1)
	if v123 != 0 {
		v108 = v125
		v109 = v123
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_DecodeNumber[0])) = int32(0)
	v145 = F_strtod(m, v34, v17+int32(12))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L37
	}
L37:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147))))
	if v148 != 0 {
		v247 = v33
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeNumber[0]))
	if v150 != 0 {
		v247 = v33
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v152 = v145
	goto L16
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(_a_F_DecodeNumber_1)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+28)) = v27
	v247 = int32(0)
	goto L2
L41:
	;
	goto L42
L42:
	;
	switch v163 - int32(1) {
	case 0, 2, 4, 6, 8, 10, 12:
		goto L45
	case 1:
		goto L49
	case 3, 7:
		goto L44
	case 5:
		goto L48
	case 9:
		goto L47
	case 11:
		v247 = v33
		goto L2
	case 13:
		goto L46
	default:
		goto L50
	}
L43:
	;
	v237 = int32(0)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v238 != int32(4) {
		v247 = v237
		goto L2
	} else {
		goto L72
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v27
	goto L43
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	v227 = F_DecodeNumberField(m, l0, l1, l3, l4, l5, l6, l7)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L12
	} else {
		goto L71
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	goto L43
L48:
	;
	if l2 != 0 {
		goto L65
	} else {
		goto L66
	}
L49:
	;
	if l2 != 0 {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	if l0 <= int32(2) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v183 != int32(1) {
		goto L44
	} else {
		goto L56
	}
L52:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeNumber[3]))
	if v183 != 0 {
		goto L51
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	goto L43
L55:
	;
	goto L54
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L43
L57:
	;
	if l0 <= int32(2) {
		goto L61
	} else {
		goto L62
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L43
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L43
L61:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_DecodeNumber[3]))
	if v196 != 0 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	goto L43
L64:
	;
	goto L63
L65:
	;
	if l0 < int32(3) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L43
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L43
L69:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7))))
	if v208 != int32(1) {
		goto L68
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v213
	v216 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v216)
	goto L43
L71:
	;
	v247 = v227 >> (uint(int32(31)) % 32) & v227
	goto L2
L72:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(base.B2i32(l0 < int32(3)))
	v247 = v237
	goto L2
}
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v109
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v108 = v105
	v109 = v103
	goto L1
L3:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v101 = v99
	v103 = int32(0)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L23
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v14 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L7
	case 2:
		goto L6
	default:
		goto L4
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_DeconstructQualifiedName[0]))
	v31 = F_get_database_name(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v101 = v17 + int32(4)
	v103 = v21
	goto L2
L8:
	;
	return
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	if base.B2i32(v35 == int32(0))|base.B2i32(v35 != v38) != 0 {
		v56 = v35
		v57 = v38
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v56-v57 == int32(0) {
		v108 = v24
		v109 = v26
		goto L1
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v41 = v28
	v42 = v31
	goto L13
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v46 == int32(0) {
		v56 = v46
		v57 = v45
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v56 = v46
	v57 = v45
	goto L11
L15:
	;
	v49 = int32(1)
	if v46 == v45 {
		v41 = v41 + v49
		v42 = v42 + v49
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v68 = F_NameListToString(m, l0)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v68
	F_errmsg(m, int32(_a_F_DeconstructQualifiedName_0), v10+int32(16))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_DeconstructQualifiedName_1), int32(3333), int32(_a_F_DeconstructQualifiedName_2))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v88 = F_NameListToString(m, l0)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v88
	F_errmsg(m, int32(_a_F_DeconstructQualifiedName_3), v10)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_DeconstructQualifiedName_1), int32(3339), int32(_a_F_DeconstructQualifiedName_2))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DeescapeQuotedString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	v2 = int32(0)
	v11 = F_strlen(m, l0)
	mBase = m.M
	v13 = v11 - int32(1)
	v14 = F_palloc(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if int32(0) < v13 {
			v21 = l0 + int32(1)
			v23 = int32(0)
			v29 = v2
			for {
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v21))))
				if v34 != int32(39) {
					if v34 != int32(92) {
						v94 = v34
						v96 = v23
					} else {
						v41 = v23 + int32(1)
						v42 = v21 + v41
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
						v45 = v43 - int32(48)
						switch v45 {
						case 0, 1, 2, 3, 4, 5, 6, 7:
							v50 = int32(0)
							if base.Ui32(int32(55)) < base.Ui32(v43) {
								v81 = v50
								v83 = v50
								v94 = v81
								v96 = v83 + v23
							} else {
								v54 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+1)))
								if v54 < int32(48) {
									v94 = v45
									v96 = v23 + int32(1)
								} else {
									if base.Ui32(int32(55)) < base.Ui32(v54) {
										v94 = v45
										v96 = v23 + int32(1)
									} else {
										v66 = int32(48)
										v67 = v54 + v45<<(uint(int32(3))%32) - v66
										v69 = int32(*(*int8)(unsafe.Add(mBase, uint32(v42)+2)))
										if base.B2i32(v69 < v66)|base.B2i32(base.Ui32(int32(55)) < base.Ui32(v69)) != 0 {
											v81 = v67
											v83 = int32(2)
										} else {
											v75 = int32(3)
											v81 = v69 + v67<<(uint(v75)%32) - int32(48)
											v83 = v75
										}
										v94 = v81
										v96 = v83 + v23
									}
								}
							}
						default:
							v94 = v43
							v96 = v41
						case 50:
							v94 = int32(8)
							v96 = v41
						case 54:
							v94 = int32(12)
							v96 = v41
						case 62:
							v94 = int32(10)
							v96 = v41
						case 66:
							v94 = int32(13)
							v96 = v41
						case 68:
							v94 = int32(9)
							v96 = v41
						}
					}
				} else {
					v85 = int32(39)
					v87 = v23 + int32(1)
					v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+v87))))
					if v89 == v85 {
						v94 = v85
						v96 = v87
					} else {
						v94 = v34
						v96 = v23
					}
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v14+v29))) = uint8(v94)
				v101 = int32(1)
				v102 = v29 + v101
				v104 = v96 + v101
				if v104 < v13 {
					v23 = v104
					v29 = v102
					continue
				} else {
					break
				}
				break
			}
			v112 = v102
		} else {
			v112 = v2
		}
		v119 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14+v112-int32(1)))) = uint8(v119)
		return v14
	}
}
func F_DelRoleMems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int64
	_ = v187
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v25 = F_check_role_grantor(m, l0, l2, l5, v9)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_table_open(m, int32(1261), int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	F_LockSharedObject(m, int32(1260), l2, int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = int32(0)
	v40 = F_SearchSysCacheList(m, int32(9), int32(1), l2, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v101 = v40 + int32(48)
	v111 = v9
	goto L13
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v42 == int32(0) {
		v91 = v9
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = F_palloc(m, v42<<(uint(int32(2))%32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v49 <= int32(0) {
		v91 = v47
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v55 = int32(0)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47+v55<<(uint(int32(2))%32)))) = int32(0)
	v78 = v55 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v78 < v79 {
		v55 = v78
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v91 = v47
	goto L5
L12:
	;
	goto L11
L13:
	;
	v121 = int32(0)
	if l3 == v121 {
		v130 = v121
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v124 <= v111 {
		v130 = v121
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v130 = v126 + v111<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if int32(0) < v276 {
		goto L48
	} else {
		goto L49
	}
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if int32(0) < v140 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if base.B2i32(v130 == int32(0))|base.B2i32(v135 <= v111) != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v138 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v146 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L46
	}
L26:
	;
	v164 = v146 << (uint(int32(2)) % 32)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v91+v164)))
	if v166 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L25
L28:
	;
	v245 = v146 + int32(1)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v245 < v246 {
		v146 = v245
		goto L26
	} else {
		goto L45
	}
L29:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v164+v101)))
	if v166 == int32(4) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+56))
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v174)+22)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174+v175)))
	F_deleteSharedDependencyRecordsFor(m, int32(1261), v177, int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v185 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = v185
	v187 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v187
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v187
	*(*int32)(unsafe.Add(mBase, uint32(v22)+27)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v22)+19)) = v185
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v185
	switch v166 - int32(1) {
	case 0:
		goto L36
	case 1:
		goto L39
	case 2:
		goto L38
	default:
		goto L37
	}
L33:
	;
	F_simple_heap_delete(m, v29, v170+int32(44))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L28
L35:
	;
	v236 = F_heap_modify_tuple(m, v170+int32(40), v31, v20+int32(-32), v20+int32(-40), v20+int32(-48))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L43
	}
L36:
	;
	v224 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = int32(0)
	goto L35
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)) = uint8(v207)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = int32(0)
	goto L35
L39:
	;
	v203 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)) = uint8(v203)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = int32(0)
	goto L35
L40:
	;
	F_errmsg_internal(m, int32(_a_F_DelRoleMems_0), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_DelRoleMems_1), int32(2089), int32(_a_F_DelRoleMems_2))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	F_CatalogTupleUpdate(m, v29, v236+int32(4), v236)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L28
L45:
	;
	goto L27
L46:
	;
	F_relation_close(m, v29, int32(0))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v22 - int32(-64)
	return
L48:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v138+v111<<(uint(int32(2))%32))))
	v286 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v359 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L64
	}
L51:
	;
	v304 = v286 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v101+v304)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+56))
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+22)))
	v309 = v307 + v308
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+8))
	if v310 != v282 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v336 = v286 + int32(1)
	if v336 != v276 {
		v286 = v336
		goto L51
	} else {
		goto L63
	}
L54:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	if v312 != v25 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v314&int32(2) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v304))) = int32(2)
	v111 = v111 + int32(1)
	goto L13
L57:
	;
	goto L58
L58:
	;
	if v314&int32(4) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v304))) = int32(3)
	v111 = v111 + int32(1)
	goto L13
L60:
	;
	goto L61
L61:
	;
	F_plan_recursive_revoke(m, v40, v91, v286, v314&int32(1), l7)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v111 = v111 + int32(1)
	goto L13
L63:
	;
	goto L52
L64:
	;
	if v359 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v361 = F_get_rolespec_name(m, v275)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L68
	}
L66:
	;
	goto L67
L67:
	;
	v111 = v111 + int32(1)
	goto L13
L68:
	;
	v364 = F_GetUserNameFromId(m, v25, int32(0))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v364
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v361
	F_errmsg(m, int32(_a_F_DelRoleMems_3), v22)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_DelRoleMems_1), int32(2027), int32(_a_F_DelRoleMems_2))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	goto L67
}
func F_DropPreparedStatement(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_DropPreparedStatement[0]))
	if v11 != 0 {
		v12 = int32(0)
		v14 = F_hash_search(m, v11, l0, v12, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = v14
			if v16 == int32(0) {
				v19 = l1
			} else {
				v19 = v3
			}
			if v19 == int32(0) {
				if v16 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
					F_DropCachedPlan(m, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _c_F_DropPreparedStatement[0]))
						v29 = F_hash_search(m, v26, v16, int32(2), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errcode(m, int32(386))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(_a_F_DropPreparedStatement_0), v7)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_DropPreparedStatement_1), int32(454), int32(_a_F_DropPreparedStatement_2))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		}
	} else {
		v16 = v3
		if v16 == int32(0) {
			v19 = l1
		} else {
			v19 = v3
		}
		if v19 == int32(0) {
			if v16 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
				F_DropCachedPlan(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_DropPreparedStatement[0]))
					v29 = F_hash_search(m, v26, v16, int32(2), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errcode(m, int32(386))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg(m, int32(_a_F_DropPreparedStatement_0), v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_DropPreparedStatement_1), int32(454), int32(_a_F_DropPreparedStatement_2))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_danish_ISO_8859_1_close_env(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_SN_close_env(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_datadir_fsync_fname(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_datadir_fsync_fname[0]))
	if v21 != 0 {
		v22 = F_GetCurrentTimestamp(m)
		mBase = m.M
		v24 = *(*int64)(unsafe.Add(mBase, _c_F_datadir_fsync_fname[1]))
		F_TimestampDifference(m, v24, v22, v18+int32(12), v18+int32(8))
		mBase = m.M
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(28)))) = v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v32
		*(*int32)(unsafe.Add(mBase, _c_F_datadir_fsync_fname[0])) = int32(0)
	} else {
	}
	m.G0 = v18 + int32(16)
	if base.B2i32(v21 != int32(0)) == int32(0) {
		v67 = F_fsync_fname_ext(m, l0, l1, int32(1), l2)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	} else {
		v47 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			if v47 == int32(0) {
				v67 = F_fsync_fname_ext(m, l0, l1, int32(1), l2)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v51
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
				v56 = base.I32_div_s(v54, int32(_a_F_datadir_fsync_fname_0))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v56
				F_errmsg(m, int32(_a_F_datadir_fsync_fname_1), v7)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_datadir_fsync_fname_2), int32(3832), int32(_a_F_datadir_fsync_fname_3))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v67 = F_fsync_fname_ext(m, l0, l1, int32(1), l2)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_date2timestamp_no_overflow(m *base.Module, l0 int32) float64 {
	if l0 == int32(-2147483648) {
		return float64(-1.7976931348623157e+308)
	} else {
		if l0 == int32(2147483647) {
			return float64(1.7976931348623157e+308)
		} else {
			return base.F64_mul(base.F64_convert_i32_s(l0), float64(8.64e+10))
		}
	}
}
func F_dcbrt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v14 int32
	_ = v14
	var v21 float64
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 float64
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v77 float64
	_ = v77
	var v83 float64
	_ = v83
	var v85 float64
	_ = v85
	var v93 float64
	_ = v93
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v14 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v7))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v14) {
		v83 = base.F64_add(v7, v7)
	} else {
		if base.Ui32(int32(_a_F_dcbrt_0)) < base.Ui32(v14) {
			v31 = v14
			v32 = v7
			v33 = int32(715094163)
			v35 = base.I32_div_u_s(v31, int32(3))
			v41 = base.F64_copysign(base.F64_reinterpret_i64(base.I64_extend_i32_u(v33+v35)<<(uint(int64(32))%64)), v32)
			v44 = base.F64_mul(base.F64_mul(v41, v41), base.F64_div(v41, v7))
			v66 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_mul(v41, base.F64_add(base.F64_mul(base.F64_mul(v44, base.F64_mul(v44, v44)), base.F64_add(base.F64_mul(v44, float64(0.14599619288661245)), float64(-0.758397934778766))), base.F64_add(base.F64_mul(v44, base.F64_add(base.F64_mul(v44, float64(1.6214297201053545)), float64(-1.8849797954337717))), float64(1.87595182427177)))))&int64(-1073741824) + int64(2147483648))
			v68 = base.F64_div(v7, base.F64_mul(v66, v66))
			v77 = base.F64_add(base.F64_mul(v66, base.F64_div(base.F64_sub(v68, v66), base.F64_add(base.F64_add(v66, v66), v68))), v66)
		} else {
			v21 = base.F64_mul(v7, float64(1.8014398509481984e+16))
			v27 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v21))>>(uint(int64(32))%64))) & int32(2147483647)
			if v27 == int32(0) {
				v77 = v7
			} else {
				v31 = v27
				v32 = v21
				v33 = int32(696219795)
				v35 = base.I32_div_u_s(v31, int32(3))
				v41 = base.F64_copysign(base.F64_reinterpret_i64(base.I64_extend_i32_u(v33+v35)<<(uint(int64(32))%64)), v32)
				v44 = base.F64_mul(base.F64_mul(v41, v41), base.F64_div(v41, v7))
				v66 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_mul(v41, base.F64_add(base.F64_mul(base.F64_mul(v44, base.F64_mul(v44, v44)), base.F64_add(base.F64_mul(v44, float64(0.14599619288661245)), float64(-0.758397934778766))), base.F64_add(base.F64_mul(v44, base.F64_add(base.F64_mul(v44, float64(1.6214297201053545)), float64(-1.8849797954337717))), float64(1.87595182427177)))))&int64(-1073741824) + int64(2147483648))
				v68 = base.F64_div(v7, base.F64_mul(v66, v66))
				v77 = base.F64_add(base.F64_mul(v66, base.F64_div(base.F64_sub(v68, v66), base.F64_add(base.F64_add(v66, v66), v68))), v66)
			}
		}
		v83 = v77
	}
	v85 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v83), v85)&base.F64_ne(base.F64_abs(v7), v85) == int32(0) {
		v93 = float64(0)
		if base.F64_eq(v83, v93)&base.F64_ne(v7, v93) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v106 = m.ExcPending
			if v106 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v98 = F_Float8GetDatum(m, v83)
			mBase = m.M
			v101 = m.ExcPending
			if v101 != 0 {
				return int32(0)
			} else {
				return v98
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v104 = m.ExcPending
		if v104 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dcosd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 float64
	_ = v10
	var v20 int32
	_ = v20
	var v30 int64
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 float64
	_ = v39
	var v42 int64
	_ = v42
	var v49 float64
	_ = v49
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v72 int32
	_ = v72
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v95 int64
	_ = v95
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int32
	_ = v106
	var v111 int64
	_ = v111
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v141 int32
	_ = v141
	var v156 int64
	_ = v156
	var v164 float64
	_ = v164
	var v168 float64
	_ = v168
	var v172 float64
	_ = v172
	var v176 float64
	_ = v176
	var v181 float64
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v201 float64
	_ = v201
	var v205 int32
	_ = v205
	var v206 float64
	_ = v206
	var v207 float64
	_ = v207
	var v212 float64
	_ = v212
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v219 float64
	_ = v219
	var v223 float64
	_ = v223
	var v227 float64
	_ = v227
	var v231 float64
	_ = v231
	var v240 float64
	_ = v240
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v260 float64
	_ = v260
	var v264 int32
	_ = v264
	var v265 float64
	_ = v265
	var v266 float64
	_ = v266
	var v272 float64
	_ = v272
	var v273 float64
	_ = v273
	var v275 float64
	_ = v275
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v288 float64
	_ = v288
	var v292 float64
	_ = v292
	var v299 float64
	_ = v299
	var v303 float64
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*float64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(base.I64_reinterpret_f64(v10)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L92
	} else {
		goto L98
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L92
	} else {
		goto L94
	}
L3:
	;
	if base.F64_eq(base.F64_abs(v10), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v303 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L5
L5:
	;
	v304 = F_Float8GetDatum(m, v303)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L92
	} else {
		goto L93
	}
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dcosd[0])))
	if v20 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_init_degree_constants(m)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	v30 = base.I64_reinterpret_f64(v10)
	v34 = int32(2047)
	v35 = base.I32_wrap_i64(int64(base.Ui64(v30)>>(uint(int64(52))%64))) & v34
	if v35 == v34 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v292), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L88
	}
L11:
	;
	if base.F64_lt(v164, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L12:
	;
	v39 = base.F64_mul(v10, float64(360))
	v164 = base.F64_div(v39, v39)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v42 = v30 << (uint(int64(1)) % 64)
	if base.Ui64(v42) <= base.Ui64(int64(-9156662467374350336)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v42 == int64(-9156662467374350336) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v35 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v49 = base.F64_mul(v10, float64(0))
	goto L20
L19:
	;
	v49 = v10
	goto L20
L20:
	;
	v164 = v49
	goto L11
L21:
	;
	if int32(1031) < v85 {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v52 = int32(0)
	v54 = v30 << (uint(int64(12)) % 64)
	if int64(0) <= v54 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v85 = v35
	v86 = v30&int64(4503599627370495) | int64(4503599627370496)
	goto L21
L25:
	;
	v58 = v54
	v61 = v52
	goto L28
L26:
	;
	v72 = v52
	goto L27
L27:
	;
	v85 = v72
	v86 = v30 << (uint(base.I64_extend_i32_u(int32(1)-v72)) % 64)
	goto L21
L28:
	;
	v63 = v61 - int32(1)
	v65 = v58 << (uint(int64(1)) % 64)
	if int64(0) <= v65 {
		v58 = v65
		v61 = v63
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v72 = v63
	goto L27
L30:
	;
	goto L29
L31:
	;
	v90 = v86
	v93 = v85
	goto L34
L32:
	;
	v111 = v86
	v114 = v85
	goto L33
L33:
	;
	v116 = v111 - int64(6333186975989760)
	if v116 < int64(0) {
		v123 = v111
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v95 = v90 - int64(6333186975989760)
	if v95 < int64(0) {
		v102 = v90
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v111 = v104
	v114 = int32(1031)
	goto L33
L36:
	;
	v104 = v102 << (uint(int64(1)) % 64)
	v106 = v93 - int32(1)
	if int32(1031) < v106 {
		v90 = v104
		v93 = v106
		goto L34
	} else {
		goto L39
	}
L37:
	;
	if v95 != int64(0) {
		v102 = v95
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v164 = base.F64_mul(v10, float64(0))
	goto L11
L39:
	;
	goto L35
L40:
	;
	if base.Ui64(v123) <= base.Ui64(int64(4503599627370495)) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v116 != int64(0) {
		v123 = v116
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v164 = base.F64_mul(v10, float64(0))
	goto L11
L43:
	;
	v127 = v123
	v130 = v114
	goto L46
L44:
	;
	v138 = v123
	v141 = v114
	goto L45
L45:
	;
	if int32(0) < v141 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v132 = v130 - int32(1)
	v134 = v127 << (uint(int64(1)) % 64)
	if base.Ui64(v127) < base.Ui64(int64(2251799813685248)) {
		v127 = v134
		v130 = v132
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v138 = v134
	v141 = v132
	goto L45
L48:
	;
	goto L47
L49:
	;
	v156 = v138 - int64(4503599627370496) | base.I64_extend_i32_u(v141)<<(uint(int64(52))%64)
	goto L51
L50:
	;
	v156 = int64(base.Ui64(v138) >> (uint(base.I64_extend_i32_u(int32(1)-v141)) % 64))
	goto L51
L51:
	;
	v164 = base.F64_reinterpret_i64(v30&int64(-9223372036854775807-1) | v156)
	goto L11
L52:
	;
	v168 = base.F64_neg(v164)
	goto L54
L53:
	;
	v168 = v164
	goto L54
L54:
	;
	if base.F64_gt(v168, float64(180)) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v172 = base.F64_sub(float64(360), v168)
	goto L57
L56:
	;
	v172 = v168
	goto L57
L57:
	;
	if base.F64_gt(v172, float64(90)) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v176 = base.F64_sub(float64(180), v172)
	goto L60
L59:
	;
	v176 = v172
	goto L60
L60:
	;
	if base.F64_le(v176, float64(60)) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v181 = base.F64_mul(v176, float64(0.017453292519943295))
	v185 = m.G0
	v187 = v185 - int32(16)
	m.G0 = v187
	v194 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v181))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v194) <= base.Ui32(int32(1072243195)) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L63
L63:
	;
	v240 = base.F64_mul(base.F64_sub(float64(90), v176), float64(0.017453292519943295))
	v244 = m.G0
	v246 = v244 - int32(16)
	m.G0 = v246
	v253 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v240))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v253) <= base.Ui32(int32(1072243195)) {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v227 = base.F64_sub(float64(1), v223)
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v227
	v231 = *(*float64)(unsafe.Add(mBase, _c_F_dcosd[1]))
	v292 = base.F64_add(base.F64_mul(base.F64_div(v227, v231), float64(-0.5)), float64(1))
	goto L10
L65:
	;
	m.G0 = v187 + int32(16)
	goto L64
L66:
	;
	if base.Ui32(v194) < base.Ui32(int32(1044816030)) {
		v223 = float64(1)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v194) {
		v223 = base.F64_sub(v181, v181)
		goto L65
	} else {
		goto L70
	}
L69:
	;
	v201 = F___cos(m, v181, float64(0))
	mBase = m.M
	v223 = v201
	goto L65
L70:
	;
	v205 = F___rem_pio2(m, v181, v187)
	mBase = m.M
	v206 = *(*float64)(unsafe.Add(mBase, uint32(v187)+8))
	v207 = *(*float64)(unsafe.Add(mBase, uint32(v187)))
	switch v205&int32(3) - int32(1) {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	default:
		goto L74
	}
L71:
	;
	v219 = F___sin(m, v207, v206, int32(1))
	mBase = m.M
	v223 = v219
	goto L65
L72:
	;
	v216 = F___cos(m, v207, v206)
	mBase = m.M
	v223 = base.F64_neg(v216)
	goto L65
L73:
	;
	v214 = F___sin(m, v207, v206, int32(1))
	mBase = m.M
	v223 = base.F64_neg(v214)
	goto L65
L74:
	;
	v212 = F___cos(m, v207, v206)
	mBase = m.M
	v223 = v212
	goto L65
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v279
	v288 = *(*float64)(unsafe.Add(mBase, _c_F_dcosd[2]))
	v292 = base.F64_mul(base.F64_div(v279, v288), float64(0.5))
	goto L10
L76:
	;
	m.G0 = v246 + int32(16)
	goto L75
L77:
	;
	if base.Ui32(v253) < base.Ui32(int32(1045430272)) {
		v279 = v240
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v253) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v260 = F___sin(m, v240, float64(0), int32(0))
	mBase = m.M
	v279 = v260
	goto L76
L81:
	;
	v279 = base.F64_sub(v240, v240)
	goto L76
L82:
	;
	goto L83
L83:
	;
	v264 = F___rem_pio2(m, v240, v246)
	mBase = m.M
	v265 = *(*float64)(unsafe.Add(mBase, uint32(v246)+8))
	v266 = *(*float64)(unsafe.Add(mBase, uint32(v246)))
	switch v264&int32(3) - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L85
	case 2:
		goto L84
	default:
		goto L87
	}
L84:
	;
	v277 = F___cos(m, v266, v265)
	mBase = m.M
	v279 = base.F64_neg(v277)
	goto L76
L85:
	;
	v275 = F___sin(m, v266, v265, int32(1))
	mBase = m.M
	v279 = base.F64_neg(v275)
	goto L76
L86:
	;
	v273 = F___cos(m, v266, v265)
	mBase = m.M
	v279 = v273
	goto L76
L87:
	;
	v272 = F___sin(m, v266, v265, int32(1))
	mBase = m.M
	v279 = v272
	goto L76
L88:
	;
	if base.F64_gt(v172, float64(90)) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v299 = base.F64_neg(v292)
	goto L91
L90:
	;
	v299 = v292
	goto L91
L91:
	;
	v303 = v299
	goto L5
L92:
	;
	return int32(0)
L93:
	;
	m.G0 = v7 + int32(16)
	return v304
L94:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_dcosd_0), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L92
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(_a_F_dcosd_1), int32(2334), int32(_a_F_dcosd_2))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L92
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dcosh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v9 float64
	_ = v9
	var v10 int64
	_ = v10
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v63 int32
	_ = v63
	var v64 float64
	_ = v64
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v102 float64
	_ = v102
	var v105 float64
	_ = v105
	var v111 float64
	_ = v111
	var v120 float64
	_ = v120
	var v135 float64
	_ = v135
	var v144 float64
	_ = v144
	var v149 float64
	_ = v149
	var v156 float64
	_ = v156
	var v159 float64
	_ = v159
	var v165 float64
	_ = v165
	var v175 float64
	_ = v175
	var v177 float64
	_ = v177
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v199 float64
	_ = v199
	var v206 float64
	_ = v206
	var v210 float64
	_ = v210
	var v215 float64
	_ = v215
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	*(*int32)(unsafe.Add(mBase, _c_F_dcosh[0])) = int32(0)
	v9 = base.F64_abs(v5)
	v10 = base.I64_reinterpret_f64(v9)
	if base.Ui64(v10) <= base.Ui64(int64(4604418530035630079)) {
		if base.Ui64(v10) < base.Ui64(int64(4490088828488384512)) {
			v215 = float64(1)
		} else {
			v22 = base.I64_reinterpret_f64(v9)
			v27 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v27) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v9)&int64(9223372036854775807)) {
					v177 = v9
					v189 = v177
				} else {
					if v22 < int64(0) {
						v189 = float64(-1)
					} else {
						if base.F64_gt(v9, float64(709.782712893384)) == int32(0) {
							v63 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v9, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v9)))
							v64 = base.F64_convert_i32_s(v63)
							v70 = v63
							v71 = base.F64_mul(v64, float64(1.9082149292705877e-10))
							v73 = base.F64_add(v9, base.F64_mul(v64, float64(-0.6931471803691238)))
							v74 = base.F64_sub(v73, v71)
							v80 = v74
							v81 = v70
							v82 = base.F64_sub(base.F64_sub(v73, v74), v71)
							v85 = base.F64_mul(v80, float64(0.5))
							v86 = base.F64_mul(v80, v85)
							v102 = base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v105 = base.F64_sub(float64(3), base.F64_mul(v102, v85))
							v111 = base.F64_mul(v86, base.F64_div(base.F64_sub(v102, v105), base.F64_sub(float64(6), base.F64_mul(v80, v105))))
							if v81 == int32(0) {
								v189 = base.F64_sub(v80, base.F64_sub(base.F64_mul(v80, v111), v86))
							} else {
								v120 = base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_sub(v111, v82)), v82), v86)
								switch v81 + int32(1) {
								case 0:
									v189 = base.F64_add(base.F64_mul(base.F64_sub(v80, v120), float64(0.5)), float64(-0.5))
								default:
									v144 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v81+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v81) {
										v149 = base.F64_add(base.F64_sub(v80, v120), float64(1))
										if v81 == int32(1024) {
											v156 = base.F64_mul(base.F64_add(v149, v149), float64(8.98846567431158e+307))
										} else {
											v156 = base.F64_mul(v149, v144)
										}
										v189 = base.F64_add(v156, float64(-1))
									} else {
										v159 = float64(1)
										v165 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v81) << (uint(int64(52)) % 64))
										if base.Ui32(v81) <= base.Ui32(int32(19)) {
											v175 = base.F64_add(base.F64_sub(v159, v165), base.F64_sub(v80, v120))
										} else {
											v175 = base.F64_add(base.F64_sub(v80, base.F64_add(v120, v165)), v159)
										}
										v177 = base.F64_mul(v175, v144)
										v189 = v177
									}
								case 2:
									if base.F64_lt(v80, float64(-0.25)) != 0 {
										v189 = base.F64_mul(base.F64_sub(v120, base.F64_add(v80, float64(0.5))), float64(-2))
									} else {
										v135 = base.F64_sub(v80, v120)
										v189 = base.F64_add(base.F64_add(v135, v135), float64(1))
									}
								}
							}
						} else {
							v189 = base.F64_mul(v9, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v27) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v27) < base.Ui32(int32(1016070144)) {
						v177 = v9
						v189 = v177
					} else {
						v80 = v9
						v81 = int32(0)
						v82 = float64(0)
						v85 = base.F64_mul(v80, float64(0.5))
						v86 = base.F64_mul(v80, v85)
						v102 = base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v105 = base.F64_sub(float64(3), base.F64_mul(v102, v85))
						v111 = base.F64_mul(v86, base.F64_div(base.F64_sub(v102, v105), base.F64_sub(float64(6), base.F64_mul(v80, v105))))
						if v81 == int32(0) {
							v189 = base.F64_sub(v80, base.F64_sub(base.F64_mul(v80, v111), v86))
						} else {
							v120 = base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_sub(v111, v82)), v82), v86)
							switch v81 + int32(1) {
							case 0:
								v189 = base.F64_add(base.F64_mul(base.F64_sub(v80, v120), float64(0.5)), float64(-0.5))
							default:
								v144 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v81+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v81) {
									v149 = base.F64_add(base.F64_sub(v80, v120), float64(1))
									if v81 == int32(1024) {
										v156 = base.F64_mul(base.F64_add(v149, v149), float64(8.98846567431158e+307))
									} else {
										v156 = base.F64_mul(v149, v144)
									}
									v189 = base.F64_add(v156, float64(-1))
								} else {
									v159 = float64(1)
									v165 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v81) << (uint(int64(52)) % 64))
									if base.Ui32(v81) <= base.Ui32(int32(19)) {
										v175 = base.F64_add(base.F64_sub(v159, v165), base.F64_sub(v80, v120))
									} else {
										v175 = base.F64_add(base.F64_sub(v80, base.F64_add(v120, v165)), v159)
									}
									v177 = base.F64_mul(v175, v144)
									v189 = v177
								}
							case 2:
								if base.F64_lt(v80, float64(-0.25)) != 0 {
									v189 = base.F64_mul(base.F64_sub(v120, base.F64_add(v80, float64(0.5))), float64(-2))
								} else {
									v135 = base.F64_sub(v80, v120)
									v189 = base.F64_add(base.F64_add(v135, v135), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v27) {
						v63 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v9, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v9)))
						v64 = base.F64_convert_i32_s(v63)
						v70 = v63
						v71 = base.F64_mul(v64, float64(1.9082149292705877e-10))
						v73 = base.F64_add(v9, base.F64_mul(v64, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v22 {
							v70 = int32(1)
							v71 = float64(1.9082149292705877e-10)
							v73 = base.F64_add(v9, float64(-0.6931471803691238))
						} else {
							v70 = int32(-1)
							v71 = float64(-1.9082149292705877e-10)
							v73 = base.F64_add(v9, float64(0.6931471803691238))
						}
					}
					v74 = base.F64_sub(v73, v71)
					v80 = v74
					v81 = v70
					v82 = base.F64_sub(base.F64_sub(v73, v74), v71)
					v85 = base.F64_mul(v80, float64(0.5))
					v86 = base.F64_mul(v80, v85)
					v102 = base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v105 = base.F64_sub(float64(3), base.F64_mul(v102, v85))
					v111 = base.F64_mul(v86, base.F64_div(base.F64_sub(v102, v105), base.F64_sub(float64(6), base.F64_mul(v80, v105))))
					if v81 == int32(0) {
						v189 = base.F64_sub(v80, base.F64_sub(base.F64_mul(v80, v111), v86))
					} else {
						v120 = base.F64_sub(base.F64_sub(base.F64_mul(v80, base.F64_sub(v111, v82)), v82), v86)
						switch v81 + int32(1) {
						case 0:
							v189 = base.F64_add(base.F64_mul(base.F64_sub(v80, v120), float64(0.5)), float64(-0.5))
						default:
							v144 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v81+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v81) {
								v149 = base.F64_add(base.F64_sub(v80, v120), float64(1))
								if v81 == int32(1024) {
									v156 = base.F64_mul(base.F64_add(v149, v149), float64(8.98846567431158e+307))
								} else {
									v156 = base.F64_mul(v149, v144)
								}
								v189 = base.F64_add(v156, float64(-1))
							} else {
								v159 = float64(1)
								v165 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v81) << (uint(int64(52)) % 64))
								if base.Ui32(v81) <= base.Ui32(int32(19)) {
									v175 = base.F64_add(base.F64_sub(v159, v165), base.F64_sub(v80, v120))
								} else {
									v175 = base.F64_add(base.F64_sub(v80, base.F64_add(v120, v165)), v159)
								}
								v177 = base.F64_mul(v175, v144)
								v189 = v177
							}
						case 2:
							if base.F64_lt(v80, float64(-0.25)) != 0 {
								v189 = base.F64_mul(base.F64_sub(v120, base.F64_add(v80, float64(0.5))), float64(-2))
							} else {
								v135 = base.F64_sub(v80, v120)
								v189 = base.F64_add(base.F64_add(v135, v135), float64(1))
							}
						}
					}
				}
			}
			v191 = float64(1)
			v192 = base.F64_add(v189, v191)
			v215 = base.F64_add(base.F64_div(base.F64_mul(v189, v189), base.F64_add(v192, v192)), v191)
		}
	} else {
		if base.Ui64(v10) <= base.Ui64(int64(4649454526309335039)) {
			v199 = F_exp(m, v9)
			mBase = m.M
			v215 = base.F64_mul(base.F64_add(v199, base.F64_div(float64(1), v199)), float64(0.5))
		} else {
			v206 = float64(2.247116418577895e+307)
			v210 = F_exp(m, base.F64_add(v9, float64(-1416.0996898839683)))
			mBase = m.M
			v215 = base.F64_mul(base.F64_mul(base.F64_mul(float64(1), v206), v210), v206)
		}
	}
	if base.F64_eq(v215, float64(0)) != 0 {
		F_float_underflow_error(m)
		mBase = m.M
		v221 = m.ExcPending
		if v221 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v222 = F_Float8GetDatum(m, v215)
		mBase = m.M
		v223 = m.ExcPending
		if v223 != 0 {
			return int32(0)
		} else {
			return v222
		}
	}
}
func F_debackslash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v8 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if int32(0) < l1 {
			v14 = l0
			v15 = l1
			v16 = v8
			for {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				v22 = int32(1)
				v24 = base.B2i32(v19 == int32(92)) & base.B2i32(v15 != v22)
				v25 = v14 + v24
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v26)
				v31 = v16 + v22
				v34 = v24 ^ int32(-1) + v15
				if int32(0) < v34 {
					v14 = v25 + v22
					v15 = v34
					v16 = v31
					continue
				} else {
					break
				}
				break
			}
			v39 = v31
		} else {
			v39 = v8
		}
		v42 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v42)
		return v8
	}
}
func F_decrypt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int64
	_ = v505
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	v11 = m.G0
	v13 = v11 - int32(208)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = int32(0)
	F_init_work(m, v13+int32(196), l1, l5, v13+int32(152))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(1)
	v26 = l2 + v25
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v31 = v29 & v25
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = v26
	goto L5
L4:
	;
	v32 = l2 + int32(4)
	goto L5
L5:
	;
	if v29 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v60 = F_mbuf_create_from_data(m, v32, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v38 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v49 = int32(1)
	if v31 != 0 {
		v59 = int32(base.Ui32(v29)>>(uint(v49)%32)) - v49
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v41 = int32(16)
	goto L12
L11:
	;
	v41 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v38-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = int32(4)
	goto L15
L14:
	;
	v48 = v41
	goto L15
L15:
	;
	v59 = v48
	goto L6
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v62 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v90 = F_mbuf_create(m, v87+int32(2048))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L29
	}
L19:
	;
	v66 = int32(18)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v68 == v66 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v79 = int32(1)
	if v62&v79 != 0 {
		v87 = int32(base.Ui32(v62) >> (uint(v79) % 32))
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v71 = v66
	goto L24
L23:
	;
	v71 = int32(2)
	goto L24
L24:
	;
	if base.Ui32((v68-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = int32(6)
	goto L27
L26:
	;
	v78 = v71
	goto L27
L27:
	;
	v87 = v78
	goto L18
L28:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v87 = int32(base.Ui32(v83) >> (uint(int32(2)) % 32))
	goto L18
L29:
	;
	v95 = F_mbuf_append(m, v90, v13+int32(204), int32(4))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	if l0 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v235 = int32(0)
	if v235 <= v231 {
		goto L89
	} else {
		goto L90
	}
L32:
	;
	if l4 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	goto L34
L34:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v183 = int32(1)
	v184 = l3 + v183
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v189 = v187 & v183
	if v189 != 0 {
		goto L71
	} else {
		goto L72
	}
L35:
	;
	v139 = int32(1)
	v140 = l3 + v139
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v145 = v143 & v139
	if v145 != 0 {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v99 = int32(0)
	v135 = v99
	v138 = v99
	goto L35
L37:
	;
	goto L38
L38:
	;
	v102 = l4 + int32(4)
	v103 = int32(1)
	v104 = l4 + v103
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v107 = v105 & v103
	if v105 == v103 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v107 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	if v107 != 0 {
		goto L51
	} else {
		goto L52
	}
L42:
	;
	v110 = v104
	goto L44
L43:
	;
	v110 = v102
	goto L44
L44:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v114 == int32(18) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v117 = int32(16)
	goto L47
L46:
	;
	v117 = int32(0)
	goto L47
L47:
	;
	if base.Ui32((v114-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v124 = int32(4)
	goto L50
L49:
	;
	v124 = v117
	goto L50
L50:
	;
	v135 = v110
	v138 = v124
	goto L35
L51:
	;
	v125 = int32(1)
	v135 = v104
	v138 = int32(base.Ui32(v105)>>(uint(v125)%32)) - v125
	goto L35
L52:
	;
	goto L53
L53:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v135 = v102
	v138 = int32(base.Ui32(v129)>>(uint(int32(2))%32)) - int32(4)
	goto L35
L54:
	;
	v146 = v140
	goto L56
L55:
	;
	v146 = l3 + int32(4)
	goto L56
L56:
	;
	if v143 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v174 = F_mbuf_create_from_data(m, v146, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L68
	}
L58:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140))))
	if v152 == int32(18) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L60
L60:
	;
	v163 = int32(1)
	if v145 != 0 {
		v173 = int32(base.Ui32(v143)>>(uint(v163)%32)) - v163
		goto L57
	} else {
		goto L67
	}
L61:
	;
	v155 = int32(16)
	goto L63
L62:
	;
	v155 = int32(0)
	goto L63
L63:
	;
	if base.Ui32((v152-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v162 = int32(4)
	goto L66
L65:
	;
	v162 = v155
	goto L66
L66:
	;
	v173 = v162
	goto L57
L67:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v173 = int32(base.Ui32(v167)>>(uint(int32(2))%32)) - int32(4)
	goto L57
L68:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v178 = F_pgp_set_pubkey(m, v176, v174, v135, v138, int32(1))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v180 = F_mbuf_free(m, v174)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v231 = v178
	goto L31
L71:
	;
	v190 = v184
	goto L73
L72:
	;
	v190 = l3 + int32(4)
	goto L73
L73:
	;
	if v187 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v218 = int32(0)
	if base.B2i32(v190 == v218)|base.B2i32(v217 <= v218) != 0 {
		goto L86
	} else {
		goto L87
	}
L75:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184))))
	if v196 == int32(18) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	v207 = int32(1)
	if v189 != 0 {
		v217 = int32(base.Ui32(v187)>>(uint(v207)%32)) - v207
		goto L74
	} else {
		goto L84
	}
L78:
	;
	v199 = int32(16)
	goto L80
L79:
	;
	v199 = int32(0)
	goto L80
L80:
	;
	if base.Ui32((v196-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v206 = int32(4)
	goto L83
L82:
	;
	v206 = v199
	goto L83
L83:
	;
	v217 = v206
	goto L74
L84:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v217 = int32(base.Ui32(v211)>>(uint(int32(2))%32)) - int32(4)
	goto L74
L85:
	;
	v231 = v229
	goto L31
L86:
	;
	v229 = int32(-13)
	goto L88
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+128)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v182)+124)) = v190
	v229 = int32(0)
	goto L88
L88:
	;
	goto L85
L89:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v239 = F_pgp_decrypt(m, v238, v60, v90)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L92
	}
L90:
	;
	v488 = v231
	v489 = v235
	goto L91
L91:
	;
	v490 = F_mbuf_free(m, v60)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L158
	}
L92:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
	if v241 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v483)+88))
	v488 = v239
	v489 = base.B2i32(v484 != int32(0))
	goto L91
L94:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v13)+160))
	if v245 < int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v13)+164))
	if v271 < int32(0) {
		goto L102
	} else {
		goto L103
	}
L96:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244)+60))
	if v245 == v248 {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v252 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v252 == int32(0) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v244)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = int32(_a_F_decrypt_internal_5)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(128))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(149), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	goto L95
L102:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
	if v297 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v244)+44))
	if v271 == v274 {
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v278 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v278 == int32(0) {
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v244)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v13)+116)) = v271
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = int32(_a_F_decrypt_internal_6)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(112))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(150), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	goto L102
L109:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v13)+176))
	if v323 < int32(0) {
		goto L116
	} else {
		goto L117
	}
L110:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	if v297 == v300 {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v304 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	if v304 == int32(0) {
		goto L109
	} else {
		goto L113
	}
L113:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v244)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v308
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = int32(_a_F_decrypt_internal_7)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(96))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(151), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	goto L109
L116:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v13)+184))
	if v349 < int32(0) {
		goto L123
	} else {
		goto L124
	}
L117:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v244)+52))
	if v323 == v326 {
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v330 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	if v330 == int32(0) {
		goto L116
	} else {
		goto L120
	}
L120:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v244)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v334
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v323
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = int32(_a_F_decrypt_internal_8)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(80))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(152), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	goto L116
L123:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v244)+76))
	if v375 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L124:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v244)+76))
	if v349 == v352 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v356 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v356 == int32(0) {
		goto L123
	} else {
		goto L127
	}
L127:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v244)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v360
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = int32(_a_F_decrypt_internal_9)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13-int32(-64))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(153), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	goto L123
L130:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
	if v405 < int32(0) {
		goto L138
	} else {
		goto L139
	}
L131:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v13)+172))
	if v378 < int32(0) {
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v244)+56))
	if v378 == v381 {
		goto L130
	} else {
		goto L133
	}
L133:
	;
	v385 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	if v385 == int32(0) {
		goto L130
	} else {
		goto L135
	}
L135:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v244)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v389
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v378
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(_a_F_decrypt_internal_10)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(48))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(155), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	goto L130
L138:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v13)+180))
	if v431 < int32(0) {
		goto L145
	} else {
		goto L146
	}
L139:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v244)+72))
	if v405 == v408 {
		goto L138
	} else {
		goto L140
	}
L140:
	;
	v412 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	if v412 == int32(0) {
		goto L138
	} else {
		goto L142
	}
L142:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v244)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v416
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v405
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(_a_F_decrypt_internal_11)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(32))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(156), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	goto L138
L145:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
	if v457 < int32(0) {
		goto L93
	} else {
		goto L152
	}
L146:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v244)+64))
	if v431 == v434 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v438 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	if v438 == int32(0) {
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v244)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v431
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_decrypt_internal_4)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13+int32(16))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(157), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	goto L145
L152:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v244)+88))
	if v457 == v460 {
		goto L93
	} else {
		goto L153
	}
L153:
	;
	v464 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	if v464 == int32(0) {
		goto L93
	} else {
		goto L155
	}
L155:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v244)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v468
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v457
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_decrypt_internal_0)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v13)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(158), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	goto L93
L158:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v493 = F_pgp_free(m, v492)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	if v488 == int32(0) {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v500 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+16)) = uint16(v500)
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(200)))) = v503
	v505 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v90)+8)) = v505
	*(*int64)(unsafe.Add(mBase, uint32(v90))) = v505
	goto L163
L161:
	;
	goto L162
L162:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decrypt_internal[0])) = int32(0)
	goto L210
L163:
	;
	v510 = F_mbuf_free(m, v90)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L164
	}
L164:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v512))) = (v502 - v503) << (uint(int32(2)) % 32)
	v516 = int32(0)
	if v489&base.B2i32(l1 != v516) == v516 {
		goto L166
	} else {
		goto L167
	}
L165:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decrypt_internal[0])) = int32(0)
	goto L209
L166:
	;
	v601 = v512
	goto L165
L167:
	;
	goto L168
L168:
	;
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_decrypt_internal[1]))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v522)+4))
	goto L169
L169:
	;
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	if v524 == int32(1) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v554 = int32(1)
	if v524&v554 != 0 {
		goto L181
	} else {
		goto L182
	}
L171:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+1)))
	if v530 == int32(18) {
		goto L174
	} else {
		goto L175
	}
L172:
	;
	goto L173
L173:
	;
	v541 = int32(1)
	if v524&v541 != 0 {
		v553 = int32(base.Ui32(v524)>>(uint(v541)%32)) - v541
		goto L170
	} else {
		goto L180
	}
L174:
	;
	v533 = int32(16)
	goto L176
L175:
	;
	v533 = int32(0)
	goto L176
L176:
	;
	if base.Ui32((v530-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v540 = int32(4)
	goto L179
L178:
	;
	v540 = v533
	goto L179
L179:
	;
	v553 = v540
	goto L170
L180:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v553 = int32(base.Ui32(v547)>>(uint(int32(2))%32)) - int32(4)
	goto L170
L181:
	;
	v558 = v554
	goto L183
L182:
	;
	v558 = int32(4)
	goto L183
L183:
	;
	v559 = v512 + v558
	v561 = F_pg_do_encoding_conversion(m, v559, v553, int32(6), v523)
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	if v561 == v559 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v601 = v512
	goto L165
L186:
	;
	goto L187
L187:
	;
	v564 = F_cstring_to_text(m, v561)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L188
	}
L188:
	;
	F_pfree(m, v561)
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L1
	} else {
		goto L189
	}
L189:
	;
	if v512 == v564 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v601 = v512
	goto L165
L191:
	;
	goto L192
L192:
	;
	v570 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512))))
	if v570 == int32(1) {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	if v595 != 0 {
		goto L205
	} else {
		goto L206
	}
L194:
	;
	v574 = int32(18)
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v512)+1)))
	if v576 == v574 {
		goto L197
	} else {
		goto L198
	}
L195:
	;
	goto L196
L196:
	;
	v587 = int32(1)
	if v570&v587 != 0 {
		v595 = int32(base.Ui32(v570) >> (uint(v587) % 32))
		goto L193
	} else {
		goto L203
	}
L197:
	;
	v579 = v574
	goto L199
L198:
	;
	v579 = int32(2)
	goto L199
L199:
	;
	if base.Ui32((v576-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v586 = int32(6)
	goto L202
L201:
	;
	v586 = v579
	goto L202
L202:
	;
	v595 = v586
	goto L193
L203:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v512)))
	v595 = int32(base.Ui32(v591) >> (uint(int32(2)) % 32))
	goto L193
L204:
	;
	F_pfree(m, v512)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L208
	}
L205:
	;
	base.MemoryFill(m, v512, int32(0), v595)
	goto L207
L206:
	;
	goto L207
L207:
	;
	goto L204
L208:
	;
	v601 = v564
	goto L165
L209:
	;
	m.G0 = v13 + int32(208)
	return v601
L210:
	;
	v613 = F_mbuf_free(m, v90)
	mBase = m.M
	v614 = m.ExcPending
	if v614 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	F_px_THROW_ERROR(m, v488)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_deparse_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v3 = l2
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v12 = v7 + int32(-16)
	F_initStringInfo(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+40)) = uint8(v3)
		v18 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v18
		v20 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v18
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+43)) = uint8(v18)
		v27 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v9)+41)) = uint16(v27)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v18
		*(*int64)(unsafe.Add(mBase, uint32(v9)+28)) = v20
		*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v12
		F_get_rule_expr(m, l0, v7+int32(-56), l3)
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+48))
			m.G0 = v9 - int32(-64)
			return v38
		}
	}
}
func F_deparse_lquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v16 = l0 + int32(16)
	v17 = int32(1)
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = v16
	v22 = v17
	v26 = v2
	goto L4
L2:
	;
	v58 = v17
	goto L3
L3:
	;
	v65 = F_palloc(m, v58)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	if v29 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v58 = v46
	goto L3
L6:
	;
	v53 = v26 + int32(1)
	if v53 != v18 {
		v20 = v20 + (v45+int32(7))&int32(_a_F_deparse_lquery_0)
		v22 = v46
		v26 = v53
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
	v31 = int32(2)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+2)))
	if v37&int32(32) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20))))
	v45 = v42
	v46 = v22 + int32(27)
	goto L6
L10:
	;
	v40 = int32(27)
	goto L12
L11:
	;
	v40 = v31
	goto L12
L12:
	;
	v45 = v30
	v46 = v30 + (v22 + v29<<(uint(v31)%32)) + v40
	goto L6
L13:
	;
	goto L5
L14:
	;
	return int32(0)
L15:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v69 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v71 = v65
	v72 = v16
	v79 = v2
	goto L19
L17:
	;
	v276 = v65
	goto L18
L18:
	;
	v285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v285)
	m.G0 = v13 - int32(-64)
	return v65
L19:
	;
	if v79 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v276 = v262
	goto L18
L21:
	;
	v80 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v80)
	v84 = v71 + int32(1)
	goto L23
L22:
	;
	v84 = v71
	goto L23
L23:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)))
	if v213&int32(32) == int32(0) {
		goto L61
	} else {
		goto L62
	}
L25:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+2)))
	if v86&int32(16) != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v199 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v199)
	v204 = v84 + int32(1)
	goto L24
L28:
	;
	v89 = int32(33)
	*(*uint8)(unsafe.Add(mBase, uint32(v84))) = uint8(v89)
	v92 = v84 + int32(1)
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v93 == int32(0) {
		v204 = v92
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v96 = v84
	goto L30
L30:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+20)))
	if v97 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v96 = v92
	goto L30
L32:
	;
	base.MemoryCopy(m, v96, v72+int32(23), v97)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+20)))
	v102 = v96 + v101
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	if v103&int32(4) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v106 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v106)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v111 = v102 + int32(1)
	v112 = v108
	goto L37
L36:
	;
	v111 = v102
	v112 = v103
	goto L37
L37:
	;
	if v112&int32(2) != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v115 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v115)
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+22)))
	v120 = v111 + int32(1)
	v121 = v117
	goto L40
L39:
	;
	v120 = v111
	v121 = v112
	goto L40
L40:
	;
	if v121&int32(1) != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v124 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v124)
	v128 = v120 + int32(1)
	goto L43
L42:
	;
	v128 = v120
	goto L43
L43:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if base.Ui32(v129) < base.Ui32(int32(2)) {
		v204 = v128
		goto L24
	} else {
		goto L44
	}
L44:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+20)))
	v142 = v128
	v144 = v72 + (v132+int32(7))&int32(_a_F_deparse_lquery_0) + int32(24)
	v147 = int32(1)
	goto L45
L45:
	;
	v151 = int32(124)
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v151)
	v154 = v142 + int32(1)
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+4)))
	if v155 != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v204 = v186
	goto L24
L47:
	;
	base.MemoryCopy(m, v154, v144+int32(7), v155)
	goto L49
L48:
	;
	goto L49
L49:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+4)))
	v160 = v154 + v159
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+6)))
	if v161&int32(4) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v164 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v160))) = uint8(v164)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+6)))
	v169 = v160 + int32(1)
	v170 = v166
	goto L52
L51:
	;
	v169 = v160
	v170 = v161
	goto L52
L52:
	;
	if v170&int32(2) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v173 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v169))) = uint8(v173)
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+6)))
	v178 = v169 + int32(1)
	v179 = v175
	goto L55
L54:
	;
	v178 = v169
	v179 = v170
	goto L55
L55:
	;
	if v179&int32(1) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v182 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v182)
	v186 = v178 + int32(1)
	goto L58
L57:
	;
	v186 = v178
	goto L58
L58:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v144)+4)))
	v196 = v147 + int32(1)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if base.Ui32(v196) < base.Ui32(v197) {
		v142 = v186
		v144 = v144 + (v187+int32(7))&int32(_a_F_deparse_lquery_0) + int32(8)
		v147 = v196
		goto L45
	} else {
		goto L59
	}
L59:
	;
	goto L46
L60:
	;
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	v272 = v79 + int32(1)
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if base.Ui32(v272) < base.Ui32(v273) {
		v71 = v262
		v72 = v72 + (v265+int32(7))&int32(_a_F_deparse_lquery_0)
		v79 = v272
		goto L19
	} else {
		goto L86
	}
L61:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v218 != 0 {
		v262 = v204
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+6)))
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+8)))
	if v219 == v220 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L63
L65:
	;
	v260 = F_strlen(m, v204)
	mBase = m.M
	v262 = v260 + v204
	goto L60
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v219
	v224 = F_pg_sprintf(m, v204, int32(_a_F_deparse_lquery_1), v13)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L14
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if v219 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L65
L70:
	;
	if v220 == int32(_a_F_deparse_lquery_2) {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L72
L72:
	;
	if v220 == int32(_a_F_deparse_lquery_2) {
		goto L81
	} else {
		goto L82
	}
L73:
	;
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72)+4)))
	if v230 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v220
	v243 = F_pg_sprintf(m, v204, int32(_a_F_deparse_lquery_3), v11+int32(-48))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L14
	} else {
		goto L80
	}
L76:
	;
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v204))) = uint8(v233)
	goto L65
L77:
	;
	goto L78
L78:
	;
	v237 = F_pg_sprintf(m, v204, int32(_a_F_deparse_lquery_4), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L14
	} else {
		goto L79
	}
L79:
	;
	goto L65
L80:
	;
	goto L65
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v219
	v251 = F_pg_sprintf(m, v204, int32(_a_F_deparse_lquery_5), v11+int32(-32))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L14
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v219
	v258 = F_pg_sprintf(m, v204, int32(_a_F_deparse_lquery_6), v11+int32(-16))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L14
	} else {
		goto L85
	}
L84:
	;
	goto L65
L85:
	;
	goto L65
L86:
	;
	goto L20
}
func F_dexp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v13 float64
	_ = v13
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 float64
	_ = v25
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v13 = float64(0)
			if base.F64_gt(v4, v13) != 0 {
				v16 = v4
			} else {
				v16 = v13
			}
			v17 = F_Float8GetDatum(m, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_dexp[0])) = int32(0)
			v25 = F_exp(m, v4)
			mBase = m.M
			if base.F64_eq(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v25, float64(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = v25
					v32 = F_Float8GetDatum(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return v32
					}
				}
			}
		}
	} else {
		v31 = v4
		v32 = F_Float8GetDatum(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			return v32
		}
	}
}
func F_difference(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v46 int32
	_ = v46
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var __phi69 int32
	_ = __phi69
	var v70 int32
	_ = v70
	var __phi70 int32
	_ = __phi70
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v73 int32
	_ = v73
	var __phi73 int32
	_ = __phi73
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v87 int32
	_ = v87
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
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v207 int32
	_ = v207
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var __phi230 int32
	_ = __phi230
	var v231 int32
	_ = v231
	var __phi231 int32
	_ = __phi231
	var v233 int32
	_ = v233
	var __phi233 int32
	_ = __phi233
	var v234 int32
	_ = v234
	var __phi234 int32
	_ = __phi234
	var v235 int32
	_ = v235
	var __phi235 int32
	_ = __phi235
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_pg_detoast_datum_packed(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_text_to_cstring(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = v12 + int32(11)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v178 = F_pg_detoast_datum_packed(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L41
	}
L5:
	;
	v61 = F_toupper(m, v32)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v61)
	v63 = int32(1)
	v65 = v12 + int32(12)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v66 != 0 {
		goto L14
	} else {
		goto L15
	}
L6:
	;
	v30 = v19
	v32 = v29
	goto L9
L7:
	;
	goto L8
L8:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v57)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v57
	goto L4
L9:
	;
	if base.Ui32(int32(229)) < base.Ui32((v32|int32(32)-int32(123))&int32(255)) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v46 != 0 {
		v30 = v30 + int32(1)
		v32 = v46
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v162))) = uint8(v167)
	goto L4
L14:
	;
	__phi69 = v30
	__phi70 = v66
	__phi72 = v65
	__phi73 = v63
	__phi74 = v30 + int32(1)
	v69 = __phi69
	v70 = __phi70
	v72 = __phi72
	v73 = __phi73
	v74 = __phi74
	goto L17
L15:
	;
	v149 = v65
	v150 = v63
	goto L16
L16:
	;
	v155 = int32(4) - v150
	if v155 != 0 {
		goto L38
	} else {
		goto L39
	}
L17:
	;
	if base.Ui32(int32(25)) < base.Ui32((v70|int32(32)-int32(97))&int32(255)) {
		v132 = v72
		v133 = v73
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if int32(3) < v133 {
		v162 = v132
		goto L13
	} else {
		goto L37
	}
L19:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	if v137 != 0 {
		goto L33
	} else {
		goto L34
	}
L20:
	;
	v87 = F_toupper(m, v70&int32(255))
	mBase = m.M
	v88 = base.I32_extend8_s(v87)
	v92 = base.B2i32(base.Ui32(int32(25)) < base.Ui32(v88-int32(65)))
	if v92 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v102 = F_toupper(m, v101)
	mBase = m.M
	v103 = base.I32_extend8_s(v102)
	if base.Ui32(v103-int32(65)) <= base.Ui32(int32(25)) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_difference[0]))))
	v98 = v97
	goto L21
L23:
	;
	goto L24
L24:
	;
	v98 = v87
	goto L21
L25:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+uint32(_c_F_difference[0]))))
	v111 = v110
	goto L27
L26:
	;
	v111 = v102
	goto L27
L27:
	;
	if v98&int32(255) == v111&int32(255) {
		v132 = v72
		v133 = v73
		goto L19
	} else {
		goto L28
	}
L28:
	;
	if v92 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+uint32(_c_F_difference[0]))))
	v120 = v119
	goto L31
L30:
	;
	v120 = v87
	goto L31
L31:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72))) = uint8(v120)
	if v120&int32(255) == int32(48) {
		v132 = v72
		v133 = v73
		goto L19
	} else {
		goto L32
	}
L32:
	;
	v126 = int32(1)
	v132 = v72 + v126
	v133 = v73 + v126
	goto L19
L33:
	;
	if v133 < int32(4) {
		__phi69 = v74
		__phi70 = v137
		__phi72 = v132
		__phi73 = v133
		__phi74 = v74 + int32(1)
		v69 = __phi69
		v70 = __phi70
		v72 = __phi72
		v73 = __phi73
		v74 = __phi74
		goto L17
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	goto L18
L36:
	;
	goto L35
L37:
	;
	v149 = v132
	v150 = v133
	goto L16
L38:
	;
	base.MemoryFill(m, v149, int32(48), v155)
	goto L40
L39:
	;
	goto L40
L40:
	;
	v162 = v155 + v149
	goto L13
L41:
	;
	v180 = F_text_to_cstring(m, v178)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v183 = v12 + int32(6)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	if v190 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v338 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	m.G0 = v12 + int32(16)
	return base.B2i32(v340 == v341) + base.B2i32(v338 == v339) + base.B2i32(v342 == v343) + base.B2i32(v344 == v345)
L44:
	;
	v222 = F_toupper(m, v193)
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, uint32(v183))) = uint8(v222)
	v224 = int32(1)
	v226 = v12 + int32(7)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v227 != 0 {
		goto L53
	} else {
		goto L54
	}
L45:
	;
	v191 = v180
	v193 = v190
	goto L48
L46:
	;
	goto L47
L47:
	;
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)) = uint8(v218)
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = v218
	goto L43
L48:
	;
	if base.Ui32(int32(229)) < base.Ui32((v193|int32(32)-int32(123))&int32(255)) {
		goto L44
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+1)))
	if v207 != 0 {
		v191 = v191 + int32(1)
		v193 = v207
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v328 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v323))) = uint8(v328)
	goto L43
L53:
	;
	__phi230 = v191
	__phi231 = v227
	__phi233 = v226
	__phi234 = v224
	__phi235 = v191 + int32(1)
	v230 = __phi230
	v231 = __phi231
	v233 = __phi233
	v234 = __phi234
	v235 = __phi235
	goto L56
L54:
	;
	v310 = v226
	v311 = v224
	goto L55
L55:
	;
	v316 = int32(4) - v311
	if v316 != 0 {
		goto L77
	} else {
		goto L78
	}
L56:
	;
	if base.Ui32(int32(25)) < base.Ui32((v231|int32(32)-int32(97))&int32(255)) {
		v293 = v233
		v294 = v234
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if int32(3) < v294 {
		v323 = v293
		goto L52
	} else {
		goto L76
	}
L58:
	;
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+1)))
	if v298 != 0 {
		goto L72
	} else {
		goto L73
	}
L59:
	;
	v248 = F_toupper(m, v231&int32(255))
	mBase = m.M
	v249 = base.I32_extend8_s(v248)
	v253 = base.B2i32(base.Ui32(int32(25)) < base.Ui32(v249-int32(65)))
	if v253 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230))))
	v263 = F_toupper(m, v262)
	mBase = m.M
	v264 = base.I32_extend8_s(v263)
	if base.Ui32(v264-int32(65)) <= base.Ui32(int32(25)) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_difference[0]))))
	v259 = v258
	goto L60
L62:
	;
	goto L63
L63:
	;
	v259 = v248
	goto L60
L64:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v264)+uint32(_c_F_difference[0]))))
	v272 = v271
	goto L66
L65:
	;
	v272 = v263
	goto L66
L66:
	;
	if v259&int32(255) == v272&int32(255) {
		v293 = v233
		v294 = v234
		goto L58
	} else {
		goto L67
	}
L67:
	;
	if v253 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_difference[0]))))
	v281 = v280
	goto L70
L69:
	;
	v281 = v248
	goto L70
L70:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v281)
	if v281&int32(255) == int32(48) {
		v293 = v233
		v294 = v234
		goto L58
	} else {
		goto L71
	}
L71:
	;
	v287 = int32(1)
	v293 = v233 + v287
	v294 = v234 + v287
	goto L58
L72:
	;
	if v294 < int32(4) {
		__phi230 = v235
		__phi231 = v298
		__phi233 = v293
		__phi234 = v294
		__phi235 = v235 + int32(1)
		v230 = __phi230
		v231 = __phi231
		v233 = __phi233
		v234 = __phi234
		v235 = __phi235
		goto L56
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	goto L57
L75:
	;
	goto L74
L76:
	;
	v310 = v293
	v311 = v294
	goto L55
L77:
	;
	base.MemoryFill(m, v310, int32(48), v316)
	goto L79
L78:
	;
	goto L79
L79:
	;
	v323 = v316 + v310
	goto L52
}
func F_dintdict_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_palloc0(m, int32(16))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
		if v16 != int32(1) {
			v28 = F_pnstrdup(m, v7, v6)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = v6
				v31 = v28
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				if v32 < v30 {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
					if v34 == int32(1) {
						F_pfree(m, v31)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
							return v10
						}
					} else {
						v43 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v31+v32))) = uint8(v43)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
						return v10
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
					return v10
				}
			}
		} else {
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			switch v19 - int32(43) {
			case 0, 2:
				v22 = int32(1)
				v25 = v6 - v22
				v26 = F_pnstrdup(m, v7+v22, v25)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v30 = v25
					v31 = v26
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v32 < v30 {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
						if v34 == int32(1) {
							F_pfree(m, v31)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
								return v10
							}
						} else {
							v43 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v31+v32))) = uint8(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
							return v10
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
						return v10
					}
				}
			default:
				v28 = F_pnstrdup(m, v7, v6)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v30 = v6
					v31 = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v32 < v30 {
						v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
						if v34 == int32(1) {
							F_pfree(m, v31)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = int32(0)
								return v10
							}
						} else {
							v43 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v31+v32))) = uint8(v43)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
							return v10
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
						return v10
					}
				}
			}
		}
	}
}
func F_dispell_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var __phi340 int32
	_ = __phi340
	var v342 int32
	_ = v342
	var __phi342 int32
	_ = __phi342
	var v344 int32
	_ = v344
	var __phi344 int32
	_ = __phi344
	var v348 int32
	_ = v348
	var __phi348 int32
	_ = __phi348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v359 int64
	_ = v359
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v390 int32
	_ = v390
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v12 <= v2 {
		v390 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v390
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = v15 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_str_tolower(m, v18, v12, int32(100))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v24 = int32(1)
	v26 = F_NormalizeSubWord(m, v17, v20, int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v28 == int32(0) {
		v73 = v2
		v74 = v2
		v79 = v24
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v86 = v2
	v87 = v2
	v92 = v24
	goto L8
L8:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+44)))
	if v96 != int32(1) {
		v326 = v87
		goto L23
	} else {
		goto L24
	}
L9:
	;
	F_pfree(m, v26)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v31 = v26
	v32 = v2
	v33 = v2
	v35 = v28
	v38 = v24
	goto L11
L11:
	;
	if v33 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v73 = v61
	v74 = v48
	v79 = v64
	goto L9
L13:
	;
	v45 = F_palloc(m, int32(_a_F_dispell_lexize_0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v47 = v32
	v48 = v33
	goto L15
L15:
	;
	v49 = v47 - v48
	if v49 <= int32(_a_F_dispell_lexize_1) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v47 = v45
	v48 = v45
	goto L15
L17:
	;
	v52 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+2)) = uint16(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v47)+12)) = v52
	*(*uint16)(unsafe.Add(mBase, uint32(v47))) = uint16(v38)
	v59 = v47 + int32(8)
	v61 = v59
	v62 = v59 - v48
	goto L19
L18:
	;
	v61 = v47
	v62 = v49
	goto L19
L19:
	;
	v64 = v38 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v65 == int32(0) {
		v73 = v61
		v74 = v48
		v79 = v64
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v62 < int32(_a_F_dispell_lexize_0) {
		v31 = v31 + int32(4)
		v32 = v61
		v33 = v48
		v35 = v65
		v38 = v64
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	v86 = v73
	v87 = v74
	v92 = v79
	goto L8
L23:
	;
	if v326 == int32(0) {
		v390 = v2
		goto L1
	} else {
		goto L78
	}
L24:
	;
	v99 = int32(0)
	v101 = F_strlen(m, v20)
	mBase = m.M
	v104 = F_SplitToVariants(m, v17, v99, v99, v20, v101, v99, int32(-1))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	if v104 == int32(0) {
		v326 = v87
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v108 = v104
	v109 = v86
	v110 = v87
	v115 = v92
	goto L27
L27:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if int32(2) <= v119 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v326 = v271
	goto L23
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+v119<<(uint(int32(2))%32)-int32(4))))
	v130 = F_NormalizeSubWord(m, v17, v128, int32(8))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	v270 = v109
	v271 = v110
	v274 = v119
	v276 = v115
	goto L31
L31:
	;
	v280 = int32(0)
	if v274 <= v280 {
		goto L68
	} else {
		goto L69
	}
L32:
	;
	if v130 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v132 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v258 = v109
	v259 = v110
	v264 = v115
	goto L35
L35:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v270 = v258
	v271 = v259
	v274 = v268
	v276 = v264
	goto L31
L36:
	;
	v133 = v132
	v134 = v109
	v135 = v110
	v137 = v130
	v140 = v115
	goto L39
L37:
	;
	v232 = v109
	v233 = v110
	v238 = v115
	goto L38
L38:
	;
	F_pfree(m, v130)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L3
	} else {
		goto L66
	}
L39:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if int32(0) < v144-int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v232 = v225
	v233 = v213
	v238 = v227
	goto L38
L41:
	;
	v150 = int32(0)
	v151 = v134
	v152 = v135
	goto L44
L42:
	;
	v196 = v133
	v197 = v134
	v198 = v135
	goto L43
L43:
	;
	if v198 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L44:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v161+v150<<(uint(int32(2))%32))))
	if v137 != v130 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v137)))
	v196 = v195
	v197 = v188
	v198 = v176
	goto L43
L46:
	;
	v167 = F_pstrdup(m, v165)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	v169 = v165
	goto L48
L48:
	;
	if v152 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v169 = v167
	goto L48
L50:
	;
	v173 = F_palloc(m, int32(_a_F_dispell_lexize_0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	v175 = v151
	v176 = v152
	goto L52
L52:
	;
	if v175-v176 <= int32(_a_F_dispell_lexize_1) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v175 = v173
	v176 = v173
	goto L52
L54:
	;
	v180 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+2)) = uint16(v180)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+4)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v175)+12)) = v180
	*(*uint16)(unsafe.Add(mBase, uint32(v175))) = uint16(v140)
	v188 = v175 + int32(8)
	goto L56
L55:
	;
	v188 = v175
	goto L56
L56:
	;
	v189 = int32(1)
	v190 = v150 + v189
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v190 < v191-v189 {
		v150 = v190
		v151 = v188
		v152 = v176
		goto L44
	} else {
		goto L57
	}
L57:
	;
	goto L45
L58:
	;
	v210 = F_palloc(m, int32(_a_F_dispell_lexize_0))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L61
	}
L59:
	;
	v212 = v197
	v213 = v198
	goto L60
L60:
	;
	if v212-v213 <= int32(_a_F_dispell_lexize_1) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v212 = v210
	v213 = v210
	goto L60
L62:
	;
	v217 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v212)+2)) = uint16(v217)
	*(*int32)(unsafe.Add(mBase, uint32(v212)+4)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v212)+12)) = v217
	*(*uint16)(unsafe.Add(mBase, uint32(v212))) = uint16(v140)
	v225 = v212 + int32(8)
	goto L64
L63:
	;
	v225 = v212
	goto L64
L64:
	;
	v227 = v140 + int32(1)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v137)+4))
	if v228 != 0 {
		v133 = v228
		v134 = v225
		v135 = v213
		v137 = v137 + int32(4)
		v140 = v227
		goto L39
	} else {
		goto L65
	}
L65:
	;
	goto L40
L66:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = int32(0)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v247+v248<<(uint(int32(2))%32)-int32(4))))
	F_pfree(m, v254)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v258 = v232
	v259 = v233
	v264 = v238
	goto L35
L68:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	F_pfree(m, v319)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L3
	} else {
		goto L75
	}
L69:
	;
	v283 = v280
	goto L70
L70:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294+v283<<(uint(int32(2))%32))))
	if v298 == int32(0) {
		goto L68
	} else {
		goto L72
	}
L71:
	;
	goto L68
L72:
	;
	F_pfree(m, v298)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v304 = v283 + int32(1)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	if v304 < v305 {
		v283 = v304
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	F_pfree(m, v108)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	if v318 != 0 {
		v108 = v318
		v109 = v270
		v110 = v271
		v115 = v276
		goto L27
	} else {
		goto L77
	}
L77:
	;
	goto L28
L78:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v326)+4))
	if v337 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	__phi340 = v326
	__phi342 = v326
	__phi344 = v326 + int32(4)
	__phi348 = v337
	v340 = __phi340
	v342 = __phi342
	v344 = __phi344
	v348 = __phi348
	goto L82
L80:
	;
	v371 = v326
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v371)+4)) = int32(0)
	v390 = v326
	goto L1
L82:
	;
	v351 = F_searchstoplist(m, v15, v348)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L85
	}
L83:
	;
	v371 = v363
	goto L81
L84:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v340)+12))
	if v366 != 0 {
		__phi340 = v340 + int32(8)
		__phi342 = v363
		__phi344 = v340 + int32(12)
		__phi348 = v366
		v340 = __phi340
		v342 = __phi342
		v344 = __phi344
		v348 = __phi348
		goto L82
	} else {
		goto L93
	}
L85:
	;
	if v351 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v344)))
	F_pfree(m, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v340 != v342 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = int32(0)
	v363 = v342
	goto L84
L90:
	;
	v359 = *(*int64)(unsafe.Add(mBase, uint32(v340)))
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = v359
	goto L92
L91:
	;
	goto L92
L92:
	;
	v363 = v342 + int32(8)
	goto L84
L93:
	;
	goto L83
}
func F_distance_chebyshev(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 float64
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 float64
	_ = v53
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v71 int32
	_ = v71
	var v72 float64
	_ = v72
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v88 int32
	_ = v88
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v111 int32
	_ = v111
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v142 float64
	_ = v142
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v168 float64
	_ = v168
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v193 float64
	_ = v193
	var v195 int32
	_ = v195
	var v205 float64
	_ = v205
	var v207 float64
	_ = v207
	var v210 int32
	_ = v210
	var v218 float64
	_ = v218
	var v219 float64
	_ = v219
	var v221 float64
	_ = v221
	var v223 int32
	_ = v223
	var v230 float64
	_ = v230
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	v2 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
			v29 = int32(2147483647)
			v30 = v28 & v29
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
			v33 = v31 & v29
			v34 = base.B2i32(base.Ui32(v30) < base.Ui32(v33))
			if base.Ui32(v30) < base.Ui32(v33) {
				v35 = v26
			} else {
				v35 = v21
			}
			if base.Ui32(v30) < base.Ui32(v33) {
				v36 = v21
			} else {
				v36 = v26
			}
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
			v39 = v37 & int32(2147483647)
			if v39 == int32(0) {
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v142 = v2
				v144 = v42
			} else {
				v43 = int32(8)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v53 = v2
				v58 = int32(0)
				for {
					v68 = v58 << (uint(int32(3)) % 32)
					v69 = v35 + v43 + v68
					v70 = *(*float64)(unsafe.Add(mBase, uint32(v69)))
					v71 = v68 + (v36 + v43)
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v71)))
					if int32(0) <= v47 {
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v69+v47<<(uint(int32(3))%32))))
						v79 = v78
					} else {
						v79 = v70
					}
					if int32(0) <= v37 {
						v85 = *(*float64)(unsafe.Add(mBase, uint32(v71+v37<<(uint(int32(3))%32))))
						v86 = v85
					} else {
						v86 = v72
					}
					v88 = int32(0)
					if base.B2i32(base.F64_le(v79, v86) == v88)|base.B2i32(base.F64_le(v70, v72) == v88)|(base.B2i32(base.F64_le(v79, v72) == v88)|base.B2i32(base.F64_ge(v86, v70) == v88)) == v88 {
						if base.F64_gt(v86, v72) != 0 {
							v105 = v72
						} else {
							v105 = v86
						}
						if base.F64_lt(v79, v70) != 0 {
							v107 = v70
						} else {
							v107 = v79
						}
						v130 = base.F64_sub(v105, v107)
					} else {
						v111 = int32(0)
						if base.B2i32(base.F64_gt(v79, v86) == v111)|base.B2i32(base.F64_gt(v70, v72) == v111)|(base.B2i32(base.F64_gt(v79, v72) == v111)|base.B2i32(base.F64_lt(v86, v70) == v111)) != 0 {
							v130 = float64(0)
						} else {
							if base.F64_gt(v79, v70) != 0 {
								v126 = v70
							} else {
								v126 = v79
							}
							if base.F64_lt(v86, v72) != 0 {
								v128 = v72
							} else {
								v128 = v86
							}
							v130 = base.F64_sub(v126, v128)
						}
					}
					v131 = base.F64_abs(v130)
					if base.F64_gt(v131, v53) != 0 {
						v133 = v131
					} else {
						v133 = v53
					}
					v135 = v58 + int32(1)
					if v135 != v39 {
						v53 = v133
						v58 = v135
						continue
					} else {
						break
					}
					break
				}
				v142 = v133
				v144 = v47
			}
			v157 = v144 & int32(2147483647)
			if base.Ui32(v39) < base.Ui32(v157) {
				v168 = v142
				v169 = v39
				for {
					v184 = v35 + int32(8) + v169<<(uint(int32(3))%32)
					v185 = *(*float64)(unsafe.Add(mBase, uint32(v184)))
					if base.B2i32(v144 < int32(0)) == int32(0) {
						v191 = *(*float64)(unsafe.Add(mBase, uint32(v184+v157<<(uint(int32(3))%32))))
						v192 = v191
					} else {
						v192 = v185
					}
					v193 = float64(0)
					v195 = int32(0)
					if base.B2i32(base.F64_le(v185, v193) == v195)|base.B2i32(base.F64_le(v192, v193) == v195) == v195 {
						if base.F64_lt(v192, v185) != 0 {
							v205 = v185
						} else {
							v205 = v192
						}
						v219 = base.F64_abs(v205)
					} else {
						v207 = float64(0)
						v210 = int32(0)
						if base.B2i32(base.F64_gt(v185, v207) == v210)|base.B2i32(base.F64_gt(v192, v207) == v210) != 0 {
							v219 = v207
						} else {
							if base.F64_gt(v192, v185) != 0 {
								v218 = v185
							} else {
								v218 = v192
							}
							v219 = v218
						}
					}
					if base.F64_gt(v219, v168) != 0 {
						v221 = v219
					} else {
						v221 = v168
					}
					v223 = v169 + int32(1)
					if v223 != v157 {
						v168 = v221
						v169 = v223
						continue
					} else {
						break
					}
					break
				}
				v230 = v221
			} else {
				v230 = v142
			}
			v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if base.Ui32(v30) < base.Ui32(v33) {
				if v244 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v248 = m.ExcPending
					if v248 != 0 {
						return int32(0)
					} else {
						v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 != v249 {
							F_pfree(m, v26)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								v258 = F_Float8GetDatum(m, v230)
								mBase = m.M
								v259 = m.ExcPending
								if v259 != 0 {
									return int32(0)
								} else {
									return v258
								}
							}
						} else {
							v258 = F_Float8GetDatum(m, v230)
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return int32(0)
							} else {
								return v258
							}
						}
					}
				} else {
					v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 != v249 {
						F_pfree(m, v26)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							v258 = F_Float8GetDatum(m, v230)
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return int32(0)
							} else {
								return v258
							}
						}
					} else {
						v258 = F_Float8GetDatum(m, v230)
						mBase = m.M
						v259 = m.ExcPending
						if v259 != 0 {
							return int32(0)
						} else {
							return v258
						}
					}
				}
			} else {
				if v244 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int32(0)
					} else {
						v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v26 == v254 {
							v258 = F_Float8GetDatum(m, v230)
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return int32(0)
							} else {
								return v258
							}
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int32(0)
							} else {
								v258 = F_Float8GetDatum(m, v230)
								mBase = m.M
								v259 = m.ExcPending
								if v259 != 0 {
									return int32(0)
								} else {
									return v258
								}
							}
						}
					}
				} else {
					v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v26 == v254 {
						v258 = F_Float8GetDatum(m, v230)
						mBase = m.M
						v259 = m.ExcPending
						if v259 != 0 {
							return int32(0)
						} else {
							return v258
						}
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int32(0)
						} else {
							v258 = F_Float8GetDatum(m, v230)
							mBase = m.M
							v259 = m.ExcPending
							if v259 != 0 {
								return int32(0)
							} else {
								return v258
							}
						}
					}
				}
			}
		}
	}
}
func F_div_var_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v96 int64
	_ = v96
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v401 int32
	_ = v401
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v457 int32
	_ = v457
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L10
	} else {
		goto L91
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = v35 - l2
	v40 = base.I32_div_s(l4+int32(3), int32(4))
	v43 = v36 + v40 + v34
	if v43 <= v34 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
	v28 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v28
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v28
	return
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	v46 = v34
	goto L14
L13:
	;
	v46 = v43
	goto L14
L14:
	;
	v47 = v46 + l5
	v52 = F_palloc(m, v47<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v54 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v52))) = uint16(v54)
	v57 = v52 + int32(2)
	v64 = l1 >> (uint(int32(31)) % 32)
	v66 = l1 ^ v64 - v64
	if base.Ui32(int32(_a_F_div_var_int_0)) <= base.Ui32(v66) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v166 != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	if v47 <= int32(0) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v47 <= int32(0) {
		goto L16
	} else {
		goto L27
	}
L20:
	;
	v71 = base.I64_extend_i32_u(v66)
	v74 = int32(0)
	v88 = int64(0)
	goto L21
L21:
	;
	v92 = v74 << (uint(int32(1)) % 32)
	if v74 < v19 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L16
L23:
	;
	v96 = int64(*(*int16)(unsafe.Add(mBase, uint32(v92+v32))))
	v98 = v96
	goto L25
L24:
	;
	v98 = int64(0)
	goto L25
L25:
	;
	v101 = v98 + v88*int64(10000)
	v102 = base.I64_div_u_s(v101, v71)
	*(*uint16)(unsafe.Add(mBase, uint32(v57+v92))) = uint16(v102)
	v107 = v74 + int32(1)
	if v107 != v47 {
		v74 = v107
		v88 = v101 - v71*v102
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v113 = int32(0)
	v120 = int32(0)
	goto L28
L28:
	;
	v131 = v113 << (uint(int32(1)) % 32)
	if v113 < v19 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L16
L30:
	;
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131+v32))))
	v137 = v135
	goto L32
L31:
	;
	v137 = int32(0)
	goto L32
L32:
	;
	v140 = v137 + v120*int32(_a_F_div_var_int_1)
	v141 = base.I32_div_u_s(v140, v66)
	*(*uint16)(unsafe.Add(mBase, uint32(v57+v131))) = uint16(v141)
	v146 = v113 + int32(1)
	if v146 != v47 {
		v113 = v146
		v120 = v140 - v66*v141
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	F_pfree(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L10
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
	if base.B2i32(v33 == v54)^base.B2i32(v54 < l1) != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	v174 = int32(_a_F_div_var_int_2)
	goto L40
L39:
	;
	v174 = int32(0)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v36
	if l5 != 0 {
		goto L45
	} else {
		goto L46
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v423
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v422
	return
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+4)) = int64(0)
	v422 = v401
	v423 = int32(0)
	goto L41
L43:
	;
	if int32(0) < v337 {
		goto L77
	} else {
		goto L78
	}
L44:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v336 = v334
	v337 = v335
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = l4
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v186 = l4 + v183<<(uint(int32(2))%32)
	if v186+int32(4) < int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = l4
	v304 = l4 + v36<<(uint(int32(2))%32)
	if v304+int32(4) <= int32(0) {
		v401 = v57
		goto L42
	} else {
		goto L74
	}
L48:
	;
	goto L44
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v197 = l4 & int32(3)
	v201 = base.I32_div_s(v186+int32(7), int32(4))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v202 <= v201 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	goto L48
L53:
	;
	if int32(0) <= v267 {
		goto L52
	} else {
		goto L73
	}
L54:
	;
	v247 = v241
	goto L67
L55:
	;
	v216 = int32(1)
	v217 = v201 - v216
	v220 = v195 + v217<<(uint(v216)%32)
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v220))))
	v222 = int32(2)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v197<<(uint(v222)%32))+uint32(_c_F_div_var_int[0])))
	v225 = base.I32_rem_s(v221, v224)
	v226 = v221 - v225
	*(*uint16)(unsafe.Add(mBase, uint32(v220))) = uint16(v226)
	v229 = base.I32_div_s(v224, v222)
	if v225 < v229 {
		v267 = v217
		goto L53
	} else {
		goto L62
	}
L56:
	;
	if base.B2i32(v197 == int32(0))|base.B2i32(v201 != v202) != 0 {
		goto L52
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v201
	if v197 != 0 {
		goto L55
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v201
	goto L55
L60:
	;
	v213 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195+v201<<(uint(int32(1))%32)))))
	if v213 <= int32(_a_F_div_var_int_3) {
		v267 = v201
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v241 = v201
	goto L54
L62:
	;
	v232 = v224 + base.I32_extend16_s(v226)
	if int32(_a_F_div_var_int_4) < v232 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v237 = v232 + int32(_a_F_div_var_int_5)
	goto L65
L64:
	;
	v237 = v232
	goto L65
L65:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v220))) = uint16(v237)
	if v232 < int32(_a_F_div_var_int_1) {
		v267 = v217
		goto L53
	} else {
		goto L66
	}
L66:
	;
	v241 = v217
	goto L54
L67:
	;
	v253 = int32(1)
	v254 = v247 - v253
	v257 = v195 + v254<<(uint(v253)%32)
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v257))))
	v262 = base.B2i32(int32(_a_F_div_var_int_6) < v260)
	if int32(_a_F_div_var_int_6) < v260 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v267 = v254
	goto L53
L69:
	;
	v263 = int32(-9999)
	goto L71
L70:
	;
	v263 = v253
	goto L71
L71:
	;
	v264 = v263 + v260
	*(*uint16)(unsafe.Add(mBase, uint32(v257))) = uint16(v264)
	if int32(_a_F_div_var_int_6) < v260 {
		v247 = v254
		goto L67
	} else {
		goto L72
	}
L72:
	;
	goto L68
L73:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v275 - int32(2)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v280 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v279 + v280
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v283 + v280
	goto L52
L74:
	;
	v312 = base.I32_div_s(v304+int32(7), int32(4))
	if v46 < v312 {
		goto L44
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v312
	v316 = l4 & int32(3)
	if v316 == int32(0) {
		v336 = v57
		v337 = v312
		goto L43
	} else {
		goto L76
	}
L76:
	;
	v322 = int32(2)
	v323 = v57 + v312<<(uint(int32(1))%32) - v322
	v324 = int32(*(*int16)(unsafe.Add(mBase, uint32(v323))))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v316<<(uint(v322)%32))+uint32(_c_F_div_var_int[0])))
	v328 = base.I32_rem_s(v324, v327)
	v329 = v324 - v328
	*(*uint16)(unsafe.Add(mBase, uint32(v323))) = uint16(v329)
	goto L44
L77:
	;
	v344 = v336
	v345 = v337
	goto L80
L78:
	;
	goto L79
L79:
	;
	if v337 != 0 {
		v422 = v336
		v423 = v337
		goto L41
	} else {
		goto L90
	}
L80:
	;
	v362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344))))
	if v362 != 0 {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	v401 = v336 + v337<<(uint(int32(1))%32)
	goto L42
L82:
	;
	v364 = v345
	goto L85
L83:
	;
	goto L84
L84:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v392 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v391 - v392
	if v392 < v345 {
		v344 = v344 + int32(2)
		v345 = v345 - v392
		goto L80
	} else {
		goto L89
	}
L85:
	;
	v386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v344+v364<<(uint(int32(1))%32)-int32(2)))))
	if v386 != 0 {
		v422 = v344
		v423 = v364
		goto L41
	} else {
		goto L87
	}
L87:
	;
	v387 = int32(1)
	if v387 < v364 {
		v364 = v364 - v387
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v401 = v344
	goto L42
L89:
	;
	goto L81
L90:
	;
	v401 = v336
	goto L42
L91:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(_a_F_div_var_int_7), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_div_var_int_8), int32(_a_F_div_var_int_9), int32(_a_F_div_var_int_10))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dmetaphone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_DoubleMetaphone(m, v12, v5+int32(8))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				if v18 != 0 {
					v20 = v18
				} else {
					v20 = int32(_a_F_dmetaphone_0)
				}
				v21 = F_cstring_to_text(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v21
				}
			}
		}
	}
}
func F_dmetaphone_alt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			F_DoubleMetaphone(m, v12, v5+int32(8))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				if v18 != 0 {
					v20 = v18
				} else {
					v20 = int32(_a_F_dmetaphone_alt_0)
				}
				v21 = F_cstring_to_text(m, v20)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 + int32(16)
					return v21
				}
			}
		}
	}
}
func F_do_setval(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
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
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int64
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v3 = l2
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	F_init_sequence(m, l0, v11+int32(108), v11+int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[0]))
	v24 = F_pg_class_aclcheck(m, v20, v22, int64(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L57
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L54
	}
L5:
	;
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = F_SearchSysCache1(m, int32(61), l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L50
	}
L9:
	;
	if v29 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+32))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+24)))
	if v41 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_PreventCommandIfReadOnly(m, int32(_a_F_do_setval_3))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_PreventCommandIfParallelMode(m, int32(_a_F_do_setval_3))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v54 = F_read_seq_tuple(m, v40, v11+int32(100), v11+int32(80))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if base.B2i32(l1 < v36)|base.B2i32(v37 < l1) != 0 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v3 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v65
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+118)))
	if v68 != int32(112) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v65 = v61
	goto L19
L21:
	;
	goto L22
L22:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v62)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = l1
	v65 = l1
	goto L19
L23:
	;
	v79 = int32(_a_F_do_setval_5)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_setval[1])) = v81 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = int64(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	F_MarkBufferDirty(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L31
	}
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[2]))
	if v72 <= int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	if v75 != 0 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v77 = F_GetTopTransactionId(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v76 != 0 {
		goto L23
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	goto L23
L31:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+118)))
	if v93 != int32(112) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v146 = int32(_a_F_do_setval_5)
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_do_setval[1])) = v148 - int32(1)
	F_UnlockReleaseBuffer(m, v89)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L48
	}
L33:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[2]))
	if v97 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	if v100 != 0 {
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	if v89 < int32(0) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v101 != 0 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L43
	}
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[3]))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105+(v89^int32(-1))<<(uint(int32(2))%32))))
	v119 = v111
	goto L39
L41:
	;
	goto L42
L42:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_do_setval[4]))
	v119 = v113 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L39
L43:
	;
	F_XLogRegisterBuffer(m, int32(0), v89, int32(6))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v126
	v128 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v128
	F_XLogRegisterData(m, v11-int32(-64), int32(12))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	F_XLogRegisterData(m, v135, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v141 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v119))) = base.I64_rotr(v141, int64(32))
	goto L32
L48:
	;
	F_relation_close(m, v40, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	m.G0 = v11 + int32(112)
	return
L50:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v168 + int32(4)
	F_errmsg(m, int32(_a_F_do_setval_6), v11+int32(48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_do_setval_1), int32(964), int32(_a_F_do_setval_2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_do_setval_0), v11)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_do_setval_1), int32(968), int32(_a_F_do_setval_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L57:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v202 + int32(4)
	F_errmsg(m, int32(_a_F_do_setval_4), v11+int32(16))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_do_setval_1), int32(993), int32(_a_F_do_setval_2))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dpi(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(3.141592653589793))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_dprintf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v10 = m.G0
	v11 = int32(144)
	v12 = v10 - v11
	m.G0 = v12
	base.MemoryFill(m, v12, int32(0), v11)
	v17 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = l0
	v20 = int32(_a_F_dprintf_0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(_a_F_dprintf_1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = v17
	v27 = F_vfprintf(m, v12, v20, l1)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return
	} else {
		m.G0 = v12 + int32(144)
		m.G0 = v7 + int32(16)
		return
	}
}
func F_drandom(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v117 int64
	_ = v117
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v156 float64
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_drandom[0])))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(16)
	v8 = int32(0)
	v12 = m.G0
	v14 = v12 - v7
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v8
	v20 = F_open(m, int32(_a_F_drandom_0), v8, v14)
	mBase = m.M
	if v20 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v130 = int32(_a_F_drandom_1)
	v133 = *(*int64)(unsafe.Add(mBase, _c_F_drandom[1]))
	v134 = *(*int64)(unsafe.Add(mBase, _c_F_drandom[2]))
	v135 = v133 ^ v134
	*(*int64)(unsafe.Add(mBase, _c_F_drandom[2])) = base.I64_rotl(v135, int64(37))
	*(*int64)(unsafe.Add(mBase, _c_F_drandom[1])) = v135<<(uint(int64(16))%64) ^ base.I64_rotl(v133, int64(24)) ^ v135
	v156 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v133*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L28
L4:
	;
	v128 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_drandom[0])) = uint8(v128)
	goto L3
L5:
	;
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	goto L10
L7:
	;
	v53 = v8
	goto L8
L8:
	;
	m.G0 = v14 + int32(16)
	goto L5
L9:
	;
	v48 = F_close(m, v20)
	mBase = m.M
	v53 = v46
	goto L8
L10:
	;
	v26 = int32(_a_F_drandom_1)
	v27 = v7
	goto L11
L11:
	;
	v32 = F_read(m, v20, v26, v27)
	mBase = m.M
	if v32 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v46 = int32(1)
	goto L9
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_drandom[3]))
	if v36 == int32(27) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v41 = v27 - v32
	if v41 != 0 {
		v26 = v26 + v32
		v27 = v41
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v46 = int32(0)
	goto L9
L17:
	;
	goto L12
L18:
	;
	v58 = int32(_a_F_drandom_1)
	v59 = *(*int64)(unsafe.Add(mBase, _c_F_drandom[1]))
	if v59 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v70 = int32(_a_F_drandom_1)
	v74 = m.G0
	v75 = int32(16)
	v76 = v74 - v75
	m.G0 = v76
	F_gettimeofday(m, v76)
	mBase = m.M
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	v80 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+8)))
	m.G0 = v76 + v75
	goto L26
L21:
	;
	goto L4
L22:
	;
	goto L21
L23:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _c_F_drandom[2]))
	if v62 != int64(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_drandom[2])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_drandom[1])) = int64(6364136223846793005)
	goto L22
L26:
	;
	v90 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_drandom[4])))
	v93 = v80 + v79*int64(1000000) - int64(946684800000000) ^ v90<<(uint(int64(32))%64)
	v96 = v93 + int64(4354685564936845354)
	v97 = int64(30)
	v100 = int64(-4658895280553007687)
	v101 = (int64(base.Ui64(v96)>>(uint(v97)%64)) ^ v96) * v100
	v102 = int64(27)
	v105 = int64(-7723592293110705685)
	v106 = (int64(base.Ui64(v101)>>(uint(v102)%64)) ^ v101) * v105
	v107 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_drandom[2])) = int64(base.Ui64(v106)>>(uint(v107)%64)) ^ v106
	v112 = v93 - int64(7046029254386353131)
	v117 = (int64(base.Ui64(v112)>>(uint(v97)%64)) ^ v112) * v100
	v122 = (int64(base.Ui64(v117)>>(uint(v102)%64)) ^ v117) * v105
	*(*int64)(unsafe.Add(mBase, _c_F_drandom[1])) = int64(base.Ui64(v122)>>(uint(v107)%64)) ^ v122
	goto L27
L27:
	;
	goto L4
L28:
	;
	v157 = F_Float8GetDatum(m, v156)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	return int32(0)
L30:
	;
	return v157
}
func F_dsimple_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc0(m, int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v19)
	if v13 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L44
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L40
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	m.G0 = v11 + int32(16)
	return v15
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v23 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = int32(0)
	v33 = v2
	v34 = v2
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v27<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v41 = int32(_a_F_dsimple_init_0)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsimple_init[0])))
	if base.B2i32(v44 == int32(0))|base.B2i32(v44 != v47) != 0 {
		v65 = v44
		v66 = v47
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L6
L11:
	;
	v109 = v27 + int32(1)
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v109 < v110 {
		v27 = v109
		v33 = v106
		v34 = v107
		goto L9
	} else {
		goto L35
	}
L12:
	;
	if v65-v66 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	goto L12
L14:
	;
	v50 = v40
	v51 = v41
	goto L15
L15:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	if v55 == int32(0) {
		v65 = v55
		v66 = v54
		goto L13
	} else {
		goto L17
	}
L16:
	;
	v65 = v55
	v66 = v54
	goto L13
L17:
	;
	v58 = int32(1)
	if v55 == v54 {
		v50 = v50 + v58
		v51 = v51 + v58
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	if v33 != 0 {
		goto L5
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v75 = int32(_a_F_dsimple_init_1)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dsimple_init[1])))
	if base.B2i32(v78 == int32(0))|base.B2i32(v78 != v81) != 0 {
		v99 = v78
		v100 = v81
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v70 = F_defGetString(m, v39)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_readstoplist(m, v70, v15)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v106 = int32(1)
	v107 = v34
	goto L11
L25:
	;
	if v99-v100 != 0 {
		goto L3
	} else {
		goto L32
	}
L26:
	;
	goto L25
L27:
	;
	v84 = v40
	v85 = v75
	goto L28
L28:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v89 == int32(0) {
		v99 = v89
		v100 = v88
		goto L26
	} else {
		goto L30
	}
L29:
	;
	v99 = v89
	v100 = v88
	goto L26
L30:
	;
	v92 = int32(1)
	if v89 == v88 {
		v84 = v84 + v92
		v85 = v85 + v92
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	if v34 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v102 = F_defGetBoolean(m, v39)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v102)
	v106 = v33
	v107 = int32(1)
	goto L11
L35:
	;
	goto L10
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_dsimple_init_2), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_dsimple_init_3), int32(50), int32(_a_F_dsimple_init_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_dsimple_init_5), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_dsimple_init_3), int32(59), int32(_a_F_dsimple_init_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v163
	F_errmsg(m, int32(_a_F_dsimple_init_6), v11)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_dsimple_init_3), int32(68), int32(_a_F_dsimple_init_4))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsynonym_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v10 <= v2 {
		v55 = v2
		m.G0 = v8 + int32(16)
		return v55
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 <= int32(0) {
			v55 = v2
			m.G0 = v8 + int32(16)
			return v55
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
			if v18 == int32(1) {
				v21 = F_pnstrdup(m, v17, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v28 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v36 = F_bsearch(m, v8, v32, v33, int32(16), int32(1148))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						F_pfree(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								v55 = v2
								m.G0 = v8 + int32(16)
								return v55
							} else {
								v44 = F_palloc0(m, int32(16))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
									v48 = F_pnstrdup(m, v46, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
										v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
										*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)) = uint16(v51)
										v55 = v44
										m.G0 = v8 + int32(16)
										return v55
									}
								}
							}
						}
					}
				}
			} else {
				v26 = F_str_tolower(m, v17, v10, int32(100))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v36 = F_bsearch(m, v8, v32, v33, int32(16), int32(1148))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						F_pfree(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								v55 = v2
								m.G0 = v8 + int32(16)
								return v55
							} else {
								v44 = F_palloc0(m, int32(16))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
									v48 = F_pnstrdup(m, v46, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
										v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
										*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)) = uint16(v51)
										v55 = v44
										m.G0 = v8 + int32(16)
										return v55
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_dt2time(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	v8 = base.I64_div_s(l0, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l1))) = uint32(v8)
	v13 = base.I64_extend32_s(v8)*int64(-3600000000) + l0
	v15 = base.I64_div_s(v13, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v15)
	v20 = base.I64_extend32_s(v15)*int64(-60000000) + v13
	v22 = base.I64_div_s(v20, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l3))) = uint32(v22)
	v26 = v22*int64(4293967296) + v20
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v26)
	return
}
func F_dtan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 float64
	_ = v34
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _c_F_dtan[0])) = int32(0)
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_dtan_0), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_dtan_1), int32(1980), int32(_a_F_dtan_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = m.G0
			v20 = v18 - int32(16)
			m.G0 = v20
			v27 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v4))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v27) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v27) < base.Ui32(int32(1044381696)) {
					v44 = v4
				} else {
					v34 = F___tan(m, v4, float64(0), int32(0))
					mBase = m.M
					v44 = v34
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v27) {
					v44 = base.F64_sub(v4, v4)
				} else {
					v38 = F___rem_pio2(m, v4, v20)
					mBase = m.M
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
					v43 = F___tan(m, v39, v40, v38&int32(1))
					mBase = m.M
					v44 = v43
				}
			}
			m.G0 = v20 + int32(16)
			v50 = v44
			v51 = F_Float8GetDatum(m, v50)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return v51
			}
		}
	} else {
		v50 = math.Float64frombits(uint64(0x7ff8000000000000))
		v51 = F_Float8GetDatum(m, v50)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			return v51
		}
	}
}
func F_dtand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v34 int64
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 float64
	_ = v43
	var v46 int64
	_ = v46
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int64
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int64
	_ = v69
	var v76 int32
	_ = v76
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v97 int32
	_ = v97
	var v99 int64
	_ = v99
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v115 int64
	_ = v115
	var v118 int32
	_ = v118
	var v120 int64
	_ = v120
	var v127 int64
	_ = v127
	var v131 int64
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v160 int64
	_ = v160
	var v168 float64
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 float64
	_ = v182
	var v185 int32
	_ = v185
	var v186 float64
	_ = v186
	var v190 float64
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v210 float64
	_ = v210
	var v214 int32
	_ = v214
	var v215 float64
	_ = v215
	var v216 float64
	_ = v216
	var v222 float64
	_ = v222
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v227 float64
	_ = v227
	var v229 float64
	_ = v229
	var v238 float64
	_ = v238
	var v246 float64
	_ = v246
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v266 float64
	_ = v266
	var v270 int32
	_ = v270
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v281 float64
	_ = v281
	var v284 float64
	_ = v284
	var v288 float64
	_ = v288
	var v292 float64
	_ = v292
	var v296 float64
	_ = v296
	var v302 float64
	_ = v302
	var v303 int32
	_ = v303
	var v308 float64
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v328 float64
	_ = v328
	var v332 int32
	_ = v332
	var v333 float64
	_ = v333
	var v334 float64
	_ = v334
	var v339 float64
	_ = v339
	var v341 float64
	_ = v341
	var v343 float64
	_ = v343
	var v346 float64
	_ = v346
	var v350 float64
	_ = v350
	var v354 float64
	_ = v354
	var v358 float64
	_ = v358
	var v367 float64
	_ = v367
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v380 int32
	_ = v380
	var v387 float64
	_ = v387
	var v391 int32
	_ = v391
	var v392 float64
	_ = v392
	var v393 float64
	_ = v393
	var v399 float64
	_ = v399
	var v400 float64
	_ = v400
	var v402 float64
	_ = v402
	var v404 float64
	_ = v404
	var v406 float64
	_ = v406
	var v415 float64
	_ = v415
	var v419 float64
	_ = v419
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v425 float64
	_ = v425
	var v428 float64
	_ = v428
	var v431 float64
	_ = v431
	var v437 float64
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
	if base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L127
	} else {
		goto L129
	}
L2:
	;
	if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v437 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L4
L4:
	;
	v438 = F_Float8GetDatum(m, v437)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L127
	} else {
		goto L128
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dtand[0])))
	if v22 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_init_degree_constants(m)
	mBase = m.M
	goto L8
L7:
	;
	goto L8
L8:
	;
	v26 = int32(0)
	v34 = base.I64_reinterpret_f64(v12)
	v38 = int32(2047)
	v39 = base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(int64(52))%64))) & v38
	if v39 == v38 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v170 = base.F64_lt(v168, float64(0))
	if v170 != 0 {
		goto L50
	} else {
		goto L51
	}
L10:
	;
	v43 = base.F64_mul(v12, float64(360))
	v168 = base.F64_div(v43, v43)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v46 = v34 << (uint(int64(1)) % 64)
	if base.Ui64(v46) <= base.Ui64(int64(-9156662467374350336)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v46 == int64(-9156662467374350336) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v39 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v53 = base.F64_mul(v12, float64(0))
	goto L18
L17:
	;
	v53 = v12
	goto L18
L18:
	;
	v168 = v53
	goto L9
L19:
	;
	if int32(1031) < v89 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v56 = int32(0)
	v58 = v34 << (uint(int64(12)) % 64)
	if int64(0) <= v58 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v89 = v39
	v90 = v34&int64(4503599627370495) | int64(4503599627370496)
	goto L19
L23:
	;
	v62 = v58
	v65 = v56
	goto L26
L24:
	;
	v76 = v56
	goto L25
L25:
	;
	v89 = v76
	v90 = v34 << (uint(base.I64_extend_i32_u(int32(1)-v76)) % 64)
	goto L19
L26:
	;
	v67 = v65 - int32(1)
	v69 = v62 << (uint(int64(1)) % 64)
	if int64(0) <= v69 {
		v62 = v69
		v65 = v67
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v76 = v67
	goto L25
L28:
	;
	goto L27
L29:
	;
	v94 = v90
	v97 = v89
	goto L32
L30:
	;
	v115 = v90
	v118 = v89
	goto L31
L31:
	;
	v120 = v115 - int64(6333186975989760)
	if v120 < int64(0) {
		v127 = v115
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v99 = v94 - int64(6333186975989760)
	if v99 < int64(0) {
		v106 = v94
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v115 = v108
	v118 = int32(1031)
	goto L31
L34:
	;
	v108 = v106 << (uint(int64(1)) % 64)
	v110 = v97 - int32(1)
	if int32(1031) < v110 {
		v94 = v108
		v97 = v110
		goto L32
	} else {
		goto L37
	}
L35:
	;
	if v99 != int64(0) {
		v106 = v99
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v168 = base.F64_mul(v12, float64(0))
	goto L9
L37:
	;
	goto L33
L38:
	;
	if base.Ui64(v127) <= base.Ui64(int64(4503599627370495)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v120 != int64(0) {
		v127 = v120
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v168 = base.F64_mul(v12, float64(0))
	goto L9
L41:
	;
	v131 = v127
	v134 = v118
	goto L44
L42:
	;
	v142 = v127
	v145 = v118
	goto L43
L43:
	;
	if int32(0) < v145 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v136 = v134 - int32(1)
	v138 = v131 << (uint(int64(1)) % 64)
	if base.Ui64(v131) < base.Ui64(int64(2251799813685248)) {
		v131 = v138
		v134 = v136
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v142 = v138
	v145 = v136
	goto L43
L46:
	;
	goto L45
L47:
	;
	v160 = v142 - int64(4503599627370496) | base.I64_extend_i32_u(v145)<<(uint(int64(52))%64)
	goto L49
L48:
	;
	v160 = int64(base.Ui64(v142) >> (uint(base.I64_extend_i32_u(int32(1)-v145)) % 64))
	goto L49
L49:
	;
	v168 = base.F64_reinterpret_i64(v34&int64(-9223372036854775807-1) | v160)
	goto L9
L50:
	;
	v171 = int32(-1)
	goto L52
L51:
	;
	v171 = int32(1)
	goto L52
L52:
	;
	if v170 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v174 = base.F64_neg(v168)
	goto L55
L54:
	;
	v174 = v168
	goto L55
L55:
	;
	v176 = base.F64_gt(v174, float64(180))
	if v176 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v177 = v26 - v171
	goto L58
L57:
	;
	v177 = v171
	goto L58
L58:
	;
	if v176 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v185 != 0 {
		goto L93
	} else {
		goto L94
	}
L60:
	;
	v182 = base.F64_sub(float64(360), v174)
	goto L62
L61:
	;
	v182 = v174
	goto L62
L62:
	;
	v185 = base.F64_gt(v182, float64(90))
	if v185 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v186 = base.F64_sub(float64(180), v182)
	goto L65
L64:
	;
	v186 = v182
	goto L65
L65:
	;
	if base.F64_le(v186, float64(30)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v190 = base.F64_mul(v186, float64(0.017453292519943295))
	v194 = m.G0
	v196 = v194 - int32(16)
	m.G0 = v196
	v203 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v190))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v203) <= base.Ui32(int32(1072243195)) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	v246 = base.F64_mul(base.F64_sub(float64(90), v186), float64(0.017453292519943295))
	v250 = m.G0
	v252 = v250 - int32(16)
	m.G0 = v252
	v259 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v246))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v259) <= base.Ui32(int32(1072243195)) {
		goto L84
	} else {
		goto L85
	}
L69:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v229
	v238 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[1]))
	v302 = base.F64_mul(base.F64_div(v229, v238), float64(0.5))
	goto L59
L70:
	;
	m.G0 = v196 + int32(16)
	goto L69
L71:
	;
	if base.Ui32(v203) < base.Ui32(int32(1045430272)) {
		v229 = v190
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v203) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v210 = F___sin(m, v190, float64(0), int32(0))
	mBase = m.M
	v229 = v210
	goto L70
L75:
	;
	v229 = base.F64_sub(v190, v190)
	goto L70
L76:
	;
	goto L77
L77:
	;
	v214 = F___rem_pio2(m, v190, v196)
	mBase = m.M
	v215 = *(*float64)(unsafe.Add(mBase, uint32(v196)+8))
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v196)))
	switch v214&int32(3) - int32(1) {
	case 0:
		goto L80
	case 1:
		goto L79
	case 2:
		goto L78
	default:
		goto L81
	}
L78:
	;
	v227 = F___cos(m, v216, v215)
	mBase = m.M
	v229 = base.F64_neg(v227)
	goto L70
L79:
	;
	v225 = F___sin(m, v216, v215, int32(1))
	mBase = m.M
	v229 = base.F64_neg(v225)
	goto L70
L80:
	;
	v223 = F___cos(m, v216, v215)
	mBase = m.M
	v229 = v223
	goto L70
L81:
	;
	v222 = F___sin(m, v216, v215, int32(1))
	mBase = m.M
	v229 = v222
	goto L70
L82:
	;
	v292 = base.F64_sub(float64(1), v288)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v292
	v296 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[2]))
	v302 = base.F64_add(base.F64_mul(base.F64_div(v292, v296), float64(-0.5)), float64(1))
	goto L59
L83:
	;
	m.G0 = v252 + int32(16)
	goto L82
L84:
	;
	if base.Ui32(v259) < base.Ui32(int32(1044816030)) {
		v288 = float64(1)
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v259) {
		v288 = base.F64_sub(v246, v246)
		goto L83
	} else {
		goto L88
	}
L87:
	;
	v266 = F___cos(m, v246, float64(0))
	mBase = m.M
	v288 = v266
	goto L83
L88:
	;
	v270 = F___rem_pio2(m, v246, v252)
	mBase = m.M
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v252)+8))
	v272 = *(*float64)(unsafe.Add(mBase, uint32(v252)))
	switch v270&int32(3) - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L92
	}
L89:
	;
	v284 = F___sin(m, v272, v271, int32(1))
	mBase = m.M
	v288 = v284
	goto L83
L90:
	;
	v281 = F___cos(m, v272, v271)
	mBase = m.M
	v288 = base.F64_neg(v281)
	goto L83
L91:
	;
	v279 = F___sin(m, v272, v271, int32(1))
	mBase = m.M
	v288 = base.F64_neg(v279)
	goto L83
L92:
	;
	v277 = F___cos(m, v272, v271)
	mBase = m.M
	v288 = v277
	goto L83
L93:
	;
	v303 = v26 - v177
	goto L95
L94:
	;
	v303 = v177
	goto L95
L95:
	;
	if base.F64_le(v186, float64(60)) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v420 = base.F64_div(v302, v419)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v420
	v422 = float64(0)
	v425 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[3]))
	v428 = base.F64_mul(base.F64_div(v420, v425), base.F64_convert_i32_s(v303))
	if base.F64_eq(v428, v422) != 0 {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	v308 = base.F64_mul(v186, float64(0.017453292519943295))
	v312 = m.G0
	v314 = v312 - int32(16)
	m.G0 = v314
	v321 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v308))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v321) <= base.Ui32(int32(1072243195)) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	goto L99
L99:
	;
	v367 = base.F64_mul(base.F64_sub(float64(90), v186), float64(0.017453292519943295))
	v371 = m.G0
	v373 = v371 - int32(16)
	m.G0 = v373
	v380 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v367))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v380) <= base.Ui32(int32(1072243195)) {
		goto L113
	} else {
		goto L114
	}
L100:
	;
	v354 = base.F64_sub(float64(1), v350)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v354
	v358 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[2]))
	v419 = base.F64_add(base.F64_mul(base.F64_div(v354, v358), float64(-0.5)), float64(1))
	goto L96
L101:
	;
	m.G0 = v314 + int32(16)
	goto L100
L102:
	;
	if base.Ui32(v321) < base.Ui32(int32(1044816030)) {
		v350 = float64(1)
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v321) {
		v350 = base.F64_sub(v308, v308)
		goto L101
	} else {
		goto L106
	}
L105:
	;
	v328 = F___cos(m, v308, float64(0))
	mBase = m.M
	v350 = v328
	goto L101
L106:
	;
	v332 = F___rem_pio2(m, v308, v314)
	mBase = m.M
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v314)+8))
	v334 = *(*float64)(unsafe.Add(mBase, uint32(v314)))
	switch v332&int32(3) - int32(1) {
	case 0:
		goto L109
	case 1:
		goto L108
	case 2:
		goto L107
	default:
		goto L110
	}
L107:
	;
	v346 = F___sin(m, v334, v333, int32(1))
	mBase = m.M
	v350 = v346
	goto L101
L108:
	;
	v343 = F___cos(m, v334, v333)
	mBase = m.M
	v350 = base.F64_neg(v343)
	goto L101
L109:
	;
	v341 = F___sin(m, v334, v333, int32(1))
	mBase = m.M
	v350 = base.F64_neg(v341)
	goto L101
L110:
	;
	v339 = F___cos(m, v334, v333)
	mBase = m.M
	v350 = v339
	goto L101
L111:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v406
	v415 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[1]))
	v419 = base.F64_mul(base.F64_div(v406, v415), float64(0.5))
	goto L96
L112:
	;
	m.G0 = v373 + int32(16)
	goto L111
L113:
	;
	if base.Ui32(v380) < base.Ui32(int32(1045430272)) {
		v406 = v367
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v380) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v387 = F___sin(m, v367, float64(0), int32(0))
	mBase = m.M
	v406 = v387
	goto L112
L117:
	;
	v406 = base.F64_sub(v367, v367)
	goto L112
L118:
	;
	goto L119
L119:
	;
	v391 = F___rem_pio2(m, v367, v373)
	mBase = m.M
	v392 = *(*float64)(unsafe.Add(mBase, uint32(v373)+8))
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v373)))
	switch v391&int32(3) - int32(1) {
	case 0:
		goto L122
	case 1:
		goto L121
	case 2:
		goto L120
	default:
		goto L123
	}
L120:
	;
	v404 = F___cos(m, v393, v392)
	mBase = m.M
	v406 = base.F64_neg(v404)
	goto L112
L121:
	;
	v402 = F___sin(m, v393, v392, int32(1))
	mBase = m.M
	v406 = base.F64_neg(v402)
	goto L112
L122:
	;
	v400 = F___cos(m, v393, v392)
	mBase = m.M
	v406 = v400
	goto L112
L123:
	;
	v399 = F___sin(m, v393, v392, int32(1))
	mBase = m.M
	v406 = v399
	goto L112
L124:
	;
	v431 = v422
	goto L126
L125:
	;
	v431 = v428
	goto L126
L126:
	;
	v437 = v431
	goto L4
L127:
	;
	return int32(0)
L128:
	;
	m.G0 = v9 + int32(16)
	return v438
L129:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(_a_F_dtand_0), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_dtand_1), int32(2512), int32(_a_F_dtand_2))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dtanh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v19 float64
	_ = v19
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v67 int32
	_ = v67
	var v68 float64
	_ = v68
	var v74 int32
	_ = v74
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v86 float64
	_ = v86
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v115 float64
	_ = v115
	var v124 float64
	_ = v124
	var v139 float64
	_ = v139
	var v148 float64
	_ = v148
	var v153 float64
	_ = v153
	var v160 float64
	_ = v160
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	var v179 float64
	_ = v179
	var v181 float64
	_ = v181
	var v193 float64
	_ = v193
	var v200 float64
	_ = v200
	var v207 int64
	_ = v207
	var v212 int32
	_ = v212
	var v248 int32
	_ = v248
	var v249 float64
	_ = v249
	var v255 int32
	_ = v255
	var v256 float64
	_ = v256
	var v258 float64
	_ = v258
	var v259 float64
	_ = v259
	var v265 float64
	_ = v265
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v270 float64
	_ = v270
	var v271 float64
	_ = v271
	var v287 float64
	_ = v287
	var v290 float64
	_ = v290
	var v296 float64
	_ = v296
	var v305 float64
	_ = v305
	var v320 float64
	_ = v320
	var v329 float64
	_ = v329
	var v334 float64
	_ = v334
	var v341 float64
	_ = v341
	var v344 float64
	_ = v344
	var v350 float64
	_ = v350
	var v360 float64
	_ = v360
	var v362 float64
	_ = v362
	var v374 float64
	_ = v374
	var v381 float64
	_ = v381
	var v388 int64
	_ = v388
	var v393 int32
	_ = v393
	var v429 int32
	_ = v429
	var v430 float64
	_ = v430
	var v436 int32
	_ = v436
	var v437 float64
	_ = v437
	var v439 float64
	_ = v439
	var v440 float64
	_ = v440
	var v446 float64
	_ = v446
	var v447 int32
	_ = v447
	var v448 float64
	_ = v448
	var v451 float64
	_ = v451
	var v452 float64
	_ = v452
	var v468 float64
	_ = v468
	var v471 float64
	_ = v471
	var v477 float64
	_ = v477
	var v486 float64
	_ = v486
	var v501 float64
	_ = v501
	var v510 float64
	_ = v510
	var v515 float64
	_ = v515
	var v522 float64
	_ = v522
	var v525 float64
	_ = v525
	var v531 float64
	_ = v531
	var v541 float64
	_ = v541
	var v543 float64
	_ = v543
	var v555 float64
	_ = v555
	var v560 float64
	_ = v560
	var v565 float64
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = base.F64_abs(v6)
	v8 = base.I64_reinterpret_f64(v7)
	if base.Ui64(int64(4603122931675955200)) <= base.Ui64(v8) {
		if base.Ui64(int64(4626322721511309312)) <= base.Ui64(v8) {
			v560 = base.F64_add(base.F64_div(math.Float64frombits(uint64(0x8000000000000000)), v7), float64(1))
		} else {
			v19 = base.F64_add(v7, v7)
			v26 = base.I64_reinterpret_f64(v19)
			v31 = base.I32_wrap_i64(int64(base.Ui64(v26)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v31) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
					v181 = v19
					v193 = v181
				} else {
					if v26 < int64(0) {
						v193 = float64(-1)
					} else {
						if base.F64_gt(v19, float64(709.782712893384)) == int32(0) {
							v67 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v19, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v19)))
							v68 = base.F64_convert_i32_s(v67)
							v74 = v67
							v75 = base.F64_mul(v68, float64(1.9082149292705877e-10))
							v77 = base.F64_add(v19, base.F64_mul(v68, float64(-0.6931471803691238)))
							v78 = base.F64_sub(v77, v75)
							v84 = v78
							v85 = v74
							v86 = base.F64_sub(base.F64_sub(v77, v78), v75)
							v89 = base.F64_mul(v84, float64(0.5))
							v90 = base.F64_mul(v84, v89)
							v106 = base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v109 = base.F64_sub(float64(3), base.F64_mul(v106, v89))
							v115 = base.F64_mul(v90, base.F64_div(base.F64_sub(v106, v109), base.F64_sub(float64(6), base.F64_mul(v84, v109))))
							if v85 == int32(0) {
								v193 = base.F64_sub(v84, base.F64_sub(base.F64_mul(v84, v115), v90))
							} else {
								v124 = base.F64_sub(base.F64_sub(base.F64_mul(v84, base.F64_sub(v115, v86)), v86), v90)
								switch v85 + int32(1) {
								case 0:
									v193 = base.F64_add(base.F64_mul(base.F64_sub(v84, v124), float64(0.5)), float64(-0.5))
								default:
									v148 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v85+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v85) {
										v153 = base.F64_add(base.F64_sub(v84, v124), float64(1))
										if v85 == int32(1024) {
											v160 = base.F64_mul(base.F64_add(v153, v153), float64(8.98846567431158e+307))
										} else {
											v160 = base.F64_mul(v153, v148)
										}
										v193 = base.F64_add(v160, float64(-1))
									} else {
										v163 = float64(1)
										v169 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v85) << (uint(int64(52)) % 64))
										if base.Ui32(v85) <= base.Ui32(int32(19)) {
											v179 = base.F64_add(base.F64_sub(v163, v169), base.F64_sub(v84, v124))
										} else {
											v179 = base.F64_add(base.F64_sub(v84, base.F64_add(v124, v169)), v163)
										}
										v181 = base.F64_mul(v179, v148)
										v193 = v181
									}
								case 2:
									if base.F64_lt(v84, float64(-0.25)) != 0 {
										v193 = base.F64_mul(base.F64_sub(v124, base.F64_add(v84, float64(0.5))), float64(-2))
									} else {
										v139 = base.F64_sub(v84, v124)
										v193 = base.F64_add(base.F64_add(v139, v139), float64(1))
									}
								}
							}
						} else {
							v193 = base.F64_mul(v19, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v31) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v31) < base.Ui32(int32(1016070144)) {
						v181 = v19
						v193 = v181
					} else {
						v84 = v19
						v85 = int32(0)
						v86 = float64(0)
						v89 = base.F64_mul(v84, float64(0.5))
						v90 = base.F64_mul(v84, v89)
						v106 = base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v109 = base.F64_sub(float64(3), base.F64_mul(v106, v89))
						v115 = base.F64_mul(v90, base.F64_div(base.F64_sub(v106, v109), base.F64_sub(float64(6), base.F64_mul(v84, v109))))
						if v85 == int32(0) {
							v193 = base.F64_sub(v84, base.F64_sub(base.F64_mul(v84, v115), v90))
						} else {
							v124 = base.F64_sub(base.F64_sub(base.F64_mul(v84, base.F64_sub(v115, v86)), v86), v90)
							switch v85 + int32(1) {
							case 0:
								v193 = base.F64_add(base.F64_mul(base.F64_sub(v84, v124), float64(0.5)), float64(-0.5))
							default:
								v148 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v85+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v85) {
									v153 = base.F64_add(base.F64_sub(v84, v124), float64(1))
									if v85 == int32(1024) {
										v160 = base.F64_mul(base.F64_add(v153, v153), float64(8.98846567431158e+307))
									} else {
										v160 = base.F64_mul(v153, v148)
									}
									v193 = base.F64_add(v160, float64(-1))
								} else {
									v163 = float64(1)
									v169 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v85) << (uint(int64(52)) % 64))
									if base.Ui32(v85) <= base.Ui32(int32(19)) {
										v179 = base.F64_add(base.F64_sub(v163, v169), base.F64_sub(v84, v124))
									} else {
										v179 = base.F64_add(base.F64_sub(v84, base.F64_add(v124, v169)), v163)
									}
									v181 = base.F64_mul(v179, v148)
									v193 = v181
								}
							case 2:
								if base.F64_lt(v84, float64(-0.25)) != 0 {
									v193 = base.F64_mul(base.F64_sub(v124, base.F64_add(v84, float64(0.5))), float64(-2))
								} else {
									v139 = base.F64_sub(v84, v124)
									v193 = base.F64_add(base.F64_add(v139, v139), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v31) {
						v67 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v19, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v19)))
						v68 = base.F64_convert_i32_s(v67)
						v74 = v67
						v75 = base.F64_mul(v68, float64(1.9082149292705877e-10))
						v77 = base.F64_add(v19, base.F64_mul(v68, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v26 {
							v74 = int32(1)
							v75 = float64(1.9082149292705877e-10)
							v77 = base.F64_add(v19, float64(-0.6931471803691238))
						} else {
							v74 = int32(-1)
							v75 = float64(-1.9082149292705877e-10)
							v77 = base.F64_add(v19, float64(0.6931471803691238))
						}
					}
					v78 = base.F64_sub(v77, v75)
					v84 = v78
					v85 = v74
					v86 = base.F64_sub(base.F64_sub(v77, v78), v75)
					v89 = base.F64_mul(v84, float64(0.5))
					v90 = base.F64_mul(v84, v89)
					v106 = base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, base.F64_add(base.F64_mul(v90, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v109 = base.F64_sub(float64(3), base.F64_mul(v106, v89))
					v115 = base.F64_mul(v90, base.F64_div(base.F64_sub(v106, v109), base.F64_sub(float64(6), base.F64_mul(v84, v109))))
					if v85 == int32(0) {
						v193 = base.F64_sub(v84, base.F64_sub(base.F64_mul(v84, v115), v90))
					} else {
						v124 = base.F64_sub(base.F64_sub(base.F64_mul(v84, base.F64_sub(v115, v86)), v86), v90)
						switch v85 + int32(1) {
						case 0:
							v193 = base.F64_add(base.F64_mul(base.F64_sub(v84, v124), float64(0.5)), float64(-0.5))
						default:
							v148 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v85+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v85) {
								v153 = base.F64_add(base.F64_sub(v84, v124), float64(1))
								if v85 == int32(1024) {
									v160 = base.F64_mul(base.F64_add(v153, v153), float64(8.98846567431158e+307))
								} else {
									v160 = base.F64_mul(v153, v148)
								}
								v193 = base.F64_add(v160, float64(-1))
							} else {
								v163 = float64(1)
								v169 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v85) << (uint(int64(52)) % 64))
								if base.Ui32(v85) <= base.Ui32(int32(19)) {
									v179 = base.F64_add(base.F64_sub(v163, v169), base.F64_sub(v84, v124))
								} else {
									v179 = base.F64_add(base.F64_sub(v84, base.F64_add(v124, v169)), v163)
								}
								v181 = base.F64_mul(v179, v148)
								v193 = v181
							}
						case 2:
							if base.F64_lt(v84, float64(-0.25)) != 0 {
								v193 = base.F64_mul(base.F64_sub(v124, base.F64_add(v84, float64(0.5))), float64(-2))
							} else {
								v139 = base.F64_sub(v84, v124)
								v193 = base.F64_add(base.F64_add(v139, v139), float64(1))
							}
						}
					}
				}
			}
			v560 = base.F64_sub(float64(1), base.F64_div(float64(2), base.F64_add(v193, float64(2))))
		}
	} else {
		if base.Ui64(int64(4598272728187797504)) <= base.Ui64(v8) {
			v200 = base.F64_add(v7, v7)
			v207 = base.I64_reinterpret_f64(v200)
			v212 = base.I32_wrap_i64(int64(base.Ui64(v207)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v212) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v200)&int64(9223372036854775807)) {
					v362 = v200
					v374 = v362
				} else {
					if v207 < int64(0) {
						v374 = float64(-1)
					} else {
						if base.F64_gt(v200, float64(709.782712893384)) == int32(0) {
							v248 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v200, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v200)))
							v249 = base.F64_convert_i32_s(v248)
							v255 = v248
							v256 = base.F64_mul(v249, float64(1.9082149292705877e-10))
							v258 = base.F64_add(v200, base.F64_mul(v249, float64(-0.6931471803691238)))
							v259 = base.F64_sub(v258, v256)
							v265 = v259
							v266 = v255
							v267 = base.F64_sub(base.F64_sub(v258, v259), v256)
							v270 = base.F64_mul(v265, float64(0.5))
							v271 = base.F64_mul(v265, v270)
							v287 = base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v290 = base.F64_sub(float64(3), base.F64_mul(v287, v270))
							v296 = base.F64_mul(v271, base.F64_div(base.F64_sub(v287, v290), base.F64_sub(float64(6), base.F64_mul(v265, v290))))
							if v266 == int32(0) {
								v374 = base.F64_sub(v265, base.F64_sub(base.F64_mul(v265, v296), v271))
							} else {
								v305 = base.F64_sub(base.F64_sub(base.F64_mul(v265, base.F64_sub(v296, v267)), v267), v271)
								switch v266 + int32(1) {
								case 0:
									v374 = base.F64_add(base.F64_mul(base.F64_sub(v265, v305), float64(0.5)), float64(-0.5))
								default:
									v329 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v266+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v266) {
										v334 = base.F64_add(base.F64_sub(v265, v305), float64(1))
										if v266 == int32(1024) {
											v341 = base.F64_mul(base.F64_add(v334, v334), float64(8.98846567431158e+307))
										} else {
											v341 = base.F64_mul(v334, v329)
										}
										v374 = base.F64_add(v341, float64(-1))
									} else {
										v344 = float64(1)
										v350 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v266) << (uint(int64(52)) % 64))
										if base.Ui32(v266) <= base.Ui32(int32(19)) {
											v360 = base.F64_add(base.F64_sub(v344, v350), base.F64_sub(v265, v305))
										} else {
											v360 = base.F64_add(base.F64_sub(v265, base.F64_add(v305, v350)), v344)
										}
										v362 = base.F64_mul(v360, v329)
										v374 = v362
									}
								case 2:
									if base.F64_lt(v265, float64(-0.25)) != 0 {
										v374 = base.F64_mul(base.F64_sub(v305, base.F64_add(v265, float64(0.5))), float64(-2))
									} else {
										v320 = base.F64_sub(v265, v305)
										v374 = base.F64_add(base.F64_add(v320, v320), float64(1))
									}
								}
							}
						} else {
							v374 = base.F64_mul(v200, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v212) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v212) < base.Ui32(int32(1016070144)) {
						v362 = v200
						v374 = v362
					} else {
						v265 = v200
						v266 = int32(0)
						v267 = float64(0)
						v270 = base.F64_mul(v265, float64(0.5))
						v271 = base.F64_mul(v265, v270)
						v287 = base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v290 = base.F64_sub(float64(3), base.F64_mul(v287, v270))
						v296 = base.F64_mul(v271, base.F64_div(base.F64_sub(v287, v290), base.F64_sub(float64(6), base.F64_mul(v265, v290))))
						if v266 == int32(0) {
							v374 = base.F64_sub(v265, base.F64_sub(base.F64_mul(v265, v296), v271))
						} else {
							v305 = base.F64_sub(base.F64_sub(base.F64_mul(v265, base.F64_sub(v296, v267)), v267), v271)
							switch v266 + int32(1) {
							case 0:
								v374 = base.F64_add(base.F64_mul(base.F64_sub(v265, v305), float64(0.5)), float64(-0.5))
							default:
								v329 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v266+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v266) {
									v334 = base.F64_add(base.F64_sub(v265, v305), float64(1))
									if v266 == int32(1024) {
										v341 = base.F64_mul(base.F64_add(v334, v334), float64(8.98846567431158e+307))
									} else {
										v341 = base.F64_mul(v334, v329)
									}
									v374 = base.F64_add(v341, float64(-1))
								} else {
									v344 = float64(1)
									v350 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v266) << (uint(int64(52)) % 64))
									if base.Ui32(v266) <= base.Ui32(int32(19)) {
										v360 = base.F64_add(base.F64_sub(v344, v350), base.F64_sub(v265, v305))
									} else {
										v360 = base.F64_add(base.F64_sub(v265, base.F64_add(v305, v350)), v344)
									}
									v362 = base.F64_mul(v360, v329)
									v374 = v362
								}
							case 2:
								if base.F64_lt(v265, float64(-0.25)) != 0 {
									v374 = base.F64_mul(base.F64_sub(v305, base.F64_add(v265, float64(0.5))), float64(-2))
								} else {
									v320 = base.F64_sub(v265, v305)
									v374 = base.F64_add(base.F64_add(v320, v320), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v212) {
						v248 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v200, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v200)))
						v249 = base.F64_convert_i32_s(v248)
						v255 = v248
						v256 = base.F64_mul(v249, float64(1.9082149292705877e-10))
						v258 = base.F64_add(v200, base.F64_mul(v249, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v207 {
							v255 = int32(1)
							v256 = float64(1.9082149292705877e-10)
							v258 = base.F64_add(v200, float64(-0.6931471803691238))
						} else {
							v255 = int32(-1)
							v256 = float64(-1.9082149292705877e-10)
							v258 = base.F64_add(v200, float64(0.6931471803691238))
						}
					}
					v259 = base.F64_sub(v258, v256)
					v265 = v259
					v266 = v255
					v267 = base.F64_sub(base.F64_sub(v258, v259), v256)
					v270 = base.F64_mul(v265, float64(0.5))
					v271 = base.F64_mul(v265, v270)
					v287 = base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, base.F64_add(base.F64_mul(v271, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v290 = base.F64_sub(float64(3), base.F64_mul(v287, v270))
					v296 = base.F64_mul(v271, base.F64_div(base.F64_sub(v287, v290), base.F64_sub(float64(6), base.F64_mul(v265, v290))))
					if v266 == int32(0) {
						v374 = base.F64_sub(v265, base.F64_sub(base.F64_mul(v265, v296), v271))
					} else {
						v305 = base.F64_sub(base.F64_sub(base.F64_mul(v265, base.F64_sub(v296, v267)), v267), v271)
						switch v266 + int32(1) {
						case 0:
							v374 = base.F64_add(base.F64_mul(base.F64_sub(v265, v305), float64(0.5)), float64(-0.5))
						default:
							v329 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v266+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v266) {
								v334 = base.F64_add(base.F64_sub(v265, v305), float64(1))
								if v266 == int32(1024) {
									v341 = base.F64_mul(base.F64_add(v334, v334), float64(8.98846567431158e+307))
								} else {
									v341 = base.F64_mul(v334, v329)
								}
								v374 = base.F64_add(v341, float64(-1))
							} else {
								v344 = float64(1)
								v350 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v266) << (uint(int64(52)) % 64))
								if base.Ui32(v266) <= base.Ui32(int32(19)) {
									v360 = base.F64_add(base.F64_sub(v344, v350), base.F64_sub(v265, v305))
								} else {
									v360 = base.F64_add(base.F64_sub(v265, base.F64_add(v305, v350)), v344)
								}
								v362 = base.F64_mul(v360, v329)
								v374 = v362
							}
						case 2:
							if base.F64_lt(v265, float64(-0.25)) != 0 {
								v374 = base.F64_mul(base.F64_sub(v305, base.F64_add(v265, float64(0.5))), float64(-2))
							} else {
								v320 = base.F64_sub(v265, v305)
								v374 = base.F64_add(base.F64_add(v320, v320), float64(1))
							}
						}
					}
				}
			}
			v560 = base.F64_div(v374, base.F64_add(v374, float64(2)))
		} else {
			if base.Ui64(v8) < base.Ui64(int64(4503599627370496)) {
				v560 = v7
			} else {
				v381 = base.F64_mul(v7, float64(-2))
				v388 = base.I64_reinterpret_f64(v381)
				v393 = base.I32_wrap_i64(int64(base.Ui64(v388)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1078159482)) <= base.Ui32(v393) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v381)&int64(9223372036854775807)) {
						v543 = v381
						v555 = v543
					} else {
						if v388 < int64(0) {
							v555 = float64(-1)
						} else {
							if base.F64_gt(v381, float64(709.782712893384)) == int32(0) {
								v429 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v381, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v381)))
								v430 = base.F64_convert_i32_s(v429)
								v436 = v429
								v437 = base.F64_mul(v430, float64(1.9082149292705877e-10))
								v439 = base.F64_add(v381, base.F64_mul(v430, float64(-0.6931471803691238)))
								v440 = base.F64_sub(v439, v437)
								v446 = v440
								v447 = v436
								v448 = base.F64_sub(base.F64_sub(v439, v440), v437)
								v451 = base.F64_mul(v446, float64(0.5))
								v452 = base.F64_mul(v446, v451)
								v468 = base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
								v471 = base.F64_sub(float64(3), base.F64_mul(v468, v451))
								v477 = base.F64_mul(v452, base.F64_div(base.F64_sub(v468, v471), base.F64_sub(float64(6), base.F64_mul(v446, v471))))
								if v447 == int32(0) {
									v555 = base.F64_sub(v446, base.F64_sub(base.F64_mul(v446, v477), v452))
								} else {
									v486 = base.F64_sub(base.F64_sub(base.F64_mul(v446, base.F64_sub(v477, v448)), v448), v452)
									switch v447 + int32(1) {
									case 0:
										v555 = base.F64_add(base.F64_mul(base.F64_sub(v446, v486), float64(0.5)), float64(-0.5))
									default:
										v510 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v447+int32(1023)) << (uint(int64(52)) % 64))
										if base.Ui32(int32(57)) <= base.Ui32(v447) {
											v515 = base.F64_add(base.F64_sub(v446, v486), float64(1))
											if v447 == int32(1024) {
												v522 = base.F64_mul(base.F64_add(v515, v515), float64(8.98846567431158e+307))
											} else {
												v522 = base.F64_mul(v515, v510)
											}
											v555 = base.F64_add(v522, float64(-1))
										} else {
											v525 = float64(1)
											v531 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v447) << (uint(int64(52)) % 64))
											if base.Ui32(v447) <= base.Ui32(int32(19)) {
												v541 = base.F64_add(base.F64_sub(v525, v531), base.F64_sub(v446, v486))
											} else {
												v541 = base.F64_add(base.F64_sub(v446, base.F64_add(v486, v531)), v525)
											}
											v543 = base.F64_mul(v541, v510)
											v555 = v543
										}
									case 2:
										if base.F64_lt(v446, float64(-0.25)) != 0 {
											v555 = base.F64_mul(base.F64_sub(v486, base.F64_add(v446, float64(0.5))), float64(-2))
										} else {
											v501 = base.F64_sub(v446, v486)
											v555 = base.F64_add(base.F64_add(v501, v501), float64(1))
										}
									}
								}
							} else {
								v555 = base.F64_mul(v381, float64(8.98846567431158e+307))
							}
						}
					}
				} else {
					if base.Ui32(v393) < base.Ui32(int32(1071001155)) {
						if base.Ui32(v393) < base.Ui32(int32(1016070144)) {
							v543 = v381
							v555 = v543
						} else {
							v446 = v381
							v447 = int32(0)
							v448 = float64(0)
							v451 = base.F64_mul(v446, float64(0.5))
							v452 = base.F64_mul(v446, v451)
							v468 = base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v471 = base.F64_sub(float64(3), base.F64_mul(v468, v451))
							v477 = base.F64_mul(v452, base.F64_div(base.F64_sub(v468, v471), base.F64_sub(float64(6), base.F64_mul(v446, v471))))
							if v447 == int32(0) {
								v555 = base.F64_sub(v446, base.F64_sub(base.F64_mul(v446, v477), v452))
							} else {
								v486 = base.F64_sub(base.F64_sub(base.F64_mul(v446, base.F64_sub(v477, v448)), v448), v452)
								switch v447 + int32(1) {
								case 0:
									v555 = base.F64_add(base.F64_mul(base.F64_sub(v446, v486), float64(0.5)), float64(-0.5))
								default:
									v510 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v447+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v447) {
										v515 = base.F64_add(base.F64_sub(v446, v486), float64(1))
										if v447 == int32(1024) {
											v522 = base.F64_mul(base.F64_add(v515, v515), float64(8.98846567431158e+307))
										} else {
											v522 = base.F64_mul(v515, v510)
										}
										v555 = base.F64_add(v522, float64(-1))
									} else {
										v525 = float64(1)
										v531 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v447) << (uint(int64(52)) % 64))
										if base.Ui32(v447) <= base.Ui32(int32(19)) {
											v541 = base.F64_add(base.F64_sub(v525, v531), base.F64_sub(v446, v486))
										} else {
											v541 = base.F64_add(base.F64_sub(v446, base.F64_add(v486, v531)), v525)
										}
										v543 = base.F64_mul(v541, v510)
										v555 = v543
									}
								case 2:
									if base.F64_lt(v446, float64(-0.25)) != 0 {
										v555 = base.F64_mul(base.F64_sub(v486, base.F64_add(v446, float64(0.5))), float64(-2))
									} else {
										v501 = base.F64_sub(v446, v486)
										v555 = base.F64_add(base.F64_add(v501, v501), float64(1))
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(1072734897)) < base.Ui32(v393) {
							v429 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v381, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v381)))
							v430 = base.F64_convert_i32_s(v429)
							v436 = v429
							v437 = base.F64_mul(v430, float64(1.9082149292705877e-10))
							v439 = base.F64_add(v381, base.F64_mul(v430, float64(-0.6931471803691238)))
						} else {
							if int64(0) <= v388 {
								v436 = int32(1)
								v437 = float64(1.9082149292705877e-10)
								v439 = base.F64_add(v381, float64(-0.6931471803691238))
							} else {
								v436 = int32(-1)
								v437 = float64(-1.9082149292705877e-10)
								v439 = base.F64_add(v381, float64(0.6931471803691238))
							}
						}
						v440 = base.F64_sub(v439, v437)
						v446 = v440
						v447 = v436
						v448 = base.F64_sub(base.F64_sub(v439, v440), v437)
						v451 = base.F64_mul(v446, float64(0.5))
						v452 = base.F64_mul(v446, v451)
						v468 = base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, base.F64_add(base.F64_mul(v452, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v471 = base.F64_sub(float64(3), base.F64_mul(v468, v451))
						v477 = base.F64_mul(v452, base.F64_div(base.F64_sub(v468, v471), base.F64_sub(float64(6), base.F64_mul(v446, v471))))
						if v447 == int32(0) {
							v555 = base.F64_sub(v446, base.F64_sub(base.F64_mul(v446, v477), v452))
						} else {
							v486 = base.F64_sub(base.F64_sub(base.F64_mul(v446, base.F64_sub(v477, v448)), v448), v452)
							switch v447 + int32(1) {
							case 0:
								v555 = base.F64_add(base.F64_mul(base.F64_sub(v446, v486), float64(0.5)), float64(-0.5))
							default:
								v510 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v447+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v447) {
									v515 = base.F64_add(base.F64_sub(v446, v486), float64(1))
									if v447 == int32(1024) {
										v522 = base.F64_mul(base.F64_add(v515, v515), float64(8.98846567431158e+307))
									} else {
										v522 = base.F64_mul(v515, v510)
									}
									v555 = base.F64_add(v522, float64(-1))
								} else {
									v525 = float64(1)
									v531 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v447) << (uint(int64(52)) % 64))
									if base.Ui32(v447) <= base.Ui32(int32(19)) {
										v541 = base.F64_add(base.F64_sub(v525, v531), base.F64_sub(v446, v486))
									} else {
										v541 = base.F64_add(base.F64_sub(v446, base.F64_add(v486, v531)), v525)
									}
									v543 = base.F64_mul(v541, v510)
									v555 = v543
								}
							case 2:
								if base.F64_lt(v446, float64(-0.25)) != 0 {
									v555 = base.F64_mul(base.F64_sub(v486, base.F64_add(v446, float64(0.5))), float64(-2))
								} else {
									v501 = base.F64_sub(v446, v486)
									v555 = base.F64_add(base.F64_add(v501, v501), float64(1))
								}
							}
						}
					}
				}
				v560 = base.F64_div(base.F64_neg(v555), base.F64_add(v555, float64(2)))
			}
		}
	}
	if base.I64_reinterpret_f64(v6) < int64(0) {
		v565 = base.F64_neg(v560)
	} else {
		v565 = v560
	}
	if base.F64_eq(base.F64_abs(v565), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v572 = m.ExcPending
		if v572 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v573 = F_Float8GetDatum(m, v565)
		mBase = m.M
		v574 = m.ExcPending
		if v574 != 0 {
			return int32(0)
		} else {
			return v573
		}
	}
}
func F_dtof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = base.F32_demote_f64(v5)
	if base.F32_eq(base.F32_abs(v6), math.Float32frombits(uint32(0x7f800000)))&base.F64_ne(base.F64_abs(v5), math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
		if base.F32_eq(v6, float32(0))&base.F64_ne(v5, float64(0)) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.I32_reinterpret_f32(v6)
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dump_cursor_direction(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = int32(_a_F_dump_cursor_direction_0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_dump_cursor_direction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dump_cursor_direction[0])) = v10 + int32(2)
	if int32(-1) <= v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v32 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	default:
		goto L10
	}
L4:
	;
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_1), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
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
	v25 = v18 + int32(1)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_dump_cursor_direction[0]))
	if v25 < v27 {
		v18 = v25
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v55 != 0 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = v32
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_2), v4+int32(-16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L6
	} else {
		goto L19
	}
L11:
	;
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_3), int32(0))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L18
	}
L12:
	;
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_4), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L17
	}
L13:
	;
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_5), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_6), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	goto L9
L17:
	;
	goto L9
L18:
	;
	goto L9
L19:
	;
	goto L9
L20:
	;
	v89 = int32(_a_F_dump_cursor_direction_0)
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_dump_cursor_direction[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_dump_cursor_direction[0])) = v91 - int32(2)
	m.G0 = v6 - int32(-64)
	return
L21:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v56
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_7), v4+int32(-32))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v82
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_8), v6)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L6
	} else {
		goto L33
	}
L24:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	if int32(0) <= v63 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v63
	if v66 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_9), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L32
	}
L28:
	;
	v70 = int32(_a_F_dump_cursor_direction_10)
	goto L30
L29:
	;
	v70 = int32(_a_F_dump_cursor_direction_11)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v70
	F_pg_printf(m, int32(_a_F_dump_cursor_direction_12), v4+int32(-48))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	goto L20
L33:
	;
	goto L20
}
func F_dumptuples(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v14 <= v13 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L21
	} else {
		goto L58
	}
L2:
	;
	m.G0 = v11 - int32(-64)
	return
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v13 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if v16 < int64(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v19&int32(1) == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	goto L3
L11:
	;
	v76 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v72 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dumptuples[0])))
	if v80 != v76 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if int32(0) < v26 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v26 == int32(2147483647) {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v72 = v26
	v73 = l0 + int32(168)
	goto L11
L16:
	;
	v36 = l0 + int32(168)
	if v26 <= int32(0) {
		v72 = v26
		v73 = v36
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v39 < v40 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v43 = F_LogicalTapeCreate(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v63 = base.I32_rem_s(v62, v39)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v62 + int32(1)
	v72 = v26
	v73 = v36
	goto L11
L21:
	;
	return
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v46+v47<<(uint(int32(2))%32)))) = v43
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v52 + v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v56 + v53
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v72 = v60
	v73 = v36
	goto L11
L23:
	;
	F_tuplesort_sort_memtuples(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L21
	} else {
		goto L30
	}
L24:
	;
	v85 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if v85 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v93 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v89
	F_errmsg_internal(m, int32(_a_F_dumptuples_0), v9+int32(-32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2351), int32(_a_F_dumptuples_2))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dumptuples[0])))
	if v113 != int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v143 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v118 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if v118 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v126 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v122
	F_errmsg_internal(m, int32(_a_F_dumptuples_3), v9+int32(-48))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2362), int32(_a_F_dumptuples_2))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	goto L31
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_MemoryContextReset(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L50
	}
L39:
	;
	v146 = int32(0)
	if v143 != int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v155 = v146
	v157 = int32(0)
	goto L43
L41:
	;
	v186 = v146
	goto L42
L42:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v198].(func(*base.Module, int32, int32, int32))(m, l0, v193, v194+v186<<(uint(int32(4))%32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L21
	} else {
		goto L49
	}
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v164 = v155 << (uint(int32(4)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v167].(func(*base.Module, int32, int32, int32))(m, l0, v162, v164+v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L45
	}
L44:
	;
	if v143&int32(1) == int32(0) {
		goto L38
	} else {
		goto L48
	}
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v175].(func(*base.Module, int32, int32, int32))(m, l0, v170, v171+v164+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v178 = int32(2)
	v179 = v155 + v178
	v181 = v157 + v178
	if v181 != v143&int32(-2) {
		v155 = v179
		v157 = v181
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v186 = v179
	goto L42
L49:
	;
	goto L38
L50:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	v217 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v214 + v217
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	F_LogicalTapeWrite(m, v220, v9+int32(-4), int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dumptuples[0])))
	if v229 != int32(1) {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v234 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	if v234 == int32(0) {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v243 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v239
	v248 = int32(1)
	v250 = base.I32_rem_s(v240-v248, v238)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v250 + v248
	F_errmsg_internal(m, int32(_a_F_dumptuples_4), v11)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2395), int32(_a_F_dumptuples_2))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	goto L2
L58:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(2147483647)
	F_errmsg(m, int32(_a_F_dumptuples_5), v9+int32(-16))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2341), int32(_a_F_dumptuples_2))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dup2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	if l0 != l1 {
		for {
			v9 = m.Env.X__syscall_dup3(m, l0, l1, int32(0))
			mBase = m.M
			if v9 == int32(-10) {
				continue
			} else {
				break
			}
			break
		}
		v30 = v9
		if base.Ui32(int32(-4095)) <= base.Ui32(v30) {
			*(*int32)(unsafe.Add(mBase, _c_F_dup2[0])) = int32(0) - v30
			v38 = int32(-1)
		} else {
			v38 = v30
		}
		v39 = v38
	} else {
		v12 = m.G0
		v14 = v12 - int32(32)
		m.G0 = v14
		v18 = m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, v14+int32(8))
		mBase = m.M
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F_dup2[0])) = v18
			v23 = int32(0)
		} else {
			v23 = int32(1)
		}
		m.G0 = v14 + int32(32)
		if v23 != 0 {
			v39 = l0
		} else {
			v30 = int32(-8)
			if base.Ui32(int32(-4095)) <= base.Ui32(v30) {
				*(*int32)(unsafe.Add(mBase, _c_F_dup2[0])) = int32(0) - v30
				v38 = int32(-1)
			} else {
				v38 = v30
			}
			v39 = v38
		}
	}
	return v39
}
func F_duptraverse(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = m.T0[v7].(func(*base.Module) int32)(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(101)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v16 = v14
	goto L8
L7:
	;
	v16 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
	return
L9:
	;
	return
L10:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	goto L11
L13:
	;
	goto L14
L14:
	;
	v20 = F_newstate(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v20
	if v20 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	v31 = v26
	goto L18
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v33 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	goto L9
L20:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	F_duptraverse(m, l0, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v39 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	F_cparc(m, l0, v31, v40, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v45 != 0 {
		v31 = v45
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L19
}
func F_dutch_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v369 int32
	_ = v369
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v643 int32
	_ = v643
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v765 int32
	_ = v765
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v893 int32
	_ = v893
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v926 int32
	_ = v926
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1015 int32
	_ = v1015
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1049 int32
	_ = v1049
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1154 int32
	_ = v1154
	var v1156 int32
	_ = v1156
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1214 int32
	_ = v1214
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1264 int32
	_ = v1264
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1282 int32
	_ = v1282
	var v1285 int32
	_ = v1285
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1323 int32
	_ = v1323
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1365 int32
	_ = v1365
	var v1366 int32
	_ = v1366
	var v1368 int32
	_ = v1368
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1380 int32
	_ = v1380
	var v1383 int32
	_ = v1383
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1425 int32
	_ = v1425
	var v1429 int32
	_ = v1429
	var v1432 int32
	_ = v1432
	var v1433 int32
	_ = v1433
	var v1436 int32
	_ = v1436
	var v1437 int32
	_ = v1437
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1459 int32
	_ = v1459
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1571 int32
	_ = v1571
	var v1575 int32
	_ = v1575
	var v1577 int32
	_ = v1577
	var v1583 int32
	_ = v1583
	var v1600 int32
	_ = v1600
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1627 int32
	_ = v1627
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1676 int32
	_ = v1676
	var v1678 int32
	_ = v1678
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1739 int32
	_ = v1739
	var v1756 int32
	_ = v1756
	var v1763 int32
	_ = v1763
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1784 int32
	_ = v1784
	var v1786 int32
	_ = v1786
	var v1792 int32
	_ = v1792
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1808 int32
	_ = v1808
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1834 int32
	_ = v1834
	var v1837 int32
	_ = v1837
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1874 int32
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1897 int32
	_ = v1897
	var v1899 int32
	_ = v1899
	var v1906 int32
	_ = v1906
	var v1909 int32
	_ = v1909
	var v1913 int32
	_ = v1913
	var v1920 int32
	_ = v1920
	var v1921 int32
	_ = v1921
	var v1935 int32
	_ = v1935
	var v1942 int32
	_ = v1942
	var v1946 int32
	_ = v1946
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v1946
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v16 = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 <= v16 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v138 {
		v157 = v7
		goto L51
	} else {
		goto L52
	}
L4:
	;
	goto L3
L5:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v135
	goto L2
L6:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L7:
	;
	v34 = F_find_among(m, l0, int32(_a_F_dutch_UTF_8_stem_0), int32(11))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v73 = v9
	v75 = v17
	goto L6
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v16))))
	if v21&int32(224) != int32(160) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if int32(1)<<(uint(v21)%32)&int32(340306450) != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v38
	switch v34 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	default:
		goto L5
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = v38
	v75 = v72
	goto L6
L15:
	;
	v68 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_1))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L28
	}
L16:
	;
	v62 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v56 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_3))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v50 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_4))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v44 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_5))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v44 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v1946 = v44
	goto L1
L22:
	;
	if int32(0) <= v50 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v1946 = v50
	goto L1
L24:
	;
	if int32(0) <= v56 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v1946 = v56
	goto L1
L26:
	;
	if int32(0) <= v62 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v1946 = v62
	goto L1
L28:
	;
	if int32(0) <= v68 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v1946 = v68
	goto L1
L30:
	;
	if v128 < int32(0) {
		goto L4
	} else {
		goto L50
	}
L32:
	;
	goto L33
L33:
	;
	goto L34
L34:
	;
	v83 = v73
	v85 = int32(1)
	goto L37
L36:
	;
	v128 = v113
	goto L30
L37:
	;
	if v75 <= v83 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v128 = int32(-1)
	goto L30
L40:
	;
	goto L41
L41:
	;
	v90 = v83 + int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v83))))
	if base.Ui32(v92) < base.Ui32(int32(192)) {
		v113 = v90
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(1)
	if v114 < v85 {
		v83 = v113
		v85 = v85 - v114
		goto L37
	} else {
		goto L49
	}
L43:
	;
	if v75 <= v90 {
		v113 = v90
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v99 = v90
	goto L45
L45:
	;
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76+v99))))
	if int32(-65) < v102 {
		v113 = v99
		goto L42
	} else {
		goto L47
	}
L46:
	;
	v113 = v75
	goto L42
L47:
	;
	v106 = v99 + int32(1)
	if v106 != v75 {
		v99 = v106
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L38
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128
	goto L5
L51:
	;
	v163 = v157
	goto L56
L52:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v7))))
	if v142 != int32(121) {
		v157 = v7
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v145 = int32(1)
	v146 = v7 + v145
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v146
	v151 = F_slice_from_s(m, l0, v145, int32(_a_F_dutch_UTF_8_stem_6))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	if v151 < int32(0) {
		v1946 = v151
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v157 = v155
	goto L51
L56:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L61
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v507)+4)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v507)+8)) = v508
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L153
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	goto L129
L59:
	;
	if v281 != 0 {
		goto L83
	} else {
		goto L84
	}
L60:
	;
	v281 = v274
	goto L59
L61:
	;
	if v176 <= v175 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v274 = int32(0)
	goto L60
L63:
	;
	v281 = int32(-1)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v192 = int32(1)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v177))))
	if base.Ui32(v194) < base.Ui32(int32(192)) {
		v251 = v194
		v252 = v192
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if int32(232) < v251 {
		v274 = v252
		goto L60
	} else {
		goto L79
	}
L67:
	;
	v198 = v175 + int32(1)
	if v198 == v176 {
		v251 = v194
		v252 = v192
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v177))))
	v203 = v201 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v194) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v177))))
	v219 = v217 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v194) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v207 = v175 + int32(2)
	if v207 != v176 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v251 = v194<<(uint(int32(6))%32)&int32(1984) | v203
	v252 = int32(2)
	goto L66
L73:
	;
	goto L72
L74:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v223))))
	v251 = v236&int32(63) | (v194<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_7) | v203<<(uint(int32(12))%32) | v219<<(uint(int32(6))%32))
	v252 = int32(4)
	goto L66
L75:
	;
	v223 = v175 + int32(3)
	if v223 != v176 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v251 = v194<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v203<<(uint(int32(6))%32) | v219
	v252 = int32(3)
	goto L66
L78:
	;
	goto L77
L79:
	;
	v256 = v251 - int32(97)
	if v256 < int32(0) {
		v274 = v252
		goto L60
	} else {
		goto L80
	}
L80:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v256)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v262)>>(uint(v256&int32(7))%32))&int32(1) == int32(0) {
		v274 = v252
		goto L60
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 + v175
	goto L82
L82:
	;
	goto L62
L83:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v448 = v282
	v449 = v283
	goto L58
L84:
	;
	goto L85
L85:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v287 == v284 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v448 = v284
	v449 = v286
	goto L58
L87:
	;
	goto L88
L88:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v286))))
	if v290 == int32(105) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	goto L56
L90:
	;
	v439 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_9))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L12
	} else {
		goto L125
	}
L91:
	;
	v294 = v284 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L96
L92:
	;
	v419 = v287
	v420 = v286
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v284
	if v284 == v419 {
		goto L119
	} else {
		goto L120
	}
L94:
	;
	if v414 == int32(0) {
		goto L90
	} else {
		goto L118
	}
L95:
	;
	v414 = v407
	goto L94
L96:
	;
	if v309 <= v294 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v407 = int32(0)
	goto L95
L98:
	;
	v414 = int32(-1)
	goto L94
L99:
	;
	goto L100
L100:
	;
	v325 = int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v310))))
	if base.Ui32(v327) < base.Ui32(int32(192)) {
		v384 = v327
		v385 = v325
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if int32(232) < v384 {
		v407 = v385
		goto L95
	} else {
		goto L114
	}
L102:
	;
	v331 = v284 + int32(2)
	if v331 == v309 {
		v384 = v327
		v385 = v325
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v310))))
	v336 = v334 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v327) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340+v310))))
	v352 = v350 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v327) {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v340 = v284 + int32(3)
	if v340 != v309 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v384 = v327<<(uint(int32(6))%32)&int32(1984) | v336
	v385 = int32(2)
	goto L101
L108:
	;
	goto L107
L109:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v356))))
	v384 = v369&int32(63) | (v327<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_7) | v336<<(uint(int32(12))%32) | v352<<(uint(int32(6))%32))
	v385 = int32(4)
	goto L101
L110:
	;
	v356 = v284 + int32(4)
	if v356 != v309 {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v384 = v327<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v336<<(uint(int32(6))%32) | v352
	v385 = int32(3)
	goto L101
L113:
	;
	goto L112
L114:
	;
	v389 = v384 - int32(97)
	if v389 < int32(0) {
		v407 = v385
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v389)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v395)>>(uint(v389&int32(7))%32))&int32(1) == int32(0) {
		v407 = v385
		goto L95
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 + v294
	goto L117
L117:
	;
	goto L97
L118:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v419 = v418
	v420 = v417
	goto L93
L119:
	;
	v448 = v284
	v449 = v420
	goto L58
L120:
	;
	goto L121
L121:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v420))))
	if v424 != int32(121) {
		v448 = v419
		v449 = v420
		goto L58
	} else {
		goto L122
	}
L122:
	;
	v427 = int32(1)
	v428 = v284 + v427
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v428
	v433 = F_slice_from_s(m, l0, v427, int32(_a_F_dutch_UTF_8_stem_10))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L12
	} else {
		goto L123
	}
L123:
	;
	if int32(0) <= v433 {
		goto L89
	} else {
		goto L124
	}
L124:
	;
	v1946 = v433
	goto L1
L125:
	;
	if v439 < int32(0) {
		v1946 = v439
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L89
L127:
	;
	if int32(0) <= v502 {
		goto L147
	} else {
		goto L148
	}
L129:
	;
	goto L130
L130:
	;
	goto L131
L131:
	;
	v457 = v163
	v459 = int32(1)
	goto L134
L133:
	;
	v502 = v487
	goto L127
L134:
	;
	if v448 <= v457 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L133
L136:
	;
	v502 = int32(-1)
	goto L127
L137:
	;
	goto L138
L138:
	;
	v464 = v457 + int32(1)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v449+v457))))
	if base.Ui32(v466) < base.Ui32(int32(192)) {
		v487 = v464
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v488 = int32(1)
	if v488 < v459 {
		v457 = v487
		v459 = v459 - v488
		goto L134
	} else {
		goto L146
	}
L140:
	;
	if v448 <= v464 {
		v487 = v464
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v473 = v464
	goto L142
L142:
	;
	v476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v449+v473))))
	if int32(-65) < v476 {
		v487 = v473
		goto L139
	} else {
		goto L144
	}
L143:
	;
	v487 = v448
	goto L139
L144:
	;
	v480 = v473 + int32(1)
	if v480 != v448 {
		v473 = v480
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	goto L135
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v502
	v163 = v502
	goto L56
L148:
	;
	goto L149
L149:
	;
	goto L57
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1071
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1071
	if v1071 <= v7 {
		v1257 = v1069
		goto L277
	} else {
		goto L278
	}
L151:
	;
	if v565 < int32(0) {
		v1069 = v565
		goto L150
	} else {
		goto L171
	}
L153:
	;
	goto L154
L154:
	;
	goto L155
L155:
	;
	v520 = v512
	v522 = int32(3)
	goto L158
L157:
	;
	v565 = v550
	goto L151
L158:
	;
	if v513 <= v520 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	v565 = int32(-1)
	goto L151
L161:
	;
	goto L162
L162:
	;
	v527 = v520 + int32(1)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v520))))
	if base.Ui32(v529) < base.Ui32(int32(192)) {
		v550 = v527
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v551 = int32(1)
	if v551 < v522 {
		v520 = v550
		v522 = v522 - v551
		goto L158
	} else {
		goto L170
	}
L164:
	;
	if v513 <= v527 {
		v550 = v527
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v536 = v527
	goto L166
L166:
	;
	v539 = int32(*(*int8)(unsafe.Add(mBase, uint32(v511+v536))))
	if int32(-65) < v539 {
		v550 = v536
		goto L163
	} else {
		goto L168
	}
L167:
	;
	v550 = v513
	goto L163
L168:
	;
	v543 = v536 + int32(1)
	if v543 != v513 {
		v536 = v543
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	goto L159
L171:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v565
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v512
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v592 = v512
	goto L174
L172:
	;
	if v687 < int32(0) {
		v1069 = v565
		goto L150
	} else {
		goto L197
	}
L173:
	;
	v687 = v659
	goto L172
L174:
	;
	if v583 <= v592 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v687 = int32(-1)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v599 = int32(1)
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592+v584))))
	if base.Ui32(v601) < base.Ui32(int32(192)) {
		v658 = v601
		v659 = v599
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if int32(232) < v658 {
		goto L192
	} else {
		goto L193
	}
L180:
	;
	v605 = v592 + int32(1)
	if v605 == v583 {
		v658 = v601
		v659 = v599
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605+v584))))
	v610 = v608 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v601) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v584))))
	v626 = v624 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v601) {
		goto L188
	} else {
		goto L189
	}
L183:
	;
	v614 = v592 + int32(2)
	if v614 != v583 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v658 = v601<<(uint(int32(6))%32)&int32(1984) | v610
	v659 = int32(2)
	goto L179
L186:
	;
	goto L185
L187:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584+v630))))
	v658 = v643&int32(63) | (v601<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_7) | v610<<(uint(int32(12))%32) | v626<<(uint(int32(6))%32))
	v659 = int32(4)
	goto L179
L188:
	;
	v630 = v592 + int32(3)
	if v630 != v583 {
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v658 = v601<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v610<<(uint(int32(6))%32) | v626
	v659 = int32(3)
	goto L179
L191:
	;
	goto L190
L192:
	;
	v676 = v659 + v592
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v676
	v592 = v676
	goto L174
L193:
	;
	v663 = v658 - int32(97)
	if v663 < int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v663)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v669)>>(uint(v663&int32(7))%32))&int32(1) != 0 {
		goto L173
	} else {
		goto L195
	}
L195:
	;
	goto L192
L197:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v691 = v690 + v687
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v691
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = v691
	goto L200
L198:
	;
	if v810 < int32(0) {
		v1069 = v565
		goto L150
	} else {
		goto L222
	}
L199:
	;
	v810 = v781
	goto L198
L200:
	;
	if v705 <= v714 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v810 = int32(-1)
	goto L198
L203:
	;
	goto L204
L204:
	;
	v721 = int32(1)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+v706))))
	if base.Ui32(v723) < base.Ui32(int32(192)) {
		v780 = v723
		v781 = v721
		goto L205
	} else {
		goto L206
	}
L205:
	;
	if int32(232) < v780 {
		goto L199
	} else {
		goto L218
	}
L206:
	;
	v727 = v714 + int32(1)
	if v727 == v705 {
		v780 = v723
		v781 = v721
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727+v706))))
	v732 = v730 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v723) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736+v706))))
	v748 = v746 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v723) {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	v736 = v714 + int32(2)
	if v736 != v705 {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v780 = v723<<(uint(int32(6))%32)&int32(1984) | v732
	v781 = int32(2)
	goto L205
L212:
	;
	goto L211
L213:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706+v752))))
	v780 = v765&int32(63) | (v723<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_7) | v732<<(uint(int32(12))%32) | v748<<(uint(int32(6))%32))
	v781 = int32(4)
	goto L205
L214:
	;
	v752 = v714 + int32(3)
	if v752 != v705 {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v780 = v723<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v732<<(uint(int32(6))%32) | v748
	v781 = int32(3)
	goto L205
L217:
	;
	goto L216
L218:
	;
	v785 = v780 - int32(97)
	if v785 < int32(0) {
		goto L199
	} else {
		goto L219
	}
L219:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v785)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v791)>>(uint(v785&int32(7))%32))&int32(1) == int32(0) {
		goto L199
	} else {
		goto L220
	}
L220:
	;
	v799 = v781 + v714
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v799
	v714 = v799
	goto L200
L222:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v814 = v813 + v810
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v814
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	if v817 < v814 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v819 = v814
	goto L225
L224:
	;
	v819 = v817
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+8)) = v819
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v842 = v832
	goto L228
L226:
	;
	if v937 < int32(0) {
		v1069 = v817
		goto L150
	} else {
		goto L251
	}
L227:
	;
	v937 = v909
	goto L226
L228:
	;
	if v833 <= v842 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v937 = int32(-1)
	goto L226
L231:
	;
	goto L232
L232:
	;
	v849 = int32(1)
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842+v834))))
	if base.Ui32(v851) < base.Ui32(int32(192)) {
		v908 = v851
		v909 = v849
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if int32(232) < v908 {
		goto L246
	} else {
		goto L247
	}
L234:
	;
	v855 = v842 + int32(1)
	if v855 == v833 {
		v908 = v851
		v909 = v849
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855+v834))))
	v860 = v858 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v851) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+v834))))
	v876 = v874 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v851) {
		goto L242
	} else {
		goto L243
	}
L237:
	;
	v864 = v842 + int32(2)
	if v864 != v833 {
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v908 = v851<<(uint(int32(6))%32)&int32(1984) | v860
	v909 = int32(2)
	goto L233
L240:
	;
	goto L239
L241:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834+v880))))
	v908 = v893&int32(63) | (v851<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_7) | v860<<(uint(int32(12))%32) | v876<<(uint(int32(6))%32))
	v909 = int32(4)
	goto L233
L242:
	;
	v880 = v842 + int32(3)
	if v880 != v833 {
		goto L241
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v908 = v851<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v860<<(uint(int32(6))%32) | v876
	v909 = int32(3)
	goto L233
L245:
	;
	goto L244
L246:
	;
	v926 = v909 + v842
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v926
	v842 = v926
	goto L228
L247:
	;
	v913 = v908 - int32(97)
	if v913 < int32(0) {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v913)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v919)>>(uint(v913&int32(7))%32))&int32(1) != 0 {
		goto L227
	} else {
		goto L249
	}
L249:
	;
	goto L246
L251:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v941 = v940 + v937
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v941
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v964 = v941
	goto L254
L252:
	;
	if v1060 < int32(0) {
		v1069 = v817
		goto L150
	} else {
		goto L276
	}
L253:
	;
	v1060 = v1031
	goto L252
L254:
	;
	if v955 <= v964 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1060 = int32(-1)
	goto L252
L257:
	;
	goto L258
L258:
	;
	v971 = int32(1)
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964+v956))))
	if base.Ui32(v973) < base.Ui32(int32(192)) {
		v1030 = v973
		v1031 = v971
		goto L259
	} else {
		goto L260
	}
L259:
	;
	if int32(232) < v1030 {
		goto L253
	} else {
		goto L272
	}
L260:
	;
	v977 = v964 + int32(1)
	if v977 == v955 {
		v1030 = v973
		v1031 = v971
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977+v956))))
	v982 = v980 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v973) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986+v956))))
	v998 = v996 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v973) {
		goto L268
	} else {
		goto L269
	}
L263:
	;
	v986 = v964 + int32(2)
	if v986 != v955 {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1030 = v973<<(uint(int32(6))%32)&int32(1984) | v982
	v1031 = int32(2)
	goto L259
L266:
	;
	goto L265
L267:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956+v1002))))
	v1030 = v1015&int32(63) | (v973<<(uint(int32(18))%32)&int32(_a_F_dutch_UTF_8_stem_7) | v982<<(uint(int32(12))%32) | v998<<(uint(int32(6))%32))
	v1031 = int32(4)
	goto L259
L268:
	;
	v1002 = v964 + int32(3)
	if v1002 != v955 {
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1030 = v973<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v982<<(uint(int32(6))%32) | v998
	v1031 = int32(3)
	goto L259
L271:
	;
	goto L270
L272:
	;
	v1035 = v1030 - int32(97)
	if v1035 < int32(0) {
		goto L253
	} else {
		goto L273
	}
L273:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1035)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v1041)>>(uint(v1035&int32(7))%32))&int32(1) == int32(0) {
		goto L253
	} else {
		goto L274
	}
L274:
	;
	v1049 = v1031 + v964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1049
	v964 = v1049
	goto L254
L276:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+4)) = v1064 + v1060
	v1069 = v817
	goto L150
L277:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1258
	v1260 = F_r_e_ending_2(m, l0)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L12
	} else {
		goto L316
	}
L278:
	;
	v1075 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1077 = int32(1)
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1075+v1071-v1077))))
	if base.B2i32(v1079&int32(224) != int32(96))|base.B2i32(v1077<<(uint(v1079)%32)&int32(_a_F_dutch_UTF_8_stem_11) == int32(0)) != 0 {
		v1257 = v1069
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v1093 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_12), int32(5))
	mBase = m.M
	v1094 = m.ExcPending
	if v1094 != 0 {
		goto L12
	} else {
		goto L280
	}
L280:
	;
	if v1093 == int32(0) {
		v1257 = v1069
		goto L277
	} else {
		goto L281
	}
L281:
	;
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1097
	switch v1093 - int32(1) {
	case 0:
		goto L284
	case 1:
		goto L283
	case 2:
		goto L282
	default:
		v1257 = v1069
		goto L277
	}
L282:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+8))
	if v1097 < v1117 {
		goto L292
	} else {
		goto L293
	}
L283:
	;
	v1112 = F_r_en_ending_2(m, l0)
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L12
	} else {
		goto L290
	}
L284:
	;
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1101)+8))
	if v1097 < v1102 {
		goto L285
	} else {
		goto L286
	}
L285:
	;
	v1257 = int32(0)
	goto L277
L286:
	;
	goto L287
L287:
	;
	v1107 = F_slice_from_s(m, l0, int32(4), int32(_a_F_dutch_UTF_8_stem_13))
	mBase = m.M
	v1108 = m.ExcPending
	if v1108 != 0 {
		goto L12
	} else {
		goto L288
	}
L288:
	;
	if v1107 < int32(0) {
		v1946 = v1107
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1257 = int32(1)
	goto L277
L290:
	;
	if int32(0) <= v1112 {
		v1257 = v1112
		goto L277
	} else {
		goto L291
	}
L291:
	;
	v1946 = v1112
	goto L1
L292:
	;
	v1257 = int32(0)
	goto L277
L293:
	;
	goto L294
L294:
	;
	v1120 = int32(1)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L297
L295:
	;
	if v1250 != 0 {
		v1257 = v1120
		goto L277
	} else {
		goto L313
	}
L296:
	;
	v1250 = v1243
	goto L295
L297:
	;
	if v1133 <= v1134 {
		v1243 = int32(-1)
		goto L296
	} else {
		goto L299
	}
L298:
	;
	v1243 = int32(0)
	goto L296
L299:
	;
	v1151 = int32(1)
	v1152 = v1133 - v1151
	v1154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1135+v1152))))
	v1156 = v1154 & int32(255)
	if base.B2i32(v1152 == v1134)|base.B2i32(int32(0) <= v1154) != 0 {
		v1214 = v1156
		v1218 = v1151
		goto L300
	} else {
		goto L301
	}
L300:
	;
	if int32(232) < v1214 {
		goto L308
	} else {
		goto L309
	}
L301:
	;
	v1163 = v1156 & int32(63)
	v1165 = v1133 - int32(2)
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135+v1165))))
	v1169 = v1167 << (uint(int32(6)) % 32)
	if base.B2i32(v1165 != v1134)&base.B2i32(base.Ui32(v1167) < base.Ui32(int32(192))) == int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	v1214 = v1169&int32(1984) | v1163
	v1218 = int32(2)
	goto L300
L303:
	;
	goto L304
L304:
	;
	v1182 = v1169&int32(4032) | v1163
	v1184 = v1133 - int32(3)
	v1186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1135+v1184))))
	if base.B2i32(v1184 != v1134)&base.B2i32(base.Ui32(v1186) < base.Ui32(int32(224))) == int32(0) {
		goto L305
	} else {
		goto L306
	}
L305:
	;
	v1214 = v1186<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v1182
	v1218 = int32(3)
	goto L300
L306:
	;
	goto L307
L307:
	;
	v1204 = int32(4)
	v1206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1133+v1135-v1204))))
	v1214 = v1186<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_14) | v1206&int32(7)<<(uint(int32(18))%32) | v1182
	v1218 = v1204
	goto L300
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1133 - v1218
	goto L312
L309:
	;
	v1220 = v1214 - int32(97)
	if v1220 < int32(0) {
		goto L308
	} else {
		goto L310
	}
L310:
	;
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1220)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[1]))))
	if int32(base.Ui32(v1226)>>(uint(v1220&int32(7))%32))&int32(1) == int32(0) {
		goto L308
	} else {
		goto L311
	}
L311:
	;
	v1250 = v1218
	goto L295
L312:
	;
	goto L298
L313:
	;
	v1251 = F_slice_del(m, l0)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L12
	} else {
		goto L314
	}
L314:
	;
	if v1251 < int32(0) {
		v1946 = v1251
		goto L1
	} else {
		goto L315
	}
L315:
	;
	v1257 = v1120
	goto L277
L316:
	;
	if v1260 < int32(0) {
		v1946 = v1260
		goto L1
	} else {
		goto L317
	}
L317:
	;
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1264
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1264
	v1267 = int32(4)
	v1269 = int32(0)
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1264-v1272 < v1267 {
		v1282 = v1269
		goto L325
	} else {
		goto L326
	}
L318:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1834
	v1837 = v1834
	goto L464
L319:
	;
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1476
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1492 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L402
L320:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1463)+4))
	if v1361 < v1464 {
		goto L319
	} else {
		goto L396
	}
L321:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(v1456)+4))
	if v1361 < v1457 {
		goto L319
	} else {
		goto L393
	}
L322:
	;
	if int32(0) <= v1451 {
		goto L318
	} else {
		goto L392
	}
L323:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1334
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1334
	v1338 = v1334 - int32(1)
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1338 <= v1339 {
		goto L319
	} else {
		goto L349
	}
L324:
	;
	if v1282 == int32(0) {
		v1333 = v1257
		goto L323
	} else {
		goto L328
	}
L325:
	;
	goto L324
L326:
	;
	v1275 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1278 = F_memcmp(m, v1275+v1264-v1267, int32(_a_F_dutch_UTF_8_stem_15), v1267)
	mBase = m.M
	if v1278 != 0 {
		v1282 = v1269
		goto L325
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1264 - v1267
	v1282 = int32(1)
	goto L325
L328:
	;
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1285
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v1287)+4))
	if v1285 < v1288 {
		v1333 = v1257
		goto L323
	} else {
		goto L329
	}
L329:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1290 < v1285 {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1292+v1285-int32(1)))))
	if v1296 == int32(99) {
		v1333 = v1257
		goto L323
	} else {
		goto L333
	}
L331:
	;
	goto L332
L332:
	;
	v1299 = F_slice_del(m, l0)
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L12
	} else {
		goto L334
	}
L333:
	;
	goto L332
L334:
	;
	if v1299 < int32(0) {
		v1946 = v1299
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1303
	v1305 = int32(2)
	v1307 = int32(0)
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1303-v1310 < v1305 {
		v1320 = v1307
		goto L337
	} else {
		goto L338
	}
L336:
	;
	if v1320 == int32(0) {
		v1333 = v1257
		goto L323
	} else {
		goto L340
	}
L337:
	;
	goto L336
L338:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1316 = F_memcmp(m, v1313+v1303-v1305, int32(_a_F_dutch_UTF_8_stem_16), v1305)
	mBase = m.M
	if v1316 != 0 {
		v1320 = v1307
		goto L337
	} else {
		goto L339
	}
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1303 - v1305
	v1320 = int32(1)
	goto L337
L340:
	;
	v1323 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1323
	v1325 = F_r_en_ending_2(m, l0)
	mBase = m.M
	v1326 = m.ExcPending
	if v1326 != 0 {
		goto L12
	} else {
		goto L341
	}
L341:
	;
	v1328 = base.B2i32(v1325 < int32(0))
	if v1325 < int32(0) {
		goto L342
	} else {
		goto L343
	}
L342:
	;
	v1329 = v1325
	goto L344
L343:
	;
	v1329 = v1257
	goto L344
L344:
	;
	if v1325 != 0 {
		goto L345
	} else {
		goto L346
	}
L345:
	;
	v1330 = v1329
	goto L347
L346:
	;
	v1330 = v1257
	goto L347
L347:
	;
	if v1325 < int32(0) {
		v1451 = v1330
		goto L322
	} else {
		goto L348
	}
L348:
	;
	v1333 = v1330
	goto L323
L349:
	;
	v1341 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1341+v1338))))
	if base.B2i32(v1343&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1343)%32)&int32(_a_F_dutch_UTF_8_stem_17) == int32(0)) != 0 {
		goto L319
	} else {
		goto L350
	}
L350:
	;
	v1357 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_18), int32(6))
	mBase = m.M
	v1358 = m.ExcPending
	if v1358 != 0 {
		goto L12
	} else {
		goto L351
	}
L351:
	;
	if v1357 == int32(0) {
		goto L319
	} else {
		goto L352
	}
L352:
	;
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1361
	switch v1357 - int32(1) {
	case 0:
		goto L355
	case 1:
		goto L354
	case 2:
		goto L353
	case 3:
		goto L321
	case 4:
		goto L320
	default:
		goto L319
	}
L353:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1437 = *(*int32)(unsafe.Add(mBase, uint32(v1436)+4))
	if v1361 < v1437 {
		goto L319
	} else {
		goto L381
	}
L354:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+4))
	if v1361 < v1421 {
		goto L319
	} else {
		goto L374
	}
L355:
	;
	v1365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(v1365)+4))
	if v1361 < v1366 {
		goto L319
	} else {
		goto L356
	}
L356:
	;
	v1368 = F_slice_del(m, l0)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L12
	} else {
		goto L357
	}
L357:
	;
	if v1368 < int32(0) {
		v1946 = v1368
		goto L1
	} else {
		goto L358
	}
L358:
	;
	v1372 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1372
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1375 = int32(2)
	v1377 = int32(0)
	v1380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1372-v1380 < v1375 {
		v1390 = v1377
		goto L361
	} else {
		goto L362
	}
L359:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1412 + (v1372 - v1374)
	v1416 = F_r_undouble_3(m, l0)
	mBase = m.M
	v1417 = m.ExcPending
	if v1417 != 0 {
		goto L12
	} else {
		goto L372
	}
L360:
	;
	if v1390 == int32(0) {
		goto L359
	} else {
		goto L364
	}
L361:
	;
	goto L360
L362:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1386 = F_memcmp(m, v1383+v1372-v1375, int32(_a_F_dutch_UTF_8_stem_19), v1375)
	mBase = m.M
	if v1386 != 0 {
		v1390 = v1377
		goto L361
	} else {
		goto L363
	}
L363:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1372 - v1375
	v1390 = int32(1)
	goto L361
L364:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1393
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v1395)+4))
	if v1393 < v1396 {
		goto L359
	} else {
		goto L365
	}
L365:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1398 < v1393 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1400 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1400+v1393-int32(1)))))
	if v1404 == int32(101) {
		goto L359
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1407 = F_slice_del(m, l0)
	mBase = m.M
	v1408 = m.ExcPending
	if v1408 != 0 {
		goto L12
	} else {
		goto L370
	}
L369:
	;
	goto L368
L370:
	;
	if int32(0) <= v1407 {
		goto L319
	} else {
		goto L371
	}
L371:
	;
	v1946 = v1407
	goto L1
L372:
	;
	if int32(0) <= v1416 {
		goto L319
	} else {
		goto L373
	}
L373:
	;
	v1946 = v1416
	goto L1
L374:
	;
	v1423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1423 < v1361 {
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1425+v1361-int32(1)))))
	if v1429 == int32(101) {
		goto L319
	} else {
		goto L378
	}
L376:
	;
	goto L377
L377:
	;
	v1432 = F_slice_del(m, l0)
	mBase = m.M
	v1433 = m.ExcPending
	if v1433 != 0 {
		goto L12
	} else {
		goto L379
	}
L378:
	;
	goto L377
L379:
	;
	if int32(0) <= v1432 {
		goto L319
	} else {
		goto L380
	}
L380:
	;
	v1946 = v1432
	goto L1
L381:
	;
	v1439 = F_slice_del(m, l0)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L12
	} else {
		goto L382
	}
L382:
	;
	if v1439 < int32(0) {
		v1946 = v1439
		goto L1
	} else {
		goto L383
	}
L383:
	;
	v1443 = F_r_e_ending_2(m, l0)
	mBase = m.M
	v1444 = m.ExcPending
	if v1444 != 0 {
		goto L12
	} else {
		goto L384
	}
L384:
	;
	if int32(0) <= v1443 {
		goto L319
	} else {
		goto L385
	}
L385:
	;
	if v1443 < int32(0) {
		goto L386
	} else {
		goto L387
	}
L386:
	;
	v1449 = v1443
	goto L388
L387:
	;
	v1449 = v1333
	goto L388
L388:
	;
	if v1443 != 0 {
		goto L389
	} else {
		goto L390
	}
L389:
	;
	v1450 = v1449
	goto L391
L390:
	;
	v1450 = v1333
	goto L391
L391:
	;
	v1451 = v1450
	goto L322
L392:
	;
	v1946 = v1451
	goto L1
L393:
	;
	v1459 = F_slice_del(m, l0)
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L12
	} else {
		goto L394
	}
L394:
	;
	if int32(0) <= v1459 {
		goto L319
	} else {
		goto L395
	}
L395:
	;
	v1946 = v1459
	goto L1
L396:
	;
	v1466 = *(*int32)(unsafe.Add(mBase, uint32(v1463)+12))
	if v1466 == int32(0) {
		goto L319
	} else {
		goto L397
	}
L397:
	;
	v1469 = F_slice_del(m, l0)
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L12
	} else {
		goto L398
	}
L398:
	;
	if v1469 < int32(0) {
		v1946 = v1469
		goto L1
	} else {
		goto L399
	}
L399:
	;
	goto L319
L400:
	;
	if v1607 != 0 {
		goto L318
	} else {
		goto L418
	}
L401:
	;
	v1607 = v1600
	goto L400
L402:
	;
	if v1476 <= v1491 {
		v1600 = int32(-1)
		goto L401
	} else {
		goto L404
	}
L403:
	;
	v1600 = int32(0)
	goto L401
L404:
	;
	v1508 = int32(1)
	v1509 = v1476 - v1508
	v1511 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1492+v1509))))
	v1513 = v1511 & int32(255)
	if base.B2i32(v1509 == v1491)|base.B2i32(int32(0) <= v1511) != 0 {
		v1571 = v1513
		v1575 = v1508
		goto L405
	} else {
		goto L406
	}
L405:
	;
	if int32(232) < v1571 {
		goto L413
	} else {
		goto L414
	}
L406:
	;
	v1520 = v1513 & int32(63)
	v1522 = v1476 - int32(2)
	v1524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492+v1522))))
	v1526 = v1524 << (uint(int32(6)) % 32)
	if base.B2i32(v1522 != v1491)&base.B2i32(base.Ui32(v1524) < base.Ui32(int32(192))) == int32(0) {
		goto L407
	} else {
		goto L408
	}
L407:
	;
	v1571 = v1526&int32(1984) | v1520
	v1575 = int32(2)
	goto L405
L408:
	;
	goto L409
L409:
	;
	v1539 = v1526&int32(4032) | v1520
	v1541 = v1476 - int32(3)
	v1543 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1492+v1541))))
	if base.B2i32(v1541 != v1491)&base.B2i32(base.Ui32(v1543) < base.Ui32(int32(224))) == int32(0) {
		goto L410
	} else {
		goto L411
	}
L410:
	;
	v1571 = v1543<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v1539
	v1575 = int32(3)
	goto L405
L411:
	;
	goto L412
L412:
	;
	v1561 = int32(4)
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1476+v1492-v1561))))
	v1571 = v1543<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_14) | v1563&int32(7)<<(uint(int32(18))%32) | v1539
	v1575 = v1561
	goto L405
L413:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1476 - v1575
	goto L417
L414:
	;
	v1577 = v1571 - int32(73)
	if v1577 < int32(0) {
		goto L413
	} else {
		goto L415
	}
L415:
	;
	v1583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1577)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[2]))))
	if int32(base.Ui32(v1583)>>(uint(v1577&int32(7))%32))&int32(1) == int32(0) {
		goto L413
	} else {
		goto L416
	}
L416:
	;
	v1607 = v1575
	goto L400
L417:
	;
	goto L403
L418:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1610 = v1608 - int32(1)
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1610 <= v1611 {
		goto L318
	} else {
		goto L419
	}
L419:
	;
	v1613 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1613+v1610))))
	if base.B2i32(v1615&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1615)%32)&int32(_a_F_dutch_UTF_8_stem_20) == int32(0)) != 0 {
		goto L318
	} else {
		goto L420
	}
L420:
	;
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1630 = F_find_among_b(m, l0, int32(_a_F_dutch_UTF_8_stem_21), int32(4))
	mBase = m.M
	v1631 = m.ExcPending
	if v1631 != 0 {
		goto L12
	} else {
		goto L421
	}
L421:
	;
	if v1630 == int32(0) {
		goto L318
	} else {
		goto L422
	}
L422:
	;
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L425
L423:
	;
	if v1763 != 0 {
		goto L318
	} else {
		goto L441
	}
L424:
	;
	v1763 = v1756
	goto L423
L425:
	;
	if v1646 <= v1647 {
		v1756 = int32(-1)
		goto L424
	} else {
		goto L427
	}
L426:
	;
	v1756 = int32(0)
	goto L424
L427:
	;
	v1664 = int32(1)
	v1665 = v1646 - v1664
	v1667 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1648+v1665))))
	v1669 = v1667 & int32(255)
	if base.B2i32(v1665 == v1647)|base.B2i32(int32(0) <= v1667) != 0 {
		v1727 = v1669
		v1731 = v1664
		goto L428
	} else {
		goto L429
	}
L428:
	;
	if int32(232) < v1727 {
		goto L436
	} else {
		goto L437
	}
L429:
	;
	v1676 = v1669 & int32(63)
	v1678 = v1646 - int32(2)
	v1680 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648+v1678))))
	v1682 = v1680 << (uint(int32(6)) % 32)
	if base.B2i32(v1678 != v1647)&base.B2i32(base.Ui32(v1680) < base.Ui32(int32(192))) == int32(0) {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1727 = v1682&int32(1984) | v1676
	v1731 = int32(2)
	goto L428
L431:
	;
	goto L432
L432:
	;
	v1695 = v1682&int32(4032) | v1676
	v1697 = v1646 - int32(3)
	v1699 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1648+v1697))))
	if base.B2i32(v1697 != v1647)&base.B2i32(base.Ui32(v1699) < base.Ui32(int32(224))) == int32(0) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v1727 = v1699<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_8) | v1695
	v1731 = int32(3)
	goto L428
L434:
	;
	goto L435
L435:
	;
	v1717 = int32(4)
	v1719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1646+v1648-v1717))))
	v1727 = v1699<<(uint(int32(12))%32)&int32(_a_F_dutch_UTF_8_stem_14) | v1719&int32(7)<<(uint(int32(18))%32) | v1695
	v1731 = v1717
	goto L428
L436:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1646 - v1731
	goto L440
L437:
	;
	v1733 = v1727 - int32(97)
	if v1733 < int32(0) {
		goto L436
	} else {
		goto L438
	}
L438:
	;
	v1739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1733)>>(uint(int32(3))%32)))+uint32(_c_F_dutch_UTF_8_stem[0]))))
	if int32(base.Ui32(v1739)>>(uint(v1733&int32(7))%32))&int32(1) == int32(0) {
		goto L436
	} else {
		goto L439
	}
L439:
	;
	v1763 = v1731
	goto L423
L440:
	;
	goto L426
L441:
	;
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1766 = v1764 + (v1608 - v1627)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1766
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1766
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L444
L442:
	;
	if v1822 < int32(0) {
		goto L318
	} else {
		goto L461
	}
L444:
	;
	goto L445
L445:
	;
	goto L446
L446:
	;
	v1777 = v1766
	v1779 = int32(1)
	goto L449
L448:
	;
	v1822 = v1804
	goto L442
L449:
	;
	if v1777 <= v1770 {
		goto L451
	} else {
		goto L452
	}
L450:
	;
	goto L448
L451:
	;
	v1822 = int32(-1)
	goto L442
L452:
	;
	goto L453
L453:
	;
	v1784 = v1777 - int32(1)
	v1786 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1769+v1784))))
	if base.B2i32(int32(0) <= v1786)|base.B2i32(v1784 <= v1770) != 0 {
		v1804 = v1784
		goto L454
	} else {
		goto L455
	}
L454:
	;
	v1808 = int32(1)
	if v1808 < v1779 {
		v1777 = v1804
		v1779 = v1779 - v1808
		goto L449
	} else {
		goto L460
	}
L455:
	;
	v1792 = v1784
	goto L456
L456:
	;
	v1797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1769+v1792))))
	if base.Ui32(int32(191)) < base.Ui32(v1797) {
		v1804 = v1792
		goto L454
	} else {
		goto L458
	}
L457:
	;
	v1804 = v1770
	goto L454
L458:
	;
	v1801 = v1792 - int32(1)
	if v1770 < v1801 {
		v1792 = v1801
		goto L456
	} else {
		goto L459
	}
L459:
	;
	goto L457
L460:
	;
	goto L450
L461:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1822
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1822
	v1827 = F_slice_del(m, l0)
	mBase = m.M
	v1828 = m.ExcPending
	if v1828 != 0 {
		goto L12
	} else {
		goto L462
	}
L462:
	;
	if v1827 < int32(0) {
		v1946 = v1827
		goto L1
	} else {
		goto L463
	}
L463:
	;
	goto L318
L464:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1837
	v1843 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1843 <= v1837 {
		goto L469
	} else {
		goto L470
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1834
	v1946 = int32(1)
	goto L1
L466:
	;
	goto L465
L467:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1837 = v1942
	goto L464
L468:
	;
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L482
L469:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1837
	v1880 = v1837
	v1881 = v1843
	goto L468
L470:
	;
	v1845 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1847 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1845+v1837))))
	v1849 = v1847 - int32(73)
	v1850 = int32(0)
	if base.B2i32(v1849 == v1850)|base.B2i32(v1849 == int32(16)) == v1850 {
		goto L469
	} else {
		goto L471
	}
L471:
	;
	v1859 = F_find_among(m, l0, int32(_a_F_dutch_UTF_8_stem_22), int32(3))
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L12
	} else {
		goto L472
	}
L472:
	;
	v1861 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1861
	switch v1859 - int32(1) {
	case 0:
		goto L474
	case 1:
		goto L473
	case 2:
		goto L475
	default:
		goto L467
	}
L473:
	;
	v1874 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_23))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L12
	} else {
		goto L478
	}
L474:
	;
	v1868 = F_slice_from_s(m, l0, int32(1), int32(_a_F_dutch_UTF_8_stem_24))
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L12
	} else {
		goto L476
	}
L475:
	;
	v1865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1880 = v1861
	v1881 = v1865
	goto L468
L476:
	;
	if int32(0) <= v1868 {
		goto L467
	} else {
		goto L477
	}
L477:
	;
	v1946 = v1868
	goto L1
L478:
	;
	if int32(0) <= v1874 {
		goto L467
	} else {
		goto L479
	}
L479:
	;
	v1946 = v1874
	goto L1
L480:
	;
	if v1935 < int32(0) {
		goto L466
	} else {
		goto L500
	}
L482:
	;
	goto L483
L483:
	;
	goto L484
L484:
	;
	v1890 = v1880
	v1892 = int32(1)
	goto L487
L486:
	;
	v1935 = v1920
	goto L480
L487:
	;
	if v1881 <= v1890 {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	goto L486
L489:
	;
	v1935 = int32(-1)
	goto L480
L490:
	;
	goto L491
L491:
	;
	v1897 = v1890 + int32(1)
	v1899 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1883+v1890))))
	if base.Ui32(v1899) < base.Ui32(int32(192)) {
		v1920 = v1897
		goto L492
	} else {
		goto L493
	}
L492:
	;
	v1921 = int32(1)
	if v1921 < v1892 {
		v1890 = v1920
		v1892 = v1892 - v1921
		goto L487
	} else {
		goto L499
	}
L493:
	;
	if v1881 <= v1897 {
		v1920 = v1897
		goto L492
	} else {
		goto L494
	}
L494:
	;
	v1906 = v1897
	goto L495
L495:
	;
	v1909 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1883+v1906))))
	if int32(-65) < v1909 {
		v1920 = v1906
		goto L492
	} else {
		goto L497
	}
L496:
	;
	v1920 = v1881
	goto L492
L497:
	;
	v1913 = v1906 + int32(1)
	if v1913 != v1881 {
		v1906 = v1913
		goto L495
	} else {
		goto L498
	}
L498:
	;
	goto L496
L499:
	;
	goto L488
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1935
	goto L467
}
