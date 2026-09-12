package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tuplesort_begin_index_btree(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
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
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v167 int32
	_ = v167
	v3 = l2
	v4 = l3
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v16 = F_tuplesort_begin_common(m, l4, l5, int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = int32(4554128)
	v21 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v16)+24))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v23
	v26 = F_palloc(m, int32(12))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[82])))
	if v29 != int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+10)))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1847)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v26
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v59)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(1848)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(1849)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = int32(1850)
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1851)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+9)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+8)) = uint8(v3)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = l0
	v74 = F__bt_mkscankey(m, l1, int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v34 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = l4
	if v3 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v41 = int32(116)
	goto L10
L9:
	;
	v41 = int32(102)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(102)
	F_errmsg_internal(m, int32(526219), v13)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(516014), int32(382), int32(429172))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	goto L4
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	v79 = F_palloc0(m, v76*int32(36))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v82 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_pfree(m, v74)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L23
	}
L16:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v79))) = v86
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v74)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v79)+4)) = v88
	v90 = int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	v95 = int32(base.Ui32(v91)>>(uint(int32(25))%32)) & v90
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+9)) = uint8(v95)
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v79)+10)) = uint16(v97)
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)))
	*(*uint8)(unsafe.Add(mBase, uint32(v79)+20)) = uint8(v99)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	F_PrepareSortSupportFromIndexRel(m, l1, int32(base.Ui32(v101&int32(16777216))>>(uint(int32(24))%32)), v79)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v108 < int32(2) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v120 = v90
	goto L19
L19:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v16)+44))
	v126 = v123 + v120*int32(36)
	v128 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v128
	v132 = v74 + int32(16) + v120*int32(48)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v132)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	v139 = int32(base.Ui32(v135)>>(uint(int32(25))%32)) & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+9)) = uint8(v139)
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v132)+4)))
	v142 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v126)+20)) = uint8(v142)
	*(*uint16)(unsafe.Add(mBase, uint32(v126)+10)) = uint16(v141)
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v132)))
	F_PrepareSortSupportFromIndexRel(m, l1, int32(base.Ui32(v145&int32(16777216))>>(uint(int32(24))%32)), v126)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L15
L21:
	;
	v153 = v120 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v16)+40))
	if v153 < v154 {
		v120 = v153
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v21
	m.G0 = v13 + int32(16)
	return v16
}
func F_tuplesort_merge_order(m *base.Module, l0 int64) int32 {
	var v4 int32
	_ = v4
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v4 = int32(6)
	v6 = base.I64_div_s(l0, int64(278528))
	v7 = base.I32_wrap_i64(v6)
	if v7 <= v4 {
		v10 = v4
	} else {
		v10 = v7
	}
	if int32(500) <= v10 {
		v13 = int32(500)
	} else {
		v13 = v10
	}
	return v13
}
func F_tuplesort_putindextuplevalues(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v15 = F_index_form_tuple_context(m, v13, l3, l4, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v10))) = v15
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)) = uint16(v18)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v20
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+52))
		v27 = F_index_getattr_2(m, v15, int32(1), v24, v10+int32(8))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v27
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
			if v30&int32(2) == int32(0) {
				v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+6)))
				v44 = (v35&int32(8191) + int32(7)) & int32(16376)
				v45 = int32(0)
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				if v46 == v45 {
					v55 = v45
				} else {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
					if v49 == int32(0) {
						v55 = v45
					} else {
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
						v55 = v52 ^ int32(1)
					}
				}
				F_tuplesort_puttuple_common(m, l0, v10, v55&int32(1), v44)
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				v42 = F_GetMemoryChunkSpace(m, v15)
				mBase = m.M
				v43 = m.ExcPending
				if v43 != 0 {
					return
				} else {
					v44 = v42
					v45 = int32(0)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					if v46 == v45 {
						v55 = v45
					} else {
						v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
						if v49 == int32(0) {
							v55 = v45
						} else {
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+8)))
							v55 = v52 ^ int32(1)
						}
					}
					F_tuplesort_puttuple_common(m, l0, v10, v55&int32(1), v44)
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return
					} else {
						m.G0 = v10 + int32(16)
						return
					}
				}
			}
		}
	}
}
func F_tuplesort_set_bound(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	v2 = l1
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+236))
	if v5 != 0 {
		if int64(1073741823) < v2 {
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
			if v8 != int32(-1) {
				*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)) = uint32(v2)
				v14 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)) = uint8(v14)
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(0)
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
				if v20 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v20
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v23 = v22
				} else {
					v23 = v19
				}
				v24 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v24
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v24
			} else {
			}
		}
	} else {
		if int64(1073741823) < v2 {
		} else {
			*(*uint32)(unsafe.Add(mBase, uint32(l0)+72)) = uint32(v2)
			v14 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+68)) = uint8(v14)
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
			if v20 != 0 {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v20
				v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
				v23 = v22
			} else {
				v23 = v19
			}
			v24 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = v24
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v24
		}
	}
	return
}
