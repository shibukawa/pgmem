package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_statext_is_compatible_clause_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	v6 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = l0
	v16 = v10
	goto L4
L1:
	;
	return v204
L2:
	;
	v204 = int32(1)
	goto L1
L3:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v190 = F_lappend(m, v189, v25)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L18
	} else {
		goto L70
	}
L4:
	;
	if v16 != int32(27) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v186 = F_lappend(m, v185, v181)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L18
	} else {
		goto L69
	}
L6:
	;
	switch v24 - int32(6) {
	case 0:
		goto L14
	default:
		goto L3
	case 11:
		goto L13
	case 14:
		goto L12
	case 15:
		goto L11
	case 46:
		goto L10
	}
L7:
	;
	v24 = v16
	v25 = v11
	goto L6
L8:
	;
	goto L9
L9:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = v23
	v25 = v22
	goto L6
L10:
	;
	v180 = int32(6)
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v181)))
	if v182 == v180 {
		v11 = v181
		v16 = v180
		goto L4
	} else {
		goto L68
	}
L11:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if base.Ui32(int32(2)) < base.Ui32(v147) {
		goto L3
	} else {
		goto L60
	}
L12:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v96 == int32(0) {
		v204 = v6
		goto L1
	} else {
		goto L42
	}
L13:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v40 == int32(0) {
		v204 = v6
		goto L1
	} else {
		goto L20
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 != l1 {
		v204 = v6
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v30 != 0 {
		v204 = v6
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v31 = int32(*(*int16)(unsafe.Add(mBase, uint32(v25)+8)))
	if v31 <= int32(0) {
		v204 = v6
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v35 = F_bms_add_member(m, v34, v31)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v35
	goto L2
L20:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40)+4))
	if v43 != int32(2) {
		v204 = v6
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v49 == int32(27) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	v53 = v52
	goto L24
L23:
	;
	v53 = v48
	goto L24
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	if v54 == int32(27) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v68 = F_get_oprrest(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L33
	}
L26:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47)+4))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v59 = v57
	v60 = v58
	goto L28
L27:
	;
	v59 = v47
	v60 = v54
	goto L28
L28:
	;
	if v60 == int32(7) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v66 = v53
	goto L25
L30:
	;
	goto L31
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v63 != int32(7) {
		v204 = v6
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v66 = v59
	goto L25
L33:
	;
	if base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v68-int32(101)))&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v68-int32(336))) != 0 {
		v204 = v6
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v79 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v83 = F_get_opcode(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L18
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v88 = int32(6)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v89 == v88 {
		v11 = v66
		v16 = v88
		goto L4
	} else {
		goto L40
	}
L38:
	;
	v85 = F_get_func_leakproof(m, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L39
	}
L39:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v85)
	goto L37
L40:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v93 = F_lappend(m, v92, v66)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L18
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v93
	goto L2
L42:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 != int32(2) {
		v204 = v6
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105 == int32(27) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v109 = v108
	goto L46
L45:
	;
	v109 = v104
	goto L46
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v110 == int32(27) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = v114
	goto L49
L48:
	;
	v115 = v110
	goto L49
L49:
	;
	if v115 != int32(7) {
		v204 = v6
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v119 = F_get_oprrest(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	if base.B2i32(base.Ui32(int32(4)) <= base.Ui32(v119-int32(101)))&base.B2i32(base.Ui32(int32(1)) < base.Ui32(v119-int32(336))) != 0 {
		v204 = v6
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v130 == int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v134 = F_get_opcode(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L18
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v139 = int32(6)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v140 == v139 {
		v11 = v109
		v16 = v139
		goto L4
	} else {
		goto L58
	}
L56:
	;
	v136 = F_get_func_leakproof(m, v134)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L18
	} else {
		goto L57
	}
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v136)
	goto L55
L58:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v144 = F_lappend(m, v143, v109)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L18
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v144
	goto L2
L60:
	;
	v150 = int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	if v151 == int32(0) {
		v204 = v150
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v154 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v155 <= v154 {
		v204 = v150
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v158 = v154
	goto L63
L63:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v151)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v167+v158<<(uint(int32(2))%32))))
	v172 = F_statext_is_compatible_clause_internal(m, v171, l1, l2, l3, l4)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L18
	} else {
		goto L65
	}
L64:
	;
	v204 = v172
	goto L1
L65:
	;
	if v172 == int32(0) {
		v204 = v172
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v177 = v158 + int32(1)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v151)+4))
	if v177 < v178 {
		v158 = v177
		goto L63
	} else {
		goto L67
	}
L67:
	;
	goto L64
L68:
	;
	goto L5
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v186
	goto L2
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v190
	goto L2
}
