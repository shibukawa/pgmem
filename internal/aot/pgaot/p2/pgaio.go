package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_io_get_op_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v3) <= base.Ui32(int32(2)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_consts[729])))
		v11 = v10
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_pgaio_io_get_state_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v3) <= base.Ui32(int32(7)) {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32))+uint32(_consts[728])))
		v11 = v10
	} else {
		v11 = int32(0)
	}
	return v11
}
func F_pgaio_submit_staged(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _consts[726]))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)))
	if v10 == int32(0) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v13 = int32(4474964)
		v15 = *(*int32)(unsafe.Add(mBase, _consts[7]))
		*(*int32)(unsafe.Add(mBase, _consts[7])) = v15 + int32(1)
		v22 = *(*int32)(unsafe.Add(mBase, _consts[727]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
		v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, v10, v9+int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = int32(4474964)
			v28 = *(*int32)(unsafe.Add(mBase, _consts[7]))
			*(*int32)(unsafe.Add(mBase, _consts[7])) = v28 - int32(1)
			v33 = *(*int32)(unsafe.Add(mBase, _consts[726]))
			v34 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v33)+22)) = uint16(v34)
			v38 = F_errstart(m, int32(11), v34)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					F_errhidestmt(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errhidecontext(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v24
							F_errmsg_internal(m, int32(173435), v6)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_errfinish(m, int32(491116), int32(1147), int32(455516))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgaio_worker_die(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	v4 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v8 = F_LWLockAcquire(m, v4+int32(6784), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, _consts[730]))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v15 = *(*int32)(unsafe.Add(mBase, _consts[731]))
		*(*int64)(unsafe.Add(mBase, uint32(v11))) = v12 & base.I64_rotl(int64(-2), base.I64_extend_i32_u(v15))
		v22 = v11 + v15<<(uint(int32(3))%32)
		v23 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v23
		*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)) = uint8(v23)
		v28 = *(*int32)(unsafe.Add(mBase, _consts[2]))
		F_LWLockRelease(m, v28+int32(6784))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			return
		}
	}
}
func F_pgaio_wref_valid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(v2 != int32(-1))
}
