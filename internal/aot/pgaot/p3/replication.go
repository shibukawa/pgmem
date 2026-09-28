package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ReplicationOriginShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemRequest[0]))
	if v8 != 0 {
		v11 = F_add_size(m, int32(0), int32(8))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationOriginShmemRequest[0]))
			v16 = F_mul_size(m, v14, int32(64))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				v18 = F_add_size(m, v11, v16)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_ReplicationOriginShmemRequest_0)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_ReplicationOriginShmemRequest_1)
					F_ShmemRequestStructWithOpts(m, v5)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						m.G0 = v5 + int32(16)
						return
					}
				}
			}
		}
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ReplicationSlotDropAtPubNode(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v269 int64
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(240)
	m.G0 = v11
	v18 = v4
	v19 = v4
	v20 = int32(-1)
	v21 = v4
	goto L1
L1:
	;
	goto L4
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	goto L2
L4:
	;
	if v20 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v268 = int32(m.ExcTag)
	v269 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v268 == int32(0) {
		goto L70
	} else {
		goto L71
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	F_load_file(m, int32(_a_F_ReplicationSlotDropAtPubNode_0), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v112 = v18
	v113 = v19
	v115 = v21
	goto L9
L9:
	;
	if v115 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	v33 = v11 + int32(216)
	F_initStringInfo(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	F_appendStringInfoString(m, v33, int32(_a_F_ReplicationSlotDropAtPubNode_1))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	F_appendStringInfoChar(m, v33, int32(34))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = l1
	v54 = v46
	goto L17
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	v87 = v11 + int32(216)
	F_appendStringInfoChar(m, v87, int32(34))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L6
	} else {
		goto L25
	}
L17:
	;
	if v54&int32(255) == int32(34) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	F_appendStringInfoChar(m, v11+int32(216), int32(34))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	F_appendStringInfoChar(m, v11+int32(216), base.I32_extend8_s(v54))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v74 = v53 + int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v75 != 0 {
		v53 = v74
		v54 = v75
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v18
	F_appendStringInfoString(m, v87, int32(_a_F_ReplicationSlotDropAtPubNode_2))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0]))
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1]))
	goto L27
L27:
	;
	v102 = v11 + int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(v102)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v11 + int32(44)
	goto L30
L28:
	;
	v112 = v98
	v113 = v100
	v115 = int32(0)
	goto L9
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v11 + int32(48)
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[2]))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	v128 = int32(0)
	v130 = m.T0[v124].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v127, v128, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L6
	} else {
		goto L37
	}
L32:
	;
	goto L33
L33:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0])) = v112
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	F_pfree(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L6
	} else {
		goto L68
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L6
	} else {
		goto L64
	}
L35:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	if v184 != 0 {
		goto L50
	} else {
		goto L51
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_errfinish(m, int32(_a_F_ReplicationSlotDropAtPubNode_3), v176, int32(_a_F_ReplicationSlotDropAtPubNode_4))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L6
	} else {
		goto L49
	}
L37:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if v132 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	v139 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L6
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if base.B2i32(l2 == int32(0))|v132 != 0 {
		goto L34
	} else {
		goto L44
	}
L41:
	;
	if v139 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F_errmsg(m, int32(_a_F_ReplicationSlotDropAtPubNode_5), v11)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L6
	} else {
		goto L43
	}
L43:
	;
	v176 = int32(2727)
	goto L36
L44:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	if v153 != int32(67137668) {
		goto L34
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	v160 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L6
	} else {
		goto L46
	}
L46:
	;
	if v160 == int32(0) {
		goto L35
	} else {
		goto L47
	}
L47:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	F_errmsg(m, int32(_a_F_ReplicationSlotDropAtPubNode_6), v11+int32(16))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L6
	} else {
		goto L48
	}
L48:
	;
	v176 = int32(2736)
	goto L36
L49:
	;
	goto L35
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_pfree(m, v184)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L6
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	if v189 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_tuplestore_end(m, v189)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	if v194 != 0 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L56
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_FreeTupleDesc(m, v194)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L6
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_pfree(m, v130)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L6
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0])) = v112
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v11)+216))
	F_pfree(m, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L6
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0])) = v112
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v113
	m.G0 = v11 + int32(240)
	return
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_errcode(m, int32(100663808))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L65
	}
L65:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v230
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	F_errmsg(m, int32(_a_F_ReplicationSlotDropAtPubNode_6), v11+int32(32))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_errfinish(m, int32(_a_F_ReplicationSlotDropAtPubNode_3), int32(2744), int32(_a_F_ReplicationSlotDropAtPubNode_4))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	goto L3
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+236)) = v113
	*(*int32)(unsafe.Add(mBase, uint32(v11)+232)) = v112
	F_pg_re_throw(m)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L6
	} else {
		goto L69
	}
L69:
	;
	goto L5
L70:
	;
	v273 = int32(v269)
	m.G0 = v11
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v276)))
	if v11+int32(44) == v279 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	m.ExcPending = 1
	goto L79
L72:
	;
	if v283 != 0 {
		goto L76
	} else {
		goto L77
	}
L73:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v276)+4))
	v283 = v281
	goto L75
L74:
	;
	v283 = int32(0)
	goto L75
L75:
	;
	goto L72
L76:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v11)+236))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v11)+232))
	v18 = v285
	v19 = v284
	v20 = v283
	v21 = v275
	goto L1
L77:
	;
	goto L78
L78:
	;
	F___wasm_longjmp(m, v276, v275)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	return
L80:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReplicationSlotRelease(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[1])))
	if v16 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = F_pstrdup(m, v13+int32(24))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v23 = int32(0)
	goto L3
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	if v24 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	v23 = v21
	goto L3
L6:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[2]))
	v117 = F_LWLockAcquire(m, v113+int32(512), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L41
	}
L7:
	;
	v27 = int32(_a_F_ReplicationSlotRelease_0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0])) = int32(0)
	F_ReplicationSlotDropPtr(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	if v14 == int32(0) {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	F_RequestDisableLogicalDecoding(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L6
L13:
	;
	v59 = m.G0
	v60 = int32(16)
	v61 = v59 - v60
	m.G0 = v61
	F_gettimeofday(m, v61)
	mBase = m.M
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v61)))
	v65 = int64(*(*int32)(unsafe.Add(mBase, uint32(v61)+8)))
	m.G0 = v61 + v60
	v73 = v65 + v64*int64(1000000) - int64(946684800000000)
	goto L21
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v44 = base.AtomicRmwXchg32(m, v13, int32(0), int32(1))
	if v44 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v13, int32(_a_F_ReplicationSlotRelease_5))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v48
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v48))
	F_ReplicationSlotsComputeRequiredXmin(m, v48)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	goto L13
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v76 = int32(0)
	v77 = base.AtomicRmwXchg32(m, v13, v76, int32(1))
	if v74 == v76 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0])) = int32(0)
	goto L6
L23:
	;
	if v77 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v77 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	F_s_lock(m, v13, int32(_a_F_ReplicationSlotRelease_5))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(-1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
	if v85 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+272)) = v73
	goto L32
L31:
	;
	goto L32
L32:
	;
	v89 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v89))
	F_ConditionVariableBroadcast(m, v13+int32(224))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L22
L34:
	;
	F_s_lock(m, v13, int32(_a_F_ReplicationSlotRelease_5))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
	if v99 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+272)) = v73
	goto L40
L39:
	;
	goto L40
L40:
	;
	v103 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v103))
	goto L22
L41:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[3]))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+36)))
	v123 = v121 & int32(-17)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+36)) = uint8(v123)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[4]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)+32))
	*(*uint8)(unsafe.Add(mBase, uint32(v127+v128))) = uint8(v123)
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[2]))
	F_LWLockRelease(m, v132+int32(512))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[1])))
	if v138 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[5])))
	if v144 != 0 {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	goto L45
L45:
	;
	m.G0 = v10 + int32(16)
	return
L46:
	;
	v145 = int32(15)
	goto L48
L47:
	;
	v145 = int32(14)
	goto L48
L48:
	;
	v147 = F_errstart(m, v145, int32(0))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v147 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v23
	if v14 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	goto L52
L52:
	;
	F_pfree(m, v23)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	v152 = int32(_a_F_ReplicationSlotRelease_1)
	goto L55
L54:
	;
	v152 = int32(_a_F_ReplicationSlotRelease_2)
	goto L55
L55:
	;
	F_errmsg(m, v152, v10)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotRelease_3), int32(847), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L57
	}
L57:
	;
	goto L52
L58:
	;
	goto L45
}
func F_copy_replication_slot(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int64
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int64
	_ = v180
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int64
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int64
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int64
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(704)
	m.G0 = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = F_get_call_result_type(m, l0, v3, v16+int32(88))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L7
	} else {
		goto L138
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L133
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L7
	} else {
		goto L129
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L7
	} else {
		goto L125
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L121
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L114
	}
L7:
	;
	return int64(0)
L8:
	;
	if v23 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_CheckSlotPermissions(m)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L7
	} else {
		goto L111
	}
L12:
	;
	if l1 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	v42 = F_LWLockAcquire(m, v38+int32(_a_F_copy_replication_slot_0), int32(1))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L7
	} else {
		goto L19
	}
L14:
	;
	F_CheckLogicalDecodingRequirements(m, int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_CheckSlotRequirements(m, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L18
	}
L17:
	;
	goto L13
L18:
	;
	goto L13
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[1]))
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[2]))
	v48 = v45 + v47
	if int32(0) < v48 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if l1 != 0 {
		goto L57
	} else {
		goto L58
	}
L21:
	;
	v180 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v182 = base.B2i32(v180 != int64(0))
	if v134 == int32(3) {
		v186 = v182
		v187 = v133
		goto L20
	} else {
		goto L55
	}
L22:
	;
	v53 = base.I32_wrap_i64(v18)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	v60 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	F_LWLockRelease(m, v157+int32(_a_F_copy_replication_slot_0))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L50
	}
L25:
	;
	v71 = v55 + v60*int32(296)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+4)))
	if v72 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v141 = v60 + int32(1)
	if v141 != v48 {
		v60 = v141
		goto L25
	} else {
		goto L49
	}
L28:
	;
	v76 = v71 + int32(24)
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v79 == int32(0))|base.B2i32(v79 != v82) != 0 {
		v100 = v79
		v101 = v82
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v100-v101 != 0 {
		goto L27
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v85 = v76
	v86 = v19
	goto L32
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	if v90 == int32(0) {
		v100 = v90
		v101 = v89
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v100 = v90
	v101 = v89
	goto L30
L34:
	;
	v93 = int32(1)
	if v90 == v89 {
		v85 = v85 + v93
		v86 = v86 + v93
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v105 = base.AtomicRmwXchg32(m, v71, int32(0), int32(1))
	if v105 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_s_lock(m, v71, int32(_a_F_copy_replication_slot_1))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	base.MemoryCopy(m, v16+int32(408), v71, int32(296))
	v113 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v71))), uint32(v113))
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	F_LWLockRelease(m, v117+int32(_a_F_copy_replication_slot_0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v16)+496))
	if l1 != base.B2i32(v122 != int32(0)) {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v126 = *(*int64)(unsafe.Add(mBase, uint32(v16)+512))
	if v126 == int64(0) {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v16)+520))
	if v129 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	if l1 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v133 = v16 + int32(545)
	goto L47
L46:
	;
	v133 = int32(0)
	goto L47
L47:
	;
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(3) <= v134 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v16)+500))
	v186 = base.B2i32(v137 == int32(2))
	v187 = v133
	goto L20
L49:
	;
	goto L26
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_2), v16+int32(80))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(698), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L7
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	v186 = v182
	v187 = v185
	goto L20
L56:
	;
	v236 = base.AtomicRmwXchg32(m, v71, int32(0), int32(1))
	if v236 != 0 {
		goto L73
	} else {
		goto L74
	}
L57:
	;
	v188 = int32(1)
	if v186 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v217 = int32(0)
	if v186 != 0 {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v191 = int32(2)
	goto L62
L61:
	;
	v191 = v188
	goto L62
L62:
	;
	v192 = int32(0)
	F_ReplicationSlotCreate(m, v53, v188, v191, v192, v192, v192, v192)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	F_EnsureLogicalDecodingEnabled(m)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = int32(414)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = int32(415)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = int32(416)
	v206 = int32(0)
	v213 = F_CreateInitDecodingContext(m, v187, v206, v206, v126, v16+int32(112), v206, v206, v206)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	F_FreeDecodingContext(m, v213)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	goto L56
L67:
	;
	v220 = int32(2)
	goto L69
L68:
	;
	v220 = v217
	goto L69
L69:
	;
	v221 = int32(0)
	F_ReplicationSlotCreate(m, v53, v217, v220, v221, v221, v221, v221)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v228)+104)) = v126
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L72
	}
L72:
	;
	goto L56
L73:
	;
	F_s_lock(m, v71, int32(_a_F_copy_replication_slot_1))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	base.MemoryCopy(m, v16+int32(112), v71, int32(296))
	v244 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v71))), uint32(v244))
	v247 = *(*int64)(unsafe.Add(mBase, uint32(v16)+216))
	if base.Ui64(v247) < base.Ui64(v126) {
		goto L3
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	v250 = int32(0)
	if base.B2i32(v249 == v250) == base.B2i32(v122 != v250) {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v255 = *(*int64)(unsafe.Add(mBase, uint32(v16)+232))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v16)+208))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v16)+132))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v16)+128))
	v261 = v16 + int32(136)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261))))
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v264 == int32(0))|base.B2i32(v264 != v267) != 0 {
		v285 = v264
		v286 = v267
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v285-v286 != 0 {
		goto L3
	} else {
		goto L86
	}
L80:
	;
	goto L79
L81:
	;
	v270 = v261
	v271 = v19
	goto L82
L82:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271)+1)))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	if v275 == int32(0) {
		v285 = v275
		v286 = v274
		goto L80
	} else {
		goto L84
	}
L83:
	;
	v285 = v275
	v286 = v274
	goto L80
L84:
	;
	v278 = int32(1)
	if v275 == v274 {
		v270 = v270 + v278
		v271 = v271 + v278
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	if v255 == int64(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v291 = v122
	goto L89
L88:
	;
	v291 = int32(0)
	goto L89
L89:
	;
	if v291 != 0 {
		goto L2
	} else {
		goto L90
	}
L90:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v16)+224))
	if v292 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v294 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[4]))
	v297 = base.AtomicRmwXchg32(m, v294, int32(0), int32(1))
	if v297 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	F_s_lock(m, v294, int32(_a_F_copy_replication_slot_1))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v302 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[4]))
	*(*int64)(unsafe.Add(mBase, uint32(v302)+120)) = v255
	*(*int64)(unsafe.Add(mBase, uint32(v302)+104)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v302)+100)) = v256
	*(*int32)(unsafe.Add(mBase, uint32(v302)+96)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v302)+20)) = v258
	*(*int32)(unsafe.Add(mBase, uint32(v302)+16)) = v259
	v309 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v302))), uint32(v309))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L7
	} else {
		goto L99
	}
L99:
	;
	if (l1^int32(-1)|v186)&int32(1) == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L7
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+96)) = v18 & int64(4294967295)
	v331 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+94)) = uint8(v331)
	v335 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[4]))
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v335)+120))
	if v336 == int64(0) {
		goto L105
	} else {
		goto L106
	}
L103:
	;
	goto L102
L104:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+95)) = uint8(v341)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v16)+88))
	v348 = F_heap_form_tuple(m, v343, v16+int32(96), v16+int32(94))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L108
	}
L105:
	;
	v341 = int32(1)
	goto L104
L106:
	;
	goto L107
L107:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+104)) = v336
	v341 = v331
	goto L104
L108:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v348)+16))
	v351 = F_HeapTupleHeaderGetDatum(m, v350)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	m.G0 = v16 + int32(704)
	return v351
L111:
	;
	F_errmsg_internal(m, int32(_a_F_copy_replication_slot_5), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(656), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L7
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
	if v122 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v382 = int32(_a_F_copy_replication_slot_6)
	goto L118
L117:
	;
	v382 = int32(_a_F_copy_replication_slot_7)
	goto L118
L118:
	;
	F_errmsg(m, v382, v16)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(713), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L7
	} else {
		goto L120
	}
L120:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L121:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_copy_replication_slot_8), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(719), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L7
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_9), v16-int32(-64))
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(726), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L7
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_10), v16+int32(16))
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	v441 = F_errdetail(m, int32(_a_F_copy_replication_slot_11), int32(0))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(819), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L7
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L133:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_12), v16+int32(32))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	F_errhint(m, int32(_a_F_copy_replication_slot_13), int32(0))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(827), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v469 = m.ExcPending
	if v469 != 0 {
		goto L7
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_14), v16+int32(48))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	v482 = F_errdetail(m, int32(_a_F_copy_replication_slot_15), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_3), int32(841), int32(_a_F_copy_replication_slot_4))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L7
	} else {
		goto L141
	}
L141:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replication_yyensure_buffer_stack(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == int32(0) {
		v8 = F_palloc(m, int32(4))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v8
			if v8 == int32(0) {
				F_yy_fatal_error_3(m, int32(_a_F_replication_yyensure_buffer_stack_0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(0)
				*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(4294967296)
				return
			}
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if base.Ui32(v18-int32(1)) <= base.Ui32(v17) {
			v23 = v18 + int32(8)
			v26 = F_repalloc(m, v4, v23<<(uint(int32(2))%32))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v26
				if v26 == int32(0) {
					F_yy_fatal_error_3(m, int32(_a_F_replication_yyensure_buffer_stack_0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					v34 = v26 + v31<<(uint(int32(2))%32)
					v35 = int64(0)
					*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v35
					*(*int64)(unsafe.Add(mBase, uint32(v34))) = v35
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v23
					return
				}
			}
		} else {
			return
		}
	}
}
func F_replication_yyerror(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_errcode(m, int32(16801924))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
			F_errmsg_internal(m, int32(_a_F_replication_yyerror_0), v5)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_replication_yyerror_1), int32(269), int32(_a_F_replication_yyerror_2))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
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
