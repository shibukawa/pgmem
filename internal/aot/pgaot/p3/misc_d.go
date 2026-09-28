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
	F_errfinish(m, int32(_a_F_DeconstructQualifiedName_1), int32(3402), int32(_a_F_DeconstructQualifiedName_2))
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
	F_errfinish(m, int32(_a_F_DeconstructQualifiedName_1), int32(3408), int32(_a_F_DeconstructQualifiedName_2))
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
	var v39 int64
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v185 int64
	_ = v185
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v382 int32
	_ = v382
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(96)
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
	v39 = int64(0)
	v41 = F_SearchSysCacheList(m, int32(9), int32(1), base.I64_extend_i32_u(l2), v39, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v101 = v41 - int32(-64)
	v111 = v9
	goto L13
L6:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if v43 == int32(0) {
		v91 = v9
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = F_palloc_mul(m, int32(4), v43)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
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
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if int32(0) < v282 {
		goto L48
	} else {
		goto L49
	}
L19:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
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
	F_ReleaseCatCacheList(m, v41)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
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
	v251 = v146 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v41)+56))
	if v251 < v252 {
		v146 = v251
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
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v170)+72))
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
	v185 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v185
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v185
	v199 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+27)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v22)+19)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v199
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
	F_simple_heap_delete(m, v29, v170+int32(60))
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
	v242 = F_heap_modify_tuple(m, v170+int32(56), v31, v22+int32(32), v22+int32(24), v22+int32(16))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L43
	}
L36:
	;
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v230)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = int64(0)
	goto L35
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)) = uint8(v213)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+80)) = int64(0)
	goto L35
L39:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)) = uint8(v209)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+72)) = int64(0)
	goto L35
L40:
	;
	F_errmsg_internal(m, int32(_a_F_DelRoleMems_0), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_DelRoleMems_1), int32(2102), int32(_a_F_DelRoleMems_2))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
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
	F_CatalogTupleUpdate(m, v29, v242+int32(4), v242)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
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
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v22 + int32(96)
	return
L48:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v138+v111<<(uint(int32(2))%32))))
	v292 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v365 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L1
	} else {
		goto L64
	}
L51:
	;
	v310 = v292 << (uint(int32(2)) % 32)
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v101+v310)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+72))
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+22)))
	v315 = v313 + v314
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v315)+8))
	if v316 != v288 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	goto L50
L53:
	;
	v342 = v292 + int32(1)
	if v342 != v282 {
		v292 = v342
		goto L51
	} else {
		goto L63
	}
L54:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v315)+12))
	if v318 != v25 {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v320&int32(2) != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v310))) = int32(2)
	v111 = v111 + int32(1)
	goto L13
L57:
	;
	goto L58
L58:
	;
	if v320&int32(4) != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91+v310))) = int32(3)
	v111 = v111 + int32(1)
	goto L13
L60:
	;
	goto L61
L61:
	;
	F_plan_recursive_revoke(m, v41, v91, v292, v320&int32(1), l7)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
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
	if v365 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v367 = F_get_rolespec_name(m, v281)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
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
	v370 = F_GetUserNameFromId(m, v25, int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v370
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v367
	F_errmsg(m, int32(_a_F_DelRoleMems_3), v22)
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_DelRoleMems_1), int32(2040), int32(_a_F_DelRoleMems_2))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
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
							F_errfinish(m, int32(_a_F_DropPreparedStatement_1), int32(456), int32(_a_F_DropPreparedStatement_2))
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
						F_errfinish(m, int32(_a_F_DropPreparedStatement_1), int32(456), int32(_a_F_DropPreparedStatement_2))
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
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	if l0 != 0 {
		v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		F_lose_s(m, v2)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			F_SN_delete_env(m, l0)
			mBase = m.M
			v6 = m.ExcPending
			if v6 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_SN_delete_env(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
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
					F_errfinish(m, int32(_a_F_datadir_fsync_fname_2), int32(3817), int32(_a_F_datadir_fsync_fname_3))
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
func F_dcbrt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v20 float64
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v76 float64
	_ = v76
	var v82 float64
	_ = v82
	var v84 float64
	_ = v84
	var v92 float64
	_ = v92
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v6 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v6))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v13) {
		v82 = base.F64_add(v6, v6)
	} else {
		if base.Ui32(int32(_a_F_dcbrt_0)) < base.Ui32(v13) {
			v30 = v13
			v31 = v6
			v32 = int32(715094163)
			v34 = base.I32_div_u_s(v30, int32(3))
			v40 = base.F64_copysign(base.F64_reinterpret_i64(base.I64_extend_i32_u(v32+v34)<<(uint(int64(32))%64)), v31)
			v43 = base.F64_mul(base.F64_mul(v40, v40), base.F64_div(v40, v6))
			v65 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_mul(v40, base.F64_add(base.F64_mul(base.F64_mul(v43, base.F64_mul(v43, v43)), base.F64_add(base.F64_mul(v43, float64(0.14599619288661245)), float64(-0.758397934778766))), base.F64_add(base.F64_mul(v43, base.F64_add(base.F64_mul(v43, float64(1.6214297201053545)), float64(-1.8849797954337717))), float64(1.87595182427177)))))&int64(-1073741824) + int64(2147483648))
			v67 = base.F64_div(v6, base.F64_mul(v65, v65))
			v76 = base.F64_add(base.F64_mul(v65, base.F64_div(base.F64_sub(v67, v65), base.F64_add(base.F64_add(v65, v65), v67))), v65)
		} else {
			v20 = base.F64_mul(v6, float64(1.8014398509481984e+16))
			v26 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v20))>>(uint(int64(32))%64))) & int32(2147483647)
			if v26 == int32(0) {
				v76 = v6
			} else {
				v30 = v26
				v31 = v20
				v32 = int32(696219795)
				v34 = base.I32_div_u_s(v30, int32(3))
				v40 = base.F64_copysign(base.F64_reinterpret_i64(base.I64_extend_i32_u(v32+v34)<<(uint(int64(32))%64)), v31)
				v43 = base.F64_mul(base.F64_mul(v40, v40), base.F64_div(v40, v6))
				v65 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_mul(v40, base.F64_add(base.F64_mul(base.F64_mul(v43, base.F64_mul(v43, v43)), base.F64_add(base.F64_mul(v43, float64(0.14599619288661245)), float64(-0.758397934778766))), base.F64_add(base.F64_mul(v43, base.F64_add(base.F64_mul(v43, float64(1.6214297201053545)), float64(-1.8849797954337717))), float64(1.87595182427177)))))&int64(-1073741824) + int64(2147483648))
				v67 = base.F64_div(v6, base.F64_mul(v65, v65))
				v76 = base.F64_add(base.F64_mul(v65, base.F64_div(base.F64_sub(v67, v65), base.F64_add(base.F64_add(v65, v65), v67))), v65)
			}
		}
		v82 = v76
	}
	v84 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v82), v84)&base.F64_ne(base.F64_abs(v6), v84) == int32(0) {
		v92 = float64(0)
		if base.F64_eq(v82, v92)&base.F64_ne(v6, v92) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v104 = m.ExcPending
			if v104 != 0 {
				return int64(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.I64_reinterpret_f64(v82)
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v102 = m.ExcPending
		if v102 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dcosd(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 float64
	_ = v41
	var v44 int64
	_ = v44
	var v51 float64
	_ = v51
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v92 int64
	_ = v92
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int32
	_ = v108
	var v113 int64
	_ = v113
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v125 int64
	_ = v125
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v140 int64
	_ = v140
	var v143 int32
	_ = v143
	var v158 int64
	_ = v158
	var v166 float64
	_ = v166
	var v170 float64
	_ = v170
	var v174 float64
	_ = v174
	var v178 float64
	_ = v178
	var v183 float64
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v203 float64
	_ = v203
	var v207 int32
	_ = v207
	var v208 float64
	_ = v208
	var v209 float64
	_ = v209
	var v214 float64
	_ = v214
	var v216 float64
	_ = v216
	var v218 float64
	_ = v218
	var v221 float64
	_ = v221
	var v225 float64
	_ = v225
	var v229 float64
	_ = v229
	var v233 float64
	_ = v233
	var v242 float64
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v262 float64
	_ = v262
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v268 float64
	_ = v268
	var v274 float64
	_ = v274
	var v275 float64
	_ = v275
	var v277 float64
	_ = v277
	var v279 float64
	_ = v279
	var v281 float64
	_ = v281
	var v290 float64
	_ = v290
	var v294 float64
	_ = v294
	var v301 float64
	_ = v301
	var v305 int64
	_ = v305
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
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v12&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
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
		goto L97
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L92
	} else {
		goto L93
	}
L3:
	;
	v17 = base.F64_reinterpret_i64(v12)
	if base.F64_eq(base.F64_abs(v17), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v305 = int64(9221120237041090560)
	goto L5
L5:
	;
	m.G0 = v9 + int32(16)
	return v305
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dcosd[0])))
	if v22 == int32(0) {
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
	v32 = base.I64_reinterpret_f64(v17)
	v36 = int32(2047)
	v37 = base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(52))%64))) & v36
	if v37 == v36 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v294), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L88
	}
L11:
	;
	if base.F64_lt(v166, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L12:
	;
	v41 = base.F64_mul(v17, float64(360))
	v166 = base.F64_div(v41, v41)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v44 = v32 << (uint(int64(1)) % 64)
	if base.Ui64(v44) <= base.Ui64(int64(-9156662467374350336)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if v44 == int64(-9156662467374350336) {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	if v37 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v51 = base.F64_mul(v17, float64(0))
	goto L20
L19:
	;
	v51 = v17
	goto L20
L20:
	;
	v166 = v51
	goto L11
L21:
	;
	if int32(1031) < v87 {
		goto L31
	} else {
		goto L32
	}
L22:
	;
	v54 = int32(0)
	v56 = v32 << (uint(int64(12)) % 64)
	if int64(0) <= v56 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v87 = v37
	v88 = v32&int64(4503599627370495) | int64(4503599627370496)
	goto L21
L25:
	;
	v60 = v56
	v63 = v54
	goto L28
L26:
	;
	v74 = v54
	goto L27
L27:
	;
	v87 = v74
	v88 = v32 << (uint(base.I64_extend_i32_u(int32(1)-v74)) % 64)
	goto L21
L28:
	;
	v65 = v63 - int32(1)
	v67 = v60 << (uint(int64(1)) % 64)
	if int64(0) <= v67 {
		v60 = v67
		v63 = v65
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v74 = v65
	goto L27
L30:
	;
	goto L29
L31:
	;
	v92 = v88
	v95 = v87
	goto L34
L32:
	;
	v113 = v88
	v116 = v87
	goto L33
L33:
	;
	v118 = v113 - int64(6333186975989760)
	if v118 < int64(0) {
		v125 = v113
		goto L40
	} else {
		goto L41
	}
L34:
	;
	v97 = v92 - int64(6333186975989760)
	if v97 < int64(0) {
		v104 = v92
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v113 = v106
	v116 = int32(1031)
	goto L33
L36:
	;
	v106 = v104 << (uint(int64(1)) % 64)
	v108 = v95 - int32(1)
	if int32(1031) < v108 {
		v92 = v106
		v95 = v108
		goto L34
	} else {
		goto L39
	}
L37:
	;
	if v97 != int64(0) {
		v104 = v97
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v166 = base.F64_mul(v17, float64(0))
	goto L11
L39:
	;
	goto L35
L40:
	;
	if base.Ui64(v125) <= base.Ui64(int64(4503599627370495)) {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	if v118 != int64(0) {
		v125 = v118
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v166 = base.F64_mul(v17, float64(0))
	goto L11
L43:
	;
	v129 = v125
	v132 = v116
	goto L46
L44:
	;
	v140 = v125
	v143 = v116
	goto L45
L45:
	;
	if int32(0) < v143 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v134 = v132 - int32(1)
	v136 = v129 << (uint(int64(1)) % 64)
	if base.Ui64(v129) < base.Ui64(int64(2251799813685248)) {
		v129 = v136
		v132 = v134
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v140 = v136
	v143 = v134
	goto L45
L48:
	;
	goto L47
L49:
	;
	v158 = v140 - int64(4503599627370496) | base.I64_extend_i32_u(v143)<<(uint(int64(52))%64)
	goto L51
L50:
	;
	v158 = int64(base.Ui64(v140) >> (uint(base.I64_extend_i32_u(int32(1)-v143)) % 64))
	goto L51
L51:
	;
	v166 = base.F64_reinterpret_i64(v32&int64(-9223372036854775807-1) | v158)
	goto L11
L52:
	;
	v170 = base.F64_neg(v166)
	goto L54
L53:
	;
	v170 = v166
	goto L54
L54:
	;
	if base.F64_gt(v170, float64(180)) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v174 = base.F64_sub(float64(360), v170)
	goto L57
L56:
	;
	v174 = v170
	goto L57
L57:
	;
	if base.F64_gt(v174, float64(90)) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v178 = base.F64_sub(float64(180), v174)
	goto L60
L59:
	;
	v178 = v174
	goto L60
L60:
	;
	if base.F64_le(v178, float64(60)) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v183 = base.F64_mul(v178, float64(0.017453292519943295))
	v187 = m.G0
	v189 = v187 - int32(16)
	m.G0 = v189
	v196 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v183))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v196) <= base.Ui32(int32(1072243195)) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L63
L63:
	;
	v242 = base.F64_mul(base.F64_sub(float64(90), v178), float64(0.017453292519943295))
	v246 = m.G0
	v248 = v246 - int32(16)
	m.G0 = v248
	v255 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v242))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v255) <= base.Ui32(int32(1072243195)) {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v229 = base.F64_sub(float64(1), v225)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v229
	v233 = *(*float64)(unsafe.Add(mBase, _c_F_dcosd[1]))
	v294 = base.F64_add(base.F64_mul(base.F64_div(v229, v233), float64(-0.5)), float64(1))
	goto L10
L65:
	;
	m.G0 = v189 + int32(16)
	goto L64
L66:
	;
	if base.Ui32(v196) < base.Ui32(int32(1044816030)) {
		v225 = float64(1)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v196) {
		v225 = base.F64_sub(v183, v183)
		goto L65
	} else {
		goto L70
	}
L69:
	;
	v203 = F___cos(m, v183, float64(0))
	mBase = m.M
	v225 = v203
	goto L65
L70:
	;
	v207 = F___rem_pio2(m, v183, v189)
	mBase = m.M
	v208 = *(*float64)(unsafe.Add(mBase, uint32(v189)+8))
	v209 = *(*float64)(unsafe.Add(mBase, uint32(v189)))
	switch v207&int32(3) - int32(1) {
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
	v221 = F___sin(m, v209, v208, int32(1))
	mBase = m.M
	v225 = v221
	goto L65
L72:
	;
	v218 = F___cos(m, v209, v208)
	mBase = m.M
	v225 = base.F64_neg(v218)
	goto L65
L73:
	;
	v216 = F___sin(m, v209, v208, int32(1))
	mBase = m.M
	v225 = base.F64_neg(v216)
	goto L65
L74:
	;
	v214 = F___cos(m, v209, v208)
	mBase = m.M
	v225 = v214
	goto L65
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v281
	v290 = *(*float64)(unsafe.Add(mBase, _c_F_dcosd[2]))
	v294 = base.F64_mul(base.F64_div(v281, v290), float64(0.5))
	goto L10
L76:
	;
	m.G0 = v248 + int32(16)
	goto L75
L77:
	;
	if base.Ui32(v255) < base.Ui32(int32(1045430272)) {
		v281 = v242
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v255) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v262 = F___sin(m, v242, float64(0), int32(0))
	mBase = m.M
	v281 = v262
	goto L76
L81:
	;
	v281 = base.F64_sub(v242, v242)
	goto L76
L82:
	;
	goto L83
L83:
	;
	v266 = F___rem_pio2(m, v242, v248)
	mBase = m.M
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v248)+8))
	v268 = *(*float64)(unsafe.Add(mBase, uint32(v248)))
	switch v266&int32(3) - int32(1) {
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
	v279 = F___cos(m, v268, v267)
	mBase = m.M
	v281 = base.F64_neg(v279)
	goto L76
L85:
	;
	v277 = F___sin(m, v268, v267, int32(1))
	mBase = m.M
	v281 = base.F64_neg(v277)
	goto L76
L86:
	;
	v275 = F___cos(m, v268, v267)
	mBase = m.M
	v281 = v275
	goto L76
L87:
	;
	v274 = F___sin(m, v268, v267, int32(1))
	mBase = m.M
	v281 = v274
	goto L76
L88:
	;
	if base.F64_gt(v174, float64(90)) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v301 = base.F64_neg(v294)
	goto L91
L90:
	;
	v301 = v294
	goto L91
L91:
	;
	v305 = base.I64_reinterpret_f64(v301)
	goto L5
L92:
	;
	return int64(0)
L93:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(_a_F_dcosd_0), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(_a_F_dcosd_1), int32(2375), int32(_a_F_dcosd_2))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L92
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dcosh(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v8 float64
	_ = v8
	var v9 int64
	_ = v9
	var v21 int64
	_ = v21
	var v26 int32
	_ = v26
	var v62 int32
	_ = v62
	var v63 float64
	_ = v63
	var v69 int32
	_ = v69
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v101 float64
	_ = v101
	var v104 float64
	_ = v104
	var v110 float64
	_ = v110
	var v119 float64
	_ = v119
	var v134 float64
	_ = v134
	var v143 float64
	_ = v143
	var v148 float64
	_ = v148
	var v155 float64
	_ = v155
	var v158 float64
	_ = v158
	var v164 float64
	_ = v164
	var v174 float64
	_ = v174
	var v176 float64
	_ = v176
	var v188 float64
	_ = v188
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v198 float64
	_ = v198
	var v205 float64
	_ = v205
	var v209 float64
	_ = v209
	var v214 float64
	_ = v214
	var v220 int32
	_ = v220
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, _c_F_dcosh[0])) = int32(0)
	v8 = base.F64_abs(v4)
	v9 = base.I64_reinterpret_f64(v8)
	if base.Ui64(v9) <= base.Ui64(int64(4604418530035630079)) {
		if base.Ui64(v9) < base.Ui64(int64(4490088828488384512)) {
			v214 = float64(1)
		} else {
			v21 = base.I64_reinterpret_f64(v8)
			v26 = base.I32_wrap_i64(int64(base.Ui64(v21)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v26) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v8)&int64(9223372036854775807)) {
					v176 = v8
					v188 = v176
				} else {
					if v21 < int64(0) {
						v188 = float64(-1)
					} else {
						if base.F64_gt(v8, float64(709.782712893384)) == int32(0) {
							v62 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v8, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v8)))
							v63 = base.F64_convert_i32_s(v62)
							v69 = v62
							v70 = base.F64_mul(v63, float64(1.9082149292705877e-10))
							v72 = base.F64_add(v8, base.F64_mul(v63, float64(-0.6931471803691238)))
							v73 = base.F64_sub(v72, v70)
							v79 = v73
							v80 = v69
							v81 = base.F64_sub(base.F64_sub(v72, v73), v70)
							v84 = base.F64_mul(v79, float64(0.5))
							v85 = base.F64_mul(v79, v84)
							v101 = base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v104 = base.F64_sub(float64(3), base.F64_mul(v101, v84))
							v110 = base.F64_mul(v85, base.F64_div(base.F64_sub(v101, v104), base.F64_sub(float64(6), base.F64_mul(v79, v104))))
							if v80 == int32(0) {
								v188 = base.F64_sub(v79, base.F64_sub(base.F64_mul(v79, v110), v85))
							} else {
								v119 = base.F64_sub(base.F64_sub(base.F64_mul(v79, base.F64_sub(v110, v81)), v81), v85)
								switch v80 + int32(1) {
								case 0:
									v188 = base.F64_add(base.F64_mul(base.F64_sub(v79, v119), float64(0.5)), float64(-0.5))
								default:
									v143 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v80+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v80) {
										v148 = base.F64_add(base.F64_sub(v79, v119), float64(1))
										if v80 == int32(1024) {
											v155 = base.F64_mul(base.F64_add(v148, v148), float64(8.98846567431158e+307))
										} else {
											v155 = base.F64_mul(v148, v143)
										}
										v188 = base.F64_add(v155, float64(-1))
									} else {
										v158 = float64(1)
										v164 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v80) << (uint(int64(52)) % 64))
										if base.Ui32(v80) <= base.Ui32(int32(19)) {
											v174 = base.F64_add(base.F64_sub(v158, v164), base.F64_sub(v79, v119))
										} else {
											v174 = base.F64_add(base.F64_sub(v79, base.F64_add(v119, v164)), v158)
										}
										v176 = base.F64_mul(v174, v143)
										v188 = v176
									}
								case 2:
									if base.F64_lt(v79, float64(-0.25)) != 0 {
										v188 = base.F64_mul(base.F64_sub(v119, base.F64_add(v79, float64(0.5))), float64(-2))
									} else {
										v134 = base.F64_sub(v79, v119)
										v188 = base.F64_add(base.F64_add(v134, v134), float64(1))
									}
								}
							}
						} else {
							v188 = base.F64_mul(v8, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v26) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v26) < base.Ui32(int32(1016070144)) {
						v176 = v8
						v188 = v176
					} else {
						v79 = v8
						v80 = int32(0)
						v81 = float64(0)
						v84 = base.F64_mul(v79, float64(0.5))
						v85 = base.F64_mul(v79, v84)
						v101 = base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v104 = base.F64_sub(float64(3), base.F64_mul(v101, v84))
						v110 = base.F64_mul(v85, base.F64_div(base.F64_sub(v101, v104), base.F64_sub(float64(6), base.F64_mul(v79, v104))))
						if v80 == int32(0) {
							v188 = base.F64_sub(v79, base.F64_sub(base.F64_mul(v79, v110), v85))
						} else {
							v119 = base.F64_sub(base.F64_sub(base.F64_mul(v79, base.F64_sub(v110, v81)), v81), v85)
							switch v80 + int32(1) {
							case 0:
								v188 = base.F64_add(base.F64_mul(base.F64_sub(v79, v119), float64(0.5)), float64(-0.5))
							default:
								v143 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v80+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v80) {
									v148 = base.F64_add(base.F64_sub(v79, v119), float64(1))
									if v80 == int32(1024) {
										v155 = base.F64_mul(base.F64_add(v148, v148), float64(8.98846567431158e+307))
									} else {
										v155 = base.F64_mul(v148, v143)
									}
									v188 = base.F64_add(v155, float64(-1))
								} else {
									v158 = float64(1)
									v164 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v80) << (uint(int64(52)) % 64))
									if base.Ui32(v80) <= base.Ui32(int32(19)) {
										v174 = base.F64_add(base.F64_sub(v158, v164), base.F64_sub(v79, v119))
									} else {
										v174 = base.F64_add(base.F64_sub(v79, base.F64_add(v119, v164)), v158)
									}
									v176 = base.F64_mul(v174, v143)
									v188 = v176
								}
							case 2:
								if base.F64_lt(v79, float64(-0.25)) != 0 {
									v188 = base.F64_mul(base.F64_sub(v119, base.F64_add(v79, float64(0.5))), float64(-2))
								} else {
									v134 = base.F64_sub(v79, v119)
									v188 = base.F64_add(base.F64_add(v134, v134), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v26) {
						v62 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v8, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v8)))
						v63 = base.F64_convert_i32_s(v62)
						v69 = v62
						v70 = base.F64_mul(v63, float64(1.9082149292705877e-10))
						v72 = base.F64_add(v8, base.F64_mul(v63, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v21 {
							v69 = int32(1)
							v70 = float64(1.9082149292705877e-10)
							v72 = base.F64_add(v8, float64(-0.6931471803691238))
						} else {
							v69 = int32(-1)
							v70 = float64(-1.9082149292705877e-10)
							v72 = base.F64_add(v8, float64(0.6931471803691238))
						}
					}
					v73 = base.F64_sub(v72, v70)
					v79 = v73
					v80 = v69
					v81 = base.F64_sub(base.F64_sub(v72, v73), v70)
					v84 = base.F64_mul(v79, float64(0.5))
					v85 = base.F64_mul(v79, v84)
					v101 = base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, base.F64_add(base.F64_mul(v85, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v104 = base.F64_sub(float64(3), base.F64_mul(v101, v84))
					v110 = base.F64_mul(v85, base.F64_div(base.F64_sub(v101, v104), base.F64_sub(float64(6), base.F64_mul(v79, v104))))
					if v80 == int32(0) {
						v188 = base.F64_sub(v79, base.F64_sub(base.F64_mul(v79, v110), v85))
					} else {
						v119 = base.F64_sub(base.F64_sub(base.F64_mul(v79, base.F64_sub(v110, v81)), v81), v85)
						switch v80 + int32(1) {
						case 0:
							v188 = base.F64_add(base.F64_mul(base.F64_sub(v79, v119), float64(0.5)), float64(-0.5))
						default:
							v143 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v80+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v80) {
								v148 = base.F64_add(base.F64_sub(v79, v119), float64(1))
								if v80 == int32(1024) {
									v155 = base.F64_mul(base.F64_add(v148, v148), float64(8.98846567431158e+307))
								} else {
									v155 = base.F64_mul(v148, v143)
								}
								v188 = base.F64_add(v155, float64(-1))
							} else {
								v158 = float64(1)
								v164 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v80) << (uint(int64(52)) % 64))
								if base.Ui32(v80) <= base.Ui32(int32(19)) {
									v174 = base.F64_add(base.F64_sub(v158, v164), base.F64_sub(v79, v119))
								} else {
									v174 = base.F64_add(base.F64_sub(v79, base.F64_add(v119, v164)), v158)
								}
								v176 = base.F64_mul(v174, v143)
								v188 = v176
							}
						case 2:
							if base.F64_lt(v79, float64(-0.25)) != 0 {
								v188 = base.F64_mul(base.F64_sub(v119, base.F64_add(v79, float64(0.5))), float64(-2))
							} else {
								v134 = base.F64_sub(v79, v119)
								v188 = base.F64_add(base.F64_add(v134, v134), float64(1))
							}
						}
					}
				}
			}
			v190 = float64(1)
			v191 = base.F64_add(v188, v190)
			v214 = base.F64_add(base.F64_div(base.F64_mul(v188, v188), base.F64_add(v191, v191)), v190)
		}
	} else {
		if base.Ui64(v9) <= base.Ui64(int64(4649454526309335039)) {
			v198 = F_exp(m, v8)
			mBase = m.M
			v214 = base.F64_mul(base.F64_add(v198, base.F64_div(float64(1), v198)), float64(0.5))
		} else {
			v205 = float64(2.247116418577895e+307)
			v209 = F_exp(m, base.F64_add(v8, float64(-1416.0996898839683)))
			mBase = m.M
			v214 = base.F64_mul(base.F64_mul(base.F64_mul(float64(1), v205), v209), v205)
		}
	}
	if base.F64_eq(v214, float64(0)) != 0 {
		F_float_underflow_error(m)
		mBase = m.M
		v220 = m.ExcPending
		if v220 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I64_reinterpret_f64(v214)
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
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
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int64
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v612 int32
	_ = v612
	v10 = m.G0
	v12 = v10 - int32(208)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+196)) = int32(0)
	F_init_work(m, v12+int32(196), l1, l5, v12+int32(152))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v28 = v26 & v24
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v29 = v24
	goto L5
L4:
	;
	v29 = int32(4)
	goto L5
L5:
	;
	if v26 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v58 = F_mbuf_create_from_data(m, l2+v29, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v36 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v47 = int32(1)
	if v28 != 0 {
		v57 = int32(base.Ui32(v26)>>(uint(v47)%32)) - v47
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v39 = int32(16)
	goto L12
L11:
	;
	v39 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = int32(4)
	goto L15
L14:
	;
	v46 = v39
	goto L15
L15:
	;
	v57 = v46
	goto L6
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v60 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v88 = F_mbuf_create(m, v85+int32(2048))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L1
	} else {
		goto L29
	}
L19:
	;
	v64 = int32(18)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+1)))
	if v66 == v64 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v77 = int32(1)
	if v60&v77 != 0 {
		v85 = int32(base.Ui32(v60) >> (uint(v77) % 32))
		goto L18
	} else {
		goto L28
	}
L22:
	;
	v69 = v64
	goto L24
L23:
	;
	v69 = int32(2)
	goto L24
L24:
	;
	if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v76 = int32(6)
	goto L27
L26:
	;
	v76 = v69
	goto L27
L27:
	;
	v85 = v76
	goto L18
L28:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v85 = int32(base.Ui32(v81) >> (uint(int32(2)) % 32))
	goto L18
L29:
	;
	v93 = F_mbuf_append(m, v88, v12+int32(204), int32(4))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
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
	v231 = int32(0)
	if v231 <= v227 {
		goto L87
	} else {
		goto L88
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
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v180 = int32(1)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v184 = v182 & v180
	if v184 != 0 {
		goto L69
	} else {
		goto L70
	}
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	if v135 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v97 = int32(0)
	v131 = v97
	v134 = v97
	goto L35
L37:
	;
	goto L38
L38:
	;
	v99 = int32(1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v103 = v101 & v99
	if v103 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v104 = v99
	goto L41
L40:
	;
	v104 = int32(4)
	goto L41
L41:
	;
	v105 = l4 + v104
	if v101 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4)+1)))
	if v111 == int32(18) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v122 = int32(1)
	if v103 != 0 {
		v131 = v105
		v134 = int32(base.Ui32(v101)>>(uint(v122)%32)) - v122
		goto L35
	} else {
		goto L51
	}
L45:
	;
	v114 = int32(16)
	goto L47
L46:
	;
	v114 = int32(0)
	goto L47
L47:
	;
	if base.Ui32((v111-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v121 = int32(4)
	goto L50
L49:
	;
	v121 = v114
	goto L50
L50:
	;
	v131 = v105
	v134 = v121
	goto L35
L51:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v131 = v105
	v134 = int32(base.Ui32(v126)>>(uint(int32(2))%32)) - int32(4)
	goto L35
L52:
	;
	v165 = int32(1)
	if v135&v165 != 0 {
		goto L63
	} else {
		goto L64
	}
L53:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v141 == int32(18) {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	v152 = int32(1)
	if v135&v152 != 0 {
		v164 = int32(base.Ui32(v135)>>(uint(v152)%32)) - v152
		goto L52
	} else {
		goto L62
	}
L56:
	;
	v144 = int32(16)
	goto L58
L57:
	;
	v144 = int32(0)
	goto L58
L58:
	;
	if base.Ui32((v141-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v151 = int32(4)
	goto L61
L60:
	;
	v151 = v144
	goto L61
L61:
	;
	v164 = v151
	goto L52
L62:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v164 = int32(base.Ui32(v158)>>(uint(int32(2))%32)) - int32(4)
	goto L52
L63:
	;
	v169 = v165
	goto L65
L64:
	;
	v169 = int32(4)
	goto L65
L65:
	;
	v171 = F_mbuf_create_from_data(m, l3+v169, v164)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v175 = F_pgp_set_pubkey(m, v173, v171, v131, v134, int32(1))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v177 = F_mbuf_free(m, v171)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v227 = v175
	goto L31
L69:
	;
	v185 = v180
	goto L71
L70:
	;
	v185 = int32(4)
	goto L71
L71:
	;
	v186 = l3 + v185
	if v182 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v214 = int32(0)
	if base.B2i32(v186 == v214)|base.B2i32(v213 <= v214) != 0 {
		goto L84
	} else {
		goto L85
	}
L73:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3)+1)))
	if v192 == int32(18) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v203 = int32(1)
	if v184 != 0 {
		v213 = int32(base.Ui32(v182)>>(uint(v203)%32)) - v203
		goto L72
	} else {
		goto L82
	}
L76:
	;
	v195 = int32(16)
	goto L78
L77:
	;
	v195 = int32(0)
	goto L78
L78:
	;
	if base.Ui32((v192-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v202 = int32(4)
	goto L81
L80:
	;
	v202 = v195
	goto L81
L81:
	;
	v213 = v202
	goto L72
L82:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v213 = int32(base.Ui32(v207)>>(uint(int32(2))%32)) - int32(4)
	goto L72
L83:
	;
	v227 = v225
	goto L31
L84:
	;
	v225 = int32(-13)
	goto L86
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v179)+132)) = v213
	*(*int32)(unsafe.Add(mBase, uint32(v179)+128)) = v186
	v225 = int32(0)
	goto L86
L86:
	;
	goto L83
L87:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v235 = F_pgp_decrypt(m, v234, v58, v88)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L90
	}
L88:
	;
	v484 = v227
	v485 = v231
	goto L89
L89:
	;
	v486 = F_mbuf_free(m, v58)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L156
	}
L90:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v12)+156))
	if v237 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+88))
	v484 = v235
	v485 = base.B2i32(v480 != int32(0))
	goto L89
L92:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
	if v241 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v12)+164))
	if v267 < int32(0) {
		goto L100
	} else {
		goto L101
	}
L94:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v240)+60))
	if v241 == v244 {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v248 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	if v248 == int32(0) {
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v240)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v252
	*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = v241
	*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = int32(_a_F_decrypt_internal_5)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(128))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(149), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L93
L100:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v12)+168))
	if v293 < int32(0) {
		goto L107
	} else {
		goto L108
	}
L101:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	if v267 == v270 {
		goto L100
	} else {
		goto L102
	}
L102:
	;
	v274 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	if v274 == int32(0) {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v240)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+120)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v12)+116)) = v267
	*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = int32(_a_F_decrypt_internal_6)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(112))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(150), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	goto L100
L107:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v12)+176))
	if v319 < int32(0) {
		goto L114
	} else {
		goto L115
	}
L108:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	if v293 == v296 {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v300 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	if v300 == int32(0) {
		goto L107
	} else {
		goto L111
	}
L111:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+104)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v12)+100)) = v293
	*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = int32(_a_F_decrypt_internal_7)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(96))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(151), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	goto L107
L114:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v12)+184))
	if v345 < int32(0) {
		goto L121
	} else {
		goto L122
	}
L115:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v240)+52))
	if v319 == v322 {
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v326 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	if v326 == int32(0) {
		goto L114
	} else {
		goto L118
	}
L118:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v240)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = v330
	*(*int32)(unsafe.Add(mBase, uint32(v12)+84)) = v319
	*(*int32)(unsafe.Add(mBase, uint32(v12)+80)) = int32(_a_F_decrypt_internal_8)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(80))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(152), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	goto L114
L121:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v240)+76))
	if v371 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L122:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v240)+76))
	if v345 == v348 {
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v352 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	if v352 == int32(0) {
		goto L121
	} else {
		goto L125
	}
L125:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v240)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v356
	*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(_a_F_decrypt_internal_9)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12-int32(-64))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(153), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L121
L128:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
	if v401 < int32(0) {
		goto L136
	} else {
		goto L137
	}
L129:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v12)+172))
	if v374 < int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v240)+56))
	if v374 == v377 {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v381 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	if v381 == int32(0) {
		goto L128
	} else {
		goto L133
	}
L133:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v240)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v385
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v374
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(_a_F_decrypt_internal_10)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(48))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(155), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	goto L128
L136:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v12)+180))
	if v427 < int32(0) {
		goto L143
	} else {
		goto L144
	}
L137:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v240)+72))
	if v401 == v404 {
		goto L136
	} else {
		goto L138
	}
L138:
	;
	v408 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	if v408 == int32(0) {
		goto L136
	} else {
		goto L140
	}
L140:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v240)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v412
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(_a_F_decrypt_internal_11)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(32))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(156), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	goto L136
L143:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v12)+192))
	if v453 < int32(0) {
		goto L91
	} else {
		goto L150
	}
L144:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v240)+64))
	if v427 == v430 {
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v434 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	if v434 == int32(0) {
		goto L143
	} else {
		goto L147
	}
L147:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v240)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v438
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v427
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(_a_F_decrypt_internal_4)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12+int32(16))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(157), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	goto L143
L150:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v240)+88))
	if v453 == v456 {
		goto L91
	} else {
		goto L151
	}
L151:
	;
	v460 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v460 == int32(0) {
		goto L91
	} else {
		goto L153
	}
L153:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v240)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v464
	*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v453
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_decrypt_internal_0)
	F_errmsg(m, int32(_a_F_decrypt_internal_1), v12)
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	F_errfinish(m, int32(_a_F_decrypt_internal_2), int32(158), int32(_a_F_decrypt_internal_3))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	goto L91
L156:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v12)+196))
	v489 = F_pgp_free(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	if v484 == int32(0) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v496 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v88)+16)) = uint16(v496)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(200)))) = v499
	v501 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v501
	*(*int64)(unsafe.Add(mBase, uint32(v88))) = v501
	goto L161
L159:
	;
	goto L160
L160:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decrypt_internal[0])) = int32(0)
	goto L208
L161:
	;
	v506 = F_mbuf_free(m, v88)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v12)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v508))) = (v498 - v499) << (uint(int32(2)) % 32)
	v512 = int32(0)
	if v485&base.B2i32(l1 != v512) == v512 {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_decrypt_internal[0])) = int32(0)
	goto L207
L164:
	;
	v597 = v508
	goto L163
L165:
	;
	goto L166
L166:
	;
	v518 = *(*int32)(unsafe.Add(mBase, _c_F_decrypt_internal[1]))
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v518)+4))
	goto L167
L167:
	;
	v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	if v520 == int32(1) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	v550 = int32(1)
	if v520&v550 != 0 {
		goto L179
	} else {
		goto L180
	}
L169:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	if v526 == int32(18) {
		goto L172
	} else {
		goto L173
	}
L170:
	;
	goto L171
L171:
	;
	v537 = int32(1)
	if v520&v537 != 0 {
		v549 = int32(base.Ui32(v520)>>(uint(v537)%32)) - v537
		goto L168
	} else {
		goto L178
	}
L172:
	;
	v529 = int32(16)
	goto L174
L173:
	;
	v529 = int32(0)
	goto L174
L174:
	;
	if base.Ui32((v526-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L175
	} else {
		goto L176
	}
L175:
	;
	v536 = int32(4)
	goto L177
L176:
	;
	v536 = v529
	goto L177
L177:
	;
	v549 = v536
	goto L168
L178:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v549 = int32(base.Ui32(v543)>>(uint(int32(2))%32)) - int32(4)
	goto L168
L179:
	;
	v554 = v550
	goto L181
L180:
	;
	v554 = int32(4)
	goto L181
L181:
	;
	v555 = v508 + v554
	v557 = F_pg_do_encoding_conversion(m, v555, v549, int32(6), v519)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	if v557 == v555 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v597 = v508
	goto L163
L184:
	;
	goto L185
L185:
	;
	v560 = F_cstring_to_text(m, v557)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	F_pfree(m, v557)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if v508 == v560 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v597 = v508
	goto L163
L189:
	;
	goto L190
L190:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508))))
	if v566 == int32(1) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	if v591 != 0 {
		goto L203
	} else {
		goto L204
	}
L192:
	;
	v570 = int32(18)
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v508)+1)))
	if v572 == v570 {
		goto L195
	} else {
		goto L196
	}
L193:
	;
	goto L194
L194:
	;
	v583 = int32(1)
	if v566&v583 != 0 {
		v591 = int32(base.Ui32(v566) >> (uint(v583) % 32))
		goto L191
	} else {
		goto L201
	}
L195:
	;
	v575 = v570
	goto L197
L196:
	;
	v575 = int32(2)
	goto L197
L197:
	;
	if base.Ui32((v572-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L198
	} else {
		goto L199
	}
L198:
	;
	v582 = int32(6)
	goto L200
L199:
	;
	v582 = v575
	goto L200
L200:
	;
	v591 = v582
	goto L191
L201:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v591 = int32(base.Ui32(v587) >> (uint(int32(2)) % 32))
	goto L191
L202:
	;
	F_pfree(m, v508)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L1
	} else {
		goto L206
	}
L203:
	;
	base.MemoryFill(m, v508, int32(0), v591)
	goto L205
L204:
	;
	goto L205
L205:
	;
	goto L202
L206:
	;
	v597 = v560
	goto L163
L207:
	;
	m.G0 = v12 + int32(208)
	return v597
L208:
	;
	v609 = F_mbuf_free(m, v88)
	mBase = m.M
	v610 = m.ExcPending
	if v610 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	F_px_THROW_ERROR(m, v484)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
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
func F_dexp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 float64
	_ = v5
	var v13 float64
	_ = v13
	var v16 float64
	_ = v16
	var v22 float64
	_ = v22
	var v28 float64
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = base.F64_reinterpret_i64(v4)
	if base.Ui64(v4&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.F64_eq(base.F64_abs(v5), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v13 = float64(0)
			if base.F64_gt(v5, v13) != 0 {
				v16 = v5
			} else {
				v16 = v13
			}
			return base.I64_reinterpret_f64(v16)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_dexp[0])) = int32(0)
			v22 = F_exp(m, v5)
			mBase = m.M
			if base.F64_eq(base.F64_abs(v22), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v22, float64(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v28 = v22
					return base.I64_reinterpret_f64(v28)
				}
			}
		}
	} else {
		v28 = v5
		return base.I64_reinterpret_f64(v28)
	}
}
func F_difference(m *base.Module, l0 int32) int64 {
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v42 int32
	_ = v42
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var __phi71 int32
	_ = __phi71
	var v72 int32
	_ = v72
	var __phi72 int32
	_ = __phi72
	var v74 int32
	_ = v74
	var __phi74 int32
	_ = __phi74
	var v75 int32
	_ = v75
	var __phi75 int32
	_ = __phi75
	var v76 int32
	_ = v76
	var __phi76 int32
	_ = __phi76
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v224 int32
	_ = v224
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var __phi244 int32
	_ = __phi244
	var v245 int32
	_ = v245
	var __phi245 int32
	_ = __phi245
	var v247 int32
	_ = v247
	var __phi247 int32
	_ = __phi247
	var v248 int32
	_ = v248
	var __phi248 int32
	_ = __phi248
	var v249 int32
	_ = v249
	var __phi249 int32
	_ = __phi249
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v290 int32
	_ = v290
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v190 = F_pg_detoast_datum_packed(m, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L49
	}
L5:
	;
	if base.Ui32((v30-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	v28 = v19
	v30 = v27
	goto L9
L7:
	;
	goto L8
L8:
	;
	v51 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+4)) = uint8(v51)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v51
	goto L4
L9:
	;
	if base.Ui32((v30&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v42 != 0 {
		v28 = v28 + int32(1)
		v30 = v42
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v63 = v30 - int32(32)
	goto L15
L14:
	;
	v63 = v30
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v63)
	v65 = int32(1)
	v67 = v12 + int32(12)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v68 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v181 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v178))) = uint8(v181)
	goto L4
L17:
	;
	__phi71 = v28
	__phi72 = v68
	__phi74 = v67
	__phi75 = v65
	__phi76 = v28 + int32(1)
	v71 = __phi71
	v72 = __phi72
	v74 = __phi74
	v75 = __phi75
	v76 = __phi76
	goto L20
L18:
	;
	v167 = v67
	v168 = v65
	goto L19
L19:
	;
	v171 = int32(4) - v168
	if v171 != 0 {
		goto L46
	} else {
		goto L47
	}
L20:
	;
	if base.Ui32(int32(25)) < base.Ui32((v72&int32(223)-int32(65))&int32(255)) {
		v152 = v74
		v153 = v75
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if int32(3) < v153 {
		v178 = v152
		goto L16
	} else {
		goto L45
	}
L22:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)))
	if v155 != 0 {
		goto L41
	} else {
		goto L42
	}
L23:
	;
	if base.Ui32((v72-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = v72 - int32(32)
	goto L26
L25:
	;
	v93 = v72
	goto L26
L26:
	;
	v99 = base.B2i32(base.Ui32(int32(25)) < base.Ui32((v93-int32(65))&int32(255)))
	if base.Ui32(int32(25)) < base.Ui32((v93-int32(65))&int32(255)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v105 = v93
	goto L29
L28:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93&int32(255))+uint32(_c_F_difference[0]))))
	v105 = v104
	goto L29
L29:
	;
	v106 = int32(255)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if base.Ui32((v108-int32(97))&v106) < base.Ui32(int32(26)) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v117 = v108 - int32(32)
	goto L32
L31:
	;
	v117 = v108
	goto L32
L32:
	;
	if base.Ui32((v117-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117&int32(255))+uint32(_c_F_difference[0]))))
	v129 = v128
	goto L35
L34:
	;
	v129 = v117
	goto L35
L35:
	;
	if v105&v106 == v129&int32(255) {
		v152 = v74
		v153 = v75
		goto L22
	} else {
		goto L36
	}
L36:
	;
	if v99 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93&int32(255))+uint32(_c_F_difference[0]))))
	v140 = v139
	goto L39
L38:
	;
	v140 = v93
	goto L39
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v140)
	if v140&int32(255) == int32(48) {
		v152 = v74
		v153 = v75
		goto L22
	} else {
		goto L40
	}
L40:
	;
	v146 = int32(1)
	v152 = v74 + v146
	v153 = v75 + v146
	goto L22
L41:
	;
	if v153 < int32(4) {
		__phi71 = v76
		__phi72 = v155
		__phi74 = v152
		__phi75 = v153
		__phi76 = v76 + int32(1)
		v71 = __phi71
		v72 = __phi72
		v74 = __phi74
		v75 = __phi75
		v76 = __phi76
		goto L20
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	goto L21
L44:
	;
	goto L43
L45:
	;
	v167 = v152
	v168 = v153
	goto L19
L46:
	;
	base.MemoryFill(m, v167, int32(48), v171)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v178 = v171 + v167
	goto L16
L49:
	;
	v192 = F_text_to_cstring(m, v190)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v195 = v12 + int32(6)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v200 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+7)))
	v363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+6)))
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+11)))
	v366 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+8)))
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)))
	v368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+9)))
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
	m.G0 = v12 + int32(16)
	return base.I64_extend_i32_u(base.B2i32(v368 == v369)) + (base.I64_extend_i32_u(base.B2i32(v366 == v367)) + (base.I64_extend_i32_u(base.B2i32(v364 == v365)) + base.I64_extend_i32_u(base.B2i32(v362 == v363))))
L52:
	;
	if base.Ui32((v203-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	v201 = v192
	v203 = v200
	goto L56
L54:
	;
	goto L55
L55:
	;
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v195)+4)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v224
	goto L51
L56:
	;
	if base.Ui32((v203&int32(223)-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	if v215 != 0 {
		v201 = v201 + int32(1)
		v203 = v215
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v236 = v203 - int32(32)
	goto L62
L61:
	;
	v236 = v203
	goto L62
L62:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v195))) = uint8(v236)
	v238 = int32(1)
	v240 = v12 + int32(7)
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+1)))
	if v241 != 0 {
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v354 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v351))) = uint8(v354)
	goto L51
L64:
	;
	__phi244 = v201
	__phi245 = v241
	__phi247 = v240
	__phi248 = v238
	__phi249 = v201 + int32(1)
	v244 = __phi244
	v245 = __phi245
	v247 = __phi247
	v248 = __phi248
	v249 = __phi249
	goto L67
L65:
	;
	v340 = v240
	v341 = v238
	goto L66
L66:
	;
	v344 = int32(4) - v341
	if v344 != 0 {
		goto L93
	} else {
		goto L94
	}
L67:
	;
	if base.Ui32(int32(25)) < base.Ui32((v245&int32(223)-int32(65))&int32(255)) {
		v325 = v247
		v326 = v248
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if int32(3) < v326 {
		v351 = v325
		goto L63
	} else {
		goto L92
	}
L69:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	if v328 != 0 {
		goto L88
	} else {
		goto L89
	}
L70:
	;
	if base.Ui32((v245-int32(97))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v266 = v245 - int32(32)
	goto L73
L72:
	;
	v266 = v245
	goto L73
L73:
	;
	v272 = base.B2i32(base.Ui32(int32(25)) < base.Ui32((v266-int32(65))&int32(255)))
	if base.Ui32(int32(25)) < base.Ui32((v266-int32(65))&int32(255)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v278 = v266
	goto L76
L75:
	;
	v277 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266&int32(255))+uint32(_c_F_difference[0]))))
	v278 = v277
	goto L76
L76:
	;
	v279 = int32(255)
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v244))))
	if base.Ui32((v281-int32(97))&v279) < base.Ui32(int32(26)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v290 = v281 - int32(32)
	goto L79
L78:
	;
	v290 = v281
	goto L79
L79:
	;
	if base.Ui32((v290-int32(65))&int32(255)) <= base.Ui32(int32(25)) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v290&int32(255))+uint32(_c_F_difference[0]))))
	v302 = v301
	goto L82
L81:
	;
	v302 = v290
	goto L82
L82:
	;
	if v278&v279 == v302&int32(255) {
		v325 = v247
		v326 = v248
		goto L69
	} else {
		goto L83
	}
L83:
	;
	if v272 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v266&int32(255))+uint32(_c_F_difference[0]))))
	v313 = v312
	goto L86
L85:
	;
	v313 = v266
	goto L86
L86:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v247))) = uint8(v313)
	if v313&int32(255) == int32(48) {
		v325 = v247
		v326 = v248
		goto L69
	} else {
		goto L87
	}
L87:
	;
	v319 = int32(1)
	v325 = v247 + v319
	v326 = v248 + v319
	goto L69
L88:
	;
	if v326 < int32(4) {
		__phi244 = v249
		__phi245 = v328
		__phi247 = v325
		__phi248 = v326
		__phi249 = v249 + int32(1)
		v244 = __phi244
		v245 = __phi245
		v247 = __phi247
		v248 = __phi248
		v249 = __phi249
		goto L67
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	goto L68
L91:
	;
	goto L90
L92:
	;
	v340 = v325
	v341 = v326
	goto L66
L93:
	;
	base.MemoryFill(m, v340, int32(48), v344)
	goto L95
L94:
	;
	goto L95
L95:
	;
	v351 = v344 + v340
	goto L63
}
func F_dintdict_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_palloc0_mul(m, int32(8), int32(2))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(0)
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
		if v17 != int32(1) {
			v29 = F_pnstrdup(m, v7, v6)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = v6
				v32 = v29
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				if v31 <= v33 {
					v44 = v32
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
					return base.I64_extend_i32_u(v11)
				} else {
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
					if v35 == int32(1) {
						F_pfree(m, v32)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int64(0)
						} else {
							v44 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
							return base.I64_extend_i32_u(v11)
						}
					} else {
						v42 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v32+v33))) = uint8(v42)
						v44 = v32
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
						return base.I64_extend_i32_u(v11)
					}
				}
			}
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			switch v20 - int32(43) {
			case 0, 2:
				v23 = int32(1)
				v26 = v6 - v23
				v27 = F_pnstrdup(m, v7+v23, v26)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int64(0)
				} else {
					v31 = v26
					v32 = v27
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v31 <= v33 {
						v44 = v32
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
						return base.I64_extend_i32_u(v11)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
						if v35 == int32(1) {
							F_pfree(m, v32)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
								return base.I64_extend_i32_u(v11)
							}
						} else {
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v32+v33))) = uint8(v42)
							v44 = v32
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
							return base.I64_extend_i32_u(v11)
						}
					}
				}
			default:
				v29 = F_pnstrdup(m, v7, v6)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int64(0)
				} else {
					v31 = v6
					v32 = v29
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					if v31 <= v33 {
						v44 = v32
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
						return base.I64_extend_i32_u(v11)
					} else {
						v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
						if v35 == int32(1) {
							F_pfree(m, v32)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								v44 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
								return base.I64_extend_i32_u(v11)
							}
						} else {
							v42 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v32+v33))) = uint8(v42)
							v44 = v32
							*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v44
							return base.I64_extend_i32_u(v11)
						}
					}
				}
			}
		}
	}
}
func F_dispell_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int64
	_ = v11
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
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
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
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var __phi343 int32
	_ = __phi343
	var v345 int32
	_ = v345
	var __phi345 int32
	_ = __phi345
	var v346 int32
	_ = v346
	var __phi346 int32
	_ = __phi346
	var v348 int32
	_ = v348
	var __phi348 int32
	_ = __phi348
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int64
	_ = v362
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v374 int32
	_ = v374
	var v396 int64
	_ = v396
	v2 = int32(0)
	v11 = int64(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v12 <= v2 {
		v396 = v11
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v396
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = v15 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return int64(0)
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
		v74 = v2
		v75 = v2
		v80 = v24
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v87 = v2
	v88 = v2
	v93 = v24
	goto L8
L8:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+44)))
	if v97 != int32(1) {
		v329 = v88
		goto L23
	} else {
		goto L24
	}
L9:
	;
	F_pfree(m, v26)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v31 = v26
	v32 = v2
	v33 = v2
	v34 = v28
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
	v74 = v62
	v75 = v49
	v80 = v65
	goto L9
L13:
	;
	v46 = F_palloc_mul(m, int32(8), int32(1024))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v48 = v32
	v49 = v33
	goto L15
L15:
	;
	v50 = v48 - v49
	if v50 <= int32(_a_F_dispell_lexize_0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v48 = v46
	v49 = v46
	goto L15
L17:
	;
	v53 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v48)+12)) = v53
	*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v38)
	v60 = v48 + int32(8)
	v62 = v60
	v63 = v60 - v49
	goto L19
L18:
	;
	v62 = v48
	v63 = v50
	goto L19
L19:
	;
	v65 = v38 + int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v66 == int32(0) {
		v74 = v62
		v75 = v49
		v80 = v65
		goto L9
	} else {
		goto L20
	}
L20:
	;
	if v63 < int32(_a_F_dispell_lexize_1) {
		v31 = v31 + int32(4)
		v32 = v62
		v33 = v49
		v34 = v66
		v38 = v65
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	v87 = v74
	v88 = v75
	v93 = v80
	goto L8
L23:
	;
	if v329 == int32(0) {
		v396 = v11
		goto L1
	} else {
		goto L78
	}
L24:
	;
	v100 = int32(0)
	v102 = F_strlen(m, v20)
	mBase = m.M
	v105 = F_SplitToVariants(m, v17, v100, v100, v20, v102, v100, int32(-1))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L3
	} else {
		goto L25
	}
L25:
	;
	if v105 == int32(0) {
		v329 = v88
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v109 = v105
	v110 = v87
	v111 = v88
	v116 = v93
	goto L27
L27:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if int32(2) <= v120 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v329 = v274
	goto L23
L29:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v123+v120<<(uint(int32(2))%32)-int32(4))))
	v131 = F_NormalizeSubWord(m, v17, v129, int32(8))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	v273 = v110
	v274 = v111
	v278 = v120
	v279 = v116
	goto L31
L31:
	;
	v283 = int32(0)
	if v278 <= v283 {
		goto L68
	} else {
		goto L69
	}
L32:
	;
	if v131 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	if v133 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v261 = v110
	v262 = v111
	v267 = v116
	goto L35
L35:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v273 = v261
	v274 = v262
	v278 = v271
	v279 = v267
	goto L31
L36:
	;
	v134 = v133
	v135 = v110
	v136 = v111
	v139 = v131
	v141 = v116
	goto L39
L37:
	;
	v235 = v110
	v236 = v111
	v241 = v116
	goto L38
L38:
	;
	F_pfree(m, v131)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L3
	} else {
		goto L66
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if int32(0) < v145-int32(1) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v235 = v228
	v236 = v216
	v241 = v230
	goto L38
L41:
	;
	v151 = int32(0)
	v152 = v135
	v153 = v136
	goto L44
L42:
	;
	v198 = v134
	v199 = v135
	v200 = v136
	goto L43
L43:
	;
	if v200 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L44:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v162+v151<<(uint(int32(2))%32))))
	if v131 != v139 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v198 = v197
	v199 = v190
	v200 = v178
	goto L43
L46:
	;
	v168 = F_pstrdup(m, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	v170 = v166
	goto L48
L48:
	;
	if v153 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v170 = v168
	goto L48
L50:
	;
	v175 = F_palloc_mul(m, int32(8), int32(1024))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	v177 = v152
	v178 = v153
	goto L52
L52:
	;
	if v177-v178 <= int32(_a_F_dispell_lexize_0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v177 = v175
	v178 = v175
	goto L52
L54:
	;
	v182 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v177)+2)) = uint16(v182)
	*(*int32)(unsafe.Add(mBase, uint32(v177)+4)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v177)+12)) = v182
	*(*uint16)(unsafe.Add(mBase, uint32(v177))) = uint16(v141)
	v190 = v177 + int32(8)
	goto L56
L55:
	;
	v190 = v177
	goto L56
L56:
	;
	v191 = int32(1)
	v192 = v151 + v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v192 < v193-v191 {
		v151 = v192
		v152 = v190
		v153 = v178
		goto L44
	} else {
		goto L57
	}
L57:
	;
	goto L45
L58:
	;
	v213 = F_palloc_mul(m, int32(8), int32(1024))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L3
	} else {
		goto L61
	}
L59:
	;
	v215 = v199
	v216 = v200
	goto L60
L60:
	;
	if v215-v216 <= int32(_a_F_dispell_lexize_0) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v215 = v213
	v216 = v213
	goto L60
L62:
	;
	v220 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+2)) = uint16(v220)
	*(*int32)(unsafe.Add(mBase, uint32(v215)+4)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v215)+12)) = v220
	*(*uint16)(unsafe.Add(mBase, uint32(v215))) = uint16(v141)
	v228 = v215 + int32(8)
	goto L64
L63:
	;
	v228 = v215
	goto L64
L64:
	;
	v230 = v141 + int32(1)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if v231 != 0 {
		v134 = v231
		v135 = v228
		v136 = v216
		v139 = v139 + int32(4)
		v141 = v230
		goto L39
	} else {
		goto L65
	}
L65:
	;
	goto L40
L66:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = int32(0)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v250+v251<<(uint(int32(2))%32)-int32(4))))
	F_pfree(m, v257)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L3
	} else {
		goto L67
	}
L67:
	;
	v261 = v235
	v262 = v236
	v267 = v241
	goto L35
L68:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	F_pfree(m, v322)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L3
	} else {
		goto L75
	}
L69:
	;
	v286 = v283
	goto L70
L70:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v297+v286<<(uint(int32(2))%32))))
	if v301 == int32(0) {
		goto L68
	} else {
		goto L72
	}
L71:
	;
	goto L68
L72:
	;
	F_pfree(m, v301)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L3
	} else {
		goto L73
	}
L73:
	;
	v307 = v286 + int32(1)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v307 < v308 {
		v286 = v307
		goto L70
	} else {
		goto L74
	}
L74:
	;
	goto L71
L75:
	;
	F_pfree(m, v109)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L3
	} else {
		goto L76
	}
L76:
	;
	if v321 != 0 {
		v109 = v321
		v110 = v273
		v111 = v274
		v116 = v279
		goto L27
	} else {
		goto L77
	}
L77:
	;
	goto L28
L78:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	if v340 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	__phi343 = v329
	__phi345 = v329
	__phi346 = v329 + int32(4)
	__phi348 = v340
	v343 = __phi343
	v345 = __phi345
	v346 = __phi346
	v348 = __phi348
	goto L82
L80:
	;
	v374 = v329
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374)+4)) = int32(0)
	v396 = base.I64_extend_i32_u(v329)
	goto L1
L82:
	;
	v354 = F_searchstoplist(m, v15, v348)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L85
	}
L83:
	;
	v374 = v366
	goto L81
L84:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v343)+12))
	if v369 != 0 {
		__phi343 = v343 + int32(8)
		__phi345 = v366
		__phi346 = v343 + int32(12)
		__phi348 = v369
		v343 = __phi343
		v345 = __phi345
		v346 = __phi346
		v348 = __phi348
		goto L82
	} else {
		goto L93
	}
L85:
	;
	if v354 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	F_pfree(m, v356)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L3
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	if v343 != v345 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = int32(0)
	v366 = v345
	goto L84
L90:
	;
	v362 = *(*int64)(unsafe.Add(mBase, uint32(v343)))
	*(*int64)(unsafe.Add(mBase, uint32(v345))) = v362
	goto L92
L91:
	;
	goto L92
L92:
	;
	v366 = v345 + int32(8)
	goto L84
L93:
	;
	goto L83
}
func F_distance_chebyshev(m *base.Module, l0 int32) int64 {
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
	v2 = float64(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
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
			v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if base.Ui32(v30) < base.Ui32(v33) {
				if v244 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v248 = m.ExcPending
					if v248 != 0 {
						return int64(0)
					} else {
						v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v26 != v249 {
							F_pfree(m, v26)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int64(0)
							} else {
								return base.I64_reinterpret_f64(v230)
							}
						} else {
							return base.I64_reinterpret_f64(v230)
						}
					}
				} else {
					v249 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v26 != v249 {
						F_pfree(m, v26)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int64(0)
						} else {
							return base.I64_reinterpret_f64(v230)
						}
					} else {
						return base.I64_reinterpret_f64(v230)
					}
				}
			} else {
				if v244 != v21 {
					F_pfree(m, v21)
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int64(0)
					} else {
						v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
						if v26 == v254 {
							return base.I64_reinterpret_f64(v230)
						} else {
							F_pfree(m, v26)
							mBase = m.M
							v257 = m.ExcPending
							if v257 != 0 {
								return int64(0)
							} else {
								return base.I64_reinterpret_f64(v230)
							}
						}
					}
				} else {
					v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v26 == v254 {
						return base.I64_reinterpret_f64(v230)
					} else {
						F_pfree(m, v26)
						mBase = m.M
						v257 = m.ExcPending
						if v257 != 0 {
							return int64(0)
						} else {
							return base.I64_reinterpret_f64(v230)
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
func F_dmetaphone(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_DoubleMetaphone(m, v12, v14, v5+int32(8))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				if v19 != 0 {
					v21 = v19
				} else {
					v21 = int32(_a_F_dmetaphone_0)
				}
				v22 = F_cstring_to_text(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					m.G0 = v5 + int32(16)
					return base.I64_extend_i32_u(v22)
				}
			}
		}
	}
}
func F_dmetaphone_alt(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = F_text_to_cstring(m, v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_DoubleMetaphone(m, v12, v14, v5+int32(8))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				if v19 != 0 {
					v21 = v19
				} else {
					v21 = int32(_a_F_dmetaphone_alt_0)
				}
				v22 = F_cstring_to_text(m, v21)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int64(0)
				} else {
					m.G0 = v5 + int32(16)
					return base.I64_extend_i32_u(v22)
				}
			}
		}
	}
}
func F_dpi(m *base.Module, l0 int32) int64 {
	return int64(4614256656552045848)
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
func F_drandom(m *base.Module, l0 int32) int64 {
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
	return base.I64_reinterpret_f64(v156)
}
func F_dsimple_init(m *base.Module, l0 int32) int64 {
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
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
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
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L44
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L1
	} else {
		goto L40
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	m.G0 = v11 + int32(16)
	return base.I64_extend_i32_u(v15)
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
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_dsimple_init_2), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_dsimple_init_3), int32(50), int32(_a_F_dsimple_init_4))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(_a_F_dsimple_init_5), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_dsimple_init_3), int32(59), int32(_a_F_dsimple_init_4))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v164
	F_errmsg(m, int32(_a_F_dsimple_init_6), v11)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_dsimple_init_3), int32(68), int32(_a_F_dsimple_init_4))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
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
func F_dsynonym_lexize(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int64
	_ = v58
	v5 = int64(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v10 <= int32(0) {
		v58 = v5
		m.G0 = v8 + int32(16)
		return v58
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 <= int32(0) {
			v58 = v5
			m.G0 = v8 + int32(16)
			return v58
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
			if v18 == int32(1) {
				v21 = F_pnstrdup(m, v17, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v28 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v28
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v38 = F_bsearch(m, v8+int32(4), v34, v35, int32(12), int32(1263))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						F_pfree(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							if v38 == int32(0) {
								v58 = v5
								m.G0 = v8 + int32(16)
								return v58
							} else {
								v47 = F_palloc0_mul(m, int32(8), int32(2))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v50 = F_pstrdup(m, v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v50
										v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+2)) = uint16(v53)
										v58 = base.I64_extend_i32_u(v47)
										m.G0 = v8 + int32(16)
										return v58
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
					return int64(0)
				} else {
					v28 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v28
					v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v38 = F_bsearch(m, v8+int32(4), v34, v35, int32(12), int32(1263))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int64(0)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						F_pfree(m, v40)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							if v38 == int32(0) {
								v58 = v5
								m.G0 = v8 + int32(16)
								return v58
							} else {
								v47 = F_palloc0_mul(m, int32(8), int32(2))
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int64(0)
								} else {
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
									v50 = F_pstrdup(m, v49)
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v50
										v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+8)))
										*(*uint16)(unsafe.Add(mBase, uint32(v47)+2)) = uint16(v53)
										v58 = base.I64_extend_i32_u(v47)
										m.G0 = v8 + int32(16)
										return v58
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
func F_dtan(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v12 float64
	_ = v12
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
	var v52 int64
	_ = v52
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v4&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _c_F_dtan[0])) = int32(0)
		v12 = base.F64_reinterpret_i64(v4)
		if base.F64_eq(base.F64_abs(v12), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_dtan_0), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_dtan_1), int32(2021), int32(_a_F_dtan_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int64(0)
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
			v27 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v12))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v27) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v27) < base.Ui32(int32(1044381696)) {
					v44 = v12
				} else {
					v34 = F___tan(m, v12, float64(0), int32(0))
					mBase = m.M
					v44 = v34
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v27) {
					v44 = base.F64_sub(v12, v12)
				} else {
					v38 = F___rem_pio2(m, v12, v20)
					mBase = m.M
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
					v43 = F___tan(m, v39, v40, v38&int32(1))
					mBase = m.M
					v44 = v43
				}
			}
			m.G0 = v20 + int32(16)
			v52 = base.I64_reinterpret_f64(v44)
			return v52
		}
	} else {
		v52 = int64(9221120237041090560)
		return v52
	}
}
func F_dtand(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int64
	_ = v14
	var v19 float64
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 float64
	_ = v45
	var v48 int64
	_ = v48
	var v55 float64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v117 int64
	_ = v117
	var v120 int32
	_ = v120
	var v122 int64
	_ = v122
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v162 int64
	_ = v162
	var v170 float64
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 float64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 float64
	_ = v184
	var v187 int32
	_ = v187
	var v188 float64
	_ = v188
	var v192 float64
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v212 float64
	_ = v212
	var v216 int32
	_ = v216
	var v217 float64
	_ = v217
	var v218 float64
	_ = v218
	var v224 float64
	_ = v224
	var v225 float64
	_ = v225
	var v227 float64
	_ = v227
	var v229 float64
	_ = v229
	var v231 float64
	_ = v231
	var v240 float64
	_ = v240
	var v248 float64
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v268 float64
	_ = v268
	var v272 int32
	_ = v272
	var v273 float64
	_ = v273
	var v274 float64
	_ = v274
	var v279 float64
	_ = v279
	var v281 float64
	_ = v281
	var v283 float64
	_ = v283
	var v286 float64
	_ = v286
	var v290 float64
	_ = v290
	var v294 float64
	_ = v294
	var v298 float64
	_ = v298
	var v304 float64
	_ = v304
	var v305 int32
	_ = v305
	var v310 float64
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v323 int32
	_ = v323
	var v330 float64
	_ = v330
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v336 float64
	_ = v336
	var v341 float64
	_ = v341
	var v343 float64
	_ = v343
	var v345 float64
	_ = v345
	var v348 float64
	_ = v348
	var v352 float64
	_ = v352
	var v356 float64
	_ = v356
	var v360 float64
	_ = v360
	var v369 float64
	_ = v369
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v389 float64
	_ = v389
	var v393 int32
	_ = v393
	var v394 float64
	_ = v394
	var v395 float64
	_ = v395
	var v401 float64
	_ = v401
	var v402 float64
	_ = v402
	var v404 float64
	_ = v404
	var v406 float64
	_ = v406
	var v408 float64
	_ = v408
	var v417 float64
	_ = v417
	var v421 float64
	_ = v421
	var v422 float64
	_ = v422
	var v424 float64
	_ = v424
	var v427 float64
	_ = v427
	var v430 float64
	_ = v430
	var v433 float64
	_ = v433
	var v439 int64
	_ = v439
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if base.Ui64(v14&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
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
		goto L128
	}
L2:
	;
	v19 = base.F64_reinterpret_i64(v14)
	if base.F64_eq(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v439 = int64(9221120237041090560)
	goto L4
L4:
	;
	m.G0 = v11 + int32(16)
	return v439
L5:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_dtand[0])))
	if v24 == int32(0) {
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
	v28 = int32(0)
	v36 = base.I64_reinterpret_f64(v19)
	v40 = int32(2047)
	v41 = base.I32_wrap_i64(int64(base.Ui64(v36)>>(uint(int64(52))%64))) & v40
	if v41 == v40 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v172 = base.F64_lt(v170, float64(0))
	if v172 != 0 {
		goto L50
	} else {
		goto L51
	}
L10:
	;
	v45 = base.F64_mul(v19, float64(360))
	v170 = base.F64_div(v45, v45)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v48 = v36 << (uint(int64(1)) % 64)
	if base.Ui64(v48) <= base.Ui64(int64(-9156662467374350336)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v48 == int64(-9156662467374350336) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	if v41 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v55 = base.F64_mul(v19, float64(0))
	goto L18
L17:
	;
	v55 = v19
	goto L18
L18:
	;
	v170 = v55
	goto L9
L19:
	;
	if int32(1031) < v91 {
		goto L29
	} else {
		goto L30
	}
L20:
	;
	v58 = int32(0)
	v60 = v36 << (uint(int64(12)) % 64)
	if int64(0) <= v60 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v91 = v41
	v92 = v36&int64(4503599627370495) | int64(4503599627370496)
	goto L19
L23:
	;
	v64 = v60
	v67 = v58
	goto L26
L24:
	;
	v78 = v58
	goto L25
L25:
	;
	v91 = v78
	v92 = v36 << (uint(base.I64_extend_i32_u(int32(1)-v78)) % 64)
	goto L19
L26:
	;
	v69 = v67 - int32(1)
	v71 = v64 << (uint(int64(1)) % 64)
	if int64(0) <= v71 {
		v64 = v71
		v67 = v69
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v78 = v69
	goto L25
L28:
	;
	goto L27
L29:
	;
	v96 = v92
	v99 = v91
	goto L32
L30:
	;
	v117 = v92
	v120 = v91
	goto L31
L31:
	;
	v122 = v117 - int64(6333186975989760)
	if v122 < int64(0) {
		v129 = v117
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v101 = v96 - int64(6333186975989760)
	if v101 < int64(0) {
		v108 = v96
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v117 = v110
	v120 = int32(1031)
	goto L31
L34:
	;
	v110 = v108 << (uint(int64(1)) % 64)
	v112 = v99 - int32(1)
	if int32(1031) < v112 {
		v96 = v110
		v99 = v112
		goto L32
	} else {
		goto L37
	}
L35:
	;
	if v101 != int64(0) {
		v108 = v101
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v170 = base.F64_mul(v19, float64(0))
	goto L9
L37:
	;
	goto L33
L38:
	;
	if base.Ui64(v129) <= base.Ui64(int64(4503599627370495)) {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	if v122 != int64(0) {
		v129 = v122
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v170 = base.F64_mul(v19, float64(0))
	goto L9
L41:
	;
	v133 = v129
	v136 = v120
	goto L44
L42:
	;
	v144 = v129
	v147 = v120
	goto L43
L43:
	;
	if int32(0) < v147 {
		goto L47
	} else {
		goto L48
	}
L44:
	;
	v138 = v136 - int32(1)
	v140 = v133 << (uint(int64(1)) % 64)
	if base.Ui64(v133) < base.Ui64(int64(2251799813685248)) {
		v133 = v140
		v136 = v138
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v144 = v140
	v147 = v138
	goto L43
L46:
	;
	goto L45
L47:
	;
	v162 = v144 - int64(4503599627370496) | base.I64_extend_i32_u(v147)<<(uint(int64(52))%64)
	goto L49
L48:
	;
	v162 = int64(base.Ui64(v144) >> (uint(base.I64_extend_i32_u(int32(1)-v147)) % 64))
	goto L49
L49:
	;
	v170 = base.F64_reinterpret_i64(v36&int64(-9223372036854775807-1) | v162)
	goto L9
L50:
	;
	v173 = int32(-1)
	goto L52
L51:
	;
	v173 = int32(1)
	goto L52
L52:
	;
	if v172 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v176 = base.F64_neg(v170)
	goto L55
L54:
	;
	v176 = v170
	goto L55
L55:
	;
	v178 = base.F64_gt(v176, float64(180))
	if v178 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v179 = v28 - v173
	goto L58
L57:
	;
	v179 = v173
	goto L58
L58:
	;
	if v178 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v187 != 0 {
		goto L93
	} else {
		goto L94
	}
L60:
	;
	v184 = base.F64_sub(float64(360), v176)
	goto L62
L61:
	;
	v184 = v176
	goto L62
L62:
	;
	v187 = base.F64_gt(v184, float64(90))
	if v187 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v188 = base.F64_sub(float64(180), v184)
	goto L65
L64:
	;
	v188 = v184
	goto L65
L65:
	;
	if base.F64_le(v188, float64(30)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = base.F64_mul(v188, float64(0.017453292519943295))
	v196 = m.G0
	v198 = v196 - int32(16)
	m.G0 = v198
	v205 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v192))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v205) <= base.Ui32(int32(1072243195)) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	v248 = base.F64_mul(base.F64_sub(float64(90), v188), float64(0.017453292519943295))
	v252 = m.G0
	v254 = v252 - int32(16)
	m.G0 = v254
	v261 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v248))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v261) <= base.Ui32(int32(1072243195)) {
		goto L84
	} else {
		goto L85
	}
L69:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v231
	v240 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[1]))
	v304 = base.F64_mul(base.F64_div(v231, v240), float64(0.5))
	goto L59
L70:
	;
	m.G0 = v198 + int32(16)
	goto L69
L71:
	;
	if base.Ui32(v205) < base.Ui32(int32(1045430272)) {
		v231 = v192
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v205) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v212 = F___sin(m, v192, float64(0), int32(0))
	mBase = m.M
	v231 = v212
	goto L70
L75:
	;
	v231 = base.F64_sub(v192, v192)
	goto L70
L76:
	;
	goto L77
L77:
	;
	v216 = F___rem_pio2(m, v192, v198)
	mBase = m.M
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v198)+8))
	v218 = *(*float64)(unsafe.Add(mBase, uint32(v198)))
	switch v216&int32(3) - int32(1) {
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
	v229 = F___cos(m, v218, v217)
	mBase = m.M
	v231 = base.F64_neg(v229)
	goto L70
L79:
	;
	v227 = F___sin(m, v218, v217, int32(1))
	mBase = m.M
	v231 = base.F64_neg(v227)
	goto L70
L80:
	;
	v225 = F___cos(m, v218, v217)
	mBase = m.M
	v231 = v225
	goto L70
L81:
	;
	v224 = F___sin(m, v218, v217, int32(1))
	mBase = m.M
	v231 = v224
	goto L70
L82:
	;
	v294 = base.F64_sub(float64(1), v290)
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v294
	v298 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[2]))
	v304 = base.F64_add(base.F64_mul(base.F64_div(v294, v298), float64(-0.5)), float64(1))
	goto L59
L83:
	;
	m.G0 = v254 + int32(16)
	goto L82
L84:
	;
	if base.Ui32(v261) < base.Ui32(int32(1044816030)) {
		v290 = float64(1)
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v261) {
		v290 = base.F64_sub(v248, v248)
		goto L83
	} else {
		goto L88
	}
L87:
	;
	v268 = F___cos(m, v248, float64(0))
	mBase = m.M
	v290 = v268
	goto L83
L88:
	;
	v272 = F___rem_pio2(m, v248, v254)
	mBase = m.M
	v273 = *(*float64)(unsafe.Add(mBase, uint32(v254)+8))
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v254)))
	switch v272&int32(3) - int32(1) {
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
	v286 = F___sin(m, v274, v273, int32(1))
	mBase = m.M
	v290 = v286
	goto L83
L90:
	;
	v283 = F___cos(m, v274, v273)
	mBase = m.M
	v290 = base.F64_neg(v283)
	goto L83
L91:
	;
	v281 = F___sin(m, v274, v273, int32(1))
	mBase = m.M
	v290 = base.F64_neg(v281)
	goto L83
L92:
	;
	v279 = F___cos(m, v274, v273)
	mBase = m.M
	v290 = v279
	goto L83
L93:
	;
	v305 = v28 - v179
	goto L95
L94:
	;
	v305 = v179
	goto L95
L95:
	;
	if base.F64_le(v188, float64(60)) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v422 = base.F64_div(v304, v421)
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v422
	v424 = float64(0)
	v427 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[3]))
	v430 = base.F64_mul(base.F64_div(v422, v427), base.F64_convert_i32_s(v305))
	if base.F64_eq(v430, v424) != 0 {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	v310 = base.F64_mul(v188, float64(0.017453292519943295))
	v314 = m.G0
	v316 = v314 - int32(16)
	m.G0 = v316
	v323 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v310))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v323) <= base.Ui32(int32(1072243195)) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	goto L99
L99:
	;
	v369 = base.F64_mul(base.F64_sub(float64(90), v188), float64(0.017453292519943295))
	v373 = m.G0
	v375 = v373 - int32(16)
	m.G0 = v375
	v382 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v369))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v382) <= base.Ui32(int32(1072243195)) {
		goto L113
	} else {
		goto L114
	}
L100:
	;
	v356 = base.F64_sub(float64(1), v352)
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v356
	v360 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[2]))
	v421 = base.F64_add(base.F64_mul(base.F64_div(v356, v360), float64(-0.5)), float64(1))
	goto L96
L101:
	;
	m.G0 = v316 + int32(16)
	goto L100
L102:
	;
	if base.Ui32(v323) < base.Ui32(int32(1044816030)) {
		v352 = float64(1)
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v323) {
		v352 = base.F64_sub(v310, v310)
		goto L101
	} else {
		goto L106
	}
L105:
	;
	v330 = F___cos(m, v310, float64(0))
	mBase = m.M
	v352 = v330
	goto L101
L106:
	;
	v334 = F___rem_pio2(m, v310, v316)
	mBase = m.M
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v316)+8))
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v316)))
	switch v334&int32(3) - int32(1) {
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
	v348 = F___sin(m, v336, v335, int32(1))
	mBase = m.M
	v352 = v348
	goto L101
L108:
	;
	v345 = F___cos(m, v336, v335)
	mBase = m.M
	v352 = base.F64_neg(v345)
	goto L101
L109:
	;
	v343 = F___sin(m, v336, v335, int32(1))
	mBase = m.M
	v352 = base.F64_neg(v343)
	goto L101
L110:
	;
	v341 = F___cos(m, v336, v335)
	mBase = m.M
	v352 = v341
	goto L101
L111:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v11)+8)) = v408
	v417 = *(*float64)(unsafe.Add(mBase, _c_F_dtand[1]))
	v421 = base.F64_mul(base.F64_div(v408, v417), float64(0.5))
	goto L96
L112:
	;
	m.G0 = v375 + int32(16)
	goto L111
L113:
	;
	if base.Ui32(v382) < base.Ui32(int32(1045430272)) {
		v408 = v369
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v382) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v389 = F___sin(m, v369, float64(0), int32(0))
	mBase = m.M
	v408 = v389
	goto L112
L117:
	;
	v408 = base.F64_sub(v369, v369)
	goto L112
L118:
	;
	goto L119
L119:
	;
	v393 = F___rem_pio2(m, v369, v375)
	mBase = m.M
	v394 = *(*float64)(unsafe.Add(mBase, uint32(v375)+8))
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v375)))
	switch v393&int32(3) - int32(1) {
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
	v406 = F___cos(m, v395, v394)
	mBase = m.M
	v408 = base.F64_neg(v406)
	goto L112
L121:
	;
	v404 = F___sin(m, v395, v394, int32(1))
	mBase = m.M
	v408 = base.F64_neg(v404)
	goto L112
L122:
	;
	v402 = F___cos(m, v395, v394)
	mBase = m.M
	v408 = v402
	goto L112
L123:
	;
	v401 = F___sin(m, v395, v394, int32(1))
	mBase = m.M
	v408 = v401
	goto L112
L124:
	;
	v433 = v424
	goto L126
L125:
	;
	v433 = v430
	goto L126
L126:
	;
	v439 = base.I64_reinterpret_f64(v433)
	goto L4
L127:
	;
	return int64(0)
L128:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	F_errmsg(m, int32(_a_F_dtand_0), int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_dtand_1), int32(2553), int32(_a_F_dtand_2))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dtanh(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 float64
	_ = v5
	var v6 float64
	_ = v6
	var v7 int64
	_ = v7
	var v18 float64
	_ = v18
	var v25 int64
	_ = v25
	var v30 int32
	_ = v30
	var v66 int32
	_ = v66
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v83 float64
	_ = v83
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v105 float64
	_ = v105
	var v108 float64
	_ = v108
	var v114 float64
	_ = v114
	var v123 float64
	_ = v123
	var v138 float64
	_ = v138
	var v147 float64
	_ = v147
	var v152 float64
	_ = v152
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v168 float64
	_ = v168
	var v178 float64
	_ = v178
	var v180 float64
	_ = v180
	var v192 float64
	_ = v192
	var v199 float64
	_ = v199
	var v206 int64
	_ = v206
	var v211 int32
	_ = v211
	var v247 int32
	_ = v247
	var v248 float64
	_ = v248
	var v254 int32
	_ = v254
	var v255 float64
	_ = v255
	var v257 float64
	_ = v257
	var v258 float64
	_ = v258
	var v264 float64
	_ = v264
	var v265 int32
	_ = v265
	var v266 float64
	_ = v266
	var v269 float64
	_ = v269
	var v270 float64
	_ = v270
	var v286 float64
	_ = v286
	var v289 float64
	_ = v289
	var v295 float64
	_ = v295
	var v304 float64
	_ = v304
	var v319 float64
	_ = v319
	var v328 float64
	_ = v328
	var v333 float64
	_ = v333
	var v340 float64
	_ = v340
	var v343 float64
	_ = v343
	var v349 float64
	_ = v349
	var v359 float64
	_ = v359
	var v361 float64
	_ = v361
	var v373 float64
	_ = v373
	var v380 float64
	_ = v380
	var v387 int64
	_ = v387
	var v392 int32
	_ = v392
	var v428 int32
	_ = v428
	var v429 float64
	_ = v429
	var v435 int32
	_ = v435
	var v436 float64
	_ = v436
	var v438 float64
	_ = v438
	var v439 float64
	_ = v439
	var v445 float64
	_ = v445
	var v446 int32
	_ = v446
	var v447 float64
	_ = v447
	var v450 float64
	_ = v450
	var v451 float64
	_ = v451
	var v467 float64
	_ = v467
	var v470 float64
	_ = v470
	var v476 float64
	_ = v476
	var v485 float64
	_ = v485
	var v500 float64
	_ = v500
	var v509 float64
	_ = v509
	var v514 float64
	_ = v514
	var v521 float64
	_ = v521
	var v524 float64
	_ = v524
	var v530 float64
	_ = v530
	var v540 float64
	_ = v540
	var v542 float64
	_ = v542
	var v554 float64
	_ = v554
	var v559 float64
	_ = v559
	var v564 float64
	_ = v564
	var v571 int32
	_ = v571
	v5 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = base.F64_abs(v5)
	v7 = base.I64_reinterpret_f64(v6)
	if base.Ui64(int64(4603122931675955200)) <= base.Ui64(v7) {
		if base.Ui64(int64(4626322721511309312)) <= base.Ui64(v7) {
			v559 = base.F64_add(base.F64_div(math.Float64frombits(uint64(0x8000000000000000)), v6), float64(1))
		} else {
			v18 = base.F64_add(v6, v6)
			v25 = base.I64_reinterpret_f64(v18)
			v30 = base.I32_wrap_i64(int64(base.Ui64(v25)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v30) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v18)&int64(9223372036854775807)) {
					v180 = v18
					v192 = v180
				} else {
					if v25 < int64(0) {
						v192 = float64(-1)
					} else {
						if base.F64_gt(v18, float64(709.782712893384)) == int32(0) {
							v66 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v18, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v18)))
							v67 = base.F64_convert_i32_s(v66)
							v73 = v66
							v74 = base.F64_mul(v67, float64(1.9082149292705877e-10))
							v76 = base.F64_add(v18, base.F64_mul(v67, float64(-0.6931471803691238)))
							v77 = base.F64_sub(v76, v74)
							v83 = v77
							v84 = v73
							v85 = base.F64_sub(base.F64_sub(v76, v77), v74)
							v88 = base.F64_mul(v83, float64(0.5))
							v89 = base.F64_mul(v83, v88)
							v105 = base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v108 = base.F64_sub(float64(3), base.F64_mul(v105, v88))
							v114 = base.F64_mul(v89, base.F64_div(base.F64_sub(v105, v108), base.F64_sub(float64(6), base.F64_mul(v83, v108))))
							if v84 == int32(0) {
								v192 = base.F64_sub(v83, base.F64_sub(base.F64_mul(v83, v114), v89))
							} else {
								v123 = base.F64_sub(base.F64_sub(base.F64_mul(v83, base.F64_sub(v114, v85)), v85), v89)
								switch v84 + int32(1) {
								case 0:
									v192 = base.F64_add(base.F64_mul(base.F64_sub(v83, v123), float64(0.5)), float64(-0.5))
								default:
									v147 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v84+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v84) {
										v152 = base.F64_add(base.F64_sub(v83, v123), float64(1))
										if v84 == int32(1024) {
											v159 = base.F64_mul(base.F64_add(v152, v152), float64(8.98846567431158e+307))
										} else {
											v159 = base.F64_mul(v152, v147)
										}
										v192 = base.F64_add(v159, float64(-1))
									} else {
										v162 = float64(1)
										v168 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v84) << (uint(int64(52)) % 64))
										if base.Ui32(v84) <= base.Ui32(int32(19)) {
											v178 = base.F64_add(base.F64_sub(v162, v168), base.F64_sub(v83, v123))
										} else {
											v178 = base.F64_add(base.F64_sub(v83, base.F64_add(v123, v168)), v162)
										}
										v180 = base.F64_mul(v178, v147)
										v192 = v180
									}
								case 2:
									if base.F64_lt(v83, float64(-0.25)) != 0 {
										v192 = base.F64_mul(base.F64_sub(v123, base.F64_add(v83, float64(0.5))), float64(-2))
									} else {
										v138 = base.F64_sub(v83, v123)
										v192 = base.F64_add(base.F64_add(v138, v138), float64(1))
									}
								}
							}
						} else {
							v192 = base.F64_mul(v18, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v30) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v30) < base.Ui32(int32(1016070144)) {
						v180 = v18
						v192 = v180
					} else {
						v83 = v18
						v84 = int32(0)
						v85 = float64(0)
						v88 = base.F64_mul(v83, float64(0.5))
						v89 = base.F64_mul(v83, v88)
						v105 = base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v108 = base.F64_sub(float64(3), base.F64_mul(v105, v88))
						v114 = base.F64_mul(v89, base.F64_div(base.F64_sub(v105, v108), base.F64_sub(float64(6), base.F64_mul(v83, v108))))
						if v84 == int32(0) {
							v192 = base.F64_sub(v83, base.F64_sub(base.F64_mul(v83, v114), v89))
						} else {
							v123 = base.F64_sub(base.F64_sub(base.F64_mul(v83, base.F64_sub(v114, v85)), v85), v89)
							switch v84 + int32(1) {
							case 0:
								v192 = base.F64_add(base.F64_mul(base.F64_sub(v83, v123), float64(0.5)), float64(-0.5))
							default:
								v147 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v84+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v84) {
									v152 = base.F64_add(base.F64_sub(v83, v123), float64(1))
									if v84 == int32(1024) {
										v159 = base.F64_mul(base.F64_add(v152, v152), float64(8.98846567431158e+307))
									} else {
										v159 = base.F64_mul(v152, v147)
									}
									v192 = base.F64_add(v159, float64(-1))
								} else {
									v162 = float64(1)
									v168 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v84) << (uint(int64(52)) % 64))
									if base.Ui32(v84) <= base.Ui32(int32(19)) {
										v178 = base.F64_add(base.F64_sub(v162, v168), base.F64_sub(v83, v123))
									} else {
										v178 = base.F64_add(base.F64_sub(v83, base.F64_add(v123, v168)), v162)
									}
									v180 = base.F64_mul(v178, v147)
									v192 = v180
								}
							case 2:
								if base.F64_lt(v83, float64(-0.25)) != 0 {
									v192 = base.F64_mul(base.F64_sub(v123, base.F64_add(v83, float64(0.5))), float64(-2))
								} else {
									v138 = base.F64_sub(v83, v123)
									v192 = base.F64_add(base.F64_add(v138, v138), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v30) {
						v66 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v18, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v18)))
						v67 = base.F64_convert_i32_s(v66)
						v73 = v66
						v74 = base.F64_mul(v67, float64(1.9082149292705877e-10))
						v76 = base.F64_add(v18, base.F64_mul(v67, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v25 {
							v73 = int32(1)
							v74 = float64(1.9082149292705877e-10)
							v76 = base.F64_add(v18, float64(-0.6931471803691238))
						} else {
							v73 = int32(-1)
							v74 = float64(-1.9082149292705877e-10)
							v76 = base.F64_add(v18, float64(0.6931471803691238))
						}
					}
					v77 = base.F64_sub(v76, v74)
					v83 = v77
					v84 = v73
					v85 = base.F64_sub(base.F64_sub(v76, v77), v74)
					v88 = base.F64_mul(v83, float64(0.5))
					v89 = base.F64_mul(v83, v88)
					v105 = base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, base.F64_add(base.F64_mul(v89, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v108 = base.F64_sub(float64(3), base.F64_mul(v105, v88))
					v114 = base.F64_mul(v89, base.F64_div(base.F64_sub(v105, v108), base.F64_sub(float64(6), base.F64_mul(v83, v108))))
					if v84 == int32(0) {
						v192 = base.F64_sub(v83, base.F64_sub(base.F64_mul(v83, v114), v89))
					} else {
						v123 = base.F64_sub(base.F64_sub(base.F64_mul(v83, base.F64_sub(v114, v85)), v85), v89)
						switch v84 + int32(1) {
						case 0:
							v192 = base.F64_add(base.F64_mul(base.F64_sub(v83, v123), float64(0.5)), float64(-0.5))
						default:
							v147 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v84+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v84) {
								v152 = base.F64_add(base.F64_sub(v83, v123), float64(1))
								if v84 == int32(1024) {
									v159 = base.F64_mul(base.F64_add(v152, v152), float64(8.98846567431158e+307))
								} else {
									v159 = base.F64_mul(v152, v147)
								}
								v192 = base.F64_add(v159, float64(-1))
							} else {
								v162 = float64(1)
								v168 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v84) << (uint(int64(52)) % 64))
								if base.Ui32(v84) <= base.Ui32(int32(19)) {
									v178 = base.F64_add(base.F64_sub(v162, v168), base.F64_sub(v83, v123))
								} else {
									v178 = base.F64_add(base.F64_sub(v83, base.F64_add(v123, v168)), v162)
								}
								v180 = base.F64_mul(v178, v147)
								v192 = v180
							}
						case 2:
							if base.F64_lt(v83, float64(-0.25)) != 0 {
								v192 = base.F64_mul(base.F64_sub(v123, base.F64_add(v83, float64(0.5))), float64(-2))
							} else {
								v138 = base.F64_sub(v83, v123)
								v192 = base.F64_add(base.F64_add(v138, v138), float64(1))
							}
						}
					}
				}
			}
			v559 = base.F64_sub(float64(1), base.F64_div(float64(2), base.F64_add(v192, float64(2))))
		}
	} else {
		if base.Ui64(int64(4598272728187797504)) <= base.Ui64(v7) {
			v199 = base.F64_add(v6, v6)
			v206 = base.I64_reinterpret_f64(v199)
			v211 = base.I32_wrap_i64(int64(base.Ui64(v206)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v211) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v199)&int64(9223372036854775807)) {
					v361 = v199
					v373 = v361
				} else {
					if v206 < int64(0) {
						v373 = float64(-1)
					} else {
						if base.F64_gt(v199, float64(709.782712893384)) == int32(0) {
							v247 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v199, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v199)))
							v248 = base.F64_convert_i32_s(v247)
							v254 = v247
							v255 = base.F64_mul(v248, float64(1.9082149292705877e-10))
							v257 = base.F64_add(v199, base.F64_mul(v248, float64(-0.6931471803691238)))
							v258 = base.F64_sub(v257, v255)
							v264 = v258
							v265 = v254
							v266 = base.F64_sub(base.F64_sub(v257, v258), v255)
							v269 = base.F64_mul(v264, float64(0.5))
							v270 = base.F64_mul(v264, v269)
							v286 = base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v289 = base.F64_sub(float64(3), base.F64_mul(v286, v269))
							v295 = base.F64_mul(v270, base.F64_div(base.F64_sub(v286, v289), base.F64_sub(float64(6), base.F64_mul(v264, v289))))
							if v265 == int32(0) {
								v373 = base.F64_sub(v264, base.F64_sub(base.F64_mul(v264, v295), v270))
							} else {
								v304 = base.F64_sub(base.F64_sub(base.F64_mul(v264, base.F64_sub(v295, v266)), v266), v270)
								switch v265 + int32(1) {
								case 0:
									v373 = base.F64_add(base.F64_mul(base.F64_sub(v264, v304), float64(0.5)), float64(-0.5))
								default:
									v328 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v265+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v265) {
										v333 = base.F64_add(base.F64_sub(v264, v304), float64(1))
										if v265 == int32(1024) {
											v340 = base.F64_mul(base.F64_add(v333, v333), float64(8.98846567431158e+307))
										} else {
											v340 = base.F64_mul(v333, v328)
										}
										v373 = base.F64_add(v340, float64(-1))
									} else {
										v343 = float64(1)
										v349 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v265) << (uint(int64(52)) % 64))
										if base.Ui32(v265) <= base.Ui32(int32(19)) {
											v359 = base.F64_add(base.F64_sub(v343, v349), base.F64_sub(v264, v304))
										} else {
											v359 = base.F64_add(base.F64_sub(v264, base.F64_add(v304, v349)), v343)
										}
										v361 = base.F64_mul(v359, v328)
										v373 = v361
									}
								case 2:
									if base.F64_lt(v264, float64(-0.25)) != 0 {
										v373 = base.F64_mul(base.F64_sub(v304, base.F64_add(v264, float64(0.5))), float64(-2))
									} else {
										v319 = base.F64_sub(v264, v304)
										v373 = base.F64_add(base.F64_add(v319, v319), float64(1))
									}
								}
							}
						} else {
							v373 = base.F64_mul(v199, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v211) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v211) < base.Ui32(int32(1016070144)) {
						v361 = v199
						v373 = v361
					} else {
						v264 = v199
						v265 = int32(0)
						v266 = float64(0)
						v269 = base.F64_mul(v264, float64(0.5))
						v270 = base.F64_mul(v264, v269)
						v286 = base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v289 = base.F64_sub(float64(3), base.F64_mul(v286, v269))
						v295 = base.F64_mul(v270, base.F64_div(base.F64_sub(v286, v289), base.F64_sub(float64(6), base.F64_mul(v264, v289))))
						if v265 == int32(0) {
							v373 = base.F64_sub(v264, base.F64_sub(base.F64_mul(v264, v295), v270))
						} else {
							v304 = base.F64_sub(base.F64_sub(base.F64_mul(v264, base.F64_sub(v295, v266)), v266), v270)
							switch v265 + int32(1) {
							case 0:
								v373 = base.F64_add(base.F64_mul(base.F64_sub(v264, v304), float64(0.5)), float64(-0.5))
							default:
								v328 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v265+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v265) {
									v333 = base.F64_add(base.F64_sub(v264, v304), float64(1))
									if v265 == int32(1024) {
										v340 = base.F64_mul(base.F64_add(v333, v333), float64(8.98846567431158e+307))
									} else {
										v340 = base.F64_mul(v333, v328)
									}
									v373 = base.F64_add(v340, float64(-1))
								} else {
									v343 = float64(1)
									v349 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v265) << (uint(int64(52)) % 64))
									if base.Ui32(v265) <= base.Ui32(int32(19)) {
										v359 = base.F64_add(base.F64_sub(v343, v349), base.F64_sub(v264, v304))
									} else {
										v359 = base.F64_add(base.F64_sub(v264, base.F64_add(v304, v349)), v343)
									}
									v361 = base.F64_mul(v359, v328)
									v373 = v361
								}
							case 2:
								if base.F64_lt(v264, float64(-0.25)) != 0 {
									v373 = base.F64_mul(base.F64_sub(v304, base.F64_add(v264, float64(0.5))), float64(-2))
								} else {
									v319 = base.F64_sub(v264, v304)
									v373 = base.F64_add(base.F64_add(v319, v319), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v211) {
						v247 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v199, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v199)))
						v248 = base.F64_convert_i32_s(v247)
						v254 = v247
						v255 = base.F64_mul(v248, float64(1.9082149292705877e-10))
						v257 = base.F64_add(v199, base.F64_mul(v248, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v206 {
							v254 = int32(1)
							v255 = float64(1.9082149292705877e-10)
							v257 = base.F64_add(v199, float64(-0.6931471803691238))
						} else {
							v254 = int32(-1)
							v255 = float64(-1.9082149292705877e-10)
							v257 = base.F64_add(v199, float64(0.6931471803691238))
						}
					}
					v258 = base.F64_sub(v257, v255)
					v264 = v258
					v265 = v254
					v266 = base.F64_sub(base.F64_sub(v257, v258), v255)
					v269 = base.F64_mul(v264, float64(0.5))
					v270 = base.F64_mul(v264, v269)
					v286 = base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, base.F64_add(base.F64_mul(v270, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v289 = base.F64_sub(float64(3), base.F64_mul(v286, v269))
					v295 = base.F64_mul(v270, base.F64_div(base.F64_sub(v286, v289), base.F64_sub(float64(6), base.F64_mul(v264, v289))))
					if v265 == int32(0) {
						v373 = base.F64_sub(v264, base.F64_sub(base.F64_mul(v264, v295), v270))
					} else {
						v304 = base.F64_sub(base.F64_sub(base.F64_mul(v264, base.F64_sub(v295, v266)), v266), v270)
						switch v265 + int32(1) {
						case 0:
							v373 = base.F64_add(base.F64_mul(base.F64_sub(v264, v304), float64(0.5)), float64(-0.5))
						default:
							v328 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v265+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v265) {
								v333 = base.F64_add(base.F64_sub(v264, v304), float64(1))
								if v265 == int32(1024) {
									v340 = base.F64_mul(base.F64_add(v333, v333), float64(8.98846567431158e+307))
								} else {
									v340 = base.F64_mul(v333, v328)
								}
								v373 = base.F64_add(v340, float64(-1))
							} else {
								v343 = float64(1)
								v349 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v265) << (uint(int64(52)) % 64))
								if base.Ui32(v265) <= base.Ui32(int32(19)) {
									v359 = base.F64_add(base.F64_sub(v343, v349), base.F64_sub(v264, v304))
								} else {
									v359 = base.F64_add(base.F64_sub(v264, base.F64_add(v304, v349)), v343)
								}
								v361 = base.F64_mul(v359, v328)
								v373 = v361
							}
						case 2:
							if base.F64_lt(v264, float64(-0.25)) != 0 {
								v373 = base.F64_mul(base.F64_sub(v304, base.F64_add(v264, float64(0.5))), float64(-2))
							} else {
								v319 = base.F64_sub(v264, v304)
								v373 = base.F64_add(base.F64_add(v319, v319), float64(1))
							}
						}
					}
				}
			}
			v559 = base.F64_div(v373, base.F64_add(v373, float64(2)))
		} else {
			if base.Ui64(v7) < base.Ui64(int64(4503599627370496)) {
				v559 = v6
			} else {
				v380 = base.F64_mul(v6, float64(-2))
				v387 = base.I64_reinterpret_f64(v380)
				v392 = base.I32_wrap_i64(int64(base.Ui64(v387)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1078159482)) <= base.Ui32(v392) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v380)&int64(9223372036854775807)) {
						v542 = v380
						v554 = v542
					} else {
						if v387 < int64(0) {
							v554 = float64(-1)
						} else {
							if base.F64_gt(v380, float64(709.782712893384)) == int32(0) {
								v428 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v380, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v380)))
								v429 = base.F64_convert_i32_s(v428)
								v435 = v428
								v436 = base.F64_mul(v429, float64(1.9082149292705877e-10))
								v438 = base.F64_add(v380, base.F64_mul(v429, float64(-0.6931471803691238)))
								v439 = base.F64_sub(v438, v436)
								v445 = v439
								v446 = v435
								v447 = base.F64_sub(base.F64_sub(v438, v439), v436)
								v450 = base.F64_mul(v445, float64(0.5))
								v451 = base.F64_mul(v445, v450)
								v467 = base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
								v470 = base.F64_sub(float64(3), base.F64_mul(v467, v450))
								v476 = base.F64_mul(v451, base.F64_div(base.F64_sub(v467, v470), base.F64_sub(float64(6), base.F64_mul(v445, v470))))
								if v446 == int32(0) {
									v554 = base.F64_sub(v445, base.F64_sub(base.F64_mul(v445, v476), v451))
								} else {
									v485 = base.F64_sub(base.F64_sub(base.F64_mul(v445, base.F64_sub(v476, v447)), v447), v451)
									switch v446 + int32(1) {
									case 0:
										v554 = base.F64_add(base.F64_mul(base.F64_sub(v445, v485), float64(0.5)), float64(-0.5))
									default:
										v509 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v446+int32(1023)) << (uint(int64(52)) % 64))
										if base.Ui32(int32(57)) <= base.Ui32(v446) {
											v514 = base.F64_add(base.F64_sub(v445, v485), float64(1))
											if v446 == int32(1024) {
												v521 = base.F64_mul(base.F64_add(v514, v514), float64(8.98846567431158e+307))
											} else {
												v521 = base.F64_mul(v514, v509)
											}
											v554 = base.F64_add(v521, float64(-1))
										} else {
											v524 = float64(1)
											v530 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v446) << (uint(int64(52)) % 64))
											if base.Ui32(v446) <= base.Ui32(int32(19)) {
												v540 = base.F64_add(base.F64_sub(v524, v530), base.F64_sub(v445, v485))
											} else {
												v540 = base.F64_add(base.F64_sub(v445, base.F64_add(v485, v530)), v524)
											}
											v542 = base.F64_mul(v540, v509)
											v554 = v542
										}
									case 2:
										if base.F64_lt(v445, float64(-0.25)) != 0 {
											v554 = base.F64_mul(base.F64_sub(v485, base.F64_add(v445, float64(0.5))), float64(-2))
										} else {
											v500 = base.F64_sub(v445, v485)
											v554 = base.F64_add(base.F64_add(v500, v500), float64(1))
										}
									}
								}
							} else {
								v554 = base.F64_mul(v380, float64(8.98846567431158e+307))
							}
						}
					}
				} else {
					if base.Ui32(v392) < base.Ui32(int32(1071001155)) {
						if base.Ui32(v392) < base.Ui32(int32(1016070144)) {
							v542 = v380
							v554 = v542
						} else {
							v445 = v380
							v446 = int32(0)
							v447 = float64(0)
							v450 = base.F64_mul(v445, float64(0.5))
							v451 = base.F64_mul(v445, v450)
							v467 = base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v470 = base.F64_sub(float64(3), base.F64_mul(v467, v450))
							v476 = base.F64_mul(v451, base.F64_div(base.F64_sub(v467, v470), base.F64_sub(float64(6), base.F64_mul(v445, v470))))
							if v446 == int32(0) {
								v554 = base.F64_sub(v445, base.F64_sub(base.F64_mul(v445, v476), v451))
							} else {
								v485 = base.F64_sub(base.F64_sub(base.F64_mul(v445, base.F64_sub(v476, v447)), v447), v451)
								switch v446 + int32(1) {
								case 0:
									v554 = base.F64_add(base.F64_mul(base.F64_sub(v445, v485), float64(0.5)), float64(-0.5))
								default:
									v509 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v446+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v446) {
										v514 = base.F64_add(base.F64_sub(v445, v485), float64(1))
										if v446 == int32(1024) {
											v521 = base.F64_mul(base.F64_add(v514, v514), float64(8.98846567431158e+307))
										} else {
											v521 = base.F64_mul(v514, v509)
										}
										v554 = base.F64_add(v521, float64(-1))
									} else {
										v524 = float64(1)
										v530 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v446) << (uint(int64(52)) % 64))
										if base.Ui32(v446) <= base.Ui32(int32(19)) {
											v540 = base.F64_add(base.F64_sub(v524, v530), base.F64_sub(v445, v485))
										} else {
											v540 = base.F64_add(base.F64_sub(v445, base.F64_add(v485, v530)), v524)
										}
										v542 = base.F64_mul(v540, v509)
										v554 = v542
									}
								case 2:
									if base.F64_lt(v445, float64(-0.25)) != 0 {
										v554 = base.F64_mul(base.F64_sub(v485, base.F64_add(v445, float64(0.5))), float64(-2))
									} else {
										v500 = base.F64_sub(v445, v485)
										v554 = base.F64_add(base.F64_add(v500, v500), float64(1))
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(1072734897)) < base.Ui32(v392) {
							v428 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(v380, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v380)))
							v429 = base.F64_convert_i32_s(v428)
							v435 = v428
							v436 = base.F64_mul(v429, float64(1.9082149292705877e-10))
							v438 = base.F64_add(v380, base.F64_mul(v429, float64(-0.6931471803691238)))
						} else {
							if int64(0) <= v387 {
								v435 = int32(1)
								v436 = float64(1.9082149292705877e-10)
								v438 = base.F64_add(v380, float64(-0.6931471803691238))
							} else {
								v435 = int32(-1)
								v436 = float64(-1.9082149292705877e-10)
								v438 = base.F64_add(v380, float64(0.6931471803691238))
							}
						}
						v439 = base.F64_sub(v438, v436)
						v445 = v439
						v446 = v435
						v447 = base.F64_sub(base.F64_sub(v438, v439), v436)
						v450 = base.F64_mul(v445, float64(0.5))
						v451 = base.F64_mul(v445, v450)
						v467 = base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, base.F64_add(base.F64_mul(v451, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v470 = base.F64_sub(float64(3), base.F64_mul(v467, v450))
						v476 = base.F64_mul(v451, base.F64_div(base.F64_sub(v467, v470), base.F64_sub(float64(6), base.F64_mul(v445, v470))))
						if v446 == int32(0) {
							v554 = base.F64_sub(v445, base.F64_sub(base.F64_mul(v445, v476), v451))
						} else {
							v485 = base.F64_sub(base.F64_sub(base.F64_mul(v445, base.F64_sub(v476, v447)), v447), v451)
							switch v446 + int32(1) {
							case 0:
								v554 = base.F64_add(base.F64_mul(base.F64_sub(v445, v485), float64(0.5)), float64(-0.5))
							default:
								v509 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v446+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v446) {
									v514 = base.F64_add(base.F64_sub(v445, v485), float64(1))
									if v446 == int32(1024) {
										v521 = base.F64_mul(base.F64_add(v514, v514), float64(8.98846567431158e+307))
									} else {
										v521 = base.F64_mul(v514, v509)
									}
									v554 = base.F64_add(v521, float64(-1))
								} else {
									v524 = float64(1)
									v530 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v446) << (uint(int64(52)) % 64))
									if base.Ui32(v446) <= base.Ui32(int32(19)) {
										v540 = base.F64_add(base.F64_sub(v524, v530), base.F64_sub(v445, v485))
									} else {
										v540 = base.F64_add(base.F64_sub(v445, base.F64_add(v485, v530)), v524)
									}
									v542 = base.F64_mul(v540, v509)
									v554 = v542
								}
							case 2:
								if base.F64_lt(v445, float64(-0.25)) != 0 {
									v554 = base.F64_mul(base.F64_sub(v485, base.F64_add(v445, float64(0.5))), float64(-2))
								} else {
									v500 = base.F64_sub(v445, v485)
									v554 = base.F64_add(base.F64_add(v500, v500), float64(1))
								}
							}
						}
					}
				}
				v559 = base.F64_div(base.F64_neg(v554), base.F64_add(v554, float64(2)))
			}
		}
	}
	if base.I64_reinterpret_f64(v5) < int64(0) {
		v564 = base.F64_neg(v559)
	} else {
		v564 = v559
	}
	if base.F64_eq(base.F64_abs(v564), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v571 = m.ExcPending
		if v571 != 0 {
			return int64(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I64_reinterpret_f64(v564)
	}
}
func F_dtof(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v5 float32
	_ = v5
	var v15 int32
	_ = v15
	var v16 float64
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	v4 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = base.F32_demote_f64(v4)
	if base.F32_ne(base.F32_abs(v5), math.Float32frombits(uint32(0x7f800000)))|base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = F_float_overflow_error_ext(m, v15)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int64(0)
		} else {
			if base.F32_ne(v5, float32(0))|base.F64_eq(v4, float64(0)) == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v28 = F_float_underflow_error_ext(m, v27)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_s(base.I32_reinterpret_f32(v5))
				}
			} else {
				return base.I64_extend_i32_s(base.I32_reinterpret_f32(v5))
			}
		}
	} else {
		if base.F32_ne(v5, float32(0))|base.F64_eq(v4, float64(0)) == int32(0) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v28 = F_float_underflow_error_ext(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_s(base.I32_reinterpret_f32(v5))
			}
		} else {
			return base.I64_extend_i32_s(base.I32_reinterpret_f32(v5))
		}
	}
}
func F_dummy_handler(m *base.Module, l0 int32, l1 int32) {
	return
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
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2248), int32(_a_F_dumptuples_2))
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
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2259), int32(_a_F_dumptuples_2))
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
	m.T0[v198].(func(*base.Module, int32, int32, int32))(m, l0, v193, v194+v186*int32(24))
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
	v164 = v155 * int32(24)
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
	m.T0[v175].(func(*base.Module, int32, int32, int32))(m, l0, v170, v171+v164+int32(24))
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
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2292), int32(_a_F_dumptuples_2))
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
	F_errfinish(m, int32(_a_F_dumptuples_1), int32(2238), int32(_a_F_dumptuples_2))
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
