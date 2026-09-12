package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cmpNodePtr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	return v3 - v4
}
func F_cmp_fxid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
}
func F_cmp_lsn(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(base.Ui64(v6) < base.Ui64(v5)) - base.B2i32(base.Ui64(v5) < base.Ui64(v6))
}
func F_cmpaffix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v196 int32
	_ = v196
	v8 = int32(1)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = v9 & v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v14 = v12 & v8
	if base.Ui32(v11) < base.Ui32(v14) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if base.Ui32(v14) < base.Ui32(v11) {
		v196 = v8
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v196
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 == int32(0) {
		v45 = v25
		v46 = v26
		goto L10
	} else {
		goto L11
	}
L7:
	;
	goto L8
L8:
	;
	if v20&int32(3) == int32(0) {
		v72 = v20
		goto L19
	} else {
		goto L20
	}
L9:
	;
	return v46 - v45
L10:
	;
	goto L9
L11:
	;
	if v25 != v26 {
		v45 = v25
		v46 = v26
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v30 = v20
	v31 = v19
	goto L13
L13:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+1)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+1)))
	if v35 == int32(0) {
		v45 = v34
		v46 = v35
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v45 = v34
	v46 = v35
	goto L10
L15:
	;
	v38 = int32(1)
	if v34 == v35 {
		v30 = v30 + v38
		v31 = v31 + v38
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	if v19&int32(3) == int32(0) {
		v129 = v19
		goto L36
	} else {
		goto L37
	}
L18:
	;
	v105 = v97 - v20
	goto L17
L19:
	;
	v76 = v72
	goto L28
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v56 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v105 = int32(0)
	goto L17
L22:
	;
	goto L23
L23:
	;
	v61 = v20
	goto L24
L24:
	;
	v65 = v61 + int32(1)
	if v65&int32(3) == int32(0) {
		v72 = v65
		goto L19
	} else {
		goto L26
	}
L25:
	;
	v97 = v65
	goto L18
L26:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v70 != 0 {
		v61 = v65
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	v85 = int32(-2139062144)
	if (int32(16843008)-v82|v82)&v85 == v85 {
		v76 = v76 + int32(4)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v91 = v76
	goto L31
L30:
	;
	goto L29
L31:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91))))
	if v95 != 0 {
		v91 = v91 + int32(1)
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v97 = v91
	goto L18
L33:
	;
	goto L32
L34:
	;
	v163 = v105
	v164 = v162
	goto L51
L35:
	;
	v162 = v154 - v19
	goto L34
L36:
	;
	v133 = v129
	goto L45
L37:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v113 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v162 = int32(0)
	goto L34
L39:
	;
	goto L40
L40:
	;
	v118 = v19
	goto L41
L41:
	;
	v122 = v118 + int32(1)
	if v122&int32(3) == int32(0) {
		v129 = v122
		goto L36
	} else {
		goto L43
	}
L42:
	;
	v154 = v122
	goto L35
L43:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v127 != 0 {
		v118 = v122
		goto L41
	} else {
		goto L44
	}
L44:
	;
	goto L42
L45:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	v142 = int32(-2139062144)
	if (int32(16843008)-v139|v139)&v142 == v142 {
		v133 = v133 + int32(4)
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v148 = v133
	goto L48
L47:
	;
	goto L46
L48:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148))))
	if v152 != 0 {
		v148 = v148 + int32(1)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v154 = v148
	goto L35
L50:
	;
	goto L49
L51:
	;
	v170 = int32(1)
	v171 = v164 - v170
	v173 = v163 - v170
	if v173 < int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v173 < v171 {
		goto L60
	} else {
		goto L61
	}
L53:
	;
	goto L52
L54:
	;
	if v171 < int32(0) {
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173+v20))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171+v19))))
	if base.Ui32(v179) < base.Ui32(v181) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(-1)
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v179) <= base.Ui32(v181) {
		v163 = v173
		v164 = v171
		goto L51
	} else {
		goto L59
	}
L59:
	;
	v196 = v8
	goto L4
L60:
	;
	return int32(-1)
L61:
	;
	goto L62
L62:
	;
	v196 = base.B2i32(v171 < v173)
	goto L4
}
