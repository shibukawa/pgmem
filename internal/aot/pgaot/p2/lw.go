package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_LWLockNewTrancheId(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
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
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v210 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_LWLockNewTrancheId[0]))), uint32(v210))
	F_errstart_cold(m, int32(21), v210)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L53
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L10
	} else {
		goto L48
	}
L3:
	;
	v10 = F_strlen(m, l0)
	mBase = m.M
	if base.Ui32(int32(64)) <= base.Ui32(v10) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L10
	} else {
		goto L44
	}
L6:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
	v17 = base.AtomicRmwXchg32(m, v14, int32(_a_F_LWLockNewTrancheId_0), int32(1))
	if v17 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_s_lock(m, v14+int32(_a_F_LWLockNewTrancheId_0), int32(_a_F_LWLockNewTrancheId_1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_LWLockNewTrancheId[2])))
	if int32(256) <= v27 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L9
L12:
	;
	v31 = v27 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_LWLockNewTrancheId[2]))) = v31
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[3])) = v31
	v36 = v27 * int32(68)
	v37 = v26 + v36
	goto L16
L13:
	;
	v157 = int32(_a_F_LWLockNewTrancheId_2)
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v158+v36)+64)) = int32(-1)
	v163 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockNewTrancheId[1]))
	v164 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v163)+uint32(_c_F_LWLockNewTrancheId[0]))), uint32(v164))
	m.G0 = v8 + int32(32)
	return v27 + int32(100)
L14:
	;
	v154 = F_strlen(m, v143)
	mBase = m.M
	goto L13
L16:
	;
	goto L17
L17:
	;
	v44 = int32(63)
	if (v37^l0)&int32(3) != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v147)
	goto L14
L19:
	;
	v128 = v123
	v129 = v124
	v130 = v125
	goto L40
L20:
	;
	if v118 == int32(0) {
		v143 = v116
		v144 = v117
		goto L18
	} else {
		goto L39
	}
L21:
	;
	v116 = l0
	v117 = v37
	v118 = v44
	goto L20
L22:
	;
	goto L23
L23:
	;
	v48 = int32(0)
	if base.B2i32(l0&int32(3) == v48)|int32(0) == v48 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if v84 == int32(0) {
		v143 = v81
		v144 = v82
		goto L18
	} else {
		goto L33
	}
L25:
	;
	v60 = l0
	v61 = v37
	v62 = v44
	goto L28
L26:
	;
	goto L27
L27:
	;
	v81 = l0
	v82 = v37
	v83 = v44
	v84 = int32(1)
	goto L24
L28:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v64)
	if v64 == int32(0) {
		v123 = v60
		v124 = v61
		v125 = v62
		goto L19
	} else {
		goto L30
	}
L29:
	;
	v81 = v75
	v82 = v69
	v83 = v71
	v84 = v73
	goto L24
L30:
	;
	v68 = int32(1)
	v69 = v61 + v68
	v71 = v62 - v68
	v72 = int32(0)
	v73 = base.B2i32(v71 != v72)
	v75 = v60 + v68
	if v75&int32(3) == v72 {
		v81 = v75
		v82 = v69
		v83 = v71
		v84 = v73
		goto L24
	} else {
		goto L31
	}
L31:
	;
	if v71 != 0 {
		v60 = v75
		v61 = v69
		v62 = v71
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if base.B2i32(v87 == int32(0))|base.B2i32(base.Ui32(v83) < base.Ui32(int32(4))) != 0 {
		v116 = v81
		v117 = v82
		v118 = v83
		goto L20
	} else {
		goto L34
	}
L34:
	;
	v94 = v81
	v95 = v82
	v96 = v83
	goto L35
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = int32(-2139062144)
	if (int32(16843008)-v99|v99)&v102 != v102 {
		v123 = v94
		v124 = v95
		v125 = v96
		goto L19
	} else {
		goto L37
	}
L36:
	;
	v116 = v110
	v117 = v108
	v118 = v112
	goto L20
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
	v107 = int32(4)
	v108 = v95 + v107
	v110 = v94 + v107
	v112 = v96 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v94 = v110
		v95 = v108
		v96 = v112
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L19
L40:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v132)
	if v132 == int32(0) {
		v143 = v128
		v144 = v129
		goto L18
	} else {
		goto L42
	}
L41:
	;
	v143 = v139
	v144 = v137
	goto L18
L42:
	;
	v136 = int32(1)
	v137 = v129 + v136
	v139 = v128 + v136
	v141 = v130 - v136
	if v141 != 0 {
		v128 = v139
		v129 = v137
		v130 = v141
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	F_errmsg(m, int32(_a_F_LWLockNewTrancheId_3), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_LWLockNewTrancheId_4), int32(569), int32(_a_F_LWLockNewTrancheId_5))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(34103428))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L10
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_LWLockNewTrancheId_6), int32(0))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(63)
	v203 = F_errdetail(m, int32(_a_F_LWLockNewTrancheId_7), v8)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_LWLockNewTrancheId_4), int32(576), int32(_a_F_LWLockNewTrancheId_5))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	F_errmsg(m, int32(_a_F_LWLockNewTrancheId_8), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L10
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(256)
	v226 = F_errdetail(m, int32(_a_F_LWLockNewTrancheId_9), v8+int32(16))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L10
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(_a_F_LWLockNewTrancheId_4), int32(587), int32(_a_F_LWLockNewTrancheId_5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L10
	} else {
		goto L56
	}
L56:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
