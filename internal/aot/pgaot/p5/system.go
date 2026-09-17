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
	if base.Ui32(v4) < base.Ui32(int32(_a_F_IsSystemRelation_0)) {
		v19 = v3
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+68))
		if v8 == int32(99) {
			v19 = v3
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, _c_F_IsSystemRelation[0]))
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
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
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
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	v2 = int32(_a_F_SystemAttributeByName_0)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SystemAttributeByName[0])))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v5 == int32(0))|base.B2i32(v5 != v8) != 0 {
		v26 = v5
		v27 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26-v27 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L2:
	;
	goto L1
L3:
	;
	v11 = v2
	v12 = l0
	goto L4
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v16 == int32(0) {
		v26 = v16
		v27 = v15
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v26 = v16
	v27 = v15
	goto L2
L6:
	;
	v19 = int32(1)
	if v16 == v15 {
		v11 = v11 + v19
		v12 = v12 + v19
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	return int32(_a_F_SystemAttributeByName_1)
L9:
	;
	goto L10
L10:
	;
	v33 = int32(_a_F_SystemAttributeByName_2)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SystemAttributeByName[1])))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v36 == int32(0))|base.B2i32(v36 != v39) != 0 {
		v57 = v36
		v58 = v39
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v57-v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	goto L11
L13:
	;
	v42 = v33
	v43 = l0
	goto L14
L14:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v47
		v58 = v46
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v57 = v47
	v58 = v46
	goto L12
L16:
	;
	v50 = int32(1)
	if v47 == v46 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	return int32(_a_F_SystemAttributeByName_3)
L19:
	;
	goto L20
L20:
	;
	v64 = int32(_a_F_SystemAttributeByName_4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SystemAttributeByName[2])))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if v88-v89 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L22:
	;
	goto L21
L23:
	;
	v73 = v64
	v74 = l0
	goto L24
L24:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v88 = v78
	v89 = v77
	goto L22
L26:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	return int32(_a_F_SystemAttributeByName_5)
L29:
	;
	goto L30
L30:
	;
	v95 = int32(_a_F_SystemAttributeByName_6)
	v98 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SystemAttributeByName[3])))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v98 == int32(0))|base.B2i32(v98 != v101) != 0 {
		v119 = v98
		v120 = v101
		goto L32
	} else {
		goto L33
	}
L31:
	;
	if v119-v120 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	goto L31
L33:
	;
	v104 = v95
	v105 = l0
	goto L34
L34:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	if v109 == int32(0) {
		v119 = v109
		v120 = v108
		goto L32
	} else {
		goto L36
	}
L35:
	;
	v119 = v109
	v120 = v108
	goto L32
L36:
	;
	v112 = int32(1)
	if v109 == v108 {
		v104 = v104 + v112
		v105 = v105 + v112
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	return int32(_a_F_SystemAttributeByName_7)
L39:
	;
	goto L40
L40:
	;
	v126 = int32(_a_F_SystemAttributeByName_8)
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SystemAttributeByName[4])))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v129 == int32(0))|base.B2i32(v129 != v132) != 0 {
		v150 = v129
		v151 = v132
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v150-v151 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L42:
	;
	goto L41
L43:
	;
	v135 = v126
	v136 = l0
	goto L44
L44:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	if v140 == int32(0) {
		v150 = v140
		v151 = v139
		goto L42
	} else {
		goto L46
	}
L45:
	;
	v150 = v140
	v151 = v139
	goto L42
L46:
	;
	v143 = int32(1)
	if v140 == v139 {
		v135 = v135 + v143
		v136 = v136 + v143
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	return int32(_a_F_SystemAttributeByName_9)
L49:
	;
	goto L50
L50:
	;
	v157 = int32(0)
	v159 = int32(_a_F_SystemAttributeByName_10)
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SystemAttributeByName[5])))
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v162 == v157)|base.B2i32(v162 != v165) != 0 {
		v183 = v162
		v184 = v165
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v183-v184 != 0 {
		goto L58
	} else {
		goto L59
	}
L52:
	;
	goto L51
L53:
	;
	v168 = v159
	v169 = l0
	goto L54
L54:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v173 == int32(0) {
		v183 = v173
		v184 = v172
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v183 = v173
	v184 = v172
	goto L52
L56:
	;
	v176 = int32(1)
	if v173 == v172 {
		v168 = v168 + v176
		v169 = v169 + v176
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v186 = v157
	goto L60
L59:
	;
	v186 = int32(_a_F_SystemAttributeByName_11)
	goto L60
L60:
	;
	return v186
}
