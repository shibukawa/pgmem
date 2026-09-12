package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HnswGetEfConstruction(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v2 == int32(0) {
		return int32(64)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		return v7
	}
}
func F_HnswGetEntryPoint(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_HnswGetMetaPageInfo(m, l0, int32(0), v5+int32(12))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
		m.G0 = v5 + int32(16)
		return v14
	}
}
func F_HnswGetTypeInfo(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	v3 = int32(4116384)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+6)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v6+v8*int32(0)<<(uint(int32(2))%32)+int32(12)-int32(4))))
	if v20 == int32(0) {
		v34 = v3
		return v34
	} else {
		v25 = F_index_getprocinfo(m, l0, int32(1), int32(3))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			if v25 == int32(0) {
				v34 = v3
				return v34
			} else {
				v31 = F_FunctionCall0Coll(m, v25)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v34 = v31
					return v34
				}
			}
		}
	}
}
func F_HnswInitElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 float64, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v46 float64
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	if l5 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
		v12 = m.T0[v11].(func(*base.Module, int32, int32) int32)(m, int32(108), v10)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v19 = v12
			v20 = int32(4637744)
			v23 = *(*int64)(unsafe.Add(mBase, _consts[189]))
			v24 = *(*int64)(unsafe.Add(mBase, _consts[190]))
			v25 = v23 ^ v24
			*(*int64)(unsafe.Add(mBase, _consts[190])) = base.I64_rotl(v25, int64(37))
			*(*int64)(unsafe.Add(mBase, _consts[189])) = v25<<(uint(int64(16))%64) ^ base.I64_rotl(v23, int64(24)) ^ v25
			v46 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v23*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
			mBase = m.M
			v47 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+64)) = uint8(v47)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v51)
			v53 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v19)+66)) = uint16(v53)
			v55 = F_log(m, v46)
			mBase = m.M
			v57 = base.F64_mul(v55, base.F64_neg(l3))
			if base.F64_lt(base.F64_abs(v57), float64(2.147483648e+09)) != 0 {
				v61 = base.I32_trunc_f64_s(v57)
				v63 = v61
			} else {
				v63 = int32(-2147483648)
			}
			if l4 < v63 {
				v65 = l4
			} else {
				v65 = v63
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+65)) = uint8(v65)
			F_HnswInitNeighbors(m, l0, v19, l2, l5)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = int32(0)
				return v19
			}
		}
	} else {
		v17 = F_palloc(m, int32(108))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = v17
			v20 = int32(4637744)
			v23 = *(*int64)(unsafe.Add(mBase, _consts[189]))
			v24 = *(*int64)(unsafe.Add(mBase, _consts[190]))
			v25 = v23 ^ v24
			*(*int64)(unsafe.Add(mBase, _consts[190])) = base.I64_rotl(v25, int64(37))
			*(*int64)(unsafe.Add(mBase, _consts[189])) = v25<<(uint(int64(16))%64) ^ base.I64_rotl(v23, int64(24)) ^ v25
			v46 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v23*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
			mBase = m.M
			v47 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+64)) = uint8(v47)
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v49
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)) = uint16(v51)
			v53 = int32(256)
			*(*uint16)(unsafe.Add(mBase, uint32(v19)+66)) = uint16(v53)
			v55 = F_log(m, v46)
			mBase = m.M
			v57 = base.F64_mul(v55, base.F64_neg(l3))
			if base.F64_lt(base.F64_abs(v57), float64(2.147483648e+09)) != 0 {
				v61 = base.I32_trunc_f64_s(v57)
				v63 = v61
			} else {
				v63 = int32(-2147483648)
			}
			if l4 < v63 {
				v65 = l4
			} else {
				v65 = v63
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+65)) = uint8(v65)
			F_HnswInitNeighbors(m, l0, v19, l2, l5)
			mBase = m.M
			v68 = m.ExcPending
			if v68 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = int32(0)
				return v19
			}
		}
	}
}
func F_HnswInsertAppendPage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
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
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	F_LockRelationForExtension(m, l0, int32(7))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = F_HnswNewBuffer(m, l0, int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
			F_UnlockRelationForExtension(m, l0, int32(7))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if l5 != 0 {
					if v17 < int32(0) {
						v21 = *(*int32)(unsafe.Add(mBase, _consts[1]))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v17^int32(-1))<<(uint(int32(2))%32))))
						v40 = v27
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, _consts[2]))
						v40 = v29 + v17<<(uint(int32(13))%32) + int32(-8192)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v40
					F_PageInit(m, v40, int32(8192), int32(8))
					mBase = m.M
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
					v46 = v40 + v45
					v47 = int32(65424)
					*(*uint16)(unsafe.Add(mBase, uint32(v46)+6)) = uint16(v47)
					*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(-1)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v51 < int32(0) {
						v55 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v51^int32(-1))<<(uint(int32(6))%32))+16))
						v70 = v61
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v51<<(uint(int32(6))%32)+int32(-64))+16))
						v70 = v69
					}
					v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+16)))
					*(*int32)(unsafe.Add(mBase, uint32(l4+v71))) = v70
					return
				} else {
					v36 = F_GenericXLogRegisterBuffer(m, l3, v17, int32(1))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v40 = v36
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v40
						F_PageInit(m, v40, int32(8192), int32(8))
						mBase = m.M
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
						v46 = v40 + v45
						v47 = int32(65424)
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+6)) = uint16(v47)
						*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(-1)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v51 < int32(0) {
							v55 = *(*int32)(unsafe.Add(mBase, _consts[3]))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v51^int32(-1))<<(uint(int32(6))%32))+16))
							v70 = v61
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, _consts[4]))
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v51<<(uint(int32(6))%32)+int32(-64))+16))
							v70 = v69
						}
						v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+16)))
						*(*int32)(unsafe.Add(mBase, uint32(l4+v71))) = v70
						return
					}
				}
			}
		}
	}
}
func F_HnswLoadElement(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l0
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+80)))
	F_HnswLoadElementImpl(m, v14, v15, l1, l2, l3, l4, l5, l6, v11+int32(12))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		m.G0 = v11 + int32(16)
		return
	}
}
func F_HnswOptionalProcInfo(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+6)))
	v10 = int32(2)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v4+v6*int32(0)<<(uint(v10)%32)+l1<<(uint(v10)%32)-int32(4))))
	if v18 == int32(0) {
		return int32(0)
	} else {
		v24 = F_index_getprocinfo(m, l0, int32(1), l1)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_HnswSetElementTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	if l0 == int32(0) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
		v16 = v7
	} else {
		v8 = int32(0)
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
		if v9 == v8 {
			v16 = v8
		} else {
			v16 = l0 + v9 - int32(1)
		}
	}
	v17 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v17)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)) = uint8(v20)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v19)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)) = uint8(v23)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
	if v25 == v20 {
		v28 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)) = uint16(v28)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = int32(-1)
		v42 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v42)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(-1)
		v54 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v54)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
		v66 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v66)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+22)) = int32(-1)
		v78 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v78)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
		v90 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v90)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = int32(-1)
		v102 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
		v114 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
		v126 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
		v138 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
	} else {
		v33 = l1 + int32(4)
		v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v33))) = v34
		v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
		*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v36)
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
		if base.Ui32(int32(1)) < base.Ui32(v38) {
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+10))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = v46
			v48 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v48)
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
			if base.Ui32(int32(2)) < base.Ui32(v50) {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v58
				v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+20)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v60)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
				if base.Ui32(int32(3)) < base.Ui32(v62) {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(l2)+22))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+22)) = v70
					v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+26)))
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v72)
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
					if base.Ui32(int32(4)) < base.Ui32(v74) {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v82
						v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+32)))
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v84)
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
						if base.Ui32(int32(5)) < base.Ui32(v86) {
							v94 = *(*int32)(unsafe.Add(mBase, uint32(l2)+34))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = v94
							v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+38)))
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v96)
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
							if base.Ui32(int32(6)) < base.Ui32(v98) {
								v106 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v106
								v108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v108)
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
								if base.Ui32(int32(7)) < base.Ui32(v110) {
									v118 = *(*int32)(unsafe.Add(mBase, uint32(l2)+46))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = v118
									v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+50)))
									*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v120)
									v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
									if base.Ui32(int32(8)) < base.Ui32(v122) {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v130
										v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+56)))
										*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v132)
										v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
										if base.Ui32(int32(9)) < base.Ui32(v134) {
											v142 = *(*int32)(unsafe.Add(mBase, uint32(l2)+58))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = v142
											v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+62)))
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v144)
										} else {
											v138 = int32(0)
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
										}
									} else {
										v126 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
										v138 = int32(0)
										*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
									}
								} else {
									v114 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
									v126 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
									v138 = int32(0)
									*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
								}
							} else {
								v102 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
								v114 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
								v126 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
								v138 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
								*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
							}
						} else {
							v90 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v90)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = int32(-1)
							v102 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
							v114 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
							v126 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
							v138 = int32(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
						}
					} else {
						v78 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v78)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
						v90 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v90)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = int32(-1)
						v102 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
						v114 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
						v126 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
						v138 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
					}
				} else {
					v66 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v66)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+22)) = int32(-1)
					v78 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v78)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
					v90 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v90)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = int32(-1)
					v102 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
					v114 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
					v126 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
					v138 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
				}
			} else {
				v54 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v54)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
				v66 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v66)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+22)) = int32(-1)
				v78 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v78)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
				v90 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v90)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = int32(-1)
				v102 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
				v114 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
				v126 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
				v138 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
			}
		} else {
			v42 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v42)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = int32(-1)
			v54 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v54)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(-1)
			v66 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v66)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+22)) = int32(-1)
			v78 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v78)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = int32(-1)
			v90 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v90)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = int32(-1)
			v102 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v102)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(-1)
			v114 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v114)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = int32(-1)
			v126 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v126)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = int32(-1)
			v138 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v138)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = int32(-1)
		}
	}
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v149 == int32(1) {
		v152 = int32(6)
		v154 = int32(18)
		v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
		if v156 == v154 {
			v159 = v154
		} else {
			v159 = int32(2)
		}
		if v156&int32(254) == int32(2) {
			v164 = v152
		} else {
			v164 = v159
		}
		if v156 == int32(1) {
			v167 = v152
		} else {
			v167 = v164
		}
		v176 = v167
	} else {
		v168 = int32(1)
		if v149&v168 != 0 {
			v176 = int32(base.Ui32(v149) >> (uint(v168) % 32))
		} else {
			v172 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v176 = int32(base.Ui32(v172) >> (uint(int32(2)) % 32))
		}
	}
	if v176 != 0 {
		v177 = F__emscripten_memcpy_bulkmem(m, l1+int32(72), v16, v176)
		mBase = m.M
	} else {
	}
	return
}
func F_HnswSetNeighborTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v121 int32
	_ = v121
	var v135 int32
	_ = v135
	v5 = int32(0)
	v17 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v17)
	v22 = l0 - int32(1)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	v31 = v23
	v33 = v5
	for {
		if l0 != 0 {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v22+v42+v31<<(uint(int32(2))%32))))
			if v47 != 0 {
				v50 = v22 + v47
			} else {
				v50 = int32(0)
			}
			v57 = v50
		} else {
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v31<<(uint(int32(2))%32))))
			v57 = v55
		}
		if base.B2i32(l3 <= v5) == int32(0) {
			v62 = int32(0)
			v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			v74 = v33
			v75 = v62
			for {
				v85 = l1 + int32(4) + v74*int32(6)
				if v75 < v65 {
					v89 = v57 + int32(8) + v75*int32(12)
					if l0 == int32(0) {
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
						v95 = v92
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
						v95 = v22 + v93
					}
					v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+80)))
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
					v99 = int32(base.Ui32(v97) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v85))) = uint16(v99)
					v105 = v96
					v106 = v97
				} else {
					v101 = int32(65535)
					*(*uint16)(unsafe.Add(mBase, uint32(v85))) = uint16(v101)
					v105 = int32(0)
					v106 = v101
				}
				v107 = int32(1)
				v108 = v74 + v107
				*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)) = uint16(v105)
				*(*uint16)(unsafe.Add(mBase, uint32(v85)+2)) = uint16(v106)
				v112 = v75 + v107
				if v112 != l3<<(uint(base.B2i32(v31 == v62))%32) {
					v74 = v108
					v75 = v112
					continue
				} else {
					break
				}
				break
			}
			v121 = v108
		} else {
			v121 = v33
		}
		if int32(0) < v31 {
			v31 = v31 - int32(1)
			v33 = v121
			continue
		} else {
			break
		}
		break
	}
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)) = uint16(v121)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v135)
	return
}
func F_HnswUpdateNeighborsOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v103 float32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v245 int32
	_ = v245
	var v246 float64
	_ = v246
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v285 int32
	_ = v285
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v443 int32
	_ = v443
	var v461 int32
	_ = v461
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v505 int32
	_ = v505
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v565 int32
	_ = v565
	var v591 int32
	_ = v591
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v649 int32
	_ = v649
	v25 = m.G0
	v27 = v25 - int32(1216)
	m.G0 = v27
	v30 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v32 = int32(131072)
	v35 = F_GenerationContextCreate(m, v30, int32(67507), v32, v32, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	v38 = l0
	v39 = l1
	v40 = l2
	v41 = l3
	v42 = l4
	v43 = l5
	v46 = v27
	v51 = v37
	v55 = v35
	goto L3
L3:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v40)+72))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v62+v51<<(uint(int32(2))%32))))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if int32(0) < v67 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_MemoryContextDelete(m, v637)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L85
	}
L5:
	;
	v70 = int32(0)
	v72 = v41 << (uint(base.B2i32(v51 == v70)) % 32)
	v94 = v70
	goto L8
L6:
	;
	v620 = v38
	v621 = v39
	v622 = v40
	v623 = v41
	v624 = v42
	v625 = v43
	v628 = v46
	v637 = v55
	goto L7
L7:
	;
	if int32(0) < v51 {
		v38 = v620
		v39 = v621
		v40 = v622
		v41 = v623
		v42 = v624
		v43 = v625
		v46 = v628
		v51 = v51 - int32(1)
		v55 = v637
		goto L3
	} else {
		goto L84
	}
L8:
	;
	v102 = v66 + int32(8) + v94*int32(12)
	v103 = *(*float32)(unsafe.Add(mBase, uint32(v102)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v105 = int32(4554240)
	v106 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = int32(-1)
	v112 = F_HnswInitNeighborArray(m, v72, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v620 = v38
	v621 = v39
	v622 = v40
	v623 = v41
	v624 = v42
	v625 = v43
	v628 = v46
	v637 = v55
	goto L7
L10:
	;
	v116 = F_HnswLoadNeighborTids(m, v104, v46+int32(16), v38, v41, v72, v51)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v119 = base.B2i32(v41 <= int32(0))
	if v41 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v199 < v72 {
		goto L21
	} else {
		goto L22
	}
L13:
	;
	if v116 == int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v131 = int32(0)
	goto L15
L15:
	;
	v153 = v46 + int32(16) + v131*int32(6)
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+4)))
	if v154 == int32(0) {
		goto L12
	} else {
		goto L17
	}
L16:
	;
	goto L12
L17:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+2)))
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153))))
	v162 = F_HnswInitElementFromBlock(m, v157|v158<<(uint(int32(16))%32), v154)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v165 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = v164 + v165
	*(*int32)(unsafe.Add(mBase, uint32(v112+int32(8)+v164*int32(12)))) = v162
	v173 = v131 + v165
	if v173 != v72 {
		v131 = v173
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v106
	F_MemoryContextReset(m, v55)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L35
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = int32(-2)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v104)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+8)) = v203
	if int32(0) < v199 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v216 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	F_HnswUpdateConnection(m, int32(0), v112, v40, v103, v72, v46+int32(12), v38, v39)
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L34
	}
L27:
	;
	v236 = v112 + int32(8) + v216*int32(12)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)))
	F_HnswLoadElement(m, v237, v46+int32(16), v46+int32(8), v38, v39, int32(1), int32(0))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L1
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v46)+16))
	*(*float32)(unsafe.Add(mBase, uint32(v236)+4)) = base.F32_demote_f64(v246)
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237)+64)))
	if v249 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v216
	goto L20
L31:
	;
	goto L32
L32:
	;
	v254 = v216 + int32(1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	if v254 < v255 {
		v216 = v254
		goto L27
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	goto L20
L35:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v314 != int32(-1) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v104)+82)))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v104)+84))
	v319 = F_ReadBuffer(m, v38, v318)
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v617 = v94 + int32(1)
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v617 < v618 {
		v94 = v617
		goto L8
	} else {
		goto L83
	}
L39:
	;
	F_LockBuffer(m, v319, int32(2))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if v43 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+65)))
	v352 = (v350 - v51) * v41
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v317<<(uint(int32(2))%32)+v349)+20))
	v359 = v349 + v356&int32(32767)
	if v42 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L42:
	;
	if v319 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	v343 = F_GenericXLogStart(m, v38)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v328+(v319^int32(-1))<<(uint(int32(2))%32))))
	v348 = int32(0)
	v349 = v334
	goto L41
L46:
	;
	goto L47
L47:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v348 = int32(0)
	v349 = v337 + v319<<(uint(int32(13))%32) + int32(-8192)
	goto L41
L48:
	;
	v346 = F_GenericXLogRegisterBuffer(m, v343, v319, int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v348 = v343
	v349 = v346
	goto L41
L50:
	;
	F_UnlockReleaseBuffer(m, v319)
	mBase = m.M
	v591 = m.ExcPending
	if v591 != 0 {
		goto L1
	} else {
		goto L82
	}
L51:
	;
	if v43 != 0 {
		goto L50
	} else {
		goto L80
	}
L52:
	;
	if v505 < int32(0) {
		goto L51
	} else {
		goto L73
	}
L53:
	;
	v505 = v352 + v314
	goto L52
L54:
	;
	if v314 == int32(-2) {
		goto L51
	} else {
		goto L72
	}
L55:
	;
	if v314 != int32(-2) {
		goto L53
	} else {
		goto L66
	}
L56:
	;
	if v41 <= int32(0) {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v371 = int32(0)
	goto L58
L58:
	;
	v392 = v359 + int32(4) + (v371+v352)*int32(6)
	v393 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+4)))
	if v393 == int32(0) {
		goto L55
	} else {
		goto L60
	}
L59:
	;
	goto L55
L60:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v40)+76))
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392)+2)))
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v392))))
	if v396 == v397|v398<<(uint(int32(16))%32) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+80)))
	if v393 == v403 {
		goto L51
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v406 = v371 + int32(1)
	if v406 != v72 {
		v371 = v406
		goto L58
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	goto L59
L66:
	;
	if v41 <= int32(0) {
		goto L51
	} else {
		goto L67
	}
L67:
	;
	v443 = int32(0)
	goto L68
L68:
	;
	v461 = v443 + v352
	v465 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v359+int32(8)+v461*int32(6)))))
	if v465 == int32(0) {
		v505 = v461
		goto L52
	} else {
		goto L70
	}
L69:
	;
	goto L51
L70:
	;
	v469 = v443 + int32(1)
	if v72 != v469 {
		v443 = v469
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L53
L73:
	;
	v524 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v359)+2)))
	if base.Ui32(v524) <= base.Ui32(v505) {
		goto L51
	} else {
		goto L74
	}
L74:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v40)+76))
	v529 = v359 + v505*int32(6)
	v530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v529)+8)) = uint16(v530)
	*(*uint16)(unsafe.Add(mBase, uint32(v529)+6)) = uint16(v526)
	v534 = int32(base.Ui32(v526) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v529)+4)) = uint16(v534)
	if v43 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_MarkBufferDirty(m, v319)
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_GenericXLogFinish(m, v348)
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L50
L79:
	;
	goto L50
L80:
	;
	F_pfree(m, v348)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L50
L82:
	;
	goto L38
L83:
	;
	goto L9
L84:
	;
	goto L4
L85:
	;
	m.G0 = v628 + int32(1216)
	return
}
func F_hnsw_halfvec_support(m *base.Module, l0 int32) int32 {
	return int32(4116400)
}
