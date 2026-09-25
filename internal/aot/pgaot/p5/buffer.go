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
	v5 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[0]))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v3 + (v5 - v6)
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[1]))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v10 + (v12 - v13)
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v19 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[2]))
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v17 + (v19 - v20)
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[3]))
	v27 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v24 + (v26 - v27)
	v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[4]))
	v34 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+32)) = v31 + (v33 - v34)
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v40 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[5]))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+40)) = v38 + (v40 - v41)
	v45 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[6]))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+48)) = v45 + (v47 - v48)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v54 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[7]))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = v52 + (v54 - v55)
	v59 = *(*int64)(unsafe.Add(mBase, uint32(l0)+64))
	v61 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[8]))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(l1)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+64)) = v59 + (v61 - v62)
	v66 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v68 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[9]))
	v69 = *(*int64)(unsafe.Add(mBase, uint32(l1)+72))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+72)) = v66 + (v68 - v69)
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	v75 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[10]))
	v76 = *(*int64)(unsafe.Add(mBase, uint32(l1)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v73 + (v75 - v76)
	v80 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v82 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[11]))
	v83 = *(*int64)(unsafe.Add(mBase, uint32(l1)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v80 + (v82 - v83)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
	v89 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[12]))
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)+96))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v87 + (v89 - v90)
	v94 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
	v96 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[13]))
	v97 = *(*int64)(unsafe.Add(mBase, uint32(l1)+104))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v94 + (v96 - v97)
	v101 = *(*int64)(unsafe.Add(mBase, uint32(l0)+112))
	v103 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[14]))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(l1)+112))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+112)) = v101 + (v103 - v104)
	v108 = *(*int64)(unsafe.Add(mBase, uint32(l0)+120))
	v110 = *(*int64)(unsafe.Add(mBase, _c_F_BufferUsageAccumDiff[15]))
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
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[0]))
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
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
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
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v26+v14)))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v28
	F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_2), int32(_a_F_CheckBufferIsPinnedOnce_3))
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
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_4)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[2]))
	if l0 == v44 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_6)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[3]))
	if l0 == v48 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_7)
	goto L11
L19:
	;
	goto L20
L20:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[4]))
	if l0 == v52 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_8)
	goto L11
L22:
	;
	goto L23
L23:
	;
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[5]))
	if l0 == v56 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_9)
	goto L11
L25:
	;
	goto L26
L26:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[6]))
	if l0 == v60 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_10)
	goto L11
L28:
	;
	goto L29
L29:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[7]))
	if l0 == v64 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_11)
	goto L11
L31:
	;
	goto L32
L32:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[8]))
	if l0 == v68 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v85 = int32(_a_F_CheckBufferIsPinnedOnce_12)
	goto L11
L34:
	;
	goto L35
L35:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[9]))
	if v72 == int32(0) {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[10]))
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
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[1]))
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
	F_errmsg_internal(m, int32(_a_F_CheckBufferIsPinnedOnce_0), v7+v151)
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
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_4)
	goto L42
L44:
	;
	goto L45
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[2]))
	if l0 == v104 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_6)
	goto L42
L47:
	;
	goto L48
L48:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[3]))
	if l0 == v108 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_7)
	goto L42
L50:
	;
	goto L51
L51:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[4]))
	if l0 == v112 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_8)
	goto L42
L53:
	;
	goto L54
L54:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[5]))
	if l0 == v116 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_9)
	goto L42
L56:
	;
	goto L57
L57:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[6]))
	if l0 == v120 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_10)
	goto L42
L59:
	;
	goto L60
L60:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[7]))
	if l0 == v124 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_11)
	goto L42
L62:
	;
	goto L63
L63:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[8]))
	if l0 == v128 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v147 = int32(_a_F_CheckBufferIsPinnedOnce_12)
	goto L42
L65:
	;
	goto L66
L66:
	;
	v131 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[9]))
	if v133 == v131 {
		v149 = v131
		goto L41
	} else {
		goto L67
	}
L67:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_CheckBufferIsPinnedOnce[10]))
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
	F_errfinish(m, int32(_a_F_CheckBufferIsPinnedOnce_1), int32(_a_F_CheckBufferIsPinnedOnce_5), int32(_a_F_CheckBufferIsPinnedOnce_3))
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
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
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
	v136 = m.ExcPending
	if v136 != 0 {
		goto L22
	} else {
		goto L41
	}
L4:
	;
	m.G0 = v7 + int32(32)
	return
L5:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[0]))
	v18 = v13 + (l0^int32(-1))<<(uint(int32(6))%32)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v19&int32(_a_F_MarkBufferDirty_0) == int32(0) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[1]))
	v39 = v34 + l0<<(uint(int32(6))%32) - int32(40)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = v40
	goto L12
L8:
	;
	goto L4
L9:
	;
	v24 = int32(_a_F_MarkBufferDirty_1)
	v26 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[2]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[2])) = v26 + int64(1)
	goto L11
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v19 | int32(_a_F_MarkBufferDirty_0)
	goto L8
L12:
	;
	if v41&int32(_a_F_MarkBufferDirty_2) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	if v98&int32(_a_F_MarkBufferDirty_0) != 0 {
		goto L4
	} else {
		goto L39
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(_a_F_MarkBufferDirty_3)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(_a_F_MarkBufferDirty_4)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_MarkBufferDirty_5)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v7)+8)) = int64(0)
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v57&int32(_a_F_MarkBufferDirty_2) != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v98 = v41
	goto L16
L16:
	;
	v105 = base.AtomicRmwCmpxchg32(m, v39, int32(0), v98, v98|int32(276824064))
	if v98 != v105 {
		goto L36
	} else {
		goto L37
	}
L17:
	;
	goto L20
L18:
	;
	v71 = v57
	goto L19
L19:
	;
	v78 = int32(_a_F_MarkBufferDirty_6)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[3]))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v7+int32(8))+8))
	if v81 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L20:
	;
	F_perform_spin_delay(m, v7+int32(8))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v71 = v68
	goto L19
L22:
	;
	return
L23:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v68&int32(_a_F_MarkBufferDirty_2) != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	goto L21
L25:
	;
	v98 = v71
	goto L16
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[3])) = v96
	goto L26
L28:
	;
	if int32(999) < v79 {
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	if v79 < int32(11) {
		goto L26
	} else {
		goto L35
	}
L31:
	;
	v86 = int32(900)
	if v86 <= v79 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v86
	goto L34
L33:
	;
	v89 = v79
	goto L34
L34:
	;
	v96 = v89 + int32(100)
	goto L27
L35:
	;
	v96 = v79 - int32(1)
	goto L27
L36:
	;
	v41 = v105
	goto L12
L37:
	;
	goto L38
L38:
	;
	goto L13
L39:
	;
	v109 = int32(_a_F_MarkBufferDirty_7)
	v111 = *(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_MarkBufferDirty[4])) = v111 + int64(1)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_MarkBufferDirty[5])))
	if v116 != int32(1) {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v119 = int32(_a_F_MarkBufferDirty_8)
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[6]))
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_MarkBufferDirty[6])) = v121 + v123
	goto L4
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(0)
	F_errmsg_internal(m, int32(_a_F_MarkBufferDirty_9), v7)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_MarkBufferDirty_5), int32(2954), int32(_a_F_MarkBufferDirty_10))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReadBufferWithoutRelcache(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v183 int64
	_ = v183
	var v187 int64
	_ = v187
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	v12 = m.G0
	v14 = v12 - int32(112)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+16)) = v18
	v21 = l3 - int32(1)
	v25 = F_smgropen(m, v14+int32(16), int32(-1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		return int32(0)
	} else {
		if l2 == int32(-1) {
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v31
			v33 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v14)+32)) = v33
			*(*int64)(unsafe.Add(mBase, uint32(v14))) = v33
			*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v31
			if base.Ui32(v21) < base.Ui32(int32(2)) {
				v43 = int32(9)
			} else {
				v43 = int32(1)
			}
			v44 = F_ExtendBufferedRel(m, v14, l1, l4, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v242 = v44
				m.G0 = v14 + int32(112)
				return v242
			}
		} else {
			if base.Ui32(v21) <= base.Ui32(int32(1)) {
				v48 = F_IOContextForStrategy(m, l4)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v52 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[0]))
					F_ResourceOwnerEnlarge(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						F_ReservePrivateRefCountEntry(m)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v57
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v14)+40)) = v61
							v66 = v14 + int32(32)
							v67 = F_BufTableHashCode(m, v66)
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return int32(0)
							} else {
								v70 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[1]))
								v77 = v70 + v67&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ReadBufferWithoutRelcache_0)
								v79 = F_LWLockAcquire(m, v77, int32(1))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									v81 = F_BufTableLookup(m, v66, v67)
									mBase = m.M
									v82 = m.ExcPending
									if v82 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v81 {
											v86 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[2]))
											v89 = v86 + v81<<(uint(int32(6))%32)
											v90 = F_PinBuffer(m, v89, l4)
											mBase = m.M
											v91 = m.ExcPending
											if v91 != 0 {
												return int32(0)
											} else {
												F_LWLockRelease(m, v77)
												mBase = m.M
												v93 = m.ExcPending
												if v93 != 0 {
													return int32(0)
												} else {
													if v90 != 0 {
														v161 = v89
														v165 = int32(_a_F_ReadBufferWithoutRelcache_1)
														v167 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3]))
														v168 = int64(1)
														*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3])) = v167 + v168
														v171 = int32(1)
														v178 = v48 << (uint(int32(6)) % 32)
														v183 = *(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[4])))
														*(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[4]))) = v183 + v168
														v187 = *(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[5])))
														*(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[5]))) = v187
														F_pgstat_count_backend_io_op(m, int32(0), v48, int32(2), v171, int64(0))
														mBase = m.M
														*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[6])) = uint8(v171)
														*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[7])) = uint8(v171)
														v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[8])))
														if v199 != int32(1) {
															v209 = v161
															v211 = v171
														} else {
															v202 = int32(_a_F_ReadBufferWithoutRelcache_2)
															v204 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9]))
															v206 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[10]))
															*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9])) = v204 + v206
															v209 = v161
															v211 = v171
														}
													} else {
														v209 = v89
														v211 = int32(0)
													}
													v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
													v216 = v214 + int32(1)
													F_ZeroAndLockBuffer(m, v216, l3, v211)
													mBase = m.M
													v218 = m.ExcPending
													if v218 != 0 {
														return int32(0)
													} else {
														v242 = v216
														m.G0 = v14 + int32(112)
														return v242
													}
												}
											}
										} else {
											F_LWLockRelease(m, v77)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												v96 = F_GetVictimBuffer(m, l4, v48)
												mBase = m.M
												v97 = m.ExcPending
												if v97 != 0 {
													return int32(0)
												} else {
													v99 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[2]))
													v101 = F_LWLockAcquire(m, v77, int32(0))
													mBase = m.M
													v102 = m.ExcPending
													if v102 != 0 {
														return int32(0)
													} else {
														v105 = v99 + v96<<(uint(int32(6))%32)
														v107 = v105 + int32(-64)
														v112 = *(*int32)(unsafe.Add(mBase, uint32(v105-int32(44))))
														v113 = F_BufTableInsert(m, v14+int32(32), v67, v112)
														mBase = m.M
														v114 = m.ExcPending
														if v114 != 0 {
															return int32(0)
														} else {
															if v113 < int32(0) {
																v117 = F_LockBufHdr(m, v107)
																mBase = m.M
																v118 = m.ExcPending
																if v118 != 0 {
																	return int32(0)
																} else {
																	v119 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
																	*(*int32)(unsafe.Add(mBase, uint32(v107)+16)) = v119
																	v121 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
																	*(*int64)(unsafe.Add(mBase, uint32(v107)+8)) = v121
																	v123 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
																	*(*int64)(unsafe.Add(mBase, uint32(v107))) = v123
																	v125 = int32(0)
																	v129 = base.AtomicRmwOr32(m, v125, int32(_a_F_ReadBufferWithoutRelcache_3), v125)
																	v134 = int32(-2113667072)
																	if l5 != 0 {
																		v137 = v134
																	} else {
																		v137 = int32(33816576)
																	}
																	if l1 == int32(3) {
																		v140 = v134
																	} else {
																		v140 = v137
																	}
																	*(*int32)(unsafe.Add(mBase, uint32(v105-int32(40)))) = v117&int32(-38010881) | v140
																	F_LWLockRelease(m, v77)
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
																		return int32(0)
																	} else {
																		v209 = v107
																		v211 = v125
																		v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
																		v216 = v214 + int32(1)
																		F_ZeroAndLockBuffer(m, v216, l3, v211)
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			v242 = v216
																			m.G0 = v14 + int32(112)
																			return v242
																		}
																	}
																}
															} else {
																F_UnpinBuffer(m, v107)
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int32(0)
																} else {
																	F_StrategyFreeBuffer(m, v107)
																	mBase = m.M
																	v148 = m.ExcPending
																	if v148 != 0 {
																		return int32(0)
																	} else {
																		v151 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[2]))
																		v154 = v151 + v113<<(uint(int32(6))%32)
																		v155 = F_PinBuffer(m, v154, l4)
																		mBase = m.M
																		v156 = m.ExcPending
																		if v156 != 0 {
																			return int32(0)
																		} else {
																			F_LWLockRelease(m, v77)
																			mBase = m.M
																			v158 = m.ExcPending
																			if v158 != 0 {
																				return int32(0)
																			} else {
																				if v155 == int32(0) {
																					v209 = v154
																					v211 = int32(0)
																				} else {
																					v161 = v154
																					v165 = int32(_a_F_ReadBufferWithoutRelcache_1)
																					v167 = *(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3]))
																					v168 = int64(1)
																					*(*int64)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[3])) = v167 + v168
																					v171 = int32(1)
																					v178 = v48 << (uint(int32(6)) % 32)
																					v183 = *(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[4])))
																					*(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[4]))) = v183 + v168
																					v187 = *(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[5])))
																					*(*int64)(unsafe.Add(mBase, uint32(v178)+uint32(_c_F_ReadBufferWithoutRelcache[5]))) = v187
																					F_pgstat_count_backend_io_op(m, int32(0), v48, int32(2), v171, int64(0))
																					mBase = m.M
																					*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[6])) = uint8(v171)
																					*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[7])) = uint8(v171)
																					v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[8])))
																					if v199 != int32(1) {
																						v209 = v161
																						v211 = v171
																					} else {
																						v202 = int32(_a_F_ReadBufferWithoutRelcache_2)
																						v204 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9]))
																						v206 = *(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[10]))
																						*(*int32)(unsafe.Add(mBase, _c_F_ReadBufferWithoutRelcache[9])) = v204 + v206
																						v209 = v161
																						v211 = v171
																					}
																				}
																				v214 = *(*int32)(unsafe.Add(mBase, uint32(v209)+20))
																				v216 = v214 + int32(1)
																				F_ZeroAndLockBuffer(m, v216, l3, v211)
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					v242 = v216
																					m.G0 = v14 + int32(112)
																					return v242
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
				*(*int32)(unsafe.Add(mBase, uint32(v14)+48)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v14)+44)) = l1
				if l5 != 0 {
					v223 = int32(112)
				} else {
					v223 = int32(117)
				}
				*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)) = uint8(v223)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v25
				v229 = v14 + int32(32)
				if l3 == int32(3) {
					v236 = int32(9)
				} else {
					v236 = int32(8)
				}
				v237 = F_StartReadBuffer(m, v229, v14+int32(28), l2, v236)
				mBase = m.M
				v238 = m.ExcPending
				if v238 != 0 {
					return int32(0)
				} else {
					if v237 != 0 {
						F_WaitReadBuffers(m, v229)
						mBase = m.M
						v240 = m.ExcPending
						if v240 != 0 {
							return int32(0)
						} else {
							v241 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
							v242 = v241
							m.G0 = v14 + int32(112)
							return v242
						}
					} else {
						v241 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
						v242 = v241
						m.G0 = v14 + int32(112)
						return v242
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
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseBuffer[0]))
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReleaseBuffer[1]))
			v17 = v14 + l0<<(uint(int32(6))%32)
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17-int32(44))))
			F_ResourceOwnerForget(m, v12, v20+int32(1), int32(_a_F_ReleaseBuffer_0))
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
			F_errmsg_internal(m, int32(_a_F_ReleaseBuffer_1), v5)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ReleaseBuffer_2), int32(_a_F_ReleaseBuffer_3), int32(_a_F_ReleaseBuffer_4))
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
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockReleaseBuffer[0]))
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
