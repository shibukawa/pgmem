package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpenPipeStream(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_reserveAllocatedDesc(m)
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
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	if v15 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L61
	}
L6:
	;
	goto L15
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	v21 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[952]))
	if v21+(v23+v15) < v19 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	goto L9
L9:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _consts[953]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
	F_LruDelete(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L6
L11:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[951]))
	v44 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	v46 = *(*int32)(unsafe.Add(mBase, _consts[952]))
	if v42 <= v44+(v46+v38) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
L14:
	;
	m.G0 = v8 + int32(16)
	return v208
L15:
	;
	v61 = F_fflush(m, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v111
	v208 = v177
	goto L14
L17:
	;
	v64 = int32(0)
	v66 = m.G0
	v68 = v66 - int32(144)
	m.G0 = v68
	switch int32(2) {
	case 0, 2:
		v78 = v64
		goto L19
	default:
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0)
	v108 = F_pgl_popen(m, l0, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L31
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v78
	F_sigemptyset(m, v68+int32(8))
	mBase = m.M
	goto L22
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[656])) = v64
	v78 = int32(4730)
	goto L19
L22:
	;
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+136)) = int32(268435456)
	v90 = v68 + int32(4)
	goto L26
L24:
	;
	m.G0 = v68 + int32(144)
	goto L18
L26:
	;
	goto L27
L27:
	;
	if v90 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v101 = F___memcpy(m, int32(4721516), v90, int32(140))
	mBase = m.M
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L24
L31:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v113 = int32(-2)
	v115 = m.G0
	v117 = v115 - int32(144)
	m.G0 = v117
	switch int32(0) {
	case 0, 2:
		v127 = v113
		goto L33
	default:
		goto L34
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v111
	if v108 != 0 {
		goto L45
	} else {
		goto L46
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+4)) = v127
	F_sigemptyset(m, v117+int32(8))
	mBase = m.M
	goto L36
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[656])) = v113
	v127 = int32(4730)
	goto L33
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v117)+136)) = int32(268435456)
	v139 = v117 + int32(4)
	goto L40
L38:
	;
	m.G0 = v117 + int32(144)
	goto L32
L40:
	;
	goto L41
L41:
	;
	if v139 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v150 = F___memcpy(m, int32(4721516), v139, int32(140))
	mBase = m.M
	goto L44
L43:
	;
	goto L44
L44:
	;
	goto L38
L45:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _consts[954]))
	v159 = *(*int32)(unsafe.Add(mBase, _consts[952]))
	v162 = v157 + v159*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	goto L48
L46:
	;
	goto L47
L47:
	;
	v177 = int32(0)
	switch v111 - int32(33) {
	case 0, 8:
		goto L49
	default:
		v208 = v177
		goto L14
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v168
	v170 = int32(4470152)
	v172 = *(*int32)(unsafe.Add(mBase, _consts[952]))
	*(*int32)(unsafe.Add(mBase, _consts[952])) = v172 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v208 = v176
	goto L14
L49:
	;
	v182 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	if v182 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[950]))
	if int32(0) < v197 {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	F_errmsg(m, int32(12894), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	F_errfinish(m, int32(523542), int32(2793), int32(305643))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	goto L53
L57:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _consts[953]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	F_LruDelete(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	goto L16
L60:
	;
	goto L15
L61:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v222 = *(*int32)(unsafe.Add(mBase, _consts[955]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v222
	F_errmsg(m, int32(755105), v8)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(523542), int32(2765), int32(305643))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_OpenTemporaryFileInTablespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	v4 = m.G0
	v6 = v4 - int32(2112)
	m.G0 = v6
	if l0 != 0 {
		v13 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(l0-int32(1663)))
	} else {
		v13 = int32(0)
	}
	if v13 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = int32(246818)
		v24 = F_pg_snprintf(m, v6+int32(1088), int32(1024), int32(187361), v6+int32(48))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v43 = int32(4470172)
			v45 = *(*int32)(unsafe.Add(mBase, _consts[949]))
			*(*int32)(unsafe.Add(mBase, _consts[949])) = v45 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(246818)
			v52 = *(*int32)(unsafe.Add(mBase, _consts[712]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v52
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v45
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v6 + int32(1088)
			v64 = F_pg_snprintf(m, v6-int32(-64), int32(1024), int32(451868), v6+int32(16))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, _consts[947]))
				v71 = F_PathNameOpenFilePerm(m, v6-int32(-64), int32(578), v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if int32(0) < v71 {
						v106 = v71
						m.G0 = v6 + int32(2112)
						return v106
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, _consts[808]))
						v79 = F_mkdir(m, v6+int32(1088), v78)
						mBase = m.M
						v84 = *(*int32)(unsafe.Add(mBase, _consts[947]))
						v85 = F_PathNameOpenFilePerm(m, v6-int32(-64), int32(578), v84)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							if l1 == int32(0) {
								v106 = v85
								m.G0 = v6 + int32(2112)
								return v106
							} else {
								if int32(0) < v85 {
									v106 = v85
									m.G0 = v6 + int32(2112)
									return v106
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v6 - int32(-64)
										F_errmsg_internal(m, int32(311184), v6)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(523542), int32(1850), int32(438996))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
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
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(246818)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = int32(587648)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = int32(510355)
		v41 = F_pg_snprintf(m, v6+int32(1088), int32(1024), int32(187257), v6+int32(32))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = int32(4470172)
			v45 = *(*int32)(unsafe.Add(mBase, _consts[949]))
			*(*int32)(unsafe.Add(mBase, _consts[949])) = v45 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(246818)
			v52 = *(*int32)(unsafe.Add(mBase, _consts[712]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+24)) = v52
			*(*int32)(unsafe.Add(mBase, uint32(v6)+28)) = v45
			*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v6 + int32(1088)
			v64 = F_pg_snprintf(m, v6-int32(-64), int32(1024), int32(451868), v6+int32(16))
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, _consts[947]))
				v71 = F_PathNameOpenFilePerm(m, v6-int32(-64), int32(578), v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if int32(0) < v71 {
						v106 = v71
						m.G0 = v6 + int32(2112)
						return v106
					} else {
						v78 = *(*int32)(unsafe.Add(mBase, _consts[808]))
						v79 = F_mkdir(m, v6+int32(1088), v78)
						mBase = m.M
						v84 = *(*int32)(unsafe.Add(mBase, _consts[947]))
						v85 = F_PathNameOpenFilePerm(m, v6-int32(-64), int32(578), v84)
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							if l1 == int32(0) {
								v106 = v85
								m.G0 = v6 + int32(2112)
								return v106
							} else {
								if int32(0) < v85 {
									v106 = v85
									m.G0 = v6 + int32(2112)
									return v106
								} else {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v94 = m.ExcPending
									if v94 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6))) = v6 - int32(-64)
										F_errmsg_internal(m, int32(311184), v6)
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(523542), int32(1850), int32(438996))
											mBase = m.M
											v105 = m.ExcPending
											if v105 != 0 {
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
func F_oauth_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = int32(550144)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[513])))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v13 == int32(0) {
		v32 = v12
		v33 = v13
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L15
	} else {
		goto L40
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L15
	} else {
		goto L36
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L15
	} else {
		goto L33
	}
L4:
	;
	if v33-v32 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	goto L4
L6:
	;
	if v12 != v13 {
		v32 = v12
		v33 = v13
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v17 = l1
	v18 = v9
	goto L8
L8:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v22 == int32(0) {
		v32 = v21
		v33 = v22
		goto L5
	} else {
		goto L10
	}
L9:
	;
	v32 = v21
	v33 = v22
	goto L5
L10:
	;
	v25 = int32(1)
	if v21 == v22 {
		v17 = v17 + v25
		v18 = v18 + v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v38 = F_palloc0(m, int32(16))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L15
	} else {
		goto L29
	}
L15:
	;
	return int32(0)
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = l0
	v43 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38))) = v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+404))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+408))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+12)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+412))
	v56 = F_load_external_function(m, v52, int32(106815), v43, v43)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if v56 == int32(0) {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	v61 = m.T0[v56].(func(*base.Module) int32)(m)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[514])) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	if v64 != int32(539296288) {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	if v67 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v72 = F_palloc0(m, int32(8))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L15
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, _consts[515])) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(180003)
	v78 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v79 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	m.T0[v79].(func(*base.Module, int32))(m, v72)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L15
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v83 = F_palloc0(m, int32(12))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L15
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(792)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = v89
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88)+4)) = uint8(v91)
	*(*int32)(unsafe.Add(mBase, uint32(v88)+40)) = v83
	goto L28
L28:
	;
	m.G0 = v7 - int32(-64)
	return v38
L29:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L15
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(300820), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L15
	} else {
		goto L31
	}
L31:
	;
	F_errfinish(m, int32(521268), int32(109), int32(106793))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(106815)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(220415)
	F_errmsg(m, int32(195867), v7)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(521268), int32(761), int32(17690))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L15
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(220415)
	F_errmsg(m, int32(339841), v5+int32(-16))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L37
	}
L37:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(539296288)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v145
	F_errdetail(m, int32(684021), v5+int32(-32))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L15
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(521268), int32(776), int32(17690))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L15
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(527733)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(220415)
	F_errmsg(m, int32(334015), v5+int32(-48))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L15
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(521268), int32(785), int32(17690))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L15
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oidlt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(base.Ui32(v2) < base.Ui32(v3))
}
func F_oidtoi8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = F_Int64GetDatum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_oidvectorle(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v2 <= int32(0))
	}
}
func F_okcolors(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
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
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v300 int32
	_ = v300
	var v313 int32
	_ = v313
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(24)
	v17 = v13 + v14*v15
	if base.Ui32(v13) < base.Ui32(v17+v15) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = v13
	v31 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
	if v33&int32(1) != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if base.Ui32(v23) < base.Ui32(v17) {
		v23 = v23 + int32(24)
		v31 = v31 + int32(1)
		goto L4
	} else {
		goto L82
	}
L7:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v23)+8)))
	if v36 == int32(-1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(65535)
	v40 = v36 & v39
	v42 = v31 & v39
	if v40 == v42 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v210 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v210)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v212+v36*int32(24))+8)) = uint16(v210)
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v218 == int32(0) {
		goto L6
	} else {
		goto L51
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v45 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v46 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v23)+8)) = uint16(v46)
	v49 = v36 * int32(24)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v50)+8)) = uint16(v46)
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v66 != 0 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v42 == int32(0) {
		goto L6
	} else {
		goto L28
	}
L15:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+36))
	if v67 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	if v79 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v71 = int32(*(*int16)(unsafe.Add(mBase, uint32(v66)+4)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v66)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v70+v71*int32(24))+12)) = v75
	v79 = v75
	goto L18
L20:
	;
	goto L21
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v67)+32)) = v77
	v79 = v77
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v79)+36)) = v67
	goto L24
L23:
	;
	goto L24
L24:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+4)) = uint16(v36)
	*(*int64)(unsafe.Add(mBase, uint32(v66)+32)) = int64(0)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v87 = v84 + v49 + int32(12)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88)+36)) = v66
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v92 = v90
	goto L27
L26:
	;
	v92 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v66
	goto L13
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v100 = base.I32_extend16_s(v31)
	v103 = v99 + v100*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v100 == v106 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v111 = v100
	goto L33
L30:
	;
	goto L31
L31:
	;
	v203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+8)) = uint16(v203)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v208 = base.I32_div_s(v103-v205, int32(24))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v208)
	goto L6
L32:
	;
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	if base.Ui32(v133) < base.Ui32(v134) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108+v111*int32(24))+20)))
	if v124&int32(1) == int32(0) {
		v133 = v111
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v133 = int32(0)
	goto L32
L35:
	;
	v130 = v111 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v130
	if v130 != 0 {
		v111 = v130
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v139 = v134
	goto L40
L38:
	;
	v157 = v134
	goto L39
L39:
	;
	if v157 <= int32(0) {
		goto L6
	} else {
		goto L43
	}
L40:
	;
	v151 = int32(*(*int16)(unsafe.Add(mBase, uint32(v108+v139*int32(24))+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v151)
	if base.Ui32(v133) < base.Ui32(v151) {
		v139 = v151
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v157 = v151
	goto L39
L42:
	;
	goto L41
L43:
	;
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v108+v157*int32(24))+8)))
	if v171 <= int32(0) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v176 = v171
	v177 = v157
	goto L45
L45:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v188 = v176 & int32(65535)
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v186+v188*int32(24))+8)))
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v193) < base.Ui32(v188) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L6
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v186+base.I32_extend16_s(v177)*int32(24))+8)) = uint16(v192)
	v200 = v177
	goto L49
L48:
	;
	v200 = v176
	goto L49
L49:
	;
	if int32(0) < v192 {
		v176 = v192
		v177 = v200
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v230 = v218
	goto L52
L52:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v230)+12))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)+8))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	v237 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v237 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L6
L54:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v233)+8))
	if v240 <= v241 {
		goto L61
	} else {
		goto L62
	}
L57:
	;
	return
L58:
	;
	goto L56
L59:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v230)+32))
	if v313 != 0 {
		v230 = v313
		goto L52
	} else {
		goto L81
	}
L60:
	;
	F_createarc(m, l0, v235, v36, v234, v233)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L57
	} else {
		goto L80
	}
L61:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v234)+20))
	if v243 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v233)+16))
	if v265 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L64:
	;
	v248 = v243
	goto L65
L65:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	if v258 != v233 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L60
L67:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v248)+16))
	if v264 != 0 {
		v248 = v264
		goto L65
	} else {
		goto L71
	}
L68:
	;
	v260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+4)))
	if v260 != v40 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v248)))
	if v262 == v235 {
		goto L59
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	goto L66
L72:
	;
	v270 = v265
	goto L73
L73:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270)+8))
	if v280 != v234 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L60
L75:
	;
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v270)+24))
	if v286 != 0 {
		v270 = v286
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v270)+4)))
	if v282 != v40 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	if v284 == v235 {
		goto L59
	} else {
		goto L78
	}
L78:
	;
	goto L75
L79:
	;
	goto L74
L80:
	;
	goto L59
L81:
	;
	goto L53
L82:
	;
	goto L5
}
func F_or_arg_index_match_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v6 < v7 {
		return int32(-1)
	} else {
		v11 = int32(1)
		if v7 < v6 {
			v40 = v11
			return v40
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v13 < v14 {
				return int32(-1)
			} else {
				if v14 < v13 {
					v40 = v11
					return v40
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v19) < base.Ui32(v20) {
						return int32(-1)
					} else {
						if base.Ui32(v20) < base.Ui32(v19) {
							v40 = v11
							return v40
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if base.Ui32(v25) < base.Ui32(v26) {
								return int32(-1)
							} else {
								if base.Ui32(v26) < base.Ui32(v25) {
									v40 = v11
								} else {
									v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
									v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
									if v32 < v33 {
										v40 = int32(-1)
									} else {
										v40 = base.B2i32(v33 < v32)
									}
								}
								return v40
							}
						}
					}
				}
			}
		}
	}
}
func F_ordered_set_transition(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 == int32(1) {
		v7 = F_ordered_set_startup(m, l0, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v12 = v7
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v13 == int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				F_tuplesort_putdatum(m, v16, v17, int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v21 + int64(1)
					return v12
				}
			} else {
				return v12
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = v11
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			F_tuplesort_putdatum(m, v16, v17, int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v21 + int64(1)
				return v12
			}
		} else {
			return v12
		}
	}
}
func F_overlaps_time(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v78 int32
	_ = v78
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	v13 = int32(1)
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v20 == v13 {
		if v16&int32(1) == int32(0) {
			v36 = v13
			v37 = v19
			v38 = v19
			if v15&int32(1) != 0 {
				if v14&int32(1) != 0 {
					v89 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
					v99 = int32(0)
					return v99
				} else {
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
					v45 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
					if v45 < v44 {
						v89 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
						v99 = int32(0)
						return v99
					} else {
						v63 = v44
						v64 = v45
						v65 = int32(1)
						if v64 <= v63 {
							v78 = int32(1)
							if v36|v65 != v78 {
								v99 = v78
							} else {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
							}
							return v99
						} else {
							if v36 != 0 {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
								return v99
							} else {
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
								if v65&base.B2i32(v67 <= v64) != 0 {
									v89 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
									v99 = int32(0)
									return v99
								} else {
									return base.B2i32(v64 < v67)
								}
							}
						}
					}
				}
			} else {
				if v14&int32(1) != 0 {
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
					v51 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
					if v51 < v50 {
						v89 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
						v99 = int32(0)
						return v99
					} else {
						v63 = v50
						v64 = v51
						v65 = int32(1)
						if v64 <= v63 {
							v78 = int32(1)
							if v36|v65 != v78 {
								v99 = v78
							} else {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
							}
							return v99
						} else {
							if v36 != 0 {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
								return v99
							} else {
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
								if v65&base.B2i32(v67 <= v64) != 0 {
									v89 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
									v99 = int32(0)
									return v99
								} else {
									return base.B2i32(v64 < v67)
								}
							}
						}
					}
				} else {
					v54 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
					v55 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
					v56 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
					v57 = base.B2i32(v56 < v55)
					if v56 < v55 {
						v58 = v17
					} else {
						v58 = v18
					}
					v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
					if v59 < v54 {
						if v56 < v55 {
							v72 = v18
						} else {
							v72 = v17
						}
						v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
						if v36&base.B2i32(v73 <= v54) != 0 {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
							return v99
						} else {
							return base.B2i32(v54 < v73)
						}
					} else {
						v63 = v54
						v64 = v59
						v65 = int32(0)
						if v64 <= v63 {
							v78 = int32(1)
							if v36|v65 != v78 {
								v99 = v78
							} else {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
							}
							return v99
						} else {
							if v36 != 0 {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
								return v99
							} else {
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
								if v65&base.B2i32(v67 <= v64) != 0 {
									v89 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
									v99 = int32(0)
									return v99
								} else {
									return base.B2i32(v64 < v67)
								}
							}
						}
					}
				}
			}
		} else {
			v89 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
			v99 = int32(0)
			return v99
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v16&int32(1) != 0 {
			v36 = v13
			v37 = v19
			v38 = v27
		} else {
			v30 = *(*int64)(unsafe.Add(mBase, uint32(v27)))
			v31 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
			v32 = base.B2i32(v31 < v30)
			if v31 < v30 {
				v33 = v19
			} else {
				v33 = v27
			}
			if v31 < v30 {
				v34 = v27
			} else {
				v34 = v19
			}
			v36 = int32(0)
			v37 = v34
			v38 = v33
		}
		if v15&int32(1) != 0 {
			if v14&int32(1) != 0 {
				v89 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
				v99 = int32(0)
				return v99
			} else {
				v44 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
				v45 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				if v45 < v44 {
					v89 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
					v99 = int32(0)
					return v99
				} else {
					v63 = v44
					v64 = v45
					v65 = int32(1)
					if v64 <= v63 {
						v78 = int32(1)
						if v36|v65 != v78 {
							v99 = v78
						} else {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
						}
						return v99
					} else {
						if v36 != 0 {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
							return v99
						} else {
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							if v65&base.B2i32(v67 <= v64) != 0 {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
								return v99
							} else {
								return base.B2i32(v64 < v67)
							}
						}
					}
				}
			}
		} else {
			if v14&int32(1) != 0 {
				v50 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
				if v51 < v50 {
					v89 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
					v99 = int32(0)
					return v99
				} else {
					v63 = v50
					v64 = v51
					v65 = int32(1)
					if v64 <= v63 {
						v78 = int32(1)
						if v36|v65 != v78 {
							v99 = v78
						} else {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
						}
						return v99
					} else {
						if v36 != 0 {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
							return v99
						} else {
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							if v65&base.B2i32(v67 <= v64) != 0 {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
								return v99
							} else {
								return base.B2i32(v64 < v67)
							}
						}
					}
				}
			} else {
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v38)))
				v55 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
				v56 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
				v57 = base.B2i32(v56 < v55)
				if v56 < v55 {
					v58 = v17
				} else {
					v58 = v18
				}
				v59 = *(*int64)(unsafe.Add(mBase, uint32(v58)))
				if v59 < v54 {
					if v56 < v55 {
						v72 = v18
					} else {
						v72 = v17
					}
					v73 = *(*int64)(unsafe.Add(mBase, uint32(v72)))
					if v36&base.B2i32(v73 <= v54) != 0 {
						v89 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
						v99 = int32(0)
						return v99
					} else {
						return base.B2i32(v54 < v73)
					}
				} else {
					v63 = v54
					v64 = v59
					v65 = int32(0)
					if v64 <= v63 {
						v78 = int32(1)
						if v36|v65 != v78 {
							v99 = v78
						} else {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
						}
						return v99
					} else {
						if v36 != 0 {
							v89 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
							v99 = int32(0)
							return v99
						} else {
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v37)))
							if v65&base.B2i32(v67 <= v64) != 0 {
								v89 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v89)
								v99 = int32(0)
								return v99
							} else {
								return base.B2i32(v64 < v67)
							}
						}
					}
				}
			}
		}
	}
}
