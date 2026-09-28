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
	var v24 int32
	_ = v24
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
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
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
			v24 = int32(0)
			if base.B2i32(base.Ui32(int32(7)) < base.Ui32(v17))|base.B2i32(int32(1)<<(uint(v17)%32)&int32(169) == v24) == v24 {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_RangeVarCallbackMaintainsTable[0]))
				v32 = F_pg_class_aclcheck(m, l1, v30, int64(16384))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return
				} else {
					if v32 == int32(0) {
						m.G0 = v8 + int32(16)
						return
					} else {
						v36 = F_get_rel_relkind(m, l1)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							switch v36 - int32(73) {
							case 0, 32:
								v47 = int32(20)
								v49 = v47
							default:
								v47 = int32(42)
								v49 = v47
							case 10:
								v49 = int32(38)
							case 29:
								v49 = int32(18)
							case 36:
								v49 = int32(23)
							case 45:
								v49 = int32(52)
							}
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							F_aclcheck_error(m, v32, v49, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				if v15 == int32(0) {
					m.G0 = v8 + int32(16)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v62
							F_errmsg(m, int32(_a_F_RangeVarCallbackMaintainsTable_0), v8)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_RangeVarCallbackMaintainsTable_1), int32(_a_F_RangeVarCallbackMaintainsTable_2), int32(_a_F_RangeVarCallbackMaintainsTable_3))
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
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
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
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
	v36 = F_makeAlias(m, int32(_a_F_addRangeTableEntryForSubquery_0), int32(0))
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
		v121 = v6
		v122 = v6
		v123 = v6
		v124 = v6
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if v41 <= v121 {
		goto L33
	} else {
		goto L34
	}
L13:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v45 <= int32(0) {
		v121 = v6
		v122 = v6
		v123 = v6
		v124 = v6
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v50 = int32(0)
	v60 = v6
	v61 = v6
	v62 = v6
	v63 = v6
	goto L15
L15:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v65+v50<<(uint(int32(2))%32))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+26)))
	if v70 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v121 = v101
	v122 = v102
	v123 = v103
	v124 = v104
	goto L12
L17:
	;
	v74 = v60 + int32(1)
	if v41 < v74 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v101 = v60
	v102 = v61
	v103 = v62
	v104 = v63
	goto L19
L19:
	;
	v107 = v50 + int32(1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v107 < v108 {
		v50 = v107
		v60 = v101
		v61 = v102
		v62 = v103
		v63 = v104
		goto L15
	} else {
		goto L32
	}
L20:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v77 = F_pstrdup(m, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v87 = F_exprType(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v80 = F_makeString(m, v77)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v82 = F_lappend(m, v79, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = v82
	goto L22
L26:
	;
	v89 = F_lappend_oid(m, v62, v87)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v92 = F_exprTypmod(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v94 = F_lappend_int(m, v63, v92)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v97 = F_exprCollation(m, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v99 = F_lappend_oid(m, v61, v97)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v101 = v74
	v102 = v99
	v103 = v89
	v104 = v94
	goto L19
L32:
	;
	goto L16
L33:
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
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L41
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v131
	if v131 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v131)+4))
	v136 = v134
	goto L39
L38:
	;
	v136 = int32(0)
	goto L39
L39:
	;
	v137 = F_buildNSItemFromLists(m, v22, v136, v123, v124, v122)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v137)+20)) = uint8(base.B2i32(l2 != int32(0)))
	m.G0 = v19 + int32(16)
	return v137
L41:
	;
	F_errcode(m, int32(_a_F_addRangeTableEntryForSubquery_1))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v153
	F_errmsg(m, int32(_a_F_addRangeTableEntryForSubquery_2), v19)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_addRangeTableEntryForSubquery_3), int32(1729), int32(_a_F_addRangeTableEntryForSubquery_4))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_get_range_multirange(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14288(m, l0, int32(55))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
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
						F_errmsg(m, int32(_a_F_makeRangeVarFromAnyName_0), v8)
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
								F_errfinish(m, int32(_a_F_makeRangeVarFromAnyName_1), int32(_a_F_makeRangeVarFromAnyName_2), int32(_a_F_makeRangeVarFromAnyName_3))
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
							F_errmsg(m, int32(_a_F_makeRangeVarFromAnyName_0), v8)
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
									F_errfinish(m, int32(_a_F_makeRangeVarFromAnyName_1), int32(_a_F_makeRangeVarFromAnyName_2), int32(_a_F_makeRangeVarFromAnyName_3))
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
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(144)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v12)>>(uint(int32(2))%32))-int32(1)))))
	if v18&int32(1) != 0 {
		v78 = v4
		m.G0 = v10 + int32(144)
		return v78
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v21 == int32(0) {
			v78 = v4
			m.G0 = v10 + int32(144)
			return v78
		} else {
			F_range_deserialize(m, l0, l1, v10+int32(128), v10+int32(112), v10+int32(79))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v37 = v10 + int32(96)
				v39 = v10 + int32(80)
				F_multirange_get_bounds(m, l0, l2, int32(0), v37, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return int32(0)
				} else {
					v42 = *(*int64)(unsafe.Add(mBase, uint32(v10)+120))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(v10)+112))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v44
					v46 = *(*int64)(unsafe.Add(mBase, uint32(v10)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v46
					v48 = *(*int64)(unsafe.Add(mBase, uint32(v10)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v48
					v55 = F_bounds_adjacent(m, l0, v10+int32(56), v10+int32(40))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						if v55 != 0 {
							v78 = int32(1)
							m.G0 = v10 + int32(144)
							return v78
						} else {
							if int32(2) <= v34 {
								F_multirange_get_bounds(m, l0, l2, v34-int32(1), v37, v39)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return int32(0)
								} else {
									v63 = *(*int64)(unsafe.Add(mBase, uint32(v10)+88))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v63
									v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v65
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v10)+128))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v67
									v69 = *(*int64)(unsafe.Add(mBase, uint32(v10)+136))
									*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v69
									v75 = F_bounds_adjacent(m, l0, v10+int32(24), v10+int32(8))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v78 = v75
										m.G0 = v10 + int32(144)
										return v78
									}
								}
							} else {
								v63 = *(*int64)(unsafe.Add(mBase, uint32(v10)+88))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v63
								v65 = *(*int64)(unsafe.Add(mBase, uint32(v10)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v65
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v10)+128))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v67
								v69 = *(*int64)(unsafe.Add(mBase, uint32(v10)+136))
								*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v69
								v75 = F_bounds_adjacent(m, l0, v10+int32(24), v10+int32(8))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return int32(0)
								} else {
									v78 = v75
									m.G0 = v10 + int32(144)
									return v78
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int64
	_ = v68
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_range_deserialize(m, l2, v11, v8-int32(-64), v8+int32(48), v8+int32(15))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l2, v10, v8+int32(32), v8+int32(16), v8+int32(14))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return int32(0)
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+14)))
			v33 = int32(1)
			if v30 != 0 {
				v36 = v30&v31 - v33
			} else {
				v36 = v33
			}
			if v30|v31&int32(1) != 0 {
				v165 = v36
				m.G0 = v8 + int32(80)
				return v165
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)))
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+72)))
				if v41 == int32(1) {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+74)))
					if v40&int32(1) != 0 {
						v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
						if v47 == v44 {
							v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
							v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)))
							if v103 == int32(1) {
								v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
								if v102&int32(1) != 0 {
									v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
									if v109 == v106 {
										v165 = int32(0)
									} else {
										v113 = int32(1)
										if v106&v113 != 0 {
											v116 = int32(-1)
										} else {
											v116 = v113
										}
										v165 = v116
									}
								} else {
									v118 = int32(1)
									if v106&v118 != 0 {
										v121 = int32(-1)
									} else {
										v121 = v118
									}
									v165 = v121
								}
								m.G0 = v8 + int32(80)
								return v165
							} else {
								if v102&int32(1) != 0 {
									v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
									if v126 != 0 {
										v127 = int32(1)
									} else {
										v127 = int32(-1)
									}
									v165 = v127
									m.G0 = v8 + int32(80)
									return v165
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
									v131 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
									v132 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
									v133 = F_FunctionCall2Coll(m, l2+int32(212), v130, v131, v132)
									mBase = m.M
									v134 = m.ExcPending
									if v134 != 0 {
										return int32(0)
									} else {
										v135 = base.I32_wrap_i64(v133)
										if v135 != 0 {
											v165 = v135
										} else {
											v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
											v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+57)))
											if v137 == int32(0) {
												v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
												if v136&int32(1) == int32(0) {
													v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
													if v145 == v140 {
														v165 = int32(0)
													} else {
														v148 = int32(1)
														if v140&v148 != 0 {
															v152 = v148
														} else {
															v152 = int32(-1)
														}
														v165 = v152
													}
												} else {
													v153 = int32(1)
													if v140&v153 != 0 {
														v157 = v153
													} else {
														v157 = int32(-1)
													}
													v165 = v157
												}
											} else {
												if v136&int32(1) != 0 {
													v165 = int32(0)
												} else {
													v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
													if v163 != 0 {
														v164 = int32(-1)
													} else {
														v164 = int32(1)
													}
													v165 = v164
												}
											}
										}
										m.G0 = v8 + int32(80)
										return v165
									}
								}
							}
						} else {
							v50 = int32(1)
							if v44&v50 != 0 {
								v53 = int32(-1)
							} else {
								v53 = v50
							}
							v165 = v53
							m.G0 = v8 + int32(80)
							return v165
						}
					} else {
						v55 = int32(1)
						if v44&v55 != 0 {
							v58 = int32(-1)
						} else {
							v58 = v55
						}
						v165 = v58
						m.G0 = v8 + int32(80)
						return v165
					}
				} else {
					if v40&int32(1) != 0 {
						v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
						if v63 != 0 {
							v64 = int32(1)
						} else {
							v64 = int32(-1)
						}
						v165 = v64
						m.G0 = v8 + int32(80)
						return v165
					} else {
						v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
						v68 = *(*int64)(unsafe.Add(mBase, uint32(v8)+64))
						v69 = *(*int64)(unsafe.Add(mBase, uint32(v8)+32))
						v70 = F_FunctionCall2Coll(m, l2+int32(212), v67, v68, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							v72 = base.I32_wrap_i64(v70)
							if v72 != 0 {
								v165 = v72
								m.G0 = v8 + int32(80)
								return v165
							} else {
								v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+41)))
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+73)))
								if v74 == int32(0) {
									v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+74)))
									if v73&int32(1) == int32(0) {
										v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
										if v82 == v77 {
											v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
											v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)))
											if v103 == int32(1) {
												v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
												if v102&int32(1) != 0 {
													v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
													if v109 == v106 {
														v165 = int32(0)
													} else {
														v113 = int32(1)
														if v106&v113 != 0 {
															v116 = int32(-1)
														} else {
															v116 = v113
														}
														v165 = v116
													}
												} else {
													v118 = int32(1)
													if v106&v118 != 0 {
														v121 = int32(-1)
													} else {
														v121 = v118
													}
													v165 = v121
												}
												m.G0 = v8 + int32(80)
												return v165
											} else {
												if v102&int32(1) != 0 {
													v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
													if v126 != 0 {
														v127 = int32(1)
													} else {
														v127 = int32(-1)
													}
													v165 = v127
													m.G0 = v8 + int32(80)
													return v165
												} else {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
													v131 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
													v132 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
													v133 = F_FunctionCall2Coll(m, l2+int32(212), v130, v131, v132)
													mBase = m.M
													v134 = m.ExcPending
													if v134 != 0 {
														return int32(0)
													} else {
														v135 = base.I32_wrap_i64(v133)
														if v135 != 0 {
															v165 = v135
														} else {
															v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
															v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+57)))
															if v137 == int32(0) {
																v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
																if v136&int32(1) == int32(0) {
																	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
																	if v145 == v140 {
																		v165 = int32(0)
																	} else {
																		v148 = int32(1)
																		if v140&v148 != 0 {
																			v152 = v148
																		} else {
																			v152 = int32(-1)
																		}
																		v165 = v152
																	}
																} else {
																	v153 = int32(1)
																	if v140&v153 != 0 {
																		v157 = v153
																	} else {
																		v157 = int32(-1)
																	}
																	v165 = v157
																}
															} else {
																if v136&int32(1) != 0 {
																	v165 = int32(0)
																} else {
																	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
																	if v163 != 0 {
																		v164 = int32(-1)
																	} else {
																		v164 = int32(1)
																	}
																	v165 = v164
																}
															}
														}
														m.G0 = v8 + int32(80)
														return v165
													}
												}
											}
										} else {
											v84 = int32(1)
											if v77&v84 != 0 {
												v88 = v84
											} else {
												v88 = int32(-1)
											}
											v165 = v88
											m.G0 = v8 + int32(80)
											return v165
										}
									} else {
										v89 = int32(1)
										if v77&v89 != 0 {
											v93 = v89
										} else {
											v93 = int32(-1)
										}
										v165 = v93
										m.G0 = v8 + int32(80)
										return v165
									}
								} else {
									if v73&int32(1) != 0 {
										v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+24)))
										v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+56)))
										if v103 == int32(1) {
											v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
											if v102&int32(1) != 0 {
												v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
												if v109 == v106 {
													v165 = int32(0)
												} else {
													v113 = int32(1)
													if v106&v113 != 0 {
														v116 = int32(-1)
													} else {
														v116 = v113
													}
													v165 = v116
												}
											} else {
												v118 = int32(1)
												if v106&v118 != 0 {
													v121 = int32(-1)
												} else {
													v121 = v118
												}
												v165 = v121
											}
											m.G0 = v8 + int32(80)
											return v165
										} else {
											if v102&int32(1) != 0 {
												v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
												if v126 != 0 {
													v127 = int32(1)
												} else {
													v127 = int32(-1)
												}
												v165 = v127
												m.G0 = v8 + int32(80)
												return v165
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+208))
												v131 = *(*int64)(unsafe.Add(mBase, uint32(v8)+48))
												v132 = *(*int64)(unsafe.Add(mBase, uint32(v8)+16))
												v133 = F_FunctionCall2Coll(m, l2+int32(212), v130, v131, v132)
												mBase = m.M
												v134 = m.ExcPending
												if v134 != 0 {
													return int32(0)
												} else {
													v135 = base.I32_wrap_i64(v133)
													if v135 != 0 {
														v165 = v135
													} else {
														v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+25)))
														v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+57)))
														if v137 == int32(0) {
															v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+58)))
															if v136&int32(1) == int32(0) {
																v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
																if v145 == v140 {
																	v165 = int32(0)
																} else {
																	v148 = int32(1)
																	if v140&v148 != 0 {
																		v152 = v148
																	} else {
																		v152 = int32(-1)
																	}
																	v165 = v152
																}
															} else {
																v153 = int32(1)
																if v140&v153 != 0 {
																	v157 = v153
																} else {
																	v157 = int32(-1)
																}
																v165 = v157
															}
														} else {
															if v136&int32(1) != 0 {
																v165 = int32(0)
															} else {
																v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+26)))
																if v163 != 0 {
																	v164 = int32(-1)
																} else {
																	v164 = int32(1)
																}
																v165 = v164
															}
														}
													}
													m.G0 = v8 + int32(80)
													return v165
												}
											}
										}
									} else {
										v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+42)))
										if v98 != 0 {
											v99 = int32(-1)
										} else {
											v99 = int32(1)
										}
										v165 = v99
										m.G0 = v8 + int32(80)
										return v165
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
func F_range_constructor3(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int64
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
	var v65 int64
	_ = v65
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int64
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
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
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v9 = m.G0
	v11 = v9 - int32(48)
	m.G0 = v11
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = F_get_fn_expr_rettype(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int64(0)
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+16))
		if v21 != 0 {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
			if v22 == v16 {
				v32 = v21
				v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
				if v33 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v113 = m.ExcPending
					if v113 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(130))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_range_constructor3_0), int32(0))
							mBase = m.M
							v120 = m.ExcPending
							if v120 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_constructor3_1), int32(428), int32(_a_F_range_constructor3_2))
								mBase = m.M
								v125 = m.ExcPending
								if v125 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					v37 = F_pg_detoast_datum_packed(m, v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int64(0)
					} else {
						v39 = F_text_to_cstring(m, v37)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int64(0)
						} else {
							v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							if v41 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(16801924))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
										mBase = m.M
										v137 = m.ExcPending
										if v137 != 0 {
											return int64(0)
										} else {
											F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
											mBase = m.M
											v141 = m.ExcPending
											if v141 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
												mBase = m.M
												v146 = m.ExcPending
												if v146 != 0 {
													return int64(0)
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
									v130 = m.ExcPending
									if v130 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int64(0)
											} else {
												F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
												mBase = m.M
												v141 = m.ExcPending
												if v141 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return int64(0)
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
										v130 = m.ExcPending
										if v130 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int64(0)
												} else {
													F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return int64(0)
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
												v150 = m.ExcPending
												if v150 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(16801924))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
														mBase = m.M
														v157 = m.ExcPending
														if v157 != 0 {
															return int64(0)
														} else {
															F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
															mBase = m.M
															v161 = m.ExcPending
															if v161 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2504), int32(_a_F_range_constructor3_5))
																mBase = m.M
																v166 = m.ExcPending
																if v166 != 0 {
																	return int64(0)
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
														v170 = m.ExcPending
														if v170 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
																} else {
																	F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2518), int32(_a_F_range_constructor3_5))
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return int64(0)
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
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
														if v62 != 0 {
															v65 = int64(0)
														} else {
															v65 = v14
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
														if v74 != 0 {
															v82 = int64(0)
														} else {
															v82 = v13
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															m.G0 = v11 + int32(48)
															return base.I64_extend_i32_u(v90)
														}
													}
												} else {
													v61 = v54
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
													if v62 != 0 {
														v65 = int64(0)
													} else {
														v65 = v14
													}
													*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
													if v74 != 0 {
														v82 = int64(0)
													} else {
														v82 = v13
													}
													*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int64(0)
													} else {
														m.G0 = v11 + int32(48)
														return base.I64_extend_i32_u(v90)
													}
												}
											}
										} else {
											v54 = int32(0)
											if v44 != int32(41) {
												if v44 != int32(93) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(16801924))
														mBase = m.M
														v173 = m.ExcPending
														if v173 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
															mBase = m.M
															v177 = m.ExcPending
															if v177 != 0 {
																return int64(0)
															} else {
																F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2518), int32(_a_F_range_constructor3_5))
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
																		return int64(0)
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
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
													if v62 != 0 {
														v65 = int64(0)
													} else {
														v65 = v14
													}
													*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
													if v74 != 0 {
														v82 = int64(0)
													} else {
														v82 = v13
													}
													*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int64(0)
													} else {
														m.G0 = v11 + int32(48)
														return base.I64_extend_i32_u(v90)
													}
												}
											} else {
												v61 = v54
												v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
												if v62 != 0 {
													v65 = int64(0)
												} else {
													v65 = v14
												}
												*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
												v67 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
												v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
												v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
												v75 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
												*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
												if v74 != 0 {
													v82 = int64(0)
												} else {
													v82 = v13
												}
												*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
												v88 = int32(0)
												v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
												mBase = m.M
												v91 = m.ExcPending
												if v91 != 0 {
													return int64(0)
												} else {
													m.G0 = v11 + int32(48)
													return base.I64_extend_i32_u(v90)
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
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
							F_errmsg_internal(m, int32(_a_F_range_constructor3_6), v11)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_constructor3_1), int32(1946), int32(_a_F_range_constructor3_7))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return int64(0)
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
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
						if v33 == int32(1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(130))
								mBase = m.M
								v116 = m.ExcPending
								if v116 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_range_constructor3_0), int32(0))
									mBase = m.M
									v120 = m.ExcPending
									if v120 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_range_constructor3_1), int32(428), int32(_a_F_range_constructor3_2))
										mBase = m.M
										v125 = m.ExcPending
										if v125 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							v37 = F_pg_detoast_datum_packed(m, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int64(0)
							} else {
								v39 = F_text_to_cstring(m, v37)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int64(0)
								} else {
									v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
									if v41 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v130 = m.ExcPending
										if v130 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int64(0)
												} else {
													F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return int64(0)
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
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(16801924))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int64(0)
													} else {
														F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return int64(0)
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
												v130 = m.ExcPending
												if v130 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(16801924))
													mBase = m.M
													v133 = m.ExcPending
													if v133 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int64(0)
														} else {
															F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
															mBase = m.M
															v141 = m.ExcPending
															if v141 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
																mBase = m.M
																v146 = m.ExcPending
																if v146 != 0 {
																	return int64(0)
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
														v150 = m.ExcPending
														if v150 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																mBase = m.M
																v157 = m.ExcPending
																if v157 != 0 {
																	return int64(0)
																} else {
																	F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																	mBase = m.M
																	v161 = m.ExcPending
																	if v161 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2504), int32(_a_F_range_constructor3_5))
																		mBase = m.M
																		v166 = m.ExcPending
																		if v166 != 0 {
																			return int64(0)
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
																v170 = m.ExcPending
																if v170 != 0 {
																	return int64(0)
																} else {
																	F_errcode(m, int32(16801924))
																	mBase = m.M
																	v173 = m.ExcPending
																	if v173 != 0 {
																		return int64(0)
																	} else {
																		F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																		mBase = m.M
																		v177 = m.ExcPending
																		if v177 != 0 {
																			return int64(0)
																		} else {
																			F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																			mBase = m.M
																			v181 = m.ExcPending
																			if v181 != 0 {
																				return int64(0)
																			} else {
																				F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2518), int32(_a_F_range_constructor3_5))
																				mBase = m.M
																				v186 = m.ExcPending
																				if v186 != 0 {
																					return int64(0)
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
																v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
																if v62 != 0 {
																	v65 = int64(0)
																} else {
																	v65 = v14
																}
																*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
																v67 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
																v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
																v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
																v75 = int32(0)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
																*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
																if v74 != 0 {
																	v82 = int64(0)
																} else {
																	v82 = v13
																}
																*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
																v88 = int32(0)
																v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
																mBase = m.M
																v91 = m.ExcPending
																if v91 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v11 + int32(48)
																	return base.I64_extend_i32_u(v90)
																}
															}
														} else {
															v61 = v54
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
															if v62 != 0 {
																v65 = int64(0)
															} else {
																v65 = v14
															}
															*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
															if v74 != 0 {
																v82 = int64(0)
															} else {
																v82 = v13
															}
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int64(0)
															} else {
																m.G0 = v11 + int32(48)
																return base.I64_extend_i32_u(v90)
															}
														}
													}
												} else {
													v54 = int32(0)
													if v44 != int32(41) {
														if v44 != int32(93) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(16801924))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int64(0)
																	} else {
																		F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																		mBase = m.M
																		v181 = m.ExcPending
																		if v181 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2518), int32(_a_F_range_constructor3_5))
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
																				return int64(0)
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
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
															if v62 != 0 {
																v65 = int64(0)
															} else {
																v65 = v14
															}
															*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
															if v74 != 0 {
																v82 = int64(0)
															} else {
																v82 = v13
															}
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int64(0)
															} else {
																m.G0 = v11 + int32(48)
																return base.I64_extend_i32_u(v90)
															}
														}
													} else {
														v61 = v54
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
														if v62 != 0 {
															v65 = int64(0)
														} else {
															v65 = v14
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
														if v74 != 0 {
															v82 = int64(0)
														} else {
															v82 = v13
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															m.G0 = v11 + int32(48)
															return base.I64_extend_i32_u(v90)
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
				return int64(0)
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
				if v27 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v16
						F_errmsg_internal(m, int32(_a_F_range_constructor3_6), v11)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_range_constructor3_1), int32(1946), int32(_a_F_range_constructor3_7))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int64(0)
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
					v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
					if v33 == int32(1) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(130))
							mBase = m.M
							v116 = m.ExcPending
							if v116 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_range_constructor3_0), int32(0))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_range_constructor3_1), int32(428), int32(_a_F_range_constructor3_2))
									mBase = m.M
									v125 = m.ExcPending
									if v125 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						v37 = F_pg_detoast_datum_packed(m, v36)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int64(0)
						} else {
							v39 = F_text_to_cstring(m, v37)
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return int64(0)
							} else {
								v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
								if v41 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v130 = m.ExcPending
									if v130 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(16801924))
										mBase = m.M
										v133 = m.ExcPending
										if v133 != 0 {
											return int64(0)
										} else {
											F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
											mBase = m.M
											v137 = m.ExcPending
											if v137 != 0 {
												return int64(0)
											} else {
												F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
												mBase = m.M
												v141 = m.ExcPending
												if v141 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
													mBase = m.M
													v146 = m.ExcPending
													if v146 != 0 {
														return int64(0)
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
										v130 = m.ExcPending
										if v130 != 0 {
											return int64(0)
										} else {
											F_errcode(m, int32(16801924))
											mBase = m.M
											v133 = m.ExcPending
											if v133 != 0 {
												return int64(0)
											} else {
												F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
												mBase = m.M
												v137 = m.ExcPending
												if v137 != 0 {
													return int64(0)
												} else {
													F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
													mBase = m.M
													v141 = m.ExcPending
													if v141 != 0 {
														return int64(0)
													} else {
														F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
														mBase = m.M
														v146 = m.ExcPending
														if v146 != 0 {
															return int64(0)
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
											v130 = m.ExcPending
											if v130 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(16801924))
												mBase = m.M
												v133 = m.ExcPending
												if v133 != 0 {
													return int64(0)
												} else {
													F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
													mBase = m.M
													v137 = m.ExcPending
													if v137 != 0 {
														return int64(0)
													} else {
														F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
														mBase = m.M
														v141 = m.ExcPending
														if v141 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2491), int32(_a_F_range_constructor3_5))
															mBase = m.M
															v146 = m.ExcPending
															if v146 != 0 {
																return int64(0)
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
													v150 = m.ExcPending
													if v150 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(16801924))
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
															mBase = m.M
															v157 = m.ExcPending
															if v157 != 0 {
																return int64(0)
															} else {
																F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2504), int32(_a_F_range_constructor3_5))
																	mBase = m.M
																	v166 = m.ExcPending
																	if v166 != 0 {
																		return int64(0)
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
															v170 = m.ExcPending
															if v170 != 0 {
																return int64(0)
															} else {
																F_errcode(m, int32(16801924))
																mBase = m.M
																v173 = m.ExcPending
																if v173 != 0 {
																	return int64(0)
																} else {
																	F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																	mBase = m.M
																	v177 = m.ExcPending
																	if v177 != 0 {
																		return int64(0)
																	} else {
																		F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																		mBase = m.M
																		v181 = m.ExcPending
																		if v181 != 0 {
																			return int64(0)
																		} else {
																			F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2518), int32(_a_F_range_constructor3_5))
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
																				return int64(0)
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
															v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
															if v62 != 0 {
																v65 = int64(0)
															} else {
																v65 = v14
															}
															*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
															v67 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
															v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
															v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
															v75 = int32(0)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
															*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
															if v74 != 0 {
																v82 = int64(0)
															} else {
																v82 = v13
															}
															*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
															v88 = int32(0)
															v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
															mBase = m.M
															v91 = m.ExcPending
															if v91 != 0 {
																return int64(0)
															} else {
																m.G0 = v11 + int32(48)
																return base.I64_extend_i32_u(v90)
															}
														}
													} else {
														v61 = v54
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
														if v62 != 0 {
															v65 = int64(0)
														} else {
															v65 = v14
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
														if v74 != 0 {
															v82 = int64(0)
														} else {
															v82 = v13
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															m.G0 = v11 + int32(48)
															return base.I64_extend_i32_u(v90)
														}
													}
												}
											} else {
												v54 = int32(0)
												if v44 != int32(41) {
													if v44 != int32(93) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(16801924))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int64(0)
															} else {
																F_errmsg(m, int32(_a_F_range_constructor3_3), int32(0))
																mBase = m.M
																v177 = m.ExcPending
																if v177 != 0 {
																	return int64(0)
																} else {
																	F_errhint(m, int32(_a_F_range_constructor3_4), int32(0))
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_range_constructor3_1), int32(2518), int32(_a_F_range_constructor3_5))
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
																			return int64(0)
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
														v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
														if v62 != 0 {
															v65 = int64(0)
														} else {
															v65 = v14
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
														v67 = int32(1)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
														v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
														v75 = int32(0)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
														*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
														if v74 != 0 {
															v82 = int64(0)
														} else {
															v82 = v13
														}
														*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
														v88 = int32(0)
														v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
														mBase = m.M
														v91 = m.ExcPending
														if v91 != 0 {
															return int64(0)
														} else {
															m.G0 = v11 + int32(48)
															return base.I64_extend_i32_u(v90)
														}
													}
												} else {
													v61 = v54
													v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+40)) = uint8(v62)
													if v62 != 0 {
														v65 = int64(0)
													} else {
														v65 = v14
													}
													*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v65
													v67 = int32(1)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+42)) = uint8(v67)
													v72 = int32(base.Ui32(v61)>>(uint(v67)%32)) & v67
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+41)) = uint8(v72)
													v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
													v75 = int32(0)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+26)) = uint8(v75)
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+25)) = uint8(base.B2i32(base.Ui32(int32(3)) < base.Ui32(v61)))
													*(*uint8)(unsafe.Add(mBase, uint32(v11)+24)) = uint8(v74)
													if v74 != 0 {
														v82 = int64(0)
													} else {
														v82 = v13
													}
													*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v82
													v88 = int32(0)
													v90 = F_make_range(m, v32, v11+int32(32), v11+int32(16), v88, v88)
													mBase = m.M
													v91 = m.ExcPending
													if v91 != 0 {
														return int64(0)
													} else {
														m.G0 = v11 + int32(48)
														return base.I64_extend_i32_u(v90)
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
func F_range_contains_elem(m *base.Module, l0 int32) int64 {
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
	var v17 int64
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
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v16 {
				v30 = v19
				v31 = F_range_contains_elem_internal(m, v30, v12, v17)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 + int32(16)
					return base.I64_extend_i32_u(v31)
				}
			} else {
				v23 = F_lookup_type_cache(m, v16, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
							F_errmsg_internal(m, int32(_a_F_range_contains_elem_0), v9)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_contains_elem_1), int32(1946), int32(_a_F_range_contains_elem_2))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int64(0)
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
						v31 = F_range_contains_elem_internal(m, v30, v12, v17)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v31)
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v16, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v16
						F_errmsg_internal(m, int32(_a_F_range_contains_elem_0), v9)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_range_contains_elem_1), int32(1946), int32(_a_F_range_contains_elem_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
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
					v31 = F_range_contains_elem_internal(m, v30, v12, v17)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v31)
					}
				}
			}
		}
	}
}
func F_range_deserialize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int64
	_ = v110
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	v11 = int64(0)
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v18 = int32(*(*int8)(unsafe.Add(mBase, uint32(v17)+11)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+10)))
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+8)))
	v22 = l1 + int32(8)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v23)>>(uint(int32(2))%32))-int32(1)))))
	if v29&int32(41) != 0 {
		v89 = v11
		v90 = v22
		goto L7
	} else {
		goto L8
	}
L1:
	;
	v178 = int32(1)
	v179 = v29 & v178
	*(*uint8)(unsafe.Add(mBase, uint32(l4))) = uint8(v179)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+10)) = uint8(v178)
	v186 = int32(base.Ui32(v29)>>(uint(v178)%32)) & v178
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+9)) = uint8(v186)
	v191 = int32(base.Ui32(v29)>>(uint(int32(3))%32)) & v178
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v191)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v177
	v194 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+10)) = uint8(v194)
	v199 = int32(base.Ui32(v29)>>(uint(int32(2))%32)) & v178
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+9)) = uint8(v199)
	v204 = int32(base.Ui32(v29)>>(uint(int32(4))%32)) & v178
	*(*uint8)(unsafe.Add(mBase, uint32(l3)+8)) = uint8(v204)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v176
	m.G0 = v15 + int32(48)
	return
L2:
	;
	if v19&int32(1) != 0 {
		goto L49
	} else {
		goto L50
	}
L3:
	;
	switch v18&int32(255) - int32(99) {
	case 0:
		v140 = int32(-1)
		v141 = v109
		goto L41
	case 1:
		goto L42
	default:
		goto L45
	case 6:
		goto L43
	case 16:
		goto L44
	}
L4:
	;
	v106 = int32(-1)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v107 != 0 {
		v143 = v106
		v145 = v104
		v146 = v105
		goto L2
	} else {
		goto L40
	}
L5:
	;
	if v29&int32(80) != 0 {
		v176 = v11
		v177 = v56
		goto L1
	} else {
		goto L39
	}
L6:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v99 = int32(base.Ui32(v95) >> (uint(int32(2)) % 32))
	goto L5
L7:
	;
	if v29&int32(81) != 0 {
		v176 = v11
		v177 = v89
		goto L1
	} else {
		goto L37
	}
L8:
	;
	if v19&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if int32(0) < v20 {
		v89 = v56
		v90 = v20 + v22
		goto L7
	} else {
		goto L23
	}
L10:
	;
	if base.I32_popcnt(v20) != int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v56 = base.I64_extend_i32_u(v22)
	goto L9
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	switch base.I32_ctz(v20) {
	case 0:
		goto L18
	case 1:
		goto L17
	case 2:
		goto L16
	case 3:
		goto L15
	default:
		goto L13
	}
L15:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v56 = v41
	goto L9
L16:
	;
	v40 = int64(*(*int32)(unsafe.Add(mBase, uint32(v22))))
	v56 = v40
	goto L9
L17:
	;
	v39 = int64(*(*int16)(unsafe.Add(mBase, uint32(v22))))
	v56 = v39
	goto L9
L18:
	;
	v38 = int64(*(*int8)(unsafe.Add(mBase, uint32(v22))))
	v56 = v38
	goto L9
L19:
	;
	return
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20
	F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_range_deserialize_1), int32(123), int32(_a_F_range_deserialize_2))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	if v20 == int32(-1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v62 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v85 = F_strlen(m, v22)
	mBase = m.M
	v89 = v56
	v90 = v85 + v22 + int32(1)
	goto L7
L27:
	;
	v66 = int32(18)
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+9)))
	if v68 == v66 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v62&int32(1) == int32(0) {
		goto L6
	} else {
		goto L36
	}
L30:
	;
	v71 = v66
	goto L32
L31:
	;
	v71 = int32(2)
	goto L32
L32:
	;
	if base.Ui32((v68-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v78 = int32(6)
	goto L35
L34:
	;
	v78 = v71
	goto L35
L35:
	;
	v99 = v78
	goto L5
L36:
	;
	v99 = int32(base.Ui32(v62) >> (uint(int32(1)) % 32))
	goto L5
L37:
	;
	if v20 == int32(-1) {
		v104 = v90
		v105 = v89
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v108 = v20
	v109 = v90
	v110 = v89
	goto L3
L39:
	;
	v104 = v99 + v22
	v105 = v56
	goto L4
L40:
	;
	v108 = v106
	v109 = v104
	v110 = v105
	goto L3
L41:
	;
	v143 = v108
	v145 = v140 & v141
	v146 = v110
	goto L2
L42:
	;
	v140 = int32(-8)
	v141 = v109 + int32(7)
	goto L41
L43:
	;
	v140 = int32(-4)
	v141 = v109 + int32(3)
	goto L41
L44:
	;
	v140 = int32(-2)
	v141 = v109 + int32(1)
	goto L41
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	F_errmsg_internal(m, int32(_a_F_range_deserialize_3), v15+int32(16))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L19
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_range_deserialize_1), int32(322), int32(_a_F_range_deserialize_4))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L19
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	if base.I32_popcnt(v143) != int32(1) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	goto L51
L51:
	;
	v176 = base.I64_extend_i32_u(v145)
	v177 = v146
	goto L1
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L19
	} else {
		goto L58
	}
L53:
	;
	switch base.I32_ctz(v143) {
	case 0:
		goto L57
	case 1:
		goto L56
	case 2:
		goto L55
	case 3:
		goto L54
	default:
		goto L52
	}
L54:
	;
	v156 = *(*int64)(unsafe.Add(mBase, uint32(v145)))
	v176 = v156
	v177 = v146
	goto L1
L55:
	;
	v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v145))))
	v176 = v155
	v177 = v146
	goto L1
L56:
	;
	v154 = int64(*(*int16)(unsafe.Add(mBase, uint32(v145))))
	v176 = v154
	v177 = v146
	goto L1
L57:
	;
	v153 = int64(*(*int8)(unsafe.Add(mBase, uint32(v145))))
	v176 = v153
	v177 = v146
	goto L1
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v143
	F_errmsg_internal(m, int32(_a_F_range_deserialize_0), v15+int32(32))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L19
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_range_deserialize_1), int32(123), int32(_a_F_range_deserialize_2))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L19
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_range_empty(m *base.Module, l0 int32) int64 {
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
	var v13 int64
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v13 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v3+int32(base.Ui32(v7)>>(uint(int32(2))%32))-int32(1)))))
		return v13 & int64(1)
	}
}
func F_range_fast_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = base.I32_wrap_i64(l0)
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = base.I32_wrap_i64(l1)
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v24 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v29 = F_lookup_type_cache(m, v27, int32(2048))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	v32 = v24
	goto L6
L6:
	;
	F_range_deserialize(m, v32, v17, v14-int32(-64), v14+int32(32), v14+int32(15))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v29
	v32 = v29
	goto L6
L8:
	;
	F_range_deserialize(m, v32, v22, v14+int32(48), v14+int32(16), v14+int32(14))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)))
	v52 = int32(1)
	if v49 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v55 = v49&v50 - v52
	goto L12
L11:
	;
	v55 = v52
	goto L12
L12:
	;
	if v49|v50&int32(1) != 0 {
		v184 = v55
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if v17 != v16 {
		goto L98
	} else {
		goto L99
	}
L14:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
	if v60 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
	if v122 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L16:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+74)))
	if v59&int32(1) != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v59&int32(1) != 0 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v66 == v63 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v74 = int32(1)
	if v63&v74 != 0 {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	v69 = int32(1)
	if v63&v69 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v72 = int32(-1)
	goto L25
L24:
	;
	v72 = v69
	goto L25
L25:
	;
	v184 = v72
	goto L13
L26:
	;
	v77 = int32(-1)
	goto L28
L27:
	;
	v77 = v74
	goto L28
L28:
	;
	v184 = v77
	goto L13
L29:
	;
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v82 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
	v89 = F_FunctionCall2Coll(m, v32+int32(212), v86, v87, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L35
	}
L32:
	;
	v83 = int32(1)
	goto L34
L33:
	;
	v83 = int32(-1)
	goto L34
L34:
	;
	v184 = v83
	goto L13
L35:
	;
	v91 = base.I32_wrap_i64(v89)
	if v91 != 0 {
		v184 = v91
		goto L13
	} else {
		goto L36
	}
L36:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+57)))
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+73)))
	if v93 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+74)))
	if v92&int32(1) == int32(0) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if v92&int32(1) != 0 {
		goto L15
	} else {
		goto L50
	}
L40:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v101 == v96 {
		goto L15
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v108 = int32(1)
	if v96&v108 != 0 {
		goto L47
	} else {
		goto L48
	}
L43:
	;
	v103 = int32(1)
	if v96&v103 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v107 = v103
	goto L46
L45:
	;
	v107 = int32(-1)
	goto L46
L46:
	;
	v184 = v107
	goto L13
L47:
	;
	v112 = v108
	goto L49
L48:
	;
	v112 = int32(-1)
	goto L49
L49:
	;
	v184 = v112
	goto L13
L50:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+58)))
	if v117 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v118 = int32(-1)
	goto L53
L52:
	;
	v118 = int32(1)
	goto L53
L53:
	;
	v184 = v118
	goto L13
L54:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+42)))
	if v121&int32(1) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if v121&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+26)))
	if v128 == v125 {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	v137 = int32(1)
	if v125&v137 != 0 {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	v184 = int32(0)
	goto L13
L61:
	;
	goto L62
L62:
	;
	v132 = int32(1)
	if v125&v132 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v135 = int32(-1)
	goto L65
L64:
	;
	v135 = v132
	goto L65
L65:
	;
	v184 = v135
	goto L13
L66:
	;
	v140 = int32(-1)
	goto L68
L67:
	;
	v140 = v137
	goto L68
L68:
	;
	v184 = v140
	goto L13
L69:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+26)))
	if v145 != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	goto L71
L71:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	v150 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
	v151 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
	v152 = F_FunctionCall2Coll(m, v32+int32(212), v149, v150, v151)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L75
	}
L72:
	;
	v146 = int32(1)
	goto L74
L73:
	;
	v146 = int32(-1)
	goto L74
L74:
	;
	v184 = v146
	goto L13
L75:
	;
	v154 = base.I32_wrap_i64(v152)
	if v154 != 0 {
		v184 = v154
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+25)))
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+41)))
	if v156 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+42)))
	if v155&int32(1) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	if v155&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L80:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+26)))
	if v164 == v159 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	v172 = int32(1)
	if v159&v172 != 0 {
		goto L89
	} else {
		goto L90
	}
L83:
	;
	v184 = int32(0)
	goto L13
L84:
	;
	goto L85
L85:
	;
	v167 = int32(1)
	if v159&v167 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v171 = v167
	goto L88
L87:
	;
	v171 = int32(-1)
	goto L88
L88:
	;
	v184 = v171
	goto L13
L89:
	;
	v176 = v172
	goto L91
L90:
	;
	v176 = int32(-1)
	goto L91
L91:
	;
	v184 = v176
	goto L13
L92:
	;
	v184 = int32(0)
	goto L13
L93:
	;
	goto L94
L94:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+26)))
	if v182 != 0 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v183 = int32(-1)
	goto L97
L96:
	;
	v183 = int32(1)
	goto L97
L97:
	;
	v184 = v183
	goto L13
L98:
	;
	F_pfree(m, v17)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	if v22 != v21 {
		goto L102
	} else {
		goto L103
	}
L101:
	;
	goto L100
L102:
	;
	F_pfree(m, v22)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	m.G0 = v14 + int32(80)
	return v184
L105:
	;
	goto L104
}
func F_range_gist_penalty(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v83 float32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v96 float32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 float32
	_ = v102
	var v103 int32
	_ = v103
	var v104 float32
	_ = v104
	var v105 float32
	_ = v105
	var v106 float32
	_ = v106
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v119 float32
	_ = v119
	var v122 int32
	_ = v122
	var v123 float32
	_ = v123
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 float32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v154 float64
	_ = v154
	var v155 float64
	_ = v155
	var v158 float64
	_ = v158
	var v162 float32
	_ = v162
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 float32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int64
	_ = v189
	var v190 int64
	_ = v190
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v197 float64
	_ = v197
	var v199 float32
	_ = v199
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 float64
	_ = v218
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v252 int64
	_ = v252
	var v253 int32
	_ = v253
	var v254 float64
	_ = v254
	var v255 float64
	_ = v255
	var v258 float64
	_ = v258
	var v262 float64
	_ = v262
	var v269 float32
	_ = v269
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	v12 = m.G0
	v14 = v12 - int32(80)
	m.G0 = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int64(0)
	} else {
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v25 = F_pg_detoast_datum(m, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
			if v27 == v28 {
				v32 = base.I32_wrap_i64(v16)
				v33 = F_range_get_typcache(m, l0, v27)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v33)+272))
					F_range_deserialize(m, v33, v20, v14-int32(-64), v14+int32(32), v14+int32(15))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						F_range_deserialize(m, v33, v25, v14+int32(48), v14+int32(16), v14+int32(14))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int64(0)
						} else {
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)))
							if v52 == int32(1) {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
								if v55 != 0 {
									v269 = float32(0)
								} else {
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									v62 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+int32(base.Ui32(v56)>>(uint(int32(2))%32))-int32(1)))))
									if v62&int32(-127) != 0 {
										v269 = float32(1)
									} else {
										v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
										v67 = int32(1)
										v69 = int32(0)
										v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
										if base.B2i32(v66&v67 == v69)|base.B2i32(v71 != v67) == v69 {
											v269 = float32(2)
										} else {
											if (v66|v71)&int32(1) != 0 {
												v83 = float32(3)
											} else {
												v83 = float32(4)
											}
											v269 = v83
										}
									}
								}
								*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
								m.G0 = v14 + int32(80)
								return v16 & int64(4294967295)
							} else {
								v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
								v85 = int32(1)
								v87 = int32(0)
								v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+56)))
								if base.B2i32(v84&v85 == v87)|base.B2i32(v89 != v85) == v87 {
									v96 = float32(2)
									v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
									v101 = v99 & int32(1)
									if v101 != 0 {
										v102 = v96
									} else {
										v102 = float32(4)
									}
									v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
									if v103 != 0 {
										v104 = v96
									} else {
										v104 = v102
									}
									if v101 != 0 {
										v105 = float32(0)
									} else {
										v105 = v104
									}
									if v103 != 0 {
										v106 = v105
									} else {
										v106 = v104
									}
									*(*float32)(unsafe.Add(mBase, uint32(v32))) = v106
									v108 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
									v114 = int32(*(*int8)(unsafe.Add(mBase, uint32(v20+int32(base.Ui32(v108)>>(uint(int32(2))%32))-int32(1)))))
									if v114&int32(-127) == int32(0) {
									} else {
										v119 = *(*float32)(unsafe.Add(mBase, uint32(v32)))
										v269 = base.F32_add(v119, float32(1))
										*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
									}
									m.G0 = v14 + int32(80)
									return v16 & int64(4294967295)
								} else {
									v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
									if v89 != 0 {
										v123 = math.Float32frombits(uint32(0x7f800000))
										if v122&int32(1) != 0 {
											v269 = v123
											*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
											m.G0 = v14 + int32(80)
											return v16 & int64(4294967295)
										} else {
											v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
											if v126&int32(1) == int32(0) {
												v269 = v123
												*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
												m.G0 = v14 + int32(80)
												return v16 & int64(4294967295)
											} else {
												v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
												if v131 != 0 {
													v269 = float32(0)
													*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
													m.G0 = v14 + int32(80)
													return v16 & int64(4294967295)
												} else {
													v139 = F_range_cmp_bounds(m, v33, v14+int32(16), v14+int32(32))
													mBase = m.M
													v140 = m.ExcPending
													if v140 != 0 {
														return int64(0)
													} else {
														v142 = base.B2i32(v139 <= int32(0))
														if v139 <= int32(0) {
															v143 = float32(0)
														} else {
															v143 = float32(1)
														}
														if v142|base.B2i32(v35 == int32(0)) != 0 {
															v269 = v143
															*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
															m.G0 = v14 + int32(80)
															return v16 & int64(4294967295)
														} else {
															v149 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
															v150 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
															v151 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
															v152 = F_FunctionCall2Coll(m, v33+int32(268), v149, v150, v151)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int64(0)
															} else {
																v154 = base.F64_reinterpret_i64(v152)
																v155 = float64(0)
																if base.F64_ge(v154, v155) != 0 {
																	v158 = v154
																} else {
																	v158 = v155
																}
																v269 = base.F32_demote_f64(v158)
																*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																m.G0 = v14 + int32(80)
																return v16 & int64(4294967295)
															}
														}
													}
												}
											}
										}
									} else {
										if v84&int32(1) != 0 {
											v162 = math.Float32frombits(uint32(0x7f800000))
											if v122&int32(1) != 0 {
												v269 = v162
												*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
												m.G0 = v14 + int32(80)
												return v16 & int64(4294967295)
											} else {
												v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
												if v165&int32(1) == int32(0) {
													v269 = v162
													*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
													m.G0 = v14 + int32(80)
													return v16 & int64(4294967295)
												} else {
													v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
													if v170 != 0 {
														v269 = float32(0)
														*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
														m.G0 = v14 + int32(80)
														return v16 & int64(4294967295)
													} else {
														v178 = F_range_cmp_bounds(m, v33, v14+int32(48), v14-int32(-64))
														mBase = m.M
														v179 = m.ExcPending
														if v179 != 0 {
															return int64(0)
														} else {
															v181 = base.B2i32(int32(0) <= v178)
															if int32(0) <= v178 {
																v182 = float32(0)
															} else {
																v182 = float32(1)
															}
															if v181|base.B2i32(v35 == int32(0)) != 0 {
																v269 = v182
																*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																m.G0 = v14 + int32(80)
																return v16 & int64(4294967295)
															} else {
																v188 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
																v189 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
																v190 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
																v191 = F_FunctionCall2Coll(m, v33+int32(268), v188, v189, v190)
																mBase = m.M
																v192 = m.ExcPending
																if v192 != 0 {
																	return int64(0)
																} else {
																	v193 = base.F64_reinterpret_i64(v191)
																	v194 = float64(0)
																	if base.F64_ge(v193, v194) != 0 {
																		v197 = v193
																	} else {
																		v197 = v194
																	}
																	v269 = base.F32_demote_f64(v197)
																	*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																	m.G0 = v14 + int32(80)
																	return v16 & int64(4294967295)
																}
															}
														}
													}
												}
											}
										} else {
											v199 = math.Float32frombits(uint32(0x7f800000))
											if v122&int32(1) != 0 {
												v269 = v199
												*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
												m.G0 = v14 + int32(80)
												return v16 & int64(4294967295)
											} else {
												v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+72)))
												if v202&int32(1) != 0 {
													v269 = v199
													*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
													m.G0 = v14 + int32(80)
													return v16 & int64(4294967295)
												} else {
													v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+40)))
													if v205&int32(1) != 0 {
														v269 = v199
														*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
														m.G0 = v14 + int32(80)
														return v16 & int64(4294967295)
													} else {
														v214 = F_range_cmp_bounds(m, v33, v14+int32(48), v14-int32(-64))
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return int64(0)
														} else {
															v217 = base.B2i32(int32(0) <= v214)
															if int32(0) <= v214 {
																v218 = float64(0)
															} else {
																v218 = float64(1)
															}
															v219 = int32(0)
															if v217|base.B2i32(v35 == v219) == v219 {
																v226 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
																v227 = *(*int64)(unsafe.Add(mBase, uint32(v14)+64))
																v228 = *(*int64)(unsafe.Add(mBase, uint32(v14)+48))
																v229 = F_FunctionCall2Coll(m, v33+int32(268), v226, v227, v228)
																mBase = m.M
																v230 = m.ExcPending
																if v230 != 0 {
																	return int64(0)
																} else {
																	v231 = base.F64_reinterpret_i64(v229)
																	v232 = float64(0)
																	if base.F64_ge(v231, v232) != 0 {
																		v235 = v231
																	} else {
																		v235 = v232
																	}
																	v238 = base.F64_add(v235, float64(0))
																	v243 = F_range_cmp_bounds(m, v33, v14+int32(16), v14+int32(32))
																	mBase = m.M
																	v244 = m.ExcPending
																	if v244 != 0 {
																		return int64(0)
																	} else {
																		if v243 <= int32(0) {
																			v262 = v238
																			v269 = base.F32_demote_f64(v262)
																			*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																			m.G0 = v14 + int32(80)
																			return v16 & int64(4294967295)
																		} else {
																			if v35 != 0 {
																				v249 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
																				v250 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
																				v251 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
																				v252 = F_FunctionCall2Coll(m, v33+int32(268), v249, v250, v251)
																				mBase = m.M
																				v253 = m.ExcPending
																				if v253 != 0 {
																					return int64(0)
																				} else {
																					v254 = base.F64_reinterpret_i64(v252)
																					v255 = float64(0)
																					if base.F64_ge(v254, v255) != 0 {
																						v258 = v254
																					} else {
																						v258 = v255
																					}
																					v262 = base.F64_add(v238, v258)
																					v269 = base.F32_demote_f64(v262)
																					*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																					m.G0 = v14 + int32(80)
																					return v16 & int64(4294967295)
																				}
																			} else {
																				v262 = base.F64_add(v238, float64(1))
																				v269 = base.F32_demote_f64(v262)
																				*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																				m.G0 = v14 + int32(80)
																				return v16 & int64(4294967295)
																			}
																		}
																	}
																}
															} else {
																v238 = v218
																v243 = F_range_cmp_bounds(m, v33, v14+int32(16), v14+int32(32))
																mBase = m.M
																v244 = m.ExcPending
																if v244 != 0 {
																	return int64(0)
																} else {
																	if v243 <= int32(0) {
																		v262 = v238
																		v269 = base.F32_demote_f64(v262)
																		*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																		m.G0 = v14 + int32(80)
																		return v16 & int64(4294967295)
																	} else {
																		if v35 != 0 {
																			v249 = *(*int32)(unsafe.Add(mBase, uint32(v33)+208))
																			v250 = *(*int64)(unsafe.Add(mBase, uint32(v14)+16))
																			v251 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
																			v252 = F_FunctionCall2Coll(m, v33+int32(268), v249, v250, v251)
																			mBase = m.M
																			v253 = m.ExcPending
																			if v253 != 0 {
																				return int64(0)
																			} else {
																				v254 = base.F64_reinterpret_i64(v252)
																				v255 = float64(0)
																				if base.F64_ge(v254, v255) != 0 {
																					v258 = v254
																				} else {
																					v258 = v255
																				}
																				v262 = base.F64_add(v238, v258)
																				v269 = base.F32_demote_f64(v262)
																				*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																				m.G0 = v14 + int32(80)
																				return v16 & int64(4294967295)
																			}
																		} else {
																			v262 = base.F64_add(v238, float64(1))
																			v269 = base.F32_demote_f64(v262)
																			*(*float32)(unsafe.Add(mBase, uint32(v32))) = v269
																			m.G0 = v14 + int32(80)
																			return v16 & int64(4294967295)
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
				v286 = m.ExcPending
				if v286 != 0 {
					return int64(0)
				} else {
					F_errmsg_internal(m, int32(_a_F_range_gist_penalty_0), int32(0))
					mBase = m.M
					v290 = m.ExcPending
					if v290 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_range_gist_penalty_1), int32(379), int32(_a_F_range_gist_penalty_2))
						mBase = m.M
						v295 = m.ExcPending
						if v295 != 0 {
							return int64(0)
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
func F_range_gist_picksplit(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v27 float32
	_ = v27
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int64
	_ = v67
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
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
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v374 float32
	_ = v374
	var v375 int32
	_ = v375
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v408 float32
	_ = v408
	var v410 float32
	_ = v410
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v468 int32
	_ = v468
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 float32
	_ = v551
	var v557 int32
	_ = v557
	var v558 int64
	_ = v558
	var v559 int64
	_ = v559
	var v560 int64
	_ = v560
	var v561 int32
	_ = v561
	var v562 float64
	_ = v562
	var v563 float64
	_ = v563
	var v566 float64
	_ = v566
	var v571 float32
	_ = v571
	var v576 int32
	_ = v576
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 float32
	_ = v590
	var v592 float32
	_ = v592
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v621 float32
	_ = v621
	var v623 float32
	_ = v623
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v662 float32
	_ = v662
	var v664 float32
	_ = v664
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v728 int32
	_ = v728
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v769 int32
	_ = v769
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v809 int32
	_ = v809
	var v811 float32
	_ = v811
	var v817 int32
	_ = v817
	var v818 int64
	_ = v818
	var v819 int64
	_ = v819
	var v820 int64
	_ = v820
	var v821 int32
	_ = v821
	var v822 float64
	_ = v822
	var v823 float64
	_ = v823
	var v826 float64
	_ = v826
	var v831 float32
	_ = v831
	var v836 int32
	_ = v836
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v850 float32
	_ = v850
	var v852 float32
	_ = v852
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v873 int32
	_ = v873
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v966 int64
	_ = v966
	var v967 int64
	_ = v967
	var v968 int64
	_ = v968
	var v969 int32
	_ = v969
	var v970 float64
	_ = v970
	var v971 float64
	_ = v971
	var v974 float64
	_ = v974
	var v975 int32
	_ = v975
	var v976 int64
	_ = v976
	var v977 int64
	_ = v977
	var v978 int64
	_ = v978
	var v979 int32
	_ = v979
	var v980 float64
	_ = v980
	var v981 float64
	_ = v981
	var v984 float64
	_ = v984
	var v988 float64
	_ = v988
	var v992 int32
	_ = v992
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1025 int32
	_ = v1025
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int64
	_ = v1176
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1218 int32
	_ = v1218
	var v1225 int64
	_ = v1225
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1245 int32
	_ = v1245
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1260 int32
	_ = v1260
	var v1262 int32
	_ = v1262
	var v1266 int32
	_ = v1266
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1322 int32
	_ = v1322
	var v1323 int32
	_ = v1323
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1346 int32
	_ = v1346
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1357 int32
	_ = v1357
	var v1394 int64
	_ = v1394
	var v1395 int64
	_ = v1395
	var v1401 int32
	_ = v1401
	var v1405 int32
	_ = v1405
	var v1440 int32
	_ = v1440
	v2 = int32(0)
	v27 = float32(0)
	v34 = m.G0
	v36 = v34 - int32(112)
	m.G0 = v36
	v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v41 = F_pg_detoast_datum(m, v40)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v46 = F_range_get_typcache(m, l0, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v48 = int32(1)
	v49 = base.I32_wrap_i64(v38)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v54 = (v50 - v48) & int32(_a_F_range_gist_picksplit_0)
	v58 = v54<<(uint(v48)%32) + int32(2)
	v59 = F_palloc(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v59
	v62 = F_palloc(m, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v62
	v65 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+80)) = v65
	v67 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+72)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v36)+64)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v36)+56)) = v67
	*(*int64)(unsafe.Add(mBase, uint32(v36)+48)) = v67
	v76 = v39 + int32(8)
	if v50&int32(_a_F_range_gist_picksplit_0) != int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v82 = v48
	goto L9
L7:
	;
	v166 = v65
	v167 = v2
	v168 = v2
	v170 = v2
	v172 = v2
	v174 = v2
	v175 = v2
	v177 = v2
	v178 = v2
	goto L8
L8:
	;
	v209 = int32(0)
	v211 = base.B2i32(v166 <= v209)
	if v166 <= v209 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v76+v82*int32(24))))
	v121 = F_pg_detoast_datum(m, v120)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v36)+80))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v36)+76))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v36)+72))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v36)+68))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v36)+60))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v36)+52))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v36)+64))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v36)+56))
	v166 = v162
	v167 = v164
	v168 = v158
	v170 = v157
	v172 = v163
	v174 = v160
	v175 = v156
	v177 = v159
	v178 = v161
	goto L8
L11:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	v129 = int32(*(*int8)(unsafe.Add(mBase, uint32(v121+int32(base.Ui32(v123)>>(uint(int32(2))%32))-int32(1)))))
	goto L12
L12:
	;
	if v129&int32(1) != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v143 = int32(8)
	goto L15
L14:
	;
	v133 = int32(3)
	v136 = int32(base.Ui32(v129)>>(uint(v133)%32)) & v133
	if v129 < int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v146 = v36 + int32(48) + v143<<(uint(int32(2))%32)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	v148 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v147 + v148
	v154 = (v82 + v148) & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v154) <= base.Ui32(v54) {
		v82 = v154
		goto L9
	} else {
		goto L19
	}
L16:
	;
	v141 = v136 | int32(4)
	goto L18
L17:
	;
	v141 = v136
	goto L18
L18:
	;
	v143 = v141
	goto L15
L19:
	;
	goto L10
L20:
	;
	v212 = int32(-1)
	goto L22
L21:
	;
	v212 = v209
	goto L22
L22:
	;
	v213 = int32(0)
	v215 = base.B2i32(v213 < v166)
	if v213 < v166 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v216 = v166
	goto L25
L24:
	;
	v216 = v213
	goto L25
L25:
	;
	v217 = base.B2i32(v216 < v178)
	if v216 < v178 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v218 = int32(1)
	goto L28
L27:
	;
	v218 = v212
	goto L28
L28:
	;
	v220 = base.B2i32(int32(0) < v178)
	if int32(0) < v178 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v221 = v218
	goto L31
L30:
	;
	v221 = v212
	goto L31
L31:
	;
	if v216 < v178 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v222 = v178
	goto L34
L33:
	;
	v222 = v216
	goto L34
L34:
	;
	if int32(0) < v178 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v223 = v222
	goto L37
L36:
	;
	v223 = v216
	goto L37
L37:
	;
	v224 = base.B2i32(v223 < v167)
	if v223 < v167 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v225 = int32(2)
	goto L40
L39:
	;
	v225 = v221
	goto L40
L40:
	;
	v227 = base.B2i32(int32(0) < v167)
	if int32(0) < v167 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v228 = v225
	goto L43
L42:
	;
	v228 = v221
	goto L43
L43:
	;
	if v223 < v167 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v229 = v167
	goto L46
L45:
	;
	v229 = v223
	goto L46
L46:
	;
	if int32(0) < v167 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v230 = v229
	goto L49
L48:
	;
	v230 = v223
	goto L49
L49:
	;
	v231 = base.B2i32(v230 < v174)
	if v230 < v174 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v232 = int32(3)
	goto L52
L51:
	;
	v232 = v228
	goto L52
L52:
	;
	v234 = base.B2i32(int32(0) < v174)
	if int32(0) < v174 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v235 = v232
	goto L55
L54:
	;
	v235 = v228
	goto L55
L55:
	;
	if v230 < v174 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v236 = v174
	goto L58
L57:
	;
	v236 = v230
	goto L58
L58:
	;
	if int32(0) < v174 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v237 = v236
	goto L61
L60:
	;
	v237 = v230
	goto L61
L61:
	;
	v238 = base.B2i32(v237 < v172)
	if v237 < v172 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v239 = int32(4)
	goto L64
L63:
	;
	v239 = v235
	goto L64
L64:
	;
	v241 = base.B2i32(int32(0) < v172)
	if int32(0) < v172 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v242 = v239
	goto L67
L66:
	;
	v242 = v235
	goto L67
L67:
	;
	if v237 < v172 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v243 = v172
	goto L70
L69:
	;
	v243 = v237
	goto L70
L70:
	;
	if int32(0) < v172 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v244 = v243
	goto L73
L72:
	;
	v244 = v237
	goto L73
L73:
	;
	v245 = base.B2i32(v244 < v177)
	if v244 < v177 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v246 = int32(5)
	goto L76
L75:
	;
	v246 = v242
	goto L76
L76:
	;
	v248 = base.B2i32(int32(0) < v177)
	if int32(0) < v177 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v249 = v246
	goto L79
L78:
	;
	v249 = v242
	goto L79
L79:
	;
	if v244 < v177 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v250 = v177
	goto L82
L81:
	;
	v250 = v244
	goto L82
L82:
	;
	if int32(0) < v177 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v251 = v250
	goto L85
L84:
	;
	v251 = v244
	goto L85
L85:
	;
	v252 = base.B2i32(v251 < v168)
	if v251 < v168 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v253 = int32(6)
	goto L88
L87:
	;
	v253 = v249
	goto L88
L88:
	;
	v255 = base.B2i32(int32(0) < v168)
	if int32(0) < v168 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v256 = v253
	goto L91
L90:
	;
	v256 = v249
	goto L91
L91:
	;
	if v251 < v168 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v257 = v168
	goto L94
L93:
	;
	v257 = v251
	goto L94
L94:
	;
	if int32(0) < v168 {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v258 = v257
	goto L97
L96:
	;
	v258 = v251
	goto L97
L97:
	;
	v259 = base.B2i32(v258 < v170)
	if v258 < v170 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v260 = int32(7)
	goto L100
L99:
	;
	v260 = v256
	goto L100
L100:
	;
	v262 = base.B2i32(int32(0) < v170)
	if int32(0) < v170 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v263 = v260
	goto L103
L102:
	;
	v263 = v256
	goto L103
L103:
	;
	if v258 < v170 {
		goto L104
	} else {
		goto L105
	}
L104:
	;
	v264 = v170
	goto L106
L105:
	;
	v264 = v258
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
	v265 = v264
	goto L109
L108:
	;
	v265 = v258
	goto L109
L109:
	;
	if v265 < v175 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v267 = int32(8)
	goto L112
L111:
	;
	v267 = v263
	goto L112
L112:
	;
	v269 = base.B2i32(int32(0) < v175)
	if int32(0) < v175 {
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v270 = v267
	goto L115
L114:
	;
	v270 = v263
	goto L115
L115:
	;
	if v220+v215+v227+v234+v241+v248+v255+v262+v269 == int32(1) {
		goto L119
	} else {
		goto L120
	}
L116:
	;
	m.G0 = v36 + int32(112)
	return v38 & int64(4294967295)
L117:
	;
	F_range_gist_fallback_split(m, v46, v39, v49)
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L1
	} else {
		goto L336
	}
L118:
	;
	F_qsort_arg(m, v291, v290, int32(32), int32(1681), v46)
	mBase = m.M
	v1401 = m.ExcPending
	if v1401 != 0 {
		goto L1
	} else {
		goto L334
	}
L119:
	;
	switch v270 & int32(-5) {
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
	v1174 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v1174
	v1176 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+24)) = v1176
	*(*int64)(unsafe.Add(mBase, uint32(v36)+16)) = v1176
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v1176
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v1176
	if v211 == v1174 {
		goto L296
	} else {
		goto L297
	}
L122:
	;
	F_range_gist_fallback_split(m, v46, v39, v49)
	mBase = m.M
	v1173 = m.ExcPending
	if v1173 != 0 {
		goto L1
	} else {
		goto L294
	}
L123:
	;
	F_range_gist_single_sorting_split(m, v46, v39, v49, int32(0))
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L1
	} else {
		goto L293
	}
L124:
	;
	F_range_gist_single_sorting_split(m, v46, v39, v49, int32(1))
	mBase = m.M
	v1168 = m.ExcPending
	if v1168 != 0 {
		goto L1
	} else {
		goto L292
	}
L125:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v46)+272))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v287 = int32(_a_F_range_gist_picksplit_0)
	v288 = v286 + v287
	v290 = v288 & v287
	v291 = F_palloc_mul(m, int32(32), v290)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v294 = F_palloc_mul(m, int32(32), v290)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	if v286&int32(_a_F_range_gist_picksplit_0) == int32(1) {
		goto L118
	} else {
		goto L128
	}
L128:
	;
	v301 = v290 - int32(1)
	v303 = int32(1)
	goto L129
L129:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v76+v303*int32(24))))
	v339 = F_pg_detoast_datum(m, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L131
	}
L130:
	;
	v356 = v290 << (uint(int32(5)) % 32)
	if v356 != 0 {
		goto L134
	} else {
		goto L135
	}
L131:
	;
	v343 = v291 + v303<<(uint(int32(5))%32)
	F_range_deserialize(m, v46, v339, v343-int32(32), v343-int32(16), v36)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	v353 = (v303 + int32(1)) & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v353) <= base.Ui32(v290) {
		v303 = v353
		goto L129
	} else {
		goto L133
	}
L133:
	;
	goto L130
L134:
	;
	base.MemoryCopy(m, v294, v291, v356)
	goto L136
L135:
	;
	goto L136
L136:
	;
	F_qsort_arg(m, v291, v290, int32(32), int32(1681), v46)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	F_qsort_arg(m, v294, v290, int32(32), int32(1682), v46)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v367 = v46 + int32(268)
	v368 = int32(1)
	v369 = int32(base.Ui32(v290) >> (uint(v368) % 32))
	v373 = int32(base.Ui32(v290+v368) >> (uint(v368) % 32))
	v374 = base.F32_convert_i32_u(v290)
	v375 = int32(0)
	v382 = v294
	v383 = v291
	v384 = v375
	v388 = v375
	v392 = v368
	v393 = v375
	v394 = v375
	v399 = v375
	v408 = v27
	v410 = v27
	goto L139
L139:
	;
	v415 = v382
	v421 = v388
	goto L142
L140:
	;
	v628 = v301 << (uint(int32(5)) % 32)
	v630 = int32(16)
	v635 = v301
	v636 = v291 + v628 + v630
	v637 = v628 + v294 + v630
	v638 = v301
	v646 = v605
	v647 = v606
	v648 = v607
	v653 = v612
	v662 = v621
	v664 = v623
	goto L184
L141:
	;
	goto L140
L142:
	;
	v449 = v291 + v421<<(uint(int32(5))%32)
	v450 = F_range_cmp_bounds(m, v46, v383, v449)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	if v290 <= v384 {
		v513 = v384
		goto L153
	} else {
		goto L154
	}
L144:
	;
	if v450 == int32(0) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v455 = v449 + int32(16)
	v456 = F_range_cmp_bounds(m, v46, v455, v415)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	goto L143
L148:
	;
	if int32(0) < v456 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v460 = v455
	goto L151
L150:
	;
	v460 = v415
	goto L151
L151:
	;
	v462 = v421 + int32(1)
	if v462 < v290 {
		v415 = v460
		v421 = v462
		goto L142
	} else {
		goto L152
	}
L152:
	;
	v605 = v392
	v606 = v393
	v607 = v394
	v612 = v399
	v621 = v408
	v623 = v410
	goto L141
L153:
	;
	if v513 < v369 {
		goto L161
	} else {
		goto L162
	}
L154:
	;
	v468 = v384
	goto L155
L155:
	;
	v503 = F_range_cmp_bounds(m, v46, v294+v468<<(uint(int32(5))%32)+int32(16), v415)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L157
	}
L156:
	;
	v513 = v290
	goto L153
L157:
	;
	if int32(0) < v503 {
		v513 = v468
		goto L153
	} else {
		goto L158
	}
L158:
	;
	v508 = v468 + int32(1)
	if v508 != v290 {
		v468 = v508
		goto L155
	} else {
		goto L159
	}
L159:
	;
	goto L156
L160:
	;
	if v421 < v290 {
		v382 = v415
		v383 = v449
		v384 = v513
		v388 = v421
		v392 = v585
		v393 = v586
		v394 = v587
		v399 = v588
		v408 = v590
		v410 = v592
		goto L139
	} else {
		goto L183
	}
L161:
	;
	v544 = v513
	goto L163
L162:
	;
	v544 = v369
	goto L163
L163:
	;
	if v421 < v373 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v546 = v544
	goto L166
L165:
	;
	v546 = v421
	goto L166
L166:
	;
	v547 = v290 - v546
	if v546 < v547 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v549 = v546
	goto L169
L168:
	;
	v549 = v547
	goto L169
L169:
	;
	v551 = base.F32_div(base.F32_convert_i32_s(v549), v374)
	if base.F64_gt(base.F64_promote_f32(v551), float64(0.3)) == int32(0) {
		v585 = v392
		v586 = v393
		v587 = v394
		v588 = v399
		v590 = v408
		v592 = v410
		goto L160
	} else {
		goto L170
	}
L170:
	;
	if v284 != 0 {
		goto L172
	} else {
		goto L173
	}
L171:
	;
	if base.F32_lt(v571, v408)|v392 == int32(0) {
		goto L179
	} else {
		goto L180
	}
L172:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v46)+208))
	v558 = *(*int64)(unsafe.Add(mBase, uint32(v415)))
	v559 = *(*int64)(unsafe.Add(mBase, uint32(v449)))
	v560 = F_FunctionCall2Coll(m, v367, v557, v558, v559)
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L175
	}
L173:
	;
	goto L174
L174:
	;
	v571 = base.F32_convert_i32_s(v513 - v421)
	goto L171
L175:
	;
	v562 = base.F64_reinterpret_i64(v560)
	v563 = float64(0)
	if base.F64_ge(v562, v563) != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v566 = v562
	goto L178
L177:
	;
	v566 = v563
	goto L178
L178:
	;
	v571 = base.F32_demote_f64(v566)
	goto L171
L179:
	;
	v576 = int32(0)
	if base.B2i32(base.F32_gt(v551, v410) == v576)|base.F32_ne(v408, v571) != 0 {
		v585 = v576
		v586 = v393
		v587 = v394
		v588 = v399
		v590 = v408
		v592 = v410
		goto L160
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v585 = int32(0)
	v586 = v415
	v587 = v449
	v588 = v513 - v546
	v590 = v571
	v592 = v551
	goto L160
L182:
	;
	goto L181
L183:
	;
	v605 = v585
	v606 = v586
	v607 = v587
	v612 = v588
	v621 = v590
	v623 = v592
	goto L141
L184:
	;
	v668 = v635
	v669 = v636
	goto L187
L185:
	;
	if v866 != 0 {
		goto L117
	} else {
		goto L231
	}
L186:
	;
	goto L185
L187:
	;
	v703 = v294 + v668<<(uint(int32(5))%32)
	v705 = v703 + int32(16)
	v706 = F_range_cmp_bounds(m, v46, v637, v705)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L1
	} else {
		goto L189
	}
L188:
	;
	if v638 < int32(0) {
		v769 = v638
		goto L198
	} else {
		goto L199
	}
L189:
	;
	if v706 == int32(0) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v710 = F_range_cmp_bounds(m, v46, v703, v669)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L1
	} else {
		goto L193
	}
L191:
	;
	goto L192
L192:
	;
	goto L188
L193:
	;
	if v710 < int32(0) {
		goto L194
	} else {
		goto L195
	}
L194:
	;
	v714 = v703
	goto L196
L195:
	;
	v714 = v669
	goto L196
L196:
	;
	if int32(0) < v668 {
		v668 = v668 - int32(1)
		v669 = v714
		goto L187
	} else {
		goto L197
	}
L197:
	;
	v866 = v646
	v867 = v647
	v868 = v648
	v873 = v653
	goto L186
L198:
	;
	v800 = v668 + int32(1)
	if v800 < v369 {
		goto L208
	} else {
		goto L209
	}
L199:
	;
	v728 = v638
	goto L200
L200:
	;
	v757 = F_range_cmp_bounds(m, v46, v291+v728<<(uint(int32(5))%32), v669)
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L202
	}
L201:
	;
	v769 = int32(-1)
	goto L198
L202:
	;
	if v757 < int32(0) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v769 = v728
	goto L198
L204:
	;
	goto L205
L205:
	;
	if int32(0) < v728 {
		v728 = v728 - int32(1)
		goto L200
	} else {
		goto L206
	}
L206:
	;
	goto L201
L207:
	;
	if int32(0) <= v668 {
		v635 = v668
		v636 = v669
		v637 = v705
		v638 = v769
		v646 = v845
		v647 = v846
		v648 = v847
		v653 = v848
		v662 = v850
		v664 = v852
		goto L184
	} else {
		goto L230
	}
L208:
	;
	v802 = v800
	goto L210
L209:
	;
	v802 = v369
	goto L210
L210:
	;
	v804 = v769 + int32(1)
	if v804 < v373 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v806 = v802
	goto L213
L212:
	;
	v806 = v804
	goto L213
L213:
	;
	v807 = v290 - v806
	if v806 < v807 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v809 = v806
	goto L216
L215:
	;
	v809 = v807
	goto L216
L216:
	;
	v811 = base.F32_div(base.F32_convert_i32_s(v809), v374)
	if base.F64_gt(base.F64_promote_f32(v811), float64(0.3)) == int32(0) {
		v845 = v646
		v846 = v647
		v847 = v648
		v848 = v653
		v850 = v662
		v852 = v664
		goto L207
	} else {
		goto L217
	}
L217:
	;
	if v284 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	if base.F32_lt(v831, v662)|v646 == int32(0) {
		goto L226
	} else {
		goto L227
	}
L219:
	;
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v46)+208))
	v818 = *(*int64)(unsafe.Add(mBase, uint32(v705)))
	v819 = *(*int64)(unsafe.Add(mBase, uint32(v669)))
	v820 = F_FunctionCall2Coll(m, v367, v817, v818, v819)
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v831 = base.F32_convert_i32_s(v668 - v769)
	goto L218
L222:
	;
	v822 = base.F64_reinterpret_i64(v820)
	v823 = float64(0)
	if base.F64_ge(v822, v823) != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v826 = v822
	goto L225
L224:
	;
	v826 = v823
	goto L225
L225:
	;
	v831 = base.F32_demote_f64(v826)
	goto L218
L226:
	;
	v836 = int32(0)
	if base.B2i32(base.F32_gt(v811, v664) == v836)|base.F32_ne(v662, v831) != 0 {
		v845 = v836
		v846 = v647
		v847 = v648
		v848 = v653
		v850 = v662
		v852 = v664
		goto L207
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v845 = int32(0)
	v846 = v705
	v847 = v669
	v848 = v800 - v806
	v850 = v831
	v852 = v811
	goto L207
L229:
	;
	goto L228
L230:
	;
	v866 = v845
	v867 = v846
	v868 = v847
	v873 = v848
	goto L186
L231:
	;
	v889 = F_palloc_mul(m, int32(2), v290)
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v889
	v893 = F_palloc_mul(m, int32(2), v290)
	mBase = m.M
	v894 = m.ExcPending
	if v894 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	v895 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v895
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v893
	v901 = F_palloc_mul(m, int32(16), v290)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	v903 = int32(1)
	v905 = int32(0)
	v908 = v903
	v909 = v903
	v910 = v905
	v913 = v905
	v917 = v905
	goto L235
L235:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v76+v909*int32(24))))
	v945 = F_pg_detoast_datum(m, v944)
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L1
	} else {
		goto L237
	}
L236:
	;
	if int32(0) < v1028 {
		goto L270
	} else {
		goto L271
	}
L237:
	;
	v948 = v36 + int32(96)
	F_range_deserialize(m, v46, v945, v36, v948, v36+int32(95))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	v953 = F_range_cmp_bounds(m, v46, v948, v867)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	v1031 = v908 + int32(1)
	v1032 = int32(_a_F_range_gist_picksplit_0)
	v1033 = v1031 & v1032
	if base.Ui32(v1033) <= base.Ui32(v288&v1032) {
		v908 = v1031
		v909 = v1033
		v910 = v1025
		v913 = v1027
		v917 = v1028
		goto L235
	} else {
		goto L269
	}
L240:
	;
	if v953 <= int32(0) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v957 = F_range_cmp_bounds(m, v46, v36, v868)
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L1
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	if v1008 <= int32(0) {
		goto L265
	} else {
		goto L266
	}
L244:
	;
	if int32(0) <= v957 {
		goto L245
	} else {
		goto L246
	}
L245:
	;
	v963 = v901 + v917<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v963))) = v909
	if v284 != 0 {
		goto L248
	} else {
		goto L249
	}
L246:
	;
	goto L247
L247:
	;
	v992 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v992 <= int32(0) {
		goto L260
	} else {
		goto L261
	}
L248:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v46)+208))
	v966 = *(*int64)(unsafe.Add(mBase, uint32(v36)))
	v967 = *(*int64)(unsafe.Add(mBase, uint32(v868)))
	v968 = F_FunctionCall2Coll(m, v367, v965, v966, v967)
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L1
	} else {
		goto L251
	}
L249:
	;
	v988 = float64(0)
	goto L250
L250:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v963)+8)) = v988
	v1025 = v910
	v1027 = v913
	v1028 = v917 + int32(1)
	goto L239
L251:
	;
	v970 = base.F64_reinterpret_i64(v968)
	v971 = float64(0)
	if base.F64_ge(v970, v971) != 0 {
		goto L252
	} else {
		goto L253
	}
L252:
	;
	v974 = v970
	goto L254
L253:
	;
	v974 = v971
	goto L254
L254:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v46)+208))
	v976 = *(*int64)(unsafe.Add(mBase, uint32(v867)))
	v977 = *(*int64)(unsafe.Add(mBase, uint32(v36)+96))
	v978 = F_FunctionCall2Coll(m, v367, v975, v976, v977)
	mBase = m.M
	v979 = m.ExcPending
	if v979 != 0 {
		goto L1
	} else {
		goto L255
	}
L255:
	;
	v980 = base.F64_reinterpret_i64(v978)
	v981 = float64(0)
	if base.F64_ge(v980, v981) != 0 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v984 = v980
	goto L258
L257:
	;
	v984 = v981
	goto L258
L258:
	;
	v988 = base.F64_sub(v974, v984)
	goto L250
L259:
	;
	v1000 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v998 + v1000
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1003+v998<<(uint(v1000)%32)))) = uint16(v908)
	v1025 = v910
	v1027 = v999
	v1028 = v917
	goto L239
L260:
	;
	v998 = v992
	v999 = v945
	goto L259
L261:
	;
	goto L262
L262:
	;
	v995 = F_range_super_union(m, v46, v913, v945)
	mBase = m.M
	v996 = m.ExcPending
	if v996 != 0 {
		goto L1
	} else {
		goto L263
	}
L263:
	;
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v998 = v997
	v999 = v995
	goto L259
L264:
	;
	v1016 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v1014 + v1016
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1019+v1014<<(uint(v1016)%32)))) = uint16(v908)
	v1025 = v1015
	v1027 = v913
	v1028 = v917
	goto L239
L265:
	;
	v1014 = v1008
	v1015 = v945
	goto L264
L266:
	;
	goto L267
L267:
	;
	v1011 = F_range_super_union(m, v46, v910, v945)
	mBase = m.M
	v1012 = m.ExcPending
	if v1012 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	v1014 = v1013
	v1015 = v1011
	goto L264
L269:
	;
	goto L236
L270:
	;
	F_pg_qsort(m, v901, v1028, int32(16), int32(1683))
	mBase = m.M
	v1042 = m.ExcPending
	if v1042 != 0 {
		goto L1
	} else {
		goto L273
	}
L271:
	;
	v1131 = v1025
	v1134 = v1027
	goto L272
L272:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = base.I64_extend_i32_u(v1131)
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = base.I64_extend_i32_u(v1134)
	goto L116
L273:
	;
	v1043 = int32(0)
	v1045 = v1043
	v1046 = v1043
	v1047 = v1025
	v1050 = v1027
	goto L274
L274:
	;
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v901+v1045<<(uint(int32(4))%32))))
	v1085 = *(*int32)(unsafe.Add(mBase, uint32(v76+v1081*int32(24))))
	v1086 = F_pg_detoast_datum(m, v1085)
	mBase = m.M
	v1087 = m.ExcPending
	if v1087 != 0 {
		goto L1
	} else {
		goto L276
	}
L275:
	;
	v1131 = v1122
	v1134 = v1123
	goto L272
L276:
	;
	if v1045 < v873 {
		goto L278
	} else {
		goto L279
	}
L277:
	;
	v1125 = v1046 + int32(1)
	v1127 = v1125 & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v1127) < base.Ui32(v1028) {
		v1045 = v1127
		v1046 = v1125
		v1047 = v1122
		v1050 = v1123
		goto L274
	} else {
		goto L291
	}
L278:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v1089 <= int32(0) {
		goto L282
	} else {
		goto L283
	}
L279:
	;
	goto L280
L280:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	if v1105 <= int32(0) {
		goto L287
	} else {
		goto L288
	}
L281:
	;
	v1097 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v1095 + v1097
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1100+v1095<<(uint(v1097)%32)))) = uint16(v1081)
	v1122 = v1047
	v1123 = v1096
	goto L277
L282:
	;
	v1095 = v1089
	v1096 = v1086
	goto L281
L283:
	;
	goto L284
L284:
	;
	v1092 = F_range_super_union(m, v46, v1050, v1086)
	mBase = m.M
	v1093 = m.ExcPending
	if v1093 != 0 {
		goto L1
	} else {
		goto L285
	}
L285:
	;
	v1094 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v1095 = v1094
	v1096 = v1092
	goto L281
L286:
	;
	v1113 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v1111 + v1113
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1116+v1111<<(uint(v1113)%32)))) = uint16(v1081)
	v1122 = v1112
	v1123 = v1050
	goto L277
L287:
	;
	v1111 = v1105
	v1112 = v1086
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1108 = F_range_super_union(m, v46, v1047, v1086)
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L1
	} else {
		goto L290
	}
L290:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	v1111 = v1110
	v1112 = v1108
	goto L286
L291:
	;
	goto L275
L292:
	;
	goto L116
L293:
	;
	goto L116
L294:
	;
	goto L116
L295:
	;
	v1238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	v1239 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v1239
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v1239
	if v1238 != int32(1) {
		goto L306
	} else {
		goto L307
	}
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = int32(1)
	goto L295
L297:
	;
	goto L298
L298:
	;
	v1190 = v166 + v178 + v167 + v174
	v1191 = v54 - v1190
	v1193 = v166 + v172 + v175
	if v1193 <= int32(0) {
		goto L299
	} else {
		goto L300
	}
L299:
	;
	v1218 = int32(0)
	if base.B2i32(v1190 <= v1218)|base.B2i32(v1191 <= v1218) == v1218 {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v1196 = v54 - v1193
	if v1196 <= int32(0) {
		goto L299
	} else {
		goto L301
	}
L301:
	;
	v1199 = v1196 - v1193
	v1200 = int32(31)
	v1201 = v1199 >> (uint(v1200) % 32)
	v1204 = v1191 - v1190
	v1206 = v1204 >> (uint(v1200) % 32)
	if v1204^v1206-v1206 < v1199^v1201-v1201 {
		goto L299
	} else {
		goto L302
	}
L302:
	;
	v1210 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+32)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v1210
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v1210
	goto L295
L303:
	;
	v1225 = int64(4294967297)
	*(*int64)(unsafe.Add(mBase, uint32(v36)+8)) = v1225
	*(*int64)(unsafe.Add(mBase, uint32(v36))) = v1225
	goto L295
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36+v270<<(uint(int32(2))%32)))) = int32(1)
	goto L295
L306:
	;
	v1245 = int32(1)
	v1250 = int32(0)
	v1253 = v1245
	v1260 = v1245
	v1262 = v1250
	v1266 = v1250
	goto L309
L307:
	;
	v1394 = int64(0)
	v1395 = int64(0)
	goto L308
L308:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v49)+32)) = v1395
	*(*int64)(unsafe.Add(mBase, uint32(v49)+8)) = v1394
	goto L116
L309:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v76+v1260*int32(24))))
	v1290 = F_pg_detoast_datum(m, v1289)
	mBase = m.M
	v1291 = m.ExcPending
	if v1291 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	v1394 = base.I64_extend_i32_u(v1353)
	v1395 = base.I64_extend_i32_u(v1352)
	goto L308
L311:
	;
	v1355 = v1253 + int32(1)
	v1357 = v1355 & int32(_a_F_range_gist_picksplit_0)
	if base.Ui32(v1357) <= base.Ui32((v1238-v1245)&int32(_a_F_range_gist_picksplit_0)) {
		v1253 = v1355
		v1260 = v1357
		v1262 = v1352
		v1266 = v1353
		goto L309
	} else {
		goto L333
	}
L312:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v1290)))
	v1298 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1290+int32(base.Ui32(v1292)>>(uint(int32(2))%32))-int32(1)))))
	goto L313
L313:
	;
	if v1298&int32(1) != 0 {
		goto L314
	} else {
		goto L315
	}
L314:
	;
	v1312 = int32(8)
	goto L316
L315:
	;
	v1302 = int32(3)
	v1305 = int32(base.Ui32(v1298)>>(uint(v1302)%32)) & v1302
	if v1298 < int32(0) {
		goto L317
	} else {
		goto L318
	}
L316:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v36+v1312<<(uint(int32(2))%32))))
	if v1316 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L317:
	;
	v1310 = v1305 | int32(4)
	goto L319
L318:
	;
	v1310 = v1305
	goto L319
L319:
	;
	v1312 = v1310
	goto L316
L320:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v1319 <= int32(0) {
		goto L324
	} else {
		goto L325
	}
L321:
	;
	goto L322
L322:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	if v1335 <= int32(0) {
		goto L329
	} else {
		goto L330
	}
L323:
	;
	v1327 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v1325 + v1327
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1330+v1325<<(uint(v1327)%32)))) = uint16(v1253)
	v1352 = v1262
	v1353 = v1326
	goto L311
L324:
	;
	v1325 = v1319
	v1326 = v1290
	goto L323
L325:
	;
	goto L326
L326:
	;
	v1322 = F_range_super_union(m, v46, v1266, v1290)
	mBase = m.M
	v1323 = m.ExcPending
	if v1323 != 0 {
		goto L1
	} else {
		goto L327
	}
L327:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	v1325 = v1324
	v1326 = v1322
	goto L323
L328:
	;
	v1343 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v1341 + v1343
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v49)+20))
	*(*uint16)(unsafe.Add(mBase, uint32(v1346+v1341<<(uint(v1343)%32)))) = uint16(v1253)
	v1352 = v1342
	v1353 = v1266
	goto L311
L329:
	;
	v1341 = v1335
	v1342 = v1290
	goto L328
L330:
	;
	goto L331
L331:
	;
	v1338 = F_range_super_union(m, v46, v1262, v1290)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L332
	}
L332:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(v49)+24))
	v1341 = v1340
	v1342 = v1338
	goto L328
L333:
	;
	goto L310
L334:
	;
	F_qsort_arg(m, v294, v290, int32(32), int32(1682), v46)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L1
	} else {
		goto L335
	}
L335:
	;
	goto L117
L336:
	;
	goto L116
}
func F_range_gt(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_range_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(int64(0) < v2))
	}
}
func F_range_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v341 int64
	_ = v341
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v24 = F_get_range_io_data(m, l0, v18, int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v27 = v17
	goto L4
L4:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v36-int32(9)))&base.B2i32(v36 != int32(32)) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v52 = v27
	v53 = int32(_a_F_range_in_0)
	v54 = int32(5)
	goto L15
L6:
	;
	v27 = v27 + int32(1)
	goto L4
L7:
	;
	goto L8
L8:
	;
	goto L5
L9:
	;
	m.G0 = v13 + int32(48)
	return v341
L10:
	;
	v297 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+18)) = uint8(v297)
	v299 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+34)) = uint8(v299)
	v304 = int32(base.Ui32(v291)>>(uint(int32(3))%32)) & v299
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+32)) = uint8(v304)
	v309 = int32(base.Ui32(v291)>>(uint(int32(2))%32)) & v299
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+17)) = uint8(v309)
	v314 = int32(base.Ui32(v291&int32(240)) >> (uint(int32(4)) % 32))
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+16)) = uint8(v314)
	v319 = int32(base.Ui32(v291)>>(uint(v299)%32)) & v299
	*(*uint8)(unsafe.Add(mBase, uint32(v13)+33)) = uint8(v319)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v328 = F_make_range(m, v321, v13+int32(24), v13+int32(8), v291&v299, v15)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L82
	}
L11:
	;
	v284 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v284)
	v341 = int64(0)
	goto L9
L12:
	;
	if v181&int32(41) != 0 {
		goto L75
	} else {
		goto L76
	}
L13:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L71
	}
L14:
	;
	if v99 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v99 = int32(0)
	goto L14
L17:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 == v58 {
		v80 = v57
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	v82 = int32(1)
	if v80 != 0 {
		v52 = v52 + v82
		v53 = v53 + v82
		v54 = v54 - v82
		goto L15
	} else {
		goto L29
	}
L21:
	;
	if base.Ui32((v57-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v68 = v57 | int32(32)
	goto L24
L23:
	;
	v68 = v57
	goto L24
L24:
	;
	if base.Ui32((v58-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v77 = v58 | int32(32)
	goto L27
L26:
	;
	v77 = v58
	goto L27
L27:
	;
	if v68 == v77 {
		v80 = v68
		goto L20
	} else {
		goto L28
	}
L28:
	;
	v99 = v68 - v77
	goto L14
L29:
	;
	goto L19
L30:
	;
	v102 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v102
	v109 = v27 + int32(5)
	goto L33
L31:
	;
	goto L32
L32:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v137 == int32(40) {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109))))
	if base.B2i32(base.Ui32(v118-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v118 == int32(32)) != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v131 = F_errsave_start(m, v15)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L41
	}
L35:
	;
	v109 = v109 + int32(1)
	goto L33
L36:
	;
	if v118 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	goto L34
L38:
	;
	v291 = int32(1)
	goto L10
L39:
	;
	goto L40
L40:
	;
	goto L37
L41:
	;
	if v131 == int32(0) {
		goto L11
	} else {
		goto L42
	}
L42:
	;
	v230 = int32(_a_F_range_in_1)
	v239 = int32(2588)
	goto L13
L43:
	;
	v223 = F_errsave_start(m, v15)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L69
	}
L44:
	;
	v217 = F_errsave_start(m, v15)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L67
	}
L45:
	;
	v211 = F_errsave_start(m, v15)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L1
	} else {
		goto L65
	}
L46:
	;
	v144 = int32(0)
	goto L48
L47:
	;
	if v137 != int32(91) {
		goto L45
	} else {
		goto L49
	}
L48:
	;
	v150 = v13 + int32(24)
	v151 = F_range_parse_bound(m, v17, v27+int32(1), v13+int32(44), v150, v15)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L50
	}
L49:
	;
	v144 = int32(2)
	goto L48
L50:
	;
	if v151 == int32(0) {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v151))))
	if v155 != int32(44) {
		goto L44
	} else {
		goto L52
	}
L52:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
	v163 = F_range_parse_bound(m, v17, v151+int32(1), v13+int32(40), v150, v15)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v163 == int32(0) {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+24)))
	v173 = v158<<(uint(int32(3))%32) | v144 | v170<<(uint(int32(4))%32)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	if v174 != int32(41) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v174 != int32(93) {
		goto L43
	} else {
		goto L58
	}
L56:
	;
	v181 = v173
	goto L57
L57:
	;
	v182 = base.I32_wrap_i64(v16)
	v184 = v163
	goto L59
L58:
	;
	v181 = v173 | int32(4)
	goto L57
L59:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	if base.B2i32(base.Ui32(v195-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v195 == int32(32)) != 0 {
		v184 = v184 + int32(1)
		goto L59
	} else {
		goto L61
	}
L60:
	;
	if v195 == int32(0) {
		goto L12
	} else {
		goto L62
	}
L61:
	;
	goto L60
L62:
	;
	v205 = F_errsave_start(m, v15)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v205 == int32(0) {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	v230 = int32(_a_F_range_in_2)
	v239 = int32(2651)
	goto L13
L65:
	;
	if v211 == int32(0) {
		goto L11
	} else {
		goto L66
	}
L66:
	;
	v230 = int32(_a_F_range_in_3)
	v239 = int32(2605)
	goto L13
L67:
	;
	if v217 == int32(0) {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v230 = int32(_a_F_range_in_4)
	v239 = int32(2620)
	goto L13
L69:
	;
	if v223 == int32(0) {
		goto L11
	} else {
		goto L70
	}
L70:
	;
	v230 = int32(_a_F_range_in_5)
	v239 = int32(2640)
	goto L13
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
	F_errmsg(m, int32(_a_F_range_in_6), v13)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v248 = F_errdetail(m, v230, int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	F_errsave_finish(m, v15, int32(_a_F_range_in_7), v239, int32(_a_F_range_in_8))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L11
L75:
	;
	if v181&int32(81) != 0 {
		v291 = v181
		goto L10
	} else {
		goto L79
	}
L76:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v262 = F_InputFunctionCallSafe(m, v24+int32(4), v258, v259, v182, v15, v13+int32(24))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v262 != 0 {
		goto L75
	} else {
		goto L78
	}
L78:
	;
	goto L11
L79:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v272 = F_InputFunctionCallSafe(m, v24+int32(4), v268, v269, v182, v15, v13+int32(8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	if v272 != 0 {
		v291 = v181
		goto L10
	} else {
		goto L81
	}
L81:
	;
	goto L11
L82:
	;
	v341 = base.I64_extend_i32_u(v328)
	goto L9
}
func F_range_lower_inf(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14367(m, l0, int64(3))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_range_ne(m *base.Module, l0 int32) int64 {
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
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
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
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v33 ^ int32(1))
					}
				} else {
					v25 = F_lookup_type_cache(m, v19, int32(2048))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int64(0)
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
						if v27 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
								F_errmsg_internal(m, int32(_a_F_range_ne_0), v9)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_range_ne_1), int32(1946), int32(_a_F_range_ne_2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int64(0)
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
								return int64(0)
							} else {
								m.G0 = v9 + int32(16)
								return base.I64_extend_i32_u(v33 ^ int32(1))
							}
						}
					}
				}
			} else {
				v25 = F_lookup_type_cache(m, v19, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v19
							F_errmsg_internal(m, int32(_a_F_range_ne_0), v9)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_range_ne_1), int32(1946), int32(_a_F_range_ne_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
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
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v33 ^ int32(1))
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
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v16 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v10)>>(uint(int32(2))%32))-int32(1)))))
	if v16&int32(1) != 0 {
		v45 = v4
		m.G0 = v8 + int32(80)
		return v45
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
		if v19 == int32(0) {
			v45 = v4
			m.G0 = v8 + int32(80)
			return v45
		} else {
			v23 = v8 - int32(-64)
			F_range_deserialize(m, l0, l1, v23, v8+int32(48), v8+int32(15))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v34 = v8 + int32(32)
				F_multirange_get_bounds(m, l0, l2, int32(0), v34, v8+int32(16))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					v39 = F_range_cmp_bounds(m, l0, v23, v34)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v45 = base.B2i32(int32(0) <= v39)
						m.G0 = v8 + int32(80)
						return v45
					}
				}
			}
		}
	}
}
func F_range_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v3)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v3)+16)) = int32(1680)
	return int64(0)
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
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	F_range_deserialize(m, l0, l1, v12-int32(-64), v12+int32(32), v12+int32(15))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		F_range_deserialize(m, l0, l2, v12+int32(48), v12+int32(16), v12+int32(14))
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
					v62 = F_datumCopy(m, base.I64_extend_i32_u(v58), int32(0), int32(-1))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return int32(0)
					} else {
						v119 = base.I32_wrap_i64(v62)
						v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
						v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
						v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
						v131 = v129 | int32(128)
						*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
						v134 = v119
						m.G0 = v12 + int32(80)
						return v134
					}
				} else {
					v134 = l2
					m.G0 = v12 + int32(80)
					return v134
				}
			} else {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+14)))
				if v53 != int32(1) {
					v69 = F_range_cmp_bounds(m, l0, v12-int32(-64), v12+int32(48))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						v72 = base.B2i32(v69 <= int32(0))
						v77 = F_range_cmp_bounds(m, l0, v12+int32(32), v12+int32(16))
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							v79 = int32(0)
							if base.B2i32(v77 < v79)|base.B2i32(v79 < v69) == v79 {
								if v38 < int32(0) {
									v134 = l1
									m.G0 = v12 + int32(80)
									return v134
								} else {
									if v45 < int32(0) {
										if v69 <= int32(0) {
											v103 = v12 - int32(-64)
										} else {
											v103 = v12 + int32(48)
										}
										if int32(0) <= v77 {
											v110 = v12 + int32(32)
										} else {
											v110 = v12 + int32(16)
										}
										v111 = int32(0)
										v113 = F_make_range(m, l0, v103, v110, v111, v111)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											if v38 < int32(0) {
												v119 = v113
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
												v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
												v131 = v129 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
												v134 = v119
											} else {
												if int32(0) <= v45 {
													v134 = v113
												} else {
													v119 = v113
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
													v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
													v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
													v131 = v129 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
													v134 = v119
												}
											}
											m.G0 = v12 + int32(80)
											return v134
										}
									} else {
										v134 = l1
										m.G0 = v12 + int32(80)
										return v134
									}
								}
							} else {
								if v72|base.B2i32(int32(0) <= v77) != 0 {
									if v69 <= int32(0) {
										v103 = v12 - int32(-64)
									} else {
										v103 = v12 + int32(48)
									}
									if int32(0) <= v77 {
										v110 = v12 + int32(32)
									} else {
										v110 = v12 + int32(16)
									}
									v111 = int32(0)
									v113 = F_make_range(m, l0, v103, v110, v111, v111)
									mBase = m.M
									v114 = m.ExcPending
									if v114 != 0 {
										return int32(0)
									} else {
										if v38 < int32(0) {
											v119 = v113
											v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
											v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
											v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
											v131 = v129 | int32(128)
											*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
											v134 = v119
										} else {
											if int32(0) <= v45 {
												v134 = v113
											} else {
												v119 = v113
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
												v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
												v131 = v129 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
												v134 = v119
											}
										}
										m.G0 = v12 + int32(80)
										return v134
									}
								} else {
									v93 = int32(0)
									if base.B2i32(v45 < v93)|base.B2i32(v93 <= v38) != 0 {
										v134 = l2
										m.G0 = v12 + int32(80)
										return v134
									} else {
										if v69 <= int32(0) {
											v103 = v12 - int32(-64)
										} else {
											v103 = v12 + int32(48)
										}
										if int32(0) <= v77 {
											v110 = v12 + int32(32)
										} else {
											v110 = v12 + int32(16)
										}
										v111 = int32(0)
										v113 = F_make_range(m, l0, v103, v110, v111, v111)
										mBase = m.M
										v114 = m.ExcPending
										if v114 != 0 {
											return int32(0)
										} else {
											if v38 < int32(0) {
												v119 = v113
												v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
												v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
												v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
												v131 = v129 | int32(128)
												*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
												v134 = v119
											} else {
												if int32(0) <= v45 {
													v134 = v113
												} else {
													v119 = v113
													v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
													v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
													v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
													v131 = v129 | int32(128)
													*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
													v134 = v119
												}
											}
											m.G0 = v12 + int32(80)
											return v134
										}
									}
								}
							}
						}
					}
				} else {
					if v38&int32(-127) != 0 {
						v134 = l1
						m.G0 = v12 + int32(80)
						return v134
					} else {
						v58 = l1
						v62 = F_datumCopy(m, base.I64_extend_i32_u(v58), int32(0), int32(-1))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v119 = base.I32_wrap_i64(v62)
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v119)))
							v128 = v119 + int32(base.Ui32(v123)>>(uint(int32(2))%32)) - int32(1)
							v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
							v131 = v129 | int32(128)
							*(*uint8)(unsafe.Add(mBase, uint32(v128))) = uint8(v131)
							v134 = v119
							m.G0 = v12 + int32(80)
							return v134
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	if l3&int32(16) != 0 {
		v7 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			if v7 != 0 {
				return int32(1)
			} else {
				v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				switch v11 {
				case 0:
					v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
					v13 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v12, l2)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return int32(0)
					} else {
						if v13 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
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
							return int32(1)
						}
					}
				case 1:
					if l3&int32(1) != 0 {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
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
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v20 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v19, l2)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return int32(0)
						} else {
							if v20 != 0 {
								return int32(1)
							} else {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return int32(1)
									} else {
										if l3&int32(32) != 0 {
											v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 != 0 {
													return int32(1)
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
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
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
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
						v25 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v24, l2)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							if v25 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return int32(1)
									} else {
										if l3&int32(32) != 0 {
											v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 != 0 {
													return int32(1)
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
								return int32(1)
							}
						}
					}
				case 3:
					v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
					v30 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v29, l2)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
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
							return int32(1)
						}
					}
				case 4:
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
					v35 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v34, l2)
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						if v35 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
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
							return int32(1)
						}
					}
				case 5:
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
					v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						if v40 == int32(0) {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
							v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								if v52 != 0 {
									return int32(1)
								} else {
									if l3&int32(32) != 0 {
										v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
										mBase = m.M
										v57 = m.ExcPending
										if v57 != 0 {
											return int32(0)
										} else {
											if v56 != 0 {
												return int32(1)
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
							return int32(1)
						}
					}
				default:
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
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
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
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
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
						v47 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v46, l2)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							if v47 == int32(0) {
								v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
								v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									if v52 != 0 {
										return int32(1)
									} else {
										if l3&int32(32) != 0 {
											v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
											mBase = m.M
											v57 = m.ExcPending
											if v57 != 0 {
												return int32(0)
											} else {
												if v56 != 0 {
													return int32(1)
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
								return int32(1)
							}
						}
					}
				}
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		switch v11 {
		case 0:
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v13 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v12, l2)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
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
					return int32(1)
				}
			}
		case 1:
			if l3&int32(1) != 0 {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return int32(1)
					} else {
						if l3&int32(32) != 0 {
							v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									return int32(1)
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
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				v20 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v19, l2)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					if v20 != 0 {
						return int32(1)
					} else {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
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
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return int32(1)
					} else {
						if l3&int32(32) != 0 {
							v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									return int32(1)
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
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v25 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v24, l2)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v25 == int32(0) {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
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
						return int32(1)
					}
				}
			}
		case 3:
			v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			v30 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v29, l2)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				if v30 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
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
					return int32(1)
				}
			}
		case 4:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v35 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v34, l2)
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				if v35 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
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
					return int32(1)
				}
			}
		case 5:
			v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			v40 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v39, l2)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return int32(0)
			} else {
				if v40 == int32(0) {
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
					v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						if v52 != 0 {
							return int32(1)
						} else {
							if l3&int32(32) != 0 {
								v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return int32(0)
								} else {
									if v56 != 0 {
										return int32(1)
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
					return int32(1)
				}
			}
		default:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				if v52 != 0 {
					return int32(1)
				} else {
					if l3&int32(32) != 0 {
						v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							if v56 != 0 {
								return int32(1)
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
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
				v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					if v52 != 0 {
						return int32(1)
					} else {
						if l3&int32(32) != 0 {
							v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								if v56 != 0 {
									return int32(1)
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
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
				v47 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v46, l2)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					if v47 == int32(0) {
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
						v52 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, v51, l2)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							if v52 != 0 {
								return int32(1)
							} else {
								if l3&int32(32) != 0 {
									v56 = m.T0[l1].(func(*base.Module, int32, int32) int32)(m, l0, l2)
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return int32(0)
									} else {
										if v56 != 0 {
											return int32(1)
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
						return int32(1)
					}
				}
			}
		}
	}
}
