package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
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
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0]))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[1])))
	if v15 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v20 = F_pstrdup(m, v12+int32(24))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = int32(_a_F_ReplicationSlotRelease_1)
	v27 = int32(0)
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	if v28 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+88))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v25 = int32(_a_F_ReplicationSlotRelease_0)
	goto L8
L7:
	;
	v25 = int32(_a_F_ReplicationSlotRelease_1)
	goto L8
L8:
	;
	v26 = v25
	v27 = v20
	goto L3
L9:
	;
	v31 = int32(_a_F_ReplicationSlotRelease_2)
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0])) = int32(0)
	F_ReplicationSlotDropPtr(m, v32)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v12)+96))
	if v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L11
L13:
	;
	v60 = m.G0
	v61 = int32(16)
	v62 = v60 - v61
	m.G0 = v62
	F_gettimeofday(m, v62)
	mBase = m.M
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v62)))
	v66 = int64(*(*int32)(unsafe.Add(mBase, uint32(v62)+8)))
	m.G0 = v62 + v61
	v74 = v66 + v65*int64(1000000) - int64(946684800000000)
	goto L21
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v39 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_s_lock(m, v12, int32(_a_F_ReplicationSlotRelease_3), int32(750), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v50 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v50
	F_ReplicationSlotsComputeRequiredXmin(m, v50)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
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
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v12)+92))
	if v78 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v110 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[0])) = v110
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[2]))
	v117 = F_LWLockAcquire(m, v113+int32(512), v110)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L4
	} else {
		goto L41
	}
L23:
	;
	if v75 != 0 {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	if v75 != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	F_s_lock(m, v12, int32(_a_F_ReplicationSlotRelease_3), int32(768), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v86 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
	if v88 == v86 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L28
L30:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+272)) = v74
	goto L32
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	F_ConditionVariableBroadcast(m, v12+int32(224))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	goto L22
L34:
	;
	F_s_lock(m, v12, int32(_a_F_ReplicationSlotRelease_5), int32(251), int32(_a_F_ReplicationSlotRelease_6))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+112))
	if v103 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L36
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v12)+272)) = v74
	goto L40
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = int32(0)
	goto L22
L41:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[3]))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+124)))
	v123 = v121 & int32(-17)
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+124)) = uint8(v123)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_ReplicationSlotRelease[4]))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v120)+48))
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
	m.G0 = v9 + int32(16)
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
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v27
	F_errmsg(m, v26, v9)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_pfree(m, v27)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L55
	}
L53:
	;
	F_errfinish(m, int32(_a_F_ReplicationSlotRelease_3), int32(792), int32(_a_F_ReplicationSlotRelease_4))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
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
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int64
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v236 int64
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int64
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v426 int32
	_ = v426
	var v432 int32
	_ = v432
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
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
	v467 = m.ExcPending
	if v467 != 0 {
		goto L7
	} else {
		goto L138
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L7
	} else {
		goto L133
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L7
	} else {
		goto L129
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L7
	} else {
		goto L125
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L7
	} else {
		goto L121
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
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
	v356 = m.ExcPending
	if v356 != 0 {
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
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v175 = base.B2i32(v173 != int32(0))
	if v127 == int32(3) {
		v179 = v175
		v180 = v126
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
	v150 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	F_LWLockRelease(m, v150+int32(_a_F_copy_replication_slot_0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
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
	v134 = v52 + int32(1)
	if v134 != v43 {
		v52 = v134
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
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(1)
	if v95 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(0)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[0]))
	F_LWLockRelease(m, v110+int32(_a_F_copy_replication_slot_0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L7
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v16)+488))
	if l1 != base.B2i32(v115 != int32(0)) {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	v119 = *(*int64)(unsafe.Add(mBase, uint32(v16)+504))
	if v119 == int64(0) {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v16)+512))
	if v122 != 0 {
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
	v126 = v16 + int32(537)
	goto L47
L46:
	;
	v126 = int32(0)
	goto L47
L47:
	;
	v127 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	if int32(3) <= v127 {
		goto L21
	} else {
		goto L48
	}
L48:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v16)+492))
	v179 = base.B2i32(v130 == int32(2))
	v180 = v126
	goto L20
L49:
	;
	goto L26
L50:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L7
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_3), v16+int32(80))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(664), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
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
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v179 = v175
	v180 = v178
	goto L20
L56:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(1)
	if v222 != 0 {
		goto L72
	} else {
		goto L73
	}
L57:
	;
	v181 = int32(1)
	if v179 != 0 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v206 = int32(0)
	if v179 != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v184 = int32(2)
	goto L62
L61:
	;
	v184 = v181
	goto L62
L62:
	;
	v185 = int32(0)
	F_ReplicationSlotCreate(m, v18, v181, v184, v185, v185, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L7
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+120)) = int32(394)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+116)) = int32(395)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+112)) = int32(396)
	v196 = int32(0)
	v202 = F_CreateInitDecodingContext(m, v180, v196, v119, v16+int32(112), v196, v196, v196)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L7
	} else {
		goto L64
	}
L64:
	;
	F_FreeDecodingContext(m, v202)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	goto L56
L66:
	;
	v209 = int32(2)
	goto L68
L67:
	;
	v209 = v206
	goto L68
L68:
	;
	v210 = int32(0)
	F_ReplicationSlotCreate(m, v18, v206, v209, v210, v210, v210)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L69
	}
L69:
	;
	v216 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v216)+104)) = v119
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L70
	}
L70:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
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
	v229 = m.ExcPending
	if v229 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(0)
	v236 = *(*int64)(unsafe.Add(mBase, uint32(v16)+216))
	if base.Ui64(v236) < base.Ui64(v119) {
		goto L3
	} else {
		goto L76
	}
L75:
	;
	goto L74
L76:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v16)+200))
	v239 = int32(0)
	if base.B2i32(v238 == v239) == base.B2i32(v115 != v239) {
		goto L3
	} else {
		goto L77
	}
L77:
	;
	v244 = *(*int64)(unsafe.Add(mBase, uint32(v16)+232))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v16)+212))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v16)+208))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v16)+132))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v16)+128))
	v250 = v16 + int32(136)
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250))))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if base.B2i32(v253 == int32(0))|base.B2i32(v253 != v256) != 0 {
		v274 = v253
		v275 = v256
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v274-v275 != 0 {
		goto L3
	} else {
		goto L85
	}
L79:
	;
	goto L78
L80:
	;
	v259 = v250
	v260 = v19
	goto L81
L81:
	;
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260)+1)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+1)))
	if v264 == int32(0) {
		v274 = v264
		v275 = v263
		goto L79
	} else {
		goto L83
	}
L82:
	;
	v274 = v264
	v275 = v263
	goto L79
L83:
	;
	v267 = int32(1)
	if v264 == v263 {
		v259 = v259 + v267
		v260 = v260 + v267
		goto L81
	} else {
		goto L84
	}
L84:
	;
	goto L82
L85:
	;
	if v244 == int64(0) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v280 = v115
	goto L88
L87:
	;
	v280 = int32(0)
	goto L88
L88:
	;
	if v280 != 0 {
		goto L2
	} else {
		goto L89
	}
L89:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v16)+224))
	if v281 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v283 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v283)))
	*(*int32)(unsafe.Add(mBase, uint32(v283))) = int32(1)
	if v284 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	F_s_lock(m, v288, int32(_a_F_copy_replication_slot_1), int32(810), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L7
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v295)+120)) = v244
	*(*int64)(unsafe.Add(mBase, uint32(v295)+104)) = v236
	*(*int32)(unsafe.Add(mBase, uint32(v295)+100)) = v245
	*(*int32)(unsafe.Add(mBase, uint32(v295)+96)) = v246
	*(*int32)(unsafe.Add(mBase, uint32(v295)+20)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v295)+16)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v295))) = int32(0)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
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
	v308 = m.ExcPending
	if v308 != 0 {
		goto L7
	} else {
		goto L96
	}
L96:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L97
	}
L97:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L7
	} else {
		goto L98
	}
L98:
	;
	if (l1^int32(-1)|v179)&int32(1) == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_ReplicationSlotPersist(m)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
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
	v323 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+102)) = uint8(v323)
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_copy_replication_slot[3]))
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v327)+120))
	if v328 == int64(0) {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L101
L103:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+103)) = uint8(v335)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
	v342 = F_heap_form_tuple(m, v337, v16+int32(104), v16+int32(102))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L7
	} else {
		goto L108
	}
L104:
	;
	v335 = int32(1)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v332 = F_Int64GetDatum(m, v328)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L7
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+108)) = v332
	v335 = v323
	goto L103
L108:
	;
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v342)+16))
	v345 = F_HeapTupleHeaderGetDatum(m, v344)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L7
	} else {
		goto L109
	}
L109:
	;
	F_ReplicationSlotRelease(m)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L7
	} else {
		goto L110
	}
L110:
	;
	m.G0 = v16 + int32(688)
	return v345
L111:
	;
	F_errmsg_internal(m, int32(_a_F_copy_replication_slot_4), int32(0))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L7
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(622), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
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
	v372 = m.ExcPending
	if v372 != 0 {
		goto L7
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v19
	if v115 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v376 = int32(_a_F_copy_replication_slot_5)
	goto L118
L117:
	;
	v376 = int32(_a_F_copy_replication_slot_6)
	goto L118
L118:
	;
	F_errmsg(m, v376, v16)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L7
	} else {
		goto L119
	}
L119:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(679), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
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
	v390 = m.ExcPending
	if v390 != 0 {
		goto L7
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(_a_F_copy_replication_slot_7), int32(0))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L7
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(685), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
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
	v406 = m.ExcPending
	if v406 != 0 {
		goto L7
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+64)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_8), v16-int32(-64))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L7
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(692), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
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
	v432 = m.ExcPending
	if v432 != 0 {
		goto L7
	} else {
		goto L130
	}
L130:
	;
	F_errdetail(m, int32(_a_F_copy_replication_slot_10), int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L7
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(785), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
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
	v448 = m.ExcPending
	if v448 != 0 {
		goto L7
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v19
	F_errmsg(m, int32(_a_F_copy_replication_slot_11), v16+int32(32))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L7
	} else {
		goto L135
	}
L135:
	;
	F_errhint(m, int32(_a_F_copy_replication_slot_12), int32(0))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L7
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(793), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
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
	v473 = m.ExcPending
	if v473 != 0 {
		goto L7
	} else {
		goto L139
	}
L139:
	;
	F_errdetail(m, int32(_a_F_copy_replication_slot_14), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L7
	} else {
		goto L140
	}
L140:
	;
	F_errfinish(m, int32(_a_F_copy_replication_slot_1), int32(807), int32(_a_F_copy_replication_slot_2))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
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
