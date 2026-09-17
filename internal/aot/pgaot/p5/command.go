package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_GetCommandLogLevel(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = l0
	goto L5
L1:
	;
	m.G0 = v10 + int32(48)
	return v191
L2:
	;
	v191 = v189
	goto L1
L3:
	;
	v191 = v185
	goto L1
L4:
	;
	v166 = int32(3)
	v169 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L21
	} else {
		goto L58
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	switch v22 - int32(67) {
	case 0:
		goto L7
	default:
		goto L4
	case 69:
		goto L14
	case 70, 71, 72, 73, 131:
		v185 = int32(2)
		goto L3
	case 74:
		goto L13
	case 77, 91, 92, 134, 135, 136, 144, 146, 155, 156, 157, 158, 164, 172, 177, 178, 179, 180, 181, 187:
		v189 = int32(3)
		goto L2
	case 78, 79, 83, 84, 85, 88, 89, 93, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 128, 129, 130, 132, 133, 137, 138, 140, 141, 143, 148, 149, 150, 151, 152, 153, 154, 159, 160, 161, 162, 163, 165, 166, 167, 168, 169, 170, 171, 175, 176, 182, 183, 184, 188, 189, 190, 191, 194, 195, 196, 197, 198:
		v191 = int32(1)
		goto L1
	case 90:
		goto L12
	case 174:
		goto L9
	case 185:
		goto L11
	case 186:
		goto L10
	case 263:
		goto L8
	}
L6:
	;
	v147 = int32(3)
	v150 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L21
	} else {
		goto L54
	}
L7:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v139 != int32(6) {
		goto L51
	} else {
		goto L52
	}
L8:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v112 != int32(6) {
		goto L43
	} else {
		goto L44
	}
L9:
	;
	v49 = int32(3)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v50 == int32(0) {
		v191 = v49
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v35 = int32(3)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v38 = F_FetchPreparedStatement(m, v36, int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L21
	} else {
		goto L22
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v12 = v34
	goto L5
L12:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if v32 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v28 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v12 = v25
	goto L5
L15:
	;
	v29 = int32(1)
	goto L17
L16:
	;
	v29 = int32(3)
	goto L17
L17:
	;
	v191 = v29
	goto L1
L18:
	;
	v33 = int32(2)
	goto L20
L19:
	;
	v33 = int32(3)
	goto L20
L20:
	;
	v191 = v33
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	if v38 == int32(0) {
		v185 = v35
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v45 == int32(0) {
		v185 = v35
		goto L3
	} else {
		goto L24
	}
L24:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v12 = v48
	goto L5
L25:
	;
	v53 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v55 <= v53 {
		v191 = v49
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v60 = v53
	v63 = v53
	goto L27
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v60<<(uint(int32(2))%32))))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v71 = int32(_a_F_GetCommandLogLevel_0)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_GetCommandLogLevel[0])))
	if base.B2i32(v74 == int32(0))|base.B2i32(v74 != v77) != 0 {
		v95 = v74
		v96 = v77
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v102&int32(1) == int32(0) {
		v191 = v49
		goto L1
	} else {
		goto L41
	}
L29:
	;
	if v95-v96 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v80 = v70
	v81 = v71
	goto L32
L32:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	if v85 == int32(0) {
		v95 = v85
		v96 = v84
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v95 = v85
	v96 = v84
	goto L30
L34:
	;
	v88 = int32(1)
	if v85 == v84 {
		v80 = v80 + v88
		v81 = v81 + v88
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v100 = F_defGetBoolean(m, v69)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L21
	} else {
		goto L39
	}
L37:
	;
	v102 = v63
	goto L38
L38:
	;
	v104 = v60 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v104 < v105 {
		v60 = v104
		v63 = v102
		goto L27
	} else {
		goto L40
	}
L39:
	;
	v102 = v100
	goto L38
L40:
	;
	goto L28
L41:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v12 = v111
	goto L5
L42:
	;
	v120 = int32(3)
	v123 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L21
	} else {
		goto L46
	}
L43:
	;
	switch v112 - int32(1) {
	case 0:
		v185 = int32(3)
		goto L3
	case 1, 2, 3, 4:
		v189 = int32(2)
		goto L2
	default:
		goto L42
	}
L44:
	;
	goto L45
L45:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v12 = v119
	goto L5
L46:
	;
	if v123 == int32(0) {
		v191 = v120
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v127
	F_errmsg_internal(m, int32(_a_F_GetCommandLogLevel_1), v10+int32(16))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_GetCommandLogLevel_2), int32(3724), int32(_a_F_GetCommandLogLevel_3))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	v191 = v120
	goto L1
L50:
	;
	goto L6
L51:
	;
	switch v139 - int32(1) {
	case 0:
		v185 = int32(3)
		goto L3
	case 1, 2, 3, 4:
		v189 = int32(2)
		goto L2
	default:
		goto L50
	}
L52:
	;
	goto L53
L53:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v12 = v146
	goto L5
L54:
	;
	if v150 == int32(0) {
		v191 = v147
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v154
	F_errmsg_internal(m, int32(_a_F_GetCommandLogLevel_1), v10+int32(32))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_GetCommandLogLevel_2), int32(3755), int32(_a_F_GetCommandLogLevel_3))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	v191 = v147
	goto L1
L58:
	;
	if v169 == int32(0) {
		v191 = v166
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v173
	F_errmsg_internal(m, int32(_a_F_GetCommandLogLevel_4), v10)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_GetCommandLogLevel_2), int32(3764), int32(_a_F_GetCommandLogLevel_3))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	v191 = v166
	goto L1
}
