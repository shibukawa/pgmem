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
		v4 = *(*int32)(unsafe.Add(mBase, _consts[27]))
		if int32(1) < v4 {
			v8 = *(*int32)(unsafe.Add(mBase, _consts[100]))
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
						F_errmsg(m, int32(243786), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errfinish(m, int32(475923), int32(128), int32(115411))
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
				v13 = int32(*(*uint8)(unsafe.Add(mBase, _consts[189])))
				if v13 == int32(1) {
					v18 = *(*int32)(unsafe.Add(mBase, _consts[190]))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+316))
					v21 = base.B2i32(v19 != int32(2))
					*(*uint8)(unsafe.Add(mBase, _consts[189])) = uint8(v21)
					v23 = v21
				} else {
					v23 = int32(0)
				}
				if v23 != 0 {
					v25 = *(*int32)(unsafe.Add(mBase, _consts[275]))
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
								F_errmsg(m, int32(16896), int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return
								} else {
									F_errfinish(m, int32(475923), int32(143), int32(115411))
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
					F_errmsg(m, int32(688537), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errfinish(m, int32(475923), int32(123), int32(115411))
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
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int64
	_ = v102
	var v104 int32
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int64
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
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
	v206 = m.ExcPending
	if v206 != 0 {
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
	F_errmsg_internal(m, int32(158268), v15+int32(16))
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
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	F_hash_seq_init(m, v15+int32(68), v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L11
	}
L9:
	;
	F_errfinish(m, int32(474476), int32(820), int32(146944))
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
	v43 = F_hash_seq_search(m, v15+int32(68))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	if v43 == int32(0) {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v50 = v43
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	if v59 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L2
L16:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+48))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+117)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v59
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+24)) = v66
	v70 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	if v62 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v186 = F_hash_seq_search(m, v15+int32(68))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L4
	} else {
		goto L37
	}
L19:
	;
	v71 = int32(0)
	goto L21
L20:
	;
	v71 = v70
	goto L21
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+28)) = v71
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+40)) = v73
	v75 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+56)) = v75
	v78 = v59 * int32(36)
	v79 = F_palloc(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v81 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v15)+92)) = v78
	v150 = F_FileWriteV(m, v143, v15+int32(88), int32(1), v142, int32(167772199))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L30
	}
L24:
	;
	v85 = v50 + int32(16)
	if v81 == v85 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = v79
	v91 = v81
	goto L26
L26:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	v101 = v91 - int32(36)
	v102 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v89))) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v101)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v101)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v101)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v101)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v110
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v91)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v113))) = v115
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v50)+24)) = v117 - int32(1)
	F_pfree(m, v101)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L28
	}
L27:
	;
	goto L23
L28:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = v123 - int32(1)
	if v99 != v85 {
		v89 = v89 + int32(36)
		v91 = v99
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if v150 != v78 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v50)+8)) = v153 + base.I64_extend_i32_u(v78)
	F_XLogBeginInsert(m)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_XLogRegisterData(m, v15+int32(24), int32(40))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	F_XLogRegisterData(m, v79, v78)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v168 = F_XLogInsert(m, int32(9), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	F_pfree(m, v79)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	goto L18
L37:
	;
	if v186 != 0 {
		v50 = v186
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
	v208 = m.ExcPending
	if v208 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v78
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v50 + int32(28)
	F_errmsg(m, int32(282451), v15)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(474476), int32(886), int32(146944))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
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
