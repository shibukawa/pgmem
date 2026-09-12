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
	var v34 int32
	_ = v34
	var v36 int64
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
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
	v9 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+440))
	v12 = int64(*(*int32)(unsafe.Add(mBase, _consts[262])))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+440)) = int32(1)
	v15 = base.I64_div_u_s(l0, v12)
	if v10 != 0 {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[190]))
		F_s_lock(m, v17+int32(440), int32(489567), int32(2685), int32(519111))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return
		} else {
			v26 = *(*int32)(unsafe.Add(mBase, _consts[190]))
			*(*int32)(unsafe.Add(mBase, uint32(v26)+440)) = int32(0)
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)+224))
			if v29 == int64(0) {
				v52 = v15
			} else {
				if base.Ui64(l0) <= base.Ui64(v29) {
					v52 = v15
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, _consts[262]))
					v36 = base.I64_div_u_s(v29, base.I64_extend_i32_s(v34))
					v38 = *(*int32)(unsafe.Add(mBase, _consts[287]))
					if v38 < int32(0) {
						v52 = v36
					} else {
						v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[288])))
						if v42 != 0 {
							v52 = v36
						} else {
							v44 = base.I32_div_s(v34, int32(1048576))
							v45 = base.I32_div_s(v38, v44)
							v46 = base.I64_extend_i32_s(v45)
							if base.Ui64(v46) < base.Ui64(v15-v36) {
								v50 = v15 - v46
							} else {
								v50 = v36
							}
							v52 = v50
						}
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
					v62 = int64(*(*int32)(unsafe.Add(mBase, _consts[262])))
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
				v69 = *(*int32)(unsafe.Add(mBase, _consts[289]))
				if v69 <= int32(0) {
					v85 = v67
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, _consts[262]))
					v75 = base.I32_div_s(v73, int32(1048576))
					v76 = base.I32_div_s(v69, v75)
					v77 = base.I64_extend_i32_s(v76)
					if base.Ui64(v77) <= base.Ui64(v15-v67) {
						v85 = v67
					} else {
						if base.Ui64(v15) <= base.Ui64(v77) {
							v83 = int64(1)
						} else {
							v83 = v15 - v77
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
		v26 = *(*int32)(unsafe.Add(mBase, _consts[190]))
		*(*int32)(unsafe.Add(mBase, uint32(v26)+440)) = int32(0)
		v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)+224))
		if v29 == int64(0) {
			v52 = v15
		} else {
			if base.Ui64(l0) <= base.Ui64(v29) {
				v52 = v15
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, _consts[262]))
				v36 = base.I64_div_u_s(v29, base.I64_extend_i32_s(v34))
				v38 = *(*int32)(unsafe.Add(mBase, _consts[287]))
				if v38 < int32(0) {
					v52 = v36
				} else {
					v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[288])))
					if v42 != 0 {
						v52 = v36
					} else {
						v44 = base.I32_div_s(v34, int32(1048576))
						v45 = base.I32_div_s(v38, v44)
						v46 = base.I64_extend_i32_s(v45)
						if base.Ui64(v46) < base.Ui64(v15-v36) {
							v50 = v15 - v46
						} else {
							v50 = v36
						}
						v52 = v50
					}
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
				v62 = int64(*(*int32)(unsafe.Add(mBase, _consts[262])))
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
			v69 = *(*int32)(unsafe.Add(mBase, _consts[289]))
			if v69 <= int32(0) {
				v85 = v67
			} else {
				v73 = *(*int32)(unsafe.Add(mBase, _consts[262]))
				v75 = base.I32_div_s(v73, int32(1048576))
				v76 = base.I32_div_s(v69, v75)
				v77 = base.I64_extend_i32_s(v76)
				if base.Ui64(v77) <= base.Ui64(v15-v67) {
					v85 = v67
				} else {
					if base.Ui64(v15) <= base.Ui64(v77) {
						v83 = int64(1)
					} else {
						v83 = v15 - v77
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
