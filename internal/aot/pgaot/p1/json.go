package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_JsonTableInitOpaque(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	v3 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+80))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v18 = F_palloc0(m, int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(418352867)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v22 == int32(0) {
		v92 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	if v101 != 0 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v15)+28))
	v28 = v3
	v29 = v3
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v28 < v39 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v45 = v41 + v28<<(uint(int32(2))%32)
	goto L9
L8:
	;
	v45 = int32(0)
	goto L9
L9:
	;
	if v25 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v92 = int32(0)
	goto L3
L11:
	;
	goto L12
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v49 <= v28 {
		v92 = v29
		goto L3
	} else {
		goto L13
	}
L13:
	;
	if v45 == int32(0) {
		v92 = v29
		goto L3
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v56 = v53 + v28<<(uint(int32(2))%32)
	if v56 == int32(0) {
		v92 = v29
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v62 = F_palloc(m, int32(24))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	v65 = F_pstrdup(m, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v65
	v68 = F_strlen(m, v65)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v62)+4)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v71 = F_exprType(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+8)) = v71
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)+24))
	v75 = F_exprTypmod(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+12)) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v59)+20))
	v82 = m.T0[v81].(func(*base.Module, int32, int32, int32) int32)(m, v59, v78, v62+int32(20))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v82
	v87 = F_lappend(m, v29, v62)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v28 = v28 + int32(1)
	v29 = v87
	goto L5
L22:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v105 = v102 << (uint(int32(2)) % 32)
	goto L24
L23:
	;
	v105 = v3
	goto L24
L24:
	;
	v106 = F_palloc(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v106
	v111 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v112 = F_JsonTableInitPlan(m, v18, v16, int32(0), v92, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v112
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v18
	return
}
func F_JsonTableInitPlan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
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
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v10 = F_palloc0(m, int32(64))
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
	*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v16 - int32(50) {
	case 0:
		goto L6
	case 1:
		goto L5
	default:
		goto L3
	}
L3:
	;
	return v10
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10+v88))) = v87
	goto L3
L5:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v74 = F_JsonTableInitPlan(m, l0, v73, l2, l3, l4)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v22
	v30 = F_AllocSetContextCreateInternal(m, l4, int32(68130), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v32 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+40)) = uint8(v32)
	v34 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v30
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v37 < v34 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v67 = int32(48)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v68 == int32(0) {
		v87 = int32(0)
		v88 = v67
		goto L4
	} else {
		goto L15
	}
L10:
	;
	v42 = v37
	goto L11
L11:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v48 < v42 {
		goto L9
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v50+v42<<(uint(int32(2))%32)))) = v10
	v56 = v42 + int32(1)
	if int32(0) <= v56 {
		v42 = v56
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v71 = F_JsonTableInitPlan(m, l0, v68, v10, l3, l4)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v87 = v71
	v88 = v67
	goto L4
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = v74
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v79 = F_JsonTableInitPlan(m, l0, v78, l2, l3, l4)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v87 = v79
	v88 = int32(56)
	goto L4
}
func F_JsonTableResetRowPattern(m *base.Module, l0 int32, l1 int32) {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	v9 = m.G0
	v11 = v9 - int32(96)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = F_pg_detoast_datum(m, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		F_MemoryContextResetOnly(m, v18)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = int32(4553888)
			v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v24
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			F_jspInit(m, v11+int32(32), v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				v34 = v14 + int32(4)
				v37 = F_JsonbExtractScalar(m, v34, v11+int32(12))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					if v37 == int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = int32(18)
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
						if v44 == int32(1) {
							v47 = int32(4)
							v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
							if v49&int32(254) == int32(2) {
								v58 = v47
							} else {
								v58 = base.B2i32(v49 == int32(18)) << (uint(v47) % 32)
							}
							if v49 == int32(1) {
								v61 = v47
							} else {
								v61 = v58
							}
							v74 = v61
						} else {
							v62 = int32(1)
							if v44&v62 != 0 {
								v74 = int32(base.Ui32(v44)>>(uint(v62)%32)) - v62
							} else {
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
								v74 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v74
					} else {
					}
					*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = int32(1409)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = v27
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
					*(*int64)(unsafe.Add(mBase, uint32(v11)+76)) = int64(0)
					v87 = int32(base.Ui32(v83) >> (uint(int32(31)) % 32))
					*(*uint8)(unsafe.Add(mBase, uint32(v11)+93)) = uint8(v87)
					*(*uint8)(unsafe.Add(mBase, uint32(v11)+92)) = uint8(v87)
					v91 = v11 + int32(12)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v91
					*(*int32)(unsafe.Add(mBase, uint32(v11)+68)) = v91
					if v27 != 0 {
						v97 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
						v100 = v97 + int32(1)
					} else {
						v100 = int32(1)
					}
					v101 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v11)+95)) = uint8(v101)
					*(*uint8)(unsafe.Add(mBase, uint32(v11)+94)) = uint8(v26)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+88)) = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+84)) = v100
					v113 = F_executeItemOptUnwrapTarget(m, v11+int32(60), v11+int32(32), v11+int32(12), l0+int32(16), v87)
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v22
						if v113 == int32(2) {
							*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = int64(0)
							v144 = l0 + int32(24)
							*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
							v150 = v144
							v152 = int32(0)
						} else {
							v124 = l0 + int32(24)
							v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							if v125 != 0 {
								v126 = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v126
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v125
								v150 = v124
								v152 = v126
							} else {
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v130 == int32(0) {
									v144 = v124
									*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
									v150 = v144
									v152 = int32(0)
								} else {
									v133 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
									v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v130
									*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v134
									v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
									if v138 < int32(2) {
										v150 = v124
										v152 = int32(0)
									} else {
										v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
										v150 = v124
										v152 = v141 + int32(4)
									}
								}
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v150)+8)) = v152
						v154 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v154
						v156 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)) = uint8(v156)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v154
						m.G0 = v11 + int32(96)
						return
					}
				}
			}
		}
	}
}
func F__equalJsonValueExpr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v3 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v6 = F_equal(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		if v6 == int32(0) {
			v22 = v3
			return v22
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v14 = F_equal(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return int32(0)
			} else {
				if v14 == int32(0) {
					v22 = v3
					return v22
				} else {
					v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v20 = F_equal(m, v18, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v22 = v20
						return v22
					}
				}
			}
		}
	}
}
func F_flattenJsonPathParseItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v307 int32
	_ = v307
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v421 int32
	_ = v421
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_check_stack_depth(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v28 = int32(*(*int8)(unsafe.Add(mBase, uint32(l3))))
	F_appendStringInfoChar(m, l0, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	goto L5
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v31 ^ int32(-1) + (v31+int32(3))&int32(-4) {
	case 0:
		goto L9
	case 1:
		goto L10
	case 2:
		goto L11
	default:
		goto L8
	}
L8:
	;
	v102 = v19 - int32(8)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	F_appendBinaryStringInfo(m, l0, v17+int32(12), int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L24
	}
L9:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v80 <= v81+int32(1) {
		goto L20
	} else {
		goto L21
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v59 <= v60+int32(1) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v39 <= v31+int32(1) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v46+v31))) = uint8(v48)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = v50 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v54+v52))) = uint8(v48)
	goto L10
L15:
	;
	goto L10
L16:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67+v60))) = uint8(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v73 = v71 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v75+v73))) = uint8(v69)
	goto L9
L19:
	;
	goto L9
L20:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v90 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v88+v81))) = uint8(v90)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v94 = v92 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v94
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*uint8)(unsafe.Add(mBase, uint32(v96+v94))) = uint8(v90)
	goto L8
L23:
	;
	goto L8
L24:
	;
	v111 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	switch v112 {
	case 0, 21, 22, 27, 31, 32, 33, 34, 35, 36, 38, 43, 44, 45, 47, 48, 49:
		goto L29
	case 1, 25, 28:
		goto L40
	case 2:
		goto L39
	case 3:
		goto L38
	case 4, 5, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 41, 46:
		goto L37
	case 6, 7, 19, 20, 30, 37, 50, 51, 52, 53:
		v222 = v111
		goto L34
	case 23:
		goto L31
	case 24:
		goto L30
	case 26:
		goto L33
	case 29:
		goto L35
	default:
		goto L28
	case 40:
		goto L32
	case 42:
		goto L36
	}
L25:
	;
	m.G0 = v17 + int32(16)
	return v421
L26:
	;
	v409 = int32(1)
	if l1 == int32(0) {
		v421 = v409
		goto L25
	} else {
		goto L110
	}
L27:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v404+v103))) = v406 - v102
	goto L26
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L107
	}
L29:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v382 == int32(0) {
		goto L26
	} else {
		goto L104
	}
L30:
	;
	F_appendBinaryStringInfo(m, l0, l3+int32(8), int32(4))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L102
	}
L31:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v283
	F_appendBinaryStringInfo(m, l0, v17+int32(12), int32(4))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L88
	}
L32:
	;
	if l5 != 0 {
		goto L29
	} else {
		goto L82
	}
L33:
	;
	if int32(0) < l4 {
		goto L29
	} else {
		goto L76
	}
L34:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	F_appendBinaryStringInfo(m, l0, v17+int32(12), int32(4))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L67
	}
L35:
	;
	v222 = int32(1)
	goto L34
L36:
	;
	F_appendBinaryStringInfo(m, l0, l3+int32(20), int32(4))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L60
	}
L37:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	F_appendBinaryStringInfo(m, l0, v17+int32(12), int32(4))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L46
	}
L38:
	;
	F_appendBinaryStringInfo(m, l0, l3+int32(8), int32(1))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L45
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	F_appendBinaryStringInfo(m, l0, v125, int32(base.Ui32(v126)>>(uint(int32(2))%32)))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L44
	}
L40:
	;
	F_appendBinaryStringInfo(m, l0, l3+int32(8), int32(4))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	F_appendBinaryStringInfo(m, l0, v118, v119)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	goto L29
L44:
	;
	goto L29
L45:
	;
	goto L29
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	F_appendBinaryStringInfo(m, l0, v17+int32(12), int32(4))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v152 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v164+v136))) = v163 - v102
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	if v168 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v102
	v163 = v102
	goto L48
L50:
	;
	goto L51
L51:
	;
	v158 = F_flattenJsonPathParseItem(m, l0, v17+int32(12), l2, v152, l4, l5)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	if v158 == int32(0) {
		v421 = v111
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v163 = v162
	goto L48
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v180+v144))) = v179 - v102
	goto L29
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v102
	v179 = v102
	goto L54
L56:
	;
	goto L57
L57:
	;
	v174 = F_flattenJsonPathParseItem(m, l0, v17+int32(12), l2, v168, l4, l5)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	if v174 == int32(0) {
		v421 = v111
		goto L25
	} else {
		goto L59
	}
L59:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v179 = v178
	goto L54
L60:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(0)
	F_appendBinaryStringInfo(m, l0, v17+int32(12), int32(4))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	F_appendBinaryStringInfo(m, l0, l3+int32(16), int32(4))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_appendBinaryStringInfo(m, l0, v202, v203)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_appendStringInfoChar(m, l0, int32(0))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v212 = F_flattenJsonPathParseItem(m, l0, v17+int32(12), l2, v211, l4, l5)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v212 == int32(0) {
		v421 = v111
		goto L25
	} else {
		goto L66
	}
L66:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v216+v189))) = v218 - v102
	goto L29
L67:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v231 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v245+v223))) = v244 - v102
	goto L29
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v102
	v244 = v102
	goto L68
L70:
	;
	goto L71
L71:
	;
	v238 = F_flattenJsonPathParseItem(m, l0, v17+int32(12), l2, v231, l4+v222, l5)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	if v238 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v421 = int32(0)
	goto L25
L74:
	;
	goto L75
L75:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v244 = v243
	goto L68
L76:
	;
	v251 = F_errsave_start(m, l2)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v251 == int32(0) {
		v421 = v111
		goto L25
	} else {
		goto L78
	}
L78:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	F_errmsg(m, int32(153438), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errsave_finish(m, l2, int32(520648), int32(389), int32(304410))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v421 = v111
	goto L25
L82:
	;
	v267 = F_errsave_start(m, l2)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	if v267 == int32(0) {
		v421 = v111
		goto L25
	} else {
		goto L84
	}
L84:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(125714), int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errsave_finish(m, l2, int32(520648), int32(395), int32(304410))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v421 = v111
	goto L25
L88:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	F_appendStringInfoSpaces(m, l0, v291<<(uint(int32(3))%32))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v296 <= int32(0) {
		goto L29
	} else {
		goto L90
	}
L90:
	;
	v307 = v111
	goto L91
L91:
	;
	v316 = v307 << (uint(int32(3)) % 32)
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v316+v317)))
	v321 = F_flattenJsonPathParseItem(m, l0, v17+int32(4), l2, v319, l4, int32(1))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	v421 = int32(0)
	goto L25
L93:
	;
	goto L92
L94:
	;
	if v321 == int32(0) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v326 = v325 - v102
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v326
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v328+v316)+4))
	if v330 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v334 = F_flattenJsonPathParseItem(m, l0, v17+int32(8), l2, v330, l4, int32(1))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	v342 = v326
	v343 = int32(0)
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v343
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v347 = v345 + v290 + v316
	*(*int32)(unsafe.Add(mBase, uint32(v347))) = v342
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v347)+4)) = v349
	v352 = v307 + int32(1)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if v352 < v353 {
		v307 = v352
		goto L91
	} else {
		goto L101
	}
L99:
	;
	if v334 == int32(0) {
		goto L93
	} else {
		goto L100
	}
L100:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v342 = v338
	v343 = v339 - v102
	goto L98
L101:
	;
	goto L29
L102:
	;
	F_appendBinaryStringInfo(m, l0, l3+int32(12), int32(4))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	goto L29
L104:
	;
	v387 = F_flattenJsonPathParseItem(m, l0, v17+int32(12), l2, v382, l4, l5)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v387 != 0 {
		goto L27
	} else {
		goto L106
	}
L106:
	;
	v421 = int32(0)
	goto L25
L107:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v394
	F_errmsg_internal(m, int32(506242), v17)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(520648), int32(462), int32(304410))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v102
	v421 = v409
	goto L25
}
func F_json_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v8 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v17 = v15 + int32(5)
			v18 = F_palloc(m, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v17 << (uint(int32(2)) % 32)
				v26 = v18 + int32(4)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v15 != 0 {
					v28 = F__emscripten_memcpy_bulkmem(m, v26, v27, v15)
					mBase = m.M
					v29 = v28
				} else {
					v29 = v26
				}
				v31 = int32(93)
				*(*uint8)(unsafe.Add(mBase, uint32(v29+v15))) = uint8(v31)
				return v18
			}
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		return int32(0)
	}
}
func F_json_array_element_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v13
		v18 = int32(1)
		v20 = F_get_worker(m, v9, int32(0), v6+int32(12), v18, v18)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			if v20 == int32(0) {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v24)
				v27 = int32(0)
			} else {
				v27 = v20
			}
			m.G0 = v6 + int32(16)
			return v27
		}
	}
}
func F_json_array_length(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v17 = F_palloc0(m, int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v21 = F_pg_detoast_datum_packed(m, v12)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = int32(1)
				v24 = v21 + v23
				v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
				v29 = v27 & v23
				if v29 != 0 {
					v30 = v24
				} else {
					v30 = v21 + int32(4)
				}
				if v27 == int32(1) {
					v33 = int32(4)
					v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
					if v35&int32(254) == int32(2) {
						v44 = v33
					} else {
						v44 = base.B2i32(v35 == int32(18)) << (uint(v33) % 32)
					}
					if v35 == int32(1) {
						v47 = v33
					} else {
						v47 = v44
					}
					v58 = v47
				} else {
					v48 = int32(1)
					if v29 != 0 {
						v58 = int32(base.Ui32(v27)>>(uint(v48)%32)) - v48
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
						v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				v60 = *(*int32)(unsafe.Add(mBase, _consts[356]))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
				v63 = F_makeJsonLexContextCstringLen(m, v9+int32(12), v30, v58, v61, int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v63
					v67 = F_palloc0(m, int32(40))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v67)+36)) = int32(1360)
						*(*int32)(unsafe.Add(mBase, uint32(v67)+4)) = int32(1361)
						*(*int32)(unsafe.Add(mBase, uint32(v67))) = v17
						*(*int32)(unsafe.Add(mBase, uint32(v67)+28)) = int32(1362)
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
						v77 = F_pg_parse_json(m, v76, v67)
						mBase = m.M
						v78 = m.ExcPending
						if v78 != 0 {
							return int32(0)
						} else {
							if v77 != 0 {
								F_json_errsave_error(m, v77, v76, int32(0))
								mBase = m.M
								v81 = m.ExcPending
								if v81 != 0 {
									return int32(0)
								} else {
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									m.G0 = v9 + int32(80)
									return v82
								}
							} else {
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
								m.G0 = v9 + int32(80)
								return v82
							}
						}
					}
				}
			}
		}
	}
}
func F_json_build_object(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v15 = F_extract_variadic_args(m, l0, v7+int32(12), v7+int32(4), v7+int32(8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v21 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
			v30 = int32(0)
			m.G0 = v7 + int32(16)
			return v30
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			v26 = int32(0)
			v28 = F_json_build_object_worker(m, v15, v23, v24, v25, v26, v26)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				v30 = v28
				m.G0 = v7 + int32(16)
				return v30
			}
		}
	}
}
func F_json_object_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v5 == int32(0) {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v8 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
			v17 = v15 + int32(6)
			v18 = F_palloc(m, v17)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v18))) = v17 << (uint(int32(2)) % 32)
				v26 = v18 + int32(4)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
				if v15 != 0 {
					v28 = F__emscripten_memcpy_bulkmem(m, v26, v27, v15)
					mBase = m.M
					v29 = v28
				} else {
					v29 = v26
				}
				v31 = int32(32032)
				*(*uint16)(unsafe.Add(mBase, uint32(v29+v15))) = uint16(v31)
				return v18
			}
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
			return int32(0)
		}
	} else {
		v10 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v10)
		return int32(0)
	}
}
func F_json_object_agg_strict_transfn(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = F_json_object_agg_transfn_worker(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_object_agg_transfn(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = F_json_object_agg_transfn_worker(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_json_object_field(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = F_text_to_cstring(m, v14)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v16
				v21 = int32(0)
				v24 = F_get_worker(m, v9, v6+int32(12), v21, int32(1), v21)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					if v24 == int32(0) {
						v28 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
						v31 = int32(0)
					} else {
						v31 = v24
					}
					m.G0 = v6 + int32(16)
					return v31
				}
			}
		}
	}
}
func F_json_object_keys(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int64
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	v11 = m.G0
	v13 = v11 - int32(80)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+16))
	if v16 == int32(0) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v20 = F_pg_detoast_datum_packed(m, v19)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			v24 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				v26 = int32(4553888)
				v27 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+24))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v29
				v32 = F_palloc(m, int32(20))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v35 = F_palloc0(m, int32(40))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						v39 = F_pg_detoast_datum_packed(m, v20)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v41 = int32(1)
							v42 = v39 + v41
							v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
							v47 = v45 & v41
							if v47 != 0 {
								v48 = v42
							} else {
								v48 = v39 + int32(4)
							}
							if v45 == int32(1) {
								v51 = int32(4)
								v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
								if v53&int32(254) == int32(2) {
									v62 = v51
								} else {
									v62 = base.B2i32(v53 == int32(18)) << (uint(v51) % 32)
								}
								if v53 == int32(1) {
									v65 = v51
								} else {
									v65 = v62
								}
								v76 = v65
							} else {
								v66 = int32(1)
								if v47 != 0 {
									v76 = int32(base.Ui32(v45)>>(uint(v66)%32)) - v66
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
									v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
								}
							}
							v78 = *(*int32)(unsafe.Add(mBase, _consts[356]))
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
							v81 = F_makeJsonLexContextCstringLen(m, v13+int32(12), v48, v76, v79, int32(1))
							mBase = m.M
							v82 = m.ExcPending
							if v82 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v32)+16)) = int32(0)
								*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = int64(256)
								*(*int32)(unsafe.Add(mBase, uint32(v32))) = v81
								v89 = F_palloc(m, int32(1024))
								mBase = m.M
								v90 = m.ExcPending
								if v90 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v89
									*(*int32)(unsafe.Add(mBase, uint32(v35)+36)) = int32(1348)
									*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(1349)
									*(*int32)(unsafe.Add(mBase, uint32(v35))) = v32
									*(*int32)(unsafe.Add(mBase, uint32(v35)+20)) = int32(1350)
									v101 = F_pg_parse_json(m, v13+int32(12), v35)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										if v101 != 0 {
											F_json_errsave_error(m, v101, v13+int32(12), int32(0))
											mBase = m.M
											v107 = m.ExcPending
											if v107 != 0 {
												return int32(0)
											} else {
												F_freeJsonLexContext(m, v13+int32(12))
												mBase = m.M
												v111 = m.ExcPending
												if v111 != 0 {
													return int32(0)
												} else {
													F_pfree(m, v35)
													mBase = m.M
													v113 = m.ExcPending
													if v113 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, _consts[3])) = v27
														*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
														v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
														v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
														v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
														v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
														if v128 < v129 {
															v131 = int32(1)
															*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v128 + v131
															v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
															v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v128<<(uint(int32(2))%32))))
															v139 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
															*(*int64)(unsafe.Add(mBase, uint32(v126))) = v139 + int64(1)
															v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v131
															v146 = F_cstring_to_text(m, v138)
															mBase = m.M
															v147 = m.ExcPending
															if v147 != 0 {
																return int32(0)
															} else {
																v157 = v146
																m.G0 = v13 + int32(80)
																return v157
															}
														} else {
															F_end_MultiFuncCall(m, l0)
															mBase = m.M
															v149 = m.ExcPending
															if v149 != 0 {
																return int32(0)
															} else {
																v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
																*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = int32(2)
																v153 = int32(1)
																*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v153)
																v157 = int32(0)
																m.G0 = v13 + int32(80)
																return v157
															}
														}
													}
												}
											}
										} else {
											F_freeJsonLexContext(m, v13+int32(12))
											mBase = m.M
											v111 = m.ExcPending
											if v111 != 0 {
												return int32(0)
											} else {
												F_pfree(m, v35)
												mBase = m.M
												v113 = m.ExcPending
												if v113 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, _consts[3])) = v27
													*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v32
													v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
													v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
													v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
													v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
													if v128 < v129 {
														v131 = int32(1)
														*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v128 + v131
														v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v128<<(uint(int32(2))%32))))
														v139 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
														*(*int64)(unsafe.Add(mBase, uint32(v126))) = v139 + int64(1)
														v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v131
														v146 = F_cstring_to_text(m, v138)
														mBase = m.M
														v147 = m.ExcPending
														if v147 != 0 {
															return int32(0)
														} else {
															v157 = v146
															m.G0 = v13 + int32(80)
															return v157
														}
													} else {
														F_end_MultiFuncCall(m, l0)
														mBase = m.M
														v149 = m.ExcPending
														if v149 != 0 {
															return int32(0)
														} else {
															v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = int32(2)
															v153 = int32(1)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v153)
															v157 = int32(0)
															m.G0 = v13 + int32(80)
															return v157
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
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)+16))
		v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
		v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+16))
		v129 = *(*int32)(unsafe.Add(mBase, uint32(v127)+12))
		if v128 < v129 {
			v131 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v127)+16)) = v128 + v131
			v134 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
			v138 = *(*int32)(unsafe.Add(mBase, uint32(v134+v128<<(uint(int32(2))%32))))
			v139 = *(*int64)(unsafe.Add(mBase, uint32(v126)))
			*(*int64)(unsafe.Add(mBase, uint32(v126))) = v139 + int64(1)
			v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v143)+20)) = v131
			v146 = F_cstring_to_text(m, v138)
			mBase = m.M
			v147 = m.ExcPending
			if v147 != 0 {
				return int32(0)
			} else {
				v157 = v146
				m.G0 = v13 + int32(80)
				return v157
			}
		} else {
			F_end_MultiFuncCall(m, l0)
			mBase = m.M
			v149 = m.ExcPending
			if v149 != 0 {
				return int32(0)
			} else {
				v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v150)+20)) = int32(2)
				v153 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v153)
				v157 = int32(0)
				m.G0 = v13 + int32(80)
				return v157
			}
		}
	}
}
func F_json_populate_recordset(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	v3 = int32(1)
	F_populate_recordset_worker(m, l0, int32(112733), v3, v3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return int32(0)
	}
}
func F_json_send(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v16 = v11 + int32(1)
		F_pq_begintypsend(m, v8)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			v23 = v21 & int32(1)
			if v23 != 0 {
				v24 = v16
			} else {
				v24 = v11 + int32(4)
			}
			if v21 == int32(1) {
				v27 = int32(4)
				v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
				if v29&int32(254) == int32(2) {
					v38 = v27
				} else {
					v38 = base.B2i32(v29 == int32(18)) << (uint(v27) % 32)
				}
				if v29 == int32(1) {
					v41 = v27
				} else {
					v41 = v38
				}
				v52 = v41
			} else {
				v42 = int32(1)
				if v23 != 0 {
					v52 = int32(base.Ui32(v21)>>(uint(v42)%32)) - v42
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			F_pq_sendtext(m, v8, v24, v52)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 << (uint(int32(2)) % 32)
				m.G0 = v8 + int32(16)
				return v56
			}
		}
	}
}
func F_json_string_to_tsvector(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = F_getTSCurrentConfig(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v14
			v17 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v7 + int32(8)
			F_iterate_json_values(m, v10, int32(2), v7+int32(24))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v31 = F_make_tsvector(m, v7+int32(8))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v33 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(32)
							return v31
						}
					} else {
						m.G0 = v7 + int32(32)
						return v31
					}
				}
			}
		}
	}
}
func F_json_unique_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
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
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
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
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
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
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
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
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v8 = int32(711645284)
	v11 = v3 - int32(1636608428) ^ v8 - int32(1455628627)
	v16 = v11 ^ int32(-1636608428) - base.I32_rotl(v11, int32(25))
	v21 = v16 ^ v8 - base.I32_rotl(v16, int32(16))
	v25 = v21 ^ v11 - base.I32_rotl(v21, int32(4))
	v29 = v25 ^ v16 - base.I32_rotl(v25, int32(14))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = v35 - int32(1636608432)
	if v34&int32(3) != 0 {
		if base.Ui32(int32(11)) < base.Ui32(v35) {
			v150 = v34
			v151 = v35
			v152 = v41
			v153 = v41
			v154 = v41
			for {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
				v157 = v156 + v153
				v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
				v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
				v161 = v160 + v154
				v163 = int32(4)
				v165 = v158 + v152 - v161 ^ base.I32_rotl(v161, v163)
				v169 = v157 - v165 ^ base.I32_rotl(v165, int32(6))
				v170 = v161 + v157
				v171 = v165 + v170
				v172 = v169 + v171
				v176 = v170 - v169 ^ base.I32_rotl(v169, int32(8))
				v180 = v171 - v176 ^ base.I32_rotl(v176, int32(16))
				v184 = v172 - v180 ^ base.I32_rotl(v180, int32(19))
				v185 = v176 + v172
				v186 = v180 + v185
				v187 = v184 + v186
				v191 = v185 - v184 ^ base.I32_rotl(v184, v163)
				v192 = int32(12)
				v193 = v150 + v192
				v195 = v151 - v192
				if base.Ui32(int32(11)) < base.Ui32(v195) {
					v150 = v193
					v151 = v195
					v152 = v186
					v153 = v187
					v154 = v191
					continue
				} else {
					break
				}
				break
			}
			v198 = v193
			v199 = v195
			v200 = v186
			v201 = v187
			v202 = v191
		} else {
			v198 = v34
			v199 = v35
			v200 = v41
			v201 = v41
			v202 = v41
		}
		switch v199 - int32(1) {
		case 0:
			v261 = v200
			v262 = v201
			v263 = v202
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 1:
			v254 = v200
			v255 = v201
			v256 = v202
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 2:
			v247 = v200
			v248 = v201
			v249 = v202
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 3:
			v241 = v201
			v242 = v202
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 4:
			v237 = v201
			v238 = v202
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 5:
			v231 = v201
			v232 = v202
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
			v237 = v233<<(uint(int32(8))%32) + v231
			v238 = v232
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 6:
			v225 = v201
			v226 = v202
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+6)))
			v231 = v227<<(uint(int32(16))%32) + v225
			v232 = v226
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
			v237 = v233<<(uint(int32(8))%32) + v231
			v238 = v232
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 7:
			v220 = v202
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+7)))
			v225 = v221<<(uint(int32(24))%32) + v201
			v226 = v220
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+6)))
			v231 = v227<<(uint(int32(16))%32) + v225
			v232 = v226
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
			v237 = v233<<(uint(int32(8))%32) + v231
			v238 = v232
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 8:
			v215 = v202
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+8)))
			v220 = v216<<(uint(int32(8))%32) + v215
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+7)))
			v225 = v221<<(uint(int32(24))%32) + v201
			v226 = v220
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+6)))
			v231 = v227<<(uint(int32(16))%32) + v225
			v232 = v226
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
			v237 = v233<<(uint(int32(8))%32) + v231
			v238 = v232
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 9:
			v210 = v202
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+9)))
			v215 = v211<<(uint(int32(16))%32) + v210
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+8)))
			v220 = v216<<(uint(int32(8))%32) + v215
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+7)))
			v225 = v221<<(uint(int32(24))%32) + v201
			v226 = v220
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+6)))
			v231 = v227<<(uint(int32(16))%32) + v225
			v232 = v226
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
			v237 = v233<<(uint(int32(8))%32) + v231
			v238 = v232
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		case 10:
			v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+10)))
			v210 = v206<<(uint(int32(24))%32) + v202
			v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+9)))
			v215 = v211<<(uint(int32(16))%32) + v210
			v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+8)))
			v220 = v216<<(uint(int32(8))%32) + v215
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+7)))
			v225 = v221<<(uint(int32(24))%32) + v201
			v226 = v220
			v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+6)))
			v231 = v227<<(uint(int32(16))%32) + v225
			v232 = v226
			v233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+5)))
			v237 = v233<<(uint(int32(8))%32) + v231
			v238 = v232
			v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+4)))
			v241 = v237 + v239
			v242 = v238
			v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+3)))
			v247 = v243<<(uint(int32(24))%32) + v200
			v248 = v241
			v249 = v242
			v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+2)))
			v254 = v250<<(uint(int32(16))%32) + v247
			v255 = v248
			v256 = v249
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198)+1)))
			v261 = v257<<(uint(int32(8))%32) + v254
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198))))
			v268 = v261 + v264
			v269 = v262
			v270 = v263
		default:
			v268 = v200
			v269 = v201
			v270 = v202
		}
	} else {
		if base.Ui32(v35) < base.Ui32(int32(12)) {
			v96 = v34
			v97 = v35
			v98 = v41
			v99 = v41
			v100 = v41
		} else {
			v48 = v34
			v49 = v35
			v50 = v41
			v51 = v41
			v52 = v41
			for {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
				v55 = v54 + v51
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
				v59 = v58 + v52
				v61 = int32(4)
				v63 = v56 + v50 - v59 ^ base.I32_rotl(v59, v61)
				v67 = v55 - v63 ^ base.I32_rotl(v63, int32(6))
				v68 = v59 + v55
				v69 = v63 + v68
				v70 = v67 + v69
				v74 = v68 - v67 ^ base.I32_rotl(v67, int32(8))
				v78 = v69 - v74 ^ base.I32_rotl(v74, int32(16))
				v82 = v70 - v78 ^ base.I32_rotl(v78, int32(19))
				v83 = v74 + v70
				v84 = v78 + v83
				v85 = v82 + v84
				v89 = v83 - v82 ^ base.I32_rotl(v82, v61)
				v90 = int32(12)
				v91 = v48 + v90
				v93 = v49 - v90
				if base.Ui32(int32(11)) < base.Ui32(v93) {
					v48 = v91
					v49 = v93
					v50 = v84
					v51 = v85
					v52 = v89
					continue
				} else {
					break
				}
				break
			}
			v96 = v91
			v97 = v93
			v98 = v84
			v99 = v85
			v100 = v89
		}
		switch v97 - int32(1) {
		case 0:
			v147 = v98
			v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
			v268 = v147 + v148
			v269 = v99
			v270 = v100
		case 1:
			v142 = v98
			v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
			v147 = v143<<(uint(int32(8))%32) + v142
			v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
			v268 = v147 + v148
			v269 = v99
			v270 = v100
		case 2:
			v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+2)))
			v142 = v138<<(uint(int32(16))%32) + v98
			v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+1)))
			v147 = v143<<(uint(int32(8))%32) + v142
			v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
			v268 = v147 + v148
			v269 = v99
			v270 = v100
		case 3:
			v135 = v99
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v268 = v136 + v98
			v269 = v135
			v270 = v100
		case 4:
			v132 = v99
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)))
			v135 = v132 + v133
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v268 = v136 + v98
			v269 = v135
			v270 = v100
		case 5:
			v127 = v99
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
			v132 = v128<<(uint(int32(8))%32) + v127
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)))
			v135 = v132 + v133
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v268 = v136 + v98
			v269 = v135
			v270 = v100
		case 6:
			v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+6)))
			v127 = v123<<(uint(int32(16))%32) + v99
			v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+5)))
			v132 = v128<<(uint(int32(8))%32) + v127
			v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+4)))
			v135 = v132 + v133
			v136 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v268 = v136 + v98
			v269 = v135
			v270 = v100
		case 7:
			v118 = v100
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
			v268 = v119 + v98
			v269 = v121 + v99
			v270 = v118
		case 8:
			v113 = v100
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+8)))
			v118 = v114<<(uint(int32(8))%32) + v113
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
			v268 = v119 + v98
			v269 = v121 + v99
			v270 = v118
		case 9:
			v108 = v100
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+9)))
			v113 = v109<<(uint(int32(16))%32) + v108
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+8)))
			v118 = v114<<(uint(int32(8))%32) + v113
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
			v268 = v119 + v98
			v269 = v121 + v99
			v270 = v118
		case 10:
			v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+10)))
			v108 = v104<<(uint(int32(24))%32) + v100
			v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+9)))
			v113 = v109<<(uint(int32(16))%32) + v108
			v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+8)))
			v118 = v114<<(uint(int32(8))%32) + v113
			v119 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
			v121 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
			v268 = v119 + v98
			v269 = v121 + v99
			v270 = v118
		default:
			v268 = v98
			v269 = v99
			v270 = v100
		}
	}
	v273 = int32(14)
	v275 = v269 ^ v270 - base.I32_rotl(v269, v273)
	v279 = v275 ^ v268 - base.I32_rotl(v275, int32(11))
	v283 = v279 ^ v269 - base.I32_rotl(v279, int32(25))
	v287 = v283 ^ v275 - base.I32_rotl(v283, int32(16))
	v291 = v287 ^ v279 - base.I32_rotl(v287, int32(4))
	v295 = v291 ^ v283 - base.I32_rotl(v291, v273)
	return v29 ^ v21 - base.I32_rotl(v29, int32(24)) ^ (v295 ^ v287 - base.I32_rotl(v295, int32(24)))
}
func F_json_unique_hash_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 != v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v6 < v5 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 != v14 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v11 = int32(1)
	goto L6
L5:
	;
	v11 = int32(-1)
	goto L6
L6:
	;
	return v11
L7:
	;
	if v14 < v13 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v13 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v19 = int32(1)
	goto L12
L11:
	;
	v19 = int32(-1)
	goto L12
L12:
	;
	return v19
L13:
	;
	return v66
L14:
	;
	v66 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v28 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v29 = v21
	v30 = v22
	v31 = v13
	v32 = v28
	goto L21
L18:
	;
	v54 = v22
	v58 = int32(0)
	goto L19
L19:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	v66 = v58 - v59
	goto L13
L20:
	;
	v54 = v49
	v58 = v51
	goto L19
L21:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v32 != v34 {
		v49 = v30
		v51 = v32
		goto L20
	} else {
		goto L23
	}
L22:
	;
	v49 = v43
	v51 = int32(0)
	goto L20
L23:
	;
	if v34 == int32(0) {
		v49 = v30
		v51 = v32
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v39 = v31 - int32(1)
	if v39 == int32(0) {
		v49 = v30
		v51 = v32
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v42 = int32(1)
	v43 = v30 + v42
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+1)))
	if v44 != 0 {
		v29 = v29 + v42
		v30 = v43
		v31 = v39
		v32 = v44
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
}
func F_makeJsonConstructorExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	v6 = l5
	v7 = l6
	v11 = m.G0
	v12 = int32(32)
	v13 = v11 - v12
	m.G0 = v13
	v16 = F_palloc0(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(45)
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+25)) = uint8(v6)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = l4
		*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = l7
		*(*uint8)(unsafe.Add(mBase, uint32(v16)+24)) = uint8(v7)
		v30 = F_palloc0(m, int32(16))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(34)
			if l3 != 0 {
				v34 = F_exprType(m, l3)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v34
					v37 = F_exprTypmod(m, l3)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v37
						v40 = F_exprCollation(m, l3)
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							v54 = v40
							*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v54
							v56 = F_exprType(m, v30)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								v58 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
								if v58 == int32(0) {
									m.G0 = v13 + int32(32)
									return v16
								} else {
									if v56 == v58 {
										m.G0 = v13 + int32(32)
										return v16
									} else {
										v62 = F_exprLocation(m, v30)
										mBase = m.M
										v63 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
										if v62 < int32(0) {
											v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
											v67 = v66
										} else {
											v67 = v62
										}
										v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
										if v69 != int32(1) {
											v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
											v101 = F_coerce_to_target_type(m, l0, v30, v56, v68, v98, int32(1), int32(2), v67)
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return int32(0)
											} else {
												if v101 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v122 = m.ExcPending
													if v122 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(101744772))
														mBase = m.M
														v125 = m.ExcPending
														if v125 != 0 {
															return int32(0)
														} else {
															v126 = F_format_type_be(m, v56)
															mBase = m.M
															v127 = m.ExcPending
															if v127 != 0 {
																return int32(0)
															} else {
																v128 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
																v129 = F_format_type_be(m, v128)
																mBase = m.M
																v130 = m.ExcPending
																if v130 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v129
																	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v126
																	F_errmsg(m, int32(192857), v13+int32(16))
																	mBase = m.M
																	v137 = m.ExcPending
																	if v137 != 0 {
																		return int32(0)
																	} else {
																		F_parser_coercion_errposition(m, l0, v67, v30)
																		mBase = m.M
																		v139 = m.ExcPending
																		if v139 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(517271), int32(3667), int32(217502))
																			mBase = m.M
																			v144 = m.ExcPending
																			if v144 != 0 {
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
												} else {
													v105 = v101
													if v105 == v30 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
													}
													m.G0 = v13 + int32(32)
													return v16
												}
											}
										} else {
											if v68 != int32(17) {
												v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
												v101 = F_coerce_to_target_type(m, l0, v30, v56, v68, v98, int32(1), int32(2), v67)
												mBase = m.M
												v102 = m.ExcPending
												if v102 != 0 {
													return int32(0)
												} else {
													if v101 == int32(0) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v122 = m.ExcPending
														if v122 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(101744772))
															mBase = m.M
															v125 = m.ExcPending
															if v125 != 0 {
																return int32(0)
															} else {
																v126 = F_format_type_be(m, v56)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	v128 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
																	v129 = F_format_type_be(m, v128)
																	mBase = m.M
																	v130 = m.ExcPending
																	if v130 != 0 {
																		return int32(0)
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v129
																		*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v126
																		F_errmsg(m, int32(192857), v13+int32(16))
																		mBase = m.M
																		v137 = m.ExcPending
																		if v137 != 0 {
																			return int32(0)
																		} else {
																			F_parser_coercion_errposition(m, l0, v67, v30)
																			mBase = m.M
																			v139 = m.ExcPending
																			if v139 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(517271), int32(3667), int32(217502))
																				mBase = m.M
																				v144 = m.ExcPending
																				if v144 != 0 {
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
													} else {
														v105 = v101
														if v105 == v30 {
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
														}
														m.G0 = v13 + int32(32)
														return v16
													}
												}
											} else {
												v76 = F_coerce_to_specific_type(m, l0, v30, int32(25), int32(553074))
												mBase = m.M
												v77 = m.ExcPending
												if v77 != 0 {
													return int32(0)
												} else {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
													v79 = F_getJsonEncodingConst(m, v78)
													mBase = m.M
													v80 = m.ExcPending
													if v80 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v79
														*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v76
														*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
														*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v79
														v91 = F_list_make2_impl(m, v13+int32(12), v13+int32(8))
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															v93 = int32(0)
															v95 = F_makeFuncExpr(m, int32(1717), int32(17), v91, v93, v93)
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v95)+32)) = v67
																v105 = v95
																if v105 == v30 {
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
																}
																m.G0 = v13 + int32(32)
																return v16
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
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = int32(-1)
				if v43 == int32(2) {
					v50 = int32(3802)
				} else {
					v50 = int32(114)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v50
				v54 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+12)) = v54
				v56 = F_exprType(m, v30)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
					if v58 == int32(0) {
						m.G0 = v13 + int32(32)
						return v16
					} else {
						if v56 == v58 {
							m.G0 = v13 + int32(32)
							return v16
						} else {
							v62 = F_exprLocation(m, v30)
							mBase = m.M
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
							if v62 < int32(0) {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)+12))
								v67 = v66
							} else {
								v67 = v62
							}
							v68 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
							if v69 != int32(1) {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
								v101 = F_coerce_to_target_type(m, l0, v30, v56, v68, v98, int32(1), int32(2), v67)
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return int32(0)
								} else {
									if v101 == int32(0) {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(101744772))
											mBase = m.M
											v125 = m.ExcPending
											if v125 != 0 {
												return int32(0)
											} else {
												v126 = F_format_type_be(m, v56)
												mBase = m.M
												v127 = m.ExcPending
												if v127 != 0 {
													return int32(0)
												} else {
													v128 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
													v129 = F_format_type_be(m, v128)
													mBase = m.M
													v130 = m.ExcPending
													if v130 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v129
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v126
														F_errmsg(m, int32(192857), v13+int32(16))
														mBase = m.M
														v137 = m.ExcPending
														if v137 != 0 {
															return int32(0)
														} else {
															F_parser_coercion_errposition(m, l0, v67, v30)
															mBase = m.M
															v139 = m.ExcPending
															if v139 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(517271), int32(3667), int32(217502))
																mBase = m.M
																v144 = m.ExcPending
																if v144 != 0 {
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
									} else {
										v105 = v101
										if v105 == v30 {
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
										}
										m.G0 = v13 + int32(32)
										return v16
									}
								}
							} else {
								if v68 != int32(17) {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
									v101 = F_coerce_to_target_type(m, l0, v30, v56, v68, v98, int32(1), int32(2), v67)
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return int32(0)
									} else {
										if v101 == int32(0) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(101744772))
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													v126 = F_format_type_be(m, v56)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														v128 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
														v129 = F_format_type_be(m, v128)
														mBase = m.M
														v130 = m.ExcPending
														if v130 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v129
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v126
															F_errmsg(m, int32(192857), v13+int32(16))
															mBase = m.M
															v137 = m.ExcPending
															if v137 != 0 {
																return int32(0)
															} else {
																F_parser_coercion_errposition(m, l0, v67, v30)
																mBase = m.M
																v139 = m.ExcPending
																if v139 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(517271), int32(3667), int32(217502))
																	mBase = m.M
																	v144 = m.ExcPending
																	if v144 != 0 {
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
										} else {
											v105 = v101
											if v105 == v30 {
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
											}
											m.G0 = v13 + int32(32)
											return v16
										}
									}
								} else {
									v76 = F_coerce_to_specific_type(m, l0, v30, int32(25), int32(553074))
									mBase = m.M
									v77 = m.ExcPending
									if v77 != 0 {
										return int32(0)
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
										v79 = F_getJsonEncodingConst(m, v78)
										mBase = m.M
										v80 = m.ExcPending
										if v80 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v79
											*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = v76
											*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v76
											*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v79
											v91 = F_list_make2_impl(m, v13+int32(12), v13+int32(8))
											mBase = m.M
											v92 = m.ExcPending
											if v92 != 0 {
												return int32(0)
											} else {
												v93 = int32(0)
												v95 = F_makeFuncExpr(m, int32(1717), int32(17), v91, v93, v93)
												mBase = m.M
												v96 = m.ExcPending
												if v96 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v95)+32)) = v67
													v105 = v95
													if v105 == v30 {
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v105
													}
													m.G0 = v13 + int32(32)
													return v16
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
func F_makeJsonFormat(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(42)
		return v6
	}
}
func F_makeJsonIsPredicate(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = l3
	v8 = F_palloc0(m, int32(24))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = l4
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+16)) = uint8(v4)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(46)
		return v8
	}
}
func F_transformJsonAggConstructor(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v12 != 0 {
		v15 = F_transformWhereClause(m, l0, v12, int32(8), int32(549270))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v20 = v15
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			if v21 != 0 {
				v23 = F_palloc0(m, int32(44))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v25 = int32(256)
					*(*uint16)(unsafe.Add(mBase, uint32(v23)+36)) = uint16(v25)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(11)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if v37 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(148698), int32(0))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
									return int32(0)
								} else {
									v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
									F_parser_errposition(m, l0, v78)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(517271), int32(3880), int32(218080))
										mBase = m.M
										v85 = m.ExcPending
										if v85 != 0 {
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
						v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
						F_transformWindowFuncCall(m, l0, v23, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							v61 = v23
							v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
							v64 = F_makeJsonConstructorExpr(m, l0, l6, int32(0), v61, l2, l7, l8, v63)
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return int32(0)
							} else {
								return v64
							}
						}
					}
				}
			} else {
				v42 = F_palloc0(m, int32(72))
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v42)+64)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v42)+56)) = int64(-4294967296)
					*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = int32(7208960)
					*(*int32)(unsafe.Add(mBase, uint32(v42)+44)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l5
					*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = l4
					*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(9)
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v42)+68)) = v55
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					F_transformAggregateCall(m, l0, v42, l3, v57, int32(0))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return int32(0)
					} else {
						v61 = v42
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v64 = F_makeJsonConstructorExpr(m, l0, l6, int32(0), v61, l2, l7, l8, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return v64
						}
					}
				}
			}
		}
	} else {
		v20 = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		if v21 != 0 {
			v23 = F_palloc0(m, int32(44))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = int32(256)
				*(*uint16)(unsafe.Add(mBase, uint32(v23)+36)) = uint16(v25)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = l3
				*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v23))) = int32(11)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v23)+40)) = v35
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v37 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(148698), int32(0))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return int32(0)
							} else {
								v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
								F_parser_errposition(m, l0, v78)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(517271), int32(3880), int32(218080))
									mBase = m.M
									v85 = m.ExcPending
									if v85 != 0 {
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
					v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
					F_transformWindowFuncCall(m, l0, v23, v38)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						v61 = v23
						v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
						v64 = F_makeJsonConstructorExpr(m, l0, l6, int32(0), v61, l2, l7, l8, v63)
						mBase = m.M
						v65 = m.ExcPending
						if v65 != 0 {
							return int32(0)
						} else {
							return v64
						}
					}
				}
			}
		} else {
			v42 = F_palloc0(m, int32(72))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v42)+64)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v42)+56)) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = int32(7208960)
				*(*int32)(unsafe.Add(mBase, uint32(v42)+44)) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v42)+8)) = l5
				*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = l4
				*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(9)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v42)+68)) = v55
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				F_transformAggregateCall(m, l0, v42, l3, v57, int32(0))
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return int32(0)
				} else {
					v61 = v42
					v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v64 = F_makeJsonConstructorExpr(m, l0, l6, int32(0), v61, l2, l7, l8, v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						return v64
					}
				}
			}
		}
	}
}
func F_transformJsonOutput(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	if l1 == int32(0) {
		v15 = F_palloc0(m, int32(16))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = int32(43)
			v21 = int32(0)
			v24 = F_makeJsonFormat(m, v21, v21, int32(-1))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v24
				v106 = v15
				m.G0 = v10 + int32(16)
				return v106
			}
		}
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v30 = F_copyObjectImpl(m, v29)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v34 = v30 + int32(8)
			F_typenameTypeIdAndMod(m, l0, v32, v34, v30+int32(12))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int32(0)
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+12)))
				if v40 == int32(1) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(149342), int32(0))
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(517271), int32(3546), int32(70198))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
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
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					v44 = F_get_typtype(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						if v44 == int32(112) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(149280), int32(0))
									mBase = m.M
									v139 = m.ExcPending
									if v139 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(517271), int32(3551), int32(70198))
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
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
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
							v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
							v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
							if v50 == int32(0) {
								if v48 == int32(3802) {
									v57 = int32(2)
								} else {
									v57 = int32(1)
								}
								*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v57
								v106 = v30
								m.G0 = v10 + int32(16)
								return v106
							} else {
								if l2 != 0 {
									v75 = v50
									if v75 != int32(1) {
										v106 = v30
										m.G0 = v10 + int32(16)
										return v106
									} else {
										v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
										if v48 != int32(17) {
											if v78 == int32(0) {
												v106 = v30
												m.G0 = v10 + int32(16)
												return v106
											} else {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v86 = m.ExcPending
												if v86 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v89 = m.ExcPending
													if v89 != 0 {
														return int32(0)
													} else {
														v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
														F_parser_errposition(m, l0, v90)
														mBase = m.M
														v92 = m.ExcPending
														if v92 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(171293), int32(0))
															mBase = m.M
															v96 = m.ExcPending
															if v96 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(517271), int32(3503), int32(119464))
																mBase = m.M
																v101 = m.ExcPending
																if v101 != 0 {
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
											if base.Ui32(int32(2)) <= base.Ui32(v78) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(1088))
													mBase = m.M
													v170 = m.ExcPending
													if v170 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(351519), int32(0))
														mBase = m.M
														v174 = m.ExcPending
														if v174 != 0 {
															return int32(0)
														} else {
															F_errhint(m, int32(674589), int32(0))
															mBase = m.M
															v178 = m.ExcPending
															if v178 != 0 {
																return int32(0)
															} else {
																v179 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																F_parser_errposition(m, l0, v179)
																mBase = m.M
																v181 = m.ExcPending
																if v181 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(517271), int32(3510), int32(119464))
																	mBase = m.M
																	v186 = m.ExcPending
																	if v186 != 0 {
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
												v106 = v30
												m.G0 = v10 + int32(16)
												return v106
											}
										}
									}
								} else {
									if v48 == int32(17) {
										v75 = v50
										if v75 != int32(1) {
											v106 = v30
											m.G0 = v10 + int32(16)
											return v106
										} else {
											v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
											if v48 != int32(17) {
												if v78 == int32(0) {
													v106 = v30
													m.G0 = v10 + int32(16)
													return v106
												} else {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v89 = m.ExcPending
														if v89 != 0 {
															return int32(0)
														} else {
															v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
															F_parser_errposition(m, l0, v90)
															mBase = m.M
															v92 = m.ExcPending
															if v92 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(171293), int32(0))
																mBase = m.M
																v96 = m.ExcPending
																if v96 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(517271), int32(3503), int32(119464))
																	mBase = m.M
																	v101 = m.ExcPending
																	if v101 != 0 {
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
												if base.Ui32(int32(2)) <= base.Ui32(v78) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v167 = m.ExcPending
													if v167 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v170 = m.ExcPending
														if v170 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(351519), int32(0))
															mBase = m.M
															v174 = m.ExcPending
															if v174 != 0 {
																return int32(0)
															} else {
																F_errhint(m, int32(674589), int32(0))
																mBase = m.M
																v178 = m.ExcPending
																if v178 != 0 {
																	return int32(0)
																} else {
																	v179 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																	F_parser_errposition(m, l0, v179)
																	mBase = m.M
																	v181 = m.ExcPending
																	if v181 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(517271), int32(3510), int32(119464))
																		mBase = m.M
																		v186 = m.ExcPending
																		if v186 != 0 {
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
													v106 = v30
													m.G0 = v10 + int32(16)
													return v106
												}
											}
										}
									} else {
										if v48 == int32(114) {
											v75 = v50
											if v75 != int32(1) {
												v106 = v30
												m.G0 = v10 + int32(16)
												return v106
											} else {
												v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
												if v48 != int32(17) {
													if v78 == int32(0) {
														v106 = v30
														m.G0 = v10 + int32(16)
														return v106
													} else {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v86 = m.ExcPending
														if v86 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v89 = m.ExcPending
															if v89 != 0 {
																return int32(0)
															} else {
																v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																F_parser_errposition(m, l0, v90)
																mBase = m.M
																v92 = m.ExcPending
																if v92 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(171293), int32(0))
																	mBase = m.M
																	v96 = m.ExcPending
																	if v96 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(517271), int32(3503), int32(119464))
																		mBase = m.M
																		v101 = m.ExcPending
																		if v101 != 0 {
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
													if base.Ui32(int32(2)) <= base.Ui32(v78) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v167 = m.ExcPending
														if v167 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return int32(0)
															} else {
																F_errmsg(m, int32(351519), int32(0))
																mBase = m.M
																v174 = m.ExcPending
																if v174 != 0 {
																	return int32(0)
																} else {
																	F_errhint(m, int32(674589), int32(0))
																	mBase = m.M
																	v178 = m.ExcPending
																	if v178 != 0 {
																		return int32(0)
																	} else {
																		v179 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																		F_parser_errposition(m, l0, v179)
																		mBase = m.M
																		v181 = m.ExcPending
																		if v181 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(517271), int32(3510), int32(119464))
																			mBase = m.M
																			v186 = m.ExcPending
																			if v186 != 0 {
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
														v106 = v30
														m.G0 = v10 + int32(16)
														return v106
													}
												}
											}
										} else {
											if v48 == int32(3802) {
												v75 = v50
												if v75 != int32(1) {
													v106 = v30
													m.G0 = v10 + int32(16)
													return v106
												} else {
													v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
													if v48 != int32(17) {
														if v78 == int32(0) {
															v106 = v30
															m.G0 = v10 + int32(16)
															return v106
														} else {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v86 = m.ExcPending
															if v86 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(1088))
																mBase = m.M
																v89 = m.ExcPending
																if v89 != 0 {
																	return int32(0)
																} else {
																	v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																	F_parser_errposition(m, l0, v90)
																	mBase = m.M
																	v92 = m.ExcPending
																	if v92 != 0 {
																		return int32(0)
																	} else {
																		F_errmsg(m, int32(171293), int32(0))
																		mBase = m.M
																		v96 = m.ExcPending
																		if v96 != 0 {
																			return int32(0)
																		} else {
																			F_errfinish(m, int32(517271), int32(3503), int32(119464))
																			mBase = m.M
																			v101 = m.ExcPending
																			if v101 != 0 {
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
														if base.Ui32(int32(2)) <= base.Ui32(v78) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v167 = m.ExcPending
															if v167 != 0 {
																return int32(0)
															} else {
																F_errcode(m, int32(1088))
																mBase = m.M
																v170 = m.ExcPending
																if v170 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(351519), int32(0))
																	mBase = m.M
																	v174 = m.ExcPending
																	if v174 != 0 {
																		return int32(0)
																	} else {
																		F_errhint(m, int32(674589), int32(0))
																		mBase = m.M
																		v178 = m.ExcPending
																		if v178 != 0 {
																			return int32(0)
																		} else {
																			v179 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																			F_parser_errposition(m, l0, v179)
																			mBase = m.M
																			v181 = m.ExcPending
																			if v181 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(517271), int32(3510), int32(119464))
																				mBase = m.M
																				v186 = m.ExcPending
																				if v186 != 0 {
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
															v106 = v30
															m.G0 = v10 + int32(16)
															return v106
														}
													}
												}
											} else {
												F_get_type_category_preferred(m, v48, v10+int32(15), v10+int32(14))
												mBase = m.M
												v70 = m.ExcPending
												if v70 != 0 {
													return int32(0)
												} else {
													v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
													if v71 != int32(83) {
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v148 = m.ExcPending
														if v148 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int32(0)
															} else {
																v152 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																F_parser_errposition(m, l0, v152)
																mBase = m.M
																v154 = m.ExcPending
																if v154 != 0 {
																	return int32(0)
																} else {
																	F_errmsg(m, int32(171241), int32(0))
																	mBase = m.M
																	v158 = m.ExcPending
																	if v158 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(517271), int32(3490), int32(119464))
																		mBase = m.M
																		v163 = m.ExcPending
																		if v163 != 0 {
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
														v74 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
														v75 = v74
														if v75 != int32(1) {
															v106 = v30
															m.G0 = v10 + int32(16)
															return v106
														} else {
															v78 = *(*int32)(unsafe.Add(mBase, uint32(v49)+8))
															if v48 != int32(17) {
																if v78 == int32(0) {
																	v106 = v30
																	m.G0 = v10 + int32(16)
																	return v106
																} else {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v86 = m.ExcPending
																	if v86 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(1088))
																		mBase = m.M
																		v89 = m.ExcPending
																		if v89 != 0 {
																			return int32(0)
																		} else {
																			v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																			F_parser_errposition(m, l0, v90)
																			mBase = m.M
																			v92 = m.ExcPending
																			if v92 != 0 {
																				return int32(0)
																			} else {
																				F_errmsg(m, int32(171293), int32(0))
																				mBase = m.M
																				v96 = m.ExcPending
																				if v96 != 0 {
																					return int32(0)
																				} else {
																					F_errfinish(m, int32(517271), int32(3503), int32(119464))
																					mBase = m.M
																					v101 = m.ExcPending
																					if v101 != 0 {
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
																if base.Ui32(int32(2)) <= base.Ui32(v78) {
																	F_errstart_cold(m, int32(21), int32(0))
																	mBase = m.M
																	v167 = m.ExcPending
																	if v167 != 0 {
																		return int32(0)
																	} else {
																		F_errcode(m, int32(1088))
																		mBase = m.M
																		v170 = m.ExcPending
																		if v170 != 0 {
																			return int32(0)
																		} else {
																			F_errmsg(m, int32(351519), int32(0))
																			mBase = m.M
																			v174 = m.ExcPending
																			if v174 != 0 {
																				return int32(0)
																			} else {
																				F_errhint(m, int32(674589), int32(0))
																				mBase = m.M
																				v178 = m.ExcPending
																				if v178 != 0 {
																					return int32(0)
																				} else {
																					v179 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
																					F_parser_errposition(m, l0, v179)
																					mBase = m.M
																					v181 = m.ExcPending
																					if v181 != 0 {
																						return int32(0)
																					} else {
																						F_errfinish(m, int32(517271), int32(3510), int32(119464))
																						mBase = m.M
																						v186 = m.ExcPending
																						if v186 != 0 {
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
																	v106 = v30
																	m.G0 = v10 + int32(16)
																	return v106
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
}
func F_transformJsonTableColumns(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
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
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
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
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v332 int32
	_ = v332
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v465 int32
	_ = v465
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v504 int32
	_ = v504
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v27 = base.B2i32(v24 == int32(1))
	goto L3
L2:
	;
	v27 = v5
	goto L3
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v30 = F_exprType(m, v29)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v36 = v35
	goto L8
L7:
	;
	v36 = v5
	goto L8
L8:
	;
	if l1 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v380 = int32(0)
	if l1 == v380 {
		v465 = v380
		goto L88
	} else {
		goto L89
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v354 = v34
	goto L12
L12:
	;
	if v354 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L13:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v354 = v349
	goto L12
L14:
	;
	v49 = v5
	v54 = v5
	goto L16
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L4
	} else {
		goto L79
	}
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v56+v49<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v61 == int32(4) {
		v292 = v54
		goto L19
	} else {
		goto L20
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L74
	}
L18:
	;
	goto L17
L19:
	;
	v295 = v49 + int32(1)
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v295 < v296 {
		v49 = v295
		v54 = v292
		goto L16
	} else {
		goto L73
	}
L20:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v66 = F_pstrdup(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	v68 = F_makeString(m, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v70 = F_lappend(m, v64, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v70
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	switch v73 {
	case 0:
		goto L27
	case 1:
		goto L26
	case 2, 3:
		goto L25
	case 4:
		v292 = v54
		goto L19
	default:
		goto L15
	}
L24:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v262 = F_lappend_oid(m, v261, v260)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L4
	} else {
		goto L69
	}
L25:
	;
	v145 = F_palloc0(m, int32(16))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L46
	}
L26:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	F_typenameTypeIdAndMod(m, v28, v82, v17+int32(-36), v17+int32(-40))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	if v54 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = int32(-1)
	v76 = int32(23)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v76
	v79 = int32(0)
	v248 = v79
	v250 = v79
	v258 = int32(1)
	v260 = v76
	goto L24
L29:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v95 = v89
	goto L31
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = int32(3)
	goto L25
L31:
	;
	v106 = F_get_typtype(m, v95)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	if v122 != 0 {
		goto L30
	} else {
		goto L44
	}
L33:
	;
	if v95 == int32(114) {
		goto L30
	} else {
		goto L34
	}
L34:
	;
	if v95 == int32(2249) {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	if v95 == int32(3802) {
		goto L30
	} else {
		goto L36
	}
L36:
	;
	v114 = F_get_element_type(m, v95)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	if v114 != 0 {
		goto L30
	} else {
		goto L38
	}
L38:
	;
	if v106 == int32(99) {
		goto L30
	} else {
		goto L39
	}
L39:
	;
	if v106 == int32(100) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v120 = F_getBaseType(m, v95)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L32
L43:
	;
	v95 = v120
	goto L31
L44:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	if v123 == int32(0) {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	goto L30
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = int32(34)
	*(*int64)(unsafe.Add(mBase, uint32(v145)+8)) = int64(4294967295)
	*(*int32)(unsafe.Add(mBase, uint32(v145)+4)) = v30
	v153 = F_palloc0(m, int32(48))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = int32(122)
	v157 = int32(2)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	if v158 == int32(1) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v163 = v157
	goto L50
L49:
	;
	v163 = base.B2i32(v158 != v157)
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+4)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	v166 = F_pstrdup(m, v165)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+8)) = v166
	v169 = int32(0)
	v173 = F_makeJsonFormat(m, v169, v169, int32(-1))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L52
	}
L52:
	;
	v175 = F_makeJsonValueExpr(m, v145, v169, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v175
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v178 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v153)+16)) = v198
	v202 = F_palloc0(m, int32(12))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L4
	} else {
		goto L62
	}
L55:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+4))
	v198 = v179
	goto L54
L56:
	;
	goto L57
L57:
	;
	F_initStringInfo(m, v17+int32(-32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L4
	} else {
		goto L58
	}
L58:
	;
	F_appendStringInfoString(m, v17+int32(-32), int32(691213))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v60)+8))
	F_escape_json(m, v17+int32(-32), v191)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v196 = F_makeStringConst(m, v194, int32(-1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v198 = v196
	goto L54
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v202))) = int32(120)
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = v202
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v60)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v202)+4)) = v207
	v210 = F_palloc0(m, int32(16))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(43)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v153)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+8)) = v210
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v153)+24))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)+8))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v60)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v217)+4)) = v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v60)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+28)) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v60)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+32)) = v222
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v60)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+40)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v60)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+36)) = v226
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v60)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+44)) = v228
	v231 = F_transformExpr(m, v28, v153, int32(5))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	F_assign_expr_collations(m, v28, v231)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L65
	}
L65:
	;
	v235 = F_exprType(m, v231)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v235
	v238 = F_exprTypmod(m, v231)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v238
	v241 = F_exprCollation(m, v231)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
	v248 = v231
	v250 = v241
	v258 = v54
	v260 = v243
	goto L24
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v262
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v267 = F_lappend_int(m, v265, v266)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L4
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v267
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v271 = F_lappend_oid(m, v270, v250)
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v271
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v275 = F_lappend(m, v274, v248)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v275
	v292 = v258
	goto L19
L73:
	;
	goto L13
L74:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(459461), int32(0))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v60)+44))
	F_parser_errposition(m, v28, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(521881), int32(294), int32(157540))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L4
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v60)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v321
	F_errmsg_internal(m, int32(506106), v17+int32(-48))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(521881), int32(343), int32(157540))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v378 = base.B2i32(v36 == int32(0))
	v379 = int32(-1)
	goto L9
L83:
	;
	goto L84
L84:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v354)+4))
	v375 = base.B2i32(v372 == v36)
	if v372 == v36 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v376 = int32(-1)
	goto L87
L86:
	;
	v376 = v372 - int32(1)
	goto L87
L87:
	;
	v378 = v375
	v379 = v376
	goto L9
L88:
	;
	v474 = F_palloc0(m, int32(24))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L4
	} else {
		goto L107
	}
L89:
	;
	v383 = int32(0)
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v384 <= v383 {
		v465 = v380
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v391 = v383
	v395 = v380
	goto L91
L91:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v403+v391<<(uint(int32(2))%32))))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v408 != int32(4) {
		v451 = v395
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v465 = v451
	goto L88
L93:
	;
	v454 = v391 + int32(1)
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v454 < v455 {
		v391 = v454
		v395 = v451
		goto L91
	} else {
		goto L106
	}
L94:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	if v412 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v415 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v415
	v424 = F_pg_snprintf(m, v17+int32(-32), int32(32), int32(486122), v19)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L4
	} else {
		goto L98
	}
L96:
	;
	v437 = v411
	goto L97
L97:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v407)+32))
	v439 = F_transformJsonTableColumns(m, l0, v438, l2, v437)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L4
	} else {
		goto L101
	}
L98:
	;
	v428 = F_pstrdup(m, v17+int32(-32))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L4
	} else {
		goto L99
	}
L99:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v431 = F_lappend(m, v430, v428)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L4
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v431
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v434)+8)) = v428
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v407)+16))
	v437 = v436
	goto L97
L101:
	;
	if v395 == int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v451 = v439
	goto L93
L103:
	;
	goto L104
L104:
	;
	v444 = F_palloc0(m, int32(12))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L4
	} else {
		goto L105
	}
L105:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v444)+8)) = v439
	*(*int32)(unsafe.Add(mBase, uint32(v444)+4)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v444))) = int32(51)
	v451 = v444
	goto L93
L106:
	;
	goto L92
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = int32(50)
	v479 = int32(-1)
	v480 = int32(0)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v484)+8))
	v486 = F_DirectFunctionCall1Coll(m, int32(488), v480, v485)
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v488 = int32(0)
	v490 = F_makeConst(m, int32(4072), v479, v480, v479, v486, v488, v488)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L4
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474))) = int32(50)
	v494 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v496 = F_palloc0(m, int32(12))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L4
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v496)+8)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v496)+4)) = v490
	*(*int32)(unsafe.Add(mBase, uint32(v496))) = int32(49)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+20)) = v379
	if v378 != 0 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v504 = int32(-1)
	goto L113
L112:
	;
	v504 = v36
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v474)+16)) = v504
	*(*int32)(unsafe.Add(mBase, uint32(v474)+12)) = v465
	*(*uint8)(unsafe.Add(mBase, uint32(v474)+8)) = uint8(v27)
	*(*int32)(unsafe.Add(mBase, uint32(v474)+4)) = v496
	m.G0 = v19 - int32(-64)
	return v474
}
