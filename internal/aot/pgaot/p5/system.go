package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_IsSystemRelation(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	v3 = int32(1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if base.Ui32(v4) < base.Ui32(int32(12000)) {
		v19 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
		if v8 == int32(99) {
			v19 = v3
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[220]))
			v19 = base.B2i32(v13 != int32(0)) & base.B2i32(v8 == v13)
		}
	}
	return v19
}
func F_SystemAttributeByName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	v2 = int32(736204)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, _consts[224])))
	if v6 == int32(0) {
		v25 = v5
		v26 = v6
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v25 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L2:
	;
	goto L1
L3:
	;
	if v5 != v6 {
		v25 = v5
		v26 = v6
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v10 = v2
	v11 = l0
	goto L5
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v15 == int32(0) {
		v25 = v14
		v26 = v15
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v25 = v14
	v26 = v15
	goto L2
L7:
	;
	v18 = int32(1)
	if v14 == v15 {
		v10 = v10 + v18
		v11 = v11 + v18
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	return int32(736200)
L10:
	;
	goto L11
L11:
	;
	v32 = int32(736304)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _consts[225])))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v56-v55 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v40 = v32
	v41 = l0
	goto L16
L16:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v55 = v44
	v56 = v45
	goto L13
L18:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	return int32(736300)
L21:
	;
	goto L22
L22:
	;
	v62 = int32(736404)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _consts[226])))
	if v66 == int32(0) {
		v85 = v65
		v86 = v66
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v86-v85 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L24:
	;
	goto L23
L25:
	;
	if v65 != v66 {
		v85 = v65
		v86 = v66
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v70 = v62
	v71 = l0
	goto L27
L27:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+1)))
	if v75 == int32(0) {
		v85 = v74
		v86 = v75
		goto L24
	} else {
		goto L29
	}
L28:
	;
	v85 = v74
	v86 = v75
	goto L24
L29:
	;
	v78 = int32(1)
	if v74 == v75 {
		v70 = v70 + v78
		v71 = v71 + v78
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	return int32(736400)
L32:
	;
	goto L33
L33:
	;
	v92 = int32(736504)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, _consts[227])))
	if v96 == int32(0) {
		v115 = v95
		v116 = v96
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v116-v115 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	goto L34
L36:
	;
	if v95 != v96 {
		v115 = v95
		v116 = v96
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v100 = v92
	v101 = l0
	goto L38
L38:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v104
		v116 = v105
		goto L35
	} else {
		goto L40
	}
L39:
	;
	v115 = v104
	v116 = v105
	goto L35
L40:
	;
	v108 = int32(1)
	if v104 == v105 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	return int32(736500)
L43:
	;
	goto L44
L44:
	;
	v122 = int32(736604)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _consts[228])))
	if v126 == int32(0) {
		v145 = v125
		v146 = v126
		goto L46
	} else {
		goto L47
	}
L45:
	;
	if v146-v145 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L46:
	;
	goto L45
L47:
	;
	if v125 != v126 {
		v145 = v125
		v146 = v126
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v130 = v122
	v131 = l0
	goto L49
L49:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+1)))
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+1)))
	if v135 == int32(0) {
		v145 = v134
		v146 = v135
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v145 = v134
	v146 = v135
	goto L46
L51:
	;
	v138 = int32(1)
	if v134 == v135 {
		v130 = v130 + v138
		v131 = v131 + v138
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	return int32(736600)
L54:
	;
	goto L55
L55:
	;
	v152 = int32(0)
	v154 = int32(736704)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, _consts[229])))
	if v158 == v152 {
		v177 = v157
		v178 = v158
		goto L57
	} else {
		goto L58
	}
L56:
	;
	if v178-v177 != 0 {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	goto L56
L58:
	;
	if v157 != v158 {
		v177 = v157
		v178 = v158
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v162 = v154
	v163 = l0
	goto L60
L60:
	;
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v162)+1)))
	if v167 == int32(0) {
		v177 = v166
		v178 = v167
		goto L57
	} else {
		goto L62
	}
L61:
	;
	v177 = v166
	v178 = v167
	goto L57
L62:
	;
	v170 = int32(1)
	if v166 == v167 {
		v162 = v162 + v170
		v163 = v163 + v170
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v180 = v152
	goto L66
L65:
	;
	v180 = int32(736700)
	goto L66
L66:
	;
	return v180
}
