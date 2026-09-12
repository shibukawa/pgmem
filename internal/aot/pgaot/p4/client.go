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
	*(*uint8)(unsafe.Add(mBase, _consts[1420])) = uint8(v8)
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
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
					v53 = *(*int32)(unsafe.Add(mBase, _consts[487]))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v54
					v57 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_consts[486])))
					*(*int32)(unsafe.Add(mBase, uint32(v5))) = v62
					F_errmsg(m, int32(439737), v5)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						F_errfinish(m, int32(490839), int32(308), int32(333954))
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
			v17 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
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
							v53 = *(*int32)(unsafe.Add(mBase, _consts[487]))
							v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
							*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v54
							v57 = *(*int32)(unsafe.Add(mBase, _consts[1421]))
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v57<<(uint(int32(3))%32))+uint32(_consts[486])))
							*(*int32)(unsafe.Add(mBase, uint32(v5))) = v62
							F_errmsg(m, int32(439737), v5)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(490839), int32(308), int32(333954))
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
					v23 = *(*int32)(unsafe.Add(mBase, _consts[487]))
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
								v31 = *(*int32)(unsafe.Add(mBase, _consts[36]))
								v33 = F_MemoryContextAlloc(m, v31, int32(28))
								mBase = m.M
								v34 = m.ExcPending
								if v34 != 0 {
									return
								} else {
									v36 = *(*int32)(unsafe.Add(mBase, _consts[36]))
									F_fmgr_info_cxt(m, v26, v33, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, _consts[1426])) = v33
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
	v4 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	v6 = *(*int32)(unsafe.Add(mBase, _consts[1199]))
	if v6 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
		return
	} else {
		if l0 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[162]))
			if v10 != 0 {
				*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
				return
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, _consts[26]))
				if v12 != 0 {
					*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
					return
				} else {
					v14 = *(*int32)(unsafe.Add(mBase, _consts[216]))
					if v14 == int32(2) {
						*(*int32)(unsafe.Add(mBase, _consts[216])) = int32(0)
					} else {
					}
					v21 = *(*int32)(unsafe.Add(mBase, _consts[0]))
					if v21 == int32(0) {
						*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
						return
					} else {
						F_ProcessInterrupts(m)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
							return
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _consts[498]))
			F_SetLatch(m, v29)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[140])) = v4
				return
			}
		}
	}
}
func F_SetClientEncoding(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
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
	var v106 int32
	_ = v106
	if base.Ui32(int32(41)) < base.Ui32(l0) {
		v106 = int32(-1)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v106
L2:
	;
	v11 = int32(0)
	v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1420])))
	if v13 == v11 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1421])) = l0
	v106 = v11
	goto L1
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
	v41 = *(*int32)(unsafe.Add(mBase, _consts[1422]))
	if v41 == int32(0) {
		v106 = int32(-1)
		goto L1
	} else {
		goto L11
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1423])) = l0<<(uint(int32(3))%32) + int32(1826112)
	v32 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[1424])) = v32
	*(*int32)(unsafe.Add(mBase, _consts[1425])) = v32
	return v32
L8:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _consts[487]))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v22 == l0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	v50 = int32(0)
	v51 = v41
	v53 = int32(0)
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
	return v102 - int32(1)
L14:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v50<<(uint(int32(2))%32))))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v63 != v22 {
		v94 = v50
		v95 = v51
		v96 = v53
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v102 = v53
	goto L16
L16:
	;
	goto L13
L17:
	;
	if v95 != 0 {
		v50 = v94 + int32(1)
		v51 = v95
		v53 = v96
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
		v96 = v53
		goto L17
	} else {
		goto L19
	}
L19:
	;
	if v53 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, _consts[1423])) = l0<<(uint(int32(3))%32) + int32(1826112)
	*(*int32)(unsafe.Add(mBase, _consts[1424])) = v62 + int32(8)
	*(*int32)(unsafe.Add(mBase, _consts[1425])) = v62 + int32(36)
	v91 = v50
	v92 = v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v79 = int32(4485036)
	v81 = *(*int32)(unsafe.Add(mBase, _consts[1422]))
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
	*(*int32)(unsafe.Add(mBase, _consts[1422])) = v82
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
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
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v253 int32
	_ = v253
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
	return v253
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
	v66 = int32(1827088)
	v67 = int32(1826448)
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
	v59 = v34 + int32(1)
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v60 != 0 {
		v34 = v59
		v35 = v60
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
		v253 = v4
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
	v122 = *(*int32)(unsafe.Add(mBase, _consts[172]))
	if int32(0) <= v122 {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v108<<(uint(int32(3))%32))+uint32(_consts[486])))
	v120 = v119
	goto L35
L34:
	;
	v120 = int32(738731)
	goto L35
L35:
	;
	goto L32
L36:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120))))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v183 == int32(0) {
		v202 = v182
		v203 = v183
		goto L60
	} else {
		goto L61
	}
L37:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, _consts[482])))
	if v126 != 0 {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v143 = F_PrepareClientEncoding(m, v108)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L43
	} else {
		goto L45
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, _consts[484])) = int32(322)
	goto L41
L41:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v132
	goto L42
L42:
	;
	v138 = F_format_elog_string(m, int32(597620), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	return int32(0)
L44:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v138
	v253 = int32(0)
	goto L1
L45:
	;
	if int32(0) <= v143 {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v148)+20))
	goto L48
L47:
	;
	*(*int32)(unsafe.Add(mBase, _consts[142])) = v175
	v253 = int32(0)
	goto L1
L48:
	;
	if v149 == int32(2) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _consts[484])) = int32(1088)
	goto L52
L50:
	;
	goto L51
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v168
	goto L56
L52:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	*(*int32)(unsafe.Add(mBase, _consts[141])) = v156
	goto L53
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _consts[487]))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v120
	v165 = F_format_elog_string(m, int32(629016), v9)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L43
	} else {
		goto L55
	}
L55:
	;
	v175 = v165
	goto L47
L56:
	;
	v173 = F_format_elog_string(m, int32(558644), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L43
	} else {
		goto L57
	}
L57:
	;
	v175 = v173
	goto L47
L58:
	;
	v244 = F_guc_malloc(m, int32(4))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L43
	} else {
		goto L80
	}
L59:
	;
	if v203-v202 == int32(0) {
		goto L58
	} else {
		goto L67
	}
L60:
	;
	goto L59
L61:
	;
	if v182 != v183 {
		v202 = v182
		v203 = v183
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v187 = v179
	v188 = v120
	goto L63
L63:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188)+1)))
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187)+1)))
	if v192 == int32(0) {
		v202 = v191
		v203 = v192
		goto L60
	} else {
		goto L65
	}
L64:
	;
	v202 = v191
	v203 = v192
	goto L60
L65:
	;
	v195 = int32(1)
	if v191 == v192 {
		v187 = v187 + v195
		v188 = v188 + v195
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v207 = int32(538740)
	v210 = int32(*(*uint8)(unsafe.Add(mBase, _consts[488])))
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	if v211 == int32(0) {
		v230 = v210
		v231 = v211
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v231-v230 == int32(0) {
		goto L58
	} else {
		goto L76
	}
L69:
	;
	goto L68
L70:
	;
	if v210 != v211 {
		v230 = v210
		v231 = v211
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v215 = v179
	v216 = v207
	goto L72
L72:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v220 == int32(0) {
		v230 = v219
		v231 = v220
		goto L69
	} else {
		goto L74
	}
L73:
	;
	v230 = v219
	v231 = v220
	goto L69
L74:
	;
	v223 = int32(1)
	if v219 == v220 {
		v215 = v215 + v223
		v216 = v216 + v223
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	F_bms_free(m, v179)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L43
	} else {
		goto L77
	}
L77:
	;
	v238 = F_guc_strdup(m, int32(15), v120)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L43
	} else {
		goto L78
	}
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v238
	if v238 != 0 {
		goto L58
	} else {
		goto L79
	}
L79:
	;
	v253 = int32(0)
	goto L1
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v244
	if v244 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v253 = int32(0)
	goto L1
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v244))) = v108
	v253 = int32(1)
	goto L1
}
