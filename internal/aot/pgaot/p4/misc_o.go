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
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v101 int32
	_ = v101
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
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
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[0]))
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
		goto L55
	}
L6:
	;
	goto L15
L7:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[1]))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[2]))
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[3]))
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
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[4]))
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
	v38 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[0]))
	if v38 <= int32(0) {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[1]))
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[2]))
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[3]))
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
	*(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[5])) = v111
	v208 = v177
	goto L14
L17:
	;
	v67 = m.G0
	v69 = v67 - int32(32)
	m.G0 = v69
	v71 = int32(2)
	switch v71 {
	case 0, 2:
		goto L19
	default:
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[5])) = int32(0)
	v108 = F_pgl_popen(m, l0, l1)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L28
	}
L19:
	;
	F_sigemptyset(m, v69+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(268435456)
	switch v71 {
	case 0:
		goto L24
	default:
		goto L22
	case 2:
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[6])) = int32(0)
	goto L19
L21:
	;
	goto L26
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = int32(_a_F_OpenPipeStream_0)
	goto L21
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = int32(0)
	goto L21
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+12)) = int32(-2)
	goto L21
L26:
	;
	goto L27
L27:
	;
	v101 = F___sigaction(m, int32(13), v69+int32(12), int32(0))
	mBase = m.M
	m.G0 = v69 + int32(32)
	goto L18
L28:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[5]))
	v114 = int32(0)
	v116 = m.G0
	v118 = v116 - int32(32)
	m.G0 = v118
	switch v114 {
	case 0, 2:
		goto L30
	default:
		goto L31
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[5])) = v111
	if v108 != 0 {
		goto L39
	} else {
		goto L40
	}
L30:
	;
	F_sigemptyset(m, v118+int32(16))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v118)+24)) = int32(268435456)
	switch v114 {
	case 0:
		goto L35
	default:
		goto L33
	case 2:
		goto L34
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[6])) = int32(-2)
	goto L30
L32:
	;
	goto L37
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+24)) = int32(268435460)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = int32(_a_F_OpenPipeStream_0)
	goto L32
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = int32(0)
	goto L32
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+12)) = int32(-2)
	goto L32
L37:
	;
	goto L38
L38:
	;
	v150 = F___sigaction(m, int32(13), v118+int32(12), int32(0))
	mBase = m.M
	m.G0 = v118 + int32(32)
	goto L29
L39:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[7]))
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[3]))
	v162 = v157 + v159*int32(12)
	*(*int32)(unsafe.Add(mBase, uint32(v162)+8)) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v162))) = int32(1)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[8]))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+8))
	goto L42
L40:
	;
	goto L41
L41:
	;
	v177 = int32(0)
	switch v111 - int32(33) {
	case 0, 8:
		goto L43
	default:
		v208 = v177
		goto L14
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v162)+4)) = v168
	v170 = int32(_a_F_OpenPipeStream_1)
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[3])) = v172 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v208 = v176
	goto L14
L43:
	;
	v182 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	if v182 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[0]))
	if int32(0) < v197 {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	F_errmsg(m, int32(_a_F_OpenPipeStream_2), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errfinish(m, int32(_a_F_OpenPipeStream_3), int32(2777), int32(_a_F_OpenPipeStream_4))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	goto L47
L51:
	;
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[4]))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v201)+16))
	F_LruDelete(m, v202)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	goto L16
L54:
	;
	goto L15
L55:
	;
	F_errcode(m, int32(197))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_OpenPipeStream[9]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v222
	F_errmsg(m, int32(_a_F_OpenPipeStream_5), v8)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_OpenPipeStream_3), int32(2749), int32(_a_F_OpenPipeStream_4))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_OpenTemporaryFileInTablespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v6 = m.G0
	v8 = v6 - int32(2112)
	m.G0 = v8
	if l0 != 0 {
		v15 = base.B2i32(base.Ui32(int32(2)) <= base.Ui32(l0-int32(1663)))
	} else {
		v15 = int32(0)
	}
	if v15 == int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(_a_F_OpenTemporaryFileInTablespace_0)
		v26 = F_pg_snprintf(m, v8+int32(1088), int32(1024), int32(_a_F_OpenTemporaryFileInTablespace_1), v8+int32(48))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v45 = int32(_a_F_OpenTemporaryFileInTablespace_2)
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[0])) = v47 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_OpenTemporaryFileInTablespace_0)
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v54
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v47
			v58 = v8 + int32(1088)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v58
			v61 = v8 - int32(-64)
			v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_OpenTemporaryFileInTablespace_3), v8+int32(16))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[2]))
				v71 = F_PathNameOpenFilePerm(m, v61, int32(578), v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if int32(0) < v71 {
						v101 = v71
						m.G0 = v8 + int32(2112)
						return v101
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[3]))
						v77 = F_mkdir(m, v58, v76)
						mBase = m.M
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[2]))
						v83 = F_PathNameOpenFilePerm(m, v61, int32(578), v82)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							if base.B2i32(l1 == int32(0))|base.B2i32(int32(0) < v83) != 0 {
								v101 = v83
								m.G0 = v8 + int32(2112)
								return v101
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
									F_errmsg_internal(m, int32(_a_F_OpenTemporaryFileInTablespace_4), v8)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_OpenTemporaryFileInTablespace_5), int32(1834), int32(_a_F_OpenTemporaryFileInTablespace_6))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(_a_F_OpenTemporaryFileInTablespace_0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_OpenTemporaryFileInTablespace_7)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(_a_F_OpenTemporaryFileInTablespace_8)
		v43 = F_pg_snprintf(m, v8+int32(1088), int32(1024), int32(_a_F_OpenTemporaryFileInTablespace_9), v8+int32(32))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			v45 = int32(_a_F_OpenTemporaryFileInTablespace_2)
			v47 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[0]))
			*(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[0])) = v47 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_OpenTemporaryFileInTablespace_0)
			v54 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[1]))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v54
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v47
			v58 = v8 + int32(1088)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v58
			v61 = v8 - int32(-64)
			v66 = F_pg_snprintf(m, v61, int32(1024), int32(_a_F_OpenTemporaryFileInTablespace_3), v8+int32(16))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				v70 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[2]))
				v71 = F_PathNameOpenFilePerm(m, v61, int32(578), v70)
				mBase = m.M
				v72 = m.ExcPending
				if v72 != 0 {
					return int32(0)
				} else {
					if int32(0) < v71 {
						v101 = v71
						m.G0 = v8 + int32(2112)
						return v101
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[3]))
						v77 = F_mkdir(m, v58, v76)
						mBase = m.M
						v82 = *(*int32)(unsafe.Add(mBase, _c_F_OpenTemporaryFileInTablespace[2]))
						v83 = F_PathNameOpenFilePerm(m, v61, int32(578), v82)
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							if base.B2i32(l1 == int32(0))|base.B2i32(int32(0) < v83) != 0 {
								v101 = v83
								m.G0 = v8 + int32(2112)
								return v101
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v91 = m.ExcPending
								if v91 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8))) = v61
									F_errmsg_internal(m, int32(_a_F_OpenTemporaryFileInTablespace_4), v8)
									mBase = m.M
									v95 = m.ExcPending
									if v95 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_OpenTemporaryFileInTablespace_5), int32(1834), int32(_a_F_OpenTemporaryFileInTablespace_6))
										mBase = m.M
										v100 = m.ExcPending
										if v100 != 0 {
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
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = int32(_a_F_oauth_init_0)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_oauth_init[0])))
	if base.B2i32(v12 == int32(0))|base.B2i32(v12 != v15) != 0 {
		v33 = v12
		v34 = v15
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L14
	} else {
		goto L39
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L14
	} else {
		goto L35
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L32
	}
L4:
	;
	if v33-v34 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	goto L4
L6:
	;
	v18 = l1
	v19 = v9
	goto L7
L7:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
	if v23 == int32(0) {
		v33 = v23
		v34 = v22
		goto L5
	} else {
		goto L9
	}
L8:
	;
	v33 = v23
	v34 = v22
	goto L5
L9:
	;
	v26 = int32(1)
	if v23 == v22 {
		v18 = v18 + v26
		v19 = v19 + v26
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_init[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_oauth_init[2])) = v40
	v43 = F_palloc0(m, int32(16))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L14
	} else {
		goto L28
	}
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = l0
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v48
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+372))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+8)) = v51
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+376))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+12)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+380))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+380))
	v61 = F_load_external_function(m, v57, int32(_a_F_oauth_init_1), v48, v48)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	if v61 == int32(0) {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v66 = m.T0[v61].(func(*base.Module) int32)(m)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_oauth_init[3])) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v69 != int32(539296288) {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	if v72 == int32(0) {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v77 = F_palloc0(m, int32(8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_oauth_init[4])) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v77))) = int32(_a_F_oauth_init_2)
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_init[3]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v84 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	m.T0[v84].(func(*base.Module, int32))(m, v77)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L14
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v88 = F_palloc0(m, int32(12))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L14
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v88))) = int32(838)
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_init[1]))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+8)) = v94
	v96 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v93)+4)) = uint8(v96)
	*(*int32)(unsafe.Add(mBase, uint32(v93)+40)) = v88
	goto L27
L27:
	;
	m.G0 = v7 - int32(-64)
	return v43
L28:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L14
	} else {
		goto L29
	}
L29:
	;
	F_errmsg(m, int32(_a_F_oauth_init_3), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_oauth_init_4), int32(116), int32(_a_F_oauth_init_5))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L14
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = int32(_a_F_oauth_init_1)
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v57
	F_errmsg(m, int32(_a_F_oauth_init_6), v7)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L14
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_oauth_init_4), int32(795), int32(_a_F_oauth_init_7))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L14
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v57
	F_errmsg(m, int32(_a_F_oauth_init_8), v5+int32(-16))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L14
	} else {
		goto L36
	}
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_oauth_init[3]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(539296288)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v146
	v153 = F_errdetail(m, int32(_a_F_oauth_init_9), v5+int32(-32))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_oauth_init_4), int32(810), int32(_a_F_oauth_init_7))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L14
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_oauth_init_10)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v57
	F_errmsg(m, int32(_a_F_oauth_init_11), v5+int32(-48))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_oauth_init_4), int32(819), int32(_a_F_oauth_init_7))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_oid8_decrement(m *base.Module, l0 int32, l1 int64, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	v4 = int64(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(base.B2i32(l1 == v4))
	return l1 - base.I64_extend_i32_u(base.B2i32(l1 != v4))
}
func F_oidlt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(base.Ui32(v2) < base.Ui32(v3)))
}
func F_oidtoi8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	v2 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+24)))
	return v2
}
func F_oidvectorle(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(base.I32_wrap_i64(v2) <= int32(0)))
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
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v299 int32
	_ = v299
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v15 = int32(24)
	v19 = v13 + v14*v15 + v15
	if base.Ui32(v13) < base.Ui32(v19) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v27 = v13
	v30 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+20)))
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
	v328 = v27 + int32(24)
	if base.Ui32(v328) < base.Ui32(v19) {
		v27 = v328
		v30 = v30 + int32(1)
		goto L4
	} else {
		goto L82
	}
L7:
	;
	v36 = int32(*(*int16)(unsafe.Add(mBase, uint32(v27)+8)))
	if v36 == int32(-1) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v39 = int32(_a_F_okcolors_0)
	v40 = v36 & v39
	v42 = v30 & v39
	if v40 == v42 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v209 = int32(_a_F_okcolors_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)) = uint16(v209)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v211+v36*int32(24))+8)) = uint16(v209)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v217 == int32(0) {
		goto L6
	} else {
		goto L51
	}
L11:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v45 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v46 = int32(_a_F_okcolors_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+8)) = uint16(v46)
	v49 = v36 * int32(24)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v49+v50)+8)) = uint16(v46)
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
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
	v85 = v84 + v49
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	if v86 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+36)) = v66
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v85)+12))
	v90 = v88
	goto L27
L26:
	;
	v90 = int32(0)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v66)+32)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v66
	goto L13
L28:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v98 = base.I32_extend16_s(v30)
	v101 = v97 + v98*int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+20)) = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v98 == v104 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v109 = v98
	goto L33
L30:
	;
	goto L31
L31:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+8)) = uint16(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v207 = base.I32_div_s(v101-v204, int32(24))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v207)
	goto L6
L32:
	;
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+16)))
	if base.Ui32(v131) < base.Ui32(v132) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106+v109*int32(24))+20)))
	if v122&int32(1) == int32(0) {
		v131 = v109
		goto L32
	} else {
		goto L35
	}
L34:
	;
	v131 = int32(0)
	goto L32
L35:
	;
	v128 = v109 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v128
	if v128 != 0 {
		v109 = v128
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v137 = v132
	goto L40
L38:
	;
	v155 = v132
	goto L39
L39:
	;
	if v155 <= int32(0) {
		goto L6
	} else {
		goto L43
	}
L40:
	;
	v149 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106+v137*int32(24))+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v149)
	if base.Ui32(v131) < base.Ui32(v149) {
		v137 = v149
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v155 = v149
	goto L39
L42:
	;
	goto L41
L43:
	;
	v169 = int32(*(*int16)(unsafe.Add(mBase, uint32(v106+v155*int32(24))+8)))
	if v169 <= int32(0) {
		goto L6
	} else {
		goto L44
	}
L44:
	;
	v174 = v169
	v175 = v155
	goto L45
L45:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v186 = v174 & int32(_a_F_okcolors_0)
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v184+v186*int32(24))+8)))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if base.Ui32(v191) < base.Ui32(v186) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L6
L47:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v184+v175&int32(_a_F_okcolors_0)*int32(24))+8)) = uint16(v190)
	v199 = v175
	goto L49
L48:
	;
	v199 = v174
	goto L49
L49:
	;
	if int32(0) < v190 {
		v174 = v190
		v175 = v199
		goto L45
	} else {
		goto L50
	}
L50:
	;
	goto L46
L51:
	;
	v223 = v217
	goto L52
L52:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v223)+12))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v223)+8))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v223)))
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_okcolors[0]))
	if v236 != 0 {
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
	v238 = m.ExcPending
	if v238 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v232)+8))
	if v239 <= v240 {
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
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v223)+32))
	if v312 != 0 {
		v223 = v312
		goto L52
	} else {
		goto L81
	}
L60:
	;
	F_createarc(m, l0, v234, v36, v233, v232)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L57
	} else {
		goto L80
	}
L61:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v233)+20))
	if v242 == int32(0) {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
	if v264 == int32(0) {
		goto L60
	} else {
		goto L72
	}
L64:
	;
	v247 = v242
	goto L65
L65:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	if v257 != v232 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	goto L60
L67:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v247)+16))
	if v263 != 0 {
		v247 = v263
		goto L65
	} else {
		goto L71
	}
L68:
	;
	v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	if v259 != v40 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v247)))
	if v261 == v234 {
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
	v269 = v264
	goto L73
L73:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v269)+8))
	if v279 != v233 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	goto L60
L75:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v269)+24))
	if v285 != 0 {
		v269 = v285
		goto L73
	} else {
		goto L79
	}
L76:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v269)+4)))
	if v281 != v40 {
		goto L75
	} else {
		goto L77
	}
L77:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v269)))
	if v283 == v234 {
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
func F_ordered_set_transition(m *base.Module, l0 int32) int64 {
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
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v3 == int32(1) {
		v7 = F_ordered_set_startup(m, l0, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v12 = v7
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
			if v13 == int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				F_tuplesort_putdatum(m, v16, v17, int32(0))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int64(0)
				} else {
					v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v21 + int64(1)
					return base.I64_extend_i32_u(v12)
				}
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = v11
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
		if v13 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
			F_tuplesort_putdatum(m, v16, v17, int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+16))
				*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v21 + int64(1)
				return base.I64_extend_i32_u(v12)
			}
		} else {
			return base.I64_extend_i32_u(v12)
		}
	}
}
func F_overlaps_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
	var v25 int64
	_ = v25
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v47 int64
	_ = v47
	var v51 int64
	_ = v51
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v65 int64
	_ = v65
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	v11 = int32(1)
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v18 == v11 {
		if v14&int32(1) == int32(0) {
			v33 = v17
			v34 = v17
			v36 = v11
			if v13&int32(1) != 0 {
				if v12&int32(1) == int32(0) {
					v47 = v15
					if v47 < v34 {
						v81 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
						v90 = int32(0)
						return base.I64_extend_i32_u(v90)
					} else {
						v54 = v47
						v56 = int32(1)
						if v54 <= v34 {
							v71 = int32(1)
							if v56|v36 != v71 {
								v90 = v71
							} else {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
								v90 = int32(0)
							}
							return base.I64_extend_i32_u(v90)
						} else {
							if v36|v56&base.B2i32(v33 <= v54) != 0 {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
								v90 = int32(0)
								return base.I64_extend_i32_u(v90)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v54 < v33))
							}
						}
					}
				} else {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
					v90 = int32(0)
					return base.I64_extend_i32_u(v90)
				}
			} else {
				if v12&int32(1) == int32(0) {
					if v16 < v15 {
						v51 = v16
					} else {
						v51 = v15
					}
					if v51 < v34 {
						if v15 < v16 {
							v65 = v16
						} else {
							v65 = v15
						}
						if v36&base.B2i32(v65 <= v34) != 0 {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
							return base.I64_extend_i32_u(v90)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v34 < v65))
						}
					} else {
						v54 = v51
						v56 = int32(0)
						if v54 <= v34 {
							v71 = int32(1)
							if v56|v36 != v71 {
								v90 = v71
							} else {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
								v90 = int32(0)
							}
							return base.I64_extend_i32_u(v90)
						} else {
							if v36|v56&base.B2i32(v33 <= v54) != 0 {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
								v90 = int32(0)
								return base.I64_extend_i32_u(v90)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v54 < v33))
							}
						}
					}
				} else {
					v47 = v16
					if v47 < v34 {
						v81 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
						v90 = int32(0)
						return base.I64_extend_i32_u(v90)
					} else {
						v54 = v47
						v56 = int32(1)
						if v54 <= v34 {
							v71 = int32(1)
							if v56|v36 != v71 {
								v90 = v71
							} else {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
								v90 = int32(0)
							}
							return base.I64_extend_i32_u(v90)
						} else {
							if v36|v56&base.B2i32(v33 <= v54) != 0 {
								v81 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
								v90 = int32(0)
								return base.I64_extend_i32_u(v90)
							} else {
								return base.I64_extend_i32_u(base.B2i32(v54 < v33))
							}
						}
					}
				}
			}
		} else {
			v81 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
			v90 = int32(0)
			return base.I64_extend_i32_u(v90)
		}
	} else {
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		if v14&int32(1) != 0 {
			v33 = v17
			v34 = v25
			v36 = v11
		} else {
			if v25 < v17 {
				v29 = v25
			} else {
				v29 = v17
			}
			if v17 < v25 {
				v31 = v25
			} else {
				v31 = v17
			}
			v33 = v31
			v34 = v29
			v36 = int32(0)
		}
		if v13&int32(1) != 0 {
			if v12&int32(1) == int32(0) {
				v47 = v15
				if v47 < v34 {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
					v90 = int32(0)
					return base.I64_extend_i32_u(v90)
				} else {
					v54 = v47
					v56 = int32(1)
					if v54 <= v34 {
						v71 = int32(1)
						if v56|v36 != v71 {
							v90 = v71
						} else {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
						}
						return base.I64_extend_i32_u(v90)
					} else {
						if v36|v56&base.B2i32(v33 <= v54) != 0 {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
							return base.I64_extend_i32_u(v90)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v54 < v33))
						}
					}
				}
			} else {
				v81 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
				v90 = int32(0)
				return base.I64_extend_i32_u(v90)
			}
		} else {
			if v12&int32(1) == int32(0) {
				if v16 < v15 {
					v51 = v16
				} else {
					v51 = v15
				}
				if v51 < v34 {
					if v15 < v16 {
						v65 = v16
					} else {
						v65 = v15
					}
					if v36&base.B2i32(v65 <= v34) != 0 {
						v81 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
						v90 = int32(0)
						return base.I64_extend_i32_u(v90)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v34 < v65))
					}
				} else {
					v54 = v51
					v56 = int32(0)
					if v54 <= v34 {
						v71 = int32(1)
						if v56|v36 != v71 {
							v90 = v71
						} else {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
						}
						return base.I64_extend_i32_u(v90)
					} else {
						if v36|v56&base.B2i32(v33 <= v54) != 0 {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
							return base.I64_extend_i32_u(v90)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v54 < v33))
						}
					}
				}
			} else {
				v47 = v16
				if v47 < v34 {
					v81 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
					v90 = int32(0)
					return base.I64_extend_i32_u(v90)
				} else {
					v54 = v47
					v56 = int32(1)
					if v54 <= v34 {
						v71 = int32(1)
						if v56|v36 != v71 {
							v90 = v71
						} else {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
						}
						return base.I64_extend_i32_u(v90)
					} else {
						if v36|v56&base.B2i32(v33 <= v54) != 0 {
							v81 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v81)
							v90 = int32(0)
							return base.I64_extend_i32_u(v90)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v54 < v33))
						}
					}
				}
			}
		}
	}
}
