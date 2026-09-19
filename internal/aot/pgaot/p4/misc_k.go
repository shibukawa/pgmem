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
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int64
	_ = v27
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int64
	_ = v50
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int64
	_ = v81
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	v9 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0])))
	v10 = base.I64_div_u_s(l0, v9)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
	v15 = base.AtomicRmwXchg32(m, v12, int32(440), int32(1))
	if v15 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
		F_s_lock(m, v17+int32(440), int32(_a_F_KeepLogSeg_0), int32(2685), int32(_a_F_KeepLogSeg_1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
			v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+224))
			v28 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+440)), uint32(v28))
			if base.B2i32(v27 == int64(0))|base.B2i32(base.Ui64(l0) <= base.Ui64(v27)) != 0 {
				v56 = v10
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
				v38 = base.I64_div_u_s(v27, base.I64_extend_i32_s(v36))
				v40 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[2]))
				if v40 < int32(0) {
					v56 = v38
				} else {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_KeepLogSeg[3])))
					if v44&int32(1) != 0 {
						v56 = v38
					} else {
						v48 = base.I32_div_s(v36, int32(_a_F_KeepLogSeg_2))
						v49 = base.I32_div_s(v40, v48)
						v50 = base.I64_extend_i32_s(v49)
						if base.Ui64(v50) < base.Ui64(v10-v38) {
							v54 = v10 - v50
						} else {
							v54 = v38
						}
						v56 = v54
					}
				}
			}
			v59 = int32(0)
			v61 = F_GetOldestUnsummarizedLSN(m, v59, v59)
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return
			} else {
				if v61 != int64(0) {
					v66 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0])))
					v67 = base.I64_div_u_s(v61, v66)
					if base.Ui64(v67) < base.Ui64(v56) {
						v69 = v67
					} else {
						v69 = v56
					}
					v71 = v69
				} else {
					v71 = v56
				}
				v73 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[4]))
				if v73 <= int32(0) {
					v89 = v71
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
					v79 = base.I32_div_s(v77, int32(_a_F_KeepLogSeg_2))
					v80 = base.I32_div_s(v73, v79)
					v81 = base.I64_extend_i32_s(v80)
					if base.Ui64(v81) <= base.Ui64(v10-v71) {
						v89 = v71
					} else {
						if base.Ui64(v10) <= base.Ui64(v81) {
							v87 = int64(1)
						} else {
							v87 = v10 - v81
						}
						v89 = v87
					}
				}
				v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui64(v89) < base.Ui64(v90) {
					*(*int64)(unsafe.Add(mBase, uint32(l1))) = v89
				} else {
				}
				return
			}
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
		v27 = *(*int64)(unsafe.Add(mBase, uint32(v26)+224))
		v28 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v26)+440)), uint32(v28))
		if base.B2i32(v27 == int64(0))|base.B2i32(base.Ui64(l0) <= base.Ui64(v27)) != 0 {
			v56 = v10
		} else {
			v36 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
			v38 = base.I64_div_u_s(v27, base.I64_extend_i32_s(v36))
			v40 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[2]))
			if v40 < int32(0) {
				v56 = v38
			} else {
				v44 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_KeepLogSeg[3])))
				if v44&int32(1) != 0 {
					v56 = v38
				} else {
					v48 = base.I32_div_s(v36, int32(_a_F_KeepLogSeg_2))
					v49 = base.I32_div_s(v40, v48)
					v50 = base.I64_extend_i32_s(v49)
					if base.Ui64(v50) < base.Ui64(v10-v38) {
						v54 = v10 - v50
					} else {
						v54 = v38
					}
					v56 = v54
				}
			}
		}
		v59 = int32(0)
		v61 = F_GetOldestUnsummarizedLSN(m, v59, v59)
		mBase = m.M
		v62 = m.ExcPending
		if v62 != 0 {
			return
		} else {
			if v61 != int64(0) {
				v66 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0])))
				v67 = base.I64_div_u_s(v61, v66)
				if base.Ui64(v67) < base.Ui64(v56) {
					v69 = v67
				} else {
					v69 = v56
				}
				v71 = v69
			} else {
				v71 = v56
			}
			v73 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[4]))
			if v73 <= int32(0) {
				v89 = v71
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
				v79 = base.I32_div_s(v77, int32(_a_F_KeepLogSeg_2))
				v80 = base.I32_div_s(v73, v79)
				v81 = base.I64_extend_i32_s(v80)
				if base.Ui64(v81) <= base.Ui64(v10-v71) {
					v89 = v71
				} else {
					if base.Ui64(v10) <= base.Ui64(v81) {
						v87 = int64(1)
					} else {
						v87 = v10 - v81
					}
					v89 = v87
				}
			}
			v90 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui64(v89) < base.Ui64(v90) {
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v89
			} else {
			}
			return
		}
	}
}
