package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_sts_initialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v6 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
	if l4&int32(3) == v6 {
		v36 = l4
		goto L3
	} else {
		goto L4
	}
L1:
	;
	if base.Ui32(v69) < base.Ui32(int32(64)) {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v69 = v61 - l4
	goto L1
L3:
	;
	v40 = v36
	goto L12
L4:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v20 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v69 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v25 = l4
	goto L8
L8:
	;
	v29 = v25 + int32(1)
	if v29&int32(3) == int32(0) {
		v36 = v29
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v61 = v29
	goto L2
L10:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v34 != 0 {
		v25 = v29
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 == v49 {
		v40 = v40 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v55 = v40
	goto L15
L14:
	;
	goto L13
L15:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		v55 = v55 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v61 = v55
	goto L2
L17:
	;
	goto L16
L18:
	;
	v73 = l0 + int32(12)
	if (l4^v73)&int32(3) != 0 {
		goto L24
	} else {
		goto L25
	}
L19:
	;
	goto L20
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L49
	} else {
		goto L51
	}
L21:
	;
	if int32(0) < l1 {
		goto L42
	} else {
		goto L43
	}
L22:
	;
	goto L21
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v127)
	if v127&int32(255) == int32(0) {
		goto L22
	} else {
		goto L38
	}
L24:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v126 = l4
	v127 = v79
	v128 = v73
	goto L23
L25:
	;
	goto L26
L26:
	;
	if l4&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v83 = l4
	v85 = v73
	goto L30
L28:
	;
	v97 = l4
	v99 = v73
	goto L29
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 != v104 {
		v126 = v97
		v127 = v101
		v128 = v99
		goto L23
	} else {
		goto L34
	}
L30:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	*(*uint8)(unsafe.Add(mBase, uint32(v85))) = uint8(v86)
	if v86 == int32(0) {
		goto L22
	} else {
		goto L32
	}
L31:
	;
	v97 = v93
	v99 = v91
	goto L29
L32:
	;
	v90 = int32(1)
	v91 = v85 + v90
	v93 = v83 + v90
	if v93&int32(3) != 0 {
		v83 = v93
		v85 = v91
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v109 = v97
	v110 = v101
	v111 = v99
	goto L35
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v110
	v113 = int32(4)
	v114 = v111 + v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v117 = v109 + v113
	v121 = int32(-2139062144)
	if (v115|(int32(16843008)-v115))&v121 == v121 {
		v109 = v117
		v110 = v115
		v111 = v114
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v126 = v117
	v127 = v115
	v128 = v114
	goto L23
L37:
	;
	goto L36
L38:
	;
	v135 = v126
	v137 = v128
	goto L39
L39:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+1)) = uint8(v138)
	v140 = int32(1)
	if v138 != 0 {
		v135 = v135 + v140
		v137 = v137 + v140
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L22
L41:
	;
	goto L40
L42:
	;
	v157 = v6
	goto L45
L43:
	;
	goto L44
L44:
	;
	v183 = F_palloc0(m, int32(68))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L49
	} else {
		goto L50
	}
L45:
	;
	v161 = l0 + int32(76) + v157*int32(28)
	v162 = int32(75)
	*(*uint16)(unsafe.Add(mBase, uint32(v161))) = uint16(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+4)) = int32(1073741824)
	*(*int64)(unsafe.Add(mBase, uint32(v161)+8)) = int64(-1)
	goto L47
L46:
	;
	goto L44
L47:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v161)+24)) = uint8(v168)
	*(*int64)(unsafe.Add(mBase, uint32(v161)+16)) = int64(0)
	v173 = v157 + int32(1)
	if v173 != l1 {
		v157 = v173
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	return int32(0)
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v183)+8)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v183))) = l2
	v191 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+12)) = v191
	return v183
L51:
	;
	F_errmsg_internal(m, int32(320759), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(487532), int32(143), int32(333881))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
