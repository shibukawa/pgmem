package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitializeClientEncoding(m *base.Module) {
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[0])) = uint8(v8)
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[1]))
	v12 = F_PrepareClientEncoding(m, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 < int32(0) {
			F_errstart_cold(m, int32(22), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[2]))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[1]))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_c_F_InitializeClientEncoding[3])))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v62
					F_errmsg(m, int32(_a_F_InitializeClientEncoding_0), v5)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitializeClientEncoding_1), int32(308), int32(_a_F_InitializeClientEncoding_2))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
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
			v17 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[1]))
			v18 = F_SetClientEncoding(m, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				if v18 < int32(0) {
					F_errstart_cold(m, int32(22), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							v53 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[2]))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v54
							v57 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[1]))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_c_F_InitializeClientEncoding[3])))
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v62
							F_errmsg(m, int32(_a_F_InitializeClientEncoding_0), v5)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InitializeClientEncoding_1), int32(308), int32(_a_F_InitializeClientEncoding_2))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
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
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[2]))
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
					switch v24 {
					case 0, 6:
						m.G0 = v5 + int32(16)
						return
					default:
						v26 = F_FindDefaultConversionProc(m, int32(6), v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							if v26 == int32(0) {
								m.G0 = v5 + int32(16)
								return
							} else {
								v31 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[4]))
								v33 = F_MemoryContextAlloc(m, v31, int32(28))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[4]))
									F_fmgr_info_cxt(m, v26, v33, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _c_F_InitializeClientEncoding[5])) = v33
										m.G0 = v5 + int32(16)
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
func F_ProcessClientWriteInterrupt(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[1]))
	if v6 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
	return
L2:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[2]))
	if v10 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[3]))
	v30 = int32(0)
	v33 = base.AtomicRmwOr32(m, v30, int32(_a_F_ProcessClientWriteInterrupt_0), v30)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v34 != 0 {
		goto L15
	} else {
		goto L16
	}
L6:
	;
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[4]))
	if v12 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[5]))
	if v14 == int32(2) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[5])) = int32(0)
	goto L10
L9:
	;
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[6]))
	if v21 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
	return
L14:
	;
	goto L1
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(1)
	v37 = int32(0)
	v40 = base.AtomicRmwOr32(m, v37, int32(_a_F_ProcessClientWriteInterrupt_0), v37)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v41 == v37 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	if v44 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[7]))
	if v48 == v44 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v50 = m.G0
	v52 = v50 - int32(16)
	m.G0 = v52
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[8]))
	if v55 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v78 = F_pgmem_kill(m, v44, int32(23))
	mBase = m.M
	goto L15
L22:
	;
	m.G0 = v52 + int32(16)
	goto L14
L23:
	;
	v58 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v52)+15)) = uint8(v58)
	goto L24
L24:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[9]))
	v66 = F_write(m, v62, v52+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v66 {
		goto L22
	} else {
		goto L26
	}
L25:
	;
	goto L22
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0]))
	if v70 == int32(27) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
func F_SetClientEncoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
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
	var v68 int32
	_ = v68
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	v8 = int32(-1)
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
		v109 = v8
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v109
L2:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetClientEncoding[0])))
	if v15 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[1])) = l0
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	if l0 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[2]))
	if v44 == int32(0) {
		v109 = v8
		goto L1
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[3])) = l0<<(uint(int32(3))%32) + int32(_a_F_SetClientEncoding_0)
	v36 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[4])) = v36
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[5])) = v36
	return v36
L8:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[6]))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v26 == l0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v26 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v53 = int32(0)
	v54 = v44
	v57 = int32(0)
	goto L12
L12:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v53 < v59 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v109 = v105 - int32(1)
	goto L1
L14:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61+v53<<(uint(int32(2))%32))))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v66 != v26 {
		v97 = v53
		v98 = v54
		v99 = v57
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v105 = v57
	goto L16
L16:
	;
	goto L13
L17:
	;
	if v98 != 0 {
		v53 = v97 + int32(1)
		v54 = v98
		v57 = v99
		goto L12
	} else {
		goto L27
	}
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v68 != l0 {
		v97 = v53
		v98 = v54
		v99 = v57
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v57 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v97 = v94
	v98 = v95
	v99 = int32(1)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[3])) = l0<<(uint(int32(3))%32) + int32(_a_F_SetClientEncoding_0)
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[4])) = v65 + int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[5])) = v65 + int32(36)
	v94 = v53
	v95 = v54
	goto L20
L22:
	;
	goto L23
L23:
	;
	v82 = int32(_a_F_SetClientEncoding_1)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[2]))
	v85 = F_list_delete_nth_cell(m, v84, v53)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[2])) = v85
	F_pfree(m, v65)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v94 = v53 - int32(1)
	v95 = v85
	goto L20
L27:
	;
	v105 = v99
	goto L16
}
func F_check_client_encoding(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = int32(-1)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = m.G0
	v25 = v23 + int32(-64)
	m.G0 = v25
	if v15 == v4 {
		v102 = v12
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v269
L2:
	;
	if v102 == int32(7) {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	m.G0 = v25 - int32(-64)
	goto L2
L4:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v30 == int32(0) {
		v102 = v12
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v33 = F_strlen(m, v15)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v33) {
		v102 = v12
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v36 = v15
	v37 = v30
	v38 = v25
	goto L7
L7:
	;
	v46 = F_isalnum(m, v37&int32(255))
	mBase = m.M
	if v46 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59))) = uint8(v63)
	v67 = int32(*(*int8)(unsafe.Add(mBase, uint32(v25))))
	v68 = int32(_a_F_check_client_encoding_0)
	v69 = int32(_a_F_check_client_encoding_1)
	goto L16
L9:
	;
	if base.Ui32((v37-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v59 = v38
	goto L11
L11:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+1)))
	if v60 != 0 {
		v36 = v36 + int32(1)
		v37 = v60
		v38 = v59
		goto L7
	} else {
		goto L15
	}
L12:
	;
	v55 = v37 | int32(32)
	goto L14
L13:
	;
	v55 = v37
	goto L14
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v38))) = uint8(v55)
	v59 = v38 + int32(1)
	goto L11
L15:
	;
	goto L8
L16:
	;
	v81 = v69 + (v68-v69)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v83 = int32(*(*int8)(unsafe.Add(mBase, uint32(v82))))
	v84 = v67 - v83
	if v84 != 0 {
		v87 = v84
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v102 = v12
	goto L3
L18:
	;
	v91 = base.B2i32(v87 < int32(0))
	if v87 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v85 = F_strcmp(m, v25, v82)
	mBase = m.M
	if v85 != 0 {
		v87 = v85
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v102 = v86
	goto L3
L21:
	;
	v92 = v81 - int32(8)
	goto L23
L22:
	;
	v92 = v68
	goto L23
L23:
	;
	if v87 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v95 = v69
	goto L26
L25:
	;
	v95 = v81 + int32(8)
	goto L26
L26:
	;
	if base.Ui32(v95) <= base.Ui32(v92) {
		v68 = v92
		v69 = v95
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	v110 = v12
	goto L30
L29:
	;
	v110 = v102
	goto L30
L30:
	;
	if int32(41) < v102 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v113 = v12
	goto L33
L32:
	;
	v113 = v110
	goto L33
L33:
	;
	if v102 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v116 = v12
	goto L36
L35:
	;
	v116 = v113
	goto L36
L36:
	;
	if v116 < int32(0) {
		v269 = v4
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if base.B2i32(v116 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(v116)) != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[0]))
	if v132 < int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v130 = int32(_a_F_check_client_encoding_2)
	goto L41
L40:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v116<<(uint(int32(3))%32))+uint32(_c_F_check_client_encoding[1])))
	v130 = v129
	goto L41
L41:
	;
	goto L38
L42:
	;
	if int32(0) <= v132 {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_client_encoding[2])))
	if v136&int32(1) != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[3])) = int32(322)
	goto L45
L45:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[5])) = v144
	goto L46
L46:
	;
	v150 = F_format_elog_string(m, int32(_a_F_check_client_encoding_3), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	return int32(0)
L48:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[6])) = v150
	v269 = int32(0)
	goto L1
L49:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	if base.B2i32(v196 == int32(0))|base.B2i32(v196 != v199) != 0 {
		v217 = v196
		v218 = v199
		goto L66
	} else {
		goto L67
	}
L50:
	;
	v157 = F_PrepareClientEncoding(m, v116)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L47
	} else {
		goto L51
	}
L51:
	;
	if int32(0) <= v157 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[7]))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+20))
	goto L54
L53:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[6])) = v189
	v269 = int32(0)
	goto L1
L54:
	;
	if v163 == int32(2) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[3])) = int32(1088)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[5])) = v182
	goto L62
L58:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[5])) = v170
	goto L59
L59:
	;
	v174 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[8]))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v130
	v179 = F_format_elog_string(m, int32(_a_F_check_client_encoding_4), v9)
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L47
	} else {
		goto L61
	}
L61:
	;
	v189 = v179
	goto L53
L62:
	;
	v187 = F_format_elog_string(m, int32(_a_F_check_client_encoding_5), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L47
	} else {
		goto L63
	}
L63:
	;
	v189 = v187
	goto L53
L64:
	;
	v260 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L47
	} else {
		goto L84
	}
L65:
	;
	if v217-v218 == int32(0) {
		goto L64
	} else {
		goto L72
	}
L66:
	;
	goto L65
L67:
	;
	v202 = v193
	v203 = v130
	goto L68
L68:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203)+1)))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202)+1)))
	if v207 == int32(0) {
		v217 = v207
		v218 = v206
		goto L66
	} else {
		goto L70
	}
L69:
	;
	v217 = v207
	v218 = v206
	goto L66
L70:
	;
	v210 = int32(1)
	if v207 == v206 {
		v202 = v202 + v210
		v203 = v203 + v210
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v222 = int32(_a_F_check_client_encoding_6)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193))))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_client_encoding[9])))
	if base.B2i32(v225 == int32(0))|base.B2i32(v225 != v228) != 0 {
		v246 = v225
		v247 = v228
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v246-v247 == int32(0) {
		goto L64
	} else {
		goto L80
	}
L74:
	;
	goto L73
L75:
	;
	v231 = v193
	v232 = v222
	goto L76
L76:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+1)))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v231)+1)))
	if v236 == int32(0) {
		v246 = v236
		v247 = v235
		goto L74
	} else {
		goto L78
	}
L77:
	;
	v246 = v236
	v247 = v235
	goto L74
L78:
	;
	v239 = int32(1)
	if v236 == v235 {
		v231 = v231 + v239
		v232 = v232 + v239
		goto L76
	} else {
		goto L79
	}
L79:
	;
	goto L77
L80:
	;
	F_bms_free(m, v193)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L47
	} else {
		goto L81
	}
L81:
	;
	v254 = F_guc_strdup(m, int32(15), v130)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L47
	} else {
		goto L82
	}
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v254
	if v254 != 0 {
		goto L64
	} else {
		goto L83
	}
L83:
	;
	v269 = int32(0)
	goto L1
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v260
	if v260 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v269 = int32(0)
	goto L1
L86:
	;
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v116
	v269 = int32(1)
	goto L1
}
