package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_HashTables(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_HashTables[0]))
	v15 = v13 - int32(1)
	if int32(0) <= v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = v15
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	v26 = v21 << (uint(int32(2)) % 32)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AtEOSubXact_HashTables[1])))
	if l1 <= v27 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if int32(0) < v21 {
		v21 = v21 - int32(1)
		goto L4
	} else {
		goto L16
	}
L9:
	;
	v49 = int32(_a_F_AtEOSubXact_HashTables_0)
	v50 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_HashTables[0]))
	v52 = v50 << (uint(int32(2)) % 32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_AtEOSubXact_HashTables[2])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AtEOSubXact_HashTables[3]))) = v55
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_c_F_AtEOSubXact_HashTables[4])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AtEOSubXact_HashTables[1]))) = v59
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_HashTables[0])) = v50 - int32(1)
	goto L8
L10:
	;
	v35 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v35 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_c_F_AtEOSubXact_HashTables[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v39
	F_errmsg_internal(m, int32(_a_F_AtEOSubXact_HashTables_1), v10)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(_a_F_AtEOSubXact_HashTables_2), int32(1956), int32(_a_F_AtEOSubXact_HashTables_3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	goto L5
}
func F_AtEOSubXact_Namespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[0]))
	if l1 == v5 {
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[0])) = l2
			return
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[1])) = uint8(v10)
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[2])) = v13
			*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[0])) = v13
			*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[3])) = v13
			*(*uint8)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[4])) = uint8(v13)
			v25 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_Namespace[5]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v13
			return
		}
	} else {
		return
	}
}
func F_AtEOSubXact_PgStat(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int64
	_ = v237
	var v249 int32
	_ = v249
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[0]))
	if v12 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v15 < l1 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[0])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = l1 - int32(1)
	v28 = v20
	goto L7
L5:
	;
	goto L6
L6:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v129 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	if v34 != 0 {
		v28 = v34
		goto L7
	} else {
		goto L31
	}
L10:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v113
	F_pfree(m, v28)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L24
	} else {
		goto L30
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	if v83 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v77 = F_pgstat_get_xact_stack_level(m, v22)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v38 != v22 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+24)))
	if v40 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L10
L18:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v43 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v62 + v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v67 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+16)) = v72 + v73
	goto L17
L21:
	;
	v53 = v35
	goto L23
L22:
	;
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v44)
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v35)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v53 = v52
	goto L23
L23:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v28)+60))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v60
	goto L17
L24:
	;
	return
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+68)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v77)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v22
	goto L9
L26:
	;
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v33)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = v94 + v93
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v33)+48))
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+48)) = v97 + v98
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v33)+56))
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v28)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+56)) = v101 + v102
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v33)+96))
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v28)+8))
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+96)) = v105 + (v106 + v107)
	goto L10
L27:
	;
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	v93 = v86
	goto L26
L28:
	;
	goto L29
L29:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v28)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v28))) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v28)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+8)) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v28)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v28)+16)) = v91
	v93 = v87
	goto L26
L30:
	;
	goto L9
L31:
	;
	goto L8
L32:
	;
	F_pfree(m, v12)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L24
	} else {
		goto L59
	}
L33:
	;
	v133 = l1 - int32(1)
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[0]))
	if v135 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v159 == int32(0) {
		goto L32
	} else {
		goto L40
	}
L35:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v136 == v133 {
		v158 = v135
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[1]))
	v141 = F_MemoryContextAlloc(m, v139, int32(24))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L24
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v143 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+16)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v133
	v147 = v141 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v141)+12)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v141)+8)) = v147
	v150 = int32(_a_F_AtEOSubXact_PgStat_0)
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v143
	*(*int32)(unsafe.Add(mBase, uint32(v141)+4)) = v151
	*(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[0])) = v141
	v158 = v141
	goto L34
L40:
	;
	v163 = v12 + int32(8)
	if v159 == v163 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v166 = v158 + int32(8)
	v168 = v159
	v175 = v3
	goto L42
L42:
	;
	v179 = *(*int64)(unsafe.Add(mBase, uint32(v168-int32(12))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v168)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+4)) = v181
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	*(*int32)(unsafe.Add(mBase, uint32(v181))) = v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v185 - int32(1)
	if l0 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v229 <= int32(0) {
		goto L32
	} else {
		goto L57
	}
L44:
	;
	if v181 != v163 {
		v168 = v181
		v175 = v229
		goto L42
	} else {
		goto L56
	}
L45:
	;
	F_pfree(m, v192)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L24
	} else {
		goto L55
	}
L46:
	;
	v192 = v168 - int32(20)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168-int32(4)))))
	if v195 != int32(1) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	if v209 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v168-int32(16))))
	v202 = F_pgstat_drop_entry(m, v198, v201, v179)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L24
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v192)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	v229 = v175 + (v202 ^ int32(1))
	goto L44
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v158)+12)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v166
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168)+4)) = v166
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v168))) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v158)+8)) = v168
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v158)+16)) = v221 + int32(1)
	v229 = v175
	goto L44
L55:
	;
	v229 = v175
	goto L44
L56:
	;
	goto L43
L57:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_AtEOSubXact_PgStat[2]))
	v237 = base.AtomicRmwAdd64(m, v234, int32(16), int64(1))
	goto L58
L58:
	;
	goto L32
L59:
	;
	goto L1
}
