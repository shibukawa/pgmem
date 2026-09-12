package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_RegisterTwoPhaseRecord(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	v1 = l0
	v10 = *(*int32)(unsafe.Add(mBase, _consts[167]))
	if base.Ui32(int32(8)) <= base.Ui32(v10) {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[168]))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
		v43 = v14
		v44 = v15
		v45 = v10
		v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
		v47 = v44 + v46
		v48 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)) = uint16(v48)
		*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)) = uint8(v1)
		*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2
		v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
		v53 = int32(8)
		v54 = v52 + v53
		*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v54
		v58 = v45 - v53
		*(*int32)(unsafe.Add(mBase, _consts[167])) = v58
		v60 = int32(4444060)
		v62 = *(*int32)(unsafe.Add(mBase, _consts[169]))
		v64 = v62 + v53
		*(*int32)(unsafe.Add(mBase, _consts[169])) = v64
		if l2 != 0 {
			v69 = (l2 + int32(7)) & int32(-8)
			if base.Ui32(v69) <= base.Ui32(v58) {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
				v104 = v64
				v105 = v43
				v106 = v58
				v107 = v71
				v108 = v54
				if l2 != 0 {
					v110 = F__emscripten_memcpy_bulkmem(m, v107+v108, l1, l2)
					mBase = m.M
				} else {
				}
				v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112 + v69
				*(*int32)(unsafe.Add(mBase, _consts[169])) = v104 + v69
				*(*int32)(unsafe.Add(mBase, _consts[167])) = v106 - v69
				return
			} else {
				v73 = F_palloc0(m, int32(12))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return
				} else {
					v75 = int32(4444048)
					v76 = *(*int32)(unsafe.Add(mBase, _consts[168]))
					*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v73
					*(*int32)(unsafe.Add(mBase, _consts[168])) = v73
					*(*int64)(unsafe.Add(mBase, uint32(v73)+4)) = int64(0)
					v83 = int32(512)
					if base.Ui32(v69) <= base.Ui32(v83) {
						v86 = v83
					} else {
						v86 = v69
					}
					*(*int32)(unsafe.Add(mBase, _consts[167])) = v86
					v88 = int32(4444052)
					v90 = *(*int32)(unsafe.Add(mBase, _consts[170]))
					*(*int32)(unsafe.Add(mBase, _consts[170])) = v90 + int32(1)
					v94 = F_palloc(m, v86)
					mBase = m.M
					v95 = m.ExcPending
					if v95 != 0 {
						return
					} else {
						v97 = *(*int32)(unsafe.Add(mBase, _consts[168]))
						*(*int32)(unsafe.Add(mBase, uint32(v97))) = v94
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
						v101 = *(*int32)(unsafe.Add(mBase, _consts[169]))
						v103 = *(*int32)(unsafe.Add(mBase, _consts[167]))
						v104 = v101
						v105 = v97
						v106 = v103
						v107 = v94
						v108 = v99
						if l2 != 0 {
							v110 = F__emscripten_memcpy_bulkmem(m, v107+v108, l1, l2)
							mBase = m.M
						} else {
						}
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112 + v69
						*(*int32)(unsafe.Add(mBase, _consts[169])) = v104 + v69
						*(*int32)(unsafe.Add(mBase, _consts[167])) = v106 - v69
						return
					}
				}
			}
		} else {
			return
		}
	} else {
		v17 = F_palloc0(m, int32(12))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return
		} else {
			v19 = int32(4444048)
			v20 = *(*int32)(unsafe.Add(mBase, _consts[168]))
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v17
			*(*int32)(unsafe.Add(mBase, _consts[168])) = v17
			*(*int64)(unsafe.Add(mBase, uint32(v17)+4)) = int64(0)
			v27 = int32(512)
			*(*int32)(unsafe.Add(mBase, _consts[167])) = v27
			v29 = int32(4444052)
			v31 = *(*int32)(unsafe.Add(mBase, _consts[170]))
			*(*int32)(unsafe.Add(mBase, _consts[170])) = v31 + int32(1)
			v36 = F_palloc(m, v27)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, _consts[168]))
				*(*int32)(unsafe.Add(mBase, uint32(v39))) = v36
				v42 = *(*int32)(unsafe.Add(mBase, _consts[167]))
				v43 = v39
				v44 = v36
				v45 = v42
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				v47 = v44 + v46
				v48 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)) = uint16(v48)
				*(*uint8)(unsafe.Add(mBase, uint32(v47)+4)) = uint8(v1)
				*(*int32)(unsafe.Add(mBase, uint32(v47))) = l2
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
				v53 = int32(8)
				v54 = v52 + v53
				*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v54
				v58 = v45 - v53
				*(*int32)(unsafe.Add(mBase, _consts[167])) = v58
				v60 = int32(4444060)
				v62 = *(*int32)(unsafe.Add(mBase, _consts[169]))
				v64 = v62 + v53
				*(*int32)(unsafe.Add(mBase, _consts[169])) = v64
				if l2 != 0 {
					v69 = (l2 + int32(7)) & int32(-8)
					if base.Ui32(v69) <= base.Ui32(v58) {
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
						v104 = v64
						v105 = v43
						v106 = v58
						v107 = v71
						v108 = v54
						if l2 != 0 {
							v110 = F__emscripten_memcpy_bulkmem(m, v107+v108, l1, l2)
							mBase = m.M
						} else {
						}
						v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112 + v69
						*(*int32)(unsafe.Add(mBase, _consts[169])) = v104 + v69
						*(*int32)(unsafe.Add(mBase, _consts[167])) = v106 - v69
						return
					} else {
						v73 = F_palloc0(m, int32(12))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return
						} else {
							v75 = int32(4444048)
							v76 = *(*int32)(unsafe.Add(mBase, _consts[168]))
							*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = v73
							*(*int32)(unsafe.Add(mBase, _consts[168])) = v73
							*(*int64)(unsafe.Add(mBase, uint32(v73)+4)) = int64(0)
							v83 = int32(512)
							if base.Ui32(v69) <= base.Ui32(v83) {
								v86 = v83
							} else {
								v86 = v69
							}
							*(*int32)(unsafe.Add(mBase, _consts[167])) = v86
							v88 = int32(4444052)
							v90 = *(*int32)(unsafe.Add(mBase, _consts[170]))
							*(*int32)(unsafe.Add(mBase, _consts[170])) = v90 + int32(1)
							v94 = F_palloc(m, v86)
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return
							} else {
								v97 = *(*int32)(unsafe.Add(mBase, _consts[168]))
								*(*int32)(unsafe.Add(mBase, uint32(v97))) = v94
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v97)+4))
								v101 = *(*int32)(unsafe.Add(mBase, _consts[169]))
								v103 = *(*int32)(unsafe.Add(mBase, _consts[167]))
								v104 = v101
								v105 = v97
								v106 = v103
								v107 = v94
								v108 = v99
								if l2 != 0 {
									v110 = F__emscripten_memcpy_bulkmem(m, v107+v108, l1, l2)
									mBase = m.M
								} else {
								}
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v112 + v69
								*(*int32)(unsafe.Add(mBase, _consts[169])) = v104 + v69
								*(*int32)(unsafe.Add(mBase, _consts[167])) = v106 - v69
								return
							}
						}
					}
				} else {
					return
				}
			}
		}
	}
}
func F_TwoPhaseGetDummyProc(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v3 = F_TwoPhaseGetGXact(m, l0, l1)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, _consts[152]))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
		return v9 + v10*int32(640)
	}
}
