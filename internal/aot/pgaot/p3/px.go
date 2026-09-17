package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_px_find_combo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
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
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	v9 = F_palloc0(m, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v13 = F_pstrdup(m, l0)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L6
	}
L3:
	;
	F_pfree(m, v13)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L72
	}
L4:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v187)))
	if v194 != 0 {
		goto L66
	} else {
		goto L67
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = int32(_a_F_px_find_combo_0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_px_find_combo_1)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(_a_F_px_find_combo_2)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(_a_F_px_find_combo_3)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(_a_F_px_find_combo_4)
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_px_find_combo_5)
	F_pfree(m, v13)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L65
	}
L6:
	;
	v15 = int32(47)
	v16 = F___strchrnul(m, v13, v15)
	mBase = m.M
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v18 == v15 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	if v22 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v22 = v16
	goto L10
L9:
	;
	v22 = int32(0)
	goto L10
L10:
	;
	goto L7
L11:
	;
	v23 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v23)
	v27 = v22 + int32(1)
	v31 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v158 = v9 + int32(24)
	v159 = F_px_find_cipher(m, v13, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L63
	}
L14:
	;
	v34 = int32(0)
	v35 = int32(47)
	v36 = F___strchrnul(m, v27, v35)
	mBase = m.M
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v38 == v35 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v94 = v9 + int32(24)
	v95 = F_px_find_cipher(m, v13, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L44
	}
L16:
	;
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v42 = v36
	goto L19
L18:
	;
	v42 = v34
	goto L19
L19:
	;
	goto L16
L20:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v42))) = uint8(v43)
	v47 = v42 + int32(1)
	goto L22
L21:
	;
	v47 = v34
	goto L22
L22:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v48 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v49 = int32(58)
	v50 = F___strchrnul(m, v27, v49)
	mBase = m.M
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	if v52 == v49 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v92 = v31
	goto L25
L25:
	;
	if v47 != 0 {
		v27 = v47
		v31 = v92
		goto L14
	} else {
		goto L43
	}
L26:
	;
	if v56 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v56 = v50
	goto L29
L28:
	;
	v56 = int32(0)
	goto L29
L29:
	;
	goto L26
L30:
	;
	v204 = int32(-6)
	goto L3
L31:
	;
	goto L32
L32:
	;
	v60 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v56))) = uint8(v60)
	v62 = int32(_a_F_px_find_combo_6)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_combo[0])))
	if base.B2i32(v65 == v60)|base.B2i32(v65 != v68) != 0 {
		v86 = v65
		v87 = v68
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v86-v87 != 0 {
		goto L40
	} else {
		goto L41
	}
L34:
	;
	goto L33
L35:
	;
	v71 = v27
	v72 = v62
	goto L36
L36:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v76
		v87 = v75
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v86 = v76
	v87 = v75
	goto L34
L38:
	;
	v79 = int32(1)
	if v76 == v75 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v204 = int32(-5)
	goto L3
L41:
	;
	goto L42
L42:
	;
	v92 = v56 + int32(1)
	goto L25
L43:
	;
	goto L15
L44:
	;
	if v95 != 0 {
		v187 = v94
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v97 = int32(1)
	if v92 == int32(0) {
		v165 = v97
		goto L5
	} else {
		goto L46
	}
L46:
	;
	v100 = int32(_a_F_px_find_combo_7)
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_combo[1])))
	if base.B2i32(v103 == int32(0))|base.B2i32(v103 != v106) != 0 {
		v124 = v103
		v125 = v106
		goto L48
	} else {
		goto L49
	}
L47:
	;
	if v124-v125 == int32(0) {
		v165 = v97
		goto L5
	} else {
		goto L54
	}
L48:
	;
	goto L47
L49:
	;
	v109 = v92
	v110 = v100
	goto L50
L50:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110)+1)))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+1)))
	if v114 == int32(0) {
		v124 = v114
		v125 = v113
		goto L48
	} else {
		goto L52
	}
L51:
	;
	v124 = v114
	v125 = v113
	goto L48
L52:
	;
	v117 = int32(1)
	if v114 == v113 {
		v109 = v109 + v117
		v110 = v110 + v117
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	v129 = int32(_a_F_px_find_combo_8)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_px_find_combo[2])))
	if base.B2i32(v132 == int32(0))|base.B2i32(v132 != v135) != 0 {
		v153 = v132
		v154 = v135
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v153-v154 != 0 {
		v187 = v94
		goto L4
	} else {
		goto L62
	}
L56:
	;
	goto L55
L57:
	;
	v138 = v92
	v139 = v129
	goto L58
L58:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v139)+1)))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138)+1)))
	if v143 == int32(0) {
		v153 = v143
		v154 = v142
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v153 = v143
	v154 = v142
	goto L56
L60:
	;
	v146 = int32(1)
	if v143 == v142 {
		v138 = v138 + v146
		v139 = v139 + v146
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v165 = int32(0)
	goto L5
L63:
	;
	if v159 != 0 {
		v187 = v158
		goto L4
	} else {
		goto L64
	}
L64:
	;
	v165 = int32(1)
	goto L5
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
	return int32(0)
L66:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+24))
	m.T0[v195].(func(*base.Module, int32))(m, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_pfree(m, v9)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	F_pfree(m, v13)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	return int32(-3)
L72:
	;
	F_pfree(m, v9)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	return v204
}
