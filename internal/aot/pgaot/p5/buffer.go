package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferUsageAccumDiff(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	var v6 int64
	_ = v6
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int64
	_ = v45
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v52 int64
	_ = v52
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int64)(unsafe.Add(mBase, _consts[343]))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v3 + (v5 - v6)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int64)(unsafe.Add(mBase, _consts[344]))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + (v12 - v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = *(*int64)(unsafe.Add(mBase, _consts[345]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v17 + (v19 - v20)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int64)(unsafe.Add(mBase, _consts[346]))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v24 + (v26 - v27)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*int64)(unsafe.Add(mBase, _consts[347]))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v31 + (v33 - v34)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int64)(unsafe.Add(mBase, _consts[348]))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v38 + (v40 - v41)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = *(*int64)(unsafe.Add(mBase, _consts[349]))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v45 + (v47 - v48)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v54 = *(*int64)(unsafe.Add(mBase, _consts[350]))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v52 + (v54 - v55)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v61 = *(*int64)(unsafe.Add(mBase, _consts[351]))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v59 + (v61 - v62)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v68 = *(*int64)(unsafe.Add(mBase, _consts[352]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v66 + (v68 - v69)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	v75 = *(*int64)(unsafe.Add(mBase, _consts[353]))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v73 + (v75 - v76)
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v82 = *(*int64)(unsafe.Add(mBase, _consts[354]))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v80 + (v82 - v83)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v89 = *(*int64)(unsafe.Add(mBase, _consts[355]))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v87 + (v89 - v90)
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v96 = *(*int64)(unsafe.Add(mBase, _consts[356]))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v94 + (v96 - v97)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v103 = *(*int64)(unsafe.Add(mBase, _consts[357]))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v101 + (v103 - v104)
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v110 = *(*int64)(unsafe.Add(mBase, _consts[358]))
	v111 = *(*int64)(unsafe.Add(mBase, uint32(l1)+120))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+120)) = v108 + (v110 - v111)
	return
}
func F_CheckBufferIsPinnedOnce(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	if l0 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(32)
	return
L2:
	;
	v14 = (l0 ^ int32(-1)) << (uint(int32(2)) % 32)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[612]))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v14+v16)))
	if v18 == int32(1) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = l0
	v40 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	if l0 == v40 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _consts[612]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26+v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
	F_errmsg_internal(m, int32(451733), v7)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(464766), int32(5653), int32(392353))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L40
	}
L11:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+4))
	if v86 == int32(1) {
		goto L1
	} else {
		goto L39
	}
L12:
	;
	v85 = int32(4341120)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[614]))
	if l0 == v44 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = int32(4341128)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	if l0 == v48 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = int32(4341136)
	goto L11
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	if l0 == v52 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v85 = int32(4341144)
	goto L11
L22:
	;
	goto L23
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	if l0 == v56 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = int32(4341152)
	goto L11
L25:
	;
	goto L26
L26:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[618]))
	if l0 == v60 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v85 = int32(4341160)
	goto L11
L28:
	;
	goto L29
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	if l0 == v64 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v85 = int32(4341168)
	goto L11
L31:
	;
	goto L32
L32:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[620]))
	if l0 == v68 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v85 = int32(4341176)
	goto L11
L34:
	;
	goto L35
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _consts[604]))
	if v72 == int32(0) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _consts[621]))
	v79 = int32(0)
	v81 = F_hash_search(m, v76, v7+int32(28), v79, v79)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	if v81 == int32(0) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v85 = v81
	goto L11
L39:
	;
	goto L10
L40:
	;
	v94 = m.G0
	v96 = v94 - int32(16)
	m.G0 = v96
	*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = l0
	v100 = *(*int32)(unsafe.Add(mBase, _consts[613]))
	if l0 == v100 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v151 = int32(16)
	m.G0 = v96 + v151
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v149
	F_errmsg_internal(m, int32(451733), v7+v151)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L70
	}
L42:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
	v149 = v148
	goto L41
L43:
	;
	v147 = int32(4341120)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[614]))
	if l0 == v104 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v147 = int32(4341128)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _consts[615]))
	if l0 == v108 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = int32(4341136)
	goto L42
L50:
	;
	goto L51
L51:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[616]))
	if l0 == v112 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v147 = int32(4341144)
	goto L42
L53:
	;
	goto L54
L54:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _consts[617]))
	if l0 == v116 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v147 = int32(4341152)
	goto L42
L56:
	;
	goto L57
L57:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _consts[618]))
	if l0 == v120 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v147 = int32(4341160)
	goto L42
L59:
	;
	goto L60
L60:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	if l0 == v124 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v147 = int32(4341168)
	goto L42
L62:
	;
	goto L63
L63:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[620]))
	if l0 == v128 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v147 = int32(4341176)
	goto L42
L65:
	;
	goto L66
L66:
	;
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _consts[604]))
	if v133 == v131 {
		v149 = v131
		goto L41
	} else {
		goto L67
	}
L67:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[621]))
	v140 = int32(0)
	v142 = F_hash_search(m, v137, v96+int32(12), v140, v140)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L6
	} else {
		goto L68
	}
L68:
	;
	if v142 == int32(0) {
		v149 = v131
		goto L41
	} else {
		goto L69
	}
L69:
	;
	v147 = v142
	goto L42
L70:
	;
	F_errfinish(m, int32(464766), int32(5659), int32(392353))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L71
	}
L71:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_MarkBufferDirty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v61 int32
	_ = v61
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 < int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L22
	} else {
		goto L44
	}
L4:
	;
	m.G0 = v8 + int32(32)
	return
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v21 = v14 + (l0^int32(-1))<<(uint(int32(6))%32) + int32(24)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v22&int32(8388608) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v42 = v37 + l0<<(uint(int32(6))%32) - int32(40)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v44 = v43
	goto L12
L8:
	;
	goto L4
L9:
	;
	v27 = int32(4323592)
	v29 = *(*int64)(unsafe.Add(mBase, _consts[349]))
	*(*int64)(unsafe.Add(mBase, _consts[349])) = v29 + int64(1)
	goto L11
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v22 | int32(8388608)
	goto L8
L12:
	;
	if v44&int32(4194304) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v104&int32(8388608) != 0 {
		goto L4
	} else {
		goto L42
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = int32(428699)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(6287)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(464766)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = int64(0)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v61&int32(4194304) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v104 = v44
	goto L16
L16:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v112 = base.B2i32(v104 == v111)
	if v104 == v111 {
		goto L36
	} else {
		goto L37
	}
L17:
	;
	goto L20
L18:
	;
	v76 = v61
	goto L19
L19:
	;
	v84 = int32(4047244)
	v85 = *(*int32)(unsafe.Add(mBase, _consts[605]))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(8))+8))
	if v87 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	F_perform_spin_delay(m, v8+int32(8))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v76 = v73
	goto L19
L22:
	;
	return
L23:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	if v73&int32(4194304) != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v104 = v76
	goto L16
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _consts[605])) = v102
	goto L26
L28:
	;
	if int32(999) < v85 {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v85 < int32(11) {
		goto L26
	} else {
		goto L35
	}
L31:
	;
	v92 = int32(900)
	if v92 <= v85 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v95 = v92
	goto L34
L33:
	;
	v95 = v85
	goto L34
L34:
	;
	v102 = v95 + int32(100)
	goto L27
L35:
	;
	v102 = v85 - int32(1)
	goto L27
L36:
	;
	v113 = v104 | int32(276824064)
	goto L38
L37:
	;
	v113 = v111
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = v113
	if v112 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v44 = v111
	goto L12
L40:
	;
	goto L41
L41:
	;
	goto L13
L42:
	;
	v119 = int32(4323560)
	v121 = *(*int64)(unsafe.Add(mBase, _consts[345]))
	*(*int64)(unsafe.Add(mBase, _consts[345])) = v121 + int64(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _consts[51])))
	if v126 != int32(1) {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v129 = int32(4420000)
	v131 = *(*int32)(unsafe.Add(mBase, _consts[50]))
	v133 = *(*int32)(unsafe.Add(mBase, _consts[624]))
	*(*int32)(unsafe.Add(mBase, _consts[50])) = v131 + v133
	goto L4
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
	F_errmsg_internal(m, int32(458380), v8)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L22
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(464766), int32(2954), int32(8141))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReadBufferWithoutRelcache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int64
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
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
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
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
	var v161 int32
	_ = v161
	var v163 int64
	_ = v163
	var v164 int64
	_ = v164
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v182 int64
	_ = v182
	var v188 int64
	_ = v188
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+16)) = v17
	v20 = l3 - int32(1)
	v24 = F_smgropen(m, v13+int32(16), int32(-1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int32(0)
	} else {
		if l2 == int32(-1) {
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v30
			*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v30
			v34 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v13)+32)) = v34
			*(*int64)(unsafe.Add(mBase, uint32(v13))) = v34
			if base.Ui32(v20) < base.Ui32(int32(2)) {
				v42 = int32(9)
			} else {
				v42 = int32(1)
			}
			v43 = F_ExtendBufferedRel(m, v13, l1, l4, v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				v243 = v43
				m.G0 = v13 + int32(112)
				return v243
			}
		} else {
			if base.Ui32(v20) <= base.Ui32(int32(1)) {
				v47 = F_IOContextForStrategy(m, l4)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v51 = *(*int32)(unsafe.Add(mBase, _consts[10]))
					F_ResourceOwnerEnlarge(m, v51)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_ReservePrivateRefCountEntry(m)
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return int32(0)
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v56
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v58
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v60
							v66 = F_BufTableHashCode(m, v13+int32(32))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return int32(0)
							} else {
								v69 = *(*int32)(unsafe.Add(mBase, _consts[7]))
								v76 = v69 + v66&int32(127)<<(uint(int32(7))%32) + int32(6912)
								v78 = F_LWLockAcquire(m, v76, int32(1))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return int32(0)
								} else {
									v82 = F_BufTableLookup(m, v13+int32(32), v66)
									mBase = m.M
									v83 = m.ExcPending
									if v83 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v82 {
											v87 = *(*int32)(unsafe.Add(mBase, _consts[16]))
											v90 = v87 + v82<<(uint(int32(6))%32)
											v91 = F_PinBuffer(m, v90, l4)
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												F_LWLockRelease(m, v76)
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													if v91 != 0 {
														v159 = v90
														v161 = int32(4323544)
														v163 = *(*int64)(unsafe.Add(mBase, _consts[343]))
														v164 = int64(1)
														*(*int64)(unsafe.Add(mBase, _consts[343])) = v163 + v164
														v167 = int32(1)
														v175 = v47 << (uint(int32(6)) % 32)
														v182 = *(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[606])))
														*(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[606]))) = v182 + v164
														v188 = *(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[607])))
														*(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[607]))) = v188
														F_pgstat_count_backend_io_op(m, int32(0), v47, int32(2), v167, int64(0))
														mBase = m.M
														*(*uint8)(unsafe.Add(mBase, _consts[89])) = uint8(v167)
														*(*uint8)(unsafe.Add(mBase, _consts[608])) = uint8(v167)
														v200 = int32(*(*uint8)(unsafe.Add(mBase, _consts[51])))
														if v200 != int32(1) {
															v211 = v159
															v212 = v167
														} else {
															v203 = int32(4420000)
															v205 = *(*int32)(unsafe.Add(mBase, _consts[50]))
															v207 = *(*int32)(unsafe.Add(mBase, _consts[609]))
															*(*int32)(unsafe.Add(mBase, _consts[50])) = v205 + v207
															v211 = v159
															v212 = v167
														}
													} else {
														v211 = v90
														v212 = int32(0)
													}
													v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
													v215 = v213 + int32(1)
													F_ZeroAndLockBuffer(m, v215, l3, v212)
													mBase = m.M
													v217 = m.ExcPending
													if v217 != 0 {
														return int32(0)
													} else {
														v243 = v215
														m.G0 = v13 + int32(112)
														return v243
													}
												}
											}
										} else {
											F_LWLockRelease(m, v76)
											mBase = m.M
											v96 = m.ExcPending
											if v96 != 0 {
												return int32(0)
											} else {
												v97 = F_GetVictimBuffer(m, l4, v47)
												mBase = m.M
												v98 = m.ExcPending
												if v98 != 0 {
													return int32(0)
												} else {
													v100 = *(*int32)(unsafe.Add(mBase, _consts[16]))
													v102 = F_LWLockAcquire(m, v76, int32(0))
													mBase = m.M
													v103 = m.ExcPending
													if v103 != 0 {
														return int32(0)
													} else {
														v106 = v100 + v97<<(uint(int32(6))%32)
														v108 = v106 + int32(-64)
														v113 = *(*int32)(unsafe.Add(mBase, uint32(v106-int32(44))))
														v114 = F_BufTableInsert(m, v13+int32(32), v66, v113)
														mBase = m.M
														v115 = m.ExcPending
														if v115 != 0 {
															return int32(0)
														} else {
															if v114 < int32(0) {
																v118 = F_LockBufHdr(m, v108)
																mBase = m.M
																v119 = m.ExcPending
																if v119 != 0 {
																	return int32(0)
																} else {
																	v120 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v120
																	v122 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
																	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v122
																	v124 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
																	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v124
																	v130 = int32(-2113667072)
																	if l5 != 0 {
																		v133 = v130
																	} else {
																		v133 = int32(33816576)
																	}
																	if l1 == int32(3) {
																		v136 = v130
																	} else {
																		v136 = v133
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v106-int32(40)))) = v118&int32(-38010881) | v136
																	F_LWLockRelease(m, v76)
																	mBase = m.M
																	v140 = m.ExcPending
																	if v140 != 0 {
																		return int32(0)
																	} else {
																		v211 = v108
																		v212 = int32(0)
																		v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
																		v215 = v213 + int32(1)
																		F_ZeroAndLockBuffer(m, v215, l3, v212)
																		mBase = m.M
																		v217 = m.ExcPending
																		if v217 != 0 {
																			return int32(0)
																		} else {
																			v243 = v215
																			m.G0 = v13 + int32(112)
																			return v243
																		}
																	}
																}
															} else {
																F_UnpinBuffer(m, v108)
																mBase = m.M
																v143 = m.ExcPending
																if v143 != 0 {
																	return int32(0)
																} else {
																	F_StrategyFreeBuffer(m, v108)
																	mBase = m.M
																	v145 = m.ExcPending
																	if v145 != 0 {
																		return int32(0)
																	} else {
																		v148 = *(*int32)(unsafe.Add(mBase, _consts[16]))
																		v151 = v148 + v114<<(uint(int32(6))%32)
																		v152 = F_PinBuffer(m, v151, l4)
																		mBase = m.M
																		v153 = m.ExcPending
																		if v153 != 0 {
																			return int32(0)
																		} else {
																			F_LWLockRelease(m, v76)
																			mBase = m.M
																			v155 = m.ExcPending
																			if v155 != 0 {
																				return int32(0)
																			} else {
																				if v152 == int32(0) {
																					v211 = v151
																					v212 = int32(0)
																				} else {
																					v159 = v151
																					v161 = int32(4323544)
																					v163 = *(*int64)(unsafe.Add(mBase, _consts[343]))
																					v164 = int64(1)
																					*(*int64)(unsafe.Add(mBase, _consts[343])) = v163 + v164
																					v167 = int32(1)
																					v175 = v47 << (uint(int32(6)) % 32)
																					v182 = *(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[606])))
																					*(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[606]))) = v182 + v164
																					v188 = *(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[607])))
																					*(*int64)(unsafe.Add(mBase, uint32(v175)+uint32(_consts[607]))) = v188
																					F_pgstat_count_backend_io_op(m, int32(0), v47, int32(2), v167, int64(0))
																					mBase = m.M
																					*(*uint8)(unsafe.Add(mBase, _consts[89])) = uint8(v167)
																					*(*uint8)(unsafe.Add(mBase, _consts[608])) = uint8(v167)
																					v200 = int32(*(*uint8)(unsafe.Add(mBase, _consts[51])))
																					if v200 != int32(1) {
																						v211 = v159
																						v212 = v167
																					} else {
																						v203 = int32(4420000)
																						v205 = *(*int32)(unsafe.Add(mBase, _consts[50]))
																						v207 = *(*int32)(unsafe.Add(mBase, _consts[609]))
																						*(*int32)(unsafe.Add(mBase, _consts[50])) = v205 + v207
																						v211 = v159
																						v212 = v167
																					}
																				}
																				v213 = *(*int32)(unsafe.Add(mBase, uint32(v211)+20))
																				v215 = v213 + int32(1)
																				F_ZeroAndLockBuffer(m, v215, l3, v212)
																				mBase = m.M
																				v217 = m.ExcPending
																				if v217 != 0 {
																					return int32(0)
																				} else {
																					v243 = v215
																					m.G0 = v13 + int32(112)
																					return v243
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = l1
				if l5 != 0 {
					v222 = int32(112)
				} else {
					v222 = int32(117)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v13)+40)) = uint8(v222)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v24
				if l3 == int32(3) {
					v235 = int32(9)
				} else {
					v235 = int32(8)
				}
				v236 = F_StartReadBuffer(m, v13+int32(32), v13+int32(28), l2, v235)
				mBase = m.M
				v237 = m.ExcPending
				if v237 != 0 {
					return int32(0)
				} else {
					if v236 != 0 {
						F_WaitReadBuffers(m, v13+int32(32))
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return int32(0)
						} else {
							v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
							v243 = v242
							m.G0 = v13 + int32(112)
							return v243
						}
					} else {
						v242 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
						v243 = v242
						m.G0 = v13 + int32(112)
						return v243
					}
				}
			}
		}
	}
}
func F_ReleaseBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	if l0 != 0 {
		if l0 < int32(0) {
			F_UnpinLocalBuffer(m, l0)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return
			} else {
				m.G0 = v5 + int32(16)
				return
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[10]))
			v14 = *(*int32)(unsafe.Add(mBase, _consts[16]))
			v17 = v14 + l0<<(uint(int32(6))%32)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(44))))
			F_ResourceOwnerForget(m, v12, v20+int32(1), int32(1577392))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_UnpinBufferNoOwner(m, v17+int32(-64))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(0)
			F_errmsg_internal(m, int32(458380), v5)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_errfinish(m, int32(464766), int32(5369), int32(213362))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_UnlockReleaseBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	if int32(0) <= l0 {
		v5 = *(*int32)(unsafe.Add(mBase, _consts[16]))
		F_LWLockRelease(m, v5+l0<<(uint(int32(6))%32)-int32(16))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_ReleaseBuffer(m, l0)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		F_ReleaseBuffer(m, l0)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			return
		}
	}
}
func F_buffer_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v6) < base.Ui32(v5) {
		v8 = int32(1)
	} else {
		v8 = int32(-1)
	}
	return v8
}
