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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
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
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_c_F_InitializeClientEncoding[3])))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v60
					F_errmsg(m, int32(_a_F_InitializeClientEncoding_0), v5)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_InitializeClientEncoding_1), int32(308), int32(_a_F_InitializeClientEncoding_2))
						mBase = m.M
						v69 = m.ExcPending
						if v69 != 0 {
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
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_c_F_InitializeClientEncoding[3])))
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v60
							F_errmsg(m, int32(_a_F_InitializeClientEncoding_0), v5)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_InitializeClientEncoding_1), int32(308), int32(_a_F_InitializeClientEncoding_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
	var v31 int32
	_ = v31
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0]))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[1]))
	if v6 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
		return
	} else {
		if l0 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[2]))
			if v10 != 0 {
				*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[3]))
				if v12 != 0 {
					*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[4]))
					if v14 == int32(2) {
						*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[4])) = int32(0)
					} else {
					}
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[5]))
					if v21 == int32(0) {
						*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
						return
					} else {
						F_ProcessInterrupts(m)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
							return
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[6]))
			F_SetLatch(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_ProcessClientWriteInterrupt[0])) = v4
				return
			}
		}
	}
}
func F_SetClientEncoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	v2 = int32(0)
	if base.Ui32(int32(41)) < base.Ui32(l0) {
		v112 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v112
L2:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_SetClientEncoding[0])))
	if v12 == int32(0) {
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
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[2]))
	if v42 == int32(0) {
		v112 = int32(-1)
		goto L1
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[3])) = l0<<(uint(int32(3))%32) + int32(_a_F_SetClientEncoding_0)
	v33 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[4])) = v33
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[5])) = v33
	return v33
L8:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[6]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == l0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v23 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v50 = v2
	v51 = v42
	v54 = v2
	goto L12
L12:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v50 < v56 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v112 = v102 - int32(1)
	goto L1
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v50<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != v23 {
		v94 = v50
		v95 = v51
		v96 = v54
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v102 = v54
	goto L16
L16:
	;
	goto L13
L17:
	;
	if v95 != 0 {
		v50 = v94 + int32(1)
		v51 = v95
		v54 = v96
		goto L12
	} else {
		goto L27
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v65 != l0 {
		v94 = v50
		v95 = v51
		v96 = v54
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v54 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v94 = v91
	v95 = v92
	v96 = int32(1)
	goto L17
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[3])) = l0<<(uint(int32(3))%32) + int32(_a_F_SetClientEncoding_0)
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[4])) = v62 + int32(8)
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[5])) = v62 + int32(36)
	v91 = v50
	v92 = v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v79 = int32(_a_F_SetClientEncoding_1)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[2]))
	v82 = F_list_delete_nth_cell(m, v81, v50)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	return int32(0)
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_SetClientEncoding[2])) = v82
	F_pfree(m, v62)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v91 = v50 - int32(1)
	v92 = v82
	goto L20
L27:
	;
	v102 = v96
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
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = int32(-1)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = m.G0
	v23 = v21 + int32(-64)
	m.G0 = v23
	if v13 == v4 {
		v100 = v12
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return v256
L2:
	;
	if base.Ui32(int32(42)) <= base.Ui32(v100) {
		goto L28
	} else {
		goto L29
	}
L3:
	;
	m.G0 = v23 - int32(-64)
	goto L2
L4:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	if v28 == int32(0) {
		v100 = v12
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v31 = F_strlen(m, v13)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v31) {
		v100 = v12
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v34 = v13
	v35 = v28
	v36 = v23
	goto L7
L7:
	;
	v44 = F_isalnum(m, v35&int32(255))
	mBase = m.M
	if v44 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v57))) = uint8(v61)
	v65 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	v66 = int32(_a_F_check_client_encoding_0)
	v67 = int32(_a_F_check_client_encoding_1)
	goto L16
L9:
	;
	if base.Ui32((v35-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v57 = v36
	goto L11
L11:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+1)))
	if v58 != 0 {
		v34 = v34 + int32(1)
		v35 = v58
		v36 = v57
		goto L7
	} else {
		goto L15
	}
L12:
	;
	v53 = v35 | int32(32)
	goto L14
L13:
	;
	v53 = v35
	goto L14
L14:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v36))) = uint8(v53)
	v57 = v36 + int32(1)
	goto L11
L15:
	;
	goto L8
L16:
	;
	v79 = v67 + (v66-v67)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v81 = int32(*(*int8)(unsafe.Add(mBase, uint32(v80))))
	v82 = v65 - v81
	if v82 != 0 {
		v85 = v82
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v100 = v12
	goto L3
L18:
	;
	v89 = base.B2i32(v85 < int32(0))
	if v85 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v83 = F_strcmp(m, v23, v80)
	mBase = m.M
	if v83 != 0 {
		v85 = v83
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v100 = v84
	goto L3
L21:
	;
	v90 = v79 - int32(8)
	goto L23
L22:
	;
	v90 = v66
	goto L23
L23:
	;
	if v85 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v93 = v67
	goto L26
L25:
	;
	v93 = v79 + int32(8)
	goto L26
L26:
	;
	if base.Ui32(v93) <= base.Ui32(v90) {
		v66 = v90
		v67 = v93
		goto L16
	} else {
		goto L27
	}
L27:
	;
	goto L17
L28:
	;
	v108 = v12
	goto L30
L29:
	;
	v108 = v100
	goto L30
L30:
	;
	if v108 < int32(0) {
		v256 = v4
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if base.Ui32(v108) <= base.Ui32(int32(41)) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[0]))
	if v119 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v108<<(uint(int32(3))%32))+uint32(_c_F_check_client_encoding[1])))
	v117 = v115
	goto L35
L34:
	;
	v117 = int32(_a_F_check_client_encoding_2)
	goto L35
L35:
	;
	goto L32
L36:
	;
	if int32(0) <= v119 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_client_encoding[2])))
	if v123&int32(1) != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[3])) = int32(322)
	goto L39
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[5])) = v131
	goto L40
L40:
	;
	v137 = F_format_elog_string(m, int32(_a_F_check_client_encoding_3), int32(0))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	return int32(0)
L42:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[6])) = v137
	v256 = int32(0)
	goto L1
L43:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117))))
	if base.B2i32(v183 == int32(0))|base.B2i32(v183 != v186) != 0 {
		v204 = v183
		v205 = v186
		goto L60
	} else {
		goto L61
	}
L44:
	;
	v144 = F_PrepareClientEncoding(m, v108)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	if int32(0) <= v144 {
		goto L43
	} else {
		goto L46
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[7]))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	goto L48
L47:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[6])) = v176
	v256 = int32(0)
	goto L1
L48:
	;
	if v150 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[3])) = int32(1088)
	goto L52
L50:
	;
	goto L51
L51:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[5])) = v169
	goto L56
L52:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[4]))
	*(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[5])) = v157
	goto L53
L53:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_check_client_encoding[8]))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v162
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v117
	v166 = F_format_elog_string(m, int32(_a_F_check_client_encoding_4), v9)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L41
	} else {
		goto L55
	}
L55:
	;
	v176 = v166
	goto L47
L56:
	;
	v174 = F_format_elog_string(m, int32(_a_F_check_client_encoding_5), int32(0))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v176 = v174
	goto L47
L58:
	;
	v247 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L41
	} else {
		goto L78
	}
L59:
	;
	if v204-v205 == int32(0) {
		goto L58
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v189 = v180
	v190 = v117
	goto L62
L62:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190)+1)))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+1)))
	if v194 == int32(0) {
		v204 = v194
		v205 = v193
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v204 = v194
	v205 = v193
	goto L60
L64:
	;
	v197 = int32(1)
	if v194 == v193 {
		v189 = v189 + v197
		v190 = v190 + v197
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v209 = int32(_a_F_check_client_encoding_6)
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180))))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_check_client_encoding[9])))
	if base.B2i32(v212 == int32(0))|base.B2i32(v212 != v215) != 0 {
		v233 = v212
		v234 = v215
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v233-v234 == int32(0) {
		goto L58
	} else {
		goto L74
	}
L68:
	;
	goto L67
L69:
	;
	v218 = v180
	v219 = v209
	goto L70
L70:
	;
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v219)+1)))
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v218)+1)))
	if v223 == int32(0) {
		v233 = v223
		v234 = v222
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v233 = v223
	v234 = v222
	goto L68
L72:
	;
	v226 = int32(1)
	if v223 == v222 {
		v218 = v218 + v226
		v219 = v219 + v226
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	F_bms_free(m, v180)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L41
	} else {
		goto L75
	}
L75:
	;
	v241 = F_guc_strdup(m, int32(15), v117)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L41
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v241
	if v241 != 0 {
		goto L58
	} else {
		goto L77
	}
L77:
	;
	v256 = int32(0)
	goto L1
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v247
	if v247 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v256 = int32(0)
	goto L1
L80:
	;
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v247))) = v108
	v256 = int32(1)
	goto L1
}
