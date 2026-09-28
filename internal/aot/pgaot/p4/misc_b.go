package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BIG5toCNS(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	if base.Ui32(l0) <= base.Ui32(int32(_a_F_BIG5toCNS_0)) {
		goto L13
	} else {
		goto L14
	}
L1:
	;
	return v322 & int32(_a_F_BIG5toCNS_1)
L2:
	;
	v317 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v317)
	v322 = int32(63)
	goto L1
L3:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v311)
	v322 = v310 | int32(-32640)
	goto L1
L4:
	;
	v305 = int32(246)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v305)
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v304)+2)))
	v322 = v307 | int32(-32640)
	goto L1
L5:
	;
	v304 = int32(_a_F_BIG5toCNS_2)
	goto L4
L6:
	;
	v304 = int32(_a_F_BIG5toCNS_3)
	goto L4
L7:
	;
	v304 = int32(_a_F_BIG5toCNS_4)
	goto L4
L8:
	;
	v304 = int32(_a_F_BIG5toCNS_5)
	goto L4
L9:
	;
	v304 = int32(_a_F_BIG5toCNS_6)
	goto L4
L10:
	;
	v304 = int32(_a_F_BIG5toCNS_7)
	goto L4
L11:
	;
	v172 = int32(46)
	v174 = int32(23)
	v178 = int32(0)
	goto L62
L12:
	;
	v310 = v159
	v311 = int32(149)
	goto L3
L13:
	;
	switch l0 - int32(_a_F_BIG5toCNS_8) {
	case 0:
		v147 = int32(_a_F_BIG5toCNS_9)
		goto L16
	case 1, 3:
		goto L20
	case 2:
		goto L19
	case 4:
		goto L18
	default:
		goto L21
	}
L14:
	;
	goto L15
L15:
	;
	switch l0 - int32(_a_F_BIG5toCNS_10) {
	case 0:
		v304 = int32(_a_F_BIG5toCNS_11)
		goto L4
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	case 5:
		goto L6
	case 6:
		goto L5
	default:
		goto L58
	}
L16:
	;
	v148 = int32(247)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v148)
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+2)))
	v322 = v150 | int32(-32640)
	goto L1
L17:
	;
	v147 = int32(_a_F_BIG5toCNS_12)
	goto L16
L18:
	;
	v147 = int32(_a_F_BIG5toCNS_13)
	goto L16
L19:
	;
	v147 = int32(_a_F_BIG5toCNS_14)
	goto L16
L20:
	;
	v21 = int32(23)
	v23 = int32(11)
	v27 = int32(0)
	goto L25
L21:
	;
	if l0 == int32(_a_F_BIG5toCNS_15) {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	if v143 != 0 {
		v159 = v143
		goto L12
	} else {
		goto L57
	}
L24:
	;
	v143 = v141 & int32(_a_F_BIG5toCNS_1)
	goto L23
L25:
	;
	v29 = v23 << (uint(int32(2)) % 32)
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_BIG5toCNS[0]))))
	v32 = base.B2i32(base.Ui32(l0) < base.Ui32(v31))
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v141 = int32(0)
	goto L24
L27:
	;
	goto L26
L28:
	;
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L50
	} else {
		goto L51
	}
L29:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_BIG5toCNS[1]))))
	if base.Ui32(v33) <= base.Ui32(l0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+uint32(_c_F_BIG5toCNS[2]))))
	if v35 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v40 = l0 - v31&int32(_a_F_BIG5toCNS_16)
	if base.Ui32(int32(_a_F_BIG5toCNS_17)) <= base.Ui32(l0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v43 = int32(255)
	v44 = l0 & v43
	v46 = v31 & v43
	v56 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v46))
	if base.Ui32(int32(160)) < base.Ui32(v46) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v83 = int32(255)
	v84 = v35 & v83
	if base.Ui32(int32(160)) < base.Ui32(v84) {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v57 = int32(0)
	goto L37
L36:
	;
	v57 = int32(-34)
	goto L37
L37:
	;
	if base.Ui32(int32(160)) < base.Ui32(v46) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v60 = int32(34)
	goto L40
L39:
	;
	v60 = int32(0)
	goto L40
L40:
	;
	if base.Ui32(int32(160)) < base.Ui32(v44) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v63 = v57
	goto L43
L42:
	;
	v63 = v60
	goto L43
L43:
	;
	v68 = int32(33)
	v69 = v44 - v46 + v40>>(uint(int32(8))%32)*int32(157) + v63 + v35&int32(255) - v68
	v70 = int32(94)
	v71 = base.I32_div_s(v69, v70)
	v141 = v69 - v71*v70 + v35&int32(_a_F_BIG5toCNS_16) + v71<<(uint(int32(8))%32) + v68
	goto L24
L44:
	;
	v100 = int32(_a_F_BIG5toCNS_18)
	goto L46
L45:
	;
	v100 = int32(_a_F_BIG5toCNS_19)
	goto L46
L46:
	;
	v101 = v84 + (l0&v83 - v31&v83 + int32(base.Ui32(v40)>>(uint(int32(8))%32))*int32(94)) + v100
	v103 = int32(157)
	v104 = base.I32_div_s(base.I32_extend16_s(v101), v103)
	v107 = v101 - v104*v103
	if int32(62) < base.I32_extend16_s(v107) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v119 = int32(98)
	goto L49
L48:
	;
	v119 = int32(64)
	goto L49
L49:
	;
	v141 = v107 + v35&int32(_a_F_BIG5toCNS_16) + v104<<(uint(int32(8))%32) + v119
	goto L24
L50:
	;
	v123 = v27
	goto L52
L51:
	;
	v123 = v23 + int32(1)
	goto L52
L52:
	;
	if base.Ui32(l0) < base.Ui32(v31) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v126 = v23 - int32(1)
	goto L55
L54:
	;
	v126 = v21
	goto L55
L55:
	;
	if v123 <= v126 {
		v21 = v126
		v23 = (v123 + v126) >> (uint(int32(1)) % 32)
		v27 = v123
		goto L25
	} else {
		goto L56
	}
L56:
	;
	goto L27
L57:
	;
	goto L2
L58:
	;
	if l0 != int32(_a_F_BIG5toCNS_20) {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v159 = int32(_a_F_BIG5toCNS_21)
	goto L12
L60:
	;
	if v294 == int32(0) {
		goto L2
	} else {
		goto L94
	}
L61:
	;
	v294 = v292 & int32(_a_F_BIG5toCNS_1)
	goto L60
L62:
	;
	v180 = v174 << (uint(int32(2)) % 32)
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_BIG5toCNS[3]))))
	v183 = base.B2i32(base.Ui32(l0) < base.Ui32(v182))
	if base.Ui32(l0) < base.Ui32(v182) {
		goto L65
	} else {
		goto L66
	}
L63:
	;
	v292 = int32(0)
	goto L61
L64:
	;
	goto L63
L65:
	;
	if base.Ui32(l0) < base.Ui32(v182) {
		goto L87
	} else {
		goto L88
	}
L66:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_BIG5toCNS[4]))))
	if base.Ui32(v184) <= base.Ui32(l0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_BIG5toCNS[5]))))
	if v186 == int32(0) {
		goto L64
	} else {
		goto L68
	}
L68:
	;
	v191 = l0 - v182&int32(_a_F_BIG5toCNS_16)
	if base.Ui32(int32(_a_F_BIG5toCNS_17)) <= base.Ui32(l0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v194 = int32(255)
	v195 = l0 & v194
	v197 = v182 & v194
	v207 = base.B2i32(base.Ui32(int32(160)) < base.Ui32(v197))
	if base.Ui32(int32(160)) < base.Ui32(v197) {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v234 = int32(255)
	v235 = v186 & v234
	if base.Ui32(int32(160)) < base.Ui32(v235) {
		goto L81
	} else {
		goto L82
	}
L72:
	;
	v208 = int32(0)
	goto L74
L73:
	;
	v208 = int32(-34)
	goto L74
L74:
	;
	if base.Ui32(int32(160)) < base.Ui32(v197) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v211 = int32(34)
	goto L77
L76:
	;
	v211 = int32(0)
	goto L77
L77:
	;
	if base.Ui32(int32(160)) < base.Ui32(v195) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v214 = v208
	goto L80
L79:
	;
	v214 = v211
	goto L80
L80:
	;
	v219 = int32(33)
	v220 = v195 - v197 + v191>>(uint(int32(8))%32)*int32(157) + v214 + v186&int32(255) - v219
	v221 = int32(94)
	v222 = base.I32_div_s(v220, v221)
	v292 = v220 - v222*v221 + v186&int32(_a_F_BIG5toCNS_16) + v222<<(uint(int32(8))%32) + v219
	goto L61
L81:
	;
	v251 = int32(_a_F_BIG5toCNS_18)
	goto L83
L82:
	;
	v251 = int32(_a_F_BIG5toCNS_19)
	goto L83
L83:
	;
	v252 = v235 + (l0&v234 - v182&v234 + int32(base.Ui32(v191)>>(uint(int32(8))%32))*int32(94)) + v251
	v254 = int32(157)
	v255 = base.I32_div_s(base.I32_extend16_s(v252), v254)
	v258 = v252 - v255*v254
	if int32(62) < base.I32_extend16_s(v258) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v270 = int32(98)
	goto L86
L85:
	;
	v270 = int32(64)
	goto L86
L86:
	;
	v292 = v258 + v186&int32(_a_F_BIG5toCNS_16) + v255<<(uint(int32(8))%32) + v270
	goto L61
L87:
	;
	v274 = v178
	goto L89
L88:
	;
	v274 = v174 + int32(1)
	goto L89
L89:
	;
	if base.Ui32(l0) < base.Ui32(v182) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v277 = v174 - int32(1)
	goto L92
L91:
	;
	v277 = v172
	goto L92
L92:
	;
	if v274 <= v277 {
		v172 = v277
		v174 = (v274 + v277) >> (uint(int32(1)) % 32)
		v178 = v274
		goto L62
	} else {
		goto L93
	}
L93:
	;
	goto L64
L94:
	;
	v310 = v294
	v311 = int32(150)
	goto L3
}
func F_BarrierArriveAndDetach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_BarrierArriveAndDetach_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = v11 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v13 != v15 {
				v17 = int32(0)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
				return base.B2i32(v13 == v17)
			} else {
				v23 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 + int32(1)
				atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v23))
				F_ConditionVariableBroadcast(m, l0+int32(24))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v13 == int32(0))
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v13 = v11 - int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v13 != v15 {
			v17 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v17))
			return base.B2i32(v13 == v17)
		} else {
			v23 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v25 + int32(1)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v23))
			F_ConditionVariableBroadcast(m, l0+int32(24))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				return base.B2i32(v13 == int32(0))
			}
		}
	}
}
func F_BarrierAttach(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_BarrierAttach_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 + int32(1)
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v16))
			return v15
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 + int32(1)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v16 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v16))
		return v15
	}
}
func F_BarrierPhase(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	return v2
}
func F_BasicOpenFile(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_BasicOpenFile[0]))
	v5 = F_BasicOpenFilePerm(m, l0, l1, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_BogusGetChunkContext(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14206(m, l0, int32(_a_F_BogusGetChunkContext_0), int32(327), int32(_a_F_BogusGetChunkContext_1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_BogusRealloc(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v16 int64
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(l0-int32(8))))
		*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v16
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
		F_errmsg_internal(m, int32(_a_F_BogusRealloc_0), v6)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			F_errfinish(m, int32(_a_F_BogusRealloc_1), int32(319), int32(_a_F_BogusRealloc_2))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_BootStrapCommitTs(m *base.Module) {
	return
}
func F_BuildDescFromLists(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	v5 = int32(0)
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v12
	goto L3
L2:
	;
	v13 = v5
	goto L3
L3:
	;
	v18 = F_palloc(m, v13*int32(108)+int32(28))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v22 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v13
	v25 = int64(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v25
	v36 = v22
	v42 = v5
	goto L6
L6:
	;
	v43 = int32(0)
	if l0 == v43 {
		v54 = v43
		goto L8
	} else {
		goto L9
	}
L8:
	;
	if l1 == int32(0) {
		v63 = v43
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v48 <= v36 {
		v54 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v54 = v50 + v36<<(uint(int32(2))%32)
	goto L8
L11:
	;
	v64 = int32(0)
	if l2 == v64 {
		v75 = v64
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v57 <= v36 {
		v63 = v43
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v63 = v59 + v36<<(uint(int32(2))%32)
	goto L11
L14:
	;
	if l3 == int32(0) {
		v84 = v64
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v69 <= v36 {
		v75 = int32(0)
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v75 = v71 + v36<<(uint(int32(2))%32)
	goto L14
L17:
	;
	v85 = int32(0)
	if v84 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v78 <= v36 {
		v84 = v64
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v84 = v80 + v36<<(uint(int32(2))%32)
	goto L17
L20:
	;
	v94 = base.B2i32(v75 == v85) | (base.B2i32(v54 == v85) | base.B2i32(v63 == v85))
	goto L22
L21:
	;
	v94 = int32(1)
	goto L22
L22:
	;
	if v94 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v95 = int32(0)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v95 < v96 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v181 = base.I32_extend16_s(v42 + int32(1))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	F_TupleDescInitEntry(m, v18, v181, v183, v184, v185, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L44
	}
L26:
	;
	v100 = v18 + int32(28)
	v105 = v96
	v107 = v95
	v111 = int32(0)
	goto L30
L27:
	;
	v165 = v95
	v174 = v96
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v165
	return v18
L29:
	;
	v165 = v160
	v174 = v139
	goto L28
L30:
	;
	v118 = v100 + v96<<(uint(int32(3))%32) + v107*int32(100)
	v121 = v100 + v107<<(uint(int32(3))%32)
	if v105 != v96 {
		v139 = v105
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v160 = v96
	goto L29
L32:
	;
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+2)))
	if v140 <= int32(0) {
		v160 = v107
		goto L29
	} else {
		goto L40
	}
L33:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+7)))
	if v123 != int32(118) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v139 = v107
	goto L32
L35:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
	if v126 != int32(1) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+6)))
	if v129&int32(6) != 0 {
		goto L34
	} else {
		goto L37
	}
L37:
	;
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v121)+2)))
	if v132 <= int32(0) {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+90)))
	if v135 != int32(118) {
		v139 = v96
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+90)))
	if v143 == int32(118) {
		v160 = v107
		goto L29
	} else {
		goto L41
	}
L41:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+5)))
	v152 = (v111 + v146 - int32(1)) & (int32(0) - v146)
	if int32(_a_F_BuildDescFromLists_0) < v152 {
		v160 = v107
		goto L29
	} else {
		goto L42
	}
L42:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v152)
	v158 = v107 + int32(1)
	if v158 != v96 {
		v105 = v139
		v107 = v158
		v111 = v152 + v140
		goto L30
	} else {
		goto L43
	}
L43:
	;
	goto L31
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v189<<(uint(int32(3))%32)+v181*int32(100))+24)) = v178
	v36 = v36 + int32(1)
	v42 = v181
	goto L6
}
func F_base32hex_dec_len(m *base.Module, l0 int32, l1 int32) int64 {
	return int64(base.Ui64(base.I64_extend_i32_u(l1)*int64(5)) >> (uint(int64(3)) % 64))
}
func F_basque_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v496 int32
	_ = v496
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v541 int32
	_ = v541
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v572 int32
	_ = v572
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v589 int32
	_ = v589
	var v598 int32
	_ = v598
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v685 int32
	_ = v685
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v702 int32
	_ = v702
	var v711 int32
	_ = v711
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v919 int32
	_ = v919
	var v924 int32
	_ = v924
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v6
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 < v10 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v509 < v10 {
		goto L149
	} else {
		goto L150
	}
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v496
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v290 < v10 {
		goto L83
	} else {
		goto L84
	}
L4:
	;
	if v63 != 0 {
		goto L3
	} else {
		goto L18
	}
L5:
	;
	v22 = v10
	goto L7
L6:
	;
	v22 = v6
	goto L7
L7:
	;
	goto L9
L8:
	;
	v63 = v58
	goto L4
L9:
	;
	if v10 == v22 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v58 = int32(0)
	goto L8
L11:
	;
	v63 = int32(-1)
	goto L4
L12:
	;
	goto L13
L13:
	;
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v10))))
	if int32(117) < v37 {
		v58 = v34
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v39 = v37 - int32(97)
	if v39 < int32(0) {
		v58 = v34
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v39)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v45)>>(uint(v39&int32(7))%32))&int32(1) == int32(0) {
		v58 = v34
		goto L8
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
	goto L17
L17:
	;
	goto L10
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v73 < v64 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v178 < v64 {
		goto L53
	} else {
		goto L54
	}
L20:
	;
	if v113 != 0 {
		goto L19
	} else {
		goto L35
	}
L21:
	;
	v75 = v64
	goto L23
L22:
	;
	v75 = v73
	goto L23
L23:
	;
	goto L25
L24:
	;
	v113 = v110
	goto L20
L25:
	;
	if v64 == v75 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v110 = int32(0)
	goto L24
L27:
	;
	v113 = int32(-1)
	goto L20
L28:
	;
	goto L29
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86+v64))))
	if int32(117) < v88 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64 + int32(1)
	goto L34
L31:
	;
	v90 = v88 - int32(97)
	if v90 < int32(0) {
		goto L30
	} else {
		goto L32
	}
L32:
	;
	v93 = int32(1)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v90)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v97)>>(uint(v90&int32(7))%32))&v93 != 0 {
		v110 = v93
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	goto L26
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v122 < v121 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v162 < int32(0) {
		goto L19
	} else {
		goto L51
	}
L37:
	;
	v124 = v121
	goto L39
L38:
	;
	v124 = v122
	goto L39
L39:
	;
	v131 = v121
	goto L41
L40:
	;
	v162 = v142
	goto L36
L41:
	;
	if v131 == v124 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v162 = int32(-1)
	goto L36
L44:
	;
	goto L45
L45:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135+v131))))
	if int32(117) < v137 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v154 = v131 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v154
	v131 = v154
	goto L41
L47:
	;
	v139 = v137 - int32(97)
	if v139 < int32(0) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v142 = int32(1)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v139)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v146)>>(uint(v139&int32(7))%32))&v142 != 0 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	goto L46
L51:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v496 = v165 + v162
	goto L2
L52:
	;
	if v221 != 0 {
		goto L3
	} else {
		goto L66
	}
L53:
	;
	v180 = v64
	goto L55
L54:
	;
	v180 = v178
	goto L55
L55:
	;
	goto L57
L56:
	;
	v221 = v216
	goto L52
L57:
	;
	if v64 == v180 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v216 = int32(0)
	goto L56
L59:
	;
	v221 = int32(-1)
	goto L52
L60:
	;
	goto L61
L61:
	;
	v192 = int32(1)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193+v64))))
	if int32(117) < v195 {
		v216 = v192
		goto L56
	} else {
		goto L62
	}
L62:
	;
	v197 = v195 - int32(97)
	if v197 < int32(0) {
		v216 = v192
		goto L56
	} else {
		goto L63
	}
L63:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v197)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v203)>>(uint(v197&int32(7))%32))&int32(1) == int32(0) {
		v216 = v192
		goto L56
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v64 + int32(1)
	goto L65
L65:
	;
	goto L58
L66:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v231 < v230 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v274 < int32(0) {
		goto L3
	} else {
		goto L81
	}
L68:
	;
	v233 = v230
	goto L70
L69:
	;
	v233 = v231
	goto L70
L70:
	;
	v239 = v230
	goto L72
L71:
	;
	v274 = int32(1)
	goto L67
L72:
	;
	if v239 == v233 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v274 = int32(-1)
	goto L67
L75:
	;
	goto L76
L76:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246+v239))))
	if int32(117) < v248 {
		goto L71
	} else {
		goto L77
	}
L77:
	;
	v250 = v248 - int32(97)
	if v250 < int32(0) {
		goto L71
	} else {
		goto L78
	}
L78:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v250)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v256)>>(uint(v250&int32(7))%32))&int32(1) == int32(0) {
		goto L71
	} else {
		goto L79
	}
L79:
	;
	v265 = v239 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v265
	v239 = v265
	goto L72
L81:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v496 = v277 + v274
	goto L2
L82:
	;
	if v330 != 0 {
		goto L1
	} else {
		goto L97
	}
L83:
	;
	v292 = v10
	goto L85
L84:
	;
	v292 = v290
	goto L85
L85:
	;
	goto L87
L86:
	;
	v330 = v327
	goto L82
L87:
	;
	if v10 == v292 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v327 = int32(0)
	goto L86
L89:
	;
	v330 = int32(-1)
	goto L82
L90:
	;
	goto L91
L91:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303+v10))))
	if int32(117) < v305 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v10 + int32(1)
	goto L96
L93:
	;
	v307 = v305 - int32(97)
	if v307 < int32(0) {
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v310 = int32(1)
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v307)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v314)>>(uint(v307&int32(7))%32))&v310 != 0 {
		v327 = v310
		goto L86
	} else {
		goto L95
	}
L95:
	;
	goto L92
L96:
	;
	goto L88
L97:
	;
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v340 < v331 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
	v445 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v445 < v331 {
		goto L132
	} else {
		goto L133
	}
L99:
	;
	if v380 != 0 {
		goto L98
	} else {
		goto L114
	}
L100:
	;
	v342 = v331
	goto L102
L101:
	;
	v342 = v340
	goto L102
L102:
	;
	goto L104
L103:
	;
	v380 = v377
	goto L99
L104:
	;
	if v331 == v342 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	v377 = int32(0)
	goto L103
L106:
	;
	v380 = int32(-1)
	goto L99
L107:
	;
	goto L108
L108:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v353+v331))))
	if int32(117) < v355 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331 + int32(1)
	goto L113
L110:
	;
	v357 = v355 - int32(97)
	if v357 < int32(0) {
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v360 = int32(1)
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v357)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v364)>>(uint(v357&int32(7))%32))&v360 != 0 {
		v377 = v360
		goto L103
	} else {
		goto L112
	}
L112:
	;
	goto L109
L113:
	;
	goto L105
L114:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v389 < v388 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v429 < int32(0) {
		goto L98
	} else {
		goto L130
	}
L116:
	;
	v391 = v388
	goto L118
L117:
	;
	v391 = v389
	goto L118
L118:
	;
	v398 = v388
	goto L120
L119:
	;
	v429 = v409
	goto L115
L120:
	;
	if v398 == v391 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v429 = int32(-1)
	goto L115
L123:
	;
	goto L124
L124:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v402+v398))))
	if int32(117) < v404 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v421 = v398 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v421
	v398 = v421
	goto L120
L126:
	;
	v406 = v404 - int32(97)
	if v406 < int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v409 = int32(1)
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v406)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v413)>>(uint(v406&int32(7))%32))&v409 != 0 {
		goto L119
	} else {
		goto L128
	}
L128:
	;
	goto L125
L130:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v496 = v432 + v429
	goto L2
L131:
	;
	if v488 != 0 {
		goto L1
	} else {
		goto L145
	}
L132:
	;
	v447 = v331
	goto L134
L133:
	;
	v447 = v445
	goto L134
L134:
	;
	goto L136
L135:
	;
	v488 = v483
	goto L131
L136:
	;
	if v331 == v447 {
		goto L138
	} else {
		goto L139
	}
L137:
	;
	v483 = int32(0)
	goto L135
L138:
	;
	v488 = int32(-1)
	goto L131
L139:
	;
	goto L140
L140:
	;
	v459 = int32(1)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460+v331))))
	if int32(117) < v462 {
		v483 = v459
		goto L135
	} else {
		goto L141
	}
L141:
	;
	v464 = v462 - int32(97)
	if v464 < int32(0) {
		v483 = v459
		goto L135
	} else {
		goto L142
	}
L142:
	;
	v470 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v464)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v470)>>(uint(v464&int32(7))%32))&int32(1) == int32(0) {
		v483 = v459
		goto L135
	} else {
		goto L143
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331 + int32(1)
	goto L144
L144:
	;
	goto L137
L145:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v490 <= v489 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	v496 = v489 + int32(1)
	goto L2
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v10
	v728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v728
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v728
	v732 = v728 - int32(1)
	if v732 <= v10 {
		goto L212
	} else {
		goto L213
	}
L148:
	;
	if v549 < int32(0) {
		goto L147
	} else {
		goto L163
	}
L149:
	;
	v511 = v10
	goto L151
L150:
	;
	v511 = v509
	goto L151
L151:
	;
	v518 = v10
	goto L153
L152:
	;
	v549 = v529
	goto L148
L153:
	;
	if v518 == v511 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v549 = int32(-1)
	goto L148
L156:
	;
	goto L157
L157:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v522+v518))))
	if int32(117) < v524 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v541 = v518 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v541
	v518 = v541
	goto L153
L159:
	;
	v526 = v524 - int32(97)
	if v526 < int32(0) {
		goto L158
	} else {
		goto L160
	}
L160:
	;
	v529 = int32(1)
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v526)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v533)>>(uint(v526&int32(7))%32))&v529 != 0 {
		goto L152
	} else {
		goto L161
	}
L161:
	;
	goto L158
L163:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v553 = v552 + v549
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v553
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v564 < v553 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	if v607 < int32(0) {
		goto L147
	} else {
		goto L178
	}
L165:
	;
	v566 = v553
	goto L167
L166:
	;
	v566 = v564
	goto L167
L167:
	;
	v572 = v553
	goto L169
L168:
	;
	v607 = int32(1)
	goto L164
L169:
	;
	if v572 == v566 {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v607 = int32(-1)
	goto L164
L172:
	;
	goto L173
L173:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v581 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579+v572))))
	if int32(117) < v581 {
		goto L168
	} else {
		goto L174
	}
L174:
	;
	v583 = v581 - int32(97)
	if v583 < int32(0) {
		goto L168
	} else {
		goto L175
	}
L175:
	;
	v589 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v583)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v589)>>(uint(v583&int32(7))%32))&int32(1) == int32(0) {
		goto L168
	} else {
		goto L176
	}
L176:
	;
	v598 = v572 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v598
	v572 = v598
	goto L169
L178:
	;
	v610 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v611 = v610 + v607
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v611
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v622 < v611 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	if v662 < int32(0) {
		goto L147
	} else {
		goto L194
	}
L180:
	;
	v624 = v611
	goto L182
L181:
	;
	v624 = v622
	goto L182
L182:
	;
	v631 = v611
	goto L184
L183:
	;
	v662 = v642
	goto L179
L184:
	;
	if v631 == v624 {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v662 = int32(-1)
	goto L179
L187:
	;
	goto L188
L188:
	;
	v635 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v637 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v635+v631))))
	if int32(117) < v637 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v654 = v631 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v654
	v631 = v654
	goto L184
L190:
	;
	v639 = v637 - int32(97)
	if v639 < int32(0) {
		goto L189
	} else {
		goto L191
	}
L191:
	;
	v642 = int32(1)
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v639)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v646)>>(uint(v639&int32(7))%32))&v642 != 0 {
		goto L183
	} else {
		goto L192
	}
L192:
	;
	goto L189
L194:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v666 = v665 + v662
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v666
	v677 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v677 < v666 {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	if v720 < int32(0) {
		goto L147
	} else {
		goto L209
	}
L196:
	;
	v679 = v666
	goto L198
L197:
	;
	v679 = v677
	goto L198
L198:
	;
	v685 = v666
	goto L200
L199:
	;
	v720 = int32(1)
	goto L195
L200:
	;
	if v685 == v679 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v720 = int32(-1)
	goto L195
L203:
	;
	goto L204
L204:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v692+v685))))
	if int32(117) < v694 {
		goto L199
	} else {
		goto L205
	}
L205:
	;
	v696 = v694 - int32(97)
	if v696 < int32(0) {
		goto L199
	} else {
		goto L206
	}
L206:
	;
	v702 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v696)>>(uint(int32(3))%32)))+uint32(_c_F_basque_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v702)>>(uint(v696&int32(7))%32))&int32(1) == int32(0) {
		goto L199
	} else {
		goto L207
	}
L207:
	;
	v711 = v685 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v711
	v685 = v711
	goto L200
L209:
	;
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v723 + v720
	goto L147
L210:
	;
	return v924
L211:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v791 = v789 + (v786 - v785)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v791
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v791
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v791 <= v794 {
		v869 = v791
		v871 = v789
		goto L229
	} else {
		goto L230
	}
L212:
	;
	v785 = v728
	v786 = v728
	goto L211
L213:
	;
	goto L214
L214:
	;
	v735 = v728
	v736 = v728
	v738 = v732
	goto L215
L215:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v741 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v739+v738))))
	if base.B2i32(v741&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v741)%32)&int32(70566434) == int32(0)) != 0 {
		v785 = v735
		v786 = v736
		goto L211
	} else {
		goto L217
	}
L216:
	;
	v785 = v779
	v786 = v777
	goto L211
L217:
	;
	v756 = F_find_among_b(m, l0, int32(_a_F_basque_ISO_8859_1_stem_0), int32(109), int32(0))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	return int32(0)
L219:
	;
	if v756 == int32(0) {
		v785 = v735
		v786 = v736
		goto L211
	} else {
		goto L220
	}
L220:
	;
	v762 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v762
	switch v756 - int32(1) {
	case 0:
		goto L223
	case 1:
		goto L222
	default:
		goto L221
	}
L221:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v777
	v779 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v781 = v777 - int32(1)
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v782 < v781 {
		v735 = v779
		v736 = v777
		v738 = v781
		goto L215
	} else {
		goto L228
	}
L222:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v762 < v771 {
		v785 = v735
		v786 = v736
		goto L211
	} else {
		goto L226
	}
L223:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v762 < v766 {
		v785 = v735
		v786 = v736
		goto L211
	} else {
		goto L224
	}
L224:
	;
	v768 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v768 {
		goto L221
	} else {
		goto L225
	}
L225:
	;
	v924 = v768
	goto L210
L226:
	;
	v773 = F_slice_del(m, l0)
	mBase = m.M
	if v773 < int32(0) {
		v924 = v773
		goto L210
	} else {
		goto L227
	}
L227:
	;
	goto L221
L228:
	;
	goto L216
L229:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v874 = v872 + (v869 - v871)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v874
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v874
	v878 = v874 - int32(1)
	v879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v878 <= v879 {
		goto L256
	} else {
		goto L257
	}
L230:
	;
	v798 = v791
	v800 = v789
	goto L231
L231:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v803 = int32(1)
	v805 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v801+v798-v803))))
	if base.B2i32(v805&int32(224) != int32(96))|base.B2i32(v803<<(uint(v805)%32)&int32(71162402) == int32(0)) != 0 {
		v869 = v798
		v871 = v800
		goto L229
	} else {
		goto L233
	}
L232:
	;
	v869 = v862
	v871 = v864
	goto L229
L233:
	;
	v820 = F_find_among_b(m, l0, int32(_a_F_basque_ISO_8859_1_stem_1), int32(295), int32(0))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L218
	} else {
		goto L234
	}
L234:
	;
	if v820 == int32(0) {
		v869 = v798
		v871 = v800
		goto L229
	} else {
		goto L235
	}
L235:
	;
	v824 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v824
	switch v820 - int32(1) {
	case 0:
		goto L242
	case 1:
		goto L241
	case 2:
		goto L240
	case 3:
		goto L239
	case 4:
		goto L238
	case 5:
		goto L237
	default:
		goto L236
	}
L236:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v862
	v864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v865 < v862 {
		v798 = v862
		v800 = v864
		goto L231
	} else {
		goto L255
	}
L237:
	;
	v857 = F_slice_from_s(m, l0, int32(6), int32(_a_F_basque_ISO_8859_1_stem_2))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L218
	} else {
		goto L253
	}
L238:
	;
	v851 = F_slice_from_s(m, l0, int32(3), int32(_a_F_basque_ISO_8859_1_stem_3))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L218
	} else {
		goto L251
	}
L239:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v824 < v844 {
		v869 = v798
		v871 = v800
		goto L229
	} else {
		goto L249
	}
L240:
	;
	v840 = F_slice_from_s(m, l0, int32(3), int32(_a_F_basque_ISO_8859_1_stem_4))
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L218
	} else {
		goto L247
	}
L241:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v824 < v833 {
		v869 = v798
		v871 = v800
		goto L229
	} else {
		goto L245
	}
L242:
	;
	v828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v824 < v828 {
		v869 = v798
		v871 = v800
		goto L229
	} else {
		goto L243
	}
L243:
	;
	v830 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v830 {
		goto L236
	} else {
		goto L244
	}
L244:
	;
	v924 = v830
	goto L210
L245:
	;
	v835 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v835 {
		goto L236
	} else {
		goto L246
	}
L246:
	;
	v924 = v835
	goto L210
L247:
	;
	if int32(0) <= v840 {
		goto L236
	} else {
		goto L248
	}
L248:
	;
	v924 = v840
	goto L210
L249:
	;
	v846 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v846 {
		goto L236
	} else {
		goto L250
	}
L250:
	;
	v924 = v846
	goto L210
L251:
	;
	if int32(0) <= v851 {
		goto L236
	} else {
		goto L252
	}
L252:
	;
	v924 = v851
	goto L210
L253:
	;
	if v857 < int32(0) {
		v924 = v857
		goto L210
	} else {
		goto L254
	}
L254:
	;
	goto L236
L255:
	;
	goto L232
L256:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v919
	v924 = int32(1)
	goto L210
L257:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v881+v878))))
	if base.B2i32(v883&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v883)%32)&int32(_a_F_basque_ISO_8859_1_stem_5) == int32(0)) != 0 {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v898 = F_find_among_b(m, l0, int32(_a_F_basque_ISO_8859_1_stem_6), int32(19), int32(0))
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L218
	} else {
		goto L259
	}
L259:
	;
	if v898 == int32(0) {
		goto L256
	} else {
		goto L260
	}
L260:
	;
	v902 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v902
	switch v898 - int32(1) {
	case 0:
		goto L262
	case 1:
		goto L261
	default:
		goto L256
	}
L261:
	;
	v913 = F_slice_from_s(m, l0, int32(1), int32(_a_F_basque_ISO_8859_1_stem_7))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L218
	} else {
		goto L265
	}
L262:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v902 < v906 {
		goto L256
	} else {
		goto L263
	}
L263:
	;
	v908 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v908 {
		goto L256
	} else {
		goto L264
	}
L264:
	;
	v924 = v908
	goto L210
L265:
	;
	if v913 < int32(0) {
		v924 = v913
		goto L210
	} else {
		goto L266
	}
L266:
	;
	goto L256
}
func F_binaryheap_add_unordered(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v5 <= v4 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_binaryheap_add_unordered_0), int32(0))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_binaryheap_add_unordered_1), int32(121), int32(_a_F_binaryheap_add_unordered_2))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v20 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v20)
		*(*int64)(unsafe.Add(mBase, uint32(l0+v4<<(uint(int32(3))%32))+24)) = l1
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v26 + int32(1)
		return
	}
}
func F_binaryheap_build(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v31 int64
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int64
	_ = v81
	var v101 int32
	_ = v101
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = v10 - int32(2)
	if int32(-1) <= v12 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = l0 + int32(24)
	v18 = base.I32_div_s(v12, int32(2))
	v20 = v18
	goto L4
L2:
	;
	goto L3
L3:
	;
	v101 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v101)
	return
L4:
	;
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v16+v20<<(uint(int32(3))%32))))
	v34 = v20
	goto L6
L5:
	;
	goto L3
L6:
	;
	v41 = int32(1)
	v42 = v34 << (uint(v41) % 32)
	v44 = v42 | v41
	v46 = v42 + int32(2)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v46 < v47 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16+v34<<(uint(int32(3))%32)))) = v31
	if int32(0) < v20 {
		v20 = v20 - int32(1)
		goto L4
	} else {
		goto L20
	}
L8:
	;
	goto L7
L9:
	;
	v49 = int32(3)
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v16+v44<<(uint(v49)%32))))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v16+v46<<(uint(v49)%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v59 = m.T0[v58].(func(*base.Module, int64, int64, int32) int32)(m, v52, v56, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v65 = v44
	v66 = v47
	goto L11
L11:
	;
	if v66 <= v44 {
		goto L8
	} else {
		goto L17
	}
L12:
	;
	return
L13:
	;
	if v59 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v63 = v46
	goto L16
L15:
	;
	v63 = v44
	goto L16
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v65 = v63
	v66 = v64
	goto L11
L17:
	;
	v70 = v16 + v65<<(uint(int32(3))%32)
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v74 = m.T0[v73].(func(*base.Module, int64, int64, int32) int32)(m, v31, v71, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	if int32(0) <= v74 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v70)))
	*(*int64)(unsafe.Add(mBase, uint32(v16+v34<<(uint(int32(3))%32)))) = v81
	v34 = v65
	goto L6
L20:
	;
	goto L5
}
func F_bitncommon(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v4 = int32(0)
	v9 = int32(8)
	v10 = base.I32_div_s(l2, v9)
	if v9 <= l2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v76 + v72<<(uint(int32(3))%32)
L2:
	;
	v58 = v49
	goto L13
L3:
	;
	v16 = v4
	goto L6
L4:
	;
	v33 = v4
	goto L5
L5:
	;
	v40 = l2 - v10<<(uint(int32(3))%32)
	if v40 == int32(0) {
		v72 = v33
		v76 = v4
		goto L1
	} else {
		goto L12
	}
L6:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v16))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v16))))
	if v22 != v24 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v33 = v10
	goto L5
L8:
	;
	v49 = int32(7)
	v50 = v16
	v52 = v22
	v53 = v24
	goto L2
L9:
	;
	goto L10
L10:
	;
	v28 = v16 + int32(1)
	if v28 != v10 {
		v16 = v28
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L7
L12:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v33))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v33))))
	v49 = v40
	v50 = v33
	v52 = v46
	v53 = v44
	goto L2
L13:
	;
	if int32(base.Ui32(v52^v53)>>(uint(int32(8)-v58)%32)) != 0 {
		v58 = v58 - int32(1)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v72 = v50
	v76 = v58
	goto L1
L15:
	;
	goto L14
}
func F_bitnot(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v14 = F_palloc(m, int32(base.Ui32(v11)>>(uint(int32(2))%32)))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			v18 = v16 & int32(-4)
			*(*int32)(unsafe.Add(mBase, uint32(v14))) = v18
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v20
			v23 = v14 + int32(8)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
			if base.Ui32(int32(36)) <= base.Ui32(v24) {
				v29 = v23
				v31 = v7 + int32(8)
				for {
					v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
					v36 = v34 ^ int32(-1)
					*(*uint8)(unsafe.Add(mBase, uint32(v29))) = uint8(v36)
					v38 = int32(1)
					v39 = v29 + v38
					v41 = v31 + v38
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if base.Ui32(v41) < base.Ui32(v7+int32(base.Ui32(v42)>>(uint(int32(2))%32))) {
						v29 = v39
						v31 = v41
						continue
					} else {
						break
					}
					break
				}
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				v49 = v39
				v51 = v47
				v54 = v48
			} else {
				v49 = v23
				v51 = v20
				v54 = v18
			}
			v61 = v54<<(uint(int32(1))%32)&int32(-8) - v51 + int32(-64)
			if int32(0) < v61 {
				v65 = v49 - int32(1)
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
				v69 = v66 & (int32(255) << (uint(v61) % 32))
				*(*uint8)(unsafe.Add(mBase, uint32(v65))) = uint8(v69)
			} else {
			}
			return base.I64_extend_i32_u(v14)
		}
	}
}
func F_bitoctetlength(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		return base.I64_extend_i32_s(int32(base.Ui32(v7)>>(uint(int32(2))%32)) - int32(8))
	}
}
func F_bitsubstr(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v10 = F_bitsubstring(m, v3, v7, v8, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v10)
		}
	}
}
func F_bitsubstring(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l3 != 0 {
		v27 = v12 + int32(1)
		v28 = int32(1)
		if l1 <= v28 {
			v31 = v28
		} else {
			v31 = l1
		}
		if base.B2i32(v31 <= v12)&base.B2i32(v31 < v27) == int32(0) {
			v38 = F_palloc(m, int32(8))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v38))) = int64(32)
				return v38
			}
		} else {
			v45 = v27 - v31
			v47 = v45 + int32(7)
			v48 = int32(8)
			v49 = base.I32_div_s(v47, v48)
			v51 = v49 + v48
			v52 = F_palloc(m, v51)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v45
				v56 = v51 << (uint(int32(2)) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v52))) = v56
				v59 = v31 - int32(1)
				v61 = v59 & int32(7)
				if v61 == int32(0) {
					if v49 == int32(0) {
						v180 = v56
						v181 = v45
					} else {
						v66 = int32(8)
						base.MemoryCopy(m, v52+v66, l0+int32(base.Ui32(v59)>>(uint(int32(3))%32))+v66, v49)
						v180 = v56
						v181 = v45
					}
				} else {
					if v47 < int32(8) {
						v180 = v56
						v181 = v45
					} else {
						v76 = int32(8)
						v77 = v76 - v61
						v79 = v52 + v76
						v84 = l0 + int32(base.Ui32(v59)>>(uint(int32(3))%32)) + v76
						if base.Ui32(v76) <= base.Ui32(v45-int32(1)) {
							v95 = v84
							v96 = v79
							v97 = int32(0)
							for {
								v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
								v106 = v105 << (uint(v61) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v106)
								v109 = v95 + int32(1)
								v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(v109) < base.Ui32(l0+int32(base.Ui32(v110)>>(uint(int32(2))%32))) {
									v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
									v117 = int32(base.Ui32(v115)>>(uint(v77)%32)) | v106
									*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v117)
								} else {
								}
								v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
								v120 = v119 << (uint(v61) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)) = uint8(v120)
								v122 = int32(2)
								v123 = v95 + v122
								v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(v123) < base.Ui32(l0+int32(base.Ui32(v124)>>(uint(v122)%32))) {
									v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
									v131 = int32(base.Ui32(v129)>>(uint(v77)%32)) | v120
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)) = uint8(v131)
								} else {
								}
								v133 = int32(2)
								v134 = v96 + v133
								v136 = v97 + v133
								if v136 != v49&int32(268435454) {
									v95 = v123
									v96 = v134
									v97 = v136
									continue
								} else {
									break
								}
								break
							}
							if v49&int32(1) == int32(0) {
							} else {
								v141 = v123
								v142 = v134
								v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
								v152 = v151 << (uint(v61) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v152)
								v155 = v141 + int32(1)
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(l0+int32(base.Ui32(v156)>>(uint(int32(2))%32))) <= base.Ui32(v155) {
								} else {
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
									v163 = int32(base.Ui32(v161)>>(uint(v77)%32)) | v152
									*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v163)
								}
							}
						} else {
							v141 = v84
							v142 = v79
							v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
							v152 = v151 << (uint(v61) % 32)
							*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v152)
							v155 = v141 + int32(1)
							v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							if base.Ui32(l0+int32(base.Ui32(v156)>>(uint(int32(2))%32))) <= base.Ui32(v155) {
							} else {
								v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
								v163 = int32(base.Ui32(v161)>>(uint(v77)%32)) | v152
								*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v163)
							}
						}
						v176 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
						v177 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
						v180 = v177
						v181 = v176
					}
				}
				v190 = int32(base.Ui32(v180) >> (uint(int32(2)) % 32))
				v195 = v190<<(uint(int32(3))%32) - v181 + int32(-64)
				if int32(0) < v195 {
					v200 = v190 + v52 - int32(1)
					v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
					v204 = v201 & (int32(255) << (uint(v195) % 32))
					*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v204)
				} else {
				}
				return v52
			}
		}
	} else {
		v16 = base.B2i32(l2 < int32(0))
		if l2 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v211 = m.ExcPending
			if v211 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(17039490))
				mBase = m.M
				v214 = m.ExcPending
				if v214 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_bitsubstring_0), int32(0))
					mBase = m.M
					v218 = m.ExcPending
					if v218 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_bitsubstring_1), int32(1081), int32(_a_F_bitsubstring_2))
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v18 = v12 + int32(1)
			v19 = l1 + l2
			if base.B2i32(v19 < l1) != v16 {
				v27 = v18
			} else {
				if v19 < v18 {
					v23 = v19
				} else {
					v23 = v18
				}
				v27 = v23
			}
			v28 = int32(1)
			if l1 <= v28 {
				v31 = v28
			} else {
				v31 = l1
			}
			if base.B2i32(v31 <= v12)&base.B2i32(v31 < v27) == int32(0) {
				v38 = F_palloc(m, int32(8))
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v38))) = int64(32)
					return v38
				}
			} else {
				v45 = v27 - v31
				v47 = v45 + int32(7)
				v48 = int32(8)
				v49 = base.I32_div_s(v47, v48)
				v51 = v49 + v48
				v52 = F_palloc(m, v51)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v45
					v56 = v51 << (uint(int32(2)) % 32)
					*(*int32)(unsafe.Add(mBase, uint32(v52))) = v56
					v59 = v31 - int32(1)
					v61 = v59 & int32(7)
					if v61 == int32(0) {
						if v49 == int32(0) {
							v180 = v56
							v181 = v45
						} else {
							v66 = int32(8)
							base.MemoryCopy(m, v52+v66, l0+int32(base.Ui32(v59)>>(uint(int32(3))%32))+v66, v49)
							v180 = v56
							v181 = v45
						}
					} else {
						if v47 < int32(8) {
							v180 = v56
							v181 = v45
						} else {
							v76 = int32(8)
							v77 = v76 - v61
							v79 = v52 + v76
							v84 = l0 + int32(base.Ui32(v59)>>(uint(int32(3))%32)) + v76
							if base.Ui32(v76) <= base.Ui32(v45-int32(1)) {
								v95 = v84
								v96 = v79
								v97 = int32(0)
								for {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95))))
									v106 = v105 << (uint(v61) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v106)
									v109 = v95 + int32(1)
									v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if base.Ui32(v109) < base.Ui32(l0+int32(base.Ui32(v110)>>(uint(int32(2))%32))) {
										v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
										v117 = int32(base.Ui32(v115)>>(uint(v77)%32)) | v106
										*(*uint8)(unsafe.Add(mBase, uint32(v96))) = uint8(v117)
									} else {
									}
									v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
									v120 = v119 << (uint(v61) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)) = uint8(v120)
									v122 = int32(2)
									v123 = v95 + v122
									v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if base.Ui32(v123) < base.Ui32(l0+int32(base.Ui32(v124)>>(uint(v122)%32))) {
										v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
										v131 = int32(base.Ui32(v129)>>(uint(v77)%32)) | v120
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)) = uint8(v131)
									} else {
									}
									v133 = int32(2)
									v134 = v96 + v133
									v136 = v97 + v133
									if v136 != v49&int32(268435454) {
										v95 = v123
										v96 = v134
										v97 = v136
										continue
									} else {
										break
									}
									break
								}
								if v49&int32(1) == int32(0) {
								} else {
									v141 = v123
									v142 = v134
									v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
									v152 = v151 << (uint(v61) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v152)
									v155 = v141 + int32(1)
									v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
									if base.Ui32(l0+int32(base.Ui32(v156)>>(uint(int32(2))%32))) <= base.Ui32(v155) {
									} else {
										v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
										v163 = int32(base.Ui32(v161)>>(uint(v77)%32)) | v152
										*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v163)
									}
								}
							} else {
								v141 = v84
								v142 = v79
								v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
								v152 = v151 << (uint(v61) % 32)
								*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v152)
								v155 = v141 + int32(1)
								v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if base.Ui32(l0+int32(base.Ui32(v156)>>(uint(int32(2))%32))) <= base.Ui32(v155) {
								} else {
									v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
									v163 = int32(base.Ui32(v161)>>(uint(v77)%32)) | v152
									*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v163)
								}
							}
							v176 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
							v177 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
							v180 = v177
							v181 = v176
						}
					}
					v190 = int32(base.Ui32(v180) >> (uint(int32(2)) % 32))
					v195 = v190<<(uint(int32(3))%32) - v181 + int32(-64)
					if int32(0) < v195 {
						v200 = v190 + v52 - int32(1)
						v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200))))
						v204 = v201 & (int32(255) << (uint(v195) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v200))) = uint8(v204)
					} else {
					}
					return v52
				}
			}
		}
	}
}
func F_boolexpr_startup_fn(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v3
	if v3 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		return
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+12))
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = v9
		return
	}
}
func F_boolsend(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(1))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*uint8)(unsafe.Add(mBase, uint32(v16+v17))) = uint8(base.B2i32(v8 != int64(0)))
			v23 = v16 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v23
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v26))) = v23 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return base.I64_extend_i32_u(v26)
		}
	}
}
func F_bottomup_sort_and_shrink_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	v5 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1))))
	if v6 < v5 {
		return int32(-1)
	} else {
		if v5 < v6 {
			return int32(1)
		} else {
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
			if v13 != v14 {
				v17 = int32(1)
				v19 = base.I32_extend16_s(v13)
				if v19&(v19-v17) != 0 {
					v26 = v17 << (uint(int32(32)-base.I32_clz(v19)) % 32)
				} else {
					v26 = v19
				}
				v27 = int32(1)
				v29 = base.I32_extend16_s(v14)
				if v29&(v29-v27) != 0 {
					v36 = v27 << (uint(int32(32)-base.I32_clz(v29)) % 32)
				} else {
					v36 = v29
				}
				if base.Ui32(v36) < base.Ui32(v26) {
					v50 = int32(-1)
				} else {
					if base.Ui32(v26) < base.Ui32(v36) {
						v50 = int32(1)
					} else {
						v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
						v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
						if v45 < v44 {
							v47 = int32(1)
						} else {
							v47 = int32(-1)
						}
						v50 = v47
					}
				}
			} else {
				v44 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
				v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
				if v45 < v44 {
					v47 = int32(1)
				} else {
					v47 = int32(-1)
				}
				v50 = v47
			}
			return v50
		}
	}
}
func F_boxes_bound_box(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v19 float64
	_ = v19
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v48 float64
	_ = v48
	var v54 float64
	_ = v54
	var v60 float64
	_ = v60
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v81 float64
	_ = v81
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = F_palloc(m, int32(32))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
		if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
				v25 = v19
			} else {
				v25 = v13
			}
			if base.F64_lt(v13, v19) != 0 {
				v27 = v19
			} else {
				v27 = v25
			}
			v29 = v27
		} else {
			v29 = v13
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9))) = v29
		v31 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v6)+16))
		if base.Ui64(base.I64_reinterpret_f64(v32)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v31)&int64(9223372036854775807)) {
				v43 = v32
			} else {
				v43 = v31
			}
			if base.F64_gt(v31, v32) != 0 {
				v45 = v32
			} else {
				v45 = v43
			}
			v46 = v45
		} else {
			v46 = v31
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+16)) = v46
		v48 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
		if base.Ui64(base.I64_reinterpret_f64(v48)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			v54 = *(*float64)(unsafe.Add(mBase, uint32(v6)+8))
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) {
				v60 = v54
			} else {
				v60 = v48
			}
			if base.F64_lt(v48, v54) != 0 {
				v62 = v54
			} else {
				v62 = v60
			}
			v64 = v62
		} else {
			v64 = v48
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v64
		v66 = *(*float64)(unsafe.Add(mBase, uint32(v7)+24))
		v67 = *(*float64)(unsafe.Add(mBase, uint32(v6)+24))
		if base.Ui64(base.I64_reinterpret_f64(v67)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) {
				v78 = v67
			} else {
				v78 = v66
			}
			if base.F64_gt(v66, v67) != 0 {
				v80 = v67
			} else {
				v80 = v78
			}
			v81 = v80
		} else {
			v81 = v66
		}
		*(*float64)(unsafe.Add(mBase, uint32(v9)+24)) = v81
		return base.I64_extend_i32_u(v9)
	}
}
func F_bpcharfastcmp_c(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	v10 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(1)
	goto L6
L5:
	;
	v22 = int32(4)
	goto L6
L6:
	;
	v23 = v22 + v17
	v24 = int32(1)
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v28 = v26 & v24
	if v28 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v29 = v24
	goto L9
L8:
	;
	v29 = int32(4)
	goto L9
L9:
	;
	v30 = v10 + v29
	if v26 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v60 = int32(-1)
	v62 = v57 - int32(1)
	if v60 <= v62 {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v36 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v47 = int32(1)
	if v28 != 0 {
		v57 = int32(base.Ui32(v26)>>(uint(v47)%32)) - v47
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v39 = int32(16)
	goto L16
L15:
	;
	v39 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = int32(4)
	goto L19
L18:
	;
	v46 = v39
	goto L19
L19:
	;
	v57 = v46
	goto L10
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v81 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v65 = v60
	goto L24
L23:
	;
	v65 = v62
	goto L24
L24:
	;
	v69 = v57
	goto L25
L25:
	;
	v73 = v69 - int32(1)
	if v73 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v80 = v69
	goto L21
L27:
	;
	v80 = v65 + int32(1)
	goto L21
L28:
	;
	goto L29
L29:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30+v73))))
	if v77 == int32(32) {
		v69 = v73
		goto L25
	} else {
		goto L30
	}
L30:
	;
	goto L26
L31:
	;
	v113 = int32(-1)
	v115 = v110 - int32(1)
	if v113 <= v115 {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v87 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v98 = int32(1)
	if v81&v98 != 0 {
		v110 = int32(base.Ui32(v81)>>(uint(v98)%32)) - v98
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v90 = int32(16)
	goto L37
L36:
	;
	v90 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v87-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v97 = int32(4)
	goto L40
L39:
	;
	v97 = v90
	goto L40
L40:
	;
	v110 = v97
	goto L31
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	v134 = base.B2i32(v80 < v133)
	if v80 < v133 {
		goto L52
	} else {
		goto L53
	}
L43:
	;
	v118 = v113
	goto L45
L44:
	;
	v118 = v115
	goto L45
L45:
	;
	v122 = v110
	goto L46
L46:
	;
	v126 = v122 - int32(1)
	if v126 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v133 = v122
	goto L42
L48:
	;
	v133 = v118 + int32(1)
	goto L42
L49:
	;
	goto L50
L50:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v126))))
	if v130 == int32(32) {
		v122 = v126
		goto L46
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	v135 = v80
	goto L54
L53:
	;
	v135 = v133
	goto L54
L54:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v135) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	if base.I64_extend_i32_u(v10) != l0 {
		goto L73
	} else {
		goto L74
	}
L56:
	;
	v197 = int32(0)
	goto L55
L57:
	;
	v171 = v166
	v172 = v167
	v173 = v168
	goto L67
L58:
	;
	if (v30|v23)&int32(3) != 0 {
		v166 = v30
		v167 = v23
		v168 = v135
		goto L57
	} else {
		goto L61
	}
L59:
	;
	v159 = v30
	v160 = v23
	v161 = v135
	goto L60
L60:
	;
	if v161 == int32(0) {
		goto L56
	} else {
		goto L66
	}
L61:
	;
	v143 = v30
	v144 = v23
	v145 = v135
	goto L62
L62:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if v148 != v149 {
		v166 = v143
		v167 = v144
		v168 = v145
		goto L57
	} else {
		goto L64
	}
L63:
	;
	v159 = v154
	v160 = v152
	v161 = v156
	goto L60
L64:
	;
	v151 = int32(4)
	v152 = v144 + v151
	v154 = v143 + v151
	v156 = v145 - v151
	if base.Ui32(int32(3)) < base.Ui32(v156) {
		v143 = v154
		v144 = v152
		v145 = v156
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v166 = v159
	v167 = v160
	v168 = v161
	goto L57
L67:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171))))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172))))
	if v176 == v177 {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v197 = v176 - v177
	goto L55
L69:
	;
	v179 = int32(1)
	v184 = v173 - v179
	if v184 != 0 {
		v171 = v171 + v179
		v172 = v172 + v179
		v173 = v184
		goto L67
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L68
L72:
	;
	goto L56
L73:
	;
	F_pfree(m, v10)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	if base.I64_extend_i32_u(v17) != l1 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	F_pfree(m, v17)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if v197 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	v208 = v197
	goto L83
L82:
	;
	v208 = base.B2i32(v133 < v80) - v134
	goto L83
L83:
	;
	return v208
}
func F_bracket(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v518 int32
	_ = v518
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v781 int32
	_ = v781
	var v782 int64
	_ = v782
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v823 int32
	_ = v823
	var v832 int32
	_ = v832
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v853 int32
	_ = v853
	var v870 int32
	_ = v870
	var v875 int32
	_ = v875
	var v882 int32
	_ = v882
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v904 int32
	_ = v904
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v983 int32
	_ = v983
	var v995 int32
	_ = v995
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+6)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v14
	v18 = F_next(m, l0)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v29 - int32(67) {
	case 0:
		goto L13
	default:
		goto L10
	case 2:
		goto L14
	case 6:
		goto L15
	case 15:
		goto L17
	case 26, 34:
		goto L9
	case 32:
		goto L11
	case 45:
		goto L16
	case 48:
		goto L12
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1043 != 0 {
		goto L348
	} else {
		goto L349
	}
L6:
	;
	v1033 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v1036 = F_range_(m, l0, v1028, v1027, v1033&int32(8))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L345
	}
L7:
	;
	v887 = F_next(m, l0)
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L292
	}
L8:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v875 != int32(82) {
		v1027 = v870
		v1028 = v870
		goto L6
	} else {
		goto L291
	}
L9:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_okcolors(m, v590, v591)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L1
	} else {
		goto L204
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v586 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v586 != 0 {
		goto L201
	} else {
		goto L202
	}
L11:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v580 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12+v578))) = uint8(v580)
	v582 = F_next(m, l0)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L200
	}
L12:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v561 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v561)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v561)+8)) = v562 | int32(1024)
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v569 = F_cclasscvec(m, l0, v560, v566&int32(8))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L1
	} else {
		goto L194
	}
L13:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v441 = F_next(m, l0)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L152
	}
L14:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v172 = F_next(m, l0)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L74
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v61 = F_next(m, l0)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L32
	}
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v39 = F_next(m, l0)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v36 = v34
	goto L20
L19:
	;
	v36 = int32(11)
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v36
	goto L3
L21:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v41 != int32(82) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
	if v44&int32(8) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v59 != 0 {
		goto L3
	} else {
		goto L31
	}
L25:
	;
	v49 = int32(_a_F_bracket_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v12)+14)) = uint16(v49)
	F_subcoloronechr(m, l0, v38, l1, l2, v12+int32(14))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L1
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v55 = F_allcases(m, l0, v38)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L3
L29:
	;
	F_subcolorcvec(m, l0, v55, l1, l2)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	goto L3
L31:
	;
	v882 = v38
	goto L7
L32:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v63 != int32(112) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v91 = F_next(m, l0)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L41
	}
L34:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v86 = v66
	goto L33
L35:
	;
	goto L36
L36:
	;
	goto L37
L37:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v77 = F_next(m, l0)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	v86 = v76
	goto L33
L39:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v79 == int32(112) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	if base.Ui32(v86) <= base.Ui32(v60) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v96 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v100 != 0 {
		goto L3
	} else {
		goto L48
	}
L45:
	;
	v98 = v96
	goto L47
L46:
	;
	v98 = int32(3)
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v98
	goto L3
L48:
	;
	v101 = v86 - v60
	if v101 == int32(4) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	v870 = v104
	goto L8
L50:
	;
	goto L51
L51:
	;
	v106 = v101 >> (uint(int32(2)) % 32)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v108 | int32(1024)
	v117 = int32(_a_F_bracket_1)
	v118 = int32(_a_F_bracket_2)
	goto L53
L52:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v169 != 0 {
		goto L3
	} else {
		goto L73
	}
L53:
	;
	v123 = F_strlen(m, v118)
	mBase = m.M
	if v123 == v106 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v165 != 0 {
		goto L70
	} else {
		goto L71
	}
L55:
	;
	if v106 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v117)+8))
	if v160 != 0 {
		v117 = v117 + int32(8)
		v118 = v160
		goto L53
	} else {
		goto L69
	}
L58:
	;
	if v157 == int32(0) {
		goto L52
	} else {
		goto L68
	}
L59:
	;
	v157 = int32(0)
	goto L58
L60:
	;
	v129 = v118
	v130 = v60
	v131 = v106
	goto L61
L61:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v134 != v135 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	goto L59
L63:
	;
	v157 = v134 - v135
	goto L58
L64:
	;
	goto L65
L65:
	;
	if v134 == int32(0) {
		goto L59
	} else {
		goto L66
	}
L66:
	;
	v140 = int32(1)
	v145 = v131 - v140
	if v145 != 0 {
		v129 = v129 + v140
		v130 = v130 + int32(4)
		v131 = v145
		goto L61
	} else {
		goto L67
	}
L67:
	;
	goto L62
L68:
	;
	goto L57
L69:
	;
	goto L54
L70:
	;
	v167 = v165
	goto L72
L71:
	;
	v167 = int32(3)
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v167
	goto L3
L73:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+4)))
	v870 = v170
	goto L8
L74:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v174 != int32(112) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v202 = F_next(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L83
	}
L76:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v197 = v177
	goto L75
L77:
	;
	goto L78
L78:
	;
	goto L79
L79:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v188 = F_next(m, l0)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L81
	}
L80:
	;
	v197 = v187
	goto L75
L81:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v190 == int32(112) {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if base.Ui32(v197) <= base.Ui32(v171) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v207 != 0 {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v211 != 0 {
		goto L3
	} else {
		goto L90
	}
L87:
	;
	v209 = v207
	goto L89
L88:
	;
	v209 = int32(3)
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v209
	goto L3
L90:
	;
	v212 = v197 - v171
	if v212 == int32(4) {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v294 = v292 & int32(8)
	v297 = int32(0)
	if base.B2i32(v292&int32(_a_F_bracket_3) == v297)|base.B2i32(v291 != int32(120)) == v297 {
		goto L118
	} else {
		goto L119
	}
L92:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v171)))
	v291 = v215
	goto L91
L93:
	;
	goto L94
L94:
	;
	v217 = v212 >> (uint(int32(2)) % 32)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v218)+8)) = v219 | int32(1024)
	v228 = int32(_a_F_bracket_1)
	v229 = int32(_a_F_bracket_2)
	goto L96
L95:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v280 != 0 {
		goto L3
	} else {
		goto L116
	}
L96:
	;
	v234 = F_strlen(m, v229)
	mBase = m.M
	if v234 == v217 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v276 != 0 {
		goto L113
	} else {
		goto L114
	}
L98:
	;
	if v217 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L99:
	;
	goto L100
L100:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v228)+8))
	if v271 != 0 {
		v228 = v228 + int32(8)
		v229 = v271
		goto L96
	} else {
		goto L112
	}
L101:
	;
	if v268 == int32(0) {
		goto L95
	} else {
		goto L111
	}
L102:
	;
	v268 = int32(0)
	goto L101
L103:
	;
	v240 = v229
	v241 = v171
	v242 = v217
	goto L104
L104:
	;
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240))))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	if v245 != v246 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L102
L106:
	;
	v268 = v245 - v246
	goto L101
L107:
	;
	goto L108
L108:
	;
	if v245 == int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	v251 = int32(1)
	v256 = v242 - v251
	if v256 != 0 {
		v240 = v240 + v251
		v241 = v241 + int32(4)
		v242 = v256
		goto L104
	} else {
		goto L110
	}
L110:
	;
	goto L105
L111:
	;
	goto L100
L112:
	;
	goto L97
L113:
	;
	v278 = v276
	goto L115
L114:
	;
	v278 = int32(3)
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v278
	goto L3
L116:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v228)+4)))
	v291 = v281
	goto L91
L117:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v437 != 0 {
		goto L3
	} else {
		goto L150
	}
L118:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v304 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	goto L120
L120:
	;
	if v294 != 0 {
		goto L137
	} else {
		goto L138
	}
L121:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v348 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v347 + v348
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	v352 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v351+v347<<(uint(v352)%32)))) = int32(120)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v357 + v348
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v361+v357<<(uint(v352)%32)))) = int32(121)
	if v294 == int32(0) {
		v436 = v346
		goto L117
	} else {
		goto L136
	}
L122:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v304)+4))
	if v305 < int32(4) {
		goto L125
	} else {
		goto L126
	}
L123:
	;
	goto L124
L124:
	;
	v321 = F_palloc_extended(m, int32(44), int32(2))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L129
	}
L125:
	;
	F_pfree(m, v304)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L128
	}
L126:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v304)+16))
	if v308 < int32(0) {
		goto L125
	} else {
		goto L127
	}
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v304)+24)) = int32(-1)
	v313 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v304)+12)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v304))) = v313
	v346 = v304
	goto L121
L128:
	;
	goto L124
L129:
	;
	if v321 != 0 {
		goto L130
	} else {
		goto L131
	}
L130:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v321))) = int64(17179869184)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v321)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+20)) = v321 + int32(44)
	*(*int32)(unsafe.Add(mBase, uint32(v321)+8)) = v321 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v321
	v346 = v321
	goto L121
L131:
	;
	goto L132
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v338 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v338
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v341 != 0 {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v343 = v341
	goto L135
L134:
	;
	v343 = int32(12)
	goto L135
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v343
	v346 = v338
	goto L121
L136:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	v370 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v369 + v370
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	v374 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v373+v369<<(uint(v374)%32)))) = int32(88)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v346)))
	*(*int32)(unsafe.Add(mBase, uint32(v346))) = v379 + v370
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v346)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v383+v379<<(uint(v374)%32)))) = int32(89)
	v436 = v346
	goto L117
L137:
	;
	v389 = F_allcases(m, l0, v291)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	if v391 != 0 {
		goto L142
	} else {
		goto L143
	}
L140:
	;
	v436 = v389
	goto L117
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v425))) = v426 + int32(1)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v425)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v430+v426<<(uint(int32(2))%32)))) = v291
	v436 = v425
	goto L117
L142:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)+4))
	if v392 <= int32(0) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L144
L144:
	;
	v409 = F_palloc_extended(m, int32(32), int32(2))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L1
	} else {
		goto L149
	}
L145:
	;
	F_pfree(m, v391)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	if v395 < int32(0) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v391)+24)) = int32(-1)
	v400 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v391)+12)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v391))) = v400
	v425 = v391
	v426 = v400
	goto L141
L148:
	;
	goto L144
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409)+24)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v409)+12)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v409))) = int64(4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+20)) = v409 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v409)+8)) = v409 + int32(28)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v409
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v425 = v409
	v426 = v424
	goto L141
L150:
	;
	F_subcolorcvec(m, l0, v436, l1, l2)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	goto L3
L152:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v443 != int32(112) {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v471 = F_next(m, l0)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L1
	} else {
		goto L161
	}
L154:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v466 = v446
	goto L153
L155:
	;
	goto L156
L156:
	;
	goto L157
L157:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v457 = F_next(m, l0)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L1
	} else {
		goto L159
	}
L158:
	;
	v466 = v456
	goto L153
L159:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v459 == int32(112) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	goto L158
L161:
	;
	if base.Ui32(v466) <= base.Ui32(v440) {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v476 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v480 != 0 {
		goto L3
	} else {
		goto L168
	}
L165:
	;
	v478 = v476
	goto L167
L166:
	;
	v478 = int32(4)
	goto L167
L167:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v478
	goto L3
L168:
	;
	v483 = (v466 - v440) >> (uint(int32(2)) % 32)
	v490 = int32(_a_F_bracket_4)
	v491 = int32(_a_F_bracket_5)
	v494 = int32(0)
	goto L170
L169:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v546 != 0 {
		goto L3
	} else {
		goto L190
	}
L170:
	;
	v496 = F_strlen(m, v490)
	mBase = m.M
	if v496 == v483 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v542 != 0 {
		goto L187
	} else {
		goto L188
	}
L172:
	;
	if v483 == int32(0) {
		goto L176
	} else {
		goto L177
	}
L173:
	;
	goto L174
L174:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	v537 = v494 + int32(1)
	if v537 != int32(14) {
		v490 = v533
		v491 = v491 + int32(4)
		v494 = v537
		goto L170
	} else {
		goto L186
	}
L175:
	;
	if v530 == int32(0) {
		goto L169
	} else {
		goto L185
	}
L176:
	;
	v530 = int32(0)
	goto L175
L177:
	;
	v502 = v490
	v503 = v440
	v504 = v483
	goto L178
L178:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502))))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v503)))
	if v507 != v508 {
		goto L180
	} else {
		goto L181
	}
L179:
	;
	goto L176
L180:
	;
	v530 = v507 - v508
	goto L175
L181:
	;
	goto L182
L182:
	;
	if v507 == int32(0) {
		goto L176
	} else {
		goto L183
	}
L183:
	;
	v513 = int32(1)
	v518 = v504 - v513
	if v518 != 0 {
		v502 = v502 + v513
		v503 = v503 + int32(4)
		v504 = v518
		goto L178
	} else {
		goto L184
	}
L184:
	;
	goto L179
L185:
	;
	goto L174
L186:
	;
	goto L171
L187:
	;
	v544 = v542
	goto L189
L188:
	;
	v544 = int32(4)
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v544
	goto L3
L190:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v547)+8)) = v548 | int32(1024)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v555 = F_cclasscvec(m, l0, v494, v552&int32(8))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v557 != 0 {
		goto L3
	} else {
		goto L192
	}
L192:
	;
	F_subcolorcvec(m, l0, v555, l1, l2)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L193
	}
L193:
	;
	goto L3
L194:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v571 == int32(0) {
		goto L195
	} else {
		goto L196
	}
L195:
	;
	F_subcolorcvec(m, l0, v569, l1, l2)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L1
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	v576 = F_next(m, l0)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L1
	} else {
		goto L199
	}
L198:
	;
	goto L197
L199:
	;
	goto L3
L200:
	;
	goto L3
L201:
	;
	v588 = v586
	goto L203
L202:
	;
	v588 = int32(15)
	goto L203
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v588
	goto L3
L204:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v594 != 0 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	m.G0 = v12 + int32(16)
	return
L206:
	;
	v595 = int32(0)
	v600 = v595
	v601 = v595
	goto L208
L207:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v628)+20))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v628)+12))
	v631 = int32(24)
	v635 = v629 + v630*v631 + v631
	v636 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v636 != 0 {
		goto L218
	} else {
		goto L219
	}
L208:
	;
	v607 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v601+v12))))
	if v607 == int32(1) {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if v600&int32(1) == int32(0) {
		goto L205
	} else {
		goto L217
	}
L210:
	;
	F_charclasscomplement(m, l0, v601, l1, l2)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v619 = v601 + int32(1)
	if v619 != int32(14) {
		v601 = v619
		goto L208
	} else {
		goto L216
	}
L213:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v612 != 0 {
		goto L205
	} else {
		goto L214
	}
L214:
	;
	v613 = int32(1)
	v615 = v601 + v613
	if v615 != int32(14) {
		v600 = v613
		v601 = v615
		goto L208
	} else {
		goto L215
	}
L215:
	;
	goto L207
L216:
	;
	goto L209
L217:
	;
	goto L207
L218:
	;
	v641 = v636
	goto L221
L219:
	;
	v662 = v629
	goto L220
L220:
	;
	if base.Ui32(v662) < base.Ui32(v635) {
		goto L224
	} else {
		goto L225
	}
L221:
	;
	v646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v646)+20))
	v648 = int32(*(*int16)(unsafe.Add(mBase, uint32(v641)+4)))
	v651 = v647 + v648*int32(24)
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v651)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v651)+20)) = v652 | int32(4)
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v641)+16))
	if v656 != 0 {
		v641 = v656
		goto L221
	} else {
		goto L223
	}
L222:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v657)+20))
	v662 = v658
	goto L220
L223:
	;
	goto L222
L224:
	;
	v673 = v662
	v675 = int32(1)
	goto L227
L225:
	;
	goto L226
L226:
	;
	goto L235
L227:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v673)+20))
	if v679&int32(4) != 0 {
		goto L230
	} else {
		goto L231
	}
L228:
	;
	if v690 == int32(0) {
		goto L205
	} else {
		goto L234
	}
L229:
	;
	v692 = v673 + int32(24)
	if base.Ui32(v692) < base.Ui32(v635) {
		v673 = v692
		v675 = v690
		goto L227
	} else {
		goto L233
	}
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v673)+20)) = v679 & int32(-5)
	v690 = v675
	goto L229
L231:
	;
	goto L232
L232:
	;
	v690 = base.B2i32(v679&int32(3) != int32(0)) & v675
	goto L229
L233:
	;
	goto L228
L234:
	;
	goto L226
L235:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v714 != 0 {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_bracket[0]))
	if v793 != 0 {
		goto L266
	} else {
		goto L267
	}
L237:
	;
	v715 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v714)+12))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v714)+8))
	v722 = int32(*(*int16)(unsafe.Add(mBase, uint32(v714)+4)))
	if v722 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L238:
	;
	goto L239
L239:
	;
	goto L236
L240:
	;
	goto L235
L241:
	;
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v714)+16))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v714)+20))
	if v757 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L242:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v714)))
	v727 = v725 - int32(97)
	if base.B2i32(base.Ui32(int32(17)) < base.Ui32(v727))|base.B2i32(int32(1)<<(uint(v727)%32)&int32(_a_F_bracket_6) == int32(0)) != 0 {
		goto L241
	} else {
		goto L243
	}
L243:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v715)+80))
	if v737 != 0 {
		goto L241
	} else {
		goto L244
	}
L244:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v714)+36))
	if v738 == int32(0) {
		goto L246
	} else {
		goto L247
	}
L245:
	;
	if v750 != 0 {
		goto L249
	} else {
		goto L250
	}
L246:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v715)+52))
	v742 = *(*int32)(unsafe.Add(mBase, uint32(v741)+20))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v714)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v742+v722*int32(24))+12)) = v746
	v750 = v746
	goto L245
L247:
	;
	goto L248
L248:
	;
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v714)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v738)+32)) = v748
	v750 = v748
	goto L245
L249:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v750)+36)) = v738
	goto L251
L250:
	;
	goto L251
L251:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v714)+32)) = int64(0)
	goto L241
L252:
	;
	if v756 != 0 {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v721)+20)) = v756
	goto L252
L254:
	;
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v757)+16)) = v756
	goto L252
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v756)+20)) = v757
	goto L258
L257:
	;
	goto L258
L258:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v721)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v721)+12)) = v763 - int32(1)
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v714)+24))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v714)+28))
	if v768 == int32(0) {
		goto L260
	} else {
		goto L261
	}
L259:
	;
	if v767 != 0 {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v720)+16)) = v767
	goto L259
L261:
	;
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v768)+24)) = v767
	goto L259
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v767)+28)) = v768
	goto L265
L264:
	;
	goto L265
L265:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v720)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v720)+8)) = v774 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v714))) = int32(0)
	v781 = v714 + int32(8)
	v782 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v781)+16)) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v781)+8)) = v782
	*(*int64)(unsafe.Add(mBase, uint32(v781))) = v782
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v715)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v714)+16)) = v788
	*(*int32)(unsafe.Add(mBase, uint32(v715)+32)) = v714
	goto L240
L266:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L1
	} else {
		goto L269
	}
L267:
	;
	goto L268
L268:
	;
	v796 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v797 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v796 <= v797 {
		goto L271
	} else {
		goto L272
	}
L269:
	;
	goto L268
L270:
	;
	F_createarc(m, v791, int32(112), int32(-2), l1, l2)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L1
	} else {
		goto L290
	}
L271:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v799 == int32(0) {
		goto L270
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v820 == int32(0) {
		goto L270
	} else {
		goto L282
	}
L274:
	;
	v802 = v799
	goto L275
L275:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v802)+12))
	if v811 != l2 {
		goto L277
	} else {
		goto L278
	}
L276:
	;
	goto L270
L277:
	;
	v819 = *(*int32)(unsafe.Add(mBase, uint32(v802)+16))
	if v819 != 0 {
		v802 = v819
		goto L275
	} else {
		goto L281
	}
L278:
	;
	v813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v802)+4)))
	if v813 != int32(_a_F_bracket_7) {
		goto L277
	} else {
		goto L279
	}
L279:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v802)))
	if v816 == int32(112) {
		goto L205
	} else {
		goto L280
	}
L280:
	;
	goto L277
L281:
	;
	goto L276
L282:
	;
	v823 = v820
	goto L283
L283:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v823)+8))
	if v832 != l1 {
		goto L285
	} else {
		goto L286
	}
L284:
	;
	goto L270
L285:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v823)+24))
	if v840 != 0 {
		v823 = v840
		goto L283
	} else {
		goto L289
	}
L286:
	;
	v834 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+4)))
	if v834 != int32(_a_F_bracket_7) {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v823)))
	if v837 == int32(112) {
		goto L205
	} else {
		goto L288
	}
L288:
	;
	goto L285
L289:
	;
	goto L284
L290:
	;
	goto L205
L291:
	;
	v882 = v870
	goto L7
L292:
	;
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	switch v889 - int32(73) {
	case 0:
		goto L294
	case 1, 2, 3, 4, 5, 6, 7, 8:
		goto L5
	case 9:
		goto L295
	default:
		goto L296
	}
L293:
	;
	if v1012 == v882 {
		goto L342
	} else {
		goto L343
	}
L294:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v899 = F_next(m, l0)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L1
	} else {
		goto L300
	}
L295:
	;
	v894 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v895 = F_next(m, l0)
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L1
	} else {
		goto L298
	}
L296:
	;
	if v889 != int32(112) {
		goto L5
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v897 != 0 {
		goto L3
	} else {
		goto L299
	}
L299:
	;
	v1012 = v894
	goto L293
L300:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v901 != int32(112) {
		goto L302
	} else {
		goto L303
	}
L301:
	;
	v929 = F_next(m, l0)
	mBase = m.M
	v930 = m.ExcPending
	if v930 != 0 {
		goto L1
	} else {
		goto L309
	}
L302:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v923 = v904
	goto L301
L303:
	;
	goto L304
L304:
	;
	goto L305
L305:
	;
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v915 = F_next(m, l0)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L1
	} else {
		goto L307
	}
L306:
	;
	v923 = v914
	goto L301
L307:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v917 == int32(112) {
		goto L305
	} else {
		goto L308
	}
L308:
	;
	goto L306
L309:
	;
	if base.Ui32(v923) <= base.Ui32(v898) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v934 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v934 != 0 {
		goto L313
	} else {
		goto L314
	}
L311:
	;
	goto L312
L312:
	;
	v938 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v938 != 0 {
		goto L3
	} else {
		goto L316
	}
L313:
	;
	v936 = v934
	goto L315
L314:
	;
	v936 = int32(3)
	goto L315
L315:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v936
	goto L3
L316:
	;
	v939 = v923 - v898
	if v939 == int32(4) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(v898)))
	v1012 = v942
	goto L293
L318:
	;
	goto L319
L319:
	;
	v944 = v939 >> (uint(int32(2)) % 32)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v945)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v945)+8)) = v946 | int32(1024)
	v955 = int32(_a_F_bracket_2)
	v959 = int32(_a_F_bracket_1)
	goto L321
L320:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1007 != 0 {
		goto L3
	} else {
		goto L341
	}
L321:
	;
	v961 = F_strlen(m, v955)
	mBase = m.M
	if v961 == v944 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1003 != 0 {
		goto L338
	} else {
		goto L339
	}
L323:
	;
	if v944 == int32(0) {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	goto L325
L325:
	;
	v998 = *(*int32)(unsafe.Add(mBase, uint32(v959)+8))
	if v998 != 0 {
		v955 = v998
		v959 = v959 + int32(8)
		goto L321
	} else {
		goto L337
	}
L326:
	;
	if v995 == int32(0) {
		goto L320
	} else {
		goto L336
	}
L327:
	;
	v995 = int32(0)
	goto L326
L328:
	;
	v967 = v955
	v968 = v898
	v969 = v944
	goto L329
L329:
	;
	v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v967))))
	v973 = *(*int32)(unsafe.Add(mBase, uint32(v968)))
	if v972 != v973 {
		goto L331
	} else {
		goto L332
	}
L330:
	;
	goto L327
L331:
	;
	v995 = v972 - v973
	goto L326
L332:
	;
	goto L333
L333:
	;
	if v972 == int32(0) {
		goto L327
	} else {
		goto L334
	}
L334:
	;
	v978 = int32(1)
	v983 = v969 - v978
	if v983 != 0 {
		v967 = v967 + v978
		v968 = v968 + int32(4)
		v969 = v983
		goto L329
	} else {
		goto L335
	}
L335:
	;
	goto L330
L336:
	;
	goto L325
L337:
	;
	goto L322
L338:
	;
	v1005 = v1003
	goto L340
L339:
	;
	v1005 = int32(3)
	goto L340
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1005
	goto L3
L341:
	;
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v959)+4)))
	v1012 = v1008
	goto L293
L342:
	;
	v1027 = v882
	v1028 = v882
	goto L6
L343:
	;
	goto L344
L344:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v1019)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+8)) = v1020 | int32(512)
	v1027 = v1012
	v1028 = v882
	goto L6
L345:
	;
	v1038 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1038 != 0 {
		goto L3
	} else {
		goto L346
	}
L346:
	;
	F_subcolorcvec(m, l0, v1036, l1, l2)
	mBase = m.M
	v1040 = m.ExcPending
	if v1040 != 0 {
		goto L1
	} else {
		goto L347
	}
L347:
	;
	goto L3
L348:
	;
	v1045 = v1043
	goto L350
L349:
	;
	v1045 = int32(11)
	goto L350
L350:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1045
	goto L3
}
func F_brinbuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	v9 = m.G0
	v10 = int32(16)
	v11 = v9 - v10
	m.G0 = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v17 = v13 | v14<<(uint(v10)%32)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	if base.Ui32(v18+v19-int32(1)) < base.Ui32(v17) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = v18
	goto L4
L2:
	;
	goto L3
L3:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v77 = F_add_values_to_range(m, l0, v75, v76, l2, l3)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L6
	} else {
		goto L12
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v38 = F_brin_form_tuple(m, v34, v27, v35, v11+int32(12))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l5)+40))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v45 = F_brin_doinsert(m, v40, v41, v42, l5+int32(24), v43, v38, v44)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
	*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v47, float64(1))
	F_pfree(m, v38)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+32)) = v53 + v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l5)+48))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l5)+44))
	F_brin_memtuple_initialize(m, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l5)+32))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l5)+28))
	if base.Ui32(v61+v62-int32(1)) < base.Ui32(v17) {
		v27 = v61
		goto L4
	} else {
		goto L11
	}
L11:
	;
	goto L5
L12:
	;
	m.G0 = v11 + int32(16)
	return
}
func F_brinbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	if l1 != 0 {
		v10 = l1
		return v10
	} else {
		v6 = F_palloc0(m, int32(40))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = v6
			return v10
		}
	}
}
func F_brinvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v16 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v208 = l1
	goto L3
L3:
	;
	m.G0 = v14 + int32(16)
	return v208
L4:
	;
	v22 = F_palloc0(m, int32(40))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v26 = l1
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = F_RelationGetNumberOfBlocksInFork(m, v27, int32(0))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v26 = v22
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v29
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+56))
	v35 = F_IndexGetRelation(m, v33, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v38 = F_table_open(m, v35, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v42
	v45 = F_RelationGetNumberOfBlocksInFork(m, v41, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v45
	v49 = int32(0)
	v54 = F_read_stream_begin_relation(m, int32(13), v40, v41, v49, int32(3), v14+int32(8), v49)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L14
L14:
	;
	v68 = F_read_stream_next_buffer(m, v54, int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	F_read_stream_end(m, v54)
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L62
	}
L16:
	;
	if v68 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[0]))
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	goto L15
L20:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v68 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	goto L22
L24:
	;
	F_ReleaseBuffer(m, v68)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L7
	} else {
		goto L61
	}
L25:
	;
	v92 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+14)))
	if v92 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[1]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v77+(v68^int32(-1))<<(uint(int32(2))%32))))
	v91 = v83
	goto L25
L27:
	;
	goto L28
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[2]))
	v91 = v85 + v68<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	F_LockRelationForExtension(m, v41, int32(5))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L7
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if v68 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L32:
	;
	F_UnlockRelationForExtension(m, v41, int32(5))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	F_LockBufferInternal(m, v68, int32(3))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L7
	} else {
		goto L34
	}
L34:
	;
	v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+14)))
	if v104 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_brin_initialize_empty_new_buffer(m, v41, v68)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_UnlockBuffer(m, v68)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	F_UnlockBuffer(m, v68)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L7
	} else {
		goto L39
	}
L39:
	;
	goto L24
L40:
	;
	goto L31
L41:
	;
	goto L24
L42:
	;
	if v142 == int32(_a_F_brinvacuumcleanup_0) {
		goto L41
	} else {
		goto L48
	}
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[1]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v68^int32(-1))<<(uint(int32(2))%32))))
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+16)))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122+v123)+6)))
	if v125 != int32(_a_F_brinvacuumcleanup_1) {
		v142 = v125
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v129 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[2]))
	v132 = v129 + v68<<(uint(int32(13))%32)
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132-int32(_a_F_brinvacuumcleanup_2)))))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132+v135-int32(_a_F_brinvacuumcleanup_3)))))
	if v139 == int32(_a_F_brinvacuumcleanup_1) {
		goto L41
	} else {
		goto L47
	}
L46:
	;
	goto L41
L47:
	;
	v142 = v139
	goto L42
L48:
	;
	if v68 < int32(0) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v164 = int32(0)
	v165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+16)))
	v166 = v91 + v165
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+6)))
	if v167 != int32(_a_F_brinvacuumcleanup_4) {
		v182 = v164
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[3]))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148+(v68^int32(-1))*int32(56))+16))
	v163 = v154
	goto L49
L51:
	;
	goto L52
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_brinvacuumcleanup[4]))
	v157 = int32(56)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v156+v68*v157-v157)+16))
	v163 = v162
	goto L49
L53:
	;
	F_RecordPageWithFreeSpace(m, v41, v163, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L60
	}
L54:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+4)))
	if v170&int32(1) != 0 {
		v182 = v164
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v173 = int32(4)
	v174 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+14)))
	v175 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v91)+12)))
	v176 = v174 - v175
	if v176 <= v173 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	v182 = v179 - int32(4)
	goto L53
L57:
	;
	v179 = v173
	goto L59
L58:
	;
	v179 = v176
	goto L59
L59:
	;
	goto L56
L60:
	;
	goto L41
L61:
	;
	goto L14
L62:
	;
	F_FreeSpaceMapVacuum(m, v41)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v201 = v26 + int32(8)
	F_brinsummarize(m, v197, v38, int32(-1), int32(0), v201, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v208 = v26
	goto L3
}
func F_btadjustmembers(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	var v7 int32
	_ = v7
	Fn14310(m, l0, l1, l2, l3, int32(403))
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_btbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int64
	_ = v26
	var v34 int32
	_ = v34
	v5 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = F_palloc(m, int32(_a_F_btbeginscan_0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = int64(-4294967296)
			*(*int64)(unsafe.Add(mBase, uint32(v10)+uint32(_c_F_btbeginscan[0]))) = v12
			*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v12
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
			if v16 <= int32(0) {
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
				v26 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v26
				*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v26
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v10
				*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v34
				return v5
			} else {
				v22 = F_palloc(m, v16*int32(56))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = v22
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v24
					v26 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v10)+44)) = v26
					*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v26
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v26
					*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v26
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+36)) = v10
					*(*int32)(unsafe.Add(mBase, uint32(v5)+48)) = v34
					return v5
				}
			}
		}
	}
}
func F_btboolskipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(194)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(195)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v2
	return v2
}
func F_btendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
	if v4 == int32(-1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btendscan[0])))
	if v22 != 0 {
		goto L10
	} else {
		goto L11
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+36))
	if int32(0) < v7 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F__bt_killitems(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+56))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	goto L5
L8:
	;
	F_ReleaseBuffer(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+56)) = int32(0)
	goto L1
L10:
	;
	F_ReleaseBuffer(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L6
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v3)+8))
	if v27 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btendscan[0]))) = int32(0)
	goto L12
L14:
	;
	F_pfree(m, v27)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v3)+28))
	if v30 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_MemoryContextDelete(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	if v33 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_pfree(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v3)+44))
	if v36 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	F_pfree(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	F_pfree(m, v3)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	return
}
func F_btfloat48cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v8 float64
	_ = v8
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 float64
	_ = v12
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v25 int32
	_ = v25
	var v36 int64
	_ = v36
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = base.F64_promote_f32(v7)
	v10 = int64(9223372036854775807)
	v11 = base.I64_reinterpret_f64(v8) & v10
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = base.I64_reinterpret_f64(v12) & v10
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v15) {
		v25 = base.B2i32(base.Ui64(v11) < base.Ui64(int64(9218868437227405313)))
		v36 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v8, v12))&v25)
	} else {
		v20 = int64(1)
		if base.F64_gt(v8, v12) != 0 {
			v36 = v20
		} else {
			if base.Ui64(int64(9218868437227405312)) < base.Ui64(v11) {
				v36 = v20
			} else {
				v25 = int32(1)
				v36 = int64(0) - base.I64_extend_i32_u((base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v15))|base.F64_lt(v8, v12))&v25)
			}
		}
	}
	return v36
}
func F_btgettuple(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	v3 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)) = uint8(v3)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	goto L2
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)+60))
	if v14 != int32(-1) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v65 == int32(0) {
		goto L1
	} else {
		goto L22
	}
L4:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v58 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L5:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+30)))
	if v17 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v50 = F__bt_first(m, l0, l1)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L13
	} else {
		goto L18
	}
L8:
	;
	v44 = F__bt_next(m, l0, l1)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L13
	} else {
		goto L16
	}
L9:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	if v20 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v25 = F_palloc_mul(m, int32(4), int32(1358))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v30 = v20
	goto L12
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if int32(1357) < v31 {
		goto L8
	} else {
		goto L15
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v25
	v30 = v25
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v31 + int32(1)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v30+v31<<(uint(int32(2))%32)))) = v40
	goto L8
L16:
	;
	if v44 == int32(0) {
		goto L4
	} else {
		goto L17
	}
L17:
	;
	return int32(1)
L18:
	;
	if v50 == int32(0) {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	return int32(1)
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v62 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v61)+18)) = uint16(v62)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+17)))
	if v64 != 0 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	goto L3
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+17)))
	if v69 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+24))
	v71 = v65 + v70
	v73 = v71 + int32(12)
	v75 = F_LWLockAcquire(m, v73, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v77 != int32(4) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = int32(4)
	F_LWLockRelease(m, v73)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_LWLockRelease(m, v73)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L13
	} else {
		goto L30
	}
L28:
	;
	F_ConditionVariableBroadcast(m, v71+int32(28))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L1
}
func F_btnametextcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = F_strlen(m, v5)
		mBase = m.M
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v16 = v14 & v12
		if v16 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		if v14 == int32(1) {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v16 != 0 {
				v45 = int32(base.Ui32(v14)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v47 = F_varstr_cmp(m, v5, v11, v7+v17, v45, v46)
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			if v49 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_s(v47)
				}
			} else {
				return base.I64_extend_i32_s(v47)
			}
		}
	}
}
func F_btoidcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(base.Ui32(v4) < base.Ui32(v3)) - base.B2i32(base.Ui32(v3) < base.Ui32(v4)))
}
func F_btoidskipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v6 int32
	_ = v6
	v2 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(206)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(207)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = v2
	return v2
}
func F_btoptions(m *base.Module, l0 int64, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(4), int32(24), int32(_a_F_btoptions_0), int32(3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_btrecordimagecmp(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v2)
	}
}
func F_bttext_pattern_sortsupport(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14237(m, l0, int32(25))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_bttextnamecmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = int32(1)
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v16 != 0 {
			v17 = v12
		} else {
			v17 = int32(4)
		}
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v14 == int32(1) {
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
			if v25 == int32(18) {
				v28 = int32(16)
			} else {
				v28 = int32(0)
			}
			if base.Ui32((v25-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v35 = int32(4)
			} else {
				v35 = v28
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v16 != 0 {
				v46 = int32(base.Ui32(v14)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = F_strlen(m, v19)
		mBase = m.M
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v49 = F_varstr_cmp(m, v8+v17, v46, v19, v47, v48)
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int64(0)
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v51 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int64(0)
				} else {
					return base.I64_extend_i32_s(v49)
				}
			} else {
				return base.I64_extend_i32_s(v49)
			}
		}
	}
}
func F_build_coercion_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	if l2 != 0 {
		v18 = F_SearchSysCache1(m, int32(47), base.I64_extend_i32_u(l2))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v150 = m.ExcPending
				if v150 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = l2
					F_errmsg_internal(m, int32(_a_F_build_coercion_expression_0), v14+int32(16))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_build_coercion_expression_1), int32(858), int32(_a_F_build_coercion_expression_2))
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+22)))
				v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24+v25)+104)))
				F_ReleaseCatCache(m, v18)
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v31 = v27
					switch l1 - int32(1) {
					case 0:
						*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l0
						v39 = F_list_make1_impl(m, int32(1), v14+int32(12))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v31 < int32(2) {
								v70 = v39
								v72 = F_makeFuncExpr(m, l2, l3, v70, int32(0), l6)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = l7
									v141 = v72
									m.G0 = v14 + int32(32)
									return v141
								}
							} else {
								v45 = int32(0)
								v50 = F_makeConst(m, int32(23), int32(-1), v45, int32(4), base.I64_extend_i32_s(l4), v45, int32(1))
								mBase = m.M
								v51 = m.ExcPending
								if v51 != 0 {
									return int32(0)
								} else {
									v52 = F_lappend(m, v39, v50)
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										if v31 != int32(3) {
											v70 = v52
											v72 = F_makeFuncExpr(m, l2, l3, v70, int32(0), l6)
											mBase = m.M
											v73 = m.ExcPending
											if v73 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = l7
												v141 = v72
												m.G0 = v14 + int32(32)
												return v141
											}
										} else {
											v58 = int32(0)
											v59 = int32(1)
											v65 = F_makeConst(m, int32(16), int32(-1), v58, v59, base.I64_extend_i32_u(base.B2i32(l5 == int32(3))), v58, v59)
											mBase = m.M
											v66 = m.ExcPending
											if v66 != 0 {
												return int32(0)
											} else {
												v67 = F_lappend(m, v52, v65)
												mBase = m.M
												v68 = m.ExcPending
												if v68 != 0 {
													return int32(0)
												} else {
													v70 = v67
													v72 = F_makeFuncExpr(m, l2, l3, v70, int32(0), l6)
													mBase = m.M
													v73 = m.ExcPending
													if v73 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = l7
														v141 = v72
														m.G0 = v14 + int32(32)
														return v141
													}
												}
											}
										}
									}
								}
							}
						}
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
							F_errmsg_internal(m, int32(_a_F_build_coercion_expression_3), v14)
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_build_coercion_expression_1), int32(1001), int32(_a_F_build_coercion_expression_2))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 2:
						v76 = F_palloc0(m, int32(32))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v76))) = int32(29)
							v81 = F_palloc0(m, int32(16))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v81))) = int32(34)
								v85 = F_exprTypmod(m, l0)
								mBase = m.M
								v86 = m.ExcPending
								if v86 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v85
									v88 = F_exprType(m, l0)
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return int32(0)
									} else {
										v92 = F_getBaseTypeAndTypmod(m, v88, v14+int32(24))
										mBase = m.M
										v93 = m.ExcPending
										if v93 != 0 {
											return int32(0)
										} else {
											v94 = F_get_element_type(m, v92)
											mBase = m.M
											v95 = m.ExcPending
											if v95 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v94
												v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
												*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v97
												v101 = F_get_element_type(m, l3)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
													v105 = F_coerce_to_target_type(m, int32(0), v81, v104, v101, l4, l5, l6, l7)
													mBase = m.M
													v106 = m.ExcPending
													if v106 != 0 {
														return int32(0)
													} else {
														if v105 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v165 = m.ExcPending
															if v165 != 0 {
																return int32(0)
															} else {
																F_errmsg_internal(m, int32(_a_F_build_coercion_expression_4), int32(0))
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_build_coercion_expression_1), int32(966), int32(_a_F_build_coercion_expression_2))
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = l3
															*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v105
															*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = l0
															v112 = F_exprTypmod(m, v105)
															mBase = m.M
															v113 = m.ExcPending
															if v113 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = l7
																*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = l6
																*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v112
																v141 = v76
																m.G0 = v14 + int32(32)
																return v141
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
					case 3:
						v131 = F_palloc0(m, int32(24))
						mBase = m.M
						v132 = m.ExcPending
						if v132 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = l7
							*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = l6
							*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(28)
							v141 = v131
							m.G0 = v14 + int32(32)
							return v141
						}
					}
				}
			}
		}
	} else {
		v31 = int32(0)
		switch l1 - int32(1) {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l0
			v39 = F_list_make1_impl(m, int32(1), v14+int32(12))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				if v31 < int32(2) {
					v70 = v39
					v72 = F_makeFuncExpr(m, l2, l3, v70, int32(0), l6)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = l7
						v141 = v72
						m.G0 = v14 + int32(32)
						return v141
					}
				} else {
					v45 = int32(0)
					v50 = F_makeConst(m, int32(23), int32(-1), v45, int32(4), base.I64_extend_i32_s(l4), v45, int32(1))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = F_lappend(m, v39, v50)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v31 != int32(3) {
								v70 = v52
								v72 = F_makeFuncExpr(m, l2, l3, v70, int32(0), l6)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = l7
									v141 = v72
									m.G0 = v14 + int32(32)
									return v141
								}
							} else {
								v58 = int32(0)
								v59 = int32(1)
								v65 = F_makeConst(m, int32(16), int32(-1), v58, v59, base.I64_extend_i32_u(base.B2i32(l5 == int32(3))), v58, v59)
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									v67 = F_lappend(m, v52, v65)
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v70 = v67
										v72 = F_makeFuncExpr(m, l2, l3, v70, int32(0), l6)
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v72)+32)) = l7
											v141 = v72
											m.G0 = v14 + int32(32)
											return v141
										}
									}
								}
							}
						}
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v14))) = l1
				F_errmsg_internal(m, int32(_a_F_build_coercion_expression_3), v14)
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_build_coercion_expression_1), int32(1001), int32(_a_F_build_coercion_expression_2))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 2:
			v76 = F_palloc0(m, int32(32))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v76))) = int32(29)
				v81 = F_palloc0(m, int32(16))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v81))) = int32(34)
					v85 = F_exprTypmod(m, l0)
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v85
						v88 = F_exprType(m, l0)
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							v92 = F_getBaseTypeAndTypmod(m, v88, v14+int32(24))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int32(0)
							} else {
								v94 = F_get_element_type(m, v92)
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v81)+4)) = v94
									v97 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
									*(*int32)(unsafe.Add(mBase, uint32(v81)+12)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v81)+8)) = v97
									v101 = F_get_element_type(m, l3)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										v104 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
										v105 = F_coerce_to_target_type(m, int32(0), v81, v104, v101, l4, l5, l6, l7)
										mBase = m.M
										v106 = m.ExcPending
										if v106 != 0 {
											return int32(0)
										} else {
											if v105 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v165 = m.ExcPending
												if v165 != 0 {
													return int32(0)
												} else {
													F_errmsg_internal(m, int32(_a_F_build_coercion_expression_4), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_build_coercion_expression_1), int32(966), int32(_a_F_build_coercion_expression_2))
														mBase = m.M
														v174 = m.ExcPending
														if v174 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v76)+12)) = l3
												*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v105
												*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = l0
												v112 = F_exprTypmod(m, v105)
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = l7
													*(*int32)(unsafe.Add(mBase, uint32(v76)+24)) = l6
													*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v112
													v141 = v76
													m.G0 = v14 + int32(32)
													return v141
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
		case 3:
			v131 = F_palloc0(m, int32(24))
			mBase = m.M
			v132 = m.ExcPending
			if v132 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v131)+20)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v131)+16)) = l6
				*(*int32)(unsafe.Add(mBase, uint32(v131)+8)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = l0
				*(*int32)(unsafe.Add(mBase, uint32(v131))) = int32(28)
				v141 = v131
				m.G0 = v14 + int32(32)
				return v141
			}
		}
	}
}
func F_build_tlist_index(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v82 int32
	_ = v82
	if l0 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = int32(12)
		v15 = v9*v10 + v10
	} else {
		v15 = int32(12)
	}
	v16 = F_palloc(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v16)+8)) = uint16(v20)
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
		v24 = v16 + int32(12)
		if l0 == v20 {
			v73 = v24
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v27 <= int32(0) {
				v73 = v24
			} else {
				v31 = v24
				v34 = int32(0)
				for {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v34<<(uint(int32(2))%32))))
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
					if v43 == int32(0) {
						v64 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)) = uint8(v64)
						v66 = v31
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						if v46 != int32(321) {
							if v46 != int32(6) {
								v64 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v16)+9)) = uint8(v64)
								v66 = v31
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v31))) = v51
								v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+8)))
								*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v53)
								v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+8)))
								*(*uint16)(unsafe.Add(mBase, uint32(v31)+6)) = uint16(v55)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v43)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v57
								v66 = v31 + int32(12)
							}
						} else {
							v61 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v61)
							v66 = v31
						}
					}
					v69 = v34 + int32(1)
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v69 < v70 {
						v31 = v66
						v34 = v69
						continue
					} else {
						break
					}
					break
				}
				v73 = v66
			}
		}
		v82 = base.I32_div_s(v73-v24, int32(12))
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v82
		return v16
	}
}
func F_byte_increment(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v3 != int32(255) {
		v7 = v3 + int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v7)
	} else {
	}
	return base.B2i32(v3 != int32(255))
}
func F_byteagt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v160 = int32(0)
	return base.I64_extend_i32_u(base.B2i32(v151 == v160)&base.B2i32(v75 < v45) | base.B2i32(v160 < v151))
L60:
	;
	goto L59
}
func F_byteain(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_strlen(m, v15)
	mBase = m.M
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v17 != int32(92) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return base.I64_extend_i32_u(v136)
L2:
	;
	v47 = F_palloc(m, v16+int32(4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L8
	}
L3:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 != int32(120) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v24 = v16 - int32(2)
	v29 = F_palloc(m, int32(base.Ui32(v24)>>(uint(int32(1))%32))+int32(4))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	return int64(0)
L6:
	;
	v37 = F_hex_decode_safe(m, v15+int32(2), v24, v29+int32(4), v14)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = base.I32_wrap_i64(v37)<<(uint(int32(2))%32) + int32(16)
	v136 = v29
	goto L1
L8:
	;
	v50 = v47 + int32(4)
	v51 = v15
	v52 = v50
	goto L10
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = (v52-v50)<<(uint(int32(2))%32) + int32(16)
	v136 = v47
	goto L1
L10:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v60 != int32(92) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v110 = int32(0)
	v111 = F_errsave_start(m, v14)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L5
	} else {
		goto L23
	}
L12:
	;
	if v60 == int32(0) {
		goto L9
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
	if v70&int32(252) == int32(48) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v60)
	v66 = int32(1)
	v51 = v51 + v66
	v52 = v52 + v66
	goto L10
L16:
	;
	goto L11
L17:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+2)))
	if v75&int32(248) != int32(48) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v70 != int32(92) {
		goto L16
	} else {
		goto L22
	}
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+3)))
	if v80&int32(248) != int32(48) {
		goto L16
	} else {
		goto L21
	}
L21:
	;
	v94 = v80 + (v75<<(uint(int32(3))%32)&int32(56) | v70<<(uint(int32(6))%32)) - int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v94)
	v51 = v51 + int32(4)
	v52 = v52 + int32(1)
	goto L10
L22:
	;
	v102 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v102)
	v51 = v51 + int32(2)
	v52 = v52 + int32(1)
	goto L10
L23:
	;
	if v111 == int32(0) {
		v136 = v110
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(_a_F_byteain_0)
	F_errmsg(m, int32(_a_F_byteain_1), v12)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L5
	} else {
		goto L26
	}
L26:
	;
	F_errsave_finish(m, v14, int32(_a_F_byteain_2), int32(258), int32(_a_F_byteain_3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v136 = v110
	goto L1
}
func F_bytealike(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = F_pg_detoast_datum_packed(m, v12)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
			if v15 == int32(1) {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+1)))
				if v21 == int32(18) {
					v24 = int32(16)
				} else {
					v24 = int32(0)
				}
				if base.Ui32((v21-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v31 = int32(4)
				} else {
					v31 = v24
				}
				v44 = v31
			} else {
				v32 = int32(1)
				if v15&v32 != 0 {
					v44 = int32(base.Ui32(v15)>>(uint(v32)%32)) - v32
				} else {
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if v45 == int32(1) {
				v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
				if v51 == int32(18) {
					v54 = int32(16)
				} else {
					v54 = int32(0)
				}
				if base.Ui32((v51-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v61 = int32(4)
				} else {
					v61 = v54
				}
				v74 = v61
			} else {
				v62 = int32(1)
				if v45&v62 != 0 {
					v74 = int32(base.Ui32(v45)>>(uint(v62)%32)) - v62
				} else {
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v75 = int32(1)
			if v15&v75 != 0 {
				v79 = v75
			} else {
				v79 = int32(4)
			}
			v81 = int32(1)
			if v45&v81 != 0 {
				v85 = v81
			} else {
				v85 = int32(4)
			}
			v88 = F_SB_MatchText(m, v8+v79, v44, v13+v85, v74, int32(0))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(base.B2i32(v88 == int32(1)))
			}
		}
	}
}
