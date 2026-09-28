package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOXact_LogicalCtl(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalCtl[0])))
	if v3 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalCtl[1]))
		v9 = F_LWLockAcquire(m, v5+int32(_a_F_AtEOXact_LogicalCtl_0), int32(1))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalCtl[2]))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			v15 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalCtl[1]))
			F_LWLockRelease(m, v15+int32(_a_F_AtEOXact_LogicalCtl_0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalCtl[3])) = uint8(v13)
				v23 = int32(0)
				*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_LogicalCtl[0])) = uint8(v23)
				return
			}
		}
	} else {
		return
	}
}
func F_AtEOXact_RelationCache(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[0])) = v2
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[1])))
	if v14 == v2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[2]))
	if int32(0) < v141 {
		goto L45
	} else {
		goto L46
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[3]))
	if v18 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v82 = v8 + int32(28)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[4]))
	F_hash_seq_init(m, v82, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L12
	} else {
		goto L25
	}
L5:
	;
	if l0 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v23 = int32(44)
	goto L8
L7:
	;
	v23 = int32(32)
	goto L8
L8:
	;
	v25 = v2
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[4]))
	v35 = int32(0)
	v37 = F_hash_search(m, v30, v25<<(uint(int32(2))%32)+int32(_a_F_AtEOXact_RelationCache_0), v35, v35)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L1
L11:
	;
	v77 = v25 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[3]))
	if v77 < v79 {
		v25 = v77
		goto L9
	} else {
		goto L24
	}
L12:
	;
	return
L13:
	;
	if v37 == int32(0) {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v41+v23)))
	v44 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v41)+40)) = v44
	*(*int64)(unsafe.Add(mBase, uint32(v41)+32)) = v44
	if v43 == int32(0) {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	if v50 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	F_RelationClearRelation(m, v41)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L12
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v57 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L12
	} else {
		goto L20
	}
L19:
	;
	goto L11
L20:
	;
	if v57 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v61 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AtEOXact_RelationCache_1), v8+int32(16))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_RelationCache_2), int32(3359), int32(_a_F_AtEOXact_RelationCache_3))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L11
L24:
	;
	goto L10
L25:
	;
	v87 = F_hash_seq_search(m, v82)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L12
	} else {
		goto L26
	}
L26:
	;
	if v87 == int32(0) {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if l0 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v93 = int32(44)
	goto L30
L29:
	;
	v93 = int32(32)
	goto L30
L30:
	;
	v95 = v87
	goto L31
L31:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99+v93)))
	v102 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v99)+40)) = v102
	*(*int64)(unsafe.Add(mBase, uint32(v99)+32)) = v102
	if v101 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L1
L33:
	;
	v133 = F_hash_seq_search(m, v8+int32(28))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L12
	} else {
		goto L43
	}
L34:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v99)+16))
	if v108 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	F_RelationClearRelation(m, v99)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L12
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v115 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L12
	} else {
		goto L39
	}
L38:
	;
	goto L33
L39:
	;
	if v115 == int32(0) {
		goto L33
	} else {
		goto L40
	}
L40:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v99)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v119 + int32(4)
	F_errmsg_internal(m, int32(_a_F_AtEOXact_RelationCache_1), v8)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L12
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_RelationCache_2), int32(3359), int32(_a_F_AtEOXact_RelationCache_3))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L12
	} else {
		goto L42
	}
L42:
	;
	goto L33
L43:
	;
	if v133 != 0 {
		v95 = v133
		goto L31
	} else {
		goto L44
	}
L44:
	;
	goto L32
L45:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[5]))
	if int32(0) < v145 {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L47
L47:
	;
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[1])) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[3])) = v185
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[5])) = v185
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[2])) = v185
	m.G0 = v8 + int32(48)
	return
L48:
	;
	v150 = int32(0)
	goto L51
L49:
	;
	goto L50
L50:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[6]))
	F_pfree(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L12
	} else {
		goto L55
	}
L51:
	;
	v155 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[6]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v150<<(uint(int32(2))%32))))
	F_FreeTupleDesc(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L53
	}
L52:
	;
	goto L50
L53:
	;
	v163 = v150 + int32(1)
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[5]))
	if v163 < v165 {
		v150 = v163
		goto L51
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_RelationCache[6])) = int32(0)
	goto L47
}
func F_AtEOXact_Snapshot(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[0]))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_pairingheap_remove(m, int32(_a_F_AtEOXact_Snapshot_0), v12+int32(52))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[0])) = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[1]))
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if int32(0) < v23 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[2]))
	if v86 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v29 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[1])) = int32(0)
	goto L8
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+v29<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v39 = F_unlink(m, v38)
	mBase = m.M
	if v39 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	F_pairingheap_remove(m, int32(_a_F_AtEOXact_Snapshot_0), v61+int32(52))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L4
	} else {
		goto L20
	}
L15:
	;
	v44 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v44 == int32(0) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v48
	F_errmsg_internal(m, int32(_a_F_AtEOXact_Snapshot_1), v9+int32(16))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_Snapshot_2), int32(1059), int32(_a_F_AtEOXact_Snapshot_3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	goto L14
L20:
	;
	v67 = v29 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v67 < v68 {
		v29 = v67
		goto L12
	} else {
		goto L21
	}
L21:
	;
	goto L13
L22:
	;
	if l0 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L23:
	;
	F_pairingheap_remove(m, int32(_a_F_AtEOXact_Snapshot_0), v86+int32(52))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[2])) = int32(0)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[3]))
	if v98 != 0 {
		goto L22
	} else {
		goto L25
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[4]))
	if v100 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v104 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[5])) = v104
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+52)) = v104
	goto L22
L27:
	;
	goto L28
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[6]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+52))
	v113 = int32(3)
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[4]))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v116-int32(48))))
	if base.B2i32(base.Ui32(v112) < base.Ui32(v113))|base.B2i32(base.Ui32(v119) < base.Ui32(v113)) == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[5])) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v111)+52)) = v119
	goto L22
L30:
	;
	if v112-v119 < int32(0) {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(v119) <= base.Ui32(v112) {
		goto L22
	} else {
		goto L34
	}
L33:
	;
	goto L22
L34:
	;
	goto L29
L35:
	;
	v187 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[4])) = v187
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[3])) = v187
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[7])) = v187
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[8])) = v187
	*(*uint8)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[9])) = uint8(v187)
	if l1 != 0 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[4]))
	if v138 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[3]))
	if v157 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L38:
	;
	v143 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	if v143 == int32(0) {
		goto L37
	} else {
		goto L40
	}
L40:
	;
	F_errmsg_internal(m, int32(_a_F_AtEOXact_Snapshot_4), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_Snapshot_2), int32(1077), int32(_a_F_AtEOXact_Snapshot_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L37
L43:
	;
	v162 = v157
	goto L44
L44:
	;
	v168 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L46
	}
L45:
	;
	goto L35
L46:
	;
	if v168 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v162
	F_errmsg_internal(m, int32(_a_F_AtEOXact_Snapshot_5), v9)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	if v179 != 0 {
		v162 = v179
		goto L44
	} else {
		goto L52
	}
L50:
	;
	F_errfinish(m, int32(_a_F_AtEOXact_Snapshot_2), int32(1081), int32(_a_F_AtEOXact_Snapshot_3))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	goto L45
L53:
	;
	v202 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[5])) = v202
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_Snapshot[6]))
	*(*int32)(unsafe.Add(mBase, uint32(v205)+52)) = v202
	goto L55
L54:
	;
	goto L55
L55:
	;
	m.G0 = v9 + int32(32)
	return
}
func F_AtEOXact_TypeCache(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v1 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_TypeCache[0]))
	if v1 < v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v1
	goto L4
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_TypeCache[0])) = int32(0)
	m.G0 = v7 + int32(16)
	return
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_TypeCache[1]))
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_TypeCache[2]))
	v24 = int32(0)
	v26 = F_hash_search(m, v18, v20+v15<<(uint(int32(2))%32), v24, v24)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L3
L6:
	;
	v56 = v15 + int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_TypeCache[0]))
	if v56 < v58 {
		v15 = v56
		goto L4
	} else {
		goto L16
	}
L7:
	;
	return
L8:
	;
	if v26 == int32(0) {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+13)))
	if v30 != int32(99) {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+312))
	if v33&int32(-1572865) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+188))
	if v38 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOXact_TypeCache[3]))
	v48 = F_hash_search(m, v42, v26+int32(16), int32(1), v7+int32(15))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L13
L15:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v52
	goto L6
L16:
	;
	goto L5
}
