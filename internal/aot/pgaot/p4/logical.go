package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckLogicalDecodingRequirements(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	F_CheckSlotRequirements(m)
	mBase = m.M
	v2 = m.ExcPending
	if v2 != 0 {
		return
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalDecodingRequirements[0]))
		if int32(1) < v4 {
			v8 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalDecodingRequirements[1]))
			if v8 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return
				} else {
					F_errcode(m, int32(325))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_CheckLogicalDecodingRequirements_0), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_CheckLogicalDecodingRequirements_1), int32(128), int32(_a_F_CheckLogicalDecodingRequirements_2))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
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
				v13 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_CheckLogicalDecodingRequirements[2])))
				if v13 == int32(1) {
					v18 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalDecodingRequirements[3]))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
					v21 = base.B2i32(v19 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _c_F_CheckLogicalDecodingRequirements[2])) = uint8(v21)
					v23 = v21
				} else {
					v23 = int32(0)
				}
				if v23 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, _c_F_CheckLogicalDecodingRequirements[4]))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+172))
					if base.Ui32(v26) <= base.Ui32(int32(1)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v67 = m.ExcPending
							if v67 != 0 {
								return
							} else {
								F_errmsg(m, int32(_a_F_CheckLogicalDecodingRequirements_3), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_CheckLogicalDecodingRequirements_1), int32(143), int32(_a_F_CheckLogicalDecodingRequirements_2))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
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
					return
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_errcode(m, int32(325))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_CheckLogicalDecodingRequirements_4), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_CheckLogicalDecodingRequirements_1), int32(123), int32(_a_F_CheckLogicalDecodingRequirements_2))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
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
func F_logical_heap_rewrite_flush_mappings(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int64
	_ = v102
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	if v17 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L4
	} else {
		goto L39
	}
L2:
	;
	m.G0 = v15 + int32(96)
	return
L3:
	;
	v22 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v24
	F_errmsg_internal(m, int32(_a_F_logical_heap_rewrite_flush_mappings_0), v15+int32(16))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v37 = v15 + int32(68)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	F_hash_seq_init(m, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	F_errfinish(m, int32(_a_F_logical_heap_rewrite_flush_mappings_1), int32(820), int32(_a_F_logical_heap_rewrite_flush_mappings_2))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v41 = F_hash_seq_search(m, v37)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v41 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v48 = v41
	goto L14
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	if v57 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L2
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)+48))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v57
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v64
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_logical_heap_rewrite_flush_mappings[0]))
	if v60 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v184 = F_hash_seq_search(m, v15+int32(68))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L37
	}
L19:
	;
	v69 = int32(0)
	goto L21
L20:
	;
	v69 = v68
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v69
	v71 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v73
	v76 = v57 * int32(36)
	v77 = F_palloc(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
	if v79 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v140 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v15)+92)) = v76
	v148 = F_FileWriteV(m, v141, v15+int32(88), int32(1), v140, int32(167772199))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L4
	} else {
		goto L30
	}
L24:
	;
	v83 = v48 + int32(16)
	if v79 == v83 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v87 = v79
	v89 = v77
	goto L26
L26:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	v99 = v87 - int32(36)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v100
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v99)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = v102
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v99)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v99)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v99)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v110)+4)) = v111
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v48)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v115 - int32(1)
	F_pfree(m, v99)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	goto L23
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v121 - int32(1)
	if v97 != v83 {
		v87 = v97
		v89 = v89 + int32(36)
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v148 != v76 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = v151 + base.I64_extend_i32_u(v76)
	F_XLogBeginInsert(m)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_XLogRegisterData(m, v15+int32(24), int32(40))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_XLogRegisterData(m, v77, v76)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v166 = F_XLogInsert(m, int32(9), int32(0))
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v77)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L18
L37:
	;
	if v184 != 0 {
		v48 = v184
		goto L14
	} else {
		goto L38
	}
L38:
	;
	goto L15
L39:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v148
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v48 + int32(28)
	F_errmsg(m, int32(_a_F_logical_heap_rewrite_flush_mappings_3), v15)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_logical_heap_rewrite_flush_mappings_1), int32(886), int32(_a_F_logical_heap_rewrite_flush_mappings_2))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
