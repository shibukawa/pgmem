package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SpecialTags(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+16))
	switch v3 - int32(6) {
	case 0:
		goto L4
	case 1:
		goto L5
	case 2:
		goto L6
	default:
		goto L1
	}
L1:
	;
	return
L2:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)) = uint8(v225)
	goto L1
L3:
	;
	v225 = int32(1)
	goto L2
L4:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v176 = v171
	v177 = int32(366936)
	v178 = int32(6)
	goto L59
L5:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v66 = v61
	v67 = int32(366943)
	v68 = int32(7)
	goto L25
L6:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v11 = v6
	v12 = int32(80412)
	v13 = int32(8)
	goto L8
L7:
	;
	if v58 != 0 {
		goto L1
	} else {
		goto L23
	}
L8:
	;
	if v13 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v58 = int32(0)
	goto L7
L10:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v16 == v17 {
		v39 = v16
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	goto L9
L13:
	;
	v41 = int32(1)
	if v39 != 0 {
		v11 = v11 + v41
		v12 = v12 + v41
		v13 = v13 - v41
		goto L8
	} else {
		goto L22
	}
L14:
	;
	if base.Ui32((v16-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v27 = v16 | int32(32)
	goto L17
L16:
	;
	v27 = v16
	goto L17
L17:
	;
	if base.Ui32((v17-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v36 = v17 | int32(32)
	goto L20
L19:
	;
	v36 = v17
	goto L20
L20:
	;
	if v27 == v36 {
		v39 = v27
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v58 = v27 - v36
	goto L7
L22:
	;
	goto L12
L23:
	;
	v225 = int32(0)
	goto L2
L24:
	;
	if v113 == int32(0) {
		v225 = int32(0)
		goto L2
	} else {
		goto L40
	}
L25:
	;
	if v68 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v113 = int32(0)
	goto L24
L27:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v71 == v72 {
		v94 = v71
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v96 = int32(1)
	if v94 != 0 {
		v66 = v66 + v96
		v67 = v67 + v96
		v68 = v68 - v96
		goto L25
	} else {
		goto L39
	}
L31:
	;
	if base.Ui32((v71-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v82 = v71 | int32(32)
	goto L34
L33:
	;
	v82 = v71
	goto L34
L34:
	;
	if base.Ui32((v72-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v91 = v72 | int32(32)
	goto L37
L36:
	;
	v91 = v72
	goto L37
L37:
	;
	if v82 == v91 {
		v94 = v82
		goto L30
	} else {
		goto L38
	}
L38:
	;
	v113 = v82 - v91
	goto L24
L39:
	;
	goto L29
L40:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v121 = v116
	v122 = int32(80404)
	v123 = int32(7)
	goto L42
L41:
	;
	if v168 == int32(0) {
		goto L3
	} else {
		goto L57
	}
L42:
	;
	if v123 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v168 = int32(0)
	goto L41
L44:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
	if v126 == v127 {
		v149 = v126
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L46
L46:
	;
	goto L43
L47:
	;
	v151 = int32(1)
	if v149 != 0 {
		v121 = v121 + v151
		v122 = v122 + v151
		v123 = v123 - v151
		goto L42
	} else {
		goto L56
	}
L48:
	;
	if base.Ui32((v126-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v137 = v126 | int32(32)
	goto L51
L50:
	;
	v137 = v126
	goto L51
L51:
	;
	if base.Ui32((v127-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v146 = v127 | int32(32)
	goto L54
L53:
	;
	v146 = v127
	goto L54
L54:
	;
	if v137 == v146 {
		v149 = v137
		goto L47
	} else {
		goto L55
	}
L55:
	;
	v168 = v137 - v146
	goto L41
L56:
	;
	goto L46
L57:
	;
	goto L1
L58:
	;
	if v223 != 0 {
		goto L1
	} else {
		goto L74
	}
L59:
	;
	if v178 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v223 = int32(0)
	goto L58
L61:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177))))
	if v181 == v182 {
		v204 = v181
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	goto L60
L64:
	;
	v206 = int32(1)
	if v204 != 0 {
		v176 = v176 + v206
		v177 = v177 + v206
		v178 = v178 - v206
		goto L59
	} else {
		goto L73
	}
L65:
	;
	if base.Ui32((v181-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v192 = v181 | int32(32)
	goto L68
L67:
	;
	v192 = v181
	goto L68
L68:
	;
	if base.Ui32((v182-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v201 = v182 | int32(32)
	goto L71
L70:
	;
	v201 = v182
	goto L71
L71:
	;
	if v192 == v201 {
		v204 = v192
		goto L64
	} else {
		goto L72
	}
L72:
	;
	v223 = v192 - v201
	goto L58
L73:
	;
	goto L63
L74:
	;
	goto L3
}
