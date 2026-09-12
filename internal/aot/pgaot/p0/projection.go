package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecAssignScanProjectionInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	F_ExecConditionalAssignProjectionInfo(m, l0, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		return
	}
}
func F_apply_projection_to_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
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
	v7 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	switch v8 - int32(332) {
	case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
		v23 = v7
		v25 = v23
	case 2:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v16 != int32(290) {
			v23 = v7
			v25 = v23
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v25 = base.B2i32(v19 == int32(0))
		}
	default:
		v23 = int32(1)
		v25 = v23
	case 23:
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+72)))
		v25 = int32(base.Ui32(v11&int32(4)) >> (uint(int32(2)) % 32))
	}
	if v25 == int32(0) {
		v28 = F_create_projection_path(m, l0, l1, l2, l3)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			return v28
		}
	} else {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)+24))
		v35 = *(*float64)(unsafe.Add(mBase, uint32(v33)+16))
		*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
		v37 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
		v38 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l2)+48)) = base.F64_add(v37, base.F64_sub(v38, v35))
		v42 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
		v43 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
		v45 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
		v47 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
		*(*float64)(unsafe.Add(mBase, uint32(l2)+56)) = base.F64_add(v42, base.F64_add(base.F64_mul(base.F64_sub(v43, v34), v45), base.F64_sub(v47, v35)))
		v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v52&int32(-2) != int32(296) {
			v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
			if v68 != int32(1) {
				return l2
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v72 = F_is_parallel_safe(m, l0, v71)
				mBase = m.M
				v73 = m.ExcPending
				if v73 != 0 {
					return int32(0)
				} else {
					if v72 != 0 {
					} else {
						v74 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)) = uint8(v74)
					}
					return l2
				}
			}
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
			v58 = F_is_parallel_safe(m, l0, v57)
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				if v58 == int32(0) {
					v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
					if v68 != int32(1) {
						return l2
					} else {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v72 = F_is_parallel_safe(m, l0, v71)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							if v72 != 0 {
							} else {
								v74 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)) = uint8(v74)
							}
							return l2
						}
					}
				} else {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
					v64 = F_create_projection_path(m, l0, v63, v62, l3)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l2)+72)) = v64
						return l2
					}
				}
			}
		}
	}
}
func F_create_projection_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
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
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v83 float64
	_ = v83
	var v84 float64
	_ = v84
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v89 float64
	_ = v89
	var v90 float64
	_ = v90
	var v96 int32
	_ = v96
	var v98 float64
	_ = v98
	var v100 int32
	_ = v100
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v107 float64
	_ = v107
	var v108 float64
	_ = v108
	var v110 float64
	_ = v110
	var v112 float64
	_ = v112
	var v113 float64
	_ = v113
	v5 = int32(0)
	v8 = F_palloc0(m, int32(80))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(301)
		*(*int32)(unsafe.Add(mBase, uint32(v8))) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
		if v14 == v12 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+72))
			v18 = v17
		} else {
			v18 = l2
		}
		v19 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+20)) = uint8(v19)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l3
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(331)
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+26)))
		if v27 != int32(1) {
			v36 = v5
			*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)) = uint8(v36)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v38
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
			*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v40
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
			v44 = int32(0)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			switch v45 - int32(332) {
			case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
				v60 = v44
				v62 = v60
			case 2:
				v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				if v53 != int32(290) {
					v60 = v44
					v62 = v60
				} else {
					v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
					v62 = base.B2i32(v56 == int32(0))
				}
			default:
				v60 = int32(1)
				v62 = v60
			case 23:
				v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+72)))
				v62 = int32(base.Ui32(v48&int32(4)) >> (uint(int32(2)) % 32))
			}
			if v62 == int32(0) {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v67 = F_equal(m, v65, v66)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return int32(0)
				} else {
					if v67 == int32(0) {
						v96 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v96)
						v98 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v98
						v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v100
						v102 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
						v103 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v102, v103)
						v107 = *(*float64)(unsafe.Add(mBase, _consts[378]))
						v108 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v110 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
						v112 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
						v113 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_add(v107, v108), v110), base.F64_add(v112, v113))
						return v8
					} else {
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v71)
						v73 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v73
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
						v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v77, base.F64_sub(v78, v79))
						v83 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v84 = *(*float64)(unsafe.Add(mBase, uint32(v43)+24))
						v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
						v88 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
						v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v90 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v83, v84), v86), base.F64_add(v88, base.F64_sub(v89, v90)))
						return v8
					}
				}
			} else {
				v71 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v71)
				v73 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
				*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v73
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
				v77 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
				v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v77, base.F64_sub(v78, v79))
				v83 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
				v84 = *(*float64)(unsafe.Add(mBase, uint32(v43)+24))
				v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
				v88 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
				v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
				v90 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
				*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v83, v84), v86), base.F64_add(v88, base.F64_sub(v89, v90)))
				return v8
			}
		} else {
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+21)))
			if v30 != int32(1) {
				v36 = v5
				*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)) = uint8(v36)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
				*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v40
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
				v44 = int32(0)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				switch v45 - int32(332) {
				case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
					v60 = v44
					v62 = v60
				case 2:
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					if v53 != int32(290) {
						v60 = v44
						v62 = v60
					} else {
						v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
						v62 = base.B2i32(v56 == int32(0))
					}
				default:
					v60 = int32(1)
					v62 = v60
				case 23:
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+72)))
					v62 = int32(base.Ui32(v48&int32(4)) >> (uint(int32(2)) % 32))
				}
				if v62 == int32(0) {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
					v67 = F_equal(m, v65, v66)
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return int32(0)
					} else {
						if v67 == int32(0) {
							v96 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v96)
							v98 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
							*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v98
							v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v100
							v102 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
							v103 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v102, v103)
							v107 = *(*float64)(unsafe.Add(mBase, _consts[378]))
							v108 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
							v110 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
							v112 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
							v113 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_add(v107, v108), v110), base.F64_add(v112, v113))
							return v8
						} else {
							v71 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v71)
							v73 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
							*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v73
							v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
							v77 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
							v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							v79 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v77, base.F64_sub(v78, v79))
							v83 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
							v84 = *(*float64)(unsafe.Add(mBase, uint32(v43)+24))
							v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
							v88 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
							v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
							v90 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
							*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v83, v84), v86), base.F64_add(v88, base.F64_sub(v89, v90)))
							return v8
						}
					}
				} else {
					v71 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v71)
					v73 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
					*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v73
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
					v77 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
					v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v77, base.F64_sub(v78, v79))
					v83 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
					v84 = *(*float64)(unsafe.Add(mBase, uint32(v43)+24))
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
					v88 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
					v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
					v90 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
					*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v83, v84), v86), base.F64_add(v88, base.F64_sub(v89, v90)))
					return v8
				}
			} else {
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
				v34 = F_is_parallel_safe(m, l0, v33)
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int32(0)
				} else {
					v36 = v34
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+21)) = uint8(v36)
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v38
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v18)+64))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v8)+64)) = v40
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
					v44 = int32(0)
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
					switch v45 - int32(332) {
					case 0, 1, 3, 4, 28, 29, 30, 31, 35, 38, 39, 40, 41:
						v60 = v44
						v62 = v60
					case 2:
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
						if v53 != int32(290) {
							v60 = v44
							v62 = v60
						} else {
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
							v62 = base.B2i32(v56 == int32(0))
						}
					default:
						v60 = int32(1)
						v62 = v60
					case 23:
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+72)))
						v62 = int32(base.Ui32(v48&int32(4)) >> (uint(int32(2)) % 32))
					}
					if v62 == int32(0) {
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
						v66 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
						v67 = F_equal(m, v65, v66)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return int32(0)
						} else {
							if v67 == int32(0) {
								v96 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v96)
								v98 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v98
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v100
								v102 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
								v103 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v102, v103)
								v107 = *(*float64)(unsafe.Add(mBase, _consts[378]))
								v108 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
								v110 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
								v112 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
								v113 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_add(v107, v108), v110), base.F64_add(v112, v113))
								return v8
							} else {
								v71 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v71)
								v73 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v73
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
								v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								v79 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v77, base.F64_sub(v78, v79))
								v83 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
								v84 = *(*float64)(unsafe.Add(mBase, uint32(v43)+24))
								v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
								v88 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
								v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
								v90 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
								*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v83, v84), v86), base.F64_add(v88, base.F64_sub(v89, v90)))
								return v8
							}
						}
					} else {
						v71 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v8)+76)) = uint8(v71)
						v73 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+32)) = v73
						v75 = *(*int32)(unsafe.Add(mBase, uint32(v18)+40))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v75
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
						v78 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+48)) = base.F64_add(v77, base.F64_sub(v78, v79))
						v83 = *(*float64)(unsafe.Add(mBase, uint32(l3)+24))
						v84 = *(*float64)(unsafe.Add(mBase, uint32(v43)+24))
						v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
						v88 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
						v89 = *(*float64)(unsafe.Add(mBase, uint32(l3)+16))
						v90 = *(*float64)(unsafe.Add(mBase, uint32(v43)+16))
						*(*float64)(unsafe.Add(mBase, uint32(v8)+56)) = base.F64_add(base.F64_mul(base.F64_sub(v83, v84), v86), base.F64_add(v88, base.F64_sub(v89, v90)))
						return v8
					}
				}
			}
		}
	}
}
