package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_RangeVarCallbackMaintainsTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 == int32(0) {
		m.G0 = v8 + int32(16)
		return
	} else {
		v12 = F_get_rel_relkind(m, l1)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v15 = v12 & int32(255)
			v17 = v15 - int32(109)
			if base.Ui32(int32(7)) < base.Ui32(v17) {
				if v15 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v55 = m.ExcPending
					if v55 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v59
							F_errmsg(m, int32(32365), v8)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								F_errfinish(m, int32(493890), int32(19516), int32(396529))
								mBase = m.M
								v68 = m.ExcPending
								if v68 != 0 {
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
			} else {
				if int32(1)<<(uint(v17)%32)&int32(169) == int32(0) {
					if v15 == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v55 = m.ExcPending
						if v55 != 0 {
							return
						} else {
							F_errcode(m, int32(151027844))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return
							} else {
								v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = v59
								F_errmsg(m, int32(32365), v8)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									F_errfinish(m, int32(493890), int32(19516), int32(396529))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
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
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, _consts[239]))
					v29 = F_pg_class_aclcheck(m, l1, v27, int64(16384))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						if v29 == int32(0) {
							m.G0 = v8 + int32(16)
							return
						} else {
							v33 = F_get_rel_relkind(m, l1)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return
							} else {
								switch v33 - int32(73) {
								case 0, 32:
									v46 = int32(20)
								default:
									v44 = int32(41)
									v46 = v44
								case 10:
									v46 = int32(37)
								case 29:
									v44 = int32(18)
									v46 = v44
								case 36:
									v46 = int32(23)
								case 45:
									v46 = int32(51)
								}
								v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								F_aclcheck_error(m, v29, v46, v47)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									m.G0 = v8 + int32(16)
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
func F_addRangeTableEntryForSubquery(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v4 = l3
	v5 = l4
	v6 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v22 = F_palloc0(m, int32(136))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(101)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l2
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v39 != 0 {
		goto L9
	} else {
		goto L10
	}
L4:
	;
	v32 = F_copyObjectImpl(m, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v36 = F_makeAlias(m, int32(15408), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v38 = v32
	goto L3
L8:
	;
	v38 = v36
	goto L3
L9:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v41 = v40
	goto L11
L10:
	;
	v41 = v6
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	if v41 <= v119 {
		goto L34
	} else {
		goto L35
	}
L13:
	;
	v119 = v6
	v121 = v6
	v122 = v6
	v123 = v6
	goto L12
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v50 = int32(0)
	v58 = v6
	v60 = v6
	v61 = v6
	v62 = v6
	goto L16
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v50<<(uint(int32(2))%32))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+26)))
	if v70 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v119 = v101
	v121 = v102
	v122 = v103
	v123 = v104
	goto L12
L18:
	;
	v74 = v58 + int32(1)
	if v41 < v74 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v101 = v58
	v102 = v60
	v103 = v61
	v104 = v62
	goto L20
L20:
	;
	v107 = v50 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v107 < v108 {
		v50 = v107
		v58 = v101
		v60 = v102
		v61 = v103
		v62 = v104
		goto L16
	} else {
		goto L33
	}
L21:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v77 = F_pstrdup(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v87 = F_exprType(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v80 = F_makeString(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v82 = F_lappend(m, v79, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v82
	goto L23
L27:
	;
	v89 = F_lappend_oid(m, v62, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v92 = F_exprTypmod(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v94 = F_lappend_int(m, v61, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v97 = F_exprCollation(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v99 = F_lappend_oid(m, v60, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v101 = v74
	v102 = v99
	v103 = v94
	v104 = v89
	goto L20
L33:
	;
	goto L17
L34:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+125)) = uint8(v5)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+124)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v38
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v131 = F_lappend(m, v130, v22)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L42
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131
	if v131 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v136 = v135
	goto L40
L39:
	;
	v136 = int32(0)
	goto L40
L40:
	;
	v137 = F_buildNSItemFromLists(m, v22, v136, v123, v122, v121)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+20)) = uint8(base.B2i32(l2 != int32(0)))
	m.G0 = v19 + int32(16)
	return v137
L42:
	;
	F_errcode(m, int32(393348))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v153
	F_errmsg(m, int32(456708), v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(495726), int32(1709), int32(15674))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_range_multirange(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	v4 = F_SearchSysCache1(m, int32(55), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13)+8))
			F_ReleaseCatCache(m, v4)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return v15
			}
		}
	}
}
func F_makeRangeVarFromAnyName(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_palloc0(m, int32(28))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(3)
		if l0 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v49 = F_NameListToString(m, l0)
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
						F_errmsg(m, int32(204519), v8)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							F_scanner_errposition(m, l1, l2)
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(26978), int32(19357), int32(381469))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
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
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			switch v19 - int32(1) {
			case 0:
				*(*int64)(unsafe.Add(mBase, uint32(v11)+4)) = int64(0)
				v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = v64
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				v69 = int32(112)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
				m.G0 = v8 + int32(16)
				return v11
			case 1:
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(0)
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v26
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = v28 + int32(4)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				v69 = int32(112)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
				m.G0 = v8 + int32(16)
				return v11
			case 2:
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v33
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v65 = v39 + int32(8)
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l1
				v69 = int32(112)
				*(*uint8)(unsafe.Add(mBase, uint32(v11)+17)) = uint8(v69)
				*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v67
				m.G0 = v8 + int32(16)
				return v11
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return int32(0)
					} else {
						v49 = F_NameListToString(m, l0)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v49
							F_errmsg(m, int32(204519), v8)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								F_scanner_errposition(m, l1, l2)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(26978), int32(19357), int32(381469))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
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
func F_range_adjacent_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v72 = v4
		m.G0 = v8 + int32(80)
		return v72
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v72 = v4
			m.G0 = v8 + int32(80)
			return v72
		} else {
			F_range_deserialize(m, l0, l1, v8+int32(72), v8-int32(-64), v8+int32(47))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				F_multirange_get_bounds(m, l0, l2, int32(0), v8+int32(56), v8+int32(48))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					v40 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v40
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+56))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v42
					v49 = F_bounds_adjacent(m, l0, v8+int32(32), v8+int32(24))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						if v49 != 0 {
							v72 = int32(1)
							m.G0 = v8 + int32(80)
							return v72
						} else {
							if int32(2) <= v32 {
								F_multirange_get_bounds(m, l0, l2, v32-int32(1), v8+int32(56), v8+int32(48))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
									return int32(0)
								} else {
									v61 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v61
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v63
									v69 = F_bounds_adjacent(m, l0, v8+int32(16), v8+int32(8))
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return int32(0)
									} else {
										v72 = v69
										m.G0 = v8 + int32(80)
										return v72
									}
								}
							} else {
								v61 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v61
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v8)+72))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v63
								v69 = F_bounds_adjacent(m, l0, v8+int32(16), v8+int32(8))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
									return int32(0)
								} else {
									v72 = v69
									m.G0 = v8 + int32(80)
									return v72
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_range_compare(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_range_deserialize(m, l2, v10, v7+int32(40), v7+int32(32), v7+int32(15))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l2, v9, v7+int32(24), v7+int32(16), v7+int32(14))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(1)
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+15)))
			if v31 == v29 {
				v174 = int32(0) - (v30^int32(1))&int32(255)
				m.G0 = v7 + int32(48)
				return v174
			} else {
				if v30&int32(1) != 0 {
					v174 = v29
					m.G0 = v7 + int32(48)
					return v174
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+28)))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+44)))
					if v43 == int32(1) {
						v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
						if v42&int32(1) != 0 {
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
							if v49 == v46&int32(255) {
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
								v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
								if v108 == int32(1) {
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
									if v107&int32(1) != 0 {
										v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
										if v114 == v111&int32(255) {
											v174 = int32(0)
										} else {
											v120 = int32(1)
											if v111&v120 != 0 {
												v123 = int32(-1)
											} else {
												v123 = v120
											}
											v174 = v123
										}
									} else {
										v125 = int32(1)
										if v111&v125 != 0 {
											v128 = int32(-1)
										} else {
											v128 = v125
										}
										v174 = v128
									}
									m.G0 = v7 + int32(48)
									return v174
								} else {
									if v107&int32(1) != 0 {
										v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
										if v133 != 0 {
											v134 = int32(1)
										} else {
											v134 = int32(-1)
										}
										v174 = v134
										m.G0 = v7 + int32(48)
										return v174
									} else {
										v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
										v138 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
										v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
										v140 = F_FunctionCall2Coll(m, l2+int32(212), v137, v138, v139)
										mBase = m.M
										v141 = m.ExcPending
										if v141 != 0 {
											return int32(0)
										} else {
											if v140 != 0 {
												v174 = v140
											} else {
												v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
												v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
												if v143 == int32(0) {
													v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
													if v142&int32(1) == int32(0) {
														v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v151 == v146&int32(255) {
															v174 = int32(0)
														} else {
															v156 = int32(1)
															if v146&v156 != 0 {
																v160 = v156
															} else {
																v160 = int32(-1)
															}
															v174 = v160
														}
													} else {
														v161 = int32(1)
														if v146&v161 != 0 {
															v165 = v161
														} else {
															v165 = int32(-1)
														}
														v174 = v165
													}
												} else {
													if v142&int32(1) != 0 {
														v174 = int32(0)
													} else {
														v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v171 != 0 {
															v172 = int32(-1)
														} else {
															v172 = int32(1)
														}
														v174 = v172
													}
												}
											}
											m.G0 = v7 + int32(48)
											return v174
										}
									}
								}
							} else {
								v54 = int32(1)
								if v46&v54 != 0 {
									v57 = int32(-1)
								} else {
									v57 = v54
								}
								v174 = v57
								m.G0 = v7 + int32(48)
								return v174
							}
						} else {
							v59 = int32(1)
							if v46&v59 != 0 {
								v62 = int32(-1)
							} else {
								v62 = v59
							}
							v174 = v62
							m.G0 = v7 + int32(48)
							return v174
						}
					} else {
						if v42&int32(1) != 0 {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
							if v67 != 0 {
								v68 = int32(1)
							} else {
								v68 = int32(-1)
							}
							v174 = v68
							m.G0 = v7 + int32(48)
							return v174
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
							v74 = F_FunctionCall2Coll(m, l2+int32(212), v71, v72, v73)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								if v74 != 0 {
									v174 = v74
									m.G0 = v7 + int32(48)
									return v174
								} else {
									v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+29)))
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+45)))
									if v77 == int32(0) {
										v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+46)))
										if v76&int32(1) == int32(0) {
											v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v85 == v80&int32(255) {
												v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
												if v108 == int32(1) {
													v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
													if v107&int32(1) != 0 {
														v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v114 == v111&int32(255) {
															v174 = int32(0)
														} else {
															v120 = int32(1)
															if v111&v120 != 0 {
																v123 = int32(-1)
															} else {
																v123 = v120
															}
															v174 = v123
														}
													} else {
														v125 = int32(1)
														if v111&v125 != 0 {
															v128 = int32(-1)
														} else {
															v128 = v125
														}
														v174 = v128
													}
													m.G0 = v7 + int32(48)
													return v174
												} else {
													if v107&int32(1) != 0 {
														v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
														if v133 != 0 {
															v134 = int32(1)
														} else {
															v134 = int32(-1)
														}
														v174 = v134
														m.G0 = v7 + int32(48)
														return v174
													} else {
														v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
														v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
														v140 = F_FunctionCall2Coll(m, l2+int32(212), v137, v138, v139)
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int32(0)
														} else {
															if v140 != 0 {
																v174 = v140
															} else {
																v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
																v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
																if v143 == int32(0) {
																	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
																	if v142&int32(1) == int32(0) {
																		v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																		if v151 == v146&int32(255) {
																			v174 = int32(0)
																		} else {
																			v156 = int32(1)
																			if v146&v156 != 0 {
																				v160 = v156
																			} else {
																				v160 = int32(-1)
																			}
																			v174 = v160
																		}
																	} else {
																		v161 = int32(1)
																		if v146&v161 != 0 {
																			v165 = v161
																		} else {
																			v165 = int32(-1)
																		}
																		v174 = v165
																	}
																} else {
																	if v142&int32(1) != 0 {
																		v174 = int32(0)
																	} else {
																		v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																		if v171 != 0 {
																			v172 = int32(-1)
																		} else {
																			v172 = int32(1)
																		}
																		v174 = v172
																	}
																}
															}
															m.G0 = v7 + int32(48)
															return v174
														}
													}
												}
											} else {
												v89 = int32(1)
												if v80&v89 != 0 {
													v93 = v89
												} else {
													v93 = int32(-1)
												}
												v174 = v93
												m.G0 = v7 + int32(48)
												return v174
											}
										} else {
											v94 = int32(1)
											if v80&v94 != 0 {
												v98 = v94
											} else {
												v98 = int32(-1)
											}
											v174 = v98
											m.G0 = v7 + int32(48)
											return v174
										}
									} else {
										if v76&int32(1) != 0 {
											v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+20)))
											v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+36)))
											if v108 == int32(1) {
												v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
												if v107&int32(1) != 0 {
													v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
													if v114 == v111&int32(255) {
														v174 = int32(0)
													} else {
														v120 = int32(1)
														if v111&v120 != 0 {
															v123 = int32(-1)
														} else {
															v123 = v120
														}
														v174 = v123
													}
												} else {
													v125 = int32(1)
													if v111&v125 != 0 {
														v128 = int32(-1)
													} else {
														v128 = v125
													}
													v174 = v128
												}
												m.G0 = v7 + int32(48)
												return v174
											} else {
												if v107&int32(1) != 0 {
													v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
													if v133 != 0 {
														v134 = int32(1)
													} else {
														v134 = int32(-1)
													}
													v174 = v134
													m.G0 = v7 + int32(48)
													return v174
												} else {
													v137 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
													v139 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
													v140 = F_FunctionCall2Coll(m, l2+int32(212), v137, v138, v139)
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return int32(0)
													} else {
														if v140 != 0 {
															v174 = v140
														} else {
															v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+21)))
															v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+37)))
															if v143 == int32(0) {
																v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+38)))
																if v142&int32(1) == int32(0) {
																	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																	if v151 == v146&int32(255) {
																		v174 = int32(0)
																	} else {
																		v156 = int32(1)
																		if v146&v156 != 0 {
																			v160 = v156
																		} else {
																			v160 = int32(-1)
																		}
																		v174 = v160
																	}
																} else {
																	v161 = int32(1)
																	if v146&v161 != 0 {
																		v165 = v161
																	} else {
																		v165 = int32(-1)
																	}
																	v174 = v165
																}
															} else {
																if v142&int32(1) != 0 {
																	v174 = int32(0)
																} else {
																	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
																	if v171 != 0 {
																		v172 = int32(-1)
																	} else {
																		v172 = int32(1)
																	}
																	v174 = v172
																}
															}
														}
														m.G0 = v7 + int32(48)
														return v174
													}
												}
											}
										} else {
											v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+30)))
											if v103 != 0 {
												v104 = int32(-1)
											} else {
												v104 = int32(1)
											}
											v174 = v104
											m.G0 = v7 + int32(48)
											return v174
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
func F_range_constructor3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_get_fn_expr_rettype(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
		if v21 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v22 == v16 {
				v32 = v21
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
				if v33 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v112 = m.ExcPending
					if v112 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(302392), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493562), int32(424), int32(551811))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
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
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v37 = F_pg_detoast_datum_packed(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v39 = F_text_to_cstring(m, v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							if v41 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(16801924))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(156657), int32(0))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(653961), int32(0))
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(493562), int32(2321), int32(156604))
												mBase = m.M
												v145 = m.ExcPending
												if v145 != 0 {
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
							} else {
								v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
								if v44 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(156657), int32(0))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(653961), int32(0))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(493562), int32(2321), int32(156604))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
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
								} else {
									v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
									if v47 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(156657), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(653961), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(493562), int32(2321), int32(156604))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
									} else {
										if v41 != int32(40) {
											if v41 != int32(91) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16801924))
													mBase = m.M
													v152 = m.ExcPending
													if v152 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(156657), int32(0))
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int32(0)
														} else {
															F_errhint(m, int32(653961), int32(0))
															mBase = m.M
															v160 = m.ExcPending
															if v160 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(493562), int32(2334), int32(156604))
																mBase = m.M
																v165 = m.ExcPending
																if v165 != 0 {
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
											} else {
												v54 = int32(2)
												if v44 != int32(41) {
													if v44 != int32(93) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(156657), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(653961), int32(0))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(493562), int32(2348), int32(156604))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
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
													} else {
														v61 = v54 | int32(4)
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
														}
													}
												} else {
													v61 = v54
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
													if v62 != 0 {
														v65 = int32(0)
													} else {
														v65 = v14
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
													if v74 != 0 {
														v82 = v75
													} else {
														v82 = v13
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v90
													}
												}
											}
										} else {
											v54 = int32(0)
											if v44 != int32(41) {
												if v44 != int32(93) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v169 = m.ExcPending
													if v169 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16801924))
														mBase = m.M
														v172 = m.ExcPending
														if v172 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(156657), int32(0))
															mBase = m.M
															v176 = m.ExcPending
															if v176 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(653961), int32(0))
																mBase = m.M
																v180 = m.ExcPending
																if v180 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(493562), int32(2348), int32(156604))
																	mBase = m.M
																	v185 = m.ExcPending
																	if v185 != 0 {
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
												} else {
													v61 = v54 | int32(4)
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
													if v62 != 0 {
														v65 = int32(0)
													} else {
														v65 = v14
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
													if v74 != 0 {
														v82 = v75
													} else {
														v82 = v13
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v90
													}
												}
											} else {
												v61 = v54
												v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
												if v62 != 0 {
													v65 = int32(0)
												} else {
													v65 = v14
												}
												*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
												v67 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
												v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
												v75 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
												if v74 != 0 {
													v82 = v75
												} else {
													v82 = v13
												}
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
												v88 = int32(0)
												v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int32(0)
												} else {
													m.G0 = v11 + int32(32)
													return v90
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
				v25 = F_lookup_type_cache(m, v16, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
							F_errmsg_internal(m, int32(369887), v11)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493562), int32(1776), int32(398373))
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
						if v33 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v112 = m.ExcPending
							if v112 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(130))
								mBase = m.M
								v115 = m.ExcPending
								if v115 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(302392), int32(0))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(493562), int32(424), int32(551811))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
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
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							v37 = F_pg_detoast_datum_packed(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								v39 = F_text_to_cstring(m, v37)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
									if v41 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(156657), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(653961), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(493562), int32(2321), int32(156604))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
									} else {
										v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
										if v44 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16801924))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(156657), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(653961), int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(493562), int32(2321), int32(156604))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
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
										} else {
											v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
											if v47 != 0 {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v129 = m.ExcPending
												if v129 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(16801924))
													mBase = m.M
													v132 = m.ExcPending
													if v132 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(156657), int32(0))
														mBase = m.M
														v136 = m.ExcPending
														if v136 != 0 {
															return int32(0)
														} else {
															F_errhint(m, int32(653961), int32(0))
															mBase = m.M
															v140 = m.ExcPending
															if v140 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(493562), int32(2321), int32(156604))
																mBase = m.M
																v145 = m.ExcPending
																if v145 != 0 {
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
											} else {
												if v41 != int32(40) {
													if v41 != int32(91) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v152 = m.ExcPending
															if v152 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(156657), int32(0))
																mBase = m.M
																v156 = m.ExcPending
																if v156 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(653961), int32(0))
																	mBase = m.M
																	v160 = m.ExcPending
																	if v160 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(493562), int32(2334), int32(156604))
																		mBase = m.M
																		v165 = m.ExcPending
																		if v165 != 0 {
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
													} else {
														v54 = int32(2)
														if v44 != int32(41) {
															if v44 != int32(93) {
																F_errstart_cold(m, int32(21), int32(0))
																mBase = m.M
																v169 = m.ExcPending
																if v169 != 0 {
																	return int32(0)
																} else {
																	F_errcode(m, int32(16801924))
																	mBase = m.M
																	v172 = m.ExcPending
																	if v172 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(156657), int32(0))
																		mBase = m.M
																		v176 = m.ExcPending
																		if v176 != 0 {
																			return int32(0)
																		} else {
																			F_errhint(m, int32(653961), int32(0))
																			mBase = m.M
																			v180 = m.ExcPending
																			if v180 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(493562), int32(2348), int32(156604))
																				mBase = m.M
																				v185 = m.ExcPending
																				if v185 != 0 {
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
															} else {
																v61 = v54 | int32(4)
																v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
																if v62 != 0 {
																	v65 = int32(0)
																} else {
																	v65 = v14
																}
																*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
																v67 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
																v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
																v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
																v75 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
																if v74 != 0 {
																	v82 = v75
																} else {
																	v82 = v13
																}
																*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
																v88 = int32(0)
																v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v11 + int32(32)
																	return v90
																}
															}
														} else {
															v61 = v54
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
															if v62 != 0 {
																v65 = int32(0)
															} else {
																v65 = v14
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
															if v74 != 0 {
																v82 = v75
															} else {
																v82 = v13
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v90
															}
														}
													}
												} else {
													v54 = int32(0)
													if v44 != int32(41) {
														if v44 != int32(93) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16801924))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(156657), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		F_errhint(m, int32(653961), int32(0))
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(493562), int32(2348), int32(156604))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
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
														} else {
															v61 = v54 | int32(4)
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
															if v62 != 0 {
																v65 = int32(0)
															} else {
																v65 = v14
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
															if v74 != 0 {
																v82 = v75
															} else {
																v82 = v13
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v90
															}
														}
													} else {
														v61 = v54
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
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
		} else {
			v25 = F_lookup_type_cache(m, v16, int32(2048))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
				if v27 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
						F_errmsg_internal(m, int32(369887), v11)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493562), int32(1776), int32(398373))
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
					v32 = v25
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
					if v33 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(130))
							mBase = m.M
							v115 = m.ExcPending
							if v115 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(302392), int32(0))
								mBase = m.M
								v119 = m.ExcPending
								if v119 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493562), int32(424), int32(551811))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
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
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v37 = F_pg_detoast_datum_packed(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							v39 = F_text_to_cstring(m, v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int32(0)
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
								if v41 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v132 = m.ExcPending
										if v132 != 0 {
											return int32(0)
										} else {
											F_errmsg(m, int32(156657), int32(0))
											mBase = m.M
											v136 = m.ExcPending
											if v136 != 0 {
												return int32(0)
											} else {
												F_errhint(m, int32(653961), int32(0))
												mBase = m.M
												v140 = m.ExcPending
												if v140 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(493562), int32(2321), int32(156604))
													mBase = m.M
													v145 = m.ExcPending
													if v145 != 0 {
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
								} else {
									v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+1)))
									if v44 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v132 = m.ExcPending
											if v132 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(156657), int32(0))
												mBase = m.M
												v136 = m.ExcPending
												if v136 != 0 {
													return int32(0)
												} else {
													F_errhint(m, int32(653961), int32(0))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(493562), int32(2321), int32(156604))
														mBase = m.M
														v145 = m.ExcPending
														if v145 != 0 {
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
									} else {
										v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+2)))
										if v47 != 0 {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v129 = m.ExcPending
											if v129 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(16801924))
												mBase = m.M
												v132 = m.ExcPending
												if v132 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(156657), int32(0))
													mBase = m.M
													v136 = m.ExcPending
													if v136 != 0 {
														return int32(0)
													} else {
														F_errhint(m, int32(653961), int32(0))
														mBase = m.M
														v140 = m.ExcPending
														if v140 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(493562), int32(2321), int32(156604))
															mBase = m.M
															v145 = m.ExcPending
															if v145 != 0 {
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
										} else {
											if v41 != int32(40) {
												if v41 != int32(91) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(16801924))
														mBase = m.M
														v152 = m.ExcPending
														if v152 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(156657), int32(0))
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(653961), int32(0))
																mBase = m.M
																v160 = m.ExcPending
																if v160 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(493562), int32(2334), int32(156604))
																	mBase = m.M
																	v165 = m.ExcPending
																	if v165 != 0 {
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
												} else {
													v54 = int32(2)
													if v44 != int32(41) {
														if v44 != int32(93) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v169 = m.ExcPending
															if v169 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(16801924))
																mBase = m.M
																v172 = m.ExcPending
																if v172 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(156657), int32(0))
																	mBase = m.M
																	v176 = m.ExcPending
																	if v176 != 0 {
																		return int32(0)
																	} else {
																		F_errhint(m, int32(653961), int32(0))
																		mBase = m.M
																		v180 = m.ExcPending
																		if v180 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(493562), int32(2348), int32(156604))
																			mBase = m.M
																			v185 = m.ExcPending
																			if v185 != 0 {
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
														} else {
															v61 = v54 | int32(4)
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
															if v62 != 0 {
																v65 = int32(0)
															} else {
																v65 = v14
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
															if v74 != 0 {
																v82 = v75
															} else {
																v82 = v13
															}
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int32(0)
															} else {
																m.G0 = v11 + int32(32)
																return v90
															}
														}
													} else {
														v61 = v54
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
														}
													}
												}
											} else {
												v54 = int32(0)
												if v44 != int32(41) {
													if v44 != int32(93) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v169 = m.ExcPending
														if v169 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v172 = m.ExcPending
															if v172 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(156657), int32(0))
																mBase = m.M
																v176 = m.ExcPending
																if v176 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(653961), int32(0))
																	mBase = m.M
																	v180 = m.ExcPending
																	if v180 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(493562), int32(2348), int32(156604))
																		mBase = m.M
																		v185 = m.ExcPending
																		if v185 != 0 {
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
													} else {
														v61 = v54 | int32(4)
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
														if v62 != 0 {
															v65 = int32(0)
														} else {
															v65 = v14
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
														if v74 != 0 {
															v82 = v75
														} else {
															v82 = v13
														}
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int32(0)
														} else {
															m.G0 = v11 + int32(32)
															return v90
														}
													}
												} else {
													v61 = v54
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)) = uint8(v62)
													if v62 != 0 {
														v65 = int32(0)
													} else {
														v65 = v14
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)) = uint8(v74)
													if v74 != 0 {
														v82 = v75
													} else {
														v82 = v13
													}
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(24), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int32(0)
													} else {
														m.G0 = v11 + int32(32)
														return v90
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
}
func F_range_contains_elem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				v31 = F_range_contains_elem_internal(m, v30, v12, v16)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					m.G0 = v9 + int32(16)
					return v31
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(369887), v9)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493562), int32(1776), int32(398373))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = F_range_contains_elem_internal(m, v30, v12, v16)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v31
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v17, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(369887), v9)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(493562), int32(1776), int32(398373))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = F_range_contains_elem_internal(m, v30, v12, v16)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v31
					}
				}
			}
		}
	}
}
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	v6 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v18 = l1 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+11)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+10)))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
	v23 = base.I32_extend16_s(v22)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v24)>>(uint(int32(2))%32))-int32(1)))))
	if v30&int32(41) != 0 {
		v91 = v18
		v92 = int32(0)
		if v30&int32(81) != 0 {
			v157 = v92
			v158 = v6
			v160 = v157
			v163 = v158
			v164 = int32(1)
			v165 = v30 & v164
			*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
			v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
			v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
			*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
			v180 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
			v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
			v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
			*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
			*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
			m.G0 = v15 + int32(32)
			return
		} else {
			if v23 == int32(-1) {
				v105 = v91
				v107 = v92
				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
				if v108 != 0 {
					v128 = v105
					v129 = int32(-1)
					v130 = v107
				} else {
					v110 = v105
					v111 = int32(-1)
					v112 = v107
					switch v20 - int32(99) {
					case 0:
						v128 = v110
						v129 = v111
						v130 = v112
					case 1:
						v128 = (v110 + int32(7)) & int32(-8)
						v129 = v111
						v130 = v112
					default:
						v128 = (v110 + int32(1)) & int32(-2)
						v129 = v111
						v130 = v112
					case 6:
						v128 = (v110 + int32(3)) & int32(-4)
						v129 = v111
						v130 = v112
					}
				}
			} else {
				v110 = v91
				v111 = v23
				v112 = v92
				switch v20 - int32(99) {
				case 0:
					v128 = v110
					v129 = v111
					v130 = v112
				case 1:
					v128 = (v110 + int32(7)) & int32(-8)
					v129 = v111
					v130 = v112
				default:
					v128 = (v110 + int32(1)) & int32(-2)
					v129 = v111
					v130 = v112
				case 6:
					v128 = (v110 + int32(3)) & int32(-4)
					v129 = v111
					v130 = v112
				}
			}
			if v21&int32(1) == int32(0) {
				v160 = v130
				v163 = v128
				v164 = int32(1)
				v165 = v30 & v164
				*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
				v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
				v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
				*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
				*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
				v180 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
				v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
				v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
				*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
				*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
				m.G0 = v15 + int32(32)
				return
			} else {
				switch v129 - int32(1) {
				case 0:
					v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
					v157 = v130
					v158 = v154
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				case 1:
					v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
					v157 = v130
					v158 = v137
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v142 = m.ExcPending
					if v142 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
						F_errmsg_internal(m, int32(482718), v15+int32(16))
						mBase = m.M
						v148 = m.ExcPending
						if v148 != 0 {
							return
						} else {
							F_errfinish(m, int32(326157), int32(70), int32(67716))
							mBase = m.M
							v153 = m.ExcPending
							if v153 != 0 {
								return
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
					v157 = v130
					v158 = v138
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				}
			}
		}
	} else {
		if v21&int32(1) != 0 {
			switch v22 - int32(1) {
			case 0:
				v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v18))))
				v91 = v18 + v23
				v92 = v38
				if v30&int32(81) != 0 {
					v157 = v92
					v158 = v6
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v105 = v91
						v107 = v92
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
						if v108 != 0 {
							v128 = v105
							v129 = int32(-1)
							v130 = v107
						} else {
							v110 = v105
							v111 = int32(-1)
							v112 = v107
							switch v20 - int32(99) {
							case 0:
								v128 = v110
								v129 = v111
								v130 = v112
							case 1:
								v128 = (v110 + int32(7)) & int32(-8)
								v129 = v111
								v130 = v112
							default:
								v128 = (v110 + int32(1)) & int32(-2)
								v129 = v111
								v130 = v112
							case 6:
								v128 = (v110 + int32(3)) & int32(-4)
								v129 = v111
								v130 = v112
							}
						}
					} else {
						v110 = v91
						v111 = v23
						v112 = v92
						switch v20 - int32(99) {
						case 0:
							v128 = v110
							v129 = v111
							v130 = v112
						case 1:
							v128 = (v110 + int32(7)) & int32(-8)
							v129 = v111
							v130 = v112
						default:
							v128 = (v110 + int32(1)) & int32(-2)
							v129 = v111
							v130 = v112
						case 6:
							v128 = (v110 + int32(3)) & int32(-4)
							v129 = v111
							v130 = v112
						}
					}
					if v21&int32(1) == int32(0) {
						v160 = v130
						v163 = v128
						v164 = int32(1)
						v165 = v30 & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
						v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
						v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
						v180 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
						v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
						v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v129 - int32(1) {
						case 0:
							v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v154
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						case 1:
							v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v137
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
								F_errmsg_internal(m, int32(482718), v15+int32(16))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_errfinish(m, int32(326157), int32(70), int32(67716))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
							v157 = v130
							v158 = v138
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			case 1:
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v18))))
				v91 = v18 + v23
				v92 = v40
				if v30&int32(81) != 0 {
					v157 = v92
					v158 = v6
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v105 = v91
						v107 = v92
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
						if v108 != 0 {
							v128 = v105
							v129 = int32(-1)
							v130 = v107
						} else {
							v110 = v105
							v111 = int32(-1)
							v112 = v107
							switch v20 - int32(99) {
							case 0:
								v128 = v110
								v129 = v111
								v130 = v112
							case 1:
								v128 = (v110 + int32(7)) & int32(-8)
								v129 = v111
								v130 = v112
							default:
								v128 = (v110 + int32(1)) & int32(-2)
								v129 = v111
								v130 = v112
							case 6:
								v128 = (v110 + int32(3)) & int32(-4)
								v129 = v111
								v130 = v112
							}
						}
					} else {
						v110 = v91
						v111 = v23
						v112 = v92
						switch v20 - int32(99) {
						case 0:
							v128 = v110
							v129 = v111
							v130 = v112
						case 1:
							v128 = (v110 + int32(7)) & int32(-8)
							v129 = v111
							v130 = v112
						default:
							v128 = (v110 + int32(1)) & int32(-2)
							v129 = v111
							v130 = v112
						case 6:
							v128 = (v110 + int32(3)) & int32(-4)
							v129 = v111
							v130 = v112
						}
					}
					if v21&int32(1) == int32(0) {
						v160 = v130
						v163 = v128
						v164 = int32(1)
						v165 = v30 & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
						v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
						v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
						v180 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
						v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
						v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v129 - int32(1) {
						case 0:
							v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v154
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						case 1:
							v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v137
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
								F_errmsg_internal(m, int32(482718), v15+int32(16))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_errfinish(m, int32(326157), int32(70), int32(67716))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
							v157 = v130
							v158 = v138
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v23
					F_errmsg_internal(m, int32(482718), v15)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						F_errfinish(m, int32(326157), int32(70), int32(67716))
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
			case 3:
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v91 = v18 + v23
				v92 = v42
				if v30&int32(81) != 0 {
					v157 = v92
					v158 = v6
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v105 = v91
						v107 = v92
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
						if v108 != 0 {
							v128 = v105
							v129 = int32(-1)
							v130 = v107
						} else {
							v110 = v105
							v111 = int32(-1)
							v112 = v107
							switch v20 - int32(99) {
							case 0:
								v128 = v110
								v129 = v111
								v130 = v112
							case 1:
								v128 = (v110 + int32(7)) & int32(-8)
								v129 = v111
								v130 = v112
							default:
								v128 = (v110 + int32(1)) & int32(-2)
								v129 = v111
								v130 = v112
							case 6:
								v128 = (v110 + int32(3)) & int32(-4)
								v129 = v111
								v130 = v112
							}
						}
					} else {
						v110 = v91
						v111 = v23
						v112 = v92
						switch v20 - int32(99) {
						case 0:
							v128 = v110
							v129 = v111
							v130 = v112
						case 1:
							v128 = (v110 + int32(7)) & int32(-8)
							v129 = v111
							v130 = v112
						default:
							v128 = (v110 + int32(1)) & int32(-2)
							v129 = v111
							v130 = v112
						case 6:
							v128 = (v110 + int32(3)) & int32(-4)
							v129 = v111
							v130 = v112
						}
					}
					if v21&int32(1) == int32(0) {
						v160 = v130
						v163 = v128
						v164 = int32(1)
						v165 = v30 & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
						v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
						v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
						v180 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
						v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
						v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v129 - int32(1) {
						case 0:
							v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v154
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						case 1:
							v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v137
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
								F_errmsg_internal(m, int32(482718), v15+int32(16))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_errfinish(m, int32(326157), int32(70), int32(67716))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
							v157 = v130
							v158 = v138
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			}
		} else {
			if int32(0) < v23 {
				v91 = v18 + v23
				v92 = v18
				if v30&int32(81) != 0 {
					v157 = v92
					v158 = v6
					v160 = v157
					v163 = v158
					v164 = int32(1)
					v165 = v30 & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
					v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
					v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
					v180 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
					v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
					v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
					*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
					*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
					m.G0 = v15 + int32(32)
					return
				} else {
					if v23 == int32(-1) {
						v105 = v91
						v107 = v92
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
						if v108 != 0 {
							v128 = v105
							v129 = int32(-1)
							v130 = v107
						} else {
							v110 = v105
							v111 = int32(-1)
							v112 = v107
							switch v20 - int32(99) {
							case 0:
								v128 = v110
								v129 = v111
								v130 = v112
							case 1:
								v128 = (v110 + int32(7)) & int32(-8)
								v129 = v111
								v130 = v112
							default:
								v128 = (v110 + int32(1)) & int32(-2)
								v129 = v111
								v130 = v112
							case 6:
								v128 = (v110 + int32(3)) & int32(-4)
								v129 = v111
								v130 = v112
							}
						}
					} else {
						v110 = v91
						v111 = v23
						v112 = v92
						switch v20 - int32(99) {
						case 0:
							v128 = v110
							v129 = v111
							v130 = v112
						case 1:
							v128 = (v110 + int32(7)) & int32(-8)
							v129 = v111
							v130 = v112
						default:
							v128 = (v110 + int32(1)) & int32(-2)
							v129 = v111
							v130 = v112
						case 6:
							v128 = (v110 + int32(3)) & int32(-4)
							v129 = v111
							v130 = v112
						}
					}
					if v21&int32(1) == int32(0) {
						v160 = v130
						v163 = v128
						v164 = int32(1)
						v165 = v30 & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
						v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
						v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
						v180 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
						v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
						v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
						m.G0 = v15 + int32(32)
						return
					} else {
						switch v129 - int32(1) {
						case 0:
							v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v154
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						case 1:
							v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
							v157 = v130
							v158 = v137
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						default:
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
								F_errmsg_internal(m, int32(482718), v15+int32(16))
								mBase = m.M
								v148 = m.ExcPending
								if v148 != 0 {
									return
								} else {
									F_errfinish(m, int32(326157), int32(70), int32(67716))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						case 3:
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
							v157 = v130
							v158 = v138
							v160 = v157
							v163 = v158
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						}
					}
				}
			} else {
				if v23 == int32(-1) {
					v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
					if v62 == int32(1) {
						v65 = int32(6)
						v67 = int32(18)
						v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
						if v69 == v67 {
							v72 = v67
						} else {
							v72 = int32(2)
						}
						if v69&int32(254) == int32(2) {
							v77 = v65
						} else {
							v77 = v72
						}
						if v69 == int32(1) {
							v80 = v65
						} else {
							v80 = v77
						}
						v101 = v80
					} else {
						if v62&int32(1) == int32(0) {
							v97 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
							v101 = int32(base.Ui32(v97) >> (uint(int32(2)) % 32))
						} else {
							v101 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
						}
					}
					if v30&int32(80) != 0 {
						v160 = v18
						v163 = v6
						v164 = int32(1)
						v165 = v30 & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
						v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
						v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
						v180 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
						v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
						v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
						m.G0 = v15 + int32(32)
						return
					} else {
						v105 = v101 + v18
						v107 = v18
						v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
						if v108 != 0 {
							v128 = v105
							v129 = int32(-1)
							v130 = v107
						} else {
							v110 = v105
							v111 = int32(-1)
							v112 = v107
							switch v20 - int32(99) {
							case 0:
								v128 = v110
								v129 = v111
								v130 = v112
							case 1:
								v128 = (v110 + int32(7)) & int32(-8)
								v129 = v111
								v130 = v112
							default:
								v128 = (v110 + int32(1)) & int32(-2)
								v129 = v111
								v130 = v112
							case 6:
								v128 = (v110 + int32(3)) & int32(-4)
								v129 = v111
								v130 = v112
							}
						}
						if v21&int32(1) == int32(0) {
							v160 = v130
							v163 = v128
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						} else {
							switch v129 - int32(1) {
							case 0:
								v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
								v157 = v130
								v158 = v154
								v160 = v157
								v163 = v158
								v164 = int32(1)
								v165 = v30 & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
								v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
								v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
								v180 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
								v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
								v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
								m.G0 = v15 + int32(32)
								return
							case 1:
								v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
								v157 = v130
								v158 = v137
								v160 = v157
								v163 = v158
								v164 = int32(1)
								v165 = v30 & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
								v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
								v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
								v180 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
								v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
								v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
								m.G0 = v15 + int32(32)
								return
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
									F_errmsg_internal(m, int32(482718), v15+int32(16))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return
									} else {
										F_errfinish(m, int32(326157), int32(70), int32(67716))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
								v157 = v130
								v158 = v138
								v160 = v157
								v163 = v158
								v164 = int32(1)
								v165 = v30 & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
								v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
								v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
								v180 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
								v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
								v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
								m.G0 = v15 + int32(32)
								return
							}
						}
					}
				} else {
					v87 = F_strlen(m, v18)
					mBase = m.M
					v91 = v87 + v18 + int32(1)
					v92 = v18
					if v30&int32(81) != 0 {
						v157 = v92
						v158 = v6
						v160 = v157
						v163 = v158
						v164 = int32(1)
						v165 = v30 & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
						v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
						v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
						v180 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
						v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
						v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
						*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
						m.G0 = v15 + int32(32)
						return
					} else {
						if v23 == int32(-1) {
							v105 = v91
							v107 = v92
							v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
							if v108 != 0 {
								v128 = v105
								v129 = int32(-1)
								v130 = v107
							} else {
								v110 = v105
								v111 = int32(-1)
								v112 = v107
								switch v20 - int32(99) {
								case 0:
									v128 = v110
									v129 = v111
									v130 = v112
								case 1:
									v128 = (v110 + int32(7)) & int32(-8)
									v129 = v111
									v130 = v112
								default:
									v128 = (v110 + int32(1)) & int32(-2)
									v129 = v111
									v130 = v112
								case 6:
									v128 = (v110 + int32(3)) & int32(-4)
									v129 = v111
									v130 = v112
								}
							}
						} else {
							v110 = v91
							v111 = v23
							v112 = v92
							switch v20 - int32(99) {
							case 0:
								v128 = v110
								v129 = v111
								v130 = v112
							case 1:
								v128 = (v110 + int32(7)) & int32(-8)
								v129 = v111
								v130 = v112
							default:
								v128 = (v110 + int32(1)) & int32(-2)
								v129 = v111
								v130 = v112
							case 6:
								v128 = (v110 + int32(3)) & int32(-4)
								v129 = v111
								v130 = v112
							}
						}
						if v21&int32(1) == int32(0) {
							v160 = v130
							v163 = v128
							v164 = int32(1)
							v165 = v30 & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
							v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
							v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
							v180 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
							v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
							v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
							*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
							m.G0 = v15 + int32(32)
							return
						} else {
							switch v129 - int32(1) {
							case 0:
								v154 = int32(*(*int8)(unsafe.Add(mBase, uint32(v128))))
								v157 = v130
								v158 = v154
								v160 = v157
								v163 = v158
								v164 = int32(1)
								v165 = v30 & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
								v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
								v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
								v180 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
								v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
								v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
								m.G0 = v15 + int32(32)
								return
							case 1:
								v137 = int32(*(*int16)(unsafe.Add(mBase, uint32(v128))))
								v157 = v130
								v158 = v137
								v160 = v157
								v163 = v158
								v164 = int32(1)
								v165 = v30 & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
								v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
								v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
								v180 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
								v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
								v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
								m.G0 = v15 + int32(32)
								return
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v142 = m.ExcPending
								if v142 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v129
									F_errmsg_internal(m, int32(482718), v15+int32(16))
									mBase = m.M
									v148 = m.ExcPending
									if v148 != 0 {
										return
									} else {
										F_errfinish(m, int32(326157), int32(70), int32(67716))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							case 3:
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v128)))
								v157 = v130
								v158 = v138
								v160 = v157
								v163 = v158
								v164 = int32(1)
								v165 = v30 & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v165)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+6)) = uint8(v164)
								v172 = int32(base.Ui32(v30)>>(uint(v164)%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v172)
								v177 = int32(base.Ui32(v30)>>(uint(int32(3))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v177)
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
								v180 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+6)) = uint8(v180)
								v185 = int32(base.Ui32(v30)>>(uint(int32(2))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+5)) = uint8(v185)
								v190 = int32(base.Ui32(v30)>>(uint(int32(4))%32)) & v164
								*(*uint8)(unsafe.Add(mBase, uint32(l3)+4)) = uint8(v190)
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v163
								m.G0 = v15 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_range_empty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(int32(2))%32))-v11))))
		return v13 & v11
	}
}
func F_range_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v24 = F_lookup_type_cache(m, v22, int32(2048))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v27 = v19
	goto L6
L6:
	;
	F_range_deserialize(m, v27, v13, v11+int32(40), v11+int32(24), v11+int32(15))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v24
	v27 = v24
	goto L6
L8:
	;
	F_range_deserialize(m, v27, v17, v11+int32(32), v11+int32(16), v11+int32(14))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v44 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	if v46 == v44 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if l0 != v13 {
		goto L98
	} else {
		goto L99
	}
L11:
	;
	v182 = int32(0) - (v45 ^ int32(1))
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v45&int32(1) != 0 {
		v182 = v44
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
	if v56 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
	if v117 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L16:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+46)))
	if v55&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v55&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v59 == v62 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v70 = int32(1)
	if v59&v70 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v65 = int32(1)
	if v59&v65 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v68 = int32(-1)
	goto L25
L24:
	;
	v68 = v65
	goto L25
L25:
	;
	v182 = v68
	goto L10
L26:
	;
	v73 = int32(-1)
	goto L28
L27:
	;
	v73 = v70
	goto L28
L28:
	;
	v182 = v73
	goto L10
L29:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v78 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
	v85 = F_FunctionCall2Coll(m, v27+int32(212), v82, v83, v84)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v79 = int32(1)
	goto L34
L33:
	;
	v79 = int32(-1)
	goto L34
L34:
	;
	v182 = v79
	goto L10
L35:
	;
	if v85 != 0 {
		v182 = v85
		goto L10
	} else {
		goto L36
	}
L36:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+37)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+45)))
	if v88 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+46)))
	if v87&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v87&int32(1) != 0 {
		goto L15
	} else {
		goto L50
	}
L40:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v91 == v96 {
		goto L15
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v103 = int32(1)
	if v91&v103 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v98 = int32(1)
	if v91&v98 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v102 = v98
	goto L46
L45:
	;
	v102 = int32(-1)
	goto L46
L46:
	;
	v182 = v102
	goto L10
L47:
	;
	v107 = v103
	goto L49
L48:
	;
	v107 = int32(-1)
	goto L49
L49:
	;
	v182 = v107
	goto L10
L50:
	;
	v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+38)))
	if v112 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v113 = int32(-1)
	goto L53
L52:
	;
	v113 = int32(1)
	goto L53
L53:
	;
	v182 = v113
	goto L10
L54:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)))
	if v116&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v116&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v123 == v120&int32(255) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v134 = int32(1)
	if v120&v134 != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v182 = int32(0)
	goto L10
L61:
	;
	goto L62
L62:
	;
	v129 = int32(1)
	if v120&v129 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v132 = int32(-1)
	goto L65
L64:
	;
	v132 = v129
	goto L65
L65:
	;
	v182 = v132
	goto L10
L66:
	;
	v137 = int32(-1)
	goto L68
L67:
	;
	v137 = v134
	goto L68
L68:
	;
	v182 = v137
	goto L10
L69:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v142 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v149 = F_FunctionCall2Coll(m, v27+int32(212), v146, v147, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	v143 = int32(1)
	goto L74
L73:
	;
	v143 = int32(-1)
	goto L74
L74:
	;
	v182 = v143
	goto L10
L75:
	;
	if v149 != 0 {
		v182 = v149
		goto L10
	} else {
		goto L76
	}
L76:
	;
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+21)))
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+29)))
	if v152 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+30)))
	if v151&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if v151&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v160 == v155&int32(255) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v170 = int32(1)
	if v155&v170 != 0 {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v182 = int32(0)
	goto L10
L84:
	;
	goto L85
L85:
	;
	v165 = int32(1)
	if v155&v165 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v169 = v165
	goto L88
L87:
	;
	v169 = int32(-1)
	goto L88
L88:
	;
	v182 = v169
	goto L10
L89:
	;
	v174 = v170
	goto L91
L90:
	;
	v174 = int32(-1)
	goto L91
L91:
	;
	v182 = v174
	goto L10
L92:
	;
	v182 = int32(0)
	goto L10
L93:
	;
	goto L94
L94:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+22)))
	if v180 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v181 = int32(-1)
	goto L97
L96:
	;
	v181 = int32(1)
	goto L97
L97:
	;
	v182 = v181
	goto L10
L98:
	;
	F_pfree(m, v13)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if l1 != v17 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	F_pfree(m, v17)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	m.G0 = v11 + int32(48)
	return v182
L105:
	;
	goto L104
}
func F_range_gist_penalty(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v92 float32
	_ = v92
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 float32
	_ = v105
	var v108 float32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 float32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v139 float64
	_ = v139
	var v141 int32
	_ = v141
	var v144 float32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 float32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v176 float64
	_ = v176
	var v178 float32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v207 float64
	_ = v207
	var v210 float64
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 float64
	_ = v226
	var v227 float64
	_ = v227
	var v230 float64
	_ = v230
	var v234 float64
	_ = v234
	var v239 float32
	_ = v239
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v22 = F_pg_detoast_datum(m, v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			if v24 == v25 {
				v27 = F_range_get_typcache(m, l0, v24)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
					F_range_deserialize(m, v27, v17, v11+int32(40), v11+int32(24), v11+int32(15))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F_range_deserialize(m, v27, v22, v11+int32(32), v11+int32(16), v11+int32(14))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+14)))
							if v46 == int32(1) {
								v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
								if v49 != 0 {
									v239 = float32(0)
								} else {
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
									v56 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v50)>>(uint(int32(2))%32))-int32(1)))))
									if v56&int32(-127) != 0 {
										v239 = float32(1)
									} else {
										v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
										v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
										if v61 == int32(1) {
											if v60&int32(1) == int32(0) {
												v239 = float32(3)
											} else {
												v239 = float32(2)
											}
										} else {
											if v60&int32(1) != 0 {
												v239 = float32(3)
											} else {
												v239 = float32(4)
											}
										}
									}
								}
								*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
								m.G0 = v11 + int32(48)
								return v13
							} else {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+20)))
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+36)))
								if v74 == int32(1) {
									if v73&int32(1) != 0 {
										v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
										v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
										if v80 == int32(1) {
											if v79&int32(1) == int32(0) {
												v92 = float32(2)
											} else {
												v92 = float32(0)
											}
										} else {
											if v79&int32(1) != 0 {
												v92 = float32(2)
											} else {
												v92 = float32(4)
											}
										}
										*(*float32)(unsafe.Add(mBase, uint32(v13))) = v92
										v94 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
										v100 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17+int32(base.Ui32(v94)>>(uint(int32(2))%32))-int32(1)))))
										if v100&int32(-127) == int32(0) {
										} else {
											v105 = *(*float32)(unsafe.Add(mBase, uint32(v13)))
											v239 = base.F32_add(v105, float32(1))
											*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
										}
										m.G0 = v11 + int32(48)
										return v13
									} else {
										v108 = math.Float32frombits(uint32(0x7f800000))
										v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
										if v109 != 0 {
											v239 = v108
											*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
											m.G0 = v11 + int32(48)
											return v13
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
											if v110 != int32(1) {
												v239 = v108
												*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
												m.G0 = v11 + int32(48)
												return v13
											} else {
												v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
												if v113 != 0 {
													v239 = float32(0)
													*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
													m.G0 = v11 + int32(48)
													return v13
												} else {
													v121 = F_range_cmp_bounds(m, v27, v11+int32(16), v11+int32(24))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														v124 = base.B2i32(v121 <= int32(0))
														if v121 <= int32(0) {
															v125 = float32(0)
														} else {
															v125 = float32(1)
														}
														if v121 <= int32(0) {
															v239 = v125
															*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
															m.G0 = v11 + int32(48)
															return v13
														} else {
															if v29 == int32(0) {
																v239 = v125
																*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																m.G0 = v11 + int32(48)
																return v13
															} else {
																v130 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
																v131 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																v132 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
																v133 = F_FunctionCall2Coll(m, v27+int32(268), v130, v131, v132)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int32(0)
																} else {
																	v135 = *(*float64)(unsafe.Add(mBase, uint32(v133)))
																	v136 = float64(0)
																	if base.F64_ge(v135, v136) != 0 {
																		v139 = v135
																	} else {
																		v139 = v136
																	}
																	v239 = base.F32_demote_f64(v139)
																	*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																	m.G0 = v11 + int32(48)
																	return v13
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
									if v73&int32(1) != 0 {
										v144 = math.Float32frombits(uint32(0x7f800000))
										if v141&int32(1) != 0 {
											v239 = v144
											*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
											m.G0 = v11 + int32(48)
											return v13
										} else {
											v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
											if v147 != int32(1) {
												v239 = v144
												*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
												m.G0 = v11 + int32(48)
												return v13
											} else {
												v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
												if v150 != 0 {
													v239 = float32(0)
													*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
													m.G0 = v11 + int32(48)
													return v13
												} else {
													v158 = F_range_cmp_bounds(m, v27, v11+int32(32), v11+int32(40))
													mBase = m.M
													v159 = m.ExcPending
													if v159 != 0 {
														return int32(0)
													} else {
														v161 = base.B2i32(int32(0) <= v158)
														if int32(0) <= v158 {
															v162 = float32(0)
														} else {
															v162 = float32(1)
														}
														if int32(0) <= v158 {
															v239 = v162
															*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
															m.G0 = v11 + int32(48)
															return v13
														} else {
															if v29 == int32(0) {
																v239 = v162
																*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																m.G0 = v11 + int32(48)
																return v13
															} else {
																v167 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
																v168 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
																v169 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
																v170 = F_FunctionCall2Coll(m, v27+int32(268), v167, v168, v169)
																mBase = m.M
																v171 = m.ExcPending
																if v171 != 0 {
																	return int32(0)
																} else {
																	v172 = *(*float64)(unsafe.Add(mBase, uint32(v170)))
																	v173 = float64(0)
																	if base.F64_ge(v172, v173) != 0 {
																		v176 = v172
																	} else {
																		v176 = v173
																	}
																	v239 = base.F32_demote_f64(v176)
																	*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																	m.G0 = v11 + int32(48)
																	return v13
																}
															}
														}
													}
												}
											}
										}
									} else {
										v178 = math.Float32frombits(uint32(0x7f800000))
										if v141&int32(1) != 0 {
											v239 = v178
											*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
											m.G0 = v11 + int32(48)
											return v13
										} else {
											v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+44)))
											if v181 != 0 {
												v239 = v178
												*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
												m.G0 = v11 + int32(48)
												return v13
											} else {
												v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+28)))
												if v182 != 0 {
													v239 = v178
													*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
													m.G0 = v11 + int32(48)
													return v13
												} else {
													v189 = F_range_cmp_bounds(m, v27, v11+int32(32), v11+int32(40))
													mBase = m.M
													v190 = m.ExcPending
													if v190 != 0 {
														return int32(0)
													} else {
														v192 = base.B2i32(int32(0) <= v189)
														if int32(0) <= v189 {
															v193 = float64(0)
														} else {
															v193 = float64(1)
														}
														if int32(0) <= v189 {
															v210 = v193
															v215 = F_range_cmp_bounds(m, v27, v11+int32(16), v11+int32(24))
															mBase = m.M
															v216 = m.ExcPending
															if v216 != 0 {
																return int32(0)
															} else {
																if v215 <= int32(0) {
																	v234 = v210
																	v239 = base.F32_demote_f64(v234)
																	*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																	m.G0 = v11 + int32(48)
																	return v13
																} else {
																	if v29 != 0 {
																		v221 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
																		v222 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																		v223 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
																		v224 = F_FunctionCall2Coll(m, v27+int32(268), v221, v222, v223)
																		mBase = m.M
																		v225 = m.ExcPending
																		if v225 != 0 {
																			return int32(0)
																		} else {
																			v226 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
																			v227 = float64(0)
																			if base.F64_ge(v226, v227) != 0 {
																				v230 = v226
																			} else {
																				v230 = v227
																			}
																			v234 = base.F64_add(v210, v230)
																			v239 = base.F32_demote_f64(v234)
																			*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																			m.G0 = v11 + int32(48)
																			return v13
																		}
																	} else {
																		v234 = base.F64_add(v210, float64(1))
																		v239 = base.F32_demote_f64(v234)
																		*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																		m.G0 = v11 + int32(48)
																		return v13
																	}
																}
															}
														} else {
															if v29 == int32(0) {
																v210 = v193
																v215 = F_range_cmp_bounds(m, v27, v11+int32(16), v11+int32(24))
																mBase = m.M
																v216 = m.ExcPending
																if v216 != 0 {
																	return int32(0)
																} else {
																	if v215 <= int32(0) {
																		v234 = v210
																		v239 = base.F32_demote_f64(v234)
																		*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																		m.G0 = v11 + int32(48)
																		return v13
																	} else {
																		if v29 != 0 {
																			v221 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
																			v222 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																			v223 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
																			v224 = F_FunctionCall2Coll(m, v27+int32(268), v221, v222, v223)
																			mBase = m.M
																			v225 = m.ExcPending
																			if v225 != 0 {
																				return int32(0)
																			} else {
																				v226 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
																				v227 = float64(0)
																				if base.F64_ge(v226, v227) != 0 {
																					v230 = v226
																				} else {
																					v230 = v227
																				}
																				v234 = base.F64_add(v210, v230)
																				v239 = base.F32_demote_f64(v234)
																				*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																				m.G0 = v11 + int32(48)
																				return v13
																			}
																		} else {
																			v234 = base.F64_add(v210, float64(1))
																			v239 = base.F32_demote_f64(v234)
																			*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																			m.G0 = v11 + int32(48)
																			return v13
																		}
																	}
																}
															} else {
																v198 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
																v199 = *(*int32)(unsafe.Add(mBase, uint32(v11)+40))
																v200 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
																v201 = F_FunctionCall2Coll(m, v27+int32(268), v198, v199, v200)
																mBase = m.M
																v202 = m.ExcPending
																if v202 != 0 {
																	return int32(0)
																} else {
																	v203 = *(*float64)(unsafe.Add(mBase, uint32(v201)))
																	v204 = float64(0)
																	if base.F64_ge(v203, v204) != 0 {
																		v207 = v203
																	} else {
																		v207 = v204
																	}
																	v210 = base.F64_add(v207, float64(0))
																	v215 = F_range_cmp_bounds(m, v27, v11+int32(16), v11+int32(24))
																	mBase = m.M
																	v216 = m.ExcPending
																	if v216 != 0 {
																		return int32(0)
																	} else {
																		if v215 <= int32(0) {
																			v234 = v210
																			v239 = base.F32_demote_f64(v234)
																			*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																			m.G0 = v11 + int32(48)
																			return v13
																		} else {
																			if v29 != 0 {
																				v221 = *(*int32)(unsafe.Add(mBase, uint32(v27)+208))
																				v222 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
																				v223 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
																				v224 = F_FunctionCall2Coll(m, v27+int32(268), v221, v222, v223)
																				mBase = m.M
																				v225 = m.ExcPending
																				if v225 != 0 {
																					return int32(0)
																				} else {
																					v226 = *(*float64)(unsafe.Add(mBase, uint32(v224)))
																					v227 = float64(0)
																					if base.F64_ge(v226, v227) != 0 {
																						v230 = v226
																					} else {
																						v230 = v227
																					}
																					v234 = base.F64_add(v210, v230)
																					v239 = base.F32_demote_f64(v234)
																					*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																					m.G0 = v11 + int32(48)
																					return v13
																				}
																			} else {
																				v234 = base.F64_add(v210, float64(1))
																				v239 = base.F32_demote_f64(v234)
																				*(*float32)(unsafe.Add(mBase, uint32(v13))) = v239
																				m.G0 = v11 + int32(48)
																				return v13
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
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v254 = m.ExcPending
				if v254 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(325081), int32(0))
					mBase = m.M
					v258 = m.ExcPending
					if v258 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(492067), int32(379), int32(9118))
						mBase = m.M
						v263 = m.ExcPending
						if v263 != 0 {
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
func F_range_gist_picksplit(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v27 float32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int64
	_ = v64
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v362 float32
	_ = v362
	var v363 int32
	_ = v363
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v395 float32
	_ = v395
	var v398 float32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v458 int32
	_ = v458
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v501 int32
	_ = v501
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v531 float32
	_ = v531
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 float64
	_ = v542
	var v543 float64
	_ = v543
	var v546 float64
	_ = v546
	var v551 float32
	_ = v551
	var v558 int32
	_ = v558
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 float32
	_ = v571
	var v573 float32
	_ = v573
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v594 int32
	_ = v594
	var v601 float32
	_ = v601
	var v604 float32
	_ = v604
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v640 float32
	_ = v640
	var v643 float32
	_ = v643
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v698 int32
	_ = v698
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v748 int32
	_ = v748
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v782 float32
	_ = v782
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 float64
	_ = v793
	var v794 float64
	_ = v794
	var v797 float64
	_ = v797
	var v802 float32
	_ = v802
	var v809 int32
	_ = v809
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v822 float32
	_ = v822
	var v824 float32
	_ = v824
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v921 int32
	_ = v921
	var v924 int32
	_ = v924
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 float64
	_ = v941
	var v942 float64
	_ = v942
	var v945 float64
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 float64
	_ = v951
	var v952 float64
	_ = v952
	var v955 float64
	_ = v955
	var v959 float64
	_ = v959
	var v963 int32
	_ = v963
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v979 int32
	_ = v979
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v990 int32
	_ = v990
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1069 int32
	_ = v1069
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1133 int32
	_ = v1133
	var v1136 int32
	_ = v1136
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1141 int64
	_ = v1141
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1187 int64
	_ = v1187
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1209 int32
	_ = v1209
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1226 int32
	_ = v1226
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1253 int32
	_ = v1253
	var v1259 int32
	_ = v1259
	var v1262 int32
	_ = v1262
	var v1265 int32
	_ = v1265
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1276 int32
	_ = v1276
	var v1279 int32
	_ = v1279
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1284 int32
	_ = v1284
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1306 int32
	_ = v1306
	var v1312 int32
	_ = v1312
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1317 int32
	_ = v1317
	var v1323 int32
	_ = v1323
	var v1330 int32
	_ = v1330
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1394 int32
	_ = v1394
	v2 = int32(0)
	v27 = float32(0)
	v32 = m.G0
	v34 = v32 - int32(96)
	m.G0 = v34
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+20))
	v39 = F_pg_detoast_datum(m, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v44 = F_range_get_typcache(m, l0, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v46 = int32(1)
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v51 = (v47 - v46) & int32(65535)
	v55 = v51<<(uint(v46)%32) + int32(2)
	v56 = F_palloc(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v56
	v59 = F_palloc(m, v55)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v59
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+80)) = v62
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+72)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v34-int32(-64)))) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v34)+56)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v34)+48)) = v64
	v75 = v37 + int32(4)
	if v47&int32(65535) != int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v81 = v46
	goto L9
L7:
	;
	v163 = v62
	v164 = v2
	v166 = v2
	v170 = v2
	v171 = v2
	v172 = v2
	v173 = v2
	v175 = v2
	v183 = v2
	goto L8
L8:
	;
	v202 = int32(0)
	v204 = base.B2i32(v163 <= v202)
	if v163 <= v202 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v75+v81<<(uint(int32(4))%32))))
	v119 = F_pg_detoast_datum(m, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v34)+80))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v34)+76))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v34)+72))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v34)+68))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v34)+60))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v34)+56))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v34)+64))
	v163 = v160
	v164 = v161
	v166 = v157
	v170 = v154
	v171 = v155
	v172 = v158
	v173 = v159
	v175 = v156
	v183 = v153
	goto L8
L11:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
	v127 = int32(*(*int8)(unsafe.Add(mBase, uint32(v119+int32(base.Ui32(v121)>>(uint(int32(2))%32))-int32(1)))))
	goto L12
L12:
	;
	if v127&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v140 = int32(8)
	goto L15
L14:
	;
	v130 = int32(3)
	v133 = int32(base.Ui32(v127)>>(uint(v130)%32)) & v130
	if v127 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v143 = v34 + int32(48) + v140<<(uint(int32(2))%32)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v145 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v144 + v145
	v151 = (v81 + v145) & int32(65535)
	if base.Ui32(v151) <= base.Ui32(v51) {
		v81 = v151
		goto L9
	} else {
		goto L19
	}
L16:
	;
	v138 = v133 | int32(4)
	goto L18
L17:
	;
	v138 = v133
	goto L18
L18:
	;
	v140 = v138
	goto L15
L19:
	;
	goto L10
L20:
	;
	v205 = int32(-1)
	goto L22
L21:
	;
	v205 = v202
	goto L22
L22:
	;
	v206 = int32(0)
	v208 = base.B2i32(v206 < v163)
	if v206 < v163 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v209 = v163
	goto L25
L24:
	;
	v209 = v206
	goto L25
L25:
	;
	v210 = base.B2i32(v209 < v173)
	if v209 < v173 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v211 = int32(1)
	goto L28
L27:
	;
	v211 = v205
	goto L28
L28:
	;
	v213 = base.B2i32(int32(0) < v173)
	if int32(0) < v173 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v214 = v211
	goto L31
L30:
	;
	v214 = v205
	goto L31
L31:
	;
	if v209 < v173 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v215 = v173
	goto L34
L33:
	;
	v215 = v209
	goto L34
L34:
	;
	if int32(0) < v173 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v216 = v215
	goto L37
L36:
	;
	v216 = v209
	goto L37
L37:
	;
	v217 = base.B2i32(v216 < v172)
	if v216 < v172 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v218 = int32(2)
	goto L40
L39:
	;
	v218 = v214
	goto L40
L40:
	;
	v220 = base.B2i32(int32(0) < v172)
	if int32(0) < v172 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v221 = v218
	goto L43
L42:
	;
	v221 = v214
	goto L43
L43:
	;
	if v216 < v172 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v222 = v172
	goto L46
L45:
	;
	v222 = v216
	goto L46
L46:
	;
	if int32(0) < v172 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v223 = v222
	goto L49
L48:
	;
	v223 = v216
	goto L49
L49:
	;
	v224 = base.B2i32(v223 < v166)
	if v223 < v166 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v225 = int32(3)
	goto L52
L51:
	;
	v225 = v221
	goto L52
L52:
	;
	v227 = base.B2i32(int32(0) < v166)
	if int32(0) < v166 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v228 = v225
	goto L55
L54:
	;
	v228 = v221
	goto L55
L55:
	;
	if v223 < v166 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v229 = v166
	goto L58
L57:
	;
	v229 = v223
	goto L58
L58:
	;
	if int32(0) < v166 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v230 = v229
	goto L61
L60:
	;
	v230 = v223
	goto L61
L61:
	;
	v231 = base.B2i32(v230 < v164)
	if v230 < v164 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v232 = int32(4)
	goto L64
L63:
	;
	v232 = v228
	goto L64
L64:
	;
	v234 = base.B2i32(int32(0) < v164)
	if int32(0) < v164 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v235 = v232
	goto L67
L66:
	;
	v235 = v228
	goto L67
L67:
	;
	if v230 < v164 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v236 = v164
	goto L70
L69:
	;
	v236 = v230
	goto L70
L70:
	;
	if int32(0) < v164 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v237 = v236
	goto L73
L72:
	;
	v237 = v230
	goto L73
L73:
	;
	v238 = base.B2i32(v237 < v175)
	if v237 < v175 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v239 = int32(5)
	goto L76
L75:
	;
	v239 = v235
	goto L76
L76:
	;
	v241 = base.B2i32(int32(0) < v175)
	if int32(0) < v175 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v242 = v239
	goto L79
L78:
	;
	v242 = v235
	goto L79
L79:
	;
	if v237 < v175 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v243 = v175
	goto L82
L81:
	;
	v243 = v237
	goto L82
L82:
	;
	if int32(0) < v175 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v244 = v243
	goto L85
L84:
	;
	v244 = v237
	goto L85
L85:
	;
	v245 = base.B2i32(v244 < v171)
	if v244 < v171 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v246 = int32(6)
	goto L88
L87:
	;
	v246 = v242
	goto L88
L88:
	;
	v248 = base.B2i32(int32(0) < v171)
	if int32(0) < v171 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v249 = v246
	goto L91
L90:
	;
	v249 = v242
	goto L91
L91:
	;
	if v244 < v171 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v250 = v171
	goto L94
L93:
	;
	v250 = v244
	goto L94
L94:
	;
	if int32(0) < v171 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v251 = v250
	goto L97
L96:
	;
	v251 = v244
	goto L97
L97:
	;
	v252 = base.B2i32(v251 < v170)
	if v251 < v170 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v253 = int32(7)
	goto L100
L99:
	;
	v253 = v249
	goto L100
L100:
	;
	v255 = base.B2i32(int32(0) < v170)
	if int32(0) < v170 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v256 = v253
	goto L103
L102:
	;
	v256 = v249
	goto L103
L103:
	;
	if v251 < v170 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v257 = v170
	goto L106
L105:
	;
	v257 = v251
	goto L106
L106:
	;
	if int32(0) < v170 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v258 = v257
	goto L109
L108:
	;
	v258 = v251
	goto L109
L109:
	;
	if v258 < v183 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v260 = int32(8)
	goto L112
L111:
	;
	v260 = v256
	goto L112
L112:
	;
	v262 = base.B2i32(int32(0) < v183)
	if int32(0) < v183 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v263 = v260
	goto L115
L114:
	;
	v263 = v256
	goto L115
L115:
	;
	if v213+v208+v220+v227+v234+v241+v248+v255+v262 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	m.G0 = v34 + int32(96)
	return v36
L117:
	;
	F_range_gist_fallback_split(m, v44, v37, v36)
	mBase = m.M
	v1394 = m.ExcPending
	if v1394 != 0 {
		goto L1
	} else {
		goto L343
	}
L118:
	;
	if v284 != 0 {
		goto L338
	} else {
		goto L339
	}
L119:
	;
	switch v263 & int32(-5) {
	case 0:
		goto L125
	case 1:
		goto L124
	case 2:
		goto L123
	default:
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v1139 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1139
	v1141 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+24)) = v1141
	*(*int64)(unsafe.Add(mBase, uint32(v34)+16)) = v1141
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v1141
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v1141
	if v204 == v1139 {
		goto L299
	} else {
		goto L300
	}
L122:
	;
	F_range_gist_fallback_split(m, v44, v37, v36)
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L1
	} else {
		goto L297
	}
L123:
	;
	F_range_gist_single_sorting_split(m, v44, v37, v36, int32(0))
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L1
	} else {
		goto L296
	}
L124:
	;
	F_range_gist_single_sorting_split(m, v44, v37, v36, int32(1))
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L1
	} else {
		goto L295
	}
L125:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v44)+272))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v279 = int32(65535)
	v280 = v278 + v279
	v282 = v280 & v279
	v284 = v282 << (uint(int32(4)) % 32)
	v285 = F_palloc(m, v284)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v287 = F_palloc(m, v284)
	mBase = m.M
	v288 = m.ExcPending
	if v288 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v278&int32(65535) == int32(1) {
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v294 = v282 - int32(1)
	v296 = int32(1)
	goto L129
L129:
	;
	v327 = v296 << (uint(int32(4)) % 32)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v75+v327)))
	v330 = F_pg_detoast_datum(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	if v284 != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	v332 = v327 + v285
	F_range_deserialize(m, v44, v330, v332-int32(16), v332-int32(8), v34)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v342 = (v296 + int32(1)) & int32(65535)
	if base.Ui32(v342) <= base.Ui32(v282) {
		v296 = v342
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	F_qsort_arg(m, v285, v282, int32(16), int32(1485), v44)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L138
	}
L135:
	;
	v344 = F__emscripten_memcpy_bulkmem(m, v287, v285, v284)
	mBase = m.M
	v345 = v344
	goto L137
L136:
	;
	v345 = v287
	goto L137
L137:
	;
	goto L134
L138:
	;
	F_qsort_arg(m, v345, v282, int32(16), int32(1486), v44)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	v355 = v44 + int32(268)
	v356 = int32(1)
	v357 = int32(base.Ui32(v282) >> (uint(v356) % 32))
	v361 = int32(base.Ui32(v282+v356) >> (uint(v356) % 32))
	v362 = base.F32_convert_i32_u(v282)
	v363 = int32(0)
	v370 = v345
	v371 = v363
	v376 = v285
	v378 = v363
	v381 = v356
	v382 = v363
	v386 = v363
	v388 = v363
	v395 = v27
	v398 = v27
	goto L140
L140:
	;
	v401 = v370
	v402 = v371
	goto L143
L141:
	;
	v607 = v294 << (uint(int32(4)) % 32)
	v609 = int32(8)
	v614 = v294
	v615 = v285 + v607 + v609
	v621 = v607 + v345 + v609
	v623 = v294
	v626 = v587
	v627 = v588
	v631 = v592
	v633 = v594
	v640 = v601
	v643 = v604
	goto L186
L142:
	;
	goto L141
L143:
	;
	v433 = v285 + v402<<(uint(int32(4))%32)
	v434 = F_range_cmp_bounds(m, v44, v376, v433)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L145
	}
L144:
	;
	if v282 <= v378 {
		v501 = v378
		goto L154
	} else {
		goto L155
	}
L145:
	;
	if v434 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v439 = v433 + int32(8)
	v440 = F_range_cmp_bounds(m, v44, v439, v401)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	goto L144
L149:
	;
	if int32(0) < v440 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v444 = v439
	goto L152
L151:
	;
	v444 = v401
	goto L152
L152:
	;
	v446 = v402 + int32(1)
	if v446 < v282 {
		v401 = v444
		v402 = v446
		goto L143
	} else {
		goto L153
	}
L153:
	;
	v587 = v381
	v588 = v382
	v592 = v386
	v594 = v388
	v601 = v395
	v604 = v398
	goto L142
L154:
	;
	if v501 < v357 {
		goto L162
	} else {
		goto L163
	}
L155:
	;
	v458 = v378
	goto L156
L156:
	;
	v485 = F_range_cmp_bounds(m, v44, v345+v458<<(uint(int32(4))%32)+int32(8), v401)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L158
	}
L157:
	;
	v501 = v282
	goto L154
L158:
	;
	if int32(0) < v485 {
		v501 = v458
		goto L154
	} else {
		goto L159
	}
L159:
	;
	v490 = v458 + int32(1)
	if v490 != v282 {
		v458 = v490
		goto L156
	} else {
		goto L160
	}
L160:
	;
	goto L157
L161:
	;
	if v402 < v282 {
		v370 = v401
		v371 = v402
		v376 = v433
		v378 = v501
		v381 = v566
		v382 = v567
		v386 = v568
		v388 = v569
		v395 = v571
		v398 = v573
		goto L140
	} else {
		goto L185
	}
L162:
	;
	v524 = v501
	goto L164
L163:
	;
	v524 = v357
	goto L164
L164:
	;
	if v402 < v361 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	v526 = v524
	goto L167
L166:
	;
	v526 = v402
	goto L167
L167:
	;
	v527 = v282 - v526
	if v526 < v527 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v529 = v526
	goto L170
L169:
	;
	v529 = v527
	goto L170
L170:
	;
	v531 = base.F32_div(base.F32_convert_i32_s(v529), v362)
	if base.F64_gt(base.F64_promote_f32(v531), float64(0.3)) == int32(0) {
		v566 = v381
		v567 = v382
		v568 = v386
		v569 = v388
		v571 = v395
		v573 = v398
		goto L161
	} else {
		goto L171
	}
L171:
	;
	if v277 != 0 {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if (v381|base.F32_lt(v551, v395))&int32(1) == int32(0) {
		goto L180
	} else {
		goto L181
	}
L173:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v401)))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v433)))
	v540 = F_FunctionCall2Coll(m, v355, v537, v538, v539)
	mBase = m.M
	v541 = m.ExcPending
	if v541 != 0 {
		goto L1
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v551 = base.F32_convert_i32_s(v501 - v402)
	goto L172
L176:
	;
	v542 = *(*float64)(unsafe.Add(mBase, uint32(v540)))
	v543 = float64(0)
	if base.F64_ge(v542, v543) != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v546 = v542
	goto L179
L178:
	;
	v546 = v543
	goto L179
L179:
	;
	v551 = base.F32_demote_f64(v546)
	goto L172
L180:
	;
	v558 = int32(0)
	if base.F32_ne(v395, v551) != 0 {
		v566 = v558
		v567 = v382
		v568 = v386
		v569 = v388
		v571 = v395
		v573 = v398
		goto L161
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v566 = int32(0)
	v567 = v401
	v568 = v433
	v569 = v501 - v526
	v571 = v551
	v573 = v531
	goto L161
L183:
	;
	if base.F32_gt(v531, v398) == int32(0) {
		v566 = v558
		v567 = v382
		v568 = v386
		v569 = v388
		v571 = v395
		v573 = v398
		goto L161
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v587 = v566
	v588 = v567
	v592 = v568
	v594 = v569
	v601 = v571
	v604 = v573
	goto L142
L186:
	;
	v645 = v614
	v646 = v615
	goto L189
L187:
	;
	if v839&int32(1) != 0 {
		goto L117
	} else {
		goto L234
	}
L188:
	;
	goto L187
L189:
	;
	v678 = v345 + v645<<(uint(int32(4))%32)
	v680 = v678 + int32(8)
	v681 = F_range_cmp_bounds(m, v44, v621, v680)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L191
	}
L190:
	;
	if v623 < int32(0) {
		v748 = v623
		goto L200
	} else {
		goto L201
	}
L191:
	;
	if v681 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v685 = F_range_cmp_bounds(m, v44, v678, v646)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	goto L190
L195:
	;
	if v685 < int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v689 = v678
	goto L198
L197:
	;
	v689 = v646
	goto L198
L198:
	;
	if int32(0) < v645 {
		v645 = v645 - int32(1)
		v646 = v689
		goto L189
	} else {
		goto L199
	}
L199:
	;
	v839 = v626
	v840 = v627
	v844 = v631
	v846 = v633
	goto L188
L200:
	;
	v771 = v645 + int32(1)
	if v771 < v357 {
		goto L210
	} else {
		goto L211
	}
L201:
	;
	v698 = v623
	goto L202
L202:
	;
	v730 = F_range_cmp_bounds(m, v44, v285+v698<<(uint(int32(4))%32), v646)
	mBase = m.M
	v731 = m.ExcPending
	if v731 != 0 {
		goto L1
	} else {
		goto L204
	}
L203:
	;
	v748 = int32(-1)
	goto L200
L204:
	;
	if v730 < int32(0) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v748 = v698
	goto L200
L206:
	;
	goto L207
L207:
	;
	if int32(0) < v698 {
		v698 = v698 - int32(1)
		goto L202
	} else {
		goto L208
	}
L208:
	;
	goto L203
L209:
	;
	if int32(0) <= v645 {
		v614 = v645
		v615 = v646
		v621 = v680
		v623 = v748
		v626 = v817
		v627 = v818
		v631 = v819
		v633 = v820
		v640 = v822
		v643 = v824
		goto L186
	} else {
		goto L233
	}
L210:
	;
	v773 = v771
	goto L212
L211:
	;
	v773 = v357
	goto L212
L212:
	;
	v775 = v748 + int32(1)
	if v775 < v361 {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	v777 = v773
	goto L215
L214:
	;
	v777 = v775
	goto L215
L215:
	;
	v778 = v282 - v777
	if v777 < v778 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v780 = v777
	goto L218
L217:
	;
	v780 = v778
	goto L218
L218:
	;
	v782 = base.F32_div(base.F32_convert_i32_s(v780), v362)
	if base.F64_gt(base.F64_promote_f32(v782), float64(0.3)) == int32(0) {
		v817 = v626
		v818 = v627
		v819 = v631
		v820 = v633
		v822 = v640
		v824 = v643
		goto L209
	} else {
		goto L219
	}
L219:
	;
	if v277 != 0 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	if (v626|base.F32_lt(v802, v640))&int32(1) == int32(0) {
		goto L228
	} else {
		goto L229
	}
L221:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v680)))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v646)))
	v791 = F_FunctionCall2Coll(m, v355, v788, v789, v790)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L1
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	v802 = base.F32_convert_i32_s(v645 - v748)
	goto L220
L224:
	;
	v793 = *(*float64)(unsafe.Add(mBase, uint32(v791)))
	v794 = float64(0)
	if base.F64_ge(v793, v794) != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v797 = v793
	goto L227
L226:
	;
	v797 = v794
	goto L227
L227:
	;
	v802 = base.F32_demote_f64(v797)
	goto L220
L228:
	;
	v809 = int32(0)
	if base.F32_ne(v640, v802) != 0 {
		v817 = v809
		v818 = v627
		v819 = v631
		v820 = v633
		v822 = v640
		v824 = v643
		goto L209
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	v817 = int32(0)
	v818 = v680
	v819 = v646
	v820 = v771 - v777
	v822 = v802
	v824 = v782
	goto L209
L231:
	;
	if base.F32_gt(v782, v643) == int32(0) {
		v817 = v809
		v818 = v627
		v819 = v631
		v820 = v633
		v822 = v640
		v824 = v643
		goto L209
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	v839 = v817
	v840 = v818
	v844 = v819
	v846 = v820
	goto L188
L234:
	;
	v861 = v282 << (uint(int32(1)) % 32)
	v862 = F_palloc(m, v861)
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v862
	v865 = F_palloc(m, v861)
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	v867 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v867
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v865
	v872 = F_palloc(m, v284)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v874 = int32(1)
	v876 = int32(0)
	v879 = v874
	v880 = v874
	v883 = v876
	v886 = v876
	v887 = v876
	goto L238
L238:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v75+v880<<(uint(int32(4))%32))))
	v914 = F_pg_detoast_datum(m, v913)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	if int32(0) < v997 {
		goto L273
	} else {
		goto L274
	}
L240:
	;
	F_range_deserialize(m, v44, v914, v34, v34+int32(88), v34+int32(87))
	mBase = m.M
	v921 = m.ExcPending
	if v921 != 0 {
		goto L1
	} else {
		goto L241
	}
L241:
	;
	v924 = F_range_cmp_bounds(m, v44, v34+int32(88), v840)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L1
	} else {
		goto L243
	}
L242:
	;
	v1002 = v879 + int32(1)
	v1003 = int32(65535)
	v1004 = v1002 & v1003
	if base.Ui32(v1004) <= base.Ui32(v280&v1003) {
		v879 = v1002
		v880 = v1004
		v883 = v997
		v886 = v998
		v887 = v999
		goto L238
	} else {
		goto L272
	}
L243:
	;
	if v924 <= int32(0) {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v928 = F_range_cmp_bounds(m, v44, v34, v844)
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L1
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	v979 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v979 <= int32(0) {
		goto L268
	} else {
		goto L269
	}
L247:
	;
	if int32(0) <= v928 {
		goto L248
	} else {
		goto L249
	}
L248:
	;
	v934 = v872 + v883<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v934))) = v880
	if v277 != 0 {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	goto L250
L250:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v963 <= int32(0) {
		goto L263
	} else {
		goto L264
	}
L251:
	;
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	v938 = *(*int32)(unsafe.Add(mBase, uint32(v844)))
	v939 = F_FunctionCall2Coll(m, v355, v936, v937, v938)
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L254
	}
L252:
	;
	v959 = float64(0)
	goto L253
L253:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v934)+8)) = v959
	v997 = v883 + int32(1)
	v998 = v886
	v999 = v887
	goto L242
L254:
	;
	v941 = *(*float64)(unsafe.Add(mBase, uint32(v939)))
	v942 = float64(0)
	if base.F64_ge(v941, v942) != 0 {
		goto L255
	} else {
		goto L256
	}
L255:
	;
	v945 = v941
	goto L257
L256:
	;
	v945 = v942
	goto L257
L257:
	;
	v946 = *(*int32)(unsafe.Add(mBase, uint32(v44)+208))
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v840)))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v34)+88))
	v949 = F_FunctionCall2Coll(m, v355, v946, v947, v948)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L258
	}
L258:
	;
	v951 = *(*float64)(unsafe.Add(mBase, uint32(v949)))
	v952 = float64(0)
	if base.F64_ge(v951, v952) != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v955 = v951
	goto L261
L260:
	;
	v955 = v952
	goto L261
L261:
	;
	v959 = base.F64_sub(v945, v955)
	goto L253
L262:
	;
	v971 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v969 + v971
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v974+v969<<(uint(v971)%32)))) = uint16(v879)
	v997 = v883
	v998 = v886
	v999 = v970
	goto L242
L263:
	;
	v969 = v963
	v970 = v914
	goto L262
L264:
	;
	goto L265
L265:
	;
	v966 = F_range_super_union(m, v44, v887, v914)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v969 = v968
	v970 = v966
	goto L262
L267:
	;
	v987 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v985 + v987
	v990 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v990+v985<<(uint(v987)%32)))) = uint16(v879)
	v997 = v883
	v998 = v986
	v999 = v887
	goto L242
L268:
	;
	v985 = v979
	v986 = v914
	goto L267
L269:
	;
	goto L270
L270:
	;
	v982 = F_range_super_union(m, v44, v886, v914)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L1
	} else {
		goto L271
	}
L271:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v985 = v984
	v986 = v982
	goto L267
L272:
	;
	goto L239
L273:
	;
	F_pg_qsort(m, v872, v997, int32(16), int32(1487))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L276
	}
L274:
	;
	v1105 = v998
	v1106 = v999
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v1105
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v1106
	goto L116
L276:
	;
	v1014 = int32(0)
	v1016 = v1014
	v1017 = v1014
	v1023 = v998
	v1024 = v999
	goto L277
L277:
	;
	v1047 = int32(4)
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v872+v1016<<(uint(v1047)%32))))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v75+v1050<<(uint(v1047)%32))))
	v1055 = F_pg_detoast_datum(m, v1054)
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L1
	} else {
		goto L279
	}
L278:
	;
	v1105 = v1091
	v1106 = v1092
	goto L275
L279:
	;
	if v1016 < v846 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1094 = v1017 + int32(1)
	v1096 = v1094 & int32(65535)
	if base.Ui32(v1096) < base.Ui32(v997) {
		v1016 = v1096
		v1017 = v1094
		v1023 = v1091
		v1024 = v1092
		goto L277
	} else {
		goto L294
	}
L281:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v1058 <= int32(0) {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	goto L283
L283:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v1074 <= int32(0) {
		goto L290
	} else {
		goto L291
	}
L284:
	;
	v1066 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v1064 + v1066
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1069+v1064<<(uint(v1066)%32)))) = uint16(v1050)
	v1091 = v1023
	v1092 = v1065
	goto L280
L285:
	;
	v1064 = v1058
	v1065 = v1055
	goto L284
L286:
	;
	goto L287
L287:
	;
	v1061 = F_range_super_union(m, v44, v1024, v1055)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L1
	} else {
		goto L288
	}
L288:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1064 = v1063
	v1065 = v1061
	goto L284
L289:
	;
	v1082 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1080 + v1082
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1085+v1080<<(uint(v1082)%32)))) = uint16(v1050)
	v1091 = v1081
	v1092 = v1024
	goto L280
L290:
	;
	v1080 = v1074
	v1081 = v1055
	goto L289
L291:
	;
	goto L292
L292:
	;
	v1077 = F_range_super_union(m, v44, v1023, v1055)
	mBase = m.M
	v1078 = m.ExcPending
	if v1078 != 0 {
		goto L1
	} else {
		goto L293
	}
L293:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v1080 = v1079
	v1081 = v1077
	goto L289
L294:
	;
	goto L278
L295:
	;
	goto L116
L296:
	;
	goto L116
L297:
	;
	goto L116
L298:
	;
	v1200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v37))))
	v1201 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1201
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v1201
	if v1200 != int32(1) {
		goto L309
	} else {
		goto L310
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(1)
	goto L298
L300:
	;
	goto L301
L301:
	;
	v1155 = v163 + v173 + v172 + v166
	v1156 = v51 - v1155
	v1158 = v163 + v164 + v183
	if v1158 <= int32(0) {
		goto L302
	} else {
		goto L303
	}
L302:
	;
	if v1155 <= int32(0) {
		goto L306
	} else {
		goto L307
	}
L303:
	;
	v1161 = v51 - v1158
	if v1161 <= int32(0) {
		goto L302
	} else {
		goto L304
	}
L304:
	;
	v1164 = v1161 - v1158
	v1165 = int32(31)
	v1166 = v1164 >> (uint(v1165) % 32)
	v1169 = v1156 - v1155
	v1171 = v1169 >> (uint(v1165) % 32)
	if v1169^v1171-v1171 < v1164^v1166-v1166 {
		goto L302
	} else {
		goto L305
	}
L305:
	;
	v1175 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v34)+32)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v34)+16)) = v1175
	*(*int32)(unsafe.Add(mBase, uint32(v34))) = v1175
	goto L298
L306:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v34+v263<<(uint(int32(2))%32)))) = int32(1)
	goto L298
L307:
	;
	if v1156 <= int32(0) {
		goto L306
	} else {
		goto L308
	}
L308:
	;
	v1187 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v34)+8)) = v1187
	*(*int64)(unsafe.Add(mBase, uint32(v34))) = v1187
	goto L298
L309:
	;
	v1209 = int32(1)
	v1215 = v1209
	v1217 = v1209
	v1219 = v1201
	v1226 = v1201
	goto L312
L310:
	;
	v1323 = v1201
	v1330 = v1201
	goto L311
L311:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v1323
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v1330
	goto L116
L312:
	;
	v1250 = *(*int32)(unsafe.Add(mBase, uint32(v75+v1217<<(uint(int32(4))%32))))
	v1251 = F_pg_detoast_datum(m, v1250)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L1
	} else {
		goto L315
	}
L313:
	;
	v1323 = v1312
	v1330 = v1313
	goto L311
L314:
	;
	v1315 = v1215 + int32(1)
	v1317 = v1315 & int32(65535)
	if base.Ui32(v1317) <= base.Ui32((v1200-v1209)&int32(65535)) {
		v1215 = v1315
		v1217 = v1317
		v1219 = v1312
		v1226 = v1313
		goto L312
	} else {
		goto L336
	}
L315:
	;
	v1253 = *(*int32)(unsafe.Add(mBase, uint32(v1251)))
	v1259 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1251+int32(base.Ui32(v1253)>>(uint(int32(2))%32))-int32(1)))))
	goto L316
L316:
	;
	if v1259&int32(1) != 0 {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1272 = int32(8)
	goto L319
L318:
	;
	v1262 = int32(3)
	v1265 = int32(base.Ui32(v1259)>>(uint(v1262)%32)) & v1262
	if v1259 < int32(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	v1276 = *(*int32)(unsafe.Add(mBase, uint32(v34+v1272<<(uint(int32(2))%32))))
	if v1276 == int32(0) {
		goto L323
	} else {
		goto L324
	}
L320:
	;
	v1270 = v1265 | int32(4)
	goto L322
L321:
	;
	v1270 = v1265
	goto L322
L322:
	;
	v1272 = v1270
	goto L319
L323:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	if v1279 <= int32(0) {
		goto L327
	} else {
		goto L328
	}
L324:
	;
	goto L325
L325:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	if v1295 <= int32(0) {
		goto L332
	} else {
		goto L333
	}
L326:
	;
	v1287 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v1285 + v1287
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1290+v1285<<(uint(v1287)%32)))) = uint16(v1215)
	v1312 = v1219
	v1313 = v1286
	goto L314
L327:
	;
	v1285 = v1279
	v1286 = v1251
	goto L326
L328:
	;
	goto L329
L329:
	;
	v1282 = F_range_super_union(m, v44, v1226, v1251)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L330
	}
L330:
	;
	v1284 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v1285 = v1284
	v1286 = v1282
	goto L326
L331:
	;
	v1303 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v1301 + v1303
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1306+v1301<<(uint(v1303)%32)))) = uint16(v1215)
	v1312 = v1302
	v1313 = v1226
	goto L314
L332:
	;
	v1301 = v1295
	v1302 = v1251
	goto L331
L333:
	;
	goto L334
L334:
	;
	v1298 = F_range_super_union(m, v44, v1219, v1251)
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	v1300 = *(*int32)(unsafe.Add(mBase, uint32(v36)+20))
	v1301 = v1300
	v1302 = v1298
	goto L331
L336:
	;
	goto L313
L337:
	;
	F_qsort_arg(m, v285, v282, int32(16), int32(1485), v44)
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L1
	} else {
		goto L341
	}
L338:
	;
	v1352 = F__emscripten_memcpy_bulkmem(m, v287, v285, v284)
	mBase = m.M
	v1353 = v1352
	goto L340
L339:
	;
	v1353 = v287
	goto L340
L340:
	;
	goto L337
L341:
	;
	F_qsort_arg(m, v1353, v282, int32(16), int32(1486), v44)
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L1
	} else {
		goto L342
	}
L342:
	;
	goto L117
L343:
	;
	goto L116
}
func F_range_gt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return base.B2i32(int32(0) < v2)
	}
}
func F_range_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_stack_depth(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = F_get_range_io_data(m, l0, v15, int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = v17
	goto L4
L4:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v34-int32(9)))&base.B2i32(v34 != int32(32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v50 = v26
	v51 = int32(8965)
	v52 = int32(5)
	goto L16
L6:
	;
	v26 = v26 + int32(1)
	goto L4
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	m.G0 = v12 + int32(32)
	return v340
L10:
	;
	v300 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)) = uint8(v300)
	v303 = int32(base.Ui32(v295) >> (uint(int32(4)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+12)) = uint8(v303)
	v305 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)) = uint8(v305)
	v310 = int32(base.Ui32(v295)>>(uint(int32(3))%32)) & v305
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+20)) = uint8(v310)
	v315 = int32(base.Ui32(v295)>>(uint(int32(2))%32)) & v305
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+13)) = uint8(v315)
	v320 = int32(base.Ui32(v295)>>(uint(v305)%32)) & v305
	*(*uint8)(unsafe.Add(mBase, uint32(v12)+21)) = uint8(v320)
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v329 = F_make_range(m, v322, v12+int32(16), v12+int32(8), v295&v305, v16)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L89
	}
L11:
	;
	v340 = int32(0)
	goto L9
L12:
	;
	if v177&int32(9) != 0 {
		goto L82
	} else {
		goto L83
	}
L13:
	;
	v255 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v255)
	goto L11
L14:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L78
	}
L15:
	;
	if v97 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	if v52 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v97 = int32(0)
	goto L15
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == v56 {
		v78 = v55
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v80 = int32(1)
	if v78 != 0 {
		v50 = v50 + v80
		v51 = v51 + v80
		v52 = v52 - v80
		goto L16
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v55-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v66 = v55 | int32(32)
	goto L25
L24:
	;
	v66 = v55
	goto L25
L25:
	;
	if base.Ui32((v56-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v75 = v56 | int32(32)
	goto L28
L27:
	;
	v75 = v56
	goto L28
L28:
	;
	if v66 == v75 {
		v78 = v66
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v97 = v66 - v75
	goto L15
L30:
	;
	goto L20
L31:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v100
	v107 = v26 + int32(5)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v131 == int32(40) {
		goto L46
	} else {
		goto L47
	}
L34:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if base.Ui32(v115-int32(9)) < base.Ui32(int32(5)) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v125 = F_errsave_start(m, v16)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L41
	}
L36:
	;
	goto L35
L37:
	;
	v107 = v107 + int32(1)
	goto L34
L38:
	;
	if v115 == int32(32) {
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if v115 != 0 {
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v295 = int32(1)
	goto L10
L41:
	;
	if v125 == int32(0) {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	v225 = int32(628685)
	v231 = int32(2418)
	goto L14
L43:
	;
	v216 = F_errsave_start(m, v16)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L1
	} else {
		goto L76
	}
L44:
	;
	v210 = F_errsave_start(m, v16)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L74
	}
L45:
	;
	v204 = F_errsave_start(m, v16)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L72
	}
L46:
	;
	v138 = int32(0)
	goto L48
L47:
	;
	if v131 != int32(91) {
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v145 = F_range_parse_bound(m, v17, v26+int32(1), v12+int32(28), v12+int32(16), v16)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v138 = int32(2)
	goto L48
L50:
	;
	if v145 == int32(0) {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 != int32(44) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	v159 = F_range_parse_bound(m, v17, v145+int32(1), v12+int32(24), v12+int32(16), v16)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v159 == int32(0) {
		goto L13
	} else {
		goto L54
	}
L54:
	;
	if v152 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v165 = v138 | int32(8)
	goto L57
L56:
	;
	v165 = v138
	goto L57
L57:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+16)))
	if v168 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v169 = v165 | int32(16)
	goto L60
L59:
	;
	v169 = v165
	goto L60
L60:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v170 != int32(41) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	if v170 != int32(93) {
		goto L43
	} else {
		goto L64
	}
L62:
	;
	v177 = v169
	goto L63
L63:
	;
	v179 = v159
	goto L65
L64:
	;
	v177 = v169 | int32(4)
	goto L63
L65:
	;
	v188 = v179 + int32(1)
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if base.Ui32(v189-int32(9)) < base.Ui32(int32(5)) {
		v179 = v188
		goto L65
	} else {
		goto L67
	}
L66:
	;
	if v189 == int32(0) {
		goto L12
	} else {
		goto L69
	}
L67:
	;
	if v189 == int32(32) {
		v179 = v188
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v198 = F_errsave_start(m, v16)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	if v198 == int32(0) {
		goto L13
	} else {
		goto L71
	}
L71:
	;
	v225 = int32(566533)
	v231 = int32(2481)
	goto L14
L72:
	;
	if v204 == int32(0) {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	v225 = int32(566574)
	v231 = int32(2435)
	goto L14
L74:
	;
	if v210 == int32(0) {
		goto L13
	} else {
		goto L75
	}
L75:
	;
	v225 = int32(629437)
	v231 = int32(2450)
	goto L14
L76:
	;
	if v216 == int32(0) {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	v225 = int32(586876)
	v231 = int32(2470)
	goto L14
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v17
	F_errmsg(m, int32(710420), v12)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errdetail(m, v225, int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errsave_finish(m, v16, int32(493562), v231, int32(360405))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L13
L82:
	;
	if v177&int32(17) != 0 {
		v295 = v177
		goto L10
	} else {
		goto L86
	}
L83:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v265 = F_InputFunctionCallSafe(m, v23+int32(4), v261, v262, v14, v16, v12+int32(16))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	if v265 != 0 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v267)
	goto L11
L86:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	v277 = F_InputFunctionCallSafe(m, v23+int32(4), v273, v274, v14, v16, v12+int32(8))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v277 != 0 {
		v295 = v177
		goto L10
	} else {
		goto L88
	}
L88:
	;
	v279 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v279)
	goto L11
L89:
	;
	v340 = v329
	goto L9
}
func F_range_lower_inf(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v11 = int32(1)
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(int32(2))%32))-v11))))
		return int32(base.Ui32(v13)>>(uint(int32(3))%32)) & v11
	}
}
func F_range_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
			if v21 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				if v22 == v19 {
					v32 = v21
					v33 = F_range_eq_internal(m, v32, v12, v17)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						m.G0 = v9 + int32(16)
						return v33 ^ int32(1)
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(369887), v9)
								mBase = m.M
								v48 = m.ExcPending
								if v48 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(493562), int32(1776), int32(398373))
									mBase = m.M
									v53 = m.ExcPending
									if v53 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
							v32 = v25
							v33 = F_range_eq_internal(m, v32, v12, v17)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v33 ^ int32(1)
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(369887), v9)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(493562), int32(1776), int32(398373))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v25
						v32 = v25
						v33 = F_range_eq_internal(m, v32, v12, v17)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v33 ^ int32(1)
						}
					}
				}
			}
		}
	}
}
func F_range_overright_multirange_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v47 = v4
		m.G0 = v8 + int32(48)
		return v47
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v47 = v4
			m.G0 = v8 + int32(48)
			return v47
		} else {
			F_range_deserialize(m, l0, l1, v8+int32(40), v8+int32(32), v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_multirange_get_bounds(m, l0, l2, int32(0), v8+int32(24), v8+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v43 = F_range_cmp_bounds(m, l0, v8+int32(40), v8+int32(24))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v47 = base.B2i32(int32(0) <= v43)
						m.G0 = v8 + int32(48)
						return v47
					}
				}
			}
		}
	}
}
func F_range_sortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(1484)
	return v3
}
func F_range_super_union(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	F_range_deserialize(m, l0, l1, v12+int32(40), v12+int32(24), v12+int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l0, l2, v12+int32(32), v12+int32(16), v12+int32(14))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v32)>>(uint(int32(2))%32))-int32(1)))))
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(l2+int32(base.Ui32(v39)>>(uint(int32(2))%32))-int32(1)))))
			v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
			if v46 == int32(1) {
				if v45&int32(-127) == int32(0) {
					v58 = l2
					v61 = F_datumCopy(m, v58, int32(0), int32(-1))
					mBase = m.M
					v62 = m.ExcPending
					if v62 != 0 {
						return int32(0)
					} else {
						v112 = v61
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
						v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
						v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
						v124 = v122 | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
						v127 = v112
						m.G0 = v12 + int32(48)
						return v127
					}
				} else {
					v127 = l2
					m.G0 = v12 + int32(48)
					return v127
				}
			} else {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
				if v53 != int32(1) {
					v67 = F_range_cmp_bounds(m, l0, v12+int32(40), v12+int32(32))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						v70 = base.B2i32(v67 <= int32(0))
						v75 = F_range_cmp_bounds(m, l0, v12+int32(24), v12+int32(16))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return int32(0)
						} else {
							if int32(0) < v67 {
								if v67 <= int32(0) {
									if v67 <= int32(0) {
										v96 = v12 + int32(40)
									} else {
										v96 = v12 + int32(32)
									}
									if int32(0) <= v75 {
										v103 = v12 + int32(24)
									} else {
										v103 = v12 + int32(16)
									}
									v104 = int32(0)
									v106 = F_make_range(m, l0, v96, v103, v104, v104)
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										if v38 < int32(0) {
											v112 = v106
											v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
											v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
											v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
											v124 = v122 | int32(128)
											*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
											v127 = v112
										} else {
											if int32(0) <= v45 {
												v127 = v106
											} else {
												v112 = v106
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
												v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
												v124 = v122 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
												v127 = v112
											}
										}
										m.G0 = v12 + int32(48)
										return v127
									}
								} else {
									if int32(0) <= v75 {
										if v67 <= int32(0) {
											v96 = v12 + int32(40)
										} else {
											v96 = v12 + int32(32)
										}
										if int32(0) <= v75 {
											v103 = v12 + int32(24)
										} else {
											v103 = v12 + int32(16)
										}
										v104 = int32(0)
										v106 = F_make_range(m, l0, v96, v103, v104, v104)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											if v38 < int32(0) {
												v112 = v106
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
												v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
												v124 = v122 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
												v127 = v112
											} else {
												if int32(0) <= v45 {
													v127 = v106
												} else {
													v112 = v106
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
													v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
													v124 = v122 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
													v127 = v112
												}
											}
											m.G0 = v12 + int32(48)
											return v127
										}
									} else {
										if v45 < int32(0) {
											v127 = l2
											m.G0 = v12 + int32(48)
											return v127
										} else {
											if int32(0) <= v38 {
												v127 = l2
												m.G0 = v12 + int32(48)
												return v127
											} else {
												if v67 <= int32(0) {
													v96 = v12 + int32(40)
												} else {
													v96 = v12 + int32(32)
												}
												if int32(0) <= v75 {
													v103 = v12 + int32(24)
												} else {
													v103 = v12 + int32(16)
												}
												v104 = int32(0)
												v106 = F_make_range(m, l0, v96, v103, v104, v104)
												mBase = m.M
												v107 = m.ExcPending
												if v107 != 0 {
													return int32(0)
												} else {
													if v38 < int32(0) {
														v112 = v106
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
														v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
														v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
														v124 = v122 | int32(128)
														*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
														v127 = v112
													} else {
														if int32(0) <= v45 {
															v127 = v106
														} else {
															v112 = v106
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
															v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
															v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
															v124 = v122 | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
															v127 = v112
														}
													}
													m.G0 = v12 + int32(48)
													return v127
												}
											}
										}
									}
								}
							} else {
								if v75 < int32(0) {
									if v67 <= int32(0) {
										if v67 <= int32(0) {
											v96 = v12 + int32(40)
										} else {
											v96 = v12 + int32(32)
										}
										if int32(0) <= v75 {
											v103 = v12 + int32(24)
										} else {
											v103 = v12 + int32(16)
										}
										v104 = int32(0)
										v106 = F_make_range(m, l0, v96, v103, v104, v104)
										mBase = m.M
										v107 = m.ExcPending
										if v107 != 0 {
											return int32(0)
										} else {
											if v38 < int32(0) {
												v112 = v106
												v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
												v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
												v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
												v124 = v122 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
												v127 = v112
											} else {
												if int32(0) <= v45 {
													v127 = v106
												} else {
													v112 = v106
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
													v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
													v124 = v122 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
													v127 = v112
												}
											}
											m.G0 = v12 + int32(48)
											return v127
										}
									} else {
										if int32(0) <= v75 {
											if v67 <= int32(0) {
												v96 = v12 + int32(40)
											} else {
												v96 = v12 + int32(32)
											}
											if int32(0) <= v75 {
												v103 = v12 + int32(24)
											} else {
												v103 = v12 + int32(16)
											}
											v104 = int32(0)
											v106 = F_make_range(m, l0, v96, v103, v104, v104)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												if v38 < int32(0) {
													v112 = v106
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
													v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
													v124 = v122 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
													v127 = v112
												} else {
													if int32(0) <= v45 {
														v127 = v106
													} else {
														v112 = v106
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
														v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
														v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
														v124 = v122 | int32(128)
														*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
														v127 = v112
													}
												}
												m.G0 = v12 + int32(48)
												return v127
											}
										} else {
											if v45 < int32(0) {
												v127 = l2
												m.G0 = v12 + int32(48)
												return v127
											} else {
												if int32(0) <= v38 {
													v127 = l2
													m.G0 = v12 + int32(48)
													return v127
												} else {
													if v67 <= int32(0) {
														v96 = v12 + int32(40)
													} else {
														v96 = v12 + int32(32)
													}
													if int32(0) <= v75 {
														v103 = v12 + int32(24)
													} else {
														v103 = v12 + int32(16)
													}
													v104 = int32(0)
													v106 = F_make_range(m, l0, v96, v103, v104, v104)
													mBase = m.M
													v107 = m.ExcPending
													if v107 != 0 {
														return int32(0)
													} else {
														if v38 < int32(0) {
															v112 = v106
															v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
															v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
															v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
															v124 = v122 | int32(128)
															*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
															v127 = v112
														} else {
															if int32(0) <= v45 {
																v127 = v106
															} else {
																v112 = v106
																v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
																v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
																v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
																v124 = v122 | int32(128)
																*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
																v127 = v112
															}
														}
														m.G0 = v12 + int32(48)
														return v127
													}
												}
											}
										}
									}
								} else {
									if v38 < int32(0) {
										v127 = l1
										m.G0 = v12 + int32(48)
										return v127
									} else {
										if v45 < int32(0) {
											if v67 <= int32(0) {
												v96 = v12 + int32(40)
											} else {
												v96 = v12 + int32(32)
											}
											if int32(0) <= v75 {
												v103 = v12 + int32(24)
											} else {
												v103 = v12 + int32(16)
											}
											v104 = int32(0)
											v106 = F_make_range(m, l0, v96, v103, v104, v104)
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												if v38 < int32(0) {
													v112 = v106
													v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
													v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
													v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
													v124 = v122 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
													v127 = v112
												} else {
													if int32(0) <= v45 {
														v127 = v106
													} else {
														v112 = v106
														v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
														v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
														v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
														v124 = v122 | int32(128)
														*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
														v127 = v112
													}
												}
												m.G0 = v12 + int32(48)
												return v127
											}
										} else {
											v127 = l1
											m.G0 = v12 + int32(48)
											return v127
										}
									}
								}
							}
						}
					}
				} else {
					if v38&int32(-127) != 0 {
						v127 = l1
						m.G0 = v12 + int32(48)
						return v127
					} else {
						v58 = l1
						v61 = F_datumCopy(m, v58, int32(0), int32(-1))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							v112 = v61
							v116 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
							v121 = v112 + int32(base.Ui32(v116)>>(uint(int32(2))%32)) - int32(1)
							v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121))))
							v124 = v122 | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v121))) = uint8(v124)
							v127 = v112
							m.G0 = v12 + int32(48)
							return v127
						}
					}
				}
			}
		}
	}
}
func F_range_table_entry_walker_impl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	if l3&int32(16) != 0 {
		v8 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			if v8 != 0 {
				v65 = int32(1)
				return v65
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				switch v12 {
				case 0:
					v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v14 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v13, l2)
					mBase = m.M
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						if v14 == int32(0) {
							v55 = int32(1)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 != 0 {
									v65 = v55
									return v65
								} else {
									if l3&int32(32) != 0 {
										v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v61 != 0 {
												v65 = v55
												return v65
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v65 = int32(1)
							return v65
						}
					}
				case 1:
					v18 = int32(1)
					if l3&v18 != 0 {
						v55 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v65 = v55
								return v65
							} else {
								if l3&int32(32) != 0 {
									v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 != 0 {
											v65 = v55
											return v65
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v22 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v21, l2)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							if v22 != 0 {
								v65 = v18
								return v65
							} else {
								v55 = int32(1)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									if v57 != 0 {
										v65 = v55
										return v65
									} else {
										if l3&int32(32) != 0 {
											v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												if v61 != 0 {
													v65 = v55
													return v65
												} else {
													return int32(0)
												}
											}
										} else {
											return int32(0)
										}
									}
								}
							}
						}
					}
				case 2:
					if l3&int32(4) != 0 {
						v55 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v65 = v55
								return v65
							} else {
								if l3&int32(32) != 0 {
									v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 != 0 {
											v65 = v55
											return v65
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v27 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v26, l2)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							if v27 == int32(0) {
								v55 = int32(1)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									if v57 != 0 {
										v65 = v55
										return v65
									} else {
										if l3&int32(32) != 0 {
											v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												if v61 != 0 {
													v65 = v55
													return v65
												} else {
													return int32(0)
												}
											}
										} else {
											return int32(0)
										}
									}
								}
							} else {
								v65 = int32(1)
								return v65
							}
						}
					}
				case 3:
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v32 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v31, l2)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						if v32 == int32(0) {
							v55 = int32(1)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 != 0 {
									v65 = v55
									return v65
								} else {
									if l3&int32(32) != 0 {
										v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v61 != 0 {
												v65 = v55
												return v65
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v65 = int32(1)
							return v65
						}
					}
				case 4:
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v37 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v36, l2)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						if v37 == int32(0) {
							v55 = int32(1)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 != 0 {
									v65 = v55
									return v65
								} else {
									if l3&int32(32) != 0 {
										v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v61 != 0 {
												v65 = v55
												return v65
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v65 = int32(1)
							return v65
						}
					}
				case 5:
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					v42 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v41, l2)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						if v42 == int32(0) {
							v55 = int32(1)
							v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								if v57 != 0 {
									v65 = v55
									return v65
								} else {
									if l3&int32(32) != 0 {
										v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
											return int32(0)
										} else {
											if v61 != 0 {
												v65 = v55
												return v65
											} else {
												return int32(0)
											}
										}
									} else {
										return int32(0)
									}
								}
							}
						} else {
							v65 = int32(1)
							return v65
						}
					}
				default:
					v55 = int32(1)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 != 0 {
							v65 = v55
							return v65
						} else {
							if l3&int32(32) != 0 {
								v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 != 0 {
										v65 = v55
										return v65
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				case 9:
					if l3&int32(256) != 0 {
						v55 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v65 = v55
								return v65
							} else {
								if l3&int32(32) != 0 {
									v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 != 0 {
											v65 = v55
											return v65
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v48, l2)
						mBase = m.M
						v50 = m.ExcPending
						if v50 != 0 {
							return int32(0)
						} else {
							if v49 == int32(0) {
								v55 = int32(1)
								v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									if v57 != 0 {
										v65 = v55
										return v65
									} else {
										if l3&int32(32) != 0 {
											v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v62 = m.ExcPending
											if v62 != 0 {
												return int32(0)
											} else {
												if v61 != 0 {
													v65 = v55
													return v65
												} else {
													return int32(0)
												}
											}
										} else {
											return int32(0)
										}
									}
								}
							} else {
								v65 = int32(1)
								return v65
							}
						}
					}
				}
			}
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		switch v12 {
		case 0:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v14 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v13, l2)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v55 = int32(1)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 != 0 {
							v65 = v55
							return v65
						} else {
							if l3&int32(32) != 0 {
								v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 != 0 {
										v65 = v55
										return v65
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v65 = int32(1)
					return v65
				}
			}
		case 1:
			v18 = int32(1)
			if l3&v18 != 0 {
				v55 = int32(1)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					if v57 != 0 {
						v65 = v55
						return v65
					} else {
						if l3&int32(32) != 0 {
							v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								if v61 != 0 {
									v65 = v55
									return v65
								} else {
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v22 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v21, l2)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					if v22 != 0 {
						v65 = v18
						return v65
					} else {
						v55 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v65 = v55
								return v65
							} else {
								if l3&int32(32) != 0 {
									v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 != 0 {
											v65 = v55
											return v65
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					}
				}
			}
		case 2:
			if l3&int32(4) != 0 {
				v55 = int32(1)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					if v57 != 0 {
						v65 = v55
						return v65
					} else {
						if l3&int32(32) != 0 {
							v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								if v61 != 0 {
									v65 = v55
									return v65
								} else {
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v27 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v26, l2)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						v55 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v65 = v55
								return v65
							} else {
								if l3&int32(32) != 0 {
									v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 != 0 {
											v65 = v55
											return v65
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v65 = int32(1)
						return v65
					}
				}
			}
		case 3:
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v32 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v31, l2)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				if v32 == int32(0) {
					v55 = int32(1)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 != 0 {
							v65 = v55
							return v65
						} else {
							if l3&int32(32) != 0 {
								v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 != 0 {
										v65 = v55
										return v65
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v65 = int32(1)
					return v65
				}
			}
		case 4:
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v37 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v36, l2)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				if v37 == int32(0) {
					v55 = int32(1)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 != 0 {
							v65 = v55
							return v65
						} else {
							if l3&int32(32) != 0 {
								v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 != 0 {
										v65 = v55
										return v65
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v65 = int32(1)
					return v65
				}
			}
		case 5:
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v42 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v41, l2)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				if v42 == int32(0) {
					v55 = int32(1)
					v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 != 0 {
							v65 = v55
							return v65
						} else {
							if l3&int32(32) != 0 {
								v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									if v61 != 0 {
										v65 = v55
										return v65
									} else {
										return int32(0)
									}
								}
							} else {
								return int32(0)
							}
						}
					}
				} else {
					v65 = int32(1)
					return v65
				}
			}
		default:
			v55 = int32(1)
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				if v57 != 0 {
					v65 = v55
					return v65
				} else {
					if l3&int32(32) != 0 {
						v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return int32(0)
						} else {
							if v61 != 0 {
								v65 = v55
								return v65
							} else {
								return int32(0)
							}
						}
					} else {
						return int32(0)
					}
				}
			}
		case 9:
			if l3&int32(256) != 0 {
				v55 = int32(1)
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return int32(0)
				} else {
					if v57 != 0 {
						v65 = v55
						return v65
					} else {
						if l3&int32(32) != 0 {
							v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return int32(0)
							} else {
								if v61 != 0 {
									v65 = v55
									return v65
								} else {
									return int32(0)
								}
							}
						} else {
							return int32(0)
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v49 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v48, l2)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return int32(0)
				} else {
					if v49 == int32(0) {
						v55 = int32(1)
						v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v57 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v56, l2)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 != 0 {
								v65 = v55
								return v65
							} else {
								if l3&int32(32) != 0 {
									v61 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										if v61 != 0 {
											v65 = v55
											return v65
										} else {
											return int32(0)
										}
									}
								} else {
									return int32(0)
								}
							}
						}
					} else {
						v65 = int32(1)
						return v65
					}
				}
			}
		}
	}
}
