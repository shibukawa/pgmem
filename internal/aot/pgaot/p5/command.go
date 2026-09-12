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
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
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
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = l0
	goto L5
L1:
	;
	m.G0 = v10 + int32(48)
	return v190
L2:
	;
	v190 = v188
	goto L1
L3:
	;
	v190 = v184
	goto L1
L4:
	;
	v165 = int32(3)
	v168 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L59
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
		v184 = int32(2)
		goto L3
	case 74:
		goto L13
	case 77, 91, 92, 134, 135, 136, 144, 146, 155, 156, 157, 158, 164, 172, 177, 178, 179, 180, 181, 187:
		v188 = int32(3)
		goto L2
	case 78, 79, 83, 84, 85, 88, 89, 93, 95, 96, 97, 98, 99, 100, 101, 102, 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, 115, 116, 117, 118, 119, 120, 121, 122, 123, 124, 125, 126, 128, 129, 130, 132, 133, 137, 138, 140, 141, 143, 148, 149, 150, 151, 152, 153, 154, 159, 160, 161, 162, 163, 165, 166, 167, 168, 169, 170, 171, 175, 176, 182, 183, 184, 188, 189, 190, 191, 194, 195, 196, 197, 198:
		v190 = int32(1)
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
	v146 = int32(3)
	v149 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L21
	} else {
		goto L55
	}
L7:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v138 != int32(6) {
		goto L52
	} else {
		goto L53
	}
L8:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v111 != int32(6) {
		goto L44
	} else {
		goto L45
	}
L9:
	;
	v49 = int32(3)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v50 == int32(0) {
		v190 = v49
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
	v190 = v29
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
	v190 = v33
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	if v38 == int32(0) {
		v184 = v35
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)+64))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	if v45 == int32(0) {
		v184 = v35
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
		v190 = v49
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
	v71 = int32(358308)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, _consts[937])))
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v75 == int32(0) {
		v94 = v74
		v95 = v75
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if v101&int32(1) == int32(0) {
		v190 = v49
		goto L1
	} else {
		goto L42
	}
L29:
	;
	if v95-v94 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L30:
	;
	goto L29
L31:
	;
	if v74 != v75 {
		v94 = v74
		v95 = v75
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v79 = v70
	v80 = v71
	goto L33
L33:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+1)))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+1)))
	if v84 == int32(0) {
		v94 = v83
		v95 = v84
		goto L30
	} else {
		goto L35
	}
L34:
	;
	v94 = v83
	v95 = v84
	goto L30
L35:
	;
	v87 = int32(1)
	if v83 == v84 {
		v79 = v79 + v87
		v80 = v80 + v87
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v99 = F_defGetBoolean(m, v69)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L21
	} else {
		goto L40
	}
L38:
	;
	v101 = v63
	goto L39
L39:
	;
	v103 = v60 + int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v103 < v104 {
		v60 = v103
		v63 = v101
		goto L27
	} else {
		goto L41
	}
L40:
	;
	v101 = v99
	goto L39
L41:
	;
	goto L28
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v12 = v110
	goto L5
L43:
	;
	v119 = int32(3)
	v122 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L21
	} else {
		goto L47
	}
L44:
	;
	switch v111 - int32(1) {
	case 0:
		v184 = int32(3)
		goto L3
	case 1, 2, 3, 4:
		v188 = int32(2)
		goto L2
	default:
		goto L43
	}
L45:
	;
	goto L46
L46:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	v12 = v118
	goto L5
L47:
	;
	if v122 == int32(0) {
		v190 = v119
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v126
	F_errmsg_internal(m, int32(510175), v10+int32(16))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(516689), int32(3724), int32(322012))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	v190 = v119
	goto L1
L51:
	;
	goto L6
L52:
	;
	switch v138 - int32(1) {
	case 0:
		v184 = int32(3)
		goto L3
	case 1, 2, 3, 4:
		v188 = int32(2)
		goto L2
	default:
		goto L51
	}
L53:
	;
	goto L54
L54:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v12 = v145
	goto L5
L55:
	;
	if v149 == int32(0) {
		v190 = v146
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v153
	F_errmsg_internal(m, int32(510175), v10+int32(32))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(516689), int32(3755), int32(322012))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L21
	} else {
		goto L58
	}
L58:
	;
	v190 = v146
	goto L1
L59:
	;
	if v168 == int32(0) {
		v190 = v165
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v172
	F_errmsg_internal(m, int32(509858), v10)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(516689), int32(3764), int32(322012))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L21
	} else {
		goto L62
	}
L62:
	;
	v190 = v165
	goto L1
}
