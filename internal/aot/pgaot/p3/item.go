package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_executeItemUnwrapTargetArray(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != int32(18) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v20
			F_errmsg_internal(m, int32(_a_F_executeItemUnwrapTargetArray_0), v9)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_executeItemUnwrapTargetArray_1), int32(1707), int32(_a_F_executeItemUnwrapTargetArray_2))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
		v31 = int32(1)
		v35 = F_executeAnyItem(m, l0, l1, v30, l3, v31, v31, v31, int32(0), l4)
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			m.G0 = v9 + int32(16)
			return v35
		}
	}
}
func F_executeNextItem(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v42 int64
	_ = v42
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v79 int32
	_ = v79
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l2 != 0 {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		if int32(0) < v12 {
			v23 = l2
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			v25 = F_executeItemOptUnwrapTarget(m, l0, v23, l3, l4, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v79 = v25
				m.G0 = v10 + int32(32)
				return v79
			}
		} else {
			v28 = int32(0)
			if l4 == v28 {
				v79 = v28
				m.G0 = v10 + int32(32)
				return v79
			} else {
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
				if v32 < v33 {
					v37 = v31 + v32<<(uint(int32(5))%32)
					v38 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v37)+40)) = v38
					v40 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
					*(*int64)(unsafe.Add(mBase, uint32(v37)+32)) = v40
					v42 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = v42
					v44 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
					*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v44
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v46 + int32(1)
					v79 = v28
					m.G0 = v10 + int32(32)
					return v79
				} else {
					v50 = int32(16)
					v52 = v33 << (uint(int32(1)) % 32)
					if v52 <= v50 {
						v55 = v50
					} else {
						v55 = v52
					}
					v60 = F_palloc(m, v55<<(uint(int32(5))%32)|int32(16))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(1)
						v67 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v67
						v69 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v69
						v71 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v71
						v73 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v60)+40)) = v73
						*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v60
						*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v60
						v79 = v28
						m.G0 = v10 + int32(32)
						return v79
					}
				}
			}
		}
	} else {
		v16 = v10 + int32(4)
		v17 = F_jspGetNext(m, l1, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 == int32(0) {
				v28 = int32(0)
				if l4 == v28 {
					v79 = v28
					m.G0 = v10 + int32(32)
					return v79
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
					if v32 < v33 {
						v37 = v31 + v32<<(uint(int32(5))%32)
						v38 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
						*(*int64)(unsafe.Add(mBase, uint32(v37)+40)) = v38
						v40 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v37)+32)) = v40
						v42 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v37)+24)) = v42
						v44 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
						*(*int64)(unsafe.Add(mBase, uint32(v37)+16)) = v44
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
						*(*int32)(unsafe.Add(mBase, uint32(v31))) = v46 + int32(1)
						v79 = v28
						m.G0 = v10 + int32(32)
						return v79
					} else {
						v50 = int32(16)
						v52 = v33 << (uint(int32(1)) % 32)
						if v52 <= v50 {
							v55 = v50
						} else {
							v55 = v52
						}
						v60 = F_palloc(m, v55<<(uint(int32(5))%32)|int32(16))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v60)+4)) = v55
							*(*int32)(unsafe.Add(mBase, uint32(v60))) = int32(1)
							v67 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+16)) = v67
							v69 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+24)) = v69
							v71 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+32)) = v71
							v73 = *(*int64)(unsafe.Add(mBase, uint32(l3)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v60)+40)) = v73
							*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v60
							*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v60
							v79 = v28
							m.G0 = v10 + int32(32)
							return v79
						}
					}
				}
			} else {
				v23 = v16
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				v25 = F_executeItemOptUnwrapTarget(m, l0, v23, l3, l4, v24)
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v79 = v25
					m.G0 = v10 + int32(32)
					return v79
				}
			}
		}
	}
}
