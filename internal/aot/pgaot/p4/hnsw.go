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
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn14214(m, l0, int32(3), int32(_a_F_HnswGetTypeInfo_0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
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
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
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
			v20 = int32(_a_F_HnswInitElement_0)
			v23 = *(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[0]))
			v24 = *(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[1]))
			v25 = v23 ^ v24
			*(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[1])) = base.I64_rotl(v25, int64(37))
			*(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[0])) = v25<<(uint(int64(16))%64) ^ base.I64_rotl(v23, int64(24)) ^ v25
			v46 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v23*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
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
			v58 = base.I32_trunc_sat_f64_s(base.F64_mul(v55, base.F64_neg(l3)))
			if l4 < v58 {
				v60 = l4
			} else {
				v60 = v58
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+65)) = uint8(v60)
			F_HnswInitNeighbors(m, l0, v19, l2, l5)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
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
			v20 = int32(_a_F_HnswInitElement_0)
			v23 = *(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[0]))
			v24 = *(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[1]))
			v25 = v23 ^ v24
			*(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[1])) = base.I64_rotl(v25, int64(37))
			*(*int64)(unsafe.Add(mBase, _c_F_HnswInitElement[0])) = v25<<(uint(int64(16))%64) ^ base.I64_rotl(v23, int64(24)) ^ v25
			v46 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v23*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
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
			v58 = base.I32_trunc_sat_f64_s(base.F64_mul(v55, base.F64_neg(l3)))
			if l4 < v58 {
				v60 = l4
			} else {
				v60 = v58
			}
			*(*uint8)(unsafe.Add(mBase, uint32(v19)+65)) = uint8(v60)
			F_HnswInitNeighbors(m, l0, v19, l2, l5)
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
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
	var v64 int32
	_ = v64
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
						v21 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertAppendPage[0]))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v21+(v17^int32(-1))<<(uint(int32(2))%32))))
						v40 = v27
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertAppendPage[1]))
						v40 = v29 + v17<<(uint(int32(13))%32) + int32(-8192)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v40
					F_PageInit(m, v40, int32(_a_F_HnswInsertAppendPage_0), int32(8))
					mBase = m.M
					v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
					v46 = v40 + v45
					v47 = int32(_a_F_HnswInsertAppendPage_1)
					*(*uint16)(unsafe.Add(mBase, uint32(v46)+6)) = uint16(v47)
					*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(-1)
					v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					if v51 < int32(0) {
						v55 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertAppendPage[2]))
						v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v51^int32(-1))*int32(56))+16))
						v70 = v61
					} else {
						v63 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertAppendPage[3]))
						v64 = int32(56)
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v51*v64-v64)+16))
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
						F_PageInit(m, v40, int32(_a_F_HnswInsertAppendPage_0), int32(8))
						mBase = m.M
						v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v40)+16)))
						v46 = v40 + v45
						v47 = int32(_a_F_HnswInsertAppendPage_1)
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+6)) = uint16(v47)
						*(*int32)(unsafe.Add(mBase, uint32(v46))) = int32(-1)
						v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if v51 < int32(0) {
							v55 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertAppendPage[2]))
							v61 = *(*int32)(unsafe.Add(mBase, uint32(v55+(v51^int32(-1))*int32(56))+16))
							v70 = v61
						} else {
							v63 = *(*int32)(unsafe.Add(mBase, _c_F_HnswInsertAppendPage[3]))
							v64 = int32(56)
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v63+v51*v64-v64)+16))
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
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
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
		v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)))
		*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)) = uint16(v34)
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v33))) = v36
		v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
		if base.Ui32(int32(1)) < base.Ui32(v38) {
			v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+14)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+14)) = uint16(v46)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+10))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+10)) = v48
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
			if base.Ui32(int32(2)) < base.Ui32(v50) {
				v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+20)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+20)) = uint16(v58)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v60
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
				if base.Ui32(int32(3)) < base.Ui32(v62) {
					v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+26)))
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+26)) = uint16(v70)
					v72 = *(*int32)(unsafe.Add(mBase, uint32(l2)+22))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+22)) = v72
					v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
					if base.Ui32(int32(4)) < base.Ui32(v74) {
						v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+32)))
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v82)
						v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v84
						v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
						if base.Ui32(int32(5)) < base.Ui32(v86) {
							v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+38)))
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+38)) = uint16(v94)
							v96 = *(*int32)(unsafe.Add(mBase, uint32(l2)+34))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+34)) = v96
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
							if base.Ui32(int32(6)) < base.Ui32(v98) {
								v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+44)))
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+44)) = uint16(v106)
								v108 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v108
								v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
								if base.Ui32(int32(7)) < base.Ui32(v110) {
									v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+50)))
									*(*uint16)(unsafe.Add(mBase, uint32(l1)+50)) = uint16(v118)
									v120 = *(*int32)(unsafe.Add(mBase, uint32(l2)+46))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+46)) = v120
									v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
									if base.Ui32(int32(8)) < base.Ui32(v122) {
										v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+56)))
										*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)) = uint16(v130)
										v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+52))
										*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v132
										v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+64)))
										if base.Ui32(int32(9)) < base.Ui32(v134) {
											v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+62)))
											*(*uint16)(unsafe.Add(mBase, uint32(l1)+62)) = uint16(v142)
											v144 = *(*int32)(unsafe.Add(mBase, uint32(l2)+58))
											*(*int32)(unsafe.Add(mBase, uint32(l1)+58)) = v144
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
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v147 == int32(1) {
		v151 = int32(18)
		v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
		if v153 == v151 {
			v156 = v151
		} else {
			v156 = int32(2)
		}
		if base.Ui32((v153-int32(1))&int32(255)) < base.Ui32(int32(3)) {
			v163 = int32(6)
		} else {
			v163 = v156
		}
		v172 = v163
	} else {
		v164 = int32(1)
		if v147&v164 != 0 {
			v172 = int32(base.Ui32(v147) >> (uint(v164) % 32))
		} else {
			v168 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
			v172 = int32(base.Ui32(v168) >> (uint(int32(2)) % 32))
		}
	}
	if v172 != 0 {
		base.MemoryCopy(m, l1+int32(72), v16, v172)
	} else {
	}
	return
}
func F_HnswSetNeighborTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
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
	var v123 int32
	_ = v123
	var v134 int32
	_ = v134
	v5 = int32(0)
	v16 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v16)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	v28 = v20
	v32 = v5
	for {
		if l0 != 0 {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0+v38+v28<<(uint(int32(2))%32)-int32(1))))
			if v45 != 0 {
				v56 = l0 + v45 - int32(1)
			} else {
				v56 = int32(0)
			}
		} else {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v47+v28<<(uint(int32(2))%32))))
			v56 = v51
		}
		if base.B2i32(l3 <= v5) == int32(0) {
			v61 = int32(0)
			v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v70 = v61
			v75 = v32
			for {
				v83 = l1 + int32(4) + v75*int32(6)
				if v70 < v64 {
					v87 = v56 + int32(8) + v70*int32(12)
					if l0 == int32(0) {
						v90 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						v95 = v90
					} else {
						v91 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
						v95 = l0 + v91 - int32(1)
					}
					v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+80)))
					v97 = *(*int32)(unsafe.Add(mBase, uint32(v95)+76))
					v99 = int32(base.Ui32(v97) >> (uint(int32(16)) % 32))
					*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v99)
					v105 = v96
					v106 = v97
				} else {
					v101 = int32(_a_F_HnswSetNeighborTuple_0)
					*(*uint16)(unsafe.Add(mBase, uint32(v83))) = uint16(v101)
					v105 = int32(0)
					v106 = v101
				}
				v107 = int32(1)
				v108 = v75 + v107
				*(*uint16)(unsafe.Add(mBase, uint32(v83)+4)) = uint16(v105)
				*(*uint16)(unsafe.Add(mBase, uint32(v83)+2)) = uint16(v106)
				v112 = v70 + v107
				if v112 != l3<<(uint(base.B2i32(v28 == v61))%32) {
					v70 = v112
					v75 = v108
					continue
				} else {
					break
				}
				break
			}
			v123 = v108
		} else {
			v123 = v32
		}
		if int32(0) < v28 {
			v28 = v28 - int32(1)
			v32 = v123
			continue
		} else {
			break
		}
		break
	}
	*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)) = uint16(v123)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)) = uint8(v134)
	return
}
func F_HnswUpdateNeighborsOnDisk(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 float32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
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
	var v129 int32
	_ = v129
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v212 int32
	_ = v212
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 float64
	_ = v239
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v277 int32
	_ = v277
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v366 int32
	_ = v366
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v427 int32
	_ = v427
	var v436 int32
	_ = v436
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v496 int32
	_ = v496
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v554 int32
	_ = v554
	var v579 int32
	_ = v579
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v624 int32
	_ = v624
	var v635 int32
	_ = v635
	v24 = m.G0
	v26 = v24 - int32(1216)
	m.G0 = v26
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateNeighborsOnDisk[0]))
	v31 = int32(_a_F_HnswUpdateNeighborsOnDisk_0)
	v34 = F_GenerationContextCreate(m, v29, int32(_a_F_HnswUpdateNeighborsOnDisk_1), v31, v31, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+65)))
	v37 = l0
	v38 = l1
	v39 = l2
	v40 = l3
	v41 = l4
	v42 = l5
	v46 = v26
	v49 = v36
	v54 = v34
	goto L3
L3:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v60+v49<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if int32(0) < v65 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_MemoryContextDelete(m, v624)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L85
	}
L5:
	;
	v68 = int32(0)
	v70 = v40 << (uint(base.B2i32(v49 == v68)) % 32)
	v92 = v68
	goto L8
L6:
	;
	v607 = v37
	v608 = v38
	v609 = v39
	v610 = v40
	v611 = v41
	v612 = v42
	v616 = v46
	v624 = v54
	goto L7
L7:
	;
	if int32(0) < v49 {
		v37 = v607
		v38 = v608
		v39 = v609
		v40 = v610
		v41 = v611
		v42 = v612
		v46 = v616
		v49 = v49 - int32(1)
		v54 = v624
		goto L3
	} else {
		goto L84
	}
L8:
	;
	v99 = v64 + int32(8) + v92*int32(12)
	v100 = *(*float32)(unsafe.Add(mBase, uint32(v99)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v102 = int32(_a_F_HnswUpdateNeighborsOnDisk_2)
	v103 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateNeighborsOnDisk[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateNeighborsOnDisk[0])) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = int32(-1)
	v109 = F_HnswInitNeighborArray(m, v70, int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v607 = v37
	v608 = v38
	v609 = v39
	v610 = v40
	v611 = v41
	v612 = v42
	v616 = v46
	v624 = v54
	goto L7
L10:
	;
	v113 = F_HnswLoadNeighborTids(m, v101, v46+int32(16), v37, v40, v70, v49)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v115 = int32(0)
	v116 = base.B2i32(v40 <= v115)
	if v116|base.B2i32(v113 == v115) != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v195 < v70 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v129 = int32(0)
	goto L14
L14:
	;
	v150 = v46 + int32(16) + v129*int32(6)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+4)))
	if v151 == int32(0) {
		goto L12
	} else {
		goto L16
	}
L15:
	;
	goto L12
L16:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+2)))
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150))))
	v159 = F_HnswInitElementFromBlock(m, v154|v155<<(uint(int32(16))%32), v151)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	v162 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v109))) = v161 + v162
	*(*int32)(unsafe.Add(mBase, uint32(v109+int32(8)+v161*int32(12)))) = v159
	v170 = v129 + v162
	if v170 != v70 {
		v129 = v170
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateNeighborsOnDisk[0])) = v103
	F_MemoryContextReset(m, v54)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L1
	} else {
		goto L34
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = int32(-2)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v199 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v101)+88)))
	*(*int64)(unsafe.Add(mBase, uint32(v46))) = v199
	if int32(0) < v195 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v212 = int32(0)
	goto L26
L24:
	;
	goto L25
L25:
	;
	F_HnswUpdateConnection(m, int32(0), v109, v39, v100, v70, v46+int32(12), v37, v38)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L33
	}
L26:
	;
	v231 = v109 + int32(8) + v212*int32(12)
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	F_HnswLoadElement(m, v232, v46+int32(16), v46, v37, v38, int32(1), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L28
	}
L27:
	;
	goto L25
L28:
	;
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v46)+16))
	*(*float32)(unsafe.Add(mBase, uint32(v231)+4)) = base.F32_demote_f64(v239)
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+64)))
	if v242 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+12)) = v212
	goto L19
L30:
	;
	goto L31
L31:
	;
	v247 = v212 + int32(1)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
	if v247 < v248 {
		v212 = v247
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	goto L19
L34:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v305 != int32(-1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+82)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v101)+84))
	v310 = F_ReadBuffer(m, v37, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v604 = v92 + int32(1)
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v604 < v605 {
		v92 = v604
		goto L8
	} else {
		goto L83
	}
L38:
	;
	F_LockBufferInternal(m, v310, int32(3))
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if v42 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+65)))
	v343 = (v341 - v49) * v40
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v340+v308<<(uint(int32(2))%32))+20))
	v350 = v340 + v347&int32(_a_F_HnswUpdateNeighborsOnDisk_3)
	if v41 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L41:
	;
	if v310 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L43
L43:
	;
	v334 = F_GenericXLogStart(m, v37)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L47
	}
L44:
	;
	v319 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateNeighborsOnDisk[1]))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v319+(v310^int32(-1))<<(uint(int32(2))%32))))
	v339 = int32(0)
	v340 = v325
	goto L40
L45:
	;
	goto L46
L46:
	;
	v328 = *(*int32)(unsafe.Add(mBase, _c_F_HnswUpdateNeighborsOnDisk[2]))
	v339 = int32(0)
	v340 = v328 + v310<<(uint(int32(13))%32) + int32(-8192)
	goto L40
L47:
	;
	v337 = F_GenericXLogRegisterBuffer(m, v334, v310, int32(0))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v339 = v334
	v340 = v337
	goto L40
L49:
	;
	F_UnlockReleaseBuffer(m, v310)
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L1
	} else {
		goto L82
	}
L50:
	;
	if v42 != 0 {
		goto L49
	} else {
		goto L80
	}
L51:
	;
	if v496 < int32(0) {
		goto L50
	} else {
		goto L73
	}
L52:
	;
	v496 = v305 + v343
	goto L51
L53:
	;
	if v305 == int32(-2) {
		goto L50
	} else {
		goto L72
	}
L54:
	;
	if v305 != int32(-2) {
		goto L52
	} else {
		goto L66
	}
L55:
	;
	if v40 <= v115 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	v366 = int32(0)
	goto L57
L57:
	;
	v384 = v350 + v343*int32(6) + int32(4) + v366*int32(6)
	if v384 == int32(0) {
		goto L54
	} else {
		goto L59
	}
L58:
	;
	goto L54
L59:
	;
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+4)))
	if v387 == int32(0) {
		goto L54
	} else {
		goto L60
	}
L60:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v39)+76))
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384)+2)))
	v392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v384))))
	if v390 == v391|v392<<(uint(int32(16))%32) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v397 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+80)))
	if v387 == v397 {
		goto L50
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v400 = v366 + int32(1)
	if v400 != v70 {
		v366 = v400
		goto L57
	} else {
		goto L65
	}
L64:
	;
	goto L63
L65:
	;
	goto L58
L66:
	;
	v427 = int32(0)
	if v40 <= v427 {
		goto L50
	} else {
		goto L67
	}
L67:
	;
	v436 = v427
	goto L68
L68:
	;
	v453 = v436 + v343
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350+v453*int32(6))+8)))
	if v457 == int32(0) {
		v496 = v453
		goto L51
	} else {
		goto L70
	}
L69:
	;
	goto L50
L70:
	;
	v461 = v436 + int32(1)
	if v70 != v461 {
		v436 = v461
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	goto L52
L73:
	;
	v514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350)+2)))
	if base.Ui32(v514) <= base.Ui32(v496) {
		goto L50
	} else {
		goto L74
	}
L74:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v39)+76))
	v519 = v350 + v496*int32(6)
	v520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v519)+8)) = uint16(v520)
	*(*uint16)(unsafe.Add(mBase, uint32(v519)+6)) = uint16(v516)
	v524 = int32(base.Ui32(v516) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v519)+4)) = uint16(v524)
	if v42 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	F_MarkBufferDirty(m, v310)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	F_GenericXLogFinish(m, v339)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	goto L49
L79:
	;
	goto L49
L80:
	;
	F_pfree(m, v339)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	goto L49
L82:
	;
	goto L37
L83:
	;
	goto L9
L84:
	;
	goto L4
L85:
	;
	m.G0 = v616 + int32(1216)
	return
}
func F_hnsw_halfvec_support(m *base.Module, l0 int32) int64 {
	return int64(4171832)
}
