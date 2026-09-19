package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_extract_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3)+8)))
	if int32(-64) <= v4 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
		return int32(base.Ui32(v11) >> (uint(int32(4)) % 32))
	}
}
func F_uuid_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return int32(base.Ui32(v162^int32(-1)) >> (uint(int32(31)) % 32))
}
func F_uuid_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int64
	_ = v6
	var v7 int64
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v41 int64
	_ = v41
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v84 int64
	_ = v84
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v152 int64
	_ = v152
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int64)(unsafe.Add(mBase, uint32(v5)))
	v7 = int64(56)
	v9 = int64(65280)
	v11 = int64(40)
	v14 = int64(16711680)
	v16 = int64(24)
	v18 = int64(4278190080)
	v20 = int64(8)
	v41 = v6<<(uint(v7)%64) | v6&v9<<(uint(v11)%64) | (v6&v14<<(uint(v16)%64) | v6&v18<<(uint(v20)%64)) | (int64(base.Ui64(v6)>>(uint(v20)%64))&v18 | int64(base.Ui64(v6)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v6)>>(uint(v11)%64))&v9 | int64(base.Ui64(v6)>>(uint(v7)%64))))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v42)))
	v78 = v43<<(uint(v7)%64) | v43&v9<<(uint(v11)%64) | (v43&v14<<(uint(v16)%64) | v43&v18<<(uint(v20)%64)) | (int64(base.Ui64(v43)>>(uint(v20)%64))&v18 | int64(base.Ui64(v43)>>(uint(v16)%64))&v14 | (int64(base.Ui64(v43)>>(uint(v11)%64))&v9 | int64(base.Ui64(v43)>>(uint(v7)%64))))
	if v41 == v78 {
		v81 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v82 = int64(56)
		v84 = int64(65280)
		v86 = int64(40)
		v89 = int64(16711680)
		v91 = int64(24)
		v93 = int64(4278190080)
		v95 = int64(8)
		v116 = v81<<(uint(v82)%64) | v81&v84<<(uint(v86)%64) | (v81&v89<<(uint(v91)%64) | v81&v93<<(uint(v95)%64)) | (int64(base.Ui64(v81)>>(uint(v95)%64))&v93 | int64(base.Ui64(v81)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v81)>>(uint(v86)%64))&v84 | int64(base.Ui64(v81)>>(uint(v82)%64))))
		v117 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
		v152 = v117<<(uint(v82)%64) | v117&v84<<(uint(v86)%64) | (v117&v89<<(uint(v91)%64) | v117&v93<<(uint(v95)%64)) | (int64(base.Ui64(v117)>>(uint(v95)%64))&v93 | int64(base.Ui64(v117)>>(uint(v91)%64))&v89 | (int64(base.Ui64(v117)>>(uint(v86)%64))&v84 | int64(base.Ui64(v117)>>(uint(v82)%64))))
		if v116 == v152 {
			v162 = int32(0)
		} else {
			v154 = v152
			v155 = v116
			if base.Ui64(v155) < base.Ui64(v154) {
				v159 = int32(-1)
			} else {
				v159 = int32(1)
			}
			v162 = v159
		}
	} else {
		v154 = v78
		v155 = v41
		if base.Ui64(v155) < base.Ui64(v154) {
			v159 = int32(-1)
		} else {
			v159 = int32(1)
		}
		v162 = v159
	}
	return int32(base.Ui32(v162) >> (uint(int32(31)) % 32))
}
func F_uuid_ns_oid(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14018(m, l0, int32(_a_F_uuid_ns_oid_0), int32(_a_F_uuid_ns_oid_1), int32(_a_F_uuid_ns_oid_2), int32(_a_F_uuid_ns_oid_3), int32(_a_F_uuid_ns_oid_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_uuid_skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = F_palloc(m, int32(16))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v13
			v17 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v17
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(1530)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(1531)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v6
			return int32(0)
		}
	}
}
func F_uuid_unparse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
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
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v23
	v57 = F_snprintf(m, l1, int32(37), int32(_a_F_uuid_unparse_0), v21)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		return
	} else {
		m.G0 = v21 - int32(-64)
		return
	}
}
