package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CopyLoadRawBuf(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
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
	var v188 int64
	_ = v188
	var v190 int64
	_ = v190
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v9 = v7 - v8
	if v8 <= int32(0) {
		v163 = v9
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+328)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v163
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+308))
	if v167 == v168 {
		goto L50
	} else {
		goto L51
	}
L2:
	;
	if v9 <= int32(0) {
		v163 = v9
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v15 = v14 + v8
	if v14 == v15 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+332))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+328))
	v163 = v160 - v161
	goto L1
L5:
	;
	goto L4
L6:
	;
	v19 = v14 + v9
	if base.Ui32(v15-v19) <= base.Ui32(int32(0)-v9<<(uint(int32(1))%32)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v26 = F___memcpy(m, v14, v15, v9)
	mBase = m.M
	goto L4
L8:
	;
	goto L9
L9:
	;
	v29 = (v14 ^ v15) & int32(3)
	if base.Ui32(v14) < base.Ui32(v15) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v131 == int32(0) {
		goto L5
	} else {
		goto L46
	}
L11:
	;
	if base.Ui32(v109) <= base.Ui32(int32(3)) {
		v130 = v108
		v131 = v109
		v132 = v110
		goto L10
	} else {
		goto L42
	}
L12:
	;
	if v29 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v29 != 0 {
		v91 = v9
		goto L25
	} else {
		goto L26
	}
L15:
	;
	v130 = v15
	v131 = v9
	v132 = v14
	goto L10
L16:
	;
	goto L17
L17:
	;
	if v14&int32(3) == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v108 = v15
	v109 = v9
	v110 = v14
	goto L11
L19:
	;
	goto L20
L20:
	;
	v36 = v15
	v37 = v9
	v38 = v14
	goto L21
L21:
	;
	if v37 == int32(0) {
		goto L5
	} else {
		goto L23
	}
L22:
	;
	v108 = v45
	v109 = v47
	v110 = v49
	goto L11
L23:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36))))
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v42)
	v44 = int32(1)
	v45 = v36 + v44
	v47 = v37 - v44
	v49 = v38 + v44
	if v49&int32(3) != 0 {
		v36 = v45
		v37 = v47
		v38 = v49
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	if v91 == int32(0) {
		goto L5
	} else {
		goto L38
	}
L26:
	;
	if v19&int32(3) != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v56 = v9
	goto L30
L28:
	;
	v71 = v9
	goto L29
L29:
	;
	if base.Ui32(v71) <= base.Ui32(int32(3)) {
		v91 = v71
		goto L25
	} else {
		goto L34
	}
L30:
	;
	if v56 == int32(0) {
		goto L5
	} else {
		goto L32
	}
L31:
	;
	v71 = v62
	goto L29
L32:
	;
	v62 = v56 - int32(1)
	v63 = v14 + v62
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v62))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v65)
	if v63&int32(3) != 0 {
		v56 = v62
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v78 = v71
	goto L35
L35:
	;
	v82 = v78 - int32(4)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v15+v82)))
	*(*int32)(unsafe.Add(mBase, uint32(v14+v82))) = v85
	if base.Ui32(int32(3)) < base.Ui32(v82) {
		v78 = v82
		goto L35
	} else {
		goto L37
	}
L36:
	;
	v91 = v82
	goto L25
L37:
	;
	goto L36
L38:
	;
	v98 = v91
	goto L39
L39:
	;
	v102 = v98 - int32(1)
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v14+v102))) = uint8(v105)
	if v102 != 0 {
		v98 = v102
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L5
L41:
	;
	goto L40
L42:
	;
	v115 = v108
	v116 = v109
	v117 = v110
	goto L43
L43:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)))
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v119
	v121 = int32(4)
	v122 = v115 + v121
	v124 = v117 + v121
	v126 = v116 - v121
	if base.Ui32(int32(3)) < base.Ui32(v126) {
		v115 = v122
		v116 = v126
		v117 = v124
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v130 = v122
	v131 = v126
	v132 = v124
	goto L10
L45:
	;
	goto L44
L46:
	;
	v137 = v130
	v138 = v131
	v139 = v132
	goto L47
L47:
	;
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v137))))
	*(*uint8)(unsafe.Add(mBase, uint32(v139))) = uint8(v141)
	v143 = int32(1)
	v148 = v138 - v143
	if v148 != 0 {
		v137 = v137 + v143
		v138 = v148
		v139 = v139 + v143
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L5
L49:
	;
	goto L48
L50:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+312))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+312)) = int32(0)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+316))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+316)) = v173 - v170
	goto L52
L51:
	;
	goto L52
L52:
	;
	v180 = F_CopyGetData(m, l0, v163+v167, int32(65536)-v163)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return
L54:
	;
	v182 = v180 + v9
	v183 = *(*int32)(unsafe.Add(mBase, uint32(l0)+324))
	v185 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v182+v183))) = uint8(v185)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+332)) = v182
	v188 = *(*int64)(unsafe.Add(mBase, uint32(l0)+344))
	v190 = v188 + base.I64_extend_i32_s(v180)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+344)) = v190
	v195 = *(*int32)(unsafe.Add(mBase, _consts[55]))
	if v195 == v185 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v180 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L55
L57:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _consts[56])))
	if v199 != int32(1) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v202 = int32(4556756)
	v204 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	v205 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v204 + v205
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v208 + v205
	*(*int64)(unsafe.Add(mBase, uint32(v195+int32(0))+232)) = v190
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	*(*int32)(unsafe.Add(mBase, uint32(v195))) = v216 + v205
	v222 = *(*int32)(unsafe.Add(mBase, _consts[26]))
	*(*int32)(unsafe.Add(mBase, _consts[26])) = v222 - v205
	goto L56
L59:
	;
	v228 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+336)) = uint8(v228)
	goto L61
L60:
	;
	goto L61
L61:
	;
	return
}
func F_get_raw_page_fork_1_9(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
			v15 = F_text_to_cstring(m, v11)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = F_forkname_to_number(m, v15)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if base.Ui64(int64(4294967295)) <= base.Ui64(v14) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v24 = m.ExcPending
						if v24 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(240379), int32(0))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(525195), int32(113), int32(575386))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
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
						v42 = F_get_raw_page_internal(m, v6, v17, base.I32_wrap_i64(v14))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return int32(0)
						} else {
							return v42
						}
					}
				}
			}
		}
	}
}
