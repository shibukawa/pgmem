package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_KeepLogSeg(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v53 int64
	_ = v53
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v65 int64
	_ = v65
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int64
	_ = v80
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+440))
	v12 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1])))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+440)) = int32(1)
	v15 = base.I64_div_u_s(l0, v12)
	if v10 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
		F_s_lock(m, v17+int32(440), int32(_a_F_KeepLogSeg_0), int32(2685), int32(_a_F_KeepLogSeg_1))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+440)) = int32(0)
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)+224))
			if base.B2i32(v29 == int64(0))|base.B2i32(base.Ui64(l0) <= base.Ui64(v29)) != 0 {
				v55 = v15
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
				v37 = base.I64_div_u_s(v29, base.I64_extend_i32_s(v35))
				v39 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[2]))
				if v39 < int32(0) {
					v55 = v37
				} else {
					v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_KeepLogSeg[3])))
					if v43&int32(1) != 0 {
						v55 = v37
					} else {
						v47 = base.I32_div_s(v35, int32(_a_F_KeepLogSeg_2))
						v48 = base.I32_div_s(v39, v47)
						v49 = base.I64_extend_i32_s(v48)
						if base.Ui64(v49) < base.Ui64(v15-v37) {
							v53 = v15 - v49
						} else {
							v53 = v37
						}
						v55 = v53
					}
				}
			}
			v58 = int32(0)
			v60 = F_GetOldestUnsummarizedLSN(m, v58, v58)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return
			} else {
				if v60 != int64(0) {
					v65 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1])))
					v66 = base.I64_div_u_s(v60, v65)
					if base.Ui64(v66) < base.Ui64(v55) {
						v68 = v66
					} else {
						v68 = v55
					}
					v70 = v68
				} else {
					v70 = v55
				}
				v72 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[4]))
				if v72 <= int32(0) {
					v88 = v70
				} else {
					v76 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
					v78 = base.I32_div_s(v76, int32(_a_F_KeepLogSeg_2))
					v79 = base.I32_div_s(v72, v78)
					v80 = base.I64_extend_i32_s(v79)
					if base.Ui64(v80) <= base.Ui64(v15-v70) {
						v88 = v70
					} else {
						if base.Ui64(v15) <= base.Ui64(v80) {
							v86 = int64(1)
						} else {
							v86 = v15 - v80
						}
						v88 = v86
					}
				}
				v89 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui64(v88) < base.Ui64(v89) {
					*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
				} else {
				}
				return
			}
		}
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v26)+440)) = int32(0)
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)+224))
		if base.B2i32(v29 == int64(0))|base.B2i32(base.Ui64(l0) <= base.Ui64(v29)) != 0 {
			v55 = v15
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
			v37 = base.I64_div_u_s(v29, base.I64_extend_i32_s(v35))
			v39 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[2]))
			if v39 < int32(0) {
				v55 = v37
			} else {
				v43 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_KeepLogSeg[3])))
				if v43&int32(1) != 0 {
					v55 = v37
				} else {
					v47 = base.I32_div_s(v35, int32(_a_F_KeepLogSeg_2))
					v48 = base.I32_div_s(v39, v47)
					v49 = base.I64_extend_i32_s(v48)
					if base.Ui64(v49) < base.Ui64(v15-v37) {
						v53 = v15 - v49
					} else {
						v53 = v37
					}
					v55 = v53
				}
			}
		}
		v58 = int32(0)
		v60 = F_GetOldestUnsummarizedLSN(m, v58, v58)
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return
		} else {
			if v60 != int64(0) {
				v65 = int64(*(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1])))
				v66 = base.I64_div_u_s(v60, v65)
				if base.Ui64(v66) < base.Ui64(v55) {
					v68 = v66
				} else {
					v68 = v55
				}
				v70 = v68
			} else {
				v70 = v55
			}
			v72 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[4]))
			if v72 <= int32(0) {
				v88 = v70
			} else {
				v76 = *(*int32)(unsafe.Add(mBase, _c_F_KeepLogSeg[1]))
				v78 = base.I32_div_s(v76, int32(_a_F_KeepLogSeg_2))
				v79 = base.I32_div_s(v72, v78)
				v80 = base.I64_extend_i32_s(v79)
				if base.Ui64(v80) <= base.Ui64(v15-v70) {
					v88 = v70
				} else {
					if base.Ui64(v15) <= base.Ui64(v80) {
						v86 = int64(1)
					} else {
						v86 = v15 - v80
					}
					v88 = v86
				}
			}
			v89 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
			if base.Ui64(v88) < base.Ui64(v89) {
				*(*int64)(unsafe.Add(mBase, uint32(l1))) = v88
			} else {
			}
			return
		}
	}
}
