package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

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
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(256)
	m.G0 = v11
	v18 = v4
	v19 = v4
	v20 = v4
	v21 = int32(-1)
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
	if v21 != int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v212 = int32(m.ExcTag)
	v213 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v212 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v19
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
	v60 = v18
	v61 = v19
	v62 = v20
	goto L9
L9:
	;
	if v62 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v18
	v33 = v11 + int32(232)
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
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v18
	v38 = F_quote_identifier(m, l1)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v38
	F_appendStringInfo(m, v33, int32(_a_F_ReplicationSlotDropAtPubNode_1), v11+int32(48))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0]))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1]))
	goto L14
L14:
	;
	v54 = v11 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v54))) = v11 + int32(60)
	goto L17
L15:
	;
	v60 = v50
	v61 = v52
	v62 = int32(0)
	goto L9
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v11 - int32(-64)
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[2]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+232))
	v76 = int32(0)
	v78 = m.T0[v72].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v75, v76, v76)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L24
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0])) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v11)+232))
	F_pfree(m, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L6
	} else {
		goto L55
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L51
	}
L22:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	if v132 != 0 {
		goto L37
	} else {
		goto L38
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_errfinish(m, int32(_a_F_ReplicationSlotDropAtPubNode_2), v124, int32(_a_F_ReplicationSlotDropAtPubNode_3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L36
	}
L24:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v80 == int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	v87 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if base.B2i32(l2 == int32(0))|v80 != 0 {
		goto L21
	} else {
		goto L31
	}
L28:
	;
	if v87 == int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F_errmsg(m, int32(_a_F_ReplicationSlotDropAtPubNode_4), v11)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	v124 = int32(1937)
	goto L23
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v101 != int32(67137668) {
		goto L21
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	v108 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	if v108 == int32(0) {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	F_errmsg(m, int32(_a_F_ReplicationSlotDropAtPubNode_5), v11+int32(16))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	v124 = int32(1946)
	goto L23
L36:
	;
	goto L22
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_pfree(m, v132)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v137 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_tuplestore_end(m, v137)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v142 != 0 {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_FreeTupleDesc(m, v142)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L6
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_pfree(m, v78)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L6
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0])) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v11)+232))
	F_pfree(m, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L6
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[0])) = v60
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotDropAtPubNode[1])) = v61
	m.G0 = v11 + int32(256)
	return
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_errcode(m, int32(100663808))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L52
	}
L52:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v178
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l1
	F_errmsg(m, int32(_a_F_ReplicationSlotDropAtPubNode_5), v11+int32(32))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_errfinish(m, int32(_a_F_ReplicationSlotDropAtPubNode_2), int32(1954), int32(_a_F_ReplicationSlotDropAtPubNode_3))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L3
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+252)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v11)+248)) = v60
	F_pg_re_throw(m)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L6
	} else {
		goto L56
	}
L56:
	;
	goto L5
L57:
	;
	v217 = int32(v213)
	m.G0 = v11
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v217)+4))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v217)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	if v11+int32(60) == v223 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	m.ExcPending = 1
	goto L66
L59:
	;
	if v227 != 0 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v227 = v225
	goto L62
L61:
	;
	v227 = int32(0)
	goto L62
L62:
	;
	goto L59
L63:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v11)+252))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v11)+248))
	v18 = v229
	v19 = v228
	v20 = v219
	v21 = v227
	goto L1
L64:
	;
	goto L65
L65:
	;
	F___wasm_longjmp(m, v220, v219)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	return
L67:
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
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
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v76 int64
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0]))
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
	v27 = int32(_a_F_ReplicationSlotRelease_1)
	v28 = int32(0)
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	if v29 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+88))
	if v25 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = int32(_a_F_ReplicationSlotRelease_0)
	goto L8
L7:
	;
	v26 = int32(_a_F_ReplicationSlotRelease_1)
	goto L8
L8:
	;
	v27 = v26
	v28 = v21
	goto L3
L9:
	;
	v32 = int32(_a_F_ReplicationSlotRelease_2)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0])) = int32(0)
	F_ReplicationSlotDropPtr(m, v33)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	if v39 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v62 = m.G0
	v63 = int32(16)
	v64 = v62 - v63
	m.G0 = v64
	F_gettimeofday(m, v64)
	mBase = m.M
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v64)))
	v68 = int64(*(*int32)(unsafe.Add(mBase, uint32(v64)+8)))
	m.G0 = v64 + v63
	v76 = v68 + v67*int64(1000000) - int64(946684800000000)
	goto L21
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v40 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v45 = base.AtomicRmwXchg32(m, v13, int32(0), int32(1))
	if v45 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v13, int32(_a_F_ReplicationSlotRelease_3), int32(750), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v51 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v51
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v51))
	F_ReplicationSlotsComputeRequiredXmin(m, v51)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	v79 = int32(0)
	v80 = base.AtomicRmwXchg32(m, v13, v79, int32(1))
	if v77 == v79 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v114 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0])) = v114
	v117 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[2]))
	v121 = F_LWLockAcquire(m, v117+int32(512), v114)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L41
	}
L23:
	;
	if v80 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v80 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	F_s_lock(m, v13, int32(_a_F_ReplicationSlotRelease_3), int32(768), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
	if v90 == v88 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+272)) = v76
	goto L32
L31:
	;
	goto L32
L32:
	;
	v94 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v94))
	F_ConditionVariableBroadcast(m, v13+int32(224))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L22
L34:
	;
	F_s_lock(m, v13, int32(_a_F_ReplicationSlotRelease_5), int32(251), int32(_a_F_ReplicationSlotRelease_6))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v13)+112))
	if v106 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+272)) = v76
	goto L40
L39:
	;
	goto L40
L40:
	;
	v110 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v13))), uint32(v110))
	goto L22
L41:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[3]))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+124)))
	v127 = v125 & int32(-17)
	*(*uint8)(unsafe.Add(mBase, uint32(v124)+124)) = uint8(v127)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[4]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v124)+48))
	*(*uint8)(unsafe.Add(mBase, uint32(v131+v132))) = uint8(v127)
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[2]))
	F_LWLockRelease(m, v136+int32(512))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[1])))
	if v142 == int32(1) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[5])))
	if v148 != 0 {
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
	v149 = int32(15)
	goto L48
L47:
	;
	v149 = int32(14)
	goto L48
L48:
	;
	v151 = F_errstart(m, v149, int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	if v151 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v28
	F_errmsg(m, v27, v10)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_pfree(m, v28)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotRelease_3), int32(792), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	goto L45
}
func F_copy_replication_slot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int64
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v197 int32
	_ = v197
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v238 int64
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int64
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v485 int32
	_ = v485
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(688)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = F_get_call_result_type(m, l0, v3, v16+int32(96))
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
	v470 = m.ExcPending
	if v470 != 0 {
		goto L7
	} else {
		goto L138
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L7
	} else {
		goto L133
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L7
	} else {
		goto L129
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L7
	} else {
		goto L125
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L7
	} else {
		goto L121
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L7
	} else {
		goto L114
	}
L7:
	;
	return int32(0)
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
	v359 = m.ExcPending
	if v359 != 0 {
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
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	v40 = F_LWLockAcquire(m, v36+int32(_a_F_copy_replication_slot_0), int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L19
	}
L14:
	;
	F_CheckLogicalDecodingRequirements(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_CheckSlotRequirements(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
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
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[1]))
	if int32(0) < v43 {
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
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v176 = base.B2i32(v174 != int32(0))
	if v128 == int32(3) {
		v180 = v176
		v181 = v127
		goto L20
	} else {
		goto L55
	}
L22:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[2]))
	v52 = v3
	goto L25
L23:
	;
	goto L24
L24:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	F_LWLockRelease(m, v151+int32(_a_F_copy_replication_slot_0))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L50
	}
L25:
	;
	v63 = v47 + v52*int32(288)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+4)))
	if v64 != int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L24
L27:
	;
	v135 = v52 + int32(1)
	if v135 != v43 {
		v52 = v135
		goto L25
	} else {
		goto L49
	}
L28:
	;
	v68 = v63 + int32(24)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v71 == int32(0))|base.B2i32(v71 != v74) != 0 {
		v92 = v71
		v93 = v74
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v92-v93 != 0 {
		goto L27
	} else {
		goto L36
	}
L30:
	;
	goto L29
L31:
	;
	v77 = v68
	v78 = v19
	goto L32
L32:
	;
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78)+1)))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	if v82 == int32(0) {
		v92 = v82
		v93 = v81
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v92 = v82
	v93 = v81
	goto L30
L34:
	;
	v85 = int32(1)
	if v82 == v81 {
		v77 = v77 + v85
		v78 = v78 + v85
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v97 = base.AtomicRmwXchg32(m, v63, int32(0), int32(1))
	if v97 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_s_lock(m, v63, int32(_a_F_copy_replication_slot_1), int32(651), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	base.MemoryCopy(m, v16+int32(400), v63, int32(288))
	v107 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v63))), uint32(v107))
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	F_LWLockRelease(m, v111+int32(_a_F_copy_replication_slot_0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L7
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v16)+488))
	if l1 != base.B2i32(v116 != int32(0)) {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v120 = *(*int64)(unsafe.Add(mBase, uint32(v16)+504))
	if v120 == int64(0) {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v16)+512))
	if v123 != 0 {
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
	v127 = v16 + int32(537)
	goto L47
L46:
	;
	v127 = int32(0)
	goto L47
L47:
	;
	v128 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(3) <= v128 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v16)+492))
	v180 = base.B2i32(v131 == int32(2))
	v181 = v127
	goto L20
L49:
	;
	goto L26
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_3), v16+int32(80))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(664), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
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
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v180 = v176
	v181 = v179
	goto L20
L56:
	;
	v225 = base.AtomicRmwXchg32(m, v63, int32(0), int32(1))
	if v225 != 0 {
		goto L72
	} else {
		goto L73
	}
L57:
	;
	v182 = int32(1)
	if v180 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v207 = int32(0)
	if v180 != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v185 = int32(2)
	goto L62
L61:
	;
	v185 = v182
	goto L62
L62:
	;
	v186 = int32(0)
	F_ReplicationSlotCreate(m, v18, v182, v185, v186, v186, v186)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = int32(396)
	v197 = int32(0)
	v203 = F_CreateInitDecodingContext(m, v181, v197, v120, v16+int32(112), v197, v197, v197)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_FreeDecodingContext(m, v203)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	goto L56
L66:
	;
	v210 = int32(2)
	goto L68
L67:
	;
	v210 = v207
	goto L68
L68:
	;
	v211 = int32(0)
	F_ReplicationSlotCreate(m, v18, v207, v210, v211, v211, v211)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v217)+104)) = v120
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L7
	} else {
		goto L71
	}
L71:
	;
	goto L56
L72:
	;
	F_s_lock(m, v63, int32(_a_F_copy_replication_slot_1), int32(753), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	base.MemoryCopy(m, v16+int32(112), v63, int32(288))
	v235 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v63))), uint32(v235))
	v238 = *(*int64)(unsafe.Add(mBase, uint32(v16)+216))
	if base.Ui64(v238) < base.Ui64(v120) {
		goto L3
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	v241 = int32(0)
	if base.B2i32(v240 == v241) == base.B2i32(v116 != v241) {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v246 = *(*int64)(unsafe.Add(mBase, uint32(v16)+232))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v16)+208))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v16)+132))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v16)+128))
	v252 = v16 + int32(136)
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v255 == int32(0))|base.B2i32(v255 != v258) != 0 {
		v276 = v255
		v277 = v258
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v276-v277 != 0 {
		goto L3
	} else {
		goto L85
	}
L79:
	;
	goto L78
L80:
	;
	v261 = v252
	v262 = v19
	goto L81
L81:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+1)))
	v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v261)+1)))
	if v266 == int32(0) {
		v276 = v266
		v277 = v265
		goto L79
	} else {
		goto L83
	}
L82:
	;
	v276 = v266
	v277 = v265
	goto L79
L83:
	;
	v269 = int32(1)
	if v266 == v265 {
		v261 = v261 + v269
		v262 = v262 + v269
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	if v246 == int64(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v282 = v116
	goto L88
L87:
	;
	v282 = int32(0)
	goto L88
L88:
	;
	if v282 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v16)+224))
	if v283 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	v288 = base.AtomicRmwXchg32(m, v285, int32(0), int32(1))
	if v288 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v290 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	F_s_lock(m, v290, int32(_a_F_copy_replication_slot_1), int32(810), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v297 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v297)+120)) = v246
	*(*int64)(unsafe.Add(mBase, uint32(v297)+104)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v297)+100)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v297)+96)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v297)+20)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v297)+16)) = v250
	v304 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v297))), uint32(v304))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L7
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	F_ReplicationSlotsComputeRequiredXmin(m, int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	if (l1^int32(-1)|v180)&int32(1) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L7
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+104)) = v18
	v326 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+102)) = uint8(v326)
	v330 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v330)+120))
	if v331 == int64(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+103)) = uint8(v338)
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v345 = F_heap_form_tuple(m, v340, v16+int32(104), v16+int32(102))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L108
	}
L104:
	;
	v338 = int32(1)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v335 = F_Int64GetDatum(m, v331)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = v335
	v338 = v326
	goto L103
L108:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v345)+16))
	v348 = F_HeapTupleHeaderGetDatum(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	m.G0 = v16 + int32(688)
	return v348
L111:
	;
	F_errmsg_internal(m, int32(_a_F_copy_replication_slot_4), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(622), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
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
	v375 = m.ExcPending
	if v375 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
	if v116 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v379 = int32(_a_F_copy_replication_slot_5)
	goto L118
L117:
	;
	v379 = int32(_a_F_copy_replication_slot_6)
	goto L118
L118:
	;
	F_errmsg(m, v379, v16)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(679), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
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
	v393 = m.ExcPending
	if v393 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_copy_replication_slot_7), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(685), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
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
	v409 = m.ExcPending
	if v409 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_8), v16-int32(-64))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(692), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
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
	F_errmsg(m, int32(_a_F_copy_replication_slot_9), v16+int32(16))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	F_errdetail(m, int32(_a_F_copy_replication_slot_10), int32(0))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(785), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
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
	v451 = m.ExcPending
	if v451 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_11), v16+int32(32))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	F_errhint(m, int32(_a_F_copy_replication_slot_12), int32(0))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(793), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
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
	F_errmsg(m, int32(_a_F_copy_replication_slot_13), v16+int32(48))
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	F_errdetail(m, int32(_a_F_copy_replication_slot_14), int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(807), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
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
				F_errfinish(m, int32(_a_F_replication_yyerror_1), int32(265), int32(_a_F_replication_yyerror_2))
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
