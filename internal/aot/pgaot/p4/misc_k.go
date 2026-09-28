package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_KeepLogSeg(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v50 int64
	_ = v50
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v58 int32
	_ = v58
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	v9 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0])))
	v10 = base.I64_div_u_s(l0, v9)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
	v15 = base.AtomicRmwXchg32(m, v12, int32(440), int32(1))
	if v15 != 0 {
		F_s_lock(m, v12+int32(440), int32(_a_F_KeepLogSeg_0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+216))
			v24 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+440)), uint32(v24))
			if base.B2i32(v23 == int64(0))|base.B2i32(base.Ui64(l0) <= base.Ui64(v23)) != 0 {
				v52 = v10
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
				v34 = base.I64_div_u_s(v23, base.I64_extend_i32_s(v32))
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[2]))
				if v36 < int32(0) {
					v52 = v34
				} else {
					v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_KeepLogSeg[3])))
					if v40&int32(1) != 0 {
						v52 = v34
					} else {
						v44 = base.I32_div_s(v32, int32(_a_F_KeepLogSeg_1))
						v45 = base.I32_div_s(v36, v44)
						v46 = base.I64_extend_i32_s(v45)
						if base.Ui64(v46) < base.Ui64(v10-v34) {
							v50 = v10 - v46
						} else {
							v50 = v34
						}
						v52 = v50
					}
				}
			}
			v55 = int32(0)
			v57 = F_GetOldestUnsummarizedLSN(m, v55, v55)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				if v57 != int64(0) {
					v62 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0])))
					v63 = base.I64_div_u_s(v57, v62)
					if base.Ui64(v63) < base.Ui64(v52) {
						v65 = v63
					} else {
						v65 = v52
					}
					v67 = v65
				} else {
					v67 = v52
				}
				v69 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[4]))
				if v69 <= int32(0) {
					v85 = v67
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
					v75 = base.I32_div_s(v73, int32(_a_F_KeepLogSeg_1))
					v76 = base.I32_div_s(v69, v75)
					v77 = base.I64_extend_i32_s(v76)
					if base.Ui64(v77) <= base.Ui64(v10-v67) {
						v85 = v67
					} else {
						if base.Ui64(v10) <= base.Ui64(v77) {
							v83 = int64(1)
						} else {
							v83 = v10 - v77
						}
						v85 = v83
					}
				}
				v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui64(v85) < base.Ui64(v86) {
					*(*int64)(unsafe.Add(mBase, uint32(l1))) = v85
				} else {
				}
				return
			}
		}
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)+216))
		v24 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v22)+440)), uint32(v24))
		if base.B2i32(v23 == int64(0))|base.B2i32(base.Ui64(l0) <= base.Ui64(v23)) != 0 {
			v52 = v10
		} else {
			v32 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
			v34 = base.I64_div_u_s(v23, base.I64_extend_i32_s(v32))
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[2]))
			if v36 < int32(0) {
				v52 = v34
			} else {
				v40 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_KeepLogSeg[3])))
				if v40&int32(1) != 0 {
					v52 = v34
				} else {
					v44 = base.I32_div_s(v32, int32(_a_F_KeepLogSeg_1))
					v45 = base.I32_div_s(v36, v44)
					v46 = base.I64_extend_i32_s(v45)
					if base.Ui64(v46) < base.Ui64(v10-v34) {
						v50 = v10 - v46
					} else {
						v50 = v34
					}
					v52 = v50
				}
			}
		}
		v55 = int32(0)
		v57 = F_GetOldestUnsummarizedLSN(m, v55, v55)
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return
		} else {
			if v57 != int64(0) {
				v62 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0])))
				v63 = base.I64_div_u_s(v57, v62)
				if base.Ui64(v63) < base.Ui64(v52) {
					v65 = v63
				} else {
					v65 = v52
				}
				v67 = v65
			} else {
				v67 = v52
			}
			v69 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[4]))
			if v69 <= int32(0) {
				v85 = v67
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
				v75 = base.I32_div_s(v73, int32(_a_F_KeepLogSeg_1))
				v76 = base.I32_div_s(v69, v75)
				v77 = base.I64_extend_i32_s(v76)
				if base.Ui64(v77) <= base.Ui64(v10-v67) {
					v85 = v67
				} else {
					if base.Ui64(v10) <= base.Ui64(v77) {
						v83 = int64(1)
					} else {
						v83 = v10 - v77
					}
					v85 = v83
				}
			}
			v86 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui64(v85) < base.Ui64(v86) {
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v85
			} else {
			}
			return
		}
	}
}
