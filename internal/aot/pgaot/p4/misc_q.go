package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QTNodeCompare(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	F_check_stack_depth(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v18 != v20 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return v152
L4:
	;
	if base.I32_extend8_s(v20) < base.I32_extend8_s(v18) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	switch v18 - int32(1) {
	case 0:
		goto L12
	case 1:
		goto L13
	default:
		goto L11
	}
L7:
	;
	v27 = int32(-1)
	goto L9
L8:
	;
	v27 = int32(1)
	goto L9
L9:
	;
	v152 = v27
	goto L3
L10:
	;
	if v40 < v39 {
		goto L71
	} else {
		goto L72
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L68
	}
L12:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v90 != v91 {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v31 = base.I32_extend8_s(v30)
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v32 != v30 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if base.I32_extend8_s(v32) < v31 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v39 != v40 {
		goto L10
	} else {
		goto L20
	}
L17:
	;
	v38 = int32(-1)
	goto L19
L18:
	;
	v38 = int32(1)
	goto L19
L19:
	;
	v152 = v38
	goto L3
L20:
	;
	if int32(0) < v39 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v48 = int32(0)
	goto L24
L22:
	;
	v71 = v31
	goto L23
L23:
	;
	v76 = int32(0)
	if v71&int32(255) != int32(4) {
		v152 = v76
		goto L3
	} else {
		goto L29
	}
L24:
	;
	v54 = v48 << (uint(int32(2)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v54+v55)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v58+v54)))
	v61 = F_QTNodeCompare(m, v57, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	v71 = v67
	goto L23
L26:
	;
	if v61 != 0 {
		v152 = v61
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v64 = v48 + int32(1)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v64 < v65 {
		v48 = v64
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v81 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17)+2)))
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+2)))
	if v81 == v82 {
		v152 = v76
		goto L3
	} else {
		goto L30
	}
L30:
	;
	if base.I32_extend16_s(v82) < base.I32_extend16_s(v81) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v89 = int32(-1)
	goto L33
L32:
	;
	v89 = int32(1)
	goto L33
L33:
	;
	v152 = v89
	goto L3
L34:
	;
	if v91 < v90 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v99 = int32(4095)
	v100 = v98 & v99
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v104 = v102 & v99
	if v100 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v96 = int32(-1)
	goto L39
L38:
	;
	v96 = int32(1)
	goto L39
L39:
	;
	v152 = v96
	goto L3
L40:
	;
	v152 = v130
	goto L3
L41:
	;
	goto L45
L42:
	;
	goto L43
L43:
	;
	if v104 == int32(0) {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	goto L46
L46:
	;
	v110 = int32(0)
	if v110 < v104 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v113 = int32(-1)
	goto L49
L48:
	;
	v113 = v110
	goto L49
L49:
	;
	v130 = v113
	goto L40
L50:
	;
	v130 = base.B2i32(int32(0) < v100)
	goto L40
L51:
	;
	goto L52
L52:
	;
	if base.Ui32(v100) < base.Ui32(v104) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v119 = v100
	goto L55
L54:
	;
	v119 = v104
	goto L55
L55:
	;
	v120 = F_memcmp(m, v97, v101, v119)
	mBase = m.M
	goto L58
L56:
	;
	v130 = v128
	goto L40
L58:
	;
	goto L59
L59:
	;
	if v120 != 0 {
		v128 = v120
		goto L56
	} else {
		goto L61
	}
L61:
	;
	if v100 == v104 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v130 = int32(0)
	goto L40
L63:
	;
	goto L64
L64:
	;
	if v100 < v104 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v127 = int32(-1)
	goto L67
L66:
	;
	v127 = int32(1)
	goto L67
L67:
	;
	v128 = v127
	goto L56
L68:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v136 = int32(*(*int8)(unsafe.Add(mBase, uint32(v135))))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v136
	F_errmsg_internal(m, int32(506305), v11)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(520231), int32(144), int32(382076))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	v149 = int32(-1)
	goto L73
L72:
	;
	v149 = int32(1)
	goto L73
L73:
	;
	v152 = v149
	goto L3
}
