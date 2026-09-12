package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SyncRepInitConfig(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = int32(*(*uint8)(unsafe.Add(mBase, _consts[683])))
	if v10 != 0 {
		v155 = v1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)+72))
	if v160 == v155 {
		goto L49
	} else {
		goto L50
	}
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _consts[685]))
	if v12 == int32(0) {
		v155 = v1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v15 == int32(0) {
		v155 = v1
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[686]))
	if v19 == int32(0) {
		v155 = v1
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	if v22 <= int32(0) {
		v155 = v1
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v28 = v19 + int32(16)
	v30 = int32(1)
	goto L7
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[687]))
	v36 = v28
	v37 = v33
	goto L11
L8:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _consts[686]))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151)+8)))
	if v152 != 0 {
		goto L46
	} else {
		goto L47
	}
L9:
	;
	goto L8
L10:
	;
	if v74 == int32(0) {
		goto L9
	} else {
		goto L23
	}
L11:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v40 == v41 {
		v63 = v40
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v74 = int32(0)
	goto L10
L13:
	;
	v65 = int32(1)
	if v63 != 0 {
		v36 = v36 + v65
		v37 = v37 + v65
		goto L11
	} else {
		goto L22
	}
L14:
	;
	if base.Ui32((v40-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v51 = v40 | int32(32)
	goto L17
L16:
	;
	v51 = v40
	goto L17
L17:
	;
	if base.Ui32((v41-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v41 | int32(32)
	goto L20
L19:
	;
	v60 = v41
	goto L20
L20:
	;
	if v51 == v60 {
		v63 = v51
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v74 = v51 - v60
	goto L10
L22:
	;
	goto L12
L23:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v77 == int32(42) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)))
	if v80 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v28&int32(3) == int32(0) {
		v106 = v28
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L26
L28:
	;
	v141 = int32(1)
	v144 = v30 + v141
	v146 = *(*int32)(unsafe.Add(mBase, _consts[686]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+12))
	if v144 <= v147 {
		v28 = v139 + v28 + v141
		v30 = v144
		goto L7
	} else {
		goto L45
	}
L29:
	;
	v139 = v131 - v28
	goto L28
L30:
	;
	v110 = v106
	goto L39
L31:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v90 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v139 = int32(0)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v95 = v28
	goto L35
L35:
	;
	v99 = v95 + int32(1)
	if v99&int32(3) == int32(0) {
		v106 = v99
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v131 = v99
	goto L29
L37:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99))))
	if v104 != 0 {
		v95 = v99
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)))
	v119 = int32(-2139062144)
	if (int32(16843008)-v116|v116)&v119 == v119 {
		v110 = v110 + int32(4)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v125 = v110
	goto L42
L41:
	;
	goto L40
L42:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v129 != 0 {
		v125 = v125 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v131 = v125
	goto L29
L44:
	;
	goto L43
L45:
	;
	v155 = v1
	goto L1
L46:
	;
	v153 = int32(1)
	goto L48
L47:
	;
	v153 = v30
	goto L48
L48:
	;
	v155 = v153
	goto L1
L49:
	;
	m.G0 = v7 + int32(16)
	return
L50:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v159)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+76)) = int32(1)
	if v162 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	F_s_lock(m, v166+int32(76), int32(472984), int32(456), int32(320955))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L53
L53:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[684]))
	v176 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175)+76)) = v176
	*(*int32)(unsafe.Add(mBase, uint32(v175)+72)) = v155
	v181 = F_errstart(m, int32(14), v176)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L54
	} else {
		goto L56
	}
L54:
	;
	return
L55:
	;
	goto L53
L56:
	;
	if v181 == int32(0) {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v155
	v187 = *(*int32)(unsafe.Add(mBase, _consts[687]))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v187
	F_errmsg_internal(m, int32(445581), v7)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(472984), int32(462), int32(320955))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L49
}
