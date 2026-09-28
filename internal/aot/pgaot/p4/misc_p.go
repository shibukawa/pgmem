package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_PMSignalShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_PMSignalShmemRequest[0]))
	if v8 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_PMSignalShmemRequest_0), int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_PMSignalShmemRequest_1), int32(84), int32(_a_F_PMSignalShmemRequest_2))
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_PMSignalShmemRequest[1])) = v8
		v28 = F_mul_size(m, v8, int32(4))
		mBase = m.M
		v29 = m.ExcPending
		if v29 != 0 {
			return
		} else {
			v30 = F_add_size(m, int32(52), v28)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = int32(_a_F_PMSignalShmemRequest_3)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_PMSignalShmemRequest_4)
				F_ShmemRequestStructWithOpts(m, v5)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					m.G0 = v5 + int32(16)
					return
				}
			}
		}
	}
}
func F_PageGetItemIdCareful_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v16 = l2 + l3<<(uint(int32(2))%32) + int32(20)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v21 = int32(base.Ui32(v17) >> (uint(int32(17)) % 32))
	if base.Ui32(v17&int32(_a_F_PageGetItemIdCareful_2_0)+v21) < base.Ui32(int32(_a_F_PageGetItemIdCareful_2_1)) {
		switch int32(base.Ui32(v17)>>(uint(int32(15))%32)) & int32(3) {
		case 0, 2:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int32(0)
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v84 + int32(4)
					F_errmsg(m, int32(_a_F_PageGetItemIdCareful_2_2), v10+int32(80))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
						*(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = int32(base.Ui32(v95)>>(uint(int32(15))%32)) & int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l3
						*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(base.Ui32(v95) >> (uint(int32(17)) % 32))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v95 & int32(_a_F_PageGetItemIdCareful_2_0)
						F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_2_3), v10+int32(48))
						mBase = m.M
						v113 = m.ExcPending
						if v113 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_PageGetItemIdCareful_2_4), int32(3522), int32(_a_F_PageGetItemIdCareful_2_5))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
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
		default:
			if v21 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
						return int32(0)
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = v84 + int32(4)
						F_errmsg(m, int32(_a_F_PageGetItemIdCareful_2_2), v10+int32(80))
						mBase = m.M
						v92 = m.ExcPending
						if v92 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
							*(*int32)(unsafe.Add(mBase, uint32(v10-int32(-64)))) = int32(base.Ui32(v95)>>(uint(int32(15))%32)) & int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v10)+60)) = int32(base.Ui32(v95) >> (uint(int32(17)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = v95 & int32(_a_F_PageGetItemIdCareful_2_0)
							F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_2_3), v10+int32(48))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_PageGetItemIdCareful_2_4), int32(3522), int32(_a_F_PageGetItemIdCareful_2_5))
								mBase = m.M
								v118 = m.ExcPending
								if v118 != 0 {
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
				m.G0 = v10 + int32(96)
				return v16
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(33557032))
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v45 + int32(4)
				F_errmsg(m, int32(_a_F_PageGetItemIdCareful_2_6), v10+int32(32))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(base.Ui32(v54)>>(uint(int32(15))%32)) & int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(base.Ui32(v54) >> (uint(int32(17)) % 32))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v54 & int32(_a_F_PageGetItemIdCareful_2_0)
					F_errdetail_internal(m, int32(_a_F_PageGetItemIdCareful_2_3), v10)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_PageGetItemIdCareful_2_4), int32(3506), int32(_a_F_PageGetItemIdCareful_2_5))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
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
func F_ParseComplexProjection(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v7 != int32(6) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return int32(0)
L2:
	;
	if v29 == int32(0) {
		goto L1
	} else {
		goto L14
	}
L3:
	;
	v27 = F_get_expr_result_tupdesc(m, l2, int32(1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L13
	}
L4:
	;
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
	if v10 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v13 = F_GetNSItemByVar(m, l0, l2)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v21 != int32(2249) {
		goto L3
	} else {
		goto L11
	}
L8:
	;
	return int32(0)
L9:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v18 = F_scanNSItemForColumn(m, l0, v13, v17, l1, l3)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	return v18
L11:
	;
	v24 = F_expandRecordVariable(m, l0, l2)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v29 = v24
	goto L2
L13:
	;
	v29 = v27
	goto L2
L14:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	if v32 <= int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v41 = int32(0)
	goto L16
L16:
	;
	v49 = v29 + v32<<(uint(int32(3))%32) + int32(28) + v41*int32(100)
	v51 = v49 + int32(4)
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if base.B2i32(v54 == int32(0))|base.B2i32(v54 != v57) != 0 {
		v75 = v54
		v76 = v57
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L1
L18:
	;
	v96 = v41 + int32(1)
	if v96 != v32 {
		v41 = v96
		goto L16
	} else {
		goto L29
	}
L19:
	;
	if v75-v76 != 0 {
		goto L18
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v60 = l1
	v61 = v51
	goto L22
L22:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
	if v65 == int32(0) {
		v75 = v65
		v76 = v64
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v75 = v65
	v76 = v64
	goto L20
L24:
	;
	v68 = int32(1)
	if v65 == v64 {
		v60 = v60 + v68
		v61 = v61 + v68
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+91)))
	if v78 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v80 = F_palloc0(m, int32(24))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v83 = v41 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v80)+8)) = uint16(v83)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = int32(25)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v88
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v49)+96))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v92
	return v80
L29:
	;
	goto L17
}
func F_ParseLongOption(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	v5 = int32(_a_F_ParseLongOption_0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(*(*int8)(unsafe.Add(mBase, _c_F_ParseLongOption[0])))
	if v13 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v208
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v211 = v210
	goto L58
L2:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v71))))
	if v73 == int32(61) {
		goto L20
	} else {
		goto L21
	}
L3:
	;
	m.G0 = v11 + int32(32)
	v71 = v64 - l0
	goto L2
L4:
	;
	F___memset(m, v11, int32(0), int32(32))
	mBase = m.M
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParseLongOption[0])))
	if v19 != 0 {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v14 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ParseLongOption[1])))
	if v14 != 0 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v15 = F___strchrnul(m, l0, v13)
	mBase = m.M
	v64 = v15
	goto L3
L8:
	;
	goto L7
L9:
	;
	v21 = v5
	v22 = v19
	goto L12
L10:
	;
	goto L11
L11:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v43 == int32(0) {
		v64 = l0
		goto L3
	} else {
		goto L15
	}
L12:
	;
	v29 = v11 + int32(base.Ui32(v22)>>(uint(int32(3))%32))&int32(28)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	v31 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v30 | v31<<(uint(v22)%32)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)))
	if v35 != 0 {
		v21 = v21 + v31
		v22 = v35
		goto L12
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	goto L13
L15:
	;
	v47 = l0
	v48 = v43
	goto L16
L16:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(base.Ui32(v48)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v56)>>(uint(v48)%32))&int32(1) != 0 {
		v64 = v47
		goto L3
	} else {
		goto L18
	}
L17:
	;
	v64 = v62
	goto L3
L18:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+1)))
	v62 = v47 + int32(1)
	if v60 != 0 {
		v47 = v62
		v48 = v60
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v77 = v71 + int32(1)
	v78 = F_palloc(m, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	v202 = F_pstrdup(m, l0)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L23
	} else {
		goto L57
	}
L23:
	;
	return
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v78
	if v77 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v200 = F_pstrdup(m, l0+v77)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L56
	}
L26:
	;
	v196 = F_strlen(m, v192)
	mBase = m.M
	goto L25
L27:
	;
	v192 = l0
	goto L26
L28:
	;
	goto L29
L29:
	;
	v86 = v77 - int32(1)
	if (v78^l0)&int32(3) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v189 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v189)
	v192 = v185
	goto L26
L31:
	;
	v170 = v165
	v171 = v166
	v172 = v167
	goto L52
L32:
	;
	if v160 == int32(0) {
		v185 = v158
		v186 = v159
		goto L30
	} else {
		goto L51
	}
L33:
	;
	v158 = l0
	v159 = v78
	v160 = v86
	goto L32
L34:
	;
	goto L35
L35:
	;
	v90 = int32(0)
	if base.B2i32(l0&int32(3) == v90)|base.B2i32(v86 == v90) == v90 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v126 == int32(0) {
		v185 = v123
		v186 = v124
		goto L30
	} else {
		goto L45
	}
L37:
	;
	v102 = l0
	v103 = v78
	v104 = v86
	goto L40
L38:
	;
	goto L39
L39:
	;
	v123 = l0
	v124 = v78
	v125 = v86
	v126 = base.B2i32(v86 != v90)
	goto L36
L40:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	*(*uint8)(unsafe.Add(mBase, uint32(v103))) = uint8(v106)
	if v106 == int32(0) {
		v165 = v102
		v166 = v103
		v167 = v104
		goto L31
	} else {
		goto L42
	}
L41:
	;
	v123 = v117
	v124 = v111
	v125 = v113
	v126 = v115
	goto L36
L42:
	;
	v110 = int32(1)
	v111 = v103 + v110
	v113 = v104 - v110
	v114 = int32(0)
	v115 = base.B2i32(v113 != v114)
	v117 = v102 + v110
	if v117&int32(3) == v114 {
		v123 = v117
		v124 = v111
		v125 = v113
		v126 = v115
		goto L36
	} else {
		goto L43
	}
L43:
	;
	if v113 != 0 {
		v102 = v117
		v103 = v111
		v104 = v113
		goto L40
	} else {
		goto L44
	}
L44:
	;
	goto L41
L45:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
	if base.B2i32(v129 == int32(0))|base.B2i32(base.Ui32(v125) < base.Ui32(int32(4))) != 0 {
		v158 = v123
		v159 = v124
		v160 = v125
		goto L32
	} else {
		goto L46
	}
L46:
	;
	v136 = v123
	v137 = v124
	v138 = v125
	goto L47
L47:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v144 = int32(-2139062144)
	if (int32(16843008)-v141|v141)&v144 != v144 {
		v165 = v136
		v166 = v137
		v167 = v138
		goto L31
	} else {
		goto L49
	}
L48:
	;
	v158 = v152
	v159 = v150
	v160 = v154
	goto L32
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v137))) = v141
	v149 = int32(4)
	v150 = v137 + v149
	v152 = v136 + v149
	v154 = v138 - v149
	if base.Ui32(int32(3)) < base.Ui32(v154) {
		v136 = v152
		v137 = v150
		v138 = v154
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v165 = v158
	v166 = v159
	v167 = v160
	goto L31
L52:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170))))
	*(*uint8)(unsafe.Add(mBase, uint32(v171))) = uint8(v174)
	if v174 == int32(0) {
		v185 = v170
		v186 = v171
		goto L30
	} else {
		goto L54
	}
L53:
	;
	v185 = v181
	v186 = v179
	goto L30
L54:
	;
	v178 = int32(1)
	v179 = v171 + v178
	v181 = v170 + v178
	v183 = v172 - v178
	if v183 != 0 {
		v170 = v181
		v171 = v179
		v172 = v183
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v208 = v200
	goto L1
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v202
	v208 = int32(0)
	goto L1
L58:
	;
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v211))))
	if v215 != int32(45) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v211 = v211 + int32(1)
	goto L58
L61:
	;
	if v215 != 0 {
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v218 = int32(95)
	*(*uint8)(unsafe.Add(mBase, uint32(v211))) = uint8(v218)
	goto L60
L64:
	;
	return
}
func F_PreCommit_CheckForSerializationFailure(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v94 int64
	_ = v94
	var v96 int64
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[0]))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	F_LWLockRelease(m, v149+int32(3584))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L6
	} else {
		goto L36
	}
L2:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	F_LWLockRelease(m, v119+int32(3584))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L6
	} else {
		goto L29
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	v17 = F_LWLockAcquire(m, v13+int32(3584), int32(0))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	return
L6:
	;
	return
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[0]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	if v21&int32(2056) == int32(8) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v93 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[2]))
	v94 = *(*int64)(unsafe.Add(mBase, uint32(v93)+32))
	v96 = v94 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+32)) = v96
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v20)+108))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+108)) = v99 | int32(2)
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_CheckForSerializationFailure[1]))
	F_LWLockRelease(m, v104+int32(3584))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L6
	} else {
		goto L28
	}
L10:
	;
	v30 = v20 + int32(40)
	if v26 == v30 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v33 = v26
	goto L12
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+108))
	if v42&int32(9) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v81 != v30 {
		v33 = v81
		goto L12
	} else {
		goto L27
	}
L15:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+44))
	if v45 == int32(0) {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v49 = v41 + int32(40)
	if v45 == v49 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	v53 = v45
	goto L18
L18:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v53)+8))
	if v20 != v60 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L14
L20:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v70 != v49 {
		v53 = v70
		goto L18
	} else {
		goto L26
	}
L21:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+108)))
	if v62&int32(41) != 0 {
		goto L20
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	if v42&int32(2) != 0 {
		goto L1
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+108)) = v42 | int32(8)
	goto L14
L26:
	;
	goto L19
L27:
	;
	goto L13
L28:
	;
	goto L5
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L6
	} else {
		goto L30
	}
L30:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_PreCommit_CheckForSerializationFailure_0), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L6
	} else {
		goto L32
	}
L32:
	;
	F_errdetail_internal(m, int32(_a_F_PreCommit_CheckForSerializationFailure_1), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L6
	} else {
		goto L33
	}
L33:
	;
	F_errhint(m, int32(_a_F_PreCommit_CheckForSerializationFailure_2), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L6
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_PreCommit_CheckForSerializationFailure_3), int32(_a_F_PreCommit_CheckForSerializationFailure_4), int32(_a_F_PreCommit_CheckForSerializationFailure_5))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L6
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L6
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(16777220))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L6
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_PreCommit_CheckForSerializationFailure_0), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L6
	} else {
		goto L39
	}
L39:
	;
	F_errdetail_internal(m, int32(_a_F_PreCommit_CheckForSerializationFailure_6), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L6
	} else {
		goto L40
	}
L40:
	;
	F_errhint(m, int32(_a_F_PreCommit_CheckForSerializationFailure_2), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_PreCommit_CheckForSerializationFailure_3), int32(_a_F_PreCommit_CheckForSerializationFailure_7), int32(_a_F_PreCommit_CheckForSerializationFailure_5))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L6
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PreCommit_Portals(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
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
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = v9 + int32(12)
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_Portals[0]))
	F_hash_seq_init(m, v12, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = F_hash_seq_search(m, v12)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L41
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L38
	}
L5:
	;
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v22 = v19
	v26 = v2
	goto L9
L7:
	;
	v91 = v2
	goto L8
L8:
	;
	m.G0 = v9 + int32(32)
	return v91
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+64))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+84)))
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v91 = v81
	goto L8
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+85)))
	if v31 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
	if v34 == int32(3) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	v84 = F_hash_seq_search(m, v9+int32(12))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L36
	}
L16:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+112))
	if v37 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v27)+20))
	v49 = int32(0)
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+76)))
	if base.B2i32(v48 == v49)|(base.B2i32(v51&int32(32) == v49)|base.B2i32(v34 != int32(2))) == v49 {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v38 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+100)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v44
	v81 = v26
	goto L15
L22:
	;
	F_UnregisterSnapshotFromOwner(m, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = int32(0)
	goto L21
L25:
	;
	goto L24
L26:
	;
	v70 = v9 + int32(12)
	F_hash_seq_term(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L34
	}
L27:
	;
	if l0 != 0 {
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v48 == int32(0) {
		v81 = v26
		goto L15
	} else {
		goto L32
	}
L30:
	;
	F_HoldPortal(m, v27)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	F_PortalDrop(m, v27, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L26
L34:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_Portals[0]))
	F_hash_seq_init(m, v70, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v81 = int32(1)
	goto L15
L36:
	;
	if v84 != 0 {
		v22 = v84
		v26 = v81
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L10
L38:
	;
	F_errmsg_internal(m, int32(_a_F_PreCommit_Portals_0), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_PreCommit_Portals_1), int32(696), int32(_a_F_PreCommit_Portals_2))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errmsg(m, int32(_a_F_PreCommit_Portals_3), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_PreCommit_Portals_1), int32(739), int32(_a_F_PreCommit_Portals_2))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
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
func F_PreCommit_on_commit_actions(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
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
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	v1 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_PreCommit_on_commit_actions[0]))
	if v12 == v1 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(16)
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if int32(0) < v15 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v19 = v1
	v20 = v1
	v23 = v1
	goto L6
L4:
	;
	v53 = v1
	v56 = v1
	goto L5
L5:
	;
	if v53 != 0 {
		goto L17
	} else {
		goto L18
	}
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v19<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v29 != 0 {
		v45 = v20
		v46 = v23
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v53 = v45
	v56 = v46
	goto L5
L8:
	;
	v48 = v19 + int32(1)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v48 < v49 {
		v19 = v48
		v20 = v45
		v23 = v46
		goto L6
	} else {
		goto L16
	}
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	switch v30 - int32(2) {
	case 0:
		goto L11
	case 1:
		goto L10
	default:
		v45 = v20
		v46 = v23
		goto L8
	}
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v43 = F_lappend_oid(m, v23, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L13
	} else {
		goto L15
	}
L11:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreCommit_on_commit_actions[1])))
	if v34&int32(1) == int32(0) {
		v45 = v20
		v46 = v23
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v40 = F_lappend_oid(m, v20, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return
L14:
	;
	v45 = v40
	v46 = v23
	goto L8
L15:
	;
	v45 = v20
	v46 = v43
	goto L8
L16:
	;
	goto L7
L17:
	;
	v57 = int32(0)
	if v53 == v57 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	goto L19
L19:
	;
	if v56 == int32(0) {
		goto L1
	} else {
		goto L50
	}
L20:
	;
	goto L19
L21:
	;
	F_heap_truncate_check_FKs(m, int32(0), int32(1))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L13
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if int32(0) < v65 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L20
L25:
	;
	v68 = v57
	v69 = v57
	goto L28
L26:
	;
	v89 = v57
	goto L27
L27:
	;
	F_heap_truncate_check_FKs(m, v89, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L13
	} else {
		goto L33
	}
L28:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v53)+12))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74+v68<<(uint(int32(2))%32))))
	v80 = F_table_open(m, v78, int32(8))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L13
	} else {
		goto L30
	}
L29:
	;
	v89 = v82
	goto L27
L30:
	;
	v82 = F_lappend(m, v69, v80)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	v85 = v68 + int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v53)+4))
	if v85 < v86 {
		v68 = v85
		v69 = v82
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	if v89 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	goto L20
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v99 <= int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v105 = int32(0)
	goto L37
L37:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109+v105<<(uint(int32(2))%32))))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+119)))
	if v115 == int32(112) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L34
L39:
	;
	F_relation_close(m, v113, int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L13
	} else {
		goto L48
	}
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v113)+188))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v118)+116))
	m.T0[v119].(func(*base.Module, int32))(m, v113)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L13
	} else {
		goto L41
	}
L41:
	;
	F_RelationTruncateIndexes(m, v113)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L42
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+112))
	if v125 == int32(0) {
		goto L39
	} else {
		goto L43
	}
L43:
	;
	v129 = F_table_open(m, v125, int32(8))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+188))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)+116))
	m.T0[v132].(func(*base.Module, int32))(m, v129)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	F_RelationTruncateIndexes(m, v129)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	F_relation_close(m, v129, int32(0))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L13
	} else {
		goto L47
	}
L47:
	;
	goto L39
L48:
	;
	v145 = v105 + int32(1)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v145 < v146 {
		v105 = v145
		goto L37
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v168 = F_new_object_addresses(m)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L51
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if int32(0) < v170 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v175 = int32(0)
	goto L55
L53:
	;
	goto L54
L54:
	;
	v204 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L13
	} else {
		goto L59
	}
L55:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = int32(1259)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v180+v175<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v186
	F_add_exact_object_address(m, v9+int32(4), v168)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L57
	}
L56:
	;
	goto L54
L57:
	;
	v195 = v175 + int32(1)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	if v195 < v196 {
		v175 = v195
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	F_PushActiveSnapshot(m, v204)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L13
	} else {
		goto L60
	}
L60:
	;
	F_performMultipleDeletions(m, v168, int32(1), int32(5))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L13
	} else {
		goto L61
	}
L61:
	;
	F_PopActiveSnapshot(m)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L62
	}
L62:
	;
	goto L1
}
func F_PrefetchSharedBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v19
	v24 = v11 + int32(12)
	v25 = F_BufTableHashCode(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_PrefetchSharedBuffer[0]))
		v35 = v28 + v25&int32(127)<<(uint(int32(7))%32) + int32(_a_F_PrefetchSharedBuffer_0)
		v37 = F_LWLockAcquire(m, v35, int32(1))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return
		} else {
			v39 = F_BufTableLookup(m, v24, v25)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				F_LWLockRelease(m, v35)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					if v39 < int32(0) {
						v46 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PrefetchSharedBuffer[1])))
						if v46&int32(1) != 0 {
							m.G0 = v11 + int32(32)
							return
						} else {
							v50 = F_smgrprefetch(m, l1, l2, l3, int32(1))
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								if v50 == int32(0) {
								} else {
									v54 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v54)
								}
								m.G0 = v11 + int32(32)
								return
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v39 + int32(1)
						m.G0 = v11 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_PrepareInplaceInvalidationState(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	v5 = F_palloc0(m, int32(20))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[0]))
		if v10 != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v11
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v14
		} else {
			v18 = int64(0)
			*(*int64)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[1])) = v18
			*(*int64)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[2])) = v18
		}
		*(*int32)(unsafe.Add(mBase, _c_F_PrepareInplaceInvalidationState[3])) = v5
		return v5
	}
}
func F_PrepareSortSupportComparisonShim(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = F_MemoryContextAlloc(m, v4, int32(88))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		F_fmgr_info_cxt(m, l0, v6, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v6)+36)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v6
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v15 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+80)) = uint8(v15)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+64)) = uint8(v15)
			v19 = int32(2)
			*(*uint16)(unsafe.Add(mBase, uint32(v6)+50)) = uint16(v19)
			*(*uint8)(unsafe.Add(mBase, uint32(v6)+48)) = uint8(v15)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = v14
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(2042)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v6
			return
		}
	}
}
func F_PreventCommandIfParallelMode(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_PreventCommandIfParallelMode[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+72))
	if v10 != 0 {
		v13 = int32(1)
	} else {
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+76)))
		v13 = v12
	}
	if v13&int32(1) != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandIfParallelMode_0), v5)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandIfParallelMode_1), int32(431), int32(_a_F_PreventCommandIfParallelMode_2))
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
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PreventCommandIfReadOnly(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v8 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_PreventCommandIfReadOnly[0])))
	if v8 == int32(1) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			F_errcode(m, int32(100663618))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
				F_errmsg(m, int32(_a_F_PreventCommandIfReadOnly_0), v5)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventCommandIfReadOnly_1), int32(413), int32(_a_F_PreventCommandIfReadOnly_2))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_PreventInTransactionBlock(m *base.Module, l0 int32, l1 int32) {
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
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_PreventInTransactionBlock[0]))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if base.Ui32(v11) < base.Ui32(int32(2)) {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
		if int32(2) <= v14 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v47 = m.ExcPending
			if v47 != 0 {
				return
			} else {
				F_errcode(m, int32(16777538))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l1
					F_errmsg(m, int32(_a_F_PreventInTransactionBlock_0), v7+int32(16))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_PreventInTransactionBlock_1), int32(3721), int32(_a_F_PreventInTransactionBlock_2))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
			if l0 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(16777538))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
						F_errmsg(m, int32(_a_F_PreventInTransactionBlock_3), v7+int32(32))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_PreventInTransactionBlock_1), int32(3731), int32(_a_F_PreventInTransactionBlock_2))
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
			} else {
				v19 = int32(_a_F_PreventInTransactionBlock_4)
				v21 = *(*int32)(unsafe.Add(mBase, _c_F_PreventInTransactionBlock[1]))
				*(*int32)(unsafe.Add(mBase, _c_F_PreventInTransactionBlock[1])) = v21 | int32(4)
				m.G0 = v7 + int32(48)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			F_errcode(m, int32(16777538))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
				F_errmsg(m, int32(_a_F_PreventInTransactionBlock_5), v7)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_PreventInTransactionBlock_1), int32(3711), int32(_a_F_PreventInTransactionBlock_2))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
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
func F_ProcessCheckpointerInterrupts(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[0]))
	if v2 != 0 {
		F_ProcessProcSignalBarrier(m)
		mBase = m.M
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1]))
			if v6 == int32(0) {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
				if v35 != 0 {
					F_ProcessLogMemoryContextInterrupt(m)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						return
					}
				} else {
					return
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1])) = int32(0)
				F_ProcessConfigFile(m, int32(2))
				mBase = m.M
				v14 = m.ExcPending
				if v14 != 0 {
					return
				} else {
					F_SyncRepUpdateSyncStandbysDefined(m)
					mBase = m.M
					v16 = m.ExcPending
					if v16 != 0 {
						return
					} else {
						F_UpdateFullPageWrites(m)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v21 = F_errstart(m, int32(13), int32(0))
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								if v21 == int32(0) {
									v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
									if v35 != 0 {
										F_ProcessLogMemoryContextInterrupt(m)
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											return
										}
									} else {
										return
									}
								} else {
									F_errmsg_internal(m, int32(_a_F_ProcessCheckpointerInterrupts_0), int32(0))
									mBase = m.M
									v28 = m.ExcPending
									if v28 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_ProcessCheckpointerInterrupts_1), int32(1515), int32(_a_F_ProcessCheckpointerInterrupts_2))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
											if v35 != 0 {
												F_ProcessLogMemoryContextInterrupt(m)
												mBase = m.M
												v37 = m.ExcPending
												if v37 != 0 {
													return
												} else {
													return
												}
											} else {
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
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1]))
		if v6 == int32(0) {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
			if v35 != 0 {
				F_ProcessLogMemoryContextInterrupt(m)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[1])) = int32(0)
			F_ProcessConfigFile(m, int32(2))
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return
			} else {
				F_SyncRepUpdateSyncStandbysDefined(m)
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					F_UpdateFullPageWrites(m)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v21 = F_errstart(m, int32(13), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							if v21 == int32(0) {
								v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
								if v35 != 0 {
									F_ProcessLogMemoryContextInterrupt(m)
									mBase = m.M
									v37 = m.ExcPending
									if v37 != 0 {
										return
									} else {
										return
									}
								} else {
									return
								}
							} else {
								F_errmsg_internal(m, int32(_a_F_ProcessCheckpointerInterrupts_0), int32(0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_ProcessCheckpointerInterrupts_1), int32(1515), int32(_a_F_ProcessCheckpointerInterrupts_2))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										v35 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCheckpointerInterrupts[2]))
										if v35 != 0 {
											F_ProcessLogMemoryContextInterrupt(m)
											mBase = m.M
											v37 = m.ExcPending
											if v37 != 0 {
												return
											} else {
												return
											}
										} else {
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
}
func F_ProcessCopyOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v475 int32
	_ = v475
	var v484 int32
	_ = v484
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v520 int32
	_ = v520
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v565 int32
	_ = v565
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v623 int32
	_ = v623
	var v632 int32
	_ = v632
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int64
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v811 int32
	_ = v811
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v892 int32
	_ = v892
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v936 int32
	_ = v936
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v971 int32
	_ = v971
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1027 int32
	_ = v1027
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1067 int32
	_ = v1067
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1084 int32
	_ = v1084
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1123 int32
	_ = v1123
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1171 int32
	_ = v1171
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1187 int32
	_ = v1187
	var v1190 int32
	_ = v1190
	var v1193 int32
	_ = v1193
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1201 int32
	_ = v1201
	var v1204 int32
	_ = v1204
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1216 int32
	_ = v1216
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1241 int32
	_ = v1241
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1255 int32
	_ = v1255
	var v1264 int32
	_ = v1264
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1276 int32
	_ = v1276
	var v1286 int32
	_ = v1286
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1300 int32
	_ = v1300
	var v1309 int32
	_ = v1309
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1330 int32
	_ = v1330
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1362 int32
	_ = v1362
	var v1365 int32
	_ = v1365
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1377 int32
	_ = v1377
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1389 int32
	_ = v1389
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1412 int32
	_ = v1412
	var v1415 int32
	_ = v1415
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1423 int32
	_ = v1423
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1440 int32
	_ = v1440
	var v1441 int32
	_ = v1441
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1460 int32
	_ = v1460
	var v1469 int32
	_ = v1469
	var v1472 int32
	_ = v1472
	var v1474 int32
	_ = v1474
	var v1483 int32
	_ = v1483
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1505 int32
	_ = v1505
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1528 int32
	_ = v1528
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1549 int32
	_ = v1549
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1572 int32
	_ = v1572
	var v1574 int32
	_ = v1574
	var v1581 int32
	_ = v1581
	var v1584 int32
	_ = v1584
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1607 int32
	_ = v1607
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1640 int64
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1642 int64
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1644 int64
	_ = v1644
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1690 int32
	_ = v1690
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1708 int32
	_ = v1708
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1721 int32
	_ = v1721
	var v1722 int32
	_ = v1722
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1771 int32
	_ = v1771
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1791 int32
	_ = v1791
	var v1795 int32
	_ = v1795
	var v1800 int32
	_ = v1800
	var v1801 int32
	_ = v1801
	var v1804 int32
	_ = v1804
	var v1810 int32
	_ = v1810
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1821 int32
	_ = v1821
	var v1825 int32
	_ = v1825
	var v1830 int32
	_ = v1830
	var v1835 int32
	_ = v1835
	var v1838 int32
	_ = v1838
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1851 int32
	_ = v1851
	var v1855 int32
	_ = v1855
	var v1858 int32
	_ = v1858
	var v1861 int32
	_ = v1861
	var v1862 int32
	_ = v1862
	var v1865 int32
	_ = v1865
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1876 int32
	_ = v1876
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1899 int32
	_ = v1899
	var v1900 int32
	_ = v1900
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1905 int32
	_ = v1905
	var v1909 int32
	_ = v1909
	var v1910 int32
	_ = v1910
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1948 int32
	_ = v1948
	var v1961 int32
	_ = v1961
	var v1962 int32
	_ = v1962
	var v1971 int32
	_ = v1971
	var v1973 int32
	_ = v1973
	var v1977 int32
	_ = v1977
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1986 int32
	_ = v1986
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2015 int32
	_ = v2015
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2040 int32
	_ = v2040
	var v2047 int32
	_ = v2047
	var v2051 int32
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2075 int32
	_ = v2075
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2119 int32
	_ = v2119
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2132 int32
	_ = v2132
	var v2135 int32
	_ = v2135
	var v2139 int32
	_ = v2139
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2148 int32
	_ = v2148
	var v2154 int32
	_ = v2154
	var v2157 int32
	_ = v2157
	var v2164 int32
	_ = v2164
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2174 int32
	_ = v2174
	var v2177 int32
	_ = v2177
	var v2180 int32
	_ = v2180
	var v2183 int32
	_ = v2183
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2199 int32
	_ = v2199
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2217 int32
	_ = v2217
	var v2221 int32
	_ = v2221
	var v2228 int32
	_ = v2228
	var v2231 int32
	_ = v2231
	var v2238 int32
	_ = v2238
	var v2243 int32
	_ = v2243
	var v2248 int32
	_ = v2248
	var v2252 int32
	_ = v2252
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2264 int32
	_ = v2264
	var v2269 int32
	_ = v2269
	var v2273 int32
	_ = v2273
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2283 int32
	_ = v2283
	var v2288 int32
	_ = v2288
	var v2292 int32
	_ = v2292
	var v2295 int32
	_ = v2295
	var v2299 int32
	_ = v2299
	var v2304 int32
	_ = v2304
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2315 int32
	_ = v2315
	var v2320 int32
	_ = v2320
	var v2324 int32
	_ = v2324
	var v2327 int32
	_ = v2327
	var v2331 int32
	_ = v2331
	var v2336 int32
	_ = v2336
	var v2340 int32
	_ = v2340
	var v2343 int32
	_ = v2343
	var v2347 int32
	_ = v2347
	var v2352 int32
	_ = v2352
	var v2356 int32
	_ = v2356
	var v2359 int32
	_ = v2359
	var v2363 int32
	_ = v2363
	var v2368 int32
	_ = v2368
	var v2372 int32
	_ = v2372
	var v2375 int32
	_ = v2375
	var v2376 int32
	_ = v2376
	var v2382 int32
	_ = v2382
	var v2387 int32
	_ = v2387
	var v2391 int32
	_ = v2391
	var v2394 int32
	_ = v2394
	var v2398 int32
	_ = v2398
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2410 int32
	_ = v2410
	var v2417 int32
	_ = v2417
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2424 int32
	_ = v2424
	var v2433 int32
	_ = v2433
	var v2436 int32
	_ = v2436
	var v2445 int32
	_ = v2445
	var v2450 int32
	_ = v2450
	var v2451 int32
	_ = v2451
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2455 int32
	_ = v2455
	var v2459 int32
	_ = v2459
	var v2464 int32
	_ = v2464
	var v2465 int32
	_ = v2465
	var v2467 int32
	_ = v2467
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2480 int32
	_ = v2480
	var v2483 int32
	_ = v2483
	var v2490 int32
	_ = v2490
	var v2492 int32
	_ = v2492
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2514 int32
	_ = v2514
	var v2515 int32
	_ = v2515
	var v2516 int32
	_ = v2516
	var v2517 int32
	_ = v2517
	var v2518 int32
	_ = v2518
	var v2520 int32
	_ = v2520
	var v2526 int32
	_ = v2526
	var v2529 int32
	_ = v2529
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2541 int32
	_ = v2541
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2558 int32
	_ = v2558
	var v2559 int64
	_ = v2559
	var v2562 int32
	_ = v2562
	var v2571 int32
	_ = v2571
	var v2574 int32
	_ = v2574
	var v2581 int32
	_ = v2581
	var v2586 int32
	_ = v2586
	var v2590 int32
	_ = v2590
	var v2593 int32
	_ = v2593
	var v2600 int32
	_ = v2600
	var v2605 int32
	_ = v2605
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2621 int32
	_ = v2621
	var v2626 int32
	_ = v2626
	var v2630 int32
	_ = v2630
	var v2633 int32
	_ = v2633
	var v2642 int32
	_ = v2642
	var v2647 int32
	_ = v2647
	var v2651 int32
	_ = v2651
	var v2654 int32
	_ = v2654
	var v2661 int32
	_ = v2661
	var v2666 int32
	_ = v2666
	var v2670 int32
	_ = v2670
	var v2673 int32
	_ = v2673
	var v2682 int32
	_ = v2682
	var v2687 int32
	_ = v2687
	var v2691 int32
	_ = v2691
	var v2694 int32
	_ = v2694
	var v2701 int32
	_ = v2701
	var v2706 int32
	_ = v2706
	var v2710 int32
	_ = v2710
	var v2713 int32
	_ = v2713
	var v2720 int32
	_ = v2720
	var v2725 int32
	_ = v2725
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2741 int32
	_ = v2741
	var v2745 int32
	_ = v2745
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2757 int32
	_ = v2757
	var v2761 int32
	_ = v2761
	var v2764 int32
	_ = v2764
	var v2775 int32
	_ = v2775
	var v2780 int32
	_ = v2780
	var v2784 int32
	_ = v2784
	var v2787 int32
	_ = v2787
	var v2796 int32
	_ = v2796
	var v2801 int32
	_ = v2801
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(528)
	m.G0 = v20
	if l1 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = F_palloc0(m, int32(120))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v27 = l1
	goto L3
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = int64(4294967295)
	if l3 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L4:
	;
	return
L5:
	;
	v27 = v25
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2784 = m.ExcPending
	if v2784 != 0 {
		goto L4
	} else {
		goto L938
	}
L7:
	;
	v2452 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1840))))
	v2453 = F___strchrnul(m, v1850, v2452)
	mBase = m.M
	v2455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2453))))
	if v2455 == v2452&int32(255) {
		goto L829
	} else {
		goto L830
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2433 = m.ExcPending
	if v2433 != 0 {
		goto L4
	} else {
		goto L814
	}
L9:
	;
	if l2 != 0 {
		v2451 = v2423
		goto L7
	} else {
		goto L812
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2407 = m.ExcPending
	if v2407 != 0 {
		goto L4
	} else {
		goto L808
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2391 = m.ExcPending
	if v2391 != 0 {
		goto L4
	} else {
		goto L804
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2372 = m.ExcPending
	if v2372 != 0 {
		goto L4
	} else {
		goto L800
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L4
	} else {
		goto L796
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2340 = m.ExcPending
	if v2340 != 0 {
		goto L4
	} else {
		goto L792
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2324 = m.ExcPending
	if v2324 != 0 {
		goto L4
	} else {
		goto L788
	}
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L4
	} else {
		goto L784
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L4
	} else {
		goto L780
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L4
	} else {
		goto L776
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2252 = m.ExcPending
	if v2252 != 0 {
		goto L4
	} else {
		goto L772
	}
L20:
	;
	F_errorConflictingDefElem(m, v56, l0)
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L4
	} else {
		goto L771
	}
L21:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v1741 == int32(0) {
		goto L570
	} else {
		goto L571
	}
L22:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v32 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v43 = v5
	v44 = v5
	v45 = v5
	v46 = v5
	v47 = v5
	v48 = v5
	v49 = v5
	v50 = v5
	goto L24
L24:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v52+v44<<(uint(int32(2))%32))))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	v58 = int32(_a_F_ProcessCopyOptions_0)
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[0])))
	if base.B2i32(v61 == int32(0))|base.B2i32(v61 != v64) != 0 {
		v82 = v61
		v83 = v64
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L21
L26:
	;
	v1721 = v44 + int32(1)
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v1721 < v1722 {
		v43 = v1712
		v44 = v1721
		v45 = v1713
		v46 = v1714
		v47 = v1715
		v48 = v1716
		v49 = v1717
		v50 = v1718
		goto L24
	} else {
		goto L569
	}
L27:
	;
	if v82-v83 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L28:
	;
	goto L27
L29:
	;
	v67 = v57
	v68 = v58
	goto L30
L30:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+1)))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+1)))
	if v72 == int32(0) {
		v82 = v72
		v83 = v71
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v82 = v72
	v83 = v71
	goto L28
L32:
	;
	v75 = int32(1)
	if v72 == v71 {
		v67 = v67 + v75
		v68 = v68 + v75
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v87 = F_defGetString(m, v56)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v240 = int32(_a_F_ProcessCopyOptions_1)
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[1])))
	if base.B2i32(v243 == int32(0))|base.B2i32(v243 != v246) != 0 {
		v264 = v243
		v265 = v246
		goto L85
	} else {
		goto L86
	}
L37:
	;
	if v43&int32(1) != 0 {
		goto L20
	} else {
		goto L38
	}
L38:
	;
	v91 = int32(_a_F_ProcessCopyOptions_2)
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[2])))
	if base.B2i32(v94 == int32(0))|base.B2i32(v94 != v97) != 0 {
		v115 = v94
		v116 = v97
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v115-v116 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L40:
	;
	goto L39
L41:
	;
	v100 = v87
	v101 = v91
	goto L42
L42:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	if v105 == int32(0) {
		v115 = v105
		v116 = v104
		goto L40
	} else {
		goto L44
	}
L43:
	;
	v115 = v105
	v116 = v104
	goto L40
L44:
	;
	v108 = int32(1)
	if v105 == v104 {
		v100 = v100 + v108
		v101 = v101 + v108
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(0)
	v1712 = int32(1)
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L47:
	;
	goto L48
L48:
	;
	v123 = int32(_a_F_ProcessCopyOptions_3)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v129 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[3])))
	if base.B2i32(v126 == int32(0))|base.B2i32(v126 != v129) != 0 {
		v147 = v126
		v148 = v129
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v147-v148 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L50:
	;
	goto L49
L51:
	;
	v132 = v87
	v133 = v123
	goto L52
L52:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132)+1)))
	if v137 == int32(0) {
		v147 = v137
		v148 = v136
		goto L50
	} else {
		goto L54
	}
L53:
	;
	v147 = v137
	v148 = v136
	goto L50
L54:
	;
	v140 = int32(1)
	if v137 == v136 {
		v132 = v132 + v140
		v133 = v133 + v140
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(2)
	v1712 = int32(1)
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L57:
	;
	goto L58
L58:
	;
	v155 = int32(_a_F_ProcessCopyOptions_4)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[4])))
	if base.B2i32(v158 == int32(0))|base.B2i32(v158 != v161) != 0 {
		v179 = v158
		v180 = v161
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v179-v180 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L60:
	;
	goto L59
L61:
	;
	v164 = v87
	v165 = v155
	goto L62
L62:
	;
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	if v169 == int32(0) {
		v179 = v169
		v180 = v168
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v179 = v169
	v180 = v168
	goto L60
L64:
	;
	v172 = int32(1)
	if v169 == v168 {
		v164 = v164 + v172
		v165 = v165 + v172
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v184 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v184
	v1712 = v184
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L67:
	;
	goto L68
L68:
	;
	v187 = int32(_a_F_ProcessCopyOptions_5)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[5])))
	if base.B2i32(v190 == int32(0))|base.B2i32(v190 != v193) != 0 {
		v211 = v190
		v212 = v193
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v211-v212 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L70:
	;
	goto L69
L71:
	;
	v196 = v87
	v197 = v187
	goto L72
L72:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+1)))
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196)+1)))
	if v201 == int32(0) {
		v211 = v201
		v212 = v200
		goto L70
	} else {
		goto L74
	}
L73:
	;
	v211 = v201
	v212 = v200
	goto L70
L74:
	;
	v204 = int32(1)
	if v201 == v200 {
		v196 = v196 + v204
		v197 = v197 + v204
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = int32(3)
	v1712 = int32(1)
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L77:
	;
	goto L78
L78:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+352)) = v87
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_6), v20+int32(352))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(627), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	if v264-v265 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	goto L84
L86:
	;
	v249 = v57
	v250 = v240
	goto L87
L87:
	;
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250)+1)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
	if v254 == int32(0) {
		v264 = v254
		v265 = v253
		goto L85
	} else {
		goto L89
	}
L88:
	;
	v264 = v254
	v265 = v253
	goto L85
L89:
	;
	v257 = int32(1)
	if v254 == v253 {
		v249 = v249 + v257
		v250 = v250 + v257
		goto L87
	} else {
		goto L90
	}
L90:
	;
	goto L88
L91:
	;
	if v45 != 0 {
		goto L20
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v273 = int32(_a_F_ProcessCopyOptions_9)
	v276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v279 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[6])))
	if base.B2i32(v276 == int32(0))|base.B2i32(v276 != v279) != 0 {
		v297 = v276
		v298 = v279
		goto L97
	} else {
		goto L98
	}
L94:
	;
	v269 = F_defGetBoolean(m, v56)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)) = uint8(v269)
	v1712 = v43
	v1713 = int32(1)
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L96:
	;
	if v297-v298 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L97:
	;
	goto L96
L98:
	;
	v282 = v57
	v283 = v273
	goto L99
L99:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282)+1)))
	if v287 == int32(0) {
		v297 = v287
		v298 = v286
		goto L97
	} else {
		goto L101
	}
L100:
	;
	v297 = v287
	v298 = v286
	goto L97
L101:
	;
	v290 = int32(1)
	if v287 == v286 {
		v282 = v282 + v290
		v283 = v283 + v290
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	if v302 != 0 {
		goto L20
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v306 = int32(_a_F_ProcessCopyOptions_10)
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[7])))
	if base.B2i32(v309 == int32(0))|base.B2i32(v309 != v312) != 0 {
		v330 = v309
		v331 = v312
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v303 = F_defGetString(m, v56)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L4
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v303
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L108:
	;
	if v330-v331 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	goto L108
L110:
	;
	v315 = v57
	v316 = v306
	goto L111
L111:
	;
	v319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v315)+1)))
	if v320 == int32(0) {
		v330 = v320
		v331 = v319
		goto L109
	} else {
		goto L113
	}
L112:
	;
	v330 = v320
	v331 = v319
	goto L109
L113:
	;
	v323 = int32(1)
	if v320 == v319 {
		v315 = v315 + v323
		v316 = v316 + v323
		goto L111
	} else {
		goto L114
	}
L114:
	;
	goto L112
L115:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v335 != 0 {
		goto L20
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v339 = int32(_a_F_ProcessCopyOptions_11)
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v345 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[8])))
	if base.B2i32(v342 == int32(0))|base.B2i32(v342 != v345) != 0 {
		v363 = v342
		v364 = v345
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v336 = F_defGetString(m, v56)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L4
	} else {
		goto L119
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v336
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L120:
	;
	if v363-v364 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L121:
	;
	goto L120
L122:
	;
	v348 = v57
	v349 = v339
	goto L123
L123:
	;
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349)+1)))
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v348)+1)))
	if v353 == int32(0) {
		v363 = v353
		v364 = v352
		goto L121
	} else {
		goto L125
	}
L124:
	;
	v363 = v353
	v364 = v352
	goto L121
L125:
	;
	v356 = int32(1)
	if v353 == v352 {
		v348 = v348 + v356
		v349 = v349 + v356
		goto L123
	} else {
		goto L126
	}
L126:
	;
	goto L124
L127:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v368 != 0 {
		goto L20
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v372 = int32(_a_F_ProcessCopyOptions_12)
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[9])))
	if base.B2i32(v375 == int32(0))|base.B2i32(v375 != v378) != 0 {
		v396 = v375
		v397 = v378
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v369 = F_defGetString(m, v56)
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L4
	} else {
		goto L131
	}
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+28)) = v369
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L132:
	;
	if v396-v397 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L133:
	;
	goto L132
L134:
	;
	v381 = v57
	v382 = v372
	goto L135
L135:
	;
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v382)+1)))
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v381)+1)))
	if v386 == int32(0) {
		v396 = v386
		v397 = v385
		goto L133
	} else {
		goto L137
	}
L136:
	;
	v396 = v386
	v397 = v385
	goto L133
L137:
	;
	v389 = int32(1)
	if v386 == v385 {
		v381 = v381 + v389
		v382 = v382 + v389
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	if v46 != 0 {
		goto L20
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	v676 = int32(_a_F_ProcessCopyOptions_13)
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v682 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[10])))
	if base.B2i32(v679 == int32(0))|base.B2i32(v679 != v682) != 0 {
		v700 = v679
		v701 = v682
		goto L233
	} else {
		goto L234
	}
L142:
	;
	v401 = int32(1)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v403 == int32(0) {
		v673 = v401
		goto L143
	} else {
		goto L144
	}
L143:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+12)) = v673
	v1712 = v43
	v1713 = v45
	v1714 = v401
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L144:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	if v406 == int32(473) {
		goto L146
	} else {
		goto L147
	}
L145:
	;
	if v667 < int32(0) {
		goto L18
	} else {
		goto L229
	}
L146:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v403)+4))
	v667 = v409
	goto L145
L147:
	;
	goto L148
L148:
	;
	v410 = F_defGetString(m, v56)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	v415 = v410
	v416 = int32(_a_F_ProcessCopyOptions_14)
	goto L151
L150:
	;
	if v453 == int32(0) {
		v673 = v401
		goto L143
	} else {
		goto L163
	}
L151:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415))))
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416))))
	if v419 == v420 {
		v442 = v419
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v453 = int32(0)
	goto L150
L153:
	;
	v444 = int32(1)
	if v442 != 0 {
		v415 = v415 + v444
		v416 = v416 + v444
		goto L151
	} else {
		goto L162
	}
L154:
	;
	if base.Ui32((v419-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v430 = v419 | int32(32)
	goto L157
L156:
	;
	v430 = v419
	goto L157
L157:
	;
	if base.Ui32((v420-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v439 = v420 | int32(32)
	goto L160
L159:
	;
	v439 = v420
	goto L160
L160:
	;
	if v430 == v439 {
		v442 = v430
		goto L153
	} else {
		goto L161
	}
L161:
	;
	v453 = v430 - v439
	goto L150
L162:
	;
	goto L152
L163:
	;
	v460 = v410
	v461 = int32(_a_F_ProcessCopyOptions_15)
	goto L165
L164:
	;
	if v498 == int32(0) {
		v673 = int32(0)
		goto L143
	} else {
		goto L177
	}
L165:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	v465 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461))))
	if v464 == v465 {
		v487 = v464
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v498 = int32(0)
	goto L164
L167:
	;
	v489 = int32(1)
	if v487 != 0 {
		v460 = v460 + v489
		v461 = v461 + v489
		goto L165
	} else {
		goto L176
	}
L168:
	;
	if base.Ui32((v464-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v475 = v464 | int32(32)
	goto L171
L170:
	;
	v475 = v464
	goto L171
L171:
	;
	if base.Ui32((v465-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v484 = v465 | int32(32)
	goto L174
L173:
	;
	v484 = v465
	goto L174
L174:
	;
	if v475 == v484 {
		v487 = v475
		goto L167
	} else {
		goto L175
	}
L175:
	;
	v498 = v475 - v484
	goto L164
L176:
	;
	goto L166
L177:
	;
	v505 = v410
	v506 = int32(_a_F_ProcessCopyOptions_16)
	goto L179
L178:
	;
	if v543 == int32(0) {
		v673 = int32(1)
		goto L143
	} else {
		goto L191
	}
L179:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v505))))
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v506))))
	if v509 == v510 {
		v532 = v509
		goto L181
	} else {
		goto L182
	}
L180:
	;
	v543 = int32(0)
	goto L178
L181:
	;
	v534 = int32(1)
	if v532 != 0 {
		v505 = v505 + v534
		v506 = v506 + v534
		goto L179
	} else {
		goto L190
	}
L182:
	;
	if base.Ui32((v509-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v520 = v509 | int32(32)
	goto L185
L184:
	;
	v520 = v509
	goto L185
L185:
	;
	if base.Ui32((v510-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L186
	} else {
		goto L187
	}
L186:
	;
	v529 = v510 | int32(32)
	goto L188
L187:
	;
	v529 = v510
	goto L188
L188:
	;
	if v520 == v529 {
		v532 = v520
		goto L181
	} else {
		goto L189
	}
L189:
	;
	v543 = v520 - v529
	goto L178
L190:
	;
	goto L180
L191:
	;
	v550 = v410
	v551 = int32(_a_F_ProcessCopyOptions_17)
	goto L193
L192:
	;
	if v588 == int32(0) {
		v673 = int32(0)
		goto L143
	} else {
		goto L205
	}
L193:
	;
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v550))))
	v555 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551))))
	if v554 == v555 {
		v577 = v554
		goto L195
	} else {
		goto L196
	}
L194:
	;
	v588 = int32(0)
	goto L192
L195:
	;
	v579 = int32(1)
	if v577 != 0 {
		v550 = v550 + v579
		v551 = v551 + v579
		goto L193
	} else {
		goto L204
	}
L196:
	;
	if base.Ui32((v554-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	v565 = v554 | int32(32)
	goto L199
L198:
	;
	v565 = v554
	goto L199
L199:
	;
	if base.Ui32((v555-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v574 = v555 | int32(32)
	goto L202
L201:
	;
	v574 = v555
	goto L202
L202:
	;
	if v565 == v574 {
		v577 = v565
		goto L195
	} else {
		goto L203
	}
L203:
	;
	v588 = v565 - v574
	goto L192
L204:
	;
	goto L194
L205:
	;
	v594 = v410
	v595 = int32(_a_F_ProcessCopyOptions_18)
	goto L207
L206:
	;
	if v632 == int32(0) {
		goto L219
	} else {
		goto L220
	}
L207:
	;
	v598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v594))))
	v599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v595))))
	if v598 == v599 {
		v621 = v598
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v632 = int32(0)
	goto L206
L209:
	;
	v623 = int32(1)
	if v621 != 0 {
		v594 = v594 + v623
		v595 = v595 + v623
		goto L207
	} else {
		goto L218
	}
L210:
	;
	if base.Ui32((v598-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v609 = v598 | int32(32)
	goto L213
L212:
	;
	v609 = v598
	goto L213
L213:
	;
	if base.Ui32((v599-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v618 = v599 | int32(32)
	goto L216
L215:
	;
	v618 = v599
	goto L216
L216:
	;
	if v609 == v618 {
		v621 = v609
		goto L209
	} else {
		goto L217
	}
L217:
	;
	v632 = v609 - v618
	goto L206
L218:
	;
	goto L208
L219:
	;
	if l2 != 0 {
		v673 = int32(-1)
		goto L143
	} else {
		goto L222
	}
L220:
	;
	goto L221
L221:
	;
	v655 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[11]))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+520)) = v655
	v658 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[12]))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+512)) = v658
	v662 = F_pg_strtoint32_safe(m, v410, v20+int32(512))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L4
	} else {
		goto L227
	}
L222:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v639 = m.ExcPending
	if v639 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L4
	} else {
		goto L224
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+384)) = v410
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_19), v20+int32(384))
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L4
	} else {
		goto L225
	}
L225:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(438), int32(_a_F_ProcessCopyOptions_20))
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L4
	} else {
		goto L226
	}
L226:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L227:
	;
	v664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+516)))
	if v664 == int32(1) {
		goto L19
	} else {
		goto L228
	}
L228:
	;
	v667 = v662
	goto L145
L229:
	;
	if l2 != 0 {
		v673 = v667
		goto L143
	} else {
		goto L230
	}
L230:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v667) {
		goto L17
	} else {
		goto L231
	}
L231:
	;
	v673 = v667
	goto L143
L232:
	;
	if v700-v701 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L233:
	;
	goto L232
L234:
	;
	v685 = v57
	v686 = v676
	goto L235
L235:
	;
	v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v686)+1)))
	v690 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v685)+1)))
	if v690 == int32(0) {
		v700 = v690
		v701 = v689
		goto L233
	} else {
		goto L237
	}
L236:
	;
	v700 = v690
	v701 = v689
	goto L233
L237:
	;
	v693 = int32(1)
	if v690 == v689 {
		v685 = v685 + v693
		v686 = v686 + v693
		goto L235
	} else {
		goto L238
	}
L238:
	;
	goto L236
L239:
	;
	v705 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v705 != 0 {
		goto L20
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v709 = int32(_a_F_ProcessCopyOptions_21)
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v715 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[13])))
	if base.B2i32(v712 == int32(0))|base.B2i32(v712 != v715) != 0 {
		v733 = v712
		v734 = v715
		goto L245
	} else {
		goto L246
	}
L242:
	;
	v706 = F_defGetString(m, v56)
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L4
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v706
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L244:
	;
	if v733-v734 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L245:
	;
	goto L244
L246:
	;
	v718 = v57
	v719 = v709
	goto L247
L247:
	;
	v722 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+1)))
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718)+1)))
	if v723 == int32(0) {
		v733 = v723
		v734 = v722
		goto L245
	} else {
		goto L249
	}
L248:
	;
	v733 = v723
	v734 = v722
	goto L245
L249:
	;
	v726 = int32(1)
	if v723 == v722 {
		v718 = v718 + v726
		v719 = v719 + v726
		goto L247
	} else {
		goto L250
	}
L250:
	;
	goto L248
L251:
	;
	v738 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v738 != 0 {
		goto L20
	} else {
		goto L254
	}
L252:
	;
	goto L253
L253:
	;
	v742 = int32(_a_F_ProcessCopyOptions_22)
	v745 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v748 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[14])))
	if base.B2i32(v745 == int32(0))|base.B2i32(v745 != v748) != 0 {
		v766 = v745
		v767 = v748
		goto L258
	} else {
		goto L259
	}
L254:
	;
	v739 = F_defGetString(m, v56)
	mBase = m.M
	v740 = m.ExcPending
	if v740 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v739
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+48)) = v775
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L257:
	;
	if v766-v767 == int32(0) {
		goto L264
	} else {
		goto L265
	}
L258:
	;
	goto L257
L259:
	;
	v751 = v57
	v752 = v742
	goto L260
L260:
	;
	v755 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v752)+1)))
	v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751)+1)))
	if v756 == int32(0) {
		v766 = v756
		v767 = v755
		goto L258
	} else {
		goto L262
	}
L261:
	;
	v766 = v756
	v767 = v755
	goto L258
L262:
	;
	v759 = int32(1)
	if v756 == v755 {
		v751 = v751 + v759
		v752 = v752 + v759
		goto L260
	} else {
		goto L263
	}
L263:
	;
	goto L261
L264:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v771 != 0 {
		goto L20
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v808 = int32(_a_F_ProcessCopyOptions_23)
	v811 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v814 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[15])))
	if base.B2i32(v811 == int32(0))|base.B2i32(v811 != v814) != 0 {
		v832 = v811
		v833 = v814
		goto L279
	} else {
		goto L280
	}
L267:
	;
	v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+52)))
	if v772 == int32(1) {
		goto L20
	} else {
		goto L268
	}
L268:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v775 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L4
	} else {
		goto L273
	}
L270:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v775)))
	if v778 == int32(1) {
		goto L256
	} else {
		goto L271
	}
L271:
	;
	if v778 != int32(77) {
		goto L269
	} else {
		goto L272
	}
L272:
	;
	v783 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+52)) = uint8(v783)
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L273:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L4
	} else {
		goto L274
	}
L274:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+416)) = v793
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_24), v20+int32(416))
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L4
	} else {
		goto L275
	}
L275:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v800)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L276
	}
L276:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(686), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L4
	} else {
		goto L277
	}
L277:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L278:
	;
	if v832-v833 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L279:
	;
	goto L278
L280:
	;
	v817 = v57
	v818 = v808
	goto L281
L281:
	;
	v821 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v818)+1)))
	v822 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v817)+1)))
	if v822 == int32(0) {
		v832 = v822
		v833 = v821
		goto L279
	} else {
		goto L283
	}
L282:
	;
	v832 = v822
	v833 = v821
	goto L279
L283:
	;
	v825 = int32(1)
	if v822 == v821 {
		v817 = v817 + v825
		v818 = v818 + v825
		goto L281
	} else {
		goto L284
	}
L284:
	;
	goto L282
L285:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v837 != 0 {
		goto L20
	} else {
		goto L288
	}
L286:
	;
	goto L287
L287:
	;
	v875 = int32(_a_F_ProcessCopyOptions_25)
	v878 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v881 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[16])))
	if base.B2i32(v878 == int32(0))|base.B2i32(v878 != v881) != 0 {
		v899 = v878
		v900 = v881
		goto L302
	} else {
		goto L303
	}
L288:
	;
	v838 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+64)))
	if v838 == int32(1) {
		goto L20
	} else {
		goto L289
	}
L289:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v841 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L4
	} else {
		goto L296
	}
L291:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v841)))
	if v844 != int32(1) {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	if v844 != int32(77) {
		goto L290
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+60)) = v841
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L295:
	;
	v849 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+64)) = uint8(v849)
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L296:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L4
	} else {
		goto L297
	}
L297:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+432)) = v860
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_24), v20+int32(432))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L4
	} else {
		goto L298
	}
L298:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v867)
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L4
	} else {
		goto L299
	}
L299:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(701), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v874 = m.ExcPending
	if v874 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L301:
	;
	if v899-v900 == int32(0) {
		goto L308
	} else {
		goto L309
	}
L302:
	;
	goto L301
L303:
	;
	v884 = v57
	v885 = v875
	goto L304
L304:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v885)+1)))
	v889 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v884)+1)))
	if v889 == int32(0) {
		v899 = v889
		v900 = v888
		goto L302
	} else {
		goto L306
	}
L305:
	;
	v899 = v889
	v900 = v888
	goto L302
L306:
	;
	v892 = int32(1)
	if v889 == v888 {
		v884 = v884 + v892
		v885 = v885 + v892
		goto L304
	} else {
		goto L307
	}
L307:
	;
	goto L305
L308:
	;
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v904 != 0 {
		goto L20
	} else {
		goto L311
	}
L309:
	;
	goto L310
L310:
	;
	v942 = int32(_a_F_ProcessCopyOptions_26)
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v948 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[17])))
	if base.B2i32(v945 == int32(0))|base.B2i32(v945 != v948) != 0 {
		v966 = v945
		v967 = v948
		goto L325
	} else {
		goto L326
	}
L311:
	;
	v905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)))
	if v905 == int32(1) {
		goto L20
	} else {
		goto L312
	}
L312:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v908 == int32(0) {
		goto L313
	} else {
		goto L314
	}
L313:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L4
	} else {
		goto L319
	}
L314:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v908)))
	if v911 != int32(1) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	if v911 != int32(77) {
		goto L313
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+76)) = v908
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L318:
	;
	v916 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)) = uint8(v916)
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L319:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v926 = m.ExcPending
	if v926 != 0 {
		goto L4
	} else {
		goto L320
	}
L320:
	;
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+448)) = v927
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_24), v20+int32(448))
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L4
	} else {
		goto L321
	}
L321:
	;
	v934 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v934)
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L4
	} else {
		goto L322
	}
L322:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(716), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L4
	} else {
		goto L323
	}
L323:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L324:
	;
	if v966-v967 == int32(0) {
		goto L331
	} else {
		goto L332
	}
L325:
	;
	goto L324
L326:
	;
	v951 = v57
	v952 = v942
	goto L327
L327:
	;
	v955 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v952)+1)))
	v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v951)+1)))
	if v956 == int32(0) {
		v966 = v956
		v967 = v955
		goto L325
	} else {
		goto L329
	}
L328:
	;
	v966 = v956
	v967 = v955
	goto L325
L329:
	;
	v959 = int32(1)
	if v956 == v955 {
		v951 = v951 + v959
		v952 = v952 + v959
		goto L327
	} else {
		goto L330
	}
L330:
	;
	goto L328
L331:
	;
	v971 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+88)))
	if v971 == int32(1) {
		goto L20
	} else {
		goto L334
	}
L332:
	;
	goto L333
L333:
	;
	v1003 = int32(_a_F_ProcessCopyOptions_27)
	v1006 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1009 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[18])))
	if base.B2i32(v1006 == int32(0))|base.B2i32(v1006 != v1009) != 0 {
		v1027 = v1006
		v1028 = v1009
		goto L346
	} else {
		goto L347
	}
L334:
	;
	v974 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+88)) = uint8(v974)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v976 != 0 {
		goto L336
	} else {
		goto L337
	}
L335:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v984 = m.ExcPending
	if v984 != 0 {
		goto L4
	} else {
		goto L340
	}
L336:
	;
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v976)))
	if v977 != int32(1) {
		goto L335
	} else {
		goto L339
	}
L337:
	;
	goto L338
L338:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+112)) = v976
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L339:
	;
	goto L338
L340:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L4
	} else {
		goto L341
	}
L341:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+464)) = v988
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_24), v20+int32(464))
	mBase = m.M
	v994 = m.ExcPending
	if v994 != 0 {
		goto L4
	} else {
		goto L342
	}
L342:
	;
	v995 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v995)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L4
	} else {
		goto L343
	}
L343:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(735), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v1002 = m.ExcPending
	if v1002 != 0 {
		goto L4
	} else {
		goto L344
	}
L344:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L345:
	;
	if v1027-v1028 == int32(0) {
		goto L352
	} else {
		goto L353
	}
L346:
	;
	goto L345
L347:
	;
	v1012 = v57
	v1013 = v1003
	goto L348
L348:
	;
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1013)+1)))
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012)+1)))
	if v1017 == int32(0) {
		v1027 = v1017
		v1028 = v1016
		goto L346
	} else {
		goto L350
	}
L349:
	;
	v1027 = v1017
	v1028 = v1016
	goto L346
L350:
	;
	v1020 = int32(1)
	if v1017 == v1016 {
		v1012 = v1012 + v1020
		v1013 = v1013 + v1020
		goto L348
	} else {
		goto L351
	}
L351:
	;
	goto L349
L352:
	;
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if int32(0) <= v1032 {
		goto L20
	} else {
		goto L355
	}
L353:
	;
	goto L354
L354:
	;
	v1154 = int32(_a_F_ProcessCopyOptions_28)
	v1157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1160 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[19])))
	if base.B2i32(v1157 == int32(0))|base.B2i32(v1157 != v1160) != 0 {
		v1178 = v1157
		v1179 = v1160
		goto L390
	} else {
		goto L391
	}
L355:
	;
	v1035 = F_defGetString(m, v56)
	mBase = m.M
	v1036 = m.ExcPending
	if v1036 != 0 {
		goto L4
	} else {
		goto L356
	}
L356:
	;
	v1044 = m.G0
	v1046 = v1044 + int32(-64)
	m.G0 = v1046
	v1048 = int32(-1)
	if v1035 == int32(0) {
		v1123 = v1048
		goto L358
	} else {
		goto L359
	}
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v1123
	if int32(0) <= v1123 {
		v1712 = v43
		v1713 = v45
		v1714 = v46
		v1715 = v47
		v1716 = v48
		v1717 = v49
		v1718 = v50
		goto L26
	} else {
		goto L383
	}
L358:
	;
	m.G0 = v1046 - int32(-64)
	goto L357
L359:
	;
	v1051 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035))))
	if v1051 == int32(0) {
		v1123 = v1048
		goto L358
	} else {
		goto L360
	}
L360:
	;
	v1054 = F_strlen(m, v1035)
	mBase = m.M
	if base.Ui32(int32(63)) < base.Ui32(v1054) {
		v1123 = v1048
		goto L358
	} else {
		goto L361
	}
L361:
	;
	v1057 = v1035
	v1058 = v1051
	v1059 = v1046
	goto L362
L362:
	;
	v1067 = F_isalnum(m, v1058&int32(255))
	mBase = m.M
	if v1067 != 0 {
		goto L364
	} else {
		goto L365
	}
L363:
	;
	v1084 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1080))) = uint8(v1084)
	v1088 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1046))))
	v1089 = int32(_a_F_ProcessCopyOptions_29)
	v1090 = int32(_a_F_ProcessCopyOptions_30)
	goto L371
L364:
	;
	if base.Ui32((v1058-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L367
	} else {
		goto L368
	}
L365:
	;
	v1080 = v1059
	goto L366
L366:
	;
	v1081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1057)+1)))
	if v1081 != 0 {
		v1057 = v1057 + int32(1)
		v1058 = v1081
		v1059 = v1080
		goto L362
	} else {
		goto L370
	}
L367:
	;
	v1076 = v1058 | int32(32)
	goto L369
L368:
	;
	v1076 = v1058
	goto L369
L369:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1059))) = uint8(v1076)
	v1080 = v1059 + int32(1)
	goto L366
L370:
	;
	goto L363
L371:
	;
	v1102 = v1090 + (v1089-v1090)>>(uint(int32(4))%32)<<(uint(int32(3))%32)
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1102)))
	v1104 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1103))))
	v1105 = v1088 - v1104
	if v1105 != 0 {
		v1108 = v1105
		goto L373
	} else {
		goto L374
	}
L372:
	;
	v1123 = v1048
	goto L358
L373:
	;
	v1112 = base.B2i32(v1108 < int32(0))
	if v1108 < int32(0) {
		goto L376
	} else {
		goto L377
	}
L374:
	;
	v1106 = F_strcmp(m, v1046, v1103)
	mBase = m.M
	if v1106 != 0 {
		v1108 = v1106
		goto L373
	} else {
		goto L375
	}
L375:
	;
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+4))
	v1123 = v1107
	goto L358
L376:
	;
	v1113 = v1102 - int32(8)
	goto L378
L377:
	;
	v1113 = v1089
	goto L378
L378:
	;
	if v1108 < int32(0) {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	v1116 = v1090
	goto L381
L380:
	;
	v1116 = v1102 + int32(8)
	goto L381
L381:
	;
	if base.Ui32(v1116) <= base.Ui32(v1113) {
		v1089 = v1113
		v1090 = v1116
		goto L371
	} else {
		goto L382
	}
L382:
	;
	goto L372
L383:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L4
	} else {
		goto L384
	}
L384:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1138 = m.ExcPending
	if v1138 != 0 {
		goto L4
	} else {
		goto L385
	}
L385:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+480)) = v1139
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_31), v20+int32(480))
	mBase = m.M
	v1145 = m.ExcPending
	if v1145 != 0 {
		goto L4
	} else {
		goto L386
	}
L386:
	;
	v1146 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1146)
	mBase = m.M
	v1148 = m.ExcPending
	if v1148 != 0 {
		goto L4
	} else {
		goto L387
	}
L387:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(747), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L4
	} else {
		goto L388
	}
L388:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L389:
	;
	if v1178-v1179 == int32(0) {
		goto L396
	} else {
		goto L397
	}
L390:
	;
	goto L389
L391:
	;
	v1163 = v57
	v1164 = v1154
	goto L392
L392:
	;
	v1167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1164)+1)))
	v1168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+1)))
	if v1168 == int32(0) {
		v1178 = v1168
		v1179 = v1167
		goto L390
	} else {
		goto L394
	}
L393:
	;
	v1178 = v1168
	v1179 = v1167
	goto L390
L394:
	;
	v1171 = int32(1)
	if v1168 == v1167 {
		v1163 = v1163 + v1171
		v1164 = v1164 + v1171
		goto L392
	} else {
		goto L395
	}
L395:
	;
	goto L393
L396:
	;
	if v47 != 0 {
		goto L20
	} else {
		goto L399
	}
L397:
	;
	goto L398
L398:
	;
	v1187 = int32(_a_F_ProcessCopyOptions_32)
	v1190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1193 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[20])))
	if base.B2i32(v1190 == int32(0))|base.B2i32(v1190 != v1193) != 0 {
		v1211 = v1190
		v1212 = v1193
		goto L402
	} else {
		goto L403
	}
L399:
	;
	v1183 = F_defGetBoolean(m, v56)
	mBase = m.M
	v1184 = m.ExcPending
	if v1184 != 0 {
		goto L4
	} else {
		goto L400
	}
L400:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)) = uint8(v1183)
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = int32(1)
	v1716 = v48
	v1717 = v49
	v1718 = v50
	goto L26
L401:
	;
	if v1211-v1212 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L402:
	;
	goto L401
L403:
	;
	v1196 = v57
	v1197 = v1187
	goto L404
L404:
	;
	v1200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1197)+1)))
	v1201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1196)+1)))
	if v1201 == int32(0) {
		v1211 = v1201
		v1212 = v1200
		goto L402
	} else {
		goto L406
	}
L405:
	;
	v1211 = v1201
	v1212 = v1200
	goto L402
L406:
	;
	v1204 = int32(1)
	if v1201 == v1200 {
		v1196 = v1196 + v1204
		v1197 = v1197 + v1204
		goto L404
	} else {
		goto L407
	}
L407:
	;
	goto L405
L408:
	;
	if v48 != 0 {
		goto L20
	} else {
		goto L411
	}
L409:
	;
	goto L410
L410:
	;
	v1406 = int32(_a_F_ProcessCopyOptions_33)
	v1409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[21])))
	if base.B2i32(v1409 == int32(0))|base.B2i32(v1409 != v1412) != 0 {
		v1430 = v1409
		v1431 = v1412
		goto L472
	} else {
		goto L473
	}
L411:
	;
	v1216 = m.G0
	v1218 = v1216 - int32(32)
	m.G0 = v1218
	v1220 = F_defGetString(m, v56)
	mBase = m.M
	v1221 = m.ExcPending
	if v1221 != 0 {
		goto L4
	} else {
		goto L413
	}
L412:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v1355
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = int32(1)
	v1717 = v49
	v1718 = v50
	goto L26
L413:
	;
	if l2 != 0 {
		goto L415
	} else {
		goto L416
	}
L414:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L4
	} else {
		goto L466
	}
L415:
	;
	v1226 = v1220
	v1227 = int32(_a_F_ProcessCopyOptions_34)
	goto L420
L416:
	;
	goto L417
L417:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1362 = m.ExcPending
	if v1362 != 0 {
		goto L4
	} else {
		goto L461
	}
L418:
	;
	m.G0 = v1218 + int32(32)
	goto L412
L419:
	;
	if v1264 == int32(0) {
		v1355 = int32(0)
		goto L418
	} else {
		goto L432
	}
L420:
	;
	v1230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226))))
	v1231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227))))
	if v1230 == v1231 {
		v1253 = v1230
		goto L422
	} else {
		goto L423
	}
L421:
	;
	v1264 = int32(0)
	goto L419
L422:
	;
	v1255 = int32(1)
	if v1253 != 0 {
		v1226 = v1226 + v1255
		v1227 = v1227 + v1255
		goto L420
	} else {
		goto L431
	}
L423:
	;
	if base.Ui32((v1230-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1241 = v1230 | int32(32)
	goto L426
L425:
	;
	v1241 = v1230
	goto L426
L426:
	;
	if base.Ui32((v1231-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1250 = v1231 | int32(32)
	goto L429
L428:
	;
	v1250 = v1231
	goto L429
L429:
	;
	if v1241 == v1250 {
		v1253 = v1241
		goto L422
	} else {
		goto L430
	}
L430:
	;
	v1264 = v1241 - v1250
	goto L419
L431:
	;
	goto L421
L432:
	;
	v1271 = v1220
	v1272 = int32(_a_F_ProcessCopyOptions_35)
	goto L434
L433:
	;
	if v1309 == int32(0) {
		v1355 = int32(1)
		goto L418
	} else {
		goto L446
	}
L434:
	;
	v1275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1271))))
	v1276 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272))))
	if v1275 == v1276 {
		v1298 = v1275
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v1309 = int32(0)
	goto L433
L436:
	;
	v1300 = int32(1)
	if v1298 != 0 {
		v1271 = v1271 + v1300
		v1272 = v1272 + v1300
		goto L434
	} else {
		goto L445
	}
L437:
	;
	if base.Ui32((v1275-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L438
	} else {
		goto L439
	}
L438:
	;
	v1286 = v1275 | int32(32)
	goto L440
L439:
	;
	v1286 = v1275
	goto L440
L440:
	;
	if base.Ui32((v1276-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	v1295 = v1276 | int32(32)
	goto L443
L442:
	;
	v1295 = v1276
	goto L443
L443:
	;
	if v1286 == v1295 {
		v1298 = v1286
		goto L436
	} else {
		goto L444
	}
L444:
	;
	v1309 = v1286 - v1295
	goto L433
L445:
	;
	goto L435
L446:
	;
	v1315 = v1220
	v1316 = int32(_a_F_ProcessCopyOptions_36)
	goto L448
L447:
	;
	if v1353 != 0 {
		goto L414
	} else {
		goto L460
	}
L448:
	;
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1315))))
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1316))))
	if v1319 == v1320 {
		v1342 = v1319
		goto L450
	} else {
		goto L451
	}
L449:
	;
	v1353 = int32(0)
	goto L447
L450:
	;
	v1344 = int32(1)
	if v1342 != 0 {
		v1315 = v1315 + v1344
		v1316 = v1316 + v1344
		goto L448
	} else {
		goto L459
	}
L451:
	;
	if base.Ui32((v1319-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L452
	} else {
		goto L453
	}
L452:
	;
	v1330 = v1319 | int32(32)
	goto L454
L453:
	;
	v1330 = v1319
	goto L454
L454:
	;
	if base.Ui32((v1320-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	v1339 = v1320 | int32(32)
	goto L457
L456:
	;
	v1339 = v1320
	goto L457
L457:
	;
	if v1330 == v1339 {
		v1342 = v1330
		goto L450
	} else {
		goto L458
	}
L458:
	;
	v1353 = v1330 - v1339
	goto L447
L459:
	;
	goto L449
L460:
	;
	v1355 = int32(2)
	goto L418
L461:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1365 = m.ExcPending
	if v1365 != 0 {
		goto L4
	} else {
		goto L462
	}
L462:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+20)) = int32(_a_F_ProcessCopyOptions_37)
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+16)) = int32(_a_F_ProcessCopyOptions_38)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_39), v1218+int32(16))
	mBase = m.M
	v1374 = m.ExcPending
	if v1374 != 0 {
		goto L4
	} else {
		goto L463
	}
L463:
	;
	v1375 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1375)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L4
	} else {
		goto L464
	}
L464:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(489), int32(_a_F_ProcessCopyOptions_40))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L4
	} else {
		goto L465
	}
L465:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L466:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L4
	} else {
		goto L467
	}
L467:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1218)+4)) = v1220
	*(*int32)(unsafe.Add(mBase, uint32(v1218))) = int32(_a_F_ProcessCopyOptions_38)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_41), v1218)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L4
	} else {
		goto L468
	}
L468:
	;
	v1396 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1396)
	mBase = m.M
	v1398 = m.ExcPending
	if v1398 != 0 {
		goto L4
	} else {
		goto L469
	}
L469:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(502), int32(_a_F_ProcessCopyOptions_40))
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L4
	} else {
		goto L470
	}
L470:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L471:
	;
	if v1430-v1431 == int32(0) {
		goto L478
	} else {
		goto L479
	}
L472:
	;
	goto L471
L473:
	;
	v1415 = v57
	v1416 = v1406
	goto L474
L474:
	;
	v1419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1416)+1)))
	v1420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1415)+1)))
	if v1420 == int32(0) {
		v1430 = v1420
		v1431 = v1419
		goto L472
	} else {
		goto L476
	}
L475:
	;
	v1430 = v1420
	v1431 = v1419
	goto L472
L476:
	;
	v1423 = int32(1)
	if v1420 == v1419 {
		v1415 = v1415 + v1423
		v1416 = v1416 + v1423
		goto L474
	} else {
		goto L477
	}
L477:
	;
	goto L475
L478:
	;
	if v49 != 0 {
		goto L20
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	v1601 = int32(_a_F_ProcessCopyOptions_42)
	v1604 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57))))
	v1607 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessCopyOptions[22])))
	if base.B2i32(v1604 == int32(0))|base.B2i32(v1604 != v1607) != 0 {
		v1625 = v1604
		v1626 = v1607
		goto L534
	} else {
		goto L535
	}
L481:
	;
	v1435 = m.G0
	v1437 = v1435 - int32(16)
	m.G0 = v1437
	v1440 = F_defGetString(m, v56)
	mBase = m.M
	v1441 = m.ExcPending
	if v1441 != 0 {
		goto L4
	} else {
		goto L485
	}
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+96)) = v1574
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = int32(1)
	v1718 = v50
	goto L26
L483:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1581 = m.ExcPending
	if v1581 != 0 {
		goto L4
	} else {
		goto L528
	}
L484:
	;
	m.G0 = v1437 + int32(16)
	goto L482
L485:
	;
	v1445 = v1440
	v1446 = int32(_a_F_ProcessCopyOptions_43)
	goto L487
L486:
	;
	if v1483 == int32(0) {
		v1574 = int32(-1)
		goto L484
	} else {
		goto L499
	}
L487:
	;
	v1449 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1445))))
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446))))
	if v1449 == v1450 {
		v1472 = v1449
		goto L489
	} else {
		goto L490
	}
L488:
	;
	v1483 = int32(0)
	goto L486
L489:
	;
	v1474 = int32(1)
	if v1472 != 0 {
		v1445 = v1445 + v1474
		v1446 = v1446 + v1474
		goto L487
	} else {
		goto L498
	}
L490:
	;
	if base.Ui32((v1449-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L491
	} else {
		goto L492
	}
L491:
	;
	v1460 = v1449 | int32(32)
	goto L493
L492:
	;
	v1460 = v1449
	goto L493
L493:
	;
	if base.Ui32((v1450-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	v1469 = v1450 | int32(32)
	goto L496
L495:
	;
	v1469 = v1450
	goto L496
L496:
	;
	if v1460 == v1469 {
		v1472 = v1460
		goto L489
	} else {
		goto L497
	}
L497:
	;
	v1483 = v1460 - v1469
	goto L486
L498:
	;
	goto L488
L499:
	;
	v1490 = v1440
	v1491 = int32(_a_F_ProcessCopyOptions_11)
	goto L501
L500:
	;
	if v1528 == int32(0) {
		v1574 = int32(0)
		goto L484
	} else {
		goto L513
	}
L501:
	;
	v1494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1490))))
	v1495 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1491))))
	if v1494 == v1495 {
		v1517 = v1494
		goto L503
	} else {
		goto L504
	}
L502:
	;
	v1528 = int32(0)
	goto L500
L503:
	;
	v1519 = int32(1)
	if v1517 != 0 {
		v1490 = v1490 + v1519
		v1491 = v1491 + v1519
		goto L501
	} else {
		goto L512
	}
L504:
	;
	if base.Ui32((v1494-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v1505 = v1494 | int32(32)
	goto L507
L506:
	;
	v1505 = v1494
	goto L507
L507:
	;
	if base.Ui32((v1495-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v1514 = v1495 | int32(32)
	goto L510
L509:
	;
	v1514 = v1495
	goto L510
L510:
	;
	if v1505 == v1514 {
		v1517 = v1505
		goto L503
	} else {
		goto L511
	}
L511:
	;
	v1528 = v1505 - v1514
	goto L500
L512:
	;
	goto L502
L513:
	;
	v1534 = v1440
	v1535 = int32(_a_F_ProcessCopyOptions_44)
	goto L515
L514:
	;
	if v1572 != 0 {
		goto L483
	} else {
		goto L527
	}
L515:
	;
	v1538 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1534))))
	v1539 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1535))))
	if v1538 == v1539 {
		v1561 = v1538
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v1572 = int32(0)
	goto L514
L517:
	;
	v1563 = int32(1)
	if v1561 != 0 {
		v1534 = v1534 + v1563
		v1535 = v1535 + v1563
		goto L515
	} else {
		goto L526
	}
L518:
	;
	if base.Ui32((v1538-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L519
	} else {
		goto L520
	}
L519:
	;
	v1549 = v1538 | int32(32)
	goto L521
L520:
	;
	v1549 = v1538
	goto L521
L521:
	;
	if base.Ui32((v1539-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v1558 = v1539 | int32(32)
	goto L524
L523:
	;
	v1558 = v1539
	goto L524
L524:
	;
	if v1549 == v1558 {
		v1561 = v1549
		goto L517
	} else {
		goto L525
	}
L525:
	;
	v1572 = v1549 - v1558
	goto L514
L526:
	;
	goto L516
L527:
	;
	v1574 = int32(1)
	goto L484
L528:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1584 = m.ExcPending
	if v1584 != 0 {
		goto L4
	} else {
		goto L529
	}
L529:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1437)+4)) = v1440
	*(*int32)(unsafe.Add(mBase, uint32(v1437))) = int32(_a_F_ProcessCopyOptions_45)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_41), v1437)
	mBase = m.M
	v1590 = m.ExcPending
	if v1590 != 0 {
		goto L4
	} else {
		goto L530
	}
L530:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1591)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L4
	} else {
		goto L531
	}
L531:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(560), int32(_a_F_ProcessCopyOptions_46))
	mBase = m.M
	v1598 = m.ExcPending
	if v1598 != 0 {
		goto L4
	} else {
		goto L532
	}
L532:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L533:
	;
	if v1625-v1626 == int32(0) {
		goto L540
	} else {
		goto L541
	}
L534:
	;
	goto L533
L535:
	;
	v1610 = v57
	v1611 = v1601
	goto L536
L536:
	;
	v1614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1611)+1)))
	v1615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1610)+1)))
	if v1615 == int32(0) {
		v1625 = v1615
		v1626 = v1614
		goto L534
	} else {
		goto L538
	}
L537:
	;
	v1625 = v1615
	v1626 = v1614
	goto L534
L538:
	;
	v1618 = int32(1)
	if v1615 == v1614 {
		v1610 = v1610 + v1618
		v1611 = v1611 + v1618
		goto L536
	} else {
		goto L539
	}
L539:
	;
	goto L537
L540:
	;
	if v50 != 0 {
		goto L20
	} else {
		goto L543
	}
L541:
	;
	goto L542
L542:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1690 = m.ExcPending
	if v1690 != 0 {
		goto L4
	} else {
		goto L564
	}
L543:
	;
	v1630 = m.G0
	v1632 = v1630 - int32(32)
	m.G0 = v1632
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v56)+12))
	if v1634 != 0 {
		goto L546
	} else {
		goto L547
	}
L544:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v27)+104)) = v1644
	v1712 = v43
	v1713 = v45
	v1714 = v46
	v1715 = v47
	v1716 = v48
	v1717 = v49
	v1718 = int32(1)
	goto L26
L545:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L4
	} else {
		goto L560
	}
L546:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1634)))
	if v1635 == int32(476) {
		goto L550
	} else {
		goto L551
	}
L547:
	;
	goto L548
L548:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1653 = m.ExcPending
	if v1653 != 0 {
		goto L4
	} else {
		goto L556
	}
L549:
	;
	if v1644 <= int64(0) {
		goto L545
	} else {
		goto L555
	}
L550:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+4))
	v1640 = F_pg_strtoint64_safe(m, v1638, int32(0))
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L4
	} else {
		goto L553
	}
L551:
	;
	goto L552
L552:
	;
	v1642 = F_defGetInt64(m, v56)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L4
	} else {
		goto L554
	}
L553:
	;
	v1644 = v1640
	goto L549
L554:
	;
	v1644 = v1642
	goto L549
L555:
	;
	m.G0 = v1632 + int32(32)
	goto L544
L556:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L4
	} else {
		goto L557
	}
L557:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1632))) = v1657
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_47), v1632)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L4
	} else {
		goto L558
	}
L558:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(522), int32(_a_F_ProcessCopyOptions_48))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L4
	} else {
		goto L559
	}
L559:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L560:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L4
	} else {
		goto L561
	}
L561:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1632)+16)) = v1644
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_49), v1632+int32(16))
	mBase = m.M
	v1679 = m.ExcPending
	if v1679 != 0 {
		goto L4
	} else {
		goto L562
	}
L562:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(532), int32(_a_F_ProcessCopyOptions_48))
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L4
	} else {
		goto L563
	}
L563:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L564:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L4
	} else {
		goto L565
	}
L565:
	;
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+496)) = v1694
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_50), v20+int32(496))
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L4
	} else {
		goto L566
	}
L566:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v56)+20))
	F_parser_errposition(m, l0, v1701)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L4
	} else {
		goto L567
	}
L567:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(782), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v1708 = m.ExcPending
	if v1708 != 0 {
		goto L4
	} else {
		goto L568
	}
L568:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L569:
	;
	goto L25
L570:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v27)+16))
	if v1771 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L571:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	switch v1744 - int32(1) {
	case 0, 2:
		goto L572
	default:
		goto L570
	}
L572:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1750 = m.ExcPending
	if v1750 != 0 {
		goto L4
	} else {
		goto L573
	}
L573:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1753 = m.ExcPending
	if v1753 != 0 {
		goto L4
	} else {
		goto L574
	}
L574:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+336)) = int32(_a_F_ProcessCopyOptions_51)
	if v1754 == int32(1) {
		goto L575
	} else {
		goto L576
	}
L575:
	;
	v1761 = int32(_a_F_ProcessCopyOptions_52)
	goto L577
L576:
	;
	v1761 = int32(_a_F_ProcessCopyOptions_53)
	goto L577
L577:
	;
	F_errmsg(m, v1761, v20+int32(336))
	mBase = m.M
	v1765 = m.ExcPending
	if v1765 != 0 {
		goto L4
	} else {
		goto L578
	}
L578:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(796), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v1770 = m.ExcPending
	if v1770 != 0 {
		goto L4
	} else {
		goto L579
	}
L579:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L580:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, uint32(v27)+28))
	if v1801 == int32(0) {
		goto L590
	} else {
		goto L591
	}
L581:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	switch v1774 - int32(1) {
	case 0, 2:
		goto L582
	default:
		goto L580
	}
L582:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L4
	} else {
		goto L583
	}
L583:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L4
	} else {
		goto L584
	}
L584:
	;
	v1784 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+320)) = int32(_a_F_ProcessCopyOptions_54)
	if v1784 == int32(1) {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v1791 = int32(_a_F_ProcessCopyOptions_52)
	goto L587
L586:
	;
	v1791 = int32(_a_F_ProcessCopyOptions_53)
	goto L587
L587:
	;
	F_errmsg(m, v1791, v20+int32(320))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L4
	} else {
		goto L588
	}
L588:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(805), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v1800 = m.ExcPending
	if v1800 != 0 {
		goto L4
	} else {
		goto L589
	}
L589:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L590:
	;
	if v1741 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L591:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	switch v1804 - int32(1) {
	case 0, 2:
		goto L592
	default:
		goto L590
	}
L592:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1810 = m.ExcPending
	if v1810 != 0 {
		goto L4
	} else {
		goto L593
	}
L593:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L4
	} else {
		goto L594
	}
L594:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+304)) = int32(_a_F_ProcessCopyOptions_55)
	if v1814 == int32(1) {
		goto L595
	} else {
		goto L596
	}
L595:
	;
	v1821 = int32(_a_F_ProcessCopyOptions_52)
	goto L597
L596:
	;
	v1821 = int32(_a_F_ProcessCopyOptions_53)
	goto L597
L597:
	;
	F_errmsg(m, v1821, v20+int32(304))
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
		goto L4
	} else {
		goto L598
	}
L598:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(814), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L4
	} else {
		goto L599
	}
L599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L600:
	;
	v1835 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v1835 == int32(2) {
		goto L603
	} else {
		goto L604
	}
L601:
	;
	v1840 = v1741
	goto L602
L602:
	;
	v1841 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v1771 == int32(0) {
		goto L606
	} else {
		goto L607
	}
L603:
	;
	v1838 = int32(_a_F_ProcessCopyOptions_56)
	goto L605
L604:
	;
	v1838 = int32(_a_F_ProcessCopyOptions_57)
	goto L605
L605:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+36)) = v1838
	v1840 = v1838
	goto L602
L606:
	;
	if v1841 == int32(2) {
		goto L609
	} else {
		goto L610
	}
L607:
	;
	v1850 = v1771
	goto L608
L608:
	;
	v1851 = F_strlen(m, v1850)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+20)) = v1851
	if v1841 != int32(2) {
		goto L612
	} else {
		goto L613
	}
L609:
	;
	v1848 = int32(_a_F_ProcessCopyOptions_58)
	goto L611
L610:
	;
	v1848 = int32(_a_F_ProcessCopyOptions_59)
	goto L611
L611:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+16)) = v1848
	v1850 = v1848
	goto L608
L612:
	;
	v1865 = F_strlen(m, v1840)
	mBase = m.M
	if v1865 != int32(1) {
		goto L16
	} else {
		goto L618
	}
L613:
	;
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v1855 == int32(0) {
		goto L614
	} else {
		goto L615
	}
L614:
	;
	v1858 = int32(_a_F_ProcessCopyOptions_60)
	*(*int32)(unsafe.Add(mBase, uint32(v27)+40)) = v1858
	v1861 = v1858
	goto L616
L615:
	;
	v1861 = v1855
	goto L616
L616:
	;
	v1862 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v1862 != 0 {
		goto L612
	} else {
		goto L617
	}
L617:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+44)) = v1861
	goto L612
L618:
	;
	v1868 = int32(13)
	v1869 = F___strchrnul(m, v1840, v1868)
	mBase = m.M
	v1871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1869))))
	if v1871 == v1868 {
		goto L620
	} else {
		goto L621
	}
L619:
	;
	if v1875 != 0 {
		goto L15
	} else {
		goto L623
	}
L620:
	;
	v1875 = v1869
	goto L622
L621:
	;
	v1875 = int32(0)
	goto L622
L622:
	;
	goto L619
L623:
	;
	v1876 = int32(10)
	v1877 = F___strchrnul(m, v1840, v1876)
	mBase = m.M
	v1879 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1877))))
	if v1879 == v1876 {
		goto L625
	} else {
		goto L626
	}
L624:
	;
	if v1883 != 0 {
		goto L15
	} else {
		goto L628
	}
L625:
	;
	v1883 = v1877
	goto L627
L626:
	;
	v1883 = int32(0)
	goto L627
L627:
	;
	goto L624
L628:
	;
	v1884 = int32(13)
	v1885 = F___strchrnul(m, v1850, v1884)
	mBase = m.M
	v1887 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1885))))
	if v1887 == v1884 {
		goto L630
	} else {
		goto L631
	}
L629:
	;
	if v1891 != 0 {
		goto L14
	} else {
		goto L633
	}
L630:
	;
	v1891 = v1885
	goto L632
L631:
	;
	v1891 = int32(0)
	goto L632
L632:
	;
	goto L629
L633:
	;
	v1892 = int32(10)
	v1893 = F___strchrnul(m, v1850, v1892)
	mBase = m.M
	v1895 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1893))))
	if v1895 == v1892 {
		goto L635
	} else {
		goto L636
	}
L634:
	;
	if v1899 != 0 {
		goto L14
	} else {
		goto L638
	}
L635:
	;
	v1899 = v1893
	goto L637
L636:
	;
	v1899 = int32(0)
	goto L637
L637:
	;
	goto L634
L638:
	;
	if v1801 != 0 {
		goto L639
	} else {
		goto L640
	}
L639:
	;
	v1900 = F_strlen(m, v1801)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v27)+32)) = v1900
	v1902 = int32(13)
	v1903 = F___strchrnul(m, v1801, v1902)
	mBase = m.M
	v1905 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1903))))
	if v1905 == v1902 {
		goto L643
	} else {
		goto L644
	}
L640:
	;
	goto L641
L641:
	;
	if v1841 != int32(2) {
		goto L652
	} else {
		goto L653
	}
L642:
	;
	if v1909 != 0 {
		goto L13
	} else {
		goto L646
	}
L643:
	;
	v1909 = v1903
	goto L645
L644:
	;
	v1909 = int32(0)
	goto L645
L645:
	;
	goto L642
L646:
	;
	v1910 = int32(10)
	v1911 = F___strchrnul(m, v1801, v1910)
	mBase = m.M
	v1913 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1911))))
	if v1913 == v1910 {
		goto L648
	} else {
		goto L649
	}
L647:
	;
	if v1917 != 0 {
		goto L13
	} else {
		goto L651
	}
L648:
	;
	v1917 = v1911
	goto L650
L649:
	;
	v1917 = int32(0)
	goto L650
L650:
	;
	goto L647
L651:
	;
	goto L641
L652:
	;
	v1921 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1840))))
	goto L659
L653:
	;
	goto L654
L654:
	;
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v27)+12))
	if v2028 == int32(0) {
		goto L681
	} else {
		goto L682
	}
L655:
	;
	if v2027 != 0 {
		goto L12
	} else {
		goto L680
	}
L656:
	;
	v2027 = int32(0)
	goto L655
L657:
	;
	v2005 = v1998
	v2007 = v2000
	goto L674
L658:
	;
	if base.B2i32(v1944 != v1945) == int32(0) {
		goto L656
	} else {
		goto L665
	}
L659:
	;
	v1936 = int32(_a_F_ProcessCopyOptions_61)
	v1938 = int32(39)
	goto L660
L660:
	;
	v1941 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1936))))
	if v1941 == v1921&int32(255) {
		v1998 = v1936
		v2000 = v1938
		goto L657
	} else {
		goto L662
	}
L661:
	;
	goto L658
L662:
	;
	v1943 = int32(1)
	v1944 = v1938 - v1943
	v1945 = int32(0)
	v1948 = v1936 + v1943
	if v1948&int32(3) == v1945 {
		goto L658
	} else {
		goto L663
	}
L663:
	;
	if v1944 != 0 {
		v1936 = v1948
		v1938 = v1944
		goto L660
	} else {
		goto L664
	}
L664:
	;
	goto L661
L665:
	;
	v1961 = v1921 & int32(255)
	v1962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1948))))
	if base.B2i32(v1961 == v1962)|base.B2i32(base.Ui32(v1944) < base.Ui32(int32(4))) == int32(0) {
		goto L666
	} else {
		goto L667
	}
L666:
	;
	v1971 = v1948
	v1973 = v1944
	goto L669
L667:
	;
	v1991 = v1948
	v1993 = v1944
	goto L668
L668:
	;
	if v1993 == int32(0) {
		goto L656
	} else {
		goto L673
	}
L669:
	;
	v1977 = *(*int32)(unsafe.Add(mBase, uint32(v1971)))
	v1978 = v1977 ^ v1961*int32(16843009)
	v1981 = int32(-2139062144)
	if (int32(16843008)-v1978|v1978)&v1981 != v1981 {
		v1998 = v1971
		v2000 = v1973
		goto L657
	} else {
		goto L671
	}
L670:
	;
	v1991 = v1986
	v1993 = v1988
	goto L668
L671:
	;
	v1985 = int32(4)
	v1986 = v1971 + v1985
	v1988 = v1973 - v1985
	if base.Ui32(int32(3)) < base.Ui32(v1988) {
		v1971 = v1986
		v1973 = v1988
		goto L669
	} else {
		goto L672
	}
L672:
	;
	goto L670
L673:
	;
	v1998 = v1991
	v2000 = v1993
	goto L657
L674:
	;
	v2010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2005))))
	if v1921&int32(255) == v2010 {
		goto L676
	} else {
		goto L677
	}
L675:
	;
	goto L656
L676:
	;
	v2027 = v2005
	goto L655
L677:
	;
	goto L678
L678:
	;
	v2012 = int32(1)
	v2015 = v2007 - v2012
	if v2015 != 0 {
		v2005 = v2005 + v2012
		v2007 = v2015
		goto L674
	} else {
		goto L679
	}
L679:
	;
	goto L675
L680:
	;
	goto L654
L681:
	;
	v2057 = *(*int32)(unsafe.Add(mBase, uint32(v27)+40))
	if v1841 != int32(2) {
		goto L700
	} else {
		goto L701
	}
L682:
	;
	switch v1841 - int32(1) {
	case 0, 2:
		goto L683
	default:
		goto L681
	}
L683:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2036 = m.ExcPending
	if v2036 != 0 {
		goto L4
	} else {
		goto L684
	}
L684:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2039 = m.ExcPending
	if v2039 != 0 {
		goto L4
	} else {
		goto L685
	}
L685:
	;
	v2040 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+272)) = int32(_a_F_ProcessCopyOptions_62)
	if v2040 == int32(1) {
		goto L686
	} else {
		goto L687
	}
L686:
	;
	v2047 = int32(_a_F_ProcessCopyOptions_52)
	goto L688
L687:
	;
	v2047 = int32(_a_F_ProcessCopyOptions_53)
	goto L688
L688:
	;
	F_errmsg(m, v2047, v20+int32(272))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L4
	} else {
		goto L689
	}
L689:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(888), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L4
	} else {
		goto L690
	}
L690:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L691:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L770
	}
L692:
	;
	if l2 == int32(0) {
		goto L758
	} else {
		goto L759
	}
L693:
	;
	if l2 == int32(0) {
		goto L753
	} else {
		goto L754
	}
L694:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L751
	}
L695:
	;
	if l2 != 0 {
		goto L739
	} else {
		goto L740
	}
L696:
	;
	v2170 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v2170 != 0 {
		goto L694
	} else {
		goto L732
	}
L697:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	if v2145 == int32(0) {
		goto L724
	} else {
		goto L725
	}
L698:
	;
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	v2126 = F_strlen(m, v2125)
	mBase = m.M
	if v2126 == int32(1) {
		goto L696
	} else {
		goto L719
	}
L699:
	;
	v2103 = *(*int32)(unsafe.Add(mBase, uint32(v27)+44))
	if v2103 == int32(0) {
		goto L697
	} else {
		goto L714
	}
L700:
	;
	if v2057 == int32(0) {
		goto L699
	} else {
		goto L703
	}
L701:
	;
	goto L702
L702:
	;
	v2081 = F_strlen(m, v2057)
	mBase = m.M
	if v2081 != int32(1) {
		goto L11
	} else {
		goto L708
	}
L703:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2065 = m.ExcPending
	if v2065 != 0 {
		goto L4
	} else {
		goto L704
	}
L704:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2068 = m.ExcPending
	if v2068 != 0 {
		goto L4
	} else {
		goto L705
	}
L705:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = int32(_a_F_ProcessCopyOptions_63)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_64), v20+int32(256))
	mBase = m.M
	v2075 = m.ExcPending
	if v2075 != 0 {
		goto L4
	} else {
		goto L706
	}
L706:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(895), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2080 = m.ExcPending
	if v2080 != 0 {
		goto L4
	} else {
		goto L707
	}
L707:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L708:
	;
	v2084 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1840))))
	v2085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2057))))
	if v2084 != v2085 {
		goto L698
	} else {
		goto L709
	}
L709:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2090 = m.ExcPending
	if v2090 != 0 {
		goto L4
	} else {
		goto L710
	}
L710:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L4
	} else {
		goto L711
	}
L711:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_65), int32(0))
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L4
	} else {
		goto L712
	}
L712:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(905), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2102 = m.ExcPending
	if v2102 != 0 {
		goto L4
	} else {
		goto L713
	}
L713:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L714:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2109 = m.ExcPending
	if v2109 != 0 {
		goto L4
	} else {
		goto L715
	}
L715:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2112 = m.ExcPending
	if v2112 != 0 {
		goto L4
	} else {
		goto L716
	}
L716:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = int32(_a_F_ProcessCopyOptions_66)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_64), v20+int32(240))
	mBase = m.M
	v2119 = m.ExcPending
	if v2119 != 0 {
		goto L4
	} else {
		goto L717
	}
L717:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(912), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2124 = m.ExcPending
	if v2124 != 0 {
		goto L4
	} else {
		goto L718
	}
L718:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L719:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2132 = m.ExcPending
	if v2132 != 0 {
		goto L4
	} else {
		goto L720
	}
L720:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L4
	} else {
		goto L721
	}
L721:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_67), int32(0))
	mBase = m.M
	v2139 = m.ExcPending
	if v2139 != 0 {
		goto L4
	} else {
		goto L722
	}
L722:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(917), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L4
	} else {
		goto L723
	}
L723:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L724:
	;
	v2148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+52)))
	if v2148 != int32(1) {
		goto L695
	} else {
		goto L727
	}
L725:
	;
	goto L726
L726:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L4
	} else {
		goto L728
	}
L727:
	;
	goto L726
L728:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2157 = m.ExcPending
	if v2157 != 0 {
		goto L4
	} else {
		goto L729
	}
L729:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = int32(_a_F_ProcessCopyOptions_68)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_64), v20+int32(192))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L4
	} else {
		goto L730
	}
L730:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(924), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2169 = m.ExcPending
	if v2169 != 0 {
		goto L4
	} else {
		goto L731
	}
L731:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L732:
	;
	if l2 != 0 {
		goto L733
	} else {
		goto L734
	}
L733:
	;
	v2171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+52)))
	if v2171&int32(1) != 0 {
		goto L10
	} else {
		goto L736
	}
L734:
	;
	goto L735
L735:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v2174 == int32(0) {
		goto L693
	} else {
		goto L737
	}
L736:
	;
	goto L735
L737:
	;
	if l2 != 0 {
		v2451 = v2057
		goto L7
	} else {
		goto L738
	}
L738:
	;
	goto L6
L739:
	;
	v2177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+52)))
	if v2177&int32(1) != 0 {
		goto L10
	} else {
		goto L742
	}
L740:
	;
	goto L741
L741:
	;
	v2180 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v2180 == int32(0) {
		goto L743
	} else {
		goto L744
	}
L742:
	;
	goto L741
L743:
	;
	v2183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+64)))
	if v2183 != int32(1) {
		goto L692
	} else {
		goto L746
	}
L744:
	;
	goto L745
L745:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L4
	} else {
		goto L747
	}
L746:
	;
	goto L745
L747:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2192 = m.ExcPending
	if v2192 != 0 {
		goto L4
	} else {
		goto L748
	}
L748:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = int32(_a_F_ProcessCopyOptions_69)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_64), v20+int32(208))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L4
	} else {
		goto L749
	}
L749:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(939), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L4
	} else {
		goto L750
	}
L750:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L751:
	;
	v2205 = *(*int32)(unsafe.Add(mBase, uint32(v27)+60))
	if v2205 != 0 {
		goto L6
	} else {
		goto L752
	}
L752:
	;
	goto L693
L753:
	;
	v2208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+64)))
	if v2208&int32(1) != 0 {
		goto L6
	} else {
		goto L756
	}
L754:
	;
	goto L755
L755:
	;
	v2211 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v2211 != 0 {
		goto L691
	} else {
		goto L757
	}
L756:
	;
	goto L755
L757:
	;
	v2423 = v2057
	goto L9
L758:
	;
	v2214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+64)))
	if v2214&int32(1) != 0 {
		goto L6
	} else {
		goto L761
	}
L759:
	;
	goto L760
L760:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(v27)+76))
	if v2217 == int32(0) {
		goto L762
	} else {
		goto L763
	}
L761:
	;
	goto L760
L762:
	;
	v2221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)))
	if v2221 != int32(1) {
		v2423 = int32(0)
		goto L9
	} else {
		goto L765
	}
L763:
	;
	goto L764
L764:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2228 = m.ExcPending
	if v2228 != 0 {
		goto L4
	} else {
		goto L766
	}
L765:
	;
	goto L764
L766:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L4
	} else {
		goto L767
	}
L767:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = int32(_a_F_ProcessCopyOptions_70)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_64), v20+int32(224))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L4
	} else {
		goto L768
	}
L768:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(955), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L4
	} else {
		goto L769
	}
L769:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L770:
	;
	v2451 = v2057
	goto L7
L771:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L772:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L4
	} else {
		goto L773
	}
L773:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+404)) = int32(_a_F_ProcessCopyOptions_18)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+400)) = v2256
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_71), v20+int32(400))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L4
	} else {
		goto L774
	}
L774:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(455), int32(_a_F_ProcessCopyOptions_20))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L4
	} else {
		goto L775
	}
L775:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L776:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L4
	} else {
		goto L777
	}
L777:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+368)) = v2277
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_72), v20+int32(368))
	mBase = m.M
	v2283 = m.ExcPending
	if v2283 != 0 {
		goto L4
	} else {
		goto L778
	}
L778:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(465), int32(_a_F_ProcessCopyOptions_20))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L4
	} else {
		goto L779
	}
L779:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L780:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2295 = m.ExcPending
	if v2295 != 0 {
		goto L4
	} else {
		goto L781
	}
L781:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_73), int32(0))
	mBase = m.M
	v2299 = m.ExcPending
	if v2299 != 0 {
		goto L4
	} else {
		goto L782
	}
L782:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(470), int32(_a_F_ProcessCopyOptions_20))
	mBase = m.M
	v2304 = m.ExcPending
	if v2304 != 0 {
		goto L4
	} else {
		goto L783
	}
L783:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L784:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L4
	} else {
		goto L785
	}
L785:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_74), int32(0))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L4
	} else {
		goto L786
	}
L786:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(836), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L4
	} else {
		goto L787
	}
L787:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L788:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L4
	} else {
		goto L789
	}
L789:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_75), int32(0))
	mBase = m.M
	v2331 = m.ExcPending
	if v2331 != 0 {
		goto L4
	} else {
		goto L790
	}
L790:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(843), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2336 = m.ExcPending
	if v2336 != 0 {
		goto L4
	} else {
		goto L791
	}
L791:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L792:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L4
	} else {
		goto L793
	}
L793:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_76), int32(0))
	mBase = m.M
	v2347 = m.ExcPending
	if v2347 != 0 {
		goto L4
	} else {
		goto L794
	}
L794:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(849), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2352 = m.ExcPending
	if v2352 != 0 {
		goto L4
	} else {
		goto L795
	}
L795:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L796:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2359 = m.ExcPending
	if v2359 != 0 {
		goto L4
	} else {
		goto L797
	}
L797:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_77), int32(0))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L4
	} else {
		goto L798
	}
L798:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(859), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2368 = m.ExcPending
	if v2368 != 0 {
		goto L4
	} else {
		goto L799
	}
L799:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L800:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2375 = m.ExcPending
	if v2375 != 0 {
		goto L4
	} else {
		goto L801
	}
L801:
	;
	v2376 = *(*int32)(unsafe.Add(mBase, uint32(v27)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+288)) = v2376
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_78), v20+int32(288))
	mBase = m.M
	v2382 = m.ExcPending
	if v2382 != 0 {
		goto L4
	} else {
		goto L802
	}
L802:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(877), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L4
	} else {
		goto L803
	}
L803:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L804:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2394 = m.ExcPending
	if v2394 != 0 {
		goto L4
	} else {
		goto L805
	}
L805:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_79), int32(0))
	mBase = m.M
	v2398 = m.ExcPending
	if v2398 != 0 {
		goto L4
	} else {
		goto L806
	}
L806:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(900), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2403 = m.ExcPending
	if v2403 != 0 {
		goto L4
	} else {
		goto L807
	}
L807:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L808:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L4
	} else {
		goto L809
	}
L809:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = int32(_a_F_ProcessCopyOptions_80)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(_a_F_ProcessCopyOptions_68)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_39), v20)
	mBase = m.M
	v2417 = m.ExcPending
	if v2417 != 0 {
		goto L4
	} else {
		goto L810
	}
L810:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(931), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2422 = m.ExcPending
	if v2422 != 0 {
		goto L4
	} else {
		goto L811
	}
L811:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L812:
	;
	v2424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+80)))
	if v2424&int32(1) == int32(0) {
		v2451 = v2423
		goto L7
	} else {
		goto L813
	}
L813:
	;
	goto L8
L814:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2436 = m.ExcPending
	if v2436 != 0 {
		goto L4
	} else {
		goto L815
	}
L815:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+164)) = int32(_a_F_ProcessCopyOptions_37)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = int32(_a_F_ProcessCopyOptions_70)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_39), v20+int32(160))
	mBase = m.M
	v2445 = m.ExcPending
	if v2445 != 0 {
		goto L4
	} else {
		goto L816
	}
L816:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(964), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2450 = m.ExcPending
	if v2450 != 0 {
		goto L4
	} else {
		goto L817
	}
L817:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L818:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2761 = m.ExcPending
	if v2761 != 0 {
		goto L4
	} else {
		goto L934
	}
L819:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2745 = m.ExcPending
	if v2745 != 0 {
		goto L4
	} else {
		goto L930
	}
L820:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2729 = m.ExcPending
	if v2729 != 0 {
		goto L4
	} else {
		goto L926
	}
L821:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2710 = m.ExcPending
	if v2710 != 0 {
		goto L4
	} else {
		goto L922
	}
L822:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2691 = m.ExcPending
	if v2691 != 0 {
		goto L4
	} else {
		goto L918
	}
L823:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2670 = m.ExcPending
	if v2670 != 0 {
		goto L4
	} else {
		goto L914
	}
L824:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L4
	} else {
		goto L910
	}
L825:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2630 = m.ExcPending
	if v2630 != 0 {
		goto L4
	} else {
		goto L906
	}
L826:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2609 = m.ExcPending
	if v2609 != 0 {
		goto L4
	} else {
		goto L902
	}
L827:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2590 = m.ExcPending
	if v2590 != 0 {
		goto L4
	} else {
		goto L898
	}
L828:
	;
	if v2459 == int32(0) {
		goto L832
	} else {
		goto L833
	}
L829:
	;
	v2459 = v2453
	goto L831
L830:
	;
	v2459 = int32(0)
	goto L831
L831:
	;
	goto L828
L832:
	;
	if v1841 == int32(2) {
		goto L835
	} else {
		goto L836
	}
L833:
	;
	goto L834
L834:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2571 = m.ExcPending
	if v2571 != 0 {
		goto L4
	} else {
		goto L894
	}
L835:
	;
	v2464 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2451))))
	v2465 = F___strchrnul(m, v1850, v2464)
	mBase = m.M
	v2467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2465))))
	if v2467 == v2464&int32(255) {
		goto L839
	} else {
		goto L840
	}
L836:
	;
	goto L837
L837:
	;
	if l2 == int32(0) {
		goto L843
	} else {
		goto L844
	}
L838:
	;
	if v2471 != 0 {
		goto L827
	} else {
		goto L842
	}
L839:
	;
	v2471 = v2465
	goto L841
L840:
	;
	v2471 = int32(0)
	goto L841
L841:
	;
	goto L838
L842:
	;
	goto L837
L843:
	;
	v2474 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+8)))
	if v2474&int32(1) != 0 {
		goto L826
	} else {
		goto L846
	}
L844:
	;
	goto L845
L845:
	;
	if v1841 == int32(3) {
		goto L847
	} else {
		goto L848
	}
L846:
	;
	goto L845
L847:
	;
	v2480 = l2
	goto L849
L848:
	;
	v2480 = int32(0)
	goto L849
L849:
	;
	if v2480 != 0 {
		goto L825
	} else {
		goto L850
	}
L850:
	;
	if v1841 != int32(3) {
		goto L851
	} else {
		goto L852
	}
L851:
	;
	v2483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+72)))
	if v2483 == int32(1) {
		goto L824
	} else {
		goto L854
	}
L852:
	;
	goto L853
L853:
	;
	if v1801 == int32(0) {
		goto L855
	} else {
		goto L856
	}
L854:
	;
	goto L853
L855:
	;
	if v1841 == int32(1) {
		goto L886
	} else {
		goto L887
	}
L856:
	;
	if l2 == int32(0) {
		goto L823
	} else {
		goto L857
	}
L857:
	;
	v2490 = F___strchrnul(m, v1801, v2452)
	mBase = m.M
	v2492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2490))))
	if v2492 == v2452&int32(255) {
		goto L859
	} else {
		goto L860
	}
L858:
	;
	if v2496 != 0 {
		goto L822
	} else {
		goto L862
	}
L859:
	;
	v2496 = v2490
	goto L861
L860:
	;
	v2496 = int32(0)
	goto L861
L861:
	;
	goto L858
L862:
	;
	if v1841 == int32(2) {
		goto L863
	} else {
		goto L864
	}
L863:
	;
	v2499 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2451))))
	v2500 = F___strchrnul(m, v1801, v2499)
	mBase = m.M
	v2502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2500))))
	if v2502 == v2499&int32(255) {
		goto L867
	} else {
		goto L868
	}
L864:
	;
	goto L865
L865:
	;
	v2507 = *(*int32)(unsafe.Add(mBase, uint32(v27)+32))
	if v1851 != v2507 {
		goto L855
	} else {
		goto L871
	}
L866:
	;
	if v2506 != 0 {
		goto L821
	} else {
		goto L870
	}
L867:
	;
	v2506 = v2500
	goto L869
L868:
	;
	v2506 = int32(0)
	goto L869
L869:
	;
	goto L866
L870:
	;
	goto L865
L871:
	;
	if v1851 == int32(0) {
		goto L873
	} else {
		goto L874
	}
L872:
	;
	if v2553 == int32(0) {
		goto L820
	} else {
		goto L885
	}
L873:
	;
	v2553 = int32(0)
	goto L872
L874:
	;
	goto L875
L875:
	;
	v2514 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1850))))
	if v2514 != 0 {
		goto L876
	} else {
		goto L877
	}
L876:
	;
	v2515 = v1850
	v2516 = v1801
	v2517 = v1851
	v2518 = v2514
	goto L880
L877:
	;
	v2541 = v1801
	v2545 = int32(0)
	goto L878
L878:
	;
	v2546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2541))))
	v2553 = v2545 - v2546
	goto L872
L879:
	;
	v2541 = v2536
	v2545 = v2538
	goto L878
L880:
	;
	v2520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2516))))
	if base.B2i32(v2518 != v2520)|base.B2i32(v2520 == int32(0)) != 0 {
		v2536 = v2516
		v2538 = v2518
		goto L879
	} else {
		goto L882
	}
L881:
	;
	v2536 = v2530
	v2538 = int32(0)
	goto L879
L882:
	;
	v2526 = v2517 - int32(1)
	if v2526 == int32(0) {
		v2536 = v2516
		v2538 = v2518
		goto L879
	} else {
		goto L883
	}
L883:
	;
	v2529 = int32(1)
	v2530 = v2516 + v2529
	v2531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2515)+1)))
	if v2531 != 0 {
		v2515 = v2515 + v2529
		v2516 = v2530
		v2517 = v2526
		v2518 = v2531
		goto L880
	} else {
		goto L884
	}
L884:
	;
	goto L881
L885:
	;
	goto L855
L886:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v2558 != 0 {
		goto L819
	} else {
		goto L889
	}
L887:
	;
	goto L888
L888:
	;
	v2559 = *(*int64)(unsafe.Add(mBase, uint32(v27)+104))
	if v2559 != int64(0) {
		goto L890
	} else {
		goto L891
	}
L889:
	;
	goto L888
L890:
	;
	v2562 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
	if v2562 != int32(1) {
		goto L818
	} else {
		goto L893
	}
L891:
	;
	goto L892
L892:
	;
	m.G0 = v20 + int32(528)
	return
L893:
	;
	goto L892
L894:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L4
	} else {
		goto L895
	}
L895:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = int32(_a_F_ProcessCopyOptions_54)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_81), v20+int32(144))
	mBase = m.M
	v2581 = m.ExcPending
	if v2581 != 0 {
		goto L4
	} else {
		goto L896
	}
L896:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(972), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2586 = m.ExcPending
	if v2586 != 0 {
		goto L4
	} else {
		goto L897
	}
L897:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L898:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2593 = m.ExcPending
	if v2593 != 0 {
		goto L4
	} else {
		goto L899
	}
L899:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = int32(_a_F_ProcessCopyOptions_54)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_82), v20+int32(128))
	mBase = m.M
	v2600 = m.ExcPending
	if v2600 != 0 {
		goto L4
	} else {
		goto L900
	}
L900:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(981), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2605 = m.ExcPending
	if v2605 != 0 {
		goto L4
	} else {
		goto L901
	}
L901:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L902:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2612 = m.ExcPending
	if v2612 != 0 {
		goto L4
	} else {
		goto L903
	}
L903:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = int32(_a_F_ProcessCopyOptions_37)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = int32(_a_F_ProcessCopyOptions_83)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_39), v20+int32(112))
	mBase = m.M
	v2621 = m.ExcPending
	if v2621 != 0 {
		goto L4
	} else {
		goto L904
	}
L904:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(990), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2626 = m.ExcPending
	if v2626 != 0 {
		goto L4
	} else {
		goto L905
	}
L905:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L906:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2633 = m.ExcPending
	if v2633 != 0 {
		goto L4
	} else {
		goto L907
	}
L907:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = int32(_a_F_ProcessCopyOptions_80)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(_a_F_ProcessCopyOptions_84)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_85), v20+int32(16))
	mBase = m.M
	v2642 = m.ExcPending
	if v2642 != 0 {
		goto L4
	} else {
		goto L908
	}
L908:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(996), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2647 = m.ExcPending
	if v2647 != 0 {
		goto L4
	} else {
		goto L909
	}
L909:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L910:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2654 = m.ExcPending
	if v2654 != 0 {
		goto L4
	} else {
		goto L911
	}
L911:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = int32(_a_F_ProcessCopyOptions_86)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_87), v20+int32(96))
	mBase = m.M
	v2661 = m.ExcPending
	if v2661 != 0 {
		goto L4
	} else {
		goto L912
	}
L912:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1001), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2666 = m.ExcPending
	if v2666 != 0 {
		goto L4
	} else {
		goto L913
	}
L913:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L914:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2673 = m.ExcPending
	if v2673 != 0 {
		goto L4
	} else {
		goto L915
	}
L915:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = int32(_a_F_ProcessCopyOptions_37)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = int32(_a_F_ProcessCopyOptions_55)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_39), v20+int32(80))
	mBase = m.M
	v2682 = m.ExcPending
	if v2682 != 0 {
		goto L4
	} else {
		goto L916
	}
L916:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1011), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2687 = m.ExcPending
	if v2687 != 0 {
		goto L4
	} else {
		goto L917
	}
L917:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L918:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2694 = m.ExcPending
	if v2694 != 0 {
		goto L4
	} else {
		goto L919
	}
L919:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = int32(_a_F_ProcessCopyOptions_55)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_81), v20-int32(-64))
	mBase = m.M
	v2701 = m.ExcPending
	if v2701 != 0 {
		goto L4
	} else {
		goto L920
	}
L920:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1019), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2706 = m.ExcPending
	if v2706 != 0 {
		goto L4
	} else {
		goto L921
	}
L921:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L922:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2713 = m.ExcPending
	if v2713 != 0 {
		goto L4
	} else {
		goto L923
	}
L923:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = int32(_a_F_ProcessCopyOptions_55)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_82), v20+int32(48))
	mBase = m.M
	v2720 = m.ExcPending
	if v2720 != 0 {
		goto L4
	} else {
		goto L924
	}
L924:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1028), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2725 = m.ExcPending
	if v2725 != 0 {
		goto L4
	} else {
		goto L925
	}
L925:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L926:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v2732 = m.ExcPending
	if v2732 != 0 {
		goto L4
	} else {
		goto L927
	}
L927:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_88), int32(0))
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L4
	} else {
		goto L928
	}
L928:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1036), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2741 = m.ExcPending
	if v2741 != 0 {
		goto L4
	} else {
		goto L929
	}
L929:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L930:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L4
	} else {
		goto L931
	}
L931:
	;
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_89), int32(0))
	mBase = m.M
	v2752 = m.ExcPending
	if v2752 != 0 {
		goto L4
	} else {
		goto L932
	}
L932:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1042), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2757 = m.ExcPending
	if v2757 != 0 {
		goto L4
	} else {
		goto L933
	}
L933:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L934:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L4
	} else {
		goto L935
	}
L935:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = int32(_a_F_ProcessCopyOptions_90)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = int32(_a_F_ProcessCopyOptions_38)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = int32(_a_F_ProcessCopyOptions_91)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_92), v20+int32(32))
	mBase = m.M
	v2775 = m.ExcPending
	if v2775 != 0 {
		goto L4
	} else {
		goto L936
	}
L936:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(1050), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2780 = m.ExcPending
	if v2780 != 0 {
		goto L4
	} else {
		goto L937
	}
L937:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L938:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v2787 = m.ExcPending
	if v2787 != 0 {
		goto L4
	} else {
		goto L939
	}
L939:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+180)) = int32(_a_F_ProcessCopyOptions_37)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = int32(_a_F_ProcessCopyOptions_69)
	F_errmsg(m, int32(_a_F_ProcessCopyOptions_39), v20+int32(176))
	mBase = m.M
	v2796 = m.ExcPending
	if v2796 != 0 {
		goto L4
	} else {
		goto L940
	}
L940:
	;
	F_errfinish(m, int32(_a_F_ProcessCopyOptions_7), int32(947), int32(_a_F_ProcessCopyOptions_8))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L4
	} else {
		goto L941
	}
L941:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ProcessPendingWrites(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int64
	_ = v49
	var v53 int32
	_ = v53
	var v55 int64
	_ = v55
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int64
	_ = v86
	var v87 int64
	_ = v87
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v106 int32
	_ = v106
	var v115 int64
	_ = v115
	var v121 int64
	_ = v121
	var v130 int64
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int64
	_ = v137
	var v141 int32
	_ = v141
	var v147 int64
	_ = v147
	var v153 int64
	_ = v153
	var v162 int64
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v267 int32
	_ = v267
	goto L2
L1:
	;
	F_WalSndShutdown(m)
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L4
	} else {
		goto L73
	}
L2:
	;
	F_ProcessRepliesIfAny(m)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[0]))
	v211 = int32(0)
	v214 = base.AtomicRmwOr32(m, v211, int32(_a_F_ProcessPendingWrites_0), v211)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	if v215 != 0 {
		goto L60
	} else {
		goto L61
	}
L4:
	;
	return
L5:
	;
	v12 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[1]))
	if v12 <= int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_WalSndCheckShutdownTimeout(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L14
	}
L7:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[2]))
	if v16 <= int32(0) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v20 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[3]))
	if v20 < base.I64_extend_i32_u(v16)*int64(1000)+v12 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v28 = F_errstart(m, int32(16), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	if v28 == int32(0) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errmsg(m, int32(_a_F_ProcessPendingWrites_1), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(_a_F_ProcessPendingWrites_2), int32(3008), int32(_a_F_ProcessPendingWrites_3))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L1
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[2]))
	if v45 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[4]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+12))
	v76 = m.T0[v75].(func(*base.Module) int32)(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L23
	}
L16:
	;
	v49 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[1]))
	if v49 <= int64(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[5])))
	if v53 != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v55 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[3]))
	if v55 < base.I64_extend_i32_u(int32(base.Ui32(v45)>>(uint(int32(1))%32)))*int64(1000)+v49 {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	F_WalSndKeepalive(m, int32(1), int64(0))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[4]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+8))
	v70 = m.T0[v69].(func(*base.Module) int32)(m)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L21
	}
L21:
	;
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	goto L15
L23:
	;
	if v76 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = m.G0
	v82 = int32(16)
	v83 = v81 - v82
	m.G0 = v83
	F_gettimeofday(m, v83)
	mBase = m.M
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v83)))
	v87 = int64(*(*int32)(unsafe.Add(mBase, uint32(v83)+8)))
	m.G0 = v83 + v82
	v95 = v87 + v86*int64(1000000) - int64(946684800000000)
	goto L27
L25:
	;
	goto L26
L26:
	;
	goto L3
L27:
	;
	v96 = int32(_a_F_ProcessPendingWrites_4)
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[2]))
	if v98 <= int32(0) {
		v134 = v96
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v137 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[6]))
	if v137 == int64(0) {
		v168 = v134
		goto L35
	} else {
		goto L36
	}
L29:
	;
	v102 = *(*int64)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[1]))
	if v102 <= int64(0) {
		v134 = v96
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[5])))
	v115 = base.I64_extend_i32_u(int32(base.Ui32(v98)>>(uint((v106^int32(-1))&int32(1))%32)))*int64(1000) + v102
	if v115 <= v95 {
		v133 = int32(0)
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v134 = v133
	goto L28
L32:
	;
	goto L31
L33:
	;
	v121 = v115 - v95
	if base.B2i32(int64(0) < v95)^base.B2i32(v121 < v115)|base.B2i32(int64(2147483646000) < v121) != 0 {
		v133 = int32(2147483647)
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v130 = base.I64_div_s(v121+int64(999), int64(1000))
	v133 = base.I32_wrap_i64(v130)
	goto L32
L35:
	;
	F_WalSndWait(m, int32(6), v168, int32(100663304))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L45
	}
L36:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[7]))
	if v141 <= int32(0) {
		v168 = v134
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v147 = base.I64_extend_i32_u(v141)*int64(1000) + v137
	if v147 <= v95 {
		v165 = int32(0)
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v165 < v134 {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	goto L38
L40:
	;
	v153 = v147 - v95
	if base.B2i32(int64(0) < v95)^base.B2i32(v153 < v147)|base.B2i32(int64(2147483646000) < v153) != 0 {
		v165 = int32(2147483647)
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v162 = base.I64_div_s(v153+int64(999), int64(1000))
	v165 = base.I32_wrap_i64(v162)
	goto L39
L42:
	;
	v167 = v165
	goto L44
L43:
	;
	v167 = v134
	goto L44
L44:
	;
	v168 = v167
	goto L35
L45:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[0]))
	v176 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v176
	v181 = base.AtomicRmwOr32(m, v176, int32(_a_F_ProcessPendingWrites_0), v176)
	goto L46
L46:
	;
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[8]))
	if v183 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v187 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[9]))
	if v187 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[4]))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	v205 = m.T0[v204].(func(*base.Module) int32)(m)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L57
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[9])) = int32(0)
	F_ProcessConfigFile(m, int32(2))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_SyncRepInitConfig(m)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[10])))
	if v199 != 0 {
		goto L51
	} else {
		goto L55
	}
L55:
	;
	F_SyncRepReleaseWaiters(m)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L4
	} else {
		goto L56
	}
L56:
	;
	goto L51
L57:
	;
	if v205 == int32(0) {
		goto L2
	} else {
		goto L58
	}
L58:
	;
	goto L1
L59:
	;
	return
L60:
	;
	goto L59
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(1)
	v218 = int32(0)
	v221 = base.AtomicRmwOr32(m, v218, int32(_a_F_ProcessPendingWrites_0), v218)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v210)+4))
	if v222 == v218 {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)+12))
	if v225 == int32(0) {
		goto L60
	} else {
		goto L63
	}
L63:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[11]))
	if v229 == v225 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v231 = m.G0
	v233 = v231 - int32(16)
	m.G0 = v233
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[12]))
	if v236 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v259 = F_pgmem_kill(m, v225, int32(23))
	mBase = m.M
	goto L60
L67:
	;
	m.G0 = v233 + int32(16)
	goto L59
L68:
	;
	v239 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v233)+15)) = uint8(v239)
	goto L69
L69:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[13]))
	v247 = F_write(m, v243, v233+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v247 {
		goto L67
	} else {
		goto L71
	}
L70:
	;
	goto L67
L71:
	;
	v251 = *(*int32)(unsafe.Add(mBase, _c_F_ProcessPendingWrites[14]))
	if v251 == int32(27) {
		goto L69
	} else {
		goto L72
	}
L72:
	;
	goto L70
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_PushCopiedSnapshot(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v42 int32
	_ = v42
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
	var v57 int32
	_ = v57
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_PushCopiedSnapshot[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = int32(2)
	v13 = int32(72)
	v18 = v9<<(uint(v11)%32) + v13
	if int32(0) < v8 {
		v21 = (v8+v9)<<(uint(v11)%32) + v13
	} else {
		v21 = v18
	}
	v22 = F_MemoryContextAlloc(m, v7, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return
	} else {
		v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+48))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+48)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v26
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v28
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+56)) = v30
		v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v32
		v34 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+16)) = v34
		v36 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v22)+8)) = v36
		v38 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v22))) = v38
		*(*int64)(unsafe.Add(mBase, uint32(v22)+64)) = int64(0)
		v42 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = v42
		*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = v42
		v46 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v22)+30)) = uint8(v46)
		v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v48 != 0 {
			v50 = v22 + int32(72)
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v50
			v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v54 = v52 << (uint(int32(2)) % 32)
			if v54 == int32(0) {
			} else {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				base.MemoryCopy(m, v50, v57, v54)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = int32(0)
		}
		v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v63 <= int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
		} else {
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+28)))
			if v66 == int32(1) {
				v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
				if v69 != int32(1) {
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
				} else {
					v72 = v22 + v18
					*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v72
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					v76 = v74 << (uint(int32(2)) % 32)
					if v76 == int32(0) {
					} else {
						v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						base.MemoryCopy(m, v72, v79, v76)
					}
				}
			} else {
				v72 = v22 + v18
				*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v72
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v76 = v74 << (uint(int32(2)) % 32)
				if v76 == int32(0) {
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					base.MemoryCopy(m, v72, v79, v76)
				}
			}
		}
		v86 = *(*int32)(unsafe.Add(mBase, _c_F_PushCopiedSnapshot[1]))
		v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+28))
		F_PushActiveSnapshotWithLevel(m, v22, v87)
		mBase = m.M
		v89 = m.ExcPending
		if v89 != 0 {
			return
		} else {
			return
		}
	}
}
func F_p_isalnum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = F_pg_database_locale(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		v11 = int32(2)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+v10<<(uint(v11)%32))))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v15 < v11 {
			v25 = F_pg_database_locale(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = F_pg_iswalnum(m, v14, v25)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = v27
					return v29
				}
			}
		} else {
			v18 = int32(1)
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+2)))
			if v19 != v18 {
				v25 = F_pg_database_locale(m)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = F_pg_iswalnum(m, v14, v25)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						v29 = v27
						return v29
					}
				}
			} else {
				if base.Ui32(int32(127)) < base.Ui32(v14) {
					v29 = v18
					return v29
				} else {
					v25 = F_pg_database_locale(m)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						v27 = F_pg_iswalnum(m, v14, v25)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							v29 = v27
							return v29
						}
					}
				}
			}
		}
	}
}
func F_p_isspecial(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	v7 = int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_p_isspecial[0]))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14*int32(28))+uint32(_c_F_p_isspecial[1])))
	v20 = m.T0[v19].(func(*base.Module, int32) int32)(m, v8+v10)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v70
L2:
	;
	return int32(0)
L3:
	;
	if v20 == int32(0) {
		v70 = v7
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_p_isspecial[0]))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	goto L5
L5:
	;
	if v28 == int32(6) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v31+v33<<(uint(int32(2))%32))))
	v40 = int32(_a_F_p_isspecial_0)
	v41 = int32(_a_F_p_isspecial_1)
	goto L9
L7:
	;
	goto L8
L8:
	;
	v70 = int32(0)
	goto L1
L9:
	;
	v51 = v40 + (v41-v40)>>(uint(int32(3))%32)<<(uint(int32(2))%32)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v52 == v37 {
		v70 = v7
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v56 = base.B2i32(base.Ui32(v52) < base.Ui32(v37))
	if base.Ui32(v52) < base.Ui32(v37) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v57 = v51 + int32(4)
	goto L14
L13:
	;
	v57 = v40
	goto L14
L14:
	;
	if base.Ui32(v52) < base.Ui32(v37) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v58 = v41
	goto L17
L16:
	;
	v58 = v51
	goto L17
L17:
	;
	if base.Ui32(v57) < base.Ui32(v58) {
		v40 = v57
		v41 = v58
		goto L9
	} else {
		goto L18
	}
L18:
	;
	goto L10
}
func F_p_isurlchar(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v24 int32
	_ = v24
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != int32(1) {
		v24 = v2
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8+v9))))
		if base.Ui32((v11-int32(127))&int32(255)) < base.Ui32(int32(162)) {
			v24 = v2
		} else {
			switch v11 - int32(60) {
			case 0, 2, 32, 34, 36, 63, 64, 65:
				v24 = v2
			case 1, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 33, 35, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62:
				v24 = int32(1)
			default:
				if v11 == int32(34) {
					v24 = v2
				} else {
					v24 = int32(1)
				}
			}
		}
	}
	return v24
}
func F_pairingheap_GISTSearchItem_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 float64
	_ = v40
	var v41 float64
	_ = v41
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v50 int64
	_ = v50
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v101 int32
	_ = v101
	v4 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v4 < v10 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v101
L2:
	;
	v101 = int32(0) - v69
	goto L1
L3:
	;
	v13 = int32(32)
	v20 = v10
	v21 = v4
	goto L6
L4:
	;
	goto L5
L5:
	;
	v84 = int32(-1)
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v86 == v84 {
		goto L26
	} else {
		goto L27
	}
L6:
	;
	v27 = v21 << (uint(int32(4)) % 32)
	v28 = l0 + v13 + v27
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	if v29 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L5
L8:
	;
	v73 = v21 + int32(1)
	if v73 < v71 {
		v20 = v71
		v21 = v73
		goto L6
	} else {
		goto L24
	}
L9:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v27)+40)))
	if v33 != 0 {
		v71 = v20
		goto L8
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v36 = v27 + (l1 + v13)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+8)))
	if v37 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	return int32(-1)
L13:
	;
	return int32(1)
L14:
	;
	goto L15
L15:
	;
	v40 = *(*float64)(unsafe.Add(mBase, uint32(v28)))
	v41 = *(*float64)(unsafe.Add(mBase, uint32(v36)))
	v46 = int64(9223372036854775807)
	v47 = base.I64_reinterpret_f64(v40) & v46
	v50 = base.I64_reinterpret_f64(v41) & v46
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(v50) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if v69 != 0 {
		goto L2
	} else {
		goto L23
	}
L17:
	;
	goto L16
L18:
	;
	v69 = int32(0) - v60&(base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v50))|base.F64_lt(v40, v41))
	goto L17
L19:
	;
	v60 = base.B2i32(base.Ui64(v47) < base.Ui64(int64(9218868437227405313)))
	goto L18
L20:
	;
	goto L21
L21:
	;
	v55 = int32(1)
	if base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v47))|base.F64_gt(v40, v41) != 0 {
		v69 = v55
		goto L17
	} else {
		goto L22
	}
L22:
	;
	v60 = v55
	goto L18
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v71 = v70
	goto L8
L24:
	;
	goto L7
L25:
	;
	return int32(0)
L26:
	;
	if v85 == int32(-1) {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v85 == int32(-1) {
		v101 = v84
		goto L1
	} else {
		goto L30
	}
L29:
	;
	return int32(1)
L30:
	;
	goto L25
}
func F_pairingheap_add(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
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
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v8 != 0 {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = m.T0[v10].(func(*base.Module, int32, int32, int32) int32)(m, v8, l1, v9)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			v14 = base.B2i32(v11 < int32(0))
			if v11 < int32(0) {
				v15 = v8
			} else {
				v15 = l1
			}
			if v11 < int32(0) {
				v16 = l1
			} else {
				v16 = v8
			}
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			if v17 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v15
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v16
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v15
			v23 = v16
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
			v28 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v28
			return
		}
	} else {
		v23 = l1
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
		v28 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v23)+8)) = v28
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v28
		return
	}
}
func F_palloc_aligned(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_palloc_aligned[0]))
	v6 = F_MemoryContextAllocAligned(m, v5, l0, l1, l2)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_palloc_mul_extended(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v6 = base.I64_extend_i32_u(l0) * base.I64_extend_i32_u(l1)
	if int64(base.Ui64(v6)>>(uint(int64(32))%64)) == int64(0) {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_palloc_mul_extended[0]))
		v13 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)) = uint8(v13)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int32)(m, v12, base.I32_wrap_i64(v6), int32(2))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v19
		}
	} else {
		F_mul_size_error(m, l0, l1)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_parse_filename_for_nontemp_relation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v34 int64
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v111 int32
	_ = v111
	var v116 int64
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v5
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32((v20-int32(58))&int32(255)) < base.Ui32(int32(247)) {
		v141 = v5
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0])) = int32(0)
		v34 = F_strtox_2(m, l0, v12+int32(8), int32(10), int64(4294967295))
		mBase = m.M
		v35 = base.I32_wrap_i64(v34)
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0]))
		if v37 != 0 {
			v141 = v5
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if base.B2i32(v35 == int32(0))|base.B2i32(l0 == v40) != 0 {
				v141 = v5
			} else {
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
				if v43 != int32(95) {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = int32(0)
					v94 = v43
					v95 = v40
					if v94&int32(255) == int32(46) {
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
						if base.Ui32((v100-int32(58))&int32(255)) < base.Ui32(int32(247)) {
							v141 = v5
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0])) = int32(0)
							v111 = v95 + int32(1)
							v116 = F_strtox_2(m, v111, v12+int32(8), int32(10), int64(4294967295))
							mBase = m.M
							v117 = base.I32_wrap_i64(v116)
							v119 = *(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0]))
							if v119 != 0 {
								v141 = v5
							} else {
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
								if base.B2i32(v117 == int32(0))|base.B2i32(v111 == v122) != 0 {
									v141 = v5
								} else {
									v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
									v128 = v117
									v129 = v125
									if v129&int32(255) != 0 {
										v141 = v5
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
										v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
										*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
										v141 = int32(1)
									}
								}
							}
						}
					} else {
						v128 = v5
						v129 = v94
						if v129&int32(255) != 0 {
							v141 = v5
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
							v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
							v141 = int32(1)
						}
					}
				} else {
					v49 = v40 + int32(1)
					v51 = v12 + int32(12)
					v54 = int32(3)
					v57 = F_strncmp(m, int32(_a_F_parse_filename_for_nontemp_relation_0), v49, v54)
					mBase = m.M
					if v57 == int32(0) {
						v79 = v54
						v80 = int32(1)
						if v51 == int32(0) {
							v86 = v79
						} else {
							v83 = v79
							v84 = v80
							*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
							v86 = v83
						}
					} else {
						v61 = int32(2)
						v65 = F_strncmp(m, int32(_a_F_parse_filename_for_nontemp_relation_1), v49, v61)
						mBase = m.M
						if v65 == int32(0) {
							v79 = v61
							v80 = v61
							if v51 == int32(0) {
								v86 = v79
							} else {
								v83 = v79
								v84 = v80
								*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
								v86 = v83
							}
						} else {
							v68 = int32(4)
							v71 = F_strncmp(m, int32(_a_F_parse_filename_for_nontemp_relation_2), v49, v68)
							mBase = m.M
							if v71 == int32(0) {
								if v51 != 0 {
									v83 = v68
									v84 = int32(3)
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
									v86 = v83
								} else {
									v86 = v68
								}
							} else {
								v76 = int32(0)
								if v51 == v76 {
									v86 = v76
								} else {
									v83 = v76
									v84 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v51))) = v84
									v86 = v83
								}
							}
						}
					}
					if v86 <= int32(0) {
						v141 = v5
					} else {
						v92 = v86 + v40 + int32(1)
						v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
						v94 = v93
						v95 = v92
						if v94&int32(255) == int32(46) {
							v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+1)))
							if base.Ui32((v100-int32(58))&int32(255)) < base.Ui32(int32(247)) {
								v141 = v5
							} else {
								*(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0])) = int32(0)
								v111 = v95 + int32(1)
								v116 = F_strtox_2(m, v111, v12+int32(8), int32(10), int64(4294967295))
								mBase = m.M
								v117 = base.I32_wrap_i64(v116)
								v119 = *(*int32)(unsafe.Add(mBase, _c_F_parse_filename_for_nontemp_relation[0]))
								if v119 != 0 {
									v141 = v5
								} else {
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
									if base.B2i32(v117 == int32(0))|base.B2i32(v111 == v122) != 0 {
										v141 = v5
									} else {
										v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
										v128 = v117
										v129 = v125
										if v129&int32(255) != 0 {
											v141 = v5
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
											v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
											*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
											v141 = int32(1)
										}
									}
								}
							}
						} else {
							v128 = v5
							v129 = v94
							if v129&int32(255) != 0 {
								v141 = v5
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1))) = v35
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l2))) = v133
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v128
								v141 = int32(1)
							}
						}
					}
				}
			}
		}
	}
	m.G0 = v12 + int32(16)
	return v141
}
func F_parse_new_len(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v15 < int32(0) {
			v106 = v15
			m.G0 = v10 + int32(16)
			return v106
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+15)))
			if base.Ui32(v22) < base.Ui32(int32(192)) {
				v100 = v22
				v101 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
				v106 = v101
				m.G0 = v10 + int32(16)
				return v106
			} else {
				if base.Ui32(v22) <= base.Ui32(int32(223)) {
					v30 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(14))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 < int32(0) {
							v106 = v30
						} else {
							v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+14)))
							v100 = v34 | v22<<(uint(int32(8))%32) - int32(_a_F_parse_new_len_0)
							v101 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
							v106 = v101
						}
						m.G0 = v10 + int32(16)
						return v106
					}
				} else {
					if v22 == int32(255) {
						v46 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(13))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							if v46 < int32(0) {
								v106 = v46
								m.G0 = v10 + int32(16)
								return v106
							} else {
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+13)))
								v54 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(12))
								mBase = m.M
								v55 = m.ExcPending
								if v55 != 0 {
									return int32(0)
								} else {
									if v54 < int32(0) {
										v106 = v54
										m.G0 = v10 + int32(16)
										return v106
									} else {
										v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+12)))
										v62 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(11))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											if v62 < int32(0) {
												v106 = v62
												m.G0 = v10 + int32(16)
												return v106
											} else {
												v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
												v70 = F_pullf_read_fixed(m, l0, int32(1), v10+int32(10))
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													if v70 < int32(0) {
														v106 = v70
														m.G0 = v10 + int32(16)
														return v106
													} else {
														v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+10)))
														v75 = int32(8)
														v88 = v74 | (v58<<(uint(v75)%32)|v50<<(uint(int32(16))%32)|v66)<<(uint(v75)%32)
														v92 = int32(1)
														if base.Ui32(v88) < base.Ui32(int32(16777217)) {
															v100 = v88
															v101 = v92
															*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
															v106 = v101
															m.G0 = v10 + int32(16)
															return v106
														} else {
															F_px_debug(m, int32(_a_F_parse_new_len_1), int32(0))
															mBase = m.M
															v98 = m.ExcPending
															if v98 != 0 {
																return int32(0)
															} else {
																v106 = int32(-100)
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
					} else {
						v88 = int32(1) << (uint(v22) % 32)
						v92 = int32(2)
						if base.Ui32(v88) < base.Ui32(int32(16777217)) {
							v100 = v88
							v101 = v92
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v100
							v106 = v101
							m.G0 = v10 + int32(16)
							return v106
						} else {
							F_px_debug(m, int32(_a_F_parse_new_len_1), int32(0))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int32(0)
							} else {
								v106 = int32(-100)
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
func F_parse_weight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	switch l0 - int32(65) {
	case 0, 32:
		v42 = int32(3)
		m.G0 = v6 + int32(32)
		return v42
	case 1, 33:
		v42 = int32(2)
		m.G0 = v6 + int32(32)
		return v42
	case 2, 34:
		v42 = int32(1)
		m.G0 = v6 + int32(32)
		return v42
	case 3, 35:
		v42 = int32(0)
		m.G0 = v6 + int32(32)
		return v42
	default:
		if base.Ui32((l0-int32(32))&int32(255)) <= base.Ui32(int32(94)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
					F_errmsg(m, int32(_a_F_parse_weight_0), v6)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_parse_weight_1), int32(239), int32(_a_F_parse_weight_2))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
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
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0 & int32(255)
					F_errmsg(m, int32(_a_F_parse_weight_3), v6+int32(16))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_parse_weight_1), int32(244), int32(_a_F_parse_weight_2))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
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
func F_parser_errposition(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	v3 = int32(0)
	if base.B2i32(l0 == v3)|base.B2i32(l1 < v3) != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v8 == int32(0) {
			return
		} else {
			v11 = F_pg_mbstrlen_with_len(m, v8, l1)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return
			} else {
				v15 = F_errposition(m, v11+int32(1))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return
				} else {
					return
				}
			}
		}
	}
}
func F_parsetext(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v35 int64
	_ = v35
	var v49 int64
	_ = v49
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var __phi150 int32
	_ = __phi150
	var v151 int32
	_ = v151
	var __phi151 int32
	_ = __phi151
	var v156 int32
	_ = v156
	var __phi156 int32
	_ = __phi156
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
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
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v306 int64
	_ = v306
	var v307 int32
	_ = v307
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v5
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v5
	v23 = F_lookup_ts_config_cache(m, l0)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v26 = F_lookup_ts_parser_cache(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v33 = F_FunctionCall2Coll(m, v26+int32(28), int32(0), base.I64_extend_i32_u(l2), base.I64_extend_i32_s(l3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+44)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v17)+28)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v17)+52)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v17)+60)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = int32(0)
	v49 = v33 & int64(4294967295)
	goto L5
L5:
	;
	v71 = F_FunctionCall3Coll(m, v26+int32(56), int32(0), v49, base.I64_extend_i32_u(v17+int32(72)), base.I64_extend_i32_u(v17+int32(76)))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	v306 = F_FunctionCall1Coll(m, v26+int32(84), int32(0), v49)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L1
	} else {
		goto L57
	}
L7:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v17)+76))
	v76 = base.I32_wrap_i64(v71)
	v77 = int32(0)
	if base.B2i32(v73 < int32(2048))|base.B2i32(v76 <= v77) == v77 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	if int32(0) < v76 {
		goto L5
	} else {
		goto L56
	}
L9:
	;
	v84 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
	v107 = F_palloc(m, int32(16))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L18
	}
L12:
	;
	if v84 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	F_errmsg(m, int32(_a_F_parsetext_0), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(2047)
	v98 = F_errdetail(m, int32(_a_F_parsetext_1), v17)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_parsetext_2), int32(389), int32(_a_F_parsetext_3))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	goto L8
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v107)+4)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v76
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	if v112 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v107
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v116
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v17)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v118
	v123 = F_LexizeExec(m, v17+int32(24), v116)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L23
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = v107
	goto L19
L21:
	;
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v107
	goto L19
L23:
	;
	if v123 == int32(0) {
		goto L8
	} else {
		goto L24
	}
L24:
	;
	v132 = v123
	goto L25
L25:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v141 + int32(1)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)+4))
	if v145 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L8
L27:
	;
	__phi150 = v132
	__phi151 = v145
	__phi156 = v132 + int32(4)
	v150 = __phi150
	v151 = __phi151
	v156 = __phi156
	goto L30
L28:
	;
	goto L29
L29:
	;
	F_pfree(m, v132)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L1
	} else {
		goto L53
	}
L30:
	;
	v162 = F_strlen(m, v151)
	mBase = m.M
	if base.Ui32(int32(2048)) <= base.Ui32(v162) {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L29
L32:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v150)+12))
	if v265 != 0 {
		__phi150 = v150 + int32(8)
		__phi151 = v265
		__phi156 = v150 + int32(12)
		v150 = __phi150
		v151 = __phi151
		v156 = __phi156
		goto L30
	} else {
		goto L52
	}
L33:
	;
	v167 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v190 == v191 {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	if v167 == int32(0) {
		goto L32
	} else {
		goto L37
	}
L37:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_parsetext_0), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(2047)
	v183 = F_errdetail(m, int32(_a_F_parsetext_1), v17+int32(16))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_parsetext_2), int32(417), int32(_a_F_parsetext_3))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	goto L32
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v190 << (uint(int32(1)) % 32)
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v199 = F_repalloc(m, v196, v190<<(uint(int32(5))%32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150)+2)))
	if v202&int32(1) != 0 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v199
	goto L44
L46:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v205 + int32(1)
	goto L48
L47:
	;
	goto L48
L48:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v211 = int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v209+v210<<(uint(v211)%32))+2)) = uint16(v162)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*int32)(unsafe.Add(mBase, uint32(v215+v216<<(uint(v211)%32))+12)) = v220
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150))))
	*(*uint16)(unsafe.Add(mBase, uint32(v222+v223<<(uint(v211)%32))+4)) = uint16(v227)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+2)))
	v236 = v234 & int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v229+v230<<(uint(v211)%32)))) = uint16(v236)
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v243 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v238+v239<<(uint(v211)%32))+6)) = uint16(v243)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v250 = int32(_a_F_parsetext_4)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v250 <= v251 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v254 = v250
	goto L51
L50:
	;
	v254 = v251
	goto L51
L51:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v245+v246<<(uint(v211)%32))+8)) = uint16(v254)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v256 + int32(1)
	goto L32
L52:
	;
	goto L31
L53:
	;
	v285 = F_LexizeExec(m, v17+int32(24), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	if v285 != 0 {
		v132 = v285
		goto L25
	} else {
		goto L55
	}
L55:
	;
	goto L26
L56:
	;
	goto L6
L57:
	;
	m.G0 = v17 + int32(80)
	return
}
func F_patternsel_common(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) float64 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v29 float64
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
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
	var v85 float32
	_ = v85
	var v89 float64
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v108 int64
	_ = v108
	var v109 int32
	_ = v109
	var v112 float64
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 float64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v148 float64
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v163 int32
	_ = v163
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v169 float64
	_ = v169
	var v170 int64
	_ = v170
	var v171 int32
	_ = v171
	var v174 float64
	_ = v174
	var v175 int32
	_ = v175
	var v177 float64
	_ = v177
	var v179 float64
	_ = v179
	var v190 float64
	_ = v190
	var v191 float64
	_ = v191
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v198 float64
	_ = v198
	var v207 float64
	_ = v207
	var v210 float64
	_ = v210
	var v218 float64
	_ = v218
	var v224 float64
	_ = v224
	var v225 int32
	_ = v225
	var v228 float64
	_ = v228
	var v239 float64
	_ = v239
	var v242 float64
	_ = v242
	var v250 float64
	_ = v250
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v267 float64
	_ = v267
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v288 float64
	_ = v288
	v19 = m.G0
	v21 = v19 - int32(96)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21)+40)) = int64(0)
	if l7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = float64(0.995)
	goto L3
L2:
	;
	v29 = float64(0.005)
	goto L3
L3:
	;
	v36 = F_get_restriction_variable(m, l0, l3, l4, v21-int32(-64), v21+int32(60), v21+int32(59))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v21 + int32(96)
	return v288
L5:
	;
	return float64(0)
L6:
	;
	if v36 == int32(0) {
		v288 = v29
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+59)))
	if v42 != int32(1) {
		v267 = v29
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v272 == int32(0) {
		v288 = v267
		goto L4
	} else {
		goto L76
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v46 != int32(7) {
		v267 = v29
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+32)))
	if v49 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v267 = float64(0)
	goto L8
L12:
	;
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v53&int32(-9) != int32(17) {
		v267 = v29
		goto L8
	} else {
		goto L14
	}
L14:
	;
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v45)+24))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v21)+80))
	switch v62 - int32(17) {
	case 0:
		goto L18
	case 1, 3, 4, 5, 6, 7:
		v267 = v29
		goto L8
	case 2:
		goto L16
	case 8:
		v77 = v62
		v78 = int32(98)
		v79 = int32(664)
		v80 = int32(667)
		goto L15
	default:
		goto L17
	}
L15:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v21)+72))
	if v81 != 0 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v77 = int32(25)
	v78 = int32(254)
	v79 = int32(255)
	v80 = int32(257)
	goto L15
L17:
	;
	if v62 != int32(1042) {
		v267 = v29
		goto L8
	} else {
		goto L19
	}
L18:
	;
	v77 = v62
	v78 = int32(1955)
	v79 = int32(1957)
	v80 = int32(1960)
	goto L15
L19:
	;
	v77 = v62
	v78 = int32(1054)
	v79 = int32(1058)
	v80 = int32(1061)
	goto L15
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+16))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+22)))
	v85 = *(*float32)(unsafe.Add(mBase, uint32(v82+v83)+8))
	v89 = base.F64_promote_f32(v85)
	goto L22
L21:
	;
	v89 = float64(0)
	goto L22
L22:
	;
	v94 = F_pattern_fixed_prefix(m, v45, l6, l5, v21+int32(52), v21+int32(40))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v96 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v94 == int32(2) {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	if v99 == v77 {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = v77
	goto L24
L27:
	;
	if v96 == int32(0) {
		v267 = v250
		goto L8
	} else {
		goto L73
	}
L28:
	;
	if l7 != 0 {
		goto L68
	} else {
		goto L69
	}
L29:
	;
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	v109 = int32(0)
	v112 = F_var_eq_const(m, v21-int32(-64), v78, l5, v108, v109, int32(1), v109)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	if l2 != 0 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v239 = v112
	goto L28
L33:
	;
	v116 = l2
	goto L35
L34:
	;
	v114 = F_get_opcode(m, l1)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L36
	}
L35:
	;
	v118 = v21 + int32(8)
	F_fmgr_info(m, v116, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L5
	} else {
		goto L37
	}
L36:
	;
	v116 = v114
	goto L35
L37:
	;
	v122 = v21 - int32(-64)
	v126 = F_histogram_selectivity(m, v122, v118, l5, v58, int32(1), v21+int32(36))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L5
	} else {
		goto L38
	}
L38:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	if int32(99) < v128 {
		v207 = v126
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v210 = float64(0.0001)
	if base.F64_lt(v207, v210) != 0 {
		v218 = v210
		goto L64
	} else {
		goto L65
	}
L40:
	;
	if v94 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v133 = m.G0
	v135 = v133 - int32(32)
	m.G0 = v135
	v137 = F_get_opcode(m, v80)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L5
	} else {
		goto L44
	}
L42:
	;
	v190 = float64(1)
	goto L43
L43:
	;
	v191 = *(*float64)(unsafe.Add(mBase, uint32(v21)+40))
	v192 = base.F64_mul(v190, v191)
	if base.F64_lt(v126, float64(0)) != 0 {
		goto L61
	} else {
		goto L62
	}
L44:
	;
	v140 = v135 + int32(4)
	F_fmgr_info(m, v137, v140)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	v144 = int32(1)
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v96)+4))
	v148 = F_ineq_histogram_selectivity(m, l0, v122, v80, v140, v144, v144, l5, v146, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	if base.F64_lt(v148, float64(0)) == int32(0) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v154 = F_get_opcode(m, v79)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L50
	}
L48:
	;
	v179 = float64(0.005)
	goto L49
L49:
	;
	m.G0 = v135 + int32(32)
	v190 = v179
	goto L43
L50:
	;
	F_fmgr_info(m, v154, v140)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L5
	} else {
		goto L51
	}
L51:
	;
	v158 = F_make_greater_string(m, v96, v140, l5)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	if v158 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v160 = int32(0)
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v158)+24))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v164 = F_ineq_histogram_selectivity(m, l0, v122, v79, v140, v160, v160, l5, v162, v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L5
	} else {
		goto L56
	}
L54:
	;
	v169 = v148
	goto L55
L55:
	;
	v170 = *(*int64)(unsafe.Add(mBase, uint32(v96)+24))
	v171 = int32(0)
	v174 = F_var_eq_const(m, v122, v78, l5, v170, v171, int32(1), v171)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L5
	} else {
		goto L57
	}
L56:
	;
	v169 = base.F64_add(base.F64_add(v148, v164), float64(-1))
	goto L55
L57:
	;
	if base.F64_lt(v174, v169) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v177 = v169
	goto L60
L59:
	;
	v177 = v174
	goto L60
L60:
	;
	v179 = v177
	goto L49
L61:
	;
	v207 = v192
	goto L39
L62:
	;
	goto L63
L63:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v198 = base.F64_div(base.F64_convert_i32_s(v195), float64(100))
	v207 = base.F64_add(base.F64_mul(v126, v198), base.F64_mul(v192, base.F64_sub(float64(1), v198)))
	goto L39
L64:
	;
	v224 = F_mcv_selectivity(m, v21-int32(-64), v21+int32(8), l5, v58, int32(1), v21)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L5
	} else {
		goto L67
	}
L65:
	;
	if base.F64_gt(v207, float64(0.9999)) == int32(0) {
		v218 = v207
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v218 = float64(0.9999)
	goto L64
L67:
	;
	v228 = *(*float64)(unsafe.Add(mBase, uint32(v21)))
	v239 = base.F64_add(v224, base.F64_mul(v218, base.F64_sub(base.F64_sub(float64(1), v89), v228)))
	goto L28
L68:
	;
	v242 = base.F64_sub(base.F64_sub(float64(1), v239), v89)
	goto L70
L69:
	;
	v242 = v239
	goto L70
L70:
	;
	if base.F64_lt(v242, float64(0)) != 0 {
		v250 = float64(0)
		goto L27
	} else {
		goto L71
	}
L71:
	;
	if base.F64_gt(v242, float64(1)) == int32(0) {
		v250 = v242
		goto L27
	} else {
		goto L72
	}
L72:
	;
	v250 = float64(1)
	goto L27
L73:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v96)+24))
	F_pfree(m, v253)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L5
	} else {
		goto L74
	}
L74:
	;
	F_pfree(m, v96)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L5
	} else {
		goto L75
	}
L75:
	;
	v267 = v250
	goto L8
L76:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v21)+76))
	m.T0[v275].(func(*base.Module, int32))(m, v272)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v288 = v267
	goto L4
}
func F_pcb_error_callback(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = F_geterrcode(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if v3 == int32(67371461) {
			return
		} else {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v7 == int32(0) {
				return
			} else {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v10 < int32(0) {
					return
				} else {
					v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					if v13 == int32(0) {
						return
					} else {
						v16 = F_pg_mbstrlen_with_len(m, v13, v10)
						mBase = m.M
						v17 = m.ExcPending
						if v17 != 0 {
							return
						} else {
							v20 = F_errposition(m, v16+int32(1))
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_pfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v4&int32(15)*int32(36))+uint32(_c_F_pfree[0])))
	m.T0[v9].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		return
	}
}
func F_pgsql_version(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_cstring_to_text(m, int32(_a_F_pgsql_version_0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v3)
	}
}
func F_pgstattuple_approx(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(16797828))
				mBase = m.M
				v16 = m.ExcPending
				if v16 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_pgstattuple_approx_0), int32(0))
					mBase = m.M
					v20 = m.ExcPending
					if v20 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_pgstattuple_approx_1), int32(286), int32(_a_F_pgstattuple_approx_2))
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
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
			v27 = F_pgstattuple_approx_internal(m, base.I32_wrap_i64(v3), l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int64(0)
			} else {
				return v27
			}
		}
	}
}
func F_pgstattuple_approx_v1_5(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pgstattuple_approx_internal(m, v2, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_plainto_tsquery_byid(m *base.Module, l0 int32) int64 {
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14390(m, l0, int32(1), int32(2))
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_portuguese_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v480 int32
	_ = v480
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v529 int32
	_ = v529
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v555 int32
	_ = v555
	var v568 int32
	_ = v568
	var v570 int32
	_ = v570
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v705 int32
	_ = v705
	var v713 int32
	_ = v713
	var v721 int32
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v744 int32
	_ = v744
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v761 int32
	_ = v761
	var v770 int32
	_ = v770
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v795 int32
	_ = v795
	var v797 int32
	_ = v797
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v816 int32
	_ = v816
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v841 int32
	_ = v841
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v897 int32
	_ = v897
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v907 int32
	_ = v907
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1020 int32
	_ = v1020
	var v1022 int32
	_ = v1022
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1074 int32
	_ = v1074
	var v1076 int32
	_ = v1076
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1114 int32
	_ = v1114
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1243 int32
	_ = v1243
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v1243
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v15 <= v9 {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v57
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v79 < v7 {
		goto L24
	} else {
		goto L25
	}
L4:
	;
	goto L3
L5:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v65
	goto L2
L6:
	;
	if v57 <= v55 {
		goto L4
	} else {
		goto L19
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v55 = v9
	v57 = v15
	goto L6
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17+v9))))
	v21 = v19 - int32(227)
	v22 = int32(0)
	if base.B2i32(v21 == v22)|base.B2i32(v21 == int32(18)) == v22 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v32 = F_find_among(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_0), int32(3), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v36
	switch v32 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	case 2:
		goto L14
	default:
		goto L5
	}
L12:
	;
	v49 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_ISO_8859_1_stem_1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L17
	}
L13:
	;
	v43 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_ISO_8859_1_stem_2))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v55 = v36
	v57 = v40
	goto L6
L15:
	;
	if int32(0) <= v43 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v1243 = v43
	goto L1
L17:
	;
	if int32(0) <= v49 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	v1243 = v49
	goto L1
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v55 + int32(1)
	goto L5
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v568 < v7 {
		goto L168
	} else {
		goto L169
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = v555
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v349 < v7 {
		goto L102
	} else {
		goto L103
	}
L23:
	;
	if v122 != 0 {
		goto L22
	} else {
		goto L37
	}
L24:
	;
	v81 = v7
	goto L26
L25:
	;
	v81 = v79
	goto L26
L26:
	;
	goto L28
L27:
	;
	v122 = v117
	goto L23
L28:
	;
	if v7 == v81 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v117 = int32(0)
	goto L27
L30:
	;
	v122 = int32(-1)
	goto L23
L31:
	;
	goto L32
L32:
	;
	v93 = int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v7))))
	if int32(250) < v96 {
		v117 = v93
		goto L27
	} else {
		goto L33
	}
L33:
	;
	v98 = v96 - int32(97)
	if v98 < int32(0) {
		v117 = v93
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v98)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v104)>>(uint(v98&int32(7))%32))&int32(1) == int32(0) {
		v117 = v93
		goto L27
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 + int32(1)
	goto L36
L36:
	;
	goto L29
L37:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v132 < v123 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v237 < v123 {
		goto L72
	} else {
		goto L73
	}
L39:
	;
	if v172 != 0 {
		goto L38
	} else {
		goto L54
	}
L40:
	;
	v134 = v123
	goto L42
L41:
	;
	v134 = v132
	goto L42
L42:
	;
	goto L44
L43:
	;
	v172 = v169
	goto L39
L44:
	;
	if v123 == v134 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v169 = int32(0)
	goto L43
L46:
	;
	v172 = int32(-1)
	goto L39
L47:
	;
	goto L48
L48:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145+v123))))
	if int32(250) < v147 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(1)
	goto L53
L50:
	;
	v149 = v147 - int32(97)
	if v149 < int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	v152 = int32(1)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v149)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v156)>>(uint(v149&int32(7))%32))&v152 != 0 {
		v169 = v152
		goto L43
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	goto L45
L54:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v181 < v180 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v221 < int32(0) {
		goto L38
	} else {
		goto L70
	}
L56:
	;
	v183 = v180
	goto L58
L57:
	;
	v183 = v181
	goto L58
L58:
	;
	v190 = v180
	goto L60
L59:
	;
	v221 = v201
	goto L55
L60:
	;
	if v190 == v183 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v221 = int32(-1)
	goto L55
L63:
	;
	goto L64
L64:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v190))))
	if int32(250) < v196 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v213 = v190 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v213
	v190 = v213
	goto L60
L66:
	;
	v198 = v196 - int32(97)
	if v198 < int32(0) {
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v201 = int32(1)
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v198)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v205)>>(uint(v198&int32(7))%32))&v201 != 0 {
		goto L59
	} else {
		goto L68
	}
L68:
	;
	goto L65
L70:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v555 = v224 + v221
	goto L21
L71:
	;
	if v280 != 0 {
		goto L22
	} else {
		goto L85
	}
L72:
	;
	v239 = v123
	goto L74
L73:
	;
	v239 = v237
	goto L74
L74:
	;
	goto L76
L75:
	;
	v280 = v275
	goto L71
L76:
	;
	if v123 == v239 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v275 = int32(0)
	goto L75
L78:
	;
	v280 = int32(-1)
	goto L71
L79:
	;
	goto L80
L80:
	;
	v251 = int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v252+v123))))
	if int32(250) < v254 {
		v275 = v251
		goto L75
	} else {
		goto L81
	}
L81:
	;
	v256 = v254 - int32(97)
	if v256 < int32(0) {
		v275 = v251
		goto L75
	} else {
		goto L82
	}
L82:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v256)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v262)>>(uint(v256&int32(7))%32))&int32(1) == int32(0) {
		v275 = v251
		goto L75
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 + int32(1)
	goto L84
L84:
	;
	goto L77
L85:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v290 < v289 {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	if v333 < int32(0) {
		goto L22
	} else {
		goto L100
	}
L87:
	;
	v292 = v289
	goto L89
L88:
	;
	v292 = v290
	goto L89
L89:
	;
	v298 = v289
	goto L91
L90:
	;
	v333 = int32(1)
	goto L86
L91:
	;
	if v298 == v292 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v333 = int32(-1)
	goto L86
L94:
	;
	goto L95
L95:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v305+v298))))
	if int32(250) < v307 {
		goto L90
	} else {
		goto L96
	}
L96:
	;
	v309 = v307 - int32(97)
	if v309 < int32(0) {
		goto L90
	} else {
		goto L97
	}
L97:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v309)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v315)>>(uint(v309&int32(7))%32))&int32(1) == int32(0) {
		goto L90
	} else {
		goto L98
	}
L98:
	;
	v324 = v298 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v324
	v298 = v324
	goto L91
L100:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v555 = v336 + v333
	goto L21
L101:
	;
	if v389 != 0 {
		goto L20
	} else {
		goto L116
	}
L102:
	;
	v351 = v7
	goto L104
L103:
	;
	v351 = v349
	goto L104
L104:
	;
	goto L106
L105:
	;
	v389 = v386
	goto L101
L106:
	;
	if v7 == v351 {
		goto L108
	} else {
		goto L109
	}
L107:
	;
	v386 = int32(0)
	goto L105
L108:
	;
	v389 = int32(-1)
	goto L101
L109:
	;
	goto L110
L110:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v364 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362+v7))))
	if int32(250) < v364 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7 + int32(1)
	goto L115
L112:
	;
	v366 = v364 - int32(97)
	if v366 < int32(0) {
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v369 = int32(1)
	v373 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v366)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v373)>>(uint(v366&int32(7))%32))&v369 != 0 {
		v386 = v369
		goto L105
	} else {
		goto L114
	}
L114:
	;
	goto L111
L115:
	;
	goto L107
L116:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v399 < v390 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v390
	v504 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v504 < v390 {
		goto L151
	} else {
		goto L152
	}
L118:
	;
	if v439 != 0 {
		goto L117
	} else {
		goto L133
	}
L119:
	;
	v401 = v390
	goto L121
L120:
	;
	v401 = v399
	goto L121
L121:
	;
	goto L123
L122:
	;
	v439 = v436
	goto L118
L123:
	;
	if v390 == v401 {
		goto L125
	} else {
		goto L126
	}
L124:
	;
	v436 = int32(0)
	goto L122
L125:
	;
	v439 = int32(-1)
	goto L118
L126:
	;
	goto L127
L127:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v412+v390))))
	if int32(250) < v414 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v390 + int32(1)
	goto L132
L129:
	;
	v416 = v414 - int32(97)
	if v416 < int32(0) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v419 = int32(1)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v416)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v423)>>(uint(v416&int32(7))%32))&v419 != 0 {
		v436 = v419
		goto L122
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	goto L124
L133:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v448 < v447 {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	if v488 < int32(0) {
		goto L117
	} else {
		goto L149
	}
L135:
	;
	v450 = v447
	goto L137
L136:
	;
	v450 = v448
	goto L137
L137:
	;
	v457 = v447
	goto L139
L138:
	;
	v488 = v468
	goto L134
L139:
	;
	if v457 == v450 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v488 = int32(-1)
	goto L134
L142:
	;
	goto L143
L143:
	;
	v461 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v461+v457))))
	if int32(250) < v463 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v480 = v457 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v480
	v457 = v480
	goto L139
L145:
	;
	v465 = v463 - int32(97)
	if v465 < int32(0) {
		goto L144
	} else {
		goto L146
	}
L146:
	;
	v468 = int32(1)
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v465)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v472)>>(uint(v465&int32(7))%32))&v468 != 0 {
		goto L138
	} else {
		goto L147
	}
L147:
	;
	goto L144
L149:
	;
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v555 = v491 + v488
	goto L21
L150:
	;
	if v547 != 0 {
		goto L20
	} else {
		goto L164
	}
L151:
	;
	v506 = v390
	goto L153
L152:
	;
	v506 = v504
	goto L153
L153:
	;
	goto L155
L154:
	;
	v547 = v542
	goto L150
L155:
	;
	if v390 == v506 {
		goto L157
	} else {
		goto L158
	}
L156:
	;
	v542 = int32(0)
	goto L154
L157:
	;
	v547 = int32(-1)
	goto L150
L158:
	;
	goto L159
L159:
	;
	v518 = int32(1)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v519+v390))))
	if int32(250) < v521 {
		v542 = v518
		goto L154
	} else {
		goto L160
	}
L160:
	;
	v523 = v521 - int32(97)
	if v523 < int32(0) {
		v542 = v518
		goto L154
	} else {
		goto L161
	}
L161:
	;
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v523)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v529)>>(uint(v523&int32(7))%32))&int32(1) == int32(0) {
		v542 = v518
		goto L154
	} else {
		goto L162
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v390 + int32(1)
	goto L163
L163:
	;
	goto L156
L164:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v549 <= v548 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	v555 = v548 + int32(1)
	goto L21
L166:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v787 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v787
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v787
	if v787-int32(2) <= v7 {
		goto L231
	} else {
		goto L232
	}
L167:
	;
	if v608 < int32(0) {
		goto L166
	} else {
		goto L182
	}
L168:
	;
	v570 = v7
	goto L170
L169:
	;
	v570 = v568
	goto L170
L170:
	;
	v577 = v7
	goto L172
L171:
	;
	v608 = v588
	goto L167
L172:
	;
	if v577 == v570 {
		goto L174
	} else {
		goto L175
	}
L174:
	;
	v608 = int32(-1)
	goto L167
L175:
	;
	goto L176
L176:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v581+v577))))
	if int32(250) < v583 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v600 = v577 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v600
	v577 = v600
	goto L172
L178:
	;
	v585 = v583 - int32(97)
	if v585 < int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v588 = int32(1)
	v592 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v585)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v592)>>(uint(v585&int32(7))%32))&v588 != 0 {
		goto L171
	} else {
		goto L180
	}
L180:
	;
	goto L177
L182:
	;
	v611 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v612 = v611 + v608
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v612
	v623 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v623 < v612 {
		goto L184
	} else {
		goto L185
	}
L183:
	;
	if v666 < int32(0) {
		goto L166
	} else {
		goto L197
	}
L184:
	;
	v625 = v612
	goto L186
L185:
	;
	v625 = v623
	goto L186
L186:
	;
	v631 = v612
	goto L188
L187:
	;
	v666 = int32(1)
	goto L183
L188:
	;
	if v631 == v625 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v666 = int32(-1)
	goto L183
L191:
	;
	goto L192
L192:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v638+v631))))
	if int32(250) < v640 {
		goto L187
	} else {
		goto L193
	}
L193:
	;
	v642 = v640 - int32(97)
	if v642 < int32(0) {
		goto L187
	} else {
		goto L194
	}
L194:
	;
	v648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v642)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v648)>>(uint(v642&int32(7))%32))&int32(1) == int32(0) {
		goto L187
	} else {
		goto L195
	}
L195:
	;
	v657 = v631 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v657
	v631 = v657
	goto L188
L197:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v670 = v669 + v666
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v670
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v670
	v681 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v681 < v670 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	if v721 < int32(0) {
		goto L166
	} else {
		goto L213
	}
L199:
	;
	v683 = v670
	goto L201
L200:
	;
	v683 = v681
	goto L201
L201:
	;
	v690 = v670
	goto L203
L202:
	;
	v721 = v701
	goto L198
L203:
	;
	if v690 == v683 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v721 = int32(-1)
	goto L198
L206:
	;
	goto L207
L207:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v696 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694+v690))))
	if int32(250) < v696 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	v713 = v690 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v713
	v690 = v713
	goto L203
L209:
	;
	v698 = v696 - int32(97)
	if v698 < int32(0) {
		goto L208
	} else {
		goto L210
	}
L210:
	;
	v701 = int32(1)
	v705 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v698)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v705)>>(uint(v698&int32(7))%32))&v701 != 0 {
		goto L202
	} else {
		goto L211
	}
L211:
	;
	goto L208
L213:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v725 = v724 + v721
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v725
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v736 < v725 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v779 < int32(0) {
		goto L166
	} else {
		goto L228
	}
L215:
	;
	v738 = v725
	goto L217
L216:
	;
	v738 = v736
	goto L217
L217:
	;
	v744 = v725
	goto L219
L218:
	;
	v779 = int32(1)
	goto L214
L219:
	;
	if v744 == v738 {
		goto L221
	} else {
		goto L222
	}
L221:
	;
	v779 = int32(-1)
	goto L214
L222:
	;
	goto L223
L223:
	;
	v751 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v751+v744))))
	if int32(250) < v753 {
		goto L218
	} else {
		goto L224
	}
L224:
	;
	v755 = v753 - int32(97)
	if v755 < int32(0) {
		goto L218
	} else {
		goto L225
	}
L225:
	;
	v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v755)>>(uint(int32(3))%32)))+uint32(_c_F_portuguese_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v761)>>(uint(v755&int32(7))%32))&int32(1) == int32(0) {
		goto L218
	} else {
		goto L226
	}
L226:
	;
	v770 = v744 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v770
	v744 = v770
	goto L219
L228:
	;
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v782 + v779
	goto L166
L229:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1114
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1114
	v1120 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_3), int32(4), int32(0))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L10
	} else {
		goto L320
	}
L230:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1083
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1083
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1083 <= v1086 {
		goto L229
	} else {
		goto L313
	}
L231:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1042
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1044 <= v1042 {
		goto L303
	} else {
		goto L304
	}
L232:
	;
	v793 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v795 = int32(1)
	v797 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v793+v787-v795))))
	if base.B2i32(v797&int32(224) != int32(96))|base.B2i32(v795<<(uint(v797)%32)&int32(_a_F_portuguese_ISO_8859_1_stem_4) == int32(0)) != 0 {
		goto L231
	} else {
		goto L233
	}
L233:
	;
	v812 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_5), int32(45), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L10
	} else {
		goto L234
	}
L234:
	;
	if v812 == int32(0) {
		goto L231
	} else {
		goto L235
	}
L235:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v816
	switch v812 - int32(1) {
	case 0:
		goto L244
	case 1:
		goto L243
	case 2:
		goto L242
	case 3:
		goto L241
	case 4:
		goto L240
	case 5:
		goto L239
	case 6:
		goto L238
	case 7:
		goto L237
	case 8:
		goto L236
	default:
		goto L230
	}
L236:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v816 < v1020 {
		goto L231
	} else {
		goto L297
	}
L237:
	;
	v988 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v988 {
		goto L231
	} else {
		goto L288
	}
L238:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v949 {
		goto L231
	} else {
		goto L280
	}
L239:
	;
	v917 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v917 {
		goto L231
	} else {
		goto L272
	}
L240:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v816 < v849 {
		goto L231
	} else {
		goto L256
	}
L241:
	;
	v841 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v841 {
		goto L231
	} else {
		goto L253
	}
L242:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v833 {
		goto L231
	} else {
		goto L250
	}
L243:
	;
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v825 {
		goto L231
	} else {
		goto L247
	}
L244:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v816 < v820 {
		goto L231
	} else {
		goto L245
	}
L245:
	;
	v822 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v822 {
		goto L230
	} else {
		goto L246
	}
L246:
	;
	v1243 = v822
	goto L1
L247:
	;
	v829 = F_slice_from_s(m, l0, int32(3), int32(_a_F_portuguese_ISO_8859_1_stem_6))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L10
	} else {
		goto L248
	}
L248:
	;
	if int32(0) <= v829 {
		goto L230
	} else {
		goto L249
	}
L249:
	;
	v1243 = v829
	goto L1
L250:
	;
	v837 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_7))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L10
	} else {
		goto L251
	}
L251:
	;
	if int32(0) <= v837 {
		goto L230
	} else {
		goto L252
	}
L252:
	;
	v1243 = v837
	goto L1
L253:
	;
	v845 = F_slice_from_s(m, l0, int32(4), int32(_a_F_portuguese_ISO_8859_1_stem_8))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L10
	} else {
		goto L254
	}
L254:
	;
	if int32(0) <= v845 {
		goto L230
	} else {
		goto L255
	}
L255:
	;
	v1243 = v845
	goto L1
L256:
	;
	v851 = F_slice_del(m, l0)
	mBase = m.M
	if v851 < int32(0) {
		v1243 = v851
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v854
	v857 = v854 - int32(1)
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v857 <= v858 {
		goto L230
	} else {
		goto L258
	}
L258:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v862 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v860+v857))))
	if base.B2i32(v862&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v862)%32)&int32(_a_F_portuguese_ISO_8859_1_stem_9) == int32(0)) != 0 {
		goto L230
	} else {
		goto L259
	}
L259:
	;
	v877 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_10), int32(4), int32(0))
	mBase = m.M
	v878 = m.ExcPending
	if v878 != 0 {
		goto L10
	} else {
		goto L260
	}
L260:
	;
	if v877 == int32(0) {
		goto L230
	} else {
		goto L261
	}
L261:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v881
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v881 < v883 {
		goto L230
	} else {
		goto L262
	}
L262:
	;
	v885 = F_slice_del(m, l0)
	mBase = m.M
	if v885 < int32(0) {
		v1243 = v885
		goto L1
	} else {
		goto L263
	}
L263:
	;
	if v877 != int32(1) {
		goto L230
	} else {
		goto L264
	}
L264:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v890
	v892 = int32(2)
	v894 = int32(0)
	v897 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v890-v897 < v892 {
		v907 = v894
		goto L266
	} else {
		goto L267
	}
L265:
	;
	if v907 == int32(0) {
		goto L230
	} else {
		goto L269
	}
L266:
	;
	goto L265
L267:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v903 = F_memcmp(m, v900+v890-v892, int32(_a_F_portuguese_ISO_8859_1_stem_11), v892)
	mBase = m.M
	if v903 != 0 {
		v907 = v894
		goto L266
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v890 - v892
	v907 = int32(1)
	goto L266
L269:
	;
	v910 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v910
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v910 < v912 {
		goto L230
	} else {
		goto L270
	}
L270:
	;
	v914 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v914 {
		goto L230
	} else {
		goto L271
	}
L271:
	;
	v1243 = v914
	goto L1
L272:
	;
	v919 = F_slice_del(m, l0)
	mBase = m.M
	if v919 < int32(0) {
		v1243 = v919
		goto L1
	} else {
		goto L273
	}
L273:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v922
	v924 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v922-int32(3) <= v924 {
		goto L230
	} else {
		goto L274
	}
L274:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v932 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v928+v922-int32(1)))))
	switch v932 - int32(101) {
	case 0, 7:
		goto L275
	default:
		goto L230
	}
L275:
	;
	v938 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_12), int32(3), int32(0))
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L10
	} else {
		goto L276
	}
L276:
	;
	if v938 == int32(0) {
		goto L230
	} else {
		goto L277
	}
L277:
	;
	v942 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v942
	v944 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v942 < v944 {
		goto L230
	} else {
		goto L278
	}
L278:
	;
	v946 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v946 {
		goto L230
	} else {
		goto L279
	}
L279:
	;
	v1243 = v946
	goto L1
L280:
	;
	v951 = F_slice_del(m, l0)
	mBase = m.M
	if v951 < int32(0) {
		v1243 = v951
		goto L1
	} else {
		goto L281
	}
L281:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v954
	v957 = v954 - int32(1)
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v957 <= v958 {
		goto L230
	} else {
		goto L282
	}
L282:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v962 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v960+v957))))
	if base.B2i32(v962&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v962)%32)&int32(_a_F_portuguese_ISO_8859_1_stem_13) == int32(0)) != 0 {
		goto L230
	} else {
		goto L283
	}
L283:
	;
	v977 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_14), int32(3), int32(0))
	mBase = m.M
	v978 = m.ExcPending
	if v978 != 0 {
		goto L10
	} else {
		goto L284
	}
L284:
	;
	if v977 == int32(0) {
		goto L230
	} else {
		goto L285
	}
L285:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v981
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v981 < v983 {
		goto L230
	} else {
		goto L286
	}
L286:
	;
	v985 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v985 {
		goto L230
	} else {
		goto L287
	}
L287:
	;
	v1243 = v985
	goto L1
L288:
	;
	v990 = F_slice_del(m, l0)
	mBase = m.M
	if v990 < int32(0) {
		v1243 = v990
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v993
	v995 = int32(2)
	v997 = int32(0)
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v993-v1000 < v995 {
		v1010 = v997
		goto L291
	} else {
		goto L292
	}
L290:
	;
	if v1010 == int32(0) {
		goto L230
	} else {
		goto L294
	}
L291:
	;
	goto L290
L292:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1006 = F_memcmp(m, v1003+v993-v995, int32(_a_F_portuguese_ISO_8859_1_stem_15), v995)
	mBase = m.M
	if v1006 != 0 {
		v1010 = v997
		goto L291
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v993 - v995
	v1010 = int32(1)
	goto L291
L294:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1013
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v1013 < v1015 {
		goto L230
	} else {
		goto L295
	}
L295:
	;
	v1017 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1017 {
		goto L230
	} else {
		goto L296
	}
L296:
	;
	v1243 = v1017
	goto L1
L297:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v816 <= v1022 {
		goto L231
	} else {
		goto L298
	}
L298:
	;
	v1024 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1028 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1024+v816-int32(1)))))
	if v1028 != int32(101) {
		goto L231
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v816 - int32(1)
	v1036 = F_slice_from_s(m, l0, int32(2), int32(_a_F_portuguese_ISO_8859_1_stem_16))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L10
	} else {
		goto L300
	}
L300:
	;
	if int32(0) <= v1036 {
		goto L230
	} else {
		goto L301
	}
L301:
	;
	v1243 = v1036
	goto L1
L302:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1074
	v1076 = F_slice_del(m, l0)
	mBase = m.M
	if v1076 < int32(0) {
		v1243 = v1076
		goto L1
	} else {
		goto L312
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1042
	v1047 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1044
	v1052 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_17), int32(120), int32(0))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L10
	} else {
		goto L306
	}
L304:
	;
	v1056 = v1042
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1056
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1056
	v1063 = F_find_among_b(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_18), int32(7), int32(0))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L10
	} else {
		goto L308
	}
L306:
	;
	if v1052 != 0 {
		goto L302
	} else {
		goto L307
	}
L307:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1047
	v1055 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1056 = v1055
	goto L305
L308:
	;
	if v1063 == int32(0) {
		goto L229
	} else {
		goto L309
	}
L309:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1067
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1067 < v1069 {
		goto L229
	} else {
		goto L310
	}
L310:
	;
	v1071 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1071 {
		goto L229
	} else {
		goto L311
	}
L311:
	;
	v1243 = v1071
	goto L1
L312:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v1047
	goto L230
L313:
	;
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1089 = v1088 + v1083
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089-int32(1)))))
	if v1092 != int32(105) {
		goto L229
	} else {
		goto L314
	}
L314:
	;
	v1096 = v1083 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1096
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1096
	if v1096 <= v1086 {
		goto L229
	} else {
		goto L315
	}
L315:
	;
	v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1089-int32(2)))))
	if v1102 != int32(99) {
		goto L229
	} else {
		goto L316
	}
L316:
	;
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1083 <= v1105 {
		goto L229
	} else {
		goto L317
	}
L317:
	;
	v1107 = F_slice_del(m, l0)
	mBase = m.M
	if v1107 < int32(0) {
		v1243 = v1107
		goto L1
	} else {
		goto L318
	}
L318:
	;
	goto L229
L319:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1187
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1191 = v1187
	v1194 = v1189
	goto L339
L320:
	;
	if v1120 == int32(0) {
		goto L319
	} else {
		goto L321
	}
L321:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1124
	switch v1120 - int32(1) {
	case 0:
		goto L323
	case 1:
		goto L322
	default:
		goto L319
	}
L322:
	;
	v1178 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_19))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L10
	} else {
		goto L337
	}
L323:
	;
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1124 < v1128 {
		goto L319
	} else {
		goto L324
	}
L324:
	;
	v1130 = F_slice_del(m, l0)
	mBase = m.M
	if v1130 < int32(0) {
		v1243 = v1130
		goto L1
	} else {
		goto L325
	}
L325:
	;
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1133
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1133 <= v1135 {
		goto L319
	} else {
		goto L326
	}
L326:
	;
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1138 = v1137 + v1133
	v1140 = v1138 - int32(1)
	v1141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140))))
	if v1141 != int32(117) {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1169
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v1169 < v1171 {
		goto L319
	} else {
		goto L335
	}
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1133
	v1156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1140))))
	if v1156 != int32(105) {
		goto L319
	} else {
		goto L332
	}
L329:
	;
	v1145 = v1133 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1145
	if v1145 <= v1135 {
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138-int32(2)))))
	if v1151 == int32(103) {
		v1169 = v1145
		goto L327
	} else {
		goto L331
	}
L331:
	;
	goto L328
L332:
	;
	v1160 = v1133 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1160
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1160
	if v1160 <= v1135 {
		goto L319
	} else {
		goto L333
	}
L333:
	;
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138-int32(2)))))
	if v1166 != int32(99) {
		goto L319
	} else {
		goto L334
	}
L334:
	;
	v1169 = v1160
	goto L327
L335:
	;
	v1173 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v1173 {
		goto L319
	} else {
		goto L336
	}
L336:
	;
	v1243 = v1173
	goto L1
L337:
	;
	if v1178 < int32(0) {
		v1243 = v1178
		goto L1
	} else {
		goto L338
	}
L338:
	;
	goto L319
L339:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1191
	v1198 = v1191 + int32(1)
	if v1198 < v1194 {
		goto L345
	} else {
		goto L346
	}
L340:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1187
	v1243 = int32(1)
	goto L1
L341:
	;
	goto L340
L342:
	;
	v1238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1191 = v1239
	v1194 = v1238
	goto L339
L343:
	;
	if v1230 <= v1228 {
		goto L341
	} else {
		goto L357
	}
L344:
	;
	v1209 = F_find_among(m, l0, int32(_a_F_portuguese_ISO_8859_1_stem_20), int32(3), int32(0))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L10
	} else {
		goto L349
	}
L345:
	;
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1200+v1198))))
	if v1202 == int32(126) {
		goto L344
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1191
	v1228 = v1191
	v1230 = v1194
	goto L343
L348:
	;
	goto L347
L349:
	;
	v1211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1211
	switch v1209 - int32(1) {
	case 0:
		goto L352
	case 1:
		goto L351
	case 2:
		goto L350
	default:
		goto L342
	}
L350:
	;
	v1227 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1228 = v1211
	v1230 = v1227
	goto L343
L351:
	;
	v1223 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_21))
	mBase = m.M
	v1224 = m.ExcPending
	if v1224 != 0 {
		goto L10
	} else {
		goto L355
	}
L352:
	;
	v1217 = F_slice_from_s(m, l0, int32(1), int32(_a_F_portuguese_ISO_8859_1_stem_22))
	mBase = m.M
	v1218 = m.ExcPending
	if v1218 != 0 {
		goto L10
	} else {
		goto L353
	}
L353:
	;
	if int32(0) <= v1217 {
		goto L342
	} else {
		goto L354
	}
L354:
	;
	v1243 = v1217
	goto L1
L355:
	;
	if int32(0) <= v1223 {
		goto L342
	} else {
		goto L356
	}
L356:
	;
	v1243 = v1223
	goto L1
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1228 + int32(1)
	goto L342
}
func F_pow(m *base.Module, l0 float64, l1 float64) float64 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v48 float64
	_ = v48
	var v52 int64
	_ = v52
	var v58 int64
	_ = v58
	var v74 float64
	_ = v74
	var v81 float64
	_ = v81
	var v92 int32
	_ = v92
	var v99 int64
	_ = v99
	var v103 int64
	_ = v103
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 float64
	_ = v119
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 float64
	_ = v167
	var v173 int32
	_ = v173
	var v179 int64
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v192 float64
	_ = v192
	var v202 float64
	_ = v202
	var v205 float64
	_ = v205
	var v213 int64
	_ = v213
	var v215 int32
	_ = v215
	var v216 int64
	_ = v216
	var v218 float64
	_ = v218
	var v220 int64
	_ = v220
	var v221 int64
	_ = v221
	var v223 float64
	_ = v223
	var v225 float64
	_ = v225
	var v233 int32
	_ = v233
	var v234 float64
	_ = v234
	var v238 int64
	_ = v238
	var v243 float64
	_ = v243
	var v244 float64
	_ = v244
	var v247 float64
	_ = v247
	var v250 float64
	_ = v250
	var v251 float64
	_ = v251
	var v253 float64
	_ = v253
	var v255 float64
	_ = v255
	var v256 float64
	_ = v256
	var v257 float64
	_ = v257
	var v262 float64
	_ = v262
	var v263 float64
	_ = v263
	var v264 float64
	_ = v264
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v273 float64
	_ = v273
	var v276 float64
	_ = v276
	var v279 float64
	_ = v279
	var v283 float64
	_ = v283
	var v286 float64
	_ = v286
	var v291 float64
	_ = v291
	var v294 float64
	_ = v294
	var v298 float64
	_ = v298
	var v299 float64
	_ = v299
	var v301 float64
	_ = v301
	var v306 float64
	_ = v306
	var v307 float64
	_ = v307
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v336 float64
	_ = v336
	var v338 float64
	_ = v338
	var v350 float64
	_ = v350
	var v352 float64
	_ = v352
	var v353 int32
	_ = v353
	var v355 float64
	_ = v355
	var v358 float64
	_ = v358
	var v359 float64
	_ = v359
	var v360 float64
	_ = v360
	var v362 float64
	_ = v362
	var v365 float64
	_ = v365
	var v369 float64
	_ = v369
	var v370 float64
	_ = v370
	var v373 float64
	_ = v373
	var v376 float64
	_ = v376
	var v380 float64
	_ = v380
	var v383 float64
	_ = v383
	var v386 int64
	_ = v386
	var v391 int32
	_ = v391
	var v392 float64
	_ = v392
	var v395 float64
	_ = v395
	var v396 int64
	_ = v396
	var v401 int64
	_ = v401
	var v410 float64
	_ = v410
	var v416 int64
	_ = v416
	var v417 float64
	_ = v417
	var v418 float64
	_ = v418
	var v419 float64
	_ = v419
	var v423 float64
	_ = v423
	var v425 int32
	_ = v425
	var v432 int32
	_ = v432
	var v443 float64
	_ = v443
	var v444 float64
	_ = v444
	var v451 float64
	_ = v451
	var v454 float64
	_ = v454
	var v458 float64
	_ = v458
	var v467 float64
	_ = v467
	var v468 float64
	_ = v468
	var v480 float64
	_ = v480
	var v483 float64
	_ = v483
	v11 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v24 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(52)) % 64)))
	v28 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l1)) >> (uint(int64(52)) % 64)))
	v29 = int32(2047)
	v30 = v28 & v29
	v32 = v30 - int32(1086)
	v33 = base.I64_reinterpret_f64(l1)
	v34 = base.I64_reinterpret_f64(l0)
	if base.B2i32(base.Ui32(int32(-129)) < base.Ui32(v32))&base.B2i32(base.Ui32(int32(-2046)) <= base.Ui32(v24-v29)) != 0 {
		v213 = v34
		v215 = v11
		v216 = int64(-134217728)
		v218 = base.F64_reinterpret_i64(v33 & v216)
		v220 = v213 - int64(4604531861337669632)
		v221 = int64(52)
		v223 = base.F64_convert_i64_s(v220 >> (uint(v221) % 64))
		v225 = *(*float64)(unsafe.Add(mBase, _c_F_pow[0]))
		v233 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
		v234 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[1])))
		v238 = v213 - v220&int64(-4503599627370496)
		v243 = base.F64_reinterpret_i64((v238 + int64(2147483648)) & int64(-4294967296))
		v244 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[2])))
		v247 = base.F64_add(base.F64_mul(v243, v244), float64(-1))
		v250 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v238), v243), v244)
		v251 = base.F64_add(v247, v250)
		v253 = *(*float64)(unsafe.Add(mBase, _c_F_pow[3]))
		v255 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[4])))
		v256 = base.F64_add(base.F64_mul(v223, v253), v255)
		v257 = base.F64_add(v251, v256)
		v262 = *(*float64)(unsafe.Add(mBase, _c_F_pow[5]))
		v263 = base.F64_mul(v251, v262)
		v264 = base.F64_mul(v247, v262)
		v268 = base.F64_mul(v247, v264)
		v269 = base.F64_add(v257, v268)
		v273 = base.F64_mul(v251, v263)
		v276 = *(*float64)(unsafe.Add(mBase, _c_F_pow[6]))
		v279 = *(*float64)(unsafe.Add(mBase, _c_F_pow[7]))
		v283 = *(*float64)(unsafe.Add(mBase, _c_F_pow[8]))
		v286 = *(*float64)(unsafe.Add(mBase, _c_F_pow[9]))
		v291 = *(*float64)(unsafe.Add(mBase, _c_F_pow[10]))
		v294 = *(*float64)(unsafe.Add(mBase, _c_F_pow[11]))
		v298 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v234), base.F64_add(v251, base.F64_sub(v256, v257))), base.F64_mul(v250, base.F64_add(v263, v264))), base.F64_add(v268, base.F64_sub(v257, v269))), base.F64_mul(base.F64_mul(v251, v273), base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v251, v276), v279)), base.F64_add(base.F64_mul(v251, v283), v286))), base.F64_add(base.F64_mul(v251, v291), v294))))
		v299 = base.F64_add(v269, v298)
		v301 = base.F64_add(v298, base.F64_sub(v269, v299))
		*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v301
		v306 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v299) & v216)
		v307 = base.F64_mul(v218, v306)
		v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(v221)%64))) & int32(2047)
		v325 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
		if base.Ui32(v320-v325) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v325) {
			v353 = v320
			v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
			v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
			v359 = base.F64_add(base.F64_mul(v307, v355), v358)
			v360 = base.F64_sub(v359, v358)
			v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
			v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
			v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
			v370 = base.F64_mul(v369, v369)
			v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
			v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
			v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
			v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
			v386 = base.I64_reinterpret_f64(v359)
			v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
			v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
			v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
			v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
			v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
			if v353 == int32(0) {
				if v386&int64(2147483648) == int64(0) {
					v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
					v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
				} else {
					v416 = v401 + int64(4602678819172646912)
					v417 = base.F64_reinterpret_i64(v416)
					v418 = base.F64_mul(v417, v395)
					v419 = base.F64_add(v418, v417)
					if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
						v423 = float64(2.2250738585072014e-308)
						v425 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
						v432 = m.G0
						*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
						if base.F64_lt(v419, float64(0)) != 0 {
							v443 = float64(-1)
						} else {
							v443 = float64(1)
						}
						v444 = base.F64_add(v419, v443)
						v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
						if base.F64_eq(v451, float64(0)) != 0 {
							v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
						} else {
							v454 = v451
						}
						v458 = v454
					} else {
						v458 = v419
					}
					v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
				}
				v480 = v467
			} else {
				v468 = base.F64_reinterpret_i64(v401)
				v480 = base.F64_add(base.F64_mul(v468, v395), v468)
			}
		} else {
			if base.Ui32(v320) < base.Ui32(v325) {
				v336 = base.F64_add(v307, float64(1))
				if v215 != 0 {
					v338 = base.F64_neg(v336)
				} else {
					v338 = v336
				}
				v480 = v338
			} else {
				if base.Ui32(v320) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
					v353 = int32(0)
					v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
					v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
					v359 = base.F64_add(base.F64_mul(v307, v355), v358)
					v360 = base.F64_sub(v359, v358)
					v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
					v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
					v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
					v370 = base.F64_mul(v369, v369)
					v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
					v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
					v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
					v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
					v386 = base.I64_reinterpret_f64(v359)
					v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
					v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
					v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
					v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
					v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
					if v353 == int32(0) {
						if v386&int64(2147483648) == int64(0) {
							v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
							v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
						} else {
							v416 = v401 + int64(4602678819172646912)
							v417 = base.F64_reinterpret_i64(v416)
							v418 = base.F64_mul(v417, v395)
							v419 = base.F64_add(v418, v417)
							if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
								v423 = float64(2.2250738585072014e-308)
								v425 = m.G0
								*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
								v432 = m.G0
								*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
								if base.F64_lt(v419, float64(0)) != 0 {
									v443 = float64(-1)
								} else {
									v443 = float64(1)
								}
								v444 = base.F64_add(v419, v443)
								v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
								if base.F64_eq(v451, float64(0)) != 0 {
									v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
								} else {
									v454 = v451
								}
								v458 = v454
							} else {
								v458 = v419
							}
							v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
						}
						v480 = v467
					} else {
						v468 = base.F64_reinterpret_i64(v401)
						v480 = base.F64_add(base.F64_mul(v468, v395), v468)
					}
				} else {
					if base.I64_reinterpret_f64(v307) < int64(0) {
						v350 = F___math_xflow(m, v215, float64(1.2882297539194267e-231))
						mBase = m.M
						v480 = v350
					} else {
						v352 = F___math_xflow(m, v215, float64(3.105036184601418e+231))
						mBase = m.M
						v480 = v352
					}
				}
			}
		}
		v483 = v480
	} else {
		if base.Ui64(v33<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)) {
			v48 = float64(1)
			if v34 == int64(4607182418800017408) {
				v483 = v48
			} else {
				v52 = v33 << (uint(int64(1)) % 64)
				if v52 == int64(0) {
					v483 = v48
				} else {
					v58 = v34 << (uint(int64(1)) % 64)
					if base.B2i32(base.Ui64(v52) < base.Ui64(int64(-9007199254740991)))&base.B2i32(base.Ui64(v58) <= base.Ui64(int64(-9007199254740992))) == int32(0) {
						v483 = base.F64_add(l0, l1)
					} else {
						if v58 == int64(9214364837600034816) {
							v483 = v48
						} else {
							if base.B2i32(v33 < int64(0))^base.B2i32(base.Ui64(v58) < base.Ui64(int64(9214364837600034816))) != 0 {
								v74 = float64(0)
							} else {
								v74 = base.F64_mul(l1, l1)
							}
							v483 = v74
						}
					}
				}
			}
		} else {
			if base.Ui64(v34<<(uint(int64(1))%64)+int64(9007199254740992)) < base.Ui64(int64(9007199254740993)) {
				v81 = base.F64_mul(l0, l0)
				if v34 < int64(0) {
					v92 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v92) < base.Ui32(int32(1023)) {
						v116 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v92) {
							v116 = int32(2)
						} else {
							v99 = int64(1)
							v103 = v99 << (uint(base.I64_extend_i32_u(int32(1075)-v92)) % 64)
							if (v103-v99)&v33 != int64(0) {
								v116 = int32(0)
							} else {
								if v33&v103 == int64(0) {
									v114 = int32(2)
								} else {
									v114 = int32(1)
								}
								v116 = v114
							}
						}
					}
					if v116 == int32(1) {
						v119 = base.F64_neg(v81)
					} else {
						v119 = v81
					}
					v120 = v119
				} else {
					v120 = v81
				}
				if int64(0) <= v33 {
					v483 = v120
				} else {
					v124 = base.F64_div(float64(1), v120)
					v126 = m.G0
					*(*float64)(unsafe.Add(mBase, uint32(v126-int32(16))+8)) = v124
					v483 = v124
				}
			} else {
				if v34 < int64(0) {
					v140 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(52))%64))) & int32(2047)
					if base.Ui32(v140) < base.Ui32(int32(1023)) {
						v164 = int32(0)
					} else {
						if base.Ui32(int32(1075)) < base.Ui32(v140) {
							v164 = int32(2)
						} else {
							v147 = int64(1)
							v151 = v147 << (uint(base.I64_extend_i32_u(int32(1075)-v140)) % 64)
							if (v151-v147)&v33 != int64(0) {
								v164 = int32(0)
							} else {
								if v33&v151 == int64(0) {
									v162 = int32(2)
								} else {
									v162 = int32(1)
								}
								v164 = v162
							}
						}
					}
					if v164 == int32(0) {
						v167 = base.F64_sub(l0, l0)
						v483 = base.F64_div(v167, v167)
					} else {
						if v164 == int32(1) {
							v173 = int32(_a_F_pow_0)
						} else {
							v173 = int32(0)
						}
						v179 = base.I64_reinterpret_f64(l0) & int64(9223372036854775807)
						v180 = v24 & int32(2047)
						v181 = v173
						if base.Ui32(v32) <= base.Ui32(int32(-129)) {
							if v179 == int64(4607182418800017408) {
								v483 = float64(1)
							} else {
								if base.Ui32(v30) <= base.Ui32(int32(957)) {
									if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
										v192 = l1
									} else {
										v192 = base.F64_neg(l1)
									}
									v483 = base.F64_add(v192, float64(1))
								} else {
									if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v28)) != base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
										v202 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
										mBase = m.M
										v483 = v202
									} else {
										v205 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
										mBase = m.M
										v483 = v205
									}
								}
							}
						} else {
							if v180 != 0 {
								v213 = v179
								v215 = v181
							} else {
								v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) - int64(234187180623265792)
								v215 = v181
							}
							v216 = int64(-134217728)
							v218 = base.F64_reinterpret_i64(v33 & v216)
							v220 = v213 - int64(4604531861337669632)
							v221 = int64(52)
							v223 = base.F64_convert_i64_s(v220 >> (uint(v221) % 64))
							v225 = *(*float64)(unsafe.Add(mBase, _c_F_pow[0]))
							v233 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
							v234 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[1])))
							v238 = v213 - v220&int64(-4503599627370496)
							v243 = base.F64_reinterpret_i64((v238 + int64(2147483648)) & int64(-4294967296))
							v244 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[2])))
							v247 = base.F64_add(base.F64_mul(v243, v244), float64(-1))
							v250 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v238), v243), v244)
							v251 = base.F64_add(v247, v250)
							v253 = *(*float64)(unsafe.Add(mBase, _c_F_pow[3]))
							v255 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[4])))
							v256 = base.F64_add(base.F64_mul(v223, v253), v255)
							v257 = base.F64_add(v251, v256)
							v262 = *(*float64)(unsafe.Add(mBase, _c_F_pow[5]))
							v263 = base.F64_mul(v251, v262)
							v264 = base.F64_mul(v247, v262)
							v268 = base.F64_mul(v247, v264)
							v269 = base.F64_add(v257, v268)
							v273 = base.F64_mul(v251, v263)
							v276 = *(*float64)(unsafe.Add(mBase, _c_F_pow[6]))
							v279 = *(*float64)(unsafe.Add(mBase, _c_F_pow[7]))
							v283 = *(*float64)(unsafe.Add(mBase, _c_F_pow[8]))
							v286 = *(*float64)(unsafe.Add(mBase, _c_F_pow[9]))
							v291 = *(*float64)(unsafe.Add(mBase, _c_F_pow[10]))
							v294 = *(*float64)(unsafe.Add(mBase, _c_F_pow[11]))
							v298 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v234), base.F64_add(v251, base.F64_sub(v256, v257))), base.F64_mul(v250, base.F64_add(v263, v264))), base.F64_add(v268, base.F64_sub(v257, v269))), base.F64_mul(base.F64_mul(v251, v273), base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v251, v276), v279)), base.F64_add(base.F64_mul(v251, v283), v286))), base.F64_add(base.F64_mul(v251, v291), v294))))
							v299 = base.F64_add(v269, v298)
							v301 = base.F64_add(v298, base.F64_sub(v269, v299))
							*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v301
							v306 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v299) & v216)
							v307 = base.F64_mul(v218, v306)
							v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(v221)%64))) & int32(2047)
							v325 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
							if base.Ui32(v320-v325) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v325) {
								v353 = v320
								v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
								v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
								v359 = base.F64_add(base.F64_mul(v307, v355), v358)
								v360 = base.F64_sub(v359, v358)
								v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
								v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
								v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
								v370 = base.F64_mul(v369, v369)
								v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
								v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
								v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
								v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
								v386 = base.I64_reinterpret_f64(v359)
								v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
								v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
								v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
								v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
								v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
								if v353 == int32(0) {
									if v386&int64(2147483648) == int64(0) {
										v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
										v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
									} else {
										v416 = v401 + int64(4602678819172646912)
										v417 = base.F64_reinterpret_i64(v416)
										v418 = base.F64_mul(v417, v395)
										v419 = base.F64_add(v418, v417)
										if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
											v423 = float64(2.2250738585072014e-308)
											v425 = m.G0
											*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
											v432 = m.G0
											*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
											if base.F64_lt(v419, float64(0)) != 0 {
												v443 = float64(-1)
											} else {
												v443 = float64(1)
											}
											v444 = base.F64_add(v419, v443)
											v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
											if base.F64_eq(v451, float64(0)) != 0 {
												v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
											} else {
												v454 = v451
											}
											v458 = v454
										} else {
											v458 = v419
										}
										v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
									}
									v480 = v467
								} else {
									v468 = base.F64_reinterpret_i64(v401)
									v480 = base.F64_add(base.F64_mul(v468, v395), v468)
								}
							} else {
								if base.Ui32(v320) < base.Ui32(v325) {
									v336 = base.F64_add(v307, float64(1))
									if v215 != 0 {
										v338 = base.F64_neg(v336)
									} else {
										v338 = v336
									}
									v480 = v338
								} else {
									if base.Ui32(v320) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
										v353 = int32(0)
										v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
										v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
										v359 = base.F64_add(base.F64_mul(v307, v355), v358)
										v360 = base.F64_sub(v359, v358)
										v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
										v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
										v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
										v370 = base.F64_mul(v369, v369)
										v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
										v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
										v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
										v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
										v386 = base.I64_reinterpret_f64(v359)
										v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
										v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
										v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
										v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
										v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
										if v353 == int32(0) {
											if v386&int64(2147483648) == int64(0) {
												v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
												v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
											} else {
												v416 = v401 + int64(4602678819172646912)
												v417 = base.F64_reinterpret_i64(v416)
												v418 = base.F64_mul(v417, v395)
												v419 = base.F64_add(v418, v417)
												if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
													v423 = float64(2.2250738585072014e-308)
													v425 = m.G0
													*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
													v432 = m.G0
													*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
													if base.F64_lt(v419, float64(0)) != 0 {
														v443 = float64(-1)
													} else {
														v443 = float64(1)
													}
													v444 = base.F64_add(v419, v443)
													v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
													if base.F64_eq(v451, float64(0)) != 0 {
														v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
													} else {
														v454 = v451
													}
													v458 = v454
												} else {
													v458 = v419
												}
												v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
											}
											v480 = v467
										} else {
											v468 = base.F64_reinterpret_i64(v401)
											v480 = base.F64_add(base.F64_mul(v468, v395), v468)
										}
									} else {
										if base.I64_reinterpret_f64(v307) < int64(0) {
											v350 = F___math_xflow(m, v215, float64(1.2882297539194267e-231))
											mBase = m.M
											v480 = v350
										} else {
											v352 = F___math_xflow(m, v215, float64(3.105036184601418e+231))
											mBase = m.M
											v480 = v352
										}
									}
								}
							}
							v483 = v480
						}
					}
				} else {
					v179 = v34
					v180 = v24
					v181 = v11
					if base.Ui32(v32) <= base.Ui32(int32(-129)) {
						if v179 == int64(4607182418800017408) {
							v483 = float64(1)
						} else {
							if base.Ui32(v30) <= base.Ui32(int32(957)) {
								if base.Ui64(int64(4607182418800017408)) < base.Ui64(v179) {
									v192 = l1
								} else {
									v192 = base.F64_neg(l1)
								}
								v483 = base.F64_add(v192, float64(1))
							} else {
								if base.B2i32(base.Ui32(int32(2047)) < base.Ui32(v28)) != base.B2i32(base.Ui64(int64(4607182418800017408)) < base.Ui64(v179)) {
									v202 = F___math_xflow(m, int32(0), float64(3.105036184601418e+231))
									mBase = m.M
									v483 = v202
								} else {
									v205 = F___math_xflow(m, int32(0), float64(1.2882297539194267e-231))
									mBase = m.M
									v483 = v205
								}
							}
						}
					} else {
						if v180 != 0 {
							v213 = v179
							v215 = v181
						} else {
							v213 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15)))&int64(9223372036854775807) - int64(234187180623265792)
							v215 = v181
						}
						v216 = int64(-134217728)
						v218 = base.F64_reinterpret_i64(v33 & v216)
						v220 = v213 - int64(4604531861337669632)
						v221 = int64(52)
						v223 = base.F64_convert_i64_s(v220 >> (uint(v221) % 64))
						v225 = *(*float64)(unsafe.Add(mBase, _c_F_pow[0]))
						v233 = base.I32_wrap_i64(int64(base.Ui64(v220)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(5)) % 32)
						v234 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[1])))
						v238 = v213 - v220&int64(-4503599627370496)
						v243 = base.F64_reinterpret_i64((v238 + int64(2147483648)) & int64(-4294967296))
						v244 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[2])))
						v247 = base.F64_add(base.F64_mul(v243, v244), float64(-1))
						v250 = base.F64_mul(base.F64_sub(base.F64_reinterpret_i64(v238), v243), v244)
						v251 = base.F64_add(v247, v250)
						v253 = *(*float64)(unsafe.Add(mBase, _c_F_pow[3]))
						v255 = *(*float64)(unsafe.Add(mBase, uint32(v233)+uint32(_c_F_pow[4])))
						v256 = base.F64_add(base.F64_mul(v223, v253), v255)
						v257 = base.F64_add(v251, v256)
						v262 = *(*float64)(unsafe.Add(mBase, _c_F_pow[5]))
						v263 = base.F64_mul(v251, v262)
						v264 = base.F64_mul(v247, v262)
						v268 = base.F64_mul(v247, v264)
						v269 = base.F64_add(v257, v268)
						v273 = base.F64_mul(v251, v263)
						v276 = *(*float64)(unsafe.Add(mBase, _c_F_pow[6]))
						v279 = *(*float64)(unsafe.Add(mBase, _c_F_pow[7]))
						v283 = *(*float64)(unsafe.Add(mBase, _c_F_pow[8]))
						v286 = *(*float64)(unsafe.Add(mBase, _c_F_pow[9]))
						v291 = *(*float64)(unsafe.Add(mBase, _c_F_pow[10]))
						v294 = *(*float64)(unsafe.Add(mBase, _c_F_pow[11]))
						v298 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_add(base.F64_mul(v223, v225), v234), base.F64_add(v251, base.F64_sub(v256, v257))), base.F64_mul(v250, base.F64_add(v263, v264))), base.F64_add(v268, base.F64_sub(v257, v269))), base.F64_mul(base.F64_mul(v251, v273), base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v273, base.F64_add(base.F64_mul(v251, v276), v279)), base.F64_add(base.F64_mul(v251, v283), v286))), base.F64_add(base.F64_mul(v251, v291), v294))))
						v299 = base.F64_add(v269, v298)
						v301 = base.F64_add(v298, base.F64_sub(v269, v299))
						*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = v301
						v306 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(v299) & v216)
						v307 = base.F64_mul(v218, v306)
						v320 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v307))>>(uint(v221)%64))) & int32(2047)
						v325 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(5.551115123125783e-17))) >> (uint(int64(52)) % 64)))
						if base.Ui32(v320-v325) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(512)))>>(uint(int64(52))%64)))-v325) {
							v353 = v320
							v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
							v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
							v359 = base.F64_add(base.F64_mul(v307, v355), v358)
							v360 = base.F64_sub(v359, v358)
							v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
							v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
							v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
							v370 = base.F64_mul(v369, v369)
							v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
							v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
							v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
							v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
							v386 = base.I64_reinterpret_f64(v359)
							v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
							v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
							v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
							v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
							v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
							if v353 == int32(0) {
								if v386&int64(2147483648) == int64(0) {
									v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
									v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
								} else {
									v416 = v401 + int64(4602678819172646912)
									v417 = base.F64_reinterpret_i64(v416)
									v418 = base.F64_mul(v417, v395)
									v419 = base.F64_add(v418, v417)
									if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
										v423 = float64(2.2250738585072014e-308)
										v425 = m.G0
										*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
										v432 = m.G0
										*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
										if base.F64_lt(v419, float64(0)) != 0 {
											v443 = float64(-1)
										} else {
											v443 = float64(1)
										}
										v444 = base.F64_add(v419, v443)
										v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
										if base.F64_eq(v451, float64(0)) != 0 {
											v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
										} else {
											v454 = v451
										}
										v458 = v454
									} else {
										v458 = v419
									}
									v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
								}
								v480 = v467
							} else {
								v468 = base.F64_reinterpret_i64(v401)
								v480 = base.F64_add(base.F64_mul(v468, v395), v468)
							}
						} else {
							if base.Ui32(v320) < base.Ui32(v325) {
								v336 = base.F64_add(v307, float64(1))
								if v215 != 0 {
									v338 = base.F64_neg(v336)
								} else {
									v338 = v336
								}
								v480 = v338
							} else {
								if base.Ui32(v320) < base.Ui32(base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(float64(1024)))>>(uint(int64(52))%64)))) {
									v353 = int32(0)
									v355 = *(*float64)(unsafe.Add(mBase, _c_F_pow[12]))
									v358 = *(*float64)(unsafe.Add(mBase, _c_F_pow[13]))
									v359 = base.F64_add(base.F64_mul(v307, v355), v358)
									v360 = base.F64_sub(v359, v358)
									v362 = *(*float64)(unsafe.Add(mBase, _c_F_pow[14]))
									v365 = *(*float64)(unsafe.Add(mBase, _c_F_pow[15]))
									v369 = base.F64_add(base.F64_add(base.F64_mul(base.F64_sub(l1, v218), v306), base.F64_mul(l1, base.F64_add(v301, base.F64_sub(v299, v306)))), base.F64_add(base.F64_mul(v360, v362), base.F64_add(base.F64_mul(v360, v365), v307)))
									v370 = base.F64_mul(v369, v369)
									v373 = *(*float64)(unsafe.Add(mBase, _c_F_pow[16]))
									v376 = *(*float64)(unsafe.Add(mBase, _c_F_pow[17]))
									v380 = *(*float64)(unsafe.Add(mBase, _c_F_pow[18]))
									v383 = *(*float64)(unsafe.Add(mBase, _c_F_pow[19]))
									v386 = base.I64_reinterpret_f64(v359)
									v391 = base.I32_wrap_i64(v386) << (uint(int32(4)) % 32) & int32(2032)
									v392 = *(*float64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[20])))
									v395 = base.F64_add(base.F64_mul(base.F64_mul(v370, v370), base.F64_add(base.F64_mul(v369, v373), v376)), base.F64_add(base.F64_mul(v370, base.F64_add(base.F64_mul(v369, v380), v383)), base.F64_add(v392, v369)))
									v396 = *(*int64)(unsafe.Add(mBase, uint32(v391)+uint32(_c_F_pow[21])))
									v401 = v396 + (v386+base.I64_extend_i32_u(v215))<<(uint(int64(45))%64)
									if v353 == int32(0) {
										if v386&int64(2147483648) == int64(0) {
											v410 = base.F64_reinterpret_i64(v401 - int64(4544132024016830464))
											v467 = base.F64_mul(base.F64_add(base.F64_mul(v410, v395), v410), float64(5.486124068793689e+303))
										} else {
											v416 = v401 + int64(4602678819172646912)
											v417 = base.F64_reinterpret_i64(v416)
											v418 = base.F64_mul(v417, v395)
											v419 = base.F64_add(v418, v417)
											if base.F64_lt(base.F64_abs(v419), float64(1)) != 0 {
												v423 = float64(2.2250738585072014e-308)
												v425 = m.G0
												*(*float64)(unsafe.Add(mBase, uint32(v425-int32(16))+8)) = v423
												v432 = m.G0
												*(*float64)(unsafe.Add(mBase, uint32(v432-int32(16))+8)) = base.F64_mul(v423, float64(2.2250738585072014e-308))
												if base.F64_lt(v419, float64(0)) != 0 {
													v443 = float64(-1)
												} else {
													v443 = float64(1)
												}
												v444 = base.F64_add(v419, v443)
												v451 = base.F64_sub(base.F64_add(v444, base.F64_add(base.F64_add(v418, base.F64_sub(v417, v419)), base.F64_add(v419, base.F64_sub(v443, v444)))), v443)
												if base.F64_eq(v451, float64(0)) != 0 {
													v454 = base.F64_reinterpret_i64(v416 & int64(-9223372036854775807-1))
												} else {
													v454 = v451
												}
												v458 = v454
											} else {
												v458 = v419
											}
											v467 = base.F64_mul(v458, float64(2.2250738585072014e-308))
										}
										v480 = v467
									} else {
										v468 = base.F64_reinterpret_i64(v401)
										v480 = base.F64_add(base.F64_mul(v468, v395), v468)
									}
								} else {
									if base.I64_reinterpret_f64(v307) < int64(0) {
										v350 = F___math_xflow(m, v215, float64(1.2882297539194267e-231))
										mBase = m.M
										v480 = v350
									} else {
										v352 = F___math_xflow(m, v215, float64(3.105036184601418e+231))
										mBase = m.M
										v480 = v352
									}
								}
							}
						}
						v483 = v480
					}
				}
			}
		}
	}
	m.G0 = v19 + int32(16)
	return v483
}
func F_pread(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
	v17 = m.Wasi_snapshot_preview1.Fd_pread(m, l0, v8+int32(8), int32(1), l3, v8+int32(4))
	mBase = m.M
	if v17 == int32(0) {
		v24 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _c_F_pread[0])) = v17
		v24 = int32(-1)
	}
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	m.G0 = v8 + int32(16)
	if v24 != 0 {
		v30 = int32(-1)
	} else {
		v30 = v25
	}
	return v30
}
func F_printtup(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
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
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v16 == v18 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v132 < v131 {
		goto L30
	} else {
		goto L31
	}
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v20 == v17 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+96))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v24 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L4
L6:
	;
	F_pfree(m, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v16
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v31
	if v17 <= v31 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	return int32(0)
L10:
	;
	goto L8
L11:
	;
	v37 = F_palloc0(m, v17*int32(40))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L9
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v37
	v47 = int32(0)
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v54 = v51 + v47*int32(40)
	if v23 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	goto L1
L15:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	F_fmgr_info(m, v111, v54+int32(12))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L9
	} else {
		goto L28
	}
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v98+v16+v47*int32(100))+96))
	F_getTypeOutputInfo(m, v104, v54, v54+int32(8))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L9
	} else {
		goto L27
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v58 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+10)) = uint16(v58)
	v98 = v57 << (uint(int32(3)) % 32)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v23+v47<<(uint(int32(1))%32)))))
	*(*uint16)(unsafe.Add(mBase, uint32(v54)+10)) = uint16(v66)
	v69 = v62 << (uint(int32(3)) % 32)
	switch v66 {
	case 0:
		v98 = v69
		goto L16
	case 1:
		goto L21
	default:
		goto L20
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L9
	} else {
		goto L23
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v69+v16+v47*int32(100))+96))
	v76 = v54 + int32(4)
	F_getTypeBinaryOutputInfo(m, v74, v76, v54+int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v109 = v76
	goto L15
L23:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L9
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = base.I32_extend16_s(v66)
	F_errmsg(m, int32(_a_F_printtup_0), v14)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L9
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_printtup_1), int32(293), int32(_a_F_printtup_2))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v109 = v54
	goto L15
L28:
	;
	v117 = v47 + int32(1)
	if v117 != v17 {
		v47 = v117
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+16))
	m.T0[v135].(func(*base.Module, int32, int32))(m, l0, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v138 = int32(_a_F_printtup_3)
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_printtup[0]))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, _c_F_printtup[0])) = v141
	v144 = l1 + int32(40)
	F_resetStringInfo(m, v144)
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = int32(68)
	goto L34
L33:
	;
	goto L32
L34:
	;
	F_enlargeStringInfo(m, v144, int32(2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L9
	} else {
		goto L35
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v154 = int32(8)
	v160 = v17<<(uint(v154)%32) | int32(base.Ui32(v17&int32(_a_F_printtup_4))>>(uint(v154)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v151+v152))) = uint16(v160)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v151 + int32(2)
	if int32(0) < v17 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v171 = int32(0)
	goto L39
L37:
	;
	goto L38
L38:
	;
	F_pq_endmessage_reuse(m, v144)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L9
	} else {
		goto L55
	}
L39:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179+v171))))
	if v181 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L38
L41:
	;
	v254 = v171 + int32(1)
	if v254 != v17 {
		v171 = v254
		goto L39
	} else {
		goto L54
	}
L42:
	;
	F_enlargeStringInfo(m, v144, int32(4))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L9
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v195+v171<<(uint(int32(3))%32))))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v203 = v200 + v171*int32(40)
	v205 = v203 + int32(12)
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v203)+10)))
	if v206 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v187+v188))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v187 + int32(4)
	goto L41
L46:
	;
	v209 = F_OutputFunctionCall(m, v205, v199)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L9
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v214 = F_SendFunctionCall(m, v205, v199)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L9
	} else {
		goto L51
	}
L49:
	;
	v211 = F_strlen(m, v209)
	mBase = m.M
	F_pq_sendcountedtext(m, v144, v209, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L9
	} else {
		goto L50
	}
L50:
	;
	goto L41
L51:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	F_enlargeStringInfo(m, v144, int32(4))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L9
	} else {
		goto L52
	}
L52:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v223 = int32(2)
	v225 = int32(4)
	v226 = int32(base.Ui32(v216)>>(uint(v223)%32)) - v225
	v227 = int32(16711935)
	*(*int32)(unsafe.Add(mBase, uint32(v220+v221))) = base.I32_rotr(v226&v227, int32(8)) | base.I32_rotr(v226, int32(24))&v227
	*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v220 + v225
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v214)))
	F_appendBinaryStringInfo(m, v144, v214+v225, int32(base.Ui32(v242)>>(uint(v223)%32))-v225)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L9
	} else {
		goto L53
	}
L53:
	;
	goto L41
L54:
	;
	goto L40
L55:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_printtup[0])) = v139
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_MemoryContextReset(m, v271)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L9
	} else {
		goto L56
	}
L56:
	;
	m.G0 = v14 + int32(16)
	return int32(1)
}
func F_printtup_startup(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = l0 + int32(40)
	F_initStringInfo(m, v7)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _c_F_printtup_startup[0]))
		v16 = F_AllocSetContextCreateInternal(m, v11, int32(_a_F_printtup_startup_0), int32(0), int32(_a_F_printtup_startup_1), int32(_a_F_printtup_startup_2))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = v16
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
			if v19 == int32(1) {
				v22 = F_FetchPortalTargetList(m, v5)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
					F_SendRowDescriptionMessage(m, v7, l2, v22, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						return
					}
				}
			} else {
				return
			}
		}
	}
}
func F_proclock_hash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_proclock_hash[0]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = F_get_hash_value(m, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		return v6 ^ v10<<(uint(int32(4))%32)
	}
}
func F_psprintf(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_psprintf[0]))
	v14 = F_palloc(m, int32(128))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_psprintf[0])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v22 = F_pvsnprintf(m, v14, int32(128), l0, l1)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(int32(128)) <= base.Ui32(v22) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = v14
	v29 = v22
	goto L7
L5:
	;
	v44 = v14
	goto L6
L6:
	;
	m.G0 = v9 + int32(16)
	return v44
L7:
	;
	F_pfree(m, v28)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v44 = v34
	goto L6
L9:
	;
	v34 = F_palloc(m, v29)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_psprintf[0])) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
	v39 = F_pvsnprintf(m, v34, v29, l0, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if base.Ui32(v29) <= base.Ui32(v39) {
		v28 = v34
		v29 = v39
		goto L7
	} else {
		goto L12
	}
L12:
	;
	goto L8
}
func F_pstrdup(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	v2 = int32(0)
	v4 = F_strlen(m, l0)
	mBase = m.M
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pstrdup[0]))
	*(*uint8)(unsafe.Add(mBase, uint32(v6)+4)) = uint8(v2)
	v10 = v4 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v14 = m.T0[v13].(func(*base.Module, int32, int32, int32) int32)(m, v6, v10, v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v10 != 0 {
			base.MemoryCopy(m, v14, l0, v10)
		} else {
		}
		return v14
	}
}
func F_pull_exec_paramids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	v3 = int32(0)
	if l0 == v3 {
		v24 = v3
		return v24
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v6 == int32(8) {
			v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v9 != int32(1) {
				v24 = v3
				return v24
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				v14 = F_bms_add_member(m, v12, v13)
				mBase = m.M
				v17 = m.ExcPending
				if v17 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v14
					return int32(0)
				}
			}
		} else {
			v22 = F_expression_tree_walker_impl(m, l0, int32(960), l1)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = v22
				return v24
			}
		}
	}
}
func F_pull_up_sublinks_jointree_recurse(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_check_stack_depth(m)
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
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(32)
	return v181
L4:
	;
	v19 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v19
	v181 = v19
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v22 - int32(63) {
	case 0:
		goto L7
	case 1:
		goto L9
	case 2:
		goto L10
	default:
		goto L8
	}
L7:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v177 = F_bms_make_singleton(m, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L45
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L42
	}
L9:
	;
	v80 = F_palloc(m, int32(40))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v25 == int32(0) {
		v62 = v4
		v63 = v4
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v66 = F_makeFromExpr(m, v63, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L20
	}
L12:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v28 <= int32(0) {
		v62 = v4
		v63 = v4
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v36 = v4
	v37 = v4
	v38 = v4
	goto L14
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v38<<(uint(int32(2))%32))))
	v46 = F_pull_up_sublinks_jointree_recurse(m, l0, v43, v11+int32(28))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v62 = v51
	v63 = v48
	goto L11
L16:
	;
	v48 = F_lappend(m, v37, v46)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v51 = F_bms_join(m, v36, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v54 = v38 + int32(1)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v54 < v55 {
		v36 = v51
		v37 = v48
		v38 = v54
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v66
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v72 = int32(0)
	v74 = F_pull_up_sublinks_qual_recurse(m, l0, v69, v11+int32(28), v62, v72, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+8)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v62
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v181 = v78
	goto L3
L22:
	;
	v82 = *(*int64)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+32)) = v82
	v84 = *(*int64)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+24)) = v84
	v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+16)) = v86
	v88 = *(*int64)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v80))) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v80
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v80)+12))
	v96 = F_pull_up_sublinks_jointree_recurse(m, l0, v93, v11+int32(28))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+12)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v102 = F_pull_up_sublinks_jointree_recurse(m, l0, v99, v11+int32(24))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v102
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	switch v105 {
	case 0:
		goto L26
	case 1:
		goto L29
	case 2:
		goto L25
	case 3:
		goto L28
	default:
		goto L27
	}
L25:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v154 = F_bms_join(m, v152, v153)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L37
	}
L26:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v145 = F_bms_union(m, v143, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L32
	}
L28:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v119 = int32(0)
	v121 = F_pull_up_sublinks_qual_recurse(m, l0, v115, v80+int32(12), v118, v119, v119)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v80)+28))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	v110 = int32(0)
	v112 = F_pull_up_sublinks_qual_recurse(m, l0, v106, v80+int32(16), v109, v110, v110)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v112
	goto L25
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v121
	goto L25
L32:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v80)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v128
	F_errmsg_internal(m, int32(_a_F_pull_up_sublinks_jointree_recurse_0), v11+int32(16))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_pull_up_sublinks_jointree_recurse_1), int32(821), int32(_a_F_pull_up_sublinks_jointree_recurse_2))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	v147 = int32(0)
	v149 = F_pull_up_sublinks_qual_recurse(m, l0, v140, v11+int32(20), v145, v147, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v149
	goto L25
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v154
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v80)+36))
	if v157 != 0 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v158 = F_bms_add_member(m, v154, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v181 = v161
	goto L3
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v158
	goto L40
L42:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v166
	F_errmsg_internal(m, int32(_a_F_pull_up_sublinks_jointree_recurse_3), v11)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_pull_up_sublinks_jointree_recurse_1), int32(840), int32(_a_F_pull_up_sublinks_jointree_recurse_2))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v177
	v181 = l1
	goto L3
}
func F_pull_up_subqueries_recurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
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
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
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
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v519 int32
	_ = v519
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v598 int64
	_ = v598
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v676 int32
	_ = v676
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v786 int32
	_ = v786
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v834 int32
	_ = v834
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v841 int32
	_ = v841
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v873 int32
	_ = v873
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v883 int32
	_ = v883
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v898 int32
	_ = v898
	var v906 int32
	_ = v906
	var v916 int32
	_ = v916
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v975 int32
	_ = v975
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1051 int32
	_ = v1051
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
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
	return int32(0)
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_pull_up_subqueries_recurse[0]))
	if v24 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v27 - int32(63) {
	case 0:
		goto L13
	case 1:
		goto L9
	case 2:
		goto L12
	default:
		goto L11
	}
L6:
	;
	goto L5
L7:
	;
	m.G0 = v17 + int32(80)
	return v1092
L8:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v581 = F_copyObjectImpl(m, v48)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L133
	}
L9:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v508 {
	case 0:
		goto L117
	case 1, 4, 5:
		goto L118
	case 2:
		goto L119
	case 3:
		goto L120
	default:
		goto L121
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L1
	} else {
		goto L114
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L1
	} else {
		goto L111
	}
L12:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v448 == int32(0) {
		v1092 = l1
		goto L7
	} else {
		goto L105
	}
L13:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+52))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32+v33<<(uint(int32(2))%32)-int32(4))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v40 != int32(1) {
		v126 = v40
		goto L18
	} else {
		goto L19
	}
L14:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v371 != int32(3) {
		v1092 = l1
		goto L7
	} else {
		goto L86
	}
L15:
	;
	v329 = F_palloc0(m, v328)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L1
	} else {
		goto L81
	}
L16:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v328 = v323<<(uint(int32(2))%32) + int32(4)
	goto L15
L17:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v235)+52))
	if v236 != 0 {
		goto L68
	} else {
		goto L69
	}
L18:
	;
	if l2|l3|base.B2i32(v126 != int32(5)) != 0 {
		goto L14
	} else {
		goto L48
	}
L19:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v44 = F_is_simple_subquery(m, l0, v43, v39, l2)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	if v100 != int32(1) {
		v126 = v100
		goto L18
	} else {
		goto L37
	}
L21:
	;
	if v44 == int32(0) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	if l3 == int32(0) {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+60))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v52 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	if v55 == int32(0) {
		goto L8
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v62 = v51
	goto L28
L27:
	;
	goto L26
L28:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	if v72 != int32(65) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	if v72 != int32(63) {
		goto L20
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	if v77 != 0 {
		goto L20
	} else {
		goto L34
	}
L33:
	;
	goto L8
L34:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)+4))
	if v78 == int32(0) {
		goto L20
	} else {
		goto L35
	}
L35:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v81 != int32(1) {
		goto L20
	} else {
		goto L36
	}
L36:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)))
	v62 = v85
	goto L28
L37:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v104 != int32(67) {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	if v107 != int32(1) {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v103)+144))
	if v110 == int32(0) {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v103)+124))
	if v113 != 0 {
		goto L14
	} else {
		goto L41
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v103)+128))
	if v114 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v103)+132))
	if v115 != 0 {
		goto L14
	} else {
		goto L43
	}
L43:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v103)+140))
	if v116 != 0 {
		goto L14
	} else {
		goto L44
	}
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v103)+48))
	if v117 != 0 {
		goto L14
	} else {
		goto L45
	}
L45:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v110)+20))
	v119 = F_is_simple_union_all_recurse(m, v110, v103, v118)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v119 != 0 {
		goto L17
	} else {
		goto L47
	}
L47:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v126 = v121
	goto L18
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	if v140 == int32(0) {
		goto L14
	} else {
		goto L49
	}
L49:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v140)+4))
	if v143 != int32(1) {
		goto L14
	} else {
		goto L50
	}
L50:
	;
	v146 = F_expression_returns_set(m, v140)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	if v146 != 0 {
		goto L14
	} else {
		goto L52
	}
L52:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v149 = F_contain_volatile_functions(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	if v149 != 0 {
		goto L14
	} else {
		goto L54
	}
L54:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v151)+52))
	if v152 == int32(0) {
		goto L14
	} else {
		goto L55
	}
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
	if v155 != int32(1) {
		goto L14
	} else {
		goto L56
	}
L56:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	if v39 != v159 {
		goto L14
	} else {
		goto L57
	}
L57:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = F_copyObjectImpl(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v220 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v161
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v220
	*(*int64)(unsafe.Add(mBase, uint32(v17)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v215
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v151 + int32(39)
	if v215 != 0 {
		goto L16
	} else {
		goto L67
	}
L59:
	;
	if v165 == int32(0) {
		v215 = v5
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v170 <= int32(0) {
		v215 = v5
		goto L58
	} else {
		goto L61
	}
L61:
	;
	v177 = int32(0)
	v178 = int32(1)
	v183 = v5
	goto L62
L62:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v188+v177<<(uint(int32(2))%32))))
	v194 = int32(0)
	v196 = F_makeTargetEntry(m, v192, base.I32_extend16_s(v178), v194, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	v215 = v198
	goto L58
L64:
	;
	v198 = F_lappend(m, v183, v196)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v200 = int32(1)
	v203 = v177 + v200
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v203 < v204 {
		v177 = v203
		v178 = v178 + v200
		v183 = v198
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v328 = int32(4)
	goto L15
L68:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+4))
	v239 = v237
	goto L70
L69:
	;
	v239 = int32(0)
	goto L70
L70:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v234)+52))
	v242 = F_copyObjectImpl(m, v241)
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v244 = m.G0
	v245 = int32(16)
	v246 = v244 - v245
	m.G0 = v246
	*(*int32)(unsafe.Add(mBase, uint32(v246)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v246)+8)) = int32(-1)
	v256 = F_range_table_walker_impl(m, v242, int32(1128), v246+int32(8), v245)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	m.G0 = v246 + int32(16)
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+124)))
	if base.B2i32(v242 == int32(0))|base.B2i32(v263 != int32(1)) != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v234)+56))
	F_CombineRangeTables(m, v310+int32(52), v310+int32(56), v242, v315)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L79
	}
L74:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v267 <= int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v275 = int32(0)
	goto L76
L76:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v285+v275<<(uint(int32(2))%32))))
	v290 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v289)+124)) = uint8(v290)
	v293 = v275 + v290
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v242)+4))
	if v293 < v294 {
		v275 = v293
		goto L76
	} else {
		goto L78
	}
L77:
	;
	goto L73
L78:
	;
	goto L77
L79:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v234)+144))
	F_pull_up_union_leaf_queries(m, v318, l0, v240, v234, v239)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+20)) = uint8(v321)
	v1092 = l1
	goto L7
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v329
	F_perform_pullup_replace_vars(m, l0, v17+int32(24), int32(0))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v338 = F_palloc0(m, int32(136))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+12)) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v338))) = int32(101)
	v346 = F_makeAlias(m, int32(_a_F_pull_up_subqueries_recurse_0), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+8)) = v346
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v338
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v338
	v354 = F_list_make1_impl(m, int32(1), v17+int32(12))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v151)+52)) = v354
	v1092 = l1
	goto L7
L86:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+72)))
	if v374 != 0 {
		v1092 = l1
		goto L7
	} else {
		goto L87
	}
L87:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v39)+68))
	if v375 == int32(0) {
		v1092 = l1
		goto L7
	} else {
		goto L88
	}
L88:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v375)+4))
	if v378 != int32(1) {
		v1092 = l1
		goto L7
	} else {
		goto L89
	}
L89:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v375)+12))
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v381)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v383)))
	if v384 != int32(7) {
		v1092 = l1
		goto L7
	} else {
		goto L90
	}
L90:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v382)+8))
	if v387 != int32(1) {
		v1092 = l1
		goto L7
	} else {
		goto L91
	}
L91:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v382)+12))
	if v390 != 0 {
		v1092 = l1
		goto L7
	} else {
		goto L92
	}
L92:
	;
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v396 = F_get_expr_result_type(m, v383, v17+int32(68), v17-int32(-64))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v396 != 0 {
		v1092 = l1
		goto L7
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l0
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v382)+4))
	v401 = int32(0)
	v403 = F_makeTargetEntry(m, v399, int32(1), v401, v401)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v403
	v410 = F_list_make1_impl(m, int32(1), v17+int32(8))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v391 + int32(39)
	v415 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v415
	*(*int64)(unsafe.Add(mBase, uint32(v17)+36)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v39
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v410
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v415
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v421
	if v410 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	v431 = v425<<(uint(int32(2))%32) + int32(4)
	goto L99
L98:
	;
	v431 = int32(4)
	goto L99
L99:
	;
	v432 = F_palloc0(m, v431)
	mBase = m.M
	v433 = m.ExcPending
	if v433 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v432
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v391)+108))
	if v435 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(1)
	goto L103
L102:
	;
	goto L103
L103:
	;
	F_perform_pullup_replace_vars(m, l0, v17+int32(24), l3)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+124)) = uint8(v442)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+68)) = v442
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = int32(8)
	v1092 = l1
	goto L7
L105:
	;
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v451 <= int32(0) {
		v1092 = l1
		goto L7
	} else {
		goto L106
	}
L106:
	;
	v458 = v5
	goto L107
L107:
	;
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	v471 = v468 + v458<<(uint(int32(2))%32)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v471)))
	v474 = F_pull_up_subqueries_recurse(m, l0, v472, l2, int32(0))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L109
	}
L108:
	;
	v1092 = l1
	goto L7
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = v474
	v478 = v458 + int32(1)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	if v478 < v479 {
		v458 = v478
		goto L107
	} else {
		goto L110
	}
L110:
	;
	goto L108
L111:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v485
	F_errmsg_internal(m, int32(_a_F_pull_up_subqueries_recurse_1), v17)
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	F_errfinish(m, int32(_a_F_pull_up_subqueries_recurse_2), int32(1398), int32(_a_F_pull_up_subqueries_recurse_3))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L114:
	;
	F_errmsg_internal(m, int32(_a_F_pull_up_subqueries_recurse_4), int32(0))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	F_errfinish(m, int32(_a_F_pull_up_subqueries_recurse_2), int32(2371), int32(_a_F_pull_up_subqueries_recurse_5))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L117:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v557 = F_pull_up_subqueries_recurse(m, l0, v555, l2, int32(0))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L131
	}
L118:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v547 = F_pull_up_subqueries_recurse(m, l0, v545, l1, int32(0))
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L1
	} else {
		goto L129
	}
L119:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v537 = F_pull_up_subqueries_recurse(m, l0, v535, l1, int32(0))
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L1
	} else {
		goto L127
	}
L120:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v527 = F_pull_up_subqueries_recurse(m, l0, v525, l1, int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L1
	} else {
		goto L125
	}
L121:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v513
	F_errmsg_internal(m, int32(_a_F_pull_up_subqueries_recurse_6), v17+int32(16))
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_pull_up_subqueries_recurse_2), int32(1392), int32(_a_F_pull_up_subqueries_recurse_3))
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v527
	v530 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v532 = F_pull_up_subqueries_recurse(m, l0, v530, l1, int32(0))
	mBase = m.M
	v533 = m.ExcPending
	if v533 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v532
	v1092 = l1
	goto L7
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v537
	v540 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v542 = F_pull_up_subqueries_recurse(m, l0, v540, l1, int32(0))
	mBase = m.M
	v543 = m.ExcPending
	if v543 != 0 {
		goto L1
	} else {
		goto L128
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v542
	v1092 = l1
	goto L7
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v547
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v552 = F_pull_up_subqueries_recurse(m, l0, v550, l1, int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L130
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v552
	v1092 = l1
	goto L7
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v557
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v562 = F_pull_up_subqueries_recurse(m, l0, v560, l2, int32(0))
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v562
	v1092 = l1
	goto L7
L133:
	;
	v584 = F_palloc0(m, int32(400))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+4)) = v581
	*(*int32)(unsafe.Add(mBase, uint32(v584))) = int32(269)
	v589 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+8)) = v589
	v591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+12)) = v591
	v593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+20)) = v593
	v595 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v584)+24)) = v595
	v597 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v598 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v584)+28)) = v598
	*(*int32)(unsafe.Add(mBase, uint32(v584)+16)) = v597
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_pull_up_subqueries_recurse[1]))
	v603 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v584)+339)) = uint8(v603)
	*(*uint16)(unsafe.Add(mBase, uint32(v584)+337)) = uint16(v603)
	*(*int32)(unsafe.Add(mBase, uint32(v584)+328)) = v603
	*(*int32)(unsafe.Add(mBase, uint32(v584)+300)) = v602
	*(*int64)(unsafe.Add(mBase, uint32(v584)+80)) = v598
	*(*int64)(unsafe.Add(mBase, uint32(v584)+88)) = v598
	*(*int64)(unsafe.Add(mBase, uint32(v584)+93)) = v598
	*(*int64)(unsafe.Add(mBase, uint32(v584)+124)) = v598
	*(*int64)(unsafe.Add(mBase, uint32(v584)+132)) = v598
	*(*int64)(unsafe.Add(mBase, uint32(v584)+140)) = v598
	base.MemoryFill(m, v584+int32(212), v603, int32(88))
	*(*int64)(unsafe.Add(mBase, uint32(v584)+360)) = int64(4294967295)
	v629 = F_preprocess_relation_rtes(m, v584)
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v584)+4)) = v629
	F_replace_empty_jointree(m, v629)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v634 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+39)))
	if v634 == int32(1) {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v637)+60))
	v641 = F_pull_up_sublinks_jointree_recurse(m, v584, v638, v17+int32(24))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v660)+52))
	if v661 == int32(0) {
		v714 = v660
		goto L146
	} else {
		goto L147
	}
L140:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
	if v643 != int32(65) {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v641
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v641
	v651 = F_list_make1_impl(m, int32(1), v17+int32(4))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L144
	}
L142:
	;
	v656 = v641
	goto L143
L143:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v657)+60)) = v656
	goto L139
L144:
	;
	v654 = F_makeFromExpr(m, v651, int32(0))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v656 = v654
	goto L143
L146:
	;
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v714)+60))
	v725 = int32(0)
	v727 = F_pull_up_subqueries_recurse(m, v584, v724, v725, v725)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L157
	}
L147:
	;
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	if v664 <= int32(0) {
		v714 = v660
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v676 = v5
	goto L149
L149:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v661)+12))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v681+v676<<(uint(int32(2))%32))))
	v686 = *(*int32)(unsafe.Add(mBase, uint32(v685)+12))
	if v686 != int32(3) {
		goto L151
	} else {
		goto L152
	}
L150:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	v714 = v709
	goto L146
L151:
	;
	v706 = v676 + int32(1)
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v661)+4))
	if v706 < v707 {
		v676 = v706
		goto L149
	} else {
		goto L156
	}
L152:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v685)+68))
	v690 = F_eval_const_expressions(m, v584, v689)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v685)+68)) = v690
	v693 = F_inline_function_in_from(m, v584, v685)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	if v693 == int32(0) {
		goto L151
	} else {
		goto L155
	}
L155:
	;
	v697 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+72)) = uint8(v697)
	*(*uint8)(unsafe.Add(mBase, uint32(v685)+40)) = uint8(v697)
	*(*int32)(unsafe.Add(mBase, uint32(v685)+36)) = v693
	*(*int32)(unsafe.Add(mBase, uint32(v685)+12)) = int32(1)
	goto L151
L156:
	;
	goto L150
L157:
	;
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v729)+60)) = v727
	v731 = F_is_simple_subquery(m, l0, v629, v39, l2)
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	if v731 == int32(0) {
		v1092 = l1
		goto L7
	} else {
		goto L159
	}
L159:
	;
	if l3 != 0 {
		goto L160
	} else {
		goto L161
	}
L160:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v629)+60))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v735)+4))
	if v736 != 0 {
		goto L164
	} else {
		goto L165
	}
L161:
	;
	goto L162
L162:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v584)+4))
	v804 = *(*int32)(unsafe.Add(mBase, uint32(v629)+76))
	v805 = F_flatten_join_alias_vars(m, v584, v803, v804)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L179
	}
L163:
	;
	if v786 == int32(0) {
		v1092 = l1
		goto L7
	} else {
		goto L178
	}
L164:
	;
	v742 = v735
	goto L168
L165:
	;
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v735)+8))
	if v737 != 0 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v786 = int32(1)
	goto L163
L167:
	;
	v786 = v771
	goto L163
L168:
	;
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v742)))
	if v754 != int32(65) {
		goto L171
	} else {
		goto L172
	}
L169:
	;
	v771 = int32(0)
	goto L167
L170:
	;
	goto L169
L171:
	;
	if v754 == int32(63) {
		v771 = int32(1)
		goto L167
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v759 = *(*int32)(unsafe.Add(mBase, uint32(v742)+8))
	if v759 != 0 {
		goto L170
	} else {
		goto L175
	}
L174:
	;
	goto L170
L175:
	;
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v742)+4))
	if v760 == int32(0) {
		goto L170
	} else {
		goto L176
	}
L176:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v760)+4))
	if v763 != int32(1) {
		goto L170
	} else {
		goto L177
	}
L177:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v760)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v766)))
	v742 = v767
	goto L168
L178:
	;
	goto L162
L179:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v629)+76)) = v805
	v808 = int32(0)
	v810 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
	if v810 != 0 {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v810)+4))
	v812 = v811
	goto L182
L181:
	;
	v812 = v808
	goto L182
L182:
	;
	F_OffsetVarNodes(m, v629, v812)
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v584)+136))
	F_OffsetVarNodes(m, v815, v812)
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L184
	}
L184:
	;
	F_IncrementVarSublevelsUp(m, v629, int32(-1), int32(1))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L1
	} else {
		goto L185
	}
L185:
	;
	v822 = *(*int32)(unsafe.Add(mBase, uint32(v584)+136))
	F_IncrementVarSublevelsUp(m, v822, int32(-1), int32(1))
	mBase = m.M
	v826 = m.ExcPending
	if v826 != 0 {
		goto L1
	} else {
		goto L186
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l0
	v828 = *(*int32)(unsafe.Add(mBase, uint32(v629)+76))
	v829 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v829
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v828
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v39
	v834 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+124)))
	if v834 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v837 = *(*int32)(unsafe.Add(mBase, uint32(v629)+60))
	v838 = int32(1)
	v840 = F_get_relids_in_jointree(m, v837, v838, v838)
	mBase = m.M
	v841 = m.ExcPending
	if v841 != 0 {
		goto L1
	} else {
		goto L190
	}
L188:
	;
	v861 = v829
	v863 = v828
	v864 = v808
	goto L189
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v579
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v861
	*(*int32)(unsafe.Add(mBase, uint32(v17)+40)) = v864
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v580 + int32(39)
	if v863 != 0 {
		goto L197
	} else {
		goto L198
	}
L190:
	;
	v844 = F_palloc(m, int32(8))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	v846 = *(*int32)(unsafe.Add(mBase, uint32(v580)+52))
	if v846 != 0 {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v847 = *(*int32)(unsafe.Add(mBase, uint32(v846)+4))
	v848 = v847
	goto L194
L193:
	;
	v848 = int32(0)
	goto L194
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844)+4)) = v848
	v853 = F_palloc0_mul(m, int32(4), v848+int32(1))
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v844))) = v853
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v580)+60))
	F_get_nullingrels_recurse(m, v856, int32(0), v844)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v629)+76))
	v861 = v844
	v863 = v860
	v864 = v840
	goto L189
L197:
	;
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v863)+4))
	v879 = v873<<(uint(int32(2))%32) + int32(4)
	goto L199
L198:
	;
	v879 = int32(4)
	goto L199
L199:
	;
	v880 = F_palloc0(m, v879)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v880
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v580)+108))
	if v883 != 0 {
		goto L201
	} else {
		goto L202
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = int32(1)
	goto L203
L202:
	;
	goto L203
L203:
	;
	F_perform_pullup_replace_vars(m, l0, v17+int32(24), l3)
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L1
	} else {
		goto L204
	}
L204:
	;
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+124)))
	if v892 != int32(1) {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v629)+52))
	v948 = *(*int32)(unsafe.Add(mBase, uint32(v629)+56))
	F_CombineRangeTables(m, v580+int32(52), v580+int32(56), v947, v948)
	mBase = m.M
	v950 = m.ExcPending
	if v950 != 0 {
		goto L1
	} else {
		goto L216
	}
L206:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v629)+52))
	if v895 == int32(0) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v898 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
	if v898 <= int32(0) {
		goto L205
	} else {
		goto L208
	}
L208:
	;
	v906 = int32(0)
	goto L209
L209:
	;
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v895)+12))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v916+v906<<(uint(int32(2))%32))))
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v920)+12))
	switch v921 {
	case 0:
		goto L213
	case 1, 3, 4, 5:
		goto L212
	default:
		goto L211
	}
L210:
	;
	goto L205
L211:
	;
	v928 = v906 + int32(1)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
	if v928 < v929 {
		v906 = v928
		goto L209
	} else {
		goto L215
	}
L212:
	;
	v925 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v920)+124)) = uint8(v925)
	goto L211
L213:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v920)+32))
	if v922 == int32(0) {
		goto L211
	} else {
		goto L214
	}
L214:
	;
	goto L212
L215:
	;
	goto L210
L216:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(v580)+140))
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v629)+140))
	v953 = F_list_concat(m, v951, v952)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v580)+140)) = v953
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)+80))
	if v957 != 0 {
		goto L219
	} else {
		goto L220
	}
L218:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(v584)+136))
	v1068 = F_list_concat(m, v1066, v1067)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L244
	}
L219:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v629)+60))
	v963 = F_get_relids_in_jointree(m, v960, int32(1), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L1
	} else {
		goto L222
	}
L220:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v958 != 0 {
		goto L219
	} else {
		goto L221
	}
L221:
	;
	v1066 = int32(0)
	goto L218
L222:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v966 = *(*int32)(unsafe.Add(mBase, uint32(v965)+80))
	if v966 != 0 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v963
	v968 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v968
	*(*int32)(unsafe.Add(mBase, uint32(v17)+68)) = v579
	v975 = F_query_or_expression_tree_walker_impl(m, v580, int32(900), v17+int32(68), v968)
	mBase = m.M
	v976 = m.ExcPending
	if v976 != 0 {
		goto L1
	} else {
		goto L226
	}
L224:
	;
	goto L225
L225:
	;
	v977 = int32(0)
	v978 = m.G0
	v980 = v978 - int32(16)
	m.G0 = v980
	v982 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v982 == v977 {
		goto L227
	} else {
		goto L228
	}
L226:
	;
	goto L225
L227:
	;
	m.G0 = v980 + int32(16)
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v1066 = v1051
	goto L218
L228:
	;
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v982)+4))
	if v985 <= int32(0) {
		goto L227
	} else {
		goto L229
	}
L229:
	;
	v992 = int32(-1)
	v998 = v977
	goto L230
L230:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v982)+12))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v1003+v998<<(uint(int32(2))%32))))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+8))
	if v579 == v1008 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	goto L227
L232:
	;
	if v992 < int32(0) {
		goto L235
	} else {
		goto L236
	}
L233:
	;
	v1016 = v992
	goto L234
L234:
	;
	v1017 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v1017)+80))
	if v1018 != 0 {
		goto L239
	} else {
		goto L240
	}
L235:
	;
	v1012 = F_bms_singleton_member(m, v963)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L238
	}
L236:
	;
	v1014 = v992
	goto L237
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1007)+8)) = v1014
	v1016 = v1014
	goto L234
L238:
	;
	v1014 = v1012
	goto L237
L239:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1007)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v980)+12)) = v963
	v1021 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v980)+8)) = v1021
	*(*int32)(unsafe.Add(mBase, uint32(v980)+4)) = v579
	v1028 = F_query_or_expression_tree_walker_impl(m, v1019, int32(900), v980+int32(4), v1021)
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v1031 = v998 + int32(1)
	v1032 = *(*int32)(unsafe.Add(mBase, uint32(v982)+4))
	if v1031 < v1032 {
		v992 = v1016
		v998 = v1031
		goto L230
	} else {
		goto L243
	}
L242:
	;
	goto L241
L243:
	;
	goto L231
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v1068
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = int32(0)
	v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+39)))
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+39)))
	v1075 = v1073 | v1074
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+39)) = uint8(v1075)
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v580)+44)))
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v629)+44)))
	v1079 = v1077 | v1078
	*(*uint8)(unsafe.Add(mBase, uint32(v580)+44)) = uint8(v1079)
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v629)+60))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+8))
	if v1082 != 0 {
		v1092 = v1081
		goto L7
	} else {
		goto L245
	}
L245:
	;
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1081)+4))
	if v1083 == int32(0) {
		v1092 = v1081
		goto L7
	} else {
		goto L246
	}
L246:
	;
	v1086 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+4))
	if v1086 != int32(1) {
		v1092 = v1081
		goto L7
	} else {
		goto L247
	}
L247:
	;
	v1089 = *(*int32)(unsafe.Add(mBase, uint32(v1083)+12))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1089)))
	v1092 = v1090
	goto L7
}
func F_pullf_create(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v41 int32
	_ = v41
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v11 != 0 {
		v14 = m.T0[v11].(func(*base.Module, int32, int32, int32) int32)(m, v9+int32(12), l2, l3)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			if int32(0) <= v14 {
				v22 = v14
				v24 = F_palloc0(m, int32(24))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
					*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
					if v22 != 0 {
						v32 = F_palloc(m, v22)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							v34 = v32
							v35 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
							*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
							*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
							v41 = v35
							m.G0 = v9 + int32(16)
							return v41
						}
					} else {
						v34 = int32(0)
						v35 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
						*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
						v41 = v35
						m.G0 = v9 + int32(16)
						return v41
					}
				}
			} else {
				v41 = v14
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l2
		v22 = int32(0)
		v24 = F_palloc0(m, int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v24)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v24)+8)) = v22
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v24))) = l3
			*(*int32)(unsafe.Add(mBase, uint32(v24)+20)) = v28
			if v22 != 0 {
				v32 = F_palloc(m, v22)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = v32
					v35 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
					*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
					v41 = v35
					m.G0 = v9 + int32(16)
					return v41
				}
			} else {
				v34 = int32(0)
				v35 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v24)+16)) = v35
				*(*int32)(unsafe.Add(mBase, uint32(v24)+12)) = v34
				*(*int32)(unsafe.Add(mBase, uint32(l0))) = v24
				v41 = v35
				m.G0 = v9 + int32(16)
				return v41
			}
		}
	}
}
func F_push_into_mbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	if int32(0) < l3 {
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+16)))
		if v9 == int32(1) {
			F_px_debug(m, int32(_a_F_push_into_mbuf_0), int32(0))
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int32(0)
			} else {
				return int32(-12)
			}
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v20) < base.Ui32(v21+l3) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				v30 = v20 - v24 + (l3+int32(_a_F_push_into_mbuf_1))&int32(2147467264)
				v31 = F_repalloc(m, v24, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v31 + v30
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v31
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					v39 = v31 + (v37 - v35)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v39
					v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v31 + (v41 - v35)
					v45 = v39
					if l3 != 0 {
						base.MemoryCopy(m, v45, l2, l3)
					} else {
					}
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v49 + l3
					return int32(0)
				}
			} else {
				v45 = v21
				if l3 != 0 {
					base.MemoryCopy(m, v45, l2, l3)
				} else {
				}
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v49 + l3
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_pushval_morph(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	v7 = int32(0)
	v14 = m.G0
	v15 = int32(16)
	v16 = v14 - v15
	m.G0 = v16
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v7
	*(*int64)(unsafe.Add(mBase, uint32(v16)+4)) = int64(4)
	v24 = F_palloc_mul(m, v15, int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_parsetext(m, v27, v16, l2, l3)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if int32(0) < v30 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v16 + int32(16)
	return
L5:
	;
	v37 = int32(0)
	v41 = v30
	v42 = v7
	v45 = v7
	goto L8
L6:
	;
	goto L7
L7:
	;
	F_pushStop(m, l1)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L56
	}
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	if v42 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	F_pfree(m, v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L55
	}
L10:
	;
	if v96 <= v37 {
		v224 = v37
		v228 = v96
		goto L23
	} else {
		goto L24
	}
L11:
	;
	v53 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v37<<(uint(int32(4))%32))+8)))
	v96 = v41
	v97 = v53
	v100 = v45
	goto L10
L12:
	;
	goto L13
L13:
	;
	v55 = v42 + int32(1)
	v57 = v37 << (uint(int32(4)) % 32)
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v47+v57)+8)))
	if base.Ui32(v59) <= base.Ui32(v55) {
		v96 = v41
		v97 = v59
		v100 = v45
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v63 = v55
	v72 = v45
	goto L15
L15:
	;
	F_pushStop(m, l1)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v96 = v88
	v97 = v86
	v100 = v81
	goto L10
L17:
	;
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pushOperator(m, l1, v76, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v80 = int32(1)
	v81 = v72 + v80
	v83 = v63 + v80
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84+v57)+8)))
	if base.Ui32(v83) < base.Ui32(v86) {
		v63 = v83
		v72 = v81
		goto L15
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	goto L16
L23:
	;
	if v100 != 0 {
		goto L50
	} else {
		goto L51
	}
L24:
	;
	v107 = v37
	v111 = v96
	v114 = int32(0)
	goto L25
L25:
	;
	v118 = v107 << (uint(int32(4)) % 32)
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v120 = v118 + v119
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+8)))
	if v97 != v121 {
		v224 = v107
		v228 = v111
		goto L23
	} else {
		goto L27
	}
L26:
	;
	v224 = v202
	v228 = v217
	goto L23
L27:
	;
	if v111 <= v107 {
		v202 = v107
		v206 = v111
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v114 != 0 {
		goto L45
	} else {
		goto L46
	}
L29:
	;
	v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+4)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v120)+12))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120)+2)))
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v120))))
	F_pushValue(m, l1, v125, v126, l4, l5|int32(base.Ui32(v127&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v135+v118)+12))
	F_pfree(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v141 = v107 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v142 <= v141 {
		v202 = v141
		v206 = v142
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v148 = v145 + v141<<(uint(int32(4))%32)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+8)))
	if v97 != v149 {
		v202 = v141
		v206 = v142
		goto L28
	} else {
		goto L33
	}
L33:
	;
	v153 = v148
	v154 = v141
	v158 = v142
	v160 = int32(1)
	goto L34
L34:
	;
	v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	if v124 != v164 {
		v202 = v154
		v206 = v158
		goto L28
	} else {
		goto L36
	}
L35:
	;
	v202 = v188
	v206 = v189
	goto L28
L36:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+2)))
	v168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
	F_pushValue(m, l1, v166, v167, l4, l5|int32(base.Ui32(v168&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176+v154<<(uint(int32(4))%32))+12))
	F_pfree(m, v180)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	if v160 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	F_pushOperator(m, l1, int32(2), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v188 = v154 + int32(1)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v189 <= v188 {
		v202 = v188
		v206 = v189
		goto L28
	} else {
		goto L43
	}
L42:
	;
	goto L41
L43:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v196 = v193 + v188<<(uint(int32(4))%32)
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v196)+8)))
	if v97 == v197 {
		v153 = v196
		v154 = v188
		v158 = v189
		v160 = v160 + int32(1)
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	F_pushOperator(m, l1, int32(3), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	v217 = v206
	goto L47
L47:
	;
	if v202 < v217 {
		v107 = v202
		v111 = v217
		v114 = v114 + int32(1)
		goto L25
	} else {
		goto L49
	}
L48:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v217 = v216
	goto L47
L49:
	;
	goto L26
L50:
	;
	v234 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+4)))
	F_pushOperator(m, l1, v234, int32(1))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	v239 = v228
	goto L52
L52:
	;
	if v224 < v239 {
		v37 = v224
		v41 = v239
		v42 = v97
		v45 = v100 + int32(1)
		goto L8
	} else {
		goto L54
	}
L53:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v239 = v238
	goto L52
L54:
	;
	goto L9
L55:
	;
	goto L4
L56:
	;
	goto L4
}
func F_pvsnprintf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_pg_vsnprintf(m, l0, l1, l2, l3)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if int32(0) <= v10 {
			if base.Ui32(l1) <= base.Ui32(v10) {
				if base.Ui32(int32(1073741823)) <= base.Ui32(v10) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_pvsnprintf_0), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_pvsnprintf_1), int32(140), int32(_a_F_pvsnprintf_2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
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
					v21 = v10 + int32(1)
					m.G0 = v8 + int32(16)
					return v21
				}
			} else {
				v21 = v10
				m.G0 = v8 + int32(16)
				return v21
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8))) = l2
				F_errmsg_internal(m, int32(_a_F_pvsnprintf_3), v8)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_pvsnprintf_1), int32(113), int32(_a_F_pvsnprintf_2))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
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
