package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckTableNotInUse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v97 int32
	_ = v97
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v201 int32
	_ = v201
	var v216 int32
	_ = v216
	var v243 int32
	_ = v243
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v20 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L62
	} else {
		goto L67
	}
L2:
	;
	v21 = int32(2)
	goto L4
L3:
	;
	v21 = int32(1)
	goto L4
L4:
	;
	if v17 == v21 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+119)))
	if v24|int32(32) != int32(105) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L62
	} else {
		goto L63
	}
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v31 = *(*int32)(unsafe.Add(mBase, _consts[399]))
	if v31 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	m.G0 = v15 + int32(32)
	return
L11:
	;
	if v243 != 0 {
		goto L1
	} else {
		goto L61
	}
L12:
	;
	v35 = v31
	goto L15
L13:
	;
	goto L14
L14:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	if v111 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L15:
	;
	v45 = v35 + int32(16)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if base.Ui32(v45) < base.Ui32(v46) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	v52 = v45
	goto L20
L18:
	;
	goto L19
L19:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
	if v97 != 0 {
		v35 = v97
		goto L15
	} else {
		goto L32
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v60 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L19
L22:
	;
	v71 = v60 & int32(939524096)
	if v71 == int32(134217728) {
		v82 = int32(24)
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v52+v60&int32(134217727))+8))
	if v66 != v29 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v243 = int32(1)
	goto L11
L25:
	;
	v83 = v82 + v52
	if base.Ui32(v83) < base.Ui32(v46) {
		v52 = v83
		goto L20
	} else {
		goto L31
	}
L26:
	;
	if v71 == int32(268435456) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = int32(12)
	goto L29
L28:
	;
	v78 = int32(4)
	goto L29
L29:
	;
	if v71 != int32(805306368) {
		v82 = v78
		goto L25
	} else {
		goto L30
	}
L30:
	;
	v82 = int32(16)
	goto L25
L31:
	;
	goto L21
L32:
	;
	goto L16
L33:
	;
	v243 = int32(0)
	goto L11
L34:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _consts[401]))
	if v115 <= int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	v126 = int32(0)
	goto L36
L36:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v119+v126*int32(20))))
	if v135 != 0 {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L33
L38:
	;
	v139 = v135
	goto L41
L39:
	;
	goto L40
L40:
	;
	if v111 <= v126 {
		goto L33
	} else {
		goto L59
	}
L41:
	;
	v149 = v139 + int32(16)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v139)+4))
	if base.Ui32(v149) < base.Ui32(v150) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	v156 = v149
	goto L46
L44:
	;
	goto L45
L45:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	if v201 != 0 {
		v139 = v201
		goto L41
	} else {
		goto L58
	}
L46:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	if v164 < int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v175 = v164 & int32(939524096)
	if v175 == int32(134217728) {
		v186 = int32(24)
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v156+v164&int32(134217727))+8))
	if v170 != v29 {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v243 = int32(1)
	goto L11
L51:
	;
	v187 = v186 + v156
	if base.Ui32(v187) < base.Ui32(v150) {
		v156 = v187
		goto L46
	} else {
		goto L57
	}
L52:
	;
	if v175 == int32(268435456) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v182 = int32(12)
	goto L55
L54:
	;
	v182 = int32(4)
	goto L55
L55:
	;
	if v175 != int32(805306368) {
		v186 = v182
		goto L51
	} else {
		goto L56
	}
L56:
	;
	v186 = int32(16)
	goto L51
L57:
	;
	goto L47
L58:
	;
	goto L42
L59:
	;
	v216 = v126 + int32(1)
	if v216 < v115 {
		v126 = v216
		goto L36
	} else {
		goto L60
	}
L60:
	;
	goto L37
L61:
	;
	goto L10
L62:
	;
	return
L63:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v266 + int32(4)
	F_errmsg(m, int32(265728), v15+int32(16))
	mBase = m.M
	v275 = m.ExcPending
	if v275 != 0 {
		goto L62
	} else {
		goto L65
	}
L65:
	;
	F_errfinish(m, int32(489058), int32(4426), int32(358588))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L62
	} else {
		goto L66
	}
L66:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L67:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v288 + int32(4)
	F_errmsg(m, int32(119155), v15)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L62
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(489058), int32(4435), int32(358588))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L62
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_CreateTableAsRelExists(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v21 int32
	_ = v21
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = F_RangeVarGetCreationNamespace(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		v17 = F_get_relname_relid(m, v16, v11)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				m.G0 = v7 + int32(32)
				return base.B2i32(v17 != int32(0))
			} else {
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+17)))
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117571716))
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v67
							F_errmsg(m, int32(115448), v7+int32(16))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(489605), int32(409), int32(116350))
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
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
					*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(1259)
					F_checkMembershipInCurrentExtension(m, v7+int32(20))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v35 = F_errstart(m, int32(18), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							if v35 == int32(0) {
								m.G0 = v7 + int32(32)
								return base.B2i32(v17 != int32(0))
							} else {
								F_errcode(m, int32(117571716))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v42 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
									v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v43
									F_errmsg(m, int32(329653), v7)
									mBase = m.M
									v47 = m.ExcPending
									if v47 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(489605), int32(424), int32(116350))
										mBase = m.M
										v52 = m.ExcPending
										if v52 != 0 {
											return int32(0)
										} else {
											m.G0 = v7 + int32(32)
											return base.B2i32(v17 != int32(0))
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
func F_TableFuncNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v5 != 0 {
		v12 = v5
		v15 = F_tuplestore_gettupleslot(m, v12, int32(1), int32(0), v4)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			return v4
		}
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
		F_tfuncFetchRows(m, l0, v6)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
			v12 = v11
			v15 = F_tuplestore_gettupleslot(m, v12, int32(1), int32(0), v4)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				return v4
			}
		}
	}
}
func F_table_am_handler_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(217546)
			F_errmsg(m, int32(190696), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(488711), int32(370), int32(66469))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
func F_table_beginscan_parallel(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+13)))
	if v5 != 0 {
		v17 = int32(4146400)
		v18 = int32(449)
		v19 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
		v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
		v23 = m.T0[v22].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v17, v19, v19, l1, v18)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			return v23
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v10 = F_RestoreSnapshot(m, l1+v8)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = F_RegisterSnapshot(m, v10)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				v17 = v10
				v18 = int32(961)
				v19 = int32(0)
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
				v23 = m.T0[v22].(func(*base.Module, int32, int32, int32, int32, int32, int32) int32)(m, l0, v17, v19, v19, l1, v18)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			}
		}
	}
}
func F_table_block_relation_size(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int64
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int64
	_ = v140
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	if l1 == int32(-1) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v12 != 0 {
			v40 = v12
			v42 = F_smgrnblocks(m, v40, int32(0))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int64(0)
			} else {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v45 != 0 {
					v71 = v45
					v73 = F_smgrnblocks(m, v71, int32(1))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int64(0)
					} else {
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v77 != 0 {
							v101 = v77
							v103 = F_smgrnblocks(m, v101, int32(2))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
								m.G0 = v8 - int32(-64)
								return v140 << (uint(int64(13)) % 64)
							}
						} else {
							v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v79
							v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
							*(*int64)(unsafe.Add(mBase, uint32(v8))) = v81
							v83 = F_smgropen(m, v8, v78)
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v83
								v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
								if v87 != 0 {
									v95 = v87
								} else {
									v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
									v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
									*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v89
									v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
									*(*int32)(unsafe.Add(mBase, uint32(v89))) = v91
									v93 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
									v95 = v93
								}
								*(*int32)(unsafe.Add(mBase, uint32(v83)+72)) = v95 + int32(1)
								v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								v101 = v99
								v103 = F_smgrnblocks(m, v101, int32(2))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int64(0)
								} else {
									v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
									m.G0 = v8 - int32(-64)
									return v140 << (uint(int64(13)) % 64)
								}
							}
						}
					}
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v47
					v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v49
					v53 = F_smgropen(m, v6+int32(-48), v46)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v53
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
						if v57 != 0 {
							v65 = v57
						} else {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
							*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
							*(*int32)(unsafe.Add(mBase, uint32(v59))) = v61
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
							v65 = v63
						}
						*(*int32)(unsafe.Add(mBase, uint32(v53)+72)) = v65 + int32(1)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v71 = v69
						v73 = F_smgrnblocks(m, v71, int32(1))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int64(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v77 != 0 {
								v101 = v77
								v103 = F_smgrnblocks(m, v101, int32(2))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int64(0)
								} else {
									v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
									m.G0 = v8 - int32(-64)
									return v140 << (uint(int64(13)) % 64)
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v79
								v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v81
								v83 = F_smgropen(m, v8, v78)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v83
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
									if v87 != 0 {
										v95 = v87
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
										*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v89
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
										*(*int32)(unsafe.Add(mBase, uint32(v89))) = v91
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
										v95 = v93
									}
									*(*int32)(unsafe.Add(mBase, uint32(v83)+72)) = v95 + int32(1)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v101 = v99
									v103 = F_smgrnblocks(m, v101, int32(2))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int64(0)
									} else {
										v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
										m.G0 = v8 - int32(-64)
										return v140 << (uint(int64(13)) % 64)
									}
								}
							}
						}
					}
				}
			}
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v14
			v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v16
			v20 = F_smgropen(m, v6+int32(-32), v13)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
				if v26 != 0 {
					v34 = v26
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v28
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v20)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v30
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
					v34 = v32
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v34 + int32(1)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v40 = v38
				v42 = F_smgrnblocks(m, v40, int32(0))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int64(0)
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v45 != 0 {
						v71 = v45
						v73 = F_smgrnblocks(m, v71, int32(1))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int64(0)
						} else {
							v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v77 != 0 {
								v101 = v77
								v103 = F_smgrnblocks(m, v101, int32(2))
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int64(0)
								} else {
									v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
									m.G0 = v8 - int32(-64)
									return v140 << (uint(int64(13)) % 64)
								}
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v79
								v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
								*(*int64)(unsafe.Add(mBase, uint32(v8))) = v81
								v83 = F_smgropen(m, v8, v78)
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v83
									v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
									if v87 != 0 {
										v95 = v87
									} else {
										v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
										v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
										*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v89
										v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
										*(*int32)(unsafe.Add(mBase, uint32(v89))) = v91
										v93 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
										v95 = v93
									}
									*(*int32)(unsafe.Add(mBase, uint32(v83)+72)) = v95 + int32(1)
									v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									v101 = v99
									v103 = F_smgrnblocks(m, v101, int32(2))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int64(0)
									} else {
										v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
										m.G0 = v8 - int32(-64)
										return v140 << (uint(int64(13)) % 64)
									}
								}
							}
						}
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v47
						v49 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
						*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v49
						v53 = F_smgropen(m, v6+int32(-48), v46)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v53
							v57 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
							if v57 != 0 {
								v65 = v57
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v53)+80))
								*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59
								v61 = *(*int32)(unsafe.Add(mBase, uint32(v53)+76))
								*(*int32)(unsafe.Add(mBase, uint32(v59))) = v61
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v53)+72))
								v65 = v63
							}
							*(*int32)(unsafe.Add(mBase, uint32(v53)+72)) = v65 + int32(1)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v71 = v69
							v73 = F_smgrnblocks(m, v71, int32(1))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int64(0)
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v77 != 0 {
									v101 = v77
									v103 = F_smgrnblocks(m, v101, int32(2))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int64(0)
									} else {
										v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
										m.G0 = v8 - int32(-64)
										return v140 << (uint(int64(13)) % 64)
									}
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v79
									v81 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
									*(*int64)(unsafe.Add(mBase, uint32(v8))) = v81
									v83 = F_smgropen(m, v8, v78)
									mBase = m.M
									v84 = m.ExcPending
									if v84 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v83
										v87 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
										if v87 != 0 {
											v95 = v87
										} else {
											v88 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
											v89 = *(*int32)(unsafe.Add(mBase, uint32(v83)+80))
											*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v89
											v91 = *(*int32)(unsafe.Add(mBase, uint32(v83)+76))
											*(*int32)(unsafe.Add(mBase, uint32(v89))) = v91
											v93 = *(*int32)(unsafe.Add(mBase, uint32(v83)+72))
											v95 = v93
										}
										*(*int32)(unsafe.Add(mBase, uint32(v83)+72)) = v95 + int32(1)
										v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										v101 = v99
										v103 = F_smgrnblocks(m, v101, int32(2))
										mBase = m.M
										v104 = m.ExcPending
										if v104 != 0 {
											return int64(0)
										} else {
											v140 = base.I64_extend_i32_u(v42) + base.I64_extend_i32_u(v73) + base.I64_extend_i32_u(v103)
											m.G0 = v8 - int32(-64)
											return v140 << (uint(int64(13)) % 64)
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
		v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v107 != 0 {
			v133 = v107
			v134 = F_smgrnblocks(m, v133, l1)
			mBase = m.M
			v135 = m.ExcPending
			if v135 != 0 {
				return int64(0)
			} else {
				v140 = base.I64_extend_i32_u(v134)
				m.G0 = v8 - int32(-64)
				return v140 << (uint(int64(13)) % 64)
			}
		} else {
			v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+56)) = v109
			v111 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v8)+48)) = v111
			v115 = F_smgropen(m, v6+int32(-16), v108)
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v115
				v119 = *(*int32)(unsafe.Add(mBase, uint32(v115)+72))
				if v119 != 0 {
					v127 = v119
				} else {
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v115)+76))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+80))
					*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v121
					v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+76))
					*(*int32)(unsafe.Add(mBase, uint32(v121))) = v123
					v125 = *(*int32)(unsafe.Add(mBase, uint32(v115)+72))
					v127 = v125
				}
				*(*int32)(unsafe.Add(mBase, uint32(v115)+72)) = v127 + int32(1)
				v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v133 = v131
				v134 = F_smgrnblocks(m, v133, l1)
				mBase = m.M
				v135 = m.ExcPending
				if v135 != 0 {
					return int64(0)
				} else {
					v140 = base.I64_extend_i32_u(v134)
					m.G0 = v8 - int32(-64)
					return v140 << (uint(int64(13)) % 64)
				}
			}
		}
	}
}
func F_transformTableConstraint(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v10 {
	case 0, 2, 10, 11, 12, 13, 14, 15:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v130 = m.ExcPending
		if v130 != 0 {
			return
		} else {
			v131 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v131
			F_errmsg_internal(m, int32(470663), v7+int32(16))
			mBase = m.M
			v137 = m.ExcPending
			if v137 != 0 {
				return
			} else {
				F_errfinish(m, int32(494160), int32(1101), int32(90216))
				mBase = m.M
				v142 = m.ExcPending
				if v142 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		v80 = int32(32)
		v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+60)))
		if v81 != int32(1) {
			v158 = v80
			v159 = l0 + v158
			v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
			v161 = F_lappend(m, v160, l1)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+17)))
			if v84 != int32(1) {
				v158 = v80
				v159 = l0 + v158
				v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
				v161 = F_lappend(m, v160, l1)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
					m.G0 = v7 + int32(32)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						F_errmsg(m, int32(514838), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return
						} else {
							F_errfinish(m, int32(494160), int32(1077), int32(90216))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
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
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v146 = m.ExcPending
		if v146 != 0 {
			return
		} else {
			v147 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v147
			F_errmsg_internal(m, int32(479115), v7)
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return
			} else {
				F_errfinish(m, int32(494160), int32(1106), int32(90216))
				mBase = m.M
				v156 = m.ExcPending
				if v156 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 5:
		v158 = int32(28)
		v159 = l0 + v158
		v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
		v161 = F_lappend(m, v160, l1)
		mBase = m.M
		v162 = m.ExcPending
		if v162 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
			m.G0 = v7 + int32(32)
			return
		}
	case 6:
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v11 != int32(1) {
			v158 = int32(40)
			v159 = l0 + v158
			v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
			v161 = F_lappend(m, v160, l1)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return
				} else {
					F_errmsg(m, int32(163938), int32(0))
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
						F_parser_errposition(m, v25, v26)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_errfinish(m, int32(494160), int32(1045), int32(90216))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
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
	case 7:
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v34 != int32(1) {
			v158 = int32(40)
			v159 = l0 + v158
			v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
			v161 = F_lappend(m, v160, l1)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					F_errmsg(m, int32(164116), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
						F_parser_errposition(m, v48, v49)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return
						} else {
							F_errfinish(m, int32(494160), int32(1055), int32(90216))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
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
	case 8:
		v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v57 != int32(1) {
			v158 = int32(40)
			v159 = l0 + v158
			v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
			v161 = F_lappend(m, v160, l1)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return
				} else {
					F_errmsg(m, int32(164058), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
						F_parser_errposition(m, v71, v72)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_errfinish(m, int32(494160), int32(1065), int32(90216))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
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
	case 9:
		v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v103 != int32(1) {
			v158 = int32(36)
			v159 = l0 + v158
			v160 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
			v161 = F_lappend(m, v160, l1)
			mBase = m.M
			v162 = m.ExcPending
			if v162 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v159))) = v161
				m.G0 = v7 + int32(32)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v113 = m.ExcPending
				if v113 != 0 {
					return
				} else {
					F_errmsg(m, int32(163998), int32(0))
					mBase = m.M
					v117 = m.ExcPending
					if v117 != 0 {
						return
					} else {
						v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l1)+104))
						F_parser_errposition(m, v118, v119)
						mBase = m.M
						v121 = m.ExcPending
						if v121 != 0 {
							return
						} else {
							F_errfinish(m, int32(494160), int32(1088), int32(90216))
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
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
}
