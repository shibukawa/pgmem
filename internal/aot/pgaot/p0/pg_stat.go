package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pg_stat_get_analyze_count(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+160))
			return v11
		}
	}
}
func F_pg_stat_get_archiver(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int64
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int64
	_ = v187
	var v190 int32
	_ = v190
	var v194 int64
	_ = v194
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int64
	_ = v208
	var v211 int32
	_ = v211
	var v215 int64
	_ = v215
	var v218 int32
	_ = v218
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v229 int32
	_ = v229
	v2 = int32(0)
	v3 = int64(0)
	v4 = m.G0
	v6 = v4 - int32(80)
	m.G0 = v6
	*(*int64)(unsafe.Add(mBase, uint32(v6)+64)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v6)+11)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v2
	v27 = F_CreateTemplateTupleDesc(m, int32(7))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_TupleDescInitEntry(m, v27, int32(1), int32(_a_F_pg_stat_get_archiver_0), int32(20), int32(-1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_TupleDescInitEntry(m, v27, int32(2), int32(_a_F_pg_stat_get_archiver_1), int32(25), int32(-1), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_TupleDescInitEntry(m, v27, int32(3), int32(_a_F_pg_stat_get_archiver_2), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_TupleDescInitEntry(m, v27, int32(4), int32(_a_F_pg_stat_get_archiver_3), int32(20), int32(-1), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	F_TupleDescInitEntry(m, v27, int32(5), int32(_a_F_pg_stat_get_archiver_4), int32(25), int32(-1), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	F_TupleDescInitEntry(m, v27, int32(6), int32(_a_F_pg_stat_get_archiver_5), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	F_TupleDescInitEntry(m, v27, int32(7), int32(_a_F_pg_stat_get_archiver_6), int32(1184), int32(-1), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v80 = int32(0)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v80 < v89 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v167 = F_BlessTupleDesc(m, v27)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L29
	}
L11:
	;
	v93 = v27 + int32(28)
	v100 = v80
	v101 = v89
	v103 = v80
	goto L15
L12:
	;
	v157 = v80
	v164 = v89
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v164
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v157
	goto L10
L14:
	;
	v157 = v151
	v164 = v130
	goto L13
L15:
	;
	v109 = v93 + v89<<(uint(int32(3))%32) + v100*int32(100)
	v112 = v93 + v100<<(uint(int32(3))%32)
	if v89 != v101 {
		v130 = v101
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v151 = v89
	goto L14
L17:
	;
	v131 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+2)))
	if v131 <= int32(0) {
		v151 = v100
		goto L14
	} else {
		goto L25
	}
L18:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+7)))
	if v114 != int32(118) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v130 = v100
	goto L17
L20:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+4)))
	if v117 != int32(1) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+6)))
	if v120&int32(6) != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v112)+2)))
	if v123 <= int32(0) {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+90)))
	if v126 != int32(118) {
		v130 = v89
		goto L17
	} else {
		goto L24
	}
L24:
	;
	goto L19
L25:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+90)))
	if v134 == int32(118) {
		v151 = v100
		goto L14
	} else {
		goto L26
	}
L26:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112)+5)))
	v143 = (v103 + v137 - int32(1)) & (int32(0) - v137)
	if int32(_a_F_pg_stat_get_archiver_7) < v143 {
		v151 = v100
		goto L14
	} else {
		goto L27
	}
L27:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v112))) = uint16(v143)
	v149 = v100 + int32(1)
	if v149 != v89 {
		v100 = v149
		v101 = v130
		v103 = v143 + v131
		goto L15
	} else {
		goto L28
	}
L28:
	;
	goto L16
L29:
	;
	F_pgstat_snapshot_fixed(m, int32(7))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v173 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+16)) = v173
	v176 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[1])))
	if v176 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v187 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[2]))
	if v187 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L32:
	;
	v179 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+9)) = uint8(v179)
	goto L31
L33:
	;
	goto L34
L34:
	;
	v182 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_8))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+24)) = base.I64_extend_i32_u(v182)
	goto L31
L36:
	;
	v194 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v6)+40)) = v194
	v197 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[4])))
	if v197 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L37:
	;
	v190 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+10)) = uint8(v190)
	goto L36
L38:
	;
	goto L39
L39:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+32)) = v187
	goto L36
L40:
	;
	v208 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[5]))
	if v208 == int64(0) {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	v200 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+12)) = uint8(v200)
	goto L40
L42:
	;
	goto L43
L43:
	;
	v203 = F_cstring_to_text(m, int32(_a_F_pg_stat_get_archiver_9))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+48)) = base.I64_extend_i32_u(v203)
	goto L40
L45:
	;
	v215 = *(*int64)(unsafe.Add(mBase, _c_F_pg_stat_get_archiver[6]))
	if v215 == int64(0) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	v211 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+13)) = uint8(v211)
	goto L45
L47:
	;
	goto L48
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+56)) = v208
	goto L45
L49:
	;
	v225 = F_heap_form_tuple(m, v27, v6+int32(16), v6+int32(8))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v218 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+14)) = uint8(v218)
	goto L49
L51:
	;
	goto L52
L52:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v6)+64)) = v215
	goto L49
L53:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v225)+16))
	v228 = F_HeapTupleHeaderGetDatum(m, v227)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	m.G0 = v6 + int32(80)
	return v228
}
func F_pg_stat_get_backend_activity(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		if v5 == int32(0) {
			v27 = int32(_a_F_pg_stat_get_backend_activity_0)
			v28 = F_pgstat_clip_activity(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				v30 = F_cstring_to_text(m, v28)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int64(0)
				} else {
					F_pfree(m, v28)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(v30)
					}
				}
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity[0]))
			v14 = F_has_privs_of_role(m, v12, int32(3375))
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int64(0)
			} else {
				if v14 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+216))
					v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
					if v24 != 0 {
						v25 = v22
					} else {
						v25 = int32(_a_F_pg_stat_get_backend_activity_1)
					}
					v27 = v25
					v28 = F_pgstat_clip_activity(m, v27)
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int64(0)
					} else {
						v30 = F_cstring_to_text(m, v28)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							F_pfree(m, v28)
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int64(0)
							} else {
								return base.I64_extend_i32_u(v30)
							}
						}
					}
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_activity[0]))
					v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
					v19 = F_has_privs_of_role(m, v17, v18)
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						if v19 != 0 {
							v22 = *(*int32)(unsafe.Add(mBase, uint32(v5)+216))
							v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
							if v24 != 0 {
								v25 = v22
							} else {
								v25 = int32(_a_F_pg_stat_get_backend_activity_1)
							}
							v27 = v25
						} else {
							v27 = int32(_a_F_pg_stat_get_backend_activity_2)
						}
						v28 = F_pgstat_clip_activity(m, v27)
						mBase = m.M
						v29 = m.ExcPending
						if v29 != 0 {
							return int64(0)
						} else {
							v30 = F_cstring_to_text(m, v28)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int64(0)
							} else {
								F_pfree(m, v28)
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int64(0)
								} else {
									return base.I64_extend_i32_u(v30)
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pg_stat_get_backend_client_port(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v186 int64
	_ = v186
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pgstat_get_beentry_by_proc_number(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(32)
	return v186
L2:
	;
	return int64(0)
L3:
	;
	if v15 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
	v186 = v9
	goto L1
L5:
	;
	goto L6
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_port[0]))
	v26 = F_has_privs_of_role(m, v24, int32(3375))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v36 = v15 + int32(56)
	v40 = (int32(-56) - v15) & int32(3)
	if v40 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	if v26 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_client_port[0]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v31 = F_has_privs_of_role(m, v29, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v31 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v33)
	v186 = v9
	goto L1
L12:
	;
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v176)
	v186 = v9
	goto L1
L13:
	;
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36))))
	switch v153 - int32(1) {
	case 0:
		v186 = int64(-1)
		goto L1
	case 1, 9:
		goto L42
	default:
		goto L43
	}
L14:
	;
	v53 = v40 + v36
	v57 = (v15 + int32(188)) & int32(-4)
	v59 = v57 - int32(28)
	if base.Ui32(v53) < base.Ui32(v59) {
		goto L21
	} else {
		goto L22
	}
L15:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	if v43 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	if v40 == int32(1) {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+57)))
	if v46 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	if v40 == int32(2) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+58)))
	if v49|base.B2i32(v40 != int32(3)) != 0 {
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L14
L21:
	;
	v62 = v40
	v63 = v53
	goto L24
L22:
	;
	v90 = v40
	goto L23
L23:
	;
	v98 = v90 + v36
	if base.Ui32(v98) < base.Ui32(v57) {
		goto L28
	} else {
		goto L29
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+28))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v63)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v70|(v71|(v72|(v73|(v74|(v75|(v76|v77)))))) != 0 {
		goto L13
	} else {
		goto L26
	}
L25:
	;
	v90 = v86
	goto L23
L26:
	;
	v86 = v62 + int32(32)
	v87 = v36 + v86
	if base.Ui32(v87) < base.Ui32(v59) {
		v62 = v86
		v63 = v87
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v101 = v90
	v102 = v98
	goto L31
L29:
	;
	v115 = v90
	goto L30
L30:
	;
	v123 = int32(132)
	if base.Ui32(v115) <= base.Ui32(v123) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	if v109 != 0 {
		goto L13
	} else {
		goto L33
	}
L32:
	;
	v115 = v111
	goto L30
L33:
	;
	v111 = v101 + int32(4)
	v112 = v36 + v111
	if base.Ui32(v112) < base.Ui32(v57) {
		v101 = v111
		v102 = v112
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v126 = v123
	goto L37
L36:
	;
	v126 = v115
	goto L37
L37:
	;
	v128 = v115
	goto L38
L38:
	;
	if v128 == v126 {
		goto L12
	} else {
		goto L40
	}
L39:
	;
	goto L13
L40:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+v36))))
	if v140 == int32(0) {
		v128 = v128 + int32(1)
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v159 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12))) = uint8(v159)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v15)+184))
	v166 = F_pg_getnameinfo_all(m, v36, v161, v159, v159, v12, int32(32), int32(3))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L2
	} else {
		goto L44
	}
L43:
	;
	v156 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v156)
	v186 = int64(0)
	goto L1
L44:
	;
	if v166 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v168 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v168)
	v186 = int64(0)
	goto L1
L46:
	;
	goto L47
L47:
	;
	v174 = F_DirectFunctionCall1Coll(m, int32(1142), int32(0), base.I64_extend_i32_u(v12))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v186 = v174
	goto L1
}
func F_pg_stat_get_backend_pid(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pgstat_get_beentry_by_proc_number(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = int64(*(*int32)(unsafe.Add(mBase, uint32(v4)+4)))
			return v14
		}
	}
}
func F_pg_stat_get_backend_wait_event_type(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pgstat_get_beentry_by_proc_number(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v92 = F_cstring_to_text(m, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L37
	}
L2:
	;
	return int64(0)
L3:
	;
	if v5 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v90 = int32(_a_F_pg_stat_get_backend_wait_event_type_0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event_type[0]))
	v15 = F_has_privs_of_role(m, v13, int32(3375))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L8
	}
L7:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v24 = F_BackendPidGetProc(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L2
	} else {
		goto L13
	}
L8:
	;
	if v15 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event_type[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+52))
	v20 = F_has_privs_of_role(m, v18, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v20 != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v90 = int32(_a_F_pg_stat_get_backend_wait_event_type_1)
	goto L1
L12:
	;
	v86 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v86)
	return int64(0)
L13:
	;
	if v24 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v29 = int32(0)
	if v28 == v29 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v67 = v24
	goto L16
L16:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+648))
	if v68 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L17:
	;
	if v64 == int32(0) {
		goto L12
	} else {
		goto L28
	}
L18:
	;
	v64 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_backend_wait_event_type[1]))
	v41 = v29
	goto L23
L21:
	;
	v64 = v58
	goto L17
L22:
	;
	v58 = v48 + int32(768)
	goto L21
L23:
	;
	v44 = v41 * int32(768)
	v45 = v37 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v46 == v28 {
		v58 = v45
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v64 = int32(0)
	goto L17
L25:
	;
	v48 = v37 + v44
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+780))
	if v49 == v28 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v52 = v41 + int32(2)
	if v52 != int32(38) {
		v41 = v52
		goto L23
	} else {
		goto L27
	}
L27:
	;
	goto L24
L28:
	;
	v67 = v64
	goto L16
L29:
	;
	if v83 != 0 {
		v90 = v83
		goto L1
	} else {
		goto L36
	}
L30:
	;
	v83 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v73 = v68 - int32(16777216)
	if base.Ui32(int32(184549375)) < base.Ui32(v73) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v83 = int32(_a_F_pg_stat_get_backend_wait_event_type_2)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v73)>>(uint(int32(22))%32))&int32(1020))+uint32(_c_F_pg_stat_get_backend_wait_event_type[2])))
	v83 = v81
	goto L29
L36:
	;
	goto L12
L37:
	;
	return base.I64_extend_i32_u(v92)
}
func F_pg_stat_get_checkpointer_num_timed(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	v2 = F_pgstat_fetch_stat_checkpointer(m)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		v6 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
		return v6
	}
}
func F_pg_stat_get_db_checksum_last_failure(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v27 int64
	_ = v27
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_get_db_checksum_last_failure[0]))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+268))
	if base.B2i32(v7 != int32(0)) == int32(0) {
		v24 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
		v27 = int64(0)
		return v27
	} else {
		v13 = F_pgstat_fetch_stat_dbentry(m, base.I32_wrap_i64(v4))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			if v13 == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v27 = int64(0)
			} else {
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v13)+160))
				if v19 != int64(0) {
					v27 = v19
				} else {
					v24 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
					v27 = int64(0)
				}
			}
			return v27
		}
	}
}
func F_pg_stat_get_db_numbackends(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pgstat_fetch_stat_numbackends(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v7 <= int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	goto L5
L5:
	;
	v15 = int32(1)
	v16 = int32(0)
	goto L6
L6:
	;
	v19 = F_pgstat_get_local_beentry_by_index(m, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	return base.I64_extend_i32_s(v23)
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	v23 = v16 + base.B2i32(v21 == v5)
	v25 = v15 + int32(1)
	if v25 <= v7 {
		v15 = v25
		v16 = v23
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_pg_stat_get_db_tuples_fetched(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+40))
			return v11
		}
	}
}
func F_pg_stat_get_db_tuples_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+56))
			return v11
		}
	}
}
func F_pg_stat_get_db_xact_rollback(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_dbentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+8))
			return v11
		}
	}
}
func F_pg_stat_get_function_calls(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pgstat_fetch_stat_funcentry(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int64(0)
		} else {
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
			return v14
		}
	}
}
func F_pg_stat_get_live_tuples(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+72))
			return v11
		}
	}
}
func F_pg_stat_get_mod_since_analyze(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+88))
			return v11
		}
	}
}
func F_pg_stat_get_numscans(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
			return v11
		}
	}
}
func F_pg_stat_get_tuples_inserted(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstat_fetch_stat_tabentry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+32))
			return v11
		}
	}
}
func F_pg_stat_get_xact_tuples_deleted(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+56))
			return v11
		}
	}
}
func F_pg_stat_get_xact_tuples_hot_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+64))
			return v11
		}
	}
}
func F_pg_stat_get_xact_tuples_newpage_updated(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int64
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_find_tabstat_entry(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		if v3 == int32(0) {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v3)+72))
			return v11
		}
	}
}
func F_pg_stat_reset_shared(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v43 int32
	_ = v43
	var v46 int64
	_ = v46
	var v48 int32
	_ = v48
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
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
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int64
	_ = v242
	var v244 int64
	_ = v244
	var v246 int32
	_ = v246
	var v247 int64
	_ = v247
	var v249 int64
	_ = v249
	var v251 int32
	_ = v251
	var v254 int64
	_ = v254
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v261 int32
	_ = v261
	var v264 int64
	_ = v264
	var v266 int32
	_ = v266
	var v269 int64
	_ = v269
	var v271 int32
	_ = v271
	var v274 int64
	_ = v274
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v362 int32
	_ = v362
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v7 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L6
	} else {
		goto L97
	}
L2:
	;
	m.G0 = v5 + int32(16)
	return int64(0)
L3:
	;
	F_pgstat_reset_of_kind(m, int32(7))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v69 = F_pg_detoast_datum_packed(m, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L6
	} else {
		goto L15
	}
L6:
	;
	return int64(0)
L7:
	;
	F_pgstat_reset_of_kind(m, int32(8))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_pgstat_reset_of_kind(m, int32(9))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	F_pgstat_reset_of_kind(m, int32(10))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	F_pgstat_reset_of_kind(m, int32(11))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	v27 = int32(_a_F_pg_stat_reset_shared_0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v29 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v31 = base.AtomicRmwXchg64(m, v28, int32(0), v29)
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v34 = int64(0)
	v36 = base.AtomicRmwXchg64(m, v33, int32(8), v34)
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v41 = base.AtomicRmwXchg64(m, v38, int32(16), v34)
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v46 = base.AtomicRmwXchg64(m, v43, int32(24), v34)
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v51 = base.AtomicRmwXchg64(m, v48, int32(32), v34)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v56 = base.AtomicRmwXchg64(m, v53, int32(40), v34)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v61 = base.AtomicRmwXchg64(m, v58, int32(48), v34)
	goto L12
L12:
	;
	F_pgstat_reset_of_kind(m, int32(12))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	F_pgstat_reset_of_kind(m, int32(13))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L2
L15:
	;
	v71 = F_text_to_cstring(m, v69)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L6
	} else {
		goto L16
	}
L16:
	;
	v73 = int32(_a_F_pg_stat_reset_shared_1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[1])))
	if base.B2i32(v76 == int32(0))|base.B2i32(v76 != v79) != 0 {
		v97 = v76
		v98 = v79
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v97-v98 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	v82 = v71
	v83 = v73
	goto L20
L20:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	if v87 == int32(0) {
		v97 = v87
		v98 = v86
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v97 = v87
	v98 = v86
	goto L18
L22:
	;
	v90 = int32(1)
	if v87 == v86 {
		v82 = v82 + v90
		v83 = v83 + v90
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	F_pgstat_reset_of_kind(m, int32(7))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L6
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v105 = int32(_a_F_pg_stat_reset_shared_2)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v111 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[2])))
	if base.B2i32(v108 == int32(0))|base.B2i32(v108 != v111) != 0 {
		v129 = v108
		v130 = v111
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L2
L28:
	;
	if v129-v130 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	goto L28
L30:
	;
	v114 = v71
	v115 = v105
	goto L31
L31:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v119 == int32(0) {
		v129 = v119
		v130 = v118
		goto L29
	} else {
		goto L33
	}
L32:
	;
	v129 = v119
	v130 = v118
	goto L29
L33:
	;
	v122 = int32(1)
	if v119 == v118 {
		v114 = v114 + v122
		v115 = v115 + v122
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	F_pgstat_reset_of_kind(m, int32(8))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v137 = int32(_a_F_pg_stat_reset_shared_3)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v143 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[3])))
	if base.B2i32(v140 == int32(0))|base.B2i32(v140 != v143) != 0 {
		v161 = v140
		v162 = v143
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L2
L39:
	;
	if v161-v162 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	v146 = v71
	v147 = v137
	goto L42
L42:
	;
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v147)+1)))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+1)))
	if v151 == int32(0) {
		v161 = v151
		v162 = v150
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v161 = v151
	v162 = v150
	goto L40
L44:
	;
	v154 = int32(1)
	if v151 == v150 {
		v146 = v146 + v154
		v147 = v147 + v154
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	F_pgstat_reset_of_kind(m, int32(9))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	if v169 != int32(105) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L2
L50:
	;
	v179 = int32(_a_F_pg_stat_reset_shared_4)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v185 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[4])))
	if base.B2i32(v182 == int32(0))|base.B2i32(v182 != v185) != 0 {
		v203 = v182
		v204 = v185
		goto L56
	} else {
		goto L57
	}
L51:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v172 != int32(111) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
	if v175 != 0 {
		goto L50
	} else {
		goto L53
	}
L53:
	;
	F_pgstat_reset_of_kind(m, int32(10))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L54
	}
L54:
	;
	goto L2
L55:
	;
	if v203-v204 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L56:
	;
	goto L55
L57:
	;
	v188 = v71
	v189 = v179
	goto L58
L58:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	if v193 == int32(0) {
		v203 = v193
		v204 = v192
		goto L56
	} else {
		goto L60
	}
L59:
	;
	v203 = v193
	v204 = v192
	goto L56
L60:
	;
	v196 = int32(1)
	if v193 == v192 {
		v188 = v188 + v196
		v189 = v189 + v196
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	F_pgstat_reset_of_kind(m, int32(11))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v211 = int32(_a_F_pg_stat_reset_shared_5)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[5])))
	if base.B2i32(v214 == int32(0))|base.B2i32(v214 != v217) != 0 {
		v235 = v214
		v236 = v217
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L2
L66:
	;
	if v235-v236 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L67:
	;
	goto L66
L68:
	;
	v220 = v71
	v221 = v211
	goto L69
L69:
	;
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+1)))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v220)+1)))
	if v225 == int32(0) {
		v235 = v225
		v236 = v224
		goto L67
	} else {
		goto L71
	}
L70:
	;
	v235 = v225
	v236 = v224
	goto L67
L71:
	;
	v228 = int32(1)
	if v225 == v224 {
		v220 = v220 + v228
		v221 = v221 + v228
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	v240 = int32(_a_F_pg_stat_reset_shared_0)
	v241 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v242 = F_GetCurrentTimestamp(m)
	mBase = m.M
	v244 = base.AtomicRmwXchg64(m, v241, int32(0), v242)
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v247 = int64(0)
	v249 = base.AtomicRmwXchg64(m, v246, int32(8), v247)
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v254 = base.AtomicRmwXchg64(m, v251, int32(16), v247)
	v256 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v259 = base.AtomicRmwXchg64(m, v256, int32(24), v247)
	v261 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v264 = base.AtomicRmwXchg64(m, v261, int32(32), v247)
	v266 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v269 = base.AtomicRmwXchg64(m, v266, int32(40), v247)
	v271 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[0]))
	v274 = base.AtomicRmwXchg64(m, v271, int32(48), v247)
	goto L76
L74:
	;
	goto L75
L75:
	;
	v275 = int32(_a_F_pg_stat_reset_shared_6)
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v281 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[6])))
	if base.B2i32(v278 == int32(0))|base.B2i32(v278 != v281) != 0 {
		v299 = v278
		v300 = v281
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L2
L77:
	;
	if v299-v300 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L78:
	;
	goto L77
L79:
	;
	v284 = v71
	v285 = v275
	goto L80
L80:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+1)))
	v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	if v289 == int32(0) {
		v299 = v289
		v300 = v288
		goto L78
	} else {
		goto L82
	}
L81:
	;
	v299 = v289
	v300 = v288
	goto L78
L82:
	;
	v292 = int32(1)
	if v289 == v288 {
		v284 = v284 + v292
		v285 = v285 + v292
		goto L80
	} else {
		goto L83
	}
L83:
	;
	goto L81
L84:
	;
	F_pgstat_reset_of_kind(m, int32(12))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L6
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v307 = int32(_a_F_pg_stat_reset_shared_7)
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	v313 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pg_stat_reset_shared[7])))
	if base.B2i32(v310 == int32(0))|base.B2i32(v310 != v313) != 0 {
		v331 = v310
		v332 = v313
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L2
L88:
	;
	if v331-v332 != 0 {
		goto L1
	} else {
		goto L95
	}
L89:
	;
	goto L88
L90:
	;
	v316 = v71
	v317 = v307
	goto L91
L91:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	if v321 == int32(0) {
		v331 = v321
		v332 = v320
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v331 = v321
	v332 = v320
	goto L89
L93:
	;
	v324 = int32(1)
	if v321 == v320 {
		v316 = v316 + v324
		v317 = v317 + v324
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	F_pgstat_reset_of_kind(m, int32(13))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L6
	} else {
		goto L96
	}
L96:
	;
	goto L2
L97:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L6
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v71
	F_errmsg(m, int32(_a_F_pg_stat_reset_shared_8), v5)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L6
	} else {
		goto L99
	}
L99:
	;
	F_errhint(m, int32(_a_F_pg_stat_reset_shared_9), int32(0))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L6
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(_a_F_pg_stat_reset_shared_10), int32(1994), int32(_a_F_pg_stat_reset_shared_11))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L6
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_pg_stat_reset_single_function_counters(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_reset_single_function_counters[0]))
	v5 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	F_pgstat_reset(m, int32(3), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_1_13(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(9), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_1_8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v9 int32
	_ = v9
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	F_pg_stat_statements_internal(m, l0, int32(4), base.B2i32(v3 != int64(0)))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		return int64(0)
	}
}
func F_pg_stat_statements_info(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	v2 = int32(0)
	v3 = int64(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v3
	*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v3
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v2)
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
	if v16 == v2 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_pg_stat_statements_info_0), int32(0))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_pg_stat_statements_info_1), int32(2046), int32(_a_F_pg_stat_statements_info_2))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[1]))
		if v20 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pg_stat_statements_info_0), int32(0))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pg_stat_statements_info_1), int32(2046), int32(_a_F_pg_stat_statements_info_2))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return int64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v26 = F_get_call_result_type(m, l0, int32(0), v7+int32(44))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				if v26 != int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return int64(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_pg_stat_statements_info_3), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_pg_stat_statements_info_1), int32(2050), int32(_a_F_pg_stat_statements_info_2))
							mBase = m.M
							v93 = m.ExcPending
							if v93 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
					v36 = base.AtomicRmwXchg32(m, v33, int32(140), int32(1))
					if v36 != 0 {
						F_s_lock(m, v33+int32(140), int32(_a_F_pg_stat_statements_info_4))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
							v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+160))
							v45 = *(*int64)(unsafe.Add(mBase, uint32(v43)+168))
							v46 = int32(0)
							atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v43)+140)), uint32(v46))
							*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v45
							*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v44
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
							v56 = F_heap_form_tuple(m, v51, v7+int32(16), v7+int32(14))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int64(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
								v59 = F_HeapTupleHeaderGetDatum(m, v58)
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int64(0)
								} else {
									m.G0 = v7 + int32(48)
									return v59
								}
							}
						}
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_pg_stat_statements_info[0]))
						v44 = *(*int64)(unsafe.Add(mBase, uint32(v43)+160))
						v45 = *(*int64)(unsafe.Add(mBase, uint32(v43)+168))
						v46 = int32(0)
						atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v43)+140)), uint32(v46))
						*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = v45
						*(*int64)(unsafe.Add(mBase, uint32(v7)+16)) = v44
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
						v56 = F_heap_form_tuple(m, v51, v7+int32(16), v7+int32(14))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+16))
							v59 = F_HeapTupleHeaderGetDatum(m, v58)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								m.G0 = v7 + int32(48)
								return v59
							}
						}
					}
				}
			}
		}
	}
}
