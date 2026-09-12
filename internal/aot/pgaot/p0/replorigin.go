package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_replorigin_by_oid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = F_SearchSysCache1(m, int32(58), l0)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v19 = F_text_to_cstring(m, v14+v15+int32(4))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v19
				F_ReleaseCatCache(m, v10)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return base.B2i32(v10 != int32(0))
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			m.G0 = v7 + int32(16)
			return base.B2i32(v10 != int32(0))
		}
	}
}
func F_replorigin_check_prerequisites(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	v2 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	if v2 != 0 {
		v5 = int32(*(*uint8)(unsafe.Add(mBase, _consts[2])))
		if v5 == int32(1) {
			v10 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+316))
			v13 = base.B2i32(v11 != int32(2))
			*(*uint8)(unsafe.Add(mBase, _consts[2])) = uint8(v13)
			v15 = v13
		} else {
			v15 = int32(0)
		}
		if v15 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return
			} else {
				F_errcode(m, int32(100663618))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errmsg(m, int32(13432), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_errfinish(m, int32(477368), int32(200), int32(151600))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
			return
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				F_errmsg(m, int32(534263), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(477368), int32(195), int32(151600))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
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
func F_replorigin_create(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	v7 = m.G0
	v9 = v7 - int32(144)
	m.G0 = v9
	if l0&int32(3) == int32(0) {
		v34 = l0
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L22
	} else {
		goto L50
	}
L2:
	;
	if base.Ui32(v67) < base.Ui32(int32(513)) {
		goto L19
	} else {
		goto L20
	}
L3:
	;
	v67 = v59 - l0
	goto L2
L4:
	;
	v38 = v34
	goto L13
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v67 = int32(0)
	goto L2
L7:
	;
	goto L8
L8:
	;
	v23 = l0
	goto L9
L9:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L4
	} else {
		goto L11
	}
L10:
	;
	v59 = v27
	goto L3
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L13
	} else {
		goto L15
	}
L14:
	;
	v53 = v38
	goto L16
L15:
	;
	goto L14
L16:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v59 = v53
	goto L3
L18:
	;
	goto L17
L19:
	;
	v70 = F_cstring_to_text(m, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L22
	} else {
		goto L45
	}
L22:
	;
	return int32(0)
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+72)) = int32(4)
	v78 = F_table_open(m, int32(6000), int32(7))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v81 = int32(1)
	goto L26
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v81
	v122 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)) = uint16(v122)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v78)+52))
	v129 = F_heap_form_tuple(m, v124, v9+int32(12), v9+int32(22))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L22
	} else {
		goto L39
	}
L26:
	;
	v88 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v88 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	F_sequence_close(m, v78, int32(7))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L22
	} else {
		goto L38
	}
L28:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L22
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	F_ScanKeyInit(m, v9+int32(24), int32(1), int32(3), int32(184), v81)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L22
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	v99 = int32(1)
	v105 = F_systable_beginscan(m, v78, int32(6001), v99, v9+int32(72), v99, v9+int32(24))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	v107 = F_systable_getnext(m, v105)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	F_systable_endscan(m, v105)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L22
	} else {
		goto L35
	}
L35:
	;
	if v107 == int32(0) {
		goto L25
	} else {
		goto L36
	}
L36:
	;
	v114 = v81 + int32(1)
	if v114 != int32(65535) {
		v81 = v114
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L27
L38:
	;
	goto L1
L39:
	;
	F_CatalogTupleInsert(m, v78, v129)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	F_sequence_close(m, v78, int32(7))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	if v129 == int32(0) {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_pfree(m, v129)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	m.G0 = v9 + int32(144)
	return v81
L45:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L22
	} else {
		goto L46
	}
L46:
	;
	F_errmsg(m, int32(314182), int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L22
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(512)
	F_errdetail(m, int32(560948), v9)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L22
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(477368), int32(277), int32(341134))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L22
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L22
	} else {
		goto L51
	}
L51:
	;
	F_errmsg(m, int32(523461), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L22
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(477368), int32(359), int32(341134))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L22
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_replorigin_identify(m *base.Module, l0 int32) int32 {
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	if l0 == int32(16) {
		v6 = int32(506885)
	} else {
		v6 = int32(0)
	}
	if l0 != 0 {
		v8 = v6
	} else {
		v8 = int32(501379)
	}
	return v8
}
func F_replorigin_session_reset(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v3 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	if v3 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_errcode(m, int32(325))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				F_errmsg(m, int32(432860), int32(0))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_errfinish(m, int32(477368), int32(1222), int32(99431))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
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
		v23 = *(*int32)(unsafe.Add(mBase, _consts[29]))
		v27 = F_LWLockAcquire(m, v23+int32(5120), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			v29 = int32(4359508)
			v30 = *(*int32)(unsafe.Add(mBase, _consts[519]))
			v31 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v30)+24)) = v31
			*(*int32)(unsafe.Add(mBase, _consts[519])) = v31
			v37 = *(*int32)(unsafe.Add(mBase, _consts[29]))
			F_LWLockRelease(m, v37+int32(5120))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				F_ConditionVariableBroadcast(m, v30+int32(28))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_replorigin_session_setup(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	v1 = l0
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v16 = int32(*(*uint8)(unsafe.Add(mBase, _consts[518])))
	if v16 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_on_shmem_exit(m, int32(1013), int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	if v27 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return
L5:
	;
	v24 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[518])) = uint8(v24)
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L53
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L48
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L44
	}
L9:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	v35 = F_LWLockAcquire(m, v31+int32(5120), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L40
	}
L12:
	;
	v38 = *(*int32)(unsafe.Add(mBase, _consts[520]))
	if v38 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	if l1 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L14:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	if base.B2i32(v85 == int32(0))&base.B2i32(v79 == int32(-1)) != 0 {
		goto L7
	} else {
		goto L31
	}
L15:
	;
	v79 = int32(-1)
	goto L14
L16:
	;
	goto L17
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v49 = v3
	v51 = int32(-1)
	goto L18
L18:
	;
	v57 = v43 + v49*int32(56)
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	if v58 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v79 = v70
	goto L14
L20:
	;
	v72 = v49 + int32(1)
	if v72 != v38 {
		v49 = v72
		v51 = v70
		goto L18
	} else {
		goto L30
	}
L21:
	;
	if v51 == int32(-1) {
		v70 = v49
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v1 != v58 {
		v70 = v51
		goto L20
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	if l1 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	if v67 != 0 {
		goto L8
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, _consts[519])) = v57
	v101 = v57
	goto L13
L29:
	;
	goto L28
L30:
	;
	goto L19
L31:
	;
	if v85 != 0 {
		v101 = v85
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _consts[521]))
	v96 = v93 + v79*int32(56)
	*(*int32)(unsafe.Add(mBase, _consts[519])) = v96
	*(*uint16)(unsafe.Add(mBase, uint32(v96))) = uint16(v1)
	v101 = v96
	goto L13
L33:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[29]))
	F_LWLockRelease(m, v117+int32(5120))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L38
	}
L34:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[522]))
	*(*int32)(unsafe.Add(mBase, uint32(v101)+24)) = v112
	goto L33
L35:
	;
	goto L36
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v101)+24))
	if v114 != l1 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v123 = *(*int32)(unsafe.Add(mBase, _consts[519]))
	F_ConditionVariableBroadcast(m, v123+int32(28))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	m.G0 = v13 + int32(48)
	return
L40:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(222509), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(477368), int32(1137), int32(222412))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57))))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v155
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v154
	F_errmsg(m, int32(461127), v13+int32(32))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(477368), int32(1167), int32(222412))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	F_errcode(m, int32(16581))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v1
	F_errmsg(m, int32(461403), v13)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(583023), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(477368), int32(1181), int32(222412))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v1
	F_errmsg_internal(m, int32(449514), v13+int32(16))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(477368), int32(1198), int32(222412))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
