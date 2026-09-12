package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChooseRelationName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
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
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	v15 = m.G0
	v17 = v15 - int32(256)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = int32(4)
	v23 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = v17 + int32(192)
	goto L6
L3:
	;
	v159 = int32(0)
	goto L35
L4:
	;
	v141 = F_strlen(m, v130)
	mBase = m.M
	goto L3
L6:
	;
	goto L7
L7:
	;
	v35 = int32(63)
	if (v28^l2)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v134 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v134)
	goto L4
L9:
	;
	v115 = v110
	v116 = v111
	v117 = v112
	goto L31
L10:
	;
	if v105 == int32(0) {
		v130 = v103
		v131 = v104
		goto L8
	} else {
		goto L30
	}
L11:
	;
	v103 = l2
	v104 = v28
	v105 = v35
	goto L10
L12:
	;
	goto L13
L13:
	;
	if l2&int32(3) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v72 == int32(0) {
		v130 = v69
		v131 = v70
		goto L8
	} else {
		goto L23
	}
L15:
	;
	v69 = l2
	v70 = v28
	v71 = v35
	v72 = int32(1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v48 = l2
	v49 = v28
	v50 = v35
	goto L18
L18:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	*(*uint8)(unsafe.Add(mBase, uint32(v49))) = uint8(v52)
	if v52 == int32(0) {
		v110 = v48
		v111 = v49
		v112 = v50
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v69 = v63
	v70 = v57
	v71 = v59
	v72 = v61
	goto L14
L20:
	;
	v56 = int32(1)
	v57 = v49 + v56
	v59 = v50 - v56
	v60 = int32(0)
	v61 = base.B2i32(v59 != v60)
	v63 = v48 + v56
	if v63&int32(3) == v60 {
		v69 = v63
		v70 = v57
		v71 = v59
		v72 = v61
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v59 != 0 {
		v48 = v63
		v49 = v57
		v50 = v59
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v75 == int32(0) {
		v103 = v69
		v104 = v70
		v105 = v71
		goto L10
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v71) < base.Ui32(int32(4)) {
		v103 = v69
		v104 = v70
		v105 = v71
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v81 = v69
	v82 = v70
	v83 = v71
	goto L26
L26:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v89 = int32(-2139062144)
	if (int32(16843008)-v86|v86)&v89 != v89 {
		v110 = v81
		v111 = v82
		v112 = v83
		goto L9
	} else {
		goto L28
	}
L27:
	;
	v103 = v97
	v104 = v95
	v105 = v99
	goto L10
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82))) = v86
	v94 = int32(4)
	v95 = v82 + v94
	v97 = v81 + v94
	v99 = v83 - v94
	if base.Ui32(int32(3)) < base.Ui32(v99) {
		v81 = v97
		v82 = v95
		v83 = v99
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v110 = v103
	v111 = v104
	v112 = v105
	goto L9
L31:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	*(*uint8)(unsafe.Add(mBase, uint32(v116))) = uint8(v119)
	if v119 == int32(0) {
		v130 = v115
		v131 = v116
		goto L8
	} else {
		goto L33
	}
L32:
	;
	v130 = v126
	v131 = v124
	goto L8
L33:
	;
	v123 = int32(1)
	v124 = v116 + v123
	v126 = v115 + v123
	v128 = v117 - v123
	if v128 != 0 {
		v115 = v126
		v116 = v124
		v117 = v128
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v167 = F_makeObjectName(m, l0, l1, v17+int32(192))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	F_sequence_close(m, v23, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L58
	}
L37:
	;
	goto L36
L38:
	;
	F_ScanKeyInit(m, v17+int32(16), int32(2), int32(3), int32(62), v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v171 = int32(3)
	F_ScanKeyInit(m, v17-int32(-64), v171, v171, int32(184), l3)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v183 = F_systable_beginscan(m, v23, int32(2663), int32(1), v17+int32(120), int32(2), v17+int32(16))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v185 = F_systable_getnext(m, v183)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_systable_endscan(m, v183)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	if v185 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	if l4 == int32(0) {
		goto L37
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	F_pfree(m, v167)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L56
	}
L47:
	;
	v193 = m.G0
	v195 = v193 - int32(96)
	m.G0 = v195
	v199 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_ScanKeyInit(m, v195, int32(2), int32(3), int32(62), v167)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v208 = int32(3)
	F_ScanKeyInit(m, v195+int32(48), v208, v208, int32(184), l3)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v217 = F_systable_beginscan(m, v199, int32(2664), int32(1), int32(0), int32(2), v195)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v219 = F_systable_getnext(m, v217)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_systable_endscan(m, v217)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_sequence_close(m, v199, int32(1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v195 + int32(96)
	if v219 == int32(0) {
		goto L37
	} else {
		goto L55
	}
L55:
	;
	goto L46
L56:
	;
	v238 = v159 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l2
	v245 = F_pg_snprintf(m, v17+int32(192), int32(64), int32(486730), v17)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v159 = v238
	goto L35
L58:
	;
	m.G0 = v17 + int32(256)
	return v167
}
func F_DropRelationFiles(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int64
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int64
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v117 int32
	_ = v117
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v16 = F_palloc(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_pfree(m, v16)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L24
	}
L4:
	;
	v24 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_smgrdounlinkall(m, v16, l1, l2)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v31 = l0 + v24*int32(12)
	v33 = v31 + int32(8)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v34
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+64)) = v36
	v41 = F_smgropen(m, v12-int32(-64), int32(-1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_smgrdounlinkall(m, v16, l1, l2)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v43
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v45
	F_XLogDropRelation(m, v12+int32(48), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v24<<(uint(int32(2))%32)))) = v41
	v82 = v24 + int32(1)
	if v82 != l1 {
		v24 = v82
		goto L7
	} else {
		goto L17
	}
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v52
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v54
	F_XLogDropRelation(m, v12+int32(32), int32(1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v63
	F_XLogDropRelation(m, v12+int32(16), int32(2))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v70
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v72
	F_XLogDropRelation(m, v12, int32(3))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	goto L8
L18:
	;
	v91 = int32(0)
	goto L19
L19:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v16+v91<<(uint(int32(2))%32))))
	F_smgrclose(m, v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L3
L21:
	;
	v103 = v91 + int32(1)
	if v103 != l1 {
		v91 = v103
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L3
L24:
	;
	m.G0 = v12 + int32(80)
	return
}
func F_LockRelationId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v9
	v21 = F_LockAcquireExtended(m, v5+int32(16), int32(1), v2, v2, v5+int32(12), v2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if v21 != int32(3) {
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+53)) = uint8(v28)
				m.G0 = v5 + int32(32)
				return
			}
		} else {
			m.G0 = v5 + int32(32)
			return
		}
	}
}
func F_RelationCacheInitFilePreInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 - int32(2080)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(106923)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v8
		v18 = F_pg_snprintf(m, v5+int32(1056), int32(1024), int32(187320), v5+int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(106923)
			v26 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(187351), v5)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _consts[86]))
				v33 = F_LWLockAcquire(m, v29+int32(2048), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _consts[379]))
					if v36 != 0 {
						F_unlink_initfile(m, v5+int32(1056), int32(21))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_unlink_initfile(m, v5+int32(32), int32(21))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								m.G0 = v5 + int32(2080)
								return
							}
						}
					} else {
						F_unlink_initfile(m, v5+int32(32), int32(21))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v5 + int32(2080)
							return
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(106923)
		v26 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(187351), v5)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _consts[86]))
			v33 = F_LWLockAcquire(m, v29+int32(2048), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _consts[379]))
				if v36 != 0 {
					F_unlink_initfile(m, v5+int32(1056), int32(21))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_unlink_initfile(m, v5+int32(32), int32(21))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v5 + int32(2080)
							return
						}
					}
				} else {
					F_unlink_initfile(m, v5+int32(32), int32(21))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v5 + int32(2080)
						return
					}
				}
			}
		}
	}
}
func F_RelationClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	v8 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if v8 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _consts[182]))
		F_ResourceOwnerForget(m, v10, l0, int32(1771968))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v15 = v14
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v16 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v24 == int32(0) {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						if v27 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v24)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
					if v19 == int32(0) {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v24 == int32(0) {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							if v27 == int32(0) {
								return
							} else {
								F_MemoryContextDeleteChildren(m, v24)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_MemoryContextDeleteChildren(m, v16)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if v24 == int32(0) {
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
								if v27 == int32(0) {
									return
								} else {
									F_MemoryContextDeleteChildren(m, v24)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v15 = v5
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v16 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v24 == int32(0) {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					if v27 == int32(0) {
						return
					} else {
						F_MemoryContextDeleteChildren(m, v24)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
				if v19 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v24 == int32(0) {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						if v27 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v24)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F_MemoryContextDeleteChildren(m, v16)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v24 == int32(0) {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							if v27 == int32(0) {
								return
							} else {
								F_MemoryContextDeleteChildren(m, v24)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v293 int32
	_ = v293
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v455 int32
	_ = v455
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)))
	if v24 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v581
L2:
	;
	v573 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v574 = F_bms_copy(m, v573)
	mBase = m.M
	v575 = m.ExcPending
	if v575 != 0 {
		goto L11
	} else {
		goto L138
	}
L3:
	;
	switch l1 {
	case 0:
		goto L2
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+116)))
	if v55 != int32(1) {
		v581 = v3
		goto L1
	} else {
		goto L19
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L16
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v39 = F_bms_copy(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L15
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v36 = F_bms_copy(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L14
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v33 = F_bms_copy(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L13
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v28 = F_bms_copy(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v581 = v28
	goto L1
L13:
	;
	v581 = v33
	goto L1
L14:
	;
	v581 = v36
	goto L1
L15:
	;
	v581 = v39
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l1
	F_errmsg_internal(m, int32(58169), v22)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(522810), int32(5333), int32(249760))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v58 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v58 == int32(0) {
		v581 = v3
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v71 = v58
	goto L27
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L11
	} else {
		goto L135
	}
L23:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v581 = v557
	goto L1
L24:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v581 = v556
	goto L1
L25:
	;
	v581 = v461
	goto L1
L26:
	;
	v581 = v462
	goto L1
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v83
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v83 < v90 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v581 = int32(0)
	goto L1
L29:
	;
	v99 = v83
	v105 = v83
	v106 = v83
	v109 = int32(0)
	goto L32
L30:
	;
	v455 = v83
	v461 = v83
	v462 = v83
	goto L31
L31:
	;
	v469 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L11
	} else {
		goto L109
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v109<<(uint(int32(2))%32))))
	v119 = F_index_open(m, v117, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v455 = v423
	v461 = v429
	v462 = v430
	goto L31
L34:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+196))
	v123 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v123 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v127 = int32(4554240)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v131 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v131
	v134 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	v186 = v123
	goto L37
L37:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v202)+18)))
	if base.Ui32(v203&int32(2044)) <= base.Ui32(int32(19)) {
		goto L48
	} else {
		goto L49
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v134)+4)) = int64(-4294965047)
	v139 = v134 + int32(20)
	v142 = int32(0)
	goto L39
L39:
	;
	v159 = int32(100)
	v160 = v142 * v159
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	goto L42
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v134
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v128
	v186 = v134
	goto L37
L41:
	;
	F_populate_compact_attribute(m, v134, v142)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L11
	} else {
		goto L45
	}
L42:
	;
	v169 = F__emscripten_memcpy_bulkmem(m, v160+(v139+v161<<(uint(int32(4))%32)), v160+int32(1790048), v159)
	mBase = m.M
	goto L44
L44:
	;
	goto L41
L45:
	;
	v174 = v142 + int32(1)
	if v174 != int32(21) {
		v142 = v174
		goto L39
	} else {
		goto L46
	}
L46:
	;
	goto L40
L47:
	;
	v219 = int32(0)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)))
	if v220 == v219 {
		goto L53
	} else {
		goto L54
	}
L48:
	;
	v211 = F_getmissingattr(m, v186, int32(20), v22+int32(23))
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L11
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v216 = F_fastgetattr_3(m, v121, int32(20), v186, v22+int32(23))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L11
	} else {
		goto L52
	}
L51:
	;
	v218 = v211
	goto L47
L52:
	;
	v218 = v216
	goto L47
L53:
	;
	v223 = F_text_to_cstring(m, v218)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L56
	}
L54:
	;
	v227 = v219
	goto L55
L55:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v119)+196))
	v230 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v230 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v225 = F_stringToNode(m, v223)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v227 = v225
	goto L55
L58:
	;
	v234 = int32(4554240)
	v235 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v238 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v238
	v241 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L11
	} else {
		goto L61
	}
L59:
	;
	v293 = v230
	goto L60
L60:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v228)+16))
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309)+18)))
	if base.Ui32(v310&int32(2047)) <= base.Ui32(int32(20)) {
		goto L71
	} else {
		goto L72
	}
L61:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v241)+4)) = int64(-4294965047)
	v246 = v241 + int32(20)
	v249 = int32(0)
	goto L62
L62:
	;
	v266 = int32(100)
	v267 = v249 * v266
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v241)))
	goto L65
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v246))) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v241
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v235
	v293 = v241
	goto L60
L64:
	;
	F_populate_compact_attribute(m, v241, v249)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L68
	}
L65:
	;
	v276 = F__emscripten_memcpy_bulkmem(m, v267+(v246+v268<<(uint(int32(4))%32)), v267+int32(1790048), v266)
	mBase = m.M
	goto L67
L67:
	;
	goto L64
L68:
	;
	v281 = v249 + int32(1)
	if v281 != int32(21) {
		v249 = v281
		goto L62
	} else {
		goto L69
	}
L69:
	;
	goto L63
L70:
	;
	v326 = int32(0)
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)))
	if v328 == v326 {
		goto L76
	} else {
		goto L77
	}
L71:
	;
	v318 = F_getmissingattr(m, v293, int32(21), v22+int32(23))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L11
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v323 = F_fastgetattr_3(m, v228, int32(21), v293, v22+int32(23))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L11
	} else {
		goto L75
	}
L74:
	;
	v325 = v318
	goto L70
L75:
	;
	v325 = v323
	goto L70
L76:
	;
	v331 = F_text_to_cstring(m, v325)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L79
	}
L77:
	;
	v335 = v326
	goto L78
L78:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v119)+204))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340)+28)))
	if v341 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v333 = F_stringToNode(m, v331)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	v335 = v333
	goto L78
L81:
	;
	v342 = v22 + int32(24)
	goto L83
L82:
	;
	v342 = v22 + int32(28)
	goto L83
L83:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v344 = int32(*(*int16)(unsafe.Add(mBase, uint32(v343)+8)))
	if int32(0) < v344 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v347 = int32(0)
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+12)))
	v356 = v343
	v357 = v326
	v359 = v99
	v365 = v105
	v366 = v106
	goto L87
L85:
	;
	v423 = v99
	v429 = v105
	v430 = v106
	goto L86
L86:
	;
	F_pull_varattnos(m, v227, int32(1), v342)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L11
	} else {
		goto L104
	}
L87:
	;
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v356+v357<<(uint(int32(1))%32))+48)))
	if v376 == int32(0) {
		v409 = v356
		v410 = v359
		v412 = v365
		v413 = v366
		goto L89
	} else {
		goto L90
	}
L88:
	;
	v423 = v410
	v429 = v412
	v430 = v413
	goto L86
L89:
	;
	v415 = v357 + int32(1)
	v416 = int32(*(*int16)(unsafe.Add(mBase, uint32(v409)+8)))
	if v415 < v416 {
		v356 = v409
		v357 = v415
		v359 = v410
		v365 = v412
		v366 = v413
		goto L87
	} else {
		goto L103
	}
L90:
	;
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	v381 = v376 + int32(7)
	v382 = F_bms_add_member(m, v379, v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342))) = v382
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	if base.B2i32(v335 == v347)&(v349&base.B2i32(v227 == v347)) == int32(0) {
		v393 = v385
		v394 = v359
		goto L92
	} else {
		goto L93
	}
L92:
	;
	if v82 != v117 {
		v401 = v393
		v402 = v366
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v385)+10)))
	if v388 <= v357 {
		v393 = v385
		v394 = v359
		goto L92
	} else {
		goto L94
	}
L94:
	;
	v390 = F_bms_add_member(m, v359, v381)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v393 = v392
	v394 = v390
	goto L92
L96:
	;
	if v81 != v117 {
		v409 = v401
		v410 = v394
		v412 = v365
		v413 = v402
		goto L89
	} else {
		goto L100
	}
L97:
	;
	v396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v393)+10)))
	if v396 <= v357 {
		v401 = v393
		v402 = v366
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v398 = F_bms_add_member(m, v366, v381)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v401 = v400
	v402 = v398
	goto L96
L100:
	;
	v404 = int32(*(*int16)(unsafe.Add(mBase, uint32(v401)+10)))
	if v404 <= v357 {
		v409 = v401
		v410 = v394
		v412 = v365
		v413 = v402
		goto L89
	} else {
		goto L101
	}
L101:
	;
	v406 = F_bms_add_member(m, v365, v381)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v409 = v408
	v410 = v394
	v412 = v406
	v413 = v402
	goto L89
L103:
	;
	goto L88
L104:
	;
	F_pull_varattnos(m, v335, int32(1), v342)
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	F_relation_close(m, v119, int32(1))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	v447 = v109 + int32(1)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v447 < v448 {
		v99 = v423
		v105 = v429
		v106 = v430
		v109 = v447
		goto L32
	} else {
		goto L107
	}
L107:
	;
	goto L33
L108:
	;
	F_list_free(m, v469)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L11
	} else {
		goto L126
	}
L109:
	;
	v471 = F_equal(m, v71, v469)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	if v471 == int32(0) {
		goto L108
	} else {
		goto L111
	}
L111:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v82 != v475 {
		goto L108
	} else {
		goto L112
	}
L112:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v81 != v477 {
		goto L108
	} else {
		goto L113
	}
L113:
	;
	F_list_free(m, v469)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	F_list_free(m, v71)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	v483 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v483)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_bms_free(m, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	v490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_bms_free(m, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v495 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_bms_free(m, v495)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	v500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	F_bms_free(m, v500)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(0)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_bms_free(m, v505)
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v510 = int32(4554240)
	v511 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v514 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v514
	v516 = F_bms_copy(m, v455)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v516
	v519 = F_bms_copy(m, v462)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L11
	} else {
		goto L122
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v519
	v522 = F_bms_copy(m, v461)
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v522
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v526 = F_bms_copy(m, v525)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v526
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v530 = F_bms_copy(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	v532 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v532)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v530
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v511
	switch l1 {
	case 0:
		v581 = v455
		goto L1
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		goto L22
	}
L126:
	;
	F_list_free(m, v71)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L11
	} else {
		goto L127
	}
L127:
	;
	F_bms_free(m, v455)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	F_bms_free(m, v462)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	F_bms_free(m, v461)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L11
	} else {
		goto L130
	}
L130:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_bms_free(m, v547)
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L11
	} else {
		goto L131
	}
L131:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	F_bms_free(m, v550)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L11
	} else {
		goto L132
	}
L132:
	;
	v553 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L11
	} else {
		goto L133
	}
L133:
	;
	if v553 != 0 {
		v71 = v553
		goto L27
	} else {
		goto L134
	}
L134:
	;
	goto L28
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(58169), v22+int32(16))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(522810), int32(5555), int32(249760))
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L11
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
	v581 = v574
	goto L1
}
func F_RelationInitIndexAccessInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v20 int32
	_ = v20
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
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
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
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
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v283 int32
	_ = v283
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v538 int32
	_ = v538
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v614 int32
	_ = v614
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v642 int32
	_ = v642
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v658 int32
	_ = v658
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int64
	_ = v688
	var v690 int32
	_ = v690
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v729 int32
	_ = v729
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v788 int32
	_ = v788
	var v793 int32
	_ = v793
	v20 = m.G0
	v22 = v20 - int32(288)
	m.G0 = v22
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v26 = F_SearchSysCache1(m, int32(34), v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L6
	} else {
		goto L161
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L6
	} else {
		goto L158
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L6
	} else {
		goto L155
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L6
	} else {
		goto L152
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L6
	} else {
		goto L149
	}
L6:
	;
	return
L7:
	;
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v28 = int32(4554240)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v32
	v34 = F_heap_copytuple(m, v26)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v702 = m.ExcPending
	if v702 != 0 {
		goto L6
	} else {
		goto L146
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v37 + v38
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v29
	F_ReleaseCatCache(m, v26)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+84))
	v48 = F_SearchSysCache1(m, int32(2), v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v48 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+22)))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52+v53)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v55
	F_ReleaseCatCache(m, v48)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+120)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+8)))
	if v60 != v62 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v61)+10)))
	v65 = int32(0)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	v72 = F_AllocSetContextCreateInternal(m, v67, int32(253880), v65, int32(1024), int32(8192))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v72
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v78 = F_MemoryContextStrdup(m, v72, v75+int32(4))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v72)+36)) = v78
	goto L19
L19:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v82 = F_GetIndexAmRoutine(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v86 = F_MemoryContextAlloc(m, v84, int32(140))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	goto L23
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v89
	F_pfree(m, v82)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L26
	}
L23:
	;
	v89 = F__emscripten_memcpy_bulkmem(m, v86, v82, int32(140))
	mBase = m.M
	goto L25
L25:
	;
	goto L22
L26:
	;
	v95 = v64 << (uint(int32(2)) % 32)
	v96 = F_MemoryContextAllocZero(m, v72, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v96
	v99 = F_MemoryContextAllocZero(m, v72, v95)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+212)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)))
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v118
	v120 = F_MemoryContextAllocZero(m, v72, v95)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L6
	} else {
		goto L35
	}
L30:
	;
	v105 = v103 * base.I32_extend16_s(v60)
	v108 = F_MemoryContextAllocZero(m, v72, v105<<(uint(int32(2))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
	v118 = v65
	goto L29
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v108
	v113 = F_MemoryContextAllocZero(m, v72, v105*int32(28))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	v118 = v113
	goto L29
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+248)) = v120
	v124 = v64 << (uint(int32(1)) % 32)
	v125 = F_MemoryContextAllocZero(m, v72, v124)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L6
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v130 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v130 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v133 = int32(4554240)
	v134 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v137 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v137
	v140 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L40
	}
L38:
	;
	v192 = v130
	goto L39
L39:
	;
	v212 = F_fastgetattr_3(m, v128, int32(17), v192, v22+int32(79))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L49
	}
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v140)+4)) = int64(-4294965047)
	v148 = int32(0)
	goto L41
L41:
	;
	v166 = int32(100)
	v167 = v148 * v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	goto L44
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v140
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v134
	v192 = v140
	goto L39
L43:
	;
	F_populate_compact_attribute(m, v140, v148)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L6
	} else {
		goto L47
	}
L44:
	;
	v176 = F__emscripten_memcpy_bulkmem(m, v167+(v140+int32(20)+v168<<(uint(int32(4))%32)), v167+int32(1790048), v166)
	mBase = m.M
	goto L46
L46:
	;
	goto L43
L47:
	;
	v181 = v148 + int32(1)
	if v181 != int32(21) {
		v148 = v181
		goto L41
	} else {
		goto L48
	}
L48:
	;
	goto L42
L49:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	if v95 != 0 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v221 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v221 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	v217 = F__emscripten_memcpy_bulkmem(m, v214, v212+int32(24), v95)
	mBase = m.M
	goto L53
L52:
	;
	goto L53
L53:
	;
	goto L50
L54:
	;
	v225 = int32(4554240)
	v226 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v229 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v229
	v232 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L6
	} else {
		goto L57
	}
L55:
	;
	v283 = v221
	goto L56
L56:
	;
	v303 = F_fastgetattr_3(m, v219, int32(18), v283, v22+int32(79))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L66
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v232)+4)) = int64(-4294965047)
	v239 = int32(0)
	goto L58
L58:
	;
	v257 = int32(100)
	v258 = v239 * v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	goto L61
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v232)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v232
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v226
	v283 = v232
	goto L56
L60:
	;
	F_populate_compact_attribute(m, v232, v239)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L6
	} else {
		goto L64
	}
L61:
	;
	v267 = F__emscripten_memcpy_bulkmem(m, v258+(v232+int32(20)+v259<<(uint(int32(4))%32)), v258+int32(1790048), v257)
	mBase = m.M
	goto L63
L63:
	;
	goto L60
L64:
	;
	v272 = v239 + int32(1)
	if v272 != int32(21) {
		v239 = v272
		goto L58
	} else {
		goto L65
	}
L65:
	;
	goto L59
L66:
	;
	if int32(0) < v64 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v311 = v103 << (uint(int32(2)) % 32)
	v325 = int32(0)
	goto L70
L68:
	;
	goto L69
L69:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v596 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v596 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L70:
	;
	v339 = v325 << (uint(int32(2)) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v303+int32(24)+v339)))
	if v341 == int32(0) {
		goto L3
	} else {
		goto L72
	}
L71:
	;
	goto L69
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+284)) = v341
	v346 = *(*int32)(unsafe.Add(mBase, _consts[1375]))
	if v346 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v350 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	if v350 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	v366 = v346
	goto L75
L75:
	;
	v372 = F_hash_search(m, v366, v22+int32(284), int32(1), v22+int32(80))
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L6
	} else {
		goto L81
	}
L76:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L6
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v22)+96)) = int64(85899345924)
	v363 = F_hash_create(m, int32(417446), int32(64), v22+int32(80), int32(40))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L6
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1375])) = v363
	v366 = v363
	goto L75
L81:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+80)))
	if v374 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v339+v308))) = v560
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v339+v307))) = v563
	if v103 != 0 {
		goto L120
	} else {
		goto L121
	}
L83:
	;
	if v103 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v377 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v372)+16)) = v377
	*(*uint16)(unsafe.Add(mBase, uint32(v372)+6)) = uint16(v103)
	*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)) = uint8(v377)
	v389 = int32(1)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v383 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)))
	if v383&int32(1) != 0 {
		goto L82
	} else {
		goto L87
	}
L87:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v389 = base.B2i32(v386 == int32(0))
	goto L83
L88:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	v402 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1361])))
	if v402 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L89:
	;
	if v389 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v395 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	v396 = F_MemoryContextAllocZero(m, v395, v311)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L6
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v372)+16)) = v396
	goto L88
L92:
	;
	v410 = base.B2i32(v400 != int32(1981)) & base.B2i32(v400 != int32(1979))
	goto L94
L93:
	;
	v410 = int32(1)
	goto L94
L94:
	;
	F_ScanKeyInit(m, v22+int32(128), int32(1), int32(3), int32(184), v400)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	v420 = F_table_open(m, int32(2616), int32(1))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	v427 = F_systable_beginscan(m, v420, int32(2687), v410, int32(0), int32(1), v22+int32(128))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L97
	}
L97:
	;
	v429 = F_systable_getnext(m, v427)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	if v429 == int32(0) {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v429)+16))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+22)))
	v435 = v433 + v434
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v372)+8)) = v436
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v435)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v372)+12)) = v438
	F_systable_endscan(m, v427)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_sequence_close(m, v420, int32(1))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	if v103 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v372)+8))
	F_ScanKeyInit(m, v22+int32(128), int32(2), int32(3), int32(184), v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L6
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v538 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)) = uint8(v538)
	goto L82
L105:
	;
	v453 = int32(3)
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	F_ScanKeyInit(m, v22+int32(176), v453, v453, int32(184), v456)
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L6
	} else {
		goto L106
	}
L106:
	;
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v372)+12))
	F_ScanKeyInit(m, v22+int32(224), int32(4), int32(3), int32(184), v462)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L6
	} else {
		goto L107
	}
L107:
	;
	v467 = F_table_open(m, int32(2603), int32(1))
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L6
	} else {
		goto L108
	}
L108:
	;
	v474 = F_systable_beginscan(m, v467, int32(2655), v410, int32(0), int32(3), v22+int32(128))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L6
	} else {
		goto L109
	}
L109:
	;
	goto L110
L110:
	;
	v495 = F_systable_getnext(m, v474)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L6
	} else {
		goto L112
	}
L111:
	;
	F_systable_endscan(m, v474)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L6
	} else {
		goto L118
	}
L112:
	;
	if v495 != 0 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v495)+16))
	v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497)+22)))
	v499 = v497 + v498
	v500 = int32(*(*int16)(unsafe.Add(mBase, uint32(v499)+16)))
	if v500 <= int32(0) {
		goto L1
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	goto L111
L116:
	;
	v504 = v500 & int32(65535)
	if base.Ui32(v103) < base.Ui32(v504) {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v499)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v506+v504<<(uint(int32(2))%32)-int32(4)))) = v512
	goto L110
L118:
	;
	F_sequence_close(m, v467, int32(1))
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L6
	} else {
		goto L119
	}
L119:
	;
	goto L104
L120:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v372)+16))
	if v311 != 0 {
		goto L124
	} else {
		goto L125
	}
L121:
	;
	goto L122
L122:
	;
	v573 = v325 + int32(1)
	if v573 != v64 {
		v325 = v573
		goto L70
	} else {
		goto L127
	}
L123:
	;
	goto L122
L124:
	;
	v570 = F__emscripten_memcpy_bulkmem(m, v309+v325*v103<<(uint(int32(2))%32), v569, v311)
	mBase = m.M
	goto L126
L125:
	;
	goto L126
L126:
	;
	goto L123
L127:
	;
	goto L71
L128:
	;
	v599 = int32(4554240)
	v600 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v603 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v603
	v606 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L6
	} else {
		goto L131
	}
L129:
	;
	v658 = v596
	goto L130
L130:
	;
	v678 = F_fastgetattr_3(m, v594, int32(19), v658, v22+int32(79))
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L6
	} else {
		goto L140
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v606)+4)) = int64(-4294965047)
	v614 = int32(0)
	goto L132
L132:
	;
	v632 = int32(100)
	v633 = v614 * v632
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v606)))
	goto L135
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v606)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v606
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v600
	v658 = v606
	goto L130
L134:
	;
	F_populate_compact_attribute(m, v606, v614)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L6
	} else {
		goto L138
	}
L135:
	;
	v642 = F__emscripten_memcpy_bulkmem(m, v633+(v606+int32(20)+v634<<(uint(int32(4))%32)), v633+int32(1790048), v632)
	mBase = m.M
	goto L137
L137:
	;
	goto L134
L138:
	;
	v647 = v614 + int32(1)
	if v647 != int32(21) {
		v614 = v647
		goto L132
	} else {
		goto L139
	}
L139:
	;
	goto L133
L140:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	if v124 != 0 {
		goto L142
	} else {
		goto L143
	}
L141:
	;
	v686 = F_RelationGetIndexAttOptions(m, l0, int32(0))
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L6
	} else {
		goto L145
	}
L142:
	;
	v683 = F__emscripten_memcpy_bulkmem(m, v680, v678+int32(24), v124)
	mBase = m.M
	goto L144
L143:
	;
	goto L144
L144:
	;
	goto L141
L145:
	;
	v688 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+228)) = v688
	v690 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v690
	*(*int64)(unsafe.Add(mBase, uint32(l0)+236)) = v688
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v690
	m.G0 = v22 + int32(288)
	return
L146:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v703
	F_errmsg_internal(m, int32(42843), v22)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L6
	} else {
		goto L147
	}
L147:
	;
	F_errfinish(m, int32(522810), int32(1471), int32(254147))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L149:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v718
	F_errmsg_internal(m, int32(58072), v22+int32(16))
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L6
	} else {
		goto L150
	}
L150:
	;
	F_errfinish(m, int32(522810), int32(1485), int32(254147))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L151
	}
L151:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L152:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+64)) = v734
	F_errmsg_internal(m, int32(42754), v22-int32(-64))
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L6
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(522810), int32(1493), int32(254147))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L154
	}
L154:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L155:
	;
	F_errmsg_internal(m, int32(401611), int32(0))
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L6
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(522810), int32(1630), int32(358363))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L6
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = v763
	F_errmsg_internal(m, int32(45877), v22+int32(32))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L6
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(522810), int32(1766), int32(254175))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L6
	} else {
		goto L160
	}
L160:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L161:
	;
	v780 = int32(*(*int16)(unsafe.Add(mBase, uint32(v499)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v780
	v782 = *(*int32)(unsafe.Add(mBase, uint32(v22)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = v782
	F_errmsg_internal(m, int32(45948), v22+int32(48))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L6
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(522810), int32(1800), int32(254175))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationMapFilenumberToOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v3 = int32(0)
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v113
L2:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
	v113 = v108
	goto L1
L3:
	;
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if v8 < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v55 = int32(0)
	v57 = *(*int32)(unsafe.Add(mBase, _consts[1260]))
	if v55 < v57 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v14 = v8
	goto L9
L7:
	;
	goto L8
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _consts[1261]))
	if v35 <= int32(0) {
		v113 = v3
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v19 = v14 << (uint(int32(3)) % 32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_consts[1384])))
	if v22 == l0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v107 = v19 + int32(4543568)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v27 = v14 + int32(1)
	if v27 != v10 {
		v14 = v27
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v40 = int32(0)
	goto L16
L16:
	;
	v45 = v40 << (uint(int32(3)) % 32)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+uint32(_consts[1385])))
	if v48 != l0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v107 = v45 + int32(4544092)
	goto L2
L18:
	;
	v51 = v40 + int32(1)
	if v35 != v51 {
		v40 = v51
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v113 = v3
	goto L1
L22:
	;
	v87 = int32(0)
	goto L32
L23:
	;
	v107 = v66 + int32(4542520)
	goto L2
L24:
	;
	v61 = v55
	goto L27
L25:
	;
	goto L26
L26:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[1262]))
	if v80 <= int32(0) {
		v113 = v3
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v66 = v61 << (uint(int32(3)) % 32)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)+uint32(_consts[1386])))
	if l0 == v69 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v72 = v61 + int32(1)
	if v72 != v57 {
		v61 = v72
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L22
L32:
	;
	v92 = v87 << (uint(int32(3)) % 32)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v92)+uint32(_consts[1387])))
	if v95 != l0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v107 = v92 + int32(4543044)
	goto L2
L34:
	;
	v98 = v87 + int32(1)
	if v80 != v98 {
		v87 = v98
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v113 = v3
	goto L1
}
func F_RelationMapOidToFilenumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v3 = int32(0)
	if l1 == v3 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v102
L2:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v102 = v97
	goto L1
L3:
	;
	v80 = int32(0)
	goto L27
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _consts[397]))
	if int32(0) < v9 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[1260]))
	if int32(0) < v37 {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	v14 = v3
	goto L10
L8:
	;
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[1261]))
	if v32 <= int32(0) {
		v102 = v3
		goto L1
	} else {
		goto L14
	}
L10:
	;
	v18 = v14 << (uint(int32(3)) % 32)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v18)+uint32(_consts[398])))
	if l0 == v21 {
		v93 = v18 + int32(4543568)
		goto L2
	} else {
		goto L12
	}
L11:
	;
	goto L9
L12:
	;
	v24 = v14 + int32(1)
	if v24 != v9 {
		v14 = v24
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L3
L15:
	;
	v42 = v3
	goto L18
L16:
	;
	goto L17
L17:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[1262]))
	if v60 <= int32(0) {
		v102 = v3
		goto L1
	} else {
		goto L22
	}
L18:
	;
	v46 = v42 << (uint(int32(3)) % 32)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+uint32(_consts[1263])))
	if l0 == v49 {
		v93 = v46 + int32(4542520)
		goto L2
	} else {
		goto L20
	}
L19:
	;
	goto L17
L20:
	;
	v52 = v42 + int32(1)
	if v52 != v37 {
		v42 = v52
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v66 = int32(0)
	goto L23
L23:
	;
	v70 = v66 << (uint(int32(3)) % 32)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v70)+uint32(_consts[1264])))
	if l0 == v73 {
		v93 = v70 + int32(4543044)
		goto L2
	} else {
		goto L25
	}
L24:
	;
	v102 = v3
	goto L1
L25:
	;
	v76 = v66 + int32(1)
	if v60 != v76 {
		v66 = v76
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v84 = v80 << (uint(int32(3)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v84)+uint32(_consts[1265])))
	if l0 == v87 {
		v93 = v84 + int32(4544092)
		goto L2
	} else {
		goto L29
	}
L28:
	;
	v102 = v3
	goto L1
L29:
	;
	v90 = v80 + int32(1)
	if v32 != v90 {
		v80 = v90
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
}
func F_RelationMapOidToFilenumberForDatabase(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(528)
	m.G0 = v9
	F_read_relmap_file(m, v9+int32(4), l0, v3, int32(21))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v19 <= int32(0) {
		v42 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(528)
	return v42
L4:
	;
	v26 = v3
	goto L6
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v42 = v39
	goto L3
L6:
	;
	v32 = v9 + int32(12) + v26<<(uint(int32(3))%32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if l1 == v33 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v42 = int32(0)
	goto L3
L8:
	;
	v36 = v26 + int32(1)
	if v36 != v19 {
		v26 = v36
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_RelationTruncateIndexes(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v329 int32
	_ = v329
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	v19 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v19 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v23 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v39 = int32(0)
	goto L6
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v19)+12))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v44+v39<<(uint(int32(2))%32))))
	v50 = F_index_open(m, v48, int32(8))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_RelationTruncate(m, v50, int32(0))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L2
	} else {
		goto L68
	}
L9:
	;
	v53 = m.G0
	v55 = v53 - int32(16)
	m.G0 = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v50)+192))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+8)))
	if base.Ui32(int32(65503)) < base.Ui32((v58-int32(33))&int32(65535)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v66 = v58 & int32(3)
	v68 = v57 + int32(48)
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v57)+10)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+84))
	v72 = int32(0)
	v73 = m.G0
	v75 = v73 - int32(16)
	m.G0 = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v50)+196))
	if v77 == v72 {
		v242 = v72
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L2
	} else {
		goto L65
	}
L13:
	;
	m.G0 = v75 + int32(16)
	v257 = int32(0)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+12)))
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+13)))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+20)))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v50)+204))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v262)+28)))
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+15)))
	v266 = F_makeIndexInfo(m, v58, v69, v71, v242, v257, v258, v259, v260, v257, v263, v258&v264)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L2
	} else {
		goto L52
	}
L14:
	;
	v82 = F_heap_attisnull(m, v77, int32(20), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v82 != 0 {
		v242 = v72
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v84 = int32(0)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v50)+196))
	v87 = *(*int32)(unsafe.Add(mBase, _consts[381]))
	if v87 == v84 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v90 = int32(4554240)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v94 = *(*int32)(unsafe.Add(mBase, _consts[382]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v94
	v97 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v146 = v87
	goto L19
L19:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+18)))
	if base.Ui32(v164&int32(2044)) <= base.Ui32(int32(19)) {
		goto L30
	} else {
		goto L31
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v97)+4)) = int64(-4294965047)
	v106 = v84
	goto L21
L21:
	;
	v121 = int32(100)
	v122 = v106 * v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	goto L24
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[381])) = v97
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v91
	v146 = v97
	goto L19
L23:
	;
	F_populate_compact_attribute(m, v97, v106)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L2
	} else {
		goto L27
	}
L24:
	;
	v131 = F__emscripten_memcpy_bulkmem(m, v122+(v97+int32(20)+v123<<(uint(int32(4))%32)), v122+int32(1790048), v121)
	mBase = m.M
	goto L26
L26:
	;
	goto L23
L27:
	;
	v136 = v106 + int32(1)
	if v136 != int32(21) {
		v106 = v136
		goto L21
	} else {
		goto L28
	}
L28:
	;
	goto L22
L29:
	;
	v180 = F_text_to_cstring(m, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L2
	} else {
		goto L35
	}
L30:
	;
	v172 = F_getmissingattr(m, v146, int32(20), v75+int32(15))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v177 = F_fastgetattr_3(m, v85, int32(20), v146, v75+int32(15))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L34
	}
L33:
	;
	v179 = v172
	goto L29
L34:
	;
	v179 = v177
	goto L29
L35:
	;
	v182 = F_stringToNode(m, v180)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_pfree(m, v180)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	if v182 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v242 = int32(0)
	goto L13
L39:
	;
	goto L40
L40:
	;
	v189 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v190 <= v189 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v242 = int32(0)
	goto L13
L42:
	;
	goto L43
L43:
	;
	v198 = v189
	v201 = int32(0)
	goto L44
L44:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v198<<(uint(int32(2))%32))))
	v218 = F_exprType(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L46
	}
L45:
	;
	v242 = v230
	goto L13
L46:
	;
	v220 = F_exprTypmod(m, v217)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v222 = F_exprCollation(m, v217)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v224 = int32(1)
	v228 = F_makeConst(m, v218, v220, v222, v224, int32(0), v224, v224)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v230 = F_lappend(m, v201, v228)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L2
	} else {
		goto L50
	}
L50:
	;
	v233 = v198 + int32(1)
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
	if v233 < v234 {
		v198 = v233
		v201 = v230
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L45
L52:
	;
	v269 = v266 + int32(12)
	v270 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v58-int32(1)) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v279 = int32(0)
	v282 = v270
	goto L56
L54:
	;
	v329 = v270
	goto L55
L55:
	;
	if v66 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	v297 = v282 << (uint(int32(1)) % 32)
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297+v68))))
	*(*uint16)(unsafe.Add(mBase, uint32(v269+v297))) = uint16(v300)
	v303 = v297 | int32(2)
	v306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303+v68))))
	*(*uint16)(unsafe.Add(mBase, uint32(v269+v303))) = uint16(v306)
	v308 = int32(4)
	v309 = v297 | v308
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309+v68))))
	*(*uint16)(unsafe.Add(mBase, uint32(v269+v309))) = uint16(v312)
	v315 = v297 | int32(6)
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v315+v68))))
	*(*uint16)(unsafe.Add(mBase, uint32(v269+v315))) = uint16(v318)
	v321 = v282 + v308
	v323 = v279 + v308
	if v323 != v58&int32(60) {
		v279 = v323
		v282 = v321
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v329 = v321
	goto L55
L58:
	;
	goto L57
L59:
	;
	v347 = v329
	v357 = int32(0)
	goto L62
L60:
	;
	goto L61
L61:
	;
	m.G0 = v55 + int32(16)
	goto L8
L62:
	;
	v361 = int32(1)
	v362 = v347 << (uint(v361) % 32)
	v365 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362+v68))))
	*(*uint16)(unsafe.Add(mBase, uint32(v269+v362))) = uint16(v365)
	v370 = v357 + v361
	if v370 != v66 {
		v347 = v347 + v361
		v357 = v370
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L61
L64:
	;
	goto L63
L65:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v58
	F_errmsg_internal(m, int32(42876), v55)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L2
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(514633), int32(2499), int32(253918))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L2
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_index_build(m, l0, v50, v266, int32(1), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L2
	} else {
		goto L69
	}
L69:
	;
	F_relation_close(m, v50, int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L2
	} else {
		goto L70
	}
L70:
	;
	v419 = v39 + int32(1)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	if v419 < v420 {
		v39 = v419
		goto L6
	} else {
		goto L71
	}
L71:
	;
	goto L7
}
func F_SetRelationTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v21)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v27 = v25 + v26
				v30 = *(*int32)(unsafe.Add(mBase, _consts[126]))
				if l1 != v30 {
					v32 = l1
				} else {
					v32 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v32
				if l2 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = l2
				} else {
				}
				F_CatalogTupleUpdate(m, v16, v11+int32(8), v19)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_UnlockTuple(m, v16, v11+int32(8), int32(7))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+119)))
						switch v45 - int32(83) {
						case 0, 22, 26, 31, 33:
							F_pfree(m, v19)
							mBase = m.M
							v76 = m.ExcPending
							if v76 != 0 {
								return
							} else {
								F_sequence_close(m, v16, int32(3))
								mBase = m.M
								v79 = m.ExcPending
								if v79 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						default:
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
							v51 = F_table_open(m, int32(1214), int32(3))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								if v48 == int32(0) {
									v63 = int32(0)
									F_shdepDropDependency(m, v51, int32(1259), v13, v63, int32(1), v63, v63, v63)
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										F_sequence_close(m, v51, int32(3))
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return
										} else {
											F_pfree(m, v19)
											mBase = m.M
											v76 = m.ExcPending
											if v76 != 0 {
												return
											} else {
												F_sequence_close(m, v16, int32(3))
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								} else {
									if v48 == int32(1663) {
										v63 = int32(0)
										F_shdepDropDependency(m, v51, int32(1259), v13, v63, int32(1), v63, v63, v63)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return
										} else {
											F_sequence_close(m, v51, int32(3))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_pfree(m, v19)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													F_sequence_close(m, v16, int32(3))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
													}
												}
											}
										}
									} else {
										F_shdepChangeDep(m, v51, int32(1259), v13, int32(1213), v48, int32(116))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return
										} else {
											F_sequence_close(m, v51, int32(3))
											mBase = m.M
											v72 = m.ExcPending
											if v72 != 0 {
												return
											} else {
												F_pfree(m, v19)
												mBase = m.M
												v76 = m.ExcPending
												if v76 != 0 {
													return
												} else {
													F_sequence_close(m, v16, int32(3))
													mBase = m.M
													v79 = m.ExcPending
													if v79 != 0 {
														return
													} else {
														m.G0 = v11 + int32(16)
														return
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
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
					F_errmsg_internal(m, int32(49952), v11)
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errfinish(m, int32(517363), int32(3767), int32(439262))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
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
	}
}
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v121 int32
	_ = v121
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v198 = F_quote_identifier(m, v25)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L3
	} else {
		goto L46
	}
L2:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v167 = F_get_namespace_name_or_temp(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L3
	} else {
		goto L41
	}
L3:
	;
	return int32(0)
L4:
	;
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = v21 + v22
	v25 = v23 + int32(4)
	if l1 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L38
	}
L8:
	;
	v134 = F_RelationIsVisible(m, l0)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L3
	} else {
		goto L35
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v31 = int32(0)
	if v31 < v28 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = v28
	goto L13
L12:
	;
	v34 = v31
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = int32(0)
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35+v42<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L8
L16:
	;
	v121 = v42 + int32(1)
	if v121 != v34 {
		v42 = v121
		goto L14
	} else {
		goto L34
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v54 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = v54
	goto L21
L20:
	;
	v60 = v57
	goto L21
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v64 = int32(0)
	goto L22
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v61+v64<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v82 == int32(0) {
		v101 = v81
		v102 = v82
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L16
L24:
	;
	if v102-v101 == int32(0) {
		goto L2
	} else {
		goto L32
	}
L25:
	;
	goto L24
L26:
	;
	if v81 != v82 {
		v101 = v81
		v102 = v82
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v86 = v78
	v87 = v25
	goto L28
L28:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v90
		v102 = v91
		goto L25
	} else {
		goto L30
	}
L29:
	;
	v101 = v90
	v102 = v91
	goto L25
L30:
	;
	v94 = int32(1)
	if v90 == v91 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v107 = v64 + int32(1)
	if v107 != v60 {
		v64 = v107
		goto L22
	} else {
		goto L33
	}
L33:
	;
	goto L23
L34:
	;
	goto L15
L35:
	;
	if v134 == int32(0) {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	F_initStringInfo(m, v14+int32(32))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	goto L1
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(49952), v14)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(516795), int32(13165), int32(396542))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_initStringInfo(m, v14+int32(32))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L3
	} else {
		goto L42
	}
L42:
	;
	if v167 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v175 = F_quote_identifier(m, v167)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v175
	F_appendStringInfo(m, v14+int32(32), int32(633500), v14+int32(16))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L3
	} else {
		goto L45
	}
L45:
	;
	goto L1
L46:
	;
	F_appendStringInfoString(m, v14+int32(32), v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	F_ReleaseCatCache(m, v17)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L3
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v14 + int32(48)
	return v202
}
func F_getRelationDescription(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			if l2 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(49952), v9)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errfinish(m, int32(516164), int32(4115), int32(259089))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = F_RelationIsVisible(m, l1)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v32 != 0 {
					v38 = int32(0)
					v41 = F_quote_qualified_identifier(m, v38, v31+int32(4))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
						switch v44 - int32(73) {
						case 0, 32:
							v55 = int32(188243)
						default:
							v55 = int32(194201)
						case 10:
							v55 = int32(206079)
						case 26:
							v55 = int32(202823)
						case 29:
							v55 = int32(205429)
						case 36:
							v55 = int32(188507)
						case 39, 41:
							v55 = int32(205463)
						case 43:
							v55 = int32(205327)
						case 45:
							v55 = int32(188520)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v41
						F_appendStringInfo(m, l0, v55, v9+int32(16))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
					v36 = F_get_namespace_name(m, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = v36
						v41 = F_quote_qualified_identifier(m, v38, v31+int32(4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
							switch v44 - int32(73) {
							case 0, 32:
								v55 = int32(188243)
							default:
								v55 = int32(194201)
							case 10:
								v55 = int32(206079)
							case 26:
								v55 = int32(202823)
							case 29:
								v55 = int32(205429)
							case 36:
								v55 = int32(188507)
							case 39, 41:
								v55 = int32(205463)
							case 43:
								v55 = int32(205327)
							case 45:
								v55 = int32(188520)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v41
							F_appendStringInfo(m, l0, v55, v9+int32(16))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_getRelationIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			if l3 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(49952), v9)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(516164), int32(6109), int32(10870))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if l2 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
				}
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
			v24 = v22 + v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
			v26 = F_get_namespace_name_or_temp(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v29 = v24 + int32(4)
				v30 = F_quote_qualified_identifier(m, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_appendStringInfoString(m, l0, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						if l2 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v26
							v35 = F_pstrdup(m, v29)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v35
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v35
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v39
								v45 = F_list_make2_impl(m, v9+int32(20), v9+int32(16))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v45
									F_ReleaseCatCache(m, v12)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_get_relation_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v15 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		F_ScanKeyInit(m, v11+int32(16), int32(9), int32(3), int32(184), l0)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_ScanKeyInit(m, v11-int32(-64), int32(10), int32(3), int32(184), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v11+int32(112), int32(2), int32(3), int32(62), l1)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v47 = F_systable_beginscan(m, v15, int32(2665), int32(1), int32(0), int32(3), v11+int32(16))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_systable_getnext(m, v47)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 != 0 {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
								v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+22)))
								v54 = *(*int32)(unsafe.Add(mBase, uint32(v51+v52)))
								v55 = v54
							} else {
								v55 = int32(0)
							}
							F_systable_endscan(m, v47)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if l2 != 0 {
									F_sequence_close(m, v15, int32(1))
									mBase = m.M
									v79 = m.ExcPending
									if v79 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(160)
										return v55
									}
								} else {
									if v55 != 0 {
										F_sequence_close(m, v15, int32(1))
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int32(0)
										} else {
											m.G0 = v11 + int32(160)
											return v55
										}
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v61 = m.ExcPending
										if v61 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(67137668))
											mBase = m.M
											v64 = m.ExcPending
											if v64 != 0 {
												return int32(0)
											} else {
												v65 = F_get_rel_name(m, l0)
												mBase = m.M
												v66 = m.ExcPending
												if v66 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v65
													*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
													F_errmsg(m, int32(78309), v11)
													mBase = m.M
													v71 = m.ExcPending
													if v71 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(515494), int32(1235), int32(453710))
														mBase = m.M
														v76 = m.ExcPending
														if v76 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
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
func F_relation_excluded_by_constraints(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
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
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v13 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v385
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v25 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _consts[642]))
	switch v67 {
	case 0:
		v385 = int32(0)
		goto L4
	case 1:
		goto L17
	case 2:
		goto L18
	default:
		v76 = v4
		goto L16
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21+v25<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v51 = v25 + int32(1)
	if v51 != v18 {
		v25 = v51
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v41 != int32(7) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
	if v45 != 0 {
		v385 = v44
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v46 == int32(0) {
		v385 = v44
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	goto L9
L16:
	;
	v78 = int32(0)
	if v78 < v18 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v76 = base.B2i32(v73 == int32(0))
	goto L16
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v68 == int32(2) {
		v76 = v4
		goto L16
	} else {
		goto L19
	}
L19:
	;
	return int32(0)
L20:
	;
	v85 = int32(0)
	v87 = v78
	goto L23
L21:
	;
	v119 = v78
	goto L22
L22:
	;
	v127 = F_predicate_refuted_by(m, v119, v119, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L25
	} else {
		goto L32
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v85<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = F_contain_mutable_functions(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v119 = v109
	goto L22
L25:
	;
	return int32(0)
L26:
	;
	if v100 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v107 = F_lappend(m, v87, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v109 = v87
	goto L29
L29:
	;
	v111 = v85 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v111 < v112 {
		v85 = v111
		v87 = v109
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v109 = v107
	goto L29
L31:
	;
	goto L24
L32:
	;
	if v127 != 0 {
		v385 = int32(1)
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v130 != 0 {
		v385 = int32(0)
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v131 = int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v132 == v131 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v138 = base.B2i32(v135 == int32(112))
	goto L37
L36:
	;
	v138 = v131
	goto L37
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v140 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v143 = F_table_open(m, v141, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	if v76 == int32(0) {
		v314 = v280
		goto L72
	} else {
		goto L73
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	if v146 == int32(0) {
		v280 = v140
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+14)))
	if v149 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	if v138 == int32(0) {
		v280 = v201
		goto L38
	} else {
		goto L60
	}
L42:
	;
	v201 = v140
	goto L41
L43:
	;
	goto L44
L44:
	;
	v156 = int32(0)
	v158 = v140
	goto L45
L45:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v168 = v165 + v156*int32(12)
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+9)))
	if v169 != int32(1) {
		v192 = v158
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v201 = v192
	goto L41
L47:
	;
	v194 = v156 + int32(1)
	if v194 != v149 {
		v156 = v194
		v158 = v192
		goto L45
	} else {
		goto L59
	}
L48:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+10)))
	if v172&v132 != 0 {
		v192 = v158
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	v175 = F_stringToNode(m, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v177 = F_eval_const_expressions(m, l0, v175)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	v180 = F_canonicalize_qual(m, v177, int32(1))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L25
	} else {
		goto L52
	}
L52:
	;
	if v139 != int32(1) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_ChangeVarNodes(m, v180, int32(1), v139)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v187 = F_make_ands_implicit(m, v180)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	v189 = F_list_concat(m, v158, v187)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L25
	} else {
		goto L58
	}
L58:
	;
	v192 = v189
	goto L47
L59:
	;
	goto L46
L60:
	;
	v210 = int32(1)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+16)))
	if v211 != v210 {
		v280 = v201
		goto L38
	} else {
		goto L61
	}
L61:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	if v215 <= int32(0) {
		v280 = v201
		goto L38
	} else {
		goto L62
	}
L62:
	;
	v221 = v210
	v223 = v201
	goto L63
L63:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v232 = v221 - int32(1)
	v235 = v230 + v232<<(uint(int32(4))%32)
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+31)))
	if v236 != int32(118) {
		v270 = v223
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v280 = v270
	goto L38
L65:
	;
	v273 = v221 + int32(1)
	if v273 <= v215 {
		v221 = v273
		v223 = v270
		goto L63
	} else {
		goto L71
	}
L66:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+29)))
	if v239 != 0 {
		v270 = v223
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v242 = F_palloc0(m, int32(20))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242))) = int32(52)
	v252 = v230 + v240<<(uint(int32(4))%32) + v232*int32(100)
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+88))
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v252)+96))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v252)+116))
	v257 = F_makeVar(m, v139, base.I32_extend16_s(v221), v253, v254, v255, int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v242)+16)) = int32(-1)
	v261 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+12)) = uint8(v261)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v242)+4)) = v257
	v266 = F_lappend(m, v223, v242)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	v270 = v266
	goto L65
L71:
	;
	goto L64
L72:
	;
	if v314 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L73:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v143)+48))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+131)))
	if v290 != int32(1) {
		v314 = v280
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v293 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	if v293 != 0 {
		v308 = v293
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v310 = F_list_concat(m, v280, v308)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L25
	} else {
		goto L86
	}
L76:
	;
	v294 = F_RelationGetPartitionQual(m, v143)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L25
	} else {
		goto L77
	}
L77:
	;
	if v294 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v308 = v298
	goto L75
L79:
	;
	goto L80
L80:
	;
	v299 = F_expression_planner(m, v294)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L25
	} else {
		goto L81
	}
L81:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v301 != int32(1) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	F_ChangeVarNodes(m, v299, int32(1), v301)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L25
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+248)) = v299
	v308 = v299
	goto L75
L85:
	;
	goto L84
L86:
	;
	v314 = v310
	goto L72
L87:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v375 = F_predicate_refuted_by(m, v368, v373, int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L25
	} else {
		goto L104
	}
L88:
	;
	v317 = int32(0)
	F_sequence_close(m, v143, v317)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L25
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v321 = int32(0)
	v322 = F_expand_generated_columns_in_expr(m, v314, v143, v139)
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L25
	} else {
		goto L92
	}
L91:
	;
	v368 = v317
	goto L87
L92:
	;
	F_sequence_close(m, v143, int32(0))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	if v322 == int32(0) {
		v368 = v321
		goto L87
	} else {
		goto L94
	}
L94:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v329 <= int32(0) {
		v368 = v321
		goto L87
	} else {
		goto L95
	}
L95:
	;
	v336 = int32(0)
	v340 = v321
	goto L96
L96:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v322)+12))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v345+v336<<(uint(int32(2))%32))))
	v350 = F_contain_mutable_functions(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L25
	} else {
		goto L98
	}
L97:
	;
	v368 = v356
	goto L87
L98:
	;
	if v350 == int32(0) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v354 = F_lappend(m, v340, v349)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L25
	} else {
		goto L102
	}
L100:
	;
	v356 = v340
	goto L101
L101:
	;
	v358 = v336 + int32(1)
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v322)+4))
	if v358 < v359 {
		v336 = v358
		v340 = v356
		goto L96
	} else {
		goto L103
	}
L102:
	;
	v356 = v354
	goto L101
L103:
	;
	goto L97
L104:
	;
	v385 = v375
	goto L4
}
func F_relation_is_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
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
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v307 int32
	_ = v307
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v377 int32
	_ = v377
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = F_try_relation_open(m, l0, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(16)
	return v530
L4:
	;
	if v24 == int32(0) {
		v530 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v29 = int32(0)
	if l1 == v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v67 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		v60 = v29
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v67 = v60
	goto L6
L11:
	;
	v38 = int32(0)
	if v38 < v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v35
	goto L14
L13:
	;
	v41 = v38
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = int32(0)
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44<<(uint(int32(2))%32))))
	v53 = base.B2i32(v52 == v28)
	if v52 == v28 {
		v60 = v53
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v60 = v53
	goto L10
L17:
	;
	v55 = v44 + int32(1)
	if v55 != v41 {
		v44 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
	switch v72 - int32(112) {
	case 0, 2:
		goto L24
	default:
		goto L23
	}
L22:
	;
	v530 = v5
	goto L3
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v79 == int32(0) {
		v185 = v5
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v530 = int32(28)
	goto L3
L26:
	;
	if l2 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v82 <= int32(0) {
		v185 = v5
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v85 = int32(1)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v82 == v85 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if v82&v85 == int32(0) {
		v172 = v147
		goto L42
	} else {
		goto L43
	}
L30:
	;
	v145 = int32(0)
	v147 = v5
	goto L29
L31:
	;
	goto L32
L32:
	;
	v93 = int32(0)
	v99 = v93
	v101 = v5
	v102 = v93
	goto L33
L33:
	;
	v111 = v87 + v99<<(uint(int32(2))%32)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+17)))
	if v113 != int32(1) {
		v123 = v101
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v145 = v137
	v147 = v135
	goto L29
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v111)+4))
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+17)))
	if v125 != int32(1) {
		v135 = v123
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)+8))
	if v116 != 0 {
		v123 = v101
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112)+4))
	v123 = int32(1)<<(uint(v118)%32)&int32(28) | v101
	goto L35
L38:
	;
	v136 = int32(2)
	v137 = v99 + v136
	v139 = v102 + v136
	if v139 != v82&int32(2147483646) {
		v99 = v137
		v101 = v135
		v102 = v139
		goto L33
	} else {
		goto L41
	}
L39:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+8))
	if v128 != 0 {
		v135 = v123
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v135 = int32(1)<<(uint(v130)%32)&int32(28) | v123
	goto L38
L41:
	;
	goto L34
L42:
	;
	v173 = int32(28)
	if v172 != v173 {
		v185 = v172
		goto L26
	} else {
		goto L46
	}
L43:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v87+v145<<(uint(int32(2))%32))))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160)+17)))
	if v161 != int32(1) {
		v172 = v147
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
	if v164 != 0 {
		v172 = v147
		goto L42
	} else {
		goto L45
	}
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	v172 = int32(1)<<(uint(v166)%32)&int32(28) | v147
	goto L42
L46:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v530 = v173
	goto L3
L48:
	;
	switch v72 - int32(102) {
	case 0:
		goto L65
	default:
		v513 = v217
		goto L63
	case 16:
		goto L64
	}
L49:
	;
	v217 = v185
	goto L48
L50:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	if v195 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+10)))
	if v200 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v201 = v185 | int32(8)
	goto L54
L53:
	;
	v201 = v185
	goto L54
L54:
	;
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+15)))
	if v204 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v205 = v201 | int32(4)
	goto L57
L56:
	;
	v205 = v201
	goto L57
L57:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+20)))
	if v208 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v209 = v205 | int32(16)
	goto L60
L59:
	;
	v209 = v205
	goto L60
L60:
	;
	if v209 != int32(28) {
		v217 = v209
		goto L48
	} else {
		goto L61
	}
L61:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v530 = int32(28)
	goto L3
L63:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L1
	} else {
		goto L176
	}
L64:
	;
	v245 = F_get_view_query(m, v24)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L82
	}
L65:
	;
	v222 = F_GetFdwRoutineForRelation(m, v24, int32(0))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L67
	}
L66:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L81
	}
L67:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v222)+84))
	if v224 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v225 = m.T0[v224].(func(*base.Module, int32) int32)(m, v24)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v222)+52))
	if v230 != 0 {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	v241 = v225 | v217
	goto L66
L72:
	;
	v231 = v217 | int32(8)
	goto L74
L73:
	;
	v231 = v217
	goto L74
L74:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v222)+64))
	if v234 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v235 = v231 | int32(4)
	goto L77
L76:
	;
	v235 = v231
	goto L77
L77:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v222)+68))
	if v238 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v239 = v235 | int32(16)
	goto L80
L79:
	;
	v239 = v235
	goto L80
L80:
	;
	v241 = v239
	goto L66
L81:
	;
	v530 = v241
	goto L3
L82:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v245)+120))
	if v253 != 0 {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v387 != 0 {
		v513 = v217
		goto L63
	} else {
		goto L142
	}
L84:
	;
	v387 = int32(667113)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v255 = int32(666910)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v245)+100))
	if v256 != 0 {
		v377 = v255
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v387 = v377
	goto L83
L88:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v245)+108))
	if v257 != 0 {
		v377 = v255
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v245)+112))
	if v258 != 0 {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v387 = int32(667227)
	goto L83
L91:
	;
	goto L92
L92:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v245)+144))
	if v260 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v387 = int32(666969)
	goto L83
L94:
	;
	goto L95
L95:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v245)+48))
	if v262 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v387 = int32(667172)
	goto L83
L97:
	;
	goto L98
L98:
	;
	v264 = int32(667047)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v245)+128))
	if v265 != 0 {
		v377 = v264
		goto L87
	} else {
		goto L99
	}
L99:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v245)+132))
	if v266 != 0 {
		v377 = v264
		goto L87
	} else {
		goto L100
	}
L100:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+36)))
	if v267 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v387 = int32(666696)
	goto L83
L102:
	;
	goto L103
L103:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+37)))
	if v269 != 0 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v387 = int32(666553)
	goto L83
L105:
	;
	goto L106
L106:
	;
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245)+38)))
	if v271 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v387 = int32(666621)
	goto L83
L108:
	;
	goto L109
L109:
	;
	v273 = int32(666467)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v245)+60))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+4))
	if v275 == int32(0) {
		v377 = v273
		goto L87
	} else {
		goto L110
	}
L110:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v278 != int32(1) {
		v377 = v273
		goto L87
	} else {
		goto L111
	}
L111:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v275)+12))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v281)))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
	if v283 != int32(63) {
		v377 = v273
		goto L87
	} else {
		goto L112
	}
L112:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v245)+52))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v286)+12))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v287+v288<<(uint(int32(2))%32)-int32(4))))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)+12))
	if v295 != 0 {
		v377 = v273
		goto L87
	} else {
		goto L113
	}
L113:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+21)))
	v298 = v296 - int32(102)
	v307 = (v298<<(uint(int32(7))%32) | int32(base.Ui32(v298&int32(254))>>(uint(int32(1))%32))) & int32(255)
	if base.Ui32(int32(8)) < base.Ui32(v307) {
		v377 = v273
		goto L87
	} else {
		goto L114
	}
L114:
	;
	if int32(1)<<(uint(v307)%32)&int32(353) == int32(0) {
		v377 = v273
		goto L87
	} else {
		goto L115
	}
L115:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v294)+32))
	if v318 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v319 = int32(667284)
	goto L118
L117:
	;
	v319 = int32(0)
	goto L118
L118:
	;
	v377 = v319
	goto L87
L142:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v245)+60))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+4))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+12))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	v394 = v17 + int32(12)
	if v394 != 0 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = int32(0)
	goto L145
L144:
	;
	goto L145
L145:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v245)+76))
	if v397 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if l3 != 0 {
		goto L163
	} else {
		goto L164
	}
L147:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v400 <= int32(0) {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v410 = int32(7)
	v415 = int32(0)
	goto L149
L149:
	;
	v419 = v410 + int32(1)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v420+v415<<(uint(int32(2))%32))))
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v424)+26)))
	if v425 != 0 {
		goto L152
	} else {
		goto L153
	}
L150:
	;
	goto L146
L151:
	;
	v451 = v415 + int32(1)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
	if v451 < v452 {
		v410 = v419
		v415 = v451
		goto L149
	} else {
		goto L162
	}
L152:
	;
	v447 = F_bms_is_member(m, base.I32_extend16_s(v419), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L160
	}
L153:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
	if v427 != int32(6) {
		goto L152
	} else {
		goto L154
	}
L154:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v392)+4))
	if v430 != v431 {
		goto L152
	} else {
		goto L155
	}
L155:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v426)+28))
	if v433 != 0 {
		goto L152
	} else {
		goto L156
	}
L156:
	;
	v434 = int32(*(*int16)(unsafe.Add(mBase, uint32(v426)+8)))
	if v434 <= int32(0) {
		goto L152
	} else {
		goto L157
	}
L157:
	;
	if v394 == int32(0) {
		goto L151
	} else {
		goto L158
	}
L158:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v441 = F_bms_add_member(m, v439, base.I32_extend16_s(v419))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v394))) = v441
	goto L151
L160:
	;
	if v447 != 0 {
		goto L146
	} else {
		goto L161
	}
L161:
	;
	goto L151
L162:
	;
	goto L150
L163:
	;
	v469 = F_bms_int_members(m, v468, l3)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L166
	}
L164:
	;
	v472 = v468
	goto L165
L165:
	;
	if v472 != 0 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v469
	v472 = v469
	goto L165
L167:
	;
	v475 = int32(28)
	goto L169
L168:
	;
	v475 = int32(16)
	goto L169
L169:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v245)+52))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v476)+12))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v245)+60))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v478)+4))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v479)+12))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v481)+4))
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v477+v482<<(uint(int32(2))%32)-int32(4))))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v488)+21)))
	switch v489 - int32(112) {
	case 0, 2:
		v507 = v475
		goto L170
	default:
		goto L171
	}
L170:
	;
	v513 = v507 | v217
	goto L63
L171:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v488)+16))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v494 = F_lappend_oid(m, l1, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L172
	}
L172:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v245)+76))
	v498 = F_adjust_view_column_set(m, v496, v497)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L1
	} else {
		goto L173
	}
L173:
	;
	v500 = F_relation_is_updatable(m, v492, v494, l2, v498)
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v502 = F_list_delete_last(m, v494)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v507 = v500 & v475
	goto L170
L176:
	;
	v530 = v513
	goto L3
}
func F_relation_statistics_update(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v38 int32
	_ = v38
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 float32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 float32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 float32
	_ = v159
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v2
	F_stats_check_required_arg(m, l0, int32(4160064), v2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		F_stats_check_required_arg(m, l0, int32(4160064), int32(1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v44 = F_text_to_cstring(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v47 = F_text_to_cstring(m, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, _consts[198])))
					if v51 == int32(1) {
						v56 = *(*int32)(unsafe.Add(mBase, _consts[199]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+316))
						v59 = base.B2i32(v57 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _consts[198])) = uint8(v59)
						v61 = v59
					} else {
						v61 = int32(0)
					}
					if v61 == int32(0) {
						v65 = F_makeRangeVar(m, v44, v47, int32(-1))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v72 = F_RangeVarGetRelidExtended(m, v65, int32(4), int32(0), int32(1060), v17+int32(-40))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
								if v74 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v78 = v77
								} else {
									v78 = v2
								}
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
								if v80 != 0 {
									v108 = v2
									v109 = float32(0)
									v110 = int32(1)
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
									if v111 == int32(0) {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v115 = v114
									} else {
										v115 = v2
									}
									v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))))
									if v118 == int32(0) {
										v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										v122 = v121
									} else {
										v122 = v2
									}
									v125 = F_table_open(m, int32(1259), int32(3))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										v128 = F_SearchSysCache1(m, int32(57), v72)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											if v128 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v254 = m.ExcPending
												if v254 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
													F_errmsg_internal(m, int32(443520), v19)
													mBase = m.M
													v258 = m.ExcPending
													if v258 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(516100), int32(144), int32(373006))
														mBase = m.M
														v263 = m.ExcPending
														if v263 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
												v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+22)))
												v134 = v132 + v133
												v135 = int32(0)
												v137 = v17 + int32(-16)
												v139 = v17 + int32(-32)
												if v74 != 0 {
													v154 = v135
													v155 = v137
													v156 = v139
												} else {
													v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+96))
													if v78 == v140 {
														v154 = v135
														v155 = v137
														v156 = v139
													} else {
														v144 = int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
														*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
														v154 = int32(1)
														v155 = v17 + int32(-16) | v144
														v156 = v17 + int32(-32) | v144
													}
												}
												if v108 == int32(0) {
													v166 = v154
												} else {
													v159 = *(*float32)(unsafe.Add(mBase, uint32(v134)+100))
													if base.F32_eq(v109, v159) != 0 {
														v166 = v154
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(11)
														*(*float32)(unsafe.Add(mBase, uint32(v156))) = v109
														v166 = v154 + int32(1)
													}
												}
												if v111 != 0 {
													v182 = v166
												} else {
													v167 = *(*int32)(unsafe.Add(mBase, uint32(v134)+104))
													if v115 == v167 {
														v182 = v166
													} else {
														v170 = v166 << (uint(int32(2)) % 32)
														*(*int32)(unsafe.Add(mBase, uint32(v170|(v17+int32(-32))))) = v115
														*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v170))) = int32(12)
														v182 = v166 + int32(1)
													}
												}
												if v118 != 0 {
													if v182 == int32(0) {
														F_ReleaseCatCache(m, v128)
														mBase = m.M
														v221 = m.ExcPending
														if v221 != 0 {
															return int32(0)
														} else {
															F_sequence_close(m, v125, int32(3))
															mBase = m.M
															v224 = m.ExcPending
															if v224 != 0 {
																return int32(0)
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v226 = m.ExcPending
																if v226 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v19 - int32(-64)
																	return v110
																}
															}
														}
													} else {
														v201 = v182
														v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
														v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v210)
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v128)
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v125, int32(3))
																		mBase = m.M
																		v224 = m.ExcPending
																		if v224 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v226 = m.ExcPending
																			if v226 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v19 - int32(-64)
																				return v110
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v184 = *(*int32)(unsafe.Add(mBase, uint32(v134)+108))
													if v122 == v184 {
														if v182 == int32(0) {
															F_ReleaseCatCache(m, v128)
															mBase = m.M
															v221 = m.ExcPending
															if v221 != 0 {
																return int32(0)
															} else {
																F_sequence_close(m, v125, int32(3))
																mBase = m.M
																v224 = m.ExcPending
																if v224 != 0 {
																	return int32(0)
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v226 = m.ExcPending
																	if v226 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v19 - int32(-64)
																		return v110
																	}
																}
															}
														} else {
															v201 = v182
															v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
															v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v210)
																	mBase = m.M
																	v217 = m.ExcPending
																	if v217 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v128)
																		mBase = m.M
																		v221 = m.ExcPending
																		if v221 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v125, int32(3))
																			mBase = m.M
																			v224 = m.ExcPending
																			if v224 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v226 = m.ExcPending
																				if v226 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v187 = v182 << (uint(int32(2)) % 32)
														*(*int32)(unsafe.Add(mBase, uint32(v187+(v17+int32(-32))))) = v122
														*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v187))) = int32(13)
														v201 = v182 + int32(1)
														v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
														v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
														mBase = m.M
														v211 = m.ExcPending
														if v211 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v210)
																mBase = m.M
																v217 = m.ExcPending
																if v217 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v128)
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v125, int32(3))
																		mBase = m.M
																		v224 = m.ExcPending
																		if v224 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v226 = m.ExcPending
																			if v226 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v19 - int32(-64)
																				return v110
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
									v81 = *(*float32)(unsafe.Add(mBase, uint32(l0)+44))
									if base.F32_lt(v81, float32(-1)) == int32(0) {
										v86 = int32(1)
										v108 = v86
										v109 = v81
										v110 = v86
										v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
										if v111 == int32(0) {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v115 = v114
										} else {
											v115 = v2
										}
										v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))))
										if v118 == int32(0) {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
											v122 = v121
										} else {
											v122 = v2
										}
										v125 = F_table_open(m, int32(1259), int32(3))
										mBase = m.M
										v126 = m.ExcPending
										if v126 != 0 {
											return int32(0)
										} else {
											v128 = F_SearchSysCache1(m, int32(57), v72)
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												if v128 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v254 = m.ExcPending
													if v254 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
														F_errmsg_internal(m, int32(443520), v19)
														mBase = m.M
														v258 = m.ExcPending
														if v258 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(516100), int32(144), int32(373006))
															mBase = m.M
															v263 = m.ExcPending
															if v263 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
													v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+22)))
													v134 = v132 + v133
													v135 = int32(0)
													v137 = v17 + int32(-16)
													v139 = v17 + int32(-32)
													if v74 != 0 {
														v154 = v135
														v155 = v137
														v156 = v139
													} else {
														v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+96))
														if v78 == v140 {
															v154 = v135
															v155 = v137
															v156 = v139
														} else {
															v144 = int32(4)
															*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
															*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
															v154 = int32(1)
															v155 = v17 + int32(-16) | v144
															v156 = v17 + int32(-32) | v144
														}
													}
													if v108 == int32(0) {
														v166 = v154
													} else {
														v159 = *(*float32)(unsafe.Add(mBase, uint32(v134)+100))
														if base.F32_eq(v109, v159) != 0 {
															v166 = v154
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(11)
															*(*float32)(unsafe.Add(mBase, uint32(v156))) = v109
															v166 = v154 + int32(1)
														}
													}
													if v111 != 0 {
														v182 = v166
													} else {
														v167 = *(*int32)(unsafe.Add(mBase, uint32(v134)+104))
														if v115 == v167 {
															v182 = v166
														} else {
															v170 = v166 << (uint(int32(2)) % 32)
															*(*int32)(unsafe.Add(mBase, uint32(v170|(v17+int32(-32))))) = v115
															*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v170))) = int32(12)
															v182 = v166 + int32(1)
														}
													}
													if v118 != 0 {
														if v182 == int32(0) {
															F_ReleaseCatCache(m, v128)
															mBase = m.M
															v221 = m.ExcPending
															if v221 != 0 {
																return int32(0)
															} else {
																F_sequence_close(m, v125, int32(3))
																mBase = m.M
																v224 = m.ExcPending
																if v224 != 0 {
																	return int32(0)
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v226 = m.ExcPending
																	if v226 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v19 - int32(-64)
																		return v110
																	}
																}
															}
														} else {
															v201 = v182
															v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
															v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v210)
																	mBase = m.M
																	v217 = m.ExcPending
																	if v217 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v128)
																		mBase = m.M
																		v221 = m.ExcPending
																		if v221 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v125, int32(3))
																			mBase = m.M
																			v224 = m.ExcPending
																			if v224 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v226 = m.ExcPending
																				if v226 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v184 = *(*int32)(unsafe.Add(mBase, uint32(v134)+108))
														if v122 == v184 {
															if v182 == int32(0) {
																F_ReleaseCatCache(m, v128)
																mBase = m.M
																v221 = m.ExcPending
																if v221 != 0 {
																	return int32(0)
																} else {
																	F_sequence_close(m, v125, int32(3))
																	mBase = m.M
																	v224 = m.ExcPending
																	if v224 != 0 {
																		return int32(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v226 = m.ExcPending
																		if v226 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v19 - int32(-64)
																			return v110
																		}
																	}
																}
															} else {
																v201 = v182
																v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
																	return int32(0)
																} else {
																	F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v210)
																		mBase = m.M
																		v217 = m.ExcPending
																		if v217 != 0 {
																			return int32(0)
																		} else {
																			F_ReleaseCatCache(m, v128)
																			mBase = m.M
																			v221 = m.ExcPending
																			if v221 != 0 {
																				return int32(0)
																			} else {
																				F_sequence_close(m, v125, int32(3))
																				mBase = m.M
																				v224 = m.ExcPending
																				if v224 != 0 {
																					return int32(0)
																				} else {
																					F_CommandCounterIncrement(m)
																					mBase = m.M
																					v226 = m.ExcPending
																					if v226 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v19 - int32(-64)
																						return v110
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v187 = v182 << (uint(int32(2)) % 32)
															*(*int32)(unsafe.Add(mBase, uint32(v187+(v17+int32(-32))))) = v122
															*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v187))) = int32(13)
															v201 = v182 + int32(1)
															v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
															v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
															mBase = m.M
															v211 = m.ExcPending
															if v211 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v210)
																	mBase = m.M
																	v217 = m.ExcPending
																	if v217 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v128)
																		mBase = m.M
																		v221 = m.ExcPending
																		if v221 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v125, int32(3))
																			mBase = m.M
																			v224 = m.ExcPending
																			if v224 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v226 = m.ExcPending
																				if v226 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
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
										v90 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											if v90 != 0 {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(173985)
													F_errmsg(m, int32(595758), v17+int32(-48))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(516100), int32(117), int32(373006))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v108 = v2
															v109 = v81
															v110 = int32(0)
															v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
															if v111 == int32(0) {
																v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
																v115 = v114
															} else {
																v115 = v2
															}
															v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))))
															if v118 == int32(0) {
																v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
																v122 = v121
															} else {
																v122 = v2
															}
															v125 = F_table_open(m, int32(1259), int32(3))
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return int32(0)
															} else {
																v128 = F_SearchSysCache1(m, int32(57), v72)
																mBase = m.M
																v129 = m.ExcPending
																if v129 != 0 {
																	return int32(0)
																} else {
																	if v128 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v254 = m.ExcPending
																		if v254 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
																			F_errmsg_internal(m, int32(443520), v19)
																			mBase = m.M
																			v258 = m.ExcPending
																			if v258 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(516100), int32(144), int32(373006))
																				mBase = m.M
																				v263 = m.ExcPending
																				if v263 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	} else {
																		v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
																		v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+22)))
																		v134 = v132 + v133
																		v135 = int32(0)
																		v137 = v17 + int32(-16)
																		v139 = v17 + int32(-32)
																		if v74 != 0 {
																			v154 = v135
																			v155 = v137
																			v156 = v139
																		} else {
																			v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+96))
																			if v78 == v140 {
																				v154 = v135
																				v155 = v137
																				v156 = v139
																			} else {
																				v144 = int32(4)
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
																				v154 = int32(1)
																				v155 = v17 + int32(-16) | v144
																				v156 = v17 + int32(-32) | v144
																			}
																		}
																		if v108 == int32(0) {
																			v166 = v154
																		} else {
																			v159 = *(*float32)(unsafe.Add(mBase, uint32(v134)+100))
																			if base.F32_eq(v109, v159) != 0 {
																				v166 = v154
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(11)
																				*(*float32)(unsafe.Add(mBase, uint32(v156))) = v109
																				v166 = v154 + int32(1)
																			}
																		}
																		if v111 != 0 {
																			v182 = v166
																		} else {
																			v167 = *(*int32)(unsafe.Add(mBase, uint32(v134)+104))
																			if v115 == v167 {
																				v182 = v166
																			} else {
																				v170 = v166 << (uint(int32(2)) % 32)
																				*(*int32)(unsafe.Add(mBase, uint32(v170|(v17+int32(-32))))) = v115
																				*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v170))) = int32(12)
																				v182 = v166 + int32(1)
																			}
																		}
																		if v118 != 0 {
																			if v182 == int32(0) {
																				F_ReleaseCatCache(m, v128)
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
																					return int32(0)
																				} else {
																					F_sequence_close(m, v125, int32(3))
																					mBase = m.M
																					v224 = m.ExcPending
																					if v224 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v226 = m.ExcPending
																						if v226 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v19 - int32(-64)
																							return v110
																						}
																					}
																				}
																			} else {
																				v201 = v182
																				v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																				v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																				mBase = m.M
																				v211 = m.ExcPending
																				if v211 != 0 {
																					return int32(0)
																				} else {
																					F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v210)
																						mBase = m.M
																						v217 = m.ExcPending
																						if v217 != 0 {
																							return int32(0)
																						} else {
																							F_ReleaseCatCache(m, v128)
																							mBase = m.M
																							v221 = m.ExcPending
																							if v221 != 0 {
																								return int32(0)
																							} else {
																								F_sequence_close(m, v125, int32(3))
																								mBase = m.M
																								v224 = m.ExcPending
																								if v224 != 0 {
																									return int32(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v226 = m.ExcPending
																									if v226 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v19 - int32(-64)
																										return v110
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v184 = *(*int32)(unsafe.Add(mBase, uint32(v134)+108))
																			if v122 == v184 {
																				if v182 == int32(0) {
																					F_ReleaseCatCache(m, v128)
																					mBase = m.M
																					v221 = m.ExcPending
																					if v221 != 0 {
																						return int32(0)
																					} else {
																						F_sequence_close(m, v125, int32(3))
																						mBase = m.M
																						v224 = m.ExcPending
																						if v224 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v226 = m.ExcPending
																							if v226 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v19 - int32(-64)
																								return v110
																							}
																						}
																					}
																				} else {
																					v201 = v182
																					v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																					v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																					mBase = m.M
																					v211 = m.ExcPending
																					if v211 != 0 {
																						return int32(0)
																					} else {
																						F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																						mBase = m.M
																						v215 = m.ExcPending
																						if v215 != 0 {
																							return int32(0)
																						} else {
																							F_pfree(m, v210)
																							mBase = m.M
																							v217 = m.ExcPending
																							if v217 != 0 {
																								return int32(0)
																							} else {
																								F_ReleaseCatCache(m, v128)
																								mBase = m.M
																								v221 = m.ExcPending
																								if v221 != 0 {
																									return int32(0)
																								} else {
																									F_sequence_close(m, v125, int32(3))
																									mBase = m.M
																									v224 = m.ExcPending
																									if v224 != 0 {
																										return int32(0)
																									} else {
																										F_CommandCounterIncrement(m)
																										mBase = m.M
																										v226 = m.ExcPending
																										if v226 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v19 - int32(-64)
																											return v110
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				v187 = v182 << (uint(int32(2)) % 32)
																				*(*int32)(unsafe.Add(mBase, uint32(v187+(v17+int32(-32))))) = v122
																				*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v187))) = int32(13)
																				v201 = v182 + int32(1)
																				v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																				v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																				mBase = m.M
																				v211 = m.ExcPending
																				if v211 != 0 {
																					return int32(0)
																				} else {
																					F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v210)
																						mBase = m.M
																						v217 = m.ExcPending
																						if v217 != 0 {
																							return int32(0)
																						} else {
																							F_ReleaseCatCache(m, v128)
																							mBase = m.M
																							v221 = m.ExcPending
																							if v221 != 0 {
																								return int32(0)
																							} else {
																								F_sequence_close(m, v125, int32(3))
																								mBase = m.M
																								v224 = m.ExcPending
																								if v224 != 0 {
																									return int32(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v226 = m.ExcPending
																									if v226 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v19 - int32(-64)
																										return v110
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
												v108 = v2
												v109 = v81
												v110 = int32(0)
												v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
												if v111 == int32(0) {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v115 = v114
												} else {
													v115 = v2
												}
												v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))))
												if v118 == int32(0) {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
													v122 = v121
												} else {
													v122 = v2
												}
												v125 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return int32(0)
												} else {
													v128 = F_SearchSysCache1(m, int32(57), v72)
													mBase = m.M
													v129 = m.ExcPending
													if v129 != 0 {
														return int32(0)
													} else {
														if v128 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v254 = m.ExcPending
															if v254 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
																F_errmsg_internal(m, int32(443520), v19)
																mBase = m.M
																v258 = m.ExcPending
																if v258 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(516100), int32(144), int32(373006))
																	mBase = m.M
																	v263 = m.ExcPending
																	if v263 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v128)+16))
															v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+22)))
															v134 = v132 + v133
															v135 = int32(0)
															v137 = v17 + int32(-16)
															v139 = v17 + int32(-32)
															if v74 != 0 {
																v154 = v135
																v155 = v137
																v156 = v139
															} else {
																v140 = *(*int32)(unsafe.Add(mBase, uint32(v134)+96))
																if v78 == v140 {
																	v154 = v135
																	v155 = v137
																	v156 = v139
																} else {
																	v144 = int32(4)
																	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
																	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
																	v154 = int32(1)
																	v155 = v17 + int32(-16) | v144
																	v156 = v17 + int32(-32) | v144
																}
															}
															if v108 == int32(0) {
																v166 = v154
															} else {
																v159 = *(*float32)(unsafe.Add(mBase, uint32(v134)+100))
																if base.F32_eq(v109, v159) != 0 {
																	v166 = v154
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v155))) = int32(11)
																	*(*float32)(unsafe.Add(mBase, uint32(v156))) = v109
																	v166 = v154 + int32(1)
																}
															}
															if v111 != 0 {
																v182 = v166
															} else {
																v167 = *(*int32)(unsafe.Add(mBase, uint32(v134)+104))
																if v115 == v167 {
																	v182 = v166
																} else {
																	v170 = v166 << (uint(int32(2)) % 32)
																	*(*int32)(unsafe.Add(mBase, uint32(v170|(v17+int32(-32))))) = v115
																	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v170))) = int32(12)
																	v182 = v166 + int32(1)
																}
															}
															if v118 != 0 {
																if v182 == int32(0) {
																	F_ReleaseCatCache(m, v128)
																	mBase = m.M
																	v221 = m.ExcPending
																	if v221 != 0 {
																		return int32(0)
																	} else {
																		F_sequence_close(m, v125, int32(3))
																		mBase = m.M
																		v224 = m.ExcPending
																		if v224 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v226 = m.ExcPending
																			if v226 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v19 - int32(-64)
																				return v110
																			}
																		}
																	}
																} else {
																	v201 = v182
																	v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																	v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v210)
																			mBase = m.M
																			v217 = m.ExcPending
																			if v217 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v128)
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
																					return int32(0)
																				} else {
																					F_sequence_close(m, v125, int32(3))
																					mBase = m.M
																					v224 = m.ExcPending
																					if v224 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v226 = m.ExcPending
																						if v226 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v19 - int32(-64)
																							return v110
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v184 = *(*int32)(unsafe.Add(mBase, uint32(v134)+108))
																if v122 == v184 {
																	if v182 == int32(0) {
																		F_ReleaseCatCache(m, v128)
																		mBase = m.M
																		v221 = m.ExcPending
																		if v221 != 0 {
																			return int32(0)
																		} else {
																			F_sequence_close(m, v125, int32(3))
																			mBase = m.M
																			v224 = m.ExcPending
																			if v224 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v226 = m.ExcPending
																				if v226 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	} else {
																		v201 = v182
																		v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																		v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																		mBase = m.M
																		v211 = m.ExcPending
																		if v211 != 0 {
																			return int32(0)
																		} else {
																			F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v210)
																				mBase = m.M
																				v217 = m.ExcPending
																				if v217 != 0 {
																					return int32(0)
																				} else {
																					F_ReleaseCatCache(m, v128)
																					mBase = m.M
																					v221 = m.ExcPending
																					if v221 != 0 {
																						return int32(0)
																					} else {
																						F_sequence_close(m, v125, int32(3))
																						mBase = m.M
																						v224 = m.ExcPending
																						if v224 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v226 = m.ExcPending
																							if v226 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v19 - int32(-64)
																								return v110
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v187 = v182 << (uint(int32(2)) % 32)
																	*(*int32)(unsafe.Add(mBase, uint32(v187+(v17+int32(-32))))) = v122
																	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v187))) = int32(13)
																	v201 = v182 + int32(1)
																	v203 = *(*int32)(unsafe.Add(mBase, uint32(v125)+52))
																	v210 = F_heap_modify_tuple_by_cols(m, v128, v203, v201, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v125, v210+int32(4), v210)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v210)
																			mBase = m.M
																			v217 = m.ExcPending
																			if v217 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v128)
																				mBase = m.M
																				v221 = m.ExcPending
																				if v221 != 0 {
																					return int32(0)
																				} else {
																					F_sequence_close(m, v125, int32(3))
																					mBase = m.M
																					v224 = m.ExcPending
																					if v224 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v226 = m.ExcPending
																						if v226 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v19 - int32(-64)
																							return v110
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v237 = m.ExcPending
							if v237 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(136074), int32(0))
								mBase = m.M
								v241 = m.ExcPending
								if v241 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(599484), int32(0))
									mBase = m.M
									v245 = m.ExcPending
									if v245 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(516100), int32(98), int32(373006))
										mBase = m.M
										v250 = m.ExcPending
										if v250 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
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
